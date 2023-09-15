package firestore

import (
	"context"
	"encoding/json"
	"log"

	"josedipietro.com/teamcoretest/functions/structs"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

func GetJob(id string) (*structs.Job, error) {
	ctx := context.Background()
	service_account_json, _ := json.Marshal(service_account)
	opt := option.WithCredentialsJSON(service_account_json)
	client, err := firestore.NewClient(ctx, "teamcoretest", opt)
	if err != nil {
		log.Fatalf("firestore new error:%s\n", err)
	}
	defer client.Close()

	snapshot, err := client.Collection("jobs").Doc(id).Get(ctx)

	if err != nil || !snapshot.Exists() {
		log.Printf("Error getting job %s", err)
		return nil, err
	}

	var job structs.Job

	snapshot.DataTo(&job)

	return &job, nil
}