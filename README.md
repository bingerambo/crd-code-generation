# crd-code-generation

Example repository for the blog post [Kubernetes Deep Dive: Code Generation for CustomResources](https://blog.openshift.com/kubernetes-deep-dive-code-generation-customresources/).

## Installation

```
export GOPATH=~/go
go get github.com/bingerambo/crd-code-generation
```

## Getting Started

First register the custom resource definition:

```
kubectl apply -f deployment/nodecache-crd.yaml
```

Then add an example of the `NodeCache` kind for Node(nodename:node1):

```
kubectl apply -f deployment/node1.yaml
```

Finally build and run the example:

```
cd ~/go/src/github.com/bingerambo/crd-code-generation/cmd/nodecache/
go build
./nodecache -kubeconfig ~/.kube/config
```
