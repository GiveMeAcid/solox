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

	db1 := db.AutoMigrate(&user.UserSettings{}, &user.UserFilters{}, &conv.Conversation{}, &conv.Messages{}, &conv.UserConversation{}, &user.User{},
		evnt.Events{}, evnt.UserEvent{})

	db.Model(&conv.Conversation{}).AddIndex("conversation_index", "conversation_id")
	db.Model(&evnt.Events{}).AddIndex("event_index", "event_id")
	db.Model(&conv.Messages{}).AddIndex("message_index", "conversation_fk")
	db.Model(&conv.UserConversation{}).AddIndex("user_conversation_index", "conversation_id")
	db.Model(&evnt.UserEvent{}).AddIndex("user_event_index", "event_id")
	db.Model(&user.UserFilters{}).AddIndex("user_filters_index", "user_filters_id")
	db.Model(&user.User{}).AddIndex("user_index", "user_id")
	db.Model(&user.UserSettings{}).AddIndex("user_settings_index", "user_settings_id")

	if db1.Error != nil {
		panic(db1.Error)
	}
}
