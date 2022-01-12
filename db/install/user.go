// Copyright 2021-2022, the SS project owners. All rights reserved.
// Please see the OWNERS and LICENSE files for details.

package dbinstall

import (
	"github.com/palchukovsky/ss"
	lambda "github.com/palchukovsky/ss/api/gateway/auth/lambda"
	"github.com/palchukovsky/ss/db"
	"github.com/palchukovsky/ss/ddb"
	ddbinstall "github.com/palchukovsky/ss/ddb/install"
)

type user struct{ ddbinstall.TableAbstraction }

func newUserTable(ddb ddbinstall.DB, log ss.Log) ddbinstall.Table {
	return user{
		TableAbstraction: ddbinstall.NewTableAbstraction(ddb, db.User{}, log),
	}
}

func (table user) Create() error {
	return table.TableAbstraction.Create(
		[]ddb.IndexRecord{&lambda.FirebaseIndex{}})
}

func (table user) Setup() error {
	return table.EnableTimeToLive("anonymExpiration")
}

func (user) InsertData() error { return nil }
