package main

import (
	"github.com/enjekt/commons"
	"fmt"
)

type BlackvaultDB interface {
	UpsertTokenPad(token *commons.Token, panda *commons.Pad)
	GetPad(token *commons.Token) *commons.Pad
}

type blackvaultDBContext struct {
	commons.Database
	tokenToPandaBucket string
}

func NewBlackvaultDB() BlackvaultDB {
	bvdb := new(blackvaultDBContext)
	bvdb.DBName = "blackvaultDB.db"
	bvdb.tokenToPandaBucket = "TokensToPads"
	return BlackvaultDB(bvdb)
}

func (bvdb *blackvaultDBContext) UpsertTokenPad(token *commons.Token, panda *commons.Pad) {
	bvdb.Upsert(bvdb.tokenToPandaBucket, token.ToString(), panda.ToString())
}

func (bvdb *blackvaultDBContext) GetPad(token *commons.Token) *commons.Pad {
	fmt.Printf("Token to retrieve pad for: %s \n", token.ToString())
	tokenStr:=token.ToString()
	padStr := bvdb.Get(bvdb.tokenToPandaBucket,tokenStr)
	fmt.Printf("Fetched from db the padStr: %s \n", padStr)
	return commons.InitPad(padStr)
}
