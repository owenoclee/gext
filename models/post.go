package models

type Post struct {
	Id        uint32
	ReplyTo   uint32
	Board     string
	Subject   string
	Body      string
	CreatedAt uint32
}
