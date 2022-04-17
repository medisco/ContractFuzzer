package event

import "github.com/gongbell/contractfuzzer/db/domain"

type InstrumentExecutionEvent struct {
	Input        string
	Instructions []uint64
	Transaction  domain.Transaction
}
