package repository

import (
	"douyin/utils"
	"strconv"
)

type Relation struct {
	Id       int64  `gorm:"column:id" json:"id"`
	UserId   int64  `gorm:"column:user_id" json:"user_id"`
	ToUserId int64  `gorm:"column:to_user_id" json:"to_user_id"`
	Token    string `gorm:"column:token" json:"token"`
}

var FollowUsers []User
var FollowerUsers []User

// AddRelation 添加关注
func AddRelation(toUserId int64, token string) error {

	userId, _ := strconv.ParseInt(token, 10, 64)
	//首先获取user的follow_count和toUser的follower_count
	var followCount, followerCount int64
	if err := db.Table("users").Where("id = ?", userId).
		Select("follow_count").Scan(&followCount).Error; err != nil {
		return err
	}
	if err := db.Table("users").Where("id = ?", toUserId).
		Select("follower_count").Scan(&followerCount).Error; err != nil {
		return err
	}

	//构建关系结构体
	relation := &Relation{
		Id:       utils.GetOnlyId(),
		UserId:   userId,
		ToUserId: toUserId,
		Token:    token,
	}

	//将关注的关系存入relations表中
	if err := db.Table("relations").Create(relation).Error; err != nil {
		return err
	}

	followCount++
	followerCount++
	//如果是关注,发起关注用户的follow_count+1,被关注用户的follower_count+1
	if err := db.Table("users").Where("id = ?", userId).
		Updates(map[string]interface{}{"follow_count": followCount, "is_follow": true}).Error; err != nil {
		return err
	}
	if err := db.Table("users").Where("id = ?", toUserId).
		Updates(map[string]interface{}{"follower_count": followerCount, "is_follow": true}).Error; err != nil {
		return err
	}

	return nil
}

// DelRelation 取消关注
func DelRelation(toUserId int64, token string) error {

	userId, _ := strconv.ParseInt(token, 10, 64)
	//首先获取user的follow_count和toUser的follower_count
	var followCount, followerCount int64
	if err := db.Table("users").Where("id = ?", userId).
		Select("follow_count").Scan(&followCount).Error; err != nil {
		return err
	}
	if err := db.Table("users").Where("id = ?", toUserId).
		Select("follower_count").Scan(&followerCount).Error; err != nil {
		return err
	}

	if err := db.Where("user_id = ?", userId).Where("to_user_id = ?", toUserId).
		Where("token = ?", token).Delete(&Relation{}).Error; err != nil {
		return err
	}

	isFollow := true
	followCount--
	if followCount < 0 {
		followCount = 0
	}
	followerCount--
	if followerCount < 0 {
		followerCount = 0
		isFollow = false
	}
	if err := db.Table("users").Where("id = ?", userId).
		Updates(map[string]interface{}{"follow_count": followCount, "is_follow": isFollow}).Error; err != nil {
		return err
	}
	if err := db.Table("users").Where("id = ?", toUserId).
		Updates(map[string]interface{}{"follower_count": followerCount, "is_follow": isFollow}).Error; err != nil {
		return err
	}
	return nil
}

// FollowList 关注列表
func FollowList(userId int64, token string) error {
	if err := db.Table("users").Where("id IN (?)",
		db.Table("relations").Select("to_user_id").
			Where("user_id = ?", userId).Where("token = ?", token)).Find(&FollowUsers).Error; err != nil {
		return err
	}
	return nil
}

// FollowerList 粉丝列表
func FollowerList(toUserId int64, token string) error {
	if err := db.Table("users").Where("id IN (?)",
		db.Table("relations").Select("user_id").
			Where("to_user_id = ?", toUserId)).Find(&FollowerUsers).Error; err != nil {
		return err
	}
	return nil
}
