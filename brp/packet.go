package brp

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strings"
)

// Ending is used to separate packets in the stream.
var Ending = []byte{'\r', '\n'}

const (
	MaxPlayerNameLen = 24
	MaxLobbyNameLen  = 24
)

type Packet struct {
	Tag     TAG
	Type    Type
	Payload []byte
}

type Type string

const (
	TypeReq   Type = "REQ"
	TypeResp  Type = "RESP"
	TypeEvent Type = "EVENT"
)

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

	return ParsePacket(data)
}

func ParsePacket(data []byte) (packet Packet, err error) {
	tagBytes := bytes.ToUpper(bytes.TrimSpace(bytes.Split(data, []byte(" "))[0]))
	packet.Payload = bytes.TrimSpace(bytes.TrimPrefix(data, tagBytes))
	packet.Tag = TAG(tagBytes)

	if strings.HasPrefix(string(tagBytes), string(TypeReq)) {
		packet.Type = TypeReq
	} else if strings.HasPrefix(string(tagBytes), string(TypeResp)) {
		packet.Type = TypeResp
	} else {
		packet.Type = TypeEvent
	}

	if _, ok := tags[packet.Tag]; !ok {
		return packet, errors.New("unknown tag: " + string(packet.Tag))
	}

	return packet, nil
}

// dropCR drops a terminal \r from the data.
func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}

func ScanCRLF(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, Ending); i >= 0 {
		// We have a full newline-terminated line.
		return i + 2, dropCR(data[0:i]), nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), dropCR(data), nil
	}
	// Request more data.
	return 0, nil, nil
}
