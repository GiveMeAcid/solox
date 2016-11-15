package data_access_layer

import (
	_ "github.com/jinzhu/gorm"
)

type Sex string

const (
	male Sex = "male"
	female Sex = "female"
)

type User struct {
	// gorm.Model
	ID        int    `gorm:"primary_key"`
	Login     string `gorm:"type:varchar(40);not null;unique"`
	FirstName string `gorm:"type:varchar(30)"`
	LastName  string `gorm:"type:varchar(50)"`
	Age       int     
	Email     string `gorm:"type:varchar(60);not null;unique"`
	PhotoPath string `gorm:"type:varchar(255)"`
	Password  string `gorm:"type:varchar(32);not null"`
	Sex       Sex    ``
}

type UserSettings struct {
	User               User `gorm:"ForeignKey:UserSettingsID;AssociationForeignKey:ID"`
	UserSettingsID     int  `gorm:"primary_key"`
	NightModeOn        bool `gorm:"not null"`
	Visibility         bool `gorm:"not null"`
	NotificationModeOn bool `gorm:"not null"`
}

type UserFilters struct {
	User          User    `gorm:"ForeignKey:UserFiltersID;AssociationForeignKey:ID"`
	UserFiltersID int     `gorm:"primary_key;AUTO_INCREMENT"`
	Age           int      
	Sex           Sex     
	SearchRadius  int     `gorm:"not null;size:5"`
}