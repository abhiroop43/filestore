package p2p

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTCPTransport(t *testing.T) {
	listenAddr := ":4000"
	transport := NewTCPTransport(listenAddr)
	assert.Equal(t, transport.listenAddress, listenAddr)
	assert.Nil(t, transport.ListenAndAccept())

	select {}
}
