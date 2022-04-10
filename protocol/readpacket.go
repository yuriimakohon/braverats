package protocol

import (
	"bufio"
	"bytes"
	"errors"
	"io"
)

func ReadPacket(reader io.Reader) (tag TAG, args []byte, err error) {
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
	args = bytes.TrimSpace(bytes.TrimPrefix(packet, tagBytes))
	tag = TAG(tagBytes)

	if _, ok := TAGs[tag]; !ok {
		return "", nil, errors.New("unknown tag: " + string(tag))
	}

	return tag, args, nil
}
