package brp

import (
	"bufio"
	"bytes"
	"errors"
	"io"
)

// Ending is used to separate packets in the stream.
var Ending = []byte{0x00, 0xFF, 0xCC}

type Packet struct {
	Tag     TAG
	Payload []byte
}

func (p Packet) String() string {
	str := "[" + string(p.Tag)
	if len(p.Payload) > 0 {
		str += ": " + string(p.Payload)
	}
	return str + "]"
}

// ReadPacket reads a packet from the given reader. Returns the tag and packet data separately.
func ReadPacket(reader io.Reader) (packet Packet, err error) {
	var data []byte
	r := bufio.NewReader(reader)

	for {
		s := ""
		s, err = r.ReadString(Ending[len(Ending)-1])
		if err != nil {
			return packet, err
		}

		data = append(data, []byte(s)...)
		if bytes.HasSuffix(data, Ending) {
			data = data[:len(data)-len(Ending)]
			break
		}
	}

	tagBytes := bytes.ToUpper(bytes.TrimSpace(bytes.Split(data, []byte(" "))[0]))
	packet.Payload = bytes.TrimSpace(bytes.TrimPrefix(data, tagBytes))
	packet.Tag = TAG(tagBytes)

	if _, ok := tags[packet.Tag]; !ok {
		return packet, errors.New("unknown tag: " + string(packet.Tag))
	}

	return packet, nil
}
