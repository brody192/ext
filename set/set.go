package set

import (
	"net/http"
	"strconv"

	"github.com/brody192/ext/extvar"
)

// sets content length header to length of length
func ContentLength[cL int | int64](w http.ResponseWriter, length cL) {
	w.Header().Set(extvar.HeaderContentLength, strconv.FormatInt(int64(length), 10))
}

// sets content type header with given mime type
func ContentType(w http.ResponseWriter, mimeType string) {
	w.Header().Set(extvar.HeaderContentType, mimeType)
}

// sets content disposition header to attachment with provided filename
func AttachmentFilename(w http.ResponseWriter, filename string) {
	w.Header().Set(extvar.HeaderContentDisposition, "attachment; filename=\""+filename+"\"")
}

// sets content disposition header to attachment
func Attachment(w http.ResponseWriter) {
	w.Header().Set(extvar.HeaderContentDisposition, "attachment")
}

// sets content disposition header to inline
func Inline(w http.ResponseWriter) {
	w.Header().Set(extvar.HeaderContentDisposition, "inline")
}

// sets content type to text plain
func PlainText(w http.ResponseWriter) {
	w.Header().Set(extvar.HeaderContentType, extvar.MIMETextPlainCharsetUTF8)
}

// sets content type to application json
func JSON(w http.ResponseWriter) {
	w.Header().Set(extvar.HeaderContentType, extvar.MIMEApplicationJSONCharsetUTF8)
}

// sets content type to text html
func HTML(w http.ResponseWriter) {
	w.Header().Set(extvar.HeaderContentType, extvar.MIMETextHTMLCharsetUTF8)
}
