import logging
import os
from dotenv import load_dotenv
import subprocess
logger = logging.getLogger(__name__)

class Manager:
    def __init__(self):
        load_dotenv()
        logger.info("Manager initialized")
    def save_codes(self, codes):
        # get the path we have to save index.py
        lambda_codes_path=os.getenv('LAMBDA_CODES_PATH')
        logger.info("Saving codes to codes.txt")
        with open(lambda_codes_path, "w") as f:
            f.write((codes))
        logger.info(F"Codes saved to {lambda_codes_path}")
    def save_requirements_txt(self, requirements):
        requirements_txt=os.getenv('REQUIREMENTS_TXT_PATH')
        with open(requirements_txt, "w") as f:
            f.write((requirements))
        logger.info(f"requirements.txt saved to {requirements_txt}")
    def save_all(self, codes, requirements):
        self.save_codes(codes)
        if requirements:
            logger.info("requirements.txt found")
            self.save_requirements_txt(requirements)
    def execute(self):
        logger.info("Executing codes")
        # use aws cdk
        activate_path=os.getenv('AWS_CDK_VENV_ACTIVATE_PATH')
        aws_cdk_folder=os.getenv('AWS_CDK_FOLDER')
        cmd = f"source {activate_path} && cd {aws_cdk_folder} && cdk deploy --require-approval never"
        subprocess.run(cmd, shell=True, executable="/bin/bash")


