// Copyright 2020 Anapaya Systems
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package grpc_test

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/scionproto/scion/go/lib/addr"
	"github.com/scionproto/scion/go/lib/metrics"
	"github.com/scionproto/scion/go/lib/scrypto/cppki"
	"github.com/scionproto/scion/go/lib/scrypto/signed"
	"github.com/scionproto/scion/go/lib/serrors"
	"github.com/scionproto/scion/go/lib/xtest"
	"github.com/scionproto/scion/go/pkg/ca/renewal"
	"github.com/scionproto/scion/go/pkg/ca/renewal/grpc"
	renewalgrpc "github.com/scionproto/scion/go/pkg/ca/renewal/grpc"
	"github.com/scionproto/scion/go/pkg/ca/renewal/grpc/mock_grpc"
	"github.com/scionproto/scion/go/pkg/ca/renewal/mock_renewal"
	cppb "github.com/scionproto/scion/go/pkg/proto/control_plane"
	"github.com/scionproto/scion/go/pkg/trust"
)

var (
	mockErr = serrors.New("send error")
	mockCSR = &x509.CertificateRequest{
		Raw: []byte("mock CSR"),
		Subject: pkix.Name{Names: []pkix.AttributeTypeAndValue{{
			Type:  cppki.OIDNameIA,
			Value: "1-ff00:0:111",
		}}},
	}
	mockChain = []*x509.Certificate{
		{
			Raw:          []byte("mock AS cert"),
			SubjectKeyId: []byte("mock cert subject key"),
		},
		{Raw: []byte("mock CA cert")},
	}
	mockIssuedChain = []*x509.Certificate{
		{Raw: []byte("mock issued AS cert")},
		{Raw: []byte("mock CA cert")},
	}
)

func TestCMSHandleCMSRequest(t *testing.T) {
	clientKey, chain := genChainCMS(t)
	signedReq, err := renewal.NewChainRenewalRequest(context.Background(), mockCSR.Raw,
		trust.Signer{
			PrivateKey: clientKey,
			Algorithm:  signed.ECDSAWithSHA256,
			ChainValidity: cppki.Validity{
				NotBefore: time.Now(),
				NotAfter:  time.Now().Add(time.Hour),
			},
			Expiration:   time.Now().Add(time.Hour - time.Minute),
			IA:           xtest.MustParseIA("1-ff00:0:111"),
			SubjectKeyID: chain[0].SubjectKeyId,
			Chain:        chain,
		},
	)
	require.NoError(t, err)

	tests := map[string]struct {
		Request      func(t *testing.T) *cppb.ChainRenewalRequest
		Verifier     func(ctrl *gomock.Controller) renewalgrpc.RenewalRequestVerifier
		ChainBuilder func(ctrl *gomock.Controller) renewalgrpc.ChainBuilder
		CMSSigner    func(ctrl *gomock.Controller) renewalgrpc.CMSSigner
		DB           func(ctrl *gomock.Controller) renewal.DB
		IA           addr.IA
		Metric       string
		Assertion    assert.ErrorAssertionFunc
		Code         codes.Code
	}{
		"dummy request": {
			Request: func(t *testing.T) *cppb.ChainRenewalRequest {
				return &cppb.ChainRenewalRequest{
					CmsSignedRequest: []byte("dummy request"),
				}
			},
			Verifier: func(ctrl *gomock.Controller) renewalgrpc.RenewalRequestVerifier {
				return mock_grpc.NewMockRenewalRequestVerifier(ctrl)
			},
			ChainBuilder: func(ctrl *gomock.Controller) renewalgrpc.ChainBuilder {
				return mock_grpc.NewMockChainBuilder(ctrl)
			},
			CMSSigner: func(ctrl *gomock.Controller) renewalgrpc.CMSSigner {
				return mock_grpc.NewMockCMSSigner(ctrl)
			},
			DB: func(ctrl *gomock.Controller) renewal.DB {
				return mock_renewal.NewMockDB(ctrl)
			},
			IA:        xtest.MustParseIA("1-ff00:0:110"),
			Assertion: assert.Error,
			Code:      codes.InvalidArgument,
			Metric:    "err_parse",
		},
		"not client": {
			Request: func(t *testing.T) *cppb.ChainRenewalRequest {
				return signedReq
			},
			Verifier: func(ctrl *gomock.Controller) renewalgrpc.RenewalRequestVerifier {
				return mock_grpc.NewMockRenewalRequestVerifier(ctrl)
			},
			ChainBuilder: func(ctrl *gomock.Controller) renewalgrpc.ChainBuilder {
				return mock_grpc.NewMockChainBuilder(ctrl)
			},
			CMSSigner: func(ctrl *gomock.Controller) renewalgrpc.CMSSigner {
				return mock_grpc.NewMockCMSSigner(ctrl)
			},
			DB: func(ctrl *gomock.Controller) renewal.DB {
				return mock_renewal.NewMockDB(ctrl)
			},
			IA:        xtest.MustParseIA("1-ff00:0:112"),
			Assertion: assert.Error,
			Code:      codes.PermissionDenied,
			Metric:    "err_notfound",
		},
		"invalid signature": {
			Request: func(t *testing.T) *cppb.ChainRenewalRequest {
				return signedReq
			},
			Verifier: func(ctrl *gomock.Controller) renewalgrpc.RenewalRequestVerifier {
				v := mock_grpc.NewMockRenewalRequestVerifier(ctrl)
				v.EXPECT().VerifyCMSSignedRenewalRequest(
					context.Background(),
					signedReq.CmsSignedRequest,
				).Return(nil, mockErr)
				return v
			},
			ChainBuilder: func(ctrl *gomock.Controller) renewalgrpc.ChainBuilder {
				return mock_grpc.NewMockChainBuilder(ctrl)
			},
			CMSSigner: func(ctrl *gomock.Controller) renewalgrpc.CMSSigner {
				return mock_grpc.NewMockCMSSigner(ctrl)
			},
			DB: func(ctrl *gomock.Controller) renewal.DB {
				return mock_renewal.NewMockDB(ctrl)
			},
			IA:        xtest.MustParseIA("1-ff00:0:110"),
			Assertion: assert.Error,
			Code:      codes.InvalidArgument,
			Metric:    "err_verify",
		},
		"failed to build chain": {
			Request: func(t *testing.T) *cppb.ChainRenewalRequest {
				return signedReq
			},
			Verifier: func(ctrl *gomock.Controller) renewalgrpc.RenewalRequestVerifier {
				v := mock_grpc.NewMockRenewalRequestVerifier(ctrl)
				v.EXPECT().VerifyCMSSignedRenewalRequest(
					context.Background(),
					signedReq.CmsSignedRequest,
				).Return(mockCSR, nil)
				return v
			},
			ChainBuilder: func(ctrl *gomock.Controller) renewalgrpc.ChainBuilder {
				cb := mock_grpc.NewMockChainBuilder(ctrl)
				cb.EXPECT().CreateChain(gomock.Any(), gomock.Any()).Return(nil, mockErr)
				return cb
			},
			CMSSigner: func(ctrl *gomock.Controller) renewalgrpc.CMSSigner {
				return mock_grpc.NewMockCMSSigner(ctrl)
			},
			DB: func(ctrl *gomock.Controller) renewal.DB {
				return mock_renewal.NewMockDB(ctrl)
			},
			IA:        xtest.MustParseIA("1-ff00:0:110"),
			Assertion: assert.Error,
			Code:      codes.Unavailable,
			Metric:    "err_internal",
		},
		"db write error": {
			Request: func(t *testing.T) *cppb.ChainRenewalRequest {
				return signedReq
			},
			Verifier: func(ctrl *gomock.Controller) renewalgrpc.RenewalRequestVerifier {
				v := mock_grpc.NewMockRenewalRequestVerifier(ctrl)
				v.EXPECT().VerifyCMSSignedRenewalRequest(
					context.Background(),
					signedReq.CmsSignedRequest,
				).Return(mockCSR, nil)
				return v
			},
			ChainBuilder: func(ctrl *gomock.Controller) renewalgrpc.ChainBuilder {
				cb := mock_grpc.NewMockChainBuilder(ctrl)
				cb.EXPECT().CreateChain(gomock.Any(), gomock.Any()).Return(mockIssuedChain, nil)
				return cb
			},
			CMSSigner: func(ctrl *gomock.Controller) renewalgrpc.CMSSigner {
				return mock_grpc.NewMockCMSSigner(ctrl)
			},
			DB: func(ctrl *gomock.Controller) renewal.DB {
				db := mock_renewal.NewMockDB(ctrl)
				db.EXPECT().InsertClientChain(gomock.Any(), mockIssuedChain).Return(false, mockErr)
				return db
			},
			IA:        xtest.MustParseIA("1-ff00:0:110"),
			Assertion: assert.Error,
			Code:      codes.Unavailable,
			Metric:    "err_database",
		},
		"valid": {
			Request: func(t *testing.T) *cppb.ChainRenewalRequest {
				return signedReq
			},
			Verifier: func(ctrl *gomock.Controller) renewalgrpc.RenewalRequestVerifier {
				v := mock_grpc.NewMockRenewalRequestVerifier(ctrl)
				v.EXPECT().VerifyCMSSignedRenewalRequest(context.Background(),
					signedReq.CmsSignedRequest).Return(mockCSR, nil)
				return v
			},
			ChainBuilder: func(ctrl *gomock.Controller) renewalgrpc.ChainBuilder {
				cb := mock_grpc.NewMockChainBuilder(ctrl)
				cb.EXPECT().CreateChain(gomock.Any(), gomock.Any()).Return(mockIssuedChain, nil)
				return cb
			},
			CMSSigner: func(ctrl *gomock.Controller) renewalgrpc.CMSSigner {
				signer := mock_grpc.NewMockCMSSigner(ctrl)
				signer.EXPECT().SignCMS(gomock.Any(), gomock.Any())
				return signer
			},
			DB: func(ctrl *gomock.Controller) renewal.DB {
				db := mock_renewal.NewMockDB(ctrl)
				db.EXPECT().InsertClientChain(gomock.Any(), mockIssuedChain)
				return db
			},
			IA:        xtest.MustParseIA("1-ff00:0:110"),
			Assertion: assert.NoError,
			Code:      codes.OK,
			Metric:    "ok_success",
		},
	}
	for name, tc := range tests {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			ctr := metrics.NewTestCounter()
			s := &renewalgrpc.CMS{
				Verifier:     tc.Verifier(ctrl),
				ChainBuilder: tc.ChainBuilder(ctrl),
				DB:           tc.DB(ctrl),
				IA:           tc.IA,
				Metrics: grpc.CMSHandlerMetrics{
					DatabaseError: ctr.With("result", "err_database"),
					InternalError: ctr.With("result", "err_internal"),
					NotFoundError: ctr.With("result", "err_notfound"),
					ParseError:    ctr.With("result", "err_parse"),
					VerifyError:   ctr.With("result", "err_verify"),
					Success:       ctr.With("result", "ok_success"),
				},
			}
			_, err := s.HandleCMSRequest(context.Background(), tc.Request(t))
			tc.Assertion(t, err)
			assert.Equal(t, tc.Code, status.Code(err))
			for _, res := range []string{
				"err_database",
				"err_internal",
				"err_unavailable",
				"err_notfound",
				"err_parse",
				"err_verify",
				"ok_success",
			} {
				expected := float64(0)
				if res == tc.Metric {
					expected = 1
				}
				assert.Equal(t, expected, metrics.CounterValue(ctr.With("result", res)), res)
			}
		})
	}
}

func genChainCMS(t *testing.T) (*ecdsa.PrivateKey, []*x509.Certificate) {
	t.Helper()
	// Generate serial numbers for certs.
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serial1, err := rand.Int(rand.Reader, serialNumberLimit)
	require.NoError(t, err)
	serial2, err := rand.Int(rand.Reader, serialNumberLimit)
	require.NoError(t, err)

	// Generate CA key and cert (to the extent needed)
	caKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)
	caCertTmpl := x509.Certificate{
		Subject: pkix.Name{ExtraNames: []pkix.AttributeTypeAndValue{{
			Type:  cppki.OIDNameIA,
			Value: "1-ff00:0:110",
		}}},
		SerialNumber: serial1,
	}
	caCertRaw, err := x509.CreateCertificate(rand.Reader, &caCertTmpl, &caCertTmpl,
		&caKey.PublicKey, caKey)
	require.NoError(t, err)
	caCert, err := x509.ParseCertificate(caCertRaw)
	require.NoError(t, err)

	// Generate client key and cert (to the extent needed)
	clientKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)
	clientCertTmpl := x509.Certificate{
		Subject: pkix.Name{ExtraNames: []pkix.AttributeTypeAndValue{{
			Type:  cppki.OIDNameIA,
			Value: "1-ff00:0:111",
		}}},
		SerialNumber: serial2,
	}
	clientCertRaw, err := x509.CreateCertificate(rand.Reader, &clientCertTmpl, caCert,
		&clientKey.PublicKey, caKey)
	require.NoError(t, err)
	clientCert, err := x509.ParseCertificate(clientCertRaw)
	require.NoError(t, err)

	return clientKey, []*x509.Certificate{clientCert, caCert}
}
