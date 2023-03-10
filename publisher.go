package p

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
)

func publish(projectID, topicID string, event any, attributes map[string]string) (error, error) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err), nil
	}

	defer client.Close()

	message, _ := json.Marshal(event)

	topic := client.Topic(topicID)
	result := topic.Publish(ctx, &pubsub.Message{
		Data:       message,
		Attributes: attributes,
	})

	_, err = result.Get(ctx)
	if err != nil {
		return fmt.Errorf("get: %v", err), nil
	}
	return nil, nil
}
