package exporters

type Variable struct {
	Name  string
	Value string
}

type Exporter[D SSMOutput | any] interface {
	Export() ([]*Variable, error)
	Parse(acc []*Variable, data D) ([]*Variable, error)
}

func NewExporter(exporter string, arg1 string) Exporter[SSMOutput] {
	switch exporter {
	case "ssm":
		return NewSSM(arg1, true)
	default:
		return nil
	}
}
