# go_sdl2
Example app to explain SDL2 using Go. 

In this app I make a screen, draw colored squares, and read out events.

Most of this app is based on this tutorial series: https://markkeeley.xyz/2016/go-sdl2-lesson-1/

SDL wrapper: https://github.com/veandco/go-sdl2


# Installation
### Install Go
This will differ per OS. I'll assume that you have followed those instructions

### Make sure sdl2, sld2_image, and sdl2_ttf are installed
In linux, you can simply search in the package manager for SDL, and install the
needed packages. The names in the title are the same in almost every distribution of linux.

### Download Go packages
You can now download the external packages that we'll be using:
```bash
go get github.com/veandco/go-sdl2/sdl
go get github.com/veandco/go-sdl2/img
go get github.com/veandco/go-sdl2/ttf
```

If you now look in your GOPATH, you'll find a pkg and a src folder.

### Clone this repository to your computer
Go to the folder where you want to clone this folder to. In my case:

```bash
cd $HOME/git/
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

You should see a white screen with black squares popping in and out of view.
There should be a red label in the top right corner, a spinning cat, and red text
saying hello. When you press a key, the key event should be printed in black for
some seconds.

To compile the application, and then run it, you can do the following:
```bash
go build
.\go_sdl2`
```

To quickly build (and clean up the previous build before hand), you can use the
provided run script. To enable it, do this once to enable the script to be executed:
```bash
chmod +x run
```

Now you can build and run the program by doing:
```bash
./run
```

# Controls
- Closing window closes application (yes you have to build this)
- Pressing a key will print out the event in the terminal and on the screen.
