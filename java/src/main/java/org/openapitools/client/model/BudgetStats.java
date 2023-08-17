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
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import org.openapitools.client.model.BudgetStatsCurrent;
import org.openapitools.client.model.BudgetStatsIntervalsInner;

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
 * BudgetStats
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-17T13:37:08.123540Z[Etc/UTC]")
public class BudgetStats {
  public static final String SERIALIZED_NAME_CURRENCY = "currency";
  @SerializedName(SERIALIZED_NAME_CURRENCY)
  private String currency;

  public static final String SERIALIZED_NAME_CONVERSION_RATE = "conversionRate";
  @SerializedName(SERIALIZED_NAME_CONVERSION_RATE)
  private Long conversionRate;

  public static final String SERIALIZED_NAME_INTERVALS = "intervals";
  @SerializedName(SERIALIZED_NAME_INTERVALS)
  private List<BudgetStatsIntervalsInner> intervals;

  public static final String SERIALIZED_NAME_CURRENT = "current";
  @SerializedName(SERIALIZED_NAME_CURRENT)
  private BudgetStatsCurrent current;

  public BudgetStats() {
  }

  public BudgetStats currency(String currency) {
    
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


  public BudgetStats conversionRate(Long conversionRate) {
    
    this.conversionRate = conversionRate;
    return this;
  }

   /**
   * Get conversionRate
   * @return conversionRate
  **/
  @javax.annotation.Nullable
  public Long getConversionRate() {
    return conversionRate;
  }


  public void setConversionRate(Long conversionRate) {
    this.conversionRate = conversionRate;
  }


  public BudgetStats intervals(List<BudgetStatsIntervalsInner> intervals) {
    
    this.intervals = intervals;
    return this;
  }

  public BudgetStats addIntervalsItem(BudgetStatsIntervalsInner intervalsItem) {
    if (this.intervals == null) {
      this.intervals = new ArrayList<>();
    }
    this.intervals.add(intervalsItem);
    return this;
  }

   /**
   * Get intervals
   * @return intervals
  **/
  @javax.annotation.Nullable
  public List<BudgetStatsIntervalsInner> getIntervals() {
    return intervals;
  }


  public void setIntervals(List<BudgetStatsIntervalsInner> intervals) {
    this.intervals = intervals;
  }


  public BudgetStats current(BudgetStatsCurrent current) {
    
    this.current = current;
    return this;
  }

   /**
   * Get current
   * @return current
  **/
  @javax.annotation.Nullable
  public BudgetStatsCurrent getCurrent() {
    return current;
  }


  public void setCurrent(BudgetStatsCurrent current) {
    this.current = current;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    BudgetStats budgetStats = (BudgetStats) o;
    return Objects.equals(this.currency, budgetStats.currency) &&
        Objects.equals(this.conversionRate, budgetStats.conversionRate) &&
        Objects.equals(this.intervals, budgetStats.intervals) &&
        Objects.equals(this.current, budgetStats.current);
  }

  @Override
  public int hashCode() {
    return Objects.hash(currency, conversionRate, intervals, current);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class BudgetStats {\n");
    sb.append("    currency: ").append(toIndentedString(currency)).append("\n");
    sb.append("    conversionRate: ").append(toIndentedString(conversionRate)).append("\n");
    sb.append("    intervals: ").append(toIndentedString(intervals)).append("\n");
    sb.append("    current: ").append(toIndentedString(current)).append("\n");
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
    openapiFields.add("currency");
    openapiFields.add("conversionRate");
    openapiFields.add("intervals");
    openapiFields.add("current");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Element and throws an exception if issues found
  *
  * @param jsonElement JSON Element
  * @throws IOException if the JSON Element is invalid with respect to BudgetStats
  */
  public static void validateJsonElement(JsonElement jsonElement) throws IOException {
      if (jsonElement == null) {
        if (!BudgetStats.openapiRequiredFields.isEmpty()) { // has required fields but JSON element is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in BudgetStats is not found in the empty JSON string", BudgetStats.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonElement.getAsJsonObject().entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!BudgetStats.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `BudgetStats` properties. JSON: %s", entry.getKey(), jsonElement.toString()));
        }
      }
        JsonObject jsonObj = jsonElement.getAsJsonObject();
      if ((jsonObj.get("currency") != null && !jsonObj.get("currency").isJsonNull()) && !jsonObj.get("currency").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `currency` to be a primitive type in the JSON string but got `%s`", jsonObj.get("currency").toString()));
      }
      if (jsonObj.get("intervals") != null && !jsonObj.get("intervals").isJsonNull()) {
        JsonArray jsonArrayintervals = jsonObj.getAsJsonArray("intervals");
        if (jsonArrayintervals != null) {
          // ensure the json data is an array
          if (!jsonObj.get("intervals").isJsonArray()) {
            throw new IllegalArgumentException(String.format("Expected the field `intervals` to be an array in the JSON string but got `%s`", jsonObj.get("intervals").toString()));
          }

          // validate the optional field `intervals` (array)
          for (int i = 0; i < jsonArrayintervals.size(); i++) {
            BudgetStatsIntervalsInner.validateJsonElement(jsonArrayintervals.get(i));
          };
        }
      }
      // validate the optional field `current`
      if (jsonObj.get("current") != null && !jsonObj.get("current").isJsonNull()) {
        BudgetStatsCurrent.validateJsonElement(jsonObj.get("current"));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!BudgetStats.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'BudgetStats' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<BudgetStats> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(BudgetStats.class));

       return (TypeAdapter<T>) new TypeAdapter<BudgetStats>() {
           @Override
           public void write(JsonWriter out, BudgetStats value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public BudgetStats read(JsonReader in) throws IOException {
             JsonElement jsonElement = elementAdapter.read(in);
             validateJsonElement(jsonElement);
             return thisAdapter.fromJsonTree(jsonElement);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of BudgetStats given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of BudgetStats
  * @throws IOException if the JSON string is invalid with respect to BudgetStats
  */
  public static BudgetStats fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, BudgetStats.class);
  }

 /**
  * Convert an instance of BudgetStats to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

