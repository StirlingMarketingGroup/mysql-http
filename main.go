package main

// #include <string.h>
// #include <stdbool.h>
// #include <mysql.h>
// #cgo CFLAGS: -O3 -I/usr/include/mysql -fno-omit-frame-pointer
import "C"
import (
	"net/http"
	"unsafe"
)

func msg(message *C.char, s string) {
	m := C.CString(s)
	defer C.free(unsafe.Pointer(m))

	C.strcpy(message, m)
}

//export http_touch_init
func http_touch_init(initid *C.UDF_INIT, args *C.UDF_ARGS, message *C.char) C.bool {
	if args.arg_count != 1 {
		msg(message, "`http_get` requires 1 parameter: the URL string")
		return C.bool(true)
	}

	argsTypes := (*[2]uint32)(unsafe.Pointer(args.arg_type))

	argsTypes[0] = C.STRING_RESULT
	initid.maybe_null = 1

	return C.bool(false)
}

//export http_touch
func http_touch(initid *C.UDF_INIT, args *C.UDF_ARGS, isNull *C.char, isError *C.char) C.longlong {
	*isNull = 1

	c := 1
	argsArgs := (*[1 << 30]*C.char)(unsafe.Pointer(args.args))[:c:c]

	a := make([]string, c, c)
	for i, argsArg := range argsArgs {
		// This should be the correct way, but lengths come through as "0"
		// for everything after the first argument, so hopefully we don't
		// encounter any URLs or param names with null bytes in them (not really that worried)
		// a[i] = C.GoStringN(argsArg, C.int(argsLengths[i]))

		a[i] = C.GoString(argsArg)
	}

	if a[0] == "" {
		return 0
	}

	http.Get(a[0])

	return 0
}

func main() {}
