package v1

import (
	"fmt"
	"net/http"
	"time"

	jwthandler "github.com/Asliddin3/exam-api-gateway/api/token"
	pbc "github.com/Asliddin3/exam-api-gateway/genproto/customer"
	"github.com/Asliddin3/exam-api-gateway/pkg/logger"

	"github.com/Asliddin3/exam-api-gateway/api/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

// @Summary login
// @Tags Auth
// @Accept json
// @Produce json
// @Param loginData body models.Login true "login data"
// @Success 200 {object} customer.LoginResponse
// @Failure 400 {object} models.Error
// @Router /login [post]
func (h *handlerV1) Login(c *gin.Context) {

	var body = &models.Login{}
	c.ShouldBindJSON(body)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	exists, err := h.serviceManager.CustomerService().CheckField(ctx, &pbc.CheckFieldRequest{Key: "username", Value: body.Username})
	if err != nil {
		h.log.Error("error while logging into ", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong, please try again",
		})
		return
	}

	if !exists.Exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong username or password",
		})
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	user, err := h.serviceManager.CustomerService().Login(ctx, &pbc.LoginRequest{UserName: body.Username, Password: body.Password})
	if err != nil {
		h.log.Error("error while logging into ", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong, please try again",
		})
		return
	}

	token := jwthandler.JWTHandler{
		Sub:       string(user.Id),
		Role:      "user",
		Iss:       "customer-api",
		SigninKey: h.cfg.SigninKey,
	}
	fmt.Println("in token ", token.SigninKey)
	access, refresh, err := token.GenerateAuthJWT()
	if err != nil {
		h.log.Error("error while generating tokens")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong, please try again",
		})
		return
	}

	res, err := jwt.Parse(access, func(t *jwt.Token) (interface{}, error) { return []byte(h.cfg.SignKey), nil })
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Claims)

	ctx, cancel = context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	h.serviceManager.CustomerService().RefreshToken(ctx,
		&pbc.RefreshTokenRequest{
			Id:           user.Id,
			RefreshToken: refresh})

	user.RefreshToken = refresh
	user.PassWord = ""
	c.JSON(http.StatusOK, user)
}
