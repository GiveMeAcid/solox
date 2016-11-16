package data_access_layer

import (
	driver "database/sql/driver"
)

type EventType string

const (
	single   EventType = "single"
	multiple EventType = "multiple"
)

type Point struct {
	X, Y int
}

func Pt(X, Y int) Point {
	return Point{X, Y}
}

type Event struct {
	EventID int `gorm:"primary_key;not null;index:idx_event_id"`
	Type    EventType
	Place   Point `gorm:"not null"`
}

type UserEvent struct {
	UserInfo []Event `gorm:"ForeignKey:UserId;AssociationForeignKey:UserFk"`
	Events   []Event `gorm:"ForeignKey:EventId;AssociationForeignKey:EventFk"`
	EventFk  uint    `gorm:"primary_key:true;index:idx_event_fk"`
	UserFk   uint    `gorm:"primary_key:true"`
	Result   bool    `gorm:"not null"`
}

func (u Point) Scan(value interface{}) error {
	u = Point(value.(Point))
	return nil
}
func (u Point) Value() (driver.Value, error) {
	return Point(u), nil
}
