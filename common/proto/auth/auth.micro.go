// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: auth.proto

/*
Package auth is a generated protocol buffer package.

It is generated from these files:
	auth.proto
	ldap.proto
	oauth.proto

It has these top-level messages:
	Token
	MatchInvalidTokenRequest
	MatchInvalidTokenResponse
	RevokeTokenRequest
	RevokeTokenResponse
	PruneTokensRequest
	PruneTokensResponse
	GetLoginRequest
	GetLoginResponse
	CreateLoginRequest
	CreateLoginResponse
	AcceptLoginRequest
	AcceptLoginResponse
	GetConsentRequest
	GetConsentResponse
	CreateConsentRequest
	CreateConsentResponse
	AcceptConsentRequest
	AcceptConsentResponse
	CreateAuthCodeRequest
	CreateAuthCodeResponse
	VerifyTokenRequest
	VerifyTokenResponse
	ExchangeRequest
	ExchangeResponse
	RefreshTokenRequest
	RefreshTokenResponse
	LdapSearchFilter
	LdapMapping
	LdapMemberOfMapping
	LdapServerConfig
	OAuth2ClientConfig
	OAuth2ConnectorCollection
	OAuth2ConnectorConfig
	OAuth2MappingRule
	OAuth2ConnectorPydioConfig
	OAuth2ConnectorOIDCConfig
	OAuth2ConnectorSAMLConfig
	OAuth2ConnectorBitbucketConfig
	OAuth2ConnectorGithubConfig
	OAuth2ConnectorGithubConfigOrg
	OAuth2ConnectorGitlabConfig
	OAuth2ConnectorLinkedinConfig
	OAuth2ConnectorMicrosoftConfig
*/
package auth

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for AuthTokenRevoker service

type AuthTokenRevokerClient interface {
	// Look for an invalid token entry in the store that match the current one
	MatchInvalid(ctx context.Context, in *MatchInvalidTokenRequest, opts ...client.CallOption) (*MatchInvalidTokenResponse, error)
	// Revoker invalidates the current token and specifies if the invalidation is due to a refresh or a revokation
	Revoke(ctx context.Context, in *RevokeTokenRequest, opts ...client.CallOption) (*RevokeTokenResponse, error)
	// PruneTokens clear revoked tokens
	PruneTokens(ctx context.Context, in *PruneTokensRequest, opts ...client.CallOption) (*PruneTokensResponse, error)
}

type authTokenRevokerClient struct {
	c           client.Client
	serviceName string
}

func NewAuthTokenRevokerClient(serviceName string, c client.Client) AuthTokenRevokerClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "auth"
	}
	return &authTokenRevokerClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *authTokenRevokerClient) MatchInvalid(ctx context.Context, in *MatchInvalidTokenRequest, opts ...client.CallOption) (*MatchInvalidTokenResponse, error) {
	req := c.c.NewRequest(c.serviceName, "AuthTokenRevoker.MatchInvalid", in)
	out := new(MatchInvalidTokenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authTokenRevokerClient) Revoke(ctx context.Context, in *RevokeTokenRequest, opts ...client.CallOption) (*RevokeTokenResponse, error) {
	req := c.c.NewRequest(c.serviceName, "AuthTokenRevoker.Revoke", in)
	out := new(RevokeTokenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authTokenRevokerClient) PruneTokens(ctx context.Context, in *PruneTokensRequest, opts ...client.CallOption) (*PruneTokensResponse, error) {
	req := c.c.NewRequest(c.serviceName, "AuthTokenRevoker.PruneTokens", in)
	out := new(PruneTokensResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthTokenRevoker service

type AuthTokenRevokerHandler interface {
	// Look for an invalid token entry in the store that match the current one
	MatchInvalid(context.Context, *MatchInvalidTokenRequest, *MatchInvalidTokenResponse) error
	// Revoker invalidates the current token and specifies if the invalidation is due to a refresh or a revokation
	Revoke(context.Context, *RevokeTokenRequest, *RevokeTokenResponse) error
	// PruneTokens clear revoked tokens
	PruneTokens(context.Context, *PruneTokensRequest, *PruneTokensResponse) error
}

func RegisterAuthTokenRevokerHandler(s server.Server, hdlr AuthTokenRevokerHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&AuthTokenRevoker{hdlr}, opts...))
}

type AuthTokenRevoker struct {
	AuthTokenRevokerHandler
}

func (h *AuthTokenRevoker) MatchInvalid(ctx context.Context, in *MatchInvalidTokenRequest, out *MatchInvalidTokenResponse) error {
	return h.AuthTokenRevokerHandler.MatchInvalid(ctx, in, out)
}

func (h *AuthTokenRevoker) Revoke(ctx context.Context, in *RevokeTokenRequest, out *RevokeTokenResponse) error {
	return h.AuthTokenRevokerHandler.Revoke(ctx, in, out)
}

func (h *AuthTokenRevoker) PruneTokens(ctx context.Context, in *PruneTokensRequest, out *PruneTokensResponse) error {
	return h.AuthTokenRevokerHandler.PruneTokens(ctx, in, out)
}

// Client API for LoginProvider service

type LoginProviderClient interface {
	GetLogin(ctx context.Context, in *GetLoginRequest, opts ...client.CallOption) (*GetLoginResponse, error)
	CreateLogin(ctx context.Context, in *CreateLoginRequest, opts ...client.CallOption) (*CreateLoginResponse, error)
	AcceptLogin(ctx context.Context, in *AcceptLoginRequest, opts ...client.CallOption) (*AcceptLoginResponse, error)
}

type loginProviderClient struct {
	c           client.Client
	serviceName string
}

func NewLoginProviderClient(serviceName string, c client.Client) LoginProviderClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "auth"
	}
	return &loginProviderClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *loginProviderClient) GetLogin(ctx context.Context, in *GetLoginRequest, opts ...client.CallOption) (*GetLoginResponse, error) {
	req := c.c.NewRequest(c.serviceName, "LoginProvider.GetLogin", in)
	out := new(GetLoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginProviderClient) CreateLogin(ctx context.Context, in *CreateLoginRequest, opts ...client.CallOption) (*CreateLoginResponse, error) {
	req := c.c.NewRequest(c.serviceName, "LoginProvider.CreateLogin", in)
	out := new(CreateLoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginProviderClient) AcceptLogin(ctx context.Context, in *AcceptLoginRequest, opts ...client.CallOption) (*AcceptLoginResponse, error) {
	req := c.c.NewRequest(c.serviceName, "LoginProvider.AcceptLogin", in)
	out := new(AcceptLoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LoginProvider service

type LoginProviderHandler interface {
	GetLogin(context.Context, *GetLoginRequest, *GetLoginResponse) error
	CreateLogin(context.Context, *CreateLoginRequest, *CreateLoginResponse) error
	AcceptLogin(context.Context, *AcceptLoginRequest, *AcceptLoginResponse) error
}

func RegisterLoginProviderHandler(s server.Server, hdlr LoginProviderHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&LoginProvider{hdlr}, opts...))
}

type LoginProvider struct {
	LoginProviderHandler
}

func (h *LoginProvider) GetLogin(ctx context.Context, in *GetLoginRequest, out *GetLoginResponse) error {
	return h.LoginProviderHandler.GetLogin(ctx, in, out)
}

func (h *LoginProvider) CreateLogin(ctx context.Context, in *CreateLoginRequest, out *CreateLoginResponse) error {
	return h.LoginProviderHandler.CreateLogin(ctx, in, out)
}

func (h *LoginProvider) AcceptLogin(ctx context.Context, in *AcceptLoginRequest, out *AcceptLoginResponse) error {
	return h.LoginProviderHandler.AcceptLogin(ctx, in, out)
}

// Client API for ConsentProvider service

type ConsentProviderClient interface {
	GetConsent(ctx context.Context, in *GetConsentRequest, opts ...client.CallOption) (*GetConsentResponse, error)
	CreateConsent(ctx context.Context, in *CreateConsentRequest, opts ...client.CallOption) (*CreateConsentResponse, error)
	AcceptConsent(ctx context.Context, in *AcceptConsentRequest, opts ...client.CallOption) (*AcceptConsentResponse, error)
}

type consentProviderClient struct {
	c           client.Client
	serviceName string
}

func NewConsentProviderClient(serviceName string, c client.Client) ConsentProviderClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "auth"
	}
	return &consentProviderClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *consentProviderClient) GetConsent(ctx context.Context, in *GetConsentRequest, opts ...client.CallOption) (*GetConsentResponse, error) {
	req := c.c.NewRequest(c.serviceName, "ConsentProvider.GetConsent", in)
	out := new(GetConsentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consentProviderClient) CreateConsent(ctx context.Context, in *CreateConsentRequest, opts ...client.CallOption) (*CreateConsentResponse, error) {
	req := c.c.NewRequest(c.serviceName, "ConsentProvider.CreateConsent", in)
	out := new(CreateConsentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consentProviderClient) AcceptConsent(ctx context.Context, in *AcceptConsentRequest, opts ...client.CallOption) (*AcceptConsentResponse, error) {
	req := c.c.NewRequest(c.serviceName, "ConsentProvider.AcceptConsent", in)
	out := new(AcceptConsentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ConsentProvider service

type ConsentProviderHandler interface {
	GetConsent(context.Context, *GetConsentRequest, *GetConsentResponse) error
	CreateConsent(context.Context, *CreateConsentRequest, *CreateConsentResponse) error
	AcceptConsent(context.Context, *AcceptConsentRequest, *AcceptConsentResponse) error
}

func RegisterConsentProviderHandler(s server.Server, hdlr ConsentProviderHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&ConsentProvider{hdlr}, opts...))
}

type ConsentProvider struct {
	ConsentProviderHandler
}

func (h *ConsentProvider) GetConsent(ctx context.Context, in *GetConsentRequest, out *GetConsentResponse) error {
	return h.ConsentProviderHandler.GetConsent(ctx, in, out)
}

func (h *ConsentProvider) CreateConsent(ctx context.Context, in *CreateConsentRequest, out *CreateConsentResponse) error {
	return h.ConsentProviderHandler.CreateConsent(ctx, in, out)
}

func (h *ConsentProvider) AcceptConsent(ctx context.Context, in *AcceptConsentRequest, out *AcceptConsentResponse) error {
	return h.ConsentProviderHandler.AcceptConsent(ctx, in, out)
}

// Client API for AuthCodeProvider service

type AuthCodeProviderClient interface {
	CreateAuthCode(ctx context.Context, in *CreateAuthCodeRequest, opts ...client.CallOption) (*CreateAuthCodeResponse, error)
}

type authCodeProviderClient struct {
	c           client.Client
	serviceName string
}

func NewAuthCodeProviderClient(serviceName string, c client.Client) AuthCodeProviderClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "auth"
	}
	return &authCodeProviderClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *authCodeProviderClient) CreateAuthCode(ctx context.Context, in *CreateAuthCodeRequest, opts ...client.CallOption) (*CreateAuthCodeResponse, error) {
	req := c.c.NewRequest(c.serviceName, "AuthCodeProvider.CreateAuthCode", in)
	out := new(CreateAuthCodeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthCodeProvider service

type AuthCodeProviderHandler interface {
	CreateAuthCode(context.Context, *CreateAuthCodeRequest, *CreateAuthCodeResponse) error
}

func RegisterAuthCodeProviderHandler(s server.Server, hdlr AuthCodeProviderHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&AuthCodeProvider{hdlr}, opts...))
}

type AuthCodeProvider struct {
	AuthCodeProviderHandler
}

func (h *AuthCodeProvider) CreateAuthCode(ctx context.Context, in *CreateAuthCodeRequest, out *CreateAuthCodeResponse) error {
	return h.AuthCodeProviderHandler.CreateAuthCode(ctx, in, out)
}

// Client API for AuthTokenVerifier service

type AuthTokenVerifierClient interface {
	// Verifies a token and returns claims
	Verify(ctx context.Context, in *VerifyTokenRequest, opts ...client.CallOption) (*VerifyTokenResponse, error)
}

type authTokenVerifierClient struct {
	c           client.Client
	serviceName string
}

func NewAuthTokenVerifierClient(serviceName string, c client.Client) AuthTokenVerifierClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "auth"
	}
	return &authTokenVerifierClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *authTokenVerifierClient) Verify(ctx context.Context, in *VerifyTokenRequest, opts ...client.CallOption) (*VerifyTokenResponse, error) {
	req := c.c.NewRequest(c.serviceName, "AuthTokenVerifier.Verify", in)
	out := new(VerifyTokenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthTokenVerifier service

type AuthTokenVerifierHandler interface {
	// Verifies a token and returns claims
	Verify(context.Context, *VerifyTokenRequest, *VerifyTokenResponse) error
}

func RegisterAuthTokenVerifierHandler(s server.Server, hdlr AuthTokenVerifierHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&AuthTokenVerifier{hdlr}, opts...))
}

type AuthTokenVerifier struct {
	AuthTokenVerifierHandler
}

func (h *AuthTokenVerifier) Verify(ctx context.Context, in *VerifyTokenRequest, out *VerifyTokenResponse) error {
	return h.AuthTokenVerifierHandler.Verify(ctx, in, out)
}

// Client API for AuthCodeExchanger service

type AuthCodeExchangerClient interface {
	Exchange(ctx context.Context, in *ExchangeRequest, opts ...client.CallOption) (*ExchangeResponse, error)
}

type authCodeExchangerClient struct {
	c           client.Client
	serviceName string
}

func NewAuthCodeExchangerClient(serviceName string, c client.Client) AuthCodeExchangerClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "auth"
	}
	return &authCodeExchangerClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *authCodeExchangerClient) Exchange(ctx context.Context, in *ExchangeRequest, opts ...client.CallOption) (*ExchangeResponse, error) {
	req := c.c.NewRequest(c.serviceName, "AuthCodeExchanger.Exchange", in)
	out := new(ExchangeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthCodeExchanger service

type AuthCodeExchangerHandler interface {
	Exchange(context.Context, *ExchangeRequest, *ExchangeResponse) error
}

func RegisterAuthCodeExchangerHandler(s server.Server, hdlr AuthCodeExchangerHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&AuthCodeExchanger{hdlr}, opts...))
}

type AuthCodeExchanger struct {
	AuthCodeExchangerHandler
}

func (h *AuthCodeExchanger) Exchange(ctx context.Context, in *ExchangeRequest, out *ExchangeResponse) error {
	return h.AuthCodeExchangerHandler.Exchange(ctx, in, out)
}

// Client API for AuthTokenRefresher service

type AuthTokenRefresherClient interface {
	Refresh(ctx context.Context, in *RefreshTokenRequest, opts ...client.CallOption) (*RefreshTokenResponse, error)
}

type authTokenRefresherClient struct {
	c           client.Client
	serviceName string
}

func NewAuthTokenRefresherClient(serviceName string, c client.Client) AuthTokenRefresherClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "auth"
	}
	return &authTokenRefresherClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *authTokenRefresherClient) Refresh(ctx context.Context, in *RefreshTokenRequest, opts ...client.CallOption) (*RefreshTokenResponse, error) {
	req := c.c.NewRequest(c.serviceName, "AuthTokenRefresher.Refresh", in)
	out := new(RefreshTokenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthTokenRefresher service

type AuthTokenRefresherHandler interface {
	Refresh(context.Context, *RefreshTokenRequest, *RefreshTokenResponse) error
}

func RegisterAuthTokenRefresherHandler(s server.Server, hdlr AuthTokenRefresherHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&AuthTokenRefresher{hdlr}, opts...))
}

type AuthTokenRefresher struct {
	AuthTokenRefresherHandler
}

func (h *AuthTokenRefresher) Refresh(ctx context.Context, in *RefreshTokenRequest, out *RefreshTokenResponse) error {
	return h.AuthTokenRefresherHandler.Refresh(ctx, in, out)
}
