package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
	"time"

	"github.com/Asliddin3/exam-api-gateway/api/models"
	pbc "github.com/Asliddin3/exam-api-gateway/genproto/customer"
	"github.com/Asliddin3/exam-api-gateway/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cast"
)

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// @BasePath /api/v1
// Register godoc
// @Summary Register for authentication
// @Tags Auth
// @Accept json
// @Produce json
// @Param userData body models.CofirmEmail true "login"
// @Success 201 {object}  models.VerifiedResponse
// @Router /confirm [post]
func (h *handlerV1) GetVerification(c *gin.Context) {
	var (
		body models.CofirmEmail
	)
	err := c.ShouldBindJSON(&body)
	fmt.Println(err, body)
	if err != nil {
		h.log.Error("error binding json")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error getting verification",
		})
		return
	}
	customer, err := h.redis.Get(body.Password)
	if err != nil {
		h.log.Error("error getting customer")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error wrong code",
		})
		return
	}
	customerReq := pbc.CustomerRequest{}
	err = json.Unmarshal([]byte(cast.ToString(customer)), &customerReq)
	if err != nil {
		h.log.Error("error while unmarshiling byte to object", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error unmarshiling byte to object",
		})
		return
	}
	if customerReq.Email != body.UserNameOrEmail {
		h.log.Error("error wrong email", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error wrong email",
		})
		return
	}
	id := uuid.New()
	h.jwthandler.Sub = id.String()
	h.jwthandler.Role = "user"
	h.jwthandler.Aud = []string{"todo-app"}

	accessToken, refreshToken, err := h.jwthandler.GenerateAuthJWT()

	if err != nil {
		h.log.Error("error occured while generating tokens")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong,please try again",
		})
		return
	}

	// Do not forget to update this part of method, it is very important

	// newUser.AccessToken = "ACCESS TOKEN NEED TO BE ASSIGNED"
	// newUser.RefreshToken = "REFRESH TOKEN ALSO NEED TO BE ASSIGNED"
	customerReq.AccessToken = accessToken
	customerReq.RefreshToken = refreshToken
	fmt.Println(customerReq.PassWord)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	customerResp, err := h.serviceManager.CustomerService().CreateCustomer(ctx, &customerReq)
	fmt.Println(err)
	if err != nil {
		h.log.Error("error while inserting customer", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error inserting customer",
		})
		return
	}

	// c.JSON(http.StatusCreated, customerResp)

	verified := models.VerifiedResponse{Id: int64(customerResp.Id), AccessToken: customerReq.AccessToken, RefreshToken: customerReq.RefreshToken}

	c.JSON(http.StatusOK, verified)
}
