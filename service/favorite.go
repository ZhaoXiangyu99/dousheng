package service

import "douyin/repository"

// AddFavorite 点赞
func AddFavorite(videoId int64, token string) error {
	if err := repository.AddFavorite(videoId, token); err != nil {
		return err
	}
	return nil
}

// DelFavorite 取消点赞
func DelFavorite(videoId int64, token string) error {
	if err := repository.DelFavorite(videoId, token); err != nil {
		return err
	}
	return nil
}

// GetFavoriteList 获取点赞列表
func GetFavoriteList(userId int64, token string) error {
	if err := repository.GetFavoriteList(userId, token); err != nil {
		return err
	}
	return nil
}
