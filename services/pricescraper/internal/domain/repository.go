package domain

import "context"

// ObservationRepository defines the interface for publishing observations
type EventStreamer interface {
	Publish(ctx context.Context, observation *AmazonProduct) error
}
