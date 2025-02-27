// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Talos.Machine.Outputs
{

    [OutputType]
    public sealed class ConfigurationMachineSecretsClusterResult
    {
        /// <summary>
        /// The cluster id
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The cluster secret
        /// </summary>
        public readonly string Secret;

        [OutputConstructor]
        private ConfigurationMachineSecretsClusterResult(
            string id,

            string secret)
        {
            Id = id;
            Secret = secret;
        }
    }
}
