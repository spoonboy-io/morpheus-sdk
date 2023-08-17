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
import org.openapitools.client.model.Approvals;
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
 * ListApprovals200Response
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-17T13:37:08.123540Z[Etc/UTC]")
public class ListApprovals200Response {
  public static final String SERIALIZED_NAME_META = "meta";
  @SerializedName(SERIALIZED_NAME_META)
  private MetaMeta meta;

  public static final String SERIALIZED_NAME_APPROVALS = "approvals";
  @SerializedName(SERIALIZED_NAME_APPROVALS)
  private List<Approvals> approvals;

  public ListApprovals200Response() {
  }

  public ListApprovals200Response meta(MetaMeta meta) {
    
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


  public ListApprovals200Response approvals(List<Approvals> approvals) {
    
    this.approvals = approvals;
    return this;
  }

  public ListApprovals200Response addApprovalsItem(Approvals approvalsItem) {
    if (this.approvals == null) {
      this.approvals = new ArrayList<>();
    }
    this.approvals.add(approvalsItem);
    return this;
  }

   /**
   * Get approvals
   * @return approvals
  **/
  @javax.annotation.Nullable
  public List<Approvals> getApprovals() {
    return approvals;
  }


  public void setApprovals(List<Approvals> approvals) {
    this.approvals = approvals;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ListApprovals200Response listApprovals200Response = (ListApprovals200Response) o;
    return Objects.equals(this.meta, listApprovals200Response.meta) &&
        Objects.equals(this.approvals, listApprovals200Response.approvals);
  }

  @Override
  public int hashCode() {
    return Objects.hash(meta, approvals);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ListApprovals200Response {\n");
    sb.append("    meta: ").append(toIndentedString(meta)).append("\n");
    sb.append("    approvals: ").append(toIndentedString(approvals)).append("\n");
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
    openapiFields.add("approvals");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Element and throws an exception if issues found
  *
  * @param jsonElement JSON Element
  * @throws IOException if the JSON Element is invalid with respect to ListApprovals200Response
  */
  public static void validateJsonElement(JsonElement jsonElement) throws IOException {
      if (jsonElement == null) {
        if (!ListApprovals200Response.openapiRequiredFields.isEmpty()) { // has required fields but JSON element is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in ListApprovals200Response is not found in the empty JSON string", ListApprovals200Response.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonElement.getAsJsonObject().entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!ListApprovals200Response.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `ListApprovals200Response` properties. JSON: %s", entry.getKey(), jsonElement.toString()));
        }
      }
        JsonObject jsonObj = jsonElement.getAsJsonObject();
      // validate the optional field `meta`
      if (jsonObj.get("meta") != null && !jsonObj.get("meta").isJsonNull()) {
        MetaMeta.validateJsonElement(jsonObj.get("meta"));
      }
      if (jsonObj.get("approvals") != null && !jsonObj.get("approvals").isJsonNull()) {
        JsonArray jsonArrayapprovals = jsonObj.getAsJsonArray("approvals");
        if (jsonArrayapprovals != null) {
          // ensure the json data is an array
          if (!jsonObj.get("approvals").isJsonArray()) {
            throw new IllegalArgumentException(String.format("Expected the field `approvals` to be an array in the JSON string but got `%s`", jsonObj.get("approvals").toString()));
          }

          // validate the optional field `approvals` (array)
          for (int i = 0; i < jsonArrayapprovals.size(); i++) {
            Approvals.validateJsonElement(jsonArrayapprovals.get(i));
          };
        }
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!ListApprovals200Response.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'ListApprovals200Response' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<ListApprovals200Response> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(ListApprovals200Response.class));

       return (TypeAdapter<T>) new TypeAdapter<ListApprovals200Response>() {
           @Override
           public void write(JsonWriter out, ListApprovals200Response value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public ListApprovals200Response read(JsonReader in) throws IOException {
             JsonElement jsonElement = elementAdapter.read(in);
             validateJsonElement(jsonElement);
             return thisAdapter.fromJsonTree(jsonElement);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of ListApprovals200Response given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of ListApprovals200Response
  * @throws IOException if the JSON string is invalid with respect to ListApprovals200Response
  */
  public static ListApprovals200Response fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, ListApprovals200Response.class);
  }

 /**
  * Convert an instance of ListApprovals200Response to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

