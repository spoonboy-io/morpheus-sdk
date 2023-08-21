/*
 * Morpheus API
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.2.1
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
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.client.model.ApiPriceSetsPriceSetZone;
import org.openapitools.client.model.ApiPriceSetsPriceSetZonePool;

/**
 * ApiPriceSetsIdPriceSet
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ApiPriceSetsIdPriceSet {
  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_CODE = "code";
  @SerializedName(SERIALIZED_NAME_CODE)
  private String code;

  public static final String SERIALIZED_NAME_REGION_CODE = "regionCode";
  @SerializedName(SERIALIZED_NAME_REGION_CODE)
  private String regionCode;

  public static final String SERIALIZED_NAME_ZONE = "zone";
  @SerializedName(SERIALIZED_NAME_ZONE)
  private ApiPriceSetsPriceSetZone zone;

  public static final String SERIALIZED_NAME_ZONE_POOL = "zonePool";
  @SerializedName(SERIALIZED_NAME_ZONE_POOL)
  private ApiPriceSetsPriceSetZonePool zonePool;

  /**
   * Price Unit
   */
  @JsonAdapter(PriceUnitEnum.Adapter.class)
  public enum PriceUnitEnum {
    MINUTE("minute"),
    
    HOUR("hour"),
    
    DAY("day"),
    
    MONTH("month"),
    
    YEAR("year"),
    
    TWO_YEAR("two year"),
    
    THREE_YEAR("three year"),
    
    FOUR_YEAR("four year"),
    
    FIVE_YEAR("five year");

    private String value;

    PriceUnitEnum(String value) {
      this.value = value;
    }

    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static PriceUnitEnum fromValue(String value) {
      for (PriceUnitEnum b : PriceUnitEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }

    public static class Adapter extends TypeAdapter<PriceUnitEnum> {
      @Override
      public void write(final JsonWriter jsonWriter, final PriceUnitEnum enumeration) throws IOException {
        jsonWriter.value(enumeration.getValue());
      }

      @Override
      public PriceUnitEnum read(final JsonReader jsonReader) throws IOException {
        String value =  jsonReader.nextString();
        return PriceUnitEnum.fromValue(value);
      }
    }
  }

  public static final String SERIALIZED_NAME_PRICE_UNIT = "priceUnit";
  @SerializedName(SERIALIZED_NAME_PRICE_UNIT)
  private PriceUnitEnum priceUnit;

  /**
   * Price set type
   */
  @JsonAdapter(TypeEnum.Adapter.class)
  public enum TypeEnum {
    FIXED("fixed"),
    
    COMPUTE_PLUS_STORAGE("compute_plus_storage"),
    
    COMPONENT("component"),
    
    LOAD_BALANCER("load_balancer"),
    
    SNAPSHOT("snapshot"),
    
    VIRTUAL_IMAGE("virtual_image"),
    
    SOFTWARE_OR_SERVICE("software_or_service");

    private String value;

    TypeEnum(String value) {
      this.value = value;
    }

    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static TypeEnum fromValue(String value) {
      for (TypeEnum b : TypeEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }

    public static class Adapter extends TypeAdapter<TypeEnum> {
      @Override
      public void write(final JsonWriter jsonWriter, final TypeEnum enumeration) throws IOException {
        jsonWriter.value(enumeration.getValue());
      }

      @Override
      public TypeEnum read(final JsonReader jsonReader) throws IOException {
        String value =  jsonReader.nextString();
        return TypeEnum.fromValue(value);
      }
    }
  }

  public static final String SERIALIZED_NAME_TYPE = "type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private TypeEnum type;

  public static final String SERIALIZED_NAME_PRICES = "prices";
  @SerializedName(SERIALIZED_NAME_PRICES)
  private List<Long> prices = null;


  public ApiPriceSetsIdPriceSet name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Price set name
   * @return name
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "testName", value = "Price set name")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public ApiPriceSetsIdPriceSet code(String code) {
    
    this.code = code;
    return this;
  }

   /**
   * Price set code. Must be unique.
   * @return code
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "priceSet1", value = "Price set code. Must be unique.")

  public String getCode() {
    return code;
  }


  public void setCode(String code) {
    this.code = code;
  }


  public ApiPriceSetsIdPriceSet regionCode(String regionCode) {
    
    this.regionCode = regionCode;
    return this;
  }

   /**
   * Price set region code
   * @return regionCode
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "region.code.1", value = "Price set region code")

  public String getRegionCode() {
    return regionCode;
  }


  public void setRegionCode(String regionCode) {
    this.regionCode = regionCode;
  }


  public ApiPriceSetsIdPriceSet zone(ApiPriceSetsPriceSetZone zone) {
    
    this.zone = zone;
    return this;
  }

   /**
   * Get zone
   * @return zone
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ApiPriceSetsPriceSetZone getZone() {
    return zone;
  }


  public void setZone(ApiPriceSetsPriceSetZone zone) {
    this.zone = zone;
  }


  public ApiPriceSetsIdPriceSet zonePool(ApiPriceSetsPriceSetZonePool zonePool) {
    
    this.zonePool = zonePool;
    return this;
  }

   /**
   * Get zonePool
   * @return zonePool
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ApiPriceSetsPriceSetZonePool getZonePool() {
    return zonePool;
  }


  public void setZonePool(ApiPriceSetsPriceSetZonePool zonePool) {
    this.zonePool = zonePool;
  }


  public ApiPriceSetsIdPriceSet priceUnit(PriceUnitEnum priceUnit) {
    
    this.priceUnit = priceUnit;
    return this;
  }

   /**
   * Price Unit
   * @return priceUnit
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Price Unit")

  public PriceUnitEnum getPriceUnit() {
    return priceUnit;
  }


  public void setPriceUnit(PriceUnitEnum priceUnit) {
    this.priceUnit = priceUnit;
  }


  public ApiPriceSetsIdPriceSet type(TypeEnum type) {
    
    this.type = type;
    return this;
  }

   /**
   * Price set type
   * @return type
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Price set type")

  public TypeEnum getType() {
    return type;
  }


  public void setType(TypeEnum type) {
    this.type = type;
  }


  public ApiPriceSetsIdPriceSet prices(List<Long> prices) {
    
    this.prices = prices;
    return this;
  }

  public ApiPriceSetsIdPriceSet addPricesItem(Long pricesItem) {
    if (this.prices == null) {
      this.prices = new ArrayList<Long>();
    }
    this.prices.add(pricesItem);
    return this;
  }

   /**
   * Get prices
   * @return prices
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Long> getPrices() {
    return prices;
  }


  public void setPrices(List<Long> prices) {
    this.prices = prices;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ApiPriceSetsIdPriceSet apiPriceSetsIdPriceSet = (ApiPriceSetsIdPriceSet) o;
    return Objects.equals(this.name, apiPriceSetsIdPriceSet.name) &&
        Objects.equals(this.code, apiPriceSetsIdPriceSet.code) &&
        Objects.equals(this.regionCode, apiPriceSetsIdPriceSet.regionCode) &&
        Objects.equals(this.zone, apiPriceSetsIdPriceSet.zone) &&
        Objects.equals(this.zonePool, apiPriceSetsIdPriceSet.zonePool) &&
        Objects.equals(this.priceUnit, apiPriceSetsIdPriceSet.priceUnit) &&
        Objects.equals(this.type, apiPriceSetsIdPriceSet.type) &&
        Objects.equals(this.prices, apiPriceSetsIdPriceSet.prices);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, code, regionCode, zone, zonePool, priceUnit, type, prices);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ApiPriceSetsIdPriceSet {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    code: ").append(toIndentedString(code)).append("\n");
    sb.append("    regionCode: ").append(toIndentedString(regionCode)).append("\n");
    sb.append("    zone: ").append(toIndentedString(zone)).append("\n");
    sb.append("    zonePool: ").append(toIndentedString(zonePool)).append("\n");
    sb.append("    priceUnit: ").append(toIndentedString(priceUnit)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    prices: ").append(toIndentedString(prices)).append("\n");
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

}

