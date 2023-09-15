package implementation

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"

	"josedipietro.com/teamcoretest/functions/firestore"
	"josedipietro.com/teamcoretest/functions/structs"
)

type PasswordGeneratorParams struct {
  PasswordLength int `json:"passwordLength"`
  MinSpecialChar int `json:"minSpecialChar"` // optional
  MinUpperCase int `json:"minUpperCase"` // optional
  MinNum int `json:"minNum"` // optional
}

var basicPasswordParams = PasswordGeneratorParams{PasswordLength: 8, MinSpecialChar: 1, MinUpperCase: 1, MinNum: 1}

var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func generateRandomPassword(passwordParams PasswordGeneratorParams) string {
	var password strings.Builder

    //Set special character
    for i := 0; i < passwordParams.MinSpecialChar; i++ {
        random := rand.Intn(len(specialCharSet))
        password.WriteString(string(specialCharSet[random]))
    }

    //Set numeric
    for i := 0; i < passwordParams.MinNum; i++ {
        random := rand.Intn(len(numberSet))
        password.WriteString(string(numberSet[random]))
    }

    //Set uppercase
    for i := 0; i < passwordParams.MinUpperCase; i++ {
        random := rand.Intn(len(upperCharSet))
        password.WriteString(string(upperCharSet[random]))
    }

    remainingLength := passwordParams.PasswordLength - passwordParams.MinSpecialChar - passwordParams.MinNum - passwordParams.MinUpperCase
    for i := 0; i < remainingLength; i++ {
        random := rand.Intn(len(allCharSet))
        password.WriteString(string(allCharSet[random]))
    }
    inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})

	return string(inRune)
}

func GenerateRandomPasswordFunction(w http.ResponseWriter, r *http.Request) {
    var job structs.Job
	if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
		fmt.Fprint(w, "Invalid Job")
		return
	}

    jsonData, _ := json.Marshal(job.Args)

    var params PasswordGeneratorParams

    json.Unmarshal(jsonData, &params)

    password := generateRandomPassword(params)

    log.Println("password")
    log.Println(password)

    job.Payload.Result = password
    job.Payload.Error = nil
    
    firestore.UpdateJob(job)

    fmt.Fprint(w, "Password generated ", password, " for jobId ", job.ID)
}
