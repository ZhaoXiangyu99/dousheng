package handlers

import (
	"douyin/http/request"
	"douyin/http/response"
	"douyin/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

// FavoriteAction 点赞和取消点赞
func FavoriteAction(c *gin.Context) {
	var param request.FavoriteActionParam
	if err := c.ShouldBind(&param); err != nil {
		fmt.Printf("%+v", err.Error())
		c.JSON(500, response.Basic{StatusCode: -1, StatusMsg: "failed to bind param "})
		return
	}

	if param.ActionType == 1 {
		if err := service.AddFavorite(param.VideoId, param.Token); err != nil {
			c.JSON(500, response.Basic{StatusCode: -1, StatusMsg: "failed to add favorite "})
			return
		}
	} else {
		if err := service.DelFavorite(param.VideoId, param.Token); err != nil {
			c.JSON(500, response.Basic{StatusCode: -1, StatusMsg: "failed to delete favorite "})
			return
		}
	}

	c.JSON(200, response.Basic{StatusCode: 0, StatusMsg: "success"})
}
