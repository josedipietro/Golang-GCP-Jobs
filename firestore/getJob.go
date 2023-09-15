package firestore

import (
	"context"
	"log"
	"teamcore/test/structs"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

func GetJob(id string) (*structs.Job, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile("service_account.json")
	client, err := firestore.NewClient(ctx, "teamcoretest", opt)
	if err != nil {
		log.Fatalf("firestore new error:%s\n", err)
	}
	defer client.Close()

	snapshot, err := client.Collection("jobs").Doc(id).Get(ctx)

	if (err != nil || !snapshot.Exists()) {
		log.Printf("Error getting job %s", err)
		return nil, err
	}

	var job structs.Job

	snapshot.DataTo(&job)

	return &job, nil
}