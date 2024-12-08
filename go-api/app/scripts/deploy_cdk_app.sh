#!/bin/bash

# verify the input parameters
if [ $# -eq 0 ]; then
	echo "Missing necessary folder name (AWS CDK FOLDER)"
	exit 1
elif [ $# -ne 1 ]; then
	echo "Too many input parameters"
	exit 1
fi	
cd $1
log_file=./deploy_cdk_app.log
echo "-----------------------------------------------" >> $log_file
echo "Deploying cdk app..."  >> $log_file
echo "go to folder: $1"      >> $log_file
base_folder=`basename $1`
source .venv/bin/activate    
cdk bootstrap                >> $log_file 2>&1
cdk synth                    >> $log_file 2>&1
cdk deploy  --require-approval never >> $log_file 2>&1
echo "Deployed cdk app... Done"  >> $log_file
echo "-----------------------------------------------" >> $log_file
endpoint=`grep https $log_file  | grep com |tail -1 | awk '{print $3}'`
echo $endpoint

