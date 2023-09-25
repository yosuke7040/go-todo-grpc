package models

const (
	TodoStatusUnspecified = iota
	TodoStatusDoing
	TodoStatusDone
)

type Todo struct {
	Id     int32
	Title  string
	Status int32
}
