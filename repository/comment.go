package repository

import (
	"douyin/utils"
	"time"
)

type Comment struct {
	Id         int64     `gorm:"column:id" json:"id"`
	User       User      `gorm:"foreignKey:Id;references:user_id" json:"user"`
	UserId     int64     `gorm:"column:user_id" json:"user_id"`
	VideoId    int64     `gorm:"column:video_id" json:"video_id"`
	Content    string    `gorm:"column:content" json:"content"`
	Token      string    `gorm:"column:token" json:"token"`
	CreateDate time.Time `gorm:"column:create_date" json:"create_date"`
}

var ActionComment Comment
var CommentList []Comment

// AddComment 添加评论
func AddComment(userId, videoId int64, token, commentText string) error {
	comment := &Comment{
		Id:         utils.GetOnlyId(),
		UserId:     userId,
		VideoId:    videoId,
		Content:    commentText,
		Token:      token,
		CreateDate: time.Now(),
	}

	//将comment添加进数据库
	if err := db.Table("comment").Create(comment).Error; err != nil {
		return err
	}

	// ===============================获取视频的评论总数,并且评论总数 + 1====================================================

	//从sql中获取comment_count
	var commentCount int64
	if err := db.Table("videos").Where("id = ?", videoId).
		Select("comment_count").Scan(&commentCount).Error; err != nil {
		return err
	}

	//comment_count加一
	commentCount++
	//更新videos表中的comment_count
	if err := db.Table("videos").Where("id = ?", videoId).
		Update("comment_count", commentCount).Error; err != nil {
		return err
	}

	//获取要返回的comment表
	if err := db.Table("comment").Where("id = ?", comment.Id).Where("user_id IN (?)",
		db.Table("users").Select("id").Where("id = ?", userId)).
		Preload("User").Find(&ActionComment).Error; err != nil {
		return err
	}

	return nil
}

// DeleteComment 删除评论
func DeleteComment(deleteId, videoId int64, token string) error {
	//从sql中删除token为给定值的评论
	if err := db.Where("token = ?", token).Delete(&Comment{}, deleteId).Error; err != nil {
		return err
	}

	// =========================================删除评论后,videos表中的comment_count要减1==================================
	//从sql中获取comment_count
	var commentCount int64
	if err := db.Table("videos").Where("id = ?", videoId).
		Select("comment_count").Scan(&commentCount).Error; err != nil {
		return err
	}

	//comment_count加一
	commentCount--
	//更新videos表中的comment_count
	if err := db.Table("videos").Where("id = ?", videoId).
		Update("comment_count", commentCount).Error; err != nil {
		return err
	}

	//获取要返回的comment
	if err := db.Table("comment").Where("id = ?", deleteId).Find(&ActionComment).Error; err != nil {
		return err
	}
	return nil
}

// GetCommentList 获取评论列表
func GetCommentList(videoId int64, token string) error {
	if err := db.Table("comment").Where("token = ?", token).
		Where("video_id = ?", videoId).Preload("User").Find(&CommentList).Error; err != nil {
		return err
	}
	return nil
}
