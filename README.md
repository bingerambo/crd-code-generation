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

Example log print info:

```
========================================
ADD NodeCache ... 
setNodeCache: {Name:node1 Datasets:4e1d6a39e550ee41&5&0;8e1d6a39e550ee4x&5&1 FreeSize:100 AllocatableSize:380} 
========================================
ADD NodeCache ... 
setNodeCache: {Name:node3 Datasets:4e1d6a39e550ee41&5&0;8e1d6a39e550ee4x&5&1 FreeSize:0 AllocatableSize:0} 
========================================
ADD NodeCache ... 
setNodeCache: {Name:node2 Datasets:4e1d6a39e550ee41&5&0 FreeSize:0 AllocatableSize:0} 
========================================
DELETE NodeCache ... 
deleteNodeCache: node3 
========================================
ADD NodeCache ... 
setNodeCache: {Name:node3 Datasets:4e1d6a39e550ee41&5&0;8e1d6a39e550ee4x&5&1 FreeSize:0 AllocatableSize:0} 


#######################################################
mem cache lister...
*NcItem Struct was not found for node1
*NcItem Struct was not found for node2
*NcItem Struct was not found for node3
#######################################################
========================================
ADD NodeCache ... 
setNodeCache: {Name:node1 Datasets:4e1d6a39e550ee41&5&0;8e1d6a39e550ee4x&5&1 FreeSize:100 AllocatableSize:380} 
setMemCache: &{Name:node1 Datasets:4e1d6a39e550ee41&5&0;8e1d6a39e550ee4x&5&1 FreeSize:100 AllocatableSize:380} 
========================================
ADD NodeCache ... 
setNodeCache: {Name:node3 Datasets:4e1d6a39e550ee41&3334445&0;8e1d6a39e550ee4x&5&1 FreeSize:0 AllocatableSize:0} 
setMemCache: &{Name:node3 Datasets:4e1d6a39e550ee41&3334445&0;8e1d6a39e550ee4x&5&1 FreeSize:0 AllocatableSize:0} 
========================================
ADD NodeCache ... 
setNodeCache: {Name:node2 Datasets:4e1d6a39e550ee41&5&0 FreeSize:0 AllocatableSize:0} 
setMemCache: &{Name:node2 Datasets:4e1d6a39e550ee41&5&0 FreeSize:0 AllocatableSize:0} 



#######################################################
mem cache lister...
setMemCache: &{Name:node1 Datasets:4e1d6a39e550ee41&5&0;8e1d6a39e550ee4x&5&1 FreeSize:100 AllocatableSize:380} 
setMemCache: &{Name:node2 Datasets:4e1d6a39e550ee41&5&0 FreeSize:0 AllocatableSize:0} 
setMemCache: &{Name:node3 Datasets:4e1d6a39e550ee41&3334445&0;8e1d6a39e550ee4x&5&1 FreeSize:0 AllocatableSize:0} 
#######################################################



#######################################################
mem cache lister...
setMemCache: &{Name:node1 Datasets:4e1d6a39e550ee41&5&0;8e1d6a39e550ee4x&5&1 FreeSize:100 AllocatableSize:380} 
setMemCache: &{Name:node2 Datasets:4e1d6a39e550ee41&5&0 FreeSize:0 AllocatableSize:0} 
setMemCache: &{Name:node3 Datasets:4e1d6a39e550ee41&3334445&0;8e1d6a39e550ee4x&5&1 FreeSize:0 AllocatableSize:0} 
#######################################################
========================================
UPDATE NodeCache ... 
setNodeCache: {Name:node3 Datasets:4e1d6a39e550ee41&5&0;8e1d6a39e550ee4x&5&1 FreeSize:0 AllocatableSize:0} 
setMemCache: &{Name:node3 Datasets:4e1d6a39e550ee41&5&0;8e1d6a39e550ee4x&5&1 FreeSize:0 AllocatableSize:0} 



#######################################################
mem cache lister...
setMemCache: &{Name:node1 Datasets:4e1d6a39e550ee41&5&0;8e1d6a39e550ee4x&5&1 FreeSize:100 AllocatableSize:380} 
setMemCache: &{Name:node2 Datasets:4e1d6a39e550ee41&5&0 FreeSize:0 AllocatableSize:0} 
setMemCache: &{Name:node3 Datasets:4e1d6a39e550ee41&5&0;8e1d6a39e550ee4x&5&1 FreeSize:0 AllocatableSize:0} 
#######################################################


```