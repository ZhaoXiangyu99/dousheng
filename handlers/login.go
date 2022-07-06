package handlers

import (
	"douyin/http/request"
	"douyin/http/response"
	"douyin/repository"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login 登录操作
func Login(c *gin.Context) {
	var param request.UserParam

	//绑定参数
	if err := c.ShouldBind(&param); err != nil {
		print("bind param err")
		c.JSON(401, response.Login{
			Basic:  response.Basic{StatusCode: -1, StatusMsg: "bind param err"},
			UserId: 0,
			Token:  "",
		})
		return
	}

	//校验参数
	if len(param.UserName) == 0 || len(param.PassWord) == 0 {
		print("param length err")
		c.JSON(402, response.Register{
			Basic:  response.Basic{StatusCode: -1, StatusMsg: "param length err"},
			UserId: 0,
			Token:  "",
		})
		return
	}

	if err := service.Login(param.UserName, param.PassWord); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.Login{
		Basic: response.Basic{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		UserId: repository.LoginUser.Id,
		Token:  repository.LoginUser.Token,
	})
}
