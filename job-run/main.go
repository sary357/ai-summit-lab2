package main

import (
      //  "github.com/gin-gonic/gin"
        "strings"
        "bytes"
        "encoding/json"
        "fmt"
        "net/http"
        "github.com/sirupsen/logrus"
        "job-run/app"
        "job-run/utils"
	"job-run/config"
)

type LockExecutor struct {
        LockExecutorID string `json:"executor_id"`
}

type JobLockInfo struct {
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

type JobLockResponse struct {
        Status  string   `json:"status"`
        JobLockInfo JobLockInfo  `json:"job_info,omitempty"`
        Msg     string   `json:"msg,omitempty"`
}


type ExeExecutor struct {
        ExecExecutorID string `json:"executor_id"`
        ExecJobID      int    `json:"job_id"`
}

type JobExecutionResponse struct {
        Status     string `json:"status"`
        JobStatus  string `json:"job_status"`
        Msg        string `json:"msg,omitempty"`
}
//  '{"job_id":1, "executor_id":"my_executor", "generated_api_endpoint":"my_endpoint","job_status":3}'
type FinishExecutor struct {
	JobID               int    `json:"job_id"`
	ExecutorID          string `json:"executor_id"`
	GeneratedAPIEndpoint string `json:"generated_api_endpoint"`
	JobStatus           int    `json:"job_status"`
}


func sendRequest(url string, requestData interface{}, responseData interface{}) error {
	// Encode the request data to JSON
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return fmt.Errorf("error marshalling request data: %w", err)
	}

	// Create a new HTTP POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	// Set the request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error: %s", resp.Status)
	}

	// Decode the response body into the provided responseData struct
	if responseData != nil {
		err = json.NewDecoder(resp.Body).Decode(responseData)
		if err != nil {
			return fmt.Errorf("error decoding response: %w", err)
		}
	}

	return nil
}

func sendJobLockRequest(executor LockExecutor, url string) (*JobLockResponse, error) {
	var jobLockResponse JobLockResponse
	err := sendRequest(url, executor, &jobLockResponse)
	if err != nil {
		return nil, fmt.Errorf("error sending job lock request: %w", err)
	}
	return &jobLockResponse, nil
}

func sendJobExecutionRequest(executor ExeExecutor, url string) (*JobExecutionResponse, error) {
	var jobExecutionResponse JobExecutionResponse
	err := sendRequest(url, executor, &jobExecutionResponse)
	if err != nil {
		return nil, fmt.Errorf("error sending job execution request: %w", err)
	}
	return &jobExecutionResponse, nil
}

func sendJobFinishRequest(executor FinishExecutor, url string) (*JobExecutionResponse, error) {
        var jobExecutionResponse JobExecutionResponse
        err := sendRequest(url, executor, &jobExecutionResponse)
        if err != nil {
                return nil, fmt.Errorf("error sending job execution request: %w", err)
        }
        return &jobExecutionResponse, nil
}


func main() {
	const JOB_SUCCESS int =3
        const JOB_FAILED int = 4
	base_url := config.JobBaseUrl
	my_executor_name := "my_executor"
        utils.LogInstance.WithFields(logrus.Fields{
		"Host": utils.GetHostname(),
	}).Info("job-run is starting. If there is no error message, it means the service is ready.")

        // Encode the executor struct to JSON
        executor := LockExecutor{LockExecutorID: my_executor_name}
        lockUrl := base_url + "/job_lock/" // lock url

        jobLockResponse, err := sendJobLockRequest(executor, lockUrl)
        if err != nil {
                panic(err)
        } else {
		fmt.Printf("Job ID: %d\n", jobLockResponse.JobLockInfo.ID)
		fmt.Printf("Job Codes: %s\n", jobLockResponse.JobLockInfo.Codes)
		fmt.Printf("Job requirements.txt: %s\n", jobLockResponse.JobLockInfo.RequirementsTxt)
		fmt.Printf("Message: %s\n", jobLockResponse.Msg)
	}


	jobId := jobLockResponse.JobLockInfo.ID
	exeExecutor := ExeExecutor{ExecExecutorID: my_executor_name, ExecJobID: jobId}
        exeUrl := base_url + "/job_execution/"

        jobExecutionResponse, err := sendJobExecutionRequest(exeExecutor, exeUrl)
        if err != nil {
                panic(err)
        } else {
		fmt.Println("Job Execution Response:", jobExecutionResponse)
	}

	jobExecResult:=app.SaveAndExec(jobLockResponse.JobLockInfo.Codes,jobLockResponse.JobLockInfo.RequirementsTxt)
	finishUrl := base_url + "/job_completion/"
	var finishExcutor FinishExecutor
        if strings.Contains(jobExecResult, "ERR") {
		finishExcutor = FinishExecutor {JobID: jobId, ExecutorID: my_executor_name, JobStatus: JOB_FAILED}
	} else {
		finishExcutor = FinishExecutor {JobID: jobId, ExecutorID: my_executor_name, GeneratedAPIEndpoint: jobExecResult, JobStatus: JOB_SUCCESS}
	}
	jobFinishResponse, err:=sendJobFinishRequest(finishExcutor, finishUrl)
        if err != nil {
                panic(err)
        } else {
		fmt.Println("Job Finish Response:", jobFinishResponse)
	}


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
