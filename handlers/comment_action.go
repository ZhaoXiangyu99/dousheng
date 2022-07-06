package handlers

import (
	"douyin/http/request"
	"douyin/http/response"
	"douyin/repository"
	"douyin/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

// Comment 添加或删除评论
func Comment(c *gin.Context) {
	var param request.CommentActionParam
	if err := c.ShouldBind(&param); err != nil {
		fmt.Printf("%+v", err.Error())
		c.JSON(500, response.CommentAction{
			Basic:   response.Basic{StatusCode: -1, StatusMsg: "failed to bind param"},
			Comment: repository.Comment{},
		})
		return
	}

	if param.ActionType == 1 {
		if err := service.AddComment(param.UserId, param.VideoId, param.Token, param.CommentText); err != nil {
			c.JSON(500, response.CommentAction{
				Basic:   response.Basic{StatusCode: -1, StatusMsg: "failed to add comment"},
				Comment: repository.Comment{},
			})
		} else {
			c.JSON(200, response.CommentAction{
				Basic:   response.Basic{StatusCode: 0, StatusMsg: "success to add comment"},
				Comment: repository.ActionComment,
			})
		}
	}
	if param.ActionType == 2 {
		if err := service.DeleteComment(param.DeleteId, param.VideoId, param.Token); err != nil {
			c.JSON(500, response.CommentAction{
				Basic:   response.Basic{StatusCode: -1, StatusMsg: "failed to delete comment"},
				Comment: repository.Comment{},
			})
		} else {
			c.JSON(200, response.CommentAction{
				Basic:   response.Basic{StatusCode: 0, StatusMsg: "success to add comment"},
				Comment: repository.ActionComment,
			})
		}
	}
}
