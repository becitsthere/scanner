module github.com/neuvector/scanner

go 1.22

replace (
	google.golang.org/grpc/security/advancedtls => github.com/holyspectral/grpc-go/security/advancedtls v0.0.0-20240202204003-24d7309846bd
	k8s.io/cri-api => k8s.io/cri-api v0.25.16
)

require (
	github.com/google/uuid v1.6.0
	github.com/jedib0t/go-pretty/v6 v6.4.6
	github.com/neuvector/neuvector v0.0.0-20240410171347-025fc1fa0041
	github.com/opencontainers/go-digest v1.0.0
	github.com/sirupsen/logrus v1.9.3
	google.golang.org/grpc v1.62.1
)

require (
	github.com/AdaLogics/go-fuzz-headers v0.0.0-20230811130428-ced1acdcaa24 // indirect
	github.com/AdamKorcz/go-118-fuzz-build v0.0.0-20230306123547-8075edf89bb0 // indirect
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/Microsoft/hcsshim v0.12.0-rc.1 // indirect
	github.com/armon/go-metrics v0.0.0-20180917152333-f0300d1749da // indirect
	github.com/aws/aws-sdk-go v1.42.36 // indirect
	github.com/codegangsta/inject v0.0.0-20150114235600-33e0aa1cb7c0 // indirect
	github.com/codeskyblue/go-sh v0.0.0-20200712050446-30169cf553fe // indirect
	github.com/containerd/cgroups/v3 v3.0.2 // indirect
	github.com/containerd/containerd v1.7.14 // indirect
	github.com/containerd/continuity v0.4.2 // indirect
	github.com/containerd/fifo v1.1.0 // indirect
	github.com/containerd/log v0.1.0 // indirect
	github.com/containerd/ttrpc v1.2.3 // indirect
	github.com/containerd/typeurl v1.0.3-0.20220422153119-7f6e6d160d67 // indirect
	github.com/containerd/typeurl/v2 v2.1.1 // indirect
	github.com/distribution/reference v0.5.0 // indirect
	github.com/docker/distribution v2.8.3+incompatible // indirect
	github.com/docker/docker v24.0.7+incompatible // indirect
	github.com/docker/go-events v0.0.0-20190806004212-e31b211e4f1c // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/docker/libtrust v0.0.0-20160708172513-aabc10ec26b7 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/go-logr/logr v1.3.0 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-immutable-radix v1.0.0 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/go-version v1.4.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/serf v0.9.7 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/klauspost/compress v1.16.5 // indirect
	github.com/knqyf263/go-rpmdb v0.0.0-20220209103220-0f7a6d951a6d // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/moby/locker v1.0.1 // indirect
	github.com/moby/sys/mountinfo v0.6.2 // indirect
	github.com/moby/sys/sequential v0.5.0 // indirect
	github.com/moby/sys/signal v0.7.0 // indirect
	github.com/moby/sys/user v0.1.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/opencontainers/image-spec v1.1.0-rc5 // indirect
	github.com/opencontainers/runtime-spec v1.1.1-0.20230922153023-c0e90434df2a // indirect
	github.com/opencontainers/selinux v1.11.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20200410134404-eec4a21b6bb0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/streadway/simpleuuid v0.0.0-20130420165545-6617b501e485 // indirect
	github.com/vishvananda/netlink v1.2.1-beta.2 // indirect
	github.com/vishvananda/netns v0.0.4 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.45.0 // indirect
	go.opentelemetry.io/otel v1.19.0 // indirect
	go.opentelemetry.io/otel/metric v1.19.0 // indirect
	go.opentelemetry.io/otel/trace v1.19.0 // indirect
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/net v0.23.0 // indirect
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/tools v0.16.1 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/genproto v0.0.0-20240123012728-ef4313101c80 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240318140521-94a12d6c2237 // indirect
	google.golang.org/grpc/security/advancedtls v0.0.0-20240408170116-ec257b4e1cec // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gotest.tools/v3 v3.5.1 // indirect
	k8s.io/api v0.29.3 // indirect
	k8s.io/apimachinery v0.29.3 // indirect
	k8s.io/cri-api v0.29.3 // indirect
	k8s.io/klog/v2 v2.110.1 // indirect
	k8s.io/utils v0.0.0-20231127182322-b307cd553661 // indirect
	lukechampine.com/uint128 v1.1.1 // indirect
	modernc.org/cc/v3 v3.35.22 // indirect
	modernc.org/ccgo/v3 v3.15.1 // indirect
	modernc.org/libc v1.14.1 // indirect
	modernc.org/mathutil v1.4.1 // indirect
	modernc.org/memory v1.0.5 // indirect
	modernc.org/opt v0.1.1 // indirect
	modernc.org/sqlite v1.14.5 // indirect
	modernc.org/strutil v1.1.1 // indirect
	modernc.org/token v1.0.0 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.4.1 // indirect
)
