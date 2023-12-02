package main

import (
	"log"
	"time"

	"github.com/bambookim/acl-agent/acl-api-server/api"
	"github.com/gin-gonic/gin"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	Run()
}

func Run() {
	// database
	// database.Connect()

	// etcd
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Panic(err)
	}
	defer client.Close()

	/* test
	aclService := acl.NewAclService(client, nil)

	req := &acl.CreateAclRequest{
		Name:            "example",
		Action:          acl.ACL_ACTION_PERMIT,
		Direction:       acl.ACL_INGRESS,
		SourceCidr:      "10.1.0.0/16",
		DestCidr:        "10.2.0.0/16",
		SourcePortStart: 10,
		SourcePortStop:  10,
		DestPortStart:   11,
		DestPortStop:    11,
		Protocol:        acl.ACL_TCP,
	}

	if err := aclService.CreateAcl(req); err != nil {
		log.Panic(err)
	}
	*/

	router := gin.Default()
	apiGroup := router.Group("/api")
	api.AclControllerRoute(apiGroup, client)

	router.Run(":8080")
}
