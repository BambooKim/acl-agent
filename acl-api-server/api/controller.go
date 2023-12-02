package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/bambookim/acl-agent/acl-api-server/domain/acl"
	"github.com/gin-gonic/gin"
	clientv3 "go.etcd.io/etcd/client/v3"
	"gorm.io/gorm"
)

var (
	ERROR_CODE = map[string]int{
		"INVALIDATE_PROTOCOL":          http.StatusBadRequest,
		gorm.ErrRecordNotFound.Error(): http.StatusBadRequest,
	}
)

func AclControllerRoute(rg *gin.RouterGroup, etcdClient *clientv3.Client) {
	aclRepository := acl.NewAclRepository()
	aclService := acl.NewAclService(etcdClient, &aclRepository)
	aclController := NewAclController(&aclService)

	aclRouterGroup := rg.Group("/acl")
	aclRouterGroup.GET("", aclController.GetAclList)
	// aclRouterGroup.GET("/:index", aclController.GetAclByIndex)
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
		AclService: *aclService,
	}
}

func (ci *AclControllerImpl) GetAclList(c *gin.Context) {
	acls, err := ci.AclService.GetAclList()
	if err != nil {
		errorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, acls)
}

func (ci *AclControllerImpl) GetAclByIndex(c *gin.Context) {

}

func (ci *AclControllerImpl) CreateAcl(c *gin.Context) {
	reqDto := &acl.CreateAclRequest{}
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		errorResponse(c, err)
		return
	}
	if err := json.Unmarshal(body, reqDto); err != nil {
		errorResponse(c, err)
		return
	}

	if err := ci.AclService.CreateAcl(reqDto); err != nil {
		errorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, "created")
}

func (ci *AclControllerImpl) ModifyAcl(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, err)
		return
	}

	reqDto := &acl.ModifyAclRequest{}
	body, err := io.ReadAll(c.Request.Body)
	if err := json.Unmarshal(body, reqDto); err != nil {
		errorResponse(c, err)
		return
	}

	if err := ci.AclService.ModifyAcl(id, reqDto); err != nil {
		errorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, "modified")
}

func (ci *AclControllerImpl) DeleteAcl(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, err)
		return
	}
	if err := ci.AclService.DeleteAcl(id); err != nil {
		errorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, "deleted")
}

func errorResponse(c *gin.Context, err error) {
	c.AbortWithStatusJSON(getStatusCode(err), err.Error())
}

func getStatusCode(err error) int {
	code, ok := ERROR_CODE[err.Error()]
	if !ok {
		return http.StatusInternalServerError
	}

	return code
}
