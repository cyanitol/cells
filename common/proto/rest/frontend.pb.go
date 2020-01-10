// Code generated by protoc-gen-go. DO NOT EDIT.
// source: frontend.proto

package rest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SettingsMenuRequest struct {
}

func (m *SettingsMenuRequest) Reset()                    { *m = SettingsMenuRequest{} }
func (m *SettingsMenuRequest) String() string            { return proto.CompactTextString(m) }
func (*SettingsMenuRequest) ProtoMessage()               {}
func (*SettingsMenuRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

type SettingsEntryMeta struct {
	IconClass string   `protobuf:"bytes,1,opt,name=IconClass,json=icon_class" json:"IconClass,omitempty"`
	Component string   `protobuf:"bytes,2,opt,name=Component,json=component" json:"Component,omitempty"`
	Props     string   `protobuf:"bytes,3,opt,name=Props,json=props" json:"Props,omitempty"`
	Advanced  bool     `protobuf:"varint,4,opt,name=Advanced,json=advanced" json:"Advanced,omitempty"`
	Indexed   []string `protobuf:"bytes,5,rep,name=Indexed,json=indexed" json:"Indexed,omitempty"`
}

func (m *SettingsEntryMeta) Reset()                    { *m = SettingsEntryMeta{} }
func (m *SettingsEntryMeta) String() string            { return proto.CompactTextString(m) }
func (*SettingsEntryMeta) ProtoMessage()               {}
func (*SettingsEntryMeta) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *SettingsEntryMeta) GetIconClass() string {
	if m != nil {
		return m.IconClass
	}
	return ""
}

func (m *SettingsEntryMeta) GetComponent() string {
	if m != nil {
		return m.Component
	}
	return ""
}

func (m *SettingsEntryMeta) GetProps() string {
	if m != nil {
		return m.Props
	}
	return ""
}

func (m *SettingsEntryMeta) GetAdvanced() bool {
	if m != nil {
		return m.Advanced
	}
	return false
}

func (m *SettingsEntryMeta) GetIndexed() []string {
	if m != nil {
		return m.Indexed
	}
	return nil
}

type SettingsEntry struct {
	Key         string             `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	Label       string             `protobuf:"bytes,2,opt,name=Label,json=LABEL" json:"Label,omitempty"`
	Description string             `protobuf:"bytes,3,opt,name=Description,json=DESCRIPTION" json:"Description,omitempty"`
	Manager     string             `protobuf:"bytes,4,opt,name=Manager,json=MANAGER" json:"Manager,omitempty"`
	Alias       string             `protobuf:"bytes,5,opt,name=Alias,json=ALIAS" json:"Alias,omitempty"`
	Metadata    *SettingsEntryMeta `protobuf:"bytes,6,opt,name=Metadata,json=METADATA" json:"Metadata,omitempty"`
}

func (m *SettingsEntry) Reset()                    { *m = SettingsEntry{} }
func (m *SettingsEntry) String() string            { return proto.CompactTextString(m) }
func (*SettingsEntry) ProtoMessage()               {}
func (*SettingsEntry) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *SettingsEntry) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SettingsEntry) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *SettingsEntry) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SettingsEntry) GetManager() string {
	if m != nil {
		return m.Manager
	}
	return ""
}

func (m *SettingsEntry) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *SettingsEntry) GetMetadata() *SettingsEntryMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type SettingsSection struct {
	Key         string           `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	Label       string           `protobuf:"bytes,2,opt,name=Label,json=LABEL" json:"Label,omitempty"`
	Description string           `protobuf:"bytes,3,opt,name=Description,json=DESCRIPTION" json:"Description,omitempty"`
	Children    []*SettingsEntry `protobuf:"bytes,4,rep,name=Children,json=CHILDREN" json:"Children,omitempty"`
}

func (m *SettingsSection) Reset()                    { *m = SettingsSection{} }
func (m *SettingsSection) String() string            { return proto.CompactTextString(m) }
func (*SettingsSection) ProtoMessage()               {}
func (*SettingsSection) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *SettingsSection) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SettingsSection) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *SettingsSection) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SettingsSection) GetChildren() []*SettingsEntry {
	if m != nil {
		return m.Children
	}
	return nil
}

type SettingsMenuResponse struct {
	RootMetadata *SettingsEntryMeta `protobuf:"bytes,1,opt,name=RootMetadata,json=__metadata__" json:"RootMetadata,omitempty"`
	Sections     []*SettingsSection `protobuf:"bytes,2,rep,name=Sections" json:"Sections,omitempty"`
}

func (m *SettingsMenuResponse) Reset()                    { *m = SettingsMenuResponse{} }
func (m *SettingsMenuResponse) String() string            { return proto.CompactTextString(m) }
func (*SettingsMenuResponse) ProtoMessage()               {}
func (*SettingsMenuResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *SettingsMenuResponse) GetRootMetadata() *SettingsEntryMeta {
	if m != nil {
		return m.RootMetadata
	}
	return nil
}

func (m *SettingsMenuResponse) GetSections() []*SettingsSection {
	if m != nil {
		return m.Sections
	}
	return nil
}

type FrontStateRequest struct {
}

func (m *FrontStateRequest) Reset()                    { *m = FrontStateRequest{} }
func (m *FrontStateRequest) String() string            { return proto.CompactTextString(m) }
func (*FrontStateRequest) ProtoMessage()               {}
func (*FrontStateRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

type FrontStateResponse struct {
}

func (m *FrontStateResponse) Reset()                    { *m = FrontStateResponse{} }
func (m *FrontStateResponse) String() string            { return proto.CompactTextString(m) }
func (*FrontStateResponse) ProtoMessage()               {}
func (*FrontStateResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

type FrontPluginsRequest struct {
	Lang string `protobuf:"bytes,1,opt,name=Lang" json:"Lang,omitempty"`
}

func (m *FrontPluginsRequest) Reset()                    { *m = FrontPluginsRequest{} }
func (m *FrontPluginsRequest) String() string            { return proto.CompactTextString(m) }
func (*FrontPluginsRequest) ProtoMessage()               {}
func (*FrontPluginsRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{7} }

func (m *FrontPluginsRequest) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

type FrontPluginsResponse struct {
}

func (m *FrontPluginsResponse) Reset()                    { *m = FrontPluginsResponse{} }
func (m *FrontPluginsResponse) String() string            { return proto.CompactTextString(m) }
func (*FrontPluginsResponse) ProtoMessage()               {}
func (*FrontPluginsResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{8} }

type FrontMessagesRequest struct {
	Lang string `protobuf:"bytes,1,opt,name=Lang" json:"Lang,omitempty"`
}

func (m *FrontMessagesRequest) Reset()                    { *m = FrontMessagesRequest{} }
func (m *FrontMessagesRequest) String() string            { return proto.CompactTextString(m) }
func (*FrontMessagesRequest) ProtoMessage()               {}
func (*FrontMessagesRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{9} }

func (m *FrontMessagesRequest) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

type FrontMessagesResponse struct {
	Messages map[string]string `protobuf:"bytes,1,rep,name=Messages" json:"Messages,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *FrontMessagesResponse) Reset()                    { *m = FrontMessagesResponse{} }
func (m *FrontMessagesResponse) String() string            { return proto.CompactTextString(m) }
func (*FrontMessagesResponse) ProtoMessage()               {}
func (*FrontMessagesResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{10} }

func (m *FrontMessagesResponse) GetMessages() map[string]string {
	if m != nil {
		return m.Messages
	}
	return nil
}

type Token struct {
	AccessToken  string `protobuf:"bytes,1,opt,name=AccessToken" json:"AccessToken,omitempty"`
	IDToken      string `protobuf:"bytes,2,opt,name=IDToken" json:"IDToken,omitempty"`
	RefreshToken string `protobuf:"bytes,3,opt,name=RefreshToken" json:"RefreshToken,omitempty"`
	ExpiresAt    string `protobuf:"bytes,4,opt,name=ExpiresAt" json:"ExpiresAt,omitempty"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (m *Token) String() string            { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{11} }

func (m *Token) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *Token) GetIDToken() string {
	if m != nil {
		return m.IDToken
	}
	return ""
}

func (m *Token) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *Token) GetExpiresAt() string {
	if m != nil {
		return m.ExpiresAt
	}
	return ""
}

type FrontSessionGetRequest struct {
}

func (m *FrontSessionGetRequest) Reset()                    { *m = FrontSessionGetRequest{} }
func (m *FrontSessionGetRequest) String() string            { return proto.CompactTextString(m) }
func (*FrontSessionGetRequest) ProtoMessage()               {}
func (*FrontSessionGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{12} }

type FrontSessionGetResponse struct {
	Token *Token `protobuf:"bytes,1,opt,name=Token" json:"Token,omitempty"`
}

func (m *FrontSessionGetResponse) Reset()                    { *m = FrontSessionGetResponse{} }
func (m *FrontSessionGetResponse) String() string            { return proto.CompactTextString(m) }
func (*FrontSessionGetResponse) ProtoMessage()               {}
func (*FrontSessionGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{13} }

func (m *FrontSessionGetResponse) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type FrontSessionRequest struct {
	// Time reference for computing jwt expiry
	ClientTime int32 `protobuf:"varint,1,opt,name=ClientTime" json:"ClientTime,omitempty"`
	// Data sent back by specific auth steps
	AuthInfo map[string]string `protobuf:"bytes,2,rep,name=AuthInfo" json:"AuthInfo,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Kill session now
	Logout bool `protobuf:"varint,3,opt,name=Logout" json:"Logout,omitempty"`
}

func (m *FrontSessionRequest) Reset()                    { *m = FrontSessionRequest{} }
func (m *FrontSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*FrontSessionRequest) ProtoMessage()               {}
func (*FrontSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{14} }

func (m *FrontSessionRequest) GetClientTime() int32 {
	if m != nil {
		return m.ClientTime
	}
	return 0
}

func (m *FrontSessionRequest) GetAuthInfo() map[string]string {
	if m != nil {
		return m.AuthInfo
	}
	return nil
}

func (m *FrontSessionRequest) GetLogout() bool {
	if m != nil {
		return m.Logout
	}
	return false
}

type FrontSessionResponse struct {
	Token *Token `protobuf:"bytes,1,opt,name=Token" json:"Token,omitempty"`
	// Trigger a specific Auth step
	Trigger string `protobuf:"bytes,2,opt,name=Trigger" json:"Trigger,omitempty"`
	// Additional data for the trigger
	TriggerInfo map[string]string `protobuf:"bytes,3,rep,name=TriggerInfo" json:"TriggerInfo,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	RedirectTo  string            `protobuf:"bytes,4,opt,name=RedirectTo" json:"RedirectTo,omitempty"`
}

func (m *FrontSessionResponse) Reset()                    { *m = FrontSessionResponse{} }
func (m *FrontSessionResponse) String() string            { return proto.CompactTextString(m) }
func (*FrontSessionResponse) ProtoMessage()               {}
func (*FrontSessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{15} }

func (m *FrontSessionResponse) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *FrontSessionResponse) GetTrigger() string {
	if m != nil {
		return m.Trigger
	}
	return ""
}

func (m *FrontSessionResponse) GetTriggerInfo() map[string]string {
	if m != nil {
		return m.TriggerInfo
	}
	return nil
}

func (m *FrontSessionResponse) GetRedirectTo() string {
	if m != nil {
		return m.RedirectTo
	}
	return ""
}

type FrontSessionDelRequest struct {
}

func (m *FrontSessionDelRequest) Reset()                    { *m = FrontSessionDelRequest{} }
func (m *FrontSessionDelRequest) String() string            { return proto.CompactTextString(m) }
func (*FrontSessionDelRequest) ProtoMessage()               {}
func (*FrontSessionDelRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{16} }

type FrontSessionDelResponse struct {
}

func (m *FrontSessionDelResponse) Reset()                    { *m = FrontSessionDelResponse{} }
func (m *FrontSessionDelResponse) String() string            { return proto.CompactTextString(m) }
func (*FrontSessionDelResponse) ProtoMessage()               {}
func (*FrontSessionDelResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{17} }

type FrontAuthRequest struct {
	RequestID string `protobuf:"bytes,1,opt,name=RequestID" json:"RequestID,omitempty"`
}

func (m *FrontAuthRequest) Reset()                    { *m = FrontAuthRequest{} }
func (m *FrontAuthRequest) String() string            { return proto.CompactTextString(m) }
func (*FrontAuthRequest) ProtoMessage()               {}
func (*FrontAuthRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{18} }

func (m *FrontAuthRequest) GetRequestID() string {
	if m != nil {
		return m.RequestID
	}
	return ""
}

type FrontAuthResponse struct {
}

func (m *FrontAuthResponse) Reset()                    { *m = FrontAuthResponse{} }
func (m *FrontAuthResponse) String() string            { return proto.CompactTextString(m) }
func (*FrontAuthResponse) ProtoMessage()               {}
func (*FrontAuthResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{19} }

type FrontEnrollAuthRequest struct {
	EnrollType string            `protobuf:"bytes,1,opt,name=EnrollType" json:"EnrollType,omitempty"`
	EnrollInfo map[string]string `protobuf:"bytes,2,rep,name=EnrollInfo" json:"EnrollInfo,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *FrontEnrollAuthRequest) Reset()                    { *m = FrontEnrollAuthRequest{} }
func (m *FrontEnrollAuthRequest) String() string            { return proto.CompactTextString(m) }
func (*FrontEnrollAuthRequest) ProtoMessage()               {}
func (*FrontEnrollAuthRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{20} }

func (m *FrontEnrollAuthRequest) GetEnrollType() string {
	if m != nil {
		return m.EnrollType
	}
	return ""
}

func (m *FrontEnrollAuthRequest) GetEnrollInfo() map[string]string {
	if m != nil {
		return m.EnrollInfo
	}
	return nil
}

type FrontEnrollAuthResponse struct {
	// Any parameters can be returned
	Info map[string]string `protobuf:"bytes,1,rep,name=Info" json:"Info,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *FrontEnrollAuthResponse) Reset()                    { *m = FrontEnrollAuthResponse{} }
func (m *FrontEnrollAuthResponse) String() string            { return proto.CompactTextString(m) }
func (*FrontEnrollAuthResponse) ProtoMessage()               {}
func (*FrontEnrollAuthResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{21} }

func (m *FrontEnrollAuthResponse) GetInfo() map[string]string {
	if m != nil {
		return m.Info
	}
	return nil
}

// Donwload binary
type FrontBinaryRequest struct {
	// Currently supported values are USER and GLOBAL
	BinaryType string `protobuf:"bytes,1,opt,name=BinaryType" json:"BinaryType,omitempty"`
	// Id of the binary
	Uuid string `protobuf:"bytes,2,opt,name=Uuid" json:"Uuid,omitempty"`
}

func (m *FrontBinaryRequest) Reset()                    { *m = FrontBinaryRequest{} }
func (m *FrontBinaryRequest) String() string            { return proto.CompactTextString(m) }
func (*FrontBinaryRequest) ProtoMessage()               {}
func (*FrontBinaryRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{22} }

func (m *FrontBinaryRequest) GetBinaryType() string {
	if m != nil {
		return m.BinaryType
	}
	return ""
}

func (m *FrontBinaryRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

// Not used, endpoint returns octet-stream
type FrontBinaryResponse struct {
}

func (m *FrontBinaryResponse) Reset()                    { *m = FrontBinaryResponse{} }
func (m *FrontBinaryResponse) String() string            { return proto.CompactTextString(m) }
func (*FrontBinaryResponse) ProtoMessage()               {}
func (*FrontBinaryResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{23} }

type FrontBootConfRequest struct {
}

func (m *FrontBootConfRequest) Reset()                    { *m = FrontBootConfRequest{} }
func (m *FrontBootConfRequest) String() string            { return proto.CompactTextString(m) }
func (*FrontBootConfRequest) ProtoMessage()               {}
func (*FrontBootConfRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{24} }

type FrontBootConfResponse struct {
	JsonData map[string]string `protobuf:"bytes,1,rep,name=JsonData" json:"JsonData,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *FrontBootConfResponse) Reset()                    { *m = FrontBootConfResponse{} }
func (m *FrontBootConfResponse) String() string            { return proto.CompactTextString(m) }
func (*FrontBootConfResponse) ProtoMessage()               {}
func (*FrontBootConfResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{25} }

func (m *FrontBootConfResponse) GetJsonData() map[string]string {
	if m != nil {
		return m.JsonData
	}
	return nil
}

func init() {
	proto.RegisterType((*SettingsMenuRequest)(nil), "rest.SettingsMenuRequest")
	proto.RegisterType((*SettingsEntryMeta)(nil), "rest.SettingsEntryMeta")
	proto.RegisterType((*SettingsEntry)(nil), "rest.SettingsEntry")
	proto.RegisterType((*SettingsSection)(nil), "rest.SettingsSection")
	proto.RegisterType((*SettingsMenuResponse)(nil), "rest.SettingsMenuResponse")
	proto.RegisterType((*FrontStateRequest)(nil), "rest.FrontStateRequest")
	proto.RegisterType((*FrontStateResponse)(nil), "rest.FrontStateResponse")
	proto.RegisterType((*FrontPluginsRequest)(nil), "rest.FrontPluginsRequest")
	proto.RegisterType((*FrontPluginsResponse)(nil), "rest.FrontPluginsResponse")
	proto.RegisterType((*FrontMessagesRequest)(nil), "rest.FrontMessagesRequest")
	proto.RegisterType((*FrontMessagesResponse)(nil), "rest.FrontMessagesResponse")
	proto.RegisterType((*Token)(nil), "rest.Token")
	proto.RegisterType((*FrontSessionGetRequest)(nil), "rest.FrontSessionGetRequest")
	proto.RegisterType((*FrontSessionGetResponse)(nil), "rest.FrontSessionGetResponse")
	proto.RegisterType((*FrontSessionRequest)(nil), "rest.FrontSessionRequest")
	proto.RegisterType((*FrontSessionResponse)(nil), "rest.FrontSessionResponse")
	proto.RegisterType((*FrontSessionDelRequest)(nil), "rest.FrontSessionDelRequest")
	proto.RegisterType((*FrontSessionDelResponse)(nil), "rest.FrontSessionDelResponse")
	proto.RegisterType((*FrontAuthRequest)(nil), "rest.FrontAuthRequest")
	proto.RegisterType((*FrontAuthResponse)(nil), "rest.FrontAuthResponse")
	proto.RegisterType((*FrontEnrollAuthRequest)(nil), "rest.FrontEnrollAuthRequest")
	proto.RegisterType((*FrontEnrollAuthResponse)(nil), "rest.FrontEnrollAuthResponse")
	proto.RegisterType((*FrontBinaryRequest)(nil), "rest.FrontBinaryRequest")
	proto.RegisterType((*FrontBinaryResponse)(nil), "rest.FrontBinaryResponse")
	proto.RegisterType((*FrontBootConfRequest)(nil), "rest.FrontBootConfRequest")
	proto.RegisterType((*FrontBootConfResponse)(nil), "rest.FrontBootConfResponse")
}

func init() { proto.RegisterFile("frontend.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 937 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xcf, 0x8e, 0x1b, 0xc5,
	0x13, 0xd6, 0xac, 0xed, 0xdd, 0x71, 0x79, 0xf3, 0xcb, 0xa6, 0xf7, 0x4f, 0xe6, 0x67, 0x85, 0xc8,
	0xcc, 0x05, 0x07, 0x90, 0x81, 0xe4, 0x00, 0x22, 0x80, 0x34, 0x6b, 0x9b, 0xc4, 0x60, 0x2f, 0xab,
	0xb1, 0x39, 0x5b, 0x93, 0x99, 0x5e, 0xef, 0x68, 0x67, 0xbb, 0xcd, 0x74, 0x3b, 0xca, 0x5e, 0x91,
	0xe0, 0x84, 0xc4, 0x03, 0xc0, 0xb3, 0x70, 0xe3, 0xcc, 0x2b, 0xa1, 0xee, 0xae, 0xf6, 0xcc, 0xd8,
	0xab, 0x88, 0x45, 0xdc, 0xa6, 0xbe, 0xaa, 0xae, 0xae, 0xfa, 0xea, 0xeb, 0xb2, 0xe1, 0x7f, 0x17,
	0x39, 0x67, 0x92, 0xb2, 0xa4, 0xb7, 0xcc, 0xb9, 0xe4, 0xa4, 0x9e, 0x53, 0x21, 0xfd, 0x63, 0x38,
	0x9c, 0x52, 0x29, 0x53, 0xb6, 0x10, 0x13, 0xca, 0x56, 0x21, 0xfd, 0x61, 0xa5, 0xe0, 0xdf, 0x1d,
	0x78, 0x60, 0xf1, 0x21, 0x93, 0xf9, 0xcd, 0x84, 0xca, 0x88, 0xbc, 0x03, 0xcd, 0x51, 0xcc, 0x59,
	0x3f, 0x8b, 0x84, 0xf0, 0x9c, 0x8e, 0xd3, 0x6d, 0x86, 0x90, 0xc6, 0x9c, 0xcd, 0x63, 0x85, 0x90,
	0x47, 0xd0, 0xec, 0xf3, 0xeb, 0x25, 0x67, 0x94, 0x49, 0x6f, 0x47, 0xbb, 0x9b, 0xb1, 0x05, 0xc8,
	0x11, 0x34, 0xce, 0x73, 0xbe, 0x14, 0x5e, 0x4d, 0x7b, 0x1a, 0x4b, 0x65, 0x90, 0x36, 0xb8, 0x41,
	0xf2, 0x3a, 0x62, 0x31, 0x4d, 0xbc, 0x7a, 0xc7, 0xe9, 0xba, 0xa1, 0x1b, 0xa1, 0x4d, 0x3c, 0xd8,
	0x1b, 0xb1, 0x84, 0xbe, 0xa1, 0x89, 0xd7, 0xe8, 0xd4, 0xba, 0xcd, 0x70, 0x2f, 0x35, 0xa6, 0xff,
	0x87, 0x03, 0xf7, 0x2a, 0xe5, 0x91, 0x03, 0xa8, 0x7d, 0x4b, 0x6f, 0xb0, 0x28, 0xf5, 0xa9, 0xee,
	0x1b, 0x47, 0xaf, 0x68, 0x86, 0x95, 0x34, 0xc6, 0xc1, 0xe9, 0x70, 0x4c, 0x3a, 0xd0, 0x1a, 0x50,
	0x11, 0xe7, 0xe9, 0x52, 0xa6, 0x9c, 0x61, 0x2d, 0xad, 0xc1, 0x70, 0xda, 0x0f, 0x47, 0xe7, 0xb3,
	0xd1, 0x77, 0x67, 0xea, 0xd6, 0x49, 0xc4, 0xa2, 0x05, 0xcd, 0x75, 0x41, 0xcd, 0x70, 0x6f, 0x12,
	0x9c, 0x05, 0x2f, 0x86, 0xa1, 0xca, 0x18, 0x64, 0x69, 0x24, 0xbc, 0x86, 0xc9, 0x18, 0x8c, 0x47,
	0xc1, 0x94, 0x3c, 0x03, 0x57, 0x91, 0x93, 0x44, 0x32, 0xf2, 0x76, 0x3b, 0x4e, 0xb7, 0xf5, 0xf4,
	0x61, 0x4f, 0x51, 0xdb, 0xdb, 0xe2, 0x2f, 0x74, 0x27, 0xc3, 0x59, 0x30, 0x08, 0x66, 0x81, 0xff,
	0x8b, 0x03, 0xf7, 0xad, 0x7f, 0x4a, 0x63, 0x55, 0xcb, 0x7f, 0xd8, 0xc2, 0x47, 0xe0, 0xf6, 0x2f,
	0xd3, 0x2c, 0xc9, 0x29, 0xf3, 0xea, 0x9d, 0x5a, 0xb7, 0xf5, 0xf4, 0xf0, 0x96, 0x92, 0x42, 0xb7,
	0xff, 0x72, 0x34, 0x1e, 0x84, 0xc3, 0x33, 0xff, 0x67, 0x07, 0x8e, 0xaa, 0x32, 0x10, 0x4b, 0xce,
	0x04, 0x25, 0xcf, 0x61, 0x3f, 0xe4, 0x5c, 0xae, 0x1b, 0x74, 0xde, 0xde, 0xe0, 0xfe, 0x7c, 0x7e,
	0x8d, 0xa1, 0xf3, 0x39, 0xf9, 0x04, 0x5c, 0xec, 0x4d, 0x78, 0x3b, 0xba, 0x8c, 0xe3, 0xea, 0x41,
	0xf4, 0x86, 0xeb, 0x30, 0xff, 0x10, 0x1e, 0x7c, 0xad, 0x64, 0x3a, 0x95, 0x91, 0xa4, 0x56, 0x8c,
	0x47, 0x40, 0xca, 0xa0, 0x29, 0xcd, 0x7f, 0x02, 0x87, 0x1a, 0x3d, 0xcf, 0x56, 0x8b, 0x94, 0x09,
	0x0c, 0x26, 0x04, 0xea, 0xe3, 0x88, 0x2d, 0x90, 0x46, 0xfd, 0xed, 0x9f, 0xc0, 0x51, 0x35, 0x14,
	0x53, 0xbc, 0x8f, 0xf8, 0x84, 0x0a, 0x11, 0x2d, 0xe8, 0x5b, 0x73, 0xfc, 0xe6, 0xc0, 0xf1, 0x46,
	0x30, 0x72, 0x34, 0x54, 0x02, 0x30, 0x98, 0xe7, 0xe8, 0x36, 0x9f, 0x98, 0x36, 0x6f, 0x0d, 0xef,
	0x59, 0x00, 0x67, 0x60, 0xcd, 0xf6, 0x73, 0xb8, 0x57, 0x71, 0x29, 0x3d, 0x5c, 0x15, 0x7a, 0xb8,
	0x32, 0x7a, 0x78, 0x1d, 0x65, 0x2b, 0x6a, 0xf5, 0xa0, 0x8d, 0xcf, 0x77, 0x3e, 0x73, 0xfc, 0x9f,
	0x1c, 0x68, 0xcc, 0xf8, 0x15, 0x65, 0x4a, 0x1d, 0x41, 0x1c, 0x53, 0x21, 0xb4, 0x89, 0xa7, 0xcb,
	0x90, 0x7e, 0x56, 0x03, 0xe3, 0x35, 0x79, 0xac, 0x49, 0x7c, 0xd8, 0x0f, 0xe9, 0x45, 0x4e, 0xc5,
	0xa5, 0x71, 0x1b, 0x69, 0x55, 0x30, 0xf5, 0xc8, 0x87, 0x6f, 0x96, 0x69, 0x4e, 0x45, 0x20, 0xf1,
	0x81, 0x14, 0x80, 0xef, 0xc1, 0x89, 0x19, 0x15, 0x15, 0x22, 0xe5, 0xec, 0x05, 0x95, 0x76, 0x88,
	0x5f, 0xc0, 0xc3, 0x2d, 0x0f, 0x12, 0xf8, 0x2e, 0xd6, 0x8e, 0xea, 0x6a, 0x19, 0xf6, 0x34, 0x14,
	0x1a, 0x8f, 0xff, 0x97, 0x83, 0xd3, 0xc6, 0xe3, 0x76, 0x52, 0x8f, 0x01, 0xfa, 0x59, 0x4a, 0x99,
	0x9c, 0xa5, 0xd7, 0x54, 0x9f, 0x6f, 0x84, 0x25, 0x84, 0xf4, 0xc1, 0x0d, 0x56, 0xf2, 0x72, 0xc4,
	0x2e, 0x38, 0x4a, 0xf0, 0xbd, 0xd2, 0x6c, 0xaa, 0xc9, 0x7a, 0x36, 0x12, 0x27, 0x63, 0x4d, 0x72,
	0x02, 0xbb, 0x63, 0xbe, 0xe0, 0x2b, 0xa9, 0x09, 0x71, 0x43, 0xb4, 0xd4, 0xc4, 0x2a, 0x47, 0xee,
	0x34, 0xb1, 0x1f, 0x77, 0x50, 0x7c, 0xeb, 0x22, 0xfe, 0x31, 0x1b, 0x6a, 0x82, 0xb3, 0x3c, 0x5d,
	0xa8, 0x15, 0x85, 0x13, 0x44, 0x93, 0x4c, 0xa0, 0x85, 0x9f, 0xba, 0xe5, 0x9a, 0x6e, 0xf9, 0x83,
	0xdb, 0x5a, 0x46, 0x35, 0x96, 0xa2, 0x4d, 0xdb, 0xe5, 0xf3, 0x8a, 0xde, 0x90, 0x26, 0x69, 0x4e,
	0x63, 0x39, 0xe3, 0x38, 0xed, 0x12, 0xd2, 0xfe, 0x0a, 0x0e, 0x36, 0x13, 0xdc, 0x89, 0x84, 0x0d,
	0xb9, 0x0c, 0x68, 0x66, 0xe5, 0xf2, 0xff, 0xaa, 0x5c, 0xb4, 0x07, 0x5f, 0xed, 0xc7, 0x70, 0xa0,
	0x5d, 0x8a, 0x7b, 0xab, 0x83, 0x47, 0xd0, 0xc4, 0xcf, 0xd1, 0x00, 0xaf, 0x2e, 0x80, 0xf5, 0x56,
	0x31, 0x27, 0x30, 0xcd, 0x9f, 0x0e, 0x5e, 0x3e, 0x64, 0x39, 0xcf, 0xb2, 0x72, 0xb6, 0xc7, 0x00,
	0x06, 0x9c, 0xdd, 0x2c, 0xa9, 0xfd, 0xa1, 0x2b, 0x10, 0x32, 0xb6, 0xfe, 0x92, 0xae, 0x3e, 0x2c,
	0x91, 0xbc, 0x95, 0xb1, 0x57, 0x84, 0x1b, 0x96, 0x4b, 0xe7, 0xdb, 0x5f, 0xc2, 0xfd, 0x0d, 0xf7,
	0x9d, 0x38, 0xfc, 0xd5, 0x41, 0xaa, 0xca, 0xb7, 0xae, 0xd7, 0x77, 0x5d, 0x97, 0xe8, 0x6c, 0x49,
	0x7f, 0x3b, 0xb8, 0x57, 0x54, 0xa7, 0x0f, 0xb5, 0x3f, 0x85, 0xe6, 0xbf, 0xab, 0xe8, 0x25, 0xee,
	0xeb, 0xd3, 0x94, 0x45, 0xf9, 0x4d, 0x89, 0x54, 0x03, 0x94, 0x49, 0x2d, 0x10, 0xb5, 0x74, 0xbf,
	0x5f, 0xa5, 0x09, 0xa6, 0xd3, 0xdf, 0xea, 0xdf, 0x49, 0x25, 0x13, 0x8e, 0xce, 0xee, 0xf3, 0x53,
	0xce, 0x65, 0x9f, 0xb3, 0x0b, 0x2b, 0x9a, 0xf5, 0x8e, 0x2e, 0x1c, 0xc5, 0x8e, 0xfe, 0x46, 0x70,
	0x36, 0x30, 0xbf, 0x61, 0x9b, 0x3b, 0x7a, 0x33, 0xbc, 0x67, 0x63, 0x71, 0x13, 0x58, 0x53, 0xbd,
	0xf8, 0x8a, 0xeb, 0x2e, 0xb4, 0xbc, 0xda, 0xd5, 0xff, 0xbb, 0x9e, 0xfd, 0x1d, 0x00, 0x00, 0xff,
	0xff, 0x2d, 0xd5, 0x3b, 0xbf, 0x89, 0x09, 0x00, 0x00,
}
