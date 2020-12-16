package model

import (
	"time"
)

type Paper struct {
	PaperID          string `gorm:"type:varchar(10);primary_key;" json:"paper_id"`
	Title            string `gorm:"type:varchar(400);not null;" json:"title"`
	PaperPublishYear string `gorm:"type:varchar(5)" json:"paper_publish_year"`
	ConferenceID     string `gorm:"type:varchar(10)" json:"conference_id"`
}

type Affiliation struct {
	AffiliationID   string `gorm:"type:varchar(10);primary_key;" json:"affiliation_id"`
	AffiliationName string `gorm:"type:varchar(150)" json:"affiliation_name"`
}

type Author struct {
	AuthorID   string `gorm:"type:varchar(10);primary_key;" json:"author_id"`
	AuthorName string `gorm:"type:varchar(100)" json:"author_name"`
}

type Conference struct {
	ConferenceID   string `gorm:"type:varchar(10);primary_key;" json:"conference_id"`
	ConferenceName string `gorm:"type:varchar(10)" json:"conference_name"`
}

type PaperAuthorAffiliation struct {
	PaperID        string `gorm:"type:varchar(10);primary_key;" json:"paper_id"`
	AuthorID       string `gorm:"type:varchar(10)" json:"author_id"`
	AffiliationID  string `gorm:"type:varchar(10)" json:"affiliation_id"`
	AuthorSequence string `gorm:"type:varchar(3);primary_key;" json:"author_sequence"`
}

type PaperReference struct {
	PaperID     string `gorm:"type:varchar(10);primary_key" json:"paper_id"`
	ReferenceID string `gorm:"type:varchar(10);primary_key" json:"reference_id"`
}

type Comment struct {
	//gorm.Model
	CommentID   uint64    `gorm:"primary_key;" json:"comment_id"`
	Username    string    `gorm:"size:15" json:"username"`
	PaperID     string    `gorm:"size:10" json:"paper_id"`
	CommentTime time.Time `json:"comment_time"`
	Content     string    `gorm:"size:255" json:"content"`
	OnTop       bool      `gorm:"default:false" json:"on_top"`
}
