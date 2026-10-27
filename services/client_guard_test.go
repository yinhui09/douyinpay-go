package services

import (
	"context"
	"strings"
	"testing"
)

func TestNewClientRSAReturnsErrorForInvalidPrivateKey(t *testing.T) {
	cli, err := NewClientRSA(context.Background(), "mch", "serial", "invalid-private-key", "invalid-cert")
	if err == nil {
		t.Fatal("expected invalid private key error")
	}
	if cli != nil {
		t.Fatal("expected nil client on invalid private key")
	}
	if !strings.Contains(err.Error(), "init rsa client") {
		t.Fatalf("expected contextual error, got %v", err)
	}
}

func TestNewAutoClientRSAReturnsErrorForInvalidPrivateKey(t *testing.T) {
	cli, err := NewAutoClientRSA(context.Background(), "mch", "serial", "invalid-private-key", "encrypt-key")
	if err == nil {
		t.Fatal("expected invalid private key error")
	}
	if cli != nil {
		t.Fatal("expected nil client on invalid private key")
	}
	if !strings.Contains(err.Error(), "init auto rsa client") {
		t.Fatalf("expected contextual error, got %v", err)
	}
}
