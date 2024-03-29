/*
 * Morpheus API
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.2.1
 * Contact: dev@morpheusdata.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;

/**
 * ApiCertificatesCertificate
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ApiCertificatesCertificate {
  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_CERT_FILE = "certFile";
  @SerializedName(SERIALIZED_NAME_CERT_FILE)
  private String certFile;

  public static final String SERIALIZED_NAME_KEY_FILE = "keyFile";
  @SerializedName(SERIALIZED_NAME_KEY_FILE)
  private String keyFile;

  public static final String SERIALIZED_NAME_DOMAIN_NAME = "domainName";
  @SerializedName(SERIALIZED_NAME_DOMAIN_NAME)
  private String domainName;

  public static final String SERIALIZED_NAME_WILDCARD = "wildcard";
  @SerializedName(SERIALIZED_NAME_WILDCARD)
  private Boolean wildcard = false;


  public ApiCertificatesCertificate name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * A unique name scoped to your account for the key
   * @return name
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "A unique name scoped to your account for the key")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public ApiCertificatesCertificate certFile(String certFile) {
    
    this.certFile = certFile;
    return this;
  }

   /**
   * The contents of the certificate file
   * @return certFile
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "The contents of the certificate file")

  public String getCertFile() {
    return certFile;
  }


  public void setCertFile(String certFile) {
    this.certFile = certFile;
  }


  public ApiCertificatesCertificate keyFile(String keyFile) {
    
    this.keyFile = keyFile;
    return this;
  }

   /**
   * The contents of the key file
   * @return keyFile
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "The contents of the key file")

  public String getKeyFile() {
    return keyFile;
  }


  public void setKeyFile(String keyFile) {
    this.keyFile = keyFile;
  }


  public ApiCertificatesCertificate domainName(String domainName) {
    
    this.domainName = domainName;
    return this;
  }

   /**
   * The domain name this certificate is tied to
   * @return domainName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "The domain name this certificate is tied to")

  public String getDomainName() {
    return domainName;
  }


  public void setDomainName(String domainName) {
    this.domainName = domainName;
  }


  public ApiCertificatesCertificate wildcard(Boolean wildcard) {
    
    this.wildcard = wildcard;
    return this;
  }

   /**
   * Wether or not this certificate is a wildcard cert
   * @return wildcard
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Wether or not this certificate is a wildcard cert")

  public Boolean getWildcard() {
    return wildcard;
  }


  public void setWildcard(Boolean wildcard) {
    this.wildcard = wildcard;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ApiCertificatesCertificate apiCertificatesCertificate = (ApiCertificatesCertificate) o;
    return Objects.equals(this.name, apiCertificatesCertificate.name) &&
        Objects.equals(this.certFile, apiCertificatesCertificate.certFile) &&
        Objects.equals(this.keyFile, apiCertificatesCertificate.keyFile) &&
        Objects.equals(this.domainName, apiCertificatesCertificate.domainName) &&
        Objects.equals(this.wildcard, apiCertificatesCertificate.wildcard);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, certFile, keyFile, domainName, wildcard);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ApiCertificatesCertificate {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    certFile: ").append(toIndentedString(certFile)).append("\n");
    sb.append("    keyFile: ").append(toIndentedString(keyFile)).append("\n");
    sb.append("    domainName: ").append(toIndentedString(domainName)).append("\n");
    sb.append("    wildcard: ").append(toIndentedString(wildcard)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

