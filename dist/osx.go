package dist

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/RaniSputnik/lovedist/builder/copy"
	"howett.net/plist"
)

type osxBuilder struct {
	pathToLoveApp string

	outputDir string // TODO: Remove me, should be common across all builders
}

func (b *osxBuilder) Build(p Project, input io.Reader) (res Result, err error) {
	// Copy the love.app
	outapp := filepath.Join(b.outputDir, "osx", fmt.Sprintf("%s.app", p.Name))
	if err = copy.Dir(b.pathToLoveApp, outapp); err != nil {
		return
	}

	// Copy .love file into love app
	// TODO we have kept this a separate step because we could
	// perform "Copy love.app" and "Create .love" steps concurrently
	finallovepath := filepath.Join(outapp, "Contents", "Resources", fmt.Sprintf("%s.love", p.Name))
	file, err := os.Create(finallovepath)
	if err != nil {
		return
	}
	defer file.Close()
	if _, err = io.Copy(file, input); err != nil {
		return
	}

	// Modify info.plist
	plistpath := filepath.Join(outapp, "Contents", "Info.plist")
	plistfile, err := os.OpenFile(plistpath, os.O_RDWR, 0666)
	defer plistfile.Close()
	if err != nil {
		return
	}
	var resPlist loveAppPlist
	decoder := plist.NewDecoder(plistfile)
	if err = decoder.Decode(&res); err != nil {
		return
	}
	resPlist.BundleName = p.Name
	if p.BundleID != "" {
		resPlist.BundleIdentifier = p.BundleID
	}
	if err = plistfile.Truncate(0); err != nil {
		return
	}
	if _, err = plistfile.Seek(0, 0); err != nil {
		return
	}
	encoder := plist.NewEncoder(plistfile)
	encoder.Indent("\t")
	if err = encoder.Encode(res); err != nil {
		return
	}
	// TODO: Set something about the build to indicate where
	// it was outputted?
	return res, nil
}
