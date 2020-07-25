package driver

import (
	"etcdadmind/config"
	"etcdadmind/log"
	pb "etcdadmind/pb/etcdadminpb"
	"etcdadmind/server/driver/client"
	"etcdadmind/server/driver/command"
	"etcdadmind/server/driver/etcdcfg"
	"etcdadmind/server/driver/etcdclient"
	"etcdadmind/utils"
	"fmt"
	"go.uber.org/zap"
)

type DriverInterface interface {
	AddMember(m *pb.AddMemberRequest_Member) error
	ManagerEtcd(cmd pb.EtcdCmd, clearwal bool, cfgs []*pb.ManagerEtcdRequest_Config) error
	ListMember() ([]*pb.ListMemberReply_Member, error)
}

type DriverImpl struct {
	logger   *zap.Logger
	portGrpc string
}

func resetEtcdConfig() error {
	etcdCfgFile := config.Init().Get("ETCD_CONF_FILE")

	m, err := etcdcfg.EtcdConfigMapInit()

	if err != nil {
		return err
	}

	return etcdcfg.EtcdConfigWrite(etcdCfgFile, m)
}

func resetEtcd(isStart bool) error {
	var err error

	// Stop etcd, ignore error
	command.CmdEtcdctlStop()

	if err := resetEtcdConfig(); err != nil {
		goto exit
	}

	if err = etcdcfg.EtcdWalDelete(); err != nil {
		goto exit
	}

exit:
	if isStart == true {
		er := command.EtcdctlStart()
		if err == nil {
			err = er
		}
	}
	return err
}

func New() *DriverImpl {
	cfg := config.Init()

	drv := &DriverImpl{
		portGrpc: cfg.Get("GRPC_PORT"),
		logger:   log.GetLogger(),
	}
	return drv
}

func (drv *DriverImpl) AddMember(m *pb.AddMemberRequest_Member) error {
	cfgStore := config.Init()
	clientPort := cfgStore.Get("ETCD_CLIENT_PORT")
	peerPort := cfgStore.Get("ETCD_PEER_PORT")

	etcdcli := etcdclient.New("127.0.0.1", clientPort)
	defer etcdclient.Release(etcdcli)

	c := client.New(m.Ip, drv.portGrpc)
	defer client.Release(c)

	drv.logger.Info(fmt.Sprintf("add member: %v %v", m.Name, m.Ip))
	ips, err := utils.GetHostIP4()
	if err != nil {
		return err
	}

	// remote clean etcd.cfg, wal and stop etcd
	cfgs, _ := etcdcfg.EtcdConfigMapInit()
	c.GrpcClientManagerEtcd(cfgs, true, client.EtcdCmdStop)

	if utils.ContainsString(ips, m.Ip) >= 0 {
		cfgs["ETCD_NAME"] = m.Name
		cfgs["ETCD_INITIAL_CLUSTER"] = fmt.Sprintf("%s=http://%s:%s", m.Name, m.Ip, peerPort)
		c.GrpcClientManagerEtcd(cfgs, false, client.EtcdCmdStart)
	} else {
		cfgs["ETCD_ADVERTISE_CLIENT_URLS"] = fmt.Sprintf("http://%s:%s", m.Ip, clientPort)
		cfgs["ETCD_INITIAL_ADVERTISE_PEER_URLS"] = fmt.Sprintf("http://%s:%s", m.Ip, peerPort)
		cfgs["ETCD_NAME"] = m.Name
		cfgs["ETCD_INITIAL_CLUSTER_STATE"] = "existing"

		members := etcdcli.MemberList()
		initCluster := ""
		for i := range members {
			initCluster = fmt.Sprintf("%s%s=http://%s:%s,", initCluster, i, members[i], peerPort)
		}
		initCluster = fmt.Sprintf("%s%s=http://%s:%s", initCluster, m.Name, m.Ip, peerPort)
		cfgs["ETCD_INITIAL_CLUSTER"] = initCluster

		// add member to cluster
		etcdcli.MemberAdd(cfgs["ETCD_INITIAL_ADVERTISE_PEER_URLS"])

		// remote writing etcd.cfg and start etcd
		c.GrpcClientManagerEtcd(cfgs, false, client.EtcdCmdStart)
	}
	return nil
}

func (drv *DriverImpl) ManagerEtcd(cmd pb.EtcdCmd, clearwal bool,
	cfgs []*pb.ManagerEtcdRequest_Config) error {
	cfgStore := config.Init()

	if cmd != pb.EtcdCmd_NONE {
		command.CmdEtcdctlStop()
	}

	if len(cfgs) > 0 {
		etcdCfgFile := cfgStore.Get("ETCD_CONF_FILE")
		m := map[string]string{}
		for _, c := range cfgs {
			m[c.Key] = c.Value
		}

		etcdcfg.EtcdConfigWrite(etcdCfgFile, m)
	}

	if clearwal == true {
		etcdcfg.EtcdWalDelete()
	}

	if cmd == pb.EtcdCmd_START || cmd == pb.EtcdCmd_RESTART {
		command.CmdEtcdctlStart()
	}
	return nil
}

func (drv *DriverImpl) ListMember() ([]*pb.ListMemberReply_Member, error) {
	memberList := []*pb.ListMemberReply_Member{}

	cfgStore := config.Init()
	clientPort := cfgStore.Get("ETCD_CLIENT_PORT")

	etcdcli := etcdclient.New("127.0.0.1", clientPort)
	defer etcdclient.Release(etcdcli)

	members := etcdcli.MemberList()
	for m := range members {
		memberInfo := &pb.ListMemberReply_Member{
			Name: m,
			Ip:   members[m],
		}
		memberList = append(memberList, memberInfo)
	}

	return memberList, nil
}
