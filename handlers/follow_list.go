package handlers

import (
	"douyin/http/request"
	"douyin/http/response"
	"douyin/repository"
	"douyin/service"
	"github.com/gin-gonic/gin"
)

// FollowList 关注列表
func FollowList(c *gin.Context) {
	var param request.RelationParam

	if err := c.ShouldBind(&param); err != nil {
		c.JSON(500, response.FollowList{
			Basic:    response.Basic{StatusCode: -1, StatusMsg: "bind param failed" + err.Error()},
			UserList: []repository.User{},
		})
		return
	}

	if err := service.FollowList(param.UserId, param.Token); err != nil {
		c.JSON(500, response.FollowList{
			Basic:    response.Basic{StatusCode: -1, StatusMsg: "failed to get follow"},
			UserList: []repository.User{},
		})
		return
	}

	c.JSON(200, response.FollowList{
		Basic:    response.Basic{StatusCode: 0, StatusMsg: "success"},
		UserList: repository.FollowUsers,
	})
}
