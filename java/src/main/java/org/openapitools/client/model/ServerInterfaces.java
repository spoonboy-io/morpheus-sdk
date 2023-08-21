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
import org.openapitools.client.model.InlineResponse20040AppDeployInstance;
import org.openapitools.client.model.InlineResponse20082LoadBalancerInstanceSslCert;
import org.openapitools.client.model.PriceSetVolumeType;

/**
 * ServerInterfaces
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ServerInterfaces {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_REF_TYPE = "refType";
  @SerializedName(SERIALIZED_NAME_REF_TYPE)
  private String refType;

  public static final String SERIALIZED_NAME_REF_ID = "refId";
  @SerializedName(SERIALIZED_NAME_REF_ID)
  private String refId;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_INTERNAL_ID = "internalId";
  @SerializedName(SERIALIZED_NAME_INTERNAL_ID)
  private String internalId;

  public static final String SERIALIZED_NAME_EXTERNAL_ID = "externalId";
  @SerializedName(SERIALIZED_NAME_EXTERNAL_ID)
  private String externalId;

  public static final String SERIALIZED_NAME_UNIQUE_ID = "uniqueId";
  @SerializedName(SERIALIZED_NAME_UNIQUE_ID)
  private String uniqueId;

  public static final String SERIALIZED_NAME_PUBLIC_IP_ADDRESS = "publicIpAddress";
  @SerializedName(SERIALIZED_NAME_PUBLIC_IP_ADDRESS)
  private String publicIpAddress;

  public static final String SERIALIZED_NAME_PUBLIC_IPV6_ADDRESS = "publicIpv6Address";
  @SerializedName(SERIALIZED_NAME_PUBLIC_IPV6_ADDRESS)
  private String publicIpv6Address;

  public static final String SERIALIZED_NAME_IP_ADDRESS = "ipAddress";
  @SerializedName(SERIALIZED_NAME_IP_ADDRESS)
  private String ipAddress;

  public static final String SERIALIZED_NAME_IPV6_ADDRESS = "ipv6Address";
  @SerializedName(SERIALIZED_NAME_IPV6_ADDRESS)
  private String ipv6Address;

  public static final String SERIALIZED_NAME_IP_SUBNET = "ipSubnet";
  @SerializedName(SERIALIZED_NAME_IP_SUBNET)
  private String ipSubnet;

  public static final String SERIALIZED_NAME_IPV6_SUBNET = "ipv6Subnet";
  @SerializedName(SERIALIZED_NAME_IPV6_SUBNET)
  private String ipv6Subnet;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_DHCP = "dhcp";
  @SerializedName(SERIALIZED_NAME_DHCP)
  private Boolean dhcp;

  public static final String SERIALIZED_NAME_ACTIVE = "active";
  @SerializedName(SERIALIZED_NAME_ACTIVE)
  private Boolean active;

  public static final String SERIALIZED_NAME_POOL_ASSIGNED = "poolAssigned";
  @SerializedName(SERIALIZED_NAME_POOL_ASSIGNED)
  private Boolean poolAssigned;

  public static final String SERIALIZED_NAME_PRIMARY_INTERFACE = "primaryInterface";
  @SerializedName(SERIALIZED_NAME_PRIMARY_INTERFACE)
  private Boolean primaryInterface;

  public static final String SERIALIZED_NAME_NETWORK = "network";
  @SerializedName(SERIALIZED_NAME_NETWORK)
  private InlineResponse20040AppDeployInstance network;

  public static final String SERIALIZED_NAME_SUBNET = "subnet";
  @SerializedName(SERIALIZED_NAME_SUBNET)
  private String subnet;

  public static final String SERIALIZED_NAME_NETWORK_GROUP = "networkGroup";
  @SerializedName(SERIALIZED_NAME_NETWORK_GROUP)
  private String networkGroup;

  public static final String SERIALIZED_NAME_NETWORK_POSITION = "networkPosition";
  @SerializedName(SERIALIZED_NAME_NETWORK_POSITION)
  private String networkPosition;

  public static final String SERIALIZED_NAME_NETWORK_POOL = "networkPool";
  @SerializedName(SERIALIZED_NAME_NETWORK_POOL)
  private InlineResponse20082LoadBalancerInstanceSslCert networkPool;

  public static final String SERIALIZED_NAME_NETWORK_DOMAIN = "networkDomain";
  @SerializedName(SERIALIZED_NAME_NETWORK_DOMAIN)
  private String networkDomain;

  public static final String SERIALIZED_NAME_TYPE = "type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private PriceSetVolumeType type;

  public static final String SERIALIZED_NAME_IP_MODE = "ipMode";
  @SerializedName(SERIALIZED_NAME_IP_MODE)
  private String ipMode;

  public static final String SERIALIZED_NAME_MAC_ADDRESS = "macAddress";
  @SerializedName(SERIALIZED_NAME_MAC_ADDRESS)
  private String macAddress;


  public ServerInterfaces id(Long id) {
    
    this.id = id;
    return this;
  }

   /**
   * Get id
   * @return id
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getId() {
    return id;
  }


  public void setId(Long id) {
    this.id = id;
  }


  public ServerInterfaces refType(String refType) {
    
    this.refType = refType;
    return this;
  }

   /**
   * Get refType
   * @return refType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRefType() {
    return refType;
  }


  public void setRefType(String refType) {
    this.refType = refType;
  }


  public ServerInterfaces refId(String refId) {
    
    this.refId = refId;
    return this;
  }

   /**
   * Get refId
   * @return refId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRefId() {
    return refId;
  }


  public void setRefId(String refId) {
    this.refId = refId;
  }


  public ServerInterfaces name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Get name
   * @return name
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public ServerInterfaces internalId(String internalId) {
    
    this.internalId = internalId;
    return this;
  }

   /**
   * Get internalId
   * @return internalId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getInternalId() {
    return internalId;
  }


  public void setInternalId(String internalId) {
    this.internalId = internalId;
  }


  public ServerInterfaces externalId(String externalId) {
    
    this.externalId = externalId;
    return this;
  }

   /**
   * Get externalId
   * @return externalId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getExternalId() {
    return externalId;
  }


  public void setExternalId(String externalId) {
    this.externalId = externalId;
  }


  public ServerInterfaces uniqueId(String uniqueId) {
    
    this.uniqueId = uniqueId;
    return this;
  }

   /**
   * Get uniqueId
   * @return uniqueId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getUniqueId() {
    return uniqueId;
  }


  public void setUniqueId(String uniqueId) {
    this.uniqueId = uniqueId;
  }


  public ServerInterfaces publicIpAddress(String publicIpAddress) {
    
    this.publicIpAddress = publicIpAddress;
    return this;
  }

   /**
   * Get publicIpAddress
   * @return publicIpAddress
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPublicIpAddress() {
    return publicIpAddress;
  }


  public void setPublicIpAddress(String publicIpAddress) {
    this.publicIpAddress = publicIpAddress;
  }


  public ServerInterfaces publicIpv6Address(String publicIpv6Address) {
    
    this.publicIpv6Address = publicIpv6Address;
    return this;
  }

   /**
   * Get publicIpv6Address
   * @return publicIpv6Address
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPublicIpv6Address() {
    return publicIpv6Address;
  }


  public void setPublicIpv6Address(String publicIpv6Address) {
    this.publicIpv6Address = publicIpv6Address;
  }


  public ServerInterfaces ipAddress(String ipAddress) {
    
    this.ipAddress = ipAddress;
    return this;
  }

   /**
   * Get ipAddress
   * @return ipAddress
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getIpAddress() {
    return ipAddress;
  }


  public void setIpAddress(String ipAddress) {
    this.ipAddress = ipAddress;
  }


  public ServerInterfaces ipv6Address(String ipv6Address) {
    
    this.ipv6Address = ipv6Address;
    return this;
  }

   /**
   * Get ipv6Address
   * @return ipv6Address
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getIpv6Address() {
    return ipv6Address;
  }


  public void setIpv6Address(String ipv6Address) {
    this.ipv6Address = ipv6Address;
  }


  public ServerInterfaces ipSubnet(String ipSubnet) {
    
    this.ipSubnet = ipSubnet;
    return this;
  }

   /**
   * Get ipSubnet
   * @return ipSubnet
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getIpSubnet() {
    return ipSubnet;
  }


  public void setIpSubnet(String ipSubnet) {
    this.ipSubnet = ipSubnet;
  }


  public ServerInterfaces ipv6Subnet(String ipv6Subnet) {
    
    this.ipv6Subnet = ipv6Subnet;
    return this;
  }

   /**
   * Get ipv6Subnet
   * @return ipv6Subnet
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getIpv6Subnet() {
    return ipv6Subnet;
  }


  public void setIpv6Subnet(String ipv6Subnet) {
    this.ipv6Subnet = ipv6Subnet;
  }


  public ServerInterfaces description(String description) {
    
    this.description = description;
    return this;
  }

   /**
   * Get description
   * @return description
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDescription() {
    return description;
  }


  public void setDescription(String description) {
    this.description = description;
  }


  public ServerInterfaces dhcp(Boolean dhcp) {
    
    this.dhcp = dhcp;
    return this;
  }

   /**
   * Get dhcp
   * @return dhcp
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getDhcp() {
    return dhcp;
  }


  public void setDhcp(Boolean dhcp) {
    this.dhcp = dhcp;
  }


  public ServerInterfaces active(Boolean active) {
    
    this.active = active;
    return this;
  }

   /**
   * Get active
   * @return active
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getActive() {
    return active;
  }


  public void setActive(Boolean active) {
    this.active = active;
  }


  public ServerInterfaces poolAssigned(Boolean poolAssigned) {
    
    this.poolAssigned = poolAssigned;
    return this;
  }

   /**
   * Get poolAssigned
   * @return poolAssigned
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getPoolAssigned() {
    return poolAssigned;
  }


  public void setPoolAssigned(Boolean poolAssigned) {
    this.poolAssigned = poolAssigned;
  }


  public ServerInterfaces primaryInterface(Boolean primaryInterface) {
    
    this.primaryInterface = primaryInterface;
    return this;
  }

   /**
   * Get primaryInterface
   * @return primaryInterface
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getPrimaryInterface() {
    return primaryInterface;
  }


  public void setPrimaryInterface(Boolean primaryInterface) {
    this.primaryInterface = primaryInterface;
  }


  public ServerInterfaces network(InlineResponse20040AppDeployInstance network) {
    
    this.network = network;
    return this;
  }

   /**
   * Get network
   * @return network
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse20040AppDeployInstance getNetwork() {
    return network;
  }


  public void setNetwork(InlineResponse20040AppDeployInstance network) {
    this.network = network;
  }


  public ServerInterfaces subnet(String subnet) {
    
    this.subnet = subnet;
    return this;
  }

   /**
   * Get subnet
   * @return subnet
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSubnet() {
    return subnet;
  }


  public void setSubnet(String subnet) {
    this.subnet = subnet;
  }


  public ServerInterfaces networkGroup(String networkGroup) {
    
    this.networkGroup = networkGroup;
    return this;
  }

   /**
   * Get networkGroup
   * @return networkGroup
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getNetworkGroup() {
    return networkGroup;
  }


  public void setNetworkGroup(String networkGroup) {
    this.networkGroup = networkGroup;
  }


  public ServerInterfaces networkPosition(String networkPosition) {
    
    this.networkPosition = networkPosition;
    return this;
  }

   /**
   * Get networkPosition
   * @return networkPosition
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getNetworkPosition() {
    return networkPosition;
  }


  public void setNetworkPosition(String networkPosition) {
    this.networkPosition = networkPosition;
  }


  public ServerInterfaces networkPool(InlineResponse20082LoadBalancerInstanceSslCert networkPool) {
    
    this.networkPool = networkPool;
    return this;
  }

   /**
   * Get networkPool
   * @return networkPool
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse20082LoadBalancerInstanceSslCert getNetworkPool() {
    return networkPool;
  }


  public void setNetworkPool(InlineResponse20082LoadBalancerInstanceSslCert networkPool) {
    this.networkPool = networkPool;
  }


  public ServerInterfaces networkDomain(String networkDomain) {
    
    this.networkDomain = networkDomain;
    return this;
  }

   /**
   * Get networkDomain
   * @return networkDomain
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getNetworkDomain() {
    return networkDomain;
  }


  public void setNetworkDomain(String networkDomain) {
    this.networkDomain = networkDomain;
  }


  public ServerInterfaces type(PriceSetVolumeType type) {
    
    this.type = type;
    return this;
  }

   /**
   * Get type
   * @return type
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public PriceSetVolumeType getType() {
    return type;
  }


  public void setType(PriceSetVolumeType type) {
    this.type = type;
  }


  public ServerInterfaces ipMode(String ipMode) {
    
    this.ipMode = ipMode;
    return this;
  }

   /**
   * Get ipMode
   * @return ipMode
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getIpMode() {
    return ipMode;
  }


  public void setIpMode(String ipMode) {
    this.ipMode = ipMode;
  }


  public ServerInterfaces macAddress(String macAddress) {
    
    this.macAddress = macAddress;
    return this;
  }

   /**
   * Get macAddress
   * @return macAddress
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getMacAddress() {
    return macAddress;
  }


  public void setMacAddress(String macAddress) {
    this.macAddress = macAddress;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ServerInterfaces serverInterfaces = (ServerInterfaces) o;
    return Objects.equals(this.id, serverInterfaces.id) &&
        Objects.equals(this.refType, serverInterfaces.refType) &&
        Objects.equals(this.refId, serverInterfaces.refId) &&
        Objects.equals(this.name, serverInterfaces.name) &&
        Objects.equals(this.internalId, serverInterfaces.internalId) &&
        Objects.equals(this.externalId, serverInterfaces.externalId) &&
        Objects.equals(this.uniqueId, serverInterfaces.uniqueId) &&
        Objects.equals(this.publicIpAddress, serverInterfaces.publicIpAddress) &&
        Objects.equals(this.publicIpv6Address, serverInterfaces.publicIpv6Address) &&
        Objects.equals(this.ipAddress, serverInterfaces.ipAddress) &&
        Objects.equals(this.ipv6Address, serverInterfaces.ipv6Address) &&
        Objects.equals(this.ipSubnet, serverInterfaces.ipSubnet) &&
        Objects.equals(this.ipv6Subnet, serverInterfaces.ipv6Subnet) &&
        Objects.equals(this.description, serverInterfaces.description) &&
        Objects.equals(this.dhcp, serverInterfaces.dhcp) &&
        Objects.equals(this.active, serverInterfaces.active) &&
        Objects.equals(this.poolAssigned, serverInterfaces.poolAssigned) &&
        Objects.equals(this.primaryInterface, serverInterfaces.primaryInterface) &&
        Objects.equals(this.network, serverInterfaces.network) &&
        Objects.equals(this.subnet, serverInterfaces.subnet) &&
        Objects.equals(this.networkGroup, serverInterfaces.networkGroup) &&
        Objects.equals(this.networkPosition, serverInterfaces.networkPosition) &&
        Objects.equals(this.networkPool, serverInterfaces.networkPool) &&
        Objects.equals(this.networkDomain, serverInterfaces.networkDomain) &&
        Objects.equals(this.type, serverInterfaces.type) &&
        Objects.equals(this.ipMode, serverInterfaces.ipMode) &&
        Objects.equals(this.macAddress, serverInterfaces.macAddress);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, refType, refId, name, internalId, externalId, uniqueId, publicIpAddress, publicIpv6Address, ipAddress, ipv6Address, ipSubnet, ipv6Subnet, description, dhcp, active, poolAssigned, primaryInterface, network, subnet, networkGroup, networkPosition, networkPool, networkDomain, type, ipMode, macAddress);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ServerInterfaces {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    refType: ").append(toIndentedString(refType)).append("\n");
    sb.append("    refId: ").append(toIndentedString(refId)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    internalId: ").append(toIndentedString(internalId)).append("\n");
    sb.append("    externalId: ").append(toIndentedString(externalId)).append("\n");
    sb.append("    uniqueId: ").append(toIndentedString(uniqueId)).append("\n");
    sb.append("    publicIpAddress: ").append(toIndentedString(publicIpAddress)).append("\n");
    sb.append("    publicIpv6Address: ").append(toIndentedString(publicIpv6Address)).append("\n");
    sb.append("    ipAddress: ").append(toIndentedString(ipAddress)).append("\n");
    sb.append("    ipv6Address: ").append(toIndentedString(ipv6Address)).append("\n");
    sb.append("    ipSubnet: ").append(toIndentedString(ipSubnet)).append("\n");
    sb.append("    ipv6Subnet: ").append(toIndentedString(ipv6Subnet)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    dhcp: ").append(toIndentedString(dhcp)).append("\n");
    sb.append("    active: ").append(toIndentedString(active)).append("\n");
    sb.append("    poolAssigned: ").append(toIndentedString(poolAssigned)).append("\n");
    sb.append("    primaryInterface: ").append(toIndentedString(primaryInterface)).append("\n");
    sb.append("    network: ").append(toIndentedString(network)).append("\n");
    sb.append("    subnet: ").append(toIndentedString(subnet)).append("\n");
    sb.append("    networkGroup: ").append(toIndentedString(networkGroup)).append("\n");
    sb.append("    networkPosition: ").append(toIndentedString(networkPosition)).append("\n");
    sb.append("    networkPool: ").append(toIndentedString(networkPool)).append("\n");
    sb.append("    networkDomain: ").append(toIndentedString(networkDomain)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    ipMode: ").append(toIndentedString(ipMode)).append("\n");
    sb.append("    macAddress: ").append(toIndentedString(macAddress)).append("\n");
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
