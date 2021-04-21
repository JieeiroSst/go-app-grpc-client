package domain

import (
	"context"
)

type User struct {
	ID int 				`gorm:"primaryKey,autoIncrement"`
	Name string 
	Address string 
	Profile Profile     `gorm:"foreignkey:UserId;references:ID"`

}

type Profile struct {
	ID int 				`gorm:"primaryKey,autoIncrement"`
	Name string 
	Email string 
	UserId int 
}

type Service interface {
	UpdateEmail(context.Context,int,Profile) (bool,string)
	CreateEmail(context.Context,Profile) (bool,string)
	DeleteEmail(context.Context,int) (bool,string)

	UserAll() (users []User, err error)
	UserById(int) (user User,err error)
	CreateUser(User) (err error)
	UpdateUser(int,User) (err error)
	DeleteUser(int) (err error)
	ProfileAll() (profiles []Profile,err error)
	ProfileById(int) (profile Profile,err error)
}

type Repository interface{
	UserAll() (users []User, err error)
	UserById(int) (user User,err error)
	CreateUser(User) (err error)
	UpdateUser(int,User) (err error)
	DeleteUser(int) (err error)

	ProfileAll() (profiles []Profile,err error)
	ProfileById(int) (profile Profile,err error)
	CreateProfile(Profile) (err error)
	UpdateProfile(int,Profile) (err error)
	DeleteProfile(int) (err error)
}