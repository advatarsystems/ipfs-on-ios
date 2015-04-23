package main

/*
// adjust LDFLAGS to produce linkable code - note Go starts first, so
// be sure to change AppDelegate main() for chaining ..
#cgo LDFLAGS: -Wl,-U,_iosmain,-U,_PopUpDialogBox
#cgo LDFLAGS: -framework UIKit -framework Foundation -framework CoreGraphics
extern char* PopUpDialogBox(char* msg);
extern void iosmain(int argc, char *argv[]);
*/
import "C"


import (
	"fmt"
	"net/http"
	"os"

    // ipfs
/*
    core "github.com/jbenet/go-ipfs/core"
	corenet "github.com/jbenet/go-ipfs/core/corenet"
	fsrepo "github.com/jbenet/go-ipfs/repo/fsrepo"
	"code.google.com/p/go.net/context"
*/
)

//export AddUsingGo
func AddUsingGo(a int, b int) int {
	return a + b
}

// basic server for debugging ..
func HandleHttpRequest(w http.ResponseWriter, r *http.Request) {
    userInput := C.GoString(C.PopUpDialogBox(C.CString(r.URL.Path[1:])))
	fmt.Fprintf(w, "Hello from Go on an iPhone! Response: %s", userInput)
}

func StartGoServer() {
	fmt.Fprintf(os.Stderr, "Starting net/http/Server ...\n")
	http.HandleFunc("/", HandleHttpRequest)
	http.ListenAndServe(":6060", nil)
}


// Go main entry point
func main() {
	fmt.Fprintf(os.Stderr, "Startup of Golang iosmain ::\n")
	go StartGoServer()
	C.iosmain(0, nil)
}
