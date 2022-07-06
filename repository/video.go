package repository

import (
	"douyin/consts"
	"douyin/utils"
	"errors"
	"strconv"
	time2 "time"
)

// Video 需要通过响应返回的Video格式
type Video struct {
	Id            int64      `gorm:"column:id" json:"id"`
	Author        User       `json:"author"`
	AuthorId      int64      `json:"author_id" gorm:"column:author_id"`
	PlayUrl       string     `gorm:"column:play_url" json:"play_url"`
	CoverUrl      string     `gorm:"column:cover_url" json:"cover_url"`
	FavoriteCount int64      `gorm:"column:favorite_count" json:"favorite_count"`
	CommentCount  int64      `gorm:"column:comment_count" json:"comment_count"`
	IsFavorite    bool       `gorm:"column:is_favorite" json:"is_favorite"`
	Title         string     `gorm:"column:title" json:"title"`
	CreateTime    time2.Time `gorm:"column:create_time" json:"create_time"`
	Token         string     `gorm:"column:token" json:"token"`
}

var LatestList []Video
var PublishList []Video
var NextTime int64

// GetLatestVideo 查询最新发布的视频,即发布时间小于给定的最新时间
func GetLatestVideo(latestTime int64, token string) error {
	t := time2.UnixMilli(latestTime)
	if err := db.Table("videos").
		Where("create_time <= ?", t).
		Preload("Author").
		Order("create_time desc").Find(&LatestList).Error; err != nil {
		return err
	}
	if len(LatestList) == 0 {
		return errors.New("no such videos")
	}

	//将数据库中的PlayUrl和CoverUrl加上前缀
	addPrefix(LatestList)

	//将VideoList中最大的create_time作为下次请求的latest_time
	for _, video := range LatestList {
		cTime := video.CreateTime.UnixMilli()
		if NextTime < cTime {
			NextTime = cTime
		}
	}
	return nil

}

// GetUserVideoList 获取用户发布视频的列表
func GetUserVideoList(userId int64, token string) error {
	if err := db.Table("videos").Where("author_id IN (?)",
		db.Table("users").Select("id").Where("id = ?", userId).
			Where("token = ?", token)).Find(&PublishList).Error; err != nil {
		return err
	}
	return nil
}

// PublishVideo 发布视频
func PublishVideo(title, token string) error {
	uid, _ := strconv.ParseInt(token, 10, 64)
	video := &Video{
		Id:         utils.GetOnlyId(),
		AuthorId:   uid,
		PlayUrl:    title + ".mp4",
		CoverUrl:   title + ".jpg",
		Title:      title,
		CreateTime: time2.Now(),
	}
	if err := db.Table("videos").Create(video).Error; err != nil {
		return err
	}
	return nil
}

//给PlayUrl和CoverUrl加上前缀
func addPrefix(videoList []Video) {
	for i := 0; i < len(videoList); i++ {
		videoList[i].PlayUrl = consts.IPAddr + "/video/" + videoList[i].PlayUrl
		videoList[i].CoverUrl = consts.IPAddr + "/cover/" + videoList[i].CoverUrl
	}
}
