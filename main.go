package main

import (
	"flag"
	"log"
	"os"

	"github.com/RaniSputnik/lovedist/builder"
)

func main() {
	pName := flag.String("name", "", "The output name of the game")
	pBuildWin := flag.Bool("win", true, "If true, a build for Windows will be attempted")
	pLoveExe := flag.String("exe", "./love/win32/love.exe", "Path to the love.exe, required for a Windows build")
	pBuildOsx := flag.Bool("osx", true, "If true, a build for OSX will be attempted")
	pLoveApp := flag.String("app", "./love/osx/love.app", "Path to the love.app, required for OSX build")
	pBundleID := flag.String("bundleid", "", "The bundle identifier of the game, usually in reverse domain form eg. com.company.product")
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		flag.Usage()
		return
	}

	logger := log.New(os.Stderr, "", 0)
	params := &builder.Params{
		Name:      *pName,
		InputDir:  args[0],
		OutputDir: args[1],
		Logger:    logger,
	}

	if *pBuildWin {
		params.WinParams = &builder.WinParams{
			PathToLoveExe: *pLoveExe,
		}
	}
	if *pBuildOsx {
		params.MacParams = &builder.MacParams{
			PathToLoveApp:    *pLoveApp,
			BundleIdentifier: *pBundleID,
		}
	}

	if err := builder.Build(params); err != nil {
		logger.Fatal(err)
	} else {
		logger.Println("Completed successfully")
	}
}
