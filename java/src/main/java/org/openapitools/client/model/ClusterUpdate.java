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
 * ClusterUpdate
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ClusterUpdate {
  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_LABELS = "labels";
  @SerializedName(SERIALIZED_NAME_LABELS)
  private List<String> labels = null;

  public static final String SERIALIZED_NAME_ENABLED = "enabled";
  @SerializedName(SERIALIZED_NAME_ENABLED)
  private Boolean enabled;

  public static final String SERIALIZED_NAME_SERVICE_URL = "serviceUrl";
  @SerializedName(SERIALIZED_NAME_SERVICE_URL)
  private String serviceUrl;

  public static final String SERIALIZED_NAME_SERVICE_TOKEN = "serviceToken";
  @SerializedName(SERIALIZED_NAME_SERVICE_TOKEN)
  private String serviceToken;

  public static final String SERIALIZED_NAME_REFRESH = "refresh";
  @SerializedName(SERIALIZED_NAME_REFRESH)
  private Boolean refresh;

  public static final String SERIALIZED_NAME_MANAGED = "managed";
  @SerializedName(SERIALIZED_NAME_MANAGED)
  private Boolean managed;


  public ClusterUpdate name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Cluster name
   * @return name
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Cluster name")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public ClusterUpdate description(String description) {
    
    this.description = description;
    return this;
  }

   /**
   * Cluster description
   * @return description
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Cluster description")

  public String getDescription() {
    return description;
  }


  public void setDescription(String description) {
    this.description = description;
  }


  public ClusterUpdate labels(List<String> labels) {
    
    this.labels = labels;
    return this;
  }

  public ClusterUpdate addLabelsItem(String labelsItem) {
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


  public ClusterUpdate enabled(Boolean enabled) {
    
    this.enabled = enabled;
    return this;
  }

   /**
   * Cluster enabled
   * @return enabled
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Cluster enabled")

  public Boolean getEnabled() {
    return enabled;
  }


  public void setEnabled(Boolean enabled) {
    this.enabled = enabled;
  }


  public ClusterUpdate serviceUrl(String serviceUrl) {
    
    this.serviceUrl = serviceUrl;
    return this;
  }

   /**
   * Cluster API Url
   * @return serviceUrl
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Cluster API Url")

  public String getServiceUrl() {
    return serviceUrl;
  }


  public void setServiceUrl(String serviceUrl) {
    this.serviceUrl = serviceUrl;
  }


  public ClusterUpdate serviceToken(String serviceToken) {
    
    this.serviceToken = serviceToken;
    return this;
  }

   /**
   * Cluster API token
   * @return serviceToken
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Cluster API token")

  public String getServiceToken() {
    return serviceToken;
  }


  public void setServiceToken(String serviceToken) {
    this.serviceToken = serviceToken;
  }


  public ClusterUpdate refresh(Boolean refresh) {
    
    this.refresh = refresh;
    return this;
  }

   /**
   * Queue cluster refresh
   * @return refresh
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Queue cluster refresh")

  public Boolean getRefresh() {
    return refresh;
  }


  public void setRefresh(Boolean refresh) {
    this.refresh = refresh;
  }


  public ClusterUpdate managed(Boolean managed) {
    
    this.managed = managed;
    return this;
  }

   /**
   * Cluster managed
   * @return managed
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Cluster managed")

  public Boolean getManaged() {
    return managed;
  }


  public void setManaged(Boolean managed) {
    this.managed = managed;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ClusterUpdate clusterUpdate = (ClusterUpdate) o;
    return Objects.equals(this.name, clusterUpdate.name) &&
        Objects.equals(this.description, clusterUpdate.description) &&
        Objects.equals(this.labels, clusterUpdate.labels) &&
        Objects.equals(this.enabled, clusterUpdate.enabled) &&
        Objects.equals(this.serviceUrl, clusterUpdate.serviceUrl) &&
        Objects.equals(this.serviceToken, clusterUpdate.serviceToken) &&
        Objects.equals(this.refresh, clusterUpdate.refresh) &&
        Objects.equals(this.managed, clusterUpdate.managed);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, description, labels, enabled, serviceUrl, serviceToken, refresh, managed);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ClusterUpdate {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    labels: ").append(toIndentedString(labels)).append("\n");
    sb.append("    enabled: ").append(toIndentedString(enabled)).append("\n");
    sb.append("    serviceUrl: ").append(toIndentedString(serviceUrl)).append("\n");
    sb.append("    serviceToken: ").append(toIndentedString(serviceToken)).append("\n");
    sb.append("    refresh: ").append(toIndentedString(refresh)).append("\n");
    sb.append("    managed: ").append(toIndentedString(managed)).append("\n");
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
