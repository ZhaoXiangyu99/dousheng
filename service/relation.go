package service

import "douyin/repository"

// AddRelation 添加关注
func AddRelation(toUserId int64, token string) error {
	if err := repository.AddRelation(toUserId, token); err != nil {
		return err
	}
	return nil
}

// DelRelation 取消关注
func DelRelation(toUserId int64, token string) error {
	if err := repository.DelRelation(toUserId, token); err != nil {
		return err
	}
	return nil
}

// FollowList 关注列表
func FollowList(userId int64, token string) error {
	if err := repository.FollowList(userId, token); err != nil {
		return err
	}
	return nil
}

// FollowerList 粉丝列表
func FollowerList(toUserId int64, token string) error {
	if err := repository.FollowerList(toUserId, token); err != nil {
		return err
	}
	return nil
}
