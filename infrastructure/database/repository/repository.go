package repository

import (
	"crypto/sha512"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// sqlExecutor sql.DB構造体のインターフェイス
type sqlExecutor interface {
	boil.ContextExecutor
	boil.ContextBeginner
}

func genHash(src string) string {
	hash := sha512.Sum512([]byte(src))
	return fmt.Sprintf("%x", hash)
}
