package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpcurlshorter/pkg/urlshorter"
	"log"
)

func main() {

	var x string
	fmt.Println("please, input url")

	fmt.Scan(&x)

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	c := urlshorter.NewUrluhorterClient(conn)

	r, err := c.Create(context.Background(), &urlshorter.LongUrl{Long: x})
	if err != nil {
		log.Printf("error %s in Create ", err)

	} else {
		res, _ := c.Get(context.Background(), &urlshorter.ShortUrl{Short: r.Short})
		fmt.Println("Get", res)
	}

	fmt.Println("Create", r)
}
