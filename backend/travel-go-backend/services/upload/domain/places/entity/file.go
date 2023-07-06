package entity

import "time"

type File struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Size        uint64    `json:"size"`
	Type        string    `json:"type"`
	ContentType string    `json:"content_type"`
	Path        string    `json:"path"`
	FullPath    string    `json:"full_path" gorm:"-"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
}
