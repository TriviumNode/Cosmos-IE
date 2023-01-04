package rest

import (
	"encoding/json"
	"strings"

	"go.uber.org/zap"

	utils "github.com/node-a-team/Cosmos-IE/utils"
)

var ()

type stakingPool struct {
	//	Height	string	`json:"height"`
	Pool struct {
		Not_bonded_tokens string `json:"not_bonded_tokens"`
		Bonded_tokens     string `json:"bonded_tokens"`
		Total_supply      float64
	}
}

type totalSupply struct {
	//	Height string	`json:"height"`
	Amount Coin
}

func getStakingPool(chain string, denom string, log *zap.Logger) stakingPool {

	var sp stakingPool

	res, err := runRESTCommand("/cosmos/staking/v1beta1/pool")
	json.Unmarshal(res, &sp)

	// log
	if strings.Contains(string(res), "not found") {
		// handle error
		log.Fatal("", zap.Bool("Success", false), zap.String("err", string(res)))
	} else if err != nil {
		log.Fatal("", zap.Bool("Success", false), zap.String("err", "Failed to connect to REST-Server"))
	} else {
		log.Info("", zap.Bool("Success", true), zap.String("Bonded tokens", sp.Pool.Bonded_tokens))
	}

	sp.Pool.Total_supply = getTotalSupply(chain, denom, log)

	return sp
}

func getTotalSupply(chain string, denom string, log *zap.Logger) float64 {
	var ts totalSupply
	var res []uint8
	switch chain {
	case "iris":
		res, _ := runRESTCommand("/cosmos/bank/v1beta1/supply/by_denom?denom=" + denom)
		json.Unmarshal(res, &ts)
	default:
		res, _ := runRESTCommand("/cosmos/bank/v1beta1/supply/" + denom)
		json.Unmarshal(res, &ts)
	}

	// log
	if ts.Amount.Amount == "" {
		// handle error
		log.Error("", zap.Bool("Success", false), zap.String("err", string(res)))
	} else {
		log.Info("", zap.Bool("Success", true), zap.String("Total Supply", ts.Amount.Amount))
	}

	return utils.StringToFloat64(ts.Amount.Amount)
}
