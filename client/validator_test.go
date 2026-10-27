package client

import (
	"net/http"
	"strings"
	"testing"

	"github.com/douyinpay/douyinpay-go/tools/consts"
)

func TestGetHeaderRejectsMissingTimestamp(t *testing.T) {
	header := http.Header{
		consts.DouyinPaySerial:    []string{"serial"},
		consts.DouyinPaySignature: []string{"signature"},
		consts.DouyinPayNonce:     []string{"nonce"},
	}

	_, err := GetHeader(header)
	if err == nil {
		t.Fatal("expected missing timestamp error")
	}
	if !strings.Contains(err.Error(), "key `Douyinpay-Timestamp` is empty") {
		t.Fatalf("expected missing timestamp error, got %v", err)
	}
}
