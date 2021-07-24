package registry

import "github.com/legato/infrastructure/database/repository"

func NewTxnInsertTrack() repository.TxnInsertTrack {
	return repository.NewTxnInsertTrack(LegatoDB)
}
