<?php
/**
 * ContainerType
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
 * ContainerType Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class ContainerType implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'containerType';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'id' => 'int',
        'account' => '\OpenAPI\Client\Model\ContainerTypeAccount',
        'name' => 'string',
        'labels' => 'string[]',
        'short_name' => 'string',
        'code' => 'string',
        'container_version' => 'string',
        'provision_type' => '\OpenAPI\Client\Model\ContainerTypeProvisionType',
        'virtual_image' => '\OpenAPI\Client\Model\ContainerTypeAccount',
        'category' => 'string',
        'config' => 'object',
        'container_ports' => '\OpenAPI\Client\Model\ContainerTypeContainerPorts[]',
        'container_scripts' => 'object[]',
        'container_templates' => 'object[]',
        'environment_variables' => 'object[]'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'id' => 'int32',
        'account' => null,
        'name' => null,
        'labels' => null,
        'short_name' => null,
        'code' => null,
        'container_version' => null,
        'provision_type' => null,
        'virtual_image' => null,
        'category' => null,
        'config' => null,
        'container_ports' => null,
        'container_scripts' => null,
        'container_templates' => null,
        'environment_variables' => null
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
        'account' => 'account',
        'name' => 'name',
        'labels' => 'labels',
        'short_name' => 'shortName',
        'code' => 'code',
        'container_version' => 'containerVersion',
        'provision_type' => 'provisionType',
        'virtual_image' => 'virtualImage',
        'category' => 'category',
        'config' => 'config',
        'container_ports' => 'containerPorts',
        'container_scripts' => 'containerScripts',
        'container_templates' => 'containerTemplates',
        'environment_variables' => 'environmentVariables'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'id' => 'setId',
        'account' => 'setAccount',
        'name' => 'setName',
        'labels' => 'setLabels',
        'short_name' => 'setShortName',
        'code' => 'setCode',
        'container_version' => 'setContainerVersion',
        'provision_type' => 'setProvisionType',
        'virtual_image' => 'setVirtualImage',
        'category' => 'setCategory',
        'config' => 'setConfig',
        'container_ports' => 'setContainerPorts',
        'container_scripts' => 'setContainerScripts',
        'container_templates' => 'setContainerTemplates',
        'environment_variables' => 'setEnvironmentVariables'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'id' => 'getId',
        'account' => 'getAccount',
        'name' => 'getName',
        'labels' => 'getLabels',
        'short_name' => 'getShortName',
        'code' => 'getCode',
        'container_version' => 'getContainerVersion',
        'provision_type' => 'getProvisionType',
        'virtual_image' => 'getVirtualImage',
        'category' => 'getCategory',
        'config' => 'getConfig',
        'container_ports' => 'getContainerPorts',
        'container_scripts' => 'getContainerScripts',
        'container_templates' => 'getContainerTemplates',
        'environment_variables' => 'getEnvironmentVariables'
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
        $this->container['account'] = $data['account'] ?? null;
        $this->container['name'] = $data['name'] ?? null;
        $this->container['labels'] = $data['labels'] ?? null;
        $this->container['short_name'] = $data['short_name'] ?? null;
        $this->container['code'] = $data['code'] ?? null;
        $this->container['container_version'] = $data['container_version'] ?? null;
        $this->container['provision_type'] = $data['provision_type'] ?? null;
        $this->container['virtual_image'] = $data['virtual_image'] ?? null;
        $this->container['category'] = $data['category'] ?? null;
        $this->container['config'] = $data['config'] ?? null;
        $this->container['container_ports'] = $data['container_ports'] ?? null;
        $this->container['container_scripts'] = $data['container_scripts'] ?? null;
        $this->container['container_templates'] = $data['container_templates'] ?? null;
        $this->container['environment_variables'] = $data['environment_variables'] ?? null;
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
     * Gets account
     *
     * @return \OpenAPI\Client\Model\ContainerTypeAccount|null
     */
    public function getAccount()
    {
        return $this->container['account'];
    }

    /**
     * Sets account
     *
     * @param \OpenAPI\Client\Model\ContainerTypeAccount|null $account account
     *
     * @return self
     */
    public function setAccount($account)
    {
        $this->container['account'] = $account;

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
     * Gets labels
     *
     * @return string[]|null
     */
    public function getLabels()
    {
        return $this->container['labels'];
    }

    /**
     * Sets labels
     *
     * @param string[]|null $labels labels
     *
     * @return self
     */
    public function setLabels($labels)
    {
        $this->container['labels'] = $labels;

        return $this;
    }

    /**
     * Gets short_name
     *
     * @return string|null
     */
    public function getShortName()
    {
        return $this->container['short_name'];
    }

    /**
     * Sets short_name
     *
     * @param string|null $short_name short_name
     *
     * @return self
     */
    public function setShortName($short_name)
    {
        $this->container['short_name'] = $short_name;

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
     * Gets container_version
     *
     * @return string|null
     */
    public function getContainerVersion()
    {
        return $this->container['container_version'];
    }

    /**
     * Sets container_version
     *
     * @param string|null $container_version container_version
     *
     * @return self
     */
    public function setContainerVersion($container_version)
    {
        $this->container['container_version'] = $container_version;

        return $this;
    }

    /**
     * Gets provision_type
     *
     * @return \OpenAPI\Client\Model\ContainerTypeProvisionType|null
     */
    public function getProvisionType()
    {
        return $this->container['provision_type'];
    }

    /**
     * Sets provision_type
     *
     * @param \OpenAPI\Client\Model\ContainerTypeProvisionType|null $provision_type provision_type
     *
     * @return self
     */
    public function setProvisionType($provision_type)
    {
        $this->container['provision_type'] = $provision_type;

        return $this;
    }

    /**
     * Gets virtual_image
     *
     * @return \OpenAPI\Client\Model\ContainerTypeAccount|null
     */
    public function getVirtualImage()
    {
        return $this->container['virtual_image'];
    }

    /**
     * Sets virtual_image
     *
     * @param \OpenAPI\Client\Model\ContainerTypeAccount|null $virtual_image virtual_image
     *
     * @return self
     */
    public function setVirtualImage($virtual_image)
    {
        $this->container['virtual_image'] = $virtual_image;

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
     * Gets config
     *
     * @return object|null
     */
    public function getConfig()
    {
        return $this->container['config'];
    }

    /**
     * Sets config
     *
     * @param object|null $config config
     *
     * @return self
     */
    public function setConfig($config)
    {
        $this->container['config'] = $config;

        return $this;
    }

    /**
     * Gets container_ports
     *
     * @return \OpenAPI\Client\Model\ContainerTypeContainerPorts[]|null
     */
    public function getContainerPorts()
    {
        return $this->container['container_ports'];
    }

    /**
     * Sets container_ports
     *
     * @param \OpenAPI\Client\Model\ContainerTypeContainerPorts[]|null $container_ports container_ports
     *
     * @return self
     */
    public function setContainerPorts($container_ports)
    {
        $this->container['container_ports'] = $container_ports;

        return $this;
    }

    /**
     * Gets container_scripts
     *
     * @return object[]|null
     */
    public function getContainerScripts()
    {
        return $this->container['container_scripts'];
    }

    /**
     * Sets container_scripts
     *
     * @param object[]|null $container_scripts container_scripts
     *
     * @return self
     */
    public function setContainerScripts($container_scripts)
    {
        $this->container['container_scripts'] = $container_scripts;

        return $this;
    }

    /**
     * Gets container_templates
     *
     * @return object[]|null
     */
    public function getContainerTemplates()
    {
        return $this->container['container_templates'];
    }

    /**
     * Sets container_templates
     *
     * @param object[]|null $container_templates container_templates
     *
     * @return self
     */
    public function setContainerTemplates($container_templates)
    {
        $this->container['container_templates'] = $container_templates;

        return $this;
    }

    /**
     * Gets environment_variables
     *
     * @return object[]|null
     */
    public function getEnvironmentVariables()
    {
        return $this->container['environment_variables'];
    }

    /**
     * Sets environment_variables
     *
     * @param object[]|null $environment_variables environment_variables
     *
     * @return self
     */
    public function setEnvironmentVariables($environment_variables)
    {
        $this->container['environment_variables'] = $environment_variables;

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


