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
import org.openapitools.client.model.AppStateInputDataInner;
import org.openapitools.client.model.AppStateInputProvidersInner;
import org.openapitools.client.model.AppStateInputVariablesInner;

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
 * AppStateInput
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-17T13:37:08.123540Z[Etc/UTC]")
public class AppStateInput {
  public static final String SERIALIZED_NAME_VARIABLES = "variables";
  @SerializedName(SERIALIZED_NAME_VARIABLES)
  private List<AppStateInputVariablesInner> variables;

  public static final String SERIALIZED_NAME_PROVIDERS = "providers";
  @SerializedName(SERIALIZED_NAME_PROVIDERS)
  private List<AppStateInputProvidersInner> providers;

  public static final String SERIALIZED_NAME_DATA = "data";
  @SerializedName(SERIALIZED_NAME_DATA)
  private List<AppStateInputDataInner> data;

  public AppStateInput() {
  }

  public AppStateInput variables(List<AppStateInputVariablesInner> variables) {
    
    this.variables = variables;
    return this;
  }

  public AppStateInput addVariablesItem(AppStateInputVariablesInner variablesItem) {
    if (this.variables == null) {
      this.variables = new ArrayList<>();
    }
    this.variables.add(variablesItem);
    return this;
  }

   /**
   * Get variables
   * @return variables
  **/
  @javax.annotation.Nullable
  public List<AppStateInputVariablesInner> getVariables() {
    return variables;
  }


  public void setVariables(List<AppStateInputVariablesInner> variables) {
    this.variables = variables;
  }


  public AppStateInput providers(List<AppStateInputProvidersInner> providers) {
    
    this.providers = providers;
    return this;
  }

  public AppStateInput addProvidersItem(AppStateInputProvidersInner providersItem) {
    if (this.providers == null) {
      this.providers = new ArrayList<>();
    }
    this.providers.add(providersItem);
    return this;
  }

   /**
   * Get providers
   * @return providers
  **/
  @javax.annotation.Nullable
  public List<AppStateInputProvidersInner> getProviders() {
    return providers;
  }


  public void setProviders(List<AppStateInputProvidersInner> providers) {
    this.providers = providers;
  }


  public AppStateInput data(List<AppStateInputDataInner> data) {
    
    this.data = data;
    return this;
  }

  public AppStateInput addDataItem(AppStateInputDataInner dataItem) {
    if (this.data == null) {
      this.data = new ArrayList<>();
    }
    this.data.add(dataItem);
    return this;
  }

   /**
   * Get data
   * @return data
  **/
  @javax.annotation.Nullable
  public List<AppStateInputDataInner> getData() {
    return data;
  }


  public void setData(List<AppStateInputDataInner> data) {
    this.data = data;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    AppStateInput appStateInput = (AppStateInput) o;
    return Objects.equals(this.variables, appStateInput.variables) &&
        Objects.equals(this.providers, appStateInput.providers) &&
        Objects.equals(this.data, appStateInput.data);
  }

  @Override
  public int hashCode() {
    return Objects.hash(variables, providers, data);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class AppStateInput {\n");
    sb.append("    variables: ").append(toIndentedString(variables)).append("\n");
    sb.append("    providers: ").append(toIndentedString(providers)).append("\n");
    sb.append("    data: ").append(toIndentedString(data)).append("\n");
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
    openapiFields.add("variables");
    openapiFields.add("providers");
    openapiFields.add("data");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Element and throws an exception if issues found
  *
  * @param jsonElement JSON Element
  * @throws IOException if the JSON Element is invalid with respect to AppStateInput
  */
  public static void validateJsonElement(JsonElement jsonElement) throws IOException {
      if (jsonElement == null) {
        if (!AppStateInput.openapiRequiredFields.isEmpty()) { // has required fields but JSON element is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in AppStateInput is not found in the empty JSON string", AppStateInput.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonElement.getAsJsonObject().entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!AppStateInput.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `AppStateInput` properties. JSON: %s", entry.getKey(), jsonElement.toString()));
        }
      }
        JsonObject jsonObj = jsonElement.getAsJsonObject();
      if (jsonObj.get("variables") != null && !jsonObj.get("variables").isJsonNull()) {
        JsonArray jsonArrayvariables = jsonObj.getAsJsonArray("variables");
        if (jsonArrayvariables != null) {
          // ensure the json data is an array
          if (!jsonObj.get("variables").isJsonArray()) {
            throw new IllegalArgumentException(String.format("Expected the field `variables` to be an array in the JSON string but got `%s`", jsonObj.get("variables").toString()));
          }

          // validate the optional field `variables` (array)
          for (int i = 0; i < jsonArrayvariables.size(); i++) {
            AppStateInputVariablesInner.validateJsonElement(jsonArrayvariables.get(i));
          };
        }
      }
      if (jsonObj.get("providers") != null && !jsonObj.get("providers").isJsonNull()) {
        JsonArray jsonArrayproviders = jsonObj.getAsJsonArray("providers");
        if (jsonArrayproviders != null) {
          // ensure the json data is an array
          if (!jsonObj.get("providers").isJsonArray()) {
            throw new IllegalArgumentException(String.format("Expected the field `providers` to be an array in the JSON string but got `%s`", jsonObj.get("providers").toString()));
          }

          // validate the optional field `providers` (array)
          for (int i = 0; i < jsonArrayproviders.size(); i++) {
            AppStateInputProvidersInner.validateJsonElement(jsonArrayproviders.get(i));
          };
        }
      }
      if (jsonObj.get("data") != null && !jsonObj.get("data").isJsonNull()) {
        JsonArray jsonArraydata = jsonObj.getAsJsonArray("data");
        if (jsonArraydata != null) {
          // ensure the json data is an array
          if (!jsonObj.get("data").isJsonArray()) {
            throw new IllegalArgumentException(String.format("Expected the field `data` to be an array in the JSON string but got `%s`", jsonObj.get("data").toString()));
          }

          // validate the optional field `data` (array)
          for (int i = 0; i < jsonArraydata.size(); i++) {
            AppStateInputDataInner.validateJsonElement(jsonArraydata.get(i));
          };
        }
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!AppStateInput.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'AppStateInput' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<AppStateInput> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(AppStateInput.class));

       return (TypeAdapter<T>) new TypeAdapter<AppStateInput>() {
           @Override
           public void write(JsonWriter out, AppStateInput value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public AppStateInput read(JsonReader in) throws IOException {
             JsonElement jsonElement = elementAdapter.read(in);
             validateJsonElement(jsonElement);
             return thisAdapter.fromJsonTree(jsonElement);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of AppStateInput given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of AppStateInput
  * @throws IOException if the JSON string is invalid with respect to AppStateInput
  */
  public static AppStateInput fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, AppStateInput.class);
  }

 /**
  * Convert an instance of AppStateInput to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

