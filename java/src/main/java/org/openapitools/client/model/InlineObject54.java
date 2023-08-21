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
 * InlineObject54
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InlineObject54 {
  public static final String SERIALIZED_NAME_SERVICE_URL = "serviceUrl";
  @SerializedName(SERIALIZED_NAME_SERVICE_URL)
  private String serviceUrl;

  public static final String SERIALIZED_NAME_SPEC_TEMPLATE = "specTemplate";
  @SerializedName(SERIALIZED_NAME_SPEC_TEMPLATE)
  private String specTemplate;

  public static final String SERIALIZED_NAME_SPEC_YAML = "specYaml";
  @SerializedName(SERIALIZED_NAME_SPEC_YAML)
  private String specYaml;


  public InlineObject54 serviceUrl(String serviceUrl) {
    
    this.serviceUrl = serviceUrl;
    return this;
  }

   /**
   * Url of desired template to apply to cluster
   * @return serviceUrl
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Url of desired template to apply to cluster")

  public String getServiceUrl() {
    return serviceUrl;
  }


  public void setServiceUrl(String serviceUrl) {
    this.serviceUrl = serviceUrl;
  }


  public InlineObject54 specTemplate(String specTemplate) {
    
    this.specTemplate = specTemplate;
    return this;
  }

   /**
   * Name or ID of desired Spec Template to apply to cluster
   * @return specTemplate
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Name or ID of desired Spec Template to apply to cluster")

  public String getSpecTemplate() {
    return specTemplate;
  }


  public void setSpecTemplate(String specTemplate) {
    this.specTemplate = specTemplate;
  }


  public InlineObject54 specYaml(String specYaml) {
    
    this.specYaml = specYaml;
    return this;
  }

   /**
   * Yaml of template to apply to cluster
   * @return specYaml
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Yaml of template to apply to cluster")

  public String getSpecYaml() {
    return specYaml;
  }


  public void setSpecYaml(String specYaml) {
    this.specYaml = specYaml;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InlineObject54 inlineObject54 = (InlineObject54) o;
    return Objects.equals(this.serviceUrl, inlineObject54.serviceUrl) &&
        Objects.equals(this.specTemplate, inlineObject54.specTemplate) &&
        Objects.equals(this.specYaml, inlineObject54.specYaml);
  }

  @Override
  public int hashCode() {
    return Objects.hash(serviceUrl, specTemplate, specYaml);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InlineObject54 {\n");
    sb.append("    serviceUrl: ").append(toIndentedString(serviceUrl)).append("\n");
    sb.append("    specTemplate: ").append(toIndentedString(specTemplate)).append("\n");
    sb.append("    specYaml: ").append(toIndentedString(specYaml)).append("\n");
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

