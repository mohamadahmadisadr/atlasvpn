package protocol

import (
	"bytes"
	"encoding/binary"
	"io"
)

func (p *Packet) SerializePacket() ([]byte, error) {

	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.BigEndian, p.Version); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.BigEndian, p.Type); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.BigEndian, p.PacketID); err != nil {
		return nil, err
	}
	p.Length = uint16(len(p.Payload))

	if err := binary.Write(&buf, binary.BigEndian, p.Length); err != nil {
		return nil, err
	}
	if err := binary.Write(&buf, binary.BigEndian, p.Payload); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func DeserializePacket(data []byte) (*Packet, error) {

	var deserializedPacket Packet
	readBuff := bytes.NewReader(data)
	if err := binary.Read(readBuff, binary.BigEndian, &deserializedPacket.Version); err != nil {
		return nil, err
	}
	if err := binary.Read(readBuff, binary.BigEndian, &deserializedPacket.Type); err != nil {
		return nil, err
	}
	if err := binary.Read(readBuff, binary.BigEndian, &deserializedPacket.PacketID); err != nil {
		return nil, err
	}
	if err := binary.Read(readBuff, binary.BigEndian, &deserializedPacket.Length); err != nil {
		return nil, err
	}
	deserializedPacket.Payload = make([]byte, int(deserializedPacket.Length))
	_, err := io.ReadFull(readBuff, deserializedPacket.Payload)
	if err != nil {
		return nil, err
	}
	return &deserializedPacket, nil

}
