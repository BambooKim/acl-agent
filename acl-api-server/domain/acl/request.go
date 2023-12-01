package acl

type CreateAclRequest struct {
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

type AclDirection string

const (
	ACL_INGRESS AclDirection = "ingress"
	ACL_EGRESS  AclDirection = "egress"
)

type AclProtocol string

const (
	ACL_ICMP AclProtocol = "icmp" // 1
	ACL_TCP  AclProtocol = "tcp"  // 6
	ACL_UDP  AclProtocol = "udp"  // 17
)
