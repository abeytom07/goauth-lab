# Go Sample Application

## Build

This uses the branch `honeycomb_only`

```sh
git clone https://github.com/abeytom07/goauth-lab.git
git checkout honeycomb_only

docker build -t local/go-lab-auth
```

## Kubernetes Manifests

After replacing the `HONEYCOMB_API_KEY` and `HONEYCOMB_DATASET`, save the contents into a file

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: auth-service
  labels:
    app: auth-service
spec:
  replicas: 1
  selector:
    matchLabels:
      app: auth-service
  template:
    metadata:
      labels:
        app: auth-service
    spec:
      containers:
      - name: auth-service
        image: localhost:32000/go-lab-auth
        ports:
        - containerPort: 8000
        imagePullPolicy: Always
        env:
          - name: HONEYCOMB_API_KEY
            value: "<honeycomb_api_key>"
          - name: HONEYCOMB_DATASET
            value: "<honeycomb_dataset>"
          - name: DATASERVICE_HOST
            value: data-service
          - name: PRICESERVICE_HOST
            value: price-service

---

apiVersion: v1
kind: Service
metadata:
  labels:
    app: auth-service
  name: auth-service
spec:
  type: NodePort
  ports:
  - name: "8000"
    port: 8000
    targetPort: 8000
    nodePort: 30000
  selector:
    app: auth-service
```

To deploy the manifest

```sh
kubectl -n $NAMESPACE apply -f /path/to/saved/file.yaml
```
