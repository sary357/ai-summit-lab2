import logging
logger = logging.getLogger(__name__)
from datetime import date
import os
from pydantic import BaseModel
from dotenv import load_dotenv
from typing import Optional

from fastapi import FastAPI, HTTPException, Request, Response, status
import json
from src.CodeGenManager import Manager

app = FastAPI()
load_dotenv()

ENCODEING="utf-8"

class Program(BaseModel):
    codes: str = None
    requirements_txt: Optional[str] = None

@app.post("/v1/code-generation")
async def code_generation(program: Program):
    logger.info('Got 1 program: %s', program)
    cwd = os.getcwd()
    logger.info(f"Current Working Directory is: {cwd}")
    code_gen_manager = Manager()
    code_gen_manager.save_all(program.codes, program.requirements_txt)
    code_gen_manager.execute()

@app.get("/v1/health-check")
async def health_check():
    return {"status": "ok"}
