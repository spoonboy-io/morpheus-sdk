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
 * BlueprintARMCreateArmGit
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class BlueprintARMCreateArmGit {
  public static final String SERIALIZED_NAME_REPO_ID = "repoId";
  @SerializedName(SERIALIZED_NAME_REPO_ID)
  private Long repoId;

  public static final String SERIALIZED_NAME_PATH = "path";
  @SerializedName(SERIALIZED_NAME_PATH)
  private String path;

  public static final String SERIALIZED_NAME_INTEGRATION_ID = "integrationId";
  @SerializedName(SERIALIZED_NAME_INTEGRATION_ID)
  private Long integrationId;

  public static final String SERIALIZED_NAME_BRANCH = "branch";
  @SerializedName(SERIALIZED_NAME_BRANCH)
  private String branch;


  public BlueprintARMCreateArmGit repoId(Long repoId) {
    
    this.repoId = repoId;
    return this;
  }

   /**
   * Morpheus SCM Repository ID
   * @return repoId
  **/
  @ApiModelProperty(required = true, value = "Morpheus SCM Repository ID")

  public Long getRepoId() {
    return repoId;
  }


  public void setRepoId(Long repoId) {
    this.repoId = repoId;
  }


  public BlueprintARMCreateArmGit path(String path) {
    
    this.path = path;
    return this;
  }

   /**
   * Path to ARM Files in the Repository
   * @return path
  **/
  @ApiModelProperty(required = true, value = "Path to ARM Files in the Repository")

  public String getPath() {
    return path;
  }


  public void setPath(String path) {
    this.path = path;
  }


  public BlueprintARMCreateArmGit integrationId(Long integrationId) {
    
    this.integrationId = integrationId;
    return this;
  }

   /**
   * Morpheus SCM Integration ID
   * @return integrationId
  **/
  @ApiModelProperty(required = true, value = "Morpheus SCM Integration ID")

  public Long getIntegrationId() {
    return integrationId;
  }


  public void setIntegrationId(Long integrationId) {
    this.integrationId = integrationId;
  }


  public BlueprintARMCreateArmGit branch(String branch) {
    
    this.branch = branch;
    return this;
  }

   /**
   * Branch Name
   * @return branch
  **/
  @ApiModelProperty(required = true, value = "Branch Name")

  public String getBranch() {
    return branch;
  }


  public void setBranch(String branch) {
    this.branch = branch;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    BlueprintARMCreateArmGit blueprintARMCreateArmGit = (BlueprintARMCreateArmGit) o;
    return Objects.equals(this.repoId, blueprintARMCreateArmGit.repoId) &&
        Objects.equals(this.path, blueprintARMCreateArmGit.path) &&
        Objects.equals(this.integrationId, blueprintARMCreateArmGit.integrationId) &&
        Objects.equals(this.branch, blueprintARMCreateArmGit.branch);
  }

  @Override
  public int hashCode() {
    return Objects.hash(repoId, path, integrationId, branch);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class BlueprintARMCreateArmGit {\n");
    sb.append("    repoId: ").append(toIndentedString(repoId)).append("\n");
    sb.append("    path: ").append(toIndentedString(path)).append("\n");
    sb.append("    integrationId: ").append(toIndentedString(integrationId)).append("\n");
    sb.append("    branch: ").append(toIndentedString(branch)).append("\n");
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

