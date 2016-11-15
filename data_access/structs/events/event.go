package data_access_layer

import _ "solox/data_access/structs/users"

type Type int

const (
	couple   Type = 2
	multiple Type = 5
)

type Events struct {
	// gorm.Model
	EventID int `gorm:"primary_key;not null"`
	Type    Type
	//Place   g_m_s `gorm:"not null"`
}

type UserEvent struct {
	// gorm.Model
	User   []Events `gorm:"AssosiationForeignKey:EventID;ForeignKey:UserID"`
	UserID  int		`gorm:""`
	EventID int 	`gorm:"not null"`
	Result  bool
}
