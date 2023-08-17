/*
 * Morpheus API
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.1.1
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
import java.io.IOException;
import java.math.BigDecimal;
import java.time.OffsetDateTime;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import org.openapitools.client.model.BillingInstanceContainersInnerUsagesInner;
import org.openapitools.jackson.nullable.JsonNullable;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.JsonArray;
import com.google.gson.JsonDeserializationContext;
import com.google.gson.JsonDeserializer;
import com.google.gson.JsonElement;
import com.google.gson.JsonObject;
import com.google.gson.JsonParseException;
import com.google.gson.TypeAdapterFactory;
import com.google.gson.reflect.TypeToken;
import com.google.gson.TypeAdapter;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;

import java.lang.reflect.Type;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Map.Entry;
import java.util.Set;

import org.openapitools.client.JSON;

/**
 * BillingInstanceContainersInner
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-17T13:37:08.123540Z[Etc/UTC]")
public class BillingInstanceContainersInner {
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
  private List<BillingInstanceContainersInnerUsagesInner> usages;

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

  public BillingInstanceContainersInner() {
  }

  public BillingInstanceContainersInner refType(String refType) {
    
    this.refType = refType;
    return this;
  }

   /**
   * Get refType
   * @return refType
  **/
  @javax.annotation.Nullable
  public String getRefType() {
    return refType;
  }


  public void setRefType(String refType) {
    this.refType = refType;
  }


  public BillingInstanceContainersInner refUUID(String refUUID) {
    
    this.refUUID = refUUID;
    return this;
  }

   /**
   * Get refUUID
   * @return refUUID
  **/
  @javax.annotation.Nullable
  public String getRefUUID() {
    return refUUID;
  }


  public void setRefUUID(String refUUID) {
    this.refUUID = refUUID;
  }


  public BillingInstanceContainersInner refId(Long refId) {
    
    this.refId = refId;
    return this;
  }

   /**
   * Get refId
   * @return refId
  **/
  @javax.annotation.Nullable
  public Long getRefId() {
    return refId;
  }


  public void setRefId(Long refId) {
    this.refId = refId;
  }


  public BillingInstanceContainersInner startDate(OffsetDateTime startDate) {
    
    this.startDate = startDate;
    return this;
  }

   /**
   * Get startDate
   * @return startDate
  **/
  @javax.annotation.Nullable
  public OffsetDateTime getStartDate() {
    return startDate;
  }


  public void setStartDate(OffsetDateTime startDate) {
    this.startDate = startDate;
  }


  public BillingInstanceContainersInner endDate(OffsetDateTime endDate) {
    
    this.endDate = endDate;
    return this;
  }

   /**
   * Get endDate
   * @return endDate
  **/
  @javax.annotation.Nullable
  public OffsetDateTime getEndDate() {
    return endDate;
  }


  public void setEndDate(OffsetDateTime endDate) {
    this.endDate = endDate;
  }


  public BillingInstanceContainersInner cost(BigDecimal cost) {
    
    this.cost = cost;
    return this;
  }

   /**
   * Get cost
   * @return cost
  **/
  @javax.annotation.Nullable
  public BigDecimal getCost() {
    return cost;
  }


  public void setCost(BigDecimal cost) {
    this.cost = cost;
  }


  public BillingInstanceContainersInner price(BigDecimal price) {
    
    this.price = price;
    return this;
  }

   /**
   * Get price
   * @return price
  **/
  @javax.annotation.Nullable
  public BigDecimal getPrice() {
    return price;
  }


  public void setPrice(BigDecimal price) {
    this.price = price;
  }


  public BillingInstanceContainersInner numUnits(BigDecimal numUnits) {
    
    this.numUnits = numUnits;
    return this;
  }

   /**
   * Get numUnits
   * @return numUnits
  **/
  @javax.annotation.Nullable
  public BigDecimal getNumUnits() {
    return numUnits;
  }


  public void setNumUnits(BigDecimal numUnits) {
    this.numUnits = numUnits;
  }


  public BillingInstanceContainersInner unit(String unit) {
    
    this.unit = unit;
    return this;
  }

   /**
   * Get unit
   * @return unit
  **/
  @javax.annotation.Nullable
  public String getUnit() {
    return unit;
  }


  public void setUnit(String unit) {
    this.unit = unit;
  }


  public BillingInstanceContainersInner currency(String currency) {
    
    this.currency = currency;
    return this;
  }

   /**
   * Get currency
   * @return currency
  **/
  @javax.annotation.Nullable
  public String getCurrency() {
    return currency;
  }


  public void setCurrency(String currency) {
    this.currency = currency;
  }


  public BillingInstanceContainersInner usages(List<BillingInstanceContainersInnerUsagesInner> usages) {
    
    this.usages = usages;
    return this;
  }

  public BillingInstanceContainersInner addUsagesItem(BillingInstanceContainersInnerUsagesInner usagesItem) {
    if (this.usages == null) {
      this.usages = new ArrayList<>();
    }
    this.usages.add(usagesItem);
    return this;
  }

   /**
   * Get usages
   * @return usages
  **/
  @javax.annotation.Nullable
  public List<BillingInstanceContainersInnerUsagesInner> getUsages() {
    return usages;
  }


  public void setUsages(List<BillingInstanceContainersInnerUsagesInner> usages) {
    this.usages = usages;
  }


  public BillingInstanceContainersInner numUsages(Long numUsages) {
    
    this.numUsages = numUsages;
    return this;
  }

   /**
   * Get numUsages
   * @return numUsages
  **/
  @javax.annotation.Nullable
  public Long getNumUsages() {
    return numUsages;
  }


  public void setNumUsages(Long numUsages) {
    this.numUsages = numUsages;
  }


  public BillingInstanceContainersInner totalUsages(Long totalUsages) {
    
    this.totalUsages = totalUsages;
    return this;
  }

   /**
   * Get totalUsages
   * @return totalUsages
  **/
  @javax.annotation.Nullable
  public Long getTotalUsages() {
    return totalUsages;
  }


  public void setTotalUsages(Long totalUsages) {
    this.totalUsages = totalUsages;
  }


  public BillingInstanceContainersInner hasMoreUsages(Boolean hasMoreUsages) {
    
    this.hasMoreUsages = hasMoreUsages;
    return this;
  }

   /**
   * Get hasMoreUsages
   * @return hasMoreUsages
  **/
  @javax.annotation.Nullable
  public Boolean getHasMoreUsages() {
    return hasMoreUsages;
  }


  public void setHasMoreUsages(Boolean hasMoreUsages) {
    this.hasMoreUsages = hasMoreUsages;
  }


  public BillingInstanceContainersInner foundPricing(Boolean foundPricing) {
    
    this.foundPricing = foundPricing;
    return this;
  }

   /**
   * Get foundPricing
   * @return foundPricing
  **/
  @javax.annotation.Nullable
  public Boolean getFoundPricing() {
    return foundPricing;
  }


  public void setFoundPricing(Boolean foundPricing) {
    this.foundPricing = foundPricing;
  }


  public BillingInstanceContainersInner name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Get name
   * @return name
  **/
  @javax.annotation.Nullable
  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public BillingInstanceContainersInner serverId(Long serverId) {
    
    this.serverId = serverId;
    return this;
  }

   /**
   * Get serverId
   * @return serverId
  **/
  @javax.annotation.Nullable
  public Long getServerId() {
    return serverId;
  }


  public void setServerId(Long serverId) {
    this.serverId = serverId;
  }


  public BillingInstanceContainersInner serverUUID(String serverUUID) {
    
    this.serverUUID = serverUUID;
    return this;
  }

   /**
   * Get serverUUID
   * @return serverUUID
  **/
  @javax.annotation.Nullable
  public String getServerUUID() {
    return serverUUID;
  }


  public void setServerUUID(String serverUUID) {
    this.serverUUID = serverUUID;
  }


  public BillingInstanceContainersInner serverUniqueId(String serverUniqueId) {
    
    this.serverUniqueId = serverUniqueId;
    return this;
  }

   /**
   * Get serverUniqueId
   * @return serverUniqueId
  **/
  @javax.annotation.Nullable
  public String getServerUniqueId() {
    return serverUniqueId;
  }


  public void setServerUniqueId(String serverUniqueId) {
    this.serverUniqueId = serverUniqueId;
  }


  public BillingInstanceContainersInner serverExternalId(String serverExternalId) {
    
    this.serverExternalId = serverExternalId;
    return this;
  }

   /**
   * Get serverExternalId
   * @return serverExternalId
  **/
  @javax.annotation.Nullable
  public String getServerExternalId() {
    return serverExternalId;
  }


  public void setServerExternalId(String serverExternalId) {
    this.serverExternalId = serverExternalId;
  }


  public BillingInstanceContainersInner serverInternalId(String serverInternalId) {
    
    this.serverInternalId = serverInternalId;
    return this;
  }

   /**
   * Get serverInternalId
   * @return serverInternalId
  **/
  @javax.annotation.Nullable
  public String getServerInternalId() {
    return serverInternalId;
  }


  public void setServerInternalId(String serverInternalId) {
    this.serverInternalId = serverInternalId;
  }


  public BillingInstanceContainersInner resourcePoolId(Long resourcePoolId) {
    
    this.resourcePoolId = resourcePoolId;
    return this;
  }

   /**
   * Get resourcePoolId
   * @return resourcePoolId
  **/
  @javax.annotation.Nullable
  public Long getResourcePoolId() {
    return resourcePoolId;
  }


  public void setResourcePoolId(Long resourcePoolId) {
    this.resourcePoolId = resourcePoolId;
  }


  public BillingInstanceContainersInner resourcePoolName(String resourcePoolName) {
    
    this.resourcePoolName = resourcePoolName;
    return this;
  }

   /**
   * Get resourcePoolName
   * @return resourcePoolName
  **/
  @javax.annotation.Nullable
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
    BillingInstanceContainersInner billingInstanceContainersInner = (BillingInstanceContainersInner) o;
    return Objects.equals(this.refType, billingInstanceContainersInner.refType) &&
        Objects.equals(this.refUUID, billingInstanceContainersInner.refUUID) &&
        Objects.equals(this.refId, billingInstanceContainersInner.refId) &&
        Objects.equals(this.startDate, billingInstanceContainersInner.startDate) &&
        Objects.equals(this.endDate, billingInstanceContainersInner.endDate) &&
        Objects.equals(this.cost, billingInstanceContainersInner.cost) &&
        Objects.equals(this.price, billingInstanceContainersInner.price) &&
        Objects.equals(this.numUnits, billingInstanceContainersInner.numUnits) &&
        Objects.equals(this.unit, billingInstanceContainersInner.unit) &&
        Objects.equals(this.currency, billingInstanceContainersInner.currency) &&
        Objects.equals(this.usages, billingInstanceContainersInner.usages) &&
        Objects.equals(this.numUsages, billingInstanceContainersInner.numUsages) &&
        Objects.equals(this.totalUsages, billingInstanceContainersInner.totalUsages) &&
        Objects.equals(this.hasMoreUsages, billingInstanceContainersInner.hasMoreUsages) &&
        Objects.equals(this.foundPricing, billingInstanceContainersInner.foundPricing) &&
        Objects.equals(this.name, billingInstanceContainersInner.name) &&
        Objects.equals(this.serverId, billingInstanceContainersInner.serverId) &&
        Objects.equals(this.serverUUID, billingInstanceContainersInner.serverUUID) &&
        Objects.equals(this.serverUniqueId, billingInstanceContainersInner.serverUniqueId) &&
        Objects.equals(this.serverExternalId, billingInstanceContainersInner.serverExternalId) &&
        Objects.equals(this.serverInternalId, billingInstanceContainersInner.serverInternalId) &&
        Objects.equals(this.resourcePoolId, billingInstanceContainersInner.resourcePoolId) &&
        Objects.equals(this.resourcePoolName, billingInstanceContainersInner.resourcePoolName);
  }

  private static <T> boolean equalsNullable(JsonNullable<T> a, JsonNullable<T> b) {
    return a == b || (a != null && b != null && a.isPresent() && b.isPresent() && Objects.deepEquals(a.get(), b.get()));
  }

  @Override
  public int hashCode() {
    return Objects.hash(refType, refUUID, refId, startDate, endDate, cost, price, numUnits, unit, currency, usages, numUsages, totalUsages, hasMoreUsages, foundPricing, name, serverId, serverUUID, serverUniqueId, serverExternalId, serverInternalId, resourcePoolId, resourcePoolName);
  }

  private static <T> int hashCodeNullable(JsonNullable<T> a) {
    if (a == null) {
      return 1;
    }
    return a.isPresent() ? Arrays.deepHashCode(new Object[]{a.get()}) : 31;
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class BillingInstanceContainersInner {\n");
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


  public static HashSet<String> openapiFields;
  public static HashSet<String> openapiRequiredFields;

  static {
    // a set of all properties/fields (JSON key names)
    openapiFields = new HashSet<String>();
    openapiFields.add("refType");
    openapiFields.add("refUUID");
    openapiFields.add("refId");
    openapiFields.add("startDate");
    openapiFields.add("endDate");
    openapiFields.add("cost");
    openapiFields.add("price");
    openapiFields.add("numUnits");
    openapiFields.add("unit");
    openapiFields.add("currency");
    openapiFields.add("usages");
    openapiFields.add("numUsages");
    openapiFields.add("totalUsages");
    openapiFields.add("hasMoreUsages");
    openapiFields.add("foundPricing");
    openapiFields.add("name");
    openapiFields.add("serverId");
    openapiFields.add("serverUUID");
    openapiFields.add("serverUniqueId");
    openapiFields.add("serverExternalId");
    openapiFields.add("serverInternalId");
    openapiFields.add("resourcePoolId");
    openapiFields.add("resourcePoolName");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Element and throws an exception if issues found
  *
  * @param jsonElement JSON Element
  * @throws IOException if the JSON Element is invalid with respect to BillingInstanceContainersInner
  */
  public static void validateJsonElement(JsonElement jsonElement) throws IOException {
      if (jsonElement == null) {
        if (!BillingInstanceContainersInner.openapiRequiredFields.isEmpty()) { // has required fields but JSON element is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in BillingInstanceContainersInner is not found in the empty JSON string", BillingInstanceContainersInner.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonElement.getAsJsonObject().entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!BillingInstanceContainersInner.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `BillingInstanceContainersInner` properties. JSON: %s", entry.getKey(), jsonElement.toString()));
        }
      }
        JsonObject jsonObj = jsonElement.getAsJsonObject();
      if ((jsonObj.get("refType") != null && !jsonObj.get("refType").isJsonNull()) && !jsonObj.get("refType").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `refType` to be a primitive type in the JSON string but got `%s`", jsonObj.get("refType").toString()));
      }
      if ((jsonObj.get("refUUID") != null && !jsonObj.get("refUUID").isJsonNull()) && !jsonObj.get("refUUID").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `refUUID` to be a primitive type in the JSON string but got `%s`", jsonObj.get("refUUID").toString()));
      }
      if ((jsonObj.get("unit") != null && !jsonObj.get("unit").isJsonNull()) && !jsonObj.get("unit").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `unit` to be a primitive type in the JSON string but got `%s`", jsonObj.get("unit").toString()));
      }
      if ((jsonObj.get("currency") != null && !jsonObj.get("currency").isJsonNull()) && !jsonObj.get("currency").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `currency` to be a primitive type in the JSON string but got `%s`", jsonObj.get("currency").toString()));
      }
      if (jsonObj.get("usages") != null && !jsonObj.get("usages").isJsonNull()) {
        JsonArray jsonArrayusages = jsonObj.getAsJsonArray("usages");
        if (jsonArrayusages != null) {
          // ensure the json data is an array
          if (!jsonObj.get("usages").isJsonArray()) {
            throw new IllegalArgumentException(String.format("Expected the field `usages` to be an array in the JSON string but got `%s`", jsonObj.get("usages").toString()));
          }

          // validate the optional field `usages` (array)
          for (int i = 0; i < jsonArrayusages.size(); i++) {
            BillingInstanceContainersInnerUsagesInner.validateJsonElement(jsonArrayusages.get(i));
          };
        }
      }
      if ((jsonObj.get("name") != null && !jsonObj.get("name").isJsonNull()) && !jsonObj.get("name").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `name` to be a primitive type in the JSON string but got `%s`", jsonObj.get("name").toString()));
      }
      if ((jsonObj.get("serverUUID") != null && !jsonObj.get("serverUUID").isJsonNull()) && !jsonObj.get("serverUUID").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `serverUUID` to be a primitive type in the JSON string but got `%s`", jsonObj.get("serverUUID").toString()));
      }
      if ((jsonObj.get("serverUniqueId") != null && !jsonObj.get("serverUniqueId").isJsonNull()) && !jsonObj.get("serverUniqueId").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `serverUniqueId` to be a primitive type in the JSON string but got `%s`", jsonObj.get("serverUniqueId").toString()));
      }
      if ((jsonObj.get("serverExternalId") != null && !jsonObj.get("serverExternalId").isJsonNull()) && !jsonObj.get("serverExternalId").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `serverExternalId` to be a primitive type in the JSON string but got `%s`", jsonObj.get("serverExternalId").toString()));
      }
      if ((jsonObj.get("serverInternalId") != null && !jsonObj.get("serverInternalId").isJsonNull()) && !jsonObj.get("serverInternalId").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `serverInternalId` to be a primitive type in the JSON string but got `%s`", jsonObj.get("serverInternalId").toString()));
      }
      if ((jsonObj.get("resourcePoolName") != null && !jsonObj.get("resourcePoolName").isJsonNull()) && !jsonObj.get("resourcePoolName").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `resourcePoolName` to be a primitive type in the JSON string but got `%s`", jsonObj.get("resourcePoolName").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!BillingInstanceContainersInner.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'BillingInstanceContainersInner' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<BillingInstanceContainersInner> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(BillingInstanceContainersInner.class));

       return (TypeAdapter<T>) new TypeAdapter<BillingInstanceContainersInner>() {
           @Override
           public void write(JsonWriter out, BillingInstanceContainersInner value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public BillingInstanceContainersInner read(JsonReader in) throws IOException {
             JsonElement jsonElement = elementAdapter.read(in);
             validateJsonElement(jsonElement);
             return thisAdapter.fromJsonTree(jsonElement);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of BillingInstanceContainersInner given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of BillingInstanceContainersInner
  * @throws IOException if the JSON string is invalid with respect to BillingInstanceContainersInner
  */
  public static BillingInstanceContainersInner fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, BillingInstanceContainersInner.class);
  }

 /**
  * Convert an instance of BillingInstanceContainersInner to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

