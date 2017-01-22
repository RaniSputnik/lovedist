# lovedist

A tool for building Love 2d games for distribution. Following the instructions [here](https://love2d.org/wiki/Game_Distribution).

### Installation

If you have [go](https://golang.org/) installed then it's as simple as cloning the repo and installing with the go tool;

```
git clone https://github.com/RaniSputnik/lovedist.git
cd lovedist
go install .
```

### Usage

```
lovedist [flags] /path/to/game /path/to/output

  -app string
    	Path to the love.app, required for OSX build (default "/Applications/love.app")
  -bundleid string
    	The bundle identifier of the game, usually in reverse domain form eg. com.company.product
  -exe string
    	Path to the love.exe, required for a Windows build (default "love.exe")
  -name string
    	The output name of the game
  -osx
    	If true, a build for OSX will be attempted
  -win
    	If true, a build for Windows will be attempted
```
