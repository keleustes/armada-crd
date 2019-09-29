package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

type InServicePolicy struct {
	// TimeoutInSecond is the maximal allowed time in second of the entire test process.
	TimeoutInSecond int64 `json:"timeoutInSecond,omitempty"`
}

// OperationalPhaseSpec defines the desired state of OperationalPhase
type OperationalPhaseSpec struct {
	PhaseSpec `json:",inline"`

	// InServicePolicy configures the policy enforcement when service is operational
	InServicePolicy *InServicePolicy `json:"inServicePolicy,omitempty"`

	// Config is the set of extra Values added to the helm renderer.
	// Config map[string]interface{} `json:"config,omitempty"`
	Config map[string]string `json:"config,omitempty"`
}

// OperationalPhaseStatus defines the observed state of OperationalPhase
type OperationalPhaseStatus struct {
	PhaseStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OperationalPhase is the Schema for the openstackdeployments API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=operationalphases,shortName=osops
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.actualState",description="State"
// +kubebuilder:printcolumn:name="Target State",type="string",JSONPath=".spec.targetState",description="Target State"
// +kubebuilder:printcolumn:name="Satisfied",type="boolean",JSONPath=".status.satisfied",description="Satisfied"
type OperationalPhase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OperationalPhaseSpec   `json:"spec,omitempty"`
	Status OperationalPhaseStatus `json:"status,omitempty"`
}

// Init is used to initialize an OperationalPhase. Namely, if the state has not been
// specified, it will be set
func (obj *OperationalPhase) Init() {
	if obj.Status.ActualState == "" {
		obj.Status.ActualState = StateUninitialized
	}
	if obj.Spec.TargetState == "" {
		obj.Spec.TargetState = StateDeployed
	}
	obj.Status.Succeeded = (obj.Spec.TargetState == obj.Status.ActualState)
}

// Return the list of dependent resources to watch
func (obj *OperationalPhase) GetDependentResources() []unstructured.Unstructured {
	var res = make([]unstructured.Unstructured, 0)
	return res
}

// Convert an unstructured.Unstructured into a typed OperationalPhase
func ToOperationalPhase(u *unstructured.Unstructured) *OperationalPhase {
	var obj *OperationalPhase
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &OperationalPhase{
			ObjectMeta: metav1.ObjectMeta{
				Name:      u.GetName(),
				Namespace: u.GetNamespace(),
			},
		}
	}
	return obj
}

// Convert a typed OperationalPhase into an unstructured.Unstructured
func (obj *OperationalPhase) FromOperationalPhase() *unstructured.Unstructured {
	u := NewOperationalPhaseVersionKind(obj.ObjectMeta.Namespace, obj.ObjectMeta.Name)
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// IsDeleted returns true if the phase has been deleted
func (obj *OperationalPhase) IsDeleted() bool {
	return obj.GetDeletionTimestamp() != nil
}

// IsTargetStateUnitialized returns true if the phase is not managed by the reconcilier
func (obj *OperationalPhase) IsTargetStateUninitialized() bool {
	return obj.Spec.TargetState == StateUninitialized
}

// IsSatisfied returns true if the phase's actual state meets its target state
func (obj *OperationalPhase) IsSatisfied() bool {
	return obj.Spec.TargetState == obj.Status.ActualState
}

// Name of the Phase
func (obj *OperationalPhase) GetName() string {
	return obj.ObjectMeta.Name
}

// Returns a GKV for OperationalPhase
func NewOperationalPhaseVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("openstacklcm.airshipit.org/v1alpha1")
	u.SetKind("OperationalPhase")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OperationalPhaseList contains a list of OperationalPhase
type OperationalPhaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OperationalPhase `json:"items"`
}

// Convert an unstructured.Unstructured into a typed OperationalPhaseList
func ToOperationalPhaseList(u *unstructured.Unstructured) *OperationalPhaseList {
	var obj *OperationalPhaseList
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &OperationalPhaseList{}
	}
	return obj
}

// Convert a typed OperationalPhaseList into an unstructured.Unstructured
func (obj *OperationalPhaseList) FromOperationalPhaseList() *unstructured.Unstructured {
	u := NewOperationalPhaseListVersionKind("", "")
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// JEB: Not sure yet if we really will need it
func (obj *OperationalPhaseList) Equivalent(other *OperationalPhaseList) bool {
	if other == nil {
		return false
	}
	return reflect.DeepEqual(obj.Items, other.Items)
}

// Returns a GKV for OperationalPhaseList
func NewOperationalPhaseListVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("openstacklcm.airshipit.org/v1alpha1")
	u.SetKind("OperationalPhaseList")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

func init() {
	SchemeBuilder.Register(&OperationalPhase{}, &OperationalPhaseList{})
}
