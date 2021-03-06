package ratelimit

import (
	"fmt"
	"time"

	envoyroute "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	"github.com/gogo/protobuf/types"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/plugins/ratelimit"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

const (
	ExtensionName      = "rate-limit"
	EnvoyExtensionName = "envoy-rate-limit"
	CustomDomain       = "custom"
	requestType        = "both"

	customStage    = 1
	DefaultTimeout = 100 * time.Millisecond

	FilterName = "envoy.rate_limit"
)

var (
	// rate limiting should happen after auth
	DefaultFilterStage = plugins.DuringStage(plugins.RateLimitStage)

	// we may want to rate limit before executing the AuthN and AuthZ stages
	// notably, AuthZ still needs to occur after AuthN
	BeforeAuthStage = plugins.BeforeStage(plugins.AuthNStage)
)

type Plugin struct {
	upstreamRef         *core.ResourceRef
	timeout             *time.Duration
	denyOnFail          bool
	rateLimitBeforeAuth bool
}

func NewPlugin() *Plugin {
	return &Plugin{}
}

func (p *Plugin) Init(params plugins.InitParams) error {
	if rlServer := params.Settings.GetRatelimitServer(); rlServer != nil {
		p.upstreamRef = rlServer.RatelimitServerRef
		p.timeout = rlServer.RequestTimeout
		p.denyOnFail = rlServer.DenyOnFail
		p.rateLimitBeforeAuth = rlServer.RateLimitBeforeAuth
	}

	return nil
}

func (p *Plugin) ProcessVirtualHost(params plugins.VirtualHostParams, in *v1.VirtualHost, out *envoyroute.VirtualHost) error {
	if rl := in.GetVirtualHostPlugins().GetRatelimit(); rl != nil {
		out.RateLimits = generateCustomEnvoyConfigForVhost(rl.RateLimits)
	}
	return nil
}

func (p *Plugin) ProcessRoute(params plugins.RouteParams, in *v1.Route, out *envoyroute.Route) error {
	var rateLimit *ratelimit.RateLimitRouteExtension
	if rl := in.GetRoutePlugins().GetRatelimit(); rl != nil {
		rateLimit = rl
	} else {
		// no rate limit route config found, nothing to do here
		return nil
	}

	ra := out.GetRoute()
	if ra != nil {
		ra.RateLimits = generateCustomEnvoyConfigForVhost(rateLimit.RateLimits)
		ra.IncludeVhRateLimits = &types.BoolValue{Value: rateLimit.IncludeVhRateLimits}
	} else {
		// TODO(yuval-k): maybe return nil here instead and just log a warning?
		return fmt.Errorf("cannot apply rate limits without a route action")
	}

	return nil
}

func (p *Plugin) HttpFilters(params plugins.Params, listener *v1.HttpListener) ([]plugins.StagedHttpFilter, error) {
	if p.upstreamRef == nil {
		return nil, nil
	}

	customConf := generateEnvoyConfigForCustomFilter(*p.upstreamRef, p.timeout, p.denyOnFail)

	filterStage := DefaultFilterStage
	if p.rateLimitBeforeAuth {
		filterStage = BeforeAuthStage
	}

	customStagedFilter, err := plugins.NewStagedFilterWithConfig(FilterName, customConf, filterStage)
	if err != nil {
		return nil, err
	}

	return []plugins.StagedHttpFilter{
		customStagedFilter,
	}, nil
}
