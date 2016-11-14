package data_access_layer

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Conversation struct {
	gorm.Model
	ConversationID int    `gorm:"primary_key;not null"`
	Name           string `gorm:"type:varchar(40)"`
}

type Message struct {
	//gorm.Model
	MessageID      int       `gorm:"primary_key;AUTO_INCREMENT"`
	UserFk         int       `gorm:"not null"`
	ConversationFk int       `gorm:"not null"`
	Text           string    `gorm:"type:varchar(255)"`
	Time           time.Time //`gorm:""`
}

type UserConversation struct {
	//gorm.Model
	UserID         int `gorm:"primary_key;AUTO_INCREMENT;not null"`
	ConversationId int `gorm:"primary_key;AUTO_INCREMENT;not null"`
}