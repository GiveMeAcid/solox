package users

import (
	driver "database/sql/driver"

	_ "github.com/jinzhu/gorm"
	//"github.com/jinzhu/gorm"
	//"fmt"
	//"os/user"
	_ "database/sql"
)

type Gender string

const (
	male   Gender = "male"
	female Gender = "female"
)

type AgeFilter string

const (
	first  AgeFilter = "18-25"
	second AgeFilter = "26-34"
	third  AgeFilter = "35-45"
	forth  AgeFilter = "45+"
)

//
//type UserInfo struct {
//	UserId    int    `gorm:"primary_key;index:idx_id;column:id"`
//	Login     string `gorm:"type:varchar(40);not null;index:idx_login"`
//	FirstName string `gorm:"type:varchar(30)"`
//	LastName  string `gorm:"type:varchar(50)"`
//	Age       int
//	Email     string `gorm:"type:varchar(60);not null"`
//	PhotoPath string `gorm:"type:varchar(255)"`
//	Password  string `gorm:"type:varchar(32);not null"`
//	Sex       Gender
//}

type UserSettings struct {
	UserInfo           UserInfo `gorm:"ForeignKey:UserSettingsID;AssociationForeignKey:ID"`
	UserSettingsID     int      `gorm:"primary_key;index:idx_user_settings_id"`
	NightModeOn        bool     `gorm:"not null"`
	Visibility         bool     `gorm:"not null"`
	NotificationModeOn bool     `gorm:"not null"`
}

type UserFilters struct {
	UserFiltersID int `gorm:"primary_key;AUTO_INCREMENT;index:idx_user_filters_id"`
	Age           AgeFilter
	Sex           Gender
	SearchRadius  int `gorm:"not null;size:5"`
}

//func (u Gender) Scan(value interface{}) error {
//	u = Gender(value.([]byte))
//	return nil
//}
//func (u Gender) Value() (driver.Value, error) {
//	return string(u), nil
//}

//var db *sql.DB

//func GetUser(id int) *UserInfo {
//	user := UserInfo{}
//	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=solox2 sslmode=disable password=31780")
//
//	if err != nil {
//		fmt.Printf("Database opening error -->%v\n", err)
//		panic("Database error")
//	}
//	user = db.Where(&UserInfo{UserId: id}).First(&UserInfo{})
//	//fmt.Println(user)
//	defer db.Close()
//	return &user
//}

type UserInfo struct {
	UserId         uint          `gorm:"primary_key:true;index:idx_user_id;auto_increment:true;column:id" json:"id"`
	Login          string        `gorm:"index:idx_user_login;not null;unique" validate:"nonzero,max=50" json:"login"`
	Password       string        `gorm:"type:varchar(64);not null" validate:"nonzero,max=50" json:"password,omitempty"`
	EMail          string        `gorm:"type:varchar(255)" validate:"nonzero,max=255,regexp=^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$" json:"eMail"`
	Description    string        `gorm:"type:varchar(255)" json:"description"`
	Gender         Gender        `json:"gender"`
	Age            int           `json:"age"`
	NickName       string        `gorm:"index:idx_user_login;not null;unique" validate:"max=50" json:"nickName"`
	UserSettingsFk uint          `gorm:"not null" json:"-"`
	UserFiltersFk  uint          `gorm:"not null" json:"-"`
	UserSettings   *UserSettings `gorm:"ForeignKey:UserSettingsId;AssociationForeignKey:UserSettingsFk" json:"userSettings,omitempty"`
	UserFilters    *UserFilters  `gorm:"ForeignKey:UserFiltersId;AssociationForeignKey:UserFiltersFk" json:"userFilters,omitempty"`
	PhotoPath      string        `gorm:"not null" json:"photoPath" json:"-"`
	PhoneFk        uint          `gorm:"not null" json:"-"`
	// Phone          *Phone        `gorm:"not null;ForeignKey:PhoneId;AssociationForeignKey:PhoneFk" validate:"nonzero" json:"phone,omitempty"`
	Token []byte `sql:"-" json:"token,omitempty"`
}

func (u Gender) Scan(value interface{}) error { u = Gender(value.(string)); return nil }
func (u Gender) Value() (driver.Value, error) { return string(u), nil }

func (user *UserInfo) CheckLogin(login string) (bool, error) {
	users := new([]UserInfo)
	err := services.DB.Where("login = ?", login).Find(&users).Error

	if len(*users) != 0 && err == nil {
		return true, nil
	}

	return false, err
}

func (user *UserInfo) Get(login string) error {
	err := services.DB.Where("login = ?", login).First(&user).Error
	if err == nil {
		user.UserSettings, err = user.UserSettings.GetUserSettigs(user.UserSettingsFk)
		if err == nil {
			user.UserFilters, err = user.UserFilters.GetUserFilters(user.UserFiltersFk)
			if err == nil {
				user.Phone, err = user.Phone.GetPhone(user.PhoneFk)
			}
		}
	}

	return err
}
