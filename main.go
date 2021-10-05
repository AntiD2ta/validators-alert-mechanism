package main

import (
	"context"
	"fmt"
	"sort"

	initConfig "github.com/AntiD2ta/validators-alert-mechanism/init"
	"github.com/AntiD2ta/validators-alert-mechanism/internal"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, in interface{}) {
	config, err := initConfig.Init()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	validators := make(chan *internal.Validator)
	defer close(validators)

	for _, validator := range config.Validators {
		go internal.GetData(validator, config.Interval, validators)
	}

	var output []string
	for i := 0; i < len(config.Validators); i++ {
		v := <-validators
		if v == nil {
			continue
		}
		output = append(output, fmt.Sprintf("Validator %s has validated %d blocks in %d minutes", v.Name, v.MinedBlocks, config.Interval))
	}

	sort.Strings(output)
	for _, o := range output {
		fmt.Println(o)
	}
}

func main() {
	lambda.Start(HandleRequest)
}
