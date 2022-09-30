package transformers

type OutputVarible struct {
	Name  string
	Value string
}

type Outputer interface {
	Transform(variable *OutputVarible) string
}

func NewOutputer(format string) Outputer {
	switch format {
	case "dotenv":
		return NewDotEnvOutput()
	default:
		return nil
	}
}
