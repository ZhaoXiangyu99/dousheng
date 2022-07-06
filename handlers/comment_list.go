package handlers

import (
	"douyin/http/request"
	"douyin/http/response"
	"douyin/repository"
	"douyin/service"
	"github.com/gin-gonic/gin"
)

// CommentList 获取评论列表
func CommentList(c *gin.Context) {
	var param request.CommentListParam
	if err := c.ShouldBindQuery(&param); err != nil {
		c.JSON(500, response.CommentList{
			Basic:       response.Basic{StatusCode: -1, StatusMsg: "failed to bind param"},
			CommentList: []repository.Comment{},
		})
		return
	}

	if err := service.GetCommentList(param.VideoId, param.Token); err != nil {
		c.JSON(500, response.CommentList{
			Basic:       response.Basic{StatusCode: -1, StatusMsg: "failed to get comment list"},
			CommentList: []repository.Comment{},
		})
		return
	}

	c.JSON(200, response.CommentList{
		Basic:       response.Basic{StatusCode: 0, StatusMsg: "success to get comment list"},
		CommentList: repository.CommentList,
	})

}
