package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class IdentitySourcesAzureADConfigConfig {
    
    String url
    
    String logoutUrl
    
    Boolean doNotIncludeSAMLRequest
    
    String saMLSignatureMode
    
    Boolean doNotValidateSignature
    
    Boolean doNotValidateStatusCode
    
    Boolean doNotValidateDestination
    
    Boolean doNotValidateIssueInstants
    
    Boolean doNotValidateAssertions
    
    String givenNameAttribute
    
    String surnameAttribute
    
    String emailAttribute
    
    String requiredAttributeValue
    
    String roleAttributeName
    
    String azureTenantId
    
    String azureAppId
    
    String azureAppSecret
    
    String roleLinkAttributeName
    
    String publicKey
    
    String azureAppSecretHash
}
