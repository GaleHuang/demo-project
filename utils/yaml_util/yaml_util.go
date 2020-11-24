package yamlutil

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

/*
	enableEnv: 开启环境变量功能，将配置文件中的${XXX}替换为OS中的名为XXX的环境变量值
 */
func UnmarshalFromFile(filePath string, v interface{}, enableEnv bool) error{

	file, err := ioutil.ReadFile(filePath)
	if err != nil{
		return err
	}
	fileStr := string(file)
	var parsedFileStr string
	if enableEnv{
		parsedFileStr =  ParseEnvFileStr(fileStr)
	}

	return yaml.Unmarshal([]byte(parsedFileStr), v)

}

/*
	将字符串中的${}替换为OS中的环境变量值
 */
func ParseEnvFileStr(fileStr string) string  {
	envStrings := findEnvStrings(fileStr)
	envSet := NewEnvSet()
	
	for _, env := range envStrings{
		if envSet.Contains(env){
			fileStr = strings.Replace(fileStr, env, ParseEnvStrToValue(env), -1)
			envSet.Add(env)
		}
	}
	return fileStr
}


type EnvSet struct {
	Container map[string]string
}

func NewEnvSet() *EnvSet  {
	return &EnvSet{
		Container: map[string]string{},
	}
}

func (e *EnvSet) Contains(str string) bool{
	if _, res := e.Container[str]; res{
		return true
	}
	return false
}

func (e *EnvSet) Add(str string)  {
	e.Container[str] = ""
}

func ParseEnvStrToValue(envStr string) string {
	r := regexp.MustCompile(`\${.*}`)
	if !r.MatchString(envStr){
		return ""
	}
	value, exists := os.LookupEnv(envStr[2:len(envStr)-1])
	if !exists{
		return envStr
	}
	return value
}



func findEnvStrings(fileStr string) []string{
	r := regexp.MustCompile(`\${.*}`)
	return r.FindAllString(fileStr, -1)
}

