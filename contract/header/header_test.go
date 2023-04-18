package header

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeaderType(t *testing.T) {
	assert.Equal(t, "fabric", HeaderTypeToString(FABRIC))
	assert.Equal(t, "ethereum", HeaderTypeToString(ETHEREUM))
	assert.Equal(t, "bitcoin", HeaderTypeToString(BITCOIN))
}

func TestFabricUnmarshaledHeader(t *testing.T) {
	data := `{"type":"fabric","publicKey":[230,184,92,89,155,62,23,91,105,229,52,154,153,127,229,108,77,101,229,141,84,169,29,41,161,67,243,81,175,199,130,96]}`
	header, err := UnmarshaledHeader([]byte(data))
	assert.NoError(t, err)
	assert.Equal(t, "*header.Header", reflect.TypeOf(header).String())
	assert.True(t, Verify(header))
	assert.Equal(t, "0x997fe56c4d65e58d54a91d29a143f351afc78260", Address(header))
}

func TestEthereumUnmarshaledHeader(t *testing.T) {
	data := `{"type":"ethereum","hash":[128,29,39,54,174,8,161,149,10,129,65,200,211,211,206,47,161,186,186,172,27,168,85,205,123,74,37,238,160,103,224,51],"signature":[20,121,85,158,189,104,217,226,97,32,13,86,243,154,197,158,28,84,235,175,113,180,64,171,227,20,127,116,192,231,223,20,32,141,151,201,192,175,147,104,127,45,129,15,80,225,35,109,205,11,100,157,203,48,127,118,159,10,46,183,53,240,85,241,1]}`
	header, err := UnmarshaledHeader([]byte(data))
	assert.NoError(t, err)
	assert.Equal(t, "*header.Header", reflect.TypeOf(header).String())
	assert.True(t, Verify(header))
	assert.Equal(t, "0x93b2Cb3061e36Ed3099d003fF78cd685b424e95b", Address(header))
}

func TestBitcoinUnmarshaledHeader(t *testing.T) {
	data := `{"type":"bitcoin","hash":[225,86,14,255,221,206,228,196,31,41,149,3,50,222,48,133,60,47,102,65,76,4,29,240,131,175,235,171,163,111,13,79],"publicKey":[3,109,140,113,83,5,95,202,169,119,150,69,17,133,187,253,98,144,65,80,49,159,82,96,232,115,54,172,36,202,224,73,53],"signature":[48,68,2,32,77,219,237,43,232,87,107,195,180,151,206,175,62,189,11,65,196,224,173,181,169,45,19,179,61,142,125,65,25,165,83,159,2,31,81,84,216,211,151,130,74,37,116,183,202,1,114,73,214,19,13,26,201,131,5,125,43,98,24,238,120,32,117,58,175]}`
	header, err := UnmarshaledHeader([]byte(data))
	assert.NoError(t, err)
	assert.Equal(t, "*header.Header", reflect.TypeOf(header).String())
	assert.True(t, Verify(header))
	assert.Equal(t, "15VDTyzYK6SiH4kCdT89bEaskB15QS79F9", Address(header))
}
