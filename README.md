# mpris-rpc

mpris-rpc is a discord rich presence status that shows media being played by apps that support MPRIS.

# TODO List
- [x] Remove the "" and [] from showing up on song and artist names
- [x] Adding some sort of whitelisting functionality, to allow choosing which apps are tracked
- [ ] add better error handeling in order to run it as background service
- [ ] disable whitelist if it's empty
- [ ] make a systemd service
- [ ] Showing song covers. Currently the image that is shown is the icon for the discord app you set
- [ ] Add an installer script to automatically compile and place the config in the right directory.

# Config
The config file mpris-rpc.kdl has to be placed in ~/.config/mpris-rpc. For now, the config has to be moved manually. mpris-rpc **will not** work without the config file.

# Installation
**You need go installed on your system for mpris-rpc to work** 

you also need a discord aplication in order for it to appear as an activity inside of discord

Its possible to either run it as is, by running `go run mpris-rpc.go` while in the files directory, or build it by running `go build mpris-rpc.go` and then running the file from a terminal.
