package acl

type GetAclResponse struct {
	AclIndex int        `json:"aclIndex"`
	AclElems []*AclElem `json:"aclElems"`
}

type AclElem struct {
	Action          AclAction `json:"action"`
	SourceCidr      string    `json:"sourceCidr"`
	SourcePortStart int       `json:"sourcePortStart"`
	SourcePortStop  int       `json:"sourcePortStop"`
	DestCidr        string    `json:"destCidr"`
	DestPortStart   int       `json:"destPortStart"`
	DestPortStop    int       `json:"destPortStop"`
	Protocol        string    `json:"protocol"`
}
