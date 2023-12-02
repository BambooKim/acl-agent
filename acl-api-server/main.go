package main

import (
	"log"
	"time"

	"github.com/bambookim/acl-agent/acl-api-server/api"
	"github.com/bambookim/acl-agent/acl-api-server/domain/acl"
	"github.com/bambookim/acl-agent/acl-api-server/global/database"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	// Init Config
	config := initConfig()

	// database
	database.Connect(config.Datasource)
	database.DB.AutoMigrate(&acl.AclEntity{})

	// etcd
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Panic(err)
	}
	defer client.Close()

	Run(client)
}

func Run(client *clientv3.Client) {

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

func initConfig() *Config {
	config := &Config{}

	viper.AddConfigPath(".")
	viper.SetConfigName("config-local")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Panic(err)
	}
	if err := viper.Unmarshal(config); err != nil {
		log.Panic(err)
	}

	return config
}

type Config struct {
	Datasource *database.Datasource `yaml:"datasource"`
}
