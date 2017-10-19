package jenkins

import (
	"log"
	"os/exec"

	"github.com/sclevine/agouti"
)

func Initial(page *agouti.Page) {
	out, err := exec.Command("cat", "/tmp/initialAdminPassword").Output()
	if err != nil {
		log.Fatal("cat password", err)
	}
	page.Screenshot("/tmp/jenkinspass1.jpg")

	password := page.FindByID("security-token")
	password.Fill(string(out))
	if err := page.FindByClass("set-security-key").Submit(); err != nil {
		log.Fatal("Failed to set password", err)
	}

	page.Screenshot("/tmp/jenkinspass2.jpg")

}
