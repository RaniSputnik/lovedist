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
	// Build accepts two arguments, a project struct containing
	// information about the project being compiled and an
	// io.Reader that will read the *.love file containing
	// the games assets.
	//
	// Build returns the result of the build and an error
	// if the build was unsuccessful for any reason.
	Build(p Project, input io.Reader) (Result, error)
}

type notImplementedBuilder struct{}

func (notImplementedBuilder) Build(p Project, input io.Reader) (res Result, err error) {
	return res, errors.New("not implemented")
}

// Win returns a builder that can build Love games for
// distribution on Windows PCs
func Win() Builder { return notImplementedBuilder{} }

// OSX returns a builder that can build Love games for
// distribution on desktop Mac.
func OSX() Builder { return notImplementedBuilder{} }

// Linux returns a builder that can build Love games for
// distribution on Linux computers.
func Linux() Builder { return notImplementedBuilder{} }

// Web returns a builder that can build Love games for
// distribution on the web using web assembly.
func Web() Builder { return notImplementedBuilder{} }
