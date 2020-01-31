package database

import (
    "github.com/jinzhu/gorm"
    "github.com/satori/go.uuid"
)

type User struct {
    gorm.Model
    UUID    uuid.UUID `gorm:"primary_key"`
    Discord string    `gorm:"unique"`
    Faceit  string    `gorm:"unique"`
    Bungie  string    `gorm:"unique"`
    Banned  bool      `gorm:"default:false"`
}

func (u *User) TableName() string {
    return "profiles"
}

func (u *User) BeforeCreate(scope *gorm.Scope) error {
    uuid := uuid.NewV4()

    return scope.SetColumn("UUID", uuid)
}
