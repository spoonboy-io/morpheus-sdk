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
 * BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-17T13:37:08.123540Z[Etc/UTC]")
public class BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_TYPE = "type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private String type;

  public static final String SERIALIZED_NAME_EXTERNAL_ID = "externalId";
  @SerializedName(SERIALIZED_NAME_EXTERNAL_ID)
  private String externalId;

  public static final String SERIALIZED_NAME_INTERNAL_ID = "internalId";
  @SerializedName(SERIALIZED_NAME_INTERNAL_ID)
  private String internalId;

  public static final String SERIALIZED_NAME_EXTERNAL_PATH = "externalPath";
  @SerializedName(SERIALIZED_NAME_EXTERNAL_PATH)
  private String externalPath;

  public BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore() {
  }

  public BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore id(Long id) {
    
    this.id = id;
    return this;
  }

   /**
   * Get id
   * @return id
  **/
  @javax.annotation.Nullable
  public Long getId() {
    return id;
  }


  public void setId(Long id) {
    this.id = id;
  }


  public BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore name(String name) {
    
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


  public BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore type(String type) {
    
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


  public BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore externalId(String externalId) {
    
    this.externalId = externalId;
    return this;
  }

   /**
   * Get externalId
   * @return externalId
  **/
  @javax.annotation.Nullable
  public String getExternalId() {
    return externalId;
  }


  public void setExternalId(String externalId) {
    this.externalId = externalId;
  }


  public BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore internalId(String internalId) {
    
    this.internalId = internalId;
    return this;
  }

   /**
   * Get internalId
   * @return internalId
  **/
  @javax.annotation.Nullable
  public String getInternalId() {
    return internalId;
  }


  public void setInternalId(String internalId) {
    this.internalId = internalId;
  }


  public BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore externalPath(String externalPath) {
    
    this.externalPath = externalPath;
    return this;
  }

   /**
   * Get externalPath
   * @return externalPath
  **/
  @javax.annotation.Nullable
  public String getExternalPath() {
    return externalPath;
  }


  public void setExternalPath(String externalPath) {
    this.externalPath = externalPath;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore billingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore = (BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) o;
    return Objects.equals(this.id, billingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore.id) &&
        Objects.equals(this.name, billingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore.name) &&
        Objects.equals(this.type, billingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore.type) &&
        Objects.equals(this.externalId, billingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore.externalId) &&
        Objects.equals(this.internalId, billingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore.internalId) &&
        Objects.equals(this.externalPath, billingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore.externalPath);
  }

  private static <T> boolean equalsNullable(JsonNullable<T> a, JsonNullable<T> b) {
    return a == b || (a != null && b != null && a.isPresent() && b.isPresent() && Objects.deepEquals(a.get(), b.get()));
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, name, type, externalId, internalId, externalPath);
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
    sb.append("class BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    externalId: ").append(toIndentedString(externalId)).append("\n");
    sb.append("    internalId: ").append(toIndentedString(internalId)).append("\n");
    sb.append("    externalPath: ").append(toIndentedString(externalPath)).append("\n");
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
    openapiFields.add("id");
    openapiFields.add("name");
    openapiFields.add("type");
    openapiFields.add("externalId");
    openapiFields.add("internalId");
    openapiFields.add("externalPath");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Element and throws an exception if issues found
  *
  * @param jsonElement JSON Element
  * @throws IOException if the JSON Element is invalid with respect to BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore
  */
  public static void validateJsonElement(JsonElement jsonElement) throws IOException {
      if (jsonElement == null) {
        if (!BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore.openapiRequiredFields.isEmpty()) { // has required fields but JSON element is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore is not found in the empty JSON string", BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonElement.getAsJsonObject().entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore` properties. JSON: %s", entry.getKey(), jsonElement.toString()));
        }
      }
        JsonObject jsonObj = jsonElement.getAsJsonObject();
      if ((jsonObj.get("name") != null && !jsonObj.get("name").isJsonNull()) && !jsonObj.get("name").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `name` to be a primitive type in the JSON string but got `%s`", jsonObj.get("name").toString()));
      }
      if ((jsonObj.get("type") != null && !jsonObj.get("type").isJsonNull()) && !jsonObj.get("type").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `type` to be a primitive type in the JSON string but got `%s`", jsonObj.get("type").toString()));
      }
      if ((jsonObj.get("externalId") != null && !jsonObj.get("externalId").isJsonNull()) && !jsonObj.get("externalId").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `externalId` to be a primitive type in the JSON string but got `%s`", jsonObj.get("externalId").toString()));
      }
      if ((jsonObj.get("internalId") != null && !jsonObj.get("internalId").isJsonNull()) && !jsonObj.get("internalId").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `internalId` to be a primitive type in the JSON string but got `%s`", jsonObj.get("internalId").toString()));
      }
      if ((jsonObj.get("externalPath") != null && !jsonObj.get("externalPath").isJsonNull()) && !jsonObj.get("externalPath").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `externalPath` to be a primitive type in the JSON string but got `%s`", jsonObj.get("externalPath").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore.class));

       return (TypeAdapter<T>) new TypeAdapter<BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore>() {
           @Override
           public void write(JsonWriter out, BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore read(JsonReader in) throws IOException {
             JsonElement jsonElement = elementAdapter.read(in);
             validateJsonElement(jsonElement);
             return thisAdapter.fromJsonTree(jsonElement);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore
  * @throws IOException if the JSON string is invalid with respect to BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore
  */
  public static BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore.class);
  }

 /**
  * Convert an instance of BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

