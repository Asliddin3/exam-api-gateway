package v1

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/Asliddin3/exam-api-gateway/genproto/review"
	l "github.com/Asliddin3/exam-api-gateway/pkg/logger"
)

// @BasePath /api/v1
// @Summary get review
// @Description this func get post review
// @Tags review
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 "success"
// @Router /review/{id} [get]
func (h *handlerV1) GetPostReview(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	id, err := strconv.ParseInt(guid, 10, 64)
	body := &review.PostId{
		Id: id,
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	response, err := h.serviceManager.ReviewService().GetPostReview(ctx, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get post review", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, response)
}

// @BasePath /api/v1
// @Summary delete review
// @Description this func delete review
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
	body := &review.PostId{
		Id: id,
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	response, err := h.serviceManager.ReviewService().DeleteReview(ctx, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
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
// @Tags review
// @Accept json
// @Produce json
// @Param review body review.Review true "Review"
// @Success 200 {object} review.Review
// @Router /review [post]
func (h *handlerV1) CreateReview(c *gin.Context) {
	var (
		body        review.Review
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
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create review", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}
