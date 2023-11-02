package instances

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/go-playground/validator/v10"
	"lookerdevelopers/boilerplate/cmd/dto"
	"lookerdevelopers/boilerplate/cmd/tasks"
)

func InitListener() {
	log.Println("☁️  Building pub/sub connection")

	ctx := context.Background()
	projectID := ""
	subscriptionName := ""

	client, err := pubsub.NewClient(ctx, projectID)

	if err != nil {
		log.Fatal("Error initializing pub/sub")
	}

	subscription := client.Subscription(subscriptionName)
	log.Println("Pub/Sub initialization successfully 🚀")

	go func() {
		err = subscription.Receive(ctx, func(ctx context.Context, message *pubsub.Message) {

			var psPayload dto.TrackingPatchDto

			data := string(message.Data)
			log.Printf("Stringify payload: %s", data)

			err := json.Unmarshal(message.Data, &psPayload)

			if err != nil {
				log.Printf("⚠️  Error unmarshalling payload: %s", err)
				message.Ack()

				return
			}

			log.Println("Payload successfully unmarshalled ⚙️")

			validate := validator.New()

			err = validate.Struct(&psPayload)

			if err != nil {
				log.Printf("❌  Error validating payload: %s", err)
				message.Ack()

				return
			}

			log.Println("Payload successfully parsed 🤖")

			affectedRows, err := tasks.SavePSPayload(DB, &psPayload)

			if err != nil {
				message.Nack()
			}

			log.Printf("Affected rows: %d", affectedRows)

			message.Ack()
		})

		if err != nil && !errors.Is(err, context.Canceled) {
			log.Printf("Errro about received message from pub/sub: %s", err)
		}

	}()

}
