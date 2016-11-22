package main

import (
	_ "database/sql"
	"fmt"
	"log"
	"net/http"

	conv "github.com/user/solox/data_access/structs/conversations"
	evnt "github.com/user/solox/data_access/structs/events"
	user "github.com/user/solox/data_access/structs/users"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

// func handler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/plain")
// 	w.Write([]byte("This is an example server.\n"))
// }

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

	db.Model(&conv.Messages{}).AddForeignKey("conversation_fk", "conversation(id)", "RESTRICT", "RESTRICT")
	db.Model(&conv.Messages{}).AddForeignKey("user_fk", "user_info(id)", "RESTRICT", "RESTRICT")
	db.Model(evnt.UserEvent{}).AddForeignKey("event_fk", "event(event_id)", "RESTRICT", "RESTRICT")
	db.Model(evnt.UserEvent{}).AddForeignKey("user_fk", "user_info(id)", "RESTRICT", "RESTRICT")
	db.Model(&user.UserSettings{}).AddForeignKey("user_settings_id", "user_info(id)", "RESTRICT", "RESTRICT")
	db.Model(&user.UserFilters{}).AddForeignKey("user_filters_id", "user_info(id)", "RESTRICT", "RESTRICT")

	db.Table("user_conversations").AddForeignKey("conversation_id", "conversation(id)", "RESTRICT", "RESTRICT")
	db.Table("user_conversations").AddForeignKey("user_info_id", "user_info(id)", "RESTRICT", "RESTRICT")

	db.Table("user_conversations").AddIndex("idx_user_conversation", "conversation_id")

	if db1.Error != nil {
		panic(db1.Error)
	}
//	
	http.HandleFunc("/", user.UserIndex)
	log.Printf("About to listen on 10443. Go to https://127.0.0.1:10443/")
	err = http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
	http.ListenAndServe(":10443", nil)
	log.Fatal(err)
}
