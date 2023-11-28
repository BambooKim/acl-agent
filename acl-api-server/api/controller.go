package api

import (
	"github.com/bambookim/acl-agent/acl-api-server/domain/acl"
	"github.com/gin-gonic/gin"
)

func AclControllerRoute(rg *gin.RouterGroup) {
	aclRepository := acl.NewAclRepository()
	aclService := acl.NewAclService(&aclRepository)
	aclController := NewAclController(&aclService)

	aclRouterGroup := rg.Group("/acl")
	aclRouterGroup.GET("", aclController.GetAclList)
	aclRouterGroup.GET("/:index", aclController.GetAclByIndex)
	aclRouterGroup.POST("", aclController.CreateAcl)
	aclRouterGroup.PUT("/:id", aclController.ModifyAcl)
	aclRouterGroup.DELETE("/:id", aclController.DeleteAcl)
}

type AclController interface {
	GetAclList(c *gin.Context)    // acl 목록 조회
	GetAclByIndex(c *gin.Context) // acl 단건 조회
	CreateAcl(c *gin.Context)     // acl 추가
	ModifyAcl(c *gin.Context)     // acl 수정
	DeleteAcl(c *gin.Context)     // acl 삭제
}

type AclControllerImpl struct {
	acl.AclService
}

func NewAclController(aclService *acl.AclService) AclController {
	return &AclControllerImpl{
		AclService: aclService,
	}
}

func (ci *AclControllerImpl) GetAclList(c *gin.Context) {

}

func (ci *AclControllerImpl) GetAclByIndex(c *gin.Context) {

}

func (ci *AclControllerImpl) CreateAcl(c *gin.Context) {

}

func (ci *AclControllerImpl) ModifyAcl(c *gin.Context) {

}

func (ci *AclControllerImpl) DeleteAcl(c *gin.Context) {

}
