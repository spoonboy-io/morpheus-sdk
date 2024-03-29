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
import java.math.BigDecimal;
import org.openapitools.client.model.BillingComputeServers;
import org.openapitools.client.model.BillingInstances;
import org.openapitools.client.model.BillingLoadBalancers;
import org.openapitools.client.model.BillingSnapshots;
import org.openapitools.client.model.BillingVirtualImages;
import org.threeten.bp.OffsetDateTime;

/**
 * BillingZone
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class BillingZone {
  public static final String SERIALIZED_NAME_ZONE_NAME = "zoneName";
  @SerializedName(SERIALIZED_NAME_ZONE_NAME)
  private String zoneName;

  public static final String SERIALIZED_NAME_ZONE_ID = "zoneId";
  @SerializedName(SERIALIZED_NAME_ZONE_ID)
  private Long zoneId;

  public static final String SERIALIZED_NAME_ZONE_U_U_I_D = "zoneUUID";
  @SerializedName(SERIALIZED_NAME_ZONE_U_U_I_D)
  private String zoneUUID;

  public static final String SERIALIZED_NAME_ZONE_CODE = "zoneCode";
  @SerializedName(SERIALIZED_NAME_ZONE_CODE)
  private String zoneCode;

  public static final String SERIALIZED_NAME_START_DATE = "startDate";
  @SerializedName(SERIALIZED_NAME_START_DATE)
  private OffsetDateTime startDate;

  public static final String SERIALIZED_NAME_END_DATE = "endDate";
  @SerializedName(SERIALIZED_NAME_END_DATE)
  private OffsetDateTime endDate;

  public static final String SERIALIZED_NAME_PRICE_UNIT = "priceUnit";
  @SerializedName(SERIALIZED_NAME_PRICE_UNIT)
  private String priceUnit;

  public static final String SERIALIZED_NAME_COMPUTE_SERVERS = "computeServers";
  @SerializedName(SERIALIZED_NAME_COMPUTE_SERVERS)
  private BillingComputeServers computeServers;

  public static final String SERIALIZED_NAME_INSTANCES = "instances";
  @SerializedName(SERIALIZED_NAME_INSTANCES)
  private BillingInstances instances;

  public static final String SERIALIZED_NAME_DISCOVERED_SERVERS = "discoveredServers";
  @SerializedName(SERIALIZED_NAME_DISCOVERED_SERVERS)
  private BillingComputeServers discoveredServers;

  public static final String SERIALIZED_NAME_LOAD_BALANCERS = "loadBalancers";
  @SerializedName(SERIALIZED_NAME_LOAD_BALANCERS)
  private BillingLoadBalancers loadBalancers;

  public static final String SERIALIZED_NAME_VIRTUAL_IMAGES = "virtualImages";
  @SerializedName(SERIALIZED_NAME_VIRTUAL_IMAGES)
  private BillingVirtualImages virtualImages;

  public static final String SERIALIZED_NAME_SNAPSHOTS = "snapshots";
  @SerializedName(SERIALIZED_NAME_SNAPSHOTS)
  private BillingSnapshots snapshots;

  public static final String SERIALIZED_NAME_PRICE = "price";
  @SerializedName(SERIALIZED_NAME_PRICE)
  private BigDecimal price;

  public static final String SERIALIZED_NAME_COST = "cost";
  @SerializedName(SERIALIZED_NAME_COST)
  private BigDecimal cost;


  public BillingZone zoneName(String zoneName) {
    
    this.zoneName = zoneName;
    return this;
  }

   /**
   * Get zoneName
   * @return zoneName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getZoneName() {
    return zoneName;
  }


  public void setZoneName(String zoneName) {
    this.zoneName = zoneName;
  }


  public BillingZone zoneId(Long zoneId) {
    
    this.zoneId = zoneId;
    return this;
  }

   /**
   * Get zoneId
   * @return zoneId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getZoneId() {
    return zoneId;
  }


  public void setZoneId(Long zoneId) {
    this.zoneId = zoneId;
  }


  public BillingZone zoneUUID(String zoneUUID) {
    
    this.zoneUUID = zoneUUID;
    return this;
  }

   /**
   * Get zoneUUID
   * @return zoneUUID
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getZoneUUID() {
    return zoneUUID;
  }


  public void setZoneUUID(String zoneUUID) {
    this.zoneUUID = zoneUUID;
  }


  public BillingZone zoneCode(String zoneCode) {
    
    this.zoneCode = zoneCode;
    return this;
  }

   /**
   * Get zoneCode
   * @return zoneCode
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getZoneCode() {
    return zoneCode;
  }


  public void setZoneCode(String zoneCode) {
    this.zoneCode = zoneCode;
  }


  public BillingZone startDate(OffsetDateTime startDate) {
    
    this.startDate = startDate;
    return this;
  }

   /**
   * Get startDate
   * @return startDate
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OffsetDateTime getStartDate() {
    return startDate;
  }


  public void setStartDate(OffsetDateTime startDate) {
    this.startDate = startDate;
  }


  public BillingZone endDate(OffsetDateTime endDate) {
    
    this.endDate = endDate;
    return this;
  }

   /**
   * Get endDate
   * @return endDate
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OffsetDateTime getEndDate() {
    return endDate;
  }


  public void setEndDate(OffsetDateTime endDate) {
    this.endDate = endDate;
  }


  public BillingZone priceUnit(String priceUnit) {
    
    this.priceUnit = priceUnit;
    return this;
  }

   /**
   * Get priceUnit
   * @return priceUnit
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPriceUnit() {
    return priceUnit;
  }


  public void setPriceUnit(String priceUnit) {
    this.priceUnit = priceUnit;
  }


  public BillingZone computeServers(BillingComputeServers computeServers) {
    
    this.computeServers = computeServers;
    return this;
  }

   /**
   * Get computeServers
   * @return computeServers
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BillingComputeServers getComputeServers() {
    return computeServers;
  }


  public void setComputeServers(BillingComputeServers computeServers) {
    this.computeServers = computeServers;
  }


  public BillingZone instances(BillingInstances instances) {
    
    this.instances = instances;
    return this;
  }

   /**
   * Get instances
   * @return instances
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BillingInstances getInstances() {
    return instances;
  }


  public void setInstances(BillingInstances instances) {
    this.instances = instances;
  }


  public BillingZone discoveredServers(BillingComputeServers discoveredServers) {
    
    this.discoveredServers = discoveredServers;
    return this;
  }

   /**
   * Get discoveredServers
   * @return discoveredServers
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BillingComputeServers getDiscoveredServers() {
    return discoveredServers;
  }


  public void setDiscoveredServers(BillingComputeServers discoveredServers) {
    this.discoveredServers = discoveredServers;
  }


  public BillingZone loadBalancers(BillingLoadBalancers loadBalancers) {
    
    this.loadBalancers = loadBalancers;
    return this;
  }

   /**
   * Get loadBalancers
   * @return loadBalancers
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BillingLoadBalancers getLoadBalancers() {
    return loadBalancers;
  }


  public void setLoadBalancers(BillingLoadBalancers loadBalancers) {
    this.loadBalancers = loadBalancers;
  }


  public BillingZone virtualImages(BillingVirtualImages virtualImages) {
    
    this.virtualImages = virtualImages;
    return this;
  }

   /**
   * Get virtualImages
   * @return virtualImages
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BillingVirtualImages getVirtualImages() {
    return virtualImages;
  }


  public void setVirtualImages(BillingVirtualImages virtualImages) {
    this.virtualImages = virtualImages;
  }


  public BillingZone snapshots(BillingSnapshots snapshots) {
    
    this.snapshots = snapshots;
    return this;
  }

   /**
   * Get snapshots
   * @return snapshots
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BillingSnapshots getSnapshots() {
    return snapshots;
  }


  public void setSnapshots(BillingSnapshots snapshots) {
    this.snapshots = snapshots;
  }


  public BillingZone price(BigDecimal price) {
    
    this.price = price;
    return this;
  }

   /**
   * Get price
   * @return price
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getPrice() {
    return price;
  }


  public void setPrice(BigDecimal price) {
    this.price = price;
  }


  public BillingZone cost(BigDecimal cost) {
    
    this.cost = cost;
    return this;
  }

   /**
   * Get cost
   * @return cost
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getCost() {
    return cost;
  }


  public void setCost(BigDecimal cost) {
    this.cost = cost;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    BillingZone billingZone = (BillingZone) o;
    return Objects.equals(this.zoneName, billingZone.zoneName) &&
        Objects.equals(this.zoneId, billingZone.zoneId) &&
        Objects.equals(this.zoneUUID, billingZone.zoneUUID) &&
        Objects.equals(this.zoneCode, billingZone.zoneCode) &&
        Objects.equals(this.startDate, billingZone.startDate) &&
        Objects.equals(this.endDate, billingZone.endDate) &&
        Objects.equals(this.priceUnit, billingZone.priceUnit) &&
        Objects.equals(this.computeServers, billingZone.computeServers) &&
        Objects.equals(this.instances, billingZone.instances) &&
        Objects.equals(this.discoveredServers, billingZone.discoveredServers) &&
        Objects.equals(this.loadBalancers, billingZone.loadBalancers) &&
        Objects.equals(this.virtualImages, billingZone.virtualImages) &&
        Objects.equals(this.snapshots, billingZone.snapshots) &&
        Objects.equals(this.price, billingZone.price) &&
        Objects.equals(this.cost, billingZone.cost);
  }

  @Override
  public int hashCode() {
    return Objects.hash(zoneName, zoneId, zoneUUID, zoneCode, startDate, endDate, priceUnit, computeServers, instances, discoveredServers, loadBalancers, virtualImages, snapshots, price, cost);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class BillingZone {\n");
    sb.append("    zoneName: ").append(toIndentedString(zoneName)).append("\n");
    sb.append("    zoneId: ").append(toIndentedString(zoneId)).append("\n");
    sb.append("    zoneUUID: ").append(toIndentedString(zoneUUID)).append("\n");
    sb.append("    zoneCode: ").append(toIndentedString(zoneCode)).append("\n");
    sb.append("    startDate: ").append(toIndentedString(startDate)).append("\n");
    sb.append("    endDate: ").append(toIndentedString(endDate)).append("\n");
    sb.append("    priceUnit: ").append(toIndentedString(priceUnit)).append("\n");
    sb.append("    computeServers: ").append(toIndentedString(computeServers)).append("\n");
    sb.append("    instances: ").append(toIndentedString(instances)).append("\n");
    sb.append("    discoveredServers: ").append(toIndentedString(discoveredServers)).append("\n");
    sb.append("    loadBalancers: ").append(toIndentedString(loadBalancers)).append("\n");
    sb.append("    virtualImages: ").append(toIndentedString(virtualImages)).append("\n");
    sb.append("    snapshots: ").append(toIndentedString(snapshots)).append("\n");
    sb.append("    price: ").append(toIndentedString(price)).append("\n");
    sb.append("    cost: ").append(toIndentedString(cost)).append("\n");
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

