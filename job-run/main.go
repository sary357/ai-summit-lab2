package main

import (
      //  "github.com/gin-gonic/gin"
        "bytes"
        "encoding/json"
        "fmt"
        "net/http"
        "github.com/sirupsen/logrus"
       // "job-run/app"
        "job-run/utils"
)

/*type CodeAndRelatedObject struct {
        Codes string `json:"codes"`
        RequirementTxt string `json:"requirementTxt"`
}*/

type Executor struct {
        ExecutorID string `json:"executor_id"`
}

type JobInfo struct {
        Status           int    `json:"status"`
        Endpoint         string `json:"endpoint,omitempty"`
        RunningExecutor  string `json:"running_executor,omitempty"`
        UpdatedAt        string `json:"updated_at"`
        RequirementsTxt  string `json:"requirements_txt,omitempty"`
        ID               int    `json:"id"`
        Codes            string `json:"codes"`
        LockExecutor     string `json:"lock_executor"`
        CreatedAt        string `json:"created_at"`
}

type JobResponse struct {
        Status  string   `json:"status"`
        JobInfo JobInfo  `json:"job_info"`
        Msg     string   `json:"msg,omitempty"`
}

// @title go API
// @version 1.0.0
// @description a sample for platform engineers to start with gin framework.

// @contact.name Fu-Ming Tsai
// @contact.email sary357@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func sendJobLockRequest(executor Executor, url string) (*JobResponse, error) {
        // Encode the executor struct to JSON
        requestBody, err := json.Marshal(executor)
        if err != nil {
                return nil, err
        }

        // Create a new HTTP POST request
        req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
        if err != nil {
                return nil, err
        }

        // Set the request headers
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Accept", "application/json")

        // Send the request
        client := &http.Client{}
        resp, err := client.Do(req)
        if err != nil {
                return nil, err
        }
        defer resp.Body.Close()

        // Check the response status code
        if resp.StatusCode != http.StatusOK {
                return nil, fmt.Errorf("Error: %s", resp.Status)
        }

        // Decode the response body
        var jobResponse JobResponse
        err = json.NewDecoder(resp.Body).Decode(&jobResponse)
        if err != nil {
                return nil, err
        }

        return &jobResponse, nil
}


func main() {
	my_executor_name := "my_executor"
        utils.LogInstance.WithFields(logrus.Fields{
		"Host": utils.GetHostname(),
	}).Info("job-run is starting. If there is no error message, it means the service is ready.")

        // Encode the executor struct to JSON
        executor := Executor{ExecutorID: my_executor_name}
        lockUrl := "http://localhost:8081/v1/job_lock/" // lock url

        jobResponse, err := sendJobLockRequest(executor, lockUrl)
        if err != nil {
                panic(err)
        }

	fmt.Printf("Job ID: %d\n", jobResponse.JobInfo.ID)
        fmt.Printf("Job Codes: %s\n", jobResponse.JobInfo.Codes)

/*	utils.LogInstance.WithFields(logrus.Fields{
                        "codes": codeAndRelatedObject.Codes,
                        "requirementstxt": codeAndRelatedObject.RequirementTxt,
                }).Info("user's inputs.")

                // start to process
         //       status:=app.SaveAndExec(codeAndRelatedObject.Codes, codeAndRelatedObject.RequirementTxt)
                //status:= "https://ihznxmqgj9.execute-api.ap-northeast-1.amazonaws.com/prod/"
                if strings.Contains(status, "ERR") {
                        retObj := ReturnObj{
                                Endpoint: "",
                                Message: status,
                        }
                        c.JSON(http.StatusBadRequest, retObj)
                } else {
                        retObj := ReturnObj{
                                Endpoint: status,
                                Message: "",
                        }
                        c.JSON(http.StatusOK, retObj)
                }
	/*///r.Run("0.0.0.0:" + strconv.Itoa(config.Port)) // listen and serve on 0.0.0.0:8080

}
