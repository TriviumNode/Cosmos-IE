package rest

import (
	"encoding/json"
	"fmt"
	"strings"

	"go.uber.org/zap"
)

type CommunityPool struct {
	Pool []Coin
}

func getCommunityPool(log *zap.Logger) []Coin {

	var p CommunityPool

	res, _ := runRESTCommand("/cosmos/distribution/v1beta1/community_pool")
	json.Unmarshal(res, &p)
	// log
	if strings.Contains(string(res), "not found") {
		// handle error
		log.Fatal("", zap.Bool("Success", false), zap.String("err", string(res)))
	} else {
		log.Info("\t", zap.Bool("Success", true), zap.String("Commission", fmt.Sprint(p.Pool)))
	}

	return p.Pool
}
