package acl

type AclRepository interface {
}

type AclRepositoryImpl struct{}

func NewAclRepository() AclRepository {
	return &AclRepositoryImpl{}
}
