package acl

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	clientv3 "go.etcd.io/etcd/client/v3"
	vpp_acl "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/acl"
)

type AclService interface {
	CreateAcl(req *CreateAclRequest) error
}

type AclServiceImpl struct {
	AclRepository
	client *clientv3.Client
}

func NewAclService(client *clientv3.Client, aclRepository *AclRepository) AclService {
	return &AclServiceImpl{
		AclRepository: aclRepository,
		client:        client,
	}
}

func (si *AclServiceImpl) CreateAcl(req *CreateAclRequest) error {
	// todo: acl repository save with transaction

	ipRule := &vpp_acl.ACL_Rule_IpRule{}
	switch req.Protocol {
	case ACL_ICMP:
		ipRule = getIcmpRule(req)
	case ACL_TCP:
		ipRule = getTcpRule(req)
	case ACL_UDP:
		ipRule = getUdpRule(req)
	default:
		return errors.New("INVALIDATE_PROTOCOL")
	}

	rule := &vpp_acl.ACL_Rule{
		Action: getAclAction(req.Action),
		IpRule: ipRule,
	}
	rules := []*vpp_acl.ACL_Rule{rule}
	newAcl := &vpp_acl.ACL{
		Name:       req.Name,
		Rules:      rules,
		Interfaces: getAclInterfaces(req.Direction),
	}

	// if err := si.broker.Put(vpp_acl.Key(req.Name), newAcl); err != nil {
	// 	return err
	// }

	newAclJson, err := json.Marshal(newAcl)
	if err != nil {
		return err
	}
	log.Printf("acl: %+v", newAcl)
	log.Printf("json: %s", string(newAclJson))
	log.Printf("key: %s", vpp_acl.Key(req.Name))
	if res, err := si.client.Put(context.Background(), "/vnf-agent/vpp1/"+vpp_acl.Key(req.Name), string(newAclJson)); err != nil {
		return err
	} else {
		log.Printf("result: %+v", res)
	}

	return nil
}

func getAclInterfaces(direction AclDirection) *vpp_acl.ACL_Interfaces {
	interfaces := &vpp_acl.ACL_Interfaces{}

	switch direction {
	case ACL_INGRESS:
		interfaces.Ingress = []string{"GigabitEthernet0/3/0"}
	case ACL_EGRESS:
		interfaces.Egress = []string{"GigabitEthernet0/3/0"}
	}

	return interfaces
}

func getAclAction(action AclAction) vpp_acl.ACL_Rule_Action {
	ret := vpp_acl.ACL_Rule_DENY

	switch action {
	case ACL_ACTION_PERMIT:
		ret = vpp_acl.ACL_Rule_PERMIT
	case ACL_ACTION_DENY:
		ret = vpp_acl.ACL_Rule_DENY
	case ACL_ACTION_PERMIT_REFLECT:
		ret = vpp_acl.ACL_Rule_REFLECT
	}

	return ret
}

func getIcmpRule(dto *CreateAclRequest) *vpp_acl.ACL_Rule_IpRule {
	return &vpp_acl.ACL_Rule_IpRule{
		Ip: &vpp_acl.ACL_Rule_IpRule_Ip{
			DestinationNetwork: dto.DestCidr,
			SourceNetwork:      dto.SourceCidr,
			Protocol:           1,
		},
		Icmp: &vpp_acl.ACL_Rule_IpRule_Icmp{
			Icmpv6: false,
		},
	}
}

func getTcpRule(dto *CreateAclRequest) *vpp_acl.ACL_Rule_IpRule {
	return &vpp_acl.ACL_Rule_IpRule{
		Ip: &vpp_acl.ACL_Rule_IpRule_Ip{
			DestinationNetwork: dto.DestCidr,
			SourceNetwork:      dto.SourceCidr,
			Protocol:           6,
		},
		Tcp: &vpp_acl.ACL_Rule_IpRule_Tcp{
			DestinationPortRange: &vpp_acl.ACL_Rule_IpRule_PortRange{
				LowerPort: uint32(dto.DestPortStart),
				UpperPort: uint32(dto.DestPortStop),
			},
			SourcePortRange: &vpp_acl.ACL_Rule_IpRule_PortRange{
				LowerPort: uint32(dto.SourcePortStart),
				UpperPort: uint32(dto.SourcePortStop),
			},
		},
	}
}

func getUdpRule(dto *CreateAclRequest) *vpp_acl.ACL_Rule_IpRule {
	return &vpp_acl.ACL_Rule_IpRule{
		Ip: &vpp_acl.ACL_Rule_IpRule_Ip{
			DestinationNetwork: dto.DestCidr,
			SourceNetwork:      dto.SourceCidr,
			Protocol:           17,
		},
		Udp: &vpp_acl.ACL_Rule_IpRule_Udp{
			DestinationPortRange: &vpp_acl.ACL_Rule_IpRule_PortRange{
				LowerPort: uint32(dto.DestPortStart),
				UpperPort: uint32(dto.DestPortStop),
			},
			SourcePortRange: &vpp_acl.ACL_Rule_IpRule_PortRange{
				LowerPort: uint32(dto.SourcePortStart),
				UpperPort: uint32(dto.SourcePortStop),
			},
		},
	}
}
