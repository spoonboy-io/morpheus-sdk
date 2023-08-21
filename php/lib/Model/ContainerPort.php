<?php
/**
 * ContainerPort
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
 * ContainerPort Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class ContainerPort implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'containerPort';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'id' => 'int',
        'index' => 'int',
        'external' => 'int',
        'internal' => 'int',
        'display_name' => 'string',
        'primary_port' => 'bool',
        'export' => 'bool',
        'visible' => 'bool',
        'export_name' => 'string',
        'load_balance_protocol' => 'string',
        'load_balance' => 'bool',
        'protocol' => 'string',
        'link' => 'bool',
        'external_ip' => 'string',
        'internal_ip' => 'string'
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
        'index' => 'int64',
        'external' => 'int64',
        'internal' => 'int64',
        'display_name' => null,
        'primary_port' => null,
        'export' => null,
        'visible' => null,
        'export_name' => null,
        'load_balance_protocol' => null,
        'load_balance' => null,
        'protocol' => null,
        'link' => null,
        'external_ip' => null,
        'internal_ip' => null
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
        'index' => 'index',
        'external' => 'external',
        'internal' => 'internal',
        'display_name' => 'displayName',
        'primary_port' => 'primaryPort',
        'export' => 'export',
        'visible' => 'visible',
        'export_name' => 'exportName',
        'load_balance_protocol' => 'loadBalanceProtocol',
        'load_balance' => 'loadBalance',
        'protocol' => 'protocol',
        'link' => 'link',
        'external_ip' => 'externalIp',
        'internal_ip' => 'internalIp'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'id' => 'setId',
        'index' => 'setIndex',
        'external' => 'setExternal',
        'internal' => 'setInternal',
        'display_name' => 'setDisplayName',
        'primary_port' => 'setPrimaryPort',
        'export' => 'setExport',
        'visible' => 'setVisible',
        'export_name' => 'setExportName',
        'load_balance_protocol' => 'setLoadBalanceProtocol',
        'load_balance' => 'setLoadBalance',
        'protocol' => 'setProtocol',
        'link' => 'setLink',
        'external_ip' => 'setExternalIp',
        'internal_ip' => 'setInternalIp'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'id' => 'getId',
        'index' => 'getIndex',
        'external' => 'getExternal',
        'internal' => 'getInternal',
        'display_name' => 'getDisplayName',
        'primary_port' => 'getPrimaryPort',
        'export' => 'getExport',
        'visible' => 'getVisible',
        'export_name' => 'getExportName',
        'load_balance_protocol' => 'getLoadBalanceProtocol',
        'load_balance' => 'getLoadBalance',
        'protocol' => 'getProtocol',
        'link' => 'getLink',
        'external_ip' => 'getExternalIp',
        'internal_ip' => 'getInternalIp'
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
        $this->container['index'] = $data['index'] ?? null;
        $this->container['external'] = $data['external'] ?? null;
        $this->container['internal'] = $data['internal'] ?? null;
        $this->container['display_name'] = $data['display_name'] ?? null;
        $this->container['primary_port'] = $data['primary_port'] ?? null;
        $this->container['export'] = $data['export'] ?? null;
        $this->container['visible'] = $data['visible'] ?? null;
        $this->container['export_name'] = $data['export_name'] ?? null;
        $this->container['load_balance_protocol'] = $data['load_balance_protocol'] ?? null;
        $this->container['load_balance'] = $data['load_balance'] ?? null;
        $this->container['protocol'] = $data['protocol'] ?? null;
        $this->container['link'] = $data['link'] ?? null;
        $this->container['external_ip'] = $data['external_ip'] ?? null;
        $this->container['internal_ip'] = $data['internal_ip'] ?? null;
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
     * Gets index
     *
     * @return int|null
     */
    public function getIndex()
    {
        return $this->container['index'];
    }

    /**
     * Sets index
     *
     * @param int|null $index index
     *
     * @return self
     */
    public function setIndex($index)
    {
        $this->container['index'] = $index;

        return $this;
    }

    /**
     * Gets external
     *
     * @return int|null
     */
    public function getExternal()
    {
        return $this->container['external'];
    }

    /**
     * Sets external
     *
     * @param int|null $external external
     *
     * @return self
     */
    public function setExternal($external)
    {
        $this->container['external'] = $external;

        return $this;
    }

    /**
     * Gets internal
     *
     * @return int|null
     */
    public function getInternal()
    {
        return $this->container['internal'];
    }

    /**
     * Sets internal
     *
     * @param int|null $internal internal
     *
     * @return self
     */
    public function setInternal($internal)
    {
        $this->container['internal'] = $internal;

        return $this;
    }

    /**
     * Gets display_name
     *
     * @return string|null
     */
    public function getDisplayName()
    {
        return $this->container['display_name'];
    }

    /**
     * Sets display_name
     *
     * @param string|null $display_name display_name
     *
     * @return self
     */
    public function setDisplayName($display_name)
    {
        $this->container['display_name'] = $display_name;

        return $this;
    }

    /**
     * Gets primary_port
     *
     * @return bool|null
     */
    public function getPrimaryPort()
    {
        return $this->container['primary_port'];
    }

    /**
     * Sets primary_port
     *
     * @param bool|null $primary_port primary_port
     *
     * @return self
     */
    public function setPrimaryPort($primary_port)
    {
        $this->container['primary_port'] = $primary_port;

        return $this;
    }

    /**
     * Gets export
     *
     * @return bool|null
     */
    public function getExport()
    {
        return $this->container['export'];
    }

    /**
     * Sets export
     *
     * @param bool|null $export export
     *
     * @return self
     */
    public function setExport($export)
    {
        $this->container['export'] = $export;

        return $this;
    }

    /**
     * Gets visible
     *
     * @return bool|null
     */
    public function getVisible()
    {
        return $this->container['visible'];
    }

    /**
     * Sets visible
     *
     * @param bool|null $visible visible
     *
     * @return self
     */
    public function setVisible($visible)
    {
        $this->container['visible'] = $visible;

        return $this;
    }

    /**
     * Gets export_name
     *
     * @return string|null
     */
    public function getExportName()
    {
        return $this->container['export_name'];
    }

    /**
     * Sets export_name
     *
     * @param string|null $export_name export_name
     *
     * @return self
     */
    public function setExportName($export_name)
    {
        $this->container['export_name'] = $export_name;

        return $this;
    }

    /**
     * Gets load_balance_protocol
     *
     * @return string|null
     */
    public function getLoadBalanceProtocol()
    {
        return $this->container['load_balance_protocol'];
    }

    /**
     * Sets load_balance_protocol
     *
     * @param string|null $load_balance_protocol load_balance_protocol
     *
     * @return self
     */
    public function setLoadBalanceProtocol($load_balance_protocol)
    {
        $this->container['load_balance_protocol'] = $load_balance_protocol;

        return $this;
    }

    /**
     * Gets load_balance
     *
     * @return bool|null
     */
    public function getLoadBalance()
    {
        return $this->container['load_balance'];
    }

    /**
     * Sets load_balance
     *
     * @param bool|null $load_balance load_balance
     *
     * @return self
     */
    public function setLoadBalance($load_balance)
    {
        $this->container['load_balance'] = $load_balance;

        return $this;
    }

    /**
     * Gets protocol
     *
     * @return string|null
     */
    public function getProtocol()
    {
        return $this->container['protocol'];
    }

    /**
     * Sets protocol
     *
     * @param string|null $protocol protocol
     *
     * @return self
     */
    public function setProtocol($protocol)
    {
        $this->container['protocol'] = $protocol;

        return $this;
    }

    /**
     * Gets link
     *
     * @return bool|null
     */
    public function getLink()
    {
        return $this->container['link'];
    }

    /**
     * Sets link
     *
     * @param bool|null $link link
     *
     * @return self
     */
    public function setLink($link)
    {
        $this->container['link'] = $link;

        return $this;
    }

    /**
     * Gets external_ip
     *
     * @return string|null
     */
    public function getExternalIp()
    {
        return $this->container['external_ip'];
    }

    /**
     * Sets external_ip
     *
     * @param string|null $external_ip external_ip
     *
     * @return self
     */
    public function setExternalIp($external_ip)
    {
        $this->container['external_ip'] = $external_ip;

        return $this;
    }

    /**
     * Gets internal_ip
     *
     * @return string|null
     */
    public function getInternalIp()
    {
        return $this->container['internal_ip'];
    }

    /**
     * Sets internal_ip
     *
     * @param string|null $internal_ip internal_ip
     *
     * @return self
     */
    public function setInternalIp($internal_ip)
    {
        $this->container['internal_ip'] = $internal_ip;

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

