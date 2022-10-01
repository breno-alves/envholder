package transformers

import (
	"fmt"

	"github.com/breno-alves/envholder/pkg/exporters"
)

type DotenvTransformer struct {
}

func NewDotEnvOutput() *DotenvTransformer {
	return &DotenvTransformer{}
}

func (*DotenvTransformer) Transform(variable *exporters.Variable) string {
	output := fmt.Sprintf("%s=%s\n", variable.Name, variable.Value)
	return output
}
