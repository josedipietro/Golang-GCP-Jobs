package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"teamcore/test/firestore"
	"teamcore/test/structs"

	"github.com/gin-gonic/gin"
)

type RecieveJobs struct {
	Jobs []structs.Job `json:"jobs"`
}


func recieveJobs(c *gin.Context) {
	var recieveJobs RecieveJobs

	if err := c.BindJSON(&recieveJobs); err != nil {
		log.Fatalf("Error in data:%s\n", err)
		return
	}

	newJobs := recieveJobs.Jobs

	sort.SliceStable(newJobs, func(i, j int) bool {
		return newJobs[i].Priority == structs.Max
	})

	for _, job := range newJobs {
		jobId := firestore.InsertNewJob(job)
		job.ID = jobId

		if (job.ExecType == structs.ExecInmediatly) {
			executeJob(job)
		}
	}

	c.IndentedJSON(http.StatusCreated, "Jobs created")
}

func executeJob(job structs.Job) {
	var url string
	switch job.JobType {
		case structs.CalculateMedian:
			url = os.Getenv("CALCULATE_MEDIAN_URL")
		case structs.GeneratePassword:
			url = os.Getenv("GENERATE_RANDOM_PASSWORD_URL")
		case structs.SumJobsType1:
			url = os.Getenv("SUM_JOB_TYPE_1_RESULTS_URL")
	}
	
	sendRequest(url, job)

	log.Printf("Job %s executed", job.ID)
}

func sendRequest(url string, body structs.Job) {
	// Create an HTTP client
	client := &http.Client{}

	jsonData, err := json.Marshal(body)

	if err != nil {

		log.Println("Error in Job json:", err)
		return
	}

	// Send the request and get the response
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	log.Println(res["json"])
}