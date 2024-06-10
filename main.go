package main

import (
	"fmt"
	"log"
	"github.com/dilshodforever/restaurant-auth/api"
	"github.com/dilshodforever/restaurant-auth/api/handler"
	pb "github.com/dilshodforever/restaurant-auth/genprotos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	UserConn, err := grpc.NewClient(fmt.Sprintf("localhost%s", ":8085"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!=nil{
		log.Fatal("Error while NEwclient: ", err.Error())
	}
	defer UserConn.Close()

	us := pb.NewUserServiceClient(UserConn)

	h := handler.NewHandler(us)
	r := api.NewGin(h)

	fmt.Println("Server started on port:8080")
	err = r.Run()
	if err!=nil{
		log.Fatal("Error while Run: ", err.Error())
	}
}
