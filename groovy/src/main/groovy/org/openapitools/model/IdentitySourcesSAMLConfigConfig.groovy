package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class IdentitySourcesSAMLConfigConfig {
    
    String roleAttributeName
    
    String requiredAttributeValue
    
    String givenNameAttribute
    
    String surnameAttribute
    
    String logoutUrl
    
    Boolean doNotIncludeSAMLRequest
    
    String publicKey
    
    String emailAttribute
    
    String url
    
    Boolean doNotValidateSignature
    
    Boolean doNotValidateStatusCode
    
    Boolean doNotValidateDestination
    
    Boolean doNotValidateIssueInstants
    
    Boolean doNotValidateAssertions
    
    Boolean doNotValidateAuthStatements
    
    Boolean doNotValidateSubject
    
    Boolean doNotValidateConditions
    
    Boolean doNotValidateAudiences
    
    Boolean doNotValidateSubjectRecipients
    
    String saMLSignatureMode
    
    Boolean forceAuthn
}
