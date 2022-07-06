package handlers

import (
	"douyin/http/request"
	"douyin/http/response"
	"douyin/repository"
	"douyin/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// Feed 视频流接口
func Feed(c *gin.Context) {
	var param request.FeedParam

	if err := c.ShouldBind(&param); err != nil {
		c.JSON(500, response.Feed{
			Basic:     response.Basic{StatusCode: -1, StatusMsg: "failed to bind params"},
			VideoList: nil,
			NextTime:  time.Now().Unix(),
		})
		fmt.Printf("bind param failed")
		return
	}

	if err := service.GetLatestVideo(param.LatestTime, param.Token); err != nil {
		c.JSON(500, response.Feed{
			Basic:     response.Basic{StatusCode: -1, StatusMsg: "failed to get latest video from sql"},
			VideoList: nil,
			NextTime:  time.Now().Unix(),
		})
		fmt.Printf("get video failed" + err.Error())
		return
	}

	c.JSON(200, response.Feed{
		Basic:     response.Basic{StatusCode: 0, StatusMsg: "success"},
		VideoList: repository.LatestList,
		NextTime:  repository.NextTime,
	})
}
