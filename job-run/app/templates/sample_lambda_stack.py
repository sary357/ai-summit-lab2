from aws_cdk import (
    Stack,
    aws_apigateway as apigw,
    aws_lambda as _lambda,
    RemovalPolicy
)
from constructs import Construct
import aws_cdk as cdk
import aws_cdk.aws_s3 as s3
import os,sys
import subprocess
import shutil
from enum import Enum
# import boto3
# from botocore.exceptions import ClientError
import os
import datetime
import random
import string

LAMBDA_FUNCTION_NAME="SampleLambdaStack"
USER_LAMBDA_LIB_NAME="lambda_layer_lib"
USER_LAMBDA_REQ_FILE_NAME="requirements.txt"
VENV_FOLDER_NAME="python"
RUNTIME_VERSION=_lambda.Runtime.PYTHON_3_9
class VenvExecStatus(Enum):
    NO_REQUIREMENTS_FILE = 1
    REQUIREMENTS_EXIST_AND_CREATE_SUCCESS = 2
    REQUIREMENTS_EXIST_BUT_FAILED = 3

class SampleLambdaStack(Stack):

    def __init__(self, scope: Construct, construct_id: str, **kwargs) -> None:
        super().__init__(scope, construct_id, **kwargs)

        # prepare necessary libraries with subprocess
        venv_status=self._generate_layer_lib()
        print("virtual environment created status: {}, {}".format(venv_status[0], venv_status[1]))

        # lambda
        if venv_status[0] == VenvExecStatus.NO_REQUIREMENTS_FILE:
            fn = _lambda.Function(
                self,
                LAMBDA_FUNCTION_NAME,
                #runtime=_lambda.Runtime.PYTHON_3_12,
                runtime=RUNTIME_VERSION,
                handler="index.lambda_handler",
                timeout=cdk.Duration.minutes(3), # 3 minutes
                memory_size=10240, # max: 10240 MB
                code=_lambda.Code.from_asset("lib/lambda-handler")
            )
        elif venv_status[0] == VenvExecStatus.REQUIREMENTS_EXIST_AND_CREATE_SUCCESS:
            local_venv_path=self._get_local_venv_path()
            print(f"Local venv: {local_venv_path}")
            # layer
            custom_lib_layer = _lambda.LayerVersion(self, "CustomLib",
                                                    #layer_version_name="custom-lib",
                                                    description="custom python packages",
                                                    compatible_runtimes = [RUNTIME_VERSION],
                                                    #code = _lambda.Code.from_bucket(s3_lib_bucket, s3_file)
                                                    #code = _lambda.Code.from_asset(local_venv_path),
                                                    code = _lambda.Code.from_asset(venv_status[1]),
                                                    compatible_architectures=[_lambda.Architecture.X86_64],
                                                    removal_policy=RemovalPolicy.DESTROY
                                                    )
            fn = _lambda.Function(
                self,
                LAMBDA_FUNCTION_NAME,
                #runtime=_lambda.Runtime.PYTHON_3_12,
                runtime=RUNTIME_VERSION,
                handler="index.lambda_handler",
                timeout=cdk.Duration.minutes(3), # 3 minutes
                memory_size=10240, # max: 10240 MB
                code=_lambda.Code.from_asset(os.getcwd()+"/lib/lambda-handler"),
                layers = [custom_lib_layer]
                )
        elif venv_status[0] == VenvExecStatus.REQUIREMENTS_EXIST_BUT_FAILED:
            print("Failed to create virtual environment. Exit!!")
            return 

        # API gateway
        enabled=True
        if enabled:
            endpoint = apigw.LambdaRestApi(
                self,
                "SampleLambdaStack_api",
                handler=fn,
                rest_api_name="SampleLambdaStack"
            )
    def _get_local_venv_path(self):
        current_work_dir=os.getcwd()
        local_zipped_venv="{}/{}/{}".format(current_work_dir,USER_LAMBDA_LIB_NAME,VENV_FOLDER_NAME)
        return local_zipped_venv
 
    def _generate_layer_lib(self) -> list:
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
                user_lambda_lib="{}/{}".format(current_work_dir,USER_LAMBDA_LIB_NAME)
                print(f"Removing virtual environment if it exists: {venv_name}")
                if os.path.exists(venv_name):
                    shutil.rmtree(venv_name)

                print(f"Creating virtual environment: {venv_name}")
                subprocess.run([sys.executable, "-m", "venv", venv_name], check=True)

                pip_path = os.path.join(venv_name, 'bin', 'pip')

                # If requirements are provided, install them
                print("Installing requirements...")
                subprocess.run([pip_path , "install", "-r", "{}".format(user_lambda_req_file), "--target={}".format(venv_name)], check=True)

                # Create zip archive of the virtual environment
                print("Creating zip archive...")
                venv_folder = "/tmp/" + generate_random_number() + "/"
                shutil.make_archive(base_name=venv_folder+VENV_FOLDER_NAME, format='zip', root_dir=user_lambda_lib)

                print(f"Virtual environment created and zipped successfully: {venv_folder}/{VENV_FOLDER_NAME}.zip")
                return (VenvExecStatus.REQUIREMENTS_EXIST_AND_CREATE_SUCCESS, f"{venv_folder}/{VENV_FOLDER_NAME}.zip")
            except subprocess.CalledProcessError as e:
                print(f"An error occurred during the process: {e}")
                return (VenvExecStatus.REQUIREMENTS_EXIST_BUT_FAILED, None)
            except Exception as e:
                print(f"An unexpected error occurred: {e}")
                return (VenvExecStatus.REQUIREMENTS_EXIST_BUT_FAILED, None)


def generate_random_number()->str:
    """Generates a random number in the specified format.

    Returns:
        str: A string representing the random number in the format "app-YYYYMMDDhhmmss-{NUMBER}".
    """

    now = datetime.datetime.now()
    year = now.strftime("%Y")
    month = now.strftime("%m")
    day = now.strftime("%d")
    time = now.strftime("%H%M%S")
    random_number = str(random.randint(10000000, 99999999))

    return f"app-{year}{month}{day}{time}-{random_number}"
