package firebaseService

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"fmt"
	"log"

	"firebase.google.com/go/v4/messaging"
)

type FireBaseService interface {
	SendToToken(ctx context.Context, data []byte)
	SendMulticastWithData(ctx context.Context, tokens []string, title, body string, data []byte) error
}

type FirebaseServiceClient struct {
	app *firebase.App
}

func NewFirebaseServiceClient(app *firebase.App) *FirebaseServiceClient {
	return &FirebaseServiceClient{app: app}
}

func (f *FirebaseServiceClient) SendToToken(ctx context.Context, data []byte) {

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
			"event": string(data),
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

func validateTokens(tokens []string) ([]string, error) {
	var validTokens []string
	for _, token := range tokens {
		if token == "" {
			log.Printf("Encountered empty FCM token, skipping")
			continue
		}
		// Add more sophisticated checks if necessary, e.g., regex matching if you know the expected format
		validTokens = append(validTokens, token)
	}
	if len(validTokens) == 0 {
		return nil, fmt.Errorf("no valid tokens provided")
	}
	return validTokens, nil
}

func (f *FirebaseServiceClient) SendMulticastWithData(ctx context.Context, tokens []string, title, body string, data []byte) error {
	client, err := f.app.Messaging(ctx)
	if err != nil {
		log.Printf("Error getting Messaging client: %v", err)
		return err
	}

	validTokens, err := validateTokens(tokens)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Valid tokens:", validTokens)

	message := &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Data: map[string]string{
			"event": string(data),
		},
		Tokens: validTokens,
	}

	response, err := client.SendEachForMulticast(ctx, message)
	if err != nil {
		log.Printf("Failed to send multicast message: %v", err)
		return err
	}

	if response.FailureCount > 0 {
		for idx, resp := range response.Responses {
			if !resp.Success {
				log.Printf("Failed to deliver to token %s: %v", validTokens[idx], resp.Error)
			}
		}
	}

	log.Printf("Successfully sent message to %d devices", response.SuccessCount)
	return nil
}
