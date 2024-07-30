package controllers

type LoginRequest struct {
	ChannelID string `json:"ChannelID" binding:"required"`
	IPAddress string `json:"IPAddress" binding:"required"`
}
