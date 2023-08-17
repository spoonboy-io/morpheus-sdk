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
 * BillingServerUsagesInnerApplicablePricesInnerPricesInner
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-17T13:37:08.123540Z[Etc/UTC]")
public class BillingServerUsagesInnerApplicablePricesInnerPricesInner {
  public static final String SERIALIZED_NAME_TYPE = "type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private String type;

  public static final String SERIALIZED_NAME_PRICE_PER_UNIT = "pricePerUnit";
  @SerializedName(SERIALIZED_NAME_PRICE_PER_UNIT)
  private BigDecimal pricePerUnit;

  public static final String SERIALIZED_NAME_COST_PER_UNIT = "costPerUnit";
  @SerializedName(SERIALIZED_NAME_COST_PER_UNIT)
  private BigDecimal costPerUnit;

  public static final String SERIALIZED_NAME_COST = "cost";
  @SerializedName(SERIALIZED_NAME_COST)
  private BigDecimal cost;

  public static final String SERIALIZED_NAME_PRICE = "price";
  @SerializedName(SERIALIZED_NAME_PRICE)
  private BigDecimal price;

  public static final String SERIALIZED_NAME_QUANTITY = "quantity";
  @SerializedName(SERIALIZED_NAME_QUANTITY)
  private Long quantity;

  public static final String SERIALIZED_NAME_DATASTORE_ID = "datastoreId";
  @SerializedName(SERIALIZED_NAME_DATASTORE_ID)
  private String datastoreId;

  public static final String SERIALIZED_NAME_VOLUME_TYPE = "volumeType";
  @SerializedName(SERIALIZED_NAME_VOLUME_TYPE)
  private String volumeType;

  public static final String SERIALIZED_NAME_DATASTORE = "datastore";
  @SerializedName(SERIALIZED_NAME_DATASTORE)
  private String datastore;

  public BillingServerUsagesInnerApplicablePricesInnerPricesInner() {
  }

  public BillingServerUsagesInnerApplicablePricesInnerPricesInner type(String type) {
    
    this.type = type;
    return this;
  }

   /**
   * Get type
   * @return type
  **/
  @javax.annotation.Nullable
  public String getType() {
    return type;
  }


  public void setType(String type) {
    this.type = type;
  }


  public BillingServerUsagesInnerApplicablePricesInnerPricesInner pricePerUnit(BigDecimal pricePerUnit) {
    
    this.pricePerUnit = pricePerUnit;
    return this;
  }

   /**
   * Get pricePerUnit
   * @return pricePerUnit
  **/
  @javax.annotation.Nullable
  public BigDecimal getPricePerUnit() {
    return pricePerUnit;
  }


  public void setPricePerUnit(BigDecimal pricePerUnit) {
    this.pricePerUnit = pricePerUnit;
  }


  public BillingServerUsagesInnerApplicablePricesInnerPricesInner costPerUnit(BigDecimal costPerUnit) {
    
    this.costPerUnit = costPerUnit;
    return this;
  }

   /**
   * Get costPerUnit
   * @return costPerUnit
  **/
  @javax.annotation.Nullable
  public BigDecimal getCostPerUnit() {
    return costPerUnit;
  }


  public void setCostPerUnit(BigDecimal costPerUnit) {
    this.costPerUnit = costPerUnit;
  }


  public BillingServerUsagesInnerApplicablePricesInnerPricesInner cost(BigDecimal cost) {
    
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


  public BillingServerUsagesInnerApplicablePricesInnerPricesInner price(BigDecimal price) {
    
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


  public BillingServerUsagesInnerApplicablePricesInnerPricesInner quantity(Long quantity) {
    
    this.quantity = quantity;
    return this;
  }

   /**
   * Get quantity
   * @return quantity
  **/
  @javax.annotation.Nullable
  public Long getQuantity() {
    return quantity;
  }


  public void setQuantity(Long quantity) {
    this.quantity = quantity;
  }


  public BillingServerUsagesInnerApplicablePricesInnerPricesInner datastoreId(String datastoreId) {
    
    this.datastoreId = datastoreId;
    return this;
  }

   /**
   * Get datastoreId
   * @return datastoreId
  **/
  @javax.annotation.Nullable
  public String getDatastoreId() {
    return datastoreId;
  }


  public void setDatastoreId(String datastoreId) {
    this.datastoreId = datastoreId;
  }


  public BillingServerUsagesInnerApplicablePricesInnerPricesInner volumeType(String volumeType) {
    
    this.volumeType = volumeType;
    return this;
  }

   /**
   * Get volumeType
   * @return volumeType
  **/
  @javax.annotation.Nullable
  public String getVolumeType() {
    return volumeType;
  }


  public void setVolumeType(String volumeType) {
    this.volumeType = volumeType;
  }


  public BillingServerUsagesInnerApplicablePricesInnerPricesInner datastore(String datastore) {
    
    this.datastore = datastore;
    return this;
  }

   /**
   * Get datastore
   * @return datastore
  **/
  @javax.annotation.Nullable
  public String getDatastore() {
    return datastore;
  }


  public void setDatastore(String datastore) {
    this.datastore = datastore;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    BillingServerUsagesInnerApplicablePricesInnerPricesInner billingServerUsagesInnerApplicablePricesInnerPricesInner = (BillingServerUsagesInnerApplicablePricesInnerPricesInner) o;
    return Objects.equals(this.type, billingServerUsagesInnerApplicablePricesInnerPricesInner.type) &&
        Objects.equals(this.pricePerUnit, billingServerUsagesInnerApplicablePricesInnerPricesInner.pricePerUnit) &&
        Objects.equals(this.costPerUnit, billingServerUsagesInnerApplicablePricesInnerPricesInner.costPerUnit) &&
        Objects.equals(this.cost, billingServerUsagesInnerApplicablePricesInnerPricesInner.cost) &&
        Objects.equals(this.price, billingServerUsagesInnerApplicablePricesInnerPricesInner.price) &&
        Objects.equals(this.quantity, billingServerUsagesInnerApplicablePricesInnerPricesInner.quantity) &&
        Objects.equals(this.datastoreId, billingServerUsagesInnerApplicablePricesInnerPricesInner.datastoreId) &&
        Objects.equals(this.volumeType, billingServerUsagesInnerApplicablePricesInnerPricesInner.volumeType) &&
        Objects.equals(this.datastore, billingServerUsagesInnerApplicablePricesInnerPricesInner.datastore);
  }

  private static <T> boolean equalsNullable(JsonNullable<T> a, JsonNullable<T> b) {
    return a == b || (a != null && b != null && a.isPresent() && b.isPresent() && Objects.deepEquals(a.get(), b.get()));
  }

  @Override
  public int hashCode() {
    return Objects.hash(type, pricePerUnit, costPerUnit, cost, price, quantity, datastoreId, volumeType, datastore);
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
    sb.append("class BillingServerUsagesInnerApplicablePricesInnerPricesInner {\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    pricePerUnit: ").append(toIndentedString(pricePerUnit)).append("\n");
    sb.append("    costPerUnit: ").append(toIndentedString(costPerUnit)).append("\n");
    sb.append("    cost: ").append(toIndentedString(cost)).append("\n");
    sb.append("    price: ").append(toIndentedString(price)).append("\n");
    sb.append("    quantity: ").append(toIndentedString(quantity)).append("\n");
    sb.append("    datastoreId: ").append(toIndentedString(datastoreId)).append("\n");
    sb.append("    volumeType: ").append(toIndentedString(volumeType)).append("\n");
    sb.append("    datastore: ").append(toIndentedString(datastore)).append("\n");
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
    openapiFields.add("type");
    openapiFields.add("pricePerUnit");
    openapiFields.add("costPerUnit");
    openapiFields.add("cost");
    openapiFields.add("price");
    openapiFields.add("quantity");
    openapiFields.add("datastoreId");
    openapiFields.add("volumeType");
    openapiFields.add("datastore");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Element and throws an exception if issues found
  *
  * @param jsonElement JSON Element
  * @throws IOException if the JSON Element is invalid with respect to BillingServerUsagesInnerApplicablePricesInnerPricesInner
  */
  public static void validateJsonElement(JsonElement jsonElement) throws IOException {
      if (jsonElement == null) {
        if (!BillingServerUsagesInnerApplicablePricesInnerPricesInner.openapiRequiredFields.isEmpty()) { // has required fields but JSON element is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in BillingServerUsagesInnerApplicablePricesInnerPricesInner is not found in the empty JSON string", BillingServerUsagesInnerApplicablePricesInnerPricesInner.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonElement.getAsJsonObject().entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!BillingServerUsagesInnerApplicablePricesInnerPricesInner.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `BillingServerUsagesInnerApplicablePricesInnerPricesInner` properties. JSON: %s", entry.getKey(), jsonElement.toString()));
        }
      }
        JsonObject jsonObj = jsonElement.getAsJsonObject();
      if ((jsonObj.get("type") != null && !jsonObj.get("type").isJsonNull()) && !jsonObj.get("type").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `type` to be a primitive type in the JSON string but got `%s`", jsonObj.get("type").toString()));
      }
      if ((jsonObj.get("datastoreId") != null && !jsonObj.get("datastoreId").isJsonNull()) && !jsonObj.get("datastoreId").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `datastoreId` to be a primitive type in the JSON string but got `%s`", jsonObj.get("datastoreId").toString()));
      }
      if ((jsonObj.get("volumeType") != null && !jsonObj.get("volumeType").isJsonNull()) && !jsonObj.get("volumeType").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `volumeType` to be a primitive type in the JSON string but got `%s`", jsonObj.get("volumeType").toString()));
      }
      if ((jsonObj.get("datastore") != null && !jsonObj.get("datastore").isJsonNull()) && !jsonObj.get("datastore").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `datastore` to be a primitive type in the JSON string but got `%s`", jsonObj.get("datastore").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!BillingServerUsagesInnerApplicablePricesInnerPricesInner.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'BillingServerUsagesInnerApplicablePricesInnerPricesInner' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<BillingServerUsagesInnerApplicablePricesInnerPricesInner> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(BillingServerUsagesInnerApplicablePricesInnerPricesInner.class));

       return (TypeAdapter<T>) new TypeAdapter<BillingServerUsagesInnerApplicablePricesInnerPricesInner>() {
           @Override
           public void write(JsonWriter out, BillingServerUsagesInnerApplicablePricesInnerPricesInner value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public BillingServerUsagesInnerApplicablePricesInnerPricesInner read(JsonReader in) throws IOException {
             JsonElement jsonElement = elementAdapter.read(in);
             validateJsonElement(jsonElement);
             return thisAdapter.fromJsonTree(jsonElement);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of BillingServerUsagesInnerApplicablePricesInnerPricesInner given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of BillingServerUsagesInnerApplicablePricesInnerPricesInner
  * @throws IOException if the JSON string is invalid with respect to BillingServerUsagesInnerApplicablePricesInnerPricesInner
  */
  public static BillingServerUsagesInnerApplicablePricesInnerPricesInner fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, BillingServerUsagesInnerApplicablePricesInnerPricesInner.class);
  }

 /**
  * Convert an instance of BillingServerUsagesInnerApplicablePricesInnerPricesInner to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

