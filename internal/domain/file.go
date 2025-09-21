package domain

import (
	"time"
)

type File struct {
	ID            string     `json:"id"`
	OriginalName  string     `json:"original_name"`
	BucketName    string     `json:"bucket_name"`
	Size          int64      `json:"size"`
	UploadTime    time.Time  `json:"upload_time"`
	DownloadCount int64      `json:"download_count"`
	ExpiryTime    *time.Time `json:"expiry_time,omitempty"`
}
