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
 * InlineResponse20040AppDeployDeploymentVersion
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InlineResponse20040AppDeployDeploymentVersion {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_USER_VERSION = "userVersion";
  @SerializedName(SERIALIZED_NAME_USER_VERSION)
  private String userVersion;

  public static final String SERIALIZED_NAME_DEPLOY_TYPE = "deployType";
  @SerializedName(SERIALIZED_NAME_DEPLOY_TYPE)
  private String deployType;


  public InlineResponse20040AppDeployDeploymentVersion id(Long id) {
    
    this.id = id;
    return this;
  }

   /**
   * Get id
   * @return id
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getId() {
    return id;
  }


  public void setId(Long id) {
    this.id = id;
  }


  public InlineResponse20040AppDeployDeploymentVersion userVersion(String userVersion) {
    
    this.userVersion = userVersion;
    return this;
  }

   /**
   * Get userVersion
   * @return userVersion
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getUserVersion() {
    return userVersion;
  }


  public void setUserVersion(String userVersion) {
    this.userVersion = userVersion;
  }


  public InlineResponse20040AppDeployDeploymentVersion deployType(String deployType) {
    
    this.deployType = deployType;
    return this;
  }

   /**
   * Get deployType
   * @return deployType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDeployType() {
    return deployType;
  }


  public void setDeployType(String deployType) {
    this.deployType = deployType;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InlineResponse20040AppDeployDeploymentVersion inlineResponse20040AppDeployDeploymentVersion = (InlineResponse20040AppDeployDeploymentVersion) o;
    return Objects.equals(this.id, inlineResponse20040AppDeployDeploymentVersion.id) &&
        Objects.equals(this.userVersion, inlineResponse20040AppDeployDeploymentVersion.userVersion) &&
        Objects.equals(this.deployType, inlineResponse20040AppDeployDeploymentVersion.deployType);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, userVersion, deployType);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InlineResponse20040AppDeployDeploymentVersion {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    userVersion: ").append(toIndentedString(userVersion)).append("\n");
    sb.append("    deployType: ").append(toIndentedString(deployType)).append("\n");
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

