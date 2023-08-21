<?php
/**
 * ServerStats
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
 * ServerStats Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class ServerStats implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'server_stats';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'ts' => '\DateTime',
        'free_swap' => 'int',
        'used_swap' => 'int',
        'cpu_idle_time' => 'int',
        'cpu_system_time' => 'int',
        'cpu_user_time' => 'int',
        'cpu_total_time' => 'int',
        'max_memory' => 'int',
        'used_memory' => 'int',
        'max_storage' => 'int',
        'used_storage' => 'int',
        'reserved_storage' => 'int',
        'cpu_usage' => 'float',
        'free_memory' => 'int',
        'net_tx_usage' => 'int',
        'net_rx_usage' => 'int',
        'network_bandwidth' => 'int'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'ts' => 'date-time',
        'free_swap' => 'int64',
        'used_swap' => 'int64',
        'cpu_idle_time' => 'int64',
        'cpu_system_time' => 'int64',
        'cpu_user_time' => 'int64',
        'cpu_total_time' => 'int64',
        'max_memory' => 'int64',
        'used_memory' => 'int64',
        'max_storage' => 'int64',
        'used_storage' => 'int64',
        'reserved_storage' => 'int64',
        'cpu_usage' => null,
        'free_memory' => 'int64',
        'net_tx_usage' => 'int64',
        'net_rx_usage' => 'int64',
        'network_bandwidth' => 'int64'
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
        'free_swap' => 'freeSwap',
        'used_swap' => 'usedSwap',
        'cpu_idle_time' => 'cpuIdleTime',
        'cpu_system_time' => 'cpuSystemTime',
        'cpu_user_time' => 'cpuUserTime',
        'cpu_total_time' => 'cpuTotalTime',
        'max_memory' => 'maxMemory',
        'used_memory' => 'usedMemory',
        'max_storage' => 'maxStorage',
        'used_storage' => 'usedStorage',
        'reserved_storage' => 'reservedStorage',
        'cpu_usage' => 'cpuUsage',
        'free_memory' => 'freeMemory',
        'net_tx_usage' => 'netTxUsage',
        'net_rx_usage' => 'netRxUsage',
        'network_bandwidth' => 'networkBandwidth'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'ts' => 'setTs',
        'free_swap' => 'setFreeSwap',
        'used_swap' => 'setUsedSwap',
        'cpu_idle_time' => 'setCpuIdleTime',
        'cpu_system_time' => 'setCpuSystemTime',
        'cpu_user_time' => 'setCpuUserTime',
        'cpu_total_time' => 'setCpuTotalTime',
        'max_memory' => 'setMaxMemory',
        'used_memory' => 'setUsedMemory',
        'max_storage' => 'setMaxStorage',
        'used_storage' => 'setUsedStorage',
        'reserved_storage' => 'setReservedStorage',
        'cpu_usage' => 'setCpuUsage',
        'free_memory' => 'setFreeMemory',
        'net_tx_usage' => 'setNetTxUsage',
        'net_rx_usage' => 'setNetRxUsage',
        'network_bandwidth' => 'setNetworkBandwidth'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'ts' => 'getTs',
        'free_swap' => 'getFreeSwap',
        'used_swap' => 'getUsedSwap',
        'cpu_idle_time' => 'getCpuIdleTime',
        'cpu_system_time' => 'getCpuSystemTime',
        'cpu_user_time' => 'getCpuUserTime',
        'cpu_total_time' => 'getCpuTotalTime',
        'max_memory' => 'getMaxMemory',
        'used_memory' => 'getUsedMemory',
        'max_storage' => 'getMaxStorage',
        'used_storage' => 'getUsedStorage',
        'reserved_storage' => 'getReservedStorage',
        'cpu_usage' => 'getCpuUsage',
        'free_memory' => 'getFreeMemory',
        'net_tx_usage' => 'getNetTxUsage',
        'net_rx_usage' => 'getNetRxUsage',
        'network_bandwidth' => 'getNetworkBandwidth'
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
        $this->container['free_swap'] = $data['free_swap'] ?? null;
        $this->container['used_swap'] = $data['used_swap'] ?? null;
        $this->container['cpu_idle_time'] = $data['cpu_idle_time'] ?? null;
        $this->container['cpu_system_time'] = $data['cpu_system_time'] ?? null;
        $this->container['cpu_user_time'] = $data['cpu_user_time'] ?? null;
        $this->container['cpu_total_time'] = $data['cpu_total_time'] ?? null;
        $this->container['max_memory'] = $data['max_memory'] ?? null;
        $this->container['used_memory'] = $data['used_memory'] ?? null;
        $this->container['max_storage'] = $data['max_storage'] ?? null;
        $this->container['used_storage'] = $data['used_storage'] ?? null;
        $this->container['reserved_storage'] = $data['reserved_storage'] ?? null;
        $this->container['cpu_usage'] = $data['cpu_usage'] ?? null;
        $this->container['free_memory'] = $data['free_memory'] ?? null;
        $this->container['net_tx_usage'] = $data['net_tx_usage'] ?? null;
        $this->container['net_rx_usage'] = $data['net_rx_usage'] ?? null;
        $this->container['network_bandwidth'] = $data['network_bandwidth'] ?? null;
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
     * @return \DateTime|null
     */
    public function getTs()
    {
        return $this->container['ts'];
    }

    /**
     * Sets ts
     *
     * @param \DateTime|null $ts ts
     *
     * @return self
     */
    public function setTs($ts)
    {
        $this->container['ts'] = $ts;

        return $this;
    }

    /**
     * Gets free_swap
     *
     * @return int|null
     */
    public function getFreeSwap()
    {
        return $this->container['free_swap'];
    }

    /**
     * Sets free_swap
     *
     * @param int|null $free_swap free_swap
     *
     * @return self
     */
    public function setFreeSwap($free_swap)
    {
        $this->container['free_swap'] = $free_swap;

        return $this;
    }

    /**
     * Gets used_swap
     *
     * @return int|null
     */
    public function getUsedSwap()
    {
        return $this->container['used_swap'];
    }

    /**
     * Sets used_swap
     *
     * @param int|null $used_swap used_swap
     *
     * @return self
     */
    public function setUsedSwap($used_swap)
    {
        $this->container['used_swap'] = $used_swap;

        return $this;
    }

    /**
     * Gets cpu_idle_time
     *
     * @return int|null
     */
    public function getCpuIdleTime()
    {
        return $this->container['cpu_idle_time'];
    }

    /**
     * Sets cpu_idle_time
     *
     * @param int|null $cpu_idle_time cpu_idle_time
     *
     * @return self
     */
    public function setCpuIdleTime($cpu_idle_time)
    {
        $this->container['cpu_idle_time'] = $cpu_idle_time;

        return $this;
    }

    /**
     * Gets cpu_system_time
     *
     * @return int|null
     */
    public function getCpuSystemTime()
    {
        return $this->container['cpu_system_time'];
    }

    /**
     * Sets cpu_system_time
     *
     * @param int|null $cpu_system_time cpu_system_time
     *
     * @return self
     */
    public function setCpuSystemTime($cpu_system_time)
    {
        $this->container['cpu_system_time'] = $cpu_system_time;

        return $this;
    }

    /**
     * Gets cpu_user_time
     *
     * @return int|null
     */
    public function getCpuUserTime()
    {
        return $this->container['cpu_user_time'];
    }

    /**
     * Sets cpu_user_time
     *
     * @param int|null $cpu_user_time cpu_user_time
     *
     * @return self
     */
    public function setCpuUserTime($cpu_user_time)
    {
        $this->container['cpu_user_time'] = $cpu_user_time;

        return $this;
    }

    /**
     * Gets cpu_total_time
     *
     * @return int|null
     */
    public function getCpuTotalTime()
    {
        return $this->container['cpu_total_time'];
    }

    /**
     * Sets cpu_total_time
     *
     * @param int|null $cpu_total_time cpu_total_time
     *
     * @return self
     */
    public function setCpuTotalTime($cpu_total_time)
    {
        $this->container['cpu_total_time'] = $cpu_total_time;

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
     * Gets free_memory
     *
     * @return int|null
     */
    public function getFreeMemory()
    {
        return $this->container['free_memory'];
    }

    /**
     * Sets free_memory
     *
     * @param int|null $free_memory free_memory
     *
     * @return self
     */
    public function setFreeMemory($free_memory)
    {
        $this->container['free_memory'] = $free_memory;

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
     * Gets network_bandwidth
     *
     * @return int|null
     */
    public function getNetworkBandwidth()
    {
        return $this->container['network_bandwidth'];
    }

    /**
     * Sets network_bandwidth
     *
     * @param int|null $network_bandwidth network_bandwidth
     *
     * @return self
     */
    public function setNetworkBandwidth($network_bandwidth)
    {
        $this->container['network_bandwidth'] = $network_bandwidth;

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


