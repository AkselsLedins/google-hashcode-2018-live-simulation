# Ubuntu Install

```
make deps
make
```

## Gopath

disclaimer: first time using Go.

```
mkdir ~/go
export GOPATH=~/go:$PWD
mkdir -p ~/go/src/github.com/AkselsLedins/
ln -s $PWD ~/go/src/github.com/AkselsLedins/google-hashcode-2018-live-simulation
```

## Troubleshooting

> make: go: Command not found

`sudo apt install golang`


> glfw/src/x11_platform.h:39:33: fatal error: X11/Xcursor/Xcursor.h: No such file or directory

`sudo apt install libxcursor-dev`



> glfw/src/x11_platform.h:42:35: fatal error: X11/extensions/Xrandr.h: No such file or directory

`sudo apt install libxrandr-dev`


> /usr/bin/ld: cannot find -lXi

`sudo apt install libxi-dev`
