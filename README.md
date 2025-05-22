# Rewrite ingress-nginx demo
This repository shows how to deploy two different pods that serve traffic at two different subpaths: /foo and /bar. We will instruct the ingress-nginx to rewrite the path to replace /foo with "" so that the backend service gets a clean rewritten URL (i.e. /foo/hello becomes /hello to the backend service).

> [!NOTE]
> Bonus points: we also set some custom headers!

# Deploy
This is a kustomize project, so you'll need `kubectl kustomize` in order to deploy this. It is also a demo, so it is assumed that it will be run under `kind`. 

The ingress-nginx is included, with a patch that makes it run at specific NodePort ports. So that we can map those two ports in `kind`.

## Requirements
- `kubectl`
- `kind`

## Create the `kind` cluster
```plaintext
❯ kind create cluster --config=./kind/kind-config.yaml
Creating cluster "ingress" ...
 ✓ Ensuring node image (kindest/node:v1.33.1) 🖼
 ✓ Preparing nodes 📦 📦
 ✓ Writing configuration 📜
 ✓ Starting control-plane 🕹️
 ✓ Installing CNI 🔌
 ✓ Installing StorageClass 💾
 ✓ Joining worker nodes 🚜
Set kubectl context to "kind-ingress"
You can now use your cluster with:

kubectl cluster-info --context kind-ingress

Not sure what to do next? 😅  Check out https://kind.sigs.k8s.io/docs/user/quick-start/
```

## Deploy!
```plaintext
❯ kubectl kustomize | kubectl apply -f-

namespace/ingress-nginx created
serviceaccount/ingress-nginx created
serviceaccount/ingress-nginx-admission created
role.rbac.authorization.k8s.io/ingress-nginx created
role.rbac.authorization.k8s.io/ingress-nginx-admission created
clusterrole.rbac.authorization.k8s.io/ingress-nginx created
clusterrole.rbac.authorization.k8s.io/ingress-nginx-admission created
rolebinding.rbac.authorization.k8s.io/ingress-nginx created
rolebinding.rbac.authorization.k8s.io/ingress-nginx-admission created
clusterrolebinding.rbac.authorization.k8s.io/ingress-nginx created
clusterrolebinding.rbac.authorization.k8s.io/ingress-nginx-admission created
configmap/custom-headers created
configmap/ingress-nginx-controller created
service/ingress-nginx-controller created
service/ingress-nginx-controller-admission created
service/bar-svc created
service/foo-svc created
deployment.apps/ingress-nginx-controller created
job.batch/ingress-nginx-admission-create created
job.batch/ingress-nginx-admission-patch created
ingress.networking.k8s.io/ingress created
ingressclass.networking.k8s.io/nginx created
pod/bar created
pod/foo created
validatingwebhookconfiguration.admissionregistration.k8s.io/ingress-nginx-admission created
```

## Profit!
Head to [https://127.0.0.1.nip.io/foo/test](https://127.0.0.1.nip.io/foo/test) and [https://127.0.0.1.nip.io/bar/test](https://127.0.0.1.nip.io/bar/test).

You will encounter a simple application telling you the service you're hitting and the URL that it is receiving.

> [!TIP] 
> You can see the source code in the `main.go` file.
