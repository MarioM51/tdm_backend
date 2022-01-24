package logs

import (
	"fmt"
	"sync"
)

type logger struct {
	isProd bool
}

var instance *logger
var once sync.Once

func New(isProd bool) *logger {
	once.Do(func() {
		instance = &logger{
			isProd: isProd,
		}
	})
	return instance
}

func (log *logger) Log(isSensible bool, line string) {
	if log.isProd && isSensible {
		return
	} else {
		fmt.Println(line)
	}
}

func (log *logger) LogF(isSensible bool, line string, args ...interface{}) {
	if log.isProd && isSensible {
		return
	} else {
		fmt.Printf(line, args...)
	}
}
