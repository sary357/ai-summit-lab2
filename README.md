# 請先讀我 
## Folder structure
```
go-api: it accepts http request and generate API gateway/lambda functions with AWS CDK. It takes 2 or 3 minutes for each request because of AWS CDK.
-----
job-submit-api: it access http request and save contents in the database.
job-run: it query job-submit-api and get job info to execute.
```
## AWS IAM setting
- The following is my IAM Policies
  - AmazonAPIGatewayAdministrator
  - AmazonEC2ContainerRegistryFullAccess
  - AmazonS3FullAccess
  - AmazonSSMReadOnlyAccess
  - AWSCloudFormationFullAccess
  - AWSLambda_FullAccess
  - IAMFullAccess
  - customize IAM access
- then create a API key 
- use aws cli to set up the environment with API key/credentail generated on AWS IAM page.

# go-api
- Prerequisite: Go version >= 1.17 & Python >= 3.9
- go to the folder `go-api`
```
$ cd go-api
```

- run `go init`
```
$ go mod init go-api
```

- run `go mod tidy`
```
$ go mod tidy
```

- ideally you can start the API server with the command
```
$ go run .
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /v1/healthcheck           --> go-api/routes.SetupHealthCheckRoute.func1 (3 handlers)
[GIN-debug] POST   /v1/genapi                --> go-api/routes.SetupAwsCdkRoute.func1 (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on 0.0.0.0:8080
```

- test with `curl` and you're supposed to see the message like the following.
```
$ curl -X POST -H "Content-Type: application/json" -d '{"codes": "import requests\n\ndef lambda_handler(event, context):\n  url = \"https://api.twelvedata.com/price?symbol=MSFT&apikey=16************e1\"\n  response = requests.get(url)\n  data = response.json()\n\n  stock_price = data[\"price\"]\n\n  return {\n    \"statusCode\": 200,\n    \"body\": f\"The latest MSFT stock price is: {stock_price}\"  }","requirementTxt":"requests"}'  http://localhost:8080/v1/genapi
{"endpoint":"https://8xrbdo625f.execute-api.ap-northeast-1.amazonaws.com/prod/","message":""}
```
- Then, you can access the endpoint `https://8xrbdo625f.execute-api.ap-northeast-1.amazonaws.com/prod/` with the browser.


# job-submit-api
- Please note: this API does NOT process any jobs that can create AWS components
- Prerequisite: Python >= 3.9 & docker & docker-compose
- go to `job-submit-api`
```
$ cd job-submit-api
```
- Prepare a PostgreSQL container with docker-compose
```
$ docker-compose up -d
```
- create database and table with the [SQL script](job-submit-api/sql/create_tables.sql)
- prepare virtual env
```
$ python3 -m venv venv
$ source venv/bin/activitate
```

- install necesssary python package
```
$ pip install -r requirements.txt
```

- start the server the the script [start.sh](job-submit-api/start.sh)
```
$ sh start.sh
```

- about how to use the API, please refer to the [doc](job-submit-api/docs/v1/README.md)
