package transformers

import (
	"fmt"
)

type DotenvOutput struct {
}

func NewDotEnvOutput() *DotenvOutput {
	return &DotenvOutput{}
}

func (*DotenvOutput) Transform(variable *OutputVarible) string {
	output := fmt.Sprintf("%s=%s\n", variable.Name, variable.Value)
	return output
}
