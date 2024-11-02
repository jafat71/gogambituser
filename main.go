package main

import (
	"context"
	"errors"
	"fmt"
	"os" //para manejo de .env

	"gogambituser/awsgo"
	"gogambituser/bd"
	"gogambituser/models"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(ExecLambda)
}

func ExecLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InitAWS()

	if !ValidateParams() {
		fmt.Println(".env no provee 'Secret Name'")
		err := errors.New("error en los params - .env 'Secret Name' ")
		return event, err
	}

	var datos models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
			fmt.Println("Email = " + datos.UserEmail)
		case "sub":
			datos.UserUUID = att
			fmt.Println("Sub  = " + datos.UserUUID)
		}
	}

	err := bd.ReadSecret()
	if err != nil {
		fmt.Println("Error al leer el Secret " + err.Error())
		return event, err
	}

	err = bd.SignUp(datos)
	return event, err
}

// Validate Params from .env
func ValidateParams() bool {
	var hasParam bool
	_, hasParam = os.LookupEnv("SecretName")
	return hasParam
}
