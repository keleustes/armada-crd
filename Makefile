# JEB: We will have to put all those tools into a docker image
# to be allow CI/CD to rebuild

KUSTOMIZE_NAME      := kustomize
PLUGINATOR_NAME     := pluginator

BINDIR              := bin
TOOLS_DIR           := tools
TOOLS_BIN_DIR       := $(TOOLS_DIR)/bin
CRD_ROOT            ?= $(MANIFEST_ROOT)/crd/bases

# Binaries.
CONTROLLER_GEN      := $(TOOLS_BIN_DIR)/controller-gen
GOLANGCI_LINT       := $(TOOLS_BIN_DIR)/golangci-lint
KUBEBUILDER         := $(TOOLS_BIN_DIR)/kubebuilder
OPENAPI_GEN         := $(TOOLS_BIN_DIR)/openapi-gen
KIND                := $(TOOLS_BIN_DIR)/kind
KUBEVAL             := $(TOOLS_BIN_DIR)/kubeval

# linting
LINTER_CMD          := $(GOLANGCI_LINT)

COVER_FILE=coverage.out

export GO111MODULE=on

all: clean generate openapi-gen swagger-gen kubeval-json

## --------------------------------------
## Tooling Binaries
## --------------------------------------

$(CONTROLLER_GEN): $(TOOLS_DIR)/go.mod # Build controller-gen from tools folder.
	cd $(TOOLS_DIR); go build -tags=tools -o $(TOOLS_BIN_DIR)/controller-gen sigs.k8s.io/controller-tools/cmd/controller-gen

$(GOLANGCI_LINT): $(TOOLS_DIR)/go.mod # Build golangci-lint from tools folder.
	cd $(TOOLS_DIR); go build -tags=tools -o $(TOOLS_BIN_DIR)/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint

$(KUBEBUILDER): $(TOOLS_DIR)/go.mod
	cd $(TOOLS_DIR); ./install_kubebuilder.sh

$(KIND): $(TOOLS_DIR)/go.mod # Build kind from tools folder.
	cd $(TOOLS_DIR); go build -tags=tools -o $(TOOLS_BIN_DIR)/kind sigs.k8s.io/kind

$(KUBEVAL): $(TOOLS_DIR)/go.mod # Build kubeval from tools folder.
	cd $(TOOLS_DIR); go build -tags=tools -o $(TOOLS_BIN_DIR)/kubeval github.com/instrumenta/kubeval

$(OPENAPI_GEN): $(TOOLS_DIR)/go.mod # Build openapi-gen from tools folder.
	cd $(TOOLS_DIR); go build -tags=tools -o $(TOOLS_BIN_DIR)/openapi-gen k8s.io/kube-openapi/cmd/openapi-gen

.PHONY: install-tools
# install-tools: $(CONTROLLER_GEN) $(GOLANGCI_LINT) $(KUBEBUILDER) $(KIND) $(KUBEVAL) $(OPENAPI_GEN)
install-tools: $(CONTROLLER_GEN) $(KUBEBUILDER) $(KIND)

## --------------------------------------
## Linting
## --------------------------------------

.PHONY: lint
lint: $(GOLANGCI_LINT) ## Lint codebase
	$(GOLANGCI_LINT) run -v

lint-full: $(GOLANGCI_LINT) ## Run slower linters to detect possible issues
	$(GOLANGCI_LINT) run -v --fast=false

# Run go fmt against code
fmt:
	go fmt ./cmd/... ./pkg/...

# Run go vet against code
vet:
	go vet ./cmd/... ./pkg/...


## --------------------------------------
## Generate
## --------------------------------------

.PHONY: modules
modules: ## Runs go mod to ensure proper vendoring.
	go mod tidy
	cd $(TOOLS_DIR); go mod tidy

.PHONY: generate
generate: ## Generate code
	$(MAKE) generate-go
	$(MAKE) generate-manifests

.PHONY: generate-go
generate-go: $(CONTROLLER_GEN)
	GO111MODULE=on $(CONTROLLER_GEN) object paths=./pkg/apis/armada/... output:object:dir=./pkg/apis/armada/v1alpha1 output:none
	GO111MODULE=on $(CONTROLLER_GEN) object paths=./pkg/apis/kubeflow/... output:object:dir=./pkg/apis/kubeflow/v1beta1 output:none
	GO111MODULE=on $(CONTROLLER_GEN) object paths=./pkg/apis/openstacklcm/... output:object:dir=./pkg/apis/openstacklcm/v1alpha1 output:none

.PHONY: generate-manifests
generate-manifests: $(CONTROLLER_GEN) ## Generate manifests e.g. CRD, RBAC etc.
	mkdir -p kubectl
	GO111MODULE=on $(CONTROLLER_GEN) crd paths=./pkg/apis/armada/... crd:trivialVersions=true output:crd:dir=./kubectl output:none
	GO111MODULE=on $(CONTROLLER_GEN) crd paths=./pkg/apis/kubeflow/... crd:trivialVersions=true output:crd:dir=./kubectl output:none
	GO111MODULE=on $(CONTROLLER_GEN) crd paths=./pkg/apis/openstacklcm/... crd:trivialVersions=true output:crd:dir=./kubectl output:none


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
.PHONY: crd-yaml
crd-yaml: $(CONTROLLER_GEN)

.PHONY: openapi-gen
openapi-gen: $(OPENAPI_GEN)
	mkdir -p $(HOME)/src/k8s.io/kube-openapi/boilerplate/
	touch $(HOME)/src/k8s.io/kube-openapi/boilerplate/boilerplate.go.txt
	mkdir -p pkg/generated
	mkdir -p swagger
	$(OPENAPI_GEN) -i "k8s.io/apimachinery/pkg/runtime,k8s.io/apimachinery/pkg/apis/meta/v1,github.com/keleustes/armada-crd/pkg/apis/armada/v1alpha1"   -o pkg   -p generated   -O openapi_generated   -r ./swagger/golden.report

.PHONY: swagger-gen
swagger-gen:
	mkdir -p swagger
	GO111MODULE=on go run cmd/builder/main.go swagger/swagger.json

.PHONY: kubeval-json
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

## --------------------------------------
## Testing
## --------------------------------------

.PHONY: test
test:
	echo "sudo systemctl stop kubelet"
	echo -e 'docker stop $$(docker ps -qa)'
	echo -e 'export PATH=$${PATH}:/usr/local/kubebuilder/bin'
	mkdir -p config/crds
	cp kubectl/* config/crds/
	GO111MODULE=on go test ./pkg/... -coverprofile=cover.out && go tool cover -html=cover.out
	rm -fr config/crds/

clusterexist=$(shell tools/bin/kind get clusters | grep armadacrd  | wc -l)
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
	$(KIND) create cluster --name armadacrd

.PHONY: delete-testcluster
delete-testcluster:
	$(KIND) delete cluster --name armadacrd
