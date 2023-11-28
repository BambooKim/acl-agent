package acl

import "github.com/bambookim/acl-agent/acl-api-server/global/common"

type AclEntity struct {
	Id                    int `gorm:"primarykey;autoIncrement"`
	AclIndex              int
	Action                string
	SourceCidr            string
	SourcePortStart       int
	SourcePortStop        int
	DestCidr              string
	DestPortStart         int
	DestPortStop          int
	Protocol              string
	common.BaseTimeEntity `gorm:"embedded"`
}
