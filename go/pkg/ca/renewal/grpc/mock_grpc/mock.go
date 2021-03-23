// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/scionproto/scion/go/pkg/ca/renewal/grpc (interfaces: CAServiceClient,ChainBuilder,CMSSigner,RenewalRequestVerifier,Signer)

// Package mock_grpc is a generated GoMock package.
package mock_grpc

import (
	context "context"
	x509 "crypto/x509"
	gomock "github.com/golang/mock/gomock"
	api "github.com/scionproto/scion/go/pkg/ca/api"
	crypto "github.com/scionproto/scion/go/pkg/proto/crypto"
	http "net/http"
	reflect "reflect"
)

// MockCAServiceClient is a mock of CAServiceClient interface
type MockCAServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockCAServiceClientMockRecorder
}

// MockCAServiceClientMockRecorder is the mock recorder for MockCAServiceClient
type MockCAServiceClientMockRecorder struct {
	mock *MockCAServiceClient
}

// NewMockCAServiceClient creates a new mock instance
func NewMockCAServiceClient(ctrl *gomock.Controller) *MockCAServiceClient {
	mock := &MockCAServiceClient{ctrl: ctrl}
	mock.recorder = &MockCAServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCAServiceClient) EXPECT() *MockCAServiceClientMockRecorder {
	return m.recorder
}

// PostCertificateRenewal mocks base method
func (m *MockCAServiceClient) PostCertificateRenewal(arg0 context.Context, arg1 int, arg2 api.AS, arg3 api.PostCertificateRenewalJSONRequestBody, arg4 ...api.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PostCertificateRenewal", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostCertificateRenewal indicates an expected call of PostCertificateRenewal
func (mr *MockCAServiceClientMockRecorder) PostCertificateRenewal(arg0, arg1, arg2, arg3 interface{}, arg4 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostCertificateRenewal", reflect.TypeOf((*MockCAServiceClient)(nil).PostCertificateRenewal), varargs...)
}

// MockChainBuilder is a mock of ChainBuilder interface
type MockChainBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockChainBuilderMockRecorder
}

// MockChainBuilderMockRecorder is the mock recorder for MockChainBuilder
type MockChainBuilderMockRecorder struct {
	mock *MockChainBuilder
}

// NewMockChainBuilder creates a new mock instance
func NewMockChainBuilder(ctrl *gomock.Controller) *MockChainBuilder {
	mock := &MockChainBuilder{ctrl: ctrl}
	mock.recorder = &MockChainBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChainBuilder) EXPECT() *MockChainBuilderMockRecorder {
	return m.recorder
}

// CreateChain mocks base method
func (m *MockChainBuilder) CreateChain(arg0 context.Context, arg1 *x509.CertificateRequest) ([]*x509.Certificate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateChain", arg0, arg1)
	ret0, _ := ret[0].([]*x509.Certificate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateChain indicates an expected call of CreateChain
func (mr *MockChainBuilderMockRecorder) CreateChain(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChain", reflect.TypeOf((*MockChainBuilder)(nil).CreateChain), arg0, arg1)
}

// MockCMSSigner is a mock of CMSSigner interface
type MockCMSSigner struct {
	ctrl     *gomock.Controller
	recorder *MockCMSSignerMockRecorder
}

// MockCMSSignerMockRecorder is the mock recorder for MockCMSSigner
type MockCMSSignerMockRecorder struct {
	mock *MockCMSSigner
}

// NewMockCMSSigner creates a new mock instance
func NewMockCMSSigner(ctrl *gomock.Controller) *MockCMSSigner {
	mock := &MockCMSSigner{ctrl: ctrl}
	mock.recorder = &MockCMSSignerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCMSSigner) EXPECT() *MockCMSSignerMockRecorder {
	return m.recorder
}

// SignCMS mocks base method
func (m *MockCMSSigner) SignCMS(arg0 context.Context, arg1 []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignCMS", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignCMS indicates an expected call of SignCMS
func (mr *MockCMSSignerMockRecorder) SignCMS(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignCMS", reflect.TypeOf((*MockCMSSigner)(nil).SignCMS), arg0, arg1)
}

// MockRenewalRequestVerifier is a mock of RenewalRequestVerifier interface
type MockRenewalRequestVerifier struct {
	ctrl     *gomock.Controller
	recorder *MockRenewalRequestVerifierMockRecorder
}

// MockRenewalRequestVerifierMockRecorder is the mock recorder for MockRenewalRequestVerifier
type MockRenewalRequestVerifierMockRecorder struct {
	mock *MockRenewalRequestVerifier
}

// NewMockRenewalRequestVerifier creates a new mock instance
func NewMockRenewalRequestVerifier(ctrl *gomock.Controller) *MockRenewalRequestVerifier {
	mock := &MockRenewalRequestVerifier{ctrl: ctrl}
	mock.recorder = &MockRenewalRequestVerifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRenewalRequestVerifier) EXPECT() *MockRenewalRequestVerifierMockRecorder {
	return m.recorder
}

// VerifyCMSSignedRenewalRequest mocks base method
func (m *MockRenewalRequestVerifier) VerifyCMSSignedRenewalRequest(arg0 context.Context, arg1 []byte) (*x509.CertificateRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyCMSSignedRenewalRequest", arg0, arg1)
	ret0, _ := ret[0].(*x509.CertificateRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyCMSSignedRenewalRequest indicates an expected call of VerifyCMSSignedRenewalRequest
func (mr *MockRenewalRequestVerifierMockRecorder) VerifyCMSSignedRenewalRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyCMSSignedRenewalRequest", reflect.TypeOf((*MockRenewalRequestVerifier)(nil).VerifyCMSSignedRenewalRequest), arg0, arg1)
}

// VerifyPbSignedRenewalRequest mocks base method
func (m *MockRenewalRequestVerifier) VerifyPbSignedRenewalRequest(arg0 context.Context, arg1 *crypto.SignedMessage, arg2 [][]*x509.Certificate) (*x509.CertificateRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyPbSignedRenewalRequest", arg0, arg1, arg2)
	ret0, _ := ret[0].(*x509.CertificateRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyPbSignedRenewalRequest indicates an expected call of VerifyPbSignedRenewalRequest
func (mr *MockRenewalRequestVerifierMockRecorder) VerifyPbSignedRenewalRequest(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyPbSignedRenewalRequest", reflect.TypeOf((*MockRenewalRequestVerifier)(nil).VerifyPbSignedRenewalRequest), arg0, arg1, arg2)
}

// MockSigner is a mock of Signer interface
type MockSigner struct {
	ctrl     *gomock.Controller
	recorder *MockSignerMockRecorder
}

// MockSignerMockRecorder is the mock recorder for MockSigner
type MockSignerMockRecorder struct {
	mock *MockSigner
}

// NewMockSigner creates a new mock instance
func NewMockSigner(ctrl *gomock.Controller) *MockSigner {
	mock := &MockSigner{ctrl: ctrl}
	mock.recorder = &MockSignerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSigner) EXPECT() *MockSignerMockRecorder {
	return m.recorder
}

// Sign mocks base method
func (m *MockSigner) Sign(arg0 context.Context, arg1 []byte, arg2 ...[]byte) (*crypto.SignedMessage, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Sign", varargs...)
	ret0, _ := ret[0].(*crypto.SignedMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sign indicates an expected call of Sign
func (mr *MockSignerMockRecorder) Sign(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockSigner)(nil).Sign), varargs...)
}
