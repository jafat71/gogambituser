package awsgo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var Ctx context.Context
var Cfg aws.Config
var err error

func InitAWS() {
	Ctx = context.TODO() // Retorno un Nonull
	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-2"))
	if err != nil {
		panic("Error al cargar la config de .aws/config" + err.Error())
	}
}
