package models

import (
    "gorm.io/gorm"
    "time"
)

type Role int 

const (
    RoleUser Role = iota
    Roledmin
)

type Users struct {
    ID          int       `json:"id" gorm:"primaryKey"`
    FullName    string    `json:"full_name" gorm:"type:text;not null"`
    Email       string    `json:"email" gorm:"type:text;not null;unique"`
    Phone       string    `json:"phone" gorm:"type:text;not null"`
    Username    string    `json:"username" gorm:"type:text;not null;unique"`
    Password    string    `json:"password" gorm:"type:text;not null"`
    Role        Role
    Status      string `json:"status" gorm:"type:text;not null;default:1"`
    Base
}
