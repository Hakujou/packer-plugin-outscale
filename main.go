package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/packer-plugin-outscale/builder/osc/bsu"
	"github.com/hashicorp/packer-plugin-outscale/builder/osc/bsusurrogate"
	"github.com/hashicorp/packer-plugin-outscale/builder/osc/bsuvolume"
	"github.com/hashicorp/packer-plugin-outscale/builder/osc/chroot"
	"github.com/hashicorp/packer-plugin-sdk/plugin"
	"github.com/hashicorp/packer-plugin-sdk/version"
)

var (
	// Version is the main version number that is being run at the moment.
	Version = "0.0.2"

	// VersionPrerelease is A pre-release marker for the Version. If this is ""
	// (empty string) then it means that it is a final release. Otherwise, this
	// is a pre-release such as "dev" (in development), "beta", "rc1", etc.
	VersionPrerelease = "dev"

	// PluginVersion is used by the plugin set to allow Packer to recognize
	// what version this plugin is.
	PluginVersion = version.InitializePluginVersion(Version, VersionPrerelease)
)

func main() {
	pps := plugin.NewSet()
	pps.SetVersion(PluginVersion)
	pps.RegisterBuilder("bsu", new(bsu.Builder))
	pps.RegisterBuilder("chroot", new(chroot.Builder))
	pps.RegisterBuilder("bsusurrogate", new(bsusurrogate.Builder))
	pps.RegisterBuilder("bsuvolume", new(bsuvolume.Builder))
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
