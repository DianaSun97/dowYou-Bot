package service

import (
	ytl "github.com/kkdai/youtube/v2"
	"io"
)

type Repository interface {
	GetVideo(url string) (*ytl.Video, error)
	DownloadVideo(video *ytl.Video, format *ytl.Format) (io.ReadCloser, int64, error)
}
