// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const TalosMachineConfigApplyMode = {
    /**
     * apply config with a reboot
     */
    REBOOT: "REBOOT",
    /**
     * automatically try to apply and reboot if only required
     */
    AUTO: "AUTO",
    /**
     * apply config without a reboot
     */
    NO_REBOOT: "NO_REBOOT",
    /**
     * apply config as staged
     */
    STAGED: "STAGED",
} as const;

export type TalosMachineConfigApplyMode = (typeof TalosMachineConfigApplyMode)[keyof typeof TalosMachineConfigApplyMode];

export const TalosMachineConfigVersion = {
    /**
     * Talos Machine Configuration Version
     */
    V1alpha1: "v1alpha1",
} as const;

export type TalosMachineConfigVersion = (typeof TalosMachineConfigVersion)[keyof typeof TalosMachineConfigVersion];
