// Code generated by generate-artifacts. DO NOT EDIT.
// +build release,!new

package neco

var CurrentArtifacts = ArtifactSet{
	Images: []ContainerImage{
		{Name: "cke", Repository: "quay.io/cybozu/cke", Tag: "0.28-1"},
		{Name: "etcd", Repository: "quay.io/cybozu/etcd", Tag: "3.3.10-1"},
		{Name: "omsa", Repository: "quay.io/cybozu/omsa", Tag: "18.11.01-1"},
		{Name: "sabakan", Repository: "quay.io/cybozu/sabakan", Tag: "0.31-1"},
		{Name: "serf", Repository: "quay.io/cybozu/serf", Tag: "0.8.1-5"},
		{Name: "vault", Repository: "quay.io/cybozu/vault", Tag: "1.0.0-1"},
		{Name: "hyperkube", Repository: "quay.io/cybozu/hyperkube", Tag: "1.13.0-1"},
		{Name: "coil", Repository: "quay.io/cybozu/coil", Tag: "0.4-1"},
		{Name: "squid", Repository: "quay.io/cybozu/squid", Tag: "3.5.27-1-3"},
	},
	Debs: []DebianPackage{
		{Name: "etcdpasswd", Owner: "cybozu-go", Repository: "etcdpasswd", Release: "v0.7"},
		{Name: "neco", Owner: "cybozu-go", Repository: "neco", Release: "release-2019.01.17-1"},
	},
	CoreOS: CoreOSImage{Channel: "stable", Version: "1967.3.0"},
}