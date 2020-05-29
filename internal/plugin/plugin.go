package plugin

import (
	"github.com/hashicorp/waypoint/sdk"
	"github.com/hashicorp/waypoint/sdk/component"
	"github.com/hashicorp/waypoint/sdk/internal-shared/mapper"

	"github.com/hashicorp/waypoint/builtin/azure/aci"
	"github.com/hashicorp/waypoint/builtin/docker"
	"github.com/hashicorp/waypoint/builtin/google/cloudrun"
	"github.com/hashicorp/waypoint/builtin/k8s"
	"github.com/hashicorp/waypoint/builtin/lambda"
	"github.com/hashicorp/waypoint/builtin/pack"
)

var (
	Builders   = mustFactory(mapper.NewFactory((*component.Builder)(nil)))
	Registries = mustFactory(mapper.NewFactory((*component.Registry)(nil)))
	Platforms  = mustFactory(mapper.NewFactory((*component.Platform)(nil)))
	Releasers  = mustFactory(mapper.NewFactory((*component.ReleaseManager)(nil)))

	// Builtins is the map of all available builtin plugins and their
	// options for launching them.
	Builtins = map[string][]sdk.Option{
		"pack":             pack.Options,
		"docker":           docker.Options,
		"google-cloud-run": cloudrun.Options,
		"azure-aci":        aci.Options,
		"kubernetes":       k8s.Options,
		"lambda":           lambda.Options,
	}
)

func init() {
	Builders.Register("docker", BuiltinFactory("docker", component.BuilderType))
	Builders.Register("pack", BuiltinFactory("pack", component.BuilderType))
	Builders.Register("lambda", BuiltinFactory("lambda", component.BuilderType))

	Registries.Register("docker", BuiltinFactory("docker", component.RegistryType))

	Platforms.Register("google-cloud-run", BuiltinFactory("google-cloud-run", component.PlatformType))
	Platforms.Register("kubernetes", BuiltinFactory("kubernetes", component.PlatformType))
	Platforms.Register("lambda", BuiltinFactory("lambda", component.PlatformType))
	Platforms.Register("azure-aci", BuiltinFactory("azure-aci", component.PlatformType))

	Releasers.Register("google-cloud-run", BuiltinFactory("google-cloud-run", component.ReleaseManagerType))
	Releasers.Register("azure-aci", BuiltinFactory("azure-aci", component.ReleaseManagerType))
	Releasers.Register("kubernetes", BuiltinFactory("kubernetes", component.ReleaseManagerType))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func mustFactory(f *mapper.Factory, err error) *mapper.Factory {
	must(err)
	return f
}
