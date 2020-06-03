# Challenge

## Requirements

To deploy this stack you will need:

* `kubectl` installed 
* `kustomize` installed 
* Access to a Kubernetes cluster

## Architecture

* Web Proxy: I decided to use Nginx Ingress solution, due the configuration of Nginx Ingress in every major cloud, it should generate a LoadBalancer with an external IP ready for use
* Web App: A golang app with http module 

## Deployment

I've generated a simple script which with Kustomize generates de config and apply it in the current cluster configured

We can pass to the scripts parameters like environment ($1) and tag ($2), so it can be easily adapted to more config

i.e
```bash
build/build.sh dev latest  
```

This should generate a manifest with all configuration needed:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: api
  namespace: demo
spec:
  ports:
  - name: http
    port: 80
    protocol: TCP
    targetPort: 8080
  selector:
    app: api
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: api
  name: api
  namespace: demo
spec:
  replicas: 1
  selector:
    matchLabels:
      app: api
  strategy:
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 1
    type: RollingUpdate
  template:
    metadata:
      labels:
        app: api
    spec:
      containers:
      - image: vmerino/go-app:latest
        imagePullPolicy: Always
        name: api-go
        ports:
        - containerPort: 8080
---
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: demo-ingress
spec:
  rules:
  - http:
      paths:
      - backend:
          serviceName: api
          servicePort: 80
        path: /

```

After everything configured you should be able to access via:

```bash
$ curl -L http://IP_ADDRESS/hello
{ "Hello": "world" }
```


## Additional notes

I've configured a trigger in docker-hub which generates the image used in this exercise `vmerino/go-app` in order to simulate build process

## Motivation

I've decided to go with Kubernetes since I've been working with K8s for some time ago, also the challenge somehow remind me as a Microservice architecture, so it can be easily scalable with adding some extra configuration.

Choosing the solution for deploy option, wasn't easy, I wanted to avoid extra manual configuration (in the end we are using a bash script, but its Kustomize who generates the manifest). My two options here were Helm or Kustomize, I decided to go with the latest because I've been working with it since some months ago and I really like it.
 

