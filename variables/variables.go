package variables

import "net/http"

const charsetUTF8 = "charset=UTF-8"

// MIME types
const (
	MIMEApplicationJSON                  = "application/json"
	MIMEApplicationJSONCharsetUTF8       = MIMEApplicationJSON + "; " + charsetUTF8
	MIMEApplicationJavaScript            = "application/javascript"
	MIMEApplicationJavaScriptCharsetUTF8 = MIMEApplicationJavaScript + "; " + charsetUTF8
	MIMEApplicationXML                   = "application/xml"
	MIMEApplicationXMLCharsetUTF8        = MIMEApplicationXML + "; " + charsetUTF8
	MIMETextXML                          = "text/xml"
	MIMETextXMLCharsetUTF8               = MIMETextXML + "; " + charsetUTF8
	MIMEApplicationForm                  = "application/x-www-form-urlencoded"
	MIMEApplicationProtobuf              = "application/protobuf"
	MIMEApplicationMsgpack               = "application/msgpack"
	MIMETextHTML                         = "text/html"
	MIMETextHTMLCharsetUTF8              = MIMETextHTML + "; " + charsetUTF8
	MIMETextPlain                        = "text/plain"
	MIMETextPlainCharsetUTF8             = MIMETextPlain + "; " + charsetUTF8
	MIMEMultipartForm                    = "multipart/form-data"
	MIMEOctetStream                      = "application/octet-stream"
	MIMETextEventStream                  = "text/event-stream"
	MIMETextEventStreamCharsetUTF8       = "text/event-stream" + "; " + charsetUTF8
)

// Headers
const (
	HeaderWildcard                        = "*"
	HeaderTrue                            = "true"
	HeaderP3P                             = "P3p"
	HeaderFeaturePolicy                   = "Feature-Policy"
	HeaderAuthorization                   = "Authorization"
	HeaderProxyAuthenticate               = "Proxy-Authenticate"
	HeaderProxyAuthorization              = "Proxy-Authorization"
	HeaderWWWAuthenticate                 = "WWW-Authenticate"
	HeaderAge                             = "Age"
	HeaderCacheControl                    = "Cache-Control"
	HeaderNoCache                         = "No-Cache"
	HeaderClearSiteData                   = "Clear-Site-Data"
	HeaderExpires                         = "Expires"
	HeaderPragma                          = "Pragma"
	HeaderWarning                         = "Warning"
	HeaderAcceptCH                        = "Accept-CH"
	HeaderAcceptCHLifetime                = "Accept-CH-Lifetime"
	HeaderContentDPR                      = "Content-DPR"
	HeaderDPR                             = "DPR"
	HeaderEarlyData                       = "Early-Data"
	HeaderSaveData                        = "Save-Data"
	HeaderViewportWidth                   = "Viewport-Width"
	HeaderWidth                           = "Width"
	HeaderETag                            = "ETag"
	HeaderIfMatch                         = "If-Match"
	HeaderIfModifiedSince                 = "If-Modified-Since"
	HeaderIfNoneMatch                     = "If-None-Match"
	HeaderIfUnmodifiedSince               = "If-Unmodified-Since"
	HeaderLastModified                    = "Last-Modified"
	HeaderVary                            = "Vary"
	HeaderConnection                      = "Connection"
	HeaderKeepAlive                       = "Keep-Alive"
	HeaderAccept                          = "Accept"
	HeaderAcceptCharset                   = "Accept-Charset"
	HeaderAcceptEncoding                  = "Accept-Encoding"
	HeaderAcceptLanguage                  = "Accept-Language"
	HeaderCookie                          = "Cookie"
	HeaderExpect                          = "Expect"
	HeaderMaxForwards                     = "Max-Forwards"
	HeaderSetCookie                       = "Set-Cookie"
	HeaderAccessControlAllowCredentials   = "Access-Control-Allow-Credentials"
	HeaderAccessControlAllowHeaders       = "Access-Control-Allow-Headers"
	HeaderAccessControlAllowMethods       = "Access-Control-Allow-Methods"
	HeaderAccessControlAllowOrigin        = "Access-Control-Allow-Origin"
	HeaderAccessControlExposeHeaders      = "Access-Control-Expose-Headers"
	HeaderAccessControlMaxAge             = "Access-Control-Max-Age"
	HeaderAccessControlRequestHeaders     = "Access-Control-Request-Headers"
	HeaderAccessControlRequestMethod      = "Access-Control-Request-Method"
	HeaderOrigin                          = "Origin"
	HeaderTimingAllowOrigin               = "Timing-Allow-Origin"
	HeaderXPermittedCrossDomainPolicies   = "X-Permitted-Cross-Domain-Policies"
	HeaderDNT                             = "DNT"
	HeaderTk                              = "Tk"
	HeaderContentDisposition              = "Content-Disposition"
	HeaderContentEncoding                 = "Content-Encoding"
	HeaderContentLanguage                 = "Content-Language"
	HeaderContentLength                   = "Content-Length"
	HeaderContentLocation                 = "Content-Location"
	HeaderContentType                     = "Content-Type"
	HeaderForwarded                       = "Forwarded"
	HeaderVia                             = "Via"
	HeaderXForwardedFor                   = "X-Forwarded-For"
	HeaderXForwardedHost                  = "X-Forwarded-Host"
	HeaderXForwardedProto                 = "X-Forwarded-Proto"
	HeaderXForwardedProtocol              = "X-Forwarded-Protocol"
	HeaderXForwardedSsl                   = "X-Forwarded-Ssl"
	HeaderXUrlScheme                      = "X-Url-Scheme"
	HeaderLocation                        = "Location"
	HeaderFrom                            = "From"
	HeaderHost                            = "Host"
	HeaderReferer                         = "Referer"
	HeaderReferrerPolicy                  = "Referrer-Policy"
	HeaderUserAgent                       = "User-Agent"
	HeaderAllow                           = "Allow"
	HeaderServer                          = "Server"
	HeaderAcceptRanges                    = "Accept-Ranges"
	HeaderContentRange                    = "Content-Range"
	HeaderIfRange                         = "If-Range"
	HeaderRange                           = "Range"
	HeaderContentSecurityPolicy           = "Content-Security-Policy"
	HeaderContentSecurityPolicyReportOnly = "Content-Security-Policy-Report-Only"
	HeaderCrossOriginResourcePolicy       = "Cross-Origin-Resource-Policy"
	HeaderExpectCT                        = "Expect-CT"
	HeaderPermissionsPolicy               = "Permissions-Policy"
	HeaderPublicKeyPins                   = "Public-Key-Pins"
	HeaderPublicKeyPinsReportOnly         = "Public-Key-Pins-Report-Only"
	HeaderStrictTransportSecurity         = "Strict-Transport-Security"
	HeaderUpgradeInsecureRequests         = "Upgrade-Insecure-Requests"
	HeaderXContentTypeOptions             = "X-Content-Type-Options"
	HeaderXDownloadOptions                = "X-Download-Options"
	HeaderXFrameOptions                   = "X-Frame-Options"
	HeaderXPoweredBy                      = "X-Powered-By"
	HeaderXXSSProtection                  = "X-XSS-Protection"
	HeaderLastEventID                     = "Last-Event-ID"
	HeaderNEL                             = "NEL"
	HeaderPingFrom                        = "Ping-From"
	HeaderPingTo                          = "Ping-To"
	HeaderReportTo                        = "Report-To"
	HeaderTE                              = "TE"
	HeaderTrailer                         = "Trailer"
	HeaderTransferEncoding                = "Transfer-Encoding"
	HeaderSecWebSocketAccept              = "Sec-WebSocket-Accept"
	HeaderSecWebSocketExtensions          = "Sec-WebSocket-Extensions"
	HeaderSecWebSocketKey                 = "Sec-WebSocket-Key"
	HeaderSecWebSocketProtocol            = "Sec-WebSocket-Protocol"
	HeaderSecWebSocketVersion             = "Sec-WebSocket-Version"
	HeaderAcceptPatch                     = "Accept-Patch"
	HeaderAcceptPushPolicy                = "Accept-Push-Policy"
	HeaderAcceptSignature                 = "Accept-Signature"
	HeaderAltSvc                          = "Alt-Svc"
	HeaderDate                            = "Date"
	HeaderIndex                           = "Index"
	HeaderLargeAllocation                 = "Large-Allocation"
	HeaderLink                            = "Link"
	HeaderPushPolicy                      = "Push-Policy"
	HeaderRetryAfter                      = "Retry-After"
	HeaderServerTiming                    = "Server-Timing"
	HeaderSignature                       = "Signature"
	HeaderSignedHeaders                   = "Signed-Headers"
	HeaderSourceMap                       = "SourceMap"
	HeaderUpgrade                         = "Upgrade"
	HeaderXDNSPrefetchControl             = "X-DNS-Prefetch-Control"
	HeaderXPingback                       = "X-Pingback"
	HeaderXRequestID                      = "X-Request-ID"
	HeaderXRequestedWith                  = "X-Requested-With"
	HeaderXRobotsTag                      = "X-Robots-Tag"
	HeaderXUACompatible                   = "X-UA-Compatible"
)

// valid http methods
var Methods = []string{
	http.MethodGet,
	http.MethodPost,
	http.MethodPut,
	http.MethodConnect,
	http.MethodDelete,
	http.MethodHead,
	http.MethodPatch,
	http.MethodOptions,
	http.MethodTrace,
}