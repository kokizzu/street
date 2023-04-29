package presentation

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/Z"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"street/conf"
	"street/domain"
)

type WebServer struct {
	AuthOltp *Tt.Adapter
	AuthOlap *Ch.Adapter
	Log      *zerolog.Logger
	Cfg      conf.WebConf
}

var requiredHeader = M.SS{
	//domain.SomeUrl: `X-CC-Webhook-Signature`,
}

// priority:
// 1. query string
// 2. body
// 3. params
func webApiParseInput(ctx *fiber.Ctx, reqCommon *domain.RequestCommon, in any, url string) error {
	body := ctx.Body()
	path := S.LeftOf(url, `?`) // without API_PREFIX
	if header, ok := requiredHeader[path]; ok {
		reqCommon.Header = ctx.Get(header)
		reqCommon.RawBody = string(body)
	}
	//L.Print(ctx.OriginalURL())
	if err := ctx.QueryParser(in); L.IsError(err, `ctx.QueryParser failed: `+url) {
		return err
	}
	if len(body) > 0 {
		retry := true
		if body[0] == '{' || ctx.Get(`content-type`) == `application/json` {
			if err := json.Unmarshal(body, in); err == nil {
				retry = false
			}
		}
		// application/x-www-form-urlencoded
		// multipart/form-data
		if retry {
			if err := ctx.BodyParser(in); L.IsError(err, `ctx.BodyParser failed: `+url) {
				return err
			}
		}
		trimBody := S.Left(string(body), 4096)
		L.Print(trimBody)
		if reqCommon.Debug && reqCommon.RawBody == `` {
			reqCommon.RawBody = trimBody
		}
	}
	return nil
}

func (w *WebServer) Start() {
	fw := fiber.New()

	d := &domain.Domain{
		AuthOltp: w.AuthOltp,
		AuthOlap: w.AuthOlap,
	}

	// load svelte templates
	views = &Views{}
	views.LoadAll()

	// assign static routes (GET)
	WebStatic(fw, d)

	// API routes (POST)
	ApiRoutes(fw, d)

	log.Err(fw.Listen(w.Cfg.ListenAddr()))
}

type Views struct {
	cache map[string]*Z.TemplateChain
}

var views *Views

func (v *Views) LoadAll() {
	if v.cache == nil {
		v.cache = map[string]*Z.TemplateChain{}
	}
	debug := conf.IsDebug()
	const svelteDir = `svelte/`
	var err error
	for svelte, html := range viewList {
		v.cache[svelte], err = Z.ParseFile(debug, debug, svelteDir+html)
		L.PanicIf(err, `failed to parse `+html)
	}
}
