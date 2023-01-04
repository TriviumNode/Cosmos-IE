package rest

import (
	"encoding/json"
	"strings"

	"go.uber.org/zap"
)

type Blocks struct {
	Block struct {
		Header struct {
			ChainID          string `json:"chain_id"`
			Height           string
			Proposer_address string
		}

		Last_commit struct {
			Signatures []struct {
				Block_id_flag     string
				Validator_address string
			}
		}
	}
}

func GetBlocks(chain string, log *zap.Logger) Blocks {
	var b Blocks
	var res []uint8

	switch chain {
	case "iris":
		res, _ = runRESTCommand("/cosmos/base/tendermint/v1beta1/blocks/latest")
	default:
		res, _ = runRESTCommand("/blocks/latest")
	}

	// handle error
	if strings.Contains(string(res), "error") {
		log.Error("", zap.Bool("Success", false), zap.String("err", string(res)))
	} else {
		json.Unmarshal(res, &b)
		// log.Info("Common Info", zap.Bool("Success", true), zap.String("Block Info", b),)
	}

	return b
}
