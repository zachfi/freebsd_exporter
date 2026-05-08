package tracing

import (
	"flag"

	"github.com/zachfi/zkit/pkg/util"
)

type Config struct {
	OtelEndpoint string `yaml:"otel_endpoint"`
	OrgID        string `yaml:"org_id"`
}

func (c *Config) RegisterFlagsAndApplyDefaults(prefix string, f *flag.FlagSet) {
	f.StringVar(&c.OtelEndpoint, util.PrefixConfig(prefix, "otel.endpoint"), "", "otel endpoint, eg: tempo:4317")
	f.StringVar(&c.OrgID, util.PrefixConfig(prefix, "org.id"), "", "org ID to use when sending traces")
}
