package rocks

import (
	"bytes"
	"encoding/binary"
	"math"
)

// Raw key:
// +key,type = value
// +name,s = "latermoon"

var (
	SEP = []byte{','}
	KEY = []byte{'+'} // Key Prefix
	SOK = []byte{'['} // Start of Key
	EOK = []byte{']'} // End of Key
)

type ElementType byte

const (
	STRING ElementType = 's'
	HASH               = 'h'
	LIST               = 'l'
	ZSET               = 'z'
)

// 字节最大范围
const MAXBYTE byte = math.MaxUint8

func rawKey(key []byte, t ElementType) []byte {
	return bytes.Join([][]byte{KEY, key, SEP, []byte{byte(t)}}, nil)
}

// 范围判断 min <= v <= max
func between(v, min, max []byte) bool {
	return bytes.Compare(v, min) >= 0 && bytes.Compare(v, max) <= 0
}

// 复制数组
func copyBytes(src []byte) []byte {
	dst := make([]byte, len(src))
	copy(dst, src)
	return dst
}

// 使用二进制存储整形
func Int64ToBytes(i int64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}
