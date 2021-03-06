package main

import (
	
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	urlshort "grpcurlshorter/pkg/mydescription"

	"grpcurlshorter/pkg/urlshorter"

	"log"

	"net"
)

func main()  {

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
