package presentation

import (
	"context"

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

	b := &domain.Domain{
		UserOltp: w.AuthOltp,
		UserOlap: w.AuthOlap,
	}

	// TODO: generate automatically
	debug := conf.IsDebug()
	const svelteDir = `svelte/`
	userRegisterUi, err := Z.ParseFile(debug, debug, svelteDir+`index.html`)
	L.PanicIf(err, `failed to parse index.html`)

	fw.Get("/", func(c *fiber.Ctx) error {
		// TODO: fetch from session
		c.Set("Content-Type", "text/html")
		return c.SendString(userRegisterUi.Str(M.SX{
			// TODO: put things to be rendered
			`title`: `Street`,
		}))
	})

	// TODO: inject routes here with codegen
	fw.Post("/"+domain.UserRegisterAction, func(c *fiber.Ctx) error {
		ctx := context.Background() // TODO: use tracer
		in := domain.UserRegisterIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserRegisterAction); err != nil {
			return err
		}
		in.FromFiberCtx(c, ctx)
		out := b.UserRegister(&in)
		out.DecorateSession(c, &in.RequestCommon, &in)
		return in.ToFiberCtx(c, out)
	})

	log.Err(fw.Listen(w.Cfg.ListenAddr()))
}
