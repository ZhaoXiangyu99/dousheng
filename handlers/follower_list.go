package handlers

import (
	"douyin/http/request"
	"douyin/http/response"
	"douyin/repository"
	"douyin/service"
	"github.com/gin-gonic/gin"
)

// FollowerList 粉丝列表
func FollowerList(c *gin.Context) {
	var param request.RelationParam

	if err := c.ShouldBind(&param); err != nil {
		c.JSON(500, response.FollowerList{
			Basic:    response.Basic{StatusCode: -1, StatusMsg: "bind param failed"},
			UserList: []repository.User{},
		})
		return
	}

	if err := service.FollowerList(param.UserId, param.Token); err != nil {
		c.JSON(500, response.FollowerList{
			Basic:    response.Basic{StatusCode: -1, StatusMsg: "failed to get follower list"},
			UserList: []repository.User{},
		})
		return
	}

	c.JSON(200, response.FollowerList{
		Basic:    response.Basic{StatusCode: 0, StatusMsg: "success"},
		UserList: repository.FollowerUsers,
	})
}
