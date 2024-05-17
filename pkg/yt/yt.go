package yt

import (
	"io"

	ytl "github.com/kkdai/youtube/v2"
)

type Youtube struct {
	yt *ytl.Client
}

func NewYoutube() *Youtube {
	ytInstance := ytl.Client{
		MaxRoutines: 10,
	}
	return &Youtube{
		yt: &ytInstance,
	}
}

func (y *Youtube) GetVideoInfo(url string) (*ytl.Video, error) {
	return y.yt.GetVideo(url)
}

func (y *Youtube) DownloadVideo(video *ytl.Video, format *ytl.Format) (io.ReadCloser, int64, error) {
	return y.yt.GetStream(video, format)
}
