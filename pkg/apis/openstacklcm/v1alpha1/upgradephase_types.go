package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

// BackupPolicy defines backup policy.
type BackupPolicy struct {
	// TimeoutInSecond is the maximal allowed time in second of the entire backup process.
	TimeoutInSecond int64 `json:"timeoutInSecond,omitempty"`
}

// BackupSource contains the supported backup sources.
type BackupSource struct {
	// Offsite defines the Offsite backup source spec.
	Offsite *OffsiteBackupSource `json:"offsite,omitempty"`
	// Ceph defines the Ceph backup source spec.
	Ceph *CephBackupSource `json:"ceph,omitempty"`
}

// OffsiteBackupSource provides the spec how to store backups on Offsite.
type OffsiteBackupSource struct {
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
	Endpoint string `json:"endpoint,omitempty"`

	// ForcePathStyle forces to use path style over the default subdomain style.
	// This is useful when you have an offsite compatible endpoint that doesn't support
	// subdomain buckets.
	ForcePathStyle bool `json:"forcePathStyle"`
}

// CephBackupSource provides the spec how to store backups on Ceph.
type CephBackupSource struct {
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

// UpgradePhaseSpec defines the desired state of UpgradePhase
type UpgradePhaseSpec struct {
	PhaseSpec `json:",inline"`

	// Should we also backup the database before upgrade
	BackupDB string `json:"backupDB,omitempty"`

	// StorageType is the armada backup storage type.
	// We need this field because CRD doesn't support validation against invalid fields
	// and we cannot verify invalid backup storage source.
	StorageType BackupStorageType `json:"storageType,omitempty"`
	// BackupPolicy configures the backup process.
	BackupPolicy *BackupPolicy `json:"backupPolicy,omitempty"`
	// BackupSource is the backup storage source.
	BackupSource `json:",inline"`

	// Config is the set of extra Values added to the helm renderer.
	// Config map[string]interface{} `json:"config,omitempty"`
	Config map[string]string `json:"config,omitempty"`
}

// UpgradePhaseStatus defines the observed state of UpgradePhase
type UpgradePhaseStatus struct {
	PhaseStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UpgradePhase is the Schema for the openstackdeployments API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=upgradephases,shortName=osupg
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.actualState",description="State"
// +kubebuilder:printcolumn:name="Target State",type="string",JSONPath=".spec.targetState",description="Target State"
// +kubebuilder:printcolumn:name="Satisfied",type="boolean",JSONPath=".status.satisfied",description="Satisfied"
type UpgradePhase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UpgradePhaseSpec   `json:"spec,omitempty"`
	Status UpgradePhaseStatus `json:"status,omitempty"`
}

// Init is used to initialize an UpgradePhase. Namely, if the state has not been
// specified, it will be set
func (obj *UpgradePhase) Init() {
	if obj.Status.ActualState == "" {
		obj.Status.ActualState = StateUninitialized
	}
	if obj.Spec.TargetState == "" {
		obj.Spec.TargetState = StateDeployed
	}
	obj.Status.Succeeded = (obj.Spec.TargetState == obj.Status.ActualState)
}

// Return the list of dependent resources to watch
func (obj *UpgradePhase) GetDependentResources() []unstructured.Unstructured {
	var res = make([]unstructured.Unstructured, 0)
	return res
}

// Convert an unstructured.Unstructured into a typed UpgradePhase
func ToUpgradePhase(u *unstructured.Unstructured) *UpgradePhase {
	var obj *UpgradePhase
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &UpgradePhase{
			ObjectMeta: metav1.ObjectMeta{
				Name:      u.GetName(),
				Namespace: u.GetNamespace(),
			},
		}
	}
	return obj
}

// Convert a typed UpgradePhase into an unstructured.Unstructured
func (obj *UpgradePhase) FromUpgradePhase() *unstructured.Unstructured {
	u := NewUpgradePhaseVersionKind(obj.ObjectMeta.Namespace, obj.ObjectMeta.Name)
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// IsDeleted returns true if the phase has been deleted
func (obj *UpgradePhase) IsDeleted() bool {
	return obj.GetDeletionTimestamp() != nil
}

// IsTargetStateUnitialized returns true if the phase is not managed by the reconcilier
func (obj *UpgradePhase) IsTargetStateUninitialized() bool {
	return obj.Spec.TargetState == StateUninitialized
}

// IsSatisfied returns true if the phase's actual state meets its target state
func (obj *UpgradePhase) IsSatisfied() bool {
	return obj.Spec.TargetState == obj.Status.ActualState
}

// Name of the Phase
func (obj *UpgradePhase) GetName() string {
	return obj.ObjectMeta.Name
}

// Returns a GKV for UpgradePhase
func NewUpgradePhaseVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("openstacklcm.airshipit.org/v1alpha1")
	u.SetKind("UpgradePhase")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UpgradePhaseList contains a list of UpgradePhase
type UpgradePhaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UpgradePhase `json:"items"`
}

// Convert an unstructured.Unstructured into a typed UpgradePhaseList
func ToUpgradePhaseList(u *unstructured.Unstructured) *UpgradePhaseList {
	var obj *UpgradePhaseList
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &UpgradePhaseList{}
	}
	return obj
}

// Convert a typed UpgradePhaseList into an unstructured.Unstructured
func (obj *UpgradePhaseList) FromUpgradePhaseList() *unstructured.Unstructured {
	u := NewUpgradePhaseListVersionKind("", "")
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// JEB: Not sure yet if we really will need it
func (obj *UpgradePhaseList) Equivalent(other *UpgradePhaseList) bool {
	if other == nil {
		return false
	}
	return reflect.DeepEqual(obj.Items, other.Items)
}

// Returns a GKV for UpgradePhaseList
func NewUpgradePhaseListVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("openstacklcm.airshipit.org/v1alpha1")
	u.SetKind("UpgradePhaseList")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

func init() {
	SchemeBuilder.Register(&UpgradePhase{}, &UpgradePhaseList{})
}
