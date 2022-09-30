package ssm

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

type OutputVarible struct {
	Name  string
	Value string
}

type SSM struct {
	client *ssm.SSM
	path   string
}

func NewSSM(path string) *SSM {
	return &SSM{
		client: ssm.New(session.Must(session.NewSession())),
		path:   path,
	}
}

func (*SSM) ParseVariables(acc []*OutputVarible, output *ssm.GetParametersByPathOutput) ([]*OutputVarible, error) {
	for _, p := range output.Parameters {
		acc = append(acc, &OutputVarible{
			Name:  *p.Name,
			Value: *p.Value,
		})
	}
	return acc, nil
}

func (instance *SSM) ExportVariables(recursive bool) (acc []*OutputVarible, err error) {
	input := &ssm.GetParametersByPathInput{
		Path:           aws.String(instance.path),
		WithDecryption: aws.Bool(true),
		Recursive:      aws.Bool(recursive),
	}

	output, err := instance.client.GetParametersByPath(input)
	if err != nil {
		return nil, err
	}

	acc, err = instance.ParseVariables(acc, output)
	if err != nil {
		return nil, err
	}

	nextToken := output.NextToken
	for nt := nextToken; nt != nil; nt = nextToken {
		input.NextToken = nt
		output, err = instance.client.GetParametersByPath(input)
		if err != nil {
			return nil, err
		}

		acc, err = instance.ParseVariables(acc, output)
		if err != nil {
			return nil, err
		}
		nextToken = output.NextToken
	}

	return acc, nil
}
