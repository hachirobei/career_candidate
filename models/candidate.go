package models

import (
    "gorm.io/gorm"
    "time"
)

type Candidate struct {
    ID           int       `json:"id" gorm:"primaryKey"`
    FullName     string    `json:"full_name" gorm:"type:text;not null"`
    Email        string    `json:"email" gorm:"type:text;not null;unique"`
    Phone        string    `json:"phone" gorm:"type:text;not null"`
    Address1     string    `json:"address_1" gorm:"type:text"`
    Address2     string    `json:"address_2" gorm:"type:text"`
    Address3     string    `json:"address_3" gorm:"type:text"` 
    Postcode     string    `json:"postcode" gorm:"type:text"`
    City         string    `json:"city" gorm:"type:text"`
    State        string    `json:"state" gorm:"type:text"` 
    Country      string    `json:"country" gorm:"type:text"` 
    Verified     int       `json:"verified" gorm:"type:int;default:1"`
    ReasonReject string    `json:"reason_reject" gorm:"type:text"`
    Status       string    `json:"status" gorm:"type:text;not null;default:1"`
    Base
}
