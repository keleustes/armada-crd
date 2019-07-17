# JEB: We will have to put all those tools into a docker image
# to be allow CI/CD to rebuild
OPENAPI_GEN      := "k8s.io/kube-openapi/cmd/openapi-gen"

.PHONY: clean

all: clean crd-yaml openapi-gen swagger-gen kubeval-json

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
        # go get sigs.k8s.io/controller-tools
	mkdir -p kubectl
	GO111MODULE=on controller-gen crd paths=github.com/keleustes/armada-operator/pkg/apis/armada/... crd:trivialVersions=true output:crd:dir=./kubectl output:none

openapi-gen:
	# go get k8s.io/kube-openapi
	mkdir -p pkg/generated
	mkdir -p swagger
	GO111MODULE=on go run ${OPENAPI_GEN} -i "k8s.io/apimachinery/pkg/apis/meta/v1,github.com/keleustes/armada-operator/pkg/apis/armada/v1alpha1"   -o pkg   -p generated   -O openapi_generated   -r ./swagger/golden.report

swagger-gen:
	mkdir -p swagger
	GO111MODULE=on go run cmd/builder/main.go swagger/swagger.json

kubeval-json:
	# JEB: Kubernetes option would be important but it does not work
	mkdir -p kubeval
	mkdir -p kubeval/master
	mkdir -p kubeval/master-local
	mkdir -p kubeval/master-standalone
	mkdir -p kubeval/master-standalone-strict
	openapi2jsonschema -o kubeval/master-standalone-strict --stand-alone --expanded --strict swagger/swagger.json
	openapi2jsonschema -o kubeval/master-standalone --stand-alone --expanded swagger/swagger.json
	openapi2jsonschema -o kubeval/master-local --expanded swagger/swagger.json
	openapi2jsonschema -o kubeval/master --expanded --prefix https://raw.githubusercontent.com/keleustes/armada-crd/master/kubeval/master/_definitions.json swagger/swagger.json


