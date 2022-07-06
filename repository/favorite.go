package repository

import (
	"douyin/utils"
	"strconv"
)

type Favorite struct {
	Id      int64  `gorm:"column:id" json:"id"`
	UserId  int64  `gorm:"column:user_id" json:"user_id"`
	VideoId int64  `gorm:"column:video_id" json:"video_id"`
	Token   string `gorm:"column:token" json:"token"`
}

var FavoriteList []Video

// AddFavorite 新增点赞操作
func AddFavorite(videoId int64, token string) error {
	//首先获取视频的点赞总数
	var favoriteCount int64
	if err := db.Table("videos").Where("id = ?", videoId).
		Select("favorite_count").Scan(&favoriteCount).Error; err != nil {
		return err
	}

	uid, _ := strconv.ParseInt(token, 10, 64)
	//构建favorite结构体
	favorite := &Favorite{
		Id:      utils.GetOnlyId(),
		UserId:  uid,
		VideoId: videoId,
		Token:   token,
	}

	//将点赞记录存入数据库
	if err := db.Table("favorite").Create(favorite).Error; err != nil {
		return err
	}

	//favoriteCount加1
	favoriteCount++

	//更新videos表中的favoriteCount和isFavorite
	if favoriteCount == 0 {
		if err := db.Table("videos").Where("id = ?", videoId).
			Updates(map[string]interface{}{"favorite_count": favoriteCount, "is_favorite": false}).Error; err != nil {
			return err
		}
	} else {
		if err := db.Table("videos").Where("id = ?", videoId).
			Updates(map[string]interface{}{"favorite_count": favoriteCount, "is_favorite": true}).Error; err != nil {
			return err
		}
	}
	return nil
}

// DelFavorite 删除favorite记录
func DelFavorite(videoId int64, token string) error {

	uid, _ := strconv.ParseInt(token, 10, 64)
	//获取视频的点赞总数
	var favoriteCount int64
	if err := db.Table("videos").Where("id = ?", videoId).
		Select("favorite_count").Scan(&favoriteCount).Error; err != nil {
		return err
	}

	//favoriteCount减1
	favoriteCount--

	//删除favorite记录
	if err := db.Table("favorite").
		Where("video_id = ?", videoId).Where("user_id", uid).Where("token = ?", token).
		Delete(&Favorite{}).Error; err != nil {
		return err
	}

	//更新favoriteCount。如果favoriteCount为0,就把is_favorite置为false
	if favoriteCount == 0 {
		if err := db.Table("videos").Where("id = ?", videoId).
			Updates(map[string]interface{}{"favorite_count": favoriteCount, "is_favorite": false}).Error; err != nil {
			return err
		}
	} else {
		if err := db.Table("videos").Where("id = ?", videoId).
			Updates(map[string]interface{}{"favorite_count": favoriteCount, "is_favorite": true}).Error; err != nil {
			return err
		}
	}

	return nil
}

// GetFavoriteList 获取点赞列表
func GetFavoriteList(userId int64, token string) error {
	if err := db.Table("videos").Where("id IN (?)",
		db.Table("favorite").Select("video_id").
			Where("user_id = ?", userId).
			Where("token = ?", token)).Preload("Author").Find(&FavoriteList).Error; err != nil {
		return err
	}
	return nil
}
