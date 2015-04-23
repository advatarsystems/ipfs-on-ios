
## ipfs-on-ios = An attempt to port the ipfs system to iOS.

###Goal:

	Get ipfs running on a mobile device, in this case: iOS (iPhone/iPad/etc.)

###Notes:

* See Internet Permanent Filesystem: [ipfs.io](http://ipfs.io)
* [Depends on the toolchain from the goios project](https://bitbucket.org/minux/goios)
* Use [homebrew](http://brew.sh/) to install go on your OSX development machine ($ brew install go)
* Currently depends on go 1.4.2 for cross-compiling the goios tools
* Will download and configure the goios ios64-new branch toolchain
* Includes a configured XCode project which builds Go code and which works with Objective-C and vice versa, to serve as a harness for an ipfs port.
* See also: [Using Go in Mobile Apps](https://medium.com/using-go-in-mobile-apps)
* When enabled (see gosources/main.go), the basic ipfs server code is used to pull in the Go dependencies, and expose missing syscall/functionality.

###Usage:

To get started, clone this repository, check the Makefile to verify GOROOT_BOOTSTRAP, type 'make' and sit back while the toolchain and then ipfsios projects are built.

###Current Status:

As of 23Apr2015, does *not* currently compile the full ipfs dependencies due to some missing syscalls.  It is set up to do so, but cannot proceed due to missing syscalls in goios implementation.

This project contains 3 components - the Makefile, the ipfsios project (for XCode), and - after the first 'make' - a local copy of the goios toolchain, configured for use on iOS.

As currently configured (see gosources/main.go), the XCode project will build the onboard gosources/main.go module successfully, and link it with a "Stock standard" XCode ViewController project to illustrate interaction between Go and C/Objective-C compiled code.  This is working 


All pull-requests welcome.

Contact: seclorum@icloud.com
