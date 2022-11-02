package v1

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/Asliddin3/exam-api-gateway/genproto/post"
	"github.com/Asliddin3/exam-api-gateway/pkg/logger"
	l "github.com/Asliddin3/exam-api-gateway/pkg/logger"
)

// @BasePath /api/v1
// @Summary get post
// @Description this func get post
// @Tags post
// @Security        BearerAuth
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
// @Security        BearerAuth
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

// // @BasePath /api/v1
// // @Summary get posts
// // @Description this func get posts
// // @Security        BearerAuth
// // @Tags post
// // @Accept json
// // @Produce json
// // @Success 200 {object} post.ListPostResp
// // @Router /post/page [get]
// func (h *handlerV1) ListPostForPage(c *gin.Context) {
// 	var (
// 		body        post.ListPostReq
// 		jspbMarshal protojson.MarshalOptions
// 	)
// 	jspbMarshal.UseProtoNames = true
// 	err := c.ShouldBindJSON(&body)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		h.log.Error("failed to bind json", l.Error(err))
// 		return
// 	}
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
// 	defer cancel()
// 	response, err := h.serviceManager.PostService().ListPost(ctx, &body)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		h.log.Error("failed to get page posts", l.Error(err))
// 		return
// 	}
// 	c.JSON(http.StatusOK, response)
// }

// @BasePath /api/v1
// @Summary delete post
// @Description this func delete post
// @Security        BearerAuth
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
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error binding id",
		})
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
	postInfo, err := h.serviceManager.PostService().GetPost(ctx, &post.PostId{Id: id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error getting post info for check customer id",
		})
		return
	}
	if claims.Sub != postInfo.CustomerId && claims.Role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"info": "You haven't access ",
		})
		return
	}

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

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	// defer cancel()
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
// @Security        BearerAuth
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
	claims, err := GetClaims(*h, c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "you are not authorized",
		})
		h.log.Error("Checking Authorozation", logger.Error(err))
		return
	}
	if claims.Sub != body.CustomerId && claims.Role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"info": "You haven't access ",
		})
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
// @Security        BearerAuth
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

// @BasePath /api/v1
// @Summary search post
// @Description this func search post
// @Security        BearerAuth
// @Tags post
// @Accept json
// @Produce json
// @Param page path int true "page"
// @Param limit path int true "limit"
// @Param parameters path []string true "parameters"
// @Param orderby path string true "orderby"
// @Success 200 "success"
// @Router /post/search/{page}/{limit}/{parameters}/{orderby} [get]
func (h *handlerV1) SearchPost(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	pageStr := c.Param("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "converting page string to int",
		})
		return
	}
	limitStr := c.Param("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "converting limit string to int",
		})
		return
	}
	parametrsStr := c.Param("parameters")
	fmt.Println(parametrsStr)
	parametrReq := make(map[string]string)
	mapParam := strings.Split(parametrsStr, ",")
	for _, param := range mapParam {
		fmt.Println(param)
		keyValSlice := strings.Split(param, ".")
		parametrReq[keyValSlice[0]] = keyValSlice[1]
	}

	orderbyStr := c.Param("orderby")
	body := post.SearchRequest{
		Limit:   int64(limit),
		Page:    int64(page),
		OrderBy: orderbyStr,
	}
	for key, val := range parametrReq {
		body.Parametrs = append(body.Parametrs, &post.KetValue{
			Key:   key,
			Value: val,
		})
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	searchResp, err := h.serviceManager.PostService().SearchOrderedPagePost(ctx, &body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error getting response from post service",
		})
		h.log.Error("error getting search result", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, searchResp)
}
