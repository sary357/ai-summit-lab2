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
LAMBDA_FUNCTION_NAME="MyFunction"
S3_BASE_LOCATION="s3://fuming-ai-summit-lab-2025/"
S3_LOCATION="{}/{}".format(S3_BASE_LOCATION, LAMBDA_FUNCTION_NAME)
USER_LAMBDA_LIB_NAME="lambda_layer_lib"
USER_LAMBDA_REQ_FILE_NAME="requirements.txt"
class HelloLambdaStack(Stack):

    def __init__(self, scope: Construct, construct_id: str, **kwargs) -> None:
        super().__init__(scope, construct_id, **kwargs)

        # prepare necessary libraries with subprocess
        zip_file_status=self._generate_layer_lib()
        print(zip_file_status)
        
        # lambda
        fn = _lambda.Function(
            self,
            LAMBDA_FUNCTION_NAME,
            runtime=_lambda.Runtime.PYTHON_3_12,
            handler="index.lambda_handler",
            timeout=cdk.Duration.minutes(1), # 3 minutes
            memory_size=10240, # max: 10240 MB
            code=_lambda.Code.from_asset("lib/lambda-handler")
        )

        # API gateway
        endpoint = apigw.LambdaRestApi(
            self,
            "ApiGwEndpoint",
            handler=fn,
            rest_api_name="HelloApi"
        )

    def _generate_layer_lib(self) -> bool:
        current_work_dir=os.getcwd()
        user_lambda_req_file="{}/{}/{}".format(current_work_dir,USER_LAMBDA_LIB_NAME,USER_LAMBDA_REQ_FILE_NAME)
        print(user_lambda_req_file)
        if not os.path.isfile(user_lambda_req_file):
            return True # no need to execute pip install
        else:
            # need to run execute pip install and check pip staus
            try:
                # Create virtual environment using subprocess
                venv_name = "{}/{}/venv".format(current_work_dir,USER_LAMBDA_LIB_NAME)
                print(f"Creating virtual environment: {venv_name}")
                subprocess.run([sys.executable, "-m", "venv", venv_name], check=True)

                pip_path = os.path.join(venv_name, 'bin', 'pip')

                # If requirements are provided, install them
                print("Installing requirements...")
                subprocess.run([pip_path + " install -r {}".format(user_lambda_req_file)], check=True)

                # Create zip archive of the virtual environment
                print("Creating zip archive...")
                shutil.make_archive(f"{venv_name}_backup", 'zip', venv_name)

                print(f"Virtual environment created and zipped successfully: {venv_name}_backup.zip")
                return True
            except subprocess.CalledProcessError as e:
                print(f"An error occurred during the process: {e}")
                return False
            except Exception as e:
                print(f"An unexpected error occurred: {e}")
                return False
