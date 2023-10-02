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

type User struct {
	Id    int32
	Name  string
	Email string
}
