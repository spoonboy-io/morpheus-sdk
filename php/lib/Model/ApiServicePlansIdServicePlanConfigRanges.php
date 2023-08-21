<?php
/**
 * ApiServicePlansIdServicePlanConfigRanges
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
 * ApiServicePlansIdServicePlanConfigRanges Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class ApiServicePlansIdServicePlanConfigRanges implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = '_api_service_plans__id__servicePlan_config_ranges';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'min_storage' => 'string',
        'max_storage' => 'string',
        'min_memory' => 'int',
        'max_memory' => 'int',
        'min_cores' => 'string',
        'max_cores' => 'string',
        'min_sockets' => 'string',
        'max_sockets' => 'string',
        'min_cores_per_socket' => 'string',
        'max_cores_per_socket' => 'string'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'min_storage' => null,
        'max_storage' => null,
        'min_memory' => 'int64',
        'max_memory' => 'int64',
        'min_cores' => null,
        'max_cores' => null,
        'min_sockets' => null,
        'max_sockets' => null,
        'min_cores_per_socket' => null,
        'max_cores_per_socket' => null
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
        'min_storage' => 'minStorage',
        'max_storage' => 'maxStorage',
        'min_memory' => 'minMemory',
        'max_memory' => 'maxMemory',
        'min_cores' => 'minCores',
        'max_cores' => 'maxCores',
        'min_sockets' => 'minSockets',
        'max_sockets' => 'maxSockets',
        'min_cores_per_socket' => 'minCoresPerSocket',
        'max_cores_per_socket' => 'maxCoresPerSocket'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'min_storage' => 'setMinStorage',
        'max_storage' => 'setMaxStorage',
        'min_memory' => 'setMinMemory',
        'max_memory' => 'setMaxMemory',
        'min_cores' => 'setMinCores',
        'max_cores' => 'setMaxCores',
        'min_sockets' => 'setMinSockets',
        'max_sockets' => 'setMaxSockets',
        'min_cores_per_socket' => 'setMinCoresPerSocket',
        'max_cores_per_socket' => 'setMaxCoresPerSocket'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'min_storage' => 'getMinStorage',
        'max_storage' => 'getMaxStorage',
        'min_memory' => 'getMinMemory',
        'max_memory' => 'getMaxMemory',
        'min_cores' => 'getMinCores',
        'max_cores' => 'getMaxCores',
        'min_sockets' => 'getMinSockets',
        'max_sockets' => 'getMaxSockets',
        'min_cores_per_socket' => 'getMinCoresPerSocket',
        'max_cores_per_socket' => 'getMaxCoresPerSocket'
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
        $this->container['min_storage'] = $data['min_storage'] ?? null;
        $this->container['max_storage'] = $data['max_storage'] ?? null;
        $this->container['min_memory'] = $data['min_memory'] ?? null;
        $this->container['max_memory'] = $data['max_memory'] ?? null;
        $this->container['min_cores'] = $data['min_cores'] ?? null;
        $this->container['max_cores'] = $data['max_cores'] ?? null;
        $this->container['min_sockets'] = $data['min_sockets'] ?? null;
        $this->container['max_sockets'] = $data['max_sockets'] ?? null;
        $this->container['min_cores_per_socket'] = $data['min_cores_per_socket'] ?? null;
        $this->container['max_cores_per_socket'] = $data['max_cores_per_socket'] ?? null;
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
     * Gets min_storage
     *
     * @return string|null
     */
    public function getMinStorage()
    {
        return $this->container['min_storage'];
    }

    /**
     * Sets min_storage
     *
     * @param string|null $min_storage Custom min storage in GB (unless `storageSizeType` set to mb)
     *
     * @return self
     */
    public function setMinStorage($min_storage)
    {
        $this->container['min_storage'] = $min_storage;

        return $this;
    }

    /**
     * Gets max_storage
     *
     * @return string|null
     */
    public function getMaxStorage()
    {
        return $this->container['max_storage'];
    }

    /**
     * Sets max_storage
     *
     * @param string|null $max_storage Custom max storage in GB (unless `storageSizeType` set to mb)
     *
     * @return self
     */
    public function setMaxStorage($max_storage)
    {
        $this->container['max_storage'] = $max_storage;

        return $this;
    }

    /**
     * Gets min_memory
     *
     * @return int|null
     */
    public function getMinMemory()
    {
        return $this->container['min_memory'];
    }

    /**
     * Sets min_memory
     *
     * @param int|null $min_memory Custom min memory in bytes
     *
     * @return self
     */
    public function setMinMemory($min_memory)
    {
        $this->container['min_memory'] = $min_memory;

        return $this;
    }

    /**
     * Gets max_memory
     *
     * @return int|null
     */
    public function getMaxMemory()
    {
        return $this->container['max_memory'];
    }

    /**
     * Sets max_memory
     *
     * @param int|null $max_memory Custom max memory in bytes
     *
     * @return self
     */
    public function setMaxMemory($max_memory)
    {
        $this->container['max_memory'] = $max_memory;

        return $this;
    }

    /**
     * Gets min_cores
     *
     * @return string|null
     */
    public function getMinCores()
    {
        return $this->container['min_cores'];
    }

    /**
     * Sets min_cores
     *
     * @param string|null $min_cores Custom min cores
     *
     * @return self
     */
    public function setMinCores($min_cores)
    {
        $this->container['min_cores'] = $min_cores;

        return $this;
    }

    /**
     * Gets max_cores
     *
     * @return string|null
     */
    public function getMaxCores()
    {
        return $this->container['max_cores'];
    }

    /**
     * Sets max_cores
     *
     * @param string|null $max_cores Custom max cores
     *
     * @return self
     */
    public function setMaxCores($max_cores)
    {
        $this->container['max_cores'] = $max_cores;

        return $this;
    }

    /**
     * Gets min_sockets
     *
     * @return string|null
     */
    public function getMinSockets()
    {
        return $this->container['min_sockets'];
    }

    /**
     * Sets min_sockets
     *
     * @param string|null $min_sockets Custom min sockets
     *
     * @return self
     */
    public function setMinSockets($min_sockets)
    {
        $this->container['min_sockets'] = $min_sockets;

        return $this;
    }

    /**
     * Gets max_sockets
     *
     * @return string|null
     */
    public function getMaxSockets()
    {
        return $this->container['max_sockets'];
    }

    /**
     * Sets max_sockets
     *
     * @param string|null $max_sockets Custom max sockets
     *
     * @return self
     */
    public function setMaxSockets($max_sockets)
    {
        $this->container['max_sockets'] = $max_sockets;

        return $this;
    }

    /**
     * Gets min_cores_per_socket
     *
     * @return string|null
     */
    public function getMinCoresPerSocket()
    {
        return $this->container['min_cores_per_socket'];
    }

    /**
     * Sets min_cores_per_socket
     *
     * @param string|null $min_cores_per_socket Custom min cores allowed per socket
     *
     * @return self
     */
    public function setMinCoresPerSocket($min_cores_per_socket)
    {
        $this->container['min_cores_per_socket'] = $min_cores_per_socket;

        return $this;
    }

    /**
     * Gets max_cores_per_socket
     *
     * @return string|null
     */
    public function getMaxCoresPerSocket()
    {
        return $this->container['max_cores_per_socket'];
    }

    /**
     * Sets max_cores_per_socket
     *
     * @param string|null $max_cores_per_socket Custom max cores allowed per socket
     *
     * @return self
     */
    public function setMaxCoresPerSocket($max_cores_per_socket)
    {
        $this->container['max_cores_per_socket'] = $max_cores_per_socket;

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


