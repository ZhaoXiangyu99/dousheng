package handlers

import (
	"douyin/consts"
	"douyin/http/request"
	"douyin/http/response"
	"douyin/service"
	"douyin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

// PublishVideo 发布视频
func PublishVideo(c *gin.Context) {
	var param request.PublishParam

	//绑定参数
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(500, response.Basic{StatusCode: -1, StatusMsg: "failed to bind params" + err.Error()})
		return
	}
	//去除title空格
	title := strings.Replace(param.Title, " ", "", -1)
	param.Title = title
	// =========================================将上传的视频保存到本地,并截取封面图===========================================

	//获取上传的视频文件
	data, _ := c.FormFile("data")
	videoPath := consts.VideoPath + param.Title + ".mp4"
	coverPath := consts.CoverPath + param.Title + ".jpg"

	//将上传视频保存到本地
	if err := c.SaveUploadedFile(data, videoPath); err != nil {
		fmt.Printf("%v", err)
		c.JSON(500, response.Basic{StatusCode: -1, StatusMsg: "failed to save file"})
		return
	}

	//将上传的视频截图,作为视频的封面
	if err := utils.GetCoverForVideo(videoPath, coverPath); err != nil {
		fmt.Printf("%v", err)
		c.JSON(500, response.Basic{StatusCode: -1, StatusMsg: "failed to get cover"})
		return
	}

	// ===========================================存入mysql=============================================================

	if err := service.PublishVideo(param.Title, param.Token); err != nil {
		c.JSON(500, response.Basic{StatusCode: -1, StatusMsg: "failed to insert into mysql"})
		return
	}

	c.JSON(200, response.Basic{StatusCode: 0, StatusMsg: "success"})
}
