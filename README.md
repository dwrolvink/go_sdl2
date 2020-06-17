# go_sdl2
Example app to explain SDL2 using Go. 

In this app I make a screen, draw colored squares, and read out events.

Most of this app is based on this tutorial series: https://markkeeley.xyz/2016/go-sdl2-lesson-1/

SDL wrapper: https://github.com/veandco/go-sdl2


# Installation
### Install Go
This will differ per OS. I'll assume that you have followed those instructions

### Make sure sdl2 and sld2_image are installed
In linux, you can simply search in the package manager for SDL.

### Set your Gopath & Download packages
Go works in a slightly annoying way. It won't be able to find packages, unless they are
in the GOPATH. I have opted to set my GOPATH to home/user/git/go. You can do the same
by adding the following to your .bashrc (or .zshrc in my case):

```bash
export GOPATH=$HOME/git/go
```

You can now download the external package that we'll be using:
```bash
go get github.com/veandco/go-sdl2/sdl
```

If you now look in your GOPATH, you'll find a pkg and a src folder.
Go into the src folder. In my case:

```bash
cd $HOME/git/go/src
```

Now we can clone this repository

```bash 
git clone https://github.com/dwrolvink/go_sdl2.git
```

### Test run
Go into the newly created folder and run the application to test it:
```bash
cd go_sdl2
go run main.go
```

You should see a white screen with black squares popping in and out of view

To compile the application, and then run it, you can do the following:
```bash
go build
.\go_sdl2`
```

# Controls
- Closing window closes application (yes you have to build this)
- Pressing a key will print out the event in the terminal
