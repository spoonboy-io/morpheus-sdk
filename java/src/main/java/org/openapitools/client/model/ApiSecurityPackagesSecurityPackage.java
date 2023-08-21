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
import java.util.ArrayList;
import java.util.List;

/**
 * ApiSecurityPackagesSecurityPackage
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ApiSecurityPackagesSecurityPackage {
  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_LABELS = "labels";
  @SerializedName(SERIALIZED_NAME_LABELS)
  private List<String> labels = null;

  public static final String SERIALIZED_NAME_TYPE = "type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private String type = "scap-package";

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_URL = "url";
  @SerializedName(SERIALIZED_NAME_URL)
  private String url;

  public static final String SERIALIZED_NAME_ENABLED = "enabled";
  @SerializedName(SERIALIZED_NAME_ENABLED)
  private Boolean enabled = true;


  public ApiSecurityPackagesSecurityPackage name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * A name for the security package
   * @return name
  **/
  @ApiModelProperty(example = "Sample Security Package", required = true, value = "A name for the security package")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public ApiSecurityPackagesSecurityPackage labels(List<String> labels) {
    
    this.labels = labels;
    return this;
  }

  public ApiSecurityPackagesSecurityPackage addLabelsItem(String labelsItem) {
    if (this.labels == null) {
      this.labels = new ArrayList<String>();
    }
    this.labels.add(labelsItem);
    return this;
  }

   /**
   * Array of label strings, can be used for filtering.
   * @return labels
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Array of label strings, can be used for filtering.")

  public List<String> getLabels() {
    return labels;
  }


  public void setLabels(List<String> labels) {
    this.labels = labels;
  }


  public ApiSecurityPackagesSecurityPackage type(String type) {
    
    this.type = type;
    return this;
  }

   /**
   * Security Package Type Code
   * @return type
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Security Package Type Code")

  public String getType() {
    return type;
  }


  public void setType(String type) {
    this.type = type;
  }


  public ApiSecurityPackagesSecurityPackage description(String description) {
    
    this.description = description;
    return this;
  }

   /**
   * A description for the security package
   * @return description
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "A description for the security package")

  public String getDescription() {
    return description;
  }


  public void setDescription(String description) {
    this.description = description;
  }


  public ApiSecurityPackagesSecurityPackage url(String url) {
    
    this.url = url;
    return this;
  }

   /**
   * URL to download the security package zip file from
   * @return url
  **/
  @ApiModelProperty(example = "http://10.0.2.2:8080/public-archives/download/SCAP/scap-security-guide-0.1.51.zip", required = true, value = "URL to download the security package zip file from")

  public String getUrl() {
    return url;
  }


  public void setUrl(String url) {
    this.url = url;
  }


  public ApiSecurityPackagesSecurityPackage enabled(Boolean enabled) {
    
    this.enabled = enabled;
    return this;
  }

   /**
   * Can be used to disable the security package
   * @return enabled
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Can be used to disable the security package")

  public Boolean getEnabled() {
    return enabled;
  }


  public void setEnabled(Boolean enabled) {
    this.enabled = enabled;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ApiSecurityPackagesSecurityPackage apiSecurityPackagesSecurityPackage = (ApiSecurityPackagesSecurityPackage) o;
    return Objects.equals(this.name, apiSecurityPackagesSecurityPackage.name) &&
        Objects.equals(this.labels, apiSecurityPackagesSecurityPackage.labels) &&
        Objects.equals(this.type, apiSecurityPackagesSecurityPackage.type) &&
        Objects.equals(this.description, apiSecurityPackagesSecurityPackage.description) &&
        Objects.equals(this.url, apiSecurityPackagesSecurityPackage.url) &&
        Objects.equals(this.enabled, apiSecurityPackagesSecurityPackage.enabled);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, labels, type, description, url, enabled);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ApiSecurityPackagesSecurityPackage {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    labels: ").append(toIndentedString(labels)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    url: ").append(toIndentedString(url)).append("\n");
    sb.append("    enabled: ").append(toIndentedString(enabled)).append("\n");
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

