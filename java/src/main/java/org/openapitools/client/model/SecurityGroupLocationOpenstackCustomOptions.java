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
 * SecurityGroupLocationOpenstackCustomOptions
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class SecurityGroupLocationOpenstackCustomOptions {
  public static final String SERIALIZED_NAME_RESOURCE_POOL_ID = "resourcePoolId";
  @SerializedName(SERIALIZED_NAME_RESOURCE_POOL_ID)
  private Long resourcePoolId;


  public SecurityGroupLocationOpenstackCustomOptions resourcePoolId(Long resourcePoolId) {
    
    this.resourcePoolId = resourcePoolId;
    return this;
  }

   /**
   * Resource Pool ID (applicable to cloud types Openstack/OpenTelekom/Huawei)
   * @return resourcePoolId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "3", value = "Resource Pool ID (applicable to cloud types Openstack/OpenTelekom/Huawei)")

  public Long getResourcePoolId() {
    return resourcePoolId;
  }


  public void setResourcePoolId(Long resourcePoolId) {
    this.resourcePoolId = resourcePoolId;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SecurityGroupLocationOpenstackCustomOptions securityGroupLocationOpenstackCustomOptions = (SecurityGroupLocationOpenstackCustomOptions) o;
    return Objects.equals(this.resourcePoolId, securityGroupLocationOpenstackCustomOptions.resourcePoolId);
  }

  @Override
  public int hashCode() {
    return Objects.hash(resourcePoolId);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SecurityGroupLocationOpenstackCustomOptions {\n");
    sb.append("    resourcePoolId: ").append(toIndentedString(resourcePoolId)).append("\n");
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

