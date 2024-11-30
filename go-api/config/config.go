package config

import (
	"github.com/spf13/viper"
)

var Port int
var LambdaCodesPath string
var RequirementsTxtPath string
var AwsCdkVenvActivatePath string
var AwsCdkFolder string
var Host string

func init() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../config")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Error happned when reading config file：" + err.Error())
	}

	Port = viper.GetInt("application.port")                        // this is the port for gin server
        Host = viper.GetString("application.publicdomainname")
	// TODO: load config. Please check the following sample codes to load config
	LambdaCodesPath = viper.GetString("application.lambda_codes_path")
	RequirementsTxtPath = viper.GetString("application.requirements_txt_path")
	AwsCdkVenvActivatePath = viper.GetString("application.aws_cdk_venv_activate_path")
	AwsCdkFolder = viper.GetString("application.aws_cdk_folder")

}
