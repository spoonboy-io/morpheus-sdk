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
 * InlineResponse200108NetworkFloatingIpCreatedBy
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InlineResponse200108NetworkFloatingIpCreatedBy {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_USERNAME = "username";
  @SerializedName(SERIALIZED_NAME_USERNAME)
  private String username;


  public InlineResponse200108NetworkFloatingIpCreatedBy id(Long id) {
    
    this.id = id;
    return this;
  }

   /**
   * User ID
   * @return id
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "User ID")

  public Long getId() {
    return id;
  }


  public void setId(Long id) {
    this.id = id;
  }


  public InlineResponse200108NetworkFloatingIpCreatedBy username(String username) {
    
    this.username = username;
    return this;
  }

   /**
   * Username
   * @return username
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Username")

  public String getUsername() {
    return username;
  }


  public void setUsername(String username) {
    this.username = username;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InlineResponse200108NetworkFloatingIpCreatedBy inlineResponse200108NetworkFloatingIpCreatedBy = (InlineResponse200108NetworkFloatingIpCreatedBy) o;
    return Objects.equals(this.id, inlineResponse200108NetworkFloatingIpCreatedBy.id) &&
        Objects.equals(this.username, inlineResponse200108NetworkFloatingIpCreatedBy.username);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, username);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InlineResponse200108NetworkFloatingIpCreatedBy {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    username: ").append(toIndentedString(username)).append("\n");
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

