// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package talos

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Generate client configuration for a Talos cluster
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-talos/sdk/go/talos"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			machineSecrets, err := talos.NewTalosMachineSecrets(ctx, "machineSecrets", nil)
//			if err != nil {
//				return err
//			}
//			_, err = talos.NewTalosClientConfiguration(ctx, "talosconfig", &talos.TalosClientConfigurationArgs{
//				ClusterName:    pulumi.String("example-cluster"),
//				MachineSecrets: machineSecrets.MachineSecrets,
//				Endpoints: pulumi.StringArray{
//					pulumi.String("10.5.0.2"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type TalosClientConfiguration struct {
	pulumi.CustomResourceState

	// The name of the cluster in the generated config
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`
	// endpoints to set in the generated config
	Endpoints pulumi.StringArrayOutput `pulumi:"endpoints"`
	// The machine secrets for the cluster
	MachineSecrets pulumi.StringOutput `pulumi:"machineSecrets"`
	// nodes to set in the generated config
	Nodes pulumi.StringArrayOutput `pulumi:"nodes"`
	// The generated talos config
	TalosConfig pulumi.StringOutput `pulumi:"talosConfig"`
}

// NewTalosClientConfiguration registers a new resource with the given unique name, arguments, and options.
func NewTalosClientConfiguration(ctx *pulumi.Context,
	name string, args *TalosClientConfigurationArgs, opts ...pulumi.ResourceOption) (*TalosClientConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.MachineSecrets == nil {
		return nil, errors.New("invalid value for required argument 'MachineSecrets'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"talosConfig",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource TalosClientConfiguration
	err := ctx.RegisterResource("talos:index/talosClientConfiguration:TalosClientConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTalosClientConfiguration gets an existing TalosClientConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTalosClientConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TalosClientConfigurationState, opts ...pulumi.ResourceOption) (*TalosClientConfiguration, error) {
	var resource TalosClientConfiguration
	err := ctx.ReadResource("talos:index/talosClientConfiguration:TalosClientConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TalosClientConfiguration resources.
type talosClientConfigurationState struct {
	// The name of the cluster in the generated config
	ClusterName *string `pulumi:"clusterName"`
	// endpoints to set in the generated config
	Endpoints []string `pulumi:"endpoints"`
	// The machine secrets for the cluster
	MachineSecrets *string `pulumi:"machineSecrets"`
	// nodes to set in the generated config
	Nodes []string `pulumi:"nodes"`
	// The generated talos config
	TalosConfig *string `pulumi:"talosConfig"`
}

type TalosClientConfigurationState struct {
	// The name of the cluster in the generated config
	ClusterName pulumi.StringPtrInput
	// endpoints to set in the generated config
	Endpoints pulumi.StringArrayInput
	// The machine secrets for the cluster
	MachineSecrets pulumi.StringPtrInput
	// nodes to set in the generated config
	Nodes pulumi.StringArrayInput
	// The generated talos config
	TalosConfig pulumi.StringPtrInput
}

func (TalosClientConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*talosClientConfigurationState)(nil)).Elem()
}

type talosClientConfigurationArgs struct {
	// The name of the cluster in the generated config
	ClusterName string `pulumi:"clusterName"`
	// endpoints to set in the generated config
	Endpoints []string `pulumi:"endpoints"`
	// The machine secrets for the cluster
	MachineSecrets string `pulumi:"machineSecrets"`
	// nodes to set in the generated config
	Nodes []string `pulumi:"nodes"`
}

// The set of arguments for constructing a TalosClientConfiguration resource.
type TalosClientConfigurationArgs struct {
	// The name of the cluster in the generated config
	ClusterName pulumi.StringInput
	// endpoints to set in the generated config
	Endpoints pulumi.StringArrayInput
	// The machine secrets for the cluster
	MachineSecrets pulumi.StringInput
	// nodes to set in the generated config
	Nodes pulumi.StringArrayInput
}

func (TalosClientConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*talosClientConfigurationArgs)(nil)).Elem()
}

type TalosClientConfigurationInput interface {
	pulumi.Input

	ToTalosClientConfigurationOutput() TalosClientConfigurationOutput
	ToTalosClientConfigurationOutputWithContext(ctx context.Context) TalosClientConfigurationOutput
}

func (*TalosClientConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**TalosClientConfiguration)(nil)).Elem()
}

func (i *TalosClientConfiguration) ToTalosClientConfigurationOutput() TalosClientConfigurationOutput {
	return i.ToTalosClientConfigurationOutputWithContext(context.Background())
}

func (i *TalosClientConfiguration) ToTalosClientConfigurationOutputWithContext(ctx context.Context) TalosClientConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TalosClientConfigurationOutput)
}

// TalosClientConfigurationArrayInput is an input type that accepts TalosClientConfigurationArray and TalosClientConfigurationArrayOutput values.
// You can construct a concrete instance of `TalosClientConfigurationArrayInput` via:
//
//	TalosClientConfigurationArray{ TalosClientConfigurationArgs{...} }
type TalosClientConfigurationArrayInput interface {
	pulumi.Input

	ToTalosClientConfigurationArrayOutput() TalosClientConfigurationArrayOutput
	ToTalosClientConfigurationArrayOutputWithContext(context.Context) TalosClientConfigurationArrayOutput
}

type TalosClientConfigurationArray []TalosClientConfigurationInput

func (TalosClientConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TalosClientConfiguration)(nil)).Elem()
}

func (i TalosClientConfigurationArray) ToTalosClientConfigurationArrayOutput() TalosClientConfigurationArrayOutput {
	return i.ToTalosClientConfigurationArrayOutputWithContext(context.Background())
}

func (i TalosClientConfigurationArray) ToTalosClientConfigurationArrayOutputWithContext(ctx context.Context) TalosClientConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TalosClientConfigurationArrayOutput)
}

// TalosClientConfigurationMapInput is an input type that accepts TalosClientConfigurationMap and TalosClientConfigurationMapOutput values.
// You can construct a concrete instance of `TalosClientConfigurationMapInput` via:
//
//	TalosClientConfigurationMap{ "key": TalosClientConfigurationArgs{...} }
type TalosClientConfigurationMapInput interface {
	pulumi.Input

	ToTalosClientConfigurationMapOutput() TalosClientConfigurationMapOutput
	ToTalosClientConfigurationMapOutputWithContext(context.Context) TalosClientConfigurationMapOutput
}

type TalosClientConfigurationMap map[string]TalosClientConfigurationInput

func (TalosClientConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TalosClientConfiguration)(nil)).Elem()
}

func (i TalosClientConfigurationMap) ToTalosClientConfigurationMapOutput() TalosClientConfigurationMapOutput {
	return i.ToTalosClientConfigurationMapOutputWithContext(context.Background())
}

func (i TalosClientConfigurationMap) ToTalosClientConfigurationMapOutputWithContext(ctx context.Context) TalosClientConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TalosClientConfigurationMapOutput)
}

type TalosClientConfigurationOutput struct{ *pulumi.OutputState }

func (TalosClientConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TalosClientConfiguration)(nil)).Elem()
}

func (o TalosClientConfigurationOutput) ToTalosClientConfigurationOutput() TalosClientConfigurationOutput {
	return o
}

func (o TalosClientConfigurationOutput) ToTalosClientConfigurationOutputWithContext(ctx context.Context) TalosClientConfigurationOutput {
	return o
}

// The name of the cluster in the generated config
func (o TalosClientConfigurationOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v *TalosClientConfiguration) pulumi.StringOutput { return v.ClusterName }).(pulumi.StringOutput)
}

// endpoints to set in the generated config
func (o TalosClientConfigurationOutput) Endpoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TalosClientConfiguration) pulumi.StringArrayOutput { return v.Endpoints }).(pulumi.StringArrayOutput)
}

// The machine secrets for the cluster
func (o TalosClientConfigurationOutput) MachineSecrets() pulumi.StringOutput {
	return o.ApplyT(func(v *TalosClientConfiguration) pulumi.StringOutput { return v.MachineSecrets }).(pulumi.StringOutput)
}

// nodes to set in the generated config
func (o TalosClientConfigurationOutput) Nodes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TalosClientConfiguration) pulumi.StringArrayOutput { return v.Nodes }).(pulumi.StringArrayOutput)
}

// The generated talos config
func (o TalosClientConfigurationOutput) TalosConfig() pulumi.StringOutput {
	return o.ApplyT(func(v *TalosClientConfiguration) pulumi.StringOutput { return v.TalosConfig }).(pulumi.StringOutput)
}

type TalosClientConfigurationArrayOutput struct{ *pulumi.OutputState }

func (TalosClientConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TalosClientConfiguration)(nil)).Elem()
}

func (o TalosClientConfigurationArrayOutput) ToTalosClientConfigurationArrayOutput() TalosClientConfigurationArrayOutput {
	return o
}

func (o TalosClientConfigurationArrayOutput) ToTalosClientConfigurationArrayOutputWithContext(ctx context.Context) TalosClientConfigurationArrayOutput {
	return o
}

func (o TalosClientConfigurationArrayOutput) Index(i pulumi.IntInput) TalosClientConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TalosClientConfiguration {
		return vs[0].([]*TalosClientConfiguration)[vs[1].(int)]
	}).(TalosClientConfigurationOutput)
}

type TalosClientConfigurationMapOutput struct{ *pulumi.OutputState }

func (TalosClientConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TalosClientConfiguration)(nil)).Elem()
}

func (o TalosClientConfigurationMapOutput) ToTalosClientConfigurationMapOutput() TalosClientConfigurationMapOutput {
	return o
}

func (o TalosClientConfigurationMapOutput) ToTalosClientConfigurationMapOutputWithContext(ctx context.Context) TalosClientConfigurationMapOutput {
	return o
}

func (o TalosClientConfigurationMapOutput) MapIndex(k pulumi.StringInput) TalosClientConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TalosClientConfiguration {
		return vs[0].(map[string]*TalosClientConfiguration)[vs[1].(string)]
	}).(TalosClientConfigurationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TalosClientConfigurationInput)(nil)).Elem(), &TalosClientConfiguration{})
	pulumi.RegisterInputType(reflect.TypeOf((*TalosClientConfigurationArrayInput)(nil)).Elem(), TalosClientConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TalosClientConfigurationMapInput)(nil)).Elem(), TalosClientConfigurationMap{})
	pulumi.RegisterOutputType(TalosClientConfigurationOutput{})
	pulumi.RegisterOutputType(TalosClientConfigurationArrayOutput{})
	pulumi.RegisterOutputType(TalosClientConfigurationMapOutput{})
}
