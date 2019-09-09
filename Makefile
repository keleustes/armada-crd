# JEB: We will have to put all those tools into a docker image
# to be allow CI/CD to rebuild
OPENAPI_GEN      := "k8s.io/kube-openapi/cmd/openapi-gen"


all: clean crd-yaml openapi-gen swagger-gen kubeval-json

.PHONY: install-tools
install-tools:
	cd /tmp && GO111MODULE=on go get sigs.k8s.io/kind@v0.5.0
	cd /tmp && GO111MODULE=on go get github.com/instrumenta/kubeval@0.13.0

clusterexist=$(shell kind get clusters | grep armadacrd  | wc -l)
ifeq ($(clusterexist), 1)
  testcluster=$(shell kind get kubeconfig-path --name="armadacrd")
  SETKUBECONFIG=KUBECONFIG=$(testcluster)
else
  SETKUBECONFIG=
endif

.PHONY: which-cluster
which-cluster:
	echo $(SETKUBECONFIG)

.PHONY: create-testcluster
create-testcluster:
	kind create cluster --name armadacrd

.PHONY: delete-testcluster
delete-testcluster:
	kind delete cluster --name armadacrd

.PHONY: clean
clean:
	rm -f kubectl/*.yaml
	rm -f pkg/generated/openapi_generated.go
	rm -f swagger/golden.report
	rm -f swagger/swagger.json
	rm -f kubeval/master/*.json
	rm -f kubeval/master-local/*.json
	rm -f kubeval/master-standalone/*.json
	rm -f kubeval/master-standalone-strict/*.json

# Generate code
crd-yaml:
	# Installation seems to fail
        # GO111MODULE=on go get -u sigs.k8s.io/controller-tools/cmd/controller-gen
	# This installation seems to work
	# cd $HOME/src/sigs.k8s.io/
 	# git clone https://github.com/kubernetes-sigs/controller-tools.git
 	# cd controller-tools/
	# GO111MODULE=on go build -o $HOME/bin/controller-gen cmd/controller-gen/main.go
	mkdir -p kubectl
	GO111MODULE=on controller-gen crd paths=./pkg/apis/armada/... crd:trivialVersions=true output:crd:dir=./kubectl output:none
	GO111MODULE=on controller-gen object paths=./pkg/apis/armada/... output:object:dir=./pkg/apis/armada/v1alpha1 output:none

openapi-gen:
	# GO111MODULE=on go get -u k8s.io/kube-openapi/cmd/openapi-gen
	# mkdir -p $HOME/src/k8s.io/kube-openapi/boilerplate/
	# touch -p $HOME/src/k8s.io/kube-openapi/boilerplate/boilerplate.go.txt
	mkdir -p pkg/generated
	mkdir -p swagger
	GO111MODULE=on go run ${OPENAPI_GEN} -i "k8s.io/apimachinery/pkg/apis/meta/v1,github.com/keleustes/armada-crd/pkg/apis/armada/v1alpha1"   -o pkg   -p generated   -O openapi_generated   -r ./swagger/golden.report

swagger-gen:
	mkdir -p swagger
	GO111MODULE=on go run cmd/builder/main.go swagger/swagger.json

kubeval-json:
	# JEB: Kubernetes option would be important but it does not work
	# GO111MODULE=on go get -u github.com/instrumenta/kubeval
	# sudo -i
	# pip install openapi2jsonschema
	mkdir -p kubeval
	mkdir -p kubeval/master
	mkdir -p kubeval/master-local
	mkdir -p kubeval/master-standalone
	mkdir -p kubeval/master-standalone-strict
	openapi2jsonschema -o kubeval/master-standalone-strict --stand-alone --expanded --strict --kubernetes swagger/swagger.json
	openapi2jsonschema -o kubeval/master-standalone --stand-alone --expanded --kubernetes swagger/swagger.json
	openapi2jsonschema -o kubeval/master-local --expanded --kubernetes swagger/swagger.json
	openapi2jsonschema -o kubeval/master --expanded --kubernetes --prefix https://raw.githubusercontent.com/keleustes/armada-crd/master/kubeval/master/_definitions.json swagger/swagger.json


