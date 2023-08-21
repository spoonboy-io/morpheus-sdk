<?php
/**
 * AppPrepareApplyData
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
 * AppPrepareApplyData Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class AppPrepareApplyData implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'appPrepareApply_data';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'image' => 'string',
        'name' => 'string',
        'auto_validate' => 'bool',
        'terraform' => '\OpenAPI\Client\Model\AppPrepareApplyDataTerraform',
        'type' => 'string',
        'config' => 'object',
        'blueprint_name' => 'string',
        'description' => 'string',
        'template_id' => 'int',
        'blueprint_id' => 'int',
        'group' => '\OpenAPI\Client\Model\AppPrepareApplyDataGroup'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'image' => null,
        'name' => null,
        'auto_validate' => null,
        'terraform' => null,
        'type' => null,
        'config' => null,
        'blueprint_name' => null,
        'description' => null,
        'template_id' => 'int64',
        'blueprint_id' => 'int64',
        'group' => null
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
        'image' => 'image',
        'name' => 'name',
        'auto_validate' => 'autoValidate',
        'terraform' => 'terraform',
        'type' => 'type',
        'config' => 'config',
        'blueprint_name' => 'blueprintName',
        'description' => 'description',
        'template_id' => 'templateId',
        'blueprint_id' => 'blueprintId',
        'group' => 'group'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'image' => 'setImage',
        'name' => 'setName',
        'auto_validate' => 'setAutoValidate',
        'terraform' => 'setTerraform',
        'type' => 'setType',
        'config' => 'setConfig',
        'blueprint_name' => 'setBlueprintName',
        'description' => 'setDescription',
        'template_id' => 'setTemplateId',
        'blueprint_id' => 'setBlueprintId',
        'group' => 'setGroup'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'image' => 'getImage',
        'name' => 'getName',
        'auto_validate' => 'getAutoValidate',
        'terraform' => 'getTerraform',
        'type' => 'getType',
        'config' => 'getConfig',
        'blueprint_name' => 'getBlueprintName',
        'description' => 'getDescription',
        'template_id' => 'getTemplateId',
        'blueprint_id' => 'getBlueprintId',
        'group' => 'getGroup'
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
        $this->container['image'] = $data['image'] ?? null;
        $this->container['name'] = $data['name'] ?? null;
        $this->container['auto_validate'] = $data['auto_validate'] ?? null;
        $this->container['terraform'] = $data['terraform'] ?? null;
        $this->container['type'] = $data['type'] ?? null;
        $this->container['config'] = $data['config'] ?? null;
        $this->container['blueprint_name'] = $data['blueprint_name'] ?? null;
        $this->container['description'] = $data['description'] ?? null;
        $this->container['template_id'] = $data['template_id'] ?? null;
        $this->container['blueprint_id'] = $data['blueprint_id'] ?? null;
        $this->container['group'] = $data['group'] ?? null;
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
     * Gets image
     *
     * @return string|null
     */
    public function getImage()
    {
        return $this->container['image'];
    }

    /**
     * Sets image
     *
     * @param string|null $image image
     *
     * @return self
     */
    public function setImage($image)
    {
        $this->container['image'] = $image;

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
     * Gets auto_validate
     *
     * @return bool|null
     */
    public function getAutoValidate()
    {
        return $this->container['auto_validate'];
    }

    /**
     * Sets auto_validate
     *
     * @param bool|null $auto_validate auto_validate
     *
     * @return self
     */
    public function setAutoValidate($auto_validate)
    {
        $this->container['auto_validate'] = $auto_validate;

        return $this;
    }

    /**
     * Gets terraform
     *
     * @return \OpenAPI\Client\Model\AppPrepareApplyDataTerraform|null
     */
    public function getTerraform()
    {
        return $this->container['terraform'];
    }

    /**
     * Sets terraform
     *
     * @param \OpenAPI\Client\Model\AppPrepareApplyDataTerraform|null $terraform terraform
     *
     * @return self
     */
    public function setTerraform($terraform)
    {
        $this->container['terraform'] = $terraform;

        return $this;
    }

    /**
     * Gets type
     *
     * @return string|null
     */
    public function getType()
    {
        return $this->container['type'];
    }

    /**
     * Sets type
     *
     * @param string|null $type type
     *
     * @return self
     */
    public function setType($type)
    {
        $this->container['type'] = $type;

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
     * Gets blueprint_name
     *
     * @return string|null
     */
    public function getBlueprintName()
    {
        return $this->container['blueprint_name'];
    }

    /**
     * Sets blueprint_name
     *
     * @param string|null $blueprint_name blueprint_name
     *
     * @return self
     */
    public function setBlueprintName($blueprint_name)
    {
        $this->container['blueprint_name'] = $blueprint_name;

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
     * Gets template_id
     *
     * @return int|null
     */
    public function getTemplateId()
    {
        return $this->container['template_id'];
    }

    /**
     * Sets template_id
     *
     * @param int|null $template_id template_id
     *
     * @return self
     */
    public function setTemplateId($template_id)
    {
        $this->container['template_id'] = $template_id;

        return $this;
    }

    /**
     * Gets blueprint_id
     *
     * @return int|null
     */
    public function getBlueprintId()
    {
        return $this->container['blueprint_id'];
    }

    /**
     * Sets blueprint_id
     *
     * @param int|null $blueprint_id blueprint_id
     *
     * @return self
     */
    public function setBlueprintId($blueprint_id)
    {
        $this->container['blueprint_id'] = $blueprint_id;

        return $this;
    }

    /**
     * Gets group
     *
     * @return \OpenAPI\Client\Model\AppPrepareApplyDataGroup|null
     */
    public function getGroup()
    {
        return $this->container['group'];
    }

    /**
     * Sets group
     *
     * @param \OpenAPI\Client\Model\AppPrepareApplyDataGroup|null $group group
     *
     * @return self
     */
    public function setGroup($group)
    {
        $this->container['group'] = $group;

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


