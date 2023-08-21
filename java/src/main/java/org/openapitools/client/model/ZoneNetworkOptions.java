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
import org.openapitools.client.model.ZoneNetworkOptionsNetworkTypes;
import org.openapitools.client.model.ZoneNetworkOptionsNetworks;

/**
 * ZoneNetworkOptions
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ZoneNetworkOptions {
  public static final String SERIALIZED_NAME_NETWORKS = "networks";
  @SerializedName(SERIALIZED_NAME_NETWORKS)
  private List<ZoneNetworkOptionsNetworks> networks = null;

  public static final String SERIALIZED_NAME_NETWORK_GROUPS = "networkGroups";
  @SerializedName(SERIALIZED_NAME_NETWORK_GROUPS)
  private List<Object> networkGroups = null;

  public static final String SERIALIZED_NAME_NETWORK_TYPES = "networkTypes";
  @SerializedName(SERIALIZED_NAME_NETWORK_TYPES)
  private List<ZoneNetworkOptionsNetworkTypes> networkTypes = null;

  public static final String SERIALIZED_NAME_NETWORK_SUBNETS = "networkSubnets";
  @SerializedName(SERIALIZED_NAME_NETWORK_SUBNETS)
  private List<Object> networkSubnets = null;

  public static final String SERIALIZED_NAME_HAS_NETWORKS = "hasNetworks";
  @SerializedName(SERIALIZED_NAME_HAS_NETWORKS)
  private Boolean hasNetworks;

  public static final String SERIALIZED_NAME_MAX_NETWORKS = "maxNetworks";
  @SerializedName(SERIALIZED_NAME_MAX_NETWORKS)
  private Long maxNetworks;

  public static final String SERIALIZED_NAME_ENABLE_NETWORK_TYPE_SELECTION = "enableNetworkTypeSelection";
  @SerializedName(SERIALIZED_NAME_ENABLE_NETWORK_TYPE_SELECTION)
  private String enableNetworkTypeSelection;

  public static final String SERIALIZED_NAME_SUPPORTS_NETWORK_SELECTION = "supportsNetworkSelection";
  @SerializedName(SERIALIZED_NAME_SUPPORTS_NETWORK_SELECTION)
  private Boolean supportsNetworkSelection;


  public ZoneNetworkOptions networks(List<ZoneNetworkOptionsNetworks> networks) {
    
    this.networks = networks;
    return this;
  }

  public ZoneNetworkOptions addNetworksItem(ZoneNetworkOptionsNetworks networksItem) {
    if (this.networks == null) {
      this.networks = new ArrayList<ZoneNetworkOptionsNetworks>();
    }
    this.networks.add(networksItem);
    return this;
  }

   /**
   * Get networks
   * @return networks
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<ZoneNetworkOptionsNetworks> getNetworks() {
    return networks;
  }


  public void setNetworks(List<ZoneNetworkOptionsNetworks> networks) {
    this.networks = networks;
  }


  public ZoneNetworkOptions networkGroups(List<Object> networkGroups) {
    
    this.networkGroups = networkGroups;
    return this;
  }

  public ZoneNetworkOptions addNetworkGroupsItem(Object networkGroupsItem) {
    if (this.networkGroups == null) {
      this.networkGroups = new ArrayList<Object>();
    }
    this.networkGroups.add(networkGroupsItem);
    return this;
  }

   /**
   * Get networkGroups
   * @return networkGroups
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Object> getNetworkGroups() {
    return networkGroups;
  }


  public void setNetworkGroups(List<Object> networkGroups) {
    this.networkGroups = networkGroups;
  }


  public ZoneNetworkOptions networkTypes(List<ZoneNetworkOptionsNetworkTypes> networkTypes) {
    
    this.networkTypes = networkTypes;
    return this;
  }

  public ZoneNetworkOptions addNetworkTypesItem(ZoneNetworkOptionsNetworkTypes networkTypesItem) {
    if (this.networkTypes == null) {
      this.networkTypes = new ArrayList<ZoneNetworkOptionsNetworkTypes>();
    }
    this.networkTypes.add(networkTypesItem);
    return this;
  }

   /**
   * Get networkTypes
   * @return networkTypes
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<ZoneNetworkOptionsNetworkTypes> getNetworkTypes() {
    return networkTypes;
  }


  public void setNetworkTypes(List<ZoneNetworkOptionsNetworkTypes> networkTypes) {
    this.networkTypes = networkTypes;
  }


  public ZoneNetworkOptions networkSubnets(List<Object> networkSubnets) {
    
    this.networkSubnets = networkSubnets;
    return this;
  }

  public ZoneNetworkOptions addNetworkSubnetsItem(Object networkSubnetsItem) {
    if (this.networkSubnets == null) {
      this.networkSubnets = new ArrayList<Object>();
    }
    this.networkSubnets.add(networkSubnetsItem);
    return this;
  }

   /**
   * Get networkSubnets
   * @return networkSubnets
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Object> getNetworkSubnets() {
    return networkSubnets;
  }


  public void setNetworkSubnets(List<Object> networkSubnets) {
    this.networkSubnets = networkSubnets;
  }


  public ZoneNetworkOptions hasNetworks(Boolean hasNetworks) {
    
    this.hasNetworks = hasNetworks;
    return this;
  }

   /**
   * Get hasNetworks
   * @return hasNetworks
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getHasNetworks() {
    return hasNetworks;
  }


  public void setHasNetworks(Boolean hasNetworks) {
    this.hasNetworks = hasNetworks;
  }


  public ZoneNetworkOptions maxNetworks(Long maxNetworks) {
    
    this.maxNetworks = maxNetworks;
    return this;
  }

   /**
   * Get maxNetworks
   * @return maxNetworks
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getMaxNetworks() {
    return maxNetworks;
  }


  public void setMaxNetworks(Long maxNetworks) {
    this.maxNetworks = maxNetworks;
  }


  public ZoneNetworkOptions enableNetworkTypeSelection(String enableNetworkTypeSelection) {
    
    this.enableNetworkTypeSelection = enableNetworkTypeSelection;
    return this;
  }

   /**
   * Get enableNetworkTypeSelection
   * @return enableNetworkTypeSelection
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getEnableNetworkTypeSelection() {
    return enableNetworkTypeSelection;
  }


  public void setEnableNetworkTypeSelection(String enableNetworkTypeSelection) {
    this.enableNetworkTypeSelection = enableNetworkTypeSelection;
  }


  public ZoneNetworkOptions supportsNetworkSelection(Boolean supportsNetworkSelection) {
    
    this.supportsNetworkSelection = supportsNetworkSelection;
    return this;
  }

   /**
   * Get supportsNetworkSelection
   * @return supportsNetworkSelection
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getSupportsNetworkSelection() {
    return supportsNetworkSelection;
  }


  public void setSupportsNetworkSelection(Boolean supportsNetworkSelection) {
    this.supportsNetworkSelection = supportsNetworkSelection;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ZoneNetworkOptions zoneNetworkOptions = (ZoneNetworkOptions) o;
    return Objects.equals(this.networks, zoneNetworkOptions.networks) &&
        Objects.equals(this.networkGroups, zoneNetworkOptions.networkGroups) &&
        Objects.equals(this.networkTypes, zoneNetworkOptions.networkTypes) &&
        Objects.equals(this.networkSubnets, zoneNetworkOptions.networkSubnets) &&
        Objects.equals(this.hasNetworks, zoneNetworkOptions.hasNetworks) &&
        Objects.equals(this.maxNetworks, zoneNetworkOptions.maxNetworks) &&
        Objects.equals(this.enableNetworkTypeSelection, zoneNetworkOptions.enableNetworkTypeSelection) &&
        Objects.equals(this.supportsNetworkSelection, zoneNetworkOptions.supportsNetworkSelection);
  }

  @Override
  public int hashCode() {
    return Objects.hash(networks, networkGroups, networkTypes, networkSubnets, hasNetworks, maxNetworks, enableNetworkTypeSelection, supportsNetworkSelection);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ZoneNetworkOptions {\n");
    sb.append("    networks: ").append(toIndentedString(networks)).append("\n");
    sb.append("    networkGroups: ").append(toIndentedString(networkGroups)).append("\n");
    sb.append("    networkTypes: ").append(toIndentedString(networkTypes)).append("\n");
    sb.append("    networkSubnets: ").append(toIndentedString(networkSubnets)).append("\n");
    sb.append("    hasNetworks: ").append(toIndentedString(hasNetworks)).append("\n");
    sb.append("    maxNetworks: ").append(toIndentedString(maxNetworks)).append("\n");
    sb.append("    enableNetworkTypeSelection: ").append(toIndentedString(enableNetworkTypeSelection)).append("\n");
    sb.append("    supportsNetworkSelection: ").append(toIndentedString(supportsNetworkSelection)).append("\n");
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
