package models

type Role struct {
    ID   uint   `gorm:"primary_key"`
    Name string `gorm:"type:varchar(50);not null"`
    Description string `gorm:"type:text;null"`
}
