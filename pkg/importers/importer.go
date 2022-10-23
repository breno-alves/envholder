package importers

type Variable struct {
	Name  string
	Value string
}

type Importer interface{
	Read() ([]*Variable, error)
}

func NewImporter(importer string, arg1 string) Importer {
	switch importer {
	case "dotenv":
		return NewDotenvImporter(arg1)
	default:
		return nil
	}
}
