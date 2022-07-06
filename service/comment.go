package service

import "douyin/repository"

// AddComment 添加评论
func AddComment(userId, videoId int64, token, commentText string) error {
	if err := repository.AddComment(userId, videoId, token, commentText); err != nil {
		return err
	}
	return nil
}

// DeleteComment 删除评论
func DeleteComment(deleteId, videoID int64, token string) error {
	if err := repository.DeleteComment(deleteId, videoID, token); err != nil {
		return err
	}
	return nil
}

// GetCommentList 获取评论列表
func GetCommentList(videoId int64, token string) error {
	if err := repository.GetCommentList(videoId, token); err != nil {
		return err
	}
	return nil
}
