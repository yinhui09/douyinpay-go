package client

import (
	"context"
	"testing"

	"github.com/douyinpay/douyinpay-go/tools/auth"
)

type guardTestSigner struct{}

func (guardTestSigner) Sign(context.Context, string) (*auth.SignatureResult, error) {
	return &auth.SignatureResult{}, nil
}

func (guardTestSigner) Algorithm() string {
	return "RSA"
}

func TestNewClientWithDialSettingsRejectsNilVerifier(t *testing.T) {
	_, err := NewClientWithDialSettings(context.Background(), &DialSettings{
		Signer: guardTestSigner{},
	})
	if err == nil {
		t.Fatal("expected error when verifier is nil")
	}
}
