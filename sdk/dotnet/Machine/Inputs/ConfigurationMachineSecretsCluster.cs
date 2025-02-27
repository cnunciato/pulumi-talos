// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Talos.Machine.Inputs
{

    public sealed class ConfigurationMachineSecretsClusterArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The cluster id
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        [Input("secret", required: true)]
        private string? _secret;

        /// <summary>
        /// The cluster secret
        /// </summary>
        public string? Secret
        {
            get => _secret;
            set => _secret = value;
        }

        public ConfigurationMachineSecretsClusterArgs()
        {
        }
        public static new ConfigurationMachineSecretsClusterArgs Empty => new ConfigurationMachineSecretsClusterArgs();
    }
}
