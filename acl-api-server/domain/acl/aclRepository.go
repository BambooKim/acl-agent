package acl

import (
	"errors"

	"github.com/bambookim/acl-agent/acl-api-server/global/database"
	"gorm.io/gorm"
)

type AclRepository interface {
	FindById(tx *gorm.DB, id int) (bool, *AclEntity, error)
	Save(tx *gorm.DB, aclEntity *AclEntity) error
}

type AclRepositoryImpl struct{}

func NewAclRepository() AclRepository {
	return &AclRepositoryImpl{}
}

func (ri *AclRepositoryImpl) FindById(tx *gorm.DB, id int) (bool, *AclEntity, error) {
	findAcl := &AclEntity{}

	if err := tx.First(findAcl, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil, nil
		}
		return false, nil, err
	}

	return true, findAcl, nil
}

func (ri *AclRepositoryImpl) Save(tx *gorm.DB, aclEntity *AclEntity) error {
	if err := database.DB.Save(aclEntity).Error; err != nil {
		return err
	}

	return nil
}
