package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeletePhaseSpec defines the desired state of DeletePhase
type DeletePhaseSpec struct {
	PhaseSpec `json:",inline"`

	// Should we also purge the database during delete
	PurgeDB string `json:"purgeDB,omitempty"`
}

// DeletePhaseStatus defines the observed state of DeletePhase
type DeletePhaseStatus struct {
	PhaseStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DeletePhase is the Schema for the openstackdeployments API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=deletephases,shortName=osdlt
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.actualState",description="State"
// +kubebuilder:printcolumn:name="Target State",type="string",JSONPath=".spec.targetState",description="Target State"
// +kubebuilder:printcolumn:name="Satisfied",type="boolean",JSONPath=".status.satisfied",description="Satisfied"
type DeletePhase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DeletePhaseSpec   `json:"spec,omitempty"`
	Status DeletePhaseStatus `json:"status,omitempty"`
}

// Init is used to initialize an DeletePhase. Namely, if the state has not been
// specified, it will be set
func (obj *DeletePhase) Init() {
	if obj.Status.ActualState == "" {
		obj.Status.ActualState = StateUninitialized
	}
	if obj.Spec.TargetState == "" {
		obj.Spec.TargetState = StateDeployed
	}
	obj.Status.Succeeded = (obj.Spec.TargetState == obj.Status.ActualState)
}

// Return the list of dependent resources to watch
func (obj *DeletePhase) GetDependentResources() []unstructured.Unstructured {
	var res = make([]unstructured.Unstructured, 0)
	return res
}

// Convert an unstructured.Unstructured into a typed DeletePhase
func ToDeletePhase(u *unstructured.Unstructured) *DeletePhase {
	var obj *DeletePhase
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &DeletePhase{
			ObjectMeta: metav1.ObjectMeta{
				Name:      u.GetName(),
				Namespace: u.GetNamespace(),
			},
		}
	}
	return obj
}

// Convert a typed DeletePhase into an unstructured.Unstructured
func (obj *DeletePhase) FromDeletePhase() *unstructured.Unstructured {
	u := NewDeletePhaseVersionKind(obj.ObjectMeta.Namespace, obj.ObjectMeta.Name)
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// IsDeleted returns true if the phase has been deleted
func (obj *DeletePhase) IsDeleted() bool {
	return obj.GetDeletionTimestamp() != nil
}

// IsTargetStateUnitialized returns true if the phase is not managed by the reconcilier
func (obj *DeletePhase) IsTargetStateUninitialized() bool {
	return obj.Spec.TargetState == StateUninitialized
}

// IsSatisfied returns true if the phase's actual state meets its target state
func (obj *DeletePhase) IsSatisfied() bool {
	return obj.Spec.TargetState == obj.Status.ActualState
}

// Name of the Phase
func (obj *DeletePhase) GetName() string {
	return obj.ObjectMeta.Name
}

// Returns a GKV for DeletePhase
func NewDeletePhaseVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("openstacklcm.airshipit.org/v1alpha1")
	u.SetKind("DeletePhase")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DeletePhaseList contains a list of DeletePhase
type DeletePhaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeletePhase `json:"items"`
}

// Convert an unstructured.Unstructured into a typed DeletePhaseList
func ToDeletePhaseList(u *unstructured.Unstructured) *DeletePhaseList {
	var obj *DeletePhaseList
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &DeletePhaseList{}
	}
	return obj
}

// Convert a typed DeletePhaseList into an unstructured.Unstructured
func (obj *DeletePhaseList) FromDeletePhaseList() *unstructured.Unstructured {
	u := NewDeletePhaseListVersionKind("", "")
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// JEB: Not sure yet if we really will need it
func (obj *DeletePhaseList) Equivalent(other *DeletePhaseList) bool {
	if other == nil {
		return false
	}
	return reflect.DeepEqual(obj.Items, other.Items)
}

// Returns a GKV for DeletePhaseList
func NewDeletePhaseListVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("openstacklcm.airshipit.org/v1alpha1")
	u.SetKind("DeletePhaseList")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

func init() {
	SchemeBuilder.Register(&DeletePhase{}, &DeletePhaseList{})
}
