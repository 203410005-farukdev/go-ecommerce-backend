package service

import "testing"

func TestNewHomeService(t *testing.T) {
	svc := NewHomeService(nil, nil, nil)
	if svc == nil {
		t.Fatal("expected home service instance")
	}
}
