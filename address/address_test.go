package address

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var args = []string{"1", "2", "3", "4"}

func TestAddress(t *testing.T) {

	fabric, err := NewAddress(
		FABRIC,
		[]byte{102, 58, 58, 86, 253, 117, 63, 136, 154, 25, 198, 237, 64, 87, 204, 66, 116, 82, 54, 102, 250, 76, 200, 142, 95, 120, 25, 59, 54, 16, 90, 51},
		nil,
		nil,
	)
	assert.NoError(t, err)
	expect := `{"type":"fabric","publicKey":"Zjo6Vv11P4iaGcbtQFfMQnRSNmb6TMiOX3gZOzYQWjM="}`
	result, err := fabric.Serialize()
	assert.NoError(t, err)
	assert.Equal(t, expect, string(result))
	mdlArgs, err := fabric.AppendArgs(args)
	assert.NoError(t, err)
	assert.Equal(t, []string{`{"type":"fabric","publicKey":"Zjo6Vv11P4iaGcbtQFfMQnRSNmb6TMiOX3gZOzYQWjM="}`, "1", "2", "3", "4"}, mdlArgs)

	ethereum, err := NewAddress(
		ETHEREUM,
		nil,
		[]byte{128, 29, 39, 54, 174, 8, 161, 149, 10, 129, 65, 200, 211, 211, 206, 47, 161, 186, 186, 172, 27, 168, 85, 205, 123, 74, 37, 238, 160, 103, 224, 51},
		[]byte{20, 121, 85, 158, 189, 104, 217, 226, 97, 32, 13, 86, 243, 154, 197, 158, 28, 84, 235, 175, 113, 180, 64, 171, 227, 20, 127, 116, 192, 231, 223, 20, 32, 141, 151, 201, 192, 175, 147, 104, 127, 45, 129, 15, 80, 225, 35, 109, 205, 11, 100, 157, 203, 48, 127, 118, 159, 10, 46, 183, 53, 240, 85, 241, 1},
	)
	t.Log(ethereum.SignatureToHex())
	assert.NoError(t, err)
	expect = `{"type":"ethereum","hash":"gB0nNq4IoZUKgUHI09POL6G6uqwbqFXNe0ol7qBn4DM=","signature":"FHlVnr1o2eJhIA1W85rFnhxU669xtECr4xR/dMDn3xQgjZfJwK+TaH8tgQ9Q4SNtzQtkncswf3afCi63NfBV8QE="}`
	result, err = ethereum.Serialize()
	assert.NoError(t, err)
	assert.Equal(t, expect, string(result))
	mdlArgs, err = ethereum.AppendArgs(args)
	assert.NoError(t, err)
	assert.Equal(t, []string{`{"type":"ethereum","hash":"gB0nNq4IoZUKgUHI09POL6G6uqwbqFXNe0ol7qBn4DM=","signature":"FHlVnr1o2eJhIA1W85rFnhxU669xtECr4xR/dMDn3xQgjZfJwK+TaH8tgQ9Q4SNtzQtkncswf3afCi63NfBV8QE="}`, "1", "2", "3", "4"}, mdlArgs)

	bitcoin, err := NewAddress(
		BITCOIN,
		[]byte{3, 109, 140, 113, 83, 5, 95, 202, 169, 119, 150, 69, 17, 133, 187, 253, 98, 144, 65, 80, 49, 159, 82, 96, 232, 115, 54, 172, 36, 202, 224, 73, 53},
		[]byte{225, 86, 14, 255, 221, 206, 228, 196, 31, 41, 149, 3, 50, 222, 48, 133, 60, 47, 102, 65, 76, 4, 29, 240, 131, 175, 235, 171, 163, 111, 13, 79},
		[]byte{48, 68, 2, 32, 77, 219, 237, 43, 232, 87, 107, 195, 180, 151, 206, 175, 62, 189, 11, 65, 196, 224, 173, 181, 169, 45, 19, 179, 61, 142, 125, 65, 25, 165, 83, 159, 2, 32, 43, 81, 84, 216, 211, 151, 130, 74, 37, 116, 183, 202, 1, 114, 73, 214, 19, 13, 26, 201, 131, 5, 125, 43, 98, 24, 238, 120, 32, 117, 58, 175},
	)
	assert.NoError(t, err)
	expect = `{"type":"bitcoin","publicKey":"A22McVMFX8qpd5ZFEYW7/WKQQVAxn1Jg6HM2rCTK4Ek1","hash":"4VYO/93O5MQfKZUDMt4whTwvZkFMBB3wg6/rq6NvDU8=","signature":"MEQCIE3b7SvoV2vDtJfOrz69C0HE4K21qS0Tsz2OfUEZpVOfAiArUVTY05eCSiV0t8oBcknWEw0ayYMFfStiGO54IHU6rw=="}`
	result, err = bitcoin.Serialize()
	assert.NoError(t, err)
	assert.Equal(t, expect, string(result))
	mdlArgs, err = bitcoin.AppendArgs(args)
	assert.NoError(t, err)
	assert.Equal(t, []string{`{"type":"bitcoin","publicKey":"A22McVMFX8qpd5ZFEYW7/WKQQVAxn1Jg6HM2rCTK4Ek1","hash":"4VYO/93O5MQfKZUDMt4whTwvZkFMBB3wg6/rq6NvDU8=","signature":"MEQCIE3b7SvoV2vDtJfOrz69C0HE4K21qS0Tsz2OfUEZpVOfAiArUVTY05eCSiV0t8oBcknWEw0ayYMFfStiGO54IHU6rw=="}`, "1", "2", "3", "4"}, mdlArgs)
}

func TestDeserialize(t *testing.T) {
	/* Fabric */
	expect := `{"type":"fabric","publicKey":"Zjo6Vv11P4iaGcbtQFfMQnRSNmb6TMiOX3gZOzYQWjM="}`
	_, err := Deserialize([]byte(expect))
	assert.NoError(t, err)

	/* Ethereum */
	expect = `{"type":"ethereum","hash":"gB0nNq4IoZUKgUHI09POL6G6uqwbqFXNe0ol7qBn4DM=","signature":"FHlVnr1o2eJhIA1W85rFnhxU669xtECr4xR/dMDn3xQgjZfJwK+TaH8tgQ9Q4SNtzQtkncswf3afCi63NfBV8QE="}`
	_, err = Deserialize([]byte(expect))
	assert.NoError(t, err)

	expect = `{"type":"bitcoin","publicKey":"A22McVMFX8qpd5ZFEYW7/WKQQVAxn1Jg6HM2rCTK4Ek1","hash":"4VYO/93O5MQfKZUDMt4whTwvZkFMBB3wg6/rq6NvDU8=","signature":"MEQCIE3b7SvoV2vDtJfOrz69C0HE4K21qS0Tsz2OfUEZpVOfAiArUVTY05eCSiV0t8oBcknWEw0ayYMFfStiGO54IHU6rw=="}`
	_, err = Deserialize([]byte(expect))
	assert.NoError(t, err)
}
