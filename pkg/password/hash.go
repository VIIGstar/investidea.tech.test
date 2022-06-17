package password

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func Hash(raw string) string {
	h := md5.New()
	io.WriteString(h, raw)
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}
