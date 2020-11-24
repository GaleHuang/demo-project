package test

import (
	"github.com/galehuang/demo-project/config"
	"github.com/galehuang/demo-project/utils/yaml_util"
	"testing"
)

func Test_Yaml_Util(t *testing.T){
	t.Logf("%v", yamlutil.ParseEnvStrToValue("${GOPATH}"))
	c := &config.Config{}
	err := yamlutil.UnmarshalFromFile(
		"C:\\Users\\dafenghuang\\Desktop\\workstation\\Go\\src\\github.com\\galehuang\\demo-project\\conf\\config.yaml", c, true)
	if err != nil{
		t.Error(err.Error())
	}
	t.Logf("%v", c)
}
