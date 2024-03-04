package models

type Photo struct {
	Id       int64  `grom:"primaryKey" json:"id"`
	Title    string `gorm:"type:varchar(50)" json:"title"`
	Caption  string `gorm:"type:varchar(100)" json:"caption"`
	PhotoUrl string `gorm:"type:varchar(200)" json:"photourl"`
	UserID   int64  `gorm:"type:int(10)" json:"userid"`
}
