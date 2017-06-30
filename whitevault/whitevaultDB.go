package main

import (
	"github.com/enjekt/commons"
)

type WhitevaultDB interface {
	UpsertTokenPanda(token *commons.Token, panda *commons.Panda)
	GetPanda(token *commons.Token) *commons.Panda
}

type WhitevaultDBContext struct {
	commons.Database
	tokenToPandaBucket string
}

func NewWhitevaultDB() WhitevaultDB {
	wvdb := new(WhitevaultDBContext)
	wvdb.DBName = "whitevault.db"
	wvdb.tokenToPandaBucket = "TokensToPandas"
	return WhitevaultDB(wvdb)
}

func (wvdb *WhitevaultDBContext) UpsertTokenPanda(token *commons.Token, panda *commons.Panda) {
	wvdb.Upsert(wvdb.tokenToPandaBucket, token.ToString(), panda.ToString())
}

func (wvdb *WhitevaultDBContext) GetPanda(token *commons.Token) *commons.Panda {
	return commons.InitPanda(wvdb.Get(wvdb.tokenToPandaBucket, token.ToString()))
}
