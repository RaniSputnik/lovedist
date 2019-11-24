package dist

import (
	"errors"
	"io"
)

type Project struct {
	Name string
	// BundleID is typically a string conveying domain
	// ownership eg. com.example.GameName.
	BundleID string
}

type Result struct {
	/* TODO: What is this? */
}

type Builder interface {
	// Build generates a compiled Love2d game.
	//
	// Build uses the given project to provide metadata about
	// the game (name, bundle id etc.) and reads the .love
	// contents from the given reader. The result of the build
	// will be written to the given output directory.
	//
	// Build returns the result of the build and an error
	// if the build was unsuccessful for any reason.
	Build(p Project, input io.Reader, output string) (Result, error)
}

type notImplementedBuilder struct{}

func (notImplementedBuilder) Build(p Project, input io.Reader, output string) (res Result, err error) {
	return res, errors.New("not implemented")
}

// Win returns a builder that can build Love games for
// distribution on Windows PCs
func Win() Builder { return notImplementedBuilder{} }

// OSX returns a builder that can build Love games for
// distribution on desktop Mac. You must provide the path
// to the love.app that will be bundled with the game.
func OSX(loveApp string) Builder {
	return &osxBuilder{
		pathToLoveApp: loveApp,
	}
}

// Linux returns a builder that can build Love games for
// distribution on Linux computers.
func Linux() Builder { return notImplementedBuilder{} }

// Web returns a builder that can build Love games for
// distribution on the web using web assembly.
func Web() Builder { return notImplementedBuilder{} }
