package secretm

import (
	"encoding/json"
	"fmt"
	"gogambituser/awsgo"
	"gogambituser/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(nombreSecret string) (models.SecretRDSJson, error) {
	var datosSecret models.SecretRDSJson
	fmt.Println(" > Load Secret " + nombreSecret)
	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nombreSecret),
	})
	if err != nil {
		fmt.Println(err.Error())
		return datosSecret, err
	}
	//Procesar -> lo que devuelve secret manager
	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println(" > Read Secret Success")
	return datosSecret, nil
}
