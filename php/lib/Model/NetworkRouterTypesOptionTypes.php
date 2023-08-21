<?php
/**
 * NetworkRouterTypesOptionTypes
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
 * NetworkRouterTypesOptionTypes Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class NetworkRouterTypesOptionTypes implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'networkRouterTypes_optionTypes';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'id' => 'int',
        'name' => 'string',
        'description' => 'string',
        'code' => 'string',
        'field_name' => 'string',
        'field_label' => 'string',
        'field_context' => 'string',
        'field_group' => 'string',
        'field_class' => 'string',
        'field_add_on' => 'string',
        'field_component' => 'string',
        'field_input' => 'string',
        'place_holder' => 'string',
        'help_block' => 'string',
        'default_value' => 'string',
        'option_source' => 'string',
        'option_list' => 'string',
        'type' => 'string',
        'advanced' => 'bool',
        'required' => 'bool',
        'editable' => 'bool',
        'creatable' => 'bool',
        'config' => 'object',
        'display_order' => 'int',
        'wrapper_class' => 'string',
        'enabled' => 'bool',
        'no_blank' => 'bool',
        'depends_on_code' => 'string',
        'contextual_default' => 'bool'
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
        'description' => null,
        'code' => null,
        'field_name' => null,
        'field_label' => null,
        'field_context' => null,
        'field_group' => null,
        'field_class' => null,
        'field_add_on' => null,
        'field_component' => null,
        'field_input' => null,
        'place_holder' => null,
        'help_block' => null,
        'default_value' => null,
        'option_source' => null,
        'option_list' => null,
        'type' => null,
        'advanced' => null,
        'required' => null,
        'editable' => null,
        'creatable' => null,
        'config' => null,
        'display_order' => 'int64',
        'wrapper_class' => null,
        'enabled' => null,
        'no_blank' => null,
        'depends_on_code' => null,
        'contextual_default' => null
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
        'description' => 'description',
        'code' => 'code',
        'field_name' => 'fieldName',
        'field_label' => 'fieldLabel',
        'field_context' => 'fieldContext',
        'field_group' => 'fieldGroup',
        'field_class' => 'fieldClass',
        'field_add_on' => 'fieldAddOn',
        'field_component' => 'fieldComponent',
        'field_input' => 'fieldInput',
        'place_holder' => 'placeHolder',
        'help_block' => 'helpBlock',
        'default_value' => 'defaultValue',
        'option_source' => 'optionSource',
        'option_list' => 'optionList',
        'type' => 'type',
        'advanced' => 'advanced',
        'required' => 'required',
        'editable' => 'editable',
        'creatable' => 'creatable',
        'config' => 'config',
        'display_order' => 'displayOrder',
        'wrapper_class' => 'wrapperClass',
        'enabled' => 'enabled',
        'no_blank' => 'noBlank',
        'depends_on_code' => 'dependsOnCode',
        'contextual_default' => 'contextualDefault'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'id' => 'setId',
        'name' => 'setName',
        'description' => 'setDescription',
        'code' => 'setCode',
        'field_name' => 'setFieldName',
        'field_label' => 'setFieldLabel',
        'field_context' => 'setFieldContext',
        'field_group' => 'setFieldGroup',
        'field_class' => 'setFieldClass',
        'field_add_on' => 'setFieldAddOn',
        'field_component' => 'setFieldComponent',
        'field_input' => 'setFieldInput',
        'place_holder' => 'setPlaceHolder',
        'help_block' => 'setHelpBlock',
        'default_value' => 'setDefaultValue',
        'option_source' => 'setOptionSource',
        'option_list' => 'setOptionList',
        'type' => 'setType',
        'advanced' => 'setAdvanced',
        'required' => 'setRequired',
        'editable' => 'setEditable',
        'creatable' => 'setCreatable',
        'config' => 'setConfig',
        'display_order' => 'setDisplayOrder',
        'wrapper_class' => 'setWrapperClass',
        'enabled' => 'setEnabled',
        'no_blank' => 'setNoBlank',
        'depends_on_code' => 'setDependsOnCode',
        'contextual_default' => 'setContextualDefault'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'id' => 'getId',
        'name' => 'getName',
        'description' => 'getDescription',
        'code' => 'getCode',
        'field_name' => 'getFieldName',
        'field_label' => 'getFieldLabel',
        'field_context' => 'getFieldContext',
        'field_group' => 'getFieldGroup',
        'field_class' => 'getFieldClass',
        'field_add_on' => 'getFieldAddOn',
        'field_component' => 'getFieldComponent',
        'field_input' => 'getFieldInput',
        'place_holder' => 'getPlaceHolder',
        'help_block' => 'getHelpBlock',
        'default_value' => 'getDefaultValue',
        'option_source' => 'getOptionSource',
        'option_list' => 'getOptionList',
        'type' => 'getType',
        'advanced' => 'getAdvanced',
        'required' => 'getRequired',
        'editable' => 'getEditable',
        'creatable' => 'getCreatable',
        'config' => 'getConfig',
        'display_order' => 'getDisplayOrder',
        'wrapper_class' => 'getWrapperClass',
        'enabled' => 'getEnabled',
        'no_blank' => 'getNoBlank',
        'depends_on_code' => 'getDependsOnCode',
        'contextual_default' => 'getContextualDefault'
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
        $this->container['description'] = $data['description'] ?? null;
        $this->container['code'] = $data['code'] ?? null;
        $this->container['field_name'] = $data['field_name'] ?? null;
        $this->container['field_label'] = $data['field_label'] ?? null;
        $this->container['field_context'] = $data['field_context'] ?? null;
        $this->container['field_group'] = $data['field_group'] ?? null;
        $this->container['field_class'] = $data['field_class'] ?? null;
        $this->container['field_add_on'] = $data['field_add_on'] ?? null;
        $this->container['field_component'] = $data['field_component'] ?? null;
        $this->container['field_input'] = $data['field_input'] ?? null;
        $this->container['place_holder'] = $data['place_holder'] ?? null;
        $this->container['help_block'] = $data['help_block'] ?? null;
        $this->container['default_value'] = $data['default_value'] ?? null;
        $this->container['option_source'] = $data['option_source'] ?? null;
        $this->container['option_list'] = $data['option_list'] ?? null;
        $this->container['type'] = $data['type'] ?? null;
        $this->container['advanced'] = $data['advanced'] ?? null;
        $this->container['required'] = $data['required'] ?? null;
        $this->container['editable'] = $data['editable'] ?? null;
        $this->container['creatable'] = $data['creatable'] ?? null;
        $this->container['config'] = $data['config'] ?? null;
        $this->container['display_order'] = $data['display_order'] ?? null;
        $this->container['wrapper_class'] = $data['wrapper_class'] ?? null;
        $this->container['enabled'] = $data['enabled'] ?? null;
        $this->container['no_blank'] = $data['no_blank'] ?? null;
        $this->container['depends_on_code'] = $data['depends_on_code'] ?? null;
        $this->container['contextual_default'] = $data['contextual_default'] ?? null;
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
     * @param string|null $field_name field_name
     *
     * @return self
     */
    public function setFieldName($field_name)
    {
        $this->container['field_name'] = $field_name;

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
     * @param string|null $field_label field_label
     *
     * @return self
     */
    public function setFieldLabel($field_label)
    {
        $this->container['field_label'] = $field_label;

        return $this;
    }

    /**
     * Gets field_context
     *
     * @return string|null
     */
    public function getFieldContext()
    {
        return $this->container['field_context'];
    }

    /**
     * Sets field_context
     *
     * @param string|null $field_context field_context
     *
     * @return self
     */
    public function setFieldContext($field_context)
    {
        $this->container['field_context'] = $field_context;

        return $this;
    }

    /**
     * Gets field_group
     *
     * @return string|null
     */
    public function getFieldGroup()
    {
        return $this->container['field_group'];
    }

    /**
     * Sets field_group
     *
     * @param string|null $field_group field_group
     *
     * @return self
     */
    public function setFieldGroup($field_group)
    {
        $this->container['field_group'] = $field_group;

        return $this;
    }

    /**
     * Gets field_class
     *
     * @return string|null
     */
    public function getFieldClass()
    {
        return $this->container['field_class'];
    }

    /**
     * Sets field_class
     *
     * @param string|null $field_class field_class
     *
     * @return self
     */
    public function setFieldClass($field_class)
    {
        $this->container['field_class'] = $field_class;

        return $this;
    }

    /**
     * Gets field_add_on
     *
     * @return string|null
     */
    public function getFieldAddOn()
    {
        return $this->container['field_add_on'];
    }

    /**
     * Sets field_add_on
     *
     * @param string|null $field_add_on field_add_on
     *
     * @return self
     */
    public function setFieldAddOn($field_add_on)
    {
        $this->container['field_add_on'] = $field_add_on;

        return $this;
    }

    /**
     * Gets field_component
     *
     * @return string|null
     */
    public function getFieldComponent()
    {
        return $this->container['field_component'];
    }

    /**
     * Sets field_component
     *
     * @param string|null $field_component field_component
     *
     * @return self
     */
    public function setFieldComponent($field_component)
    {
        $this->container['field_component'] = $field_component;

        return $this;
    }

    /**
     * Gets field_input
     *
     * @return string|null
     */
    public function getFieldInput()
    {
        return $this->container['field_input'];
    }

    /**
     * Sets field_input
     *
     * @param string|null $field_input field_input
     *
     * @return self
     */
    public function setFieldInput($field_input)
    {
        $this->container['field_input'] = $field_input;

        return $this;
    }

    /**
     * Gets place_holder
     *
     * @return string|null
     */
    public function getPlaceHolder()
    {
        return $this->container['place_holder'];
    }

    /**
     * Sets place_holder
     *
     * @param string|null $place_holder place_holder
     *
     * @return self
     */
    public function setPlaceHolder($place_holder)
    {
        $this->container['place_holder'] = $place_holder;

        return $this;
    }

    /**
     * Gets help_block
     *
     * @return string|null
     */
    public function getHelpBlock()
    {
        return $this->container['help_block'];
    }

    /**
     * Sets help_block
     *
     * @param string|null $help_block help_block
     *
     * @return self
     */
    public function setHelpBlock($help_block)
    {
        $this->container['help_block'] = $help_block;

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
     * @param string|null $default_value default_value
     *
     * @return self
     */
    public function setDefaultValue($default_value)
    {
        $this->container['default_value'] = $default_value;

        return $this;
    }

    /**
     * Gets option_source
     *
     * @return string|null
     */
    public function getOptionSource()
    {
        return $this->container['option_source'];
    }

    /**
     * Sets option_source
     *
     * @param string|null $option_source option_source
     *
     * @return self
     */
    public function setOptionSource($option_source)
    {
        $this->container['option_source'] = $option_source;

        return $this;
    }

    /**
     * Gets option_list
     *
     * @return string|null
     */
    public function getOptionList()
    {
        return $this->container['option_list'];
    }

    /**
     * Sets option_list
     *
     * @param string|null $option_list option_list
     *
     * @return self
     */
    public function setOptionList($option_list)
    {
        $this->container['option_list'] = $option_list;

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
     * Gets advanced
     *
     * @return bool|null
     */
    public function getAdvanced()
    {
        return $this->container['advanced'];
    }

    /**
     * Sets advanced
     *
     * @param bool|null $advanced advanced
     *
     * @return self
     */
    public function setAdvanced($advanced)
    {
        $this->container['advanced'] = $advanced;

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
     * @param bool|null $required required
     *
     * @return self
     */
    public function setRequired($required)
    {
        $this->container['required'] = $required;

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
     * Gets wrapper_class
     *
     * @return string|null
     */
    public function getWrapperClass()
    {
        return $this->container['wrapper_class'];
    }

    /**
     * Sets wrapper_class
     *
     * @param string|null $wrapper_class wrapper_class
     *
     * @return self
     */
    public function setWrapperClass($wrapper_class)
    {
        $this->container['wrapper_class'] = $wrapper_class;

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
     * Gets no_blank
     *
     * @return bool|null
     */
    public function getNoBlank()
    {
        return $this->container['no_blank'];
    }

    /**
     * Sets no_blank
     *
     * @param bool|null $no_blank no_blank
     *
     * @return self
     */
    public function setNoBlank($no_blank)
    {
        $this->container['no_blank'] = $no_blank;

        return $this;
    }

    /**
     * Gets depends_on_code
     *
     * @return string|null
     */
    public function getDependsOnCode()
    {
        return $this->container['depends_on_code'];
    }

    /**
     * Sets depends_on_code
     *
     * @param string|null $depends_on_code depends_on_code
     *
     * @return self
     */
    public function setDependsOnCode($depends_on_code)
    {
        $this->container['depends_on_code'] = $depends_on_code;

        return $this;
    }

    /**
     * Gets contextual_default
     *
     * @return bool|null
     */
    public function getContextualDefault()
    {
        return $this->container['contextual_default'];
    }

    /**
     * Sets contextual_default
     *
     * @param bool|null $contextual_default contextual_default
     *
     * @return self
     */
    public function setContextualDefault($contextual_default)
    {
        $this->container['contextual_default'] = $contextual_default;

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

