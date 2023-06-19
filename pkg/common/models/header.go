package models

type Header struct {
	UserID   string `header:"user-id" binding:"required"`
	ClientID string `header:"client-id"`
}
