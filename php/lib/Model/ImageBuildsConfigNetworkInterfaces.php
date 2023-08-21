<?php
/**
 * ImageBuildsConfigNetworkInterfaces
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
 * ImageBuildsConfigNetworkInterfaces Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class ImageBuildsConfigNetworkInterfaces implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'imageBuilds_config_networkInterfaces';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'ip_mode' => 'string',
        'primary_interface' => 'bool',
        'show_network_pool_label' => 'bool',
        'show_network_dhcp_label' => 'bool',
        'network' => '\OpenAPI\Client\Model\ImageBuildsConfigNetwork',
        'network_interface_type_id' => 'int',
        'network_interface_type_id_name' => 'string'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'ip_mode' => null,
        'primary_interface' => null,
        'show_network_pool_label' => null,
        'show_network_dhcp_label' => null,
        'network' => null,
        'network_interface_type_id' => 'int64',
        'network_interface_type_id_name' => null
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
        'ip_mode' => 'ipMode',
        'primary_interface' => 'primaryInterface',
        'show_network_pool_label' => 'showNetworkPoolLabel',
        'show_network_dhcp_label' => 'showNetworkDhcpLabel',
        'network' => 'network',
        'network_interface_type_id' => 'networkInterfaceTypeId',
        'network_interface_type_id_name' => 'networkInterfaceTypeIdName'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'ip_mode' => 'setIpMode',
        'primary_interface' => 'setPrimaryInterface',
        'show_network_pool_label' => 'setShowNetworkPoolLabel',
        'show_network_dhcp_label' => 'setShowNetworkDhcpLabel',
        'network' => 'setNetwork',
        'network_interface_type_id' => 'setNetworkInterfaceTypeId',
        'network_interface_type_id_name' => 'setNetworkInterfaceTypeIdName'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'ip_mode' => 'getIpMode',
        'primary_interface' => 'getPrimaryInterface',
        'show_network_pool_label' => 'getShowNetworkPoolLabel',
        'show_network_dhcp_label' => 'getShowNetworkDhcpLabel',
        'network' => 'getNetwork',
        'network_interface_type_id' => 'getNetworkInterfaceTypeId',
        'network_interface_type_id_name' => 'getNetworkInterfaceTypeIdName'
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
        $this->container['ip_mode'] = $data['ip_mode'] ?? null;
        $this->container['primary_interface'] = $data['primary_interface'] ?? null;
        $this->container['show_network_pool_label'] = $data['show_network_pool_label'] ?? null;
        $this->container['show_network_dhcp_label'] = $data['show_network_dhcp_label'] ?? null;
        $this->container['network'] = $data['network'] ?? null;
        $this->container['network_interface_type_id'] = $data['network_interface_type_id'] ?? null;
        $this->container['network_interface_type_id_name'] = $data['network_interface_type_id_name'] ?? null;
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
     * Gets show_network_pool_label
     *
     * @return bool|null
     */
    public function getShowNetworkPoolLabel()
    {
        return $this->container['show_network_pool_label'];
    }

    /**
     * Sets show_network_pool_label
     *
     * @param bool|null $show_network_pool_label show_network_pool_label
     *
     * @return self
     */
    public function setShowNetworkPoolLabel($show_network_pool_label)
    {
        $this->container['show_network_pool_label'] = $show_network_pool_label;

        return $this;
    }

    /**
     * Gets show_network_dhcp_label
     *
     * @return bool|null
     */
    public function getShowNetworkDhcpLabel()
    {
        return $this->container['show_network_dhcp_label'];
    }

    /**
     * Sets show_network_dhcp_label
     *
     * @param bool|null $show_network_dhcp_label show_network_dhcp_label
     *
     * @return self
     */
    public function setShowNetworkDhcpLabel($show_network_dhcp_label)
    {
        $this->container['show_network_dhcp_label'] = $show_network_dhcp_label;

        return $this;
    }

    /**
     * Gets network
     *
     * @return \OpenAPI\Client\Model\ImageBuildsConfigNetwork|null
     */
    public function getNetwork()
    {
        return $this->container['network'];
    }

    /**
     * Sets network
     *
     * @param \OpenAPI\Client\Model\ImageBuildsConfigNetwork|null $network network
     *
     * @return self
     */
    public function setNetwork($network)
    {
        $this->container['network'] = $network;

        return $this;
    }

    /**
     * Gets network_interface_type_id
     *
     * @return int|null
     */
    public function getNetworkInterfaceTypeId()
    {
        return $this->container['network_interface_type_id'];
    }

    /**
     * Sets network_interface_type_id
     *
     * @param int|null $network_interface_type_id network_interface_type_id
     *
     * @return self
     */
    public function setNetworkInterfaceTypeId($network_interface_type_id)
    {
        $this->container['network_interface_type_id'] = $network_interface_type_id;

        return $this;
    }

    /**
     * Gets network_interface_type_id_name
     *
     * @return string|null
     */
    public function getNetworkInterfaceTypeIdName()
    {
        return $this->container['network_interface_type_id_name'];
    }

    /**
     * Sets network_interface_type_id_name
     *
     * @param string|null $network_interface_type_id_name network_interface_type_id_name
     *
     * @return self
     */
    public function setNetworkInterfaceTypeIdName($network_interface_type_id_name)
    {
        $this->container['network_interface_type_id_name'] = $network_interface_type_id_name;

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


