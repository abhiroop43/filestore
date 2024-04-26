package p2p

// Peer is a remote node
type Peer interface {
}

// Transport is to handle communication
// between nodes, this can be TCP, UDP etc.
type Transport interface {
	ListenAndAccept() error
}
