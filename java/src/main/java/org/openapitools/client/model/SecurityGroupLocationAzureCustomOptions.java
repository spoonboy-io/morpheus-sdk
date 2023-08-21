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
 * SecurityGroupLocationAzureCustomOptions
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class SecurityGroupLocationAzureCustomOptions {
  public static final String SERIALIZED_NAME_RESOURCE_GROUP = "resourceGroup";
  @SerializedName(SERIALIZED_NAME_RESOURCE_GROUP)
  private String resourceGroup;


  public SecurityGroupLocationAzureCustomOptions resourceGroup(String resourceGroup) {
    
    this.resourceGroup = resourceGroup;
    return this;
  }

   /**
   * External ID of Azure Resource Group
   * @return resourceGroup
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "demo-lab", value = "External ID of Azure Resource Group")

  public String getResourceGroup() {
    return resourceGroup;
  }


  public void setResourceGroup(String resourceGroup) {
    this.resourceGroup = resourceGroup;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SecurityGroupLocationAzureCustomOptions securityGroupLocationAzureCustomOptions = (SecurityGroupLocationAzureCustomOptions) o;
    return Objects.equals(this.resourceGroup, securityGroupLocationAzureCustomOptions.resourceGroup);
  }

  @Override
  public int hashCode() {
    return Objects.hash(resourceGroup);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SecurityGroupLocationAzureCustomOptions {\n");
    sb.append("    resourceGroup: ").append(toIndentedString(resourceGroup)).append("\n");
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

