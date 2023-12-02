package app

import (
	"go.ligato.io/cn-infra/v2/datasync/kvdbsync"
	"go.ligato.io/cn-infra/v2/db/keyval/etcd"
	"go.ligato.io/cn-infra/v2/health/statuscheck"
	"go.ligato.io/cn-infra/v2/infra"
	"go.ligato.io/cn-infra/v2/rpc/rest"
	"go.ligato.io/cn-infra/v2/servicelabel"
	"go.ligato.io/vpp-agent/v3/plugins/orchestrator"
	"go.ligato.io/vpp-agent/v3/plugins/orchestrator/watcher"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/aclplugin"
)

type AclAgent struct {
	Deps
}

type Deps struct {
	infra.PluginDeps

	AclPlugin *aclplugin.ACLPlugin

	REST         rest.HTTPHandlers
	Orchestrator *orchestrator.Plugin
	ServiceLabel servicelabel.ReaderAPI
	StatusCheck  *statuscheck.Plugin
}

func NewAclAgent() *AclAgent {
	a := &AclAgent{}

	a.AclPlugin = &aclplugin.DefaultPlugin
	a.REST = &rest.DefaultPlugin
	a.Orchestrator = &orchestrator.DefaultPlugin
	a.ServiceLabel = &servicelabel.DefaultPlugin
	a.StatusCheck = &statuscheck.DefaultPlugin

	a.SetName("acl-agent")
	a.SetupLog()

	etcdDataSync := kvdbsync.NewPlugin(kvdbsync.UseKV(&etcd.DefaultPlugin))
	statuscheck.DefaultPlugin.Transport = etcdDataSync
	a.Orchestrator.Watcher = watcher.NewPlugin(watcher.UseWatchers(
		etcdDataSync,
	))
	a.Orchestrator.StatusPublisher = etcdDataSync

	a.PluginDeps.Setup()

	return a
}

func (a *AclAgent) Init() error {
	a.StatusCheck.Register(a.PluginName, nil)
	a.StatusCheck.ReportStateChange(a.PluginName, statuscheck.Init, nil)

	return nil
}

func (a *AclAgent) AfterInit() error {
	if err := orchestrator.DefaultPlugin.InitialSync(); err != nil {
		return err
	}
	a.StatusCheck.ReportStateChange(a.PluginName, statuscheck.OK, nil)

	agentPrefix := a.ServiceLabel.GetAgentPrefix()
	a.Log.Infof("agentPrefix: %s", agentPrefix)

	return nil
}
