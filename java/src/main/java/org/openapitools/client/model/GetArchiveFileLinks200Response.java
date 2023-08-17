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
import org.openapitools.client.model.ArchiveFileLinks;

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
 * GetArchiveFileLinks200Response
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-17T13:37:08.123540Z[Etc/UTC]")
public class GetArchiveFileLinks200Response {
  public static final String SERIALIZED_NAME_ARCHIVE_FILE_LINKS = "archiveFileLinks";
  @SerializedName(SERIALIZED_NAME_ARCHIVE_FILE_LINKS)
  private List<ArchiveFileLinks> archiveFileLinks;

  public GetArchiveFileLinks200Response() {
  }

  public GetArchiveFileLinks200Response archiveFileLinks(List<ArchiveFileLinks> archiveFileLinks) {
    
    this.archiveFileLinks = archiveFileLinks;
    return this;
  }

  public GetArchiveFileLinks200Response addArchiveFileLinksItem(ArchiveFileLinks archiveFileLinksItem) {
    if (this.archiveFileLinks == null) {
      this.archiveFileLinks = new ArrayList<>();
    }
    this.archiveFileLinks.add(archiveFileLinksItem);
    return this;
  }

   /**
   * Get archiveFileLinks
   * @return archiveFileLinks
  **/
  @javax.annotation.Nullable
  public List<ArchiveFileLinks> getArchiveFileLinks() {
    return archiveFileLinks;
  }


  public void setArchiveFileLinks(List<ArchiveFileLinks> archiveFileLinks) {
    this.archiveFileLinks = archiveFileLinks;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    GetArchiveFileLinks200Response getArchiveFileLinks200Response = (GetArchiveFileLinks200Response) o;
    return Objects.equals(this.archiveFileLinks, getArchiveFileLinks200Response.archiveFileLinks);
  }

  @Override
  public int hashCode() {
    return Objects.hash(archiveFileLinks);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class GetArchiveFileLinks200Response {\n");
    sb.append("    archiveFileLinks: ").append(toIndentedString(archiveFileLinks)).append("\n");
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
    openapiFields.add("archiveFileLinks");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Element and throws an exception if issues found
  *
  * @param jsonElement JSON Element
  * @throws IOException if the JSON Element is invalid with respect to GetArchiveFileLinks200Response
  */
  public static void validateJsonElement(JsonElement jsonElement) throws IOException {
      if (jsonElement == null) {
        if (!GetArchiveFileLinks200Response.openapiRequiredFields.isEmpty()) { // has required fields but JSON element is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in GetArchiveFileLinks200Response is not found in the empty JSON string", GetArchiveFileLinks200Response.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonElement.getAsJsonObject().entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!GetArchiveFileLinks200Response.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `GetArchiveFileLinks200Response` properties. JSON: %s", entry.getKey(), jsonElement.toString()));
        }
      }
        JsonObject jsonObj = jsonElement.getAsJsonObject();
      if (jsonObj.get("archiveFileLinks") != null && !jsonObj.get("archiveFileLinks").isJsonNull()) {
        JsonArray jsonArrayarchiveFileLinks = jsonObj.getAsJsonArray("archiveFileLinks");
        if (jsonArrayarchiveFileLinks != null) {
          // ensure the json data is an array
          if (!jsonObj.get("archiveFileLinks").isJsonArray()) {
            throw new IllegalArgumentException(String.format("Expected the field `archiveFileLinks` to be an array in the JSON string but got `%s`", jsonObj.get("archiveFileLinks").toString()));
          }

          // validate the optional field `archiveFileLinks` (array)
          for (int i = 0; i < jsonArrayarchiveFileLinks.size(); i++) {
            ArchiveFileLinks.validateJsonElement(jsonArrayarchiveFileLinks.get(i));
          };
        }
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!GetArchiveFileLinks200Response.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'GetArchiveFileLinks200Response' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<GetArchiveFileLinks200Response> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(GetArchiveFileLinks200Response.class));

       return (TypeAdapter<T>) new TypeAdapter<GetArchiveFileLinks200Response>() {
           @Override
           public void write(JsonWriter out, GetArchiveFileLinks200Response value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public GetArchiveFileLinks200Response read(JsonReader in) throws IOException {
             JsonElement jsonElement = elementAdapter.read(in);
             validateJsonElement(jsonElement);
             return thisAdapter.fromJsonTree(jsonElement);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of GetArchiveFileLinks200Response given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of GetArchiveFileLinks200Response
  * @throws IOException if the JSON string is invalid with respect to GetArchiveFileLinks200Response
  */
  public static GetArchiveFileLinks200Response fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, GetArchiveFileLinks200Response.class);
  }

 /**
  * Convert an instance of GetArchiveFileLinks200Response to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}
