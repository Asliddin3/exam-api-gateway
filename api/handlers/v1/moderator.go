package v1

import (
	"fmt"
	"net/http"

	"github.com/Asliddin3/exam-api-gateway/api/models"
	jwthandler "github.com/Asliddin3/exam-api-gateway/api/token"
	"github.com/Asliddin3/exam-api-gateway/pkg/logger"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type ModeratorLogin struct {
	UserName string
	PassWord string
}

// @BasePath /api/v1

// PingExample godoc
// @Summary login moderator
// @Description this func login moderator
// @Security        BearerAuth
// @Tags Admin
// @Accept json
// @Produce json
// @Param moderator body models.ModeratorRequest true "Moderator"
// @Success 201 {object} models.ModeratorResponse
// @Router /moderator [post]
func (h *handlerV1) LoginModerator(c *gin.Context) {

	var body = &models.ModeratorRequest{}
	err := c.ShouldBindJSON(body)
	fmt.Println(err, body)
	if err != nil {
		h.log.Error("error binding json", logger.Error(err))
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "error binding json",
		})
		return
	}

	exists, err := h.storage.LoginModerator(body)
	if err != nil {
		h.log.Error("error while logging into ", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong, please try again",
		})
		return
	}
	if body.PassWord != exists.PassWord {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong password",
		})
	}
	fmt.Println("Moderator login", h.cfg.SigninKey)
	token := jwthandler.JWTHandler{
		Sub:       string(exists.Id),
		Role:      "moderator",
		Iss:       "moderator",
		SigninKey: "supersecret",
		Aud:       []string{"Moderator-app"},
	}
	fmt.Println("in Moderator", token.SigninKey)
	access, _, err := token.GenerateAuthJWT()
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
	Moderator := models.ModeratorResponse{}
	Moderator.AccessToken = access
	Moderator.UserName = body.UserName
	c.JSON(http.StatusOK, Moderator)
}
