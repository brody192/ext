package extrespond

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/brody192/ext/extset"
	"github.com/brody192/ext/extvar"
)

// accepts a byte slice
//
// sets content length of v
//
// sets content type to mimeType
//
// writes v to w
func Blob(w http.ResponseWriter, mimeType string, v []byte, code int) {
	extset.ContentLength(w, len(v))
	w.Header().Set(extvar.HeaderContentType, mimeType)
	w.WriteHeader(code)
	w.Write(v)
}

// accepts a string
//
// sets content length of v
//
// sets content type to text plain
//
// writes v to w
func PlainText(w http.ResponseWriter, v string, code int) {
	Blob(w, extvar.MIMETextPlainCharsetUTF8, []byte(v), code)
}

// accepts a byte slice
//
// sets content length of v
//
// sets content type to text plain
//
// writes v to w
func PlainTextBlob(w http.ResponseWriter, v []byte, code int) {
	Blob(w, extvar.MIMETextPlainCharsetUTF8, v, code)
}

// accepts a string
//
// sets content length of v
//
// sets content type to text html
//
// writes v to w
func HTML(w http.ResponseWriter, v string, code int) {
	Blob(w, extvar.MIMETextHTMLCharsetUTF8, []byte(v), code)
}

// accepts a byte slice
//
// sets content length of v
//
// sets content type to text html
//
// writes v to w
func HTMLBlob(w http.ResponseWriter, v []byte, code int) {
	Blob(w, extvar.MIMETextHTMLCharsetUTF8, v, code)
}

// accepts a string
//
// sets content length of v
//
// sets content type to application json
//
// writes v to w
func JSONString(w http.ResponseWriter, v string, code int) {
	Blob(w, extvar.MIMEApplicationJSONCharsetUTF8, []byte(v), code)
}

// accepts a byte slice
//
// sets content length of v
//
// sets content type to application json
//
// writes v to w
func JSONBlob(w http.ResponseWriter, v []byte, code int) {
	Blob(w, extvar.MIMEApplicationJSONCharsetUTF8, v, code)
}

// encodes v to a buffer
//
// sets content length of the buffer
//
// sets content type to application json
//
// writes buffer to w
//
// resets buffer
func JSON(w http.ResponseWriter, v any, code int) {
	jsonIndented(w, v, "", code)
}

// encodes v to a buffer with an indent of two spaces
//
// sets content length of the buffer
//
// sets content type to application json
//
// writes buffer to w
//
// resets buffer
func JSONIndented(w http.ResponseWriter, v any, code int) {
	jsonIndented(w, v, "  ", code)
}

func jsonIndented(w http.ResponseWriter, v any, indent string, code int) {
	var buf = &bytes.Buffer{}
	var enc = json.NewEncoder(buf)
	enc.SetIndent("", indent)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONBlob(w, buf.Bytes(), code)

	buf.Reset()
}

// if no content type header was previously set MIMEOctetStream will be used
//
// copy's re to w
func Stream(w http.ResponseWriter, re io.Reader, code int) error {
	if ct := w.Header().Get(extvar.HeaderContentType); ct != "" {
		w.Header().Set(extvar.HeaderContentType, extvar.MIMEOctetStream)
	}

	w.WriteHeader(code)

	_, err := io.Copy(w, re)
	return err
}

// checks for existence of template by given name
//
// returns ErrTemplateNotFound if given template name was not found in t
//
// executes the given template name in t to a buffer
//
// writes buffer to w
//
// resets buffer
func Template(t *template.Template, w http.ResponseWriter, name string, data any, code int) error {
	var templateToExecute = t.Lookup(name)
	if templateToExecute == nil {
		return fmt.Errorf("%w: %s", ErrTemplateNotFound, name)
	}

	var buf = &bytes.Buffer{}
	var err = templateToExecute.ExecuteTemplate(buf, name, data)
	if err != nil {
		return err
	}

	HTMLBlob(w, buf.Bytes(), code)

	buf.Reset()

	return err
}
