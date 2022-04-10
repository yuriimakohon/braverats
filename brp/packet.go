package brp

import (
	"bufio"
	"bytes"
	"errors"
	"io"
)

// Ending is used to separate packets in the stream.
var Ending = []byte{0x00, 0xFF, 0xCC}

// ReadPacket reads a packet from the given reader. Returns the tag and packet data separately.
func ReadPacket(reader io.Reader) (tag TAG, data []byte, err error) {
	var packet []byte
	r := bufio.NewReader(reader)

	for {
		s := ""
		s, err = r.ReadString(Ending[len(Ending)-1])
		if err != nil {
			return
		}

		packet = append(packet, []byte(s)...)
		if bytes.HasSuffix(packet, Ending) {
			packet = packet[:len(packet)-len(Ending)]
			break
		}
	}

	tagBytes := bytes.ToUpper(bytes.TrimSpace(bytes.Split(packet, []byte(" "))[0]))
	data = bytes.TrimSpace(bytes.TrimPrefix(packet, tagBytes))
	tag = TAG(tagBytes)

	if _, ok := tags[tag]; !ok {
		return "", nil, errors.New("unknown tag: " + string(tag))
	}

	return tag, data, nil
}
