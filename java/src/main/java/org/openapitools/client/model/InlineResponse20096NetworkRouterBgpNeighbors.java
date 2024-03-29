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
import org.openapitools.client.model.InlineResponse20096Config;
import org.threeten.bp.OffsetDateTime;

/**
 * InlineResponse20096NetworkRouterBgpNeighbors
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InlineResponse20096NetworkRouterBgpNeighbors {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_IP_ADDRESS = "ipAddress";
  @SerializedName(SERIALIZED_NAME_IP_ADDRESS)
  private String ipAddress;

  public static final String SERIALIZED_NAME_FORWARDING_ADDRESS = "forwardingAddress";
  @SerializedName(SERIALIZED_NAME_FORWARDING_ADDRESS)
  private String forwardingAddress;

  public static final String SERIALIZED_NAME_PROTOCOL_ADDRESS = "protocolAddress";
  @SerializedName(SERIALIZED_NAME_PROTOCOL_ADDRESS)
  private String protocolAddress;

  public static final String SERIALIZED_NAME_REMOTE_AS = "remoteAs";
  @SerializedName(SERIALIZED_NAME_REMOTE_AS)
  private String remoteAs;

  public static final String SERIALIZED_NAME_WEIGHT = "weight";
  @SerializedName(SERIALIZED_NAME_WEIGHT)
  private Long weight;

  public static final String SERIALIZED_NAME_KEEP_ALIVE = "keepAlive";
  @SerializedName(SERIALIZED_NAME_KEEP_ALIVE)
  private Long keepAlive;

  public static final String SERIALIZED_NAME_HOLD_DOWN = "holdDown";
  @SerializedName(SERIALIZED_NAME_HOLD_DOWN)
  private Long holdDown;

  public static final String SERIALIZED_NAME_PASSWORD = "password";
  @SerializedName(SERIALIZED_NAME_PASSWORD)
  private String password;

  public static final String SERIALIZED_NAME_ROUTE_FILTERING_TYPE = "routeFilteringType";
  @SerializedName(SERIALIZED_NAME_ROUTE_FILTERING_TYPE)
  private String routeFilteringType;

  public static final String SERIALIZED_NAME_ROUTE_FILTERING_IN = "routeFilteringIn";
  @SerializedName(SERIALIZED_NAME_ROUTE_FILTERING_IN)
  private String routeFilteringIn;

  public static final String SERIALIZED_NAME_ROUTE_FILTERING_OUT = "routeFilteringOut";
  @SerializedName(SERIALIZED_NAME_ROUTE_FILTERING_OUT)
  private String routeFilteringOut;

  public static final String SERIALIZED_NAME_BFD_ENABLED = "bfdEnabled";
  @SerializedName(SERIALIZED_NAME_BFD_ENABLED)
  private Boolean bfdEnabled;

  public static final String SERIALIZED_NAME_BFD_INTERVAL = "bfdInterval";
  @SerializedName(SERIALIZED_NAME_BFD_INTERVAL)
  private Long bfdInterval;

  public static final String SERIALIZED_NAME_BFD_MULTIPLE = "bfdMultiple";
  @SerializedName(SERIALIZED_NAME_BFD_MULTIPLE)
  private Long bfdMultiple;

  public static final String SERIALIZED_NAME_ALLOW_AS_IN = "allowAsIn";
  @SerializedName(SERIALIZED_NAME_ALLOW_AS_IN)
  private Boolean allowAsIn;

  public static final String SERIALIZED_NAME_HOP_LIMIT = "hopLimit";
  @SerializedName(SERIALIZED_NAME_HOP_LIMIT)
  private Long hopLimit;

  public static final String SERIALIZED_NAME_RESTART_MODE = "restartMode";
  @SerializedName(SERIALIZED_NAME_RESTART_MODE)
  private String restartMode;

  public static final String SERIALIZED_NAME_PROVIDER_ID = "providerId";
  @SerializedName(SERIALIZED_NAME_PROVIDER_ID)
  private String providerId;

  public static final String SERIALIZED_NAME_SYNC_SOURCE = "syncSource";
  @SerializedName(SERIALIZED_NAME_SYNC_SOURCE)
  private String syncSource;

  public static final String SERIALIZED_NAME_INTERNAL_ID = "internalId";
  @SerializedName(SERIALIZED_NAME_INTERNAL_ID)
  private String internalId;

  public static final String SERIALIZED_NAME_EXTERNAL_ID = "externalId";
  @SerializedName(SERIALIZED_NAME_EXTERNAL_ID)
  private String externalId;

  public static final String SERIALIZED_NAME_REF_TYPE = "refType";
  @SerializedName(SERIALIZED_NAME_REF_TYPE)
  private String refType;

  public static final String SERIALIZED_NAME_REF_ID = "refId";
  @SerializedName(SERIALIZED_NAME_REF_ID)
  private String refId;

  public static final String SERIALIZED_NAME_CONFIG = "config";
  @SerializedName(SERIALIZED_NAME_CONFIG)
  private InlineResponse20096Config config;

  public static final String SERIALIZED_NAME_DATE_CREATED = "dateCreated";
  @SerializedName(SERIALIZED_NAME_DATE_CREATED)
  private OffsetDateTime dateCreated;

  public static final String SERIALIZED_NAME_LAST_UPDATED = "lastUpdated";
  @SerializedName(SERIALIZED_NAME_LAST_UPDATED)
  private OffsetDateTime lastUpdated;


  public InlineResponse20096NetworkRouterBgpNeighbors id(Long id) {
    
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


  public InlineResponse20096NetworkRouterBgpNeighbors ipAddress(String ipAddress) {
    
    this.ipAddress = ipAddress;
    return this;
  }

   /**
   * Get ipAddress
   * @return ipAddress
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getIpAddress() {
    return ipAddress;
  }


  public void setIpAddress(String ipAddress) {
    this.ipAddress = ipAddress;
  }


  public InlineResponse20096NetworkRouterBgpNeighbors forwardingAddress(String forwardingAddress) {
    
    this.forwardingAddress = forwardingAddress;
    return this;
  }

   /**
   * Get forwardingAddress
   * @return forwardingAddress
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getForwardingAddress() {
    return forwardingAddress;
  }


  public void setForwardingAddress(String forwardingAddress) {
    this.forwardingAddress = forwardingAddress;
  }


  public InlineResponse20096NetworkRouterBgpNeighbors protocolAddress(String protocolAddress) {
    
    this.protocolAddress = protocolAddress;
    return this;
  }

   /**
   * Get protocolAddress
   * @return protocolAddress
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getProtocolAddress() {
    return protocolAddress;
  }


  public void setProtocolAddress(String protocolAddress) {
    this.protocolAddress = protocolAddress;
  }


  public InlineResponse20096NetworkRouterBgpNeighbors remoteAs(String remoteAs) {
    
    this.remoteAs = remoteAs;
    return this;
  }

   /**
   * Get remoteAs
   * @return remoteAs
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRemoteAs() {
    return remoteAs;
  }


  public void setRemoteAs(String remoteAs) {
    this.remoteAs = remoteAs;
  }


  public InlineResponse20096NetworkRouterBgpNeighbors weight(Long weight) {
    
    this.weight = weight;
    return this;
  }

   /**
   * Get weight
   * @return weight
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getWeight() {
    return weight;
  }


  public void setWeight(Long weight) {
    this.weight = weight;
  }


  public InlineResponse20096NetworkRouterBgpNeighbors keepAlive(Long keepAlive) {
    
    this.keepAlive = keepAlive;
    return this;
  }

   /**
   * Get keepAlive
   * @return keepAlive
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getKeepAlive() {
    return keepAlive;
  }


  public void setKeepAlive(Long keepAlive) {
    this.keepAlive = keepAlive;
  }


  public InlineResponse20096NetworkRouterBgpNeighbors holdDown(Long holdDown) {
    
    this.holdDown = holdDown;
    return this;
  }

   /**
   * Get holdDown
   * @return holdDown
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getHoldDown() {
    return holdDown;
  }


  public void setHoldDown(Long holdDown) {
    this.holdDown = holdDown;
  }


  public InlineResponse20096NetworkRouterBgpNeighbors password(String password) {
    
    this.password = password;
    return this;
  }

   /**
   * Get password
   * @return password
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPassword() {
    return password;
  }


  public void setPassword(String password) {
    this.password = password;
  }


  public InlineResponse20096NetworkRouterBgpNeighbors routeFilteringType(String routeFilteringType) {
    
    this.routeFilteringType = routeFilteringType;
    return this;
  }

   /**
   * Get routeFilteringType
   * @return routeFilteringType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRouteFilteringType() {
    return routeFilteringType;
  }


  public void setRouteFilteringType(String routeFilteringType) {
    this.routeFilteringType = routeFilteringType;
  }


  public InlineResponse20096NetworkRouterBgpNeighbors routeFilteringIn(String routeFilteringIn) {
    
    this.routeFilteringIn = routeFilteringIn;
    return this;
  }

   /**
   * Get routeFilteringIn
   * @return routeFilteringIn
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRouteFilteringIn() {
    return routeFilteringIn;
  }


  public void setRouteFilteringIn(String routeFilteringIn) {
    this.routeFilteringIn = routeFilteringIn;
  }


  public InlineResponse20096NetworkRouterBgpNeighbors routeFilteringOut(String routeFilteringOut) {
    
    this.routeFilteringOut = routeFilteringOut;
    return this;
  }

   /**
   * Get routeFilteringOut
   * @return routeFilteringOut
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRouteFilteringOut() {
    return routeFilteringOut;
  }


  public void setRouteFilteringOut(String routeFilteringOut) {
    this.routeFilteringOut = routeFilteringOut;
  }


  public InlineResponse20096NetworkRouterBgpNeighbors bfdEnabled(Boolean bfdEnabled) {
    
    this.bfdEnabled = bfdEnabled;
    return this;
  }

   /**
   * Get bfdEnabled
   * @return bfdEnabled
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getBfdEnabled() {
    return bfdEnabled;
  }


  public void setBfdEnabled(Boolean bfdEnabled) {
    this.bfdEnabled = bfdEnabled;
  }


  public InlineResponse20096NetworkRouterBgpNeighbors bfdInterval(Long bfdInterval) {
    
    this.bfdInterval = bfdInterval;
    return this;
  }

   /**
   * Get bfdInterval
   * @return bfdInterval
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getBfdInterval() {
    return bfdInterval;
  }


  public void setBfdInterval(Long bfdInterval) {
    this.bfdInterval = bfdInterval;
  }


  public InlineResponse20096NetworkRouterBgpNeighbors bfdMultiple(Long bfdMultiple) {
    
    this.bfdMultiple = bfdMultiple;
    return this;
  }

   /**
   * Get bfdMultiple
   * @return bfdMultiple
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getBfdMultiple() {
    return bfdMultiple;
  }


  public void setBfdMultiple(Long bfdMultiple) {
    this.bfdMultiple = bfdMultiple;
  }


  public InlineResponse20096NetworkRouterBgpNeighbors allowAsIn(Boolean allowAsIn) {
    
    this.allowAsIn = allowAsIn;
    return this;
  }

   /**
   * Get allowAsIn
   * @return allowAsIn
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getAllowAsIn() {
    return allowAsIn;
  }


  public void setAllowAsIn(Boolean allowAsIn) {
    this.allowAsIn = allowAsIn;
  }


  public InlineResponse20096NetworkRouterBgpNeighbors hopLimit(Long hopLimit) {
    
    this.hopLimit = hopLimit;
    return this;
  }

   /**
   * Get hopLimit
   * @return hopLimit
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getHopLimit() {
    return hopLimit;
  }


  public void setHopLimit(Long hopLimit) {
    this.hopLimit = hopLimit;
  }


  public InlineResponse20096NetworkRouterBgpNeighbors restartMode(String restartMode) {
    
    this.restartMode = restartMode;
    return this;
  }

   /**
   * Get restartMode
   * @return restartMode
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRestartMode() {
    return restartMode;
  }


  public void setRestartMode(String restartMode) {
    this.restartMode = restartMode;
  }


  public InlineResponse20096NetworkRouterBgpNeighbors providerId(String providerId) {
    
    this.providerId = providerId;
    return this;
  }

   /**
   * Get providerId
   * @return providerId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getProviderId() {
    return providerId;
  }


  public void setProviderId(String providerId) {
    this.providerId = providerId;
  }


  public InlineResponse20096NetworkRouterBgpNeighbors syncSource(String syncSource) {
    
    this.syncSource = syncSource;
    return this;
  }

   /**
   * Get syncSource
   * @return syncSource
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSyncSource() {
    return syncSource;
  }


  public void setSyncSource(String syncSource) {
    this.syncSource = syncSource;
  }


  public InlineResponse20096NetworkRouterBgpNeighbors internalId(String internalId) {
    
    this.internalId = internalId;
    return this;
  }

   /**
   * Get internalId
   * @return internalId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getInternalId() {
    return internalId;
  }


  public void setInternalId(String internalId) {
    this.internalId = internalId;
  }


  public InlineResponse20096NetworkRouterBgpNeighbors externalId(String externalId) {
    
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


  public InlineResponse20096NetworkRouterBgpNeighbors refType(String refType) {
    
    this.refType = refType;
    return this;
  }

   /**
   * Get refType
   * @return refType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRefType() {
    return refType;
  }


  public void setRefType(String refType) {
    this.refType = refType;
  }


  public InlineResponse20096NetworkRouterBgpNeighbors refId(String refId) {
    
    this.refId = refId;
    return this;
  }

   /**
   * Get refId
   * @return refId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRefId() {
    return refId;
  }


  public void setRefId(String refId) {
    this.refId = refId;
  }


  public InlineResponse20096NetworkRouterBgpNeighbors config(InlineResponse20096Config config) {
    
    this.config = config;
    return this;
  }

   /**
   * Get config
   * @return config
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse20096Config getConfig() {
    return config;
  }


  public void setConfig(InlineResponse20096Config config) {
    this.config = config;
  }


  public InlineResponse20096NetworkRouterBgpNeighbors dateCreated(OffsetDateTime dateCreated) {
    
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


  public InlineResponse20096NetworkRouterBgpNeighbors lastUpdated(OffsetDateTime lastUpdated) {
    
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
    InlineResponse20096NetworkRouterBgpNeighbors inlineResponse20096NetworkRouterBgpNeighbors = (InlineResponse20096NetworkRouterBgpNeighbors) o;
    return Objects.equals(this.id, inlineResponse20096NetworkRouterBgpNeighbors.id) &&
        Objects.equals(this.ipAddress, inlineResponse20096NetworkRouterBgpNeighbors.ipAddress) &&
        Objects.equals(this.forwardingAddress, inlineResponse20096NetworkRouterBgpNeighbors.forwardingAddress) &&
        Objects.equals(this.protocolAddress, inlineResponse20096NetworkRouterBgpNeighbors.protocolAddress) &&
        Objects.equals(this.remoteAs, inlineResponse20096NetworkRouterBgpNeighbors.remoteAs) &&
        Objects.equals(this.weight, inlineResponse20096NetworkRouterBgpNeighbors.weight) &&
        Objects.equals(this.keepAlive, inlineResponse20096NetworkRouterBgpNeighbors.keepAlive) &&
        Objects.equals(this.holdDown, inlineResponse20096NetworkRouterBgpNeighbors.holdDown) &&
        Objects.equals(this.password, inlineResponse20096NetworkRouterBgpNeighbors.password) &&
        Objects.equals(this.routeFilteringType, inlineResponse20096NetworkRouterBgpNeighbors.routeFilteringType) &&
        Objects.equals(this.routeFilteringIn, inlineResponse20096NetworkRouterBgpNeighbors.routeFilteringIn) &&
        Objects.equals(this.routeFilteringOut, inlineResponse20096NetworkRouterBgpNeighbors.routeFilteringOut) &&
        Objects.equals(this.bfdEnabled, inlineResponse20096NetworkRouterBgpNeighbors.bfdEnabled) &&
        Objects.equals(this.bfdInterval, inlineResponse20096NetworkRouterBgpNeighbors.bfdInterval) &&
        Objects.equals(this.bfdMultiple, inlineResponse20096NetworkRouterBgpNeighbors.bfdMultiple) &&
        Objects.equals(this.allowAsIn, inlineResponse20096NetworkRouterBgpNeighbors.allowAsIn) &&
        Objects.equals(this.hopLimit, inlineResponse20096NetworkRouterBgpNeighbors.hopLimit) &&
        Objects.equals(this.restartMode, inlineResponse20096NetworkRouterBgpNeighbors.restartMode) &&
        Objects.equals(this.providerId, inlineResponse20096NetworkRouterBgpNeighbors.providerId) &&
        Objects.equals(this.syncSource, inlineResponse20096NetworkRouterBgpNeighbors.syncSource) &&
        Objects.equals(this.internalId, inlineResponse20096NetworkRouterBgpNeighbors.internalId) &&
        Objects.equals(this.externalId, inlineResponse20096NetworkRouterBgpNeighbors.externalId) &&
        Objects.equals(this.refType, inlineResponse20096NetworkRouterBgpNeighbors.refType) &&
        Objects.equals(this.refId, inlineResponse20096NetworkRouterBgpNeighbors.refId) &&
        Objects.equals(this.config, inlineResponse20096NetworkRouterBgpNeighbors.config) &&
        Objects.equals(this.dateCreated, inlineResponse20096NetworkRouterBgpNeighbors.dateCreated) &&
        Objects.equals(this.lastUpdated, inlineResponse20096NetworkRouterBgpNeighbors.lastUpdated);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, ipAddress, forwardingAddress, protocolAddress, remoteAs, weight, keepAlive, holdDown, password, routeFilteringType, routeFilteringIn, routeFilteringOut, bfdEnabled, bfdInterval, bfdMultiple, allowAsIn, hopLimit, restartMode, providerId, syncSource, internalId, externalId, refType, refId, config, dateCreated, lastUpdated);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InlineResponse20096NetworkRouterBgpNeighbors {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    ipAddress: ").append(toIndentedString(ipAddress)).append("\n");
    sb.append("    forwardingAddress: ").append(toIndentedString(forwardingAddress)).append("\n");
    sb.append("    protocolAddress: ").append(toIndentedString(protocolAddress)).append("\n");
    sb.append("    remoteAs: ").append(toIndentedString(remoteAs)).append("\n");
    sb.append("    weight: ").append(toIndentedString(weight)).append("\n");
    sb.append("    keepAlive: ").append(toIndentedString(keepAlive)).append("\n");
    sb.append("    holdDown: ").append(toIndentedString(holdDown)).append("\n");
    sb.append("    password: ").append(toIndentedString(password)).append("\n");
    sb.append("    routeFilteringType: ").append(toIndentedString(routeFilteringType)).append("\n");
    sb.append("    routeFilteringIn: ").append(toIndentedString(routeFilteringIn)).append("\n");
    sb.append("    routeFilteringOut: ").append(toIndentedString(routeFilteringOut)).append("\n");
    sb.append("    bfdEnabled: ").append(toIndentedString(bfdEnabled)).append("\n");
    sb.append("    bfdInterval: ").append(toIndentedString(bfdInterval)).append("\n");
    sb.append("    bfdMultiple: ").append(toIndentedString(bfdMultiple)).append("\n");
    sb.append("    allowAsIn: ").append(toIndentedString(allowAsIn)).append("\n");
    sb.append("    hopLimit: ").append(toIndentedString(hopLimit)).append("\n");
    sb.append("    restartMode: ").append(toIndentedString(restartMode)).append("\n");
    sb.append("    providerId: ").append(toIndentedString(providerId)).append("\n");
    sb.append("    syncSource: ").append(toIndentedString(syncSource)).append("\n");
    sb.append("    internalId: ").append(toIndentedString(internalId)).append("\n");
    sb.append("    externalId: ").append(toIndentedString(externalId)).append("\n");
    sb.append("    refType: ").append(toIndentedString(refType)).append("\n");
    sb.append("    refId: ").append(toIndentedString(refId)).append("\n");
    sb.append("    config: ").append(toIndentedString(config)).append("\n");
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

