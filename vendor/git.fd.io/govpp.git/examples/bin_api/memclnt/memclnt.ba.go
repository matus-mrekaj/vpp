// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: memclnt.api.json

/*
 Package memclnt is a generated from VPP binary API module 'memclnt'.

 It contains following objects:
	 13 services
	  2 types
	 22 messages
*/
package memclnt

import api "git.fd.io/govpp.git/api"
import struc "github.com/lunixbochs/struc"
import bytes "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = struc.Pack
var _ = bytes.NewBuffer

// Services represents VPP binary API services:
type Services interface {
	APIVersions(*APIVersions) (*APIVersionsReply, error)
	GetFirstMsgID(*GetFirstMsgID) (*GetFirstMsgIDReply, error)
	MemclntCreate(*MemclntCreate) (*MemclntCreateReply, error)
	MemclntDelete(*MemclntDelete) (*MemclntDeleteReply, error)
	MemclntKeepalive(*MemclntKeepalive) (*MemclntKeepaliveReply, error)
	MemclntReadTimeout(*MemclntReadTimeout) error
	MemclntRxThreadSuspend(*MemclntRxThreadSuspend) error
	RPCCall(*RPCCall) (*RPCCallReply, error)
	RxThreadExit(*RxThreadExit) error
	SockInitShm(*SockInitShm) (*SockInitShmReply, error)
	SockclntCreate(*SockclntCreate) (*SockclntCreateReply, error)
	SockclntDelete(*SockclntDelete) (*SockclntDeleteReply, error)
	TracePluginMsgIds(*TracePluginMsgIds) error
}

/* Types */

// MessageTableEntry represents VPP binary API type 'message_table_entry':
type MessageTableEntry struct {
	Index uint16
	Name  []byte `struc:"[64]byte"`
}

func (*MessageTableEntry) GetTypeName() string {
	return "message_table_entry"
}
func (*MessageTableEntry) GetCrcString() string {
	return "913bf1c6"
}

// ModuleVersion represents VPP binary API type 'module_version':
type ModuleVersion struct {
	Major uint32
	Minor uint32
	Patch uint32
	Name  []byte `struc:"[64]byte"`
}

func (*ModuleVersion) GetTypeName() string {
	return "module_version"
}
func (*ModuleVersion) GetCrcString() string {
	return "4b6da11a"
}

/* Messages */

// APIVersions represents VPP binary API message 'api_versions':
type APIVersions struct{}

func (*APIVersions) GetMessageName() string {
	return "api_versions"
}
func (*APIVersions) GetCrcString() string {
	return "51077d14"
}
func (*APIVersions) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// APIVersionsReply represents VPP binary API message 'api_versions_reply':
type APIVersionsReply struct {
	Retval      int32
	Count       uint32 `struc:"sizeof=APIVersions"`
	APIVersions []ModuleVersion
}

func (*APIVersionsReply) GetMessageName() string {
	return "api_versions_reply"
}
func (*APIVersionsReply) GetCrcString() string {
	return "90a39195"
}
func (*APIVersionsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// GetFirstMsgID represents VPP binary API message 'get_first_msg_id':
type GetFirstMsgID struct {
	Name []byte `struc:"[64]byte"`
}

func (*GetFirstMsgID) GetMessageName() string {
	return "get_first_msg_id"
}
func (*GetFirstMsgID) GetCrcString() string {
	return "0cb71b0e"
}
func (*GetFirstMsgID) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// GetFirstMsgIDReply represents VPP binary API message 'get_first_msg_id_reply':
type GetFirstMsgIDReply struct {
	Retval     int32
	FirstMsgID uint16
}

func (*GetFirstMsgIDReply) GetMessageName() string {
	return "get_first_msg_id_reply"
}
func (*GetFirstMsgIDReply) GetCrcString() string {
	return "7d337472"
}
func (*GetFirstMsgIDReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MemclntCreate represents VPP binary API message 'memclnt_create':
type MemclntCreate struct {
	CtxQuota    int32
	InputQueue  uint64
	Name        []byte   `struc:"[64]byte"`
	APIVersions []uint32 `struc:"[8]uint32"`
}

func (*MemclntCreate) GetMessageName() string {
	return "memclnt_create"
}
func (*MemclntCreate) GetCrcString() string {
	return "6d33c5ea"
}
func (*MemclntCreate) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MemclntCreateReply represents VPP binary API message 'memclnt_create_reply':
type MemclntCreateReply struct {
	Response     int32
	Handle       uint64
	Index        uint32
	MessageTable uint64
}

func (*MemclntCreateReply) GetMessageName() string {
	return "memclnt_create_reply"
}
func (*MemclntCreateReply) GetCrcString() string {
	return "42ec4560"
}
func (*MemclntCreateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MemclntDelete represents VPP binary API message 'memclnt_delete':
type MemclntDelete struct {
	Index     uint32
	Handle    uint64
	DoCleanup uint8
}

func (*MemclntDelete) GetMessageName() string {
	return "memclnt_delete"
}
func (*MemclntDelete) GetCrcString() string {
	return "4dd351e9"
}
func (*MemclntDelete) GetMessageType() api.MessageType {
	return api.OtherMessage
}

// MemclntDeleteReply represents VPP binary API message 'memclnt_delete_reply':
type MemclntDeleteReply struct {
	Response int32
	Handle   uint64
}

func (*MemclntDeleteReply) GetMessageName() string {
	return "memclnt_delete_reply"
}
func (*MemclntDeleteReply) GetCrcString() string {
	return "3d3b6312"
}
func (*MemclntDeleteReply) GetMessageType() api.MessageType {
	return api.OtherMessage
}

// MemclntKeepalive represents VPP binary API message 'memclnt_keepalive':
type MemclntKeepalive struct{}

func (*MemclntKeepalive) GetMessageName() string {
	return "memclnt_keepalive"
}
func (*MemclntKeepalive) GetCrcString() string {
	return "51077d14"
}
func (*MemclntKeepalive) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// MemclntKeepaliveReply represents VPP binary API message 'memclnt_keepalive_reply':
type MemclntKeepaliveReply struct {
	Retval int32
}

func (*MemclntKeepaliveReply) GetMessageName() string {
	return "memclnt_keepalive_reply"
}
func (*MemclntKeepaliveReply) GetCrcString() string {
	return "e8d4e804"
}
func (*MemclntKeepaliveReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// MemclntReadTimeout represents VPP binary API message 'memclnt_read_timeout':
type MemclntReadTimeout struct {
	Dummy uint8
}

func (*MemclntReadTimeout) GetMessageName() string {
	return "memclnt_read_timeout"
}
func (*MemclntReadTimeout) GetCrcString() string {
	return "c3a3a452"
}
func (*MemclntReadTimeout) GetMessageType() api.MessageType {
	return api.OtherMessage
}

// MemclntRxThreadSuspend represents VPP binary API message 'memclnt_rx_thread_suspend':
type MemclntRxThreadSuspend struct {
	Dummy uint8
}

func (*MemclntRxThreadSuspend) GetMessageName() string {
	return "memclnt_rx_thread_suspend"
}
func (*MemclntRxThreadSuspend) GetCrcString() string {
	return "c3a3a452"
}
func (*MemclntRxThreadSuspend) GetMessageType() api.MessageType {
	return api.OtherMessage
}

// RPCCall represents VPP binary API message 'rpc_call':
type RPCCall struct {
	Function        uint64
	Multicast       uint8
	NeedBarrierSync uint8
	SendReply       uint8
	DataLen         uint32 `struc:"sizeof=Data"`
	Data            []byte
}

func (*RPCCall) GetMessageName() string {
	return "rpc_call"
}
func (*RPCCall) GetCrcString() string {
	return "7e8a2c95"
}
func (*RPCCall) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// RPCCallReply represents VPP binary API message 'rpc_call_reply':
type RPCCallReply struct {
	Retval int32
}

func (*RPCCallReply) GetMessageName() string {
	return "rpc_call_reply"
}
func (*RPCCallReply) GetCrcString() string {
	return "e8d4e804"
}
func (*RPCCallReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// RxThreadExit represents VPP binary API message 'rx_thread_exit':
type RxThreadExit struct {
	Dummy uint8
}

func (*RxThreadExit) GetMessageName() string {
	return "rx_thread_exit"
}
func (*RxThreadExit) GetCrcString() string {
	return "c3a3a452"
}
func (*RxThreadExit) GetMessageType() api.MessageType {
	return api.OtherMessage
}

// SockInitShm represents VPP binary API message 'sock_init_shm':
type SockInitShm struct {
	RequestedSize uint32
	Nitems        uint8 `struc:"sizeof=Configs"`
	Configs       []uint64
}

func (*SockInitShm) GetMessageName() string {
	return "sock_init_shm"
}
func (*SockInitShm) GetCrcString() string {
	return "51646d92"
}
func (*SockInitShm) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SockInitShmReply represents VPP binary API message 'sock_init_shm_reply':
type SockInitShmReply struct {
	Retval int32
}

func (*SockInitShmReply) GetMessageName() string {
	return "sock_init_shm_reply"
}
func (*SockInitShmReply) GetCrcString() string {
	return "e8d4e804"
}
func (*SockInitShmReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// SockclntCreate represents VPP binary API message 'sockclnt_create':
type SockclntCreate struct {
	Name []byte `struc:"[64]byte"`
}

func (*SockclntCreate) GetMessageName() string {
	return "sockclnt_create"
}
func (*SockclntCreate) GetCrcString() string {
	return "df2cf94d"
}
func (*SockclntCreate) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// SockclntCreateReply represents VPP binary API message 'sockclnt_create_reply':
type SockclntCreateReply struct {
	Response     int32
	Index        uint32
	Count        uint16 `struc:"sizeof=MessageTable"`
	MessageTable []MessageTableEntry
}

func (*SockclntCreateReply) GetMessageName() string {
	return "sockclnt_create_reply"
}
func (*SockclntCreateReply) GetCrcString() string {
	return "a134a8a8"
}
func (*SockclntCreateReply) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SockclntDelete represents VPP binary API message 'sockclnt_delete':
type SockclntDelete struct {
	Index uint32
}

func (*SockclntDelete) GetMessageName() string {
	return "sockclnt_delete"
}
func (*SockclntDelete) GetCrcString() string {
	return "8ac76db6"
}
func (*SockclntDelete) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SockclntDeleteReply represents VPP binary API message 'sockclnt_delete_reply':
type SockclntDeleteReply struct {
	Response int32
}

func (*SockclntDeleteReply) GetMessageName() string {
	return "sockclnt_delete_reply"
}
func (*SockclntDeleteReply) GetCrcString() string {
	return "8f38b1ee"
}
func (*SockclntDeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// TracePluginMsgIds represents VPP binary API message 'trace_plugin_msg_ids':
type TracePluginMsgIds struct {
	PluginName []byte `struc:"[128]byte"`
	FirstMsgID uint16
	LastMsgID  uint16
}

func (*TracePluginMsgIds) GetMessageName() string {
	return "trace_plugin_msg_ids"
}
func (*TracePluginMsgIds) GetCrcString() string {
	return "64af79f9"
}
func (*TracePluginMsgIds) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func init() {
	api.RegisterMessage((*APIVersions)(nil), "memclnt.APIVersions")
	api.RegisterMessage((*APIVersionsReply)(nil), "memclnt.APIVersionsReply")
	api.RegisterMessage((*GetFirstMsgID)(nil), "memclnt.GetFirstMsgID")
	api.RegisterMessage((*GetFirstMsgIDReply)(nil), "memclnt.GetFirstMsgIDReply")
	api.RegisterMessage((*MemclntCreate)(nil), "memclnt.MemclntCreate")
	api.RegisterMessage((*MemclntCreateReply)(nil), "memclnt.MemclntCreateReply")
	api.RegisterMessage((*MemclntDelete)(nil), "memclnt.MemclntDelete")
	api.RegisterMessage((*MemclntDeleteReply)(nil), "memclnt.MemclntDeleteReply")
	api.RegisterMessage((*MemclntKeepalive)(nil), "memclnt.MemclntKeepalive")
	api.RegisterMessage((*MemclntKeepaliveReply)(nil), "memclnt.MemclntKeepaliveReply")
	api.RegisterMessage((*MemclntReadTimeout)(nil), "memclnt.MemclntReadTimeout")
	api.RegisterMessage((*MemclntRxThreadSuspend)(nil), "memclnt.MemclntRxThreadSuspend")
	api.RegisterMessage((*RPCCall)(nil), "memclnt.RPCCall")
	api.RegisterMessage((*RPCCallReply)(nil), "memclnt.RPCCallReply")
	api.RegisterMessage((*RxThreadExit)(nil), "memclnt.RxThreadExit")
	api.RegisterMessage((*SockInitShm)(nil), "memclnt.SockInitShm")
	api.RegisterMessage((*SockInitShmReply)(nil), "memclnt.SockInitShmReply")
	api.RegisterMessage((*SockclntCreate)(nil), "memclnt.SockclntCreate")
	api.RegisterMessage((*SockclntCreateReply)(nil), "memclnt.SockclntCreateReply")
	api.RegisterMessage((*SockclntDelete)(nil), "memclnt.SockclntDelete")
	api.RegisterMessage((*SockclntDeleteReply)(nil), "memclnt.SockclntDeleteReply")
	api.RegisterMessage((*TracePluginMsgIds)(nil), "memclnt.TracePluginMsgIds")
}

var Messages = []api.Message{
	(*APIVersions)(nil),
	(*APIVersionsReply)(nil),
	(*GetFirstMsgID)(nil),
	(*GetFirstMsgIDReply)(nil),
	(*MemclntCreate)(nil),
	(*MemclntCreateReply)(nil),
	(*MemclntDelete)(nil),
	(*MemclntDeleteReply)(nil),
	(*MemclntKeepalive)(nil),
	(*MemclntKeepaliveReply)(nil),
	(*MemclntReadTimeout)(nil),
	(*MemclntRxThreadSuspend)(nil),
	(*RPCCall)(nil),
	(*RPCCallReply)(nil),
	(*RxThreadExit)(nil),
	(*SockInitShm)(nil),
	(*SockInitShmReply)(nil),
	(*SockclntCreate)(nil),
	(*SockclntCreateReply)(nil),
	(*SockclntDelete)(nil),
	(*SockclntDeleteReply)(nil),
	(*TracePluginMsgIds)(nil),
}
