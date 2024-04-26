package p2p

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTCPTransport(t *testing.T) {
	// Test if NewTCPTransport returns a pointer to a TCPTransport
	listenAddr := ":4000"
	transport := NewTCPTransport(listenAddr)
	assert.Equal(t, transport.listenAddress, listenAddr)

	// Server

}
