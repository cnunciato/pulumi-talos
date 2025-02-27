// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package machine

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/pulumiverse/pulumi-talos/sdk/go/talos/internal"
)

// Generate machine secrets for Talos cluster.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-talos/sdk/go/talos/machine"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := machine.NewSecrets(ctx, "machineSecrets", nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// terraform machine secrets can be imported from an existing secrets file
//
// ```sh
//
//	$ pulumi import talos:machine/secrets:Secrets this <path-to-secrets.yaml>
//
// ```
type Secrets struct {
	pulumi.CustomResourceState

	// The generated client configuration data
	ClientConfiguration SecretsClientConfigurationOutput `pulumi:"clientConfiguration"`
	// The secrets for the talos cluster
	MachineSecrets SecretsMachineSecretsOutput `pulumi:"machineSecrets"`
	// The version of talos features to use in generated machine configuration
	TalosVersion pulumi.StringOutput `pulumi:"talosVersion"`
}

// NewSecrets registers a new resource with the given unique name, arguments, and options.
func NewSecrets(ctx *pulumi.Context,
	name string, args *SecretsArgs, opts ...pulumi.ResourceOption) (*Secrets, error) {
	if args == nil {
		args = &SecretsArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Secrets
	err := ctx.RegisterResource("talos:machine/secrets:Secrets", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecrets gets an existing Secrets resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecrets(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretsState, opts ...pulumi.ResourceOption) (*Secrets, error) {
	var resource Secrets
	err := ctx.ReadResource("talos:machine/secrets:Secrets", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Secrets resources.
type secretsState struct {
	// The generated client configuration data
	ClientConfiguration *SecretsClientConfiguration `pulumi:"clientConfiguration"`
	// The secrets for the talos cluster
	MachineSecrets *SecretsMachineSecrets `pulumi:"machineSecrets"`
	// The version of talos features to use in generated machine configuration
	TalosVersion *string `pulumi:"talosVersion"`
}

type SecretsState struct {
	// The generated client configuration data
	ClientConfiguration SecretsClientConfigurationPtrInput
	// The secrets for the talos cluster
	MachineSecrets SecretsMachineSecretsPtrInput
	// The version of talos features to use in generated machine configuration
	TalosVersion pulumi.StringPtrInput
}

func (SecretsState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretsState)(nil)).Elem()
}

type secretsArgs struct {
	// The version of talos features to use in generated machine configuration
	TalosVersion *string `pulumi:"talosVersion"`
}

// The set of arguments for constructing a Secrets resource.
type SecretsArgs struct {
	// The version of talos features to use in generated machine configuration
	TalosVersion pulumi.StringPtrInput
}

func (SecretsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretsArgs)(nil)).Elem()
}

type SecretsInput interface {
	pulumi.Input

	ToSecretsOutput() SecretsOutput
	ToSecretsOutputWithContext(ctx context.Context) SecretsOutput
}

func (*Secrets) ElementType() reflect.Type {
	return reflect.TypeOf((**Secrets)(nil)).Elem()
}

func (i *Secrets) ToSecretsOutput() SecretsOutput {
	return i.ToSecretsOutputWithContext(context.Background())
}

func (i *Secrets) ToSecretsOutputWithContext(ctx context.Context) SecretsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretsOutput)
}

func (i *Secrets) ToOutput(ctx context.Context) pulumix.Output[*Secrets] {
	return pulumix.Output[*Secrets]{
		OutputState: i.ToSecretsOutputWithContext(ctx).OutputState,
	}
}

// SecretsArrayInput is an input type that accepts SecretsArray and SecretsArrayOutput values.
// You can construct a concrete instance of `SecretsArrayInput` via:
//
//	SecretsArray{ SecretsArgs{...} }
type SecretsArrayInput interface {
	pulumi.Input

	ToSecretsArrayOutput() SecretsArrayOutput
	ToSecretsArrayOutputWithContext(context.Context) SecretsArrayOutput
}

type SecretsArray []SecretsInput

func (SecretsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Secrets)(nil)).Elem()
}

func (i SecretsArray) ToSecretsArrayOutput() SecretsArrayOutput {
	return i.ToSecretsArrayOutputWithContext(context.Background())
}

func (i SecretsArray) ToSecretsArrayOutputWithContext(ctx context.Context) SecretsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretsArrayOutput)
}

func (i SecretsArray) ToOutput(ctx context.Context) pulumix.Output[[]*Secrets] {
	return pulumix.Output[[]*Secrets]{
		OutputState: i.ToSecretsArrayOutputWithContext(ctx).OutputState,
	}
}

// SecretsMapInput is an input type that accepts SecretsMap and SecretsMapOutput values.
// You can construct a concrete instance of `SecretsMapInput` via:
//
//	SecretsMap{ "key": SecretsArgs{...} }
type SecretsMapInput interface {
	pulumi.Input

	ToSecretsMapOutput() SecretsMapOutput
	ToSecretsMapOutputWithContext(context.Context) SecretsMapOutput
}

type SecretsMap map[string]SecretsInput

func (SecretsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Secrets)(nil)).Elem()
}

func (i SecretsMap) ToSecretsMapOutput() SecretsMapOutput {
	return i.ToSecretsMapOutputWithContext(context.Background())
}

func (i SecretsMap) ToSecretsMapOutputWithContext(ctx context.Context) SecretsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretsMapOutput)
}

func (i SecretsMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Secrets] {
	return pulumix.Output[map[string]*Secrets]{
		OutputState: i.ToSecretsMapOutputWithContext(ctx).OutputState,
	}
}

type SecretsOutput struct{ *pulumi.OutputState }

func (SecretsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Secrets)(nil)).Elem()
}

func (o SecretsOutput) ToSecretsOutput() SecretsOutput {
	return o
}

func (o SecretsOutput) ToSecretsOutputWithContext(ctx context.Context) SecretsOutput {
	return o
}

func (o SecretsOutput) ToOutput(ctx context.Context) pulumix.Output[*Secrets] {
	return pulumix.Output[*Secrets]{
		OutputState: o.OutputState,
	}
}

// The generated client configuration data
func (o SecretsOutput) ClientConfiguration() SecretsClientConfigurationOutput {
	return o.ApplyT(func(v *Secrets) SecretsClientConfigurationOutput { return v.ClientConfiguration }).(SecretsClientConfigurationOutput)
}

// The secrets for the talos cluster
func (o SecretsOutput) MachineSecrets() SecretsMachineSecretsOutput {
	return o.ApplyT(func(v *Secrets) SecretsMachineSecretsOutput { return v.MachineSecrets }).(SecretsMachineSecretsOutput)
}

// The version of talos features to use in generated machine configuration
func (o SecretsOutput) TalosVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Secrets) pulumi.StringOutput { return v.TalosVersion }).(pulumi.StringOutput)
}

type SecretsArrayOutput struct{ *pulumi.OutputState }

func (SecretsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Secrets)(nil)).Elem()
}

func (o SecretsArrayOutput) ToSecretsArrayOutput() SecretsArrayOutput {
	return o
}

func (o SecretsArrayOutput) ToSecretsArrayOutputWithContext(ctx context.Context) SecretsArrayOutput {
	return o
}

func (o SecretsArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Secrets] {
	return pulumix.Output[[]*Secrets]{
		OutputState: o.OutputState,
	}
}

func (o SecretsArrayOutput) Index(i pulumi.IntInput) SecretsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Secrets {
		return vs[0].([]*Secrets)[vs[1].(int)]
	}).(SecretsOutput)
}

type SecretsMapOutput struct{ *pulumi.OutputState }

func (SecretsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Secrets)(nil)).Elem()
}

func (o SecretsMapOutput) ToSecretsMapOutput() SecretsMapOutput {
	return o
}

func (o SecretsMapOutput) ToSecretsMapOutputWithContext(ctx context.Context) SecretsMapOutput {
	return o
}

func (o SecretsMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Secrets] {
	return pulumix.Output[map[string]*Secrets]{
		OutputState: o.OutputState,
	}
}

func (o SecretsMapOutput) MapIndex(k pulumi.StringInput) SecretsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Secrets {
		return vs[0].(map[string]*Secrets)[vs[1].(string)]
	}).(SecretsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretsInput)(nil)).Elem(), &Secrets{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretsArrayInput)(nil)).Elem(), SecretsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretsMapInput)(nil)).Elem(), SecretsMap{})
	pulumi.RegisterOutputType(SecretsOutput{})
	pulumi.RegisterOutputType(SecretsArrayOutput{})
	pulumi.RegisterOutputType(SecretsMapOutput{})
}
