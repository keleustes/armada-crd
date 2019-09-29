package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

type TrafficDrainStrategy struct {
	// TimeoutInSecond is the maximal allowed time in second of the entire trafficdrain process.
	TimeoutInSecond int64 `json:"timeoutInSecond,omitempty"`
}

// TrafficDrainPhaseSpec defines the desired state of TrafficDrainPhase
type TrafficDrainPhaseSpec struct {
	PhaseSpec `json:",inline"`

	// TrafficDrainStrategy configures the strategy during drain process.
	TrafficDrainStrategy *TrafficDrainStrategy `json:"trafficDrainStrategy,omitempty"`
}

// TrafficDrainPhaseStatus defines the observed state of TrafficDrainPhase
type TrafficDrainPhaseStatus struct {
	PhaseStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TrafficDrainPhase is the Schema for the openstackdeployments API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=trafficdrainphases,shortName=osdrain
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.actualState",description="State"
// +kubebuilder:printcolumn:name="Target State",type="string",JSONPath=".spec.targetState",description="Target State"
// +kubebuilder:printcolumn:name="Satisfied",type="boolean",JSONPath=".status.satisfied",description="Satisfied"
type TrafficDrainPhase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TrafficDrainPhaseSpec   `json:"spec,omitempty"`
	Status TrafficDrainPhaseStatus `json:"status,omitempty"`
}

// Init is used to initialize an TrafficDrainPhase. Namely, if the state has not been
// specified, it will be set
func (obj *TrafficDrainPhase) Init() {
	if obj.Status.ActualState == "" {
		obj.Status.ActualState = StateUninitialized
	}
	if obj.Spec.TargetState == "" {
		obj.Spec.TargetState = StateDeployed
	}
	obj.Status.Succeeded = (obj.Spec.TargetState == obj.Status.ActualState)
}

// Return the list of dependent resources to watch
func (obj *TrafficDrainPhase) GetDependentResources() []unstructured.Unstructured {
	var res = make([]unstructured.Unstructured, 0)
	return res
}

// Convert an unstructured.Unstructured into a typed TrafficDrainPhase
func ToTrafficDrainPhase(u *unstructured.Unstructured) *TrafficDrainPhase {
	var obj *TrafficDrainPhase
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &TrafficDrainPhase{
			ObjectMeta: metav1.ObjectMeta{
				Name:      u.GetName(),
				Namespace: u.GetNamespace(),
			},
		}
	}
	return obj
}

// Convert a typed TrafficDrainPhase into an unstructured.Unstructured
func (obj *TrafficDrainPhase) FromTrafficDrainPhase() *unstructured.Unstructured {
	u := NewTrafficDrainPhaseVersionKind(obj.ObjectMeta.Namespace, obj.ObjectMeta.Name)
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// IsDeleted returns true if the phase has been deleted
func (obj *TrafficDrainPhase) IsDeleted() bool {
	return obj.GetDeletionTimestamp() != nil
}

// IsTargetStateUnitialized returns true if the phase is not managed by the reconcilier
func (obj *TrafficDrainPhase) IsTargetStateUninitialized() bool {
	return obj.Spec.TargetState == StateUninitialized
}

// IsSatisfied returns true if the phase's actual state meets its target state
func (obj *TrafficDrainPhase) IsSatisfied() bool {
	return obj.Spec.TargetState == obj.Status.ActualState
}

// Name of the Phase
func (obj *TrafficDrainPhase) GetName() string {
	return obj.ObjectMeta.Name
}

// Returns a GKV for TrafficDrainPhase
func NewTrafficDrainPhaseVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("openstacklcm.airshipit.org/v1alpha1")
	u.SetKind("TrafficDrainPhase")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TrafficDrainPhaseList contains a list of TrafficDrainPhase
type TrafficDrainPhaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrafficDrainPhase `json:"items"`
}

// Convert an unstructured.Unstructured into a typed TrafficDrainPhaseList
func ToTrafficDrainPhaseList(u *unstructured.Unstructured) *TrafficDrainPhaseList {
	var obj *TrafficDrainPhaseList
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &TrafficDrainPhaseList{}
	}
	return obj
}

// Convert a typed TrafficDrainPhaseList into an unstructured.Unstructured
func (obj *TrafficDrainPhaseList) FromTrafficDrainPhaseList() *unstructured.Unstructured {
	u := NewTrafficDrainPhaseListVersionKind("", "")
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// JEB: Not sure yet if we really will need it
func (obj *TrafficDrainPhaseList) Equivalent(other *TrafficDrainPhaseList) bool {
	if other == nil {
		return false
	}
	return reflect.DeepEqual(obj.Items, other.Items)
}

// Returns a GKV for TrafficDrainPhaseList
func NewTrafficDrainPhaseListVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("openstacklcm.airshipit.org/v1alpha1")
	u.SetKind("TrafficDrainPhaseList")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

func init() {
	SchemeBuilder.Register(&TrafficDrainPhase{}, &TrafficDrainPhaseList{})
}
