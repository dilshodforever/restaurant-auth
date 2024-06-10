package handler

import pb "github.com/dilshodforever/restaurant-auth/genprotos"

type Handler struct {
	User   pb.UserServiceClient
}

func NewHandler(us pb.UserServiceClient) *Handler {
	return &Handler{us}
}
