package main

import (
	"log"

	"github.com/sclevine/agouti"
	"github.com/tjtjtjtj/go-headless/jenkins"
)

var id = "xxxx"
var passwd = "xxxx"

func main() {
	jenkins_env := new(jenkins.Jenkins_env)
	jenkins_env.Setup()

	driver := agouti.ChromeDriver(agouti.Desired(agouti.Capabilities{
		"chromeOptions": map[string][]string{
			"args": []string{
				"headless",
			},
		},
	}))
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage()
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}

	if err := page.Navigate("http://192.168.20.41:8080/login?from=%2F"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}

	page.Screenshot("/tmp/jenkins1.jpg")

	identity := page.FindByID("j_username")
	password := page.FindByName("j_password")
	identity.Fill(id)
	password.Fill(passwd)
	if err := page.FindByID("yui-gen1-button").Submit(); err != nil {
		log.Fatalf("Failed to login:%v", err)
	}

	page.Screenshot("/tmp/jenkins2.jpg")
}
