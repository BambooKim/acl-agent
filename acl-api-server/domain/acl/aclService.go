package acl

type AclService interface {
}

type AclServiceImpl struct {
	AclRepository
}

func NewAclService(aclRepository *AclRepository) AclService {
	return &AclServiceImpl{
		AclRepository: aclRepository,
	}
}
