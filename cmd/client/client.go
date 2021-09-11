package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpcurlshorter/pkg/urlshorter"
)

func main() {

	var x string
	fmt.Println("please, input url")

	fmt.Scan(&x)

	conn, _ := grpc.Dial(":8080", grpc.WithInsecure())
	c := urlshorter.NewUrluhorterClient(conn)

	r ,_ :=c.Create(context.Background(), &urlshorter.LongUrl{Long: x})
	res, _ :=c.Get(context.Background(), &urlshorter.ShortUrl{Short:r.Short})

	fmt.Println(res)
	fmt.Println(r)
}
