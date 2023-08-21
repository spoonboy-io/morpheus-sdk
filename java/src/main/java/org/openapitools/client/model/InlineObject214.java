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
import org.openapitools.client.model.ApiSecurityGroupsIdSecurityGroup;

/**
 * InlineObject214
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InlineObject214 {
  public static final String SERIALIZED_NAME_SECURITY_GROUP = "securityGroup";
  @SerializedName(SERIALIZED_NAME_SECURITY_GROUP)
  private ApiSecurityGroupsIdSecurityGroup securityGroup;


  public InlineObject214 securityGroup(ApiSecurityGroupsIdSecurityGroup securityGroup) {
    
    this.securityGroup = securityGroup;
    return this;
  }

   /**
   * Get securityGroup
   * @return securityGroup
  **/
  @ApiModelProperty(required = true, value = "")

  public ApiSecurityGroupsIdSecurityGroup getSecurityGroup() {
    return securityGroup;
  }


  public void setSecurityGroup(ApiSecurityGroupsIdSecurityGroup securityGroup) {
    this.securityGroup = securityGroup;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InlineObject214 inlineObject214 = (InlineObject214) o;
    return Objects.equals(this.securityGroup, inlineObject214.securityGroup);
  }

  @Override
  public int hashCode() {
    return Objects.hash(securityGroup);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InlineObject214 {\n");
    sb.append("    securityGroup: ").append(toIndentedString(securityGroup)).append("\n");
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
