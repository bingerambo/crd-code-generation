package main

import (
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"time"

	"github.com/golang/glog"

	"k8s.io/client-go/tools/clientcmd"

	"github.com/bingerambo/crd-code-generation/cmd/nodecache/sdhedulerCache"

	ncver "github.com/bingerambo/crd-code-generation/pkg/nodecache_client/clientset/versioned"
	ncinfo "github.com/bingerambo/crd-code-generation/pkg/nodecache_client/informers/externalversions"
)

var (
	kuberconfig = flag.String("kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	master      = flag.String("master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
)

func main() {
	flag.Parse()

	kuberconfig := "D:\\GO_projects\\src\\github.com\\openshift-evangelists\\crd-code-generation\\cmd\\nodecache\\config"

	//nodecaecheClientSet(kuberconfig)
	//fmt.Println()
	//fmt.Println()
	//fmt.Println()
	//fmt.Println()
	//fmt.Println()

	nodecacheInformer(kuberconfig)

	// block main process
	stopCh := make(chan struct{})
	for {
		select {
		case <-stopCh:
			return
		default:

		}
	}

}

func nodecaecheClientSet(kuberconfig string) {
	cfg, err := clientcmd.BuildConfigFromFlags(*master, kuberconfig)
	if err != nil {
		glog.Fatalf("Error building kubeconfig: %v", err)
	}

	//exampleClient, err := examplecomclientset.NewForConfig(cfg)
	nodecacheclientset, err := ncver.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building nodecache clientset: %v", err)
	}

	//list, err := exampleClient.ExampleV1().Databases("default").List(metav1.ListOptions{})
	list, err := nodecacheclientset.InspurV1().NodeCaches("").List(metav1.ListOptions{})
	if err != nil {
		glog.Fatalf("Error listing all databases: %v", err)
	}

	for _, nodecache := range list.Items {
		fmt.Printf("nodecache %s with datasets: %q\n", nodecache.Name, nodecache.Spec.Datasets)
		fmt.Printf("nodecache %s with freesize: %d\n", nodecache.Name, nodecache.Spec.FreeSize)
		fmt.Printf("nodecache %s with allocatablesize: %d\n", nodecache.Name, nodecache.Spec.AllocatableSize)
	}

}

func nodecacheInformer(kuberconfig string) {
	cfg, err := clientcmd.BuildConfigFromFlags(*master, kuberconfig)
	if err != nil {
		glog.Fatalf("Error building kubeconfig: %v", err)
	}
	var ncclient *ncver.Clientset
	ncclient, err = ncver.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building nodecache clientset: %v", err)
	}

	gsc := sdhedulerCache.GSchedulerCache
	ncinformer := ncinfo.NewSharedInformerFactory(ncclient, 0)
	// create informer for NodeCaches(v1) information
	//ncInformerv1 := ncinformer.Inspur().V1().NodeCaches()
	//ncInformerv1.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
	//	AddFunc:    gsc.AddNodeCacheV1,
	//	//UpdateFunc: UpdatePodGroupAlpha1,
	//	//DeleteFunc: DeletePodGroupAlpha1,
	//})

	gsc.NCInformerv1 = ncinformer.Inspur().V1().NodeCaches()
	gsc.NCInformerv1.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    gsc.AddNodeCacheV1,
		UpdateFunc: gsc.UpdateNodeCacheV1,
		DeleteFunc: gsc.DeleteNodeCacheV1,
	})

	stopCh := make(chan struct{})

	// .0 start up mem cache lister...
	go wait.Until(func() {
		mc := gsc.MC
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println("#######################################################")
		fmt.Println("mem cache lister...")
		for index := 1; index <= 3; index++ {
			x, found := mc.Get(fmt.Sprintf("node%d", index))
			if !found {
				fmt.Printf("*NcItem Struct was not found for node%d\n",index)
				continue
			}
			fmt.Printf("setMemCache: %++v \n", x.(*sdhedulerCache.NcItem))
		}

		fmt.Println("#######################################################")
	}, 5*time.Second, stopCh)

	// .1 start all informer, list & watch

	go gsc.Run(stopCh)
	// .2 sync resources from apiserver by list...
	gsc.WaitForCacheSync(stopCh)

}
