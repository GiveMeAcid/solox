package data_access_layer

import (
	_ "github.com/jinzhu/gorm"
)

type Sex string

const (
	male   Sex = "male"
	female Sex = "female"
)

type AgeFilter string

const (
	first  AgeFilter = "18-25"
	second AgeFilter = "26-34"
	third  AgeFilter = "35-45"
	forth  AgeFilter = "45+"
)

type UserInfo struct {
	// gorm.Model
	UserId       int    `gorm:"primary_key;index:idx_id;column:id"`
	Login        string `gorm:"type:varchar(40);not null;unique"`
	FirstName    string `gorm:"type:varchar(30)"`
	LastName     string `gorm:"type:varchar(50)"`
	Age          int
	Email        string `gorm:"type:varchar(60);not null;unique"`
	PhotoPath    string `gorm:"type:varchar(255)"`
	Password     string `gorm:"type:varchar(32);not null"`
	Sex          Sex
}

type UserSettings struct {
	UserInfo           UserInfo `gorm:"ForeignKey:UserSettingsID;AssociationForeignKey:ID"`
	UserSettingsID     int      `gorm:"primary_key;index:idx_user_settings_id"`
	NightModeOn        bool     `gorm:"not null"`
	Visibility         bool     `gorm:"not null"`
	NotificationModeOn bool     `gorm:"not null"`
}

type UserFilters struct {
	// User          User `gorm:"ForeignKey:UserFiltersID;AssociationForeignKey:ID"`
	UserFiltersID int `gorm:"primary_key;AUTO_INCREMENT;index:idx_user_filters_id"`
	Age           AgeFilter
	Sex           Sex
	SearchRadius  int `gorm:"not null;size:5"`
}
