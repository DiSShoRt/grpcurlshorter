package interf

import (
	"context"
	"fmt"
	"grpcurlshorter/pkg/urlshorter"
	"grpcurlshorter/storage"

	"log"
)

type GRPCServer struct{}

func (s *GRPCServer) Create(ctx context.Context, lurl *urlshorter.LongUrl) (*urlshorter.ShortUrl, error) {
	short, err := storage.AddUrlToDb(lurl.Long)
	if err != nil {
		log.Println(err)
		return &urlshorter.ShortUrl{}, err
	}

	return &urlshorter.ShortUrl{Short: short}, nil
}

func (s *GRPCServer) Get(ctx context.Context, surl *urlshorter.ShortUrl) (*urlshorter.LongUrl, error) {
	long, err := storage.GetUrlFromDb(surl.Short)
	if err != nil {
		fmt.Println(err)
		return &urlshorter.LongUrl{}, err
	}
	return &urlshorter.LongUrl{Long: long}, nil
}

