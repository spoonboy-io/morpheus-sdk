package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class UserSourceCreateSaml {
    /* Login Redirect URL */
    String url
    /* Exclude SAML Request Parameter */
    Boolean doNotIncludeSAMLRequest = false
    /* Lougout Post URL */
    String logoutUrl
    /* SAML Request signing mode */
    String saMLSignatureMode = SaMLSignatureModeEnum.NOSIGNATURE
    /* Only applies when `SAMLSignatureMode=CustomSignature` */
    String request509Certificate
    /* RSA Private Key. Only applies when `SAMLSignatureMode=CustomSignature` */
    String requestPrivateKey
    /* SAML Response Signing Flag */
    Boolean doNotValidateSignature = true
    /* Signing Public Key. Only applies when `doNotValidateSignature=true` */
    String publicKey
    /* Encryption RSA Private Key */
    String privateKey
    /* Given Name Attribute Name */
    String givenNameAttribute
    /* Surname Attribute Name */
    String surnameAttribute
    /* Role Attribute Name */
    String roleAttributeName
    /* Role Attibute Required Value */
    String requiredAttributeValue
}
