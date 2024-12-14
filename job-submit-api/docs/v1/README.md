
# Health check
- please verify the status with the command `curl`.

```
$ curl -v http://localhost:8081/v1/health-check
* Host localhost:8081 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
*   Trying [::1]:8081...
* connect to ::1 port 8081 from ::1 port 42578 failed: Connection refused
*   Trying 127.0.0.1:8081...
* Connected to localhost (127.0.0.1) port 8081
> GET /v1/health-check HTTP/1.1
> Host: localhost:8081
> User-Agent: curl/8.5.0
> Accept: */*
>
< HTTP/1.1 200 OK
< date: Sat, 14 Dec 2024 09:13:13 GMT
< server: uvicorn
< content-length: 15
< content-type: application/json
<
* Connection #0 to host localhost left intact
{"status":"ok"}
```

# Submit a job
- Please submit the job with the command `curl`. (no requirements.txt)
```
$ curl --verbose -H 'accept: application/json' -H 'Content-Type: application/json' -X POST http://localhost:8081/v1/job_submission/ -d '{"codes":"there are lots of codes"}'
Note: Unnecessary use of -X or --request, POST is already inferred.
* Host localhost:8081 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
*   Trying [::1]:8081...
* connect to ::1 port 8081 from ::1 port 50342 failed: Connection refused
*   Trying 127.0.0.1:8081...
* Connected to localhost (127.0.0.1) port 8081
> POST /v1/job_submission/ HTTP/1.1
> Host: localhost:8081
> User-Agent: curl/8.5.0
> accept: application/json
> Content-Type: application/json
> Content-Length: 35
>
< HTTP/1.1 200 OK
< date: Sat, 14 Dec 2024 09:14:57 GMT
< server: uvicorn
< content-length: 12
< content-type: application/json
<
* Connection #0 to host localhost left intact
{"job_id":1}
```
-  Please submit the job with the command `curl`. (with requirements.txt)
```
$ curl --verbose -H 'accept: application/json' -H 'Content-Type: application/json' -X POST http://localhost:8081/v1/job_submission/ -d '{"codes": "import requests\n\ndef lambda_handler(event, context):\n  url = \"https://api.twelvedata.com/price?symbol=MSFT&apikey=XXXXX\"\n  return {\n    \"statusCode\": 200,\n    \"body\": f\"URL is: {url}\"  }","requirements_txt":"requests"}'
Note: Unnecessary use of -X or --request, POST is already inferred.
* Host localhost:8081 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
*   Trying [::1]:8081...
* connect to ::1 port 8081 from ::1 port 52454 failed: Connection refused
*   Trying 127.0.0.1:8081...
* Connected to localhost (127.0.0.1) port 8081
> POST /v1/job_submission/ HTTP/1.1
> Host: localhost:8081
> User-Agent: curl/8.5.0
> accept: application/json
> Content-Type: application/json
> Content-Length: 243
>
< HTTP/1.1 200 OK
< date: Sat, 14 Dec 2024 09:25:02 GMT
< server: uvicorn
< content-length: 12
< content-type: application/json
<
* Connection #0 to host localhost left intact
{"job_id":3}
```

# query a job info
-  Please query the job info with the command `curl`. (job id: 3)
```
$ curl -v  http://localhost:8081/v1/job_info/3
* Host localhost:8081 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
*   Trying [::1]:8081...
* connect to ::1 port 8081 from ::1 port 40306 failed: Connection refused
*   Trying 127.0.0.1:8081...
* Connected to localhost (127.0.0.1) port 8081
> GET /v1/job_info/3 HTTP/1.1
> Host: localhost:8081
> User-Agent: curl/8.5.0
> Accept: */*
>
< HTTP/1.1 200 OK
< date: Sat, 14 Dec 2024 09:54:07 GMT
< server: uvicorn
< content-length: 444
< content-type: application/json
<
* Connection #0 to host localhost left intact
{"status":"ok","job_info":{"status":0,"endpoint":null,"running_executor":null,"updated_at":"2024-12-14T09:25:03.321980+00:00","requirements_txt":"requests","id":3,"codes":"import requests\n\ndef lambda_handler(event, context):\n  url = \"https://api.twelvedata.com/price?symbol=MSFT&apikey=XXXXX\"\n  return {\n    \"statusCode\": 200,\n    \"body\": f\"URL is: {url}\"  }","lock_executor":null,"created_at":"2024-12-14T09:25:03.321977+00:00"}}
```

# lock a job
- please locak a job with the command `curl`.
```
$ curl --verbose -H 'accept: application/json' -H 'Content-Type: application/json' -X POST http://localhost:8081/v1/job_lock/ -d '{"executor_id":"my_executor"}'
Note: Unnecessary use of -X or --request, POST is already inferred.
* Host localhost:8081 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
*   Trying [::1]:8081...
* connect to ::1 port 8081 from ::1 port 44966 failed: Connection refused
*   Trying 127.0.0.1:8081...
* Connected to localhost (127.0.0.1) port 8081
> POST /v1/job_lock/ HTTP/1.1
> Host: localhost:8081
> User-Agent: curl/8.5.0
> accept: application/json
> Content-Type: application/json
> Content-Length: 29
>
< HTTP/1.1 200 OK
< date: Sat, 14 Dec 2024 10:35:15 GMT
< server: uvicorn
< content-length: 281
< content-type: application/json
<
* Connection #0 to host localhost left intact
{"status":"ok","job_info":{"status":1,"endpoint":null,"running_executor":null,"updated_at":"2024-12-14T10:33:28.789674+00:00","requirements_txt":null,"id":1,"codes":"there are lots of codes","lock_executor":"my_executor","created_at":"2024-12-14T09:14:58.181847+00:00"},"msg":null}
```

# execute a job
```
$ curl --verbose -H 'accept: application/json' -H 'Content-Type: application/json' -X POST http://localhost:8081/v1/job_execution/ -d '{"executor_id":"my_executor", "job_id":1}'
Note: Unnecessary use of -X or --request, POST is already inferred.
* Host localhost:8081 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
*   Trying [::1]:8081...
* connect to ::1 port 8081 from ::1 port 40110 failed: Connection refused
*   Trying 127.0.0.1:8081...
* Connected to localhost (127.0.0.1) port 8081
> POST /v1/job_execution/ HTTP/1.1
> Host: localhost:8081
> User-Agent: curl/8.5.0
> accept: application/json
> Content-Type: application/json
> Content-Length: 41
>
< HTTP/1.1 200 OK
< date: Sat, 14 Dec 2024 10:45:32 GMT
< server: uvicorn
< content-length: 49
< content-type: application/json
<
* Connection #0 to host localhost left intact
{"status":"ok","job_status":"running","msg":null}
```

# complete a job (successful)
- job status=3
```
$ curl --verbose -H 'accept: application/json' -H 'Content-Type: application/json' -X POST http://localhost:8081/v1/job_completion/ -d '{"job_id":1, "executor_id":"my_executor", "generated_api_endpoint":"my_endpoint","job_status":3}'
Note: Unnecessary use of -X or --request, POST is already inferred.
* Host localhost:8081 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
*   Trying [::1]:8081...
* connect to ::1 port 8081 from ::1 port 40752 failed: Connection refused
*   Trying 127.0.0.1:8081...
* Connected to localhost (127.0.0.1) port 8081
> POST /v1/job_completion/ HTTP/1.1
> Host: localhost:8081
> User-Agent: curl/8.5.0
> accept: application/json
> Content-Type: application/json
> Content-Length: 96
>
< HTTP/1.1 200 OK
< date: Sat, 14 Dec 2024 11:29:38 GMT
< server: uvicorn
< content-length: 50
< content-type: application/json
<
* Connection #0 to host localhost left intact
{"status":"ok","job_status":"finished","msg":null}
```

# complere a job (failed)
- job status = 4
```
$ curl --verbose -H 'accept: application/json' -H 'Content-Type: application/json' -X POST http://localhost:8081/v1/job_completion/ -d '{"job_id":1, "executor_id":"my_executor", "generated_api_endpoint":"my_endpoint","job_status":4}'
Note: Unnecessary use of -X or --request, POST is already inferred.
* Host localhost:8081 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
*   Trying [::1]:8081...
* connect to ::1 port 8081 from ::1 port 56280 failed: Connection refused
*   Trying 127.0.0.1:8081...
* Connected to localhost (127.0.0.1) port 8081
> POST /v1/job_completion/ HTTP/1.1
> Host: localhost:8081
> User-Agent: curl/8.5.0
> accept: application/json
> Content-Type: application/json
> Content-Length: 96
>
< HTTP/1.1 200 OK
< date: Sat, 14 Dec 2024 11:31:12 GMT
< server: uvicorn
< content-length: 50
< content-type: application/json
<
* Connection #0 to host localhost left intact
{"status":"ok","job_status":"finished","msg":null}
```
