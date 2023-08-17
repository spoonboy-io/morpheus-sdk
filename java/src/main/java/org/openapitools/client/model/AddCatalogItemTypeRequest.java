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
import org.openapitools.client.model.AddCatalogItemTypeRequestCatalogItemType;

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
 * AddCatalogItemTypeRequest
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-17T13:37:08.123540Z[Etc/UTC]")
public class AddCatalogItemTypeRequest {
  public static final String SERIALIZED_NAME_CATALOG_ITEM_TYPE = "catalogItemType";
  @SerializedName(SERIALIZED_NAME_CATALOG_ITEM_TYPE)
  private AddCatalogItemTypeRequestCatalogItemType catalogItemType;

  public AddCatalogItemTypeRequest() {
  }

  public AddCatalogItemTypeRequest catalogItemType(AddCatalogItemTypeRequestCatalogItemType catalogItemType) {
    
    this.catalogItemType = catalogItemType;
    return this;
  }

   /**
   * Get catalogItemType
   * @return catalogItemType
  **/
  @javax.annotation.Nullable
  public AddCatalogItemTypeRequestCatalogItemType getCatalogItemType() {
    return catalogItemType;
  }


  public void setCatalogItemType(AddCatalogItemTypeRequestCatalogItemType catalogItemType) {
    this.catalogItemType = catalogItemType;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    AddCatalogItemTypeRequest addCatalogItemTypeRequest = (AddCatalogItemTypeRequest) o;
    return Objects.equals(this.catalogItemType, addCatalogItemTypeRequest.catalogItemType);
  }

  @Override
  public int hashCode() {
    return Objects.hash(catalogItemType);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class AddCatalogItemTypeRequest {\n");
    sb.append("    catalogItemType: ").append(toIndentedString(catalogItemType)).append("\n");
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
    openapiFields.add("catalogItemType");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Element and throws an exception if issues found
  *
  * @param jsonElement JSON Element
  * @throws IOException if the JSON Element is invalid with respect to AddCatalogItemTypeRequest
  */
  public static void validateJsonElement(JsonElement jsonElement) throws IOException {
      if (jsonElement == null) {
        if (!AddCatalogItemTypeRequest.openapiRequiredFields.isEmpty()) { // has required fields but JSON element is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in AddCatalogItemTypeRequest is not found in the empty JSON string", AddCatalogItemTypeRequest.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonElement.getAsJsonObject().entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!AddCatalogItemTypeRequest.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `AddCatalogItemTypeRequest` properties. JSON: %s", entry.getKey(), jsonElement.toString()));
        }
      }
        JsonObject jsonObj = jsonElement.getAsJsonObject();
      // validate the optional field `catalogItemType`
      if (jsonObj.get("catalogItemType") != null && !jsonObj.get("catalogItemType").isJsonNull()) {
        AddCatalogItemTypeRequestCatalogItemType.validateJsonElement(jsonObj.get("catalogItemType"));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!AddCatalogItemTypeRequest.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'AddCatalogItemTypeRequest' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<AddCatalogItemTypeRequest> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(AddCatalogItemTypeRequest.class));

       return (TypeAdapter<T>) new TypeAdapter<AddCatalogItemTypeRequest>() {
           @Override
           public void write(JsonWriter out, AddCatalogItemTypeRequest value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public AddCatalogItemTypeRequest read(JsonReader in) throws IOException {
             JsonElement jsonElement = elementAdapter.read(in);
             validateJsonElement(jsonElement);
             return thisAdapter.fromJsonTree(jsonElement);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of AddCatalogItemTypeRequest given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of AddCatalogItemTypeRequest
  * @throws IOException if the JSON string is invalid with respect to AddCatalogItemTypeRequest
  */
  public static AddCatalogItemTypeRequest fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, AddCatalogItemTypeRequest.class);
  }

 /**
  * Convert an instance of AddCatalogItemTypeRequest to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

