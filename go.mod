module street

go 1.25.5

require (
	github.com/ClickHouse/clickhouse-go/v2 v2.47.0
	github.com/allegro/bigcache/v3 v3.1.0
	github.com/davecgh/go-spew v1.1.1
	github.com/fatih/color v1.19.0
	github.com/gabriel-vasile/mimetype v1.4.2
	github.com/goccy/go-json v0.10.6
	github.com/gofiber/fiber/v2 v2.52.13
	github.com/golang-jwt/jwt/v5 v5.3.1
	github.com/hexops/autogold/v2 v2.1.0
	github.com/joho/godotenv v1.4.0
	github.com/jxskiss/base62 v1.1.0
	github.com/kokizzu/ch-timed-buffer v1.2025.1416
	github.com/kokizzu/gotro v1.3503.236
	github.com/kokizzu/id64 v1.2829.1452
	github.com/kokizzu/lexid v1.2423.1347
	github.com/kpango/fastime v1.1.10
	github.com/mailjet/mailjet-apiv3-go/v4 v4.0.1
	github.com/mitchellh/mapstructure v1.5.0
	github.com/mojura/enkodo v0.5.6
	github.com/ory/dockertest/v3 v3.12.0
	github.com/pierrec/lz4/v4 v4.1.27
	github.com/rs/zerolog v1.29.1
	github.com/segmentio/fasthash v1.0.3
	github.com/sendgrid/sendgrid-go v3.12.0+incompatible
	github.com/stretchr/testify v1.11.1
	github.com/tarantool/go-tarantool v1.12.2
	github.com/tidwall/gjson v1.17.0
	github.com/valyala/tsvreader v1.0.0
	github.com/vburenin/nsync v0.0.0-20160822015540-9a75d1c80410
	github.com/wneessen/go-mail v0.7.1
	github.com/xuri/excelize/v2 v2.11.0
	github.com/yosuke-furukawa/json5 v0.1.1
	github.com/zeebo/xxh3 v1.1.0
	golang.org/x/exp v0.0.0-20260611194520-c48552f49976
	golang.org/x/oauth2 v0.27.0
	golang.org/x/sync v0.21.0
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
)

require github.com/tiendc/go-deepcopy v1.7.2 // indirect

replace github.com/docker/cli => github.com/docker/cli v29.6.1+incompatible

require (
	cloud.google.com/go v0.67.0 // indirect
	dario.cat/mergo v1.0.2 // indirect
	github.com/Azure/go-ansiterm v0.0.0-20250102033503-faa5f7b0171c // indirect
	github.com/ClickHouse/ch-go v0.73.0 // indirect
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/Nvveen/Gotty v0.0.0-20120604004816-cd527374f1e5 // indirect
	github.com/OneOfOne/cmap v0.0.0-20170825200327-ccaef7657ab8 // indirect
	github.com/alitto/pond v1.9.2 // indirect
	github.com/allegro/bigcache v1.2.1 // indirect
	github.com/andybalholm/brotli v1.2.2 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/clipperhouse/uax29/v2 v2.7.0 // indirect
	github.com/containerd/continuity v0.4.5 // indirect
	github.com/containerd/errdefs v1.0.0 // indirect
	github.com/containerd/errdefs/pkg v0.3.0 // indirect
	github.com/distribution/reference v0.6.0 // indirect
	github.com/docker/cli v29.6.1+incompatible // indirect
	github.com/docker/go-connections v0.7.0 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-faster/city v1.0.1 // indirect
	github.com/go-faster/errors v0.7.1 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.5.0 // indirect
	github.com/goccy/go-yaml v1.19.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hexops/gotextdiff v1.0.3 // indirect
	github.com/hexops/valast v1.4.3 // indirect
	github.com/klauspost/compress v1.19.0 // indirect
	github.com/klauspost/cpuid/v2 v2.4.0 // indirect
	github.com/kokizzu/rand v0.0.0-20221021123447-6043c55a8bad // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mailjet/mailjet-apiv3-go/v3 v3.2.0 // indirect
	github.com/mattn/go-colorable v0.1.15 // indirect
	github.com/mattn/go-isatty v0.0.22 // indirect
	github.com/mattn/go-pointer v0.0.1 // indirect
	github.com/mattn/go-runewidth v0.0.21 // indirect
	github.com/moby/docker-image-spec v1.3.1 // indirect
	github.com/moby/moby/api v1.55.0 // indirect
	github.com/moby/moby/client v0.5.0 // indirect
	github.com/moby/sys/user v0.4.0 // indirect
	github.com/moby/term v0.5.2 // indirect
	github.com/nightlyone/lockfile v1.0.0 // indirect
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.1 // indirect
	github.com/opencontainers/runc v1.3.6 // indirect
	github.com/orcaman/concurrent-map v1.0.0 // indirect
	github.com/paulmach/orb v0.13.0 // indirect
	github.com/philhofer/fwd v1.1.3-0.20240916144458-20a13a1f6b7c // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/richardlehane/mscfb v1.0.7 // indirect
	github.com/richardlehane/msoleps v1.0.6 // indirect
	github.com/rogpeppe/go-internal v1.15.0 // indirect
	github.com/segmentio/asm v1.2.1 // indirect
	github.com/sendgrid/rest v2.6.9+incompatible // indirect
	github.com/shopspring/decimal v1.4.0 // indirect
	github.com/sirupsen/logrus v1.9.4 // indirect
	github.com/spacemonkeygo/spacelog v0.0.0-20180420211403-2296661a0572 // indirect
	github.com/tarantool/go-openssl v1.1.1 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/tinylib/msgp v1.2.5 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.72.0 // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	github.com/xuri/efp v0.0.1 // indirect
	github.com/xuri/nfp v0.0.2-0.20250530014748-2ddeb826f9a9 // indirect
	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.69.0 // indirect
	go.opentelemetry.io/otel v1.44.0 // indirect
	go.opentelemetry.io/otel/metric v1.44.0 // indirect
	go.opentelemetry.io/otel/trace v1.44.0 // indirect
	go.yaml.in/yaml/v3 v3.0.4 // indirect
	golang.org/x/crypto v0.53.0 // indirect
	golang.org/x/image v0.43.0
	golang.org/x/mod v0.37.0 // indirect
	golang.org/x/net v0.56.0 // indirect
	golang.org/x/sys v0.46.0 // indirect
	golang.org/x/text v0.39.0 // indirect
	golang.org/x/tools v0.47.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
	gopkg.in/vmihailenco/msgpack.v2 v2.9.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	mvdan.cc/gofumpt v0.4.0 // indirect
)

exclude (
	github.com/docker/docker v24.0.7+incompatible
	github.com/docker/docker v27.1.1+incompatible
)
