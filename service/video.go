package service

import "douyin/repository"

// ==============================================获取最新投稿的20个视频===============================================、

func GetLatestVideo(latestTime int64, token string) error {
	if err := repository.GetLatestVideo(latestTime, token); err != nil {
		return err
	}
	return nil
}

// ==============================================获取用户发布的视频列表===============================================

func PublishList(userId int64, token string) error {
	if err := repository.GetUserVideoList(userId, token); err != nil {
		return err
	}
	return nil
}

// ===================================================发布视频=====================================================

func PublishVideo(title, token string) error {
	if err := repository.PublishVideo(title, token); err != nil {
		return err
	}
	return nil
}
