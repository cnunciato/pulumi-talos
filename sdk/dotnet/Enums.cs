// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.Talos
{
    [EnumType]
    public readonly struct TalosMachineConfigApplyMode : IEquatable<TalosMachineConfigApplyMode>
    {
        private readonly string _value;

        private TalosMachineConfigApplyMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// apply config with a reboot
        /// </summary>
        public static TalosMachineConfigApplyMode REBOOT { get; } = new TalosMachineConfigApplyMode("REBOOT");
        /// <summary>
        /// automatically try to apply and reboot if only required
        /// </summary>
        public static TalosMachineConfigApplyMode AUTO { get; } = new TalosMachineConfigApplyMode("AUTO");
        /// <summary>
        /// apply config without a reboot
        /// </summary>
        public static TalosMachineConfigApplyMode NO_REBOOT { get; } = new TalosMachineConfigApplyMode("NO_REBOOT");
        /// <summary>
        /// apply config as staged
        /// </summary>
        public static TalosMachineConfigApplyMode STAGED { get; } = new TalosMachineConfigApplyMode("STAGED");

        public static bool operator ==(TalosMachineConfigApplyMode left, TalosMachineConfigApplyMode right) => left.Equals(right);
        public static bool operator !=(TalosMachineConfigApplyMode left, TalosMachineConfigApplyMode right) => !left.Equals(right);

        public static explicit operator string(TalosMachineConfigApplyMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TalosMachineConfigApplyMode other && Equals(other);
        public bool Equals(TalosMachineConfigApplyMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct TalosMachineConfigVersion : IEquatable<TalosMachineConfigVersion>
    {
        private readonly string _value;

        private TalosMachineConfigVersion(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Talos Machine Configuration Version
        /// </summary>
        public static TalosMachineConfigVersion V1alpha1 { get; } = new TalosMachineConfigVersion("v1alpha1");

        public static bool operator ==(TalosMachineConfigVersion left, TalosMachineConfigVersion right) => left.Equals(right);
        public static bool operator !=(TalosMachineConfigVersion left, TalosMachineConfigVersion right) => !left.Equals(right);

        public static explicit operator string(TalosMachineConfigVersion value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TalosMachineConfigVersion other && Equals(other);
        public bool Equals(TalosMachineConfigVersion other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
