package bridge

import "math/big"

type Msg struct {
	FromChainId string   `json:"fromChainId"`
	FromTxHash  string   `json:"fromTxHash"`
	ToTokenAddr string   `json:"toTokenAddr"`
	ToUserAddr  string   `json:"toUserAddr"`
	Amount      *big.Int `json:"amount"`
}

func multiSignJson() string {
	json := `[{
			"inputs": [
				{"name": "signs", "type": "bytes[]"}								
			],
			"type": "constructor"
		}]`
	return json
}

func signJson() string {
	return `[{	
			"inputs": [
				{ "name": "r", "type": "bytes32" },
				{ "name": "s", "type": "bytes32" },
				{ "name": "v", "type": "uint8" }				
			],
			"type": "constructor"
		}]`
}

func msgJson() string {
	json := `[{	
			"inputs": [
				{"name": "fromChainId", "type": "string"},
				{"name": "fromTxHash", "type": "string"},
				{"name": "toTokenAddr", "type": "string"},
				{"name": "toUserAddr", "type": "string"},
				{"name": "amount", "type": "uint256"}
			],
			"type": "constructor"
		}]`
	return json
}
