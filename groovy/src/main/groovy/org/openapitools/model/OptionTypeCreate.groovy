package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.OptionTypeCreateOptionList;

@Canonical
class OptionTypeCreate {
    /* The name of the option type for handy reference */
    String name
    /* Short description of the option type */
    String description
    /* Array of label strings, can be used for filtering. */
    List<String> labels = new ArrayList<String>()
    /* Field Name, the name for user input. This along with fieldContext determines the configuration property name.  The property key for when posting this option type to a JSON POST request */
    String fieldName
    /* Type, the type of input. eg. text, checkbox, select, etc. */
    String type = "text"
    /* Field Label, the label for user input. */
    String fieldLabel
    /* Any placeholder text when nothing is yet entered */
    String placeholder
    /* Verify Pattern, A regexp string that validates the input, use (?i) to make the matcher case insensitive */
    String verifyPattern
    /* The default value if no user entry is specified. This value should be passed to the desired JSON Map if nothing else is entered */
    String defaultValue
    /* Is this field entry required for the request */
    Boolean required = false
    /* Export as Tag */
    Boolean exportMeta = false
    /* Used primarily on tasks and workflows. Basically wether or not the field can be overridden optionally when the object is run */
    Boolean editable = false
    
    OptionTypeCreateOptionList optionList
}
