package acl

type CreateAclRequest struct {
	Action          AclAction `json:"action"`
	SourceCidr      string    `json:"sourceCidr"`
	SourcePortStart int       `json:"sourcePortStart"`
	SourcePortStop  int       `json:"sourcePortStop"`
	DestCidr        string    `json:"destCidr"`
	DestPortStart   int       `json:"destPortStart"`
	DestPortStop    int       `json:"destPortStop"`
	Protocol        string    `json:"protocol"`
}

type ModifyAclRequest struct {
	Action          AclAction `json:"action"`
	SourceCidr      string    `json:"sourceCidr"`
	SourcePortStart int       `json:"sourcePortStart"`
	SourcePortStop  int       `json:"sourcePortStop"`
	DestCidr        string    `json:"destCidr"`
	DestPortStart   int       `json:"destPortStart"`
	DestPortStop    int       `json:"destPortStop"`
	Protocol        string    `json:"protocol"`
}

type AclAction string

const (
	ACL_ACTION_PERMIT         AclAction = "permit"
	ACL_ACTION_DENY           AclAction = "deny"
	ACL_ACTION_PERMIT_REFLECT AclAction = "permit+reflect"
)
