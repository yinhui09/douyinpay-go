package verifiers

import (
	"context"
	"strings"
	"testing"

	"github.com/douyinpay/douyinpay-go/tools/auth"
)

func TestNilVerifierImplementsVerifier(t *testing.T) {
	var verifier auth.Verifier = &NilVerifier{}

	if err := verifier.Verify(context.Background(), "serial", "message", "signature"); err != nil {
		t.Fatalf("expected nil verifier to skip verification, got %v", err)
	}

	_, err := verifier.GetSerial(context.Background())
	if err == nil {
		t.Fatal("expected nil verifier serial error")
	}
	if !strings.Contains(err.Error(), "NilVerifier has no serial") {
		t.Fatalf("unexpected error: %v", err)
	}
}
