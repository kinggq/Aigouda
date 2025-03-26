package service

import (
	"aigouda/internal/model"
	"aigouda/internal/repository"
	"aigouda/pkg/utils"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

// Login 用户登录
func Login(username, password string) (string, error) {
	// 获取用户信息
	user, err := repository.GetUserByUsername(username)
	if err != nil {
		return "", errors.New("用户不存在")
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("密码错误")
	}

	// 生成token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Register 用户注册
func Register(username, password string) error {
	// 检查用户名是否已存在
	_, err := repository.GetUserByUsername(username)
	if err == nil {
		return errors.New("用户名已存在")
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 创建用户
	user := &model.User{
		Username: username,
		Password: string(hashedPassword),
	}

	return repository.CreateUser(user)
}

// GetProfile 获取用户信息
func GetProfile(userID uint) (*model.User, error) {
	return repository.GetUserByID(userID)
}

// UpdateProfile 更新用户信息
func UpdateProfile(userID uint, username string) error {
	user, err := repository.GetUserByID(userID)
	if err != nil {
		return err
	}

	user.Username = username
	return repository.UpdateUser(user)
}

// UpdatePassword 更新密码
func UpdatePassword(userID uint, oldPassword, newPassword string) error {
	user, err := repository.GetUserByID(userID)
	if err != nil {
		return err
	}

	// 验证旧密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
	if err != nil {
		return errors.New("旧密码错误")
	}

	// 加密新密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return repository.UpdateUser(user)
}

func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	var user model.User
	if err := repository.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
} 
