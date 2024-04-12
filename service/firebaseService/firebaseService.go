package firebaseService

import (
	"KUNoti/internal/controller/firebase/repository"
	"KUNoti/sqlc"
	"context"
	firebase "firebase.google.com/go/v4"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"

	"firebase.google.com/go/v4/messaging"
)

type FireBaseService interface {
	SendToToken(ctx context.Context)
	SendMulticastWithData(ctx context.Context, tokens []string, title, body string, data []byte) error
	Notification(ctx context.Context, token string, title, body string, data []byte) error
	Notifications(ctx context.Context, token string) ([]repository.Notification, error)
}

type FirebaseServiceClient struct {
	app                *firebase.App
	firebaseRepository *repository.FirebaseRepository
}

func NewFirebaseServiceClient(app *firebase.App, db *pgxpool.Pool) *FirebaseServiceClient {
	queries := sqlc.New(db)
	return &FirebaseServiceClient{
		app:                app,
		firebaseRepository: repository.NewFirebaseRepository(db, queries),
	}
}

func (f *FirebaseServiceClient) SendToToken(ctx context.Context) {

	client, err := f.app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	// This registration token comes from the client FCM SDKs.
	registrationToken := ""

	eventData := `{"id":12,"title":"Tech Seminar","latitude":13.846300381568938,"longitude":100.56847292643977,"start_date":"2024-04-11T10:00:00Z","end_date":"2024-04-11T18:00:00Z","created_at":"2024-04-08T03:39:22.53Z","updated_at":"2024-04-08T03:39:22.53Z","price":0,"rating":0,"image":"https://keventimage.s3.amazonaws.com/event/9552ef2a-f559-11ee-aaff-acde48001122-jpeg","creator":16,"detail":"Web3, Blockchain","location_name":"E17 building","need_regis":false,"tag":"KU","regis_amount":0,"regis_max":0}`

	// See documentation on defining a message payload.
	message := &messaging.Message{

		Notification: &messaging.Notification{
			Title: "New Event: Tech Seminar",
			Body:  "Check out this new event happening soon!",
		},
		Data: map[string]string{
			"event": eventData,
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

func (f *FirebaseServiceClient) Notification(ctx context.Context, token string, title, body string, data []byte) error {
	err := f.firebaseRepository.Create(ctx, repository.CreateNoti{
		Title: title,
		Body:  body,
		Data:  data,
		Token: token,
	})
	if err != nil {
		log.Println("Create notification to db fail")
		return err
	}
	return nil
}

func (f *FirebaseServiceClient) Notifications(ctx context.Context, token string) ([]repository.Notification, error) {
	notis, err := f.firebaseRepository.FindByToken(ctx, token)
	if err != nil {
		return nil, err
	}
	return notis, nil
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
	} else {
		for _, tk := range validTokens {
			err = f.firebaseRepository.Create(ctx, repository.CreateNoti{
				Title: title,
				Body:  body,
				Data:  data,
				Token: tk,
			})
			if err != nil {
				log.Println("Create notification to db fail")
				return err
			}
		}
	}

	log.Printf("Successfully sent message to %d devices", response.SuccessCount)
	return nil
}
