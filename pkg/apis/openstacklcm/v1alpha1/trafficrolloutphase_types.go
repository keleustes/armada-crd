package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

// TrafficRolloutaStategy defines the strategy used during TrafficRollout
type TrafficRolloutStrategy struct {
	// TimeoutInSecond is the maximal allowed time in second of the entire trafficdrain process.
	TimeoutInSecond int64 `json:"timeoutInSecond,omitempty"`
}

// TrafficRolloutPhaseSpec defines the desired state of TrafficRolloutPhase
type TrafficRolloutPhaseSpec struct {
	PhaseSpec `json:",inline"`

	// TrafficRolloutStrategy configures the strateg during rollout process.
	TrafficRolloutStrategy *TrafficRolloutStrategy `json:"trafficRolloutStrategy,omitempty"`
}

// TrafficRolloutPhaseStatus defines the observed state of TrafficRolloutPhase
type TrafficRolloutPhaseStatus struct {
	PhaseStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TrafficRolloutPhase is the Schema for the openstackdeployments API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=trafficrolloutphases,shortName=osroll
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.actualState",description="State"
// +kubebuilder:printcolumn:name="Target State",type="string",JSONPath=".spec.targetState",description="Target State"
// +kubebuilder:printcolumn:name="Satisfied",type="boolean",JSONPath=".status.satisfied",description="Satisfied"
type TrafficRolloutPhase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TrafficRolloutPhaseSpec   `json:"spec,omitempty"`
	Status TrafficRolloutPhaseStatus `json:"status,omitempty"`
}

// Init is used to initialize an TrafficRolloutPhase. Namely, if the state has not been
// specified, it will be set
func (obj *TrafficRolloutPhase) Init() {
	if obj.Status.ActualState == "" {
		obj.Status.ActualState = StateUninitialized
	}
	if obj.Spec.TargetState == "" {
		obj.Spec.TargetState = StateDeployed
	}
	obj.Status.Succeeded = (obj.Spec.TargetState == obj.Status.ActualState)
}

// Return the list of dependent resources to watch
func (obj *TrafficRolloutPhase) GetDependentResources() []unstructured.Unstructured {
	var res = make([]unstructured.Unstructured, 0)
	return res
}

// Convert an unstructured.Unstructured into a typed TrafficRolloutPhase
func ToTrafficRolloutPhase(u *unstructured.Unstructured) *TrafficRolloutPhase {
	var obj *TrafficRolloutPhase
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &TrafficRolloutPhase{
			ObjectMeta: metav1.ObjectMeta{
				Name:      u.GetName(),
				Namespace: u.GetNamespace(),
			},
		}
	}
	return obj
}

// Convert a typed TrafficRolloutPhase into an unstructured.Unstructured
func (obj *TrafficRolloutPhase) FromTrafficRolloutPhase() *unstructured.Unstructured {
	u := NewTrafficRolloutPhaseVersionKind(obj.ObjectMeta.Namespace, obj.ObjectMeta.Name)
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// IsDeleted returns true if the phase has been deleted
func (obj *TrafficRolloutPhase) IsDeleted() bool {
	return obj.GetDeletionTimestamp() != nil
}

// IsTargetStateUnitialized returns true if the phase is not managed by the reconcilier
func (obj *TrafficRolloutPhase) IsTargetStateUninitialized() bool {
	return obj.Spec.TargetState == StateUninitialized
}

// IsSatisfied returns true if the phase's actual state meets its target state
func (obj *TrafficRolloutPhase) IsSatisfied() bool {
	return obj.Spec.TargetState == obj.Status.ActualState
}

// Name of the Phase
func (obj *TrafficRolloutPhase) GetName() string {
	return obj.ObjectMeta.Name
}

// Returns a GKV for TrafficRolloutPhase
func NewTrafficRolloutPhaseVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("openstacklcm.airshipit.org/v1alpha1")
	u.SetKind("TrafficRolloutPhase")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TrafficRolloutPhaseList contains a list of TrafficRolloutPhase
type TrafficRolloutPhaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrafficRolloutPhase `json:"items"`
}

// Convert an unstructured.Unstructured into a typed TrafficRolloutPhaseList
func ToTrafficRolloutPhaseList(u *unstructured.Unstructured) *TrafficRolloutPhaseList {
	var obj *TrafficRolloutPhaseList
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &TrafficRolloutPhaseList{}
	}
	return obj
}

// Convert a typed TrafficRolloutPhaseList into an unstructured.Unstructured
func (obj *TrafficRolloutPhaseList) FromTrafficRolloutPhaseList() *unstructured.Unstructured {
	u := NewTrafficRolloutPhaseListVersionKind("", "")
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// JEB: Not sure yet if we really will need it
func (obj *TrafficRolloutPhaseList) Equivalent(other *TrafficRolloutPhaseList) bool {
	if other == nil {
		return false
	}
	return reflect.DeepEqual(obj.Items, other.Items)
}

// Returns a GKV for TrafficRolloutPhaseList
func NewTrafficRolloutPhaseListVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("openstacklcm.airshipit.org/v1alpha1")
	u.SetKind("TrafficRolloutPhaseList")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

func init() {
	SchemeBuilder.Register(&TrafficRolloutPhase{}, &TrafficRolloutPhaseList{})
}
