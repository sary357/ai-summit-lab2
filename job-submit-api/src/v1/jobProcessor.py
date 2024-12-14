from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy import Column, Integer, String,DateTime
from sqlalchemy import select
from datetime import datetime
from src.v1 import SessionGenerator
from sqlalchemy.orm.exc import NoResultFound
import logging
import copy

from src.v1.const import *

Base = declarative_base()
logger = logging.getLogger(__name__)

'''
JOB_PENDING=0
JOB_RUNNING=1
JOB_SUCCESS=2
JOB_FAILED=3

ERR_INTERNAL_ERROR=-101

'''

class job(Base):
    __tablename__ = 'jobs'
    '''
    create table job_submit.public.jobs (
    id SERIAL,
    codes text not null,
    requirements_txt text,
    status int not null default 0,
    endpoint varchar(65535),
    lock_executor varchar(65535),
    running_executor varchar(65536),
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp,
    PRIMARY KEY (id)
     )  TABLESPACE pg_default;

     '''

    id = Column(Integer, primary_key=True)
    codes = Column(String(65535),nullable=False)
    requirements_txt = Column(String(65535))
    status = Column(Integer(), default=JOB_ACCEPTED)
    endpoint = Column(String(65535))
    lock_executor = Column(String(65535))
    running_executor = Column(String(65535))
    created_at = Column(DateTime, default=datetime.utcnow)
    updated_at = Column(DateTime, default=datetime.utcnow, onupdate=datetime.utcnow)


def save_job(input_codes:str, input_requirements_txt:str, input_status:int=JOB_ACCEPTED)->int:
    user_job = job(codes = input_codes, requirements_txt=input_requirements_txt, status=input_status)
    session = SessionGenerator.sessionGenerator().get_session()
    try:
        session.add(user_job)
        session.commit()
        user_job_id = user_job.id
        return user_job_id
    except Exception as e:
        logger.error(e.__class__.__name__)
        logger.error("Error in save_job: "+str(e))
        return ERR_INTERNAL_ERROR
    finally:
        session.close()

def lock_job(lock_executor: str)-> job:
    session = SessionGenerator.sessionGenerator().get_session()
    try:
        user_job_to_update = session.query(job) \
                                    .filter_by(status=JOB_ACCEPTED) \
                                    .order_by(job.id).first()

        if user_job_to_update is None:
            logger.error("Cannot locate any job need to be executed.")
            return None

        user_job_to_update.status = JOB_LOCKING
        user_job_to_update.lock_executor = lock_executor
        ret_job=copy.deepcopy(user_job_to_update)
        session.commit()
        return ret_job

    except IntegrityError as e:
        # Handle potential race condition where another transaction updates
        # the job before the lock can be acquired
        logger.error("Error acquiring lock for the lock job: " + str(e))
        session.rollback()
        return None

    except Exception as e:
        logger.error(e.__class__.__name__)
        logger.error("Error in lock_job: " + str(e))
        session.rollback()
        return None

    finally:
        session.close()

def execute_job(id: int, running_executor:str)->bool:
    session = SessionGenerator.sessionGenerator().get_session()
    try:
        user_job_to_update = session.query(job) \
                                    .filter_by(id=id, lock_executor=running_executor, status=JOB_LOCKING) \
                                    .with_for_update(nowait=True) \
                                    .one_or_none()

        if user_job_to_update is None:
            logger.error("Cannot find the job id: \"" + str(id) + "\" or job is not in \"JOB_LOCKING\"")
            return False

        user_job_to_update.status = JOB_EXECUTING
        user_job_to_update.running_executor = running_executor
        session.commit()
        return True

    except IntegrityError as e:
        # Handle potential race condition where another transaction updates
        # the job before the lock can be acquired
        logger.error("Error acquiring lock for job ID " + str(id) + ": " + str(e))
        session.rollback()
        return False

    except Exception as e:
        logger.error(e.__class__.__name__)
        logger.error("Error in executr_job: " + str(e))
        session.rollback()
        return False

    finally:
        session.close()

def finish_job(id: int, running_executor:str, status: int, endpoint:str=None)-> bool:
    session = SessionGenerator.sessionGenerator().get_session()
    try:
        user_job_to_update = session.query(job) \
                                    .filter_by(id=id, lock_executor=running_executor, running_executor=running_executor, status=JOB_EXECUTING) \
                                    .with_for_update(nowait=True) \
                                    .one_or_none()

        if user_job_to_update is None:
            logger.error("Cannot find the job id: \"" + str(id) + "\" or job is not in \"JOB_EXECUTING\"")
            return False
        if status not in (JOB_SUCCESS, JOB_FAILED):
            logger.error("Not a valid status code: "+str(status)+"/job id: \"" + str(id) + "\"")
            return False
        user_job_to_update.status = status
        user_job_to_update.endpoint = endpoint
        session.commit()
        return True

    except IntegrityError as e:
        # Handle potential race condition where another transaction updates
        # the job before the lock can be acquired
        logger.error("Error acquiring lock for job ID " + str(id) + ": " + str(e))
        session.rollback()
        return False

    except Exception as e:
        logger.error(e.__class__.__name__)
        logger.error("Error in finish_job: " + str(e))
        session.rollback()
        return False

    finally:
        session.close()

def success_job(id: int, running_executor:str,endpoint)-> bool:
    return finish_job(id, running_executor, JOB_SUCCESS, endpoint)

def fail_job(id: int, running_executor:str)-> bool:
    return finish_job(id, running_executor, JOB_FAILED)

def get_job_by_id(id:int)-> job:
    session = SessionGenerator.sessionGenerator().get_session()
    try:
        user_job=session.query(job).filter_by(id=id).one()
        return user_job
    except NoResultFound as e:
        logger.error("Cannot find the id: \""+str(id)+"\"")
        return None
    except Exception as e:
        logger.error(e.__class__.__name__)
        logger.error("Error in get_job_by_id: "+str(e))
        return None
    finally:
        session.close()

def get_job_status_by_id(id:int)-> int:
    user_job = get_job_by_id(id)
    if user_job is not None:
        return user_job.status
    else:
        return ERR_INTERNAL_ERROR

def health_check()->bool:
    session = SessionGenerator.sessionGenerator().get_session()
    try:
        session.execute(select(1))
        return True
    except Exception as e:
        logger.error("Error in health check: "+str(e))
        return False
    finally:
        session.close()


