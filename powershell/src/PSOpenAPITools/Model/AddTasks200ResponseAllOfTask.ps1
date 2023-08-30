#
# Morpheus API
# Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
# Version: 6.2.1
# Contact: dev@morpheusdata.com
# Generated by OpenAPI Generator: https://openapi-generator.tech
#

<#
.SYNOPSIS

No summary available.

.DESCRIPTION

No description available.

.PARAMETER Json

JSON object

.OUTPUTS

AddTasks200ResponseAllOfTask<PSCustomObject>
#>
function ConvertFrom-JsonToAddTasks200ResponseAllOfTask {
    [CmdletBinding()]
    Param (
        [AllowEmptyString()]
        [string]$Json
    )

    Process {
        $match = 0
        $matchType = $null
        $matchInstance = $null

        # try to match TaskAnsiblePlaybookConfig defined in the oneOf schemas
        try {
            $matchInstance = ConvertFrom-JsonToTaskAnsiblePlaybookConfig $Json

            foreach($property in $matchInstance.PsObject.Properties) {
                if ($null -ne $property.Value) {
                    $matchType = "TaskAnsiblePlaybookConfig"
                    $match++
                    break
                }
            }
        } catch {
            # fail to match the schema defined in oneOf, proceed to the next one
            Write-Debug "Failed to match 'TaskAnsiblePlaybookConfig' defined in oneOf (AddTasks200ResponseAllOfTask). Proceeding to the next one if any."
        }

        # try to match TaskAnsibleTowerConfig defined in the oneOf schemas
        try {
            $matchInstance = ConvertFrom-JsonToTaskAnsibleTowerConfig $Json

            foreach($property in $matchInstance.PsObject.Properties) {
                if ($null -ne $property.Value) {
                    $matchType = "TaskAnsibleTowerConfig"
                    $match++
                    break
                }
            }
        } catch {
            # fail to match the schema defined in oneOf, proceed to the next one
            Write-Debug "Failed to match 'TaskAnsibleTowerConfig' defined in oneOf (AddTasks200ResponseAllOfTask). Proceeding to the next one if any."
        }

        # try to match TaskChefBootstrapConfig defined in the oneOf schemas
        try {
            $matchInstance = ConvertFrom-JsonToTaskChefBootstrapConfig $Json

            foreach($property in $matchInstance.PsObject.Properties) {
                if ($null -ne $property.Value) {
                    $matchType = "TaskChefBootstrapConfig"
                    $match++
                    break
                }
            }
        } catch {
            # fail to match the schema defined in oneOf, proceed to the next one
            Write-Debug "Failed to match 'TaskChefBootstrapConfig' defined in oneOf (AddTasks200ResponseAllOfTask). Proceeding to the next one if any."
        }

        # try to match TaskEmailConfig defined in the oneOf schemas
        try {
            $matchInstance = ConvertFrom-JsonToTaskEmailConfig $Json

            foreach($property in $matchInstance.PsObject.Properties) {
                if ($null -ne $property.Value) {
                    $matchType = "TaskEmailConfig"
                    $match++
                    break
                }
            }
        } catch {
            # fail to match the schema defined in oneOf, proceed to the next one
            Write-Debug "Failed to match 'TaskEmailConfig' defined in oneOf (AddTasks200ResponseAllOfTask). Proceeding to the next one if any."
        }

        # try to match TaskGroovyConfig defined in the oneOf schemas
        try {
            $matchInstance = ConvertFrom-JsonToTaskGroovyConfig $Json

            foreach($property in $matchInstance.PsObject.Properties) {
                if ($null -ne $property.Value) {
                    $matchType = "TaskGroovyConfig"
                    $match++
                    break
                }
            }
        } catch {
            # fail to match the schema defined in oneOf, proceed to the next one
            Write-Debug "Failed to match 'TaskGroovyConfig' defined in oneOf (AddTasks200ResponseAllOfTask). Proceeding to the next one if any."
        }

        # try to match TaskHttpConfig defined in the oneOf schemas
        try {
            $matchInstance = ConvertFrom-JsonToTaskHttpConfig $Json

            foreach($property in $matchInstance.PsObject.Properties) {
                if ($null -ne $property.Value) {
                    $matchType = "TaskHttpConfig"
                    $match++
                    break
                }
            }
        } catch {
            # fail to match the schema defined in oneOf, proceed to the next one
            Write-Debug "Failed to match 'TaskHttpConfig' defined in oneOf (AddTasks200ResponseAllOfTask). Proceeding to the next one if any."
        }

        # try to match TaskJavaConfig defined in the oneOf schemas
        try {
            $matchInstance = ConvertFrom-JsonToTaskJavaConfig $Json

            foreach($property in $matchInstance.PsObject.Properties) {
                if ($null -ne $property.Value) {
                    $matchType = "TaskJavaConfig"
                    $match++
                    break
                }
            }
        } catch {
            # fail to match the schema defined in oneOf, proceed to the next one
            Write-Debug "Failed to match 'TaskJavaConfig' defined in oneOf (AddTasks200ResponseAllOfTask). Proceeding to the next one if any."
        }

        # try to match TaskJrubyConfig defined in the oneOf schemas
        try {
            $matchInstance = ConvertFrom-JsonToTaskJrubyConfig $Json

            foreach($property in $matchInstance.PsObject.Properties) {
                if ($null -ne $property.Value) {
                    $matchType = "TaskJrubyConfig"
                    $match++
                    break
                }
            }
        } catch {
            # fail to match the schema defined in oneOf, proceed to the next one
            Write-Debug "Failed to match 'TaskJrubyConfig' defined in oneOf (AddTasks200ResponseAllOfTask). Proceeding to the next one if any."
        }

        # try to match TaskLibraryScriptConfig defined in the oneOf schemas
        try {
            $matchInstance = ConvertFrom-JsonToTaskLibraryScriptConfig $Json

            foreach($property in $matchInstance.PsObject.Properties) {
                if ($null -ne $property.Value) {
                    $matchType = "TaskLibraryScriptConfig"
                    $match++
                    break
                }
            }
        } catch {
            # fail to match the schema defined in oneOf, proceed to the next one
            Write-Debug "Failed to match 'TaskLibraryScriptConfig' defined in oneOf (AddTasks200ResponseAllOfTask). Proceeding to the next one if any."
        }

        # try to match TaskLibraryTemplateConfig defined in the oneOf schemas
        try {
            $matchInstance = ConvertFrom-JsonToTaskLibraryTemplateConfig $Json

            foreach($property in $matchInstance.PsObject.Properties) {
                if ($null -ne $property.Value) {
                    $matchType = "TaskLibraryTemplateConfig"
                    $match++
                    break
                }
            }
        } catch {
            # fail to match the schema defined in oneOf, proceed to the next one
            Write-Debug "Failed to match 'TaskLibraryTemplateConfig' defined in oneOf (AddTasks200ResponseAllOfTask). Proceeding to the next one if any."
        }

        # try to match TaskNestedWorkflowConfig defined in the oneOf schemas
        try {
            $matchInstance = ConvertFrom-JsonToTaskNestedWorkflowConfig $Json

            foreach($property in $matchInstance.PsObject.Properties) {
                if ($null -ne $property.Value) {
                    $matchType = "TaskNestedWorkflowConfig"
                    $match++
                    break
                }
            }
        } catch {
            # fail to match the schema defined in oneOf, proceed to the next one
            Write-Debug "Failed to match 'TaskNestedWorkflowConfig' defined in oneOf (AddTasks200ResponseAllOfTask). Proceeding to the next one if any."
        }

        # try to match TaskPowershellConfig defined in the oneOf schemas
        try {
            $matchInstance = ConvertFrom-JsonToTaskPowershellConfig $Json

            foreach($property in $matchInstance.PsObject.Properties) {
                if ($null -ne $property.Value) {
                    $matchType = "TaskPowershellConfig"
                    $match++
                    break
                }
            }
        } catch {
            # fail to match the schema defined in oneOf, proceed to the next one
            Write-Debug "Failed to match 'TaskPowershellConfig' defined in oneOf (AddTasks200ResponseAllOfTask). Proceeding to the next one if any."
        }

        # try to match TaskPuppetConfig defined in the oneOf schemas
        try {
            $matchInstance = ConvertFrom-JsonToTaskPuppetConfig $Json

            foreach($property in $matchInstance.PsObject.Properties) {
                if ($null -ne $property.Value) {
                    $matchType = "TaskPuppetConfig"
                    $match++
                    break
                }
            }
        } catch {
            # fail to match the schema defined in oneOf, proceed to the next one
            Write-Debug "Failed to match 'TaskPuppetConfig' defined in oneOf (AddTasks200ResponseAllOfTask). Proceeding to the next one if any."
        }

        # try to match TaskPythonConfig defined in the oneOf schemas
        try {
            $matchInstance = ConvertFrom-JsonToTaskPythonConfig $Json

            foreach($property in $matchInstance.PsObject.Properties) {
                if ($null -ne $property.Value) {
                    $matchType = "TaskPythonConfig"
                    $match++
                    break
                }
            }
        } catch {
            # fail to match the schema defined in oneOf, proceed to the next one
            Write-Debug "Failed to match 'TaskPythonConfig' defined in oneOf (AddTasks200ResponseAllOfTask). Proceeding to the next one if any."
        }

        # try to match TaskRestartConfig defined in the oneOf schemas
        try {
            $matchInstance = ConvertFrom-JsonToTaskRestartConfig $Json

            foreach($property in $matchInstance.PsObject.Properties) {
                if ($null -ne $property.Value) {
                    $matchType = "TaskRestartConfig"
                    $match++
                    break
                }
            }
        } catch {
            # fail to match the schema defined in oneOf, proceed to the next one
            Write-Debug "Failed to match 'TaskRestartConfig' defined in oneOf (AddTasks200ResponseAllOfTask). Proceeding to the next one if any."
        }

        # try to match TaskShellConfig defined in the oneOf schemas
        try {
            $matchInstance = ConvertFrom-JsonToTaskShellConfig $Json

            foreach($property in $matchInstance.PsObject.Properties) {
                if ($null -ne $property.Value) {
                    $matchType = "TaskShellConfig"
                    $match++
                    break
                }
            }
        } catch {
            # fail to match the schema defined in oneOf, proceed to the next one
            Write-Debug "Failed to match 'TaskShellConfig' defined in oneOf (AddTasks200ResponseAllOfTask). Proceeding to the next one if any."
        }

        # try to match TaskVroConfig defined in the oneOf schemas
        try {
            $matchInstance = ConvertFrom-JsonToTaskVroConfig $Json

            foreach($property in $matchInstance.PsObject.Properties) {
                if ($null -ne $property.Value) {
                    $matchType = "TaskVroConfig"
                    $match++
                    break
                }
            }
        } catch {
            # fail to match the schema defined in oneOf, proceed to the next one
            Write-Debug "Failed to match 'TaskVroConfig' defined in oneOf (AddTasks200ResponseAllOfTask). Proceeding to the next one if any."
        }

        # try to match TaskWriteAttributesConfig defined in the oneOf schemas
        try {
            $matchInstance = ConvertFrom-JsonToTaskWriteAttributesConfig $Json

            foreach($property in $matchInstance.PsObject.Properties) {
                if ($null -ne $property.Value) {
                    $matchType = "TaskWriteAttributesConfig"
                    $match++
                    break
                }
            }
        } catch {
            # fail to match the schema defined in oneOf, proceed to the next one
            Write-Debug "Failed to match 'TaskWriteAttributesConfig' defined in oneOf (AddTasks200ResponseAllOfTask). Proceeding to the next one if any."
        }

        if ($match -gt 1) {
            throw "Error! The JSON payload matches more than one type defined in oneOf schemas ([TaskAnsiblePlaybookConfig, TaskAnsibleTowerConfig, TaskChefBootstrapConfig, TaskEmailConfig, TaskGroovyConfig, TaskHttpConfig, TaskJavaConfig, TaskJrubyConfig, TaskLibraryScriptConfig, TaskLibraryTemplateConfig, TaskNestedWorkflowConfig, TaskPowershellConfig, TaskPuppetConfig, TaskPythonConfig, TaskRestartConfig, TaskShellConfig, TaskVroConfig, TaskWriteAttributesConfig]). JSON Payload: $($Json)"
        } elseif ($match -eq 1) {
            return [PSCustomObject]@{
                "ActualType" = ${matchType}
                "ActualInstance" = ${matchInstance}
                "OneOfSchemas" = @("TaskAnsiblePlaybookConfig", "TaskAnsibleTowerConfig", "TaskChefBootstrapConfig", "TaskEmailConfig", "TaskGroovyConfig", "TaskHttpConfig", "TaskJavaConfig", "TaskJrubyConfig", "TaskLibraryScriptConfig", "TaskLibraryTemplateConfig", "TaskNestedWorkflowConfig", "TaskPowershellConfig", "TaskPuppetConfig", "TaskPythonConfig", "TaskRestartConfig", "TaskShellConfig", "TaskVroConfig", "TaskWriteAttributesConfig")
            }
        } else {
            throw "Error! The JSON payload doesn't matches any type defined in oneOf schemas ([TaskAnsiblePlaybookConfig, TaskAnsibleTowerConfig, TaskChefBootstrapConfig, TaskEmailConfig, TaskGroovyConfig, TaskHttpConfig, TaskJavaConfig, TaskJrubyConfig, TaskLibraryScriptConfig, TaskLibraryTemplateConfig, TaskNestedWorkflowConfig, TaskPowershellConfig, TaskPuppetConfig, TaskPythonConfig, TaskRestartConfig, TaskShellConfig, TaskVroConfig, TaskWriteAttributesConfig]). JSON Payload: $($Json)"
        }
    }
}
