package protocol

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

func (p *Packet) SerializePacket() ([]byte, error) {

	if err := p.HeaderValidate(); err != nil {
		return nil, err
	}

	p.Length = uint16(len(p.Payload))

	if err := p.Validate(); err != nil {
		return nil, err
	}

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

	if err := deserializedPacket.HeaderValidate(); err != nil {
		return nil, err
	}

	if err := binary.Read(readBuff, binary.BigEndian, &deserializedPacket.Length); err != nil {
		return nil, err
	}
	if deserializedPacket.Length > MTU {
		return nil, fmt.Errorf("packet exceeded the MTU")
	}

	deserializedPacket.Payload = make([]byte, int(deserializedPacket.Length))

	_, err := io.ReadFull(readBuff, deserializedPacket.Payload)
	if err != nil {
		return nil, err
	}

	return &deserializedPacket, nil

}

func (p *Packet) HeaderValidate() error {
	if p.Version > 2 {
		return fmt.Errorf("not supported version")
	}

	if !p.Type.isValidPacket() {
		return fmt.Errorf("not a valid packet type")
	}
	return nil
}

func (p *Packet) Validate() error {
	if p.Length > MTU {
		return fmt.Errorf("packet exceeded the MTU")
	}
	return nil
}

func (t PacketType) isValidPacket() bool {
	return t == PacketTypeHandshake || t == PacketTypePing || t == PacketTypeData
}
