package jenkins

import (
	"fmt"
	"log"
	"os"

	"encoding/json"
)

type Jenkins_env struct {
	User_id  string `json:"user_id"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

func (j *Jenkins_env) Setup() {
	file, err := os.Open("./jenkins/jenkins_env.json")
	if err != nil { // エラー処理
		log.Fatal(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(j)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*j)
}
