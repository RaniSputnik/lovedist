package builder

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	plist "github.com/DHowett/go-plist"
	"github.com/RaniSputnik/lovedist/builder/copy"
	"github.com/pierrre/archivefile/zip"
)

// Params are provided to build to specify the build requirements
//
// Name refers to the game name that the user will see
// InputDir should be the directory that contains the main.lua file
// OuputDir will be the output location for the build
// Logger can be used to get detailed information about the build as
// it proceeds. Can be safely left as nil.
type Params struct {
	Name      string
	InputDir  string
	OutputDir string
	Logger    *log.Logger

	*MacParams
	*WinParams
}

// MacParams are used to provide details of a build for the OSX target
type MacParams struct {
	PathToLoveApp    string
	BundleIdentifier string
}

// WinParams are used to provide details of a build for the Windows target
type WinParams struct {
	PathToLoveExe string
}

// Build creates a .love file from a given love source directory.
//
// By default a build will only create the .love file. Additional
// platforms will be built ONLY if their parameters are provided.
// Provide MacParams to build and osx .app file, and provide WinParams
// to build a Windows 32 .exe file.
func Build(params *Params) error {

	// Validate params
	if params.Logger == nil {
		params.Logger = doNotLogger()
	}
	if params.InputDir == params.OutputDir {
		return fmt.Errorf("Input directory must != output directory")
	}
	if !strings.HasSuffix(params.InputDir, "/") {
		params.InputDir += "/"
	}
	if params.Name == "" {
		params.Name = filepath.Base(params.InputDir)
	}

	// Create the .love file
	outfilename := fmt.Sprintf("%s.love", params.Name)
	outfile := filepath.Join(params.OutputDir, outfilename)
	params.Logger.Printf("Outputting to %s", outfile)
	fw, err := os.Create(outfile)
	if err != nil {
		return err
	}
	err = zip.Archive(params.InputDir, fw, func(archivePath string) {
		params.Logger.Printf("Zipping %s", archivePath)
	})
	if err != nil {
		return err
	}

	// TODO perform these steps in parallel

	// Build OSX if there are mac params
	if params.MacParams != nil {
		if err := buildMac(outfile, params); err != nil {
			return err
		}
	}

	// Build Windows if there are win params
	if params.WinParams != nil {
		if err := buildWin(outfile, params); err != nil {
			return err
		}
	}

	return nil
}

func buildMac(lovefile string, params *Params) error {
	params.Logger.Print("Starting build for osx")
	lovefilename := filepath.Base(lovefile)

	// Copy the love.app
	outapp := filepath.Join(params.OutputDir, "osx", fmt.Sprintf("%s.app", params.Name))
	if err := copy.Dir(params.PathToLoveApp, outapp); err != nil {
		return err
	}

	// Copy .love file into love app
	// TODO we have kept this a separate step because we could
	// perform "Copy love.app" and "Create .love" steps concurrently
	finallovepath := filepath.Join(outapp, "Contents", "Resources", lovefilename)
	if err := copy.File(lovefile, finallovepath); err != nil {
		return err
	}

	// Modify info.plist
	plistpath := filepath.Join(outapp, "Contents", "info.plist")
	plistfile, err := os.OpenFile(plistpath, os.O_RDWR, 0666)
	defer plistfile.Close()
	if err != nil {
		return err
	}
	var res loveAppPlist
	decoder := plist.NewDecoder(plistfile)
	if err := decoder.Decode(&res); err != nil {
		return err
	}
	res.BundleName = params.Name
	if params.BundleIdentifier != "" {
		res.BundleIdentifier = params.BundleIdentifier
	}
	if err := plistfile.Truncate(0); err != nil {
		return err
	}
	if _, err := plistfile.Seek(0, 0); err != nil {
		return err
	}
	encoder := plist.NewEncoder(plistfile)
	encoder.Indent("\t")
	if err := encoder.Encode(res); err != nil {
		return err
	}

	return nil
}

func buildWin(lovepath string, params *Params) error {
	params.Logger.Print("Starting build for win32")

	// Copy over dlls
	outpath := filepath.Join(params.OutputDir, "win32")
	params.Logger.Printf("Outputting to %s", outpath)
	exePath := filepath.Dir(params.PathToLoveExe)
	params.Logger.Printf("Copying files from %s", exePath)
	if err := copy.DirFilter(exePath, outpath, filterRequiredWinFiles); err != nil {
		return err
	}

	// Open the lovefile and love exe
	loveexe, err := os.Open(params.PathToLoveExe)
	defer loveexe.Close()
	if err != nil {
		return err
	}
	lovefile, err := os.Open(lovepath)
	defer lovefile.Close()
	if err != nil {
		return err
	}

	// Create the exe and copy the two files in
	outexe := filepath.Join(outpath, fmt.Sprintf("%s.exe", params.Name))
	f, err := os.Create(outexe)
	defer f.Close()
	if _, err := io.Copy(f, loveexe); err != nil {
		return err
	}
	if _, err := io.Copy(f, lovefile); err != nil {
		return err
	}

	return nil
}

func filterRequiredWinFiles(entry os.FileInfo) bool {
	if entry.IsDir() {
		return true
	}
	name := entry.Name()
	if name == "license.txt" {
		return true
	}
	ext := filepath.Ext(name)
	if ext == ".dll" || ext == ".ico" {
		return true
	}
	return false
}
