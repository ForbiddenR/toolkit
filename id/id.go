package id

import (
	"fmt"
	"sync"
	"time"
)

type IdGenerator interface {
	// Get returns an id produced by IdGenerator
	Get(string, string) string
}

const timeString = "060102150405"

type TransactionIdGenerator struct {
	sequence uint16
	locker   sync.Locker
}

func NewTransactionIdGenerator() *TransactionIdGenerator {
	return &TransactionIdGenerator{
		locker: &sync.Mutex{},
	}
}

func (g *TransactionIdGenerator) Get(sn, cid string) string {
	g.locker.Lock()
	defer g.locker.Unlock()
	now := time.Now()
	transactionId := fmt.Sprintf("%s%02s%s%04d", sn, cid, now.Format(timeString), g.sequence%10000)
	g.sequence++
	return transactionId
}
