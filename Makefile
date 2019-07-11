
# Image URL to use all building/pushing image targets
COMPONENT        ?= armada-operator
VERSION_V2       ?= 2.14.1
VERSION_V3       ?= 3.0.0
DHUBREPO         ?= keleustes/${COMPONENT}-dev
DOCKER_NAMESPACE ?= keleustes
IMG_V2           ?= ${DHUBREPO}:v${VERSION_V2}
IMG_V3           ?= ${DHUBREPO}:v${VERSION_V3}

OPENAPI_GEN      := "k8s.io/kube-openapi/cmd/openapi-gen"

clean:
	rm -fr vendor
	rm -fr cover.out
	rm -fr build/_output
	rm -fr config/crds

# Generate code
crd-yaml:
        # go get sigs.k8s.io/controller-tools
	GO111MODULE=on controller-gen crd paths=github.com/keleustes/armada-operator/pkg/apis/armada/... crd:trivialVersions=true output:crd:dir=./kubectl output:none

openapi-gen:
	# go get k8s.io/kube-openapi
	GO111MODULE=on go run ${OPENAPI_GEN} -i "k8s.io/apimachinery/pkg/apis/meta/v1,github.com/keleustes/armada-operator/pkg/apis/armada/v1alpha1"   -o pkg   -p generated   -O openapi_generated   -r ./swagger/golden.report

swagger-gen:
	GO111MODULE=on go run cmd/builder/main.go swagger/swagger.json

kubeval-json:
	REPO="keleustes/armada-operator"
	schema=swagger/golden.json
	prefix=https://raw.githubusercontent.com/${REPO}/master/${version}/
	openapi2jsonschema -o "${version}" --prefix "${prefix}" "${schema}"
