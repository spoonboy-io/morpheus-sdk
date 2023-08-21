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
 * InstanceConfigReplicationGroup
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InstanceConfigReplicationGroup {
  public static final String SERIALIZED_NAME_PROVIDER_METHOD = "providerMethod";
  @SerializedName(SERIALIZED_NAME_PROVIDER_METHOD)
  private String providerMethod;


  public InstanceConfigReplicationGroup providerMethod(String providerMethod) {
    
    this.providerMethod = providerMethod;
    return this;
  }

   /**
   * Get providerMethod
   * @return providerMethod
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getProviderMethod() {
    return providerMethod;
  }


  public void setProviderMethod(String providerMethod) {
    this.providerMethod = providerMethod;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InstanceConfigReplicationGroup instanceConfigReplicationGroup = (InstanceConfigReplicationGroup) o;
    return Objects.equals(this.providerMethod, instanceConfigReplicationGroup.providerMethod);
  }

  @Override
  public int hashCode() {
    return Objects.hash(providerMethod);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InstanceConfigReplicationGroup {\n");
    sb.append("    providerMethod: ").append(toIndentedString(providerMethod)).append("\n");
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
