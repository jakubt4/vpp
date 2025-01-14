// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: /usr/share/vpp/api/dhcp.api.json

/*
Package dhcp is a generated VPP binary API for 'dhcp' module.

It consists of:
	  5 types
	 25 messages
	 11 services
*/
package dhcp

import (
	bytes "bytes"
	context "context"
	api "git.fd.io/govpp.git/api"
	struc "github.com/lunixbochs/struc"
	io "io"
	strconv "strconv"
)

const (
	// ModuleName is the name of this module.
	ModuleName = "dhcp"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x8c76d62a
)

// DHCP6AddressInfo represents VPP binary API type 'dhcp6_address_info'.
type DHCP6AddressInfo struct {
	Address       []byte `struc:"[16]byte"`
	ValidTime     uint32
	PreferredTime uint32
}

func (*DHCP6AddressInfo) GetTypeName() string {
	return "dhcp6_address_info"
}
func (*DHCP6AddressInfo) GetCrcString() string {
	return "f3d501e2"
}

// DHCP6PdPrefixInfo represents VPP binary API type 'dhcp6_pd_prefix_info'.
type DHCP6PdPrefixInfo struct {
	Prefix        []byte `struc:"[16]byte"`
	PrefixLength  uint8
	ValidTime     uint32
	PreferredTime uint32
}

func (*DHCP6PdPrefixInfo) GetTypeName() string {
	return "dhcp6_pd_prefix_info"
}
func (*DHCP6PdPrefixInfo) GetCrcString() string {
	return "c459690e"
}

// DHCPClient represents VPP binary API type 'dhcp_client'.
type DHCPClient struct {
	SwIfIndex        uint32
	Hostname         []byte `struc:"[64]byte"`
	ID               []byte `struc:"[64]byte"`
	WantDHCPEvent    uint8
	SetBroadcastFlag uint8
	PID              uint32
}

func (*DHCPClient) GetTypeName() string {
	return "dhcp_client"
}
func (*DHCPClient) GetCrcString() string {
	return "527f7935"
}

// DHCPLease represents VPP binary API type 'dhcp_lease'.
type DHCPLease struct {
	SwIfIndex     uint32
	State         uint8
	Hostname      []byte `struc:"[64]byte"`
	IsIPv6        uint8
	MaskWidth     uint8
	HostAddress   []byte `struc:"[16]byte"`
	RouterAddress []byte `struc:"[16]byte"`
	HostMac       []byte `struc:"[6]byte"`
}

func (*DHCPLease) GetTypeName() string {
	return "dhcp_lease"
}
func (*DHCPLease) GetCrcString() string {
	return "61090276"
}

// DHCPServer represents VPP binary API type 'dhcp_server'.
type DHCPServer struct {
	ServerVrfID uint32
	DHCPServer  []byte `struc:"[16]byte"`
}

func (*DHCPServer) GetTypeName() string {
	return "dhcp_server"
}
func (*DHCPServer) GetCrcString() string {
	return "f16506c4"
}

// DHCP6ClientsEnableDisable represents VPP binary API message 'dhcp6_clients_enable_disable'.
type DHCP6ClientsEnableDisable struct {
	Enable uint8
}

func (*DHCP6ClientsEnableDisable) GetMessageName() string {
	return "dhcp6_clients_enable_disable"
}
func (*DHCP6ClientsEnableDisable) GetCrcString() string {
	return "8050327d"
}
func (*DHCP6ClientsEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DHCP6ClientsEnableDisableReply represents VPP binary API message 'dhcp6_clients_enable_disable_reply'.
type DHCP6ClientsEnableDisableReply struct {
	Retval int32
}

func (*DHCP6ClientsEnableDisableReply) GetMessageName() string {
	return "dhcp6_clients_enable_disable_reply"
}
func (*DHCP6ClientsEnableDisableReply) GetCrcString() string {
	return "e8d4e804"
}
func (*DHCP6ClientsEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// DHCP6DuidLlSet represents VPP binary API message 'dhcp6_duid_ll_set'.
type DHCP6DuidLlSet struct {
	DuidLl []byte `struc:"[10]byte"`
}

func (*DHCP6DuidLlSet) GetMessageName() string {
	return "dhcp6_duid_ll_set"
}
func (*DHCP6DuidLlSet) GetCrcString() string {
	return "0f6ca323"
}
func (*DHCP6DuidLlSet) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DHCP6DuidLlSetReply represents VPP binary API message 'dhcp6_duid_ll_set_reply'.
type DHCP6DuidLlSetReply struct {
	Retval int32
}

func (*DHCP6DuidLlSetReply) GetMessageName() string {
	return "dhcp6_duid_ll_set_reply"
}
func (*DHCP6DuidLlSetReply) GetCrcString() string {
	return "e8d4e804"
}
func (*DHCP6DuidLlSetReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// DHCP6PdReplyEvent represents VPP binary API message 'dhcp6_pd_reply_event'.
type DHCP6PdReplyEvent struct {
	PID             uint32
	SwIfIndex       uint32
	ServerIndex     uint32
	MsgType         uint8
	T1              uint32
	T2              uint32
	InnerStatusCode uint16
	StatusCode      uint16
	Preference      uint8
	NPrefixes       uint32 `struc:"sizeof=Prefixes"`
	Prefixes        []DHCP6PdPrefixInfo
}

func (*DHCP6PdReplyEvent) GetMessageName() string {
	return "dhcp6_pd_reply_event"
}
func (*DHCP6PdReplyEvent) GetCrcString() string {
	return "48e73c36"
}
func (*DHCP6PdReplyEvent) GetMessageType() api.MessageType {
	return api.EventMessage
}

// DHCP6PdSendClientMessage represents VPP binary API message 'dhcp6_pd_send_client_message'.
type DHCP6PdSendClientMessage struct {
	SwIfIndex   uint32
	ServerIndex uint32
	Irt         uint32
	Mrt         uint32
	Mrc         uint32
	Mrd         uint32
	Stop        uint8
	MsgType     uint8
	T1          uint32
	T2          uint32
	NPrefixes   uint32 `struc:"sizeof=Prefixes"`
	Prefixes    []DHCP6PdPrefixInfo
}

func (*DHCP6PdSendClientMessage) GetMessageName() string {
	return "dhcp6_pd_send_client_message"
}
func (*DHCP6PdSendClientMessage) GetCrcString() string {
	return "5a063fbc"
}
func (*DHCP6PdSendClientMessage) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DHCP6PdSendClientMessageReply represents VPP binary API message 'dhcp6_pd_send_client_message_reply'.
type DHCP6PdSendClientMessageReply struct {
	Retval int32
}

func (*DHCP6PdSendClientMessageReply) GetMessageName() string {
	return "dhcp6_pd_send_client_message_reply"
}
func (*DHCP6PdSendClientMessageReply) GetCrcString() string {
	return "e8d4e804"
}
func (*DHCP6PdSendClientMessageReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// DHCP6ReplyEvent represents VPP binary API message 'dhcp6_reply_event'.
type DHCP6ReplyEvent struct {
	PID             uint32
	SwIfIndex       uint32
	ServerIndex     uint32
	MsgType         uint8
	T1              uint32
	T2              uint32
	InnerStatusCode uint16
	StatusCode      uint16
	Preference      uint8
	NAddresses      uint32 `struc:"sizeof=Addresses"`
	Addresses       []DHCP6AddressInfo
}

func (*DHCP6ReplyEvent) GetMessageName() string {
	return "dhcp6_reply_event"
}
func (*DHCP6ReplyEvent) GetCrcString() string {
	return "ac4563f9"
}
func (*DHCP6ReplyEvent) GetMessageType() api.MessageType {
	return api.EventMessage
}

// DHCP6SendClientMessage represents VPP binary API message 'dhcp6_send_client_message'.
type DHCP6SendClientMessage struct {
	SwIfIndex   uint32
	ServerIndex uint32
	Irt         uint32
	Mrt         uint32
	Mrc         uint32
	Mrd         uint32
	Stop        uint8
	MsgType     uint8
	T1          uint32
	T2          uint32
	NAddresses  uint32 `struc:"sizeof=Addresses"`
	Addresses   []DHCP6AddressInfo
}

func (*DHCP6SendClientMessage) GetMessageName() string {
	return "dhcp6_send_client_message"
}
func (*DHCP6SendClientMessage) GetCrcString() string {
	return "a13ae8c4"
}
func (*DHCP6SendClientMessage) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DHCP6SendClientMessageReply represents VPP binary API message 'dhcp6_send_client_message_reply'.
type DHCP6SendClientMessageReply struct {
	Retval int32
}

func (*DHCP6SendClientMessageReply) GetMessageName() string {
	return "dhcp6_send_client_message_reply"
}
func (*DHCP6SendClientMessageReply) GetCrcString() string {
	return "e8d4e804"
}
func (*DHCP6SendClientMessageReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// DHCPClientConfig represents VPP binary API message 'dhcp_client_config'.
type DHCPClientConfig struct {
	IsAdd  uint8
	Client DHCPClient
}

func (*DHCPClientConfig) GetMessageName() string {
	return "dhcp_client_config"
}
func (*DHCPClientConfig) GetCrcString() string {
	return "c32ccdfe"
}
func (*DHCPClientConfig) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DHCPClientConfigReply represents VPP binary API message 'dhcp_client_config_reply'.
type DHCPClientConfigReply struct {
	Retval int32
}

func (*DHCPClientConfigReply) GetMessageName() string {
	return "dhcp_client_config_reply"
}
func (*DHCPClientConfigReply) GetCrcString() string {
	return "e8d4e804"
}
func (*DHCPClientConfigReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// DHCPClientDetails represents VPP binary API message 'dhcp_client_details'.
type DHCPClientDetails struct {
	Client DHCPClient
	Lease  DHCPLease
}

func (*DHCPClientDetails) GetMessageName() string {
	return "dhcp_client_details"
}
func (*DHCPClientDetails) GetCrcString() string {
	return "7ea3a745"
}
func (*DHCPClientDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// DHCPClientDump represents VPP binary API message 'dhcp_client_dump'.
type DHCPClientDump struct{}

func (*DHCPClientDump) GetMessageName() string {
	return "dhcp_client_dump"
}
func (*DHCPClientDump) GetCrcString() string {
	return "51077d14"
}
func (*DHCPClientDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DHCPComplEvent represents VPP binary API message 'dhcp_compl_event'.
type DHCPComplEvent struct {
	PID   uint32
	Lease DHCPLease
}

func (*DHCPComplEvent) GetMessageName() string {
	return "dhcp_compl_event"
}
func (*DHCPComplEvent) GetCrcString() string {
	return "2105c31b"
}
func (*DHCPComplEvent) GetMessageType() api.MessageType {
	return api.EventMessage
}

// DHCPProxyConfig represents VPP binary API message 'dhcp_proxy_config'.
type DHCPProxyConfig struct {
	RxVrfID        uint32
	ServerVrfID    uint32
	IsIPv6         uint8
	IsAdd          uint8
	DHCPServer     []byte `struc:"[16]byte"`
	DHCPSrcAddress []byte `struc:"[16]byte"`
}

func (*DHCPProxyConfig) GetMessageName() string {
	return "dhcp_proxy_config"
}
func (*DHCPProxyConfig) GetCrcString() string {
	return "6af4b645"
}
func (*DHCPProxyConfig) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DHCPProxyConfigReply represents VPP binary API message 'dhcp_proxy_config_reply'.
type DHCPProxyConfigReply struct {
	Retval int32
}

func (*DHCPProxyConfigReply) GetMessageName() string {
	return "dhcp_proxy_config_reply"
}
func (*DHCPProxyConfigReply) GetCrcString() string {
	return "e8d4e804"
}
func (*DHCPProxyConfigReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// DHCPProxyDetails represents VPP binary API message 'dhcp_proxy_details'.
type DHCPProxyDetails struct {
	RxVrfID        uint32
	VssOui         uint32
	VssFibID       uint32
	VssType        uint8
	VssVPNAsciiID  []byte `struc:"[129]byte"`
	IsIPv6         uint8
	DHCPSrcAddress []byte `struc:"[16]byte"`
	Count          uint8  `struc:"sizeof=Servers"`
	Servers        []DHCPServer
}

func (*DHCPProxyDetails) GetMessageName() string {
	return "dhcp_proxy_details"
}
func (*DHCPProxyDetails) GetCrcString() string {
	return "a5f2ad84"
}
func (*DHCPProxyDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// DHCPProxyDump represents VPP binary API message 'dhcp_proxy_dump'.
type DHCPProxyDump struct {
	IsIP6 uint8
}

func (*DHCPProxyDump) GetMessageName() string {
	return "dhcp_proxy_dump"
}
func (*DHCPProxyDump) GetCrcString() string {
	return "6fe91190"
}
func (*DHCPProxyDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DHCPProxySetVss represents VPP binary API message 'dhcp_proxy_set_vss'.
type DHCPProxySetVss struct {
	TblID      uint32
	VssType    uint8
	VPNAsciiID []byte `struc:"[129]byte"`
	Oui        uint32
	VPNIndex   uint32
	IsIPv6     uint8
	IsAdd      uint8
}

func (*DHCPProxySetVss) GetMessageName() string {
	return "dhcp_proxy_set_vss"
}
func (*DHCPProxySetVss) GetCrcString() string {
	return "606535aa"
}
func (*DHCPProxySetVss) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// DHCPProxySetVssReply represents VPP binary API message 'dhcp_proxy_set_vss_reply'.
type DHCPProxySetVssReply struct {
	Retval int32
}

func (*DHCPProxySetVssReply) GetMessageName() string {
	return "dhcp_proxy_set_vss_reply"
}
func (*DHCPProxySetVssReply) GetCrcString() string {
	return "e8d4e804"
}
func (*DHCPProxySetVssReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// WantDHCP6PdReplyEvents represents VPP binary API message 'want_dhcp6_pd_reply_events'.
type WantDHCP6PdReplyEvents struct {
	EnableDisable uint8
	PID           uint32
}

func (*WantDHCP6PdReplyEvents) GetMessageName() string {
	return "want_dhcp6_pd_reply_events"
}
func (*WantDHCP6PdReplyEvents) GetCrcString() string {
	return "05b454b5"
}
func (*WantDHCP6PdReplyEvents) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// WantDHCP6PdReplyEventsReply represents VPP binary API message 'want_dhcp6_pd_reply_events_reply'.
type WantDHCP6PdReplyEventsReply struct {
	Retval int32
}

func (*WantDHCP6PdReplyEventsReply) GetMessageName() string {
	return "want_dhcp6_pd_reply_events_reply"
}
func (*WantDHCP6PdReplyEventsReply) GetCrcString() string {
	return "e8d4e804"
}
func (*WantDHCP6PdReplyEventsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// WantDHCP6ReplyEvents represents VPP binary API message 'want_dhcp6_reply_events'.
type WantDHCP6ReplyEvents struct {
	EnableDisable uint8
	PID           uint32
}

func (*WantDHCP6ReplyEvents) GetMessageName() string {
	return "want_dhcp6_reply_events"
}
func (*WantDHCP6ReplyEvents) GetCrcString() string {
	return "05b454b5"
}
func (*WantDHCP6ReplyEvents) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// WantDHCP6ReplyEventsReply represents VPP binary API message 'want_dhcp6_reply_events_reply'.
type WantDHCP6ReplyEventsReply struct {
	Retval int32
}

func (*WantDHCP6ReplyEventsReply) GetMessageName() string {
	return "want_dhcp6_reply_events_reply"
}
func (*WantDHCP6ReplyEventsReply) GetCrcString() string {
	return "e8d4e804"
}
func (*WantDHCP6ReplyEventsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*DHCP6ClientsEnableDisable)(nil), "dhcp.DHCP6ClientsEnableDisable")
	api.RegisterMessage((*DHCP6ClientsEnableDisableReply)(nil), "dhcp.DHCP6ClientsEnableDisableReply")
	api.RegisterMessage((*DHCP6DuidLlSet)(nil), "dhcp.DHCP6DuidLlSet")
	api.RegisterMessage((*DHCP6DuidLlSetReply)(nil), "dhcp.DHCP6DuidLlSetReply")
	api.RegisterMessage((*DHCP6PdReplyEvent)(nil), "dhcp.DHCP6PdReplyEvent")
	api.RegisterMessage((*DHCP6PdSendClientMessage)(nil), "dhcp.DHCP6PdSendClientMessage")
	api.RegisterMessage((*DHCP6PdSendClientMessageReply)(nil), "dhcp.DHCP6PdSendClientMessageReply")
	api.RegisterMessage((*DHCP6ReplyEvent)(nil), "dhcp.DHCP6ReplyEvent")
	api.RegisterMessage((*DHCP6SendClientMessage)(nil), "dhcp.DHCP6SendClientMessage")
	api.RegisterMessage((*DHCP6SendClientMessageReply)(nil), "dhcp.DHCP6SendClientMessageReply")
	api.RegisterMessage((*DHCPClientConfig)(nil), "dhcp.DHCPClientConfig")
	api.RegisterMessage((*DHCPClientConfigReply)(nil), "dhcp.DHCPClientConfigReply")
	api.RegisterMessage((*DHCPClientDetails)(nil), "dhcp.DHCPClientDetails")
	api.RegisterMessage((*DHCPClientDump)(nil), "dhcp.DHCPClientDump")
	api.RegisterMessage((*DHCPComplEvent)(nil), "dhcp.DHCPComplEvent")
	api.RegisterMessage((*DHCPProxyConfig)(nil), "dhcp.DHCPProxyConfig")
	api.RegisterMessage((*DHCPProxyConfigReply)(nil), "dhcp.DHCPProxyConfigReply")
	api.RegisterMessage((*DHCPProxyDetails)(nil), "dhcp.DHCPProxyDetails")
	api.RegisterMessage((*DHCPProxyDump)(nil), "dhcp.DHCPProxyDump")
	api.RegisterMessage((*DHCPProxySetVss)(nil), "dhcp.DHCPProxySetVss")
	api.RegisterMessage((*DHCPProxySetVssReply)(nil), "dhcp.DHCPProxySetVssReply")
	api.RegisterMessage((*WantDHCP6PdReplyEvents)(nil), "dhcp.WantDHCP6PdReplyEvents")
	api.RegisterMessage((*WantDHCP6PdReplyEventsReply)(nil), "dhcp.WantDHCP6PdReplyEventsReply")
	api.RegisterMessage((*WantDHCP6ReplyEvents)(nil), "dhcp.WantDHCP6ReplyEvents")
	api.RegisterMessage((*WantDHCP6ReplyEventsReply)(nil), "dhcp.WantDHCP6ReplyEventsReply")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*DHCP6ClientsEnableDisable)(nil),
		(*DHCP6ClientsEnableDisableReply)(nil),
		(*DHCP6DuidLlSet)(nil),
		(*DHCP6DuidLlSetReply)(nil),
		(*DHCP6PdReplyEvent)(nil),
		(*DHCP6PdSendClientMessage)(nil),
		(*DHCP6PdSendClientMessageReply)(nil),
		(*DHCP6ReplyEvent)(nil),
		(*DHCP6SendClientMessage)(nil),
		(*DHCP6SendClientMessageReply)(nil),
		(*DHCPClientConfig)(nil),
		(*DHCPClientConfigReply)(nil),
		(*DHCPClientDetails)(nil),
		(*DHCPClientDump)(nil),
		(*DHCPComplEvent)(nil),
		(*DHCPProxyConfig)(nil),
		(*DHCPProxyConfigReply)(nil),
		(*DHCPProxyDetails)(nil),
		(*DHCPProxyDump)(nil),
		(*DHCPProxySetVss)(nil),
		(*DHCPProxySetVssReply)(nil),
		(*WantDHCP6PdReplyEvents)(nil),
		(*WantDHCP6PdReplyEventsReply)(nil),
		(*WantDHCP6ReplyEvents)(nil),
		(*WantDHCP6ReplyEventsReply)(nil),
	}
}

// RPCService represents RPC service API for dhcp module.
type RPCService interface {
	DumpDHCPClient(ctx context.Context, in *DHCPClientDump) (RPCService_DumpDHCPClientClient, error)
	DumpDHCPProxy(ctx context.Context, in *DHCPProxyDump) (RPCService_DumpDHCPProxyClient, error)
	DHCP6ClientsEnableDisable(ctx context.Context, in *DHCP6ClientsEnableDisable) (*DHCP6ClientsEnableDisableReply, error)
	DHCP6DuidLlSet(ctx context.Context, in *DHCP6DuidLlSet) (*DHCP6DuidLlSetReply, error)
	DHCP6PdSendClientMessage(ctx context.Context, in *DHCP6PdSendClientMessage) (*DHCP6PdSendClientMessageReply, error)
	DHCP6SendClientMessage(ctx context.Context, in *DHCP6SendClientMessage) (*DHCP6SendClientMessageReply, error)
	DHCPClientConfig(ctx context.Context, in *DHCPClientConfig) (*DHCPClientConfigReply, error)
	DHCPProxyConfig(ctx context.Context, in *DHCPProxyConfig) (*DHCPProxyConfigReply, error)
	DHCPProxySetVss(ctx context.Context, in *DHCPProxySetVss) (*DHCPProxySetVssReply, error)
	WantDHCP6PdReplyEvents(ctx context.Context, in *WantDHCP6PdReplyEvents) (*WantDHCP6PdReplyEventsReply, error)
	WantDHCP6ReplyEvents(ctx context.Context, in *WantDHCP6ReplyEvents) (*WantDHCP6ReplyEventsReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpDHCPClient(ctx context.Context, in *DHCPClientDump) (RPCService_DumpDHCPClientClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpDHCPClientClient{stream}
	return x, nil
}

type RPCService_DumpDHCPClientClient interface {
	Recv() (*DHCPClientDetails, error)
}

type serviceClient_DumpDHCPClientClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpDHCPClientClient) Recv() (*DHCPClientDetails, error) {
	m := new(DHCPClientDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpDHCPProxy(ctx context.Context, in *DHCPProxyDump) (RPCService_DumpDHCPProxyClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpDHCPProxyClient{stream}
	return x, nil
}

type RPCService_DumpDHCPProxyClient interface {
	Recv() (*DHCPProxyDetails, error)
}

type serviceClient_DumpDHCPProxyClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpDHCPProxyClient) Recv() (*DHCPProxyDetails, error) {
	m := new(DHCPProxyDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DHCP6ClientsEnableDisable(ctx context.Context, in *DHCP6ClientsEnableDisable) (*DHCP6ClientsEnableDisableReply, error) {
	out := new(DHCP6ClientsEnableDisableReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DHCP6DuidLlSet(ctx context.Context, in *DHCP6DuidLlSet) (*DHCP6DuidLlSetReply, error) {
	out := new(DHCP6DuidLlSetReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DHCP6PdSendClientMessage(ctx context.Context, in *DHCP6PdSendClientMessage) (*DHCP6PdSendClientMessageReply, error) {
	out := new(DHCP6PdSendClientMessageReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DHCP6SendClientMessage(ctx context.Context, in *DHCP6SendClientMessage) (*DHCP6SendClientMessageReply, error) {
	out := new(DHCP6SendClientMessageReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DHCPClientConfig(ctx context.Context, in *DHCPClientConfig) (*DHCPClientConfigReply, error) {
	out := new(DHCPClientConfigReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DHCPProxyConfig(ctx context.Context, in *DHCPProxyConfig) (*DHCPProxyConfigReply, error) {
	out := new(DHCPProxyConfigReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DHCPProxySetVss(ctx context.Context, in *DHCPProxySetVss) (*DHCPProxySetVssReply, error) {
	out := new(DHCPProxySetVssReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) WantDHCP6PdReplyEvents(ctx context.Context, in *WantDHCP6PdReplyEvents) (*WantDHCP6PdReplyEventsReply, error) {
	out := new(WantDHCP6PdReplyEventsReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) WantDHCP6ReplyEvents(ctx context.Context, in *WantDHCP6ReplyEvents) (*WantDHCP6ReplyEventsReply, error) {
	out := new(WantDHCP6ReplyEventsReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion1 // please upgrade the GoVPP api package

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = bytes.NewBuffer
var _ = context.Background
var _ = io.Copy
var _ = strconv.Itoa
var _ = struc.Pack
