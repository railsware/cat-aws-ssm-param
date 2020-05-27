package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: cat-aws-ssm-param [parameter-name-with-full-path]")
		os.Exit(1)
	}

	paramName := os.Args[1]

	session, err := session.NewSession()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating AWS session: %v\n", err)
		os.Exit(1)
	}

	ssmClient := ssm.New(session)

	getParamResponse, err := ssmClient.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(paramName),
		WithDecryption: aws.Bool(true),
	})

	if err != nil {
		if _, ok := err.(*ssm.ParameterNotFound); ok {
			fmt.Fprintln(os.Stderr, "Parameter not found")
		} else {
			fmt.Fprintf(os.Stderr, "Error retrieving parameter: %v\n", err)
		}
		os.Exit(1)
	}

	os.Stdout.WriteString(*getParamResponse.Parameter.Value)
}
