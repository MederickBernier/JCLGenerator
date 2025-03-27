package internals

import (
	"JCLGenerator/jcl"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadYAML(path string) (*jcl.JCLParameters, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var params jcl.JCLParameters
	err = yaml.Unmarshal(data, &params)
	if err != nil {
		return nil, err
	}

	return &params, nil
}
