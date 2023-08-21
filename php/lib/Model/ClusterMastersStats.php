<?php
/**
 * ClusterMastersStats
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
 * ClusterMastersStats Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class ClusterMastersStats implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'clusterMasters_stats';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'used_storage' => 'int',
        'reserved_storage' => 'int',
        'max_storage' => 'int',
        'used_memory' => 'int',
        'reserved_memory' => 'int',
        'max_memory' => 'int',
        'cpu_usage' => 'float'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'used_storage' => 'int64',
        'reserved_storage' => 'int64',
        'max_storage' => 'int64',
        'used_memory' => 'int64',
        'reserved_memory' => 'int64',
        'max_memory' => 'int64',
        'cpu_usage' => null
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
        'used_storage' => 'usedStorage',
        'reserved_storage' => 'reservedStorage',
        'max_storage' => 'maxStorage',
        'used_memory' => 'usedMemory',
        'reserved_memory' => 'reservedMemory',
        'max_memory' => 'maxMemory',
        'cpu_usage' => 'cpuUsage'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'used_storage' => 'setUsedStorage',
        'reserved_storage' => 'setReservedStorage',
        'max_storage' => 'setMaxStorage',
        'used_memory' => 'setUsedMemory',
        'reserved_memory' => 'setReservedMemory',
        'max_memory' => 'setMaxMemory',
        'cpu_usage' => 'setCpuUsage'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'used_storage' => 'getUsedStorage',
        'reserved_storage' => 'getReservedStorage',
        'max_storage' => 'getMaxStorage',
        'used_memory' => 'getUsedMemory',
        'reserved_memory' => 'getReservedMemory',
        'max_memory' => 'getMaxMemory',
        'cpu_usage' => 'getCpuUsage'
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
        $this->container['used_storage'] = $data['used_storage'] ?? null;
        $this->container['reserved_storage'] = $data['reserved_storage'] ?? null;
        $this->container['max_storage'] = $data['max_storage'] ?? null;
        $this->container['used_memory'] = $data['used_memory'] ?? null;
        $this->container['reserved_memory'] = $data['reserved_memory'] ?? null;
        $this->container['max_memory'] = $data['max_memory'] ?? null;
        $this->container['cpu_usage'] = $data['cpu_usage'] ?? null;
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
     * Gets used_storage
     *
     * @return int|null
     */
    public function getUsedStorage()
    {
        return $this->container['used_storage'];
    }

    /**
     * Sets used_storage
     *
     * @param int|null $used_storage used_storage
     *
     * @return self
     */
    public function setUsedStorage($used_storage)
    {
        $this->container['used_storage'] = $used_storage;

        return $this;
    }

    /**
     * Gets reserved_storage
     *
     * @return int|null
     */
    public function getReservedStorage()
    {
        return $this->container['reserved_storage'];
    }

    /**
     * Sets reserved_storage
     *
     * @param int|null $reserved_storage reserved_storage
     *
     * @return self
     */
    public function setReservedStorage($reserved_storage)
    {
        $this->container['reserved_storage'] = $reserved_storage;

        return $this;
    }

    /**
     * Gets max_storage
     *
     * @return int|null
     */
    public function getMaxStorage()
    {
        return $this->container['max_storage'];
    }

    /**
     * Sets max_storage
     *
     * @param int|null $max_storage max_storage
     *
     * @return self
     */
    public function setMaxStorage($max_storage)
    {
        $this->container['max_storage'] = $max_storage;

        return $this;
    }

    /**
     * Gets used_memory
     *
     * @return int|null
     */
    public function getUsedMemory()
    {
        return $this->container['used_memory'];
    }

    /**
     * Sets used_memory
     *
     * @param int|null $used_memory used_memory
     *
     * @return self
     */
    public function setUsedMemory($used_memory)
    {
        $this->container['used_memory'] = $used_memory;

        return $this;
    }

    /**
     * Gets reserved_memory
     *
     * @return int|null
     */
    public function getReservedMemory()
    {
        return $this->container['reserved_memory'];
    }

    /**
     * Sets reserved_memory
     *
     * @param int|null $reserved_memory reserved_memory
     *
     * @return self
     */
    public function setReservedMemory($reserved_memory)
    {
        $this->container['reserved_memory'] = $reserved_memory;

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
     * @param int|null $max_memory max_memory
     *
     * @return self
     */
    public function setMaxMemory($max_memory)
    {
        $this->container['max_memory'] = $max_memory;

        return $this;
    }

    /**
     * Gets cpu_usage
     *
     * @return float|null
     */
    public function getCpuUsage()
    {
        return $this->container['cpu_usage'];
    }

    /**
     * Sets cpu_usage
     *
     * @param float|null $cpu_usage cpu_usage
     *
     * @return self
     */
    public function setCpuUsage($cpu_usage)
    {
        $this->container['cpu_usage'] = $cpu_usage;

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

