# JEB: We will have to put all those tools into a docker image
# to be allow CI/CD to rebuild
OPENAPI_GEN      := "k8s.io/kube-openapi/cmd/openapi-gen"

.PHONY: clean

all: crd-yaml openapi-gen swagger-gen kubeval-json standalone-json

clean:
	rm -f kubectl/*.yaml
	rm -f pkg/generated/openapi_generated.go
	rm -f swagger/golden.report
	rm -f swagger/swagger.json
	rm -f kubeval/*.json
	rm -f standalone/*.json

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
	# openapi2jsonschema -o kubeval -p https://raw.githubusercontent.com/keleustes/armada-crd/master/kubeval/ --expanded --kubernetes swagger/swagger.json
	openapi2jsonschema -o kubeval -p https://raw.githubusercontent.com/keleustes/armada-crd/master/kubeval/ --expanded swagger/swagger.json

standalone-json:
	# JEB: Note really sure if it will be usable
	mkdir -p standalone
	openapi2jsonschema -o standalone --stand-alone --expanded swagger/swagger.json

