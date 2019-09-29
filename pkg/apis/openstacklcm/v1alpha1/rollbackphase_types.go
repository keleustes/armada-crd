package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

type RestoreSource struct {
	// Offsite tells where on Offsite the backup is saved and how to fetch the backup.
	Offsite *OffsiteRestoreSource `json:"offsite,omitempty"`

	// Ceph tells where on Ceph the backup is saved and how to fetch the backup.
	Ceph *CephRestoreSource `json:"ceph,omitempty"`
}

type OffsiteRestoreSource struct {
	// Path is the full offsite path where the backup is saved.
	// The format of the path must be: "<offsite-bucket-name>/<path-to-backup-file>"
	// e.g: "mybucket/armada.backup"
	Path string `json:"path"`

	// The name of the secret object that stores the Offsite credential and config files.
	// The file name of the credential MUST be 'credentials'.
	// The file name of the config MUST be 'config'.
	// The profile to use in both files will be 'default'.
	//
	// OffsiteSecret overwrites the default armada operator wide Offsite credential and config.
	OffsiteSecret string `json:"offsiteSecret"`

	// Endpoint if blank points to offsite. If specified, can point to offsite compatible object
	// stores.
	Endpoint string `json:"endpoint"`

	// ForcePathStyle forces to use path style over the default subdomain style.
	// This is useful when you have an offsite compatible endpoint that doesn't support
	// subdomain buckets.
	ForcePathStyle bool `json:"forcePathStyle"`
}

type CephRestoreSource struct {
	// Path is the full Ceph path where the backup is saved.
	// The format of the path must be: "<ceph-bucket-name>/<path-to-backup-file>"
	// e.g: "mycephbucket/armada.backup"
	Path string `json:"path"`

	// The name of the secret object that stores the Google storage credential
	// containing at most ONE of the following:
	// An access token with file name of 'access-token'.
	// JSON credentials with file name of 'credentials.json'.
	//
	// If omitted, client will use the default application credentials.
	CephSecret string `json:"cephSecret,omitempty"`
}

// RollbackPhaseSpec defines the desired state of RollbackPhase
type RollbackPhaseSpec struct {
	PhaseSpec `json:",inline"`

	// Should we also restore the database during rollback
	RestoreDB string `json:"restoreDB,omitempty"`

	// StorageType is the type of the backup storage which is used as RestoreSource.
	StorageType BackupStorageType `json:"storageType,omitempty"`
	// RestoreSource tells the where to get the backup and restore from.
	RestoreSource `json:",inline"`
}

// RollbackPhaseStatus defines the observed state of RollbackPhase
type RollbackPhaseStatus struct {
	PhaseStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RollbackPhase is the Schema for the openstackdeployments API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=rollbackphases,shortName=osrbck
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.actualState",description="State"
// +kubebuilder:printcolumn:name="Target State",type="string",JSONPath=".spec.targetState",description="Target State"
// +kubebuilder:printcolumn:name="Satisfied",type="boolean",JSONPath=".status.satisfied",description="Satisfied"
type RollbackPhase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RollbackPhaseSpec   `json:"spec,omitempty"`
	Status RollbackPhaseStatus `json:"status,omitempty"`
}

// Init is used to initialize an RollbackPhase. Namely, if the state has not been
// specified, it will be set
func (obj *RollbackPhase) Init() {
	if obj.Status.ActualState == "" {
		obj.Status.ActualState = StateUninitialized
	}
	if obj.Spec.TargetState == "" {
		obj.Spec.TargetState = StateDeployed
	}
	obj.Status.Succeeded = (obj.Spec.TargetState == obj.Status.ActualState)
}

// Return the list of dependent resources to watch
func (obj *RollbackPhase) GetDependentResources() []unstructured.Unstructured {
	var res = make([]unstructured.Unstructured, 0)
	return res
}

// Convert an unstructured.Unstructured into a typed RollbackPhase
func ToRollbackPhase(u *unstructured.Unstructured) *RollbackPhase {
	var obj *RollbackPhase
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &RollbackPhase{
			ObjectMeta: metav1.ObjectMeta{
				Name:      u.GetName(),
				Namespace: u.GetNamespace(),
			},
		}
	}
	return obj
}

// Convert a typed RollbackPhase into an unstructured.Unstructured
func (obj *RollbackPhase) FromRollbackPhase() *unstructured.Unstructured {
	u := NewRollbackPhaseVersionKind(obj.ObjectMeta.Namespace, obj.ObjectMeta.Name)
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// IsDeleted returns true if the phase has been deleted
func (obj *RollbackPhase) IsDeleted() bool {
	return obj.GetDeletionTimestamp() != nil
}

// IsTargetStateUnitialized returns true if the phase is not managed by the reconcilier
func (obj *RollbackPhase) IsTargetStateUninitialized() bool {
	return obj.Spec.TargetState == StateUninitialized
}

// IsSatisfied returns true if the phase's actual state meets its target state
func (obj *RollbackPhase) IsSatisfied() bool {
	return obj.Spec.TargetState == obj.Status.ActualState
}

// Name of the Phase
func (obj *RollbackPhase) GetName() string {
	return obj.ObjectMeta.Name
}

// Returns a GKV for RollbackPhase
func NewRollbackPhaseVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("openstacklcm.airshipit.org/v1alpha1")
	u.SetKind("RollbackPhase")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RollbackPhaseList contains a list of RollbackPhase
type RollbackPhaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RollbackPhase `json:"items"`
}

// Convert an unstructured.Unstructured into a typed RollbackPhaseList
func ToRollbackPhaseList(u *unstructured.Unstructured) *RollbackPhaseList {
	var obj *RollbackPhaseList
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &RollbackPhaseList{}
	}
	return obj
}

// Convert a typed RollbackPhaseList into an unstructured.Unstructured
func (obj *RollbackPhaseList) FromRollbackPhaseList() *unstructured.Unstructured {
	u := NewRollbackPhaseListVersionKind("", "")
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// JEB: Not sure yet if we really will need it
func (obj *RollbackPhaseList) Equivalent(other *RollbackPhaseList) bool {
	if other == nil {
		return false
	}
	return reflect.DeepEqual(obj.Items, other.Items)
}

// Returns a GKV for RollbackPhaseList
func NewRollbackPhaseListVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("openstacklcm.airshipit.org/v1alpha1")
	u.SetKind("RollbackPhaseList")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

func init() {
	SchemeBuilder.Register(&RollbackPhase{}, &RollbackPhaseList{})
}
