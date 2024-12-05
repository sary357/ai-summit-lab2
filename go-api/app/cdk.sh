#!/bin/bash


if [ $# -eq 0 ]; then
	echo "Missing necessary folder name"
	exit 1
elif [ $# -ne 1 ]; then
	echo "Too many input parameters"
	exit 1
else
	echo "start to generate AWS CDK folder: $1"
fi	

cd $1

echo "go to folder: $1"

cdk init app --language python

source .venv/bin/activate

python -m pip install -r requirements.txt


