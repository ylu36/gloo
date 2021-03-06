package kubernetes_test

import (
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	gloov1kube "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/kubernetes"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/solo-io/gloo/projects/gloo/pkg/plugins/kubernetes"
)

var _ = Describe("Uds", func() {

	It("should preseve ssl config when updating upstreams", func() {
		desired := &gloov1.Upstream{
			UpstreamSpec: &gloov1.UpstreamSpec{
				UpstreamType: &gloov1.UpstreamSpec_Kube{
					Kube: &gloov1kube.UpstreamSpec{
						ServiceName: "test",
					},
				},
			},
		}
		original := &gloov1.Upstream{
			UpstreamSpec: &gloov1.UpstreamSpec{
				UpstreamType: &gloov1.UpstreamSpec_Kube{
					Kube: &gloov1kube.UpstreamSpec{},
				},
				SslConfig: &gloov1.UpstreamSslConfig{Sni: "testsni"},
			},
		}
		updated, err := UpdateUpstream(original, desired)
		Expect(err).NotTo(HaveOccurred())
		Expect(updated).To(BeTrue())
		Expect(desired.UpstreamSpec.SslConfig).To(BeIdenticalTo(original.UpstreamSpec.SslConfig))
	})

	It("should update ssl config when one is desired", func() {
		desiredSslConfig := &gloov1.UpstreamSslConfig{
			SslSecrets: &gloov1.UpstreamSslConfig_SecretRef{
				SecretRef: &core.ResourceRef{"hi", "there"},
			},
		}
		desired := &gloov1.Upstream{
			UpstreamSpec: &gloov1.UpstreamSpec{
				UpstreamType: &gloov1.UpstreamSpec_Kube{
					Kube: &gloov1kube.UpstreamSpec{
						ServiceName: "test",
					},
				},
				SslConfig: desiredSslConfig,
			},
		}
		original := &gloov1.Upstream{
			UpstreamSpec: &gloov1.UpstreamSpec{
				UpstreamType: &gloov1.UpstreamSpec_Kube{
					Kube: &gloov1kube.UpstreamSpec{},
				},
				SslConfig: &gloov1.UpstreamSslConfig{Sni: "testsni"},
			},
		}
		updated, err := UpdateUpstream(original, desired)
		Expect(err).NotTo(HaveOccurred())
		Expect(updated).To(BeTrue())
		Expect(desired.UpstreamSpec.SslConfig).To(BeIdenticalTo(desiredSslConfig))
	})

})
