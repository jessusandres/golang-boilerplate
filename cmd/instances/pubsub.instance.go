package instances

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	llog "lookerdevelopers/boilerplate/cmd/logger"

	"lookerdevelopers/boilerplate/cmd/dto"
	"lookerdevelopers/boilerplate/cmd/tasks"

	"cloud.google.com/go/pubsub"
	"github.com/go-playground/validator/v10"
)

func InitListener() {
	llog.Logger.Infoln("☁️  Building pub/sub connection")

	ctx := context.Background()
	projectID := ""
	subscriptionName := ""

	client, err := pubsub.NewClient(ctx, projectID)

	if err != nil {
		log.Fatal("Error initializing pub/sub")
	}

	subscription := client.Subscription(subscriptionName)
	llog.Logger.Infoln("Pub/Sub initialization successfully 🚀")

	go func() {
		err = subscription.Receive(ctx, func(ctx context.Context, message *pubsub.Message) {

			var psPayload dto.IncidentPatchDto

			data := string(message.Data)
			log.Printf("Stringify payload: %s", data)

			err := json.Unmarshal(message.Data, &psPayload)

			if err != nil {
				log.Printf("⚠️  Error unmarshalling payload: %s", err)
				message.Ack()

				return
			}

			llog.Logger.Infoln("Payload successfully unmarshalled ⚙️")

			validate := validator.New()

			err = validate.Struct(&psPayload)

			if err != nil {
				log.Printf("❌  Error validating payload: %s", err)
				message.Ack()

				return
			}

			llog.Logger.Infoln("Payload successfully parsed 🤖")

			result := tasks.SaveIncident(DB, &psPayload)

			if result.Error != nil {
				message.Nack()
			}

			log.Printf("Affected rows: %d", result.RowsAffected)

			message.Ack()
		})

		if err != nil && !errors.Is(err, context.Canceled) {
			log.Printf("Errro about received message from pub/sub: %s", err)
		}

	}()

}
