package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ProvisioningSettingsUpdateCloudInitKeyPair;
import org.openapitools.model.ProvisioningSettingsUpdateDefaultTemplateType;
import org.openapitools.model.ProvisioningSettingsUpdateDeployStorageProvider;

@Canonical
class ProvisioningSettingsUpdate {
    /* Use this to enable / disable allowing cloud selection */
    Boolean allowZoneSelection
    /* Use this to enable / disable allowing host selection */
    Boolean allowServerSelection
    /* Use this to enable / disable requiring environment selection */
    Boolean requireEnvironments
    /* Use this to enable / disable showing pricing */
    Boolean showPricing
    /* Use this to enable / disable hiding datastore stats */
    Boolean hideDatastoreStats
    /* Use this to enable / disable cross-tenant naming policies */
    Boolean crossTenantNamingPolicies
    /* Use this to enable / disable reusing naming sequence numbers */
    Boolean reuseSequence
    /* Cloud-init username */
    String cloudInitUsername
    /* Cloud-init password */
    String cloudInitPassword
    
    ProvisioningSettingsUpdateCloudInitKeyPair cloudInitKeyPair
    
    ProvisioningSettingsUpdateDeployStorageProvider deployStorageProvider
    /* Windows administrator password */
    String windowsPassword
    /* PXE Boot default root password */
    String pxeRootPassword
    
    ProvisioningSettingsUpdateDefaultTemplateType defaultTemplateType
}
