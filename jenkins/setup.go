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

func Setup(jenkins_env *Jenkins_env) {
	file, err := os.Open("./jenkins/jenkins_env.json")
	if err != nil { // エラー処理
		log.Fatal(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(jenkins_env)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*jenkins_env)
}
