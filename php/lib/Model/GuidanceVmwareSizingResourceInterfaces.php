<?php
/**
 * GuidanceVmwareSizingResourceInterfaces
 *
 * PHP version 7.2
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */

/**
 * Morpheus API
 *
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.2.1
 * Contact: dev@morpheusdata.com
 * Generated by: https://openapi-generator.tech
 * OpenAPI Generator version: 5.0.0
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

namespace OpenAPI\Client\Model;

use \ArrayAccess;
use \OpenAPI\Client\ObjectSerializer;

/**
 * GuidanceVmwareSizingResourceInterfaces Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class GuidanceVmwareSizingResourceInterfaces implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'guidanceVmwareSizing_resource_interfaces';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'id' => 'int',
        'ref_type' => 'string',
        'ref_id' => 'string',
        'name' => 'string',
        'internal_id' => 'string',
        'external_id' => 'string',
        'unique_id' => 'string',
        'public_ip_address' => 'string',
        'public_ipv6_address' => 'string',
        'ip_address' => 'string',
        'ipv6_address' => 'string',
        'ip_subnet' => 'string',
        'ipv6_subnet' => 'string',
        'description' => 'string',
        'dhcp' => 'bool',
        'active' => 'bool',
        'pool_assigned' => 'bool',
        'primary_interface' => 'bool',
        'network' => '\OpenAPI\Client\Model\InlineResponse20040AppDeployInstance',
        'subnet' => 'string',
        'network_group' => 'string',
        'network_position' => 'string',
        'network_pool' => '\OpenAPI\Client\Model\InlineResponse20040AppDeployInstance',
        'network_domain' => 'string',
        'type' => '\OpenAPI\Client\Model\InlineResponse20079LoadBalancerMonitorLoadBalancerType',
        'ip_mode' => 'string',
        'mac_address' => 'string'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'id' => 'int64',
        'ref_type' => null,
        'ref_id' => null,
        'name' => null,
        'internal_id' => null,
        'external_id' => null,
        'unique_id' => null,
        'public_ip_address' => null,
        'public_ipv6_address' => null,
        'ip_address' => null,
        'ipv6_address' => null,
        'ip_subnet' => null,
        'ipv6_subnet' => null,
        'description' => null,
        'dhcp' => null,
        'active' => null,
        'pool_assigned' => null,
        'primary_interface' => null,
        'network' => null,
        'subnet' => null,
        'network_group' => null,
        'network_position' => null,
        'network_pool' => null,
        'network_domain' => null,
        'type' => null,
        'ip_mode' => null,
        'mac_address' => null
    ];

    /**
     * Array of property to type mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPITypes()
    {
        return self::$openAPITypes;
    }

    /**
     * Array of property to format mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPIFormats()
    {
        return self::$openAPIFormats;
    }

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @var string[]
     */
    protected static $attributeMap = [
        'id' => 'id',
        'ref_type' => 'refType',
        'ref_id' => 'refId',
        'name' => 'name',
        'internal_id' => 'internalId',
        'external_id' => 'externalId',
        'unique_id' => 'uniqueId',
        'public_ip_address' => 'publicIpAddress',
        'public_ipv6_address' => 'publicIpv6Address',
        'ip_address' => 'ipAddress',
        'ipv6_address' => 'ipv6Address',
        'ip_subnet' => 'ipSubnet',
        'ipv6_subnet' => 'ipv6Subnet',
        'description' => 'description',
        'dhcp' => 'dhcp',
        'active' => 'active',
        'pool_assigned' => 'poolAssigned',
        'primary_interface' => 'primaryInterface',
        'network' => 'network',
        'subnet' => 'subnet',
        'network_group' => 'networkGroup',
        'network_position' => 'networkPosition',
        'network_pool' => 'networkPool',
        'network_domain' => 'networkDomain',
        'type' => 'type',
        'ip_mode' => 'ipMode',
        'mac_address' => 'macAddress'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'id' => 'setId',
        'ref_type' => 'setRefType',
        'ref_id' => 'setRefId',
        'name' => 'setName',
        'internal_id' => 'setInternalId',
        'external_id' => 'setExternalId',
        'unique_id' => 'setUniqueId',
        'public_ip_address' => 'setPublicIpAddress',
        'public_ipv6_address' => 'setPublicIpv6Address',
        'ip_address' => 'setIpAddress',
        'ipv6_address' => 'setIpv6Address',
        'ip_subnet' => 'setIpSubnet',
        'ipv6_subnet' => 'setIpv6Subnet',
        'description' => 'setDescription',
        'dhcp' => 'setDhcp',
        'active' => 'setActive',
        'pool_assigned' => 'setPoolAssigned',
        'primary_interface' => 'setPrimaryInterface',
        'network' => 'setNetwork',
        'subnet' => 'setSubnet',
        'network_group' => 'setNetworkGroup',
        'network_position' => 'setNetworkPosition',
        'network_pool' => 'setNetworkPool',
        'network_domain' => 'setNetworkDomain',
        'type' => 'setType',
        'ip_mode' => 'setIpMode',
        'mac_address' => 'setMacAddress'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'id' => 'getId',
        'ref_type' => 'getRefType',
        'ref_id' => 'getRefId',
        'name' => 'getName',
        'internal_id' => 'getInternalId',
        'external_id' => 'getExternalId',
        'unique_id' => 'getUniqueId',
        'public_ip_address' => 'getPublicIpAddress',
        'public_ipv6_address' => 'getPublicIpv6Address',
        'ip_address' => 'getIpAddress',
        'ipv6_address' => 'getIpv6Address',
        'ip_subnet' => 'getIpSubnet',
        'ipv6_subnet' => 'getIpv6Subnet',
        'description' => 'getDescription',
        'dhcp' => 'getDhcp',
        'active' => 'getActive',
        'pool_assigned' => 'getPoolAssigned',
        'primary_interface' => 'getPrimaryInterface',
        'network' => 'getNetwork',
        'subnet' => 'getSubnet',
        'network_group' => 'getNetworkGroup',
        'network_position' => 'getNetworkPosition',
        'network_pool' => 'getNetworkPool',
        'network_domain' => 'getNetworkDomain',
        'type' => 'getType',
        'ip_mode' => 'getIpMode',
        'mac_address' => 'getMacAddress'
    ];

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @return array
     */
    public static function attributeMap()
    {
        return self::$attributeMap;
    }

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @return array
     */
    public static function setters()
    {
        return self::$setters;
    }

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @return array
     */
    public static function getters()
    {
        return self::$getters;
    }

    /**
     * The original name of the model.
     *
     * @return string
     */
    public function getModelName()
    {
        return self::$openAPIModelName;
    }

    

    

    /**
     * Associative array for storing property values
     *
     * @var mixed[]
     */
    protected $container = [];

    /**
     * Constructor
     *
     * @param mixed[] $data Associated array of property values
     *                      initializing the model
     */
    public function __construct(array $data = null)
    {
        $this->container['id'] = $data['id'] ?? null;
        $this->container['ref_type'] = $data['ref_type'] ?? null;
        $this->container['ref_id'] = $data['ref_id'] ?? null;
        $this->container['name'] = $data['name'] ?? null;
        $this->container['internal_id'] = $data['internal_id'] ?? null;
        $this->container['external_id'] = $data['external_id'] ?? null;
        $this->container['unique_id'] = $data['unique_id'] ?? null;
        $this->container['public_ip_address'] = $data['public_ip_address'] ?? null;
        $this->container['public_ipv6_address'] = $data['public_ipv6_address'] ?? null;
        $this->container['ip_address'] = $data['ip_address'] ?? null;
        $this->container['ipv6_address'] = $data['ipv6_address'] ?? null;
        $this->container['ip_subnet'] = $data['ip_subnet'] ?? null;
        $this->container['ipv6_subnet'] = $data['ipv6_subnet'] ?? null;
        $this->container['description'] = $data['description'] ?? null;
        $this->container['dhcp'] = $data['dhcp'] ?? null;
        $this->container['active'] = $data['active'] ?? null;
        $this->container['pool_assigned'] = $data['pool_assigned'] ?? null;
        $this->container['primary_interface'] = $data['primary_interface'] ?? null;
        $this->container['network'] = $data['network'] ?? null;
        $this->container['subnet'] = $data['subnet'] ?? null;
        $this->container['network_group'] = $data['network_group'] ?? null;
        $this->container['network_position'] = $data['network_position'] ?? null;
        $this->container['network_pool'] = $data['network_pool'] ?? null;
        $this->container['network_domain'] = $data['network_domain'] ?? null;
        $this->container['type'] = $data['type'] ?? null;
        $this->container['ip_mode'] = $data['ip_mode'] ?? null;
        $this->container['mac_address'] = $data['mac_address'] ?? null;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        return $invalidProperties;
    }

    /**
     * Validate all the properties in the model
     * return true if all passed
     *
     * @return bool True if all properties are valid
     */
    public function valid()
    {
        return count($this->listInvalidProperties()) === 0;
    }


    /**
     * Gets id
     *
     * @return int|null
     */
    public function getId()
    {
        return $this->container['id'];
    }

    /**
     * Sets id
     *
     * @param int|null $id id
     *
     * @return self
     */
    public function setId($id)
    {
        $this->container['id'] = $id;

        return $this;
    }

    /**
     * Gets ref_type
     *
     * @return string|null
     */
    public function getRefType()
    {
        return $this->container['ref_type'];
    }

    /**
     * Sets ref_type
     *
     * @param string|null $ref_type ref_type
     *
     * @return self
     */
    public function setRefType($ref_type)
    {
        $this->container['ref_type'] = $ref_type;

        return $this;
    }

    /**
     * Gets ref_id
     *
     * @return string|null
     */
    public function getRefId()
    {
        return $this->container['ref_id'];
    }

    /**
     * Sets ref_id
     *
     * @param string|null $ref_id ref_id
     *
     * @return self
     */
    public function setRefId($ref_id)
    {
        $this->container['ref_id'] = $ref_id;

        return $this;
    }

    /**
     * Gets name
     *
     * @return string|null
     */
    public function getName()
    {
        return $this->container['name'];
    }

    /**
     * Sets name
     *
     * @param string|null $name name
     *
     * @return self
     */
    public function setName($name)
    {
        $this->container['name'] = $name;

        return $this;
    }

    /**
     * Gets internal_id
     *
     * @return string|null
     */
    public function getInternalId()
    {
        return $this->container['internal_id'];
    }

    /**
     * Sets internal_id
     *
     * @param string|null $internal_id internal_id
     *
     * @return self
     */
    public function setInternalId($internal_id)
    {
        $this->container['internal_id'] = $internal_id;

        return $this;
    }

    /**
     * Gets external_id
     *
     * @return string|null
     */
    public function getExternalId()
    {
        return $this->container['external_id'];
    }

    /**
     * Sets external_id
     *
     * @param string|null $external_id external_id
     *
     * @return self
     */
    public function setExternalId($external_id)
    {
        $this->container['external_id'] = $external_id;

        return $this;
    }

    /**
     * Gets unique_id
     *
     * @return string|null
     */
    public function getUniqueId()
    {
        return $this->container['unique_id'];
    }

    /**
     * Sets unique_id
     *
     * @param string|null $unique_id unique_id
     *
     * @return self
     */
    public function setUniqueId($unique_id)
    {
        $this->container['unique_id'] = $unique_id;

        return $this;
    }

    /**
     * Gets public_ip_address
     *
     * @return string|null
     */
    public function getPublicIpAddress()
    {
        return $this->container['public_ip_address'];
    }

    /**
     * Sets public_ip_address
     *
     * @param string|null $public_ip_address public_ip_address
     *
     * @return self
     */
    public function setPublicIpAddress($public_ip_address)
    {
        $this->container['public_ip_address'] = $public_ip_address;

        return $this;
    }

    /**
     * Gets public_ipv6_address
     *
     * @return string|null
     */
    public function getPublicIpv6Address()
    {
        return $this->container['public_ipv6_address'];
    }

    /**
     * Sets public_ipv6_address
     *
     * @param string|null $public_ipv6_address public_ipv6_address
     *
     * @return self
     */
    public function setPublicIpv6Address($public_ipv6_address)
    {
        $this->container['public_ipv6_address'] = $public_ipv6_address;

        return $this;
    }

    /**
     * Gets ip_address
     *
     * @return string|null
     */
    public function getIpAddress()
    {
        return $this->container['ip_address'];
    }

    /**
     * Sets ip_address
     *
     * @param string|null $ip_address ip_address
     *
     * @return self
     */
    public function setIpAddress($ip_address)
    {
        $this->container['ip_address'] = $ip_address;

        return $this;
    }

    /**
     * Gets ipv6_address
     *
     * @return string|null
     */
    public function getIpv6Address()
    {
        return $this->container['ipv6_address'];
    }

    /**
     * Sets ipv6_address
     *
     * @param string|null $ipv6_address ipv6_address
     *
     * @return self
     */
    public function setIpv6Address($ipv6_address)
    {
        $this->container['ipv6_address'] = $ipv6_address;

        return $this;
    }

    /**
     * Gets ip_subnet
     *
     * @return string|null
     */
    public function getIpSubnet()
    {
        return $this->container['ip_subnet'];
    }

    /**
     * Sets ip_subnet
     *
     * @param string|null $ip_subnet ip_subnet
     *
     * @return self
     */
    public function setIpSubnet($ip_subnet)
    {
        $this->container['ip_subnet'] = $ip_subnet;

        return $this;
    }

    /**
     * Gets ipv6_subnet
     *
     * @return string|null
     */
    public function getIpv6Subnet()
    {
        return $this->container['ipv6_subnet'];
    }

    /**
     * Sets ipv6_subnet
     *
     * @param string|null $ipv6_subnet ipv6_subnet
     *
     * @return self
     */
    public function setIpv6Subnet($ipv6_subnet)
    {
        $this->container['ipv6_subnet'] = $ipv6_subnet;

        return $this;
    }

    /**
     * Gets description
     *
     * @return string|null
     */
    public function getDescription()
    {
        return $this->container['description'];
    }

    /**
     * Sets description
     *
     * @param string|null $description description
     *
     * @return self
     */
    public function setDescription($description)
    {
        $this->container['description'] = $description;

        return $this;
    }

    /**
     * Gets dhcp
     *
     * @return bool|null
     */
    public function getDhcp()
    {
        return $this->container['dhcp'];
    }

    /**
     * Sets dhcp
     *
     * @param bool|null $dhcp dhcp
     *
     * @return self
     */
    public function setDhcp($dhcp)
    {
        $this->container['dhcp'] = $dhcp;

        return $this;
    }

    /**
     * Gets active
     *
     * @return bool|null
     */
    public function getActive()
    {
        return $this->container['active'];
    }

    /**
     * Sets active
     *
     * @param bool|null $active active
     *
     * @return self
     */
    public function setActive($active)
    {
        $this->container['active'] = $active;

        return $this;
    }

    /**
     * Gets pool_assigned
     *
     * @return bool|null
     */
    public function getPoolAssigned()
    {
        return $this->container['pool_assigned'];
    }

    /**
     * Sets pool_assigned
     *
     * @param bool|null $pool_assigned pool_assigned
     *
     * @return self
     */
    public function setPoolAssigned($pool_assigned)
    {
        $this->container['pool_assigned'] = $pool_assigned;

        return $this;
    }

    /**
     * Gets primary_interface
     *
     * @return bool|null
     */
    public function getPrimaryInterface()
    {
        return $this->container['primary_interface'];
    }

    /**
     * Sets primary_interface
     *
     * @param bool|null $primary_interface primary_interface
     *
     * @return self
     */
    public function setPrimaryInterface($primary_interface)
    {
        $this->container['primary_interface'] = $primary_interface;

        return $this;
    }

    /**
     * Gets network
     *
     * @return \OpenAPI\Client\Model\InlineResponse20040AppDeployInstance|null
     */
    public function getNetwork()
    {
        return $this->container['network'];
    }

    /**
     * Sets network
     *
     * @param \OpenAPI\Client\Model\InlineResponse20040AppDeployInstance|null $network network
     *
     * @return self
     */
    public function setNetwork($network)
    {
        $this->container['network'] = $network;

        return $this;
    }

    /**
     * Gets subnet
     *
     * @return string|null
     */
    public function getSubnet()
    {
        return $this->container['subnet'];
    }

    /**
     * Sets subnet
     *
     * @param string|null $subnet subnet
     *
     * @return self
     */
    public function setSubnet($subnet)
    {
        $this->container['subnet'] = $subnet;

        return $this;
    }

    /**
     * Gets network_group
     *
     * @return string|null
     */
    public function getNetworkGroup()
    {
        return $this->container['network_group'];
    }

    /**
     * Sets network_group
     *
     * @param string|null $network_group network_group
     *
     * @return self
     */
    public function setNetworkGroup($network_group)
    {
        $this->container['network_group'] = $network_group;

        return $this;
    }

    /**
     * Gets network_position
     *
     * @return string|null
     */
    public function getNetworkPosition()
    {
        return $this->container['network_position'];
    }

    /**
     * Sets network_position
     *
     * @param string|null $network_position network_position
     *
     * @return self
     */
    public function setNetworkPosition($network_position)
    {
        $this->container['network_position'] = $network_position;

        return $this;
    }

    /**
     * Gets network_pool
     *
     * @return \OpenAPI\Client\Model\InlineResponse20040AppDeployInstance|null
     */
    public function getNetworkPool()
    {
        return $this->container['network_pool'];
    }

    /**
     * Sets network_pool
     *
     * @param \OpenAPI\Client\Model\InlineResponse20040AppDeployInstance|null $network_pool network_pool
     *
     * @return self
     */
    public function setNetworkPool($network_pool)
    {
        $this->container['network_pool'] = $network_pool;

        return $this;
    }

    /**
     * Gets network_domain
     *
     * @return string|null
     */
    public function getNetworkDomain()
    {
        return $this->container['network_domain'];
    }

    /**
     * Sets network_domain
     *
     * @param string|null $network_domain network_domain
     *
     * @return self
     */
    public function setNetworkDomain($network_domain)
    {
        $this->container['network_domain'] = $network_domain;

        return $this;
    }

    /**
     * Gets type
     *
     * @return \OpenAPI\Client\Model\InlineResponse20079LoadBalancerMonitorLoadBalancerType|null
     */
    public function getType()
    {
        return $this->container['type'];
    }

    /**
     * Sets type
     *
     * @param \OpenAPI\Client\Model\InlineResponse20079LoadBalancerMonitorLoadBalancerType|null $type type
     *
     * @return self
     */
    public function setType($type)
    {
        $this->container['type'] = $type;

        return $this;
    }

    /**
     * Gets ip_mode
     *
     * @return string|null
     */
    public function getIpMode()
    {
        return $this->container['ip_mode'];
    }

    /**
     * Sets ip_mode
     *
     * @param string|null $ip_mode ip_mode
     *
     * @return self
     */
    public function setIpMode($ip_mode)
    {
        $this->container['ip_mode'] = $ip_mode;

        return $this;
    }

    /**
     * Gets mac_address
     *
     * @return string|null
     */
    public function getMacAddress()
    {
        return $this->container['mac_address'];
    }

    /**
     * Sets mac_address
     *
     * @param string|null $mac_address mac_address
     *
     * @return self
     */
    public function setMacAddress($mac_address)
    {
        $this->container['mac_address'] = $mac_address;

        return $this;
    }
    /**
     * Returns true if offset exists. False otherwise.
     *
     * @param integer $offset Offset
     *
     * @return boolean
     */
    public function offsetExists($offset)
    {
        return isset($this->container[$offset]);
    }

    /**
     * Gets offset.
     *
     * @param integer $offset Offset
     *
     * @return mixed|null
     */
    public function offsetGet($offset)
    {
        return $this->container[$offset] ?? null;
    }

    /**
     * Sets value based on offset.
     *
     * @param int|null $offset Offset
     * @param mixed    $value  Value to be set
     *
     * @return void
     */
    public function offsetSet($offset, $value)
    {
        if (is_null($offset)) {
            $this->container[] = $value;
        } else {
            $this->container[$offset] = $value;
        }
    }

    /**
     * Unsets offset.
     *
     * @param integer $offset Offset
     *
     * @return void
     */
    public function offsetUnset($offset)
    {
        unset($this->container[$offset]);
    }

    /**
     * Serializes the object to a value that can be serialized natively by json_encode().
     * @link https://www.php.net/manual/en/jsonserializable.jsonserialize.php
     *
     * @return mixed Returns data which can be serialized by json_encode(), which is a value
     * of any type other than a resource.
     */
    public function jsonSerialize()
    {
       return ObjectSerializer::sanitizeForSerialization($this);
    }

    /**
     * Gets the string presentation of the object
     *
     * @return string
     */
    public function __toString()
    {
        return json_encode(
            ObjectSerializer::sanitizeForSerialization($this),
            JSON_PRETTY_PRINT
        );
    }

    /**
     * Gets a header-safe presentation of the object
     *
     * @return string
     */
    public function toHeaderValue()
    {
        return json_encode(ObjectSerializer::sanitizeForSerialization($this));
    }
}


