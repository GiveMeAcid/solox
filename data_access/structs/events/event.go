package data_access_layer

import "github.com/jinzhu/gorm"

type Event struct {
	gorm.Model
	EventID int   `gorm:"primary_key;not null"`
	//Type 	type  `gorm:"type:type"`
	Place   g_m_s `gorm:"not null"`
}

type UserEvent struct {
	gorm.Model

	UserID  []int  `gorm:"primary_key;AUTO_INCREMENT"`
	EventId []int  `gorm:"primary_key;not null"`
	Result  bool //`gorm:""`
}

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

type Sex string

const (
	male Sex = "male"
	female Sex = "female"
)

type g_m_s int
