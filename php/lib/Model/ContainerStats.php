<?php
/**
 * ContainerStats
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
 * ContainerStats Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class ContainerStats implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'container_stats';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'ts' => 'string',
        'running' => 'bool',
        'user_cpu_usage' => 'float',
        'system_cpu_usage' => 'float',
        'used_memory' => 'int',
        'max_memory' => 'int',
        'cache_memory' => 'int',
        'max_storage' => 'int',
        'used_storage' => 'int',
        'read_iops' => 'int',
        'write_iops' => 'float',
        'total_iops' => 'float',
        'net_tx_usage' => 'int',
        'net_rx_usage' => 'int'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'ts' => null,
        'running' => null,
        'user_cpu_usage' => null,
        'system_cpu_usage' => null,
        'used_memory' => null,
        'max_memory' => null,
        'cache_memory' => null,
        'max_storage' => null,
        'used_storage' => null,
        'read_iops' => null,
        'write_iops' => null,
        'total_iops' => null,
        'net_tx_usage' => null,
        'net_rx_usage' => null
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
        'ts' => 'ts',
        'running' => 'running',
        'user_cpu_usage' => 'userCpuUsage',
        'system_cpu_usage' => 'systemCpuUsage',
        'used_memory' => 'usedMemory',
        'max_memory' => 'maxMemory',
        'cache_memory' => 'cacheMemory',
        'max_storage' => 'maxStorage',
        'used_storage' => 'usedStorage',
        'read_iops' => 'readIOPS',
        'write_iops' => 'writeIOPS',
        'total_iops' => 'totalIOPS',
        'net_tx_usage' => 'netTxUsage',
        'net_rx_usage' => 'netRxUsage'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'ts' => 'setTs',
        'running' => 'setRunning',
        'user_cpu_usage' => 'setUserCpuUsage',
        'system_cpu_usage' => 'setSystemCpuUsage',
        'used_memory' => 'setUsedMemory',
        'max_memory' => 'setMaxMemory',
        'cache_memory' => 'setCacheMemory',
        'max_storage' => 'setMaxStorage',
        'used_storage' => 'setUsedStorage',
        'read_iops' => 'setReadIops',
        'write_iops' => 'setWriteIops',
        'total_iops' => 'setTotalIops',
        'net_tx_usage' => 'setNetTxUsage',
        'net_rx_usage' => 'setNetRxUsage'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'ts' => 'getTs',
        'running' => 'getRunning',
        'user_cpu_usage' => 'getUserCpuUsage',
        'system_cpu_usage' => 'getSystemCpuUsage',
        'used_memory' => 'getUsedMemory',
        'max_memory' => 'getMaxMemory',
        'cache_memory' => 'getCacheMemory',
        'max_storage' => 'getMaxStorage',
        'used_storage' => 'getUsedStorage',
        'read_iops' => 'getReadIops',
        'write_iops' => 'getWriteIops',
        'total_iops' => 'getTotalIops',
        'net_tx_usage' => 'getNetTxUsage',
        'net_rx_usage' => 'getNetRxUsage'
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
        $this->container['ts'] = $data['ts'] ?? null;
        $this->container['running'] = $data['running'] ?? null;
        $this->container['user_cpu_usage'] = $data['user_cpu_usage'] ?? null;
        $this->container['system_cpu_usage'] = $data['system_cpu_usage'] ?? null;
        $this->container['used_memory'] = $data['used_memory'] ?? null;
        $this->container['max_memory'] = $data['max_memory'] ?? null;
        $this->container['cache_memory'] = $data['cache_memory'] ?? null;
        $this->container['max_storage'] = $data['max_storage'] ?? null;
        $this->container['used_storage'] = $data['used_storage'] ?? null;
        $this->container['read_iops'] = $data['read_iops'] ?? null;
        $this->container['write_iops'] = $data['write_iops'] ?? null;
        $this->container['total_iops'] = $data['total_iops'] ?? null;
        $this->container['net_tx_usage'] = $data['net_tx_usage'] ?? null;
        $this->container['net_rx_usage'] = $data['net_rx_usage'] ?? null;
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
     * Gets ts
     *
     * @return string|null
     */
    public function getTs()
    {
        return $this->container['ts'];
    }

    /**
     * Sets ts
     *
     * @param string|null $ts ts
     *
     * @return self
     */
    public function setTs($ts)
    {
        $this->container['ts'] = $ts;

        return $this;
    }

    /**
     * Gets running
     *
     * @return bool|null
     */
    public function getRunning()
    {
        return $this->container['running'];
    }

    /**
     * Sets running
     *
     * @param bool|null $running running
     *
     * @return self
     */
    public function setRunning($running)
    {
        $this->container['running'] = $running;

        return $this;
    }

    /**
     * Gets user_cpu_usage
     *
     * @return float|null
     */
    public function getUserCpuUsage()
    {
        return $this->container['user_cpu_usage'];
    }

    /**
     * Sets user_cpu_usage
     *
     * @param float|null $user_cpu_usage user_cpu_usage
     *
     * @return self
     */
    public function setUserCpuUsage($user_cpu_usage)
    {
        $this->container['user_cpu_usage'] = $user_cpu_usage;

        return $this;
    }

    /**
     * Gets system_cpu_usage
     *
     * @return float|null
     */
    public function getSystemCpuUsage()
    {
        return $this->container['system_cpu_usage'];
    }

    /**
     * Sets system_cpu_usage
     *
     * @param float|null $system_cpu_usage system_cpu_usage
     *
     * @return self
     */
    public function setSystemCpuUsage($system_cpu_usage)
    {
        $this->container['system_cpu_usage'] = $system_cpu_usage;

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
     * Gets cache_memory
     *
     * @return int|null
     */
    public function getCacheMemory()
    {
        return $this->container['cache_memory'];
    }

    /**
     * Sets cache_memory
     *
     * @param int|null $cache_memory cache_memory
     *
     * @return self
     */
    public function setCacheMemory($cache_memory)
    {
        $this->container['cache_memory'] = $cache_memory;

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
     * Gets read_iops
     *
     * @return int|null
     */
    public function getReadIops()
    {
        return $this->container['read_iops'];
    }

    /**
     * Sets read_iops
     *
     * @param int|null $read_iops read_iops
     *
     * @return self
     */
    public function setReadIops($read_iops)
    {
        $this->container['read_iops'] = $read_iops;

        return $this;
    }

    /**
     * Gets write_iops
     *
     * @return float|null
     */
    public function getWriteIops()
    {
        return $this->container['write_iops'];
    }

    /**
     * Sets write_iops
     *
     * @param float|null $write_iops write_iops
     *
     * @return self
     */
    public function setWriteIops($write_iops)
    {
        $this->container['write_iops'] = $write_iops;

        return $this;
    }

    /**
     * Gets total_iops
     *
     * @return float|null
     */
    public function getTotalIops()
    {
        return $this->container['total_iops'];
    }

    /**
     * Sets total_iops
     *
     * @param float|null $total_iops total_iops
     *
     * @return self
     */
    public function setTotalIops($total_iops)
    {
        $this->container['total_iops'] = $total_iops;

        return $this;
    }

    /**
     * Gets net_tx_usage
     *
     * @return int|null
     */
    public function getNetTxUsage()
    {
        return $this->container['net_tx_usage'];
    }

    /**
     * Sets net_tx_usage
     *
     * @param int|null $net_tx_usage net_tx_usage
     *
     * @return self
     */
    public function setNetTxUsage($net_tx_usage)
    {
        $this->container['net_tx_usage'] = $net_tx_usage;

        return $this;
    }

    /**
     * Gets net_rx_usage
     *
     * @return int|null
     */
    public function getNetRxUsage()
    {
        return $this->container['net_rx_usage'];
    }

    /**
     * Sets net_rx_usage
     *
     * @param int|null $net_rx_usage net_rx_usage
     *
     * @return self
     */
    public function setNetRxUsage($net_rx_usage)
    {
        $this->container['net_rx_usage'] = $net_rx_usage;

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


