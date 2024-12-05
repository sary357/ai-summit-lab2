#!/bin/bash

# verify the input parameters
if [ $# -eq 0 ]; then
	echo "Missing necessary folder name"
	exit 1
elif [ $# -ne 1 ]; then
	echo "Too many input parameters"
	exit 1
else
	echo "start to generate AWS CDK folder: $1"
fi	

TMP_FOLDER=/tmp
REPLACE_TARGET="SampleLambdaStack"

# generate python stack file
echo "1. generating python stack file..."
current_dir=$(dirname "$0")
template_py=$current_dir/template/sample_lambda_stack.py
tmp_python_stack_file=`basename $1`
generated_python_stack_file="${tmp_python_stack_file//-/_}"
cp $(dirname "$0")/template/sample_lambda_stack.py $TMP_FOLDER/$generated_python_stack_file\_stack.py
replacement_string="${tmp_python_stack_file//-/}"Stack
sed  -i "s/$REPLACE_TARGET/$replacement_string/g" $TMP_FOLDER/$generated_python_stack_file\_stack.py
sed  -i "s/class\ app/class\ App/g" $TMP_FOLDER/$generated_python_stack_file\_stack.py
echo "1. generating python stack file: Done..."

# generate aws cdk files
echo "2. generating aws cdk files..."
mkdir -p $1
cd $1
echo "go to folder: $1"
cdk init app --language python
source .venv/bin/activate
python -m pip install -r requirements.txt
echo "2. generating aws cdk files: Done..."

# copy python stack file to aws cdk folder and replace the corresponding string ("SampleLambdaStack")
cp $TMP_FOLDER/$generated_python_stack_file\_stack.py ./$generated_python_stack_file/
