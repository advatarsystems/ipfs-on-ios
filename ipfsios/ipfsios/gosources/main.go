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
	"code.google.com/p/go.net/context"
	core "github.com/jbenet/go-ipfs/core"
	corenet "github.com/jbenet/go-ipfs/core/corenet"
	fsrepo "github.com/jbenet/go-ipfs/repo/fsrepo"
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

func ipfs_client() {
	if len(os.Args) < 2 {
		fmt.Println("Please give a peer ID as an argument")
		return
	}
	target, err := peer.IDB58Decode(os.Args[1])
	if err != nil {
		panic(err)
	}

	// Basic ipfsnode setup
	r := fsrepo.At("~/.go-ipfs")
	if err := r.Open(); err != nil {
		panic(err)
	}

	nb := core.NewNodeBuilder().Online()
	nb.SetRepo(r)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nd, err := nb.Build(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("I am peer %s dialing %s\n", nd.Identity, target)

	con, err := corenet.Dial(nd, target, "/app/whyrusleeping")
	if err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(os.Stdout, con)
}

/*
func ipfs_server() {
	// Basic ipfsnode setup
	r := fsrepo.At("~/.go-ipfs")
	if err := r.Open(); err != nil {
		panic(err)
	}

	nb := core.NewNodeBuilder().Online()
	nb.SetRepo(r)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nd, err := nb.Build(ctx)
	if err != nil {
		panic(err)
	}

	list, err := corenet.Listen(nd, "/app/whyrusleeping")
	if err != nil {
		panic(err)
	}
	fmt.Printf("I am peer: %s\n", nd.Identity)

	for {
		con, err := list.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer con.Close()

		fmt.Fprintln(con, "Hello! This is whyrusleepings awesome ipfs service")
		fmt.Printf("Connection from: %s\n", con.Conn().RemotePeer())
	}
}
*/

// Go main entry point
func main() {
	fmt.Fprintf(os.Stderr, "Startup of Golang iosmain ::\n")
	go StartGoServer()
	C.iosmain(0, nil)
}
