package repository

import (
	"crypto/sha512"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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

func appendLimitOffsetMods(mods []qm.QueryMod, limit, offset *int) []qm.QueryMod {
	if limit != nil {
		mods = append(mods, qm.Limit(*limit))
	}
	if offset != nil {
		mods = append(mods, qm.Offset(*offset))
	}
	return mods
}
