package v1

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/Asliddin3/exam-api-gateway/genproto/post"
	l "github.com/Asliddin3/exam-api-gateway/pkg/logger"
)

// @BasePath /api/v1
// @Summary get post
// @Description this func get post
// @Tags post
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} post.PostResponseCustomer
// @Router /post/{id} [get]
func (h *handlerV1) GetPost(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	id, err := strconv.ParseInt(guid, 10, 64)
	body := &post.PostId{
		Id: id,
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to convert string to int", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	response, err := h.serviceManager.PostService().GetPost(ctx, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get post", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, response)
}

// @BasePath /api/v1
// @Summary get posts
// @Description this func get posts
// @Tags post
// @Accept json
// @Produce json
// @Success 200 {object} post.ListAllPostResponse
// @Router /post/list [get]
func (h *handlerV1) GetListPosts(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	response, err := h.serviceManager.PostService().GetListPosts(ctx, &post.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get posts", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, response)
}

// @BasePath /api/v1
// @Summary delete post
// @Description this func delete post
// @Tags post
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 "success"
// @Router /post/delete/{id} [delete]
func (h *handlerV1) DeletePost(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	id, err := strconv.ParseInt(guid, 10, 64)
	body := &post.PostId{
		Id: id,
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to convert string to int", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	response, err := h.serviceManager.PostService().DeletePost(ctx, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete post", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, response)
}

// @BasePath /api/v1
// @Summary update post
// @Description this func update post
// @Tags post
// @Accept json
// @Produce json
// @Param post body post.PostUpdate true "Post"
// @Success 200 {object} post.PostResponse
// @Router /post/update [patch]
func (h *handlerV1) UpdatePost(c *gin.Context) {
	var (
		body        post.PostUpdate
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
	response, err := h.serviceManager.PostService().UpdatePost(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update post", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, response)
}

// @BasePath /api/v1

// PingExample godoc
// @Summary create post with info
// @Description this func create post
// @Tags post
// @Accept json
// @Produce json
// @Param customer body post.PostRequest true "Post"
// @Success 201 {object} post.PostResponse
// @Router /post [post]
func (h *handlerV1) CreatePost(c *gin.Context) {
	var (
		body        post.PostRequest
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
	response, err := h.serviceManager.PostService().CreatePost(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create post", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}
