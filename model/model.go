package model

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Id        int64      `json:"id" gorm:"column:id;primaryKey"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;default:NULL"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

type Rka struct {
	gorm.Model
	RkaId     int64      `json:"rka_id" gorm:"column:id;primaryKey"`
	NameRka   string     `json:"name_rka"`
	TargetRka int32      `json:"target_rka"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;default:NULL"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

type RKAUsers struct {
	gorm.Model
	RkaUserId int64      `json:"rka_user_id" gorm:"column:id;primaryKey"`
	RkaId     int64      `json:"rka_id" gorm:"column:rka_id;references:rka_id"`
	UserId    int64      `json:"user_id" gorm:"column:user_id;references:id"`
	Visited   int64      `json:"visited"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;default:NULL"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

type AccumulationRKA struct {
	Name    string `json:"name" gorm:"column:name"`
	RkaName string `json:"rka_name" gorm:"column:rka_name"`
	Visited int64  `json:"visited" gorm:"column:visited"`
	Target  int64  `json:"target" gorm:"column:target"`
	Month   int32  `json:"month" gorm:"column:month"`
	Year    int32  `json:"year" gorm:"column:year"`
}
