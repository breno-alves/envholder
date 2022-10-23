package exporters

type Variable struct {
	Name  string
	Value string
}



type Exporter[D SSMOutput | DotenvOutput] interface {
	Export() ([]*Variable, error)
	Parse(acc []*Variable, data D) ([]*Variable, error)
}


func NewExporter(exporter string, arg1 string) Exporter[D SSMOutput | DotenvOutput] {
	switch exporter {
	case "ssm":
		return NewSSM(arg1, true)
	case "dotenv":
		return NewDotenv(arg1, true)
	default:
		panic("Invalid exporter")
	}
}

// func (instance *SSMExporter) Export() ([]*Variable, error) {
// 	return instance.Read()
// }

// func (instance *SSMExporter) Parse(acc []*Variable, data SSMOutput) ([]*Variable, error) {
// 	for idx := range data.Parameters {
// 		parameter := data.Parameters[idx]
// 		acc = append(acc, &Variable{
// 			Name:  *parameter.Name,
// 			Value: *parameter.Value,
// 		})
// 	}

// 	return acc, nil
// }
