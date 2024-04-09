module github.com/neuvector/scanner

go 1.14

replace (
	github.com/containerd/containerd => github.com/containerd/containerd v1.3.10
	github.com/containerd/cri => github.com/containerd/cri v1.19.0
	github.com/cri-o/cri-o => github.com/cri-o/cri-o v1.15.4
	github.com/docker/distribution => github.com/docker/distribution v2.8.0-beta.1+incompatible
	github.com/kubernetes/cri-api => k8s.io/cri-api v0.22.3
	github.com/opencontainers/runc => github.com/opencontainers/runc v1.0.0-rc5
	golang.org/x/net => golang.org/x/net v0.0.0-20170421174939-0b588ed7a0cd
	golang.org/x/sys => golang.org/x/sys v0.0.0-20220209214540-3681064d5158
	google.golang.org/grpc => google.golang.org/grpc v1.30.1
	k8s.io/api => k8s.io/api v0.20.15
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.20.15
	k8s.io/apimachinery => k8s.io/apimachinery v0.17.17
	k8s.io/apiserver => k8s.io/apiserver v0.20.15
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.20.15
	k8s.io/client-go => k8s.io/client-go v0.20.15
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.20.15
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.20.15
	k8s.io/code-generator => k8s.io/code-generator v0.20.15
	k8s.io/component-base => k8s.io/component-base v0.20.15
	k8s.io/component-helpers => k8s.io/component-helpers v0.20.14
	k8s.io/controller-manager => k8s.io/controller-manager v0.20.14
	k8s.io/cri-api => k8s.io/cri-api v0.20.15
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.20.15
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.20.15
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.20.15
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.20.15
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.20.15
	k8s.io/kubectl => k8s.io/kubectl v0.20.15
	k8s.io/kubelet => k8s.io/kubelet v0.20.15
	k8s.io/kubernetes => k8s.io/kubernetes v1.23.1
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.20.15
	k8s.io/metrics => k8s.io/metrics v0.20.15
	k8s.io/mount-utils => k8s.io/mount-utils v0.20.14
	k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.22.5
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.20.15
)

require (
	github.com/aws/aws-sdk-go v1.42.36 // indirect
	github.com/docker/docker v20.10.12+incompatible // indirect
	github.com/google/uuid v1.3.0
	github.com/jedib0t/go-pretty/v6 v6.4.6
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/neuvector/neuvector v0.0.0-20240409044813-8e34c951bba0
	github.com/opencontainers/go-digest v1.0.0
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.8.2 // indirect
	google.golang.org/grpc v1.40.0
)
