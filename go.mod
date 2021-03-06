module github.com/talos-systems/sidero

go 1.14

require (
	github.com/evanphx/json-patch v4.9.0+incompatible
	github.com/ghodss/yaml v1.0.0
	github.com/go-logr/logr v0.2.1-0.20200730175230-ee2de8da5be6
	github.com/go-logr/zapr v0.2.0 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/hashicorp/go-multierror v1.1.0
	github.com/onsi/ginkgo v1.12.1
	github.com/onsi/gomega v1.10.1
	github.com/pensando/goipmi v0.0.0-20200303170213-e858ec1cf0b5
	github.com/pin/tftp v2.1.1-0.20200117065540-2f79be2dba4e+incompatible
	github.com/talos-systems/cluster-api-bootstrap-provider-talos v0.2.0-alpha.6
	github.com/talos-systems/cluster-api-control-plane-provider-talos v0.1.0-alpha.8
	github.com/talos-systems/go-blockdevice v0.1.1-0.20201022135759-bb8ac5d6a25e
	github.com/talos-systems/go-procfs v0.0.0-20200219015357-57c7311fdd45
	github.com/talos-systems/go-retry v0.1.1-0.20200922131245-752f081252cf
	github.com/talos-systems/go-smbios v0.0.0-20200807005123-80196199691e
	github.com/talos-systems/net v0.2.0
	github.com/talos-systems/talos v0.7.0-alpha.7
	github.com/talos-systems/talos/pkg/machinery v0.0.0-20201020161939-d2583e228288
	golang.org/x/sync v0.0.0-20201008141435-b3e1573b7520
	golang.org/x/sys v0.0.0-20201018230417-eeed37f84f13
	google.golang.org/grpc v1.29.1
	k8s.io/api v0.19.2
	k8s.io/apiextensions-apiserver v0.19.1
	k8s.io/apimachinery v0.19.2
	k8s.io/client-go v0.19.2
	k8s.io/utils v0.0.0-20200729134348-d5654de09c73
	sigs.k8s.io/cluster-api v0.3.9
	sigs.k8s.io/controller-runtime v0.6.3
)
