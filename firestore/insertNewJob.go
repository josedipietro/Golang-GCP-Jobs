package firestore

import (
	"context"
	"log"
	"teamcore/test/structs"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

func InsertNewJob(job structs.Job) string {
	ctx := context.Background()
	opt := option.WithCredentialsFile("service_account.json")
	client, err := firestore.NewClient(ctx, "teamcoretest", opt)
	if err != nil {
		log.Fatalf("firestore new error:%s\n", err)
	}
	defer client.Close()

	jobRef := client.Collection("jobs").NewDoc()
	job.ID = jobRef.ID
	job.CreatedAt = time.Now()
	wr, _ := jobRef.Set(ctx, job)

	log.Println("Job ",  jobRef.ID, " created in ", wr.UpdateTime)
	return jobRef.ID
}