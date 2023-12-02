package main

import (
	"log"

	"github.com/bambookim/acl-agent/acl-agent/app"
	"go.ligato.io/cn-infra/v2/agent"
)

func main() {
	aclAgent := app.NewAclAgent()

	a := agent.NewAgent(
		agent.AllPlugins(aclAgent),
	)
	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}
