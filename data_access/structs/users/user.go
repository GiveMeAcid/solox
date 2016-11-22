package users

import (
	driver "database/sql/driver"

	_ "github.com/jinzhu/gorm"
)

type Gender string

const (
	male   Gender = "male"
	female Gender = "female"
)

type AgeFilter string

const (
	first  AgeFilter = "18-25"
	second AgeFilter = "26-34"
	third  AgeFilter = "35-45"
	forth  AgeFilter = "45+"
)

type UserInfo struct {
	UserId    int    `gorm:"primary_key;index:idx_id;column:id"`
	Login     string `gorm:"type:varchar(40);not null;unique;index:idx_login"`
	FirstName string `gorm:"type:varchar(30)"`
	LastName  string `gorm:"type:varchar(50)"`
	Age       int
	Email     string `gorm:"type:varchar(60);not null;unique"`
	PhotoPath string `gorm:"type:varchar(255)"`
	Password  string `gorm:"type:varchar(32);not null"`
	Sex       Gender
}

type UserSettings struct {
	UserInfo           UserInfo `gorm:"ForeignKey:UserSettingsID;AssociationForeignKey:ID"`
	UserSettingsID     int      `gorm:"primary_key;index:idx_user_settings_id"`
	NightModeOn        bool     `gorm:"not null"`
	Visibility         bool     `gorm:"not null"`
	NotificationModeOn bool     `gorm:"not null"`
}

type UserFilters struct {
	UserFiltersID int `gorm:"primary_key;AUTO_INCREMENT;index:idx_user_filters_id"`
	Age           AgeFilter
	Sex           Gender
	SearchRadius  int `gorm:"not null;size:5"`
}

func (u Gender) Scan(value interface{}) error {
	u = Gender(value.([]byte))
	return nil
}
func (u Gender) Value() (driver.Value, error) {
	return string(u), nil
}
