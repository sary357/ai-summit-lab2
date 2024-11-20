from aws_cdk import (
    Stack,
    aws_apigateway as apigw,
    aws_lambda as _lambda
)
from constructs import Construct
import aws_cdk as cdk
import os,sys
import subprocess
import shutil
from enum import Enum
from .util import generate_unique_id
import boto3
from botocore.exceptions import ClientError
import os


LAMBDA_FUNCTION_NAME="MyFunction"
S3_BUCKET="fuming-ai-summit-lab-2025"
#S3_LOCATION="{}/{}".format(S3_BASE_LOCATION, LAMBDA_FUNCTION_NAME)
USER_LAMBDA_LIB_NAME="lambda_layer_lib"
USER_LAMBDA_REQ_FILE_NAME="requirements.txt"
VENV_FOLDER_NAME="venv"

class VenvExecStatus(Enum):
    NO_REQUIREMENTS_FILE = 1
    REQUIREMENTS_EXIST_AND_CREATE_SUCCESS = 2
    REQUIREMENTS_EXIST_BUT_FAILED = 3

class HelloLambdaStack(Stack):

    def __init__(self, scope: Construct, construct_id: str, **kwargs) -> None:
        super().__init__(scope, construct_id, **kwargs)

        # prepare necessary libraries with subprocess
        venv_status=self._generate_layer_lib()
        print(f"virtual environment created status: {venv_status}")
 
        # lambda
        if venv_status == VenvExecStatus.NO_REQUIREMENTS_FILE:
            fn = _lambda.Function(
                self,
                LAMBDA_FUNCTION_NAME,
                #runtime=_lambda.Runtime.PYTHON_3_12,
                runtime=_lambda.Runtime.PYTHON_3_9,
                handler="index.lambda_handler",
                timeout=cdk.Duration.minutes(1), # 3 minutes
                memory_size=10240, # max: 10240 MB
                code=_lambda.Code.from_asset("lib/lambda-handler")
            )
        elif venv_status == VenvExecStatus.REQUIREMENTS_EXIST_AND_CREATE_SUCCESS:
            s3_file=self._upload_zip_file_to_s3()
            # layer
            custom_lib_layer = _lambda.LayerVersion(self, "CustomLib",
                                                    layer_version_name="custom-lib",
                                                    compatible_runtimes = [_lambda.Runtime.PYTHON_3_9],
                                                    code = _lambda.Code.from_bucket(S3_BUCKET, s3_file)
                                                    )
            fn = _lambda.Function(
                self,
                LAMBDA_FUNCTION_NAME,
                #runtime=_lambda.Runtime.PYTHON_3_12,
                runtime=_lambda.Runtime.PYTHON_3_9,
                handler="index.lambda_handler",
                timeout=cdk.Duration.minutes(1), # 3 minutes
                memory_size=10240, # max: 10240 MB
                code=_lambda.Code.from_asset("lib/lambda-handler"),
                layers = [custom_lib_layer]
                )
        elif venv_status == VenvExecStatus.REQUIREMENTS_EXIST_BUT_FAILED:
            print("Failed to create virtual environment. Exit!!")
            return 

        # API gateway
        endpoint = apigw.LambdaRestApi(
            self,
            "ApiGwEndpoint",
            handler=fn,
            rest_api_name="HelloApi"
        )
    def _upload_zip_file_to_s3(self) -> str:
        current_work_dir=os.getcwd()
        uniq_id=generate_unique_id()
        local_zipped_venv_file="{}/{}/{}.zip".format(current_work_dir,USER_LAMBDA_LIB_NAME,VENV_FOLDER_NAME)
        s3_zipped_venv_file="{}-{}.zip".format(VENV_FOLDER_NAME, uniq_id)
        s3_client = boto3.client('s3')
        try:
            response = s3_client.upload_file(local_zipped_venv_file, S3_BUCKET, s3_zipped_venv_file)
        except ClientError as e:
            logging.error(e)
            return None
        return f"{s3_zipped_venv_file}"
    
 
    def _generate_layer_lib(self) -> VenvExecStatus:
        current_work_dir=os.getcwd()
        user_lambda_req_file="{}/{}/{}".format(current_work_dir,USER_LAMBDA_LIB_NAME,USER_LAMBDA_REQ_FILE_NAME)
        print(f"Requirements for user's lambda function: {user_lambda_req_file}")
        if not os.path.isfile(user_lambda_req_file):
            return VenvExecStatus.NO_REQUIREMENTS_FILE # no need to execute pip install
        else:
            # need to run execute pip install and check pip staus
            try:
                # Create virtual environment using subprocess
                venv_name = "{}/{}/{}".format(current_work_dir,USER_LAMBDA_LIB_NAME,VENV_FOLDER_NAME)
                print(f"Removing virtual environment if it exists: {venv_name}")
                if os.path.exists(venv_name):
                    shutil.rmtree(venv_name)
                print(f"Removing zipped virtual environment if it exists: {venv_name}.zip")
                if os.path.exists(f"{venv_name}.zip"):
                    os.remove(f"{venv_name}.zip")

                print(f"Creating virtual environment: {venv_name}")
                subprocess.run([sys.executable, "-m", "venv", venv_name], check=True)

                pip_path = os.path.join(venv_name, 'bin', 'pip')

                # If requirements are provided, install them
                print("Installing requirements...")
                subprocess.run([pip_path , "install", "-r", "{}".format(user_lambda_req_file)], check=True)

                # Create zip archive of the virtual environment
                print("Creating zip archive...")
                shutil.make_archive(f"{venv_name}", 'zip', venv_name)

                print(f"Virtual environment created and zipped successfully: {venv_name}.zip")
                return VenvExecStatus.REQUIREMENTS_EXIST_AND_CREATE_SUCCESS
            except subprocess.CalledProcessError as e:
                print(f"An error occurred during the process: {e}")
                return VenvExecStatus.REQUIREMENTS_EXIST_BUT_FAILED
            except Exception as e:
                print(f"An unexpected error occurred: {e}")
                return VenvExecStatus.REQUIREMENTS_EXIST_BUT_FAILED

