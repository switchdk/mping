# Multi Ping - mping
Tiny tool to ping multiple targets at the same time.

**Work in Progress**

# Install
with Go installed and working
1. Create a folder called `mping`
1. Download `main.go` into this folder
1. Execute `go build mping/main.go`
1. If the Go environment is correctly configured, `mping` should now be available

# Usage
* `mping` will automatically ping [Quad9](https://quad9.com/)
* `mping <hostname1> <hostname2> ... <hostnameN>`
* `mping -sleep <seconds> <hostname1> <hostname2> ... <hostnameN>` where `sleep` is the wait time between pings
