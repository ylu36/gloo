// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"fmt"

	github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes "github.com/solo-io/solo-kit/pkg/api/v1/resources/common/kubernetes"

	"github.com/solo-io/go-utils/hashutils"
	"go.uber.org/zap"
)

type DiscoverySnapshot struct {
	Upstreams      UpstreamList
	Kubenamespaces github_com_solo_io_solo_kit_pkg_api_v1_resources_common_kubernetes.KubeNamespaceList
	Secrets        SecretList
}

func (s DiscoverySnapshot) Clone() DiscoverySnapshot {
	return DiscoverySnapshot{
		Upstreams:      s.Upstreams.Clone(),
		Kubenamespaces: s.Kubenamespaces.Clone(),
		Secrets:        s.Secrets.Clone(),
	}
}

func (s DiscoverySnapshot) Hash() uint64 {
	return hashutils.HashAll(
		s.hashUpstreams(),
		s.hashKubenamespaces(),
		s.hashSecrets(),
	)
}

func (s DiscoverySnapshot) hashUpstreams() uint64 {
	return hashutils.HashAll(s.Upstreams.AsInterfaces()...)
}

func (s DiscoverySnapshot) hashKubenamespaces() uint64 {
	return hashutils.HashAll(s.Kubenamespaces.AsInterfaces()...)
}

func (s DiscoverySnapshot) hashSecrets() uint64 {
	return hashutils.HashAll(s.Secrets.AsInterfaces()...)
}

func (s DiscoverySnapshot) HashFields() []zap.Field {
	var fields []zap.Field
	fields = append(fields, zap.Uint64("upstreams", s.hashUpstreams()))
	fields = append(fields, zap.Uint64("kubenamespaces", s.hashKubenamespaces()))
	fields = append(fields, zap.Uint64("secrets", s.hashSecrets()))

	return append(fields, zap.Uint64("snapshotHash", s.Hash()))
}

type DiscoverySnapshotStringer struct {
	Version        uint64
	Upstreams      []string
	Kubenamespaces []string
	Secrets        []string
}

func (ss DiscoverySnapshotStringer) String() string {
	s := fmt.Sprintf("DiscoverySnapshot %v\n", ss.Version)

	s += fmt.Sprintf("  Upstreams %v\n", len(ss.Upstreams))
	for _, name := range ss.Upstreams {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Kubenamespaces %v\n", len(ss.Kubenamespaces))
	for _, name := range ss.Kubenamespaces {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Secrets %v\n", len(ss.Secrets))
	for _, name := range ss.Secrets {
		s += fmt.Sprintf("    %v\n", name)
	}

	return s
}

func (s DiscoverySnapshot) Stringer() DiscoverySnapshotStringer {
	return DiscoverySnapshotStringer{
		Version:        s.Hash(),
		Upstreams:      s.Upstreams.NamespacesDotNames(),
		Kubenamespaces: s.Kubenamespaces.Names(),
		Secrets:        s.Secrets.NamespacesDotNames(),
	}
}
