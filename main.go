package main

import (
	_ "database/sql"
	"fmt"
	sc "solox/data_access/structs/conversations"
	se "solox/data_access/structs/events"
	su "solox/data_access/structs/users"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=solox2 sslmode=disable password=31780")

	if err != nil {
		fmt.Printf("Database opening error -->%v\n", err)
		//panic("Database error")
	}
	defer db.Close()

	db1 := db.AutoMigrate(&su.UserSettings{}, &su.UserFilters{}, &sc.Conversation{}, &sc.Messages{}, &sc.UserConversation{}, &su.User{},
		se.Events{}, se.UserEvent{})

	db.Model(&sc.Conversation{}).AddIndex("conversation_index", "conversation_id")
	db.Model(&se.Events{}).AddIndex("event_index", "event_id")
	db.Model(&sc.Messages{}).AddIndex("message_index", "conversation_fk")
	db.Model(&sc.UserConversation{}).AddIndex("user_conversation_index", "conversation_id")
	db.Model(&se.UserEvent{}).AddIndex("user_event_index", "event_id")
	db.Model(&su.UserFilters{}).AddIndex("user_filters_index", "user_filters_id")
	db.Model(&su.User{}).AddIndex("user_index", "user_id")
	db.Model(&su.UserSettings{}).AddIndex("user_settings_index", "user_settings_id")

	if db1.Error != nil {
		panic(db1.Error)
	}
}
