package database

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type DBClient struct {
    Username string
    Password string
    Host     string
}

func (c *DBClient) Connect() (error, *gorm.DB) {
    db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@/(%s)charset=utf8mb4&parseTime=True&loc=Local",c.Username, c.Password, c.Host))
    return err, db
}

func New(username, password, host string) *DBClient {
    return &DBClient{
        username,
        password,
        host,
    }
}

func (c *DBClient) Init() error {
    err, db := c.Connect()
    if err != nil {
        return err
    }

    db.AutoMigrate(&User{})
    return nil
}
