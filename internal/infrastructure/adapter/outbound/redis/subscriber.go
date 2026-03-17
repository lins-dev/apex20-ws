package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// RedisSubscriber implements the application port using Redis.
// In Go, this implementation satisfies the RedisSubscriber interface implicitly.
type RedisSubscriber struct {
	client *redis.Client
}

// NewRedisSubscriber creates a new implementation of the Redis subscriber.
func NewRedisSubscriber(client *redis.Client) *RedisSubscriber {
	return &RedisSubscriber{
		client: client,
	}
}

// Subscribe listens to a specific channel and manages the subscription lifecycle.
func (s *RedisSubscriber) Subscribe(ctx context.Context, channel string) (<-chan string, error) {
	pubsub := s.client.Subscribe(ctx, channel)
	
	// Pre-check for initial errors
	_, err := pubsub.Receive(ctx)
	if err != nil {
		return nil, err
	}

	msgChan := make(chan string)
	
	go func() {
		defer pubsub.Close()
		defer close(msgChan)
		
		ch := pubsub.Channel()
		for {
			select {
			case <-ctx.Done():
				return
			case msg, ok := <-ch:
				if !ok {
					return
				}
				msgChan <- msg.Payload
			}
		}
	}()

	return msgChan, nil
}
