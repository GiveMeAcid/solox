package data_access_layer

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	ID        int    `gorm:"primary_key;AUTO_INCREMENT"`
	Login     string `gorm:"type:varchar(40);not null;unique"`
	FirstName string `gorm:"type:varchar(30)"`
	LastName  string `gorm:"type:varchar(50)"`
	Age       int                      //`gorm:""`
	Email     string `gorm:"type:varchar(60);not null;unique"`
	PhotoPath string `gorm:"type:varchar(255)"`
	Password  string `gorm:"type:varchar(32);not null"`
	Sex       Sex    `gorm:"type:sex"` //sex type
}

type UserSettings struct {
	gorm.Model
	User               User `gorm:"ForeignKey:UserSettingsID;AssociationForeignKey:ID"`
	UserSettingsID     int  `gorm:"primary_key;AUTO_INCREMENT"`
	NightModeOn        bool `gorm:"not null"`
	Visibility         bool `gorm:"not null"`
	NotificationModeOn bool `gorm:"not null"`
}

type UserFilters struct {
	gorm.Model
	User          User    `gorm:"ForeignKey:UserFiltersID;AssociationForeignKey:ID"`
	UserFiltersID int     `gorm:"primary_key;AUTO_INCREMENT"`
	Age           int                       //`gorm:""`
	Sex           Sex     `gorm:"type:sex"` //sex type
	SearchRadius  int     `gorm:"not null;size:5"`
}

//type Sex func() bool
type Sex string

const (
	male Sex = "male"
	female Sex = "female"
)




//type Sex int
//
//const (
//	Male Sex = 1 + iota
//	Female Sex = iota
//)
//
//func (s Sex) Gender() string {
//	switch s {
//	case Female:
//		return "female"
//	case Male:
//		return "male"
//	}
//	return
//}

//func (uf UserFilters) Age() int {
//
//}