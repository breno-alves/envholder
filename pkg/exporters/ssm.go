package exporters

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

type SSMOutput = *ssm.GetParametersByPathOutput

type SSM[D SSMOutput] struct {
	client    *ssm.SSM
	Path      string
	Recursive bool
}

func NewSSM(path string, recursive bool) Exporter[SSMOutput] {
	return &SSM[SSMOutput]{
		client:    ssm.New(session.Must(session.NewSession())),
		Path:      path,
		Recursive: recursive,
	}
}

func (instance *SSM[D]) Parse(acc []*Variable, ot D) ([]*Variable, error) {
	for _, p := range (*ot).Parameters {
		pathArr := strings.Split(*p.Name, "/")
		name := pathArr[len(pathArr)-1]
		value := *p.Value
		acc = append(acc, &Variable{
			Name:  name,
			Value: value,
		})
	}
	return acc, nil
}

func (instance *SSM[D]) Export() ([]*Variable, error) {
	acc := make([]*Variable, 0)
	input := &ssm.GetParametersByPathInput{
		Path:           aws.String(instance.Path),
		WithDecryption: aws.Bool(true),
		Recursive:      aws.Bool(instance.Recursive),
	}

	output, err := instance.client.GetParametersByPath(input)
	if err != nil {
		return nil, err
	}

	acc, err = instance.Parse(acc, output)
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

		acc, err = instance.Parse(acc, output)
		if err != nil {
			return nil, err
		}
		nextToken = output.NextToken
	}

	return acc, nil
}
