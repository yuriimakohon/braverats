package brp

import (
	"bytes"
	"strconv"
)

func NewEventJoinedLobby(name string) []byte {
	resp := append([]byte(EventJoinedLobby), ' ')
	resp = append(resp, []byte(name)...)
	return append(resp, Ending...)
}

func NewEventLeftLobby(name string) []byte {
	resp := append([]byte(EventLeftLobby), ' ')
	resp = append(resp, []byte(name)...)
	return append(resp, Ending...)
}

func NewEventLobbyClosed() []byte {
	return append([]byte(EventLobbyClosed), Ending...)
}

func NewEventPlayerReadiness(ready bool) []byte {
	resp := append([]byte(EventPlayerReadiness), ' ')
	resp = append(resp, []byte(strconv.FormatBool(ready))...)
	return append(resp, Ending...)
}

func NewEventMatchStarted() []byte {
	return append([]byte(EventMatchStarted), Ending...)
}

func NewEventCardPut(faceUp bool, card CardID) []byte {
	resp := append([]byte(EventCardPut), ' ')
	resp = append(resp, []byte(strconv.FormatBool(faceUp))...)
	if faceUp {
		resp = append(resp, ' ')
		resp = append(resp, []byte(strconv.Itoa(int(card)))...)
	}
	return append(resp, Ending...)
}

func ParseEventCardPut(packet Packet) (faceUp bool, card CardID, err error) {
	if packet.Tag != EventCardPut {
		return faceUp, card, ErrPacketTagMismatch
	}

	args := bytes.Split(packet.Payload, []byte(" "))

	faceUp, err = strconv.ParseBool(string(args[0]))
	if err != nil {
		return faceUp, card, err
	}

	if faceUp {
		id, err := strconv.Atoi(string(args[1]))
		if err != nil {
			return faceUp, card, err
		}
		card = CardID(id)
	}

	return faceUp, card, nil
}

func NewEventRoundEnded(result RoundResult, card CardID) []byte {
	resp := append([]byte(EventRoundEnded), ' ')
	resp = append(resp, []byte(strconv.Itoa(result.Int()))...)
	resp = append(resp, ' ')
	resp = append(resp, []byte(strconv.Itoa(int(card)))...)
	return append(resp, Ending...)
}

func ParseEventRoundEnded(packet Packet) (result RoundResult, card CardID, err error) {
	if packet.Tag != EventRoundEnded {
		return result, card, ErrPacketTagMismatch
	}

	args := bytes.Split(packet.Payload, []byte(" "))

	resultInt, err := strconv.Atoi(string(args[0]))
	if err != nil {
		return result, card, err
	}
	result = RoundResult(resultInt)

	cardInt, err := strconv.Atoi(string(args[1]))
	if err != nil {
		return result, card, err
	}
	card = CardID(cardInt)

	return result, card, nil
}
