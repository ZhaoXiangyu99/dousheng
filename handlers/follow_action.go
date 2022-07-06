package handlers

import (
	"douyin/http/request"
	"douyin/http/response"
	"douyin/service"
	"github.com/gin-gonic/gin"
)

// FollowAction 添加关注或取消关注
func FollowAction(c *gin.Context) {

	var param request.FollowActionParam

	//绑定参数
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(500, response.Basic{StatusCode: -1, StatusMsg: "failed to bind param"})
		return
	}

	//如果action_type为1,就添加关注,否则取消关注
	if param.ActionType == 1 {
		if err := service.AddRelation(param.ToUserId, param.Token); err != nil {
			c.JSON(500, response.Basic{StatusCode: -1, StatusMsg: "failed of add relation"})
			return
		}
	} else {
		if err := service.DelRelation(param.ToUserId, param.Token); err != nil {
			c.JSON(500, response.Basic{StatusCode: -1, StatusMsg: "failed of delete relation"})
			return
		}
	}

	c.JSON(200, response.Basic{StatusCode: 0, StatusMsg: "success"})

}
