package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class UserSourceCreateCustomExternal {
    /* External Login URL */
    String loginUrl
    /* Do not include SAMLRequest */
    Boolean doNotIncludeSAMLRequest = false
    /* External Logout URL */
    String logout
    /* Encryption Algorithm */
    String encryptionAlgo
    /* Encryption Key */
    String encryptionKey
}
