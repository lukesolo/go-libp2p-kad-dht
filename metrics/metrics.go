package metrics

import (
	pb "github.com/lukesolo/go-libp2p-kad-dht/pb"
	"go.opencensus.io/tag"
)

// Keys
var (
	KeyMessageType, _ = tag.NewKey("message_type")
	KeyPeerID, _      = tag.NewKey("peer_id")
	// KeyInstanceID identifies a dht instance by the pointer address.
	// Useful for differentiating between different dhts that have the same peer id.
	KeyInstanceID, _ = tag.NewKey("instance_id")
)

// UpsertMessageType is a convenience upserts the message type
// of a pb.Message into the KeyMessageType.
func UpsertMessageType(m *pb.Message) tag.Mutator {
	return tag.Upsert(KeyMessageType, m.Type.String())
}
