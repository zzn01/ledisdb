package store

import (
	"github.com/zzn01/ledisdb/store/driver"
)

type Slice interface {
	driver.ISlice
}
