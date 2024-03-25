package models

import (
	"time"
)

type ImgInfo struct {
	FileName string
	Size     uint32
	CreateDt time.Time
	UpdateDt time.Time
}
