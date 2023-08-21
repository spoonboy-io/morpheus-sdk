package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ApiKeyPairsKeyPair {
    
    String name
    
    String publicKey
    
    String privateKey
    
    String passphrase
}
