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

type AdminLogin struct {
	UserName string
	PassWord string
}

// @BasePath /api/v1

// PingExample godoc
// @Summary create customer with info
// @Description this func login admin
// @Security        BearerAuth
// @Tags Admin
// @Accept json
// @Produce json
// @Param admin body models.AdminRequest true "Admin"
// @Success 201 {object} models.AdminResponse
// @Router /admin [post]
func (h *handlerV1) LoginAdmin(c *gin.Context) {

	var body = &models.AdminRequest{}
	err := c.ShouldBindJSON(body)
	fmt.Println(err, body)
	if err != nil {
		h.log.Error("error binding json", logger.Error(err))
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "error binding json",
		})
		return
	}

	exists, err := h.storage.LoginAdmin(body)
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
	fmt.Println("admin login", h.cfg.SigninKey)
	token := jwthandler.JWTHandler{
		Sub:       string(exists.Id),
		Role:      "admin",
		Iss:       "admin",
		SigninKey: "supersecret",
		Aud:       []string{"admin-app"},
	}
	fmt.Println("in admin", token.SigninKey)
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
	Admin := models.AdminResponse{}
	Admin.AccessToken = access
	Admin.UserName = body.UserName
	c.JSON(http.StatusOK, Admin)
}
