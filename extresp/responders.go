package extresp

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
// sets content length of v
// sets content type to mimeType
// writes v to w
func Blob(w http.ResponseWriter, mimeType string, v []byte) {
	extset.ContentLength(w, len(v))
	w.Header().Set(extvar.HeaderContentType, mimeType)
	w.Write(v)
}

// accepts a string
// sets content length of v
// sets content type to text plain
// writes v to w
func PlainText(w http.ResponseWriter, v string) {
	Blob(w, extvar.MIMETextPlainCharsetUTF8, []byte(v))
}

// accepts a byte slice
// sets content length of v
// sets content type to text plain
// writes v to w
func PlainTextBlob(w http.ResponseWriter, v []byte) {
	Blob(w, extvar.MIMETextPlainCharsetUTF8, v)
}

// accepts a string
// sets content length of v
// sets content type to text html
// writes v to w
func HTML(w http.ResponseWriter, v string) {
	Blob(w, extvar.MIMETextHTMLCharsetUTF8, []byte(v))
}

// accepts a byte slice
// sets content length of v
// sets content type to text html
// writes v to w
func HTMLBlob(w http.ResponseWriter, v []byte) {
	Blob(w, extvar.MIMETextHTMLCharsetUTF8, v)
}

// accepts a string
// sets content length of v
// sets content type to application json
// writes v to w
func JSONString(w http.ResponseWriter, v string) {
	Blob(w, extvar.MIMEApplicationJSONCharsetUTF8, []byte(v))
}

// accepts a byte slice
// sets content length of v
// sets content type to application json
// writes v to w
func JSONBlob(w http.ResponseWriter, v []byte) {
	Blob(w, extvar.MIMEApplicationJSONCharsetUTF8, v)
}

// encodes v to a buffer
// sets content length of the buffer
// sets content type to application json
// writes buffer to w
// resets buffer
func JSON(w http.ResponseWriter, v any) {
	jsonIndented(w, v, "")
}

// encodes v to a buffer with an indent of two spaces
// sets content length of the buffer
// sets content type to application json
// writes buffer to w
// resets buffer
func JSONIndented(w http.ResponseWriter, v any) {
	jsonIndented(w, v, "  ")
}

func jsonIndented(w http.ResponseWriter, v any, indent string) {
	var buf = &bytes.Buffer{}
	var enc = json.NewEncoder(buf)
	enc.SetIndent("", indent)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	extset.ContentLength(w, buf.Len())
	w.Header().Set(extvar.HeaderContentType, extvar.MIMEApplicationJSONCharsetUTF8)
	w.Write(buf.Bytes())
	buf.Reset()
}

// if no content type header was previously set MIMEOctetStream will be used
// copy's re to w
func Stream(w http.ResponseWriter, re io.Reader) error {
	if ct := w.Header().Get(extvar.HeaderContentType); ct != "" {
		w.Header().Set(extvar.HeaderContentType, extvar.MIMEOctetStream)
	}

	_, err := io.Copy(w, re)
	return err
}

// checks for existence of template by given name
// returns ErrTemplateNotFound if given template name was not found in t
// executes the given template name in t to a buffer
// writes buffer to w
// resets buffer
func Template(t *template.Template, w http.ResponseWriter, name string, data any) error {
	var templateToExecute = t.Lookup(name)
	if templateToExecute == nil {
		return fmt.Errorf("%w: %s", ErrTemplateNotFound, name)
	}

	var buf = &bytes.Buffer{}
	var err = templateToExecute.ExecuteTemplate(buf, name, data)
	if err != nil {
		return err
	}

	extset.ContentLength(w, buf.Len())
	extset.HTML(w)

	_, err = io.Copy(w, buf)

	buf.Reset()

	return err
}
