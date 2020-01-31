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
    DBName   string
}

func (c *DBClient) Connect() (*gorm.DB, error) {
    connectionString := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.Username, c.Password, c.Host, c.DBName)
    db, err := gorm.Open("mysql", connectionString)
    return db, err
}

func New(username, password, host, dbname string) *DBClient {
    return &DBClient{
        username,
        password,
        host,
        dbname,
    }
}

func (c *DBClient) Init() error {
    db, err := c.Connect()
    if err != nil {
        return err
    }

    defer db.Close()

    db.AutoMigrate(&User{})

    return nil
}
