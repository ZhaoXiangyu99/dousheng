package handlers

import (
	"douyin/http/request"
	"douyin/http/response"
	"douyin/repository"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register 注册
func Register(c *gin.Context) {
	var param request.UserParam

	//绑定参数
	if err := c.ShouldBind(&param); err != nil {
		print("bind param err")
		c.JSON(401, response.Register{
			Basic: response.Basic{
				StatusCode: -1,
				StatusMsg:  "bind param err",
			},
			UserId: 0,
			Token:  "",
		})
		return
	}

	//校验参数
	if len(param.PassWord) == 0 || len(param.UserName) == 0 {
		print("param length err")
		c.JSON(402, response.Register{
			Basic: response.Basic{
				StatusCode: -1,
				StatusMsg:  "param length err",
			},
			UserId: 0,
			Token:  "",
		})
		return
	}

	if err := service.Register(param.UserName, param.PassWord); err != nil {
		c.JSON(500, response.Register{
			Basic: response.Basic{
				StatusCode: -1,
				StatusMsg:  "register failed",
			},
			UserId: 0,
			Token:  "",
		})
		return
	}

	c.JSON(http.StatusOK, response.Register{
		Basic: response.Basic{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		UserId: repository.RUser.Id,
		Token:  repository.RUser.Token,
	})
}
