package domain

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/jessevdk/go-flags"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/id64"
	"github.com/kokizzu/lexid"
	"github.com/kpango/fastime"
	"github.com/rs/zerolog/log"

	"street/conf"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file common.go
//go:generate replacer -afterprefix 'Id" form' 'Id,string" form' type common.go
//go:generate replacer -afterprefix 'json:"id"' 'json:"id,string"' type common.go
//go:generate replacer -afterprefix 'By" form' 'By,string" form' type common.go
// go:generate msgp -tests=false -file common.go -o  common__MSG.GEN.go
//go:generate farify doublequote --file common.go

type RequestCommon struct {
	TracerContext context.Context   `json:"-" form:"tracerContext" query:"tracerContext" long:"tracerContext" msg:"-"`
	RequestId     string            `json:"requestId" form:"requestId" query:"requestId" long:"requestId" msg:"requestId"`
	SessionToken  string            `json:"sessionToken" form:"sessionToken" query:"sessionToken" long:"sessionToken" msg:"sessionToken"`
	UserAgent     string            `json:"userAgent" form:"userAgent" query:"userAgent" long:"userAgent" msg:"userAgent"`
	IpAddress     string            `json:"ipAddress" form:"ipAddress" query:"ipAddress" long:"ipAddress" msg:"ipAddress"`
	OutputFormat  string            `json:"outputFormat,omitempty" form:"outputFormat" query:"outputFormat" long:"outputFormat" msg:"outputFormat"` // defaults to json
	Uploads       map[string]string `json:"uploads,omitempty" form:"uploads" query:"uploads" long:"uploads" msg:"uploads"`                          // form key and temporary file path
	Debug         bool              `json:"debug,omitempty" form:"debug" query:"debug" long:"debug" msg:"debug"`
	Header        string            `json:"header,omitempty" form:"header" query:"header" long:"header" msg:"header"`
	RawBody       string            `json:"rawBody,omitempty" form:"rawBody" query:"rawBody" long:"rawBody" msg:"rawBody"`
	Host          string            `json:"host" form:"host" query:"host" long:"host" msg:"host"`
	Action        string            `json:"action" form:"action" query:"action" long:"action" msg:"action"`

	// in seconds
	now int64 `json:"-" form:"now" query:"now" long:"now" msg:"-"`
}

func (l *RequestCommon) ToFiberCtx(ctx *fiber.Ctx, out any, rc *ResponseCommon, in any) error {
	defer l.deleteTempFiles()
	if rc.StatusCode != http.StatusOK {
		ctx.Status(rc.StatusCode)
	}
	rc.DecorateSession(ctx)
	switch l.OutputFormat {
	case ``, `json`, fiber.MIMEApplicationJSON:
		ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		if l.Debug {
			rc.Debug = in
		}
		byt, err := json.Marshal(out)
		if L.IsError(err, `json.Marshal: %#v`, out) {
			return err
		}
		_, err = ctx.Write(byt)
		if L.IsError(err, `ctx.Write failed: `+string(byt)) {
			return err
		}
		// TODO: log size/bytes written
		if l.Debug {
			log.Print(string(byt))
		}
	case `html`:
		// do nothing
	default:
		return errors.New(`ToFiberCtx unhandled format: ` + l.OutputFormat)
	}
	return nil
}

func (l *RequestCommon) UnixNow() int64 {
	if l.now == 0 {
		l.now = fastime.UnixNow()
	}
	return l.now
}

func (i *RequestCommon) TimeNow() time.Time {
	return time.Unix(i.UnixNow(), 0)
}

func (l *RequestCommon) FromFiberCtx(ctx *fiber.Ctx, tracerCtx context.Context) {
	l.RequestId = lexid.ID()
	l.SessionToken = ctx.Cookies(conf.CookieName, l.SessionToken)
	l.UserAgent = string(ctx.Request().Header.UserAgent())
	l.Host = ctx.Protocol() + `://` + ctx.Hostname()
	// from nginx reverse proxy
	l.IpAddress = ctx.IP()
	if l.IpAddress == `` {
		l.IpAddress = `0.0.0.0`
	}
	// "Accept":"*/*", "Connection":"close", "Content-Length":"0", "Host":"admin.hapstr.xyz", "User-Agent":"curl/7.81.0", "X-Forwarded-For":"182.253.163.10", "X-Forwarded-Proto":"https", "X-Real-Ip":"182.253.163.10"
	l.now = fastime.UnixNow()
	file, err := ctx.FormFile(`fileBinary`)
	if err == nil {
		l.Uploads = map[string]string{}
		target := `/tmp/` + id64.SID() + filepath.Ext(file.Filename)
		err = ctx.SaveFile(file, target)
		if !L.IsError(err, `failed saving upload to: `+target) {
			l.Uploads[file.Filename] = target
		}
	}
	l.TracerContext = tracerCtx
}

func (l *RequestCommon) ToCli(file *os.File, out any) {
	defer l.deleteTempFiles()
	switch l.OutputFormat {
	case ``, `json`, fiber.MIMEApplicationJSON:
		byt, err := json.Marshal(out)
		if L.IsError(err, `json.Marshal: %#v`, out) {
			return
		}
		_, err = file.Write(byt)
		if L.IsError(err, `file.Write failed: `+string(byt)) {
			return
		}
		// TODO: log size/bytes written
	default:
		L.Print(`ToCli unhandled format: ` + l.OutputFormat)
	}
}

func (l *RequestCommon) FromCli(file *os.File, ctx context.Context) {
	l.RequestId = lexid.ID()
	// TODO: read from args/stdin/config-file?
	// l.SessionToken =
	_, err := flags.Parse(&l)
	L.PanicIf(err, `flags.Parse`)
	l.UserAgent = `CLI` // TODO: add input format combination, eg. json-stdin
	l.IpAddress = `127.0.0.1`
	l.TracerContext = ctx
	l.now = fastime.UnixNow()
}

func (l *RequestCommon) deleteTempFiles() {
	for _, tmpPath := range l.Uploads {
		// TODO: delete temporary uploads
		_ = tmpPath
	}
}

type ResponseCommon struct {
	SessionToken string `json:"sessionToken" form:"sessionToken" query:"sessionToken" long:"sessionToken" msg:"sessionToken"`
	Error        string `json:"error" form:"error" query:"error" long:"error" msg:"error"`
	StatusCode   int    `json:"status" form:"statusCode" query:"statusCode" long:"statusCode" msg:"statusCode"`
	Debug        any    `json:"debug,omitempty" form:"debug" query:"debug" long:"debug" msg:"debug"`
	Redirect     string `json:"redirect,omitempty" form:"redirect" query:"redirect" long:"redirect" msg:"redirect"`

	// action trace
	traces []string
	actor  uint64
}

func (o *ResponseCommon) HasError() bool {
	return o.StatusCode >= 400 || len(o.Error) > 0
}

func (o *ResponseCommon) SetRedirect(to string) {
	o.StatusCode = 302
	o.Redirect = to
}

func (o *ResponseCommon) SetError(code int, errStr string) {
	o.StatusCode = code
	o.Error = errStr
}

func (o *ResponseCommon) SetErrorf(code int, errFmt string, args ...any) {
	o.StatusCode = code
	o.Error = fmt.Sprintf(errFmt, args...)
}

func (l *ResponseCommon) DecorateSession(ctx *fiber.Ctx) {
	if l.SessionToken != `` {
		if l.SessionToken == conf.CookieLogoutValue {
			ctx.ClearCookie(conf.CookieName)
			ctx.Cookie(&fiber.Cookie{
				Name:    conf.CookieName,
				Expires: time.Unix(0, 0),
			})
			return
		}
		ctx.Cookie(&fiber.Cookie{
			Name:    conf.CookieName,
			Value:   l.SessionToken,
			Expires: time.Now().AddDate(0, 0, conf.CookieDays),
		})
	}
}

func (o *ResponseCommon) AddTrace(act string) {
	o.traces = append(o.traces, act)
}

func (o *ResponseCommon) Traces() string {
	return strings.Join(o.traces, `|`)
}
