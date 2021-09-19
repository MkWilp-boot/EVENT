package structs

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name     string        `json:"name" gorm:"<-"`
	Image    string        `json:"image" gorm:"<-"`
	When     time.Time     `json:"when" gorm:"<-"`
	Where    string        `json:"where" gorm:"<-"`
	Duration time.Duration `json:"duration" gorm:"<-"`
}
