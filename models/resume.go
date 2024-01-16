package models

import (
    "gorm.io/gorm"
    "time"
)

type Resume struct {
    ID          int       `json:"id" gorm:"primaryKey"`
    Candidate_id    string    `json:"full_name" gorm:"type:text;not null"`
    Attachment_Path       string    `json:"email" gorm:"type:text;not null;unique"`
    Status          string `json:"status" gorm:"type:text;not null;default:1"`
    Base
}
