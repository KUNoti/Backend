package firebaseService

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"fmt"
	"log"

	"firebase.google.com/go/v4/messaging"
)

type FireBaseService interface {
	SendToToken(ctx context.Context)
}

type FirebaseServiceClient struct {
	app *firebase.App
}

func NewFirebaseServiceClient(app *firebase.App) *FirebaseServiceClient {
	return &FirebaseServiceClient{app: app}
}

func (f *FirebaseServiceClient) SendToToken(ctx context.Context) {

	ctx = context.Background()
	client, err := f.app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	// This registration token comes from the client FCM SDKs.
	registrationToken := ""

	// See documentation on defining a message payload.
	message := &messaging.Message{

		Notification: &messaging.Notification{
			Title: "test",
			Body:  "testBody",
		},
		Data: map[string]string{
			"score": "850",
			"time":  "2:45",
		},
		Token: registrationToken,
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)
	// [END send_to_token_golang]
}

func sendMulticastAndHandleErrors(client *messaging.Client) {
	// [START send_multicast_error]
	// Create a list containing up to 500 registration tokens.
	// This registration tokens come from the client FCM SDKs.
	registrationTokens := []string{
		"YOUR_REGISTRATION_TOKEN_1",
		// ...
		"YOUR_REGISTRATION_TOKEN_n",
	}
	message := &messaging.MulticastMessage{
		Data: map[string]string{
			"score": "850",
			"time":  "2:45",
		},
		Tokens: registrationTokens,
	}

	br, err := client.SendEachForMulticast(context.Background(), message)
	if err != nil {
		log.Fatalln(err)
	}

	if br.FailureCount > 0 {
		var failedTokens []string
		for idx, resp := range br.Responses {
			if !resp.Success {
				// The order of responses corresponds to the order of the registration tokens.
				failedTokens = append(failedTokens, registrationTokens[idx])
			}
		}

		fmt.Printf("List of tokens that caused failures: %v\n", failedTokens)
	}
	// [END send_multicast_error]
}

//func sendToTopic(ctx context.Context, client *messaging.Client) {
//	// [START send_to_topic_golang]
//	// The topic name can be optionally prefixed with "/topics/".
//	topic := "highScores"
//
//	// See documentation on defining a message payload.
//	message := &messaging.Message{
//		Data: map[string]string{
//			"score": "850",
//			"time":  "2:45",
//		},
//		Topic: topic,
//	}
//
//	// Send a message to the devices subscribed to the provided topic.
//	response, err := client.Send(ctx, message)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	// Response is a message ID string.
//	fmt.Println("Successfully sent message:", response)
//	// [END send_to_topic_golang]
//}
