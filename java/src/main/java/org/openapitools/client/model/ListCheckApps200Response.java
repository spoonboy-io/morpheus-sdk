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
import org.openapitools.client.model.CheckApp;
import org.openapitools.client.model.MetaMeta;

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
 * ListCheckApps200Response
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-17T13:37:08.123540Z[Etc/UTC]")
public class ListCheckApps200Response {
  public static final String SERIALIZED_NAME_META = "meta";
  @SerializedName(SERIALIZED_NAME_META)
  private MetaMeta meta;

  public static final String SERIALIZED_NAME_MONITOR_APPS = "monitorApps";
  @SerializedName(SERIALIZED_NAME_MONITOR_APPS)
  private List<CheckApp> monitorApps;

  public ListCheckApps200Response() {
  }

  public ListCheckApps200Response meta(MetaMeta meta) {
    
    this.meta = meta;
    return this;
  }

   /**
   * Get meta
   * @return meta
  **/
  @javax.annotation.Nullable
  public MetaMeta getMeta() {
    return meta;
  }


  public void setMeta(MetaMeta meta) {
    this.meta = meta;
  }


  public ListCheckApps200Response monitorApps(List<CheckApp> monitorApps) {
    
    this.monitorApps = monitorApps;
    return this;
  }

  public ListCheckApps200Response addMonitorAppsItem(CheckApp monitorAppsItem) {
    if (this.monitorApps == null) {
      this.monitorApps = new ArrayList<>();
    }
    this.monitorApps.add(monitorAppsItem);
    return this;
  }

   /**
   * Get monitorApps
   * @return monitorApps
  **/
  @javax.annotation.Nullable
  public List<CheckApp> getMonitorApps() {
    return monitorApps;
  }


  public void setMonitorApps(List<CheckApp> monitorApps) {
    this.monitorApps = monitorApps;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ListCheckApps200Response listCheckApps200Response = (ListCheckApps200Response) o;
    return Objects.equals(this.meta, listCheckApps200Response.meta) &&
        Objects.equals(this.monitorApps, listCheckApps200Response.monitorApps);
  }

  @Override
  public int hashCode() {
    return Objects.hash(meta, monitorApps);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ListCheckApps200Response {\n");
    sb.append("    meta: ").append(toIndentedString(meta)).append("\n");
    sb.append("    monitorApps: ").append(toIndentedString(monitorApps)).append("\n");
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
    openapiFields.add("meta");
    openapiFields.add("monitorApps");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Element and throws an exception if issues found
  *
  * @param jsonElement JSON Element
  * @throws IOException if the JSON Element is invalid with respect to ListCheckApps200Response
  */
  public static void validateJsonElement(JsonElement jsonElement) throws IOException {
      if (jsonElement == null) {
        if (!ListCheckApps200Response.openapiRequiredFields.isEmpty()) { // has required fields but JSON element is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in ListCheckApps200Response is not found in the empty JSON string", ListCheckApps200Response.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonElement.getAsJsonObject().entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!ListCheckApps200Response.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `ListCheckApps200Response` properties. JSON: %s", entry.getKey(), jsonElement.toString()));
        }
      }
        JsonObject jsonObj = jsonElement.getAsJsonObject();
      // validate the optional field `meta`
      if (jsonObj.get("meta") != null && !jsonObj.get("meta").isJsonNull()) {
        MetaMeta.validateJsonElement(jsonObj.get("meta"));
      }
      if (jsonObj.get("monitorApps") != null && !jsonObj.get("monitorApps").isJsonNull()) {
        JsonArray jsonArraymonitorApps = jsonObj.getAsJsonArray("monitorApps");
        if (jsonArraymonitorApps != null) {
          // ensure the json data is an array
          if (!jsonObj.get("monitorApps").isJsonArray()) {
            throw new IllegalArgumentException(String.format("Expected the field `monitorApps` to be an array in the JSON string but got `%s`", jsonObj.get("monitorApps").toString()));
          }

          // validate the optional field `monitorApps` (array)
          for (int i = 0; i < jsonArraymonitorApps.size(); i++) {
            CheckApp.validateJsonElement(jsonArraymonitorApps.get(i));
          };
        }
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!ListCheckApps200Response.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'ListCheckApps200Response' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<ListCheckApps200Response> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(ListCheckApps200Response.class));

       return (TypeAdapter<T>) new TypeAdapter<ListCheckApps200Response>() {
           @Override
           public void write(JsonWriter out, ListCheckApps200Response value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public ListCheckApps200Response read(JsonReader in) throws IOException {
             JsonElement jsonElement = elementAdapter.read(in);
             validateJsonElement(jsonElement);
             return thisAdapter.fromJsonTree(jsonElement);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of ListCheckApps200Response given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of ListCheckApps200Response
  * @throws IOException if the JSON string is invalid with respect to ListCheckApps200Response
  */
  public static ListCheckApps200Response fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, ListCheckApps200Response.class);
  }

 /**
  * Convert an instance of ListCheckApps200Response to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}
