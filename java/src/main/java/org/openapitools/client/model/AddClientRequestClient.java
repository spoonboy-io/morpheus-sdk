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
 * Payload for creating a new oauth client
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-17T13:37:08.123540Z[Etc/UTC]")
public class AddClientRequestClient {
  public static final String SERIALIZED_NAME_CLIENT_ID = "clientId";
  @SerializedName(SERIALIZED_NAME_CLIENT_ID)
  private String clientId;

  public static final String SERIALIZED_NAME_CLIENT_SECRET = "clientSecret";
  @SerializedName(SERIALIZED_NAME_CLIENT_SECRET)
  private String clientSecret;

  public static final String SERIALIZED_NAME_ACCESS_TOKEN_VALIDITY_SECONDS = "accessTokenValiditySeconds";
  @SerializedName(SERIALIZED_NAME_ACCESS_TOKEN_VALIDITY_SECONDS)
  private Integer accessTokenValiditySeconds;

  public static final String SERIALIZED_NAME_REFRESH_TOKEN_VALIDITY_SECONDS = "refreshTokenValiditySeconds";
  @SerializedName(SERIALIZED_NAME_REFRESH_TOKEN_VALIDITY_SECONDS)
  private Integer refreshTokenValiditySeconds;

  public AddClientRequestClient() {
  }

  public AddClientRequestClient clientId(String clientId) {
    
    this.clientId = clientId;
    return this;
  }

   /**
   * ClientId
   * @return clientId
  **/
  @javax.annotation.Nonnull
  public String getClientId() {
    return clientId;
  }


  public void setClientId(String clientId) {
    this.clientId = clientId;
  }


  public AddClientRequestClient clientSecret(String clientSecret) {
    
    this.clientSecret = clientSecret;
    return this;
  }

   /**
   * ClientSecret
   * @return clientSecret
  **/
  @javax.annotation.Nullable
  public String getClientSecret() {
    return clientSecret;
  }


  public void setClientSecret(String clientSecret) {
    this.clientSecret = clientSecret;
  }


  public AddClientRequestClient accessTokenValiditySeconds(Integer accessTokenValiditySeconds) {
    
    this.accessTokenValiditySeconds = accessTokenValiditySeconds;
    return this;
  }

   /**
   * Length of time accessToken is valid in seconds.
   * @return accessTokenValiditySeconds
  **/
  @javax.annotation.Nullable
  public Integer getAccessTokenValiditySeconds() {
    return accessTokenValiditySeconds;
  }


  public void setAccessTokenValiditySeconds(Integer accessTokenValiditySeconds) {
    this.accessTokenValiditySeconds = accessTokenValiditySeconds;
  }


  public AddClientRequestClient refreshTokenValiditySeconds(Integer refreshTokenValiditySeconds) {
    
    this.refreshTokenValiditySeconds = refreshTokenValiditySeconds;
    return this;
  }

   /**
   * Length of time refreshToken is valid in seconds.
   * @return refreshTokenValiditySeconds
  **/
  @javax.annotation.Nullable
  public Integer getRefreshTokenValiditySeconds() {
    return refreshTokenValiditySeconds;
  }


  public void setRefreshTokenValiditySeconds(Integer refreshTokenValiditySeconds) {
    this.refreshTokenValiditySeconds = refreshTokenValiditySeconds;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    AddClientRequestClient addClientRequestClient = (AddClientRequestClient) o;
    return Objects.equals(this.clientId, addClientRequestClient.clientId) &&
        Objects.equals(this.clientSecret, addClientRequestClient.clientSecret) &&
        Objects.equals(this.accessTokenValiditySeconds, addClientRequestClient.accessTokenValiditySeconds) &&
        Objects.equals(this.refreshTokenValiditySeconds, addClientRequestClient.refreshTokenValiditySeconds);
  }

  @Override
  public int hashCode() {
    return Objects.hash(clientId, clientSecret, accessTokenValiditySeconds, refreshTokenValiditySeconds);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class AddClientRequestClient {\n");
    sb.append("    clientId: ").append(toIndentedString(clientId)).append("\n");
    sb.append("    clientSecret: ").append(toIndentedString(clientSecret)).append("\n");
    sb.append("    accessTokenValiditySeconds: ").append(toIndentedString(accessTokenValiditySeconds)).append("\n");
    sb.append("    refreshTokenValiditySeconds: ").append(toIndentedString(refreshTokenValiditySeconds)).append("\n");
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
    openapiFields.add("clientId");
    openapiFields.add("clientSecret");
    openapiFields.add("accessTokenValiditySeconds");
    openapiFields.add("refreshTokenValiditySeconds");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
    openapiRequiredFields.add("clientId");
    openapiRequiredFields.add("accessTokenValiditySeconds");
    openapiRequiredFields.add("refreshTokenValiditySeconds");
  }

 /**
  * Validates the JSON Element and throws an exception if issues found
  *
  * @param jsonElement JSON Element
  * @throws IOException if the JSON Element is invalid with respect to AddClientRequestClient
  */
  public static void validateJsonElement(JsonElement jsonElement) throws IOException {
      if (jsonElement == null) {
        if (!AddClientRequestClient.openapiRequiredFields.isEmpty()) { // has required fields but JSON element is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in AddClientRequestClient is not found in the empty JSON string", AddClientRequestClient.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonElement.getAsJsonObject().entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!AddClientRequestClient.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `AddClientRequestClient` properties. JSON: %s", entry.getKey(), jsonElement.toString()));
        }
      }

      // check to make sure all required properties/fields are present in the JSON string
      for (String requiredField : AddClientRequestClient.openapiRequiredFields) {
        if (jsonElement.getAsJsonObject().get(requiredField) == null) {
          throw new IllegalArgumentException(String.format("The required field `%s` is not found in the JSON string: %s", requiredField, jsonElement.toString()));
        }
      }
        JsonObject jsonObj = jsonElement.getAsJsonObject();
      if (!jsonObj.get("clientId").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `clientId` to be a primitive type in the JSON string but got `%s`", jsonObj.get("clientId").toString()));
      }
      if ((jsonObj.get("clientSecret") != null && !jsonObj.get("clientSecret").isJsonNull()) && !jsonObj.get("clientSecret").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `clientSecret` to be a primitive type in the JSON string but got `%s`", jsonObj.get("clientSecret").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!AddClientRequestClient.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'AddClientRequestClient' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<AddClientRequestClient> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(AddClientRequestClient.class));

       return (TypeAdapter<T>) new TypeAdapter<AddClientRequestClient>() {
           @Override
           public void write(JsonWriter out, AddClientRequestClient value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public AddClientRequestClient read(JsonReader in) throws IOException {
             JsonElement jsonElement = elementAdapter.read(in);
             validateJsonElement(jsonElement);
             return thisAdapter.fromJsonTree(jsonElement);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of AddClientRequestClient given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of AddClientRequestClient
  * @throws IOException if the JSON string is invalid with respect to AddClientRequestClient
  */
  public static AddClientRequestClient fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, AddClientRequestClient.class);
  }

 /**
  * Convert an instance of AddClientRequestClient to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

