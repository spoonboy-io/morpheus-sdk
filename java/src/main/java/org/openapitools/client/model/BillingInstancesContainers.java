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
import java.util.ArrayList;
import java.util.List;
import org.openapitools.client.model.BillingInstancesUsages;
import org.threeten.bp.OffsetDateTime;

/**
 * BillingInstancesContainers
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class BillingInstancesContainers {
  public static final String SERIALIZED_NAME_REF_TYPE = "refType";
  @SerializedName(SERIALIZED_NAME_REF_TYPE)
  private String refType;

  public static final String SERIALIZED_NAME_REF_U_U_I_D = "refUUID";
  @SerializedName(SERIALIZED_NAME_REF_U_U_I_D)
  private String refUUID;

  public static final String SERIALIZED_NAME_REF_ID = "refId";
  @SerializedName(SERIALIZED_NAME_REF_ID)
  private Long refId;

  public static final String SERIALIZED_NAME_START_DATE = "startDate";
  @SerializedName(SERIALIZED_NAME_START_DATE)
  private OffsetDateTime startDate;

  public static final String SERIALIZED_NAME_END_DATE = "endDate";
  @SerializedName(SERIALIZED_NAME_END_DATE)
  private OffsetDateTime endDate;

  public static final String SERIALIZED_NAME_COST = "cost";
  @SerializedName(SERIALIZED_NAME_COST)
  private BigDecimal cost;

  public static final String SERIALIZED_NAME_PRICE = "price";
  @SerializedName(SERIALIZED_NAME_PRICE)
  private BigDecimal price;

  public static final String SERIALIZED_NAME_NUM_UNITS = "numUnits";
  @SerializedName(SERIALIZED_NAME_NUM_UNITS)
  private BigDecimal numUnits;

  public static final String SERIALIZED_NAME_UNIT = "unit";
  @SerializedName(SERIALIZED_NAME_UNIT)
  private String unit;

  public static final String SERIALIZED_NAME_CURRENCY = "currency";
  @SerializedName(SERIALIZED_NAME_CURRENCY)
  private String currency;

  public static final String SERIALIZED_NAME_USAGES = "usages";
  @SerializedName(SERIALIZED_NAME_USAGES)
  private List<BillingInstancesUsages> usages = null;

  public static final String SERIALIZED_NAME_NUM_USAGES = "numUsages";
  @SerializedName(SERIALIZED_NAME_NUM_USAGES)
  private Long numUsages;

  public static final String SERIALIZED_NAME_TOTAL_USAGES = "totalUsages";
  @SerializedName(SERIALIZED_NAME_TOTAL_USAGES)
  private Long totalUsages;

  public static final String SERIALIZED_NAME_HAS_MORE_USAGES = "hasMoreUsages";
  @SerializedName(SERIALIZED_NAME_HAS_MORE_USAGES)
  private Boolean hasMoreUsages;

  public static final String SERIALIZED_NAME_FOUND_PRICING = "foundPricing";
  @SerializedName(SERIALIZED_NAME_FOUND_PRICING)
  private Boolean foundPricing;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_SERVER_ID = "serverId";
  @SerializedName(SERIALIZED_NAME_SERVER_ID)
  private Long serverId;

  public static final String SERIALIZED_NAME_SERVER_U_U_I_D = "serverUUID";
  @SerializedName(SERIALIZED_NAME_SERVER_U_U_I_D)
  private String serverUUID;

  public static final String SERIALIZED_NAME_SERVER_UNIQUE_ID = "serverUniqueId";
  @SerializedName(SERIALIZED_NAME_SERVER_UNIQUE_ID)
  private String serverUniqueId;

  public static final String SERIALIZED_NAME_SERVER_EXTERNAL_ID = "serverExternalId";
  @SerializedName(SERIALIZED_NAME_SERVER_EXTERNAL_ID)
  private String serverExternalId;

  public static final String SERIALIZED_NAME_SERVER_INTERNAL_ID = "serverInternalId";
  @SerializedName(SERIALIZED_NAME_SERVER_INTERNAL_ID)
  private String serverInternalId;

  public static final String SERIALIZED_NAME_RESOURCE_POOL_ID = "resourcePoolId";
  @SerializedName(SERIALIZED_NAME_RESOURCE_POOL_ID)
  private Long resourcePoolId;

  public static final String SERIALIZED_NAME_RESOURCE_POOL_NAME = "resourcePoolName";
  @SerializedName(SERIALIZED_NAME_RESOURCE_POOL_NAME)
  private String resourcePoolName;


  public BillingInstancesContainers refType(String refType) {
    
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


  public BillingInstancesContainers refUUID(String refUUID) {
    
    this.refUUID = refUUID;
    return this;
  }

   /**
   * Get refUUID
   * @return refUUID
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRefUUID() {
    return refUUID;
  }


  public void setRefUUID(String refUUID) {
    this.refUUID = refUUID;
  }


  public BillingInstancesContainers refId(Long refId) {
    
    this.refId = refId;
    return this;
  }

   /**
   * Get refId
   * @return refId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getRefId() {
    return refId;
  }


  public void setRefId(Long refId) {
    this.refId = refId;
  }


  public BillingInstancesContainers startDate(OffsetDateTime startDate) {
    
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


  public BillingInstancesContainers endDate(OffsetDateTime endDate) {
    
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


  public BillingInstancesContainers cost(BigDecimal cost) {
    
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


  public BillingInstancesContainers price(BigDecimal price) {
    
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


  public BillingInstancesContainers numUnits(BigDecimal numUnits) {
    
    this.numUnits = numUnits;
    return this;
  }

   /**
   * Get numUnits
   * @return numUnits
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public BigDecimal getNumUnits() {
    return numUnits;
  }


  public void setNumUnits(BigDecimal numUnits) {
    this.numUnits = numUnits;
  }


  public BillingInstancesContainers unit(String unit) {
    
    this.unit = unit;
    return this;
  }

   /**
   * Get unit
   * @return unit
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getUnit() {
    return unit;
  }


  public void setUnit(String unit) {
    this.unit = unit;
  }


  public BillingInstancesContainers currency(String currency) {
    
    this.currency = currency;
    return this;
  }

   /**
   * Get currency
   * @return currency
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCurrency() {
    return currency;
  }


  public void setCurrency(String currency) {
    this.currency = currency;
  }


  public BillingInstancesContainers usages(List<BillingInstancesUsages> usages) {
    
    this.usages = usages;
    return this;
  }

  public BillingInstancesContainers addUsagesItem(BillingInstancesUsages usagesItem) {
    if (this.usages == null) {
      this.usages = new ArrayList<BillingInstancesUsages>();
    }
    this.usages.add(usagesItem);
    return this;
  }

   /**
   * Get usages
   * @return usages
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<BillingInstancesUsages> getUsages() {
    return usages;
  }


  public void setUsages(List<BillingInstancesUsages> usages) {
    this.usages = usages;
  }


  public BillingInstancesContainers numUsages(Long numUsages) {
    
    this.numUsages = numUsages;
    return this;
  }

   /**
   * Get numUsages
   * @return numUsages
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getNumUsages() {
    return numUsages;
  }


  public void setNumUsages(Long numUsages) {
    this.numUsages = numUsages;
  }


  public BillingInstancesContainers totalUsages(Long totalUsages) {
    
    this.totalUsages = totalUsages;
    return this;
  }

   /**
   * Get totalUsages
   * @return totalUsages
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getTotalUsages() {
    return totalUsages;
  }


  public void setTotalUsages(Long totalUsages) {
    this.totalUsages = totalUsages;
  }


  public BillingInstancesContainers hasMoreUsages(Boolean hasMoreUsages) {
    
    this.hasMoreUsages = hasMoreUsages;
    return this;
  }

   /**
   * Get hasMoreUsages
   * @return hasMoreUsages
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getHasMoreUsages() {
    return hasMoreUsages;
  }


  public void setHasMoreUsages(Boolean hasMoreUsages) {
    this.hasMoreUsages = hasMoreUsages;
  }


  public BillingInstancesContainers foundPricing(Boolean foundPricing) {
    
    this.foundPricing = foundPricing;
    return this;
  }

   /**
   * Get foundPricing
   * @return foundPricing
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getFoundPricing() {
    return foundPricing;
  }


  public void setFoundPricing(Boolean foundPricing) {
    this.foundPricing = foundPricing;
  }


  public BillingInstancesContainers name(String name) {
    
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


  public BillingInstancesContainers serverId(Long serverId) {
    
    this.serverId = serverId;
    return this;
  }

   /**
   * Get serverId
   * @return serverId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getServerId() {
    return serverId;
  }


  public void setServerId(Long serverId) {
    this.serverId = serverId;
  }


  public BillingInstancesContainers serverUUID(String serverUUID) {
    
    this.serverUUID = serverUUID;
    return this;
  }

   /**
   * Get serverUUID
   * @return serverUUID
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getServerUUID() {
    return serverUUID;
  }


  public void setServerUUID(String serverUUID) {
    this.serverUUID = serverUUID;
  }


  public BillingInstancesContainers serverUniqueId(String serverUniqueId) {
    
    this.serverUniqueId = serverUniqueId;
    return this;
  }

   /**
   * Get serverUniqueId
   * @return serverUniqueId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getServerUniqueId() {
    return serverUniqueId;
  }


  public void setServerUniqueId(String serverUniqueId) {
    this.serverUniqueId = serverUniqueId;
  }


  public BillingInstancesContainers serverExternalId(String serverExternalId) {
    
    this.serverExternalId = serverExternalId;
    return this;
  }

   /**
   * Get serverExternalId
   * @return serverExternalId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getServerExternalId() {
    return serverExternalId;
  }


  public void setServerExternalId(String serverExternalId) {
    this.serverExternalId = serverExternalId;
  }


  public BillingInstancesContainers serverInternalId(String serverInternalId) {
    
    this.serverInternalId = serverInternalId;
    return this;
  }

   /**
   * Get serverInternalId
   * @return serverInternalId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getServerInternalId() {
    return serverInternalId;
  }


  public void setServerInternalId(String serverInternalId) {
    this.serverInternalId = serverInternalId;
  }


  public BillingInstancesContainers resourcePoolId(Long resourcePoolId) {
    
    this.resourcePoolId = resourcePoolId;
    return this;
  }

   /**
   * Get resourcePoolId
   * @return resourcePoolId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getResourcePoolId() {
    return resourcePoolId;
  }


  public void setResourcePoolId(Long resourcePoolId) {
    this.resourcePoolId = resourcePoolId;
  }


  public BillingInstancesContainers resourcePoolName(String resourcePoolName) {
    
    this.resourcePoolName = resourcePoolName;
    return this;
  }

   /**
   * Get resourcePoolName
   * @return resourcePoolName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getResourcePoolName() {
    return resourcePoolName;
  }


  public void setResourcePoolName(String resourcePoolName) {
    this.resourcePoolName = resourcePoolName;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    BillingInstancesContainers billingInstancesContainers = (BillingInstancesContainers) o;
    return Objects.equals(this.refType, billingInstancesContainers.refType) &&
        Objects.equals(this.refUUID, billingInstancesContainers.refUUID) &&
        Objects.equals(this.refId, billingInstancesContainers.refId) &&
        Objects.equals(this.startDate, billingInstancesContainers.startDate) &&
        Objects.equals(this.endDate, billingInstancesContainers.endDate) &&
        Objects.equals(this.cost, billingInstancesContainers.cost) &&
        Objects.equals(this.price, billingInstancesContainers.price) &&
        Objects.equals(this.numUnits, billingInstancesContainers.numUnits) &&
        Objects.equals(this.unit, billingInstancesContainers.unit) &&
        Objects.equals(this.currency, billingInstancesContainers.currency) &&
        Objects.equals(this.usages, billingInstancesContainers.usages) &&
        Objects.equals(this.numUsages, billingInstancesContainers.numUsages) &&
        Objects.equals(this.totalUsages, billingInstancesContainers.totalUsages) &&
        Objects.equals(this.hasMoreUsages, billingInstancesContainers.hasMoreUsages) &&
        Objects.equals(this.foundPricing, billingInstancesContainers.foundPricing) &&
        Objects.equals(this.name, billingInstancesContainers.name) &&
        Objects.equals(this.serverId, billingInstancesContainers.serverId) &&
        Objects.equals(this.serverUUID, billingInstancesContainers.serverUUID) &&
        Objects.equals(this.serverUniqueId, billingInstancesContainers.serverUniqueId) &&
        Objects.equals(this.serverExternalId, billingInstancesContainers.serverExternalId) &&
        Objects.equals(this.serverInternalId, billingInstancesContainers.serverInternalId) &&
        Objects.equals(this.resourcePoolId, billingInstancesContainers.resourcePoolId) &&
        Objects.equals(this.resourcePoolName, billingInstancesContainers.resourcePoolName);
  }

  @Override
  public int hashCode() {
    return Objects.hash(refType, refUUID, refId, startDate, endDate, cost, price, numUnits, unit, currency, usages, numUsages, totalUsages, hasMoreUsages, foundPricing, name, serverId, serverUUID, serverUniqueId, serverExternalId, serverInternalId, resourcePoolId, resourcePoolName);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class BillingInstancesContainers {\n");
    sb.append("    refType: ").append(toIndentedString(refType)).append("\n");
    sb.append("    refUUID: ").append(toIndentedString(refUUID)).append("\n");
    sb.append("    refId: ").append(toIndentedString(refId)).append("\n");
    sb.append("    startDate: ").append(toIndentedString(startDate)).append("\n");
    sb.append("    endDate: ").append(toIndentedString(endDate)).append("\n");
    sb.append("    cost: ").append(toIndentedString(cost)).append("\n");
    sb.append("    price: ").append(toIndentedString(price)).append("\n");
    sb.append("    numUnits: ").append(toIndentedString(numUnits)).append("\n");
    sb.append("    unit: ").append(toIndentedString(unit)).append("\n");
    sb.append("    currency: ").append(toIndentedString(currency)).append("\n");
    sb.append("    usages: ").append(toIndentedString(usages)).append("\n");
    sb.append("    numUsages: ").append(toIndentedString(numUsages)).append("\n");
    sb.append("    totalUsages: ").append(toIndentedString(totalUsages)).append("\n");
    sb.append("    hasMoreUsages: ").append(toIndentedString(hasMoreUsages)).append("\n");
    sb.append("    foundPricing: ").append(toIndentedString(foundPricing)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    serverId: ").append(toIndentedString(serverId)).append("\n");
    sb.append("    serverUUID: ").append(toIndentedString(serverUUID)).append("\n");
    sb.append("    serverUniqueId: ").append(toIndentedString(serverUniqueId)).append("\n");
    sb.append("    serverExternalId: ").append(toIndentedString(serverExternalId)).append("\n");
    sb.append("    serverInternalId: ").append(toIndentedString(serverInternalId)).append("\n");
    sb.append("    resourcePoolId: ").append(toIndentedString(resourcePoolId)).append("\n");
    sb.append("    resourcePoolName: ").append(toIndentedString(resourcePoolName)).append("\n");
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

