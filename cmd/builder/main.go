/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/emicklei/go-restful"
	"github.com/go-openapi/spec"
	"github.com/keleustes/armada-crd/pkg/generated"
	"k8s.io/kube-openapi/pkg/builder"
	"k8s.io/kube-openapi/pkg/common"
	"k8s.io/kube-openapi/pkg/util"
)

// TODO: Change this to output the generated swagger to stdout.
const defaultSwaggerFile = "generated.json"

func main() {
	// Get the name of the generated swagger file from the args
	// if it exists; otherwise use the default file name.
	swaggerFilename := defaultSwaggerFile
	if len(os.Args) > 1 {
		swaggerFilename = os.Args[1]
	}

	// Generate the definition names from the map keys returned
	// from GetOpenAPIDefinitions. Anonymous function returning empty
	// Ref is not used.
	var defNames []string
	for name, _ := range generated.GetOpenAPIDefinitions(func(name string) spec.Ref {
		return spec.Ref{}
	}) {
		defNames = append(defNames, name)
	}

	// Create a minimal builder config, then call the builder with the definition names.
	config := createOpenAPIBuilderConfig()
	config.GetDefinitions = generated.GetOpenAPIDefinitions
	// Build the Paths using a simple WebService for the final spec
	swagger, serr := builder.BuildOpenAPISpec(createWebServices(), config)
	if serr != nil {
		log.Fatalf("ERROR: %s", serr.Error())
	}

	// Marshal the swagger spec into JSON, then write it out.
	specBytes, err := json.MarshalIndent(swagger, " ", " ")
	if err != nil {
		log.Fatalf("json marshal error: %s", err.Error())
	}
	err = ioutil.WriteFile(swaggerFilename, specBytes, 0644)
	if err != nil {
		log.Fatalf("stdout write error: %s", err.Error())
	}
}

// CreateOpenAPIBuilderConfig hard-codes some values in the API builder
// config for testing.
func createOpenAPIBuilderConfig() *common.Config {
	return &common.Config{
		ProtocolList:   []string{"https"},
		IgnorePrefixes: []string{"/swaggerapi"},
		Info: &spec.Info{
			InfoProps: spec.InfoProps{
				Title:   "Armada Operator",
				Version: "1.0",
			},
		},
		ResponseDefinitions: map[string]spec.Response{
			"NotFound": spec.Response{
				ResponseProps: spec.ResponseProps{
					Description: "Entity not found.",
				},
			},
		},
		CommonResponses: map[int]spec.Response{
			401: *spec.NewResponse().WithDescription("Unauthorized"),
			// 404: *spec.ResponseRef("#/responses/NotFound"),
		},
		GetOperationIDAndTags: func(r *restful.Route) (string, []string, error) {
			return r.Operation, nil, nil
		},
		GetDefinitionName: func(name string) (string, spec.Extensions) {
			gvk := name[strings.LastIndex(name, "/")+1:]
			kind := gvk[strings.LastIndex(gvk, ".")+1:]
			version := gvk[:strings.LastIndex(gvk, ".")]
			log.Printf("kind %s version %s", kind, version)

			prefix := "org.airshipit.armada."
			if strings.Contains(name, "apimachinery") {
				prefix = "io.k8s.apimachinery.pkg.apis.meta."
			}
			friendlyName := prefix + gvk
			var extensions spec.Extensions
			switch kind {
			case "ArmadaChart", "ArmadaChartGroup", "ArmadaManifest":
				extensions = spec.Extensions{"x-kubernetes-group-version-kind": map[string]interface{}{
					"group":   "armada.airshipit.org",
					"kind":    kind,
					"version": version,
				}}
			default:
				extensions = spec.Extensions{}
			}
			return friendlyName, extensions
		},
	}
}

// createWebServices hard-codes a simple WebService which only defines a GET path
// for testing.
func createWebServices() []*restful.WebService {
	w := new(restful.WebService)
	w.Route(buildRouteForType(w, "v1alpha1", "ArmadaChart"))
	w.Route(buildRouteForType(w, "v1alpha1", "ArmadaChartGroup"))
	w.Route(buildRouteForType(w, "v1alpha1", "ArmadaManifest"))
	return []*restful.WebService{w}
}

// Implements OpenAPICanonicalTypeNamer
var _ = util.OpenAPICanonicalTypeNamer(&typeNamer{})

type typeNamer struct {
	pkg  string
	name string
}

func (t *typeNamer) OpenAPICanonicalTypeName() string {
	return fmt.Sprintf("github.com/keleustes/armada-crd/pkg/apis/armada/%s.%s", t.pkg, t.name)
}

func buildRouteForType(ws *restful.WebService, pkg, name string) *restful.RouteBuilder {
	namer := typeNamer{
		pkg:  pkg,
		name: name,
	}
	return ws.GET(fmt.Sprintf("/api/armada.airshipit.org/%s/namespaces/{namespace}/%ss/{name}", pkg, strings.ToLower(name))).
		Doc(fmt.Sprintf("read the status of the specified %s", name)).
		Operation(fmt.Sprintf("readArmada%s%s", pkg, name)).
		Param(ws.PathParameter("name", fmt.Sprintf("name of the %s", name)).DataType("string")).
		Param(ws.PathParameter("namespace", "object name and auth scope, such as for teams and projects").DataType("string")).
		Produces("application/json", "application/yaml", "application/vnd.kubernetes.protobuf").
		Consumes("*/*").
		To(func(*restful.Request, *restful.Response) {}).
		Writes(&namer)
}
