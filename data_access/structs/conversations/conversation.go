package data_access_layer

import (
	user "solox/data_access/structs/users"
	"time"

	_ "github.com/jinzhu/gorm"
)

type Conversation struct {
	//gorm.Model
	ConversationId uint `gorm:"primary_key:true;index:idx_conversation_id;auto_increment:true;column:id"`
	Name           string
	Users          []user.UserInfo `gorm:"many2many:user_conversations;AssociationForeignKey:UserId;ForeignKey:ConversationId;"`
}

type Messages struct {
	MessageID      int    `gorm:"primary_key;AUTO_INCREMENT;"`
	UserFk         int    `gorm:"not null"`
	ConversationFk int    `gorm:"not null;index:idx_conversation_fk"`
	Text           string `gorm:"type:varchar(255)"`
	Time           time.Time
}
