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
KUBEBUILDER         := $(TOOLS_BIN_DIR)/kubebuilder
OPENAPI_GEN         := $(TOOLS_BIN_DIR)/openapi-gen
# Taken from PATH — see the note by the tooling recipes below.
GOLANGCI_LINT       := golangci-lint
KIND                := kind
KUBEVAL             := kubeval

# linting
LINTER_CMD          := $(GOLANGCI_LINT)

COVER_FILE=coverage.out

export GO111MODULE=on

all: clean generate openapi-gen swagger-gen kubeval-json

## --------------------------------------
## Tooling Binaries
## --------------------------------------

# NOTE: these recipes `cd $(TOOLS_DIR)` first, so the -o path must be relative
# to tools/ (bin/x), NOT $(TOOLS_BIN_DIR) (tools/bin/x) — the latter nests the
# binary at tools/tools/bin/x, where nothing looks for it.

$(CONTROLLER_GEN): $(TOOLS_DIR)/go.mod # Build controller-gen from tools folder.
	cd $(TOOLS_DIR); go build -tags=tools -o bin/controller-gen sigs.k8s.io/controller-tools/cmd/controller-gen

$(KUBEBUILDER): $(TOOLS_DIR)/go.mod
	cd $(TOOLS_DIR); ./install_kubebuilder.sh

$(OPENAPI_GEN): $(TOOLS_DIR)/go.mod # Build openapi-gen from tools folder.
	cd $(TOOLS_DIR); go build -tags=tools -o bin/openapi-gen k8s.io/kube-openapi/cmd/openapi-gen

# golangci-lint, kind and kubeval are no longer pinned in tools/tools.go: their
# transitive deps pull the legacy google.golang.org/genproto monolith, which
# collides with the split genproto/googleapis/{api,rpc} modules that
# controller-tools v0.21.0 needs ("ambiguous import" on `go mod tidy`). Use the
# system binaries; CI pins golangci-lint in .github/workflows/ci.yml.

.PHONY: install-tools
install-tools: $(CONTROLLER_GEN) $(OPENAPI_GEN)

## --------------------------------------
## Linting
## --------------------------------------

.PHONY: lint
lint: ## Lint codebase (system golangci-lint; CI pins the version)
	$(LINTER_CMD) run -v

lint-full: ## Run with the repeat caps lifted — golangci truncates repeats by default
	$(LINTER_CMD) run -v --max-same-issues 0 --max-issues-per-linter 0

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
	# crd:trivialVersions was dropped in controller-gen v0.6; v0.21.0 emits
	# apiextensions.k8s.io/v1 CRDs (v1beta1 was removed in Kubernetes 1.22).
	GO111MODULE=on $(CONTROLLER_GEN) crd paths=./pkg/apis/armada/... output:crd:dir=./kubectl output:none
	GO111MODULE=on $(CONTROLLER_GEN) crd paths=./pkg/apis/kubeflow/... output:crd:dir=./kubectl output:none
	GO111MODULE=on $(CONTROLLER_GEN) crd paths=./pkg/apis/openstacklcm/... output:crd:dir=./kubectl output:none


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

# The envtest suite reads its CRDs straight from kubectl/ now, so this no longer
# stages a throwaway config/crds copy — which meant the tests only ever passed
# via `make test` and always failed under a plain `go test ./...` (what CI runs).
# Set KUBEBUILDER_ASSETS if the control-plane binaries are not already on PATH:
#   go install sigs.k8s.io/controller-runtime/tools/setup-envtest@latest
#   export KUBEBUILDER_ASSETS="$$(setup-envtest use -p path)"
.PHONY: test
test:
	GO111MODULE=on go test ./pkg/... -coverprofile=cover.out && go tool cover -html=cover.out

# Evaluated on every make invocation, so it must stay quiet when kind is absent.
clusterexist=$(shell $(KIND) get clusters 2>/dev/null | grep -c armadacrd)
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
