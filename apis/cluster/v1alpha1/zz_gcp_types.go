/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type GCPBackupPolicyObservation struct {
}

type GCPBackupPolicyParameters struct {

	// +kubebuilder:validation:Required
	BackupLocationID *string `json:"backupLocationId" tf:"backup_location_id,omitempty"`

	// +kubebuilder:validation:Required
	ExpiryInHour *float64 `json:"expiryInHour" tf:"expiry_in_hour,omitempty"`

	// +kubebuilder:validation:Optional
	IncludeClusterResources *bool `json:"includeClusterResources,omitempty" tf:"include_cluster_resources,omitempty"`

	// +kubebuilder:validation:Optional
	IncludeDisks *bool `json:"includeDisks,omitempty" tf:"include_disks,omitempty"`

	// +kubebuilder:validation:Optional
	Namespaces []*string `json:"namespaces,omitempty" tf:"namespaces,omitempty"`

	// +kubebuilder:validation:Required
	Prefix *string `json:"prefix" tf:"prefix,omitempty"`

	// +kubebuilder:validation:Required
	Schedule *string `json:"schedule" tf:"schedule,omitempty"`
}

type GCPCloudConfigObservation struct {
}

type GCPCloudConfigParameters struct {

	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// +kubebuilder:validation:Required
	Project *string `json:"project" tf:"project,omitempty"`

	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`
}

type GCPClusterProfileObservation struct {
}

type GCPClusterProfilePackManifestObservation struct {
}

type GCPClusterProfilePackManifestParameters struct {

	// +kubebuilder:validation:Required
	Content *string `json:"content" tf:"content,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type GCPClusterProfilePackObservation struct {
}

type GCPClusterProfilePackParameters struct {

	// +kubebuilder:validation:Optional
	Manifest []GCPClusterProfilePackManifestParameters `json:"manifest,omitempty" tf:"manifest,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	RegistryUID *string `json:"registryUid,omitempty" tf:"registry_uid,omitempty"`

	// +kubebuilder:validation:Optional
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// +kubebuilder:validation:Required
	Values *string `json:"values" tf:"values,omitempty"`
}

type GCPClusterProfileParameters struct {

	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// +kubebuilder:validation:Optional
	Pack []GCPClusterProfilePackParameters `json:"pack,omitempty" tf:"pack,omitempty"`
}

type GCPClusterRbacBindingObservation struct {
}

type GCPClusterRbacBindingParameters struct {

	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// +kubebuilder:validation:Optional
	Role map[string]*string `json:"role,omitempty" tf:"role,omitempty"`

	// +kubebuilder:validation:Optional
	Subjects []GCPClusterRbacBindingSubjectsParameters `json:"subjects,omitempty" tf:"subjects,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type GCPClusterRbacBindingSubjectsObservation struct {
}

type GCPClusterRbacBindingSubjectsParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type GCPHostConfigObservation struct {
}

type GCPHostConfigParameters struct {

	// The external traffic policy for the cluster.
	// +kubebuilder:validation:Optional
	ExternalTrafficPolicy *string `json:"externalTrafficPolicy,omitempty" tf:"external_traffic_policy,omitempty"`

	// The type of endpoint for the cluster. Can be either 'Ingress' or 'LoadBalancer'. The default is 'Ingress'.
	// +kubebuilder:validation:Optional
	HostEndpointType *string `json:"hostEndpointType,omitempty" tf:"host_endpoint_type,omitempty"`

	// The host for the Ingress endpoint. Required if 'host_endpoint_type' is set to 'Ingress'.
	// +kubebuilder:validation:Optional
	IngressHost *string `json:"ingressHost,omitempty" tf:"ingress_host,omitempty"`

	// The source ranges for the load balancer. Required if 'host_endpoint_type' is set to 'LoadBalancer'.
	// +kubebuilder:validation:Optional
	LoadBalancerSourceRanges *string `json:"loadBalancerSourceRanges,omitempty" tf:"load_balancer_source_ranges,omitempty"`
}

type GCPLocationConfigObservation struct {
	CountryCode *string `json:"countryCode,omitempty" tf:"country_code,omitempty"`

	CountryName *string `json:"countryName,omitempty" tf:"country_name,omitempty"`

	Latitude *float64 `json:"latitude,omitempty" tf:"latitude,omitempty"`

	Longitude *float64 `json:"longitude,omitempty" tf:"longitude,omitempty"`

	RegionCode *string `json:"regionCode,omitempty" tf:"region_code,omitempty"`

	RegionName *string `json:"regionName,omitempty" tf:"region_name,omitempty"`
}

type GCPLocationConfigParameters struct {
}

type GCPMachinePoolObservation struct {
}

type GCPMachinePoolParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalLabels map[string]*string `json:"additionalLabels,omitempty" tf:"additional_labels,omitempty"`

	// +kubebuilder:validation:Required
	Azs []*string `json:"azs" tf:"azs,omitempty"`

	// +kubebuilder:validation:Optional
	ControlPlane *bool `json:"controlPlane,omitempty" tf:"control_plane,omitempty"`

	// +kubebuilder:validation:Optional
	ControlPlaneAsWorker *bool `json:"controlPlaneAsWorker,omitempty" tf:"control_plane_as_worker,omitempty"`

	// +kubebuilder:validation:Required
	Count *float64 `json:"count" tf:"count,omitempty"`

	// +kubebuilder:validation:Optional
	DiskSizeGb *float64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`

	// +kubebuilder:validation:Required
	InstanceType *string `json:"instanceType" tf:"instance_type,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Taints []GCPMachinePoolTaintsParameters `json:"taints,omitempty" tf:"taints,omitempty"`

	// +kubebuilder:validation:Optional
	UpdateStrategy *string `json:"updateStrategy,omitempty" tf:"update_strategy,omitempty"`
}

type GCPMachinePoolTaintsObservation struct {
}

type GCPMachinePoolTaintsParameters struct {

	// +kubebuilder:validation:Required
	Effect *string `json:"effect" tf:"effect,omitempty"`

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type GCPNamespacesObservation struct {
}

type GCPNamespacesParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceAllocation map[string]*string `json:"resourceAllocation" tf:"resource_allocation,omitempty"`
}

type GCPObservation struct {
	CloudConfigID *string `json:"cloudConfigId,omitempty" tf:"cloud_config_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Kubeconfig *string `json:"kubeconfig,omitempty" tf:"kubeconfig,omitempty"`

	LocationConfig []GCPLocationConfigObservation `json:"locationConfig,omitempty" tf:"location_config,omitempty"`
}

type GCPPackObservation struct {
}

type GCPPackParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	RegistryUID *string `json:"registryUid,omitempty" tf:"registry_uid,omitempty"`

	// +kubebuilder:validation:Required
	Tag *string `json:"tag" tf:"tag,omitempty"`

	// +kubebuilder:validation:Required
	Values *string `json:"values" tf:"values,omitempty"`
}

type GCPParameters struct {

	// +kubebuilder:validation:Optional
	ApplySetting *string `json:"applySetting,omitempty" tf:"apply_setting,omitempty"`

	// +kubebuilder:validation:Optional
	BackupPolicy []GCPBackupPolicyParameters `json:"backupPolicy,omitempty" tf:"backup_policy,omitempty"`

	// +kubebuilder:validation:Required
	CloudAccountID *string `json:"cloudAccountId" tf:"cloud_account_id,omitempty"`

	// +kubebuilder:validation:Required
	CloudConfig []GCPCloudConfigParameters `json:"cloudConfig" tf:"cloud_config,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterProfile []GCPClusterProfileParameters `json:"clusterProfile,omitempty" tf:"cluster_profile,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterRbacBinding []GCPClusterRbacBindingParameters `json:"clusterRbacBinding,omitempty" tf:"cluster_rbac_binding,omitempty"`

	// +kubebuilder:validation:Optional
	HostConfig []GCPHostConfigParameters `json:"hostConfig,omitempty" tf:"host_config,omitempty"`

	// +kubebuilder:validation:Required
	MachinePool []GCPMachinePoolParameters `json:"machinePool" tf:"machine_pool,omitempty"`

	// +kubebuilder:validation:Optional
	Namespaces []GCPNamespacesParameters `json:"namespaces,omitempty" tf:"namespaces,omitempty"`

	// +kubebuilder:validation:Optional
	OsPatchAfter *string `json:"osPatchAfter,omitempty" tf:"os_patch_after,omitempty"`

	// +kubebuilder:validation:Optional
	OsPatchOnBoot *bool `json:"osPatchOnBoot,omitempty" tf:"os_patch_on_boot,omitempty"`

	// +kubebuilder:validation:Optional
	OsPatchSchedule *string `json:"osPatchSchedule,omitempty" tf:"os_patch_schedule,omitempty"`

	// +kubebuilder:validation:Optional
	Pack []GCPPackParameters `json:"pack,omitempty" tf:"pack,omitempty"`

	// +kubebuilder:validation:Optional
	ScanPolicy []GCPScanPolicyParameters `json:"scanPolicy,omitempty" tf:"scan_policy,omitempty"`

	// +kubebuilder:validation:Optional
	SkipCompletion *bool `json:"skipCompletion,omitempty" tf:"skip_completion,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type GCPScanPolicyObservation struct {
}

type GCPScanPolicyParameters struct {

	// +kubebuilder:validation:Required
	ConfigurationScanSchedule *string `json:"configurationScanSchedule" tf:"configuration_scan_schedule,omitempty"`

	// +kubebuilder:validation:Required
	ConformanceScanSchedule *string `json:"conformanceScanSchedule" tf:"conformance_scan_schedule,omitempty"`

	// +kubebuilder:validation:Required
	PenetrationScanSchedule *string `json:"penetrationScanSchedule" tf:"penetration_scan_schedule,omitempty"`
}

// GCPSpec defines the desired state of GCP
type GCPSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GCPParameters `json:"forProvider"`
}

// GCPStatus defines the observed state of GCP.
type GCPStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GCPObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GCP is the Schema for the GCPs API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,palette}
type GCP struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GCPSpec   `json:"spec"`
	Status            GCPStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GCPList contains a list of GCPs
type GCPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GCP `json:"items"`
}

// Repository type metadata.
var (
	GCP_Kind             = "GCP"
	GCP_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GCP_Kind}.String()
	GCP_KindAPIVersion   = GCP_Kind + "." + CRDGroupVersion.String()
	GCP_GroupVersionKind = CRDGroupVersion.WithKind(GCP_Kind)
)

func init() {
	SchemeBuilder.Register(&GCP{}, &GCPList{})
}
