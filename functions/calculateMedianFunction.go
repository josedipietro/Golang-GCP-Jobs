package implementation

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"

	"josedipietro.com/teamcoretest/functions/firestore"
	"josedipietro.com/teamcoretest/functions/structs"
)

type CalculateMedianArgs struct {
    Numbers []float64 `json:"numbers"`
}

var service_account = map[string]interface{}{
    "type": "service_account",
    "project_id": "teamcoretest",
    "private_key_id": "64804e2702398bb62456ee462f6b283412523e4b",
    "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCeCJzjLbvPCVU9\nMhbOVxG4gsrB5qcWdWHe+ZRsrLCdnAfQ8Sd/sAZZFz+7mUh74Jhm3/Q8um0Aiy2/\nWpc0kYb5vze5SOEHlT0GMNm273ySJPk+vuRuYlttu3+bUnNnmBXzDnDen+bQKivC\nBNWA9RgxC3rY666CWyt9Y1j08iSv5+VCJJzFgGV704rAWvIzQmi1XkHpNPbuA9pP\nJQwgGVUtPp0D+5zgI3Za4QoYmH/Wx/Y4f4cYVzYSHivPeRFz4A45It546dRfbHSK\n3dbT6TneQF2xr/x/7bSirY4NXS5//K/qyFvcNl1Zl2lWRwKGDH8Dy0NFNqaclMdA\nfe/dyFCBAgMBAAECggEAEyrqg3vzLvD7+ZaLEVcab0GO396O1tChMKsfG2DFMbKg\neJFdV/WEyZbCLEny0pKazYR9kkWOzb1zzKUrWI0LnNxaYQLjx1iMrKTtbSymNHAf\nWoLLAE/19KzklPVWycJ2rBKs8jWdCFPLF6bgMfpYRy+UJH4OfiFaeKR9gEUj2cFi\n/0grS4HgdrAWH1kdFLMHy+I8lnAhDRIHeM0lGPOUsyjNpTOQPODFcTGVaDGSAEqI\n22c4bss52bdNtTUKqdHA1Qn8LsJl4gTT2DmT2qFIrQe797qng9xgpzogOW93sGf5\nreccI77wcZ4RLRqr794JTSsGD1jFmJNsJioP0weuyQKBgQDL7yJBd65lLjKp1FEy\nfHwLEY3MaR6kgd7NcvgP1x30oiJff7CfOheOCETUOT1HMr0UuIn5khNwrb3ZjzhZ\nG4Hg4YorkqMaoNxOCqAXmHiQcw58uazGWL3vUI3MD4IdnJ1EWg6ZSsknW20fn/ti\nGFCwwQEWKSBuZsIiJFadOFBCqQKBgQDGYXwA67UzutUsxNHk6u7Ag7dSwbg9tznZ\njLEpUmJi9pKt8GawlP09UanPZIg6BJ1Q4KqvAYgJ6W37qEHqd0e0ZFHTqkMdKppQ\n02QdGO9m5krZtir4RORR9dE1mLDqp+hO6P0JX4iB3SdkXGTNX+f2ZhvrV3CBouhI\nt3zPPFIeGQKBgAYjs99OsBRhuKq+NpeTgdR/ecpem8qOElwTCv7HFiLIQsqnOHUC\ntbTT8OaGtp6PG3wlNhoqWKV6xY1oL2UXW+ieQZ+gMYLatucukLVFJNQMcrI0kMwD\n3ev3e91Z1iv4DBADug7JXpbtvLJICbRhUQSROuwk3tIUC+IlP+pJyLjJAoGBAJ8K\ntnySe7IePhtnK5MoGgLzVjydnBS2WUWlOr8TEleesJeMXPeCasgHOWlQgrpoyqp1\ng8FMLAEuSINyMG5F9JGVv9g+7xFp/09/Ogrt27iWNjn2htqFeLqQpYofgO6PcHoa\n5gnmsizS3Wrje9j+45ux3v49GrCDp0/s8r298WO5AoGAVHf59JfSK9jUzhOMzEv/\n+Qdj8wWyTdzIgmXaFikKoOzvRZSIP67S9nCxz4Xj2WUctOGhEesUSv8dMtUrf76E\nyWgZ3XOzIf3r0pkBCYPwuJoTuM3CJwdSdj9FgcfPzjNYMsyTqpsS9niD/i9QFWW7\n3jnGK+1YFR/+2pyBqO/dnEY=\n-----END PRIVATE KEY-----\n",
    "client_email": "teamcoretest@appspot.gserviceaccount.com",
    "client_id": "100726675943388623815",
    "auth_uri": "https://accounts.google.com/o/oauth2/auth",
    "token_uri": "https://oauth2.googleapis.com/token",
    "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
    "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/teamcoretest%40appspot.gserviceaccount.com",
    "universe_domain": "googleapis.com",
  }

func calculateMedian(data []float64) float64 {
	dataCopy := make([]float64, len(data))
    copy(dataCopy, data)

    sort.Float64s(dataCopy)

    var median float64
    l := len(dataCopy)
    if l == 0 {
        return 0
    } else if l%2 == 0 {
        median = (dataCopy[l/2-1] + dataCopy[l/2]) / 2
    } else {
        median = dataCopy[l/2]
    }

    return median
}

func CalculateMedianFunction(w http.ResponseWriter, r *http.Request) {
    var job structs.Job
    if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
        fmt.Fprint(w, "Invalid Job")
        return
    }

    jsonData, _ := json.Marshal(job.Args)

    log.Println("jsonData")
    log.Println(jsonData)

    var params CalculateMedianArgs

    json.Unmarshal(jsonData, &params)

    if len(params.Numbers) == 0 {
        fmt.Fprint(w, "No numbers to calculate median")
        return
    }

    median:= calculateMedian(params.Numbers)

    log.Println("median")
    log.Println(median)

    job.Payload.Result = median
    job.Payload.Error = nil

    firestore.UpdateJob(job)

    fmt.Fprint(w, "Median calculated ", median, " for jobId ", job.ID)
}