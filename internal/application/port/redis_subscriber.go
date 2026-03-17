package port

import "context"

// RedisSubscriber defines the contract for listening to real-time events.
// This is a Port in Hexagonal Architecture.
type RedisSubscriber interface {
	// Subscribe starts listening to a channel and returns a stream of messages.
	// The lifecycle of the subscription is tied to the provided context.
	Subscribe(ctx context.Context, channel string) (<-chan string, error)
}
