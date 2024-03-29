<?php
/**
 * GroupConfig
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
 * GroupConfig Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class GroupConfig implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'group_config';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'dns_integration_id' => 'string',
        'config_cmdb_id' => 'string',
        'config_cm_id' => 'string',
        'service_registry_id' => 'string',
        'config_management_id' => 'string',
        'config_cmdb_discovery' => 'bool'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'dns_integration_id' => null,
        'config_cmdb_id' => null,
        'config_cm_id' => null,
        'service_registry_id' => null,
        'config_management_id' => null,
        'config_cmdb_discovery' => null
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
        'dns_integration_id' => 'dnsIntegrationId',
        'config_cmdb_id' => 'configCmdbId',
        'config_cm_id' => 'configCmId',
        'service_registry_id' => 'serviceRegistryId',
        'config_management_id' => 'configManagementId',
        'config_cmdb_discovery' => 'configCmdbDiscovery'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'dns_integration_id' => 'setDnsIntegrationId',
        'config_cmdb_id' => 'setConfigCmdbId',
        'config_cm_id' => 'setConfigCmId',
        'service_registry_id' => 'setServiceRegistryId',
        'config_management_id' => 'setConfigManagementId',
        'config_cmdb_discovery' => 'setConfigCmdbDiscovery'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'dns_integration_id' => 'getDnsIntegrationId',
        'config_cmdb_id' => 'getConfigCmdbId',
        'config_cm_id' => 'getConfigCmId',
        'service_registry_id' => 'getServiceRegistryId',
        'config_management_id' => 'getConfigManagementId',
        'config_cmdb_discovery' => 'getConfigCmdbDiscovery'
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
        $this->container['dns_integration_id'] = $data['dns_integration_id'] ?? null;
        $this->container['config_cmdb_id'] = $data['config_cmdb_id'] ?? null;
        $this->container['config_cm_id'] = $data['config_cm_id'] ?? null;
        $this->container['service_registry_id'] = $data['service_registry_id'] ?? null;
        $this->container['config_management_id'] = $data['config_management_id'] ?? null;
        $this->container['config_cmdb_discovery'] = $data['config_cmdb_discovery'] ?? null;
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
     * Gets dns_integration_id
     *
     * @return string|null
     */
    public function getDnsIntegrationId()
    {
        return $this->container['dns_integration_id'];
    }

    /**
     * Sets dns_integration_id
     *
     * @param string|null $dns_integration_id dns_integration_id
     *
     * @return self
     */
    public function setDnsIntegrationId($dns_integration_id)
    {
        $this->container['dns_integration_id'] = $dns_integration_id;

        return $this;
    }

    /**
     * Gets config_cmdb_id
     *
     * @return string|null
     */
    public function getConfigCmdbId()
    {
        return $this->container['config_cmdb_id'];
    }

    /**
     * Sets config_cmdb_id
     *
     * @param string|null $config_cmdb_id config_cmdb_id
     *
     * @return self
     */
    public function setConfigCmdbId($config_cmdb_id)
    {
        $this->container['config_cmdb_id'] = $config_cmdb_id;

        return $this;
    }

    /**
     * Gets config_cm_id
     *
     * @return string|null
     */
    public function getConfigCmId()
    {
        return $this->container['config_cm_id'];
    }

    /**
     * Sets config_cm_id
     *
     * @param string|null $config_cm_id config_cm_id
     *
     * @return self
     */
    public function setConfigCmId($config_cm_id)
    {
        $this->container['config_cm_id'] = $config_cm_id;

        return $this;
    }

    /**
     * Gets service_registry_id
     *
     * @return string|null
     */
    public function getServiceRegistryId()
    {
        return $this->container['service_registry_id'];
    }

    /**
     * Sets service_registry_id
     *
     * @param string|null $service_registry_id service_registry_id
     *
     * @return self
     */
    public function setServiceRegistryId($service_registry_id)
    {
        $this->container['service_registry_id'] = $service_registry_id;

        return $this;
    }

    /**
     * Gets config_management_id
     *
     * @return string|null
     */
    public function getConfigManagementId()
    {
        return $this->container['config_management_id'];
    }

    /**
     * Sets config_management_id
     *
     * @param string|null $config_management_id config_management_id
     *
     * @return self
     */
    public function setConfigManagementId($config_management_id)
    {
        $this->container['config_management_id'] = $config_management_id;

        return $this;
    }

    /**
     * Gets config_cmdb_discovery
     *
     * @return bool|null
     */
    public function getConfigCmdbDiscovery()
    {
        return $this->container['config_cmdb_discovery'];
    }

    /**
     * Sets config_cmdb_discovery
     *
     * @param bool|null $config_cmdb_discovery config_cmdb_discovery
     *
     * @return self
     */
    public function setConfigCmdbDiscovery($config_cmdb_discovery)
    {
        $this->container['config_cmdb_discovery'] = $config_cmdb_discovery;

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


