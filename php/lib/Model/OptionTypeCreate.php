<?php
/**
 * OptionTypeCreate
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
 * OptionTypeCreate Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class OptionTypeCreate implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'optionTypeCreate';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'name' => 'string',
        'description' => 'string',
        'labels' => 'string[]',
        'field_name' => 'string',
        'type' => 'string',
        'field_label' => 'string',
        'placeholder' => 'string',
        'verify_pattern' => 'string',
        'default_value' => 'string',
        'required' => 'bool',
        'export_meta' => 'bool',
        'editable' => 'bool',
        'option_list' => '\OpenAPI\Client\Model\OptionTypeCreateOptionList'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'name' => null,
        'description' => null,
        'labels' => null,
        'field_name' => null,
        'type' => null,
        'field_label' => null,
        'placeholder' => null,
        'verify_pattern' => null,
        'default_value' => null,
        'required' => null,
        'export_meta' => null,
        'editable' => null,
        'option_list' => null
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
        'name' => 'name',
        'description' => 'description',
        'labels' => 'labels',
        'field_name' => 'fieldName',
        'type' => 'type',
        'field_label' => 'fieldLabel',
        'placeholder' => 'placeholder',
        'verify_pattern' => 'verifyPattern',
        'default_value' => 'defaultValue',
        'required' => 'required',
        'export_meta' => 'exportMeta',
        'editable' => 'editable',
        'option_list' => 'optionList'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'name' => 'setName',
        'description' => 'setDescription',
        'labels' => 'setLabels',
        'field_name' => 'setFieldName',
        'type' => 'setType',
        'field_label' => 'setFieldLabel',
        'placeholder' => 'setPlaceholder',
        'verify_pattern' => 'setVerifyPattern',
        'default_value' => 'setDefaultValue',
        'required' => 'setRequired',
        'export_meta' => 'setExportMeta',
        'editable' => 'setEditable',
        'option_list' => 'setOptionList'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'name' => 'getName',
        'description' => 'getDescription',
        'labels' => 'getLabels',
        'field_name' => 'getFieldName',
        'type' => 'getType',
        'field_label' => 'getFieldLabel',
        'placeholder' => 'getPlaceholder',
        'verify_pattern' => 'getVerifyPattern',
        'default_value' => 'getDefaultValue',
        'required' => 'getRequired',
        'export_meta' => 'getExportMeta',
        'editable' => 'getEditable',
        'option_list' => 'getOptionList'
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
        $this->container['name'] = $data['name'] ?? null;
        $this->container['description'] = $data['description'] ?? null;
        $this->container['labels'] = $data['labels'] ?? null;
        $this->container['field_name'] = $data['field_name'] ?? null;
        $this->container['type'] = $data['type'] ?? 'text';
        $this->container['field_label'] = $data['field_label'] ?? null;
        $this->container['placeholder'] = $data['placeholder'] ?? null;
        $this->container['verify_pattern'] = $data['verify_pattern'] ?? null;
        $this->container['default_value'] = $data['default_value'] ?? null;
        $this->container['required'] = $data['required'] ?? false;
        $this->container['export_meta'] = $data['export_meta'] ?? false;
        $this->container['editable'] = $data['editable'] ?? false;
        $this->container['option_list'] = $data['option_list'] ?? null;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        if ($this->container['name'] === null) {
            $invalidProperties[] = "'name' can't be null";
        }
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
     * Gets name
     *
     * @return string
     */
    public function getName()
    {
        return $this->container['name'];
    }

    /**
     * Sets name
     *
     * @param string $name The name of the option type for handy reference
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
     * @param string|null $description Short description of the option type
     *
     * @return self
     */
    public function setDescription($description)
    {
        $this->container['description'] = $description;

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
     * @param string[]|null $labels Array of label strings, can be used for filtering.
     *
     * @return self
     */
    public function setLabels($labels)
    {
        $this->container['labels'] = $labels;

        return $this;
    }

    /**
     * Gets field_name
     *
     * @return string|null
     */
    public function getFieldName()
    {
        return $this->container['field_name'];
    }

    /**
     * Sets field_name
     *
     * @param string|null $field_name Field Name, the name for user input. This along with fieldContext determines the configuration property name.  The property key for when posting this option type to a JSON POST request
     *
     * @return self
     */
    public function setFieldName($field_name)
    {
        $this->container['field_name'] = $field_name;

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
     * @param string|null $type Type, the type of input. eg. text, checkbox, select, etc.
     *
     * @return self
     */
    public function setType($type)
    {
        $this->container['type'] = $type;

        return $this;
    }

    /**
     * Gets field_label
     *
     * @return string|null
     */
    public function getFieldLabel()
    {
        return $this->container['field_label'];
    }

    /**
     * Sets field_label
     *
     * @param string|null $field_label Field Label, the label for user input.
     *
     * @return self
     */
    public function setFieldLabel($field_label)
    {
        $this->container['field_label'] = $field_label;

        return $this;
    }

    /**
     * Gets placeholder
     *
     * @return string|null
     */
    public function getPlaceholder()
    {
        return $this->container['placeholder'];
    }

    /**
     * Sets placeholder
     *
     * @param string|null $placeholder Any placeholder text when nothing is yet entered
     *
     * @return self
     */
    public function setPlaceholder($placeholder)
    {
        $this->container['placeholder'] = $placeholder;

        return $this;
    }

    /**
     * Gets verify_pattern
     *
     * @return string|null
     */
    public function getVerifyPattern()
    {
        return $this->container['verify_pattern'];
    }

    /**
     * Sets verify_pattern
     *
     * @param string|null $verify_pattern Verify Pattern, A regexp string that validates the input, use (?i) to make the matcher case insensitive
     *
     * @return self
     */
    public function setVerifyPattern($verify_pattern)
    {
        $this->container['verify_pattern'] = $verify_pattern;

        return $this;
    }

    /**
     * Gets default_value
     *
     * @return string|null
     */
    public function getDefaultValue()
    {
        return $this->container['default_value'];
    }

    /**
     * Sets default_value
     *
     * @param string|null $default_value The default value if no user entry is specified. This value should be passed to the desired JSON Map if nothing else is entered
     *
     * @return self
     */
    public function setDefaultValue($default_value)
    {
        $this->container['default_value'] = $default_value;

        return $this;
    }

    /**
     * Gets required
     *
     * @return bool|null
     */
    public function getRequired()
    {
        return $this->container['required'];
    }

    /**
     * Sets required
     *
     * @param bool|null $required Is this field entry required for the request
     *
     * @return self
     */
    public function setRequired($required)
    {
        $this->container['required'] = $required;

        return $this;
    }

    /**
     * Gets export_meta
     *
     * @return bool|null
     */
    public function getExportMeta()
    {
        return $this->container['export_meta'];
    }

    /**
     * Sets export_meta
     *
     * @param bool|null $export_meta Export as Tag
     *
     * @return self
     */
    public function setExportMeta($export_meta)
    {
        $this->container['export_meta'] = $export_meta;

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
     * @param bool|null $editable Used primarily on tasks and workflows. Basically wether or not the field can be overridden optionally when the object is run
     *
     * @return self
     */
    public function setEditable($editable)
    {
        $this->container['editable'] = $editable;

        return $this;
    }

    /**
     * Gets option_list
     *
     * @return \OpenAPI\Client\Model\OptionTypeCreateOptionList|null
     */
    public function getOptionList()
    {
        return $this->container['option_list'];
    }

    /**
     * Sets option_list
     *
     * @param \OpenAPI\Client\Model\OptionTypeCreateOptionList|null $option_list option_list
     *
     * @return self
     */
    public function setOptionList($option_list)
    {
        $this->container['option_list'] = $option_list;

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


