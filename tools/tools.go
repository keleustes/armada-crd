//go:build tools

package tools

// These imports pin the code generators used by the Makefile so that
// `go build -tags=tools` in this directory produces the exact versions
// recorded in tools/go.mod.
//
// Deliberately limited to the generators. golangci-lint, kind and kubeval used
// to be pinned here too, but their transitive deps drag in the legacy
// google.golang.org/genproto monolith, which collides with the split
// genproto/googleapis/{api,rpc} modules that current controller-tools requires
// — `go mod tidy` fails with "ambiguous import". CI and local runs use the
// system golangci-lint instead (the CI pin lives in .github/workflows/ci.yml).
import (
	_ "k8s.io/kube-openapi/cmd/openapi-gen"
	_ "sigs.k8s.io/controller-tools/cmd/controller-gen"
)
