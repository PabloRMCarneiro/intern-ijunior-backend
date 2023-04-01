package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Phone    string `gorm:"not null"`
	Role     string `gorm:"not null default:'user'"`
}

func (u *User) TableName() string {
	return "users"
}

type UserRepository struct {
	DB *gorm.DB
}

func (ur *UserRepository) CreateUser(user *User) error {
	return ur.DB.Create(user).Error
}

func (ur *UserRepository) GetUserById(id string) (*User, error) {
	user := &User{}
	err := ur.DB.First(user, id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*User, error) {
	user := &User{}
	err := ur.DB.Where("email = ?", email).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) GetUsers() ([]*User, error) {
	users := []*User{}
	err := ur.DB.First(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepository) DeleteUser(user *User) error {
	return ur.DB.Delete(user).Error
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}
