package handlers

import (
	"douyin/http/request"
	"douyin/http/response"
	"douyin/repository"
	"douyin/service"
	"github.com/gin-gonic/gin"
)

// PublishList 视频发布列表
func PublishList(c *gin.Context) {
	var param request.PublishListParam

	if err := c.ShouldBind(&param); err != nil {
		c.JSON(500, response.PublishList{
			Basic:     response.Basic{StatusCode: -1, StatusMsg: "failed to bind param"},
			VideoList: []repository.Video{},
		})
		return
	}

	if err := service.PublishList(param.UserId, param.Token); err != nil {
		c.JSON(500, response.PublishList{
			Basic:     response.Basic{StatusCode: -1, StatusMsg: "failed to get publish list"},
			VideoList: []repository.Video{},
		})
		return
	}
	
	c.JSON(200, response.PublishList{
		Basic:     response.Basic{StatusCode: 0, StatusMsg: "success"},
		VideoList: repository.PublishList,
	})
}
