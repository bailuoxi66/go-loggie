package util

import (
	"bailuoxi66/go-loggie/pkg/core/log"
	"github.com/xhit/go-str2duration/v2"
	"reflect"
	"time"
	"unsafe"
)

func ByteToStringUnsafe(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{Data: bh.Data, Len: bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

func StringToByteUnsafe(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{Data: sh.Data, Len: sh.Len, Cap: sh.Len}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func StringToDuration(ds string) time.Duration {
	duration, err := str2duration.ParseDuration(ds)
	if err != nil {
		log.Info("parse string(%s) to time.Duration error. err: %v", ds, err)
	}
	return duration
}
