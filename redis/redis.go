package redis

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"os"

	"github.com/cameronbrill/brill-wtf-go/model"
	"github.com/go-redis/redis/v8"
)

type Storage struct {
	client *redis.Client
}

func (s *Storage) Connect() error {
	s.client = redis.NewClient(&redis.Options{
		Addr:      os.Getenv("REDIS_ADDR"),
		Password:  os.Getenv("REDIS_PASSWORD"),
		Username:  os.Getenv("REDIS_USERNAME"),
		TLSConfig: &tls.Config{},
	})
	_, err := s.client.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) Get(key string) (model.Link, error) {
	link, err := s.client.Get(context.Background(), key).Result()
	if err != nil {
		fmt.Printf("Error getting link: %+v\n", link)
		return model.Link{}, err
	}
	var l model.Link
	err = json.Unmarshal([]byte(link), &l)
	if err != nil {
		return model.Link{}, err
	}
	return l, nil
}

func (s *Storage) Set(key string, value model.Link) error {
	valB, err := json.Marshal(value)
	if err != nil {
		return err
	}
	_, err = s.client.Set(context.Background(), key, string(valB), value.TTL).Result()
	if err != nil {
		return err
	}
	return nil
}
