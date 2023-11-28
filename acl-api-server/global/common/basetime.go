package common

import "time"

type BaseTimeEntity struct {
	CreatedAt  time.Time `gorm:"<-:create"`
	ModifiedAt time.Time `gorm:"autoUpdateTime"`
}
