package handlers

import (
	"douyin/http/request"
	"douyin/http/response"
	"douyin/repository"
	"douyin/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

// FavoriteList 获取点赞列表
func FavoriteList(c *gin.Context) {
	var param request.FavoriteListParam
	if err := c.ShouldBind(&param); err != nil {
		fmt.Printf("%+v", err.Error())
		c.JSON(500, response.FavoriteList{
			Basic:     response.Basic{StatusCode: -1, StatusMsg: "failed to bind param"},
			VideoList: []repository.Video{},
		})
		return
	}

	if err := service.GetFavoriteList(param.UserId, param.Token); err != nil {
		c.JSON(500, response.FavoriteList{
			Basic:     response.Basic{StatusCode: -1, StatusMsg: "failed to get favorite list"},
			VideoList: []repository.Video{},
		})
		return
	}

	c.JSON(200, response.FavoriteList{
		Basic:     response.Basic{StatusCode: 0, StatusMsg: "success to get favorite list"},
		VideoList: repository.FavoriteList,
	})
}
