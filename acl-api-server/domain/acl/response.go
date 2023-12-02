package acl

import "time"

type GetAclResponse struct {
	Id              int          `json:"id"`
	Name            string       `json:"name"`
	Action          AclAction    `json:"action"`
	Direction       AclDirection `json:"direction"`
	SourceCidr      string       `json:"sourceCidr"`
	SourcePortStart int          `json:"sourcePortStart"`
	SourcePortStop  int          `json:"sourcePortStop"`
	DestCidr        string       `json:"destCidr"`
	DestPortStart   int          `json:"destPortStart"`
	DestPortStop    int          `json:"destPortStop"`
	Protocol        AclProtocol  `json:"protocol"`
	CreatedAt       time.Time    `json:"createdAt"`
	ModifiedAt      time.Time    `json:"modifiedAt"`
}
