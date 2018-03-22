// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Config struct {
	Title           string      `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Loglevel        string      `protobuf:"bytes,2,opt,name=loglevel" json:"loglevel,omitempty"`
	LogConsoleLevel string      `protobuf:"bytes,3,opt,name=logConsoleLevel" json:"logConsoleLevel,omitempty"`
	LogFile         string      `protobuf:"bytes,4,opt,name=logFile" json:"logFile,omitempty"`
	Store           *Store      `protobuf:"bytes,6,opt,name=store" json:"store,omitempty"`
	LocalStore      *LocalStore `protobuf:"bytes,7,opt,name=localStore" json:"localStore,omitempty"`
	Consensus       *Consensus  `protobuf:"bytes,8,opt,name=consensus" json:"consensus,omitempty"`
	MemPool         *MemPool    `protobuf:"bytes,9,opt,name=memPool" json:"memPool,omitempty"`
	BlockChain      *BlockChain `protobuf:"bytes,10,opt,name=blockChain" json:"blockChain,omitempty"`
	Wallet          *Wallet     `protobuf:"bytes,11,opt,name=wallet" json:"wallet,omitempty"`
	P2P             *P2P        `protobuf:"bytes,12,opt,name=p2p" json:"p2p,omitempty"`
	Rpc             *Rpc        `protobuf:"bytes,13,opt,name=rpc" json:"rpc,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Config) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Config) GetLoglevel() string {
	if m != nil {
		return m.Loglevel
	}
	return ""
}

func (m *Config) GetLogConsoleLevel() string {
	if m != nil {
		return m.LogConsoleLevel
	}
	return ""
}

func (m *Config) GetLogFile() string {
	if m != nil {
		return m.LogFile
	}
	return ""
}

func (m *Config) GetStore() *Store {
	if m != nil {
		return m.Store
	}
	return nil
}

func (m *Config) GetLocalStore() *LocalStore {
	if m != nil {
		return m.LocalStore
	}
	return nil
}

func (m *Config) GetConsensus() *Consensus {
	if m != nil {
		return m.Consensus
	}
	return nil
}

func (m *Config) GetMemPool() *MemPool {
	if m != nil {
		return m.MemPool
	}
	return nil
}

func (m *Config) GetBlockChain() *BlockChain {
	if m != nil {
		return m.BlockChain
	}
	return nil
}

func (m *Config) GetWallet() *Wallet {
	if m != nil {
		return m.Wallet
	}
	return nil
}

func (m *Config) GetP2P() *P2P {
	if m != nil {
		return m.P2P
	}
	return nil
}

func (m *Config) GetRpc() *Rpc {
	if m != nil {
		return m.Rpc
	}
	return nil
}

type MemPool struct {
	PoolCacheSize int64 `protobuf:"varint,1,opt,name=poolCacheSize" json:"poolCacheSize,omitempty"`
	MinTxFee      int64 `protobuf:"varint,2,opt,name=minTxFee" json:"minTxFee,omitempty"`
}

func (m *MemPool) Reset()                    { *m = MemPool{} }
func (m *MemPool) String() string            { return proto.CompactTextString(m) }
func (*MemPool) ProtoMessage()               {}
func (*MemPool) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *MemPool) GetPoolCacheSize() int64 {
	if m != nil {
		return m.PoolCacheSize
	}
	return 0
}

func (m *MemPool) GetMinTxFee() int64 {
	if m != nil {
		return m.MinTxFee
	}
	return 0
}

type Consensus struct {
	Name             string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Genesis          string `protobuf:"bytes,2,opt,name=genesis" json:"genesis,omitempty"`
	Minerstart       bool   `protobuf:"varint,3,opt,name=minerstart" json:"minerstart,omitempty"`
	GenesisBlockTime int64  `protobuf:"varint,4,opt,name=genesisBlockTime" json:"genesisBlockTime,omitempty"`
	HotkeyAddr       string `protobuf:"bytes,5,opt,name=hotkeyAddr" json:"hotkeyAddr,omitempty"`
}

func (m *Consensus) Reset()                    { *m = Consensus{} }
func (m *Consensus) String() string            { return proto.CompactTextString(m) }
func (*Consensus) ProtoMessage()               {}
func (*Consensus) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *Consensus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Consensus) GetGenesis() string {
	if m != nil {
		return m.Genesis
	}
	return ""
}

func (m *Consensus) GetMinerstart() bool {
	if m != nil {
		return m.Minerstart
	}
	return false
}

func (m *Consensus) GetGenesisBlockTime() int64 {
	if m != nil {
		return m.GenesisBlockTime
	}
	return 0
}

func (m *Consensus) GetHotkeyAddr() string {
	if m != nil {
		return m.HotkeyAddr
	}
	return ""
}

type Wallet struct {
	MinFee   int64  `protobuf:"varint,1,opt,name=minFee" json:"minFee,omitempty"`
	Driver   string `protobuf:"bytes,2,opt,name=driver" json:"driver,omitempty"`
	DbPath   string `protobuf:"bytes,3,opt,name=dbPath" json:"dbPath,omitempty"`
	SignType string `protobuf:"bytes,4,opt,name=signType" json:"signType,omitempty"`
}

func (m *Wallet) Reset()                    { *m = Wallet{} }
func (m *Wallet) String() string            { return proto.CompactTextString(m) }
func (*Wallet) ProtoMessage()               {}
func (*Wallet) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *Wallet) GetMinFee() int64 {
	if m != nil {
		return m.MinFee
	}
	return 0
}

func (m *Wallet) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *Wallet) GetDbPath() string {
	if m != nil {
		return m.DbPath
	}
	return ""
}

func (m *Wallet) GetSignType() string {
	if m != nil {
		return m.SignType
	}
	return ""
}

type Store struct {
	Driver string `protobuf:"bytes,1,opt,name=driver" json:"driver,omitempty"`
	DbPath string `protobuf:"bytes,2,opt,name=dbPath" json:"dbPath,omitempty"`
}

func (m *Store) Reset()                    { *m = Store{} }
func (m *Store) String() string            { return proto.CompactTextString(m) }
func (*Store) ProtoMessage()               {}
func (*Store) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *Store) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *Store) GetDbPath() string {
	if m != nil {
		return m.DbPath
	}
	return ""
}

type LocalStore struct {
	Driver string `protobuf:"bytes,1,opt,name=driver" json:"driver,omitempty"`
	DbPath string `protobuf:"bytes,2,opt,name=dbPath" json:"dbPath,omitempty"`
}

func (m *LocalStore) Reset()                    { *m = LocalStore{} }
func (m *LocalStore) String() string            { return proto.CompactTextString(m) }
func (*LocalStore) ProtoMessage()               {}
func (*LocalStore) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *LocalStore) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *LocalStore) GetDbPath() string {
	if m != nil {
		return m.DbPath
	}
	return ""
}

type BlockChain struct {
	DefCacheSize        int64  `protobuf:"varint,1,opt,name=defCacheSize" json:"defCacheSize,omitempty"`
	MaxFetchBlockNum    int64  `protobuf:"varint,2,opt,name=maxFetchBlockNum" json:"maxFetchBlockNum,omitempty"`
	TimeoutSeconds      int64  `protobuf:"varint,3,opt,name=timeoutSeconds" json:"timeoutSeconds,omitempty"`
	BatchBlockNum       int64  `protobuf:"varint,4,opt,name=batchBlockNum" json:"batchBlockNum,omitempty"`
	Driver              string `protobuf:"bytes,5,opt,name=driver" json:"driver,omitempty"`
	DbPath              string `protobuf:"bytes,6,opt,name=dbPath" json:"dbPath,omitempty"`
	IsStrongConsistency bool   `protobuf:"varint,7,opt,name=isStrongConsistency" json:"isStrongConsistency,omitempty"`
	SingleMode          bool   `protobuf:"varint,8,opt,name=singleMode" json:"singleMode,omitempty"`
}

func (m *BlockChain) Reset()                    { *m = BlockChain{} }
func (m *BlockChain) String() string            { return proto.CompactTextString(m) }
func (*BlockChain) ProtoMessage()               {}
func (*BlockChain) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

func (m *BlockChain) GetDefCacheSize() int64 {
	if m != nil {
		return m.DefCacheSize
	}
	return 0
}

func (m *BlockChain) GetMaxFetchBlockNum() int64 {
	if m != nil {
		return m.MaxFetchBlockNum
	}
	return 0
}

func (m *BlockChain) GetTimeoutSeconds() int64 {
	if m != nil {
		return m.TimeoutSeconds
	}
	return 0
}

func (m *BlockChain) GetBatchBlockNum() int64 {
	if m != nil {
		return m.BatchBlockNum
	}
	return 0
}

func (m *BlockChain) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *BlockChain) GetDbPath() string {
	if m != nil {
		return m.DbPath
	}
	return ""
}

func (m *BlockChain) GetIsStrongConsistency() bool {
	if m != nil {
		return m.IsStrongConsistency
	}
	return false
}

func (m *BlockChain) GetSingleMode() bool {
	if m != nil {
		return m.SingleMode
	}
	return false
}

type P2P struct {
	SeedPort     int32    `protobuf:"varint,1,opt,name=seedPort" json:"seedPort,omitempty"`
	DbPath       string   `protobuf:"bytes,2,opt,name=dbPath" json:"dbPath,omitempty"`
	GrpcLogFile  string   `protobuf:"bytes,3,opt,name=grpcLogFile" json:"grpcLogFile,omitempty"`
	IsSeed       bool     `protobuf:"varint,4,opt,name=isSeed" json:"isSeed,omitempty"`
	Seeds        []string `protobuf:"bytes,5,rep,name=seeds" json:"seeds,omitempty"`
	Enable       bool     `protobuf:"varint,6,opt,name=enable" json:"enable,omitempty"`
	MsgCacheSize int32    `protobuf:"varint,7,opt,name=msgCacheSize" json:"msgCacheSize,omitempty"`
	Version      int32    `protobuf:"varint,8,opt,name=version" json:"version,omitempty"`
	VerMix       int32    `protobuf:"varint,9,opt,name=verMix" json:"verMix,omitempty"`
	VerMax       int32    `protobuf:"varint,10,opt,name=verMax" json:"verMax,omitempty"`
}

func (m *P2P) Reset()                    { *m = P2P{} }
func (m *P2P) String() string            { return proto.CompactTextString(m) }
func (*P2P) ProtoMessage()               {}
func (*P2P) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

func (m *P2P) GetSeedPort() int32 {
	if m != nil {
		return m.SeedPort
	}
	return 0
}

func (m *P2P) GetDbPath() string {
	if m != nil {
		return m.DbPath
	}
	return ""
}

func (m *P2P) GetGrpcLogFile() string {
	if m != nil {
		return m.GrpcLogFile
	}
	return ""
}

func (m *P2P) GetIsSeed() bool {
	if m != nil {
		return m.IsSeed
	}
	return false
}

func (m *P2P) GetSeeds() []string {
	if m != nil {
		return m.Seeds
	}
	return nil
}

func (m *P2P) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *P2P) GetMsgCacheSize() int32 {
	if m != nil {
		return m.MsgCacheSize
	}
	return 0
}

func (m *P2P) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *P2P) GetVerMix() int32 {
	if m != nil {
		return m.VerMix
	}
	return 0
}

func (m *P2P) GetVerMax() int32 {
	if m != nil {
		return m.VerMax
	}
	return 0
}

type Rpc struct {
	JrpcBindAddr string   `protobuf:"bytes,1,opt,name=jrpcBindAddr" json:"jrpcBindAddr,omitempty"`
	GrpcBindAddr string   `protobuf:"bytes,2,opt,name=grpcBindAddr" json:"grpcBindAddr,omitempty"`
	Whitlist     []string `protobuf:"bytes,3,rep,name=whitlist" json:"whitlist,omitempty"`
}

func (m *Rpc) Reset()                    { *m = Rpc{} }
func (m *Rpc) String() string            { return proto.CompactTextString(m) }
func (*Rpc) ProtoMessage()               {}
func (*Rpc) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{8} }

func (m *Rpc) GetJrpcBindAddr() string {
	if m != nil {
		return m.JrpcBindAddr
	}
	return ""
}

func (m *Rpc) GetGrpcBindAddr() string {
	if m != nil {
		return m.GrpcBindAddr
	}
	return ""
}

func (m *Rpc) GetWhitlist() []string {
	if m != nil {
		return m.Whitlist
	}
	return nil
}

func init() {
	proto.RegisterType((*Config)(nil), "types.Config")
	proto.RegisterType((*MemPool)(nil), "types.MemPool")
	proto.RegisterType((*Consensus)(nil), "types.Consensus")
	proto.RegisterType((*Wallet)(nil), "types.Wallet")
	proto.RegisterType((*Store)(nil), "types.Store")
	proto.RegisterType((*LocalStore)(nil), "types.LocalStore")
	proto.RegisterType((*BlockChain)(nil), "types.BlockChain")
	proto.RegisterType((*P2P)(nil), "types.P2P")
	proto.RegisterType((*Rpc)(nil), "types.Rpc")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 731 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xcf, 0x6e, 0xfb, 0x44,
	0x10, 0x56, 0xe2, 0x9f, 0xf3, 0x67, 0x9a, 0x96, 0xb2, 0x20, 0x64, 0x21, 0x84, 0x22, 0x0b, 0x50,
	0xc4, 0x21, 0x82, 0x70, 0xe0, 0xc2, 0x85, 0x46, 0xea, 0x85, 0x16, 0x45, 0x9b, 0x4a, 0x9c, 0x1d,
	0x7b, 0xea, 0x2c, 0x5d, 0xef, 0x5a, 0xde, 0x6d, 0x9a, 0xf0, 0x12, 0x3c, 0x05, 0x37, 0x5e, 0x83,
	0xf7, 0x42, 0x3b, 0x5e, 0x3b, 0x76, 0x4b, 0x0f, 0xdc, 0xf2, 0x7d, 0xf3, 0xed, 0x78, 0x67, 0xe6,
	0xdb, 0x09, 0xcc, 0x52, 0xad, 0x1e, 0x45, 0xbe, 0x2c, 0x2b, 0x6d, 0x35, 0x0b, 0xed, 0xa9, 0x44,
	0x13, 0xff, 0x13, 0xc0, 0x68, 0x4d, 0x3c, 0xfb, 0x14, 0x42, 0x2b, 0xac, 0xc4, 0x68, 0x30, 0x1f,
	0x2c, 0xa6, 0xbc, 0x06, 0xec, 0x73, 0x98, 0x48, 0x9d, 0x4b, 0x3c, 0xa0, 0x8c, 0x86, 0x14, 0x68,
	0x31, 0x5b, 0xc0, 0x47, 0x52, 0xe7, 0x6b, 0xad, 0x8c, 0x96, 0x78, 0x47, 0x92, 0x80, 0x24, 0xaf,
	0x69, 0x16, 0xc1, 0x58, 0xea, 0xfc, 0x56, 0x48, 0x8c, 0x3e, 0x90, 0xa2, 0x81, 0x2c, 0x86, 0xd0,
	0x58, 0x5d, 0x61, 0x34, 0x9a, 0x0f, 0x16, 0x17, 0xab, 0xd9, 0x92, 0xee, 0xb5, 0xdc, 0x3a, 0x8e,
	0xd7, 0x21, 0xf6, 0x3d, 0x80, 0xd4, 0x69, 0x22, 0x89, 0x8c, 0xc6, 0x24, 0xfc, 0xd8, 0x0b, 0xef,
	0xda, 0x00, 0xef, 0x88, 0xd8, 0x12, 0xa6, 0xa9, 0x56, 0x06, 0x95, 0x79, 0x36, 0xd1, 0x84, 0x4e,
	0x5c, 0xfb, 0x13, 0xeb, 0x86, 0xe7, 0x67, 0x09, 0x5b, 0xc0, 0xb8, 0xc0, 0x62, 0xa3, 0xb5, 0x8c,
	0xa6, 0xa4, 0xbe, 0xf2, 0xea, 0xfb, 0x9a, 0xe5, 0x4d, 0xd8, 0x5d, 0x66, 0x27, 0x75, 0xfa, 0xb4,
	0xde, 0x27, 0x42, 0x45, 0xd0, 0xbb, 0xcc, 0x4d, 0x1b, 0xe0, 0x1d, 0x11, 0xfb, 0x1a, 0x46, 0x2f,
	0x89, 0x94, 0x68, 0xa3, 0x0b, 0x92, 0x5f, 0x7a, 0xf9, 0x6f, 0x44, 0x72, 0x1f, 0x64, 0x5f, 0x40,
	0x50, 0xae, 0xca, 0x68, 0x46, 0x1a, 0xf0, 0x9a, 0xcd, 0x6a, 0xc3, 0x1d, 0xed, 0xa2, 0x55, 0x99,
	0x46, 0x97, 0xbd, 0x28, 0x2f, 0x53, 0xee, 0xe8, 0xf8, 0x17, 0x18, 0xfb, 0x9b, 0xb2, 0xaf, 0xe0,
	0xb2, 0xd4, 0x5a, 0xae, 0x93, 0x74, 0x8f, 0x5b, 0xf1, 0x47, 0x3d, 0xcf, 0x80, 0xf7, 0x49, 0x37,
	0xd7, 0x42, 0xa8, 0x87, 0xe3, 0x2d, 0x22, 0xcd, 0x35, 0xe0, 0x2d, 0x8e, 0xff, 0x1a, 0xc0, 0xb4,
	0xed, 0x12, 0x63, 0xf0, 0x41, 0x25, 0x45, 0x63, 0x0b, 0xfa, 0xed, 0xe6, 0x99, 0xa3, 0x42, 0x23,
	0x8c, 0x37, 0x45, 0x03, 0xd9, 0x97, 0x00, 0x85, 0x50, 0x58, 0x19, 0x9b, 0x54, 0x96, 0xec, 0x30,
	0xe1, 0x1d, 0x86, 0x7d, 0x0b, 0xd7, 0x5e, 0x4a, 0xcd, 0x7a, 0x10, 0x45, 0x6d, 0x89, 0x80, 0xbf,
	0xe1, 0x5d, 0xae, 0xbd, 0xb6, 0x4f, 0x78, 0xfa, 0x39, 0xcb, 0xaa, 0x28, 0xa4, 0x0f, 0x75, 0x98,
	0x58, 0xc2, 0xa8, 0x6e, 0x21, 0xfb, 0x0c, 0x46, 0x85, 0x50, 0xae, 0x96, 0xba, 0x58, 0x8f, 0x1c,
	0x9f, 0x55, 0xe2, 0x80, 0x95, 0xbf, 0xa6, 0x47, 0xc4, 0xef, 0x36, 0x89, 0xdd, 0x7b, 0xc3, 0x7a,
	0xe4, 0xba, 0x62, 0x44, 0xae, 0x1e, 0x4e, 0x65, 0x63, 0xd4, 0x16, 0xc7, 0x3f, 0x42, 0x58, 0x7b,
	0xeb, 0x9c, 0x74, 0xf0, 0x4e, 0xd2, 0x61, 0x37, 0x69, 0xfc, 0x13, 0xc0, 0xd9, 0xa5, 0xff, 0xfb,
	0xf4, 0xdf, 0x43, 0x80, 0xb3, 0xaf, 0x58, 0x0c, 0xb3, 0x0c, 0x1f, 0x5f, 0x0f, 0xb7, 0xc7, 0xb9,
	0x1e, 0x17, 0xc9, 0xf1, 0x16, 0x6d, 0xba, 0xa7, 0x93, 0xbf, 0x3e, 0x17, 0x7e, 0xc6, 0x6f, 0x78,
	0xf6, 0x0d, 0x5c, 0x59, 0x51, 0xa0, 0x7e, 0xb6, 0x5b, 0x4c, 0xb5, 0xca, 0x0c, 0x75, 0x24, 0xe0,
	0xaf, 0x58, 0xe7, 0xaa, 0x5d, 0xd2, 0x4d, 0x58, 0x0f, 0xad, 0x4f, 0x76, 0x8a, 0x0b, 0xdf, 0x29,
	0x6e, 0xd4, 0xeb, 0xf7, 0x77, 0xf0, 0x89, 0x30, 0x5b, 0x5b, 0x69, 0x45, 0xfb, 0x42, 0x18, 0x8b,
	0x2a, 0x3d, 0xd1, 0x13, 0x9f, 0xf0, 0xff, 0x0a, 0x39, 0x4f, 0x18, 0xa1, 0x72, 0x89, 0xf7, 0x3a,
	0x43, 0x7a, 0xd9, 0x13, 0xde, 0x61, 0xe2, 0x3f, 0x87, 0x10, 0x6c, 0x56, 0x1b, 0x9a, 0x24, 0x62,
	0xb6, 0xd1, 0x95, 0xa5, 0x1e, 0x85, 0xbc, 0xc5, 0xef, 0xb5, 0x9a, 0xcd, 0xe1, 0x22, 0xaf, 0xca,
	0xf4, 0xce, 0x6f, 0xaa, 0xda, 0x1a, 0x5d, 0xca, 0x9d, 0x14, 0x66, 0x8b, 0x98, 0x51, 0xf9, 0x13,
	0xee, 0x91, 0xdb, 0x9d, 0x2e, 0xbb, 0x89, 0xc2, 0x79, 0xe0, 0x76, 0x27, 0x01, 0xa7, 0x46, 0x95,
	0xec, 0x64, 0xbd, 0xdc, 0x26, 0xdc, 0x23, 0x37, 0xc3, 0xc2, 0xe4, 0xe7, 0x19, 0x8e, 0xe9, 0x7e,
	0x3d, 0xce, 0xbd, 0xb0, 0x03, 0x56, 0x46, 0x68, 0x45, 0x45, 0x86, 0xbc, 0x81, 0x2e, 0xeb, 0x01,
	0xab, 0x7b, 0x71, 0xa4, 0x4d, 0x15, 0x72, 0x8f, 0x1a, 0x3e, 0x39, 0xd2, 0x52, 0xf2, 0x7c, 0x72,
	0x8c, 0x05, 0x04, 0xbc, 0x4c, 0xdd, 0x47, 0x7f, 0xaf, 0xca, 0xf4, 0x46, 0xa8, 0x8c, 0x9e, 0x53,
	0xed, 0xbe, 0x1e, 0xe7, 0x34, 0x79, 0x57, 0x53, 0xb7, 0xa7, 0xc7, 0xb9, 0xc6, 0xbe, 0xec, 0x85,
	0x95, 0xc2, 0xb8, 0xe7, 0xed, 0xaa, 0x6d, 0xf1, 0x6e, 0x44, 0xff, 0x2d, 0x3f, 0xfc, 0x1b, 0x00,
	0x00, 0xff, 0xff, 0xb7, 0xe0, 0xc4, 0xa2, 0x6b, 0x06, 0x00, 0x00,
}