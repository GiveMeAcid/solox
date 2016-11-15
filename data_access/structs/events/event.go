package data_access_layer

import _ "solox/data_access/structs/users"

type EventType string

const (
	single   EventType = "single"
	multiple EventType = "multiple"
)

type g_m_s string

const ()

type Event struct {
	// gorm.Model
	EventID int `gorm:"primary_key;not null;index:idx_event_id"`
	Type    EventType
	Place   g_m_s `gorm:"not null"`
}

type UserEvent struct {
	UserInfo []Event `gorm:"ForeignKey:UserId;AssociationForeignKey:UserFk"`
	Events   []Event `gorm:"ForeignKey:EventId;AssociationForeignKey:EventFk"`
	EventFk  uint    `gorm:"primary_key:true;index:idx_event_fk"`
	UserFk   uint    `gorm:"primary_key:true"`
	Result   bool    `gorm:"not null"`
}

// type UserEvent struct {
// 	// gorm.Model
// 	User    []Events `gorm:"AssosiationForeignKey:EventID;ForeignKey:UserID"`
// 	UserID  int      //`gorm:""`
// 	EventID int      `gorm:"not null"`
// 	Result  bool
// }
