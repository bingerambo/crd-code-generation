package sdhedulerCache

import (
	"encoding/json"
	"fmt"
	ncv1 "github.com/bingerambo/crd-code-generation/pkg/apis/inspur.com/v1"
	ncinfov1 "github.com/bingerambo/crd-code-generation/pkg/nodecache_client/informers/externalversions/inspur/v1"
	"github.com/golang/glog"
	"k8s.io/client-go/tools/cache"
	"sync"
)

func init() {
	GSchedulerCache = &SchedulerCache{
		NC: NodeCache{
			NCMap: map[string]*NcItem{},
		},
	}
}

const (
	LOG_LEVEL_DEBUG = 100
	//PodGroupVersionV1Alpha1 represents PodGroupVersion of V1Alpha1
	NodeCacheGroup     string = "inspur.com"
	NodeCacheVersionV1 string = "v1"
)

type NcItem struct {
	Name     string
	Datasets string `json:"datasets,omitempty" protobuf:"bytes,1,opt,name=datasets"`
	// Disk size unit :GB
	FreeSize int64 `json:"freesize,omitempty" protobuf:"bytes,2,opt,name=freesize"`
	// Disk size unit :GB
	AllocatableSize int64 `json:"allocatablesize,omitempty" protobuf:"bytes,2,opt,name=allocatablesize"`
}
type NodeCache struct {
	// nodename -> nodecache
	NCMap map[string]*NcItem
}

type SchedulerCache struct {
	sync.Mutex
	NCInformerv1 ncinfov1.NodeCacheInformer
	NC      NodeCache
	Version string
}

var GSchedulerCache *SchedulerCache

// AddNodeCacheV1 add NodeCache to scheduler cache
func (sc *SchedulerCache) AddNodeCacheV1(obj interface{}) {
	nc, ok := obj.(*ncv1.NodeCache)
	if !ok {
		glog.Errorf("Cannot convert to *ncv1.NodeCache: %v", obj)
		return
	}

	marshalled, err := json.Marshal(*nc)
	if err != nil {
		glog.Errorf("Failed to Marshal NodeCache %s with error: %v", nc.Name, err)
	}

	local_nc := &ncv1.NodeCache{}
	err = json.Unmarshal(marshalled, local_nc)
	if err != nil {
		glog.Errorf("Failed to Unmarshal Data into api.NodeCache type with error: %v", err)
	}
	//nc.Version = NodeCacheVersionV1

	sc.Mutex.Lock()
	defer sc.Mutex.Unlock()

	glog.V(LOG_LEVEL_DEBUG).Infof("Add NodeCache(%s) into cache, spec(%#v)", nc.Name, nc.Spec)

	err = sc.addNodeCache(nc)
	if err != nil {
		glog.Errorf("Failed to add NodeCache %s into cache: %v", nc.Name, err)
		return
	}
	return
}




// UpdateNodeCacheV1 add NodeCache to scheduler cache
func (sc *SchedulerCache) UpdateNodeCacheV1(oldObj, newObj interface{}) {
	oldNC, ok := oldObj.(*ncv1.NodeCache)
	if !ok {
		glog.Errorf("Cannot convert oldObj to *ncv1.NodeCache: %v", oldObj)
		return
	}
	newNC, ok := newObj.(*ncv1.NodeCache)
	if !ok {
		glog.Errorf("Cannot convert newObj to *ncv1.NodeCache: %v", newObj)
		return
	}

	oldMarshalled, err := json.Marshal(*oldNC)
	if err != nil {
		glog.Errorf("Failed to Marshal NodeCache %s with error: %v", oldNC.Name, err)
	}

	local_oldNc := &ncv1.NodeCache{}
	//oldNc.Version = api.PodGroupVersionV1Alpha1

	err = json.Unmarshal(oldMarshalled, local_oldNc)
	if err != nil {
		glog.Errorf("Failed to Unmarshal Data into NodeCache type with error: %v", err)
	}

	newMarshalled, err := json.Marshal(*newNC)
	if err != nil {
		glog.Errorf("Failed to Marshal NodeCache %s with error: %v", newNC.Name, err)
	}

	local_newNc := &ncv1.NodeCache{}

	err = json.Unmarshal(newMarshalled, local_newNc)
	if err != nil {
		glog.Errorf("Failed to Unmarshal Data into NodeCache type with error: %v", err)
	}

	sc.Mutex.Lock()
	defer sc.Mutex.Unlock()

	err = sc.updateNodeCache(local_oldNc, local_newNc)
	if err != nil {
		glog.Errorf("Failed to update NodeCache %s into cache: %v", local_oldNc.Name, err)
		return
	}
	return
}

// DeleteNodeCacheV1 delete podgroup from scheduler cache
func (sc *SchedulerCache) DeleteNodeCacheV1(obj interface{}) {
	var nc *ncv1.NodeCache
	switch t := obj.(type) {
	case *ncv1.NodeCache:
		nc = t
	case cache.DeletedFinalStateUnknown:
		var ok bool
		nc, ok = t.Obj.(*ncv1.NodeCache)
		if !ok {
			glog.Errorf("Cannot convert to *ncv1.NodeCache: %v", t.Obj)
			return
		}
	default:
		glog.Errorf("Cannot convert to *ncv1.NodeCache: %v", t)
		return
	}

	marshalled, err := json.Marshal(*nc)
	if err != nil {
		glog.Errorf("Failed to Marshal NodeCache %s with error: %v", nc.Name, err)
	}

	local_nc := &ncv1.NodeCache{}

	err = json.Unmarshal(marshalled, local_nc)
	if err != nil {
		glog.Errorf("Failed to Unmarshal Data into NodeCache type with error: %v", err)
	}

	sc.Mutex.Lock()
	defer sc.Mutex.Unlock()

	err = sc.deleteNodeCache(local_nc)
	if err != nil {
		glog.Errorf("Failed to delete NodeCache %s from cache: %v", local_nc.Name, err)
		return
	}
	return
}

// Assumes that lock is already acquired.
func (sc *SchedulerCache) setNodeCache(nc *ncv1.NodeCache) error {
	ncItem := &NcItem{
		Name:            nc.Name,
		Datasets:        nc.Spec.Datasets,
		FreeSize:        nc.Spec.FreeSize,
		AllocatableSize: nc.Spec.AllocatableSize,
	}
	sc.NC.NCMap[nc.Name] = ncItem
	fmt.Printf("setNodeCache: %++v \n", *ncItem)


	return nil

}
// Assumes that lock is already acquired.
func (sc *SchedulerCache) addNodeCache(nc *ncv1.NodeCache) error {
	fmt.Println("========================================")
	fmt.Println("ADD NodeCache ... ")
	return sc.setNodeCache(nc)

}

// Assumes that lock is already acquired.
func (sc *SchedulerCache) updateNodeCache(oldNc, newNc *ncv1.NodeCache) error {
	fmt.Println("========================================")
	fmt.Println("UPDATE NodeCache ... ")

	//sc.deletePodGroup(oldQueue)
	return sc.setNodeCache(newNc)
}

// Assumes that lock is already acquired.
func (sc *SchedulerCache) deleteNodeCache(nc *ncv1.NodeCache) error {
	fmt.Println("========================================")
	fmt.Println("DELETE NodeCache ... ")
	fmt.Printf("deleteNodeCache: %s \n", nc.Name)
	delete(sc.NC.NCMap, nc.Name)
	return nil
}

// Run  starts the schedulerCache

func (sc *SchedulerCache) Run(stopCh <-chan struct{}) {
	//delete pdb support
	//go sc.pdbInformer.Informer().Run(stopCh)
	// sc.xxxInformer.Informer() 实际就是 SharedIndexInformer 接口
	// Informer() 实际上就是之前生成的 sc.xxxInformer.Informer()
	// Informer.Run 就是 informer的controller.run
	go sc.NCInformerv1.Informer().Run(stopCh)
}

func (sc *SchedulerCache) WaitForCacheSync(stopCh <-chan struct{}) bool {

	return cache.WaitForCacheSync(stopCh,
		func() []cache.InformerSynced {
			informerSynced := []cache.InformerSynced{
				//delete pdb support
				//sc.pdbInformer.Informer().HasSynced,
				//sc.podInformer.Informer().HasSynced,
				//sc.podGroupInformerv1alpha1.Informer().HasSynced,
				//sc.podGroupInformerv1alpha2.Informer().HasSynced,
				//sc.nodeInformer.Informer().HasSynced,
				//sc.pvInformer.Informer().HasSynced,
				//sc.pvcInformer.Informer().HasSynced,
				//sc.scInformer.Informer().HasSynced,
				//sc.queueInformerv1alpha1.Informer().HasSynced,
				//sc.queueInformerv1alpha2.Informer().HasSynced,
				sc.NCInformerv1.Informer().HasSynced,
			}

			return informerSynced
		}()...,
	)
}

