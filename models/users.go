package models

import (
	"time"

	"gorm.io/gorm"
)

type UserType int

const (
	Admin UserType = iota
	Manager
	Moderator
	Customer
	DeliveryMan
)

type User struct {
	gorm.Model
	Username       string      `json:"username"`
	Password       string      `json:"password"`
	Email          string      `gorm:"unique" json:"email"`
	SuscribedSince *time.Time  `json:"suscribed_since"`
	SuscribtionID  *uint       `json:"suscribtion_id"`
	Suscribtion    Suscribtion `gorm:"foreignKey:SuscribtionID"`
	UserType       UserType    `json:"user_type"`

	// Associations
	Adresses   []UserAdress
	Phones     []UserPhone
	Cards      []UserCard
	Identities []UserIdentity
}

func (u *User) TableName() string {
	return "users.identifier"
}

type UserIdentity struct {
	gorm.Model
	UserID   uint   `json:"user_id"`
	User     User   `gorm:"foreignKey:UserID"`
	Provider string `json:"provider"`
	UID      string `json:"uid"`
}

func (u *UserIdentity) TableName() string {
	return "users.identities"
}

type UserAdress struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	User    User   `gorm:"foreignKey:UserID"`
	Street  string `json:"street"`
	City    string `json:"city"`
	Zip     string `json:"zip"`
	Country string `json:"country"`
}

func (u *UserAdress) TableName() string {
	return "users.adresses"
}

type UserPhone struct {
	gorm.Model
	UserID      uint   `json:"user_id"`
	User        User   `gorm:"foreignKey:UserID"`
	Phone       string `json:"phone"`
	CountryCode string `json:"country_code"`
}

func (u *UserPhone) TableName() string {
	return "users.phones"
}

type UserCard struct {
	gorm.Model
	UserID       uint   `json:"user_id"`
	User         User   `gorm:"foreignKey:UserID"`
	CardName     string `json:"card_name"`
	CardLast4    string `json:"card_last4"`
	CardBrand    string `json:"card_brand"`
	CardExpMonth string `json:"card_exp_month"`
	CardExpYear  string `json:"card_exp_year"`
	CardStripeID string `json:"card_stripe_id"`
}

func (u *UserCard) TableName() string {
	return "users.cards"
}
