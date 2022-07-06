package response

import "douyin/repository"

//===========================================封装所有http请求的响应体======================================================

// Basic 基础响应:所有响应都要返回的字段
type Basic struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

// Feed 获取所有视频信息
type Feed struct {
	Basic
	VideoList []repository.Video `json:"video_list"`
	NextTime  int64              `json:"next_time"`
}

// Register 注册响应
type Register struct {
	Basic
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

// Login 登陆响应
type Login struct {
	Basic
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

// Info 用户信息
type Info struct {
	Basic
	User repository.User `json:"user"`
}

// FollowList 关注列表
type FollowList struct {
	Basic
	UserList []repository.User `json:"user_list"`
}

// FollowerList 粉丝列表
type FollowerList struct {
	Basic
	UserList []repository.User `json:"user_list"`
}

// PublishList 用户发布视频的列表
type PublishList struct {
	Basic
	VideoList []repository.Video `json:"video_list"`
}

// FavoriteList 用户点赞的视频列表
type FavoriteList struct {
	Basic
	VideoList []repository.Video `json:"video_list"`
}

// CommentAction 用户评论
type CommentAction struct {
	Basic
	Comment repository.Comment `json:"comment"`
}

type CommentList struct {
	Basic
	CommentList []repository.Comment `json:"comment_list"`
}
