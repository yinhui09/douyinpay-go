package downloader

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/tjfoc/gmsm/x509"
)

func TestCertificateMapGetReturnsCopy(t *testing.T) {
	serial := "01"
	originalNotBefore := time.Unix(100, 0)
	m := NewCertificateMap(map[string]*x509.Certificate{
		serial: {
			SerialNumber: big.NewInt(1),
			NotBefore:    originalNotBefore,
		},
	})

	cert, ok := m.Get(context.Background(), serial)
	if !ok {
		t.Fatal("expected certificate")
	}
	cert.NotBefore = time.Unix(200, 0)

	cached, ok := m.Get(context.Background(), serial)
	if !ok {
		t.Fatal("expected certificate")
	}
	if !cached.NotBefore.Equal(originalNotBefore) {
		t.Fatalf("cached certificate was mutated, got %v want %v", cached.NotBefore, originalNotBefore)
	}
}
