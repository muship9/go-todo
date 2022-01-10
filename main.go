package main

import (
	"fmt"
	"go-todo/app/controllers"
	"go-todo/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
}
