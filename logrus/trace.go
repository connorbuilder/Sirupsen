package logrus

import (
	"runtime"
	"strings"
	"errors"
)


func Trace(depth int) (Fields, error) {
	pc, file, line, ok := runtime.Caller(depth)
	if !ok {
		return nil, errors.New("runtime caller failed.")
	}
	
	path := strings.Split(file, "/")
	fname := path[len(path)-1]
	funcname := runtime.FuncForPC(pc).Name()
	
	fields := make(Fields, 3)
	fields["func"] = funcname
	fields["file"] = fname
	fields["line"] = line
	
	return fields, nil
}