// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/k8s/v1/node_group.proto

package k8s

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

type IpVersion int32

const (
	IpVersion_IP_VERSION_UNSPECIFIED IpVersion = 0
	// IPv4 address, for example 192.168.0.0.
	IpVersion_IPV4 IpVersion = 1
	// IPv6 address, not available yet.
	IpVersion_IPV6 IpVersion = 2
)

var IpVersion_name = map[int32]string{
	0: "IP_VERSION_UNSPECIFIED",
	1: "IPV4",
	2: "IPV6",
}

var IpVersion_value = map[string]int32{
	"IP_VERSION_UNSPECIFIED": 0,
	"IPV4":                   1,
	"IPV6":                   2,
}

func (x IpVersion) String() string {
	return proto.EnumName(IpVersion_name, int32(x))
}

func (IpVersion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d5c271823cf8486a, []int{0}
}

type NodeGroup_Status int32

const (
	NodeGroup_STATUS_UNSPECIFIED NodeGroup_Status = 0
	// Node group is waiting for resources to be allocated.
	NodeGroup_PROVISIONING NodeGroup_Status = 1
	// Node group is running.
	NodeGroup_RUNNING NodeGroup_Status = 2
	// Node group is waiting for some work to be done, such as upgrading node software.
	NodeGroup_RECONCILING NodeGroup_Status = 3
	// Node group is being stopped.
	NodeGroup_STOPPING NodeGroup_Status = 4
	// Node group stopped.
	NodeGroup_STOPPED NodeGroup_Status = 5
	// Node group is being deleted.
	NodeGroup_DELETING NodeGroup_Status = 6
)

var NodeGroup_Status_name = map[int32]string{
	0: "STATUS_UNSPECIFIED",
	1: "PROVISIONING",
	2: "RUNNING",
	3: "RECONCILING",
	4: "STOPPING",
	5: "STOPPED",
	6: "DELETING",
}

var NodeGroup_Status_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"PROVISIONING":       1,
	"RUNNING":            2,
	"RECONCILING":        3,
	"STOPPING":           4,
	"STOPPED":            5,
	"DELETING":           6,
}

func (x NodeGroup_Status) String() string {
	return proto.EnumName(NodeGroup_Status_name, int32(x))
}

func (NodeGroup_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d5c271823cf8486a, []int{0, 0}
}

type NodeGroup struct {
	// ID of the node group.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the cluster that the node group belongs to.
	ClusterId string `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Creation timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the node group.
	// The name is unique within the folder.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the node group. 0-256 characters long.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Resource labels as `key:value` pairs. Мaximum of 64 per resource.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Status of the node group.
	Status NodeGroup_Status `protobuf:"varint,7,opt,name=status,proto3,enum=yandex.cloud.k8s.v1.NodeGroup_Status" json:"status,omitempty"`
	// Node template that specifies parameters of the compute instances for the node group.
	NodeTemplate *NodeTemplate `protobuf:"bytes,8,opt,name=node_template,json=nodeTemplate,proto3" json:"node_template,omitempty"`
	// Scale policy of the node group.  For more information, see [Scaling policy](/docs/compute/concepts/instance-groups/policies#scale-policy).
	ScalePolicy *ScalePolicy `protobuf:"bytes,9,opt,name=scale_policy,json=scalePolicy,proto3" json:"scale_policy,omitempty"`
	// Allocation policy by which resources for node group are allocated to zones and regions.
	AllocationPolicy *NodeGroupAllocationPolicy `protobuf:"bytes,10,opt,name=allocation_policy,json=allocationPolicy,proto3" json:"allocation_policy,omitempty"`
	// ID of the managed instance group associated with this node group.
	InstanceGroupId string `protobuf:"bytes,11,opt,name=instance_group_id,json=instanceGroupId,proto3" json:"instance_group_id,omitempty"`
	// Version of Kubernetes components that runs on the nodes.
	// Deprecated. Use version_info.current_version.
	NodeVersion          string                      `protobuf:"bytes,12,opt,name=node_version,json=nodeVersion,proto3" json:"node_version,omitempty"`
	VersionInfo          *VersionInfo                `protobuf:"bytes,13,opt,name=version_info,json=versionInfo,proto3" json:"version_info,omitempty"`
	MaintenancePolicy    *NodeGroupMaintenancePolicy `protobuf:"bytes,14,opt,name=maintenance_policy,json=maintenancePolicy,proto3" json:"maintenance_policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *NodeGroup) Reset()         { *m = NodeGroup{} }
func (m *NodeGroup) String() string { return proto.CompactTextString(m) }
func (*NodeGroup) ProtoMessage()    {}
func (*NodeGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5c271823cf8486a, []int{0}
}

func (m *NodeGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeGroup.Unmarshal(m, b)
}
func (m *NodeGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeGroup.Marshal(b, m, deterministic)
}
func (m *NodeGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeGroup.Merge(m, src)
}
func (m *NodeGroup) XXX_Size() int {
	return xxx_messageInfo_NodeGroup.Size(m)
}
func (m *NodeGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeGroup.DiscardUnknown(m)
}

var xxx_messageInfo_NodeGroup proto.InternalMessageInfo

func (m *NodeGroup) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NodeGroup) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *NodeGroup) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *NodeGroup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NodeGroup) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *NodeGroup) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *NodeGroup) GetStatus() NodeGroup_Status {
	if m != nil {
		return m.Status
	}
	return NodeGroup_STATUS_UNSPECIFIED
}

func (m *NodeGroup) GetNodeTemplate() *NodeTemplate {
	if m != nil {
		return m.NodeTemplate
	}
	return nil
}

func (m *NodeGroup) GetScalePolicy() *ScalePolicy {
	if m != nil {
		return m.ScalePolicy
	}
	return nil
}

func (m *NodeGroup) GetAllocationPolicy() *NodeGroupAllocationPolicy {
	if m != nil {
		return m.AllocationPolicy
	}
	return nil
}

func (m *NodeGroup) GetInstanceGroupId() string {
	if m != nil {
		return m.InstanceGroupId
	}
	return ""
}

func (m *NodeGroup) GetNodeVersion() string {
	if m != nil {
		return m.NodeVersion
	}
	return ""
}

func (m *NodeGroup) GetVersionInfo() *VersionInfo {
	if m != nil {
		return m.VersionInfo
	}
	return nil
}

func (m *NodeGroup) GetMaintenancePolicy() *NodeGroupMaintenancePolicy {
	if m != nil {
		return m.MaintenancePolicy
	}
	return nil
}

type NodeTemplate struct {
	// ID of the hardware platform configuration for the node.
	PlatformId string `protobuf:"bytes,1,opt,name=platform_id,json=platformId,proto3" json:"platform_id,omitempty"`
	// Computing resources of the node such as the amount of memory and number of cores.
	ResourcesSpec *ResourcesSpec `protobuf:"bytes,2,opt,name=resources_spec,json=resourcesSpec,proto3" json:"resources_spec,omitempty"`
	// The metadata as `key:value` pairs assigned to this instance template. This includes custom metadata and predefined keys.
	//
	// For example, you may use the metadata in order to provide your public SSH key to the node.
	// For more information, see [Metadata](/docs/compute/concepts/vm-metadata).
	BootDiskSpec *DiskSpec         `protobuf:"bytes,3,opt,name=boot_disk_spec,json=bootDiskSpec,proto3" json:"boot_disk_spec,omitempty"`
	Metadata     map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Specification for the create network interfaces for the node group compute instances.
	V4AddressSpec *NodeAddressSpec `protobuf:"bytes,5,opt,name=v4_address_spec,json=v4AddressSpec,proto3" json:"v4_address_spec,omitempty"`
	// Scheduling policy configuration.
	SchedulingPolicy     *SchedulingPolicy `protobuf:"bytes,6,opt,name=scheduling_policy,json=schedulingPolicy,proto3" json:"scheduling_policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *NodeTemplate) Reset()         { *m = NodeTemplate{} }
func (m *NodeTemplate) String() string { return proto.CompactTextString(m) }
func (*NodeTemplate) ProtoMessage()    {}
func (*NodeTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5c271823cf8486a, []int{1}
}

func (m *NodeTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeTemplate.Unmarshal(m, b)
}
func (m *NodeTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeTemplate.Marshal(b, m, deterministic)
}
func (m *NodeTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeTemplate.Merge(m, src)
}
func (m *NodeTemplate) XXX_Size() int {
	return xxx_messageInfo_NodeTemplate.Size(m)
}
func (m *NodeTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_NodeTemplate proto.InternalMessageInfo

func (m *NodeTemplate) GetPlatformId() string {
	if m != nil {
		return m.PlatformId
	}
	return ""
}

func (m *NodeTemplate) GetResourcesSpec() *ResourcesSpec {
	if m != nil {
		return m.ResourcesSpec
	}
	return nil
}

func (m *NodeTemplate) GetBootDiskSpec() *DiskSpec {
	if m != nil {
		return m.BootDiskSpec
	}
	return nil
}

func (m *NodeTemplate) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *NodeTemplate) GetV4AddressSpec() *NodeAddressSpec {
	if m != nil {
		return m.V4AddressSpec
	}
	return nil
}

func (m *NodeTemplate) GetSchedulingPolicy() *SchedulingPolicy {
	if m != nil {
		return m.SchedulingPolicy
	}
	return nil
}

type NodeAddressSpec struct {
	// One-to-one NAT configuration. Setting up one-to-one NAT ensures that public IP addresses are assigned to nodes, and therefore internet is accessible for all nodes of the node group. If the field is not set, NAT will not be set up.
	OneToOneNatSpec      *OneToOneNatSpec `protobuf:"bytes,1,opt,name=one_to_one_nat_spec,json=oneToOneNatSpec,proto3" json:"one_to_one_nat_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *NodeAddressSpec) Reset()         { *m = NodeAddressSpec{} }
func (m *NodeAddressSpec) String() string { return proto.CompactTextString(m) }
func (*NodeAddressSpec) ProtoMessage()    {}
func (*NodeAddressSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5c271823cf8486a, []int{2}
}

func (m *NodeAddressSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeAddressSpec.Unmarshal(m, b)
}
func (m *NodeAddressSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeAddressSpec.Marshal(b, m, deterministic)
}
func (m *NodeAddressSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeAddressSpec.Merge(m, src)
}
func (m *NodeAddressSpec) XXX_Size() int {
	return xxx_messageInfo_NodeAddressSpec.Size(m)
}
func (m *NodeAddressSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeAddressSpec.DiscardUnknown(m)
}

var xxx_messageInfo_NodeAddressSpec proto.InternalMessageInfo

func (m *NodeAddressSpec) GetOneToOneNatSpec() *OneToOneNatSpec {
	if m != nil {
		return m.OneToOneNatSpec
	}
	return nil
}

type OneToOneNatSpec struct {
	// IP version for the public IP address.
	IpVersion            IpVersion `protobuf:"varint,1,opt,name=ip_version,json=ipVersion,proto3,enum=yandex.cloud.k8s.v1.IpVersion" json:"ip_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *OneToOneNatSpec) Reset()         { *m = OneToOneNatSpec{} }
func (m *OneToOneNatSpec) String() string { return proto.CompactTextString(m) }
func (*OneToOneNatSpec) ProtoMessage()    {}
func (*OneToOneNatSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5c271823cf8486a, []int{3}
}

func (m *OneToOneNatSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OneToOneNatSpec.Unmarshal(m, b)
}
func (m *OneToOneNatSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OneToOneNatSpec.Marshal(b, m, deterministic)
}
func (m *OneToOneNatSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OneToOneNatSpec.Merge(m, src)
}
func (m *OneToOneNatSpec) XXX_Size() int {
	return xxx_messageInfo_OneToOneNatSpec.Size(m)
}
func (m *OneToOneNatSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_OneToOneNatSpec.DiscardUnknown(m)
}

var xxx_messageInfo_OneToOneNatSpec proto.InternalMessageInfo

func (m *OneToOneNatSpec) GetIpVersion() IpVersion {
	if m != nil {
		return m.IpVersion
	}
	return IpVersion_IP_VERSION_UNSPECIFIED
}

type ResourcesSpec struct {
	// Amount of memory available to the node, specified in bytes.
	Memory int64 `protobuf:"varint,1,opt,name=memory,proto3" json:"memory,omitempty"`
	// Number of cores available to the node.
	Cores int64 `protobuf:"varint,2,opt,name=cores,proto3" json:"cores,omitempty"`
	// Baseline level of CPU performance with the possibility to burst performance above that baseline level.
	// This field sets baseline performance for each core.
	CoreFraction         int64    `protobuf:"varint,3,opt,name=core_fraction,json=coreFraction,proto3" json:"core_fraction,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourcesSpec) Reset()         { *m = ResourcesSpec{} }
func (m *ResourcesSpec) String() string { return proto.CompactTextString(m) }
func (*ResourcesSpec) ProtoMessage()    {}
func (*ResourcesSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5c271823cf8486a, []int{4}
}

func (m *ResourcesSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourcesSpec.Unmarshal(m, b)
}
func (m *ResourcesSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourcesSpec.Marshal(b, m, deterministic)
}
func (m *ResourcesSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourcesSpec.Merge(m, src)
}
func (m *ResourcesSpec) XXX_Size() int {
	return xxx_messageInfo_ResourcesSpec.Size(m)
}
func (m *ResourcesSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourcesSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ResourcesSpec proto.InternalMessageInfo

func (m *ResourcesSpec) GetMemory() int64 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *ResourcesSpec) GetCores() int64 {
	if m != nil {
		return m.Cores
	}
	return 0
}

func (m *ResourcesSpec) GetCoreFraction() int64 {
	if m != nil {
		return m.CoreFraction
	}
	return 0
}

type DiskSpec struct {
	// ID of the disk type.
	DiskTypeId string `protobuf:"bytes,1,opt,name=disk_type_id,json=diskTypeId,proto3" json:"disk_type_id,omitempty"`
	// Size of the disk, specified in bytes.
	DiskSize             int64    `protobuf:"varint,2,opt,name=disk_size,json=diskSize,proto3" json:"disk_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiskSpec) Reset()         { *m = DiskSpec{} }
func (m *DiskSpec) String() string { return proto.CompactTextString(m) }
func (*DiskSpec) ProtoMessage()    {}
func (*DiskSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5c271823cf8486a, []int{5}
}

func (m *DiskSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiskSpec.Unmarshal(m, b)
}
func (m *DiskSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiskSpec.Marshal(b, m, deterministic)
}
func (m *DiskSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiskSpec.Merge(m, src)
}
func (m *DiskSpec) XXX_Size() int {
	return xxx_messageInfo_DiskSpec.Size(m)
}
func (m *DiskSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_DiskSpec.DiscardUnknown(m)
}

var xxx_messageInfo_DiskSpec proto.InternalMessageInfo

func (m *DiskSpec) GetDiskTypeId() string {
	if m != nil {
		return m.DiskTypeId
	}
	return ""
}

func (m *DiskSpec) GetDiskSize() int64 {
	if m != nil {
		return m.DiskSize
	}
	return 0
}

type ScalePolicy struct {
	// Types that are valid to be assigned to ScaleType:
	//	*ScalePolicy_FixedScale_
	ScaleType            isScalePolicy_ScaleType `protobuf_oneof:"scale_type"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ScalePolicy) Reset()         { *m = ScalePolicy{} }
func (m *ScalePolicy) String() string { return proto.CompactTextString(m) }
func (*ScalePolicy) ProtoMessage()    {}
func (*ScalePolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5c271823cf8486a, []int{6}
}

func (m *ScalePolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScalePolicy.Unmarshal(m, b)
}
func (m *ScalePolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScalePolicy.Marshal(b, m, deterministic)
}
func (m *ScalePolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScalePolicy.Merge(m, src)
}
func (m *ScalePolicy) XXX_Size() int {
	return xxx_messageInfo_ScalePolicy.Size(m)
}
func (m *ScalePolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_ScalePolicy.DiscardUnknown(m)
}

var xxx_messageInfo_ScalePolicy proto.InternalMessageInfo

type isScalePolicy_ScaleType interface {
	isScalePolicy_ScaleType()
}

type ScalePolicy_FixedScale_ struct {
	FixedScale *ScalePolicy_FixedScale `protobuf:"bytes,1,opt,name=fixed_scale,json=fixedScale,proto3,oneof"`
}

func (*ScalePolicy_FixedScale_) isScalePolicy_ScaleType() {}

func (m *ScalePolicy) GetScaleType() isScalePolicy_ScaleType {
	if m != nil {
		return m.ScaleType
	}
	return nil
}

func (m *ScalePolicy) GetFixedScale() *ScalePolicy_FixedScale {
	if x, ok := m.GetScaleType().(*ScalePolicy_FixedScale_); ok {
		return x.FixedScale
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ScalePolicy) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ScalePolicy_FixedScale_)(nil),
	}
}

type ScalePolicy_FixedScale struct {
	// Number of nodes in the node group.
	Size                 int64    `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScalePolicy_FixedScale) Reset()         { *m = ScalePolicy_FixedScale{} }
func (m *ScalePolicy_FixedScale) String() string { return proto.CompactTextString(m) }
func (*ScalePolicy_FixedScale) ProtoMessage()    {}
func (*ScalePolicy_FixedScale) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5c271823cf8486a, []int{6, 0}
}

func (m *ScalePolicy_FixedScale) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScalePolicy_FixedScale.Unmarshal(m, b)
}
func (m *ScalePolicy_FixedScale) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScalePolicy_FixedScale.Marshal(b, m, deterministic)
}
func (m *ScalePolicy_FixedScale) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScalePolicy_FixedScale.Merge(m, src)
}
func (m *ScalePolicy_FixedScale) XXX_Size() int {
	return xxx_messageInfo_ScalePolicy_FixedScale.Size(m)
}
func (m *ScalePolicy_FixedScale) XXX_DiscardUnknown() {
	xxx_messageInfo_ScalePolicy_FixedScale.DiscardUnknown(m)
}

var xxx_messageInfo_ScalePolicy_FixedScale proto.InternalMessageInfo

func (m *ScalePolicy_FixedScale) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type NodeGroupAllocationPolicy struct {
	// List of locations where resources for the node group will be allocated.
	Locations            []*NodeGroupLocation `protobuf:"bytes,1,rep,name=locations,proto3" json:"locations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *NodeGroupAllocationPolicy) Reset()         { *m = NodeGroupAllocationPolicy{} }
func (m *NodeGroupAllocationPolicy) String() string { return proto.CompactTextString(m) }
func (*NodeGroupAllocationPolicy) ProtoMessage()    {}
func (*NodeGroupAllocationPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5c271823cf8486a, []int{7}
}

func (m *NodeGroupAllocationPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeGroupAllocationPolicy.Unmarshal(m, b)
}
func (m *NodeGroupAllocationPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeGroupAllocationPolicy.Marshal(b, m, deterministic)
}
func (m *NodeGroupAllocationPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeGroupAllocationPolicy.Merge(m, src)
}
func (m *NodeGroupAllocationPolicy) XXX_Size() int {
	return xxx_messageInfo_NodeGroupAllocationPolicy.Size(m)
}
func (m *NodeGroupAllocationPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeGroupAllocationPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_NodeGroupAllocationPolicy proto.InternalMessageInfo

func (m *NodeGroupAllocationPolicy) GetLocations() []*NodeGroupLocation {
	if m != nil {
		return m.Locations
	}
	return nil
}

type NodeGroupLocation struct {
	// ID of the availability zone where the nodes may reside.
	ZoneId string `protobuf:"bytes,1,opt,name=zone_id,json=zoneId,proto3" json:"zone_id,omitempty"`
	// ID of the subnet. If a network chosen for the Kubernetes cluster has only one subnet in the specified zone, subnet ID may be omitted.
	SubnetId             string   `protobuf:"bytes,2,opt,name=subnet_id,json=subnetId,proto3" json:"subnet_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeGroupLocation) Reset()         { *m = NodeGroupLocation{} }
func (m *NodeGroupLocation) String() string { return proto.CompactTextString(m) }
func (*NodeGroupLocation) ProtoMessage()    {}
func (*NodeGroupLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5c271823cf8486a, []int{8}
}

func (m *NodeGroupLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeGroupLocation.Unmarshal(m, b)
}
func (m *NodeGroupLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeGroupLocation.Marshal(b, m, deterministic)
}
func (m *NodeGroupLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeGroupLocation.Merge(m, src)
}
func (m *NodeGroupLocation) XXX_Size() int {
	return xxx_messageInfo_NodeGroupLocation.Size(m)
}
func (m *NodeGroupLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeGroupLocation.DiscardUnknown(m)
}

var xxx_messageInfo_NodeGroupLocation proto.InternalMessageInfo

func (m *NodeGroupLocation) GetZoneId() string {
	if m != nil {
		return m.ZoneId
	}
	return ""
}

func (m *NodeGroupLocation) GetSubnetId() string {
	if m != nil {
		return m.SubnetId
	}
	return ""
}

type SchedulingPolicy struct {
	// True for preemptible compute instances. Default value is false. Preemptible compute instances are stopped at least once every 24 hours, and can be stopped at any time
	// if their resources are needed by Compute.
	// For more information, see [Preemptible Virtual Machines](/docs/compute/concepts/preemptible-vm).
	Preemptible          bool     `protobuf:"varint,1,opt,name=preemptible,proto3" json:"preemptible,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SchedulingPolicy) Reset()         { *m = SchedulingPolicy{} }
func (m *SchedulingPolicy) String() string { return proto.CompactTextString(m) }
func (*SchedulingPolicy) ProtoMessage()    {}
func (*SchedulingPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5c271823cf8486a, []int{9}
}

func (m *SchedulingPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SchedulingPolicy.Unmarshal(m, b)
}
func (m *SchedulingPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SchedulingPolicy.Marshal(b, m, deterministic)
}
func (m *SchedulingPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SchedulingPolicy.Merge(m, src)
}
func (m *SchedulingPolicy) XXX_Size() int {
	return xxx_messageInfo_SchedulingPolicy.Size(m)
}
func (m *SchedulingPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_SchedulingPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_SchedulingPolicy proto.InternalMessageInfo

func (m *SchedulingPolicy) GetPreemptible() bool {
	if m != nil {
		return m.Preemptible
	}
	return false
}

type NodeGroupMaintenancePolicy struct {
	AutoUpgrade          bool               `protobuf:"varint,1,opt,name=auto_upgrade,json=autoUpgrade,proto3" json:"auto_upgrade,omitempty"`
	AutoRepair           bool               `protobuf:"varint,2,opt,name=auto_repair,json=autoRepair,proto3" json:"auto_repair,omitempty"`
	MaintenanceWindow    *MaintenanceWindow `protobuf:"bytes,3,opt,name=maintenance_window,json=maintenanceWindow,proto3" json:"maintenance_window,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *NodeGroupMaintenancePolicy) Reset()         { *m = NodeGroupMaintenancePolicy{} }
func (m *NodeGroupMaintenancePolicy) String() string { return proto.CompactTextString(m) }
func (*NodeGroupMaintenancePolicy) ProtoMessage()    {}
func (*NodeGroupMaintenancePolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5c271823cf8486a, []int{10}
}

func (m *NodeGroupMaintenancePolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeGroupMaintenancePolicy.Unmarshal(m, b)
}
func (m *NodeGroupMaintenancePolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeGroupMaintenancePolicy.Marshal(b, m, deterministic)
}
func (m *NodeGroupMaintenancePolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeGroupMaintenancePolicy.Merge(m, src)
}
func (m *NodeGroupMaintenancePolicy) XXX_Size() int {
	return xxx_messageInfo_NodeGroupMaintenancePolicy.Size(m)
}
func (m *NodeGroupMaintenancePolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeGroupMaintenancePolicy.DiscardUnknown(m)
}

var xxx_messageInfo_NodeGroupMaintenancePolicy proto.InternalMessageInfo

func (m *NodeGroupMaintenancePolicy) GetAutoUpgrade() bool {
	if m != nil {
		return m.AutoUpgrade
	}
	return false
}

func (m *NodeGroupMaintenancePolicy) GetAutoRepair() bool {
	if m != nil {
		return m.AutoRepair
	}
	return false
}

func (m *NodeGroupMaintenancePolicy) GetMaintenanceWindow() *MaintenanceWindow {
	if m != nil {
		return m.MaintenanceWindow
	}
	return nil
}

func init() {
	proto.RegisterEnum("yandex.cloud.k8s.v1.IpVersion", IpVersion_name, IpVersion_value)
	proto.RegisterEnum("yandex.cloud.k8s.v1.NodeGroup_Status", NodeGroup_Status_name, NodeGroup_Status_value)
	proto.RegisterType((*NodeGroup)(nil), "yandex.cloud.k8s.v1.NodeGroup")
	proto.RegisterMapType((map[string]string)(nil), "yandex.cloud.k8s.v1.NodeGroup.LabelsEntry")
	proto.RegisterType((*NodeTemplate)(nil), "yandex.cloud.k8s.v1.NodeTemplate")
	proto.RegisterMapType((map[string]string)(nil), "yandex.cloud.k8s.v1.NodeTemplate.MetadataEntry")
	proto.RegisterType((*NodeAddressSpec)(nil), "yandex.cloud.k8s.v1.NodeAddressSpec")
	proto.RegisterType((*OneToOneNatSpec)(nil), "yandex.cloud.k8s.v1.OneToOneNatSpec")
	proto.RegisterType((*ResourcesSpec)(nil), "yandex.cloud.k8s.v1.ResourcesSpec")
	proto.RegisterType((*DiskSpec)(nil), "yandex.cloud.k8s.v1.DiskSpec")
	proto.RegisterType((*ScalePolicy)(nil), "yandex.cloud.k8s.v1.ScalePolicy")
	proto.RegisterType((*ScalePolicy_FixedScale)(nil), "yandex.cloud.k8s.v1.ScalePolicy.FixedScale")
	proto.RegisterType((*NodeGroupAllocationPolicy)(nil), "yandex.cloud.k8s.v1.NodeGroupAllocationPolicy")
	proto.RegisterType((*NodeGroupLocation)(nil), "yandex.cloud.k8s.v1.NodeGroupLocation")
	proto.RegisterType((*SchedulingPolicy)(nil), "yandex.cloud.k8s.v1.SchedulingPolicy")
	proto.RegisterType((*NodeGroupMaintenancePolicy)(nil), "yandex.cloud.k8s.v1.NodeGroupMaintenancePolicy")
}

func init() {
	proto.RegisterFile("yandex/cloud/k8s/v1/node_group.proto", fileDescriptor_d5c271823cf8486a)
}

var fileDescriptor_d5c271823cf8486a = []byte{
	// 1406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0x4f, 0x6f, 0xdb, 0xc6,
	0x12, 0x0f, 0x6d, 0x59, 0x96, 0x46, 0xb2, 0x2d, 0x6f, 0x1e, 0xf2, 0xf4, 0xf4, 0x9e, 0x5f, 0x1c,
	0x21, 0x29, 0x0c, 0x87, 0xa2, 0x48, 0x9a, 0x91, 0xe5, 0x3a, 0x2e, 0xe0, 0x3f, 0x72, 0x4a, 0xc0,
	0x91, 0x05, 0x4a, 0x76, 0x81, 0x06, 0x2d, 0xb1, 0x26, 0xd7, 0x0a, 0x61, 0x89, 0x4b, 0x90, 0x94,
	0x12, 0x1b, 0xb9, 0xb4, 0xc7, 0x7e, 0x88, 0x1e, 0xfa, 0x11, 0x72, 0xe9, 0xad, 0x68, 0x7b, 0xb1,
	0x3f, 0x4a, 0xcf, 0x3d, 0xfa, 0x54, 0xec, 0x92, 0x92, 0x65, 0x55, 0x76, 0xd0, 0x93, 0x76, 0x7e,
	0xf3, 0x9b, 0xd1, 0xfe, 0x76, 0x67, 0x76, 0x08, 0x4f, 0xcf, 0xb1, 0x6b, 0x93, 0xf7, 0x65, 0xab,
	0x43, 0x7b, 0x76, 0xf9, 0xac, 0x1a, 0x94, 0xfb, 0x4a, 0xd9, 0xa5, 0x36, 0x31, 0xdb, 0x3e, 0xed,
	0x79, 0x92, 0xe7, 0xd3, 0x90, 0xa2, 0x87, 0x11, 0x4b, 0xe2, 0x2c, 0xe9, 0xac, 0x1a, 0x48, 0x7d,
	0xa5, 0xf0, 0xb8, 0x4d, 0x69, 0xbb, 0x43, 0xca, 0x9c, 0x72, 0xd2, 0x3b, 0x2d, 0x87, 0x4e, 0x97,
	0x04, 0x21, 0xee, 0xc6, 0x51, 0x85, 0x67, 0x93, 0x72, 0x77, 0xb1, 0xe3, 0x86, 0xc4, 0xc5, 0xae,
	0x45, 0x62, 0xda, 0x93, 0x49, 0xb4, 0x3e, 0xf1, 0x03, 0x87, 0xba, 0x31, 0x65, 0xe9, 0x16, 0xa5,
	0x8f, 0x3b, 0x8e, 0x8d, 0xc3, 0xa1, 0xbb, 0xf8, 0xfb, 0x2c, 0xa4, 0xeb, 0xd4, 0x26, 0xaf, 0xd8,
	0x96, 0xd1, 0x3c, 0x4c, 0x39, 0x76, 0x5e, 0x58, 0x16, 0x56, 0xd2, 0xc6, 0x94, 0x63, 0xa3, 0x25,
	0x00, 0xab, 0xd3, 0x0b, 0x42, 0xe2, 0x9b, 0x8e, 0x9d, 0x9f, 0xe2, 0x78, 0x3a, 0x46, 0x74, 0x1b,
	0x6d, 0x00, 0x58, 0x3e, 0xc1, 0x21, 0xb1, 0x4d, 0x1c, 0xe6, 0xa7, 0x97, 0x85, 0x95, 0x8c, 0x5a,
	0x90, 0x22, 0x6d, 0xd2, 0x40, 0x9b, 0xd4, 0x1a, 0x68, 0x33, 0xd2, 0x31, 0x7b, 0x3b, 0x44, 0x08,
	0x12, 0x2e, 0xee, 0x92, 0x7c, 0x82, 0xe7, 0xe4, 0x6b, 0xb4, 0x0c, 0x19, 0x9b, 0x04, 0x96, 0xef,
	0x78, 0x6c, 0x83, 0xf9, 0x19, 0xee, 0x1a, 0x85, 0xd0, 0x0e, 0x24, 0x3b, 0xf8, 0x84, 0x74, 0x82,
	0x7c, 0x72, 0x79, 0x7a, 0x25, 0xa3, 0xae, 0x4a, 0x13, 0x4e, 0x57, 0x1a, 0xea, 0x91, 0x0e, 0x38,
	0xb9, 0xe6, 0x86, 0xfe, 0xb9, 0x11, 0x47, 0xa2, 0x2d, 0x48, 0x06, 0x21, 0x0e, 0x7b, 0x41, 0x7e,
	0x76, 0x59, 0x58, 0x99, 0x57, 0x9f, 0x7d, 0x22, 0x47, 0x93, 0x93, 0x8d, 0x38, 0x08, 0xed, 0xc3,
	0x1c, 0xbf, 0xe3, 0x90, 0x74, 0xbd, 0x0e, 0x0e, 0x49, 0x3e, 0xc5, 0x65, 0x3f, 0xb9, 0x33, 0x4b,
	0x2b, 0x26, 0x1a, 0x59, 0x77, 0xc4, 0x42, 0xbb, 0x90, 0x0d, 0x2c, 0xdc, 0x21, 0xa6, 0x47, 0x3b,
	0x8e, 0x75, 0x9e, 0x4f, 0xf3, 0x34, 0xcb, 0x13, 0xd3, 0x34, 0x19, 0xb1, 0xc1, 0x79, 0x46, 0x26,
	0xb8, 0x31, 0xd0, 0x1b, 0x58, 0xc4, 0x9d, 0x0e, 0xb5, 0xf8, 0x8d, 0x0e, 0x32, 0x01, 0xcf, 0x24,
	0xdd, 0x2f, 0x6b, 0x7b, 0x18, 0x16, 0xe7, 0xcd, 0xe1, 0x31, 0x04, 0xad, 0xc2, 0xa2, 0xe3, 0x06,
	0x21, 0x2b, 0xb7, 0xa8, 0xa2, 0x59, 0x0d, 0x64, 0xf8, 0xa5, 0x2c, 0x0c, 0x1c, 0x3c, 0x97, 0x6e,
	0xa3, 0x27, 0xc0, 0xd5, 0x99, 0x71, 0xed, 0xe5, 0xb3, 0xd1, 0xdd, 0x31, 0xec, 0x38, 0x82, 0x98,
	0xe0, 0xd8, 0x6b, 0x3a, 0xee, 0x29, 0xcd, 0xcf, 0xdd, 0x23, 0x38, 0x8e, 0xd1, 0xdd, 0x53, 0x6a,
	0x64, 0xfa, 0x37, 0x06, 0xfa, 0x16, 0xd0, 0x48, 0x17, 0x0c, 0x14, 0xcf, 0xf3, 0x54, 0xe5, 0xfb,
	0x15, 0xbf, 0xbe, 0x89, 0x8b, 0x25, 0x2f, 0x76, 0xc7, 0xa1, 0xc2, 0x06, 0x64, 0x46, 0x6a, 0x06,
	0xe5, 0x60, 0xfa, 0x8c, 0x9c, 0xc7, 0x0d, 0xc1, 0x96, 0xe8, 0x5f, 0x30, 0xd3, 0xc7, 0x9d, 0x1e,
	0x89, 0x9b, 0x21, 0x32, 0x3e, 0x9f, 0xaa, 0x0a, 0xc5, 0x73, 0x48, 0x46, 0xa5, 0x82, 0x1e, 0x01,
	0x6a, 0xb6, 0xb6, 0x5b, 0x47, 0x4d, 0xf3, 0xa8, 0xde, 0x6c, 0xd4, 0x76, 0xf5, 0x7d, 0xbd, 0xb6,
	0x97, 0x7b, 0x80, 0x72, 0x90, 0x6d, 0x18, 0x87, 0xc7, 0x7a, 0x53, 0x3f, 0xac, 0xeb, 0xf5, 0x57,
	0x39, 0x01, 0x65, 0x60, 0xd6, 0x38, 0xaa, 0x73, 0x63, 0x0a, 0x2d, 0x40, 0xc6, 0xa8, 0xed, 0x1e,
	0xd6, 0x77, 0xf5, 0x03, 0x06, 0x4c, 0xa3, 0x2c, 0xa4, 0x9a, 0xad, 0xc3, 0x46, 0x83, 0x59, 0x09,
	0xc6, 0xe5, 0x56, 0x6d, 0x2f, 0x37, 0xc3, 0x5c, 0x7b, 0xb5, 0x83, 0x5a, 0x8b, 0xb9, 0x92, 0xc5,
	0x9f, 0x12, 0x90, 0x1d, 0x2d, 0x35, 0xf4, 0x18, 0x32, 0x6c, 0x71, 0x4a, 0xfd, 0xae, 0x39, 0x6c,
	0x68, 0x18, 0x40, 0xba, 0x8d, 0x74, 0x98, 0xf7, 0x49, 0x40, 0x7b, 0xbe, 0x45, 0x02, 0x33, 0xf0,
	0x88, 0xc5, 0xf5, 0x64, 0xd4, 0xe2, 0xc4, 0x33, 0x34, 0x06, 0xd4, 0xa6, 0x47, 0x2c, 0x63, 0xce,
	0x1f, 0x35, 0xd1, 0x2e, 0xcc, 0x9f, 0x50, 0x1a, 0x9a, 0xb6, 0x13, 0x9c, 0x45, 0xa9, 0xa2, 0x87,
	0x60, 0x69, 0x62, 0xaa, 0x3d, 0x27, 0x38, 0xe3, 0x59, 0xb2, 0x2c, 0x68, 0x60, 0xa1, 0xef, 0x04,
	0x48, 0x75, 0x49, 0x88, 0x6d, 0x1c, 0xe2, 0x7c, 0x82, 0xf7, 0x76, 0xf9, 0x93, 0x1d, 0x25, 0xbd,
	0x8e, 0x23, 0xf8, 0x65, 0xed, 0x68, 0xdf, 0x5f, 0x29, 0x89, 0x97, 0x5b, 0x15, 0xed, 0x87, 0x2b,
	0x25, 0xf5, 0x72, 0x4b, 0x59, 0x53, 0xe4, 0x75, 0xf5, 0xe3, 0x95, 0x92, 0x2c, 0x24, 0x94, 0x52,
	0x65, 0xed, 0xe3, 0x95, 0x82, 0x50, 0xee, 0x0d, 0x2e, 0x5d, 0x7c, 0xf3, 0xa6, 0x64, 0xca, 0xa5,
	0x0d, 0xb6, 0x5a, 0x35, 0x86, 0x7f, 0x8b, 0x0e, 0x60, 0xa1, 0xaf, 0x99, 0xd8, 0xb6, 0x7d, 0x12,
	0xc4, 0x87, 0x32, 0xc3, 0x95, 0x3c, 0xbd, 0x73, 0x27, 0xdb, 0x11, 0x39, 0x3a, 0x96, 0xbe, 0x36,
	0x62, 0x22, 0x03, 0x16, 0x03, 0xeb, 0x2d, 0xb1, 0x7b, 0x1d, 0xc7, 0x6d, 0x0f, 0x0a, 0x35, 0xc9,
	0xf3, 0x3d, 0xbb, 0xa3, 0xc9, 0x07, 0xec, 0x41, 0x47, 0x06, 0x63, 0x48, 0x61, 0x13, 0xe6, 0x6e,
	0x49, 0xfe, 0x47, 0xf5, 0x49, 0x60, 0x61, 0x6c, 0xcb, 0xc8, 0x80, 0x87, 0xd4, 0x25, 0x66, 0x48,
	0x4d, 0xf6, 0xe3, 0xe2, 0x30, 0x52, 0x2d, 0xdc, 0xa3, 0xfa, 0xd0, 0x25, 0x2d, 0x7a, 0xe8, 0x92,
	0x3a, 0x0e, 0xb9, 0xea, 0x05, 0x7a, 0x1b, 0x28, 0x36, 0x60, 0x61, 0x8c, 0x83, 0xb6, 0x00, 0x1c,
	0x6f, 0xf8, 0x34, 0x08, 0xfc, 0xd5, 0xfd, 0xff, 0xc4, 0xec, 0xba, 0x17, 0x77, 0xbe, 0x91, 0x76,
	0x06, 0xcb, 0xe2, 0x2f, 0x02, 0xcc, 0xdd, 0xaa, 0x40, 0xb4, 0x0a, 0xc9, 0x2e, 0xe9, 0x52, 0x3f,
	0x52, 0x3e, 0xbd, 0x83, 0xae, 0x2f, 0x95, 0x79, 0xb9, 0xa4, 0xae, 0x6b, 0xd5, 0xf5, 0xf5, 0x0d,
	0xb9, 0xb2, 0xa1, 0x69, 0x46, 0xcc, 0x40, 0x3a, 0xcc, 0x58, 0xd4, 0x27, 0x01, 0x3f, 0x90, 0xe9,
	0x9d, 0xb5, 0xeb, 0x4b, 0xa5, 0x2c, 0x8b, 0x8a, 0xa8, 0x8a, 0x9a, 0x58, 0x11, 0xab, 0xa2, 0x22,
	0x8b, 0x8a, 0x2a, 0x2a, 0x9a, 0xa8, 0x54, 0x44, 0xa5, 0x2a, 0xaa, 0xb2, 0xa8, 0xaa, 0xa2, 0xaa,
	0x89, 0x6a, 0x45, 0x54, 0xab, 0xe2, 0x9a, 0x2c, 0xae, 0xa9, 0x46, 0x94, 0x01, 0x55, 0x60, 0x8e,
	0x2d, 0xcc, 0x53, 0x1f, 0x5b, 0x7c, 0x42, 0x4d, 0xf3, 0x94, 0x8b, 0xd7, 0x97, 0xca, 0x9c, 0x2c,
	0xbe, 0x60, 0xb1, 0x2f, 0x64, 0x51, 0x91, 0x65, 0x23, 0xcb, 0x78, 0xfb, 0x31, 0xad, 0xf8, 0x01,
	0x52, 0xc3, 0x42, 0xff, 0x02, 0xb2, 0xbc, 0x51, 0xc2, 0x73, 0x8f, 0x0c, 0x5b, 0x73, 0xe7, 0x7f,
	0x7f, 0x5e, 0x2a, 0xf9, 0x0f, 0x2e, 0x09, 0xdf, 0x51, 0xff, 0xac, 0x14, 0x04, 0xf6, 0x70, 0xfd,
	0xd6, 0xb6, 0x0d, 0x60, 0x11, 0xad, 0x73, 0x8f, 0xe8, 0x36, 0x92, 0x21, 0x1d, 0x35, 0x9a, 0x73,
	0x41, 0x62, 0x49, 0x0f, 0xaf, 0x2f, 0x95, 0x05, 0xb9, 0xa4, 0xad, 0x6d, 0x54, 0x65, 0xad, 0xf2,
	0x42, 0x51, 0x14, 0x59, 0x33, 0x52, 0x8c, 0xd5, 0x74, 0x2e, 0x48, 0xf1, 0x47, 0x01, 0x32, 0x23,
	0x03, 0x04, 0xd5, 0x21, 0x73, 0xea, 0xbc, 0x27, 0xb6, 0xc9, 0x07, 0x49, 0x7c, 0xd9, 0xcf, 0x3f,
	0x35, 0x77, 0xa4, 0x7d, 0x16, 0xc3, 0x81, 0x2f, 0x1f, 0x18, 0x70, 0x3a, 0xb4, 0x0a, 0xcf, 0x01,
	0x6e, 0x7c, 0x68, 0x09, 0x12, 0x7c, 0x6b, 0xd1, 0xc5, 0xa4, 0xaf, 0x2f, 0x95, 0x19, 0xb9, 0xc4,
	0x8e, 0x84, 0xc3, 0x3b, 0x08, 0x20, 0x9a, 0x7a, 0x4c, 0x3f, 0x4a, 0xfc, 0xfa, 0x9b, 0x22, 0x14,
	0x31, 0xfc, 0xe7, 0xce, 0xb1, 0x84, 0xf6, 0x20, 0x3d, 0x40, 0x82, 0xbc, 0xc0, 0x1f, 0x86, 0xcf,
	0xee, 0x7f, 0xe7, 0x0f, 0x62, 0xba, 0x71, 0x13, 0x58, 0x3c, 0x84, 0xc5, 0xbf, 0xf9, 0xd1, 0x12,
	0xcc, 0x5e, 0xb0, 0xba, 0x1f, 0xde, 0x42, 0xe2, 0x8f, 0x4b, 0x45, 0x30, 0x92, 0x0c, 0xd4, 0x6d,
	0xf4, 0x5f, 0x48, 0x07, 0xbd, 0x13, 0x97, 0x84, 0x37, 0x9f, 0x3e, 0xa9, 0x08, 0xd0, 0xed, 0xa2,
	0x06, 0xb9, 0xf1, 0x7e, 0x65, 0x9f, 0x2f, 0x9e, 0x4f, 0x48, 0xd7, 0x0b, 0x9d, 0x93, 0xf8, 0x60,
	0x53, 0xc6, 0x28, 0x54, 0xfc, 0x59, 0x80, 0xc2, 0xdd, 0xf3, 0x88, 0x0d, 0x51, 0xdc, 0x0b, 0xa9,
	0xd9, 0xf3, 0xda, 0x3e, 0xb6, 0x87, 0x19, 0x18, 0x76, 0x14, 0x41, 0xec, 0x61, 0xe7, 0x14, 0x9f,
	0x78, 0xd8, 0xf1, 0xf9, 0xb6, 0x52, 0x06, 0x30, 0xc8, 0xe0, 0x08, 0x3a, 0xba, 0x3d, 0x20, 0xdf,
	0x39, 0xae, 0x4d, 0xdf, 0xc5, 0x2f, 0xf2, 0xe4, 0x83, 0x1b, 0xd9, 0xc7, 0x57, 0x9c, 0x7d, 0x6b,
	0x2e, 0x46, 0xd0, 0xea, 0x26, 0xa4, 0x87, 0xbd, 0x89, 0x0a, 0xf0, 0x48, 0x6f, 0x98, 0xc7, 0x35,
	0x83, 0x0d, 0xb2, 0xb1, 0x19, 0x97, 0x82, 0x84, 0xde, 0x38, 0xd6, 0x72, 0x42, 0xbc, 0xaa, 0xe4,
	0xa6, 0x76, 0x8e, 0xe1, 0xdf, 0xb7, 0xfe, 0x18, 0x7b, 0x4e, 0xfc, 0xe7, 0x5f, 0x6f, 0xb6, 0x9d,
	0xf0, 0x6d, 0xef, 0x44, 0xb2, 0x68, 0xb7, 0x1c, 0x71, 0x4a, 0xd1, 0x87, 0x6a, 0x9b, 0x96, 0xda,
	0xc4, 0xe5, 0xdf, 0x90, 0xe5, 0x09, 0x1f, 0xb9, 0x9b, 0x67, 0xd5, 0xe0, 0x24, 0xc9, 0xdd, 0x6b,
	0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0xf3, 0xe0, 0xdf, 0xb9, 0x89, 0x0b, 0x00, 0x00,
}