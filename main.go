package main

import (
	"log"
	"strings"

	"github.com/sclevine/agouti"
	"github.com/tjtjtjtj/go-headless/jenkins"
)

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

	page_html, _ := page.HTML()
	if strings.Contains(page_html, "Unlock Jenkins") {
		jenkins.Initial(page)
	} else {

		page.Screenshot("/tmp/jenkins1.jpg")

		identity := page.FindByID("j_username")
		password := page.FindByName("j_password")
		identity.Fill(jenkins_env.User_id)
		password.Fill(jenkins_env.Password)
		if err := page.FindByID("yui-gen1-button").Submit(); err != nil {
			log.Fatalf("Failed to login:%v", err)
		}

		page.Screenshot("/tmp/jenkins2.jpg")
	}
}
