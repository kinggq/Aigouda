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

func (s *UserService) Register(username, password string) error {
	// 检查用户名是否已存在
	var existingUser model.User
	if err := repository.DB.Where("username = ?", username).First(&existingUser).Error; err == nil {
		return errors.New("用户名已存在")
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := model.User{
		Username: username,
		Password: string(hashedPassword),
		Role:     "admin", // 默认注册为管理员
	}

	return repository.DB.Create(&user).Error
}

func (s *UserService) Login(username, password string) (string, *model.User, error) {
	var user model.User
	if err := repository.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return "", nil, errors.New("用户不存在")
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", nil, errors.New("密码错误")
	}

	// 生成token
	token, err := utils.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		return "", nil, err
	}

	return token, &user, nil
}

func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	var user model.User
	if err := repository.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
} 
