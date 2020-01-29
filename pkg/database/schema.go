package database

import (
    "github.com/jinzhu/gorm"
    "github.com/satori/go.uuid"
)

type User struct {
    gorm.Model
    ID      string `gorm:"type:uuid;primary_key"`
    Discord string `gorm:"unique"`
    Faceit  string `gorm:"unique"`
    Bungie  string `gorm:"unique"`
    Banned  bool   `gorm:"default:false"`
}

func (User) TableName() string {
    return "profiles"
}

func (User) BeforeCreate(scope *gorm.Scope) error {
    uuid, err := uuid.NewV4()
    if err != nil {
        return err
    }

    return scope.SetColumn("ID", uuid)
}
