package dist

type Output struct {
	/* TODO: What is this? */
}

type Builder interface {
	// Build accepts an io.Reader that will read the *.love
	// file containing the games assets. Build returns the
	// result of the build and an error if the build was
	// unsuccessful for any reason.
	Build(input io.Reader) (Output, error)
}

type notImplementedBuilder struct {}

func (notImplementedBuilder) Build(input io.Reader) (res Output, err error) {
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
