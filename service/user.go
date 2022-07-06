package service

import (
	"douyin/repository"
)

// ==================================================用户注册功能模块=====================================================

func Register(name, password string) error {
	if err := repository.CreateUser(name, password); err != nil {
		return err
	}
	return nil
}

// ================================================用户登录功能模块=======================================================

func Login(name, password string) error {
	if err := repository.GetUser(name, password); err != nil {
		return err
	}
	return nil
}

// ================================================获取登录用户信息==========================================================

func GetUserInfo(uid int64) error {
	if err := repository.GetLoginUser(uid); err != nil {
		return err
	}
	return nil
}
