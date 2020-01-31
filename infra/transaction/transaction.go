package transaction

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/8treenet/freedom"
	"github.com/jinzhu/gorm"
)

func init() {
	freedom.Booting(func(initiator freedom.Initiator) {
		initiator.BindInfra(false, func() *Transaction {
			return &Transaction{}
		})
	})
}

// Transaction .
type Transaction struct {
	freedom.Infra
	db *gorm.DB
}

// BeginRequest
func (t *Transaction) BeginRequest(rt freedom.Runtime) {
	t.db = nil
	t.Infra.BeginRequest(rt)
}

func (t *Transaction) Execute(fun func() error) (e error) {
	return t.execute(fun, nil, nil)
}

// Execute Execute local transaction.
func (t *Transaction) ExecuteTx(fun func() error, ctx context.Context, opts *sql.TxOptions) (e error) {
	return t.execute(fun, ctx, opts)
}

// Execute Execute local transaction.
func (t *Transaction) execute(fun func() error, ctx context.Context, opts *sql.TxOptions) (e error) {
	if t.db != nil {
		panic("unknown error")
	}
	if ctx != nil && opts != nil {
		t.db = t.DB().BeginTx(ctx, opts)
	} else {
		t.db = t.DB().Begin()
	}

	t.Runtime.Store().Set("freedom_local_transaction_db", t.db)

	defer func() {
		if perr := recover(); perr != nil {
			t.db.Rollback()
			t.db = nil
			e = errors.New(fmt.Sprint(perr))
			t.Runtime.Store().Remove("freedom_local_transaction_db")
			return
		}

		deferdb := t.db
		t.Runtime.Store().Remove("freedom_local_transaction_db")
		t.db = nil
		if e != nil {
			e2 := deferdb.Rollback()
			if e2.Error != nil {
				e = errors.New(e.Error() + "," + e2.Error.Error())
			}
			return
		}
		e = deferdb.Commit().Error
	}()
	e = fun()
	return
}
