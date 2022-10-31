package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"time"

	"github.com/Asliddin3/exam-api-gateway/genproto/customer"
	"github.com/Asliddin3/exam-api-gateway/pkg/logger"

	"github.com/Asliddin3/exam-api-gateway/pkg/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

// @BasePath /api/v1
// Register godoc
// @Summary Register for authentication
// @Tags Auth
// @Accept json
// @Produce json
// @Param userData body models.Register true "user data"
// @Success 201 "success"
// @Router /register [post]
func (h *handlerV1) Register(c *gin.Context) {

	newUser := &customer.CustomerRequest{}
	err := c.ShouldBindJSON(newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error binding json to object",
		})
		return
	}
	err = utils.IsValidMail(newUser.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email address",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	exists, err := h.serviceManager.CustomerService().CheckField(ctx, &customer.CheckFieldRequest{Key: "username", Value: newUser.UserName})
	if err != nil {
		h.log.Error("error while checking username existance", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while cheking username existance",
		})
		return
	}

	if exists.Exists {
		c.JSON(http.StatusOK, gin.H{
			"message": "such username already exists",
		})
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	exists, err = h.serviceManager.CustomerService().CheckField(ctx, &customer.CheckFieldRequest{Key: "email", Value: newUser.Email})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while checking email",
		})
		return
	}
	if exists.Exists {
		c.JSON(http.StatusOK, gin.H{
			"message": "such email already exists",
		})
		return
	}

	val, err := h.redis.Get(newUser.Email)
	if val != nil {
		h.log.Error("error email already registered", logger.Error(err))
		c.JSON(http.StatusAlreadyReported, gin.H{
			"error": "error email already exists",
		})
		return
	}

	val, err = h.redis.Get(newUser.UserName)
	fmt.Println("if getting username", val, err)
	if val != nil {
		h.log.Error("error username already registered", logger.Error(err))
		c.JSON(http.StatusAlreadyReported, gin.H{
			"error": "error username already exists",
		})
		return
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(newUser.PassWord), 10)

	if err != nil {
		h.log.Error("error while hashing password", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}

	newUser.PassWord = string(hashPass)
	fmt.Println(newUser.PassWord)
	code := utils.RandomNum()

	_, err = h.redis.Get(fmt.Sprint(code))
	for err != nil {
		code := utils.RandomNum()
		_, err = h.redis.Get(fmt.Sprint(code))
	}
	userJSON, err := json.Marshal(newUser)
	if err != nil {
		h.log.Error("error whi.e marshiling customer")
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "error while marshiling customer",
		})
		return
	}
	if err = h.redis.SetWithTTL(fmt.Sprint(code), string(userJSON), 86000); err != nil {
		h.log.Error("error while inserting new username into redis")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong, please try again",
		})
		return
	}

	res, err := EmailVerification("Verefication", fmt.Sprint(code), newUser.Email)
	if err != nil {
		h.log.Error("error while sending verification code to new customer", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong, please try again",
		})
		return
	}

	c.JSON(http.StatusOK, res)

}

func EmailVerification(subject, code, email string) (string, error) {

	// Sender data.
	from := "asliddinvstalim@gmail.com"
	password := "gnradbxvloedrkti"

	// Receiver email address.
	to := []string{
		email,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte(fmt.Sprintf("%s %s", subject, code))

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return "Error with sending message", err
	}
	return "Message sended to your email succesfully", nil
}
