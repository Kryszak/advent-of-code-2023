package common

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"
)

func getFunctionName(f func(string) int) string {
	return strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[2]
}

func TimeMethodCall(path string, f func(string) int) {
	name := getFunctionName(f)

	start := time.Now()
	result := f(path)
	elapsed := time.Since(start)

	fmt.Printf("%v  %20v, took %s\n", name, result, elapsed.Round(time.Microsecond))
}
