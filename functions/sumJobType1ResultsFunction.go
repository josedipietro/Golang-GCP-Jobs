package implementation

import (
	"encoding/json"
	"fmt"
	"net/http"

	"josedipietro.com/teamcoretest/functions/firestore"
	"josedipietro.com/teamcoretest/functions/structs"
)

type SumJobsType1ResultsArgs struct {
	JobsIds []string `json:"jobsIds"`
}


func sumJobType1Results(jobsId []string) <-chan float64 {
	channel := make(chan float64)

	sum := 0.0
	go func ()  {
		defer close(channel)
		for _, jobId := range jobsId {
			job, err := firestore.GetJob(jobId)

			if (err != nil) {
				continue;
			}

			if (job.JobType != structs.CalculateMedian) {
				continue;
			}
			
			if (job.Payload.Error != nil) {
				continue
			}

			result := job.Payload.Result.(float64)

			sum += result
		}

		channel <- sum
	}()

	return channel
}

func SumJobType1ResultsFunction(w http.ResponseWriter, r *http.Request) {
	var job structs.Job
	if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
		fmt.Fprint(w, "Invalid Job")
		return
	}

	jsonData, _ := json.Marshal(job.Args)

	var params SumJobsType1ResultsArgs

	json.Unmarshal(jsonData, &params)

	sum := <- sumJobType1Results(params.JobsIds)

	job.Payload.Result = sum
	job.Payload.Error = nil

	firestore.UpdateJob(job)
	fmt.Fprint(w, "SumJobs ", sum, " for jobId ", job.ID)
}
