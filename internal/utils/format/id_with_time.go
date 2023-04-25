package format

import (
	"fmt"
	"time"
)

type Id string

func NewItemID(codeId string, now time.Time) Id {
	obj := Id(fmt.Sprintf("%s-%s-%o-%s", codeId, now.Format("060102"), time.Now().Nanosecond(), "MTR"))
	return obj
}

func (r Id) String() string {
	return string(r)
}
