package aclapiserver

import (
	"github.com/bambookim/acl-agent/acl-api-server/api"
	"github.com/gin-gonic/gin"
)

func main() {

}

func Run() {
	router := gin.Default()
	apiGroup := router.Group("/api")
	api.AclControllerRoute(apiGroup)

	router.Run(":8080")
}
