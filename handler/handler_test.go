package handler_test

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/RaniSputnik/lovedist/handler"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	w, r := recordRequest(http.MethodGet, "/_ah/ping", nil)

	runHandler(w, r)

	assert.Equal(t, http.StatusOK, w.Code, "Status code should be '200 - OK'")
	assert.Equal(t, "pong", w.Body.String(), "Body should be 'pong'")
}

func TestInfo(t *testing.T) {
	w, r := recordRequest(http.MethodGet, "/_ah/info", nil)
	runHandler(w, r)

	type infoResponse struct {
		Love struct {
			SupportedVersions []string `json:"supported_versions"`
		} `json:"love"`
	}

	expectedVersions := []string{"11.2.0", "0.10.2"}

	var res infoResponse
	err := json.NewDecoder(w.Body).Decode(&res)

	assert.NoError(t, err, "Expected to decode body successfully")
	assert.Equal(t, http.StatusOK, w.Code, "Status code should be '200 - OK'")
	assert.Equal(t, expectedVersions, res.Love.SupportedVersions)
}

func TestBuild(t *testing.T) {
	const buildPath = "/build"

	t.Run("FailsWhenBodyEmpty", func(t *testing.T) {
		w, r := recordRequest(http.MethodPost, buildPath, nil)

		runHandler(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code, "Status should be '400 - Bad Request'")
		if assert.Contains(t, w.Header().Get("Content-Type"), "text/html", "Content-Type should be HTML") {
			assert.Contains(t, w.Body.String(), "Bad Request", "Body should contain the text 'Bad Request'")
			assert.Contains(t, w.Body.String(), `href="/"`, "Body should contain a link back to the index page")
		}
	})

	t.Run("FailsWhenMissingUploadFile", func(t *testing.T) {
		contentType, body := mustCreateForm(map[string]io.Reader{
			"foo": strings.NewReader("bar"),
		})
		w, r := recordRequest(http.MethodPost, buildPath, body)
		r.Header.Add("Content-Type", contentType)

		runHandler(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code, "Status should be '400 - Bad Request'")
		if assert.Contains(t, w.Header().Get("Content-Type"), "text/html", "Content-Type should be HTML") {
			assert.Contains(t, w.Body.String(), "Bad Request", "Body should contain the text 'Bad Request'")
			assert.Contains(t, w.Body.String(), `href="/"`, "Body should contain a link back to the index page")
		}
	})

	t.Run("FailsWhenLoveNotFound", func(t *testing.T) {
		contentType, body := formWithValidUpload()
		w, r := recordRequest(http.MethodPost, buildPath, body)
		r.Header.Add("Content-Type", contentType)

		handler := handler.New("./theWrongDirectory", "./theWrongDirectory")
		handler.ServeHTTP(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Code, "Status should be '500 - Internal Server Error'")
		if assert.Contains(t, w.Header().Get("Content-Type"), "text/html", "Content-Type should be HTML") {
			assert.Contains(t, w.Body.String(), "Something went wrong", "Body should contain the text 'Something went wrong'")
			assert.Contains(t, w.Body.String(), `href="/"`, "Body should contain a link back to the index page")
		}
	})

	t.Run("FailsWhenUnsupportedLoveVersionSpecified", func(t *testing.T) {
		contentType, body := formWithValidUploadAndLoveVersion("invalid")
		w, r := recordRequest(http.MethodPost, buildPath, body)
		r.Header.Add("Content-Type", contentType)

		runHandler(w, r)

		assert.Equal(t, http.StatusBadRequest, w.Code, "Status should be '400 - Bad Request'")
		if assert.Contains(t, w.Header().Get("Content-Type"), "text/html", "Content-Type should be HTML") {
			assert.Contains(t, w.Body.String(), "Bad Request", "Body should contain the text 'Bad Request'")
			assert.Contains(t, w.Body.String(), `href="/"`, "Body should contain a link back to the index page")
		}
	})

	// TODO: How do we test the correct love version was used when specified?

	t.Run("ReturnsZip", func(t *testing.T) {
		contentType, body := formWithValidUpload()
		w, r := recordRequest(http.MethodPost, buildPath, body)
		r.Header.Add("Content-Type", contentType)

		runHandler(w, r)

		assert.Equal(t, http.StatusOK, w.Code, "Status should be '200 - OK'")

		gotContentDisposition := w.Header().Get("Content-Disposition")
		assert.Contains(t, gotContentDisposition, "attachment;", "Should be an attachment")
		assert.Contains(t, gotContentDisposition, ".zip", "Should be a .zip file")
	})
}

func recordRequest(method, target string, body io.Reader) (*httptest.ResponseRecorder, *http.Request) {
	return httptest.NewRecorder(), httptest.NewRequest(method, target, body)
}

func runHandler(w http.ResponseWriter, r *http.Request) http.Handler {
	// TODO set these with env vars, otherwise tests can not be run
	// from the main package
	const buildDirectory = "../build"
	const loveDirectory = "../love"

	handler := handler.New(buildDirectory, loveDirectory)
	handler.ServeHTTP(w, r)
	return handler
}

func formWithValidUpload() (contentType string, body *bytes.Buffer) {
	file, err := os.Open("./fixture.love")
	if err != nil {
		panic(err)
	}
	return mustCreateForm(map[string]io.Reader{
		"uploadfile": file,
	})
}

func formWithValidUploadAndLoveVersion(version string) (contentType string, body *bytes.Buffer) {
	file, err := os.Open("./fixture.love")
	if err != nil {
		panic(err)
	}

	return mustCreateForm(map[string]io.Reader{
		"uploadfile":  file,
		"loveversion": strings.NewReader(version),
	})
}

func mustCreateForm(values map[string]io.Reader) (contentType string, body *bytes.Buffer) {
	var b bytes.Buffer
	var err error
	w := multipart.NewWriter(&b)
	for key, r := range values {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}

		if x, ok := r.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
				panic(err)
			}
		} else {
			if fw, err = w.CreateFormField(key); err != nil {
				panic(err)
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			panic(err)
		}
	}
	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()

	return w.FormDataContentType(), &b
}
