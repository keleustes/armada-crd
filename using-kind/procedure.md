# Procedure followed to retrieve API info

## Using treasuremap


```bash
make create-testcluster
export KUBECONFIG="$(kind get kubeconfig-path --name="treasuremap")"
kubectl cluster-info
make install
make install-operators
kubectl proxy &
curl http://localhost:8001/openapi/v2 | python3 -m json.tool > api.15.0.json
```

