// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package machine

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/pulumiverse/pulumi-talos/sdk/go/talos/internal"
)

// The machine bootstrap resource allows you to bootstrap a Talos node.
//
// ## Import
//
// terraform machine bootstrap can be imported to let terraform know that the machine is already bootstrapped
//
// ```sh
//
//	$ pulumi import talos:machine/bootstrap:Bootstrap this <any id>
//
// ```
type Bootstrap struct {
	pulumi.CustomResourceState

	// The client configuration data
	ClientConfiguration BootstrapClientConfigurationOutput `pulumi:"clientConfiguration"`
	// The endpoint of the machine to bootstrap
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// The name of the node to bootstrap
	Node     pulumi.StringOutput        `pulumi:"node"`
	Timeouts BootstrapTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewBootstrap registers a new resource with the given unique name, arguments, and options.
func NewBootstrap(ctx *pulumi.Context,
	name string, args *BootstrapArgs, opts ...pulumi.ResourceOption) (*Bootstrap, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'ClientConfiguration'")
	}
	if args.Node == nil {
		return nil, errors.New("invalid value for required argument 'Node'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Bootstrap
	err := ctx.RegisterResource("talos:machine/bootstrap:Bootstrap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBootstrap gets an existing Bootstrap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBootstrap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BootstrapState, opts ...pulumi.ResourceOption) (*Bootstrap, error) {
	var resource Bootstrap
	err := ctx.ReadResource("talos:machine/bootstrap:Bootstrap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Bootstrap resources.
type bootstrapState struct {
	// The client configuration data
	ClientConfiguration *BootstrapClientConfiguration `pulumi:"clientConfiguration"`
	// The endpoint of the machine to bootstrap
	Endpoint *string `pulumi:"endpoint"`
	// The name of the node to bootstrap
	Node     *string            `pulumi:"node"`
	Timeouts *BootstrapTimeouts `pulumi:"timeouts"`
}

type BootstrapState struct {
	// The client configuration data
	ClientConfiguration BootstrapClientConfigurationPtrInput
	// The endpoint of the machine to bootstrap
	Endpoint pulumi.StringPtrInput
	// The name of the node to bootstrap
	Node     pulumi.StringPtrInput
	Timeouts BootstrapTimeoutsPtrInput
}

func (BootstrapState) ElementType() reflect.Type {
	return reflect.TypeOf((*bootstrapState)(nil)).Elem()
}

type bootstrapArgs struct {
	// The client configuration data
	ClientConfiguration BootstrapClientConfiguration `pulumi:"clientConfiguration"`
	// The endpoint of the machine to bootstrap
	Endpoint *string `pulumi:"endpoint"`
	// The name of the node to bootstrap
	Node     string             `pulumi:"node"`
	Timeouts *BootstrapTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a Bootstrap resource.
type BootstrapArgs struct {
	// The client configuration data
	ClientConfiguration BootstrapClientConfigurationInput
	// The endpoint of the machine to bootstrap
	Endpoint pulumi.StringPtrInput
	// The name of the node to bootstrap
	Node     pulumi.StringInput
	Timeouts BootstrapTimeoutsPtrInput
}

func (BootstrapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bootstrapArgs)(nil)).Elem()
}

type BootstrapInput interface {
	pulumi.Input

	ToBootstrapOutput() BootstrapOutput
	ToBootstrapOutputWithContext(ctx context.Context) BootstrapOutput
}

func (*Bootstrap) ElementType() reflect.Type {
	return reflect.TypeOf((**Bootstrap)(nil)).Elem()
}

func (i *Bootstrap) ToBootstrapOutput() BootstrapOutput {
	return i.ToBootstrapOutputWithContext(context.Background())
}

func (i *Bootstrap) ToBootstrapOutputWithContext(ctx context.Context) BootstrapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BootstrapOutput)
}

func (i *Bootstrap) ToOutput(ctx context.Context) pulumix.Output[*Bootstrap] {
	return pulumix.Output[*Bootstrap]{
		OutputState: i.ToBootstrapOutputWithContext(ctx).OutputState,
	}
}

// BootstrapArrayInput is an input type that accepts BootstrapArray and BootstrapArrayOutput values.
// You can construct a concrete instance of `BootstrapArrayInput` via:
//
//	BootstrapArray{ BootstrapArgs{...} }
type BootstrapArrayInput interface {
	pulumi.Input

	ToBootstrapArrayOutput() BootstrapArrayOutput
	ToBootstrapArrayOutputWithContext(context.Context) BootstrapArrayOutput
}

type BootstrapArray []BootstrapInput

func (BootstrapArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Bootstrap)(nil)).Elem()
}

func (i BootstrapArray) ToBootstrapArrayOutput() BootstrapArrayOutput {
	return i.ToBootstrapArrayOutputWithContext(context.Background())
}

func (i BootstrapArray) ToBootstrapArrayOutputWithContext(ctx context.Context) BootstrapArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BootstrapArrayOutput)
}

func (i BootstrapArray) ToOutput(ctx context.Context) pulumix.Output[[]*Bootstrap] {
	return pulumix.Output[[]*Bootstrap]{
		OutputState: i.ToBootstrapArrayOutputWithContext(ctx).OutputState,
	}
}

// BootstrapMapInput is an input type that accepts BootstrapMap and BootstrapMapOutput values.
// You can construct a concrete instance of `BootstrapMapInput` via:
//
//	BootstrapMap{ "key": BootstrapArgs{...} }
type BootstrapMapInput interface {
	pulumi.Input

	ToBootstrapMapOutput() BootstrapMapOutput
	ToBootstrapMapOutputWithContext(context.Context) BootstrapMapOutput
}

type BootstrapMap map[string]BootstrapInput

func (BootstrapMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Bootstrap)(nil)).Elem()
}

func (i BootstrapMap) ToBootstrapMapOutput() BootstrapMapOutput {
	return i.ToBootstrapMapOutputWithContext(context.Background())
}

func (i BootstrapMap) ToBootstrapMapOutputWithContext(ctx context.Context) BootstrapMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BootstrapMapOutput)
}

func (i BootstrapMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Bootstrap] {
	return pulumix.Output[map[string]*Bootstrap]{
		OutputState: i.ToBootstrapMapOutputWithContext(ctx).OutputState,
	}
}

type BootstrapOutput struct{ *pulumi.OutputState }

func (BootstrapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Bootstrap)(nil)).Elem()
}

func (o BootstrapOutput) ToBootstrapOutput() BootstrapOutput {
	return o
}

func (o BootstrapOutput) ToBootstrapOutputWithContext(ctx context.Context) BootstrapOutput {
	return o
}

func (o BootstrapOutput) ToOutput(ctx context.Context) pulumix.Output[*Bootstrap] {
	return pulumix.Output[*Bootstrap]{
		OutputState: o.OutputState,
	}
}

// The client configuration data
func (o BootstrapOutput) ClientConfiguration() BootstrapClientConfigurationOutput {
	return o.ApplyT(func(v *Bootstrap) BootstrapClientConfigurationOutput { return v.ClientConfiguration }).(BootstrapClientConfigurationOutput)
}

// The endpoint of the machine to bootstrap
func (o BootstrapOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Bootstrap) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// The name of the node to bootstrap
func (o BootstrapOutput) Node() pulumi.StringOutput {
	return o.ApplyT(func(v *Bootstrap) pulumi.StringOutput { return v.Node }).(pulumi.StringOutput)
}

func (o BootstrapOutput) Timeouts() BootstrapTimeoutsPtrOutput {
	return o.ApplyT(func(v *Bootstrap) BootstrapTimeoutsPtrOutput { return v.Timeouts }).(BootstrapTimeoutsPtrOutput)
}

type BootstrapArrayOutput struct{ *pulumi.OutputState }

func (BootstrapArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Bootstrap)(nil)).Elem()
}

func (o BootstrapArrayOutput) ToBootstrapArrayOutput() BootstrapArrayOutput {
	return o
}

func (o BootstrapArrayOutput) ToBootstrapArrayOutputWithContext(ctx context.Context) BootstrapArrayOutput {
	return o
}

func (o BootstrapArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Bootstrap] {
	return pulumix.Output[[]*Bootstrap]{
		OutputState: o.OutputState,
	}
}

func (o BootstrapArrayOutput) Index(i pulumi.IntInput) BootstrapOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Bootstrap {
		return vs[0].([]*Bootstrap)[vs[1].(int)]
	}).(BootstrapOutput)
}

type BootstrapMapOutput struct{ *pulumi.OutputState }

func (BootstrapMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Bootstrap)(nil)).Elem()
}

func (o BootstrapMapOutput) ToBootstrapMapOutput() BootstrapMapOutput {
	return o
}

func (o BootstrapMapOutput) ToBootstrapMapOutputWithContext(ctx context.Context) BootstrapMapOutput {
	return o
}

func (o BootstrapMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Bootstrap] {
	return pulumix.Output[map[string]*Bootstrap]{
		OutputState: o.OutputState,
	}
}

func (o BootstrapMapOutput) MapIndex(k pulumi.StringInput) BootstrapOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Bootstrap {
		return vs[0].(map[string]*Bootstrap)[vs[1].(string)]
	}).(BootstrapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BootstrapInput)(nil)).Elem(), &Bootstrap{})
	pulumi.RegisterInputType(reflect.TypeOf((*BootstrapArrayInput)(nil)).Elem(), BootstrapArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BootstrapMapInput)(nil)).Elem(), BootstrapMap{})
	pulumi.RegisterOutputType(BootstrapOutput{})
	pulumi.RegisterOutputType(BootstrapArrayOutput{})
	pulumi.RegisterOutputType(BootstrapMapOutput{})
}
