// Code generated by protoc-gen-go. DO NOT EDIT.
// source: operator/v1alpha1/component.proto

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	v1 "k8s.io/api/core/v1"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// IstioComponentSpec defines the desired installed state of Istio components.
type IstioComponentSetSpec struct {
	Base                 *BaseComponentSpec `protobuf:"bytes,29,opt,name=base,proto3" json:"base,omitempty"`
	Pilot                *ComponentSpec     `protobuf:"bytes,30,opt,name=pilot,proto3" json:"pilot,omitempty"`
	Proxy                *ComponentSpec     `protobuf:"bytes,31,opt,name=proxy,proto3" json:"proxy,omitempty"`
	SidecarInjector      *ComponentSpec     `protobuf:"bytes,32,opt,name=sidecar_injector,json=sidecarInjector,proto3" json:"sidecarInjector,omitempty"`
	Policy               *ComponentSpec     `protobuf:"bytes,33,opt,name=policy,proto3" json:"policy,omitempty"`
	Telemetry            *ComponentSpec     `protobuf:"bytes,34,opt,name=telemetry,proto3" json:"telemetry,omitempty"`
	Citadel              *ComponentSpec     `protobuf:"bytes,35,opt,name=citadel,proto3" json:"citadel,omitempty"`
	NodeAgent            *ComponentSpec     `protobuf:"bytes,36,opt,name=node_agent,json=nodeAgent,proto3" json:"nodeAgent,omitempty"`
	Galley               *ComponentSpec     `protobuf:"bytes,37,opt,name=galley,proto3" json:"galley,omitempty"`
	Cni                  *ComponentSpec     `protobuf:"bytes,38,opt,name=cni,proto3" json:"cni,omitempty"`
	IstiodRemote         *ComponentSpec     `protobuf:"bytes,39,opt,name=istiod_remote,json=istiodRemote,proto3" json:"istiodRemote,omitempty"`
	IngressGateways      []*GatewaySpec     `protobuf:"bytes,40,rep,name=ingress_gateways,json=ingressGateways,proto3" json:"ingressGateways,omitempty"`
	EgressGateways       []*GatewaySpec     `protobuf:"bytes,41,rep,name=egress_gateways,json=egressGateways,proto3" json:"egressGateways,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *IstioComponentSetSpec) Reset()         { *m = IstioComponentSetSpec{} }
func (m *IstioComponentSetSpec) String() string { return proto.CompactTextString(m) }
func (*IstioComponentSetSpec) ProtoMessage()    {}
func (*IstioComponentSetSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ed34a579e9b43a2, []int{0}
}

func (m *IstioComponentSetSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IstioComponentSetSpec.Unmarshal(m, b)
}
func (m *IstioComponentSetSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IstioComponentSetSpec.Marshal(b, m, deterministic)
}
func (m *IstioComponentSetSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IstioComponentSetSpec.Merge(m, src)
}
func (m *IstioComponentSetSpec) XXX_Size() int {
	return xxx_messageInfo_IstioComponentSetSpec.Size(m)
}
func (m *IstioComponentSetSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_IstioComponentSetSpec.DiscardUnknown(m)
}

var xxx_messageInfo_IstioComponentSetSpec proto.InternalMessageInfo

func (m *IstioComponentSetSpec) GetBase() *BaseComponentSpec {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *IstioComponentSetSpec) GetPilot() *ComponentSpec {
	if m != nil {
		return m.Pilot
	}
	return nil
}

func (m *IstioComponentSetSpec) GetProxy() *ComponentSpec {
	if m != nil {
		return m.Proxy
	}
	return nil
}

func (m *IstioComponentSetSpec) GetSidecarInjector() *ComponentSpec {
	if m != nil {
		return m.SidecarInjector
	}
	return nil
}

func (m *IstioComponentSetSpec) GetPolicy() *ComponentSpec {
	if m != nil {
		return m.Policy
	}
	return nil
}

func (m *IstioComponentSetSpec) GetTelemetry() *ComponentSpec {
	if m != nil {
		return m.Telemetry
	}
	return nil
}

func (m *IstioComponentSetSpec) GetCitadel() *ComponentSpec {
	if m != nil {
		return m.Citadel
	}
	return nil
}

func (m *IstioComponentSetSpec) GetNodeAgent() *ComponentSpec {
	if m != nil {
		return m.NodeAgent
	}
	return nil
}

func (m *IstioComponentSetSpec) GetGalley() *ComponentSpec {
	if m != nil {
		return m.Galley
	}
	return nil
}

func (m *IstioComponentSetSpec) GetCni() *ComponentSpec {
	if m != nil {
		return m.Cni
	}
	return nil
}

func (m *IstioComponentSetSpec) GetIstiodRemote() *ComponentSpec {
	if m != nil {
		return m.IstiodRemote
	}
	return nil
}

func (m *IstioComponentSetSpec) GetIngressGateways() []*GatewaySpec {
	if m != nil {
		return m.IngressGateways
	}
	return nil
}

func (m *IstioComponentSetSpec) GetEgressGateways() []*GatewaySpec {
	if m != nil {
		return m.EgressGateways
	}
	return nil
}

// Configuration for base component.
type BaseComponentSpec struct {
	// Selects whether this component is installed.
	Enabled *BoolValueForPB `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Kubernetes resource spec.
	K8S                  *KubernetesResourcesSpec `protobuf:"bytes,50,opt,name=k8s,proto3" json:"k8s,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *BaseComponentSpec) Reset()         { *m = BaseComponentSpec{} }
func (m *BaseComponentSpec) String() string { return proto.CompactTextString(m) }
func (*BaseComponentSpec) ProtoMessage()    {}
func (*BaseComponentSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ed34a579e9b43a2, []int{1}
}

func (m *BaseComponentSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseComponentSpec.Unmarshal(m, b)
}
func (m *BaseComponentSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseComponentSpec.Marshal(b, m, deterministic)
}
func (m *BaseComponentSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseComponentSpec.Merge(m, src)
}
func (m *BaseComponentSpec) XXX_Size() int {
	return xxx_messageInfo_BaseComponentSpec.Size(m)
}
func (m *BaseComponentSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseComponentSpec.DiscardUnknown(m)
}

var xxx_messageInfo_BaseComponentSpec proto.InternalMessageInfo


func (m *BaseComponentSpec) GetK8S() *KubernetesResourcesSpec {
	if m != nil {
		return m.K8S
	}
	return nil
}

// Configuration for internal components.
type ComponentSpec struct {
	// Selects whether this component is installed.
	Enabled *BoolValueForPB `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Namespace for the component.
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Hub for the component (overrides top level hub setting).
	Hub string `protobuf:"bytes,10,opt,name=hub,proto3" json:"hub,omitempty"`
	// Tag for the component (overrides top level tag setting).
	Tag interface{} `protobuf:"bytes,11,opt,name=tag,proto3" json:"tag,omitempty"`
	// Arbitrary install time configuration for the component.
	Spec interface{} `protobuf:"bytes,30,opt,name=spec,proto3" json:"spec,omitempty"`
	// Kubernetes resource spec.
	K8S                  *KubernetesResourcesSpec `protobuf:"bytes,50,opt,name=k8s,proto3" json:"k8s,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ComponentSpec) Reset()         { *m = ComponentSpec{} }
func (m *ComponentSpec) String() string { return proto.CompactTextString(m) }
func (*ComponentSpec) ProtoMessage()    {}
func (*ComponentSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ed34a579e9b43a2, []int{2}
}

func (m *ComponentSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComponentSpec.Unmarshal(m, b)
}
func (m *ComponentSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComponentSpec.Marshal(b, m, deterministic)
}
func (m *ComponentSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComponentSpec.Merge(m, src)
}
func (m *ComponentSpec) XXX_Size() int {
	return xxx_messageInfo_ComponentSpec.Size(m)
}
func (m *ComponentSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ComponentSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ComponentSpec proto.InternalMessageInfo


func (m *ComponentSpec) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ComponentSpec) GetHub() string {
	if m != nil {
		return m.Hub
	}
	return ""
}



func (m *ComponentSpec) GetK8S() *KubernetesResourcesSpec {
	if m != nil {
		return m.K8S
	}
	return nil
}

// Configuration for external components.
type ExternalComponentSpec struct {
	// Selects whether this component is installed.
	Enabled *BoolValueForPB `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Namespace for the component.
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Arbitrary install time configuration for the component.
	Spec interface{} `protobuf:"bytes,10,opt,name=spec,proto3" json:"spec,omitempty"`
	// Chart path for addon components.
	ChartPath string `protobuf:"bytes,30,opt,name=chart_path,json=chartPath,proto3" json:"chartPath,omitempty"`
	// Optional schema to validate spec against.
	Schema *any.Any `protobuf:"bytes,35,opt,name=schema,proto3" json:"schema,omitempty"`
	// Kubernetes resource spec.
	K8S                  *KubernetesResourcesSpec `protobuf:"bytes,50,opt,name=k8s,proto3" json:"k8s,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ExternalComponentSpec) Reset()         { *m = ExternalComponentSpec{} }
func (m *ExternalComponentSpec) String() string { return proto.CompactTextString(m) }
func (*ExternalComponentSpec) ProtoMessage()    {}
func (*ExternalComponentSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ed34a579e9b43a2, []int{3}
}

func (m *ExternalComponentSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExternalComponentSpec.Unmarshal(m, b)
}
func (m *ExternalComponentSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExternalComponentSpec.Marshal(b, m, deterministic)
}
func (m *ExternalComponentSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExternalComponentSpec.Merge(m, src)
}
func (m *ExternalComponentSpec) XXX_Size() int {
	return xxx_messageInfo_ExternalComponentSpec.Size(m)
}
func (m *ExternalComponentSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ExternalComponentSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ExternalComponentSpec proto.InternalMessageInfo


func (m *ExternalComponentSpec) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}


func (m *ExternalComponentSpec) GetChartPath() string {
	if m != nil {
		return m.ChartPath
	}
	return ""
}

func (m *ExternalComponentSpec) GetSchema() *any.Any {
	if m != nil {
		return m.Schema
	}
	return nil
}

func (m *ExternalComponentSpec) GetK8S() *KubernetesResourcesSpec {
	if m != nil {
		return m.K8S
	}
	return nil
}

// Configuration for gateways.
type GatewaySpec struct {
	// Selects whether this gateway is installed.
	Enabled *BoolValueForPB `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Namespace for the gateway.
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Name for the gateway.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Labels for the gateway.
	Label map[string]string `protobuf:"bytes,4,rep,name=label,proto3" json:"label,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Hub for the component (overrides top level hub setting).
	Hub string `protobuf:"bytes,10,opt,name=hub,proto3" json:"hub,omitempty"`
	// Tag for the component (overrides top level tag setting).
	Tag interface{} `protobuf:"bytes,11,opt,name=tag,proto3" json:"tag,omitempty"`
	// Kubernetes resource spec.
	K8S                  *KubernetesResourcesSpec `protobuf:"bytes,50,opt,name=k8s,proto3" json:"k8s,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *GatewaySpec) Reset()         { *m = GatewaySpec{} }
func (m *GatewaySpec) String() string { return proto.CompactTextString(m) }
func (*GatewaySpec) ProtoMessage()    {}
func (*GatewaySpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ed34a579e9b43a2, []int{4}
}

func (m *GatewaySpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewaySpec.Unmarshal(m, b)
}
func (m *GatewaySpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewaySpec.Marshal(b, m, deterministic)
}
func (m *GatewaySpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewaySpec.Merge(m, src)
}
func (m *GatewaySpec) XXX_Size() int {
	return xxx_messageInfo_GatewaySpec.Size(m)
}
func (m *GatewaySpec) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewaySpec.DiscardUnknown(m)
}

var xxx_messageInfo_GatewaySpec proto.InternalMessageInfo


func (m *GatewaySpec) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *GatewaySpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GatewaySpec) GetLabel() map[string]string {
	if m != nil {
		return m.Label
	}
	return nil
}

func (m *GatewaySpec) GetHub() string {
	if m != nil {
		return m.Hub
	}
	return ""
}


func (m *GatewaySpec) GetK8S() *KubernetesResourcesSpec {
	if m != nil {
		return m.K8S
	}
	return nil
}

// KubernetesResourcesConfig is a common set of k8s resource configs for components.
type KubernetesResourcesSpec struct {
	// k8s affinity.
	// [https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#affinity-and-anti-affinity](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#affinity-and-anti-affinity)
	Affinity *Affinity `protobuf:"bytes,1,opt,name=affinity,proto3" json:"affinity,omitempty"`
	// Deployment environment variables.
	// [https://kubernetes.io/docs/tasks/inject-data-application/define-environment-variable-container/](https://kubernetes.io/docs/tasks/inject-data-application/define-environment-variable-container/)
	Env []*EnvVar `protobuf:"bytes,2,rep,name=env,proto3" json:"env,omitempty"`
	// k8s HorizontalPodAutoscaler settings.
	// [https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/)
	HpaSpec *HorizontalPodAutoscalerSpec `protobuf:"bytes,3,opt,name=hpa_spec,json=hpaSpec,proto3" json:"hpaSpec,omitempty"`
	// k8s imagePullPolicy.
	// [https://kubernetes.io/docs/concepts/containers/images/](https://kubernetes.io/docs/concepts/containers/images/)
	ImagePullPolicy string `protobuf:"bytes,4,opt,name=image_pull_policy,json=imagePullPolicy,proto3" json:"imagePullPolicy,omitempty"`
	// k8s nodeSelector.
	// [https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#nodeselector](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#nodeselector)
	NodeSelector map[string]string `protobuf:"bytes,5,rep,name=node_selector,json=nodeSelector,proto3" json:"nodeSelector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// k8s PodDisruptionBudget settings.
	// [https://kubernetes.io/docs/concepts/workloads/pods/disruptions/#how-disruption-budgets-work](https://kubernetes.io/docs/concepts/workloads/pods/disruptions/#how-disruption-budgets-work)
	PodDisruptionBudget *PodDisruptionBudgetSpec `protobuf:"bytes,6,opt,name=pod_disruption_budget,json=podDisruptionBudget,proto3" json:"podDisruptionBudget,omitempty"`
	// k8s pod annotations.
	// [https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/](https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/)
	PodAnnotations map[string]string `protobuf:"bytes,7,rep,name=pod_annotations,json=podAnnotations,proto3" json:"podAnnotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// k8s priority_class_name. Default for all resources unless overridden.
	// [https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/#priorityclass](https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/#priorityclass)
	PriorityClassName string `protobuf:"bytes,8,opt,name=priority_class_name,json=priorityClassName,proto3" json:"priorityClassName,omitempty"`
	// k8s readinessProbe settings.
	// [https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-probes/](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-probes/)
	// k8s.io.api.core.v1.Probe readiness_probe = 9;
	ReadinessProbe *ReadinessProbe `protobuf:"bytes,9,opt,name=readiness_probe,json=readinessProbe,proto3" json:"readinessProbe,omitempty"`
	// k8s Deployment replicas setting.
	// [https://kubernetes.io/docs/concepts/workloads/controllers/deployment/](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)
	ReplicaCount uint32 `protobuf:"varint,10,opt,name=replica_count,json=replicaCount,proto3" json:"replicaCount,omitempty"`
	// k8s resources settings.
	// [https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/#resource-requests-and-limits-of-pod-and-container](https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/#resource-requests-and-limits-of-pod-and-container)
	Resources *Resources `protobuf:"bytes,11,opt,name=resources,proto3" json:"resources,omitempty"`
	// k8s Service settings.
	// [https://kubernetes.io/docs/concepts/services-networking/service/](https://kubernetes.io/docs/concepts/services-networking/service/)
	Service *ServiceSpec `protobuf:"bytes,12,opt,name=service,proto3" json:"service,omitempty"`
	// k8s deployment strategy.
	// [https://kubernetes.io/docs/concepts/workloads/controllers/deployment/](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)
	Strategy *DeploymentStrategy `protobuf:"bytes,13,opt,name=strategy,proto3" json:"strategy,omitempty"`
	// k8s toleration
	// [https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/](https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/)
	Tolerations []*v1.Toleration `protobuf:"bytes,14,rep,name=tolerations,proto3" json:"tolerations,omitempty"`
	// k8s service annotations.
	// [https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/](https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/)
	ServiceAnnotations map[string]string `protobuf:"bytes,15,rep,name=service_annotations,json=serviceAnnotations,proto3" json:"serviceAnnotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Overlays for k8s resources in rendered manifests.
	Overlays             []*K8SObjectOverlay `protobuf:"bytes,100,rep,name=overlays,proto3" json:"overlays,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *KubernetesResourcesSpec) Reset()         { *m = KubernetesResourcesSpec{} }
func (m *KubernetesResourcesSpec) String() string { return proto.CompactTextString(m) }
func (*KubernetesResourcesSpec) ProtoMessage()    {}
func (*KubernetesResourcesSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ed34a579e9b43a2, []int{5}
}

func (m *KubernetesResourcesSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KubernetesResourcesSpec.Unmarshal(m, b)
}
func (m *KubernetesResourcesSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KubernetesResourcesSpec.Marshal(b, m, deterministic)
}
func (m *KubernetesResourcesSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KubernetesResourcesSpec.Merge(m, src)
}
func (m *KubernetesResourcesSpec) XXX_Size() int {
	return xxx_messageInfo_KubernetesResourcesSpec.Size(m)
}
func (m *KubernetesResourcesSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_KubernetesResourcesSpec.DiscardUnknown(m)
}

var xxx_messageInfo_KubernetesResourcesSpec proto.InternalMessageInfo

func (m *KubernetesResourcesSpec) GetAffinity() *Affinity {
	if m != nil {
		return m.Affinity
	}
	return nil
}

func (m *KubernetesResourcesSpec) GetEnv() []*EnvVar {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *KubernetesResourcesSpec) GetHpaSpec() *HorizontalPodAutoscalerSpec {
	if m != nil {
		return m.HpaSpec
	}
	return nil
}

func (m *KubernetesResourcesSpec) GetImagePullPolicy() string {
	if m != nil {
		return m.ImagePullPolicy
	}
	return ""
}

func (m *KubernetesResourcesSpec) GetNodeSelector() map[string]string {
	if m != nil {
		return m.NodeSelector
	}
	return nil
}

func (m *KubernetesResourcesSpec) GetPodDisruptionBudget() *PodDisruptionBudgetSpec {
	if m != nil {
		return m.PodDisruptionBudget
	}
	return nil
}

func (m *KubernetesResourcesSpec) GetPodAnnotations() map[string]string {
	if m != nil {
		return m.PodAnnotations
	}
	return nil
}

func (m *KubernetesResourcesSpec) GetPriorityClassName() string {
	if m != nil {
		return m.PriorityClassName
	}
	return ""
}

func (m *KubernetesResourcesSpec) GetReadinessProbe() *ReadinessProbe {
	if m != nil {
		return m.ReadinessProbe
	}
	return nil
}

func (m *KubernetesResourcesSpec) GetReplicaCount() uint32 {
	if m != nil {
		return m.ReplicaCount
	}
	return 0
}

func (m *KubernetesResourcesSpec) GetResources() *Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *KubernetesResourcesSpec) GetService() *ServiceSpec {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *KubernetesResourcesSpec) GetStrategy() *DeploymentStrategy {
	if m != nil {
		return m.Strategy
	}
	return nil
}

func (m *KubernetesResourcesSpec) GetTolerations() []*v1.Toleration {
	if m != nil {
		return m.Tolerations
	}
	return nil
}

func (m *KubernetesResourcesSpec) GetServiceAnnotations() map[string]string {
	if m != nil {
		return m.ServiceAnnotations
	}
	return nil
}

func (m *KubernetesResourcesSpec) GetOverlays() []*K8SObjectOverlay {
	if m != nil {
		return m.Overlays
	}
	return nil
}

// Patch for an existing k8s resource.
type K8SObjectOverlay struct {
	// Resource API version.
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"apiVersion,omitempty"`
	// Resource kind.
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// Name of resource.
	// Namespace is always the component namespace.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// List of patches to apply to resource.
	Patches              []*K8SObjectOverlay_PathValue `protobuf:"bytes,4,rep,name=patches,proto3" json:"patches,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *K8SObjectOverlay) Reset()         { *m = K8SObjectOverlay{} }
func (m *K8SObjectOverlay) String() string { return proto.CompactTextString(m) }
func (*K8SObjectOverlay) ProtoMessage()    {}
func (*K8SObjectOverlay) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ed34a579e9b43a2, []int{6}
}

func (m *K8SObjectOverlay) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_K8SObjectOverlay.Unmarshal(m, b)
}
func (m *K8SObjectOverlay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_K8SObjectOverlay.Marshal(b, m, deterministic)
}
func (m *K8SObjectOverlay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_K8SObjectOverlay.Merge(m, src)
}
func (m *K8SObjectOverlay) XXX_Size() int {
	return xxx_messageInfo_K8SObjectOverlay.Size(m)
}
func (m *K8SObjectOverlay) XXX_DiscardUnknown() {
	xxx_messageInfo_K8SObjectOverlay.DiscardUnknown(m)
}

var xxx_messageInfo_K8SObjectOverlay proto.InternalMessageInfo

func (m *K8SObjectOverlay) GetApiVersion() string {
	if m != nil {
		return m.ApiVersion
	}
	return ""
}

func (m *K8SObjectOverlay) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *K8SObjectOverlay) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *K8SObjectOverlay) GetPatches() []*K8SObjectOverlay_PathValue {
	if m != nil {
		return m.Patches
	}
	return nil
}

type K8SObjectOverlay_PathValue struct {
	// Path of the form a.[key1:value1].b.[:value2]
	// Where [key1:value1] is a selector for a key-value pair to identify a list element and [:value] is a value
	// selector to identify a list element in a leaf list.
	// All path intermediate nodes must exist.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Value to add, delete or replace.
	// For add, the path should be a new leaf.
	// For delete, value should be unset.
	// For replace, path should reference an existing node.
	// All values are strings but are converted into appropriate type based on schema.
	Value                interface{} `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *K8SObjectOverlay_PathValue) Reset()         { *m = K8SObjectOverlay_PathValue{} }
func (m *K8SObjectOverlay_PathValue) String() string { return proto.CompactTextString(m) }
func (*K8SObjectOverlay_PathValue) ProtoMessage()    {}
func (*K8SObjectOverlay_PathValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ed34a579e9b43a2, []int{6, 0}
}

func (m *K8SObjectOverlay_PathValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_K8SObjectOverlay_PathValue.Unmarshal(m, b)
}
func (m *K8SObjectOverlay_PathValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_K8SObjectOverlay_PathValue.Marshal(b, m, deterministic)
}
func (m *K8SObjectOverlay_PathValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_K8SObjectOverlay_PathValue.Merge(m, src)
}
func (m *K8SObjectOverlay_PathValue) XXX_Size() int {
	return xxx_messageInfo_K8SObjectOverlay_PathValue.Size(m)
}
func (m *K8SObjectOverlay_PathValue) XXX_DiscardUnknown() {
	xxx_messageInfo_K8SObjectOverlay_PathValue.DiscardUnknown(m)
}

var xxx_messageInfo_K8SObjectOverlay_PathValue proto.InternalMessageInfo

func (m *K8SObjectOverlay_PathValue) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}





func init() {
	proto.RegisterType((*IstioComponentSetSpec)(nil), "istio.operator.v1alpha1.IstioComponentSetSpec")
	proto.RegisterType((*BaseComponentSpec)(nil), "istio.operator.v1alpha1.BaseComponentSpec")
	proto.RegisterType((*ComponentSpec)(nil), "istio.operator.v1alpha1.ComponentSpec")
	proto.RegisterType((*ExternalComponentSpec)(nil), "istio.operator.v1alpha1.ExternalComponentSpec")
	proto.RegisterType((*GatewaySpec)(nil), "istio.operator.v1alpha1.GatewaySpec")
	proto.RegisterMapType((map[string]string)(nil), "istio.operator.v1alpha1.GatewaySpec.LabelEntry")
	proto.RegisterType((*KubernetesResourcesSpec)(nil), "istio.operator.v1alpha1.KubernetesResourcesSpec")
	proto.RegisterMapType((map[string]string)(nil), "istio.operator.v1alpha1.KubernetesResourcesSpec.NodeSelectorEntry")
	proto.RegisterMapType((map[string]string)(nil), "istio.operator.v1alpha1.KubernetesResourcesSpec.PodAnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "istio.operator.v1alpha1.KubernetesResourcesSpec.ServiceAnnotationsEntry")
	proto.RegisterType((*K8SObjectOverlay)(nil), "istio.operator.v1alpha1.K8sObjectOverlay")
	proto.RegisterType((*K8SObjectOverlay_PathValue)(nil), "istio.operator.v1alpha1.K8sObjectOverlay.PathValue")
}

func init() { proto.RegisterFile("operator/v1alpha1/component.proto", fileDescriptor_6ed34a579e9b43a2) }

var fileDescriptor_6ed34a579e9b43a2 = []byte{
	// 1264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x97, 0xdd, 0x72, 0x1b, 0x35,
	0x14, 0xc7, 0x27, 0x71, 0xbe, 0x7c, 0xf2, 0xe1, 0x44, 0x69, 0xe9, 0x92, 0xa1, 0x6d, 0xea, 0x96,
	0x36, 0x2d, 0xcc, 0x9a, 0xb4, 0x5c, 0x78, 0x3a, 0x50, 0x1a, 0x37, 0xa6, 0xed, 0x94, 0x36, 0x66,
	0xd3, 0xe9, 0x05, 0x33, 0xcc, 0x8e, 0xbc, 0x7b, 0x62, 0x8b, 0xac, 0x25, 0x8d, 0x24, 0x9b, 0x2e,
	0x2f, 0xc0, 0xf0, 0x04, 0xdc, 0x30, 0x3c, 0x0b, 0xb7, 0xbc, 0x15, 0x23, 0xed, 0xda, 0xae, 0xeb,
	0x9a, 0x74, 0x99, 0x00, 0x77, 0xbb, 0x67, 0xcf, 0xef, 0x2f, 0x9d, 0xb3, 0x47, 0x3a, 0x12, 0x5c,
	0x13, 0x12, 0x15, 0x35, 0x42, 0xd5, 0x06, 0xfb, 0x34, 0x91, 0x5d, 0xba, 0x5f, 0x8b, 0x44, 0x4f,
	0x0a, 0x8e, 0xdc, 0xf8, 0x52, 0x09, 0x23, 0xc8, 0x25, 0xa6, 0x0d, 0x13, 0xfe, 0xd0, 0xd1, 0x1f,
	0x3a, 0xee, 0x7c, 0xd8, 0x11, 0xa2, 0x93, 0x60, 0xcd, 0xb9, 0xb5, 0xfb, 0x27, 0x35, 0xca, 0xd3,
	0x8c, 0xd9, 0xa9, 0x9e, 0xd6, 0xb5, 0xcf, 0x44, 0x8d, 0x4a, 0x56, 0x8b, 0x84, 0xc2, 0xda, 0x60,
	0xbf, 0xd6, 0x41, 0x6e, 0x15, 0x30, 0x1e, 0xfa, 0x4c, 0x0f, 0x7d, 0xda, 0x6f, 0xa3, 0xe2, 0x68,
	0x50, 0x67, 0x3e, 0xd5, 0xdf, 0x96, 0xe1, 0xe2, 0x53, 0x3b, 0xfc, 0xa3, 0xe1, 0xa4, 0x8e, 0xd1,
	0x1c, 0x4b, 0x8c, 0xc8, 0x03, 0x58, 0x68, 0x53, 0x8d, 0xde, 0xe5, 0xdd, 0xb9, 0xbd, 0xd5, 0xbb,
	0x77, 0xfc, 0x19, 0x93, 0xf4, 0x1b, 0x54, 0xe3, 0x18, 0x96, 0x18, 0x05, 0x8e, 0x23, 0x5f, 0xc0,
	0xa2, 0x64, 0x89, 0x30, 0xde, 0x15, 0x27, 0x70, 0x73, 0xa6, 0xc0, 0x24, 0x9c, 0x41, 0x8e, 0x56,
	0xe2, 0x75, 0xea, 0x5d, 0x2d, 0x48, 0x5b, 0x88, 0x7c, 0x0b, 0x9b, 0x9a, 0xc5, 0x18, 0x51, 0x15,
	0x32, 0xfe, 0x03, 0x46, 0x46, 0x28, 0x6f, 0xb7, 0x90, 0x50, 0x25, 0xe7, 0x9f, 0xe6, 0x38, 0x79,
	0x00, 0x4b, 0x52, 0x24, 0x2c, 0x4a, 0xbd, 0x6b, 0x85, 0x84, 0x72, 0x8a, 0x1c, 0x42, 0xd9, 0x60,
	0x82, 0x3d, 0x34, 0x2a, 0xf5, 0xaa, 0x85, 0x24, 0xc6, 0x20, 0x79, 0x08, 0xcb, 0x11, 0x33, 0x34,
	0xc6, 0xc4, 0xbb, 0x5e, 0x48, 0x63, 0x88, 0x91, 0x26, 0x00, 0x17, 0x31, 0x86, 0xb4, 0x83, 0xdc,
	0x78, 0x37, 0x8a, 0x4d, 0xc4, 0x92, 0x07, 0x16, 0xb4, 0xe9, 0xe8, 0xd0, 0x24, 0xc1, 0xd4, 0xfb,
	0xb8, 0x58, 0x3a, 0x32, 0x8a, 0xd4, 0xa1, 0x14, 0x71, 0xe6, 0xdd, 0x2c, 0x04, 0x5b, 0x84, 0x3c,
	0x83, 0x75, 0xe7, 0x1d, 0x87, 0x0a, 0x7b, 0xc2, 0xa0, 0x77, 0xab, 0x90, 0xc6, 0x5a, 0x06, 0x07,
	0x8e, 0x25, 0x47, 0xb0, 0xc9, 0x78, 0x47, 0xa1, 0xd6, 0x61, 0x87, 0x1a, 0xfc, 0x91, 0xa6, 0xda,
	0xdb, 0xdb, 0x2d, 0xed, 0xad, 0xde, 0xbd, 0x31, 0x53, 0xef, 0x71, 0xe6, 0x98, 0x95, 0x49, 0x4e,
	0xe7, 0x36, 0x4d, 0x9e, 0x43, 0x05, 0xdf, 0xd2, 0xbb, 0x5d, 0x40, 0x6f, 0x03, 0x27, 0xe4, 0xaa,
	0xbf, 0xcf, 0xc1, 0xd6, 0xd4, 0x02, 0x23, 0x4d, 0x58, 0x46, 0x4e, 0xdb, 0x09, 0xc6, 0xde, 0x9c,
	0x0b, 0xfe, 0x93, 0x99, 0xe2, 0x2f, 0x53, 0x89, 0x0d, 0x21, 0x92, 0x57, 0x34, 0xe9, 0xe3, 0xd7,
	0x42, 0xb5, 0x1a, 0xc1, 0x90, 0x25, 0x0d, 0x28, 0x9d, 0xd6, 0xb5, 0x77, 0xd7, 0x49, 0x7c, 0x36,
	0x53, 0xe2, 0xd9, 0x68, 0xcf, 0x08, 0x50, 0x8b, 0xbe, 0x8a, 0x50, 0x67, 0x7f, 0xe3, 0xb4, 0xae,
	0xab, 0x7f, 0xcc, 0xc3, 0xfa, 0xbf, 0x32, 0xb9, 0x8f, 0xa0, 0xcc, 0x69, 0x0f, 0xb5, 0xa4, 0x11,
	0x7a, 0xf3, 0xbb, 0x73, 0x7b, 0xe5, 0x60, 0x6c, 0x20, 0x9b, 0x50, 0xea, 0xf6, 0xdb, 0x1e, 0x38,
	0xbb, 0x7d, 0xb4, 0x05, 0x65, 0x68, 0xc7, 0x5b, 0x3d, 0xa3, 0x18, 0xec, 0x90, 0x4f, 0xb9, 0x41,
	0x75, 0x42, 0x23, 0x0c, 0x2c, 0x42, 0xee, 0xc3, 0x82, 0x96, 0x18, 0x9d, 0xb9, 0x4f, 0x4d, 0xa2,
	0x8e, 0x39, 0x97, 0x14, 0xfe, 0x39, 0x0f, 0x17, 0x9b, 0xaf, 0x0d, 0x2a, 0x4e, 0x93, 0xff, 0x21,
	0x95, 0xc3, 0xf0, 0xe1, 0x1f, 0x84, 0x7f, 0x19, 0x20, 0xea, 0x52, 0x65, 0x42, 0x49, 0x4d, 0xd7,
	0x25, 0xb0, 0x1c, 0x94, 0x9d, 0xa5, 0x45, 0x4d, 0x97, 0x7c, 0x0a, 0x4b, 0x3a, 0xea, 0x62, 0x8f,
	0xe6, 0x9b, 0xd5, 0x05, 0x3f, 0x6b, 0x68, 0xfe, 0xb0, 0xa1, 0xf9, 0x07, 0x3c, 0x0d, 0x72, 0x9f,
	0x73, 0xc9, 0xe5, 0xaf, 0x25, 0x58, 0x7d, 0x63, 0x3d, 0xfd, 0x37, 0x19, 0x24, 0xb0, 0x60, 0x5f,
	0xbc, 0x92, 0xfb, 0xe0, 0x9e, 0x49, 0x13, 0x16, 0x13, 0xda, 0xc6, 0xc4, 0x5b, 0x70, 0xab, 0xbf,
	0xf6, 0x3e, 0xab, 0xdf, 0xff, 0xc6, 0x12, 0x4d, 0x6e, 0x54, 0x1a, 0x64, 0xf4, 0xb9, 0xd6, 0xf9,
	0x39, 0xe4, 0x77, 0xa7, 0x0e, 0x30, 0x9e, 0xa4, 0x9d, 0xdd, 0x29, 0xa6, 0x2e, 0xb3, 0xe5, 0xc0,
	0x3e, 0x92, 0x0b, 0xb0, 0x38, 0xb0, 0xf9, 0xcb, 0x93, 0x94, 0xbd, 0xdc, 0x9f, 0xaf, 0xcf, 0x55,
	0x7f, 0x5e, 0x85, 0x4b, 0x33, 0xa4, 0xc9, 0x97, 0xb0, 0x42, 0x4f, 0x4e, 0x18, 0x67, 0x26, 0xcd,
	0x7f, 0xd3, 0xb5, 0x99, 0xd3, 0x3b, 0xc8, 0x1d, 0x83, 0x11, 0x42, 0xf6, 0xa1, 0x84, 0x7c, 0xe0,
	0xcd, 0xbb, 0x4c, 0x5f, 0x9d, 0x49, 0x36, 0xf9, 0xe0, 0x15, 0x55, 0x81, 0xf5, 0x25, 0x47, 0xb0,
	0xd2, 0x95, 0x34, 0x74, 0x85, 0x5f, 0x72, 0x23, 0x7e, 0x3e, 0x93, 0x7b, 0x22, 0x14, 0xfb, 0x49,
	0x70, 0x43, 0x93, 0x96, 0x88, 0x0f, 0xfa, 0x46, 0xe8, 0x88, 0x26, 0xa8, 0xb2, 0xb6, 0xda, 0x95,
	0xd4, 0x85, 0x70, 0x07, 0xb6, 0x58, 0x8f, 0x76, 0x30, 0x94, 0xfd, 0x24, 0x09, 0xf3, 0x93, 0xc2,
	0x82, 0x4b, 0x42, 0xc5, 0x7d, 0x68, 0xf5, 0x93, 0xa4, 0x95, 0x1d, 0x05, 0x3a, 0xb0, 0xee, 0x5a,
	0xb0, 0xc6, 0x24, 0x3b, 0x9a, 0x2c, 0xba, 0x99, 0x37, 0x8a, 0xfe, 0x12, 0xff, 0x85, 0x88, 0xf1,
	0x38, 0x17, 0xc9, 0xca, 0x66, 0x8d, 0xbf, 0x61, 0x22, 0x31, 0x5c, 0x94, 0x22, 0x0e, 0x63, 0xa6,
	0x55, 0x5f, 0x1a, 0x26, 0x78, 0xd8, 0xee, 0xc7, 0x1d, 0x34, 0xde, 0xd2, 0x19, 0x35, 0xd0, 0x12,
	0xf1, 0xe1, 0x08, 0x6a, 0x38, 0xc6, 0x85, 0xbb, 0x2d, 0xa7, 0x3f, 0x90, 0x1e, 0x54, 0xec, 0x28,
	0x94, 0x73, 0x61, 0xa8, 0xb5, 0x6b, 0x6f, 0xd9, 0x05, 0x74, 0x58, 0x38, 0x20, 0x9b, 0xe0, 0xb1,
	0x4c, 0x16, 0xd2, 0x86, 0x9c, 0x30, 0x12, 0x1f, 0xb6, 0xa5, 0x62, 0x42, 0x31, 0x93, 0x86, 0x51,
	0x42, 0xb5, 0x0e, 0xdd, 0xe2, 0x5b, 0x71, 0xb9, 0xde, 0x1a, 0x7e, 0x7a, 0x64, 0xbf, 0xbc, 0xb0,
	0x2b, 0xb1, 0x05, 0x15, 0x85, 0x34, 0x66, 0xdc, 0x36, 0x65, 0xa9, 0x44, 0x1b, 0xbd, 0xb2, 0x0b,
	0xff, 0xd6, 0xcc, 0xe9, 0x05, 0x43, 0xff, 0x96, 0x75, 0x0f, 0x36, 0xd4, 0xc4, 0x3b, 0xb9, 0x0e,
	0xeb, 0x0a, 0x65, 0xc2, 0x22, 0x1a, 0x46, 0xa2, 0xcf, 0x8d, 0x5b, 0x9e, 0xeb, 0xc1, 0x5a, 0x6e,
	0x7c, 0x64, 0x6d, 0xe4, 0x21, 0x94, 0xd5, 0x30, 0xb6, 0x7c, 0xb5, 0x56, 0xff, 0x66, 0xc0, 0xdc,
	0x33, 0x18, 0x43, 0xe4, 0x01, 0x2c, 0x6b, 0x54, 0x03, 0x16, 0xa1, 0xb7, 0xe6, 0xf8, 0xd9, 0x47,
	0x88, 0xe3, 0xcc, 0x2f, 0x2b, 0xc9, 0x1c, 0x22, 0x8f, 0x61, 0x45, 0x1b, 0x7b, 0x1f, 0xe8, 0xa4,
	0xde, 0xfa, 0x19, 0x9b, 0xdf, 0x21, 0xca, 0x44, 0xa4, 0x3d, 0xdb, 0x78, 0x72, 0x24, 0x18, 0xc1,
	0xe4, 0x21, 0xac, 0x1a, 0x91, 0x58, 0xc4, 0xfd, 0xdc, 0x0d, 0xf7, 0x73, 0xaf, 0xf8, 0xd9, 0x0d,
	0xc4, 0xa7, 0x92, 0xf9, 0xf6, 0x06, 0xe2, 0x0f, 0xf6, 0xfd, 0x97, 0x23, 0xb7, 0xe0, 0x4d, 0x84,
	0xa4, 0xb0, 0x9d, 0xcf, 0x6a, 0xa2, 0x4c, 0x2a, 0x4e, 0xe9, 0x49, 0xe1, 0x32, 0xc9, 0xc3, 0x9d,
	0x2a, 0x15, 0xa2, 0xa7, 0x3e, 0x90, 0x26, 0xac, 0x88, 0x01, 0xaa, 0xc4, 0x9e, 0xc4, 0x62, 0x37,
	0xde, 0xed, 0xd9, 0xe3, 0xd5, 0xf5, 0x51, 0xdb, 0x9e, 0xf7, 0x8f, 0x32, 0x22, 0x18, 0xa1, 0x3b,
	0x5f, 0xc1, 0xd6, 0xd4, 0x6a, 0x2b, 0xb2, 0xff, 0xed, 0x1c, 0xc0, 0xf6, 0x3b, 0xaa, 0xbb, 0x90,
	0x44, 0x13, 0x2e, 0xcd, 0x88, 0xbc, 0xd0, 0x4e, 0xfc, 0xcb, 0x3c, 0x6c, 0xbe, 0x1d, 0x29, 0xb9,
	0x0a, 0xab, 0x54, 0xb2, 0x70, 0x80, 0x4a, 0x33, 0xc1, 0x73, 0x21, 0xa0, 0x92, 0xbd, 0xca, 0x2c,
	0xb6, 0xc9, 0x9d, 0x32, 0x1e, 0xe7, 0x72, 0xee, 0xf9, 0x9d, 0x8d, 0xef, 0x39, 0x2c, 0x4b, 0x6a,
	0xa2, 0x2e, 0xea, 0xbc, 0xf5, 0xdd, 0x7b, 0xef, 0x74, 0xfb, 0xf6, 0xd0, 0xe0, 0xda, 0x6f, 0x30,
	0xd4, 0xd8, 0xf9, 0x1e, 0xca, 0x23, 0xab, 0x1d, 0xcf, 0x1d, 0x34, 0xb2, 0xd9, 0xb9, 0x67, 0x7b,
	0x51, 0x1c, 0xc7, 0xf9, 0xfe, 0x1d, 0x31, 0x83, 0xaa, 0x1e, 0x7c, 0x60, 0xed, 0xcf, 0xa9, 0x3c,
	0x36, 0x8a, 0xf1, 0xce, 0xc8, 0xa1, 0x5a, 0x81, 0xf5, 0x09, 0xa2, 0x7a, 0x01, 0xc8, 0xf4, 0x11,
	0xa1, 0xb1, 0xfb, 0xdd, 0x95, 0x6c, 0xc0, 0xfc, 0x2e, 0x3e, 0x75, 0xe5, 0x6e, 0x2f, 0xb9, 0xc3,
	0xce, 0xbd, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x7d, 0x25, 0x54, 0x09, 0x10, 0x00, 0x00,
}