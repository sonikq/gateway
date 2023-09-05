# gateway.

### Примерные креды
>$REGISTRY_USER = user
>$REGISTRY_PASSWORD = *****

### На локале
```
docker login -u $REGISTRY_USER -p $REGISTRY_PASSWORD https://registry.geogracom.com
docker build -t registry.geogracom.com/skdf/skdf-abac-go:latest .
docker push registry.geogracom.com/skdf/skdf-abac-go:latest
```

### На 45ом хосте

```
docker login -u $REGISTRY_USER -p $REGISTRY_PASSWORD https://registry.geogracom.com
docker pull registry.geogracom.com/skdf/skdf-excel-abac-go:latest
docker container rm -f skdf-excel-abac-go || true
docker run -d -p 3002:3002 --name skdf-excel-abac-go registry.geogracom.com/skdf/skdf-abac-go:latest
``````