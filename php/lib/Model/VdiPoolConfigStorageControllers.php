<?php
/**
 * VdiPoolConfigStorageControllers
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
 * VdiPoolConfigStorageControllers Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class VdiPoolConfigStorageControllers implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'vdiPool_config_storageControllers';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'id' => 'int',
        'name' => 'string',
        'active' => 'bool',
        'type_id' => 'int',
        'type_name' => 'string',
        'unit_number' => 'string',
        'bus_number' => 'string',
        'max_devices' => 'float',
        'removable' => 'bool',
        'editable' => 'bool',
        'reserved_unit_number' => 'float',
        'category' => 'string',
        'display_order' => 'float'
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
        'name' => null,
        'active' => null,
        'type_id' => 'int64',
        'type_name' => null,
        'unit_number' => null,
        'bus_number' => null,
        'max_devices' => null,
        'removable' => null,
        'editable' => null,
        'reserved_unit_number' => null,
        'category' => null,
        'display_order' => null
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
        'name' => 'name',
        'active' => 'active',
        'type_id' => 'typeId',
        'type_name' => 'typeName',
        'unit_number' => 'unitNumber',
        'bus_number' => 'busNumber',
        'max_devices' => 'maxDevices',
        'removable' => 'removable',
        'editable' => 'editable',
        'reserved_unit_number' => 'reservedUnitNumber',
        'category' => 'category',
        'display_order' => 'displayOrder'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'id' => 'setId',
        'name' => 'setName',
        'active' => 'setActive',
        'type_id' => 'setTypeId',
        'type_name' => 'setTypeName',
        'unit_number' => 'setUnitNumber',
        'bus_number' => 'setBusNumber',
        'max_devices' => 'setMaxDevices',
        'removable' => 'setRemovable',
        'editable' => 'setEditable',
        'reserved_unit_number' => 'setReservedUnitNumber',
        'category' => 'setCategory',
        'display_order' => 'setDisplayOrder'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'id' => 'getId',
        'name' => 'getName',
        'active' => 'getActive',
        'type_id' => 'getTypeId',
        'type_name' => 'getTypeName',
        'unit_number' => 'getUnitNumber',
        'bus_number' => 'getBusNumber',
        'max_devices' => 'getMaxDevices',
        'removable' => 'getRemovable',
        'editable' => 'getEditable',
        'reserved_unit_number' => 'getReservedUnitNumber',
        'category' => 'getCategory',
        'display_order' => 'getDisplayOrder'
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
        $this->container['name'] = $data['name'] ?? null;
        $this->container['active'] = $data['active'] ?? null;
        $this->container['type_id'] = $data['type_id'] ?? null;
        $this->container['type_name'] = $data['type_name'] ?? null;
        $this->container['unit_number'] = $data['unit_number'] ?? null;
        $this->container['bus_number'] = $data['bus_number'] ?? null;
        $this->container['max_devices'] = $data['max_devices'] ?? null;
        $this->container['removable'] = $data['removable'] ?? null;
        $this->container['editable'] = $data['editable'] ?? null;
        $this->container['reserved_unit_number'] = $data['reserved_unit_number'] ?? null;
        $this->container['category'] = $data['category'] ?? null;
        $this->container['display_order'] = $data['display_order'] ?? null;
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
     * Gets type_id
     *
     * @return int|null
     */
    public function getTypeId()
    {
        return $this->container['type_id'];
    }

    /**
     * Sets type_id
     *
     * @param int|null $type_id type_id
     *
     * @return self
     */
    public function setTypeId($type_id)
    {
        $this->container['type_id'] = $type_id;

        return $this;
    }

    /**
     * Gets type_name
     *
     * @return string|null
     */
    public function getTypeName()
    {
        return $this->container['type_name'];
    }

    /**
     * Sets type_name
     *
     * @param string|null $type_name type_name
     *
     * @return self
     */
    public function setTypeName($type_name)
    {
        $this->container['type_name'] = $type_name;

        return $this;
    }

    /**
     * Gets unit_number
     *
     * @return string|null
     */
    public function getUnitNumber()
    {
        return $this->container['unit_number'];
    }

    /**
     * Sets unit_number
     *
     * @param string|null $unit_number unit_number
     *
     * @return self
     */
    public function setUnitNumber($unit_number)
    {
        $this->container['unit_number'] = $unit_number;

        return $this;
    }

    /**
     * Gets bus_number
     *
     * @return string|null
     */
    public function getBusNumber()
    {
        return $this->container['bus_number'];
    }

    /**
     * Sets bus_number
     *
     * @param string|null $bus_number bus_number
     *
     * @return self
     */
    public function setBusNumber($bus_number)
    {
        $this->container['bus_number'] = $bus_number;

        return $this;
    }

    /**
     * Gets max_devices
     *
     * @return float|null
     */
    public function getMaxDevices()
    {
        return $this->container['max_devices'];
    }

    /**
     * Sets max_devices
     *
     * @param float|null $max_devices max_devices
     *
     * @return self
     */
    public function setMaxDevices($max_devices)
    {
        $this->container['max_devices'] = $max_devices;

        return $this;
    }

    /**
     * Gets removable
     *
     * @return bool|null
     */
    public function getRemovable()
    {
        return $this->container['removable'];
    }

    /**
     * Sets removable
     *
     * @param bool|null $removable removable
     *
     * @return self
     */
    public function setRemovable($removable)
    {
        $this->container['removable'] = $removable;

        return $this;
    }

    /**
     * Gets editable
     *
     * @return bool|null
     */
    public function getEditable()
    {
        return $this->container['editable'];
    }

    /**
     * Sets editable
     *
     * @param bool|null $editable editable
     *
     * @return self
     */
    public function setEditable($editable)
    {
        $this->container['editable'] = $editable;

        return $this;
    }

    /**
     * Gets reserved_unit_number
     *
     * @return float|null
     */
    public function getReservedUnitNumber()
    {
        return $this->container['reserved_unit_number'];
    }

    /**
     * Sets reserved_unit_number
     *
     * @param float|null $reserved_unit_number reserved_unit_number
     *
     * @return self
     */
    public function setReservedUnitNumber($reserved_unit_number)
    {
        $this->container['reserved_unit_number'] = $reserved_unit_number;

        return $this;
    }

    /**
     * Gets category
     *
     * @return string|null
     */
    public function getCategory()
    {
        return $this->container['category'];
    }

    /**
     * Sets category
     *
     * @param string|null $category category
     *
     * @return self
     */
    public function setCategory($category)
    {
        $this->container['category'] = $category;

        return $this;
    }

    /**
     * Gets display_order
     *
     * @return float|null
     */
    public function getDisplayOrder()
    {
        return $this->container['display_order'];
    }

    /**
     * Sets display_order
     *
     * @param float|null $display_order display_order
     *
     * @return self
     */
    public function setDisplayOrder($display_order)
    {
        $this->container['display_order'] = $display_order;

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


