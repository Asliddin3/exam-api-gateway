package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"time"

	pbc "github.com/Asliddin3/exam-api-gateway/genproto/customer"
	"github.com/Asliddin3/exam-api-gateway/pkg/logger"

	"github.com/Asliddin3/exam-api-gateway/pkg/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

// Register godoc
// @Summary Register for authentication
// @Tags Auth
// @Accept json
// @Produce json
// @Param userData body models.Register true "user data"
// @Success 201 {object} user.User
// @Failure 400 {object} models.Error
// @Router /register [post]
func (h *handlerV1) Register(c *gin.Context) {

	newUser := &pbc.CustomerRequest{}
	c.ShouldBindJSON(newUser)

	email, err := utils.IsValidMail(newUser.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email address",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	exists, err := h.serviceManager.CustomerService().CheckField(ctx, &pbc.CheckFieldRequest{Key: "username", Value: newUser.UserName})

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

	exists, err = h.serviceManager.CustomerService().CheckField(ctx, &pbc.CheckFieldRequest{Key: "email", Value: newUser.Email})

	if err != nil {
		h.log.Error("error while checking email existance", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while cheking email existance",
		})
		return
	}

	if exists.Exists {
		c.JSON(http.StatusOK, gin.H{
			"message": "such email already exists",
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

	jsNewUser, err := json.Marshal(newUser)
	if err != nil {
		h.log.Error("error while marshaling new user, inorder to insert it to redis", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while creating user",
		})
		return
	}

	code := utils.RandomNum()

	_, err = h.redis.Get(fmt.Sprint(code))
	if err == nil {
		code = utils.RandomNum()
	}

	if err = h.redis.SetWithTTL(fmt.Sprint(code), string(jsNewUser), 86000); err != nil {
		fmt.Println(err)
		h.log.Error("error while inserting new user into redis")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong, please try again",
		})
		return
	}

	newUser.Email = email

	ctx, cancel = context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	res, err := EmailVerification("Verigication", fmt.Sprint(code), email)
	// res, err := h.serviceManager.EmailService().Send(ctx, &pbe.Email{Subject: "Verification.", Body: fmt.Sprint(code), Recipients: []*pbe.Recipient{{Email: email}}})
	if err != nil {
		h.log.Error("error while sending verification code to new user", logger.Error(err))
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
		"rahimzanovmuhammadumar@gmail.com",
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
