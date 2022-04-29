package repository

import (
	"project-stokku/entity"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userModel struct {
	DB *gorm.DB
}

func NewUserModel(db *gorm.DB) *userModel {
	return &userModel{
		DB: db,
	}
}

func (u *userModel) CheckDuplicate(email string) bool {
	var user entity.User

	u.DB.Unscoped().Where("email = ?", email).Find(&user)
	u.DB.Where("email = ?", email).Find(&user)

	if user.Email == email {
		return true
	} 
	
	return false
}

func (u *userModel) Create(user *entity.User) error {
	user.Name = strings.Title(strings.ToLower(user.Name))
	user.Email = strings.ToLower(user.Email)
	
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	}

    user.Password = string(bytes)
	
	if err := u.DB.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func (u *userModel) Login(email, password string) (bool, entity.User) {
	user := entity.User{}

	u.DB.Where("email = ? ", email).Find(&user)

	if CheckPasswordHash(password, user.Password) {
		return true, user
	}
	
	return false, user
}