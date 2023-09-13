package bridge

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func DecodeMsg(msg string) (*Msg, error) {
	unpackedMsg, err := unpacked(msgJson(), msg)
	if err != nil {
		return nil, err
	}

	msgLen := len(unpackedMsg)
	if msgLen < 5 {
		return nil, fmt.Errorf("msg len : %d", msgLen)
	}

	var decodeMsg Msg
	decodeMsg.FromChainId = unpackedMsg[0].(string)
	decodeMsg.FromTxHash = unpackedMsg[1].(string)
	decodeMsg.ToTokenAddr = unpackedMsg[2].(string)
	decodeMsg.ToUserAddr = unpackedMsg[3].(string)
	decodeMsg.Amount = unpackedMsg[4].(*big.Int)

	return &decodeMsg, nil
}

// multisig -> hexSign []string
func DecodeMultiSign(multiSign string) ([]string, error) {
	unpackedMultiSign, err := unpacked(multiSignJson(), multiSign)
	if err != nil {
		return nil, err
	}

	var hexSignList []string
	for _, unpackedSign := range unpackedMultiSign {
		byteSignList := unpackedSign.([][]uint8)
		for _, byteSign := range byteSignList {
			hexSign := hex.EncodeToString(byteSign)
			hexSignList = append(hexSignList, hexSign)
		}
	}

	return hexSignList, nil
}

// hexSign list -> byteSign list
func hexSignListToByteSignList(hexSignList []string) ([][]byte, error) {
	var byteSignList [][]byte

	for _, hexSign := range hexSignList {
		byteSign, err := decodeSign(hexSign)
		if err != nil {
			return nil, err
		}
		byteSignList = append(byteSignList, byteSign)
	}

	return byteSignList, nil
}

// hexSign -> byteSign
func decodeSign(hexSign string) ([]byte, error) {
	unpackedSign, err := unpacked(signJson(), hexSign)
	if err != nil {
		return nil, err
	}

	signLen := len(unpackedSign)
	if signLen < 3 {
		return nil, fmt.Errorf("sign len : %d", signLen)
	}

	r := unpackedSign[0].([32]byte)
	s := unpackedSign[1].([32]byte)
	v := unpackedSign[2].(uint8)

	rs := append(r[:], s[:]...)
	signByte := append(rs, v)

	return signByte, nil
}

func unpacked(json string, data string) ([]interface{}, error) {
	jsonAbi, err := abi.JSON(strings.NewReader(json))
	if err != nil {
		return nil, err
	}

	if strings.HasPrefix(data, "0x") {
		data = strings.Replace(data, "0x", "", 1)
	}

	multiSignByte, err := hex.DecodeString(data)
	if err != nil {
		return nil, err
	}

	unpackedMultiSign, err := jsonAbi.Constructor.Inputs.Unpack(multiSignByte)
	if err != nil {
		return nil, err
	}

	return unpackedMultiSign, nil
}
