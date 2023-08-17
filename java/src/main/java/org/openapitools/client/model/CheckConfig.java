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



import java.io.IOException;
import java.lang.reflect.Type;
import java.util.logging.Level;
import java.util.logging.Logger;
import java.util.ArrayList;
import java.util.Collections;
import java.util.HashSet;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.JsonParseException;
import com.google.gson.TypeAdapter;
import com.google.gson.TypeAdapterFactory;
import com.google.gson.reflect.TypeToken;
import com.google.gson.JsonPrimitive;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import com.google.gson.JsonDeserializationContext;
import com.google.gson.JsonDeserializer;
import com.google.gson.JsonSerializationContext;
import com.google.gson.JsonSerializer;
import com.google.gson.JsonElement;
import com.google.gson.JsonObject;
import com.google.gson.JsonArray;
import com.google.gson.JsonParseException;

import org.openapitools.client.JSON;

@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-17T13:37:08.123540Z[Etc/UTC]")
public class CheckConfig extends AbstractOpenApiSchema {
    private static final Logger log = Logger.getLogger(CheckConfig.class.getName());

    public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
        @SuppressWarnings("unchecked")
        @Override
        public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
            if (!CheckConfig.class.isAssignableFrom(type.getRawType())) {
                return null; // this class only serializes 'CheckConfig' and its subtypes
            }
            final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
            final TypeAdapter<Object> adapterObject = gson.getDelegateAdapter(this, TypeToken.get(Object.class));
            final TypeAdapter<Object> adapterObject = gson.getDelegateAdapter(this, TypeToken.get(Object.class));
            final TypeAdapter<Object> adapterObject = gson.getDelegateAdapter(this, TypeToken.get(Object.class));
            final TypeAdapter<Object> adapterObject = gson.getDelegateAdapter(this, TypeToken.get(Object.class));
            final TypeAdapter<Object> adapterObject = gson.getDelegateAdapter(this, TypeToken.get(Object.class));
            final TypeAdapter<Object> adapterObject = gson.getDelegateAdapter(this, TypeToken.get(Object.class));

            return (TypeAdapter<T>) new TypeAdapter<CheckConfig>() {
                @Override
                public void write(JsonWriter out, CheckConfig value) throws IOException {
                    if (value == null || value.getActualInstance() == null) {
                        elementAdapter.write(out, null);
                        return;
                    }

                    // check if the actual instance is of the type `Object`
                    if (value.getActualInstance() instanceof Object) {
                      JsonPrimitive primitive = adapterObject.toJsonTree((Object)value.getActualInstance()).getAsJsonPrimitive();
                      elementAdapter.write(out, primitive);
                      return;
                    }
                    // check if the actual instance is of the type `Object`
                    if (value.getActualInstance() instanceof Object) {
                      JsonPrimitive primitive = adapterObject.toJsonTree((Object)value.getActualInstance()).getAsJsonPrimitive();
                      elementAdapter.write(out, primitive);
                      return;
                    }
                    // check if the actual instance is of the type `Object`
                    if (value.getActualInstance() instanceof Object) {
                      JsonPrimitive primitive = adapterObject.toJsonTree((Object)value.getActualInstance()).getAsJsonPrimitive();
                      elementAdapter.write(out, primitive);
                      return;
                    }
                    // check if the actual instance is of the type `Object`
                    if (value.getActualInstance() instanceof Object) {
                      JsonPrimitive primitive = adapterObject.toJsonTree((Object)value.getActualInstance()).getAsJsonPrimitive();
                      elementAdapter.write(out, primitive);
                      return;
                    }
                    // check if the actual instance is of the type `Object`
                    if (value.getActualInstance() instanceof Object) {
                      JsonPrimitive primitive = adapterObject.toJsonTree((Object)value.getActualInstance()).getAsJsonPrimitive();
                      elementAdapter.write(out, primitive);
                      return;
                    }
                    // check if the actual instance is of the type `Object`
                    if (value.getActualInstance() instanceof Object) {
                      JsonPrimitive primitive = adapterObject.toJsonTree((Object)value.getActualInstance()).getAsJsonPrimitive();
                      elementAdapter.write(out, primitive);
                      return;
                    }
                    throw new IOException("Failed to serialize as the type doesn't match anyOf schemae: ERRORUNKNOWN");
                }

                @Override
                public CheckConfig read(JsonReader in) throws IOException {
                    Object deserialized = null;
                    JsonElement jsonElement = elementAdapter.read(in);

                    ArrayList<String> errorMessages = new ArrayList<>();
                    TypeAdapter actualAdapter = elementAdapter;

                    // deserialize Object
                    try {
                      // validate the JSON object to see if any exception is thrown
                      if(!jsonElement.getAsJsonPrimitive().isNumber()) {
                        throw new IllegalArgumentException(String.format("Expected json element to be of type Number in the JSON string but got `%s`", jsonElement.toString()));
                      }
                      actualAdapter = adapterObject;
                      CheckConfig ret = new CheckConfig();
                      ret.setActualInstance(actualAdapter.fromJsonTree(jsonElement));
                      return ret;
                    } catch (Exception e) {
                      // deserialization failed, continue
                      errorMessages.add(String.format("Deserialization for Object failed with `%s`.", e.getMessage()));
                      log.log(Level.FINER, "Input data does not match schema 'Object'", e);
                    }
                    // deserialize Object
                    try {
                      // validate the JSON object to see if any exception is thrown
                      if(!jsonElement.getAsJsonPrimitive().isNumber()) {
                        throw new IllegalArgumentException(String.format("Expected json element to be of type Number in the JSON string but got `%s`", jsonElement.toString()));
                      }
                      actualAdapter = adapterObject;
                      CheckConfig ret = new CheckConfig();
                      ret.setActualInstance(actualAdapter.fromJsonTree(jsonElement));
                      return ret;
                    } catch (Exception e) {
                      // deserialization failed, continue
                      errorMessages.add(String.format("Deserialization for Object failed with `%s`.", e.getMessage()));
                      log.log(Level.FINER, "Input data does not match schema 'Object'", e);
                    }
                    // deserialize Object
                    try {
                      // validate the JSON object to see if any exception is thrown
                      if(!jsonElement.getAsJsonPrimitive().isNumber()) {
                        throw new IllegalArgumentException(String.format("Expected json element to be of type Number in the JSON string but got `%s`", jsonElement.toString()));
                      }
                      actualAdapter = adapterObject;
                      CheckConfig ret = new CheckConfig();
                      ret.setActualInstance(actualAdapter.fromJsonTree(jsonElement));
                      return ret;
                    } catch (Exception e) {
                      // deserialization failed, continue
                      errorMessages.add(String.format("Deserialization for Object failed with `%s`.", e.getMessage()));
                      log.log(Level.FINER, "Input data does not match schema 'Object'", e);
                    }
                    // deserialize Object
                    try {
                      // validate the JSON object to see if any exception is thrown
                      if(!jsonElement.getAsJsonPrimitive().isNumber()) {
                        throw new IllegalArgumentException(String.format("Expected json element to be of type Number in the JSON string but got `%s`", jsonElement.toString()));
                      }
                      actualAdapter = adapterObject;
                      CheckConfig ret = new CheckConfig();
                      ret.setActualInstance(actualAdapter.fromJsonTree(jsonElement));
                      return ret;
                    } catch (Exception e) {
                      // deserialization failed, continue
                      errorMessages.add(String.format("Deserialization for Object failed with `%s`.", e.getMessage()));
                      log.log(Level.FINER, "Input data does not match schema 'Object'", e);
                    }
                    // deserialize Object
                    try {
                      // validate the JSON object to see if any exception is thrown
                      if(!jsonElement.getAsJsonPrimitive().isNumber()) {
                        throw new IllegalArgumentException(String.format("Expected json element to be of type Number in the JSON string but got `%s`", jsonElement.toString()));
                      }
                      actualAdapter = adapterObject;
                      CheckConfig ret = new CheckConfig();
                      ret.setActualInstance(actualAdapter.fromJsonTree(jsonElement));
                      return ret;
                    } catch (Exception e) {
                      // deserialization failed, continue
                      errorMessages.add(String.format("Deserialization for Object failed with `%s`.", e.getMessage()));
                      log.log(Level.FINER, "Input data does not match schema 'Object'", e);
                    }
                    // deserialize Object
                    try {
                      // validate the JSON object to see if any exception is thrown
                      if(!jsonElement.getAsJsonPrimitive().isNumber()) {
                        throw new IllegalArgumentException(String.format("Expected json element to be of type Number in the JSON string but got `%s`", jsonElement.toString()));
                      }
                      actualAdapter = adapterObject;
                      CheckConfig ret = new CheckConfig();
                      ret.setActualInstance(actualAdapter.fromJsonTree(jsonElement));
                      return ret;
                    } catch (Exception e) {
                      // deserialization failed, continue
                      errorMessages.add(String.format("Deserialization for Object failed with `%s`.", e.getMessage()));
                      log.log(Level.FINER, "Input data does not match schema 'Object'", e);
                    }

                    throw new IOException(String.format("Failed deserialization for CheckConfig: no class matches result, expected at least 1. Detailed failure message for anyOf schemas: %s. JSON: %s", errorMessages, jsonElement.toString()));
                }
            }.nullSafe();
        }
    }

    // store a list of schema names defined in anyOf
    public static final Map<String, Class<?>> schemas = new HashMap<String, Class<?>>();

    public CheckConfig() {
        super("anyOf", Boolean.FALSE);
    }

    public CheckConfig(ERRORUNKNOWN o) {
        super("anyOf", Boolean.FALSE);
        setActualInstance(o);
    }

    static {
        schemas.put("Object", Object.class);
        schemas.put("Object", Object.class);
        schemas.put("Object", Object.class);
        schemas.put("Object", Object.class);
        schemas.put("Object", Object.class);
        schemas.put("Object", Object.class);
    }

    @Override
    public Map<String, Class<?>> getSchemas() {
        return CheckConfig.schemas;
    }

    /**
     * Set the instance that matches the anyOf child schema, check
     * the instance parameter is valid against the anyOf child schemas:
     * ERRORUNKNOWN
     *
     * It could be an instance of the 'anyOf' schemas.
     */
    @Override
    public void setActualInstance(Object instance) {
        if (instance instanceof Object) {
            super.setActualInstance(instance);
            return;
        }

        if (instance instanceof Object) {
            super.setActualInstance(instance);
            return;
        }

        if (instance instanceof Object) {
            super.setActualInstance(instance);
            return;
        }

        if (instance instanceof Object) {
            super.setActualInstance(instance);
            return;
        }

        if (instance instanceof Object) {
            super.setActualInstance(instance);
            return;
        }

        if (instance instanceof Object) {
            super.setActualInstance(instance);
            return;
        }

        throw new RuntimeException("Invalid instance type. Must be ERRORUNKNOWN");
    }

    /**
     * Get the actual instance, which can be the following:
     * ERRORUNKNOWN
     *
     * @return The actual instance (ERRORUNKNOWN)
     */
    @Override
    public Object getActualInstance() {
        return super.getActualInstance();
    }

    /**
     * Get the actual instance of `Object`. If the actual instance is not `Object`,
     * the ClassCastException will be thrown.
     *
     * @return The actual instance of `Object`
     * @throws ClassCastException if the instance is not `Object`
     */
    public Object getObject() throws ClassCastException {
        return (Object)super.getActualInstance();
    }
    /**
     * Get the actual instance of `Object`. If the actual instance is not `Object`,
     * the ClassCastException will be thrown.
     *
     * @return The actual instance of `Object`
     * @throws ClassCastException if the instance is not `Object`
     */
    public Object getObject() throws ClassCastException {
        return (Object)super.getActualInstance();
    }
    /**
     * Get the actual instance of `Object`. If the actual instance is not `Object`,
     * the ClassCastException will be thrown.
     *
     * @return The actual instance of `Object`
     * @throws ClassCastException if the instance is not `Object`
     */
    public Object getObject() throws ClassCastException {
        return (Object)super.getActualInstance();
    }
    /**
     * Get the actual instance of `Object`. If the actual instance is not `Object`,
     * the ClassCastException will be thrown.
     *
     * @return The actual instance of `Object`
     * @throws ClassCastException if the instance is not `Object`
     */
    public Object getObject() throws ClassCastException {
        return (Object)super.getActualInstance();
    }
    /**
     * Get the actual instance of `Object`. If the actual instance is not `Object`,
     * the ClassCastException will be thrown.
     *
     * @return The actual instance of `Object`
     * @throws ClassCastException if the instance is not `Object`
     */
    public Object getObject() throws ClassCastException {
        return (Object)super.getActualInstance();
    }
    /**
     * Get the actual instance of `Object`. If the actual instance is not `Object`,
     * the ClassCastException will be thrown.
     *
     * @return The actual instance of `Object`
     * @throws ClassCastException if the instance is not `Object`
     */
    public Object getObject() throws ClassCastException {
        return (Object)super.getActualInstance();
    }

 /**
  * Validates the JSON Element and throws an exception if issues found
  *
  * @param jsonElement JSON Element
  * @throws IOException if the JSON Element is invalid with respect to CheckConfig
  */
  public static void validateJsonElement(JsonElement jsonElement) throws IOException {
    // validate anyOf schemas one by one
    ArrayList<String> errorMessages = new ArrayList<>();
    // validate the json string with Object
    try {
      if(!jsonElement.getAsJsonPrimitive().isNumber()) {
        throw new IllegalArgumentException(String.format("Expected json element to be of type Number in the JSON string but got `%s`", jsonElement.toString()));
      }
      return;
    } catch (Exception e) {
      errorMessages.add(String.format("Deserialization for Object failed with `%s`.", e.getMessage()));
      // continue to the next one
    }
    // validate the json string with Object
    try {
      if(!jsonElement.getAsJsonPrimitive().isNumber()) {
        throw new IllegalArgumentException(String.format("Expected json element to be of type Number in the JSON string but got `%s`", jsonElement.toString()));
      }
      return;
    } catch (Exception e) {
      errorMessages.add(String.format("Deserialization for Object failed with `%s`.", e.getMessage()));
      // continue to the next one
    }
    // validate the json string with Object
    try {
      if(!jsonElement.getAsJsonPrimitive().isNumber()) {
        throw new IllegalArgumentException(String.format("Expected json element to be of type Number in the JSON string but got `%s`", jsonElement.toString()));
      }
      return;
    } catch (Exception e) {
      errorMessages.add(String.format("Deserialization for Object failed with `%s`.", e.getMessage()));
      // continue to the next one
    }
    // validate the json string with Object
    try {
      if(!jsonElement.getAsJsonPrimitive().isNumber()) {
        throw new IllegalArgumentException(String.format("Expected json element to be of type Number in the JSON string but got `%s`", jsonElement.toString()));
      }
      return;
    } catch (Exception e) {
      errorMessages.add(String.format("Deserialization for Object failed with `%s`.", e.getMessage()));
      // continue to the next one
    }
    // validate the json string with Object
    try {
      if(!jsonElement.getAsJsonPrimitive().isNumber()) {
        throw new IllegalArgumentException(String.format("Expected json element to be of type Number in the JSON string but got `%s`", jsonElement.toString()));
      }
      return;
    } catch (Exception e) {
      errorMessages.add(String.format("Deserialization for Object failed with `%s`.", e.getMessage()));
      // continue to the next one
    }
    // validate the json string with Object
    try {
      if(!jsonElement.getAsJsonPrimitive().isNumber()) {
        throw new IllegalArgumentException(String.format("Expected json element to be of type Number in the JSON string but got `%s`", jsonElement.toString()));
      }
      return;
    } catch (Exception e) {
      errorMessages.add(String.format("Deserialization for Object failed with `%s`.", e.getMessage()));
      // continue to the next one
    }
    throw new IOException(String.format("The JSON string is invalid for CheckConfig with anyOf schemas: ERRORUNKNOWN. no class match the result, expected at least 1. Detailed failure message for anyOf schemas: %s. JSON: %s", errorMessages, jsonElement.toString()));
    
  }

 /**
  * Create an instance of CheckConfig given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of CheckConfig
  * @throws IOException if the JSON string is invalid with respect to CheckConfig
  */
  public static CheckConfig fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, CheckConfig.class);
  }

 /**
  * Convert an instance of CheckConfig to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

