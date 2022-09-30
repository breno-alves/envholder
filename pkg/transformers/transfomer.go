package transformers

import "github.com/breno-alves/envholder/pkg/exporters"

type Transfomer interface {
	Transform(variable *exporters.Variable) string
}

func NewTransformer(format string) Transfomer {
	switch format {
	case "dotenv":
		return NewDotEnvOutput()
	default:
		return nil
	}
}
