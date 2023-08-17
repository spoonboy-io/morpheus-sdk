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
import org.openapitools.client.model.BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore;

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
 * BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-17T13:37:08.123540Z[Etc/UTC]")
public class BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner {
  public static final String SERIALIZED_NAME_SIZE = "size";
  @SerializedName(SERIALIZED_NAME_SIZE)
  private Long size;

  public static final String SERIALIZED_NAME_TYPE_CODE = "typeCode";
  @SerializedName(SERIALIZED_NAME_TYPE_CODE)
  private String typeCode;

  public static final String SERIALIZED_NAME_DATASTORE = "datastore";
  @SerializedName(SERIALIZED_NAME_DATASTORE)
  private BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore datastore;

  public BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner() {
  }

  public BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner size(Long size) {
    
    this.size = size;
    return this;
  }

   /**
   * Get size
   * @return size
  **/
  @javax.annotation.Nullable
  public Long getSize() {
    return size;
  }


  public void setSize(Long size) {
    this.size = size;
  }


  public BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner typeCode(String typeCode) {
    
    this.typeCode = typeCode;
    return this;
  }

   /**
   * Get typeCode
   * @return typeCode
  **/
  @javax.annotation.Nullable
  public String getTypeCode() {
    return typeCode;
  }


  public void setTypeCode(String typeCode) {
    this.typeCode = typeCode;
  }


  public BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner datastore(BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore datastore) {
    
    this.datastore = datastore;
    return this;
  }

   /**
   * Get datastore
   * @return datastore
  **/
  @javax.annotation.Nullable
  public BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore getDatastore() {
    return datastore;
  }


  public void setDatastore(BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore datastore) {
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
    BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner billingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner = (BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner) o;
    return Objects.equals(this.size, billingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner.size) &&
        Objects.equals(this.typeCode, billingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner.typeCode) &&
        Objects.equals(this.datastore, billingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner.datastore);
  }

  @Override
  public int hashCode() {
    return Objects.hash(size, typeCode, datastore);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner {\n");
    sb.append("    size: ").append(toIndentedString(size)).append("\n");
    sb.append("    typeCode: ").append(toIndentedString(typeCode)).append("\n");
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
    openapiFields.add("size");
    openapiFields.add("typeCode");
    openapiFields.add("datastore");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Element and throws an exception if issues found
  *
  * @param jsonElement JSON Element
  * @throws IOException if the JSON Element is invalid with respect to BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner
  */
  public static void validateJsonElement(JsonElement jsonElement) throws IOException {
      if (jsonElement == null) {
        if (!BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner.openapiRequiredFields.isEmpty()) { // has required fields but JSON element is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner is not found in the empty JSON string", BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonElement.getAsJsonObject().entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner` properties. JSON: %s", entry.getKey(), jsonElement.toString()));
        }
      }
        JsonObject jsonObj = jsonElement.getAsJsonObject();
      if ((jsonObj.get("typeCode") != null && !jsonObj.get("typeCode").isJsonNull()) && !jsonObj.get("typeCode").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `typeCode` to be a primitive type in the JSON string but got `%s`", jsonObj.get("typeCode").toString()));
      }
      // validate the optional field `datastore`
      if (jsonObj.get("datastore") != null && !jsonObj.get("datastore").isJsonNull()) {
        BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore.validateJsonElement(jsonObj.get("datastore"));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner.class));

       return (TypeAdapter<T>) new TypeAdapter<BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner>() {
           @Override
           public void write(JsonWriter out, BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner read(JsonReader in) throws IOException {
             JsonElement jsonElement = elementAdapter.read(in);
             validateJsonElement(jsonElement);
             return thisAdapter.fromJsonTree(jsonElement);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner
  * @throws IOException if the JSON string is invalid with respect to BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner
  */
  public static BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner.class);
  }

 /**
  * Convert an instance of BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

