package rest

import (
)

type commitInfo struct {
        ChainId                  string
        ValidatorPrecommitStatus float64 // [0]: false, [1]: true
//        ValidatorProposingStatus float64 // [0]: false, [1]: true
}

func getCommit(blockData Blocks, consHexAddr string) commitInfo {
//func getCommit(blockData Blocks) commitInfo {

        var cInfo commitInfo

//        blockProposer := blockData.Block.Header.Proposer_address

        cInfo.ChainId = blockData.Block.Header.ChainID
        cInfo.ValidatorPrecommitStatus  = 0.0
//        cInfo.ValidatorPrecommitStatus, cInfo.ValidatorProposingStatus = 0.0, 0.0


        for _, v := range blockData.Block.Last_commit.Signatures {

                func() {
                        defer func() {

                                if r := recover(); r != nil {
                                        // precommit failure validator
                                }
                        }()

                        if consHexAddr == v.Validator_address {

                                        cInfo.ValidatorPrecommitStatus = 1.0
                        }
/*
                        if consHexAddr == blockProposer {
                                cInfo.ValidatorProposingStatus = 1.0
                        }
*/
                }()

        }

        return cInfo
}
