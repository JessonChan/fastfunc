package fastfunc

import "time"

var (
	runTimes int64 = 1e8
)

func SetRunTimes(rt int64) {
	runTimes = rt
}

func Run(f func()) int64 {
	t := time.Now()
	var i int64
	for ; i < runTimes; i++ {
		f()
	}
	e := time.Now()
	return e.UnixNano() - t.UnixNano()
}
