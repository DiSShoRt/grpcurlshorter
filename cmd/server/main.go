package main

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	urlshort "grpcurlshorter/pkg/mydescription"

	"grpcurlshorter/pkg/urlshorter"

	"log"

	"net"
)

func main()  {
	if err := godotenv.Load(); err != nil {
		log.Fatal("env")
	}

	if err := InitCinfig(); err != nil {
		log.Fatal("error init config", err)
	}
	s := grpc.NewServer()
	srv := &urlshort.GRPCServer{}
	urlshorter.RegisterUrluhorterServer(s,srv)
	l,_ := net.Listen("tcp",":8080")
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}

}


func InitCinfig() error {
		viper.AddConfigPath("configs")
		viper.SetConfigName("config")

		return viper.ReadInConfig()
}
