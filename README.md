
## ipfs-on-ios = An attempt to port the ipfs system to iOS.

###Goal:

	Get ipfs running on a mobile device, in this case: iOS (iPhone/iPad/etc.)

###Notes:

* See Internet Permanent Filesystem: [ipfs.io](http://ipfs.io)
* [Depends on the toolchain from the goios project](https://bitbucket.org/minux/goios)
* Use [homebrew](http://brew.sh/) to install go on your OSX development machine ($ brew install go)
* Currently depends on go 1.4.2 for cross-compiling the goios tools
* Will download and configure the goios ios64-new branch toolchain
* Includes a configured XCode project 

###Usage:

	To get started, clone this repository, check the Makefile to verify GOROOT_BOOTSTRAP, type 'make' and sit back while the toolchain and then ipfsios projects are built.

###Current Status:

Currently builds a working cgo toolchain from a fresh goios repository, and then during XCode build, calls it from the ipfsios/ipfsios/build-go.sh script, successfully compiling Go code with the Objective-C of iOS.

However as of 23Apr2015, does *not* currently compile the full ipfs dependencies due to some missing syscalls.  

Once the toolchain and ifpsios projects have built successfully, modify ipfsios/ipsios/gosources/main.go to include a simple ipfs-server, and rebuild to see current ipfs state.

All pull-requests welcome.

Contact: seclorum@icloud.com
