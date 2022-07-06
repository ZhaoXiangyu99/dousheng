package handlers

import (
	"douyin/http/response"
	"douyin/repository"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

// UserInfo 登录后获取用户详细信息
func UserInfo(c *gin.Context) {
	uid, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)

	if err := service.GetUserInfo(uid); err != nil {
		c.JSON(500, response.Info{
			Basic: response.Basic{StatusCode: -1, StatusMsg: "get user info failed "},
			User:  repository.User{},
		})
		return
	}

	c.JSON(200, response.Info{
		Basic: response.Basic{StatusCode: 0, StatusMsg: "success"},
		User:  repository.LoginUser,
	})
}
