import logging
logger = logging.getLogger(__name__)
from datetime import date
import os
from pydantic import BaseModel

from typing import Optional

from fastapi import FastAPI, HTTPException, Request, Response, status
import json
from src.v1 import jobProcessor
from src.v1 import SessionGenerator
from src.v1.const import *
app = FastAPI() 


ENCODEING="utf-8"

# Class 1: Job
# field: input_codes(str), input_requirements_txt(str)
class Job(BaseModel):
    codes: str
    requirements_txt: Optional[str] = None

# Class 2: executor
class Executor(BaseModel):
    executor_id: str
    job_id: Optional[int] = -1
    job_status: Optional[int] = -1
    generated_api_endpoint:  Optional[str] = None
    
@app.post("/v1/job_submission/")
async def job_submission(job: Job):
    logger.info('Got 1 job: %s', job) # (input_codes:str, input_requirements_txt:str, input_status:int=JOB_ACCEPTED)
    if job.codes:
        job_id=jobProcessor.save_job(input_codes=job.codes, input_requirements_txt=job.requirements_txt)

        if job_id == ERR_INTERNAL_ERROR:
            raise HTTPException(status_code=status.HTTP_500_INTERNAL_SERVER_ERROR, detail="Internal server error")
        else:
            return {"job_id": job_id}
    else:
        raise HTTPException(status_code=status.HTTP_400_BAD_REQUEST, detail="Invalid request body: missing \"codes\"")

@app.get("/v1/job_info/{job_id}")
async def get_job(job_id: int):
    logger.info('Got 1 jobId: %d', job_id)
    job=jobProcessor.get_job_by_id(id=job_id)
    if job:
        return {"status": "ok","job_info":job}
    else:
        raise HTTPException(status_code=status.HTTP_500_INTERNAL_SERVER_ERROR, detail="Internal server error")

@app.post("/v1/job_lock/")
async def lock_job(executor: Executor): 
    logger.info('executor : %s', executor)
    job=jobProcessor.lock_job(lock_executor=executor.executor_id)
    if job:
        return {"status": "ok", "job_info": job, "msg": None}
    else:
        return {"status": "ok", "job_info": None, "msg": "no accepted job"}

@app.post("/v1/job_execution/")
async def execute_job(executor: Executor): 
    logger.info('executor : %s', executor)
    job_status=jobProcessor.execute_job(id=executor.job_id, running_executor=executor.executor_id)
    if job_status:
        return {"status": "ok", "job_status": "running", "msg": None}
    else:
        return {"status": "ok", "job_status": None, "msg": f"job status or job id may not be valid. Please conact the administrator. job id: {executor.job_id}"}

@app.post("/v1/job_completion/")
async def complete_job(executor: Executor):  # finish_job(id: int, running_executor:str, status: int)
    logger.info('executor : %s', executor)
    job_status=jobProcessor.finish_job(id=executor.job_id, running_executor=executor.executor_id, status=executor.job_status, endpoint=executor.generated_api_endpoint)
    if job_status:
        return {"status": "ok", "job_status": "finished", "msg": None}
    else:
        return {"status": "ok", "job_status": None, "msg": f"job status or job id may not be valid. Please conact the administrator. job id: {executor.job_id}"}


@app.get("/v1/health-check")
async def health_check():
    if jobProcessor.health_check():
        return {"status": "ok"}
    else:
        raise HTTPException(status_code=status.HTTP_500_INTERNAL_SERVER_ERROR)


