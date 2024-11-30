package config

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

/*
*
You can put 1 or more test cases in each test function
*
*/
func TestLambdaCodesPath(t *testing.T) {
	//TODO: Please check the variable name and value according to your requirement
	assert.Equal(t, LambdaCodesPath, "../TEMPLATE/lib/lambda-handler/index.py")
}
func TestRequirementsTxtPath(t *testing.T) {
	//TODO: Please check the variable name and value according to your requirement
	assert.Equal(t, RequirementsTxtPath, "../TEMPLATE/lambda_layer_lib/requirements.txt")
}
func TestAwsCdkVenvActivatePath(t *testing.T) {
	//TODO: Please check the variable name and value according to your requirement
	assert.Equal(t, AwsCdkVenvActivatePath, "../TEMPLATE/.venv/bin/activate")
}
func TestAwsCdkFolder(t *testing.T) {
	//TODO: Please check the variable name and value according to your requirement
	assert.Equal(t, AwsCdkFolder, "../TEMPLATE/")
}

//  lambda_codes_path: "../TEMPLATE/lib/lambda-handler/index.py"
//  requirements_txt_path: "../TEMPLATE/lambda_layer_lib/requirements.txt"
//  aws_cdk_venv_activate_path: "../TEMPLATE/.venv/bin/activate"
//  aws_cdk_folder: "../TEMPLATE/"
//var LambdaCodesPath string
//var RequirementsTxtPath string
//var AwsCdkVenvActivatePath string
//var AwsCdkFolder string

