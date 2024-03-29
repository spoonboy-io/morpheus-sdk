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
import org.threeten.bp.OffsetDateTime;

/**
 * DeploymentVersions
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class DeploymentVersions {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_DEPLOY_TYPE = "deployType";
  @SerializedName(SERIALIZED_NAME_DEPLOY_TYPE)
  private String deployType;

  public static final String SERIALIZED_NAME_FETCH_URL = "fetchUrl";
  @SerializedName(SERIALIZED_NAME_FETCH_URL)
  private String fetchUrl;

  public static final String SERIALIZED_NAME_GIT_URL = "gitUrl";
  @SerializedName(SERIALIZED_NAME_GIT_URL)
  private String gitUrl;

  public static final String SERIALIZED_NAME_GIT_REF = "gitRef";
  @SerializedName(SERIALIZED_NAME_GIT_REF)
  private String gitRef;

  public static final String SERIALIZED_NAME_USER_VERSION = "userVersion";
  @SerializedName(SERIALIZED_NAME_USER_VERSION)
  private String userVersion;

  public static final String SERIALIZED_NAME_VERSION = "version";
  @SerializedName(SERIALIZED_NAME_VERSION)
  private String version;

  public static final String SERIALIZED_NAME_STATUS = "status";
  @SerializedName(SERIALIZED_NAME_STATUS)
  private String status;

  public static final String SERIALIZED_NAME_DATE_CREATED = "dateCreated";
  @SerializedName(SERIALIZED_NAME_DATE_CREATED)
  private OffsetDateTime dateCreated;

  public static final String SERIALIZED_NAME_LAST_UPDATED = "lastUpdated";
  @SerializedName(SERIALIZED_NAME_LAST_UPDATED)
  private OffsetDateTime lastUpdated;


  public DeploymentVersions id(Long id) {
    
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


  public DeploymentVersions deployType(String deployType) {
    
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


  public DeploymentVersions fetchUrl(String fetchUrl) {
    
    this.fetchUrl = fetchUrl;
    return this;
  }

   /**
   * Get fetchUrl
   * @return fetchUrl
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getFetchUrl() {
    return fetchUrl;
  }


  public void setFetchUrl(String fetchUrl) {
    this.fetchUrl = fetchUrl;
  }


  public DeploymentVersions gitUrl(String gitUrl) {
    
    this.gitUrl = gitUrl;
    return this;
  }

   /**
   * Get gitUrl
   * @return gitUrl
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getGitUrl() {
    return gitUrl;
  }


  public void setGitUrl(String gitUrl) {
    this.gitUrl = gitUrl;
  }


  public DeploymentVersions gitRef(String gitRef) {
    
    this.gitRef = gitRef;
    return this;
  }

   /**
   * Get gitRef
   * @return gitRef
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getGitRef() {
    return gitRef;
  }


  public void setGitRef(String gitRef) {
    this.gitRef = gitRef;
  }


  public DeploymentVersions userVersion(String userVersion) {
    
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


  public DeploymentVersions version(String version) {
    
    this.version = version;
    return this;
  }

   /**
   * Get version
   * @return version
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getVersion() {
    return version;
  }


  public void setVersion(String version) {
    this.version = version;
  }


  public DeploymentVersions status(String status) {
    
    this.status = status;
    return this;
  }

   /**
   * Get status
   * @return status
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getStatus() {
    return status;
  }


  public void setStatus(String status) {
    this.status = status;
  }


  public DeploymentVersions dateCreated(OffsetDateTime dateCreated) {
    
    this.dateCreated = dateCreated;
    return this;
  }

   /**
   * Get dateCreated
   * @return dateCreated
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OffsetDateTime getDateCreated() {
    return dateCreated;
  }


  public void setDateCreated(OffsetDateTime dateCreated) {
    this.dateCreated = dateCreated;
  }


  public DeploymentVersions lastUpdated(OffsetDateTime lastUpdated) {
    
    this.lastUpdated = lastUpdated;
    return this;
  }

   /**
   * Get lastUpdated
   * @return lastUpdated
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OffsetDateTime getLastUpdated() {
    return lastUpdated;
  }


  public void setLastUpdated(OffsetDateTime lastUpdated) {
    this.lastUpdated = lastUpdated;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    DeploymentVersions deploymentVersions = (DeploymentVersions) o;
    return Objects.equals(this.id, deploymentVersions.id) &&
        Objects.equals(this.deployType, deploymentVersions.deployType) &&
        Objects.equals(this.fetchUrl, deploymentVersions.fetchUrl) &&
        Objects.equals(this.gitUrl, deploymentVersions.gitUrl) &&
        Objects.equals(this.gitRef, deploymentVersions.gitRef) &&
        Objects.equals(this.userVersion, deploymentVersions.userVersion) &&
        Objects.equals(this.version, deploymentVersions.version) &&
        Objects.equals(this.status, deploymentVersions.status) &&
        Objects.equals(this.dateCreated, deploymentVersions.dateCreated) &&
        Objects.equals(this.lastUpdated, deploymentVersions.lastUpdated);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, deployType, fetchUrl, gitUrl, gitRef, userVersion, version, status, dateCreated, lastUpdated);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class DeploymentVersions {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    deployType: ").append(toIndentedString(deployType)).append("\n");
    sb.append("    fetchUrl: ").append(toIndentedString(fetchUrl)).append("\n");
    sb.append("    gitUrl: ").append(toIndentedString(gitUrl)).append("\n");
    sb.append("    gitRef: ").append(toIndentedString(gitRef)).append("\n");
    sb.append("    userVersion: ").append(toIndentedString(userVersion)).append("\n");
    sb.append("    version: ").append(toIndentedString(version)).append("\n");
    sb.append("    status: ").append(toIndentedString(status)).append("\n");
    sb.append("    dateCreated: ").append(toIndentedString(dateCreated)).append("\n");
    sb.append("    lastUpdated: ").append(toIndentedString(lastUpdated)).append("\n");
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

