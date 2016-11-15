package main

import (
	_ "database/sql"
	"fmt"
	conv "solox/data_access/structs/conversations"
	evnt "solox/data_access/structs/events"
	user "solox/data_access/structs/users"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=solox2 sslmode=disable password=31780")

	if err != nil {
		fmt.Printf("Database opening error -->%v\n", err)
		panic("Database error")
	}
	defer db.Close()

	db.SingularTable(true)

	db1 := db.AutoMigrate(&user.UserSettings{}, &user.UserFilters{}, &conv.Conversation{}, &conv.Messages{}, &user.UserInfo{},
		evnt.Event{}, evnt.UserEvent{})

	// db.Model(&conv.UserConversation{}).AddForeignKey("conversation_id", "conversation(id)", "RESTRICT", "RESTRICT")
	// db.Model(&conv.UserConversation{}).AddForeignKey("id", "user_info(id)", "RESTRICT", "RESTRICT")
	db.Model(&conv.Messages{}).AddForeignKey("conversation_fk", "conversation(id)", "RESTRICT", "RESTRICT")
	db.Model(&conv.Messages{}).AddForeignKey("user_fk", "user_info(id)", "RESTRICT", "RESTRICT")
	db.Model(evnt.UserEvent{}).AddForeignKey("event_fk", "event(event_id)", "RESTRICT", "RESTRICT")
	db.Model(evnt.UserEvent{}).AddForeignKey("user_fk", "user_info(id)", "RESTRICT", "RESTRICT")
	db.Model(&user.UserSettings{}).AddForeignKey("user_settings_id", "user_info(id)", "RESTRICT", "RESTRICT")
	db.Model(&user.UserFilters{}).AddForeignKey("user_filters_id", "user_info(id)", "RESTRICT", "RESTRICT")
	db.Model(`gorm:"many2many:user_conversations;AssociationForeignKey:UserId;ForeignKey:ConversationId;"`).AddForeignKey("conversation_id", "conversation(id)", "RESTRICT", "RESTRICT")

	//db.Model(&conv.UserConversation{}).AddIndex("user_conversation_index", "conversation_id")
	if db1.Error != nil {
		panic(db1.Error)
	}
}
