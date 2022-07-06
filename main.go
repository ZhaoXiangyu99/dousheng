package main

import (
	"douyin/consts"
	"douyin/handlers"
	"douyin/repository"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	//初始化数据库
	if err := Init(); err != nil {
		os.Exit(-1)
	}

	//初始化路由
	Router()
}

func Init() error {
	if err := repository.Init(); err != nil {
		return err
	}
	return nil
}

func Router() {
	r := gin.Default()

	//设置静态资源路径
	r.Static("/cover/", consts.CoverPath)
	r.Static("/video/", consts.VideoPath)

	//公共前置路由douyin
	douyin := r.Group("/douyin")

	//用户接口
	user := douyin.Group("/user")

	user.POST("/login/", handlers.Login)       //登录
	user.POST("/register/", handlers.Register) //注册
	user.GET("/", handlers.UserInfo)           //登录之后的处理

	//视频接口
	video := douyin

	video.GET("/feed/", handlers.Feed)

	//视频发布接口
	publish := video.Group("/publish")

	publish.POST("/action/", handlers.PublishVideo)
	publish.GET("/list/", handlers.PublishList)

	//视频点赞接口
	favorite := video.Group("/favorite")

	favorite.POST("/action/", handlers.FavoriteAction)
	favorite.GET("/list/", handlers.FavoriteList)

	//视频评论接口
	comment := video.Group("/comment")

	comment.POST("/action/", handlers.Comment)
	comment.GET("/list/", handlers.CommentList)

	//关系接口
	relation := douyin.Group("/relation")

	relation.POST("/action/", handlers.FollowAction)
	relation.GET("/follow/list/", handlers.FollowList)
	relation.GET("/follower/list/", handlers.FollowerList)

	if err := http.ListenAndServe("0.0.0.0:8080", r); err != nil {
		klog.Fatal(err)
	}
}
