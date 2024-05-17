package domain

type State int

const (
	StateWaitingVideo State = iota + 1
	StateWaitingVideoDescription
)
