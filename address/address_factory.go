package address

func NewUserAddress(addrType string) Address {
	switch addrType {
	case "MSP":
		return NewAddressNormal()
	case "ETH":
		return NewAddressEth()
	}

	return NewAddressNormal()
}
