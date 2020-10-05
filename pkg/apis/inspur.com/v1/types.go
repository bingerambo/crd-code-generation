package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NodeCacheStatus represents the status of NodeCche.
type NodeCacheStatus struct {
	// The number of 'Unknonw' NodeCache in this Node.
	Unknown int32 `json:"unknown,omitempty" protobuf:"bytes,1,opt,name=unknown"`
	// The number of 'Pending' NodeCache in this queue.
	Pending int32 `json:"pending,omitempty" protobuf:"bytes,2,opt,name=pending"`
	// The number of 'Running' NodeCache in this queue.
	Running int32 `json:"running,omitempty" protobuf:"bytes,3,opt,name=running"`
}


// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type NodeCache struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec NodeCacheSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	// The status of NCDataSet.
	// +optional
	Status NodeCacheStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type NodeCacheSpec struct {
	// dataset list
	//Datasets []DataSetSummary `json:"datasets,omitempty" protobuf:"bytes,1,opt,name=datasets"`
	Datasets string `json:"datasets,omitempty" protobuf:"bytes,1,opt,name=datasets"`
	// Disk size unit :GB
	FreeSize int64 `json:"freesize,omitempty" protobuf:"bytes,2,opt,name=freesize"`
	// Disk size unit :GB
	AllocatableSize int64 `json:"allocatablesize,omitempty" protobuf:"bytes,2,opt,name=allocatablesize"`
}

//// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//
//type DataSetSummary struct {
//	Digest string `json:"digest,omitempty" protobuf:"bytes,1,opt,name=digest"`
//	Size   int64  `json:"size,omitempty" protobuf:"bytes,2,opt,name=size"`
//	InUse  bool   `json:"inuse,omitempty" protobuf:"bytes,3,opt,name=inuse"`
//}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeCacheList is a collection of NodeCache.
type NodeCacheList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// items is the list of NCDataSet
	Items []NodeCache `json:"items" protobuf:"bytes,2,rep,name=items"`
}