package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class UserSourceCreateCustomApi {
    /* API Endpoint */
    String endpoint
    /* API Style */
    String apiStyle
    /* Encryption Algorithm */
    String encryptionAlgo
    /* Encryption Key */
    String encryptionKey
}
