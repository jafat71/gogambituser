package bd

import (
	"fmt"
	"gogambituser/models"
	"gogambituser/tools"

	_ "github.com/go-sql-driver/mysql"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Comienza Registro")
	err := DbConnect()
	if err != nil {
		return err
	}

	defer Db.Close()
	sentences := "INSERT INTO users (User_Email, User_UUID,User_DateAdd) VALUES ('" + sig.UserEmail + "','" + sig.UserUUID + "','" + tools.FechaMySQL() + "')"
	fmt.Println(sentences)

	_, err = Db.Exec(sentences)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("> Registro Exitoso")
	return nil
}
