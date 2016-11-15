package data_access_layer

import (
	_ "solox/data_access/structs/users"
	"time"

	_ "github.com/jinzhu/gorm"
)

type Conversation struct {
	// gorm.Model
	ConversationID int    `gorm:"primary_key;not null"`
	Name           string `gorm:"type:varchar(40)"`
}

type Messages struct {
	MessageID      int    `gorm:"primary_key;AUTO_INCREMENT"`
	UserFk         int    `gorm:"not null"`
	ConversationFk int    `gorm:"not null"`
	Text           string `gorm:"type:varchar(255)"`
	Time           time.Time
}

type UserConversation struct {
	User           []Conversation `gorm:"many2many:user_conversations;AssosiationForeignKey:ConversationID;ForeignKey:UserID"`
	UserID         int            `gorm:"not null;"`
	ConversationID int            `gorm:"not null;"`
}
