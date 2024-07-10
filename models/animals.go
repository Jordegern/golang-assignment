package models

import (
	"gorm.io/gorm"
	"time"
)

type Animal struct {
	gorm.Model
	Birthday time.Time
	Gender   string
	Status   string
	Race     string
	Weight   int
}
