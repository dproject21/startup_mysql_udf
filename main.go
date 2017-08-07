package main

import (
	"unsafe"
	/*
	  #cgo CFLAGS: -I/usr/include/mysql
	  #include <mysql.h>
	  #include <string.h>
	  #include <stdlib.h>
	  #include <stdio.h>
	*/
	"C"
	"os"
	"unicode/utf8"

	"github.com/dproject21/startup_mysql_udf/command"
)

//export myexec_init
func myexec_init(initid *C.UDF_INIT, args *C.UDF_ARGS, message *C.char) C.my_bool {
	if int(args.arg_count) != 1 {
		return 1
	} else {
		return 0
	}
}

//export myexec
func myexec(initid *C.UDF_INIT, args *C.UDF_ARGS, result *C.char, length *C.ulong, is_null *C.char, error *C.char) *C.char {
	cArgs := C.GoString(*args.args)

	out, err := command.Exec(cArgs)
	if err != nil {
		os.Exit(1)
	}

	result = C.CString(string(out))
	*length = C.ulong(utf8.RuneCountInString(C.GoString(result)))
	return result
}

//export myexec_deinit
func myexec_deinit(initid *C.UDF_INIT) {
	C.free(unsafe.Pointer(initid.ptr))
}

func main() {
}
