package model

import "time"

type BaseModel struct {
    DeletedAt *time.Time `sql:"index"`
} 

