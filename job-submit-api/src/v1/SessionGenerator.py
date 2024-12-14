from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker
import os
import logging
logger = logging.getLogger(__name__)
class sessionGenerator:
    _instance = None 
    def __new__(cls, *args, **kwargs): 
        if cls._instance is None: 
            cls._instance = super().__new__(cls) 
            cls._instance.engine = create_engine(cls._instance._get_db_conn_str(), echo=True)
        return cls._instance 
         
    def _get_db_conn_str(self)->str:
        if self._get_db_conn_str_from_vault():
            logger.info("Get database connection string from vault")
            return self._get_db_conn_str_from_vault()
        elif self._get_db_conn_str_from_env():
            logger.info("Get database connection string from environment variable (DB_CONN_STR)")
            return self._get_db_conn_str_from_env()
        logger.info("Get database connection string from default setting (local database)")
        return self._get_local_db_conn_str()
    
    def _get_db_conn_str_from_vault(self)->str:
        # TODO: please replace this part to get the database connection string from vault
        return None

    def _get_db_conn_str_from_env(self)->str:
        return os.environ['DB_CONN_STR'].strip() if 'DB_CONN_STR' in os.environ else None
    
    def _get_local_db_conn_str(self)->str:
        return 'postgresql+psycopg2://postgres:CHANGE_ME@localhost:5432/job_submit'

    def get_session(self):
        return sessionmaker(bind=self.engine)()

