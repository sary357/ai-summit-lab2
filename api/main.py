import logging
logger = logging.getLogger(__name__)
from datetime import date
import os
from pydantic import BaseModel

from typing import Optional

from fastapi import FastAPI, HTTPException, Request, Response, status
import json
app = FastAPI() 

ENCODEING="utf-8"

#@app.post("/v2/vote/")
#async def save_vote(vote: Vote):
#    logger.info('Got 1 vote: %s', vote)
#    save_status=False
#    if vote.id:
#        save_status=voteProcessor.save_vote_by_id(id=vote.id, user_vote=vote.vote)
#    elif vote.question and vote.phone_number and vote.answer:
#        save_status=voteProcessor.save_vote(user_input_question=vote.question, user_phone=vote.phone_number, 
#                                user_vote=vote.vote, user_input_answer=vote.answer)
#    else:
#        raise HTTPException(status_code=status.HTTP_400_BAD_REQUEST, detail="Invalid request body: missing id or (phone_number + question + response)")
#    if not save_status:
#        raise HTTPException(status_code=status.HTTP_500_INTERNAL_SERVER_ERROR, detail="Internal server error")
#    else:
#        return {"status": "ok"}
#
#@app.post("/v2/qa/")
#async def save_query(qa: QA):
#    logger.info('Got 1 Question/Answer: %s', qa)
#    id=voteProcessor.save_query(user_input_question=qa.question, user_input_answer=qa.answer, user_phone=qa.phone_number)
#    if id == -1:
#        raise HTTPException(status_code=status.HTTP_500_INTERNAL_SERVER_ERROR, detail="Internal server error")
#    else:
#        return {"id":id, "status": "ok"}
#
#@app.get("/v2/health-check")
#async def health_check():
#    if voteProcessor.health_check():
#        return {"status": "ok"}
#    else:
#        raise HTTPException(status_code=status.HTTP_500_INTERNAL_SERVER_ERROR)
class Program(BaseModel):
    codes: str = None
    requirements_txt: Optional[str] = None

@app.post("/v1/code-generation")
async def code_generation(program: Program):
    logger.info('Got 1 program: %s', program)
    cwd = os.getcwd()
    logger.info(f"Current Working Directory is: {cwd}")


@app.get("/v1/health-check")
async def health_check():
    return {"status": "ok"}

