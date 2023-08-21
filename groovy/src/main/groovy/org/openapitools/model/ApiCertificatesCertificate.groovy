package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ApiCertificatesCertificate {
    /* A unique name scoped to your account for the key */
    String name
    /* The contents of the certificate file */
    String certFile
    /* The contents of the key file */
    String keyFile
    /* The domain name this certificate is tied to */
    String domainName
    /* Wether or not this certificate is a wildcard cert */
    Boolean wildcard = false
}
