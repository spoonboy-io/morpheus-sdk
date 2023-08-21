<?php
/**
 * ZoneTypeServerTypes
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
 * ZoneTypeServerTypes Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class ZoneTypeServerTypes implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'zoneType_serverTypes';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'id' => 'int',
        'code' => 'string',
        'name' => 'string',
        'description' => 'string',
        'node_type' => 'string',
        'platform' => 'string',
        'enabled' => 'bool',
        'selectable' => 'bool',
        'external_delete' => 'bool',
        'managed' => 'bool',
        'control_power' => 'bool',
        'control_suspend' => 'bool',
        'creatable' => 'bool',
        'has_agent' => 'bool',
        'vm_hypervisor' => 'bool',
        'container_hypervisor' => 'bool',
        'bare_metal_host' => 'bool',
        'guest_vm' => 'bool',
        'has_automation' => 'bool',
        'provision_type' => '\OpenAPI\Client\Model\ZoneTypeProvisionType',
        'option_types' => '\OpenAPI\Client\Model\ZoneTypeOptionTypes[]',
        'display_order' => 'int'
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
        'code' => null,
        'name' => null,
        'description' => null,
        'node_type' => null,
        'platform' => null,
        'enabled' => null,
        'selectable' => null,
        'external_delete' => null,
        'managed' => null,
        'control_power' => null,
        'control_suspend' => null,
        'creatable' => null,
        'has_agent' => null,
        'vm_hypervisor' => null,
        'container_hypervisor' => null,
        'bare_metal_host' => null,
        'guest_vm' => null,
        'has_automation' => null,
        'provision_type' => null,
        'option_types' => null,
        'display_order' => 'int64'
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
        'code' => 'code',
        'name' => 'name',
        'description' => 'description',
        'node_type' => 'nodeType',
        'platform' => 'platform',
        'enabled' => 'enabled',
        'selectable' => 'selectable',
        'external_delete' => 'externalDelete',
        'managed' => 'managed',
        'control_power' => 'controlPower',
        'control_suspend' => 'controlSuspend',
        'creatable' => 'creatable',
        'has_agent' => 'hasAgent',
        'vm_hypervisor' => 'vmHypervisor',
        'container_hypervisor' => 'containerHypervisor',
        'bare_metal_host' => 'bareMetalHost',
        'guest_vm' => 'guestVm',
        'has_automation' => 'hasAutomation',
        'provision_type' => 'provisionType',
        'option_types' => 'optionTypes',
        'display_order' => 'displayOrder'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'id' => 'setId',
        'code' => 'setCode',
        'name' => 'setName',
        'description' => 'setDescription',
        'node_type' => 'setNodeType',
        'platform' => 'setPlatform',
        'enabled' => 'setEnabled',
        'selectable' => 'setSelectable',
        'external_delete' => 'setExternalDelete',
        'managed' => 'setManaged',
        'control_power' => 'setControlPower',
        'control_suspend' => 'setControlSuspend',
        'creatable' => 'setCreatable',
        'has_agent' => 'setHasAgent',
        'vm_hypervisor' => 'setVmHypervisor',
        'container_hypervisor' => 'setContainerHypervisor',
        'bare_metal_host' => 'setBareMetalHost',
        'guest_vm' => 'setGuestVm',
        'has_automation' => 'setHasAutomation',
        'provision_type' => 'setProvisionType',
        'option_types' => 'setOptionTypes',
        'display_order' => 'setDisplayOrder'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'id' => 'getId',
        'code' => 'getCode',
        'name' => 'getName',
        'description' => 'getDescription',
        'node_type' => 'getNodeType',
        'platform' => 'getPlatform',
        'enabled' => 'getEnabled',
        'selectable' => 'getSelectable',
        'external_delete' => 'getExternalDelete',
        'managed' => 'getManaged',
        'control_power' => 'getControlPower',
        'control_suspend' => 'getControlSuspend',
        'creatable' => 'getCreatable',
        'has_agent' => 'getHasAgent',
        'vm_hypervisor' => 'getVmHypervisor',
        'container_hypervisor' => 'getContainerHypervisor',
        'bare_metal_host' => 'getBareMetalHost',
        'guest_vm' => 'getGuestVm',
        'has_automation' => 'getHasAutomation',
        'provision_type' => 'getProvisionType',
        'option_types' => 'getOptionTypes',
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
        $this->container['code'] = $data['code'] ?? null;
        $this->container['name'] = $data['name'] ?? null;
        $this->container['description'] = $data['description'] ?? null;
        $this->container['node_type'] = $data['node_type'] ?? null;
        $this->container['platform'] = $data['platform'] ?? null;
        $this->container['enabled'] = $data['enabled'] ?? null;
        $this->container['selectable'] = $data['selectable'] ?? null;
        $this->container['external_delete'] = $data['external_delete'] ?? null;
        $this->container['managed'] = $data['managed'] ?? null;
        $this->container['control_power'] = $data['control_power'] ?? null;
        $this->container['control_suspend'] = $data['control_suspend'] ?? null;
        $this->container['creatable'] = $data['creatable'] ?? null;
        $this->container['has_agent'] = $data['has_agent'] ?? null;
        $this->container['vm_hypervisor'] = $data['vm_hypervisor'] ?? null;
        $this->container['container_hypervisor'] = $data['container_hypervisor'] ?? null;
        $this->container['bare_metal_host'] = $data['bare_metal_host'] ?? null;
        $this->container['guest_vm'] = $data['guest_vm'] ?? null;
        $this->container['has_automation'] = $data['has_automation'] ?? null;
        $this->container['provision_type'] = $data['provision_type'] ?? null;
        $this->container['option_types'] = $data['option_types'] ?? null;
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
     * Gets code
     *
     * @return string|null
     */
    public function getCode()
    {
        return $this->container['code'];
    }

    /**
     * Sets code
     *
     * @param string|null $code code
     *
     * @return self
     */
    public function setCode($code)
    {
        $this->container['code'] = $code;

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
     * Gets node_type
     *
     * @return string|null
     */
    public function getNodeType()
    {
        return $this->container['node_type'];
    }

    /**
     * Sets node_type
     *
     * @param string|null $node_type node_type
     *
     * @return self
     */
    public function setNodeType($node_type)
    {
        $this->container['node_type'] = $node_type;

        return $this;
    }

    /**
     * Gets platform
     *
     * @return string|null
     */
    public function getPlatform()
    {
        return $this->container['platform'];
    }

    /**
     * Sets platform
     *
     * @param string|null $platform platform
     *
     * @return self
     */
    public function setPlatform($platform)
    {
        $this->container['platform'] = $platform;

        return $this;
    }

    /**
     * Gets enabled
     *
     * @return bool|null
     */
    public function getEnabled()
    {
        return $this->container['enabled'];
    }

    /**
     * Sets enabled
     *
     * @param bool|null $enabled enabled
     *
     * @return self
     */
    public function setEnabled($enabled)
    {
        $this->container['enabled'] = $enabled;

        return $this;
    }

    /**
     * Gets selectable
     *
     * @return bool|null
     */
    public function getSelectable()
    {
        return $this->container['selectable'];
    }

    /**
     * Sets selectable
     *
     * @param bool|null $selectable selectable
     *
     * @return self
     */
    public function setSelectable($selectable)
    {
        $this->container['selectable'] = $selectable;

        return $this;
    }

    /**
     * Gets external_delete
     *
     * @return bool|null
     */
    public function getExternalDelete()
    {
        return $this->container['external_delete'];
    }

    /**
     * Sets external_delete
     *
     * @param bool|null $external_delete external_delete
     *
     * @return self
     */
    public function setExternalDelete($external_delete)
    {
        $this->container['external_delete'] = $external_delete;

        return $this;
    }

    /**
     * Gets managed
     *
     * @return bool|null
     */
    public function getManaged()
    {
        return $this->container['managed'];
    }

    /**
     * Sets managed
     *
     * @param bool|null $managed managed
     *
     * @return self
     */
    public function setManaged($managed)
    {
        $this->container['managed'] = $managed;

        return $this;
    }

    /**
     * Gets control_power
     *
     * @return bool|null
     */
    public function getControlPower()
    {
        return $this->container['control_power'];
    }

    /**
     * Sets control_power
     *
     * @param bool|null $control_power control_power
     *
     * @return self
     */
    public function setControlPower($control_power)
    {
        $this->container['control_power'] = $control_power;

        return $this;
    }

    /**
     * Gets control_suspend
     *
     * @return bool|null
     */
    public function getControlSuspend()
    {
        return $this->container['control_suspend'];
    }

    /**
     * Sets control_suspend
     *
     * @param bool|null $control_suspend control_suspend
     *
     * @return self
     */
    public function setControlSuspend($control_suspend)
    {
        $this->container['control_suspend'] = $control_suspend;

        return $this;
    }

    /**
     * Gets creatable
     *
     * @return bool|null
     */
    public function getCreatable()
    {
        return $this->container['creatable'];
    }

    /**
     * Sets creatable
     *
     * @param bool|null $creatable creatable
     *
     * @return self
     */
    public function setCreatable($creatable)
    {
        $this->container['creatable'] = $creatable;

        return $this;
    }

    /**
     * Gets has_agent
     *
     * @return bool|null
     */
    public function getHasAgent()
    {
        return $this->container['has_agent'];
    }

    /**
     * Sets has_agent
     *
     * @param bool|null $has_agent has_agent
     *
     * @return self
     */
    public function setHasAgent($has_agent)
    {
        $this->container['has_agent'] = $has_agent;

        return $this;
    }

    /**
     * Gets vm_hypervisor
     *
     * @return bool|null
     */
    public function getVmHypervisor()
    {
        return $this->container['vm_hypervisor'];
    }

    /**
     * Sets vm_hypervisor
     *
     * @param bool|null $vm_hypervisor vm_hypervisor
     *
     * @return self
     */
    public function setVmHypervisor($vm_hypervisor)
    {
        $this->container['vm_hypervisor'] = $vm_hypervisor;

        return $this;
    }

    /**
     * Gets container_hypervisor
     *
     * @return bool|null
     */
    public function getContainerHypervisor()
    {
        return $this->container['container_hypervisor'];
    }

    /**
     * Sets container_hypervisor
     *
     * @param bool|null $container_hypervisor container_hypervisor
     *
     * @return self
     */
    public function setContainerHypervisor($container_hypervisor)
    {
        $this->container['container_hypervisor'] = $container_hypervisor;

        return $this;
    }

    /**
     * Gets bare_metal_host
     *
     * @return bool|null
     */
    public function getBareMetalHost()
    {
        return $this->container['bare_metal_host'];
    }

    /**
     * Sets bare_metal_host
     *
     * @param bool|null $bare_metal_host bare_metal_host
     *
     * @return self
     */
    public function setBareMetalHost($bare_metal_host)
    {
        $this->container['bare_metal_host'] = $bare_metal_host;

        return $this;
    }

    /**
     * Gets guest_vm
     *
     * @return bool|null
     */
    public function getGuestVm()
    {
        return $this->container['guest_vm'];
    }

    /**
     * Sets guest_vm
     *
     * @param bool|null $guest_vm guest_vm
     *
     * @return self
     */
    public function setGuestVm($guest_vm)
    {
        $this->container['guest_vm'] = $guest_vm;

        return $this;
    }

    /**
     * Gets has_automation
     *
     * @return bool|null
     */
    public function getHasAutomation()
    {
        return $this->container['has_automation'];
    }

    /**
     * Sets has_automation
     *
     * @param bool|null $has_automation has_automation
     *
     * @return self
     */
    public function setHasAutomation($has_automation)
    {
        $this->container['has_automation'] = $has_automation;

        return $this;
    }

    /**
     * Gets provision_type
     *
     * @return \OpenAPI\Client\Model\ZoneTypeProvisionType|null
     */
    public function getProvisionType()
    {
        return $this->container['provision_type'];
    }

    /**
     * Sets provision_type
     *
     * @param \OpenAPI\Client\Model\ZoneTypeProvisionType|null $provision_type provision_type
     *
     * @return self
     */
    public function setProvisionType($provision_type)
    {
        $this->container['provision_type'] = $provision_type;

        return $this;
    }

    /**
     * Gets option_types
     *
     * @return \OpenAPI\Client\Model\ZoneTypeOptionTypes[]|null
     */
    public function getOptionTypes()
    {
        return $this->container['option_types'];
    }

    /**
     * Sets option_types
     *
     * @param \OpenAPI\Client\Model\ZoneTypeOptionTypes[]|null $option_types option_types
     *
     * @return self
     */
    public function setOptionTypes($option_types)
    {
        $this->container['option_types'] = $option_types;

        return $this;
    }

    /**
     * Gets display_order
     *
     * @return int|null
     */
    public function getDisplayOrder()
    {
        return $this->container['display_order'];
    }

    /**
     * Sets display_order
     *
     * @param int|null $display_order display_order
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

