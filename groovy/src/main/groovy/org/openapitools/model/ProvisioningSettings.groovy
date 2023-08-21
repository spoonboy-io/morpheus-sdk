package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InlineResponse20040AppDeployInstance;
import org.openapitools.model.InlineResponse20094Network;

@Canonical
class ProvisioningSettings {
    
    Boolean allowZoneSelection
    
    Boolean allowServerSelection
    
    Boolean requireEnvironments
    
    Boolean showPricing
    
    Boolean hideDatastoreStats
    
    Boolean crossTenantNamingPolicies
    
    Boolean reuseSequence
    
    String cloudInitUsername
    
    String cloudInitPassword
    
    InlineResponse20040AppDeployInstance cloudInitKeyPair
    
    String windowsPassword
    
    String pxeRootPassword
    
    InlineResponse20094Network defaultTemplateType
    
    InlineResponse20040AppDeployInstance deployStorageProvider
}
