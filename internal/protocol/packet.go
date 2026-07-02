package protocol

type PacketType uint8

const (
	PacketTypeHandshake PacketType = iota + 1
	PacketTypeData
	PacketTypePing
)

type Packet struct {
	Version  uint8
	Type     PacketType
	PacketID uint32
	Length   uint16
	Payload  []byte
}
