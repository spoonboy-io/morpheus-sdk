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
import org.openapitools.client.model.ArchiveBucketUpdate;

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
 * UpdateArchiveBucketRequest
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-17T13:37:08.123540Z[Etc/UTC]")
public class UpdateArchiveBucketRequest {
  public static final String SERIALIZED_NAME_ARCHIVE_BUCKET = "archiveBucket";
  @SerializedName(SERIALIZED_NAME_ARCHIVE_BUCKET)
  private ArchiveBucketUpdate archiveBucket;

  public UpdateArchiveBucketRequest() {
  }

  public UpdateArchiveBucketRequest archiveBucket(ArchiveBucketUpdate archiveBucket) {
    
    this.archiveBucket = archiveBucket;
    return this;
  }

   /**
   * Get archiveBucket
   * @return archiveBucket
  **/
  @javax.annotation.Nullable
  public ArchiveBucketUpdate getArchiveBucket() {
    return archiveBucket;
  }


  public void setArchiveBucket(ArchiveBucketUpdate archiveBucket) {
    this.archiveBucket = archiveBucket;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    UpdateArchiveBucketRequest updateArchiveBucketRequest = (UpdateArchiveBucketRequest) o;
    return Objects.equals(this.archiveBucket, updateArchiveBucketRequest.archiveBucket);
  }

  @Override
  public int hashCode() {
    return Objects.hash(archiveBucket);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class UpdateArchiveBucketRequest {\n");
    sb.append("    archiveBucket: ").append(toIndentedString(archiveBucket)).append("\n");
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
    openapiFields.add("archiveBucket");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Element and throws an exception if issues found
  *
  * @param jsonElement JSON Element
  * @throws IOException if the JSON Element is invalid with respect to UpdateArchiveBucketRequest
  */
  public static void validateJsonElement(JsonElement jsonElement) throws IOException {
      if (jsonElement == null) {
        if (!UpdateArchiveBucketRequest.openapiRequiredFields.isEmpty()) { // has required fields but JSON element is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in UpdateArchiveBucketRequest is not found in the empty JSON string", UpdateArchiveBucketRequest.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonElement.getAsJsonObject().entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!UpdateArchiveBucketRequest.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `UpdateArchiveBucketRequest` properties. JSON: %s", entry.getKey(), jsonElement.toString()));
        }
      }
        JsonObject jsonObj = jsonElement.getAsJsonObject();
      // validate the optional field `archiveBucket`
      if (jsonObj.get("archiveBucket") != null && !jsonObj.get("archiveBucket").isJsonNull()) {
        ArchiveBucketUpdate.validateJsonElement(jsonObj.get("archiveBucket"));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!UpdateArchiveBucketRequest.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'UpdateArchiveBucketRequest' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<UpdateArchiveBucketRequest> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(UpdateArchiveBucketRequest.class));

       return (TypeAdapter<T>) new TypeAdapter<UpdateArchiveBucketRequest>() {
           @Override
           public void write(JsonWriter out, UpdateArchiveBucketRequest value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public UpdateArchiveBucketRequest read(JsonReader in) throws IOException {
             JsonElement jsonElement = elementAdapter.read(in);
             validateJsonElement(jsonElement);
             return thisAdapter.fromJsonTree(jsonElement);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of UpdateArchiveBucketRequest given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of UpdateArchiveBucketRequest
  * @throws IOException if the JSON string is invalid with respect to UpdateArchiveBucketRequest
  */
  public static UpdateArchiveBucketRequest fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, UpdateArchiveBucketRequest.class);
  }

 /**
  * Convert an instance of UpdateArchiveBucketRequest to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}
