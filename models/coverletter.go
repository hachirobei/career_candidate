package models

import (
    "gorm.io/gorm"
    "time"
)

type Coverletter struct {
    ID              int    `json:"id" gorm:"primaryKey"`
    CandidateID     string `json:"candidate_id" gorm:"type:text;not null"`
    AttachmentPath  string `json:"attachment_path" gorm:"type:text;not null;unique"`
    Status          string `json:"status" gorm:"type:text;not null;default:1"`
    Base
}
