package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class KeyPair {
    
    Long id
    
    String name
    
    Long accountId
    
    String publicKey
    
    Boolean hasPrivateKey
    
    String privateKeyHash
    /* Only present in response to generate */
    String privateKey
    
    String fingerprint
    
    Date dateCreated
    
    Date lastUpdated
}
