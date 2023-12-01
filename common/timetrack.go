package common

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"
)

func getFunctionName(f func() int) string {
	return strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[2]
}

func TimeMethodCall(f func() int) {
	name := getFunctionName(f)

	start := time.Now()
	result := f()
	elapsed := time.Since(start)

	fmt.Printf("%v  %20v, took %s\n", name, result, elapsed.Round(time.Microsecond))
}
