package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

// Phase of the Openstack Service Life Cyle
type OslcPhase string

// Describe the Phase of the Openstack Service Life Cycle
const (
	PhasePlanning       OslcPhase = "planning"
	PhaseInstall        OslcPhase = "install"
	PhaseTest           OslcPhase = "test"
	PhaseTrafficRollout OslcPhase = "trafficrollout"
	PhaseOperational    OslcPhase = "operational"
	PhaseTrafficDrain   OslcPhase = "trafficdrain"
	PhaseUpgrade        OslcPhase = "upgrade"
	PhaseRollback       OslcPhase = "rollback"
	PhaseDelete         OslcPhase = "delete"
)

// String converts a OslcPhase to a printable string
func (x OslcPhase) String() string { return string(x) }

// Phase of the Openstack Service Life Cyle
type OslcFlowKind string

// Describe the Kind of the Openstack Service Life Cycle Flow
const (
	KindInstall   OslcFlowKind = "install"
	KindUpgrade   OslcFlowKind = "upgrade"
	KindRollback  OslcFlowKind = "rollback"
	KindUninstall OslcFlowKind = "uninstall"
)

// String converts a OslcPhase to a printable string
func (x OslcFlowKind) String() string { return string(x) }

// FlowSource describe the location of the CR to create during a Flow of an
// Openstack Service Life Cycle.
type FlowSource struct {
	// ``url`` or ``path`` to the flow's parent directory
	Location string `json:"location"`
	// source to build the flow: ``generated``, ``local``, or ``tar``
	Type string `json:"type"`
}

// OslcSpec defines the desired state of Oslc
type OslcSpec struct {
	// Openstack Service Name
	ServiceName string `json:"serviceName"`
	// Openstack Service EndPoint
	ServiceEndPoint string `json:"serviceEndPoint,omitempty"`

	// Points to chart or location
	Source FlowSource `json:"source"`
	// Kind of flow applied to the OpenstackService.
	FlowKind OslcFlowKind `json:"flowKind"`

	// Target state of the Lcm Custom Resources
	TargetState LcmResourceState `json:"targetState"`
	// revisionHistoryLimit is the maximum number of revisions that will
	// be maintained in the DeletePhase's revision history. The revision history
	// consists of all revisions not represented by a currently applied
	// DeletePhaseSpec version. The default value is 10.
	RevisionHistoryLimit *int32 `json:"revisionHistoryLimit,omitempty"`
}

// OslcStatus defines the observed state of Oslc
type OslcStatus struct {
	// Actual phase of the OpenstackService
	ActualPhase OslcPhase `json:"actualPhase"`

	OpenstackLcmStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Oslc is the Schema for the openstackdeployments API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=oslcs,shortName=oslc
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.actualState",description="State"
// +kubebuilder:printcolumn:name="Target State",type="string",JSONPath=".spec.targetState",description="Target State"
// +kubebuilder:printcolumn:name="Satisfied",type="boolean",JSONPath=".status.satisfied",description="Satisfied"
type Oslc struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OslcSpec   `json:"spec,omitempty"`
	Status OslcStatus `json:"status,omitempty"`
}

// Init is used to initialize an Oslc. Namely, if the state has not been
// specified, it will be set
func (obj *Oslc) Init() {
	if obj.Status.ActualState == "" {
		obj.Status.ActualState = StateUninitialized
	}
	if obj.Spec.TargetState == "" {
		obj.Spec.TargetState = StateDeployed
	}
	obj.Status.Succeeded = (obj.Spec.TargetState == obj.Status.ActualState)
}

// Return the list of dependent resources to watch
func (obj *Oslc) GetDependentResources() []unstructured.Unstructured {
	var res = make([]unstructured.Unstructured, 0)
	return res
}

// Convert an unstructured.Unstructured into a typed Oslc
func ToOslc(u *unstructured.Unstructured) *Oslc {
	var obj *Oslc
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &Oslc{
			ObjectMeta: metav1.ObjectMeta{
				Name:      u.GetName(),
				Namespace: u.GetNamespace(),
			},
		}
	}
	return obj
}

// Convert a typed Oslc into an unstructured.Unstructured
func (obj *Oslc) FromOslc() *unstructured.Unstructured {
	u := NewOslcVersionKind(obj.ObjectMeta.Namespace, obj.ObjectMeta.Name)
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// IsDeleted returns true if the oscl has been deleted
func (obj *Oslc) IsDeleted() bool {
	return obj.GetDeletionTimestamp() != nil
}

// IsTargetStateUnitialized returns true if the oscl is not managed by the reconcilier
func (obj *Oslc) IsTargetStateUninitialized() bool {
	return obj.Spec.TargetState == StateUninitialized
}

// IsSatisfied returns true if the oscl's actual state meets its target state
func (obj *Oslc) IsSatisfied() bool {
	return obj.Spec.TargetState == obj.Status.ActualState
}

func (obj *Oslc) GetName() string {
	return obj.ObjectMeta.Name
}

// Returns a GKV for Oslc
func NewOslcVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("openstacklcm.airshipit.org/v1alpha1")
	u.SetKind("Oslc")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OslcList contains a list of Oslc
type OslcList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Oslc `json:"items"`
}

// Convert an unstructured.Unstructured into a typed OslcList
func ToOslcList(u *unstructured.Unstructured) *OslcList {
	var obj *OslcList
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &OslcList{}
	}
	return obj
}

// Convert a typed OslcList into an unstructured.Unstructured
func (obj *OslcList) FromOslcList() *unstructured.Unstructured {
	u := NewOslcListVersionKind("", "")
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// JEB: Not sure yet if we really will need it
func (obj *OslcList) Equivalent(other *OslcList) bool {
	if other == nil {
		return false
	}
	return reflect.DeepEqual(obj.Items, other.Items)
}

// Returns a GKV for OslcList
func NewOslcListVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("openstacklcm.airshipit.org/v1alpha1")
	u.SetKind("OslcList")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

// LifecycleFlow represent the list of Phase for that Service
type LifecycleFlow struct {
	Name      string
	Namespace string
	FlowKind  OslcFlowKind

	// Main workflow
	Main *unstructured.Unstructured
	// Phases organized by name
	Phases map[string](unstructured.Unstructured)
}

// Returns the Name for the LifecycleFlow
func (obj *LifecycleFlow) GetName() string {
	return obj.Name
}

// Returns the Namespace for this LifecycleFlow
func (obj *LifecycleFlow) GetNamespace() string {
	return obj.Namespace
}

// Returns the FlowKind for this LifecycleFlow
func (obj *LifecycleFlow) GetFlowKind() OslcFlowKind {
	return obj.FlowKind
}

// Returns the DependentResource for this LifecycleFlow
func (obj *LifecycleFlow) GetDependentResources() []unstructured.Unstructured {
	res := make([]unstructured.Unstructured, 0)

	// Add the Worflow itself
	if obj.Main != nil {
		res = append(res, *obj.Main)
	}

	// Add all the depending Phases
	for _, item := range obj.Phases {
		res = append(res, item)
	}
	return res
}

// JEB: Not sure yet if we really will need it
func (obj *LifecycleFlow) Equivalent(other *LifecycleFlow) bool {
	if other == nil {
		return false
	}

	// Add the Worflow itself
	var equal bool
	if other.Main == nil && obj.Main == nil {
		equal = true
	} else if other.Main != nil && obj.Main != nil {
		equal = reflect.DeepEqual(obj.Main, other.Main)
	} else {
		equal = false
	}

	if equal {
		return reflect.DeepEqual(obj.Phases, other.Phases)
	} else {
		return false
	}
}

// Let's check the reference are setup properly.
func (obj *LifecycleFlow) CheckOwnerReference(refs []metav1.OwnerReference) bool {

	// Check main Worflow is owned by the cycle
	if obj.Main != nil && !reflect.DeepEqual(obj.Main.GetOwnerReferences(), refs) {
		log.Info("OwnerReference issue: ", "kind", obj.Main.GetKind(), "name", obj.Main.GetName())
		return false
	}

	// Checki that each phase is owned by the cycle
	for _, item := range obj.Phases {
		if !reflect.DeepEqual(item.GetOwnerReferences(), refs) {
			log.Info("OwnerReference issue: ", "kind", item.GetKind(), "name", item.GetName())
			return false
		}
	}

	return true
}

// Check the state of a service
func (obj *LifecycleFlow) IsReady() bool {

	dep := &KubernetesDependency{}

	// Check the state of the Main workflow to figure out
	// if the phase is still running
	if obj.Main != nil && !dep.IsUnstructuredReady(obj.Main) {
		return false
	}

	// The state of the main workflow should be enough to
	// reflect the state of the LifecycleFlow. Also we
	// have access to the Phases, don't need to monitor
	// their state
	// for _, item := range obj.Phases {
	// }

	return true
}

func (obj *LifecycleFlow) IsFailedOrError() bool {

	dep := &KubernetesDependency{}

	// Check the state of the Main workflow to figure out
	// if the phase is still running
	if obj.Main != nil && dep.IsUnstructuredFailedOrError(obj.Main) {
		return true
	}

	// The state of the main workflow should be enough to
	// reflect the state of the LifecycleFlow. Also we
	// have access to the Phases, don't need to monitor
	// their state
	// for _, item := range obj.Phases {
	// }

	return false
}

// Returns a new LifecycleFlow
func NewLifecycleFlow(namespace string, name string) *LifecycleFlow {
	res := &LifecycleFlow{Namespace: namespace, Name: name}
	res.Main = nil
	res.Phases = make(map[string]unstructured.Unstructured)
	return res
}

func init() {
	SchemeBuilder.Register(&Oslc{}, &OslcList{})
}
