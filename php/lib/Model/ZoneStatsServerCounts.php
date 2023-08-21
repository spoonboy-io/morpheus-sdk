<?php
/**
 * ZoneStatsServerCounts
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
 * ZoneStatsServerCounts Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class ZoneStatsServerCounts implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'zone_stats_serverCounts';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'all' => 'int',
        'host' => 'int',
        'hypervisor' => 'int',
        'container_host' => 'int',
        'vm' => 'int',
        'baremetal' => 'int',
        'unmanaged' => 'int'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'all' => 'int64',
        'host' => 'int64',
        'hypervisor' => 'int64',
        'container_host' => 'int64',
        'vm' => 'int64',
        'baremetal' => 'int64',
        'unmanaged' => 'int64'
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
        'all' => 'all',
        'host' => 'host',
        'hypervisor' => 'hypervisor',
        'container_host' => 'containerHost',
        'vm' => 'vm',
        'baremetal' => 'baremetal',
        'unmanaged' => 'unmanaged'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'all' => 'setAll',
        'host' => 'setHost',
        'hypervisor' => 'setHypervisor',
        'container_host' => 'setContainerHost',
        'vm' => 'setVm',
        'baremetal' => 'setBaremetal',
        'unmanaged' => 'setUnmanaged'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'all' => 'getAll',
        'host' => 'getHost',
        'hypervisor' => 'getHypervisor',
        'container_host' => 'getContainerHost',
        'vm' => 'getVm',
        'baremetal' => 'getBaremetal',
        'unmanaged' => 'getUnmanaged'
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
        $this->container['all'] = $data['all'] ?? null;
        $this->container['host'] = $data['host'] ?? null;
        $this->container['hypervisor'] = $data['hypervisor'] ?? null;
        $this->container['container_host'] = $data['container_host'] ?? null;
        $this->container['vm'] = $data['vm'] ?? null;
        $this->container['baremetal'] = $data['baremetal'] ?? null;
        $this->container['unmanaged'] = $data['unmanaged'] ?? null;
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
     * Gets all
     *
     * @return int|null
     */
    public function getAll()
    {
        return $this->container['all'];
    }

    /**
     * Sets all
     *
     * @param int|null $all all
     *
     * @return self
     */
    public function setAll($all)
    {
        $this->container['all'] = $all;

        return $this;
    }

    /**
     * Gets host
     *
     * @return int|null
     */
    public function getHost()
    {
        return $this->container['host'];
    }

    /**
     * Sets host
     *
     * @param int|null $host host
     *
     * @return self
     */
    public function setHost($host)
    {
        $this->container['host'] = $host;

        return $this;
    }

    /**
     * Gets hypervisor
     *
     * @return int|null
     */
    public function getHypervisor()
    {
        return $this->container['hypervisor'];
    }

    /**
     * Sets hypervisor
     *
     * @param int|null $hypervisor hypervisor
     *
     * @return self
     */
    public function setHypervisor($hypervisor)
    {
        $this->container['hypervisor'] = $hypervisor;

        return $this;
    }

    /**
     * Gets container_host
     *
     * @return int|null
     */
    public function getContainerHost()
    {
        return $this->container['container_host'];
    }

    /**
     * Sets container_host
     *
     * @param int|null $container_host container_host
     *
     * @return self
     */
    public function setContainerHost($container_host)
    {
        $this->container['container_host'] = $container_host;

        return $this;
    }

    /**
     * Gets vm
     *
     * @return int|null
     */
    public function getVm()
    {
        return $this->container['vm'];
    }

    /**
     * Sets vm
     *
     * @param int|null $vm vm
     *
     * @return self
     */
    public function setVm($vm)
    {
        $this->container['vm'] = $vm;

        return $this;
    }

    /**
     * Gets baremetal
     *
     * @return int|null
     */
    public function getBaremetal()
    {
        return $this->container['baremetal'];
    }

    /**
     * Sets baremetal
     *
     * @param int|null $baremetal baremetal
     *
     * @return self
     */
    public function setBaremetal($baremetal)
    {
        $this->container['baremetal'] = $baremetal;

        return $this;
    }

    /**
     * Gets unmanaged
     *
     * @return int|null
     */
    public function getUnmanaged()
    {
        return $this->container['unmanaged'];
    }

    /**
     * Sets unmanaged
     *
     * @param int|null $unmanaged unmanaged
     *
     * @return self
     */
    public function setUnmanaged($unmanaged)
    {
        $this->container['unmanaged'] = $unmanaged;

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


