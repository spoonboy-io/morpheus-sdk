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
import org.openapitools.client.model.InlineResponse20082LoadBalancerInstanceSslCert;

/**
 * SecurityGroupLocations
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class SecurityGroupLocations {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_EXTERNAL_ID = "externalId";
  @SerializedName(SERIALIZED_NAME_EXTERNAL_ID)
  private String externalId;

  public static final String SERIALIZED_NAME_IAC_ID = "iacId";
  @SerializedName(SERIALIZED_NAME_IAC_ID)
  private String iacId;

  public static final String SERIALIZED_NAME_ZONE = "zone";
  @SerializedName(SERIALIZED_NAME_ZONE)
  private InlineResponse20082LoadBalancerInstanceSslCert zone;

  public static final String SERIALIZED_NAME_ZONE_POOL = "zonePool";
  @SerializedName(SERIALIZED_NAME_ZONE_POOL)
  private InlineResponse20082LoadBalancerInstanceSslCert zonePool;

  public static final String SERIALIZED_NAME_STATUS = "status";
  @SerializedName(SERIALIZED_NAME_STATUS)
  private String status;

  public static final String SERIALIZED_NAME_PRIORITY = "priority";
  @SerializedName(SERIALIZED_NAME_PRIORITY)
  private String priority;

  public static final String SERIALIZED_NAME_GROUP_LAYER = "groupLayer";
  @SerializedName(SERIALIZED_NAME_GROUP_LAYER)
  private String groupLayer;


  public SecurityGroupLocations id(Long id) {
    
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


  public SecurityGroupLocations name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Get name
   * @return name
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public SecurityGroupLocations description(String description) {
    
    this.description = description;
    return this;
  }

   /**
   * Get description
   * @return description
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDescription() {
    return description;
  }


  public void setDescription(String description) {
    this.description = description;
  }


  public SecurityGroupLocations externalId(String externalId) {
    
    this.externalId = externalId;
    return this;
  }

   /**
   * Get externalId
   * @return externalId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getExternalId() {
    return externalId;
  }


  public void setExternalId(String externalId) {
    this.externalId = externalId;
  }


  public SecurityGroupLocations iacId(String iacId) {
    
    this.iacId = iacId;
    return this;
  }

   /**
   * Get iacId
   * @return iacId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getIacId() {
    return iacId;
  }


  public void setIacId(String iacId) {
    this.iacId = iacId;
  }


  public SecurityGroupLocations zone(InlineResponse20082LoadBalancerInstanceSslCert zone) {
    
    this.zone = zone;
    return this;
  }

   /**
   * Get zone
   * @return zone
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse20082LoadBalancerInstanceSslCert getZone() {
    return zone;
  }


  public void setZone(InlineResponse20082LoadBalancerInstanceSslCert zone) {
    this.zone = zone;
  }


  public SecurityGroupLocations zonePool(InlineResponse20082LoadBalancerInstanceSslCert zonePool) {
    
    this.zonePool = zonePool;
    return this;
  }

   /**
   * Get zonePool
   * @return zonePool
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse20082LoadBalancerInstanceSslCert getZonePool() {
    return zonePool;
  }


  public void setZonePool(InlineResponse20082LoadBalancerInstanceSslCert zonePool) {
    this.zonePool = zonePool;
  }


  public SecurityGroupLocations status(String status) {
    
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


  public SecurityGroupLocations priority(String priority) {
    
    this.priority = priority;
    return this;
  }

   /**
   * Get priority
   * @return priority
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPriority() {
    return priority;
  }


  public void setPriority(String priority) {
    this.priority = priority;
  }


  public SecurityGroupLocations groupLayer(String groupLayer) {
    
    this.groupLayer = groupLayer;
    return this;
  }

   /**
   * Get groupLayer
   * @return groupLayer
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getGroupLayer() {
    return groupLayer;
  }


  public void setGroupLayer(String groupLayer) {
    this.groupLayer = groupLayer;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SecurityGroupLocations securityGroupLocations = (SecurityGroupLocations) o;
    return Objects.equals(this.id, securityGroupLocations.id) &&
        Objects.equals(this.name, securityGroupLocations.name) &&
        Objects.equals(this.description, securityGroupLocations.description) &&
        Objects.equals(this.externalId, securityGroupLocations.externalId) &&
        Objects.equals(this.iacId, securityGroupLocations.iacId) &&
        Objects.equals(this.zone, securityGroupLocations.zone) &&
        Objects.equals(this.zonePool, securityGroupLocations.zonePool) &&
        Objects.equals(this.status, securityGroupLocations.status) &&
        Objects.equals(this.priority, securityGroupLocations.priority) &&
        Objects.equals(this.groupLayer, securityGroupLocations.groupLayer);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, name, description, externalId, iacId, zone, zonePool, status, priority, groupLayer);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SecurityGroupLocations {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    externalId: ").append(toIndentedString(externalId)).append("\n");
    sb.append("    iacId: ").append(toIndentedString(iacId)).append("\n");
    sb.append("    zone: ").append(toIndentedString(zone)).append("\n");
    sb.append("    zonePool: ").append(toIndentedString(zonePool)).append("\n");
    sb.append("    status: ").append(toIndentedString(status)).append("\n");
    sb.append("    priority: ").append(toIndentedString(priority)).append("\n");
    sb.append("    groupLayer: ").append(toIndentedString(groupLayer)).append("\n");
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
