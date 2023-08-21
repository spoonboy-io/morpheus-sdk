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
 * DeploymentVersionCreate
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class DeploymentVersionCreate {
  public static final String SERIALIZED_NAME_VERSION = "version";
  @SerializedName(SERIALIZED_NAME_VERSION)
  private String version;

  public static final String SERIALIZED_NAME_USER_VERSION = "userVersion";
  @SerializedName(SERIALIZED_NAME_USER_VERSION)
  private String userVersion;

  /**
   * Deploy Type, eg. file, git, fetch
   */
  @JsonAdapter(DeployTypeEnum.Adapter.class)
  public enum DeployTypeEnum {
    FILE("file"),
    
    GIT("git"),
    
    FETCH("fetch");

    private String value;

    DeployTypeEnum(String value) {
      this.value = value;
    }

    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static DeployTypeEnum fromValue(String value) {
      for (DeployTypeEnum b : DeployTypeEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }

    public static class Adapter extends TypeAdapter<DeployTypeEnum> {
      @Override
      public void write(final JsonWriter jsonWriter, final DeployTypeEnum enumeration) throws IOException {
        jsonWriter.value(enumeration.getValue());
      }

      @Override
      public DeployTypeEnum read(final JsonReader jsonReader) throws IOException {
        String value =  jsonReader.nextString();
        return DeployTypeEnum.fromValue(value);
      }
    }
  }

  public static final String SERIALIZED_NAME_DEPLOY_TYPE = "deployType";
  @SerializedName(SERIALIZED_NAME_DEPLOY_TYPE)
  private DeployTypeEnum deployType;

  public static final String SERIALIZED_NAME_GIT_URL = "gitUrl";
  @SerializedName(SERIALIZED_NAME_GIT_URL)
  private String gitUrl;

  public static final String SERIALIZED_NAME_GIT_REF = "gitRef";
  @SerializedName(SERIALIZED_NAME_GIT_REF)
  private String gitRef;

  public static final String SERIALIZED_NAME_FETCH_URL = "fetchUrl";
  @SerializedName(SERIALIZED_NAME_FETCH_URL)
  private String fetchUrl;


  public DeploymentVersionCreate version(String version) {
    
    this.version = version;
    return this;
  }

   /**
   * Version number (userVersion), a unique version identifier for the deployment version.
   * @return version
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Version number (userVersion), a unique version identifier for the deployment version.")

  public String getVersion() {
    return version;
  }


  public void setVersion(String version) {
    this.version = version;
  }


  public DeploymentVersionCreate userVersion(String userVersion) {
    
    this.userVersion = userVersion;
    return this;
  }

   /**
   * Alias for version
   * @return userVersion
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Alias for version")

  public String getUserVersion() {
    return userVersion;
  }


  public void setUserVersion(String userVersion) {
    this.userVersion = userVersion;
  }


  public DeploymentVersionCreate deployType(DeployTypeEnum deployType) {
    
    this.deployType = deployType;
    return this;
  }

   /**
   * Deploy Type, eg. file, git, fetch
   * @return deployType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Deploy Type, eg. file, git, fetch")

  public DeployTypeEnum getDeployType() {
    return deployType;
  }


  public void setDeployType(DeployTypeEnum deployType) {
    this.deployType = deployType;
  }


  public DeploymentVersionCreate gitUrl(String gitUrl) {
    
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


  public DeploymentVersionCreate gitRef(String gitRef) {
    
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


  public DeploymentVersionCreate fetchUrl(String fetchUrl) {
    
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


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    DeploymentVersionCreate deploymentVersionCreate = (DeploymentVersionCreate) o;
    return Objects.equals(this.version, deploymentVersionCreate.version) &&
        Objects.equals(this.userVersion, deploymentVersionCreate.userVersion) &&
        Objects.equals(this.deployType, deploymentVersionCreate.deployType) &&
        Objects.equals(this.gitUrl, deploymentVersionCreate.gitUrl) &&
        Objects.equals(this.gitRef, deploymentVersionCreate.gitRef) &&
        Objects.equals(this.fetchUrl, deploymentVersionCreate.fetchUrl);
  }

  @Override
  public int hashCode() {
    return Objects.hash(version, userVersion, deployType, gitUrl, gitRef, fetchUrl);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class DeploymentVersionCreate {\n");
    sb.append("    version: ").append(toIndentedString(version)).append("\n");
    sb.append("    userVersion: ").append(toIndentedString(userVersion)).append("\n");
    sb.append("    deployType: ").append(toIndentedString(deployType)).append("\n");
    sb.append("    gitUrl: ").append(toIndentedString(gitUrl)).append("\n");
    sb.append("    gitRef: ").append(toIndentedString(gitRef)).append("\n");
    sb.append("    fetchUrl: ").append(toIndentedString(fetchUrl)).append("\n");
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
