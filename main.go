package main

import (
	"fmt"
	"math/rand"

	dbpostgresql "github.com/jullmi/dbPostgreSQL"
)

var MIN = 0
var MAX = 26

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func getString(length int64) string {
	starChart := "A"
	temp := ""
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(starChart[0] + byte(myRand))
		temp = temp + newChar
		if i == length {
			break
		}
		i++
	}
	return temp
}

func main() {
	dbpostgresql.Hostname = "localhost"
	dbpostgresql.Port = 5432
	dbpostgresql.Username = "postgres"
	dbpostgresql.Password = "postgres"
	dbpostgresql.Database = "go"
	dbpostgresql.Sslmode = "disable"

	randomUsername := getString(5)

	temp := dbpostgresql.Userdata{
		Username:    randomUsername,
		Name:        "Julia",
		Surname:     "Panova",
		Description: "Hello, It's me!",
	}

	id := dbpostgresql.AddUser(temp)
	if id == -1 {
		fmt.Println("There was an error adding user", temp.Username)
	}


	err := dbpostgresql.DeleteUser(id)
	if err != nil {
		fmt.Println(err)
	}

	//try delete user again
	err = dbpostgresql.DeleteUser(id)
	if err != nil {
		fmt.Println(err)
	}

	id = dbpostgresql.AddUser(temp)
	if id == -1 {
		fmt.Println("There was an error adding user", temp.Username)
	}


	temp = dbpostgresql.Userdata{
		Username:    randomUsername,
		Name:        "Julia",
		Surname:     "Panova",
		Description: "This minght not be me!",
	}

	err = dbpostgresql.UpdateUser(temp)
	if err != nil {
		fmt.Println(err)
	}

	data, err := dbpostgresql.ListUsers()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range data {
		fmt.Println(v)
	}

}
