// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package talos

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Talos cluster config resource
type ClusterConfig struct {
	pulumi.CustomResourceState

	// additional Subject-Alt-Names for the APIServer certificate
	AdditionalSans pulumi.StringArrayOutput `pulumi:"additionalSans"`
	// cluster discovery feature
	ClusterDiscovery pulumi.BoolOutput `pulumi:"clusterDiscovery"`
	// The cluster endpoint
	ClusterEndpoint pulumi.StringOutput `pulumi:"clusterEndpoint"`
	// cluster name
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`
	// generated machineconfigs (applied to all node types)
	ConfigPatches ConfigPatchesOutput `pulumi:"configPatches"`
	// generated machineconfigs (applied to 'controlplane' types)
	ConfigPatchesControlPlane ConfigPatchesOutput `pulumi:"configPatchesControlPlane"`
	// generated machineconfigs (applied to 'worker' type)
	ConfigPatchesWorker ConfigPatchesOutput `pulumi:"configPatchesWorker"`
	// the desired machine config version to refer to
	ConfigVersion pulumi.StringOutput `pulumi:"configVersion"`
	// Talos Controlplane Config
	ControlplaneConfig pulumi.StringOutput `pulumi:"controlplaneConfig"`
	// the dns domain to use for cluster
	DnsDomain pulumi.StringOutput `pulumi:"dnsDomain"`
	// machine config documentation enabled
	Docs pulumi.BoolOutput `pulumi:"docs"`
	// machine config examples enabled
	Examples pulumi.BoolOutput `pulumi:"examples"`
	// the disk to install to
	InstallDisk pulumi.StringOutput `pulumi:"installDisk"`
	// the image used to perform an installation
	InstallImage pulumi.StringOutput `pulumi:"installImage"`
	// desired kubernetes version to run
	KubernetesVersion pulumi.StringOutput `pulumi:"kubernetesVersion"`
	// kubespan enabled
	Kubespan pulumi.BoolOutput `pulumi:"kubespan"`
	// persist value for configs
	Persist pulumi.BoolOutput `pulumi:"persist"`
	// list of registry mirrors
	RegistryMirrors pulumi.StringArrayOutput `pulumi:"registryMirrors"`
	// Talos Secrets Bundle
	Secrets SecretsBundleOutput `pulumi:"secrets"`
	// Talos Config
	TalosConfig pulumi.StringOutput `pulumi:"talosConfig"`
	// the desired Talos version
	TalosVersion pulumi.StringOutput `pulumi:"talosVersion"`
	// Talos Worker Config
	WorkerConfig pulumi.StringOutput `pulumi:"workerConfig"`
}

// NewClusterConfig registers a new resource with the given unique name, arguments, and options.
func NewClusterConfig(ctx *pulumi.Context,
	name string, args *ClusterConfigArgs, opts ...pulumi.ResourceOption) (*ClusterConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterEndpoint == nil {
		return nil, errors.New("invalid value for required argument 'ClusterEndpoint'")
	}
	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.Secrets == nil {
		return nil, errors.New("invalid value for required argument 'Secrets'")
	}
	if isZero(args.ClusterDiscovery) {
		args.ClusterDiscovery = pulumi.BoolPtr(true)
	}
	if isZero(args.DnsDomain) {
		args.DnsDomain = pulumi.StringPtr("cluster.local")
	}
	if isZero(args.Docs) {
		args.Docs = pulumi.BoolPtr(true)
	}
	if isZero(args.Examples) {
		args.Examples = pulumi.BoolPtr(true)
	}
	if isZero(args.InstallDisk) {
		args.InstallDisk = pulumi.StringPtr("/dev/sda")
	}
	if isZero(args.InstallImage) {
		args.InstallImage = pulumi.StringPtr("ghcr.io/talos-systems/installer:v0.14.2")
	}
	if isZero(args.KubernetesVersion) {
		args.KubernetesVersion = pulumi.StringPtr("1.23.6")
	}
	if isZero(args.Persist) {
		args.Persist = pulumi.BoolPtr(true)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"controlplaneConfig",
		"secrets",
		"talosConfig",
		"workerConfig",
	})
	opts = append(opts, secrets)
	var resource ClusterConfig
	err := ctx.RegisterResource("talos:index:clusterConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterConfig gets an existing ClusterConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterConfigState, opts ...pulumi.ResourceOption) (*ClusterConfig, error) {
	var resource ClusterConfig
	err := ctx.ReadResource("talos:index:clusterConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterConfig resources.
type clusterConfigState struct {
}

type ClusterConfigState struct {
}

func (ClusterConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterConfigState)(nil)).Elem()
}

type clusterConfigArgs struct {
	// additional Subject-Alt-Names for the APIServer certificate
	AdditionalSans []string `pulumi:"additionalSans"`
	// enable cluster discovery feature (default true)
	ClusterDiscovery *bool `pulumi:"clusterDiscovery"`
	// The cluster endpoint is the URL for the Kubernetes API. If you decide to use
	// a control plane node, common in a single node control plane setup, use port 6443 as
	// this is the port that the API server binds to on every control plane node. For an HA
	// setup, usually involving a load balancer, use the IP and port of the load balancer.
	ClusterEndpoint string `pulumi:"clusterEndpoint"`
	// cluster name
	ClusterName string `pulumi:"clusterName"`
	// patch generated machineconfigs (applied to all node types)
	ConfigPatches *ConfigPatches `pulumi:"configPatches"`
	// patch generated machineconfigs (applied to 'controlplane' types)
	ConfigPatchesControlPlane *ConfigPatches `pulumi:"configPatchesControlPlane"`
	// patch generated machineconfigs (applied to 'worker' type)
	ConfigPatchesWorker *ConfigPatches `pulumi:"configPatchesWorker"`
	// the desired machine config version to refer to
	ConfigVersion *string `pulumi:"configVersion"`
	// the dns domain to use for cluster (default "cluster.local")
	DnsDomain *string `pulumi:"dnsDomain"`
	// renders all machine configs adding the documentation for each field (default true)
	Docs *bool `pulumi:"docs"`
	// renders all machine configs with the commented examples (default true)
	Examples *bool `pulumi:"examples"`
	// the disk to install to (default "/dev/sda")
	InstallDisk *string `pulumi:"installDisk"`
	// the image used to perform an installation (default "ghcr.io/talos-systems/installer:v0.14.2")
	InstallImage *string `pulumi:"installImage"`
	// desired kubernetes version to run (default "1.23.6")
	KubernetesVersion *string `pulumi:"kubernetesVersion"`
	// enable kubespan feature
	Kubespan *bool `pulumi:"kubespan"`
	// the desired persist value for configs (default true)
	Persist *bool `pulumi:"persist"`
	// list of registry mirrors to use in format: <registry host>=<mirror URL>
	RegistryMirrors []string `pulumi:"registryMirrors"`
	// Talos Secrets Bundle
	Secrets SecretsBundle `pulumi:"secrets"`
	// the desired Talos version to refer to
	TalosVersion *string `pulumi:"talosVersion"`
}

// The set of arguments for constructing a ClusterConfig resource.
type ClusterConfigArgs struct {
	// additional Subject-Alt-Names for the APIServer certificate
	AdditionalSans pulumi.StringArrayInput
	// enable cluster discovery feature (default true)
	ClusterDiscovery pulumi.BoolPtrInput
	// The cluster endpoint is the URL for the Kubernetes API. If you decide to use
	// a control plane node, common in a single node control plane setup, use port 6443 as
	// this is the port that the API server binds to on every control plane node. For an HA
	// setup, usually involving a load balancer, use the IP and port of the load balancer.
	ClusterEndpoint pulumi.StringInput
	// cluster name
	ClusterName pulumi.StringInput
	// patch generated machineconfigs (applied to all node types)
	ConfigPatches ConfigPatchesPtrInput
	// patch generated machineconfigs (applied to 'controlplane' types)
	ConfigPatchesControlPlane ConfigPatchesPtrInput
	// patch generated machineconfigs (applied to 'worker' type)
	ConfigPatchesWorker ConfigPatchesPtrInput
	// the desired machine config version to refer to
	ConfigVersion pulumi.StringPtrInput
	// the dns domain to use for cluster (default "cluster.local")
	DnsDomain pulumi.StringPtrInput
	// renders all machine configs adding the documentation for each field (default true)
	Docs pulumi.BoolPtrInput
	// renders all machine configs with the commented examples (default true)
	Examples pulumi.BoolPtrInput
	// the disk to install to (default "/dev/sda")
	InstallDisk pulumi.StringPtrInput
	// the image used to perform an installation (default "ghcr.io/talos-systems/installer:v0.14.2")
	InstallImage pulumi.StringPtrInput
	// desired kubernetes version to run (default "1.23.6")
	KubernetesVersion pulumi.StringPtrInput
	// enable kubespan feature
	Kubespan pulumi.BoolPtrInput
	// the desired persist value for configs (default true)
	Persist pulumi.BoolPtrInput
	// list of registry mirrors to use in format: <registry host>=<mirror URL>
	RegistryMirrors pulumi.StringArrayInput
	// Talos Secrets Bundle
	Secrets SecretsBundleInput
	// the desired Talos version to refer to
	TalosVersion pulumi.StringPtrInput
}

func (ClusterConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterConfigArgs)(nil)).Elem()
}

type ClusterConfigInput interface {
	pulumi.Input

	ToClusterConfigOutput() ClusterConfigOutput
	ToClusterConfigOutputWithContext(ctx context.Context) ClusterConfigOutput
}

func (*ClusterConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterConfig)(nil)).Elem()
}

func (i *ClusterConfig) ToClusterConfigOutput() ClusterConfigOutput {
	return i.ToClusterConfigOutputWithContext(context.Background())
}

func (i *ClusterConfig) ToClusterConfigOutputWithContext(ctx context.Context) ClusterConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterConfigOutput)
}

// ClusterConfigArrayInput is an input type that accepts ClusterConfigArray and ClusterConfigArrayOutput values.
// You can construct a concrete instance of `ClusterConfigArrayInput` via:
//
//          ClusterConfigArray{ ClusterConfigArgs{...} }
type ClusterConfigArrayInput interface {
	pulumi.Input

	ToClusterConfigArrayOutput() ClusterConfigArrayOutput
	ToClusterConfigArrayOutputWithContext(context.Context) ClusterConfigArrayOutput
}

type ClusterConfigArray []ClusterConfigInput

func (ClusterConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterConfig)(nil)).Elem()
}

func (i ClusterConfigArray) ToClusterConfigArrayOutput() ClusterConfigArrayOutput {
	return i.ToClusterConfigArrayOutputWithContext(context.Background())
}

func (i ClusterConfigArray) ToClusterConfigArrayOutputWithContext(ctx context.Context) ClusterConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterConfigArrayOutput)
}

// ClusterConfigMapInput is an input type that accepts ClusterConfigMap and ClusterConfigMapOutput values.
// You can construct a concrete instance of `ClusterConfigMapInput` via:
//
//          ClusterConfigMap{ "key": ClusterConfigArgs{...} }
type ClusterConfigMapInput interface {
	pulumi.Input

	ToClusterConfigMapOutput() ClusterConfigMapOutput
	ToClusterConfigMapOutputWithContext(context.Context) ClusterConfigMapOutput
}

type ClusterConfigMap map[string]ClusterConfigInput

func (ClusterConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterConfig)(nil)).Elem()
}

func (i ClusterConfigMap) ToClusterConfigMapOutput() ClusterConfigMapOutput {
	return i.ToClusterConfigMapOutputWithContext(context.Background())
}

func (i ClusterConfigMap) ToClusterConfigMapOutputWithContext(ctx context.Context) ClusterConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterConfigMapOutput)
}

type ClusterConfigOutput struct{ *pulumi.OutputState }

func (ClusterConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterConfig)(nil)).Elem()
}

func (o ClusterConfigOutput) ToClusterConfigOutput() ClusterConfigOutput {
	return o
}

func (o ClusterConfigOutput) ToClusterConfigOutputWithContext(ctx context.Context) ClusterConfigOutput {
	return o
}

// additional Subject-Alt-Names for the APIServer certificate
func (o ClusterConfigOutput) AdditionalSans() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ClusterConfig) pulumi.StringArrayOutput { return v.AdditionalSans }).(pulumi.StringArrayOutput)
}

// cluster discovery feature
func (o ClusterConfigOutput) ClusterDiscovery() pulumi.BoolOutput {
	return o.ApplyT(func(v *ClusterConfig) pulumi.BoolOutput { return v.ClusterDiscovery }).(pulumi.BoolOutput)
}

// The cluster endpoint
func (o ClusterConfigOutput) ClusterEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterConfig) pulumi.StringOutput { return v.ClusterEndpoint }).(pulumi.StringOutput)
}

// cluster name
func (o ClusterConfigOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterConfig) pulumi.StringOutput { return v.ClusterName }).(pulumi.StringOutput)
}

// generated machineconfigs (applied to all node types)
func (o ClusterConfigOutput) ConfigPatches() ConfigPatchesOutput {
	return o.ApplyT(func(v *ClusterConfig) ConfigPatchesOutput { return v.ConfigPatches }).(ConfigPatchesOutput)
}

// generated machineconfigs (applied to 'controlplane' types)
func (o ClusterConfigOutput) ConfigPatchesControlPlane() ConfigPatchesOutput {
	return o.ApplyT(func(v *ClusterConfig) ConfigPatchesOutput { return v.ConfigPatchesControlPlane }).(ConfigPatchesOutput)
}

// generated machineconfigs (applied to 'worker' type)
func (o ClusterConfigOutput) ConfigPatchesWorker() ConfigPatchesOutput {
	return o.ApplyT(func(v *ClusterConfig) ConfigPatchesOutput { return v.ConfigPatchesWorker }).(ConfigPatchesOutput)
}

// the desired machine config version to refer to
func (o ClusterConfigOutput) ConfigVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterConfig) pulumi.StringOutput { return v.ConfigVersion }).(pulumi.StringOutput)
}

// Talos Controlplane Config
func (o ClusterConfigOutput) ControlplaneConfig() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterConfig) pulumi.StringOutput { return v.ControlplaneConfig }).(pulumi.StringOutput)
}

// the dns domain to use for cluster
func (o ClusterConfigOutput) DnsDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterConfig) pulumi.StringOutput { return v.DnsDomain }).(pulumi.StringOutput)
}

// machine config documentation enabled
func (o ClusterConfigOutput) Docs() pulumi.BoolOutput {
	return o.ApplyT(func(v *ClusterConfig) pulumi.BoolOutput { return v.Docs }).(pulumi.BoolOutput)
}

// machine config examples enabled
func (o ClusterConfigOutput) Examples() pulumi.BoolOutput {
	return o.ApplyT(func(v *ClusterConfig) pulumi.BoolOutput { return v.Examples }).(pulumi.BoolOutput)
}

// the disk to install to
func (o ClusterConfigOutput) InstallDisk() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterConfig) pulumi.StringOutput { return v.InstallDisk }).(pulumi.StringOutput)
}

// the image used to perform an installation
func (o ClusterConfigOutput) InstallImage() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterConfig) pulumi.StringOutput { return v.InstallImage }).(pulumi.StringOutput)
}

// desired kubernetes version to run
func (o ClusterConfigOutput) KubernetesVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterConfig) pulumi.StringOutput { return v.KubernetesVersion }).(pulumi.StringOutput)
}

// kubespan enabled
func (o ClusterConfigOutput) Kubespan() pulumi.BoolOutput {
	return o.ApplyT(func(v *ClusterConfig) pulumi.BoolOutput { return v.Kubespan }).(pulumi.BoolOutput)
}

// persist value for configs
func (o ClusterConfigOutput) Persist() pulumi.BoolOutput {
	return o.ApplyT(func(v *ClusterConfig) pulumi.BoolOutput { return v.Persist }).(pulumi.BoolOutput)
}

// list of registry mirrors
func (o ClusterConfigOutput) RegistryMirrors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ClusterConfig) pulumi.StringArrayOutput { return v.RegistryMirrors }).(pulumi.StringArrayOutput)
}

// Talos Secrets Bundle
func (o ClusterConfigOutput) Secrets() SecretsBundleOutput {
	return o.ApplyT(func(v *ClusterConfig) SecretsBundleOutput { return v.Secrets }).(SecretsBundleOutput)
}

// Talos Config
func (o ClusterConfigOutput) TalosConfig() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterConfig) pulumi.StringOutput { return v.TalosConfig }).(pulumi.StringOutput)
}

// the desired Talos version
func (o ClusterConfigOutput) TalosVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterConfig) pulumi.StringOutput { return v.TalosVersion }).(pulumi.StringOutput)
}

// Talos Worker Config
func (o ClusterConfigOutput) WorkerConfig() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterConfig) pulumi.StringOutput { return v.WorkerConfig }).(pulumi.StringOutput)
}

type ClusterConfigArrayOutput struct{ *pulumi.OutputState }

func (ClusterConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterConfig)(nil)).Elem()
}

func (o ClusterConfigArrayOutput) ToClusterConfigArrayOutput() ClusterConfigArrayOutput {
	return o
}

func (o ClusterConfigArrayOutput) ToClusterConfigArrayOutputWithContext(ctx context.Context) ClusterConfigArrayOutput {
	return o
}

func (o ClusterConfigArrayOutput) Index(i pulumi.IntInput) ClusterConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClusterConfig {
		return vs[0].([]*ClusterConfig)[vs[1].(int)]
	}).(ClusterConfigOutput)
}

type ClusterConfigMapOutput struct{ *pulumi.OutputState }

func (ClusterConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterConfig)(nil)).Elem()
}

func (o ClusterConfigMapOutput) ToClusterConfigMapOutput() ClusterConfigMapOutput {
	return o
}

func (o ClusterConfigMapOutput) ToClusterConfigMapOutputWithContext(ctx context.Context) ClusterConfigMapOutput {
	return o
}

func (o ClusterConfigMapOutput) MapIndex(k pulumi.StringInput) ClusterConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClusterConfig {
		return vs[0].(map[string]*ClusterConfig)[vs[1].(string)]
	}).(ClusterConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterConfigInput)(nil)).Elem(), &ClusterConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterConfigArrayInput)(nil)).Elem(), ClusterConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterConfigMapInput)(nil)).Elem(), ClusterConfigMap{})
	pulumi.RegisterOutputType(ClusterConfigOutput{})
	pulumi.RegisterOutputType(ClusterConfigArrayOutput{})
	pulumi.RegisterOutputType(ClusterConfigMapOutput{})
}
