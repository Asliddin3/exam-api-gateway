package v1

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/Asliddin3/exam-api-gateway/genproto/review"
	"github.com/Asliddin3/exam-api-gateway/pkg/logger"
	l "github.com/Asliddin3/exam-api-gateway/pkg/logger"
)

// @BasePath /api/v1
// @Summary delete review
// @Description this func delete review
// @Security        BearerAuth
// @Tags review
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 "success"
// @Router /review/delete/{id} [delete]
func (h *handlerV1) DeleteReview(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	id, err := strconv.ParseInt(guid, 10, 64)
	body := &review.ReviewId{
		Id: id,
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	claims, err := GetClaims(*h, c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "you are not authorized",
		})
		h.log.Error("Checking Authorozation", logger.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	reviewInfo, err := h.serviceManager.ReviewService().GetReviewById(ctx, &review.ReviewId{Id: id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error checking review customer id",
		})
		return
	}
	if claims.Sub != reviewInfo.CustomerId && claims.Role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"info": "You haven't access ",
		})
		return
	}

	response, err := h.serviceManager.ReviewService().DeleteReview(ctx, body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete review", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, response)
}

// PingExample godoc
// @Summary create review
// @Description this func create review
// @Security        BearerAuth
// @Tags review
// @Accept json
// @Produce json
// @Param review body review.ReviewRequest true "Review"
// @Success 200 {object} review.Review
// @Router /review [post]
func (h *handlerV1) CreateReview(c *gin.Context) {
	var (
		body        review.ReviewRequest
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	response, err := h.serviceManager.ReviewService().CreateReview(ctx, &body)
	if response.Id == 0 && response.PostId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"info": "this review alreadi exists",
		})
		h.log.Error("failed to create review cause of exists review", l.Error(err))
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create review", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}

//@Summary get review by id
//@Description this func get review by id
// @Security        BearerAuth
//@Tags review
//@Accept json
//@Produce json
//@Param id path int true "id"
//@Success 200 {object} review.Review
//@Router /review/{id} [get]
func (h *handlerV1) GetReviewById(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	idReq := c.Param("id")
	id, err := strconv.ParseInt(idReq, 10, 64)
	body := &review.ReviewId{
		Id: id,
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to convert int to string", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	response, err := h.serviceManager.ReviewService().GetReviewById(ctx, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get review", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, response)
}
