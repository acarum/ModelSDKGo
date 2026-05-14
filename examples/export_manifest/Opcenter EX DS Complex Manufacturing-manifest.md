# Manifest Report: Opcenter EX DS Complex Manufacturing

**Mendix Version:** 11.6.4  
**MPR File:** C:\Workspaces\Mendix\Complex\main\Opcenter EX DS Complex Manufacturing.mpr  
**Generated:** 2026-05-08 18:49:11  

---

## Summary

- **External Entities:** 87 (across 6 modules)
- **Microflow/Action Calls:** 82
- **Signal Manager Subscriptions:** 64 subscription(s)
- **Navigation Items:** 3
- **Pages with Commands:** 1 page(s) analyzed
- **Command Buttons:** 13 total, 0 with extracted commands

---

## 1. External Entities

External OData entities used in the project, grouped by module.

### Module: EXFN_DocumentViewer

#### Entity: Document

**Published From:** Document

| Attribute | Type |
|-----------|------|
| _Id | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| EntityType | String |
| IsLocked | Boolean |
| ToBeCleaned | String |
| Revision | String |
| SourceRevision | String |
| IsCurrent | Boolean |
| NId | String |
| Name | String |
| Description | String |
| MIMEType | String |
| FileName | String |
| IconName | String |
| Icon_Id_Id | String |
| LocalFile_Id_Id | String |
| Url | String |
| RepositoryType | String |

#### Entity: File

**Published From:** Document

| Attribute | Type |
|-----------|------|
| _Id | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| EntityType | String |
| IsLocked | Boolean |
| ToBeCleaned | String |
| Name | String |
| _Type | String |
| PayloadSize | String |
| CreatedBy | String |
| LastUpdatedBy | DateTime |
| Folder_Id | String |
| Contents | String |

#### Entity: Category

**Published From:** Document

| Attribute | Type |
|-----------|------|
| _Id | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| EntityType | String |
| IsLocked | Boolean |
| ToBeCleaned | String |
| NId | String |
| Document_Id | String |

#### Entity: DocumentEntityLink

**Published From:** Document

| Attribute | Type |
|-----------|------|
| _Id | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| EntityType | String |
| IsLocked | Boolean |
| ToBeCleaned | String |
| DocumentNId | String |
| DocumentRevision | Boolean |
| LinkedEntityType | String |
| LinkedEntityId | String |

#### Entity: DocumentCategory

**Published From:** Document

| Attribute | Type |
|-----------|------|
| _Id | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| EntityType | String |
| IsLocked | Boolean |
| ToBeCleaned | String |
| NId | String |
| Name | String |
| Description | String |

#### Entity: DocumentSet

**Published From:** Document

| Attribute | Type |
|-----------|------|
| _Id | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| EntityType | String |
| IsLocked | Boolean |
| ToBeCleaned | String |
| Revision | String |
| SourceRevision | String |
| IsCurrent | Boolean |
| NId | String |
| Name | String |
| Description | String |

#### Entity: DocumentSetEntityLink

**Published From:** Document

| Attribute | Type |
|-----------|------|
| _Id | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| EntityType | String |
| IsLocked | Boolean |
| ToBeCleaned | String |
| DocumentSetNId | String |
| DocumentSetRevision | Boolean |
| LinkedEntityType | String |
| EffectivityExpression | String |
| LinkedEntityId | String |

### Module: EXFN_Quality

#### Entity: RMI_TimeBasedEngineDetails

**Published From:** Quality

| Attribute | Type |
|-----------|------|
| MaterialNId | String |
| MaterialRevision | String |
| InspectionDefinitionNId | String |
| InspectionDefinitionRevision | String |
| EngineState | String |
| CurrentTimerExpiration | DateTime |
| ContainerNId | String |
| _Id | String |
| EngineId | String |
| ExpiresWhen | String |

#### Entity: RMI_VisualDetectedFailuresCoordinates

**Published From:** Quality

| Attribute | Type |
|-----------|------|
| RuntimeInspectionDefinitionNId | String |
| RuntimeInspectionDefinitionNIdInContext | DateTime |
| MaterialNId | String |
| EquipmentNId | String |
| InspectionSampleId | String |
| VisualDetectedFailureId | String |
| SampleId | String |
| MaterialTrackingUnitId | String |
| FailureNId | String |
| XCoordinate | String |
| YCoordinate | String |
| Color | String |
| _Id | String |
| RuntimeInspectionDefinitionId | String |

#### Entity: RMI_RuntimeInspectionDefinition

**Published From:** Quality

| Attribute | Type |
|-----------|------|
| RuntimeInspectionDefinitionContainerNId | String |
| RuntimeInspectionDefinitionNId | String |
| InspectionDefinitionNId | String |
| InspectionDefinitionRevision | String |
| InspectionDefinitionName | String |
| InspectionDefinitionCanBeSkipped | String |
| InspectionDefinitionSampleSize | String |
| MandatoryExecutionsCompleted | String |
| InspectionDefinitionFrequencyNId | String |
| CharacteristicNId | String |
| CharacteristicName | String |
| CharacteristicDescription | String |
| CharacteristicCriticality | String |
| CharacteristicType | String |
| AttributiveNOKDescription | String |
| AttributiveOKDescription | String |
| SketchRows | String |
| SketchColumns | String |
| Id_Sketch | String |
| VariableLowerToleranceUoM | String |
| VariableLowerToleranceValue | Integer/Decimal |
| VariableUpperToleranceUoM | String |
| VariableUpperToleranceValue | Integer/Decimal |
| VariableNominalUoM | Integer/Decimal |
| VariableNominalValue | Integer/Decimal |
| InspectionAcquisitionContextId | String |
| InspectionSampleId | String |
| AnyViolation | String |
| InspectionSampleTimestamp | DateTime |
| InspectionSampleUser | String |
| InspectionSampleIsConfirmed | String |
| InspectionValueId | String |
| MaterialTrackingUnitNId | String |
| MeasuredAttributeValue | Integer/Decimal |
| MeasuredVariableValue | Integer/Decimal |
| InspectionValueTimestamp | DateTime |
| InspectionValueUser | Integer/Decimal |
| FailureNId | String |
| FailureRevision | String |
| ScenarioInstanceId | String |
| ScenarioConfigurationNId | String |
| VisualCount | Integer |
| _Id | String |
| RuntimeInspectionDefinitionContainerId | String |
| RuntimeInspectionDefinitionId | String |
| IsSignaturePending | Boolean |
| IsSPCCalculated | Boolean |
| PMI | String |
| IsFAIRequired | Boolean |

#### Entity: RMI_PartAndUnitBasedEngineDetails

**Published From:** Quality

| Attribute | Type |
|-----------|------|
| MaterialNId | String |
| MaterialRevision | String |
| EquipmentNId | String |
| CharacteristicRepresentationNId | String |
| CharacteristicRepresentationRevision | String |
| FrequencyType | String |
| RuntimeNumber | DateTime |
| ContainerNId | String |
| _Id | String |
| Serial | String |
| CriteriaNId | String |

#### Entity: RMI_PotentialFailureDetails

**Published From:** Quality

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Revision | String |
| CharacteristicSpecificationNId | String |
| IsCurrent | Boolean |
| _Id | String |
| FailureId | String |
| NumberOfChildren | String |
| NumberOfParents | String |
| FailureParentNId | String |
| FailureParentRevision | String |

#### Entity: InspectionValue

**Published From:** WorkInstruction

| Attribute | Type |
|-----------|------|
| _Id | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| EntityType | String |
| IsLocked | Boolean |
| ToBeCleaned | String |
| MeasuredAttributeValue | Integer/Decimal |
| MeasuredVariableValue | Integer/Decimal |
| MTUNId | String |
| Timestamp | DateTime |
| User | String |
| FailureNId | String |
| FailureRevision | String |
| InspectionSample_Id | String |

#### Entity: RMI_InspectionAcquisitionContext

**Published From:** Quality

| Attribute | Type |
|-----------|------|
| MaterialNId | String |
| MaterialRevision | String |
| EquipmentNId | String |
| RuntimeInspectionDefinitionNId | String |
| CalculationId | String |
| CalculationName | String |
| CalculationResult | String |
| CalculationResultId | String |
| InspectionDefinitionNId | String |
| ControlChartId | String |
| ControlChartNId | String |
| ControlChartName | String |
| _Id | String |
| AcquisitionContextId | String |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |

#### Entity: RMI_InspectionValue

**Published From:** WorkInstruction

| Attribute | Type |
|-----------|------|
| MeasuredVariableValue | Integer/Decimal |
| MTUNId | String |
| Timestamp | DateTime |
| User | String |
| FailureNId | String |
| FailureRevision | String |
| MTUCode | String |
| MeasuredAttributeValue | Integer/Decimal |
| CharacteristicTypeNId | String |
| _Id | String |
| InspectionSampleId | String |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |

### Module: EXFN_WorkInstruction

#### Entity: WorkInstructionDefinition

**Published From:** WorkInstruction

| Attribute | Type |
|-----------|------|
| _Id | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| EntityType | String |
| IsLocked | Boolean |
| ToBeCleaned | String |
| Revision | String |
| SourceRevision | String |
| IsCurrent | Boolean |
| NId | String |
| Name | String |
| Description | String |
| IsExecutable | Boolean |
| WorkInstructionTemplateNId | String |

### Module: OpcenterEXDS_EXFN_Quality

#### Entity: RM_RuntimeInspectionDefinition

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| RuntimeInspectionDefinitionContainerNId | String |
| RuntimeInspectionDefinitionNId | String |
| InspectionDefinitionNId | String |
| InspectionDefinitionRevision | String |
| InspectionDefinitionCanBeSkipped | String |
| InspectionDefinitionSampleSize | String |
| MandatoryExecutionsCompleted | String |
| InspectionDefinitionFrequencyNId | String |
| InspectionDefinitionFrequencyName | String |
| InspectionDefinitionFrequencyDescription | String |
| CharacteristicNId | String |
| CharacteristicName | String |
| CharacteristicDescription | String |
| CharacteristicContext | String |
| CharacteristicCriticality | String |
| CharacteristicType | String |
| AttributiveNOKDescription | String |
| AttributiveOKDescription | String |
| SketchRows | String |
| SketchColumns | String |
| SketchMimeType | String |
| Id_Sketch | String |
| VariableLowerToleranceUoM | String |
| VariableLowerToleranceValue | Integer/Decimal |
| VariableUpperToleranceUoM | String |
| VariableUpperToleranceValue | Integer/Decimal |
| VariableNominalUoM | Integer/Decimal |
| VariableNominalValue | Integer/Decimal |
| InspectionAcquisitionContextId | String |
| MaterialNId | String |
| MaterialRevision | String |
| InspectionSampleId | String |
| EquipmentNId | String |
| AnyViolation | String |
| InspectionSampleTimestamp | DateTime |
| InspectionSampleUser | String |
| InspectionSampleIsConfirmed | String |
| InspectionValueId | String |
| MaterialTrackingUnitNId | String |
| MeasuredAttributeValue | Integer/Decimal |
| MeasuredVariableValue | Integer/Decimal |
| InspectionValueTimestamp | DateTime |
| InspectionValueUser | Integer/Decimal |
| FailureNId | String |
| FailureRevision | String |
| InspectionValueLastUpdatedOn | DateTime |
| _Id | String |
| RuntimeInspectionDefinitionContainerId | String |
| RuntimeInspectionDefinitionId | String |
| InspectionDefinitionName | String |
| ScenarioInstanceId | String |
| ScenarioConfigurationNId | String |

#### Entity: RM_ToBeUsedInspection

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| RuntimeCharacteristicContainerId | String |
| WorkOrderOperationId | String |
| WorkOrderStepId | String |
| _Id | String |

#### Entity: RM_PotentialFailureDetails

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Revision | String |
| Description | String |
| CharacteristicSpecificationNId | String |
| CharacteristicSpecificationRevision | String |
| DocumentNId | String |
| DocumentRevision | Boolean |
| _Id | String |
| FailureId | String |
| CharacteristicSpecificationId | String |

#### Entity: RM_VisualDetectedFailuresCoordinates

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| RuntimeInspectionDefinitionNId | String |
| InspectionAcquisitionContextId | String |
| MaterialNId | String |
| MaterialRevision | String |
| EquipmentNId | String |
| AnyViolation | String |
| InspectionSampleTimestamp | DateTime |
| InspectionSampleUser | String |
| InspectionSampleIsConfirmed | String |
| InspectionSampleId | String |
| VisualDetectedFailureId | String |
| SampleId | String |
| MaterialTrackingUnitId | String |
| FailureId | String |
| FailureNId | String |
| XCoordinate | String |
| YCoordinate | String |
| Color | String |
| _Id | String |
| RuntimeInspectionDefinitionId | String |

### Module: OpcenterEXDS_OperatorLanding

#### Entity: RM_FullQuantity_SingleSerialized_ActualConsumedMaterial

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| ActualConsumedMaterialDM_MTUId | String |
| Code | String |
| ConsumptionDateTime | DateTime |
| MaterialItemAssembledQty | String |
| NId | String |
| TotalAssembledQty | String |
| WorkOrderOperationId | String |
| WorkOrderStepId | String |
| ProducedMaterialId | String |
| ActualConsumedMaterialId | String |
| ToBeConsumedMaterialId | String |
| _Id | String |

#### Entity: RM_FullQuantity_SingleSerialized_ToBeConsumedMaterial

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| DM_MaterialId | String |
| IsNotConsumed | Boolean |
| HasToBeConsumedBarcodeRule | Boolean |
| IsCustomSpecificationType | Boolean |
| IsPartiallyConsumed | Boolean |
| IsProducedMaterialNotConsumed | Boolean |
| IsProducedMaterialPartiallyConsumed | Boolean |
| IsProducedMaterialTotallyConsumed | Boolean |
| IsTotallyConsumed | Boolean |
| LogicalPosition | String |
| MaterialSpecificationType | String |
| Name | String |
| NId | String |
| Quantity | String |
| RemainingToBeConsumedQuantity | String |
| Revision | String |
| Sequence | Integer |
| SerialNumberProfile | String |
| ToBeConsumedCustomMaterialCount | Integer |
| ToBeConsumedMaterialCount | Integer |
| UoMNId | String |
| WorkOrderStepId | String |
| IsPrekitted | Boolean |
| PrekittedQuantity | String |
| PrekittedDM_MTUCount | Integer |
| GroupId | String |
| AllPrekitsToBeValidated | DateTime |
| AlternativeSelected | String |
| OccurrenceId | String |
| MaterialId | String |
| ToBeConsumedMaterialId | String |
| ProducedMaterialItemId | String |
| WorkOrderOperationId | String |
| CanAcquire | String |
| SelectedFitTargetQty | String |
| SelectedFitActualQty | String |

#### Entity: RM_HoldReason

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| NId | String |
| _Type | String |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| HoldReasonId | String |

#### Entity: RM_PauseReason

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| NId | String |
| Location | String |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |

#### Entity: RM_FullQuantity_SingleSerialized_WorkOrderOperation_AllowedRuntimeActionByEquipment

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| WorkOrderOperationNId | String |
| ActiveUserOnMachine | String |
| EquipmentNId | String |
| WorkOrderOperationId | String |
| IsPreferredEquipment | Boolean |
| CanStart | String |
| CanPause | String |
| CanComplete | String |
| EquipmentName | String |
| EquipmentLevelNId | String |
| CanCreateCP | String |
| DM_MaterialTrackingUnitId | String |
| MaterialTrackingUnitCode | String |
| MaterialTrackingUnitNId | String |
| MaterialTrackingUnitName | String |
| WorkOrderOperationStatusNId | String |
| MaterialNId | String |
| PartProgram | String |
| CanSkip | String |

#### Entity: RM_ToBeUsedTool

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| TimesToBeUsed | DateTime |
| ActualUsageCounter | Integer |
| ToolDefinitionId | String |
| ToolDefinitionIsConsumable | String |
| ToolDefinitionName | String |
| WorkOrderStepId | String |
| ProducedMaterialItemId | String |
| ToBeUsedToolId | String |
| WorkOrderOperationId | String |
| CanAcquire | String |
| _Id | String |

#### Entity: RM_FullQuantity_SingleSerialized_WorkOrderOperation_AllowedRuntimeAction

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| WorkOrderOperationId | String |
| ActiveUser | String |
| CanStart | String |
| CanPause | String |
| CanComplete | String |
| CanCreateCP | String |
| CanSkip | String |

#### Entity: RM_ToolHistory

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| ActionNId | String |
| UsageDuration | String |
| UsedToolCount | Integer |
| CreatedOn | DateTime |
| ToolDefinitionId | String |
| WorkOrderOperationId | String |
| WorkOrdeStepId | String |
| ToolNId | String |
| DM_MaterialTrackingUnitId | String |

#### Entity: RM_Tool

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| UsageCounter | Integer |
| ToolDefinitionId | String |
| ToolDefinitionNId | String |
| ToolDefinitionRevision | String |
| Status | String |
| IsScrap | Boolean |

#### Entity: RM_WorkOrderOperationSkill

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| WorkOrderOperationId | String |
| WorkOrderOperationNId | String |
| WorkOrderStepId | String |
| WorkOrderStepNId | String |
| SkillNId | String |
| SkillName | String |
| SkillColor | String |
| SkillId | String |

#### Entity: RM_FullQuantity_SingleSerialized_WorkOrderOperationOrStep

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| WorkOrderNId | String |
| WorkOrderName | String |
| ProductionTypeNId | String |
| MaterialNId | String |
| MaterialRevision | String |
| ProducedMaterialItemId | String |
| DM_MaterialTrackingUnitId | String |
| MaterialTrackingUnitNId | String |
| MaterialTrackingUnitCode | String |
| NId | String |
| Name | String |
| Description | String |
| StatusNId | String |
| EstimatedStartTime | DateTime |
| EstimatedEndTime | DateTime |
| ActualStartTime | DateTime |
| ActualEndTime | DateTime |
| EstimatedDuration | String |
| LastPauseTime | DateTime |
| Sequence | Integer |
| IsReady | Boolean |
| Revision | String |
| PauseDuration | String |
| AvailableQuantity | String |
| TargetQuantity | String |
| ProducedQuantity | String |
| WorkOrderOperationId | String |
| PartialWorkedQuantity | String |
| WorkOrderStepId | String |
| WorkOrderOperationName | String |
| ToBeCollectedDocument | String |
| HasToBeConsumedMaterials | Boolean |
| HasToBeUsedTools | Boolean |
| HasQualityInspections | Boolean |
| HasOpWorkInstructionDefinition | Boolean |
| HasSnWorkInstructionDefinition | Boolean |
| Plant | String |
| ActiveNonConformanceNr | String |
| ReworkedQuantity | String |
| ScrappedQuantity | String |
| WorkOrderId | String |
| MaterialTrackingUnitStatusNId | String |
| WorkOrderInitialQuantity | String |
| MaterialTrackingUnitId | String |
| MaterialId | String |
| MaterialsConsumptionInProgress | String |
| MaterialsConsumptionCompleted | String |
| ToolsUsageInProgress | String |
| ToolsUsageCompleted | String |
| WorkInstructionsInProgress | String |
| WorkInstructionsCompleted | String |
| HasChangePackage | Boolean |
| WorkOrderOperationNId | String |
| ParentWorkOrderOperationFolderId | String |
| ParentWorkOrderOperationFolderName | String |
| HasBuyOffQualityInspections | Boolean |
| IsFlexibleWOClosed | Boolean |
| WorkOrderStatusNId | String |
| WorkOrderProducedQuantity | String |
| IsDynamic | Boolean |
| ElectronicSignatureStart | String |
| ElectronicSignaturePause | String |
| ElectronicSignatureComplete | String |
| HasSnOnDemandWorkInstructionDefinition | Boolean |
| HasToBeCoProducedMaterials | Boolean |
| HasChildrenWorkOrderSteps | Boolean |
| _Id | String |
| IsCNC | Boolean |
| DM_MaterialId | String |
| IsWOFAIRelevant | Boolean |
| IsWOFAICandidate | Boolean |
| IsMaterialFAIRequired | Boolean |
| HasToBeUsedScrewingTools | Boolean |
| IsSkippable | Boolean |

#### Entity: RM_ComplexManufacturing_WorkOrderOperation

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| WorkOrderNId | String |
| WorkOrderName | String |
| ProductionTypeNId | String |
| MaterialNId | String |
| MaterialRevision | String |
| MaterialTrackingUnitCode | String |
| Name | String |
| StatusNId | String |
| Sequence | Integer |
| IsReady | Boolean |
| AvailableQuantity | String |
| ProducedQuantity | String |
| PartialWorkedQuantity | String |
| WorkOrderStatusNId | String |
| WorkOrderId | String |
| ActiveNonConformanceNr | String |
| ReworkedQuantity | String |
| HasChangePackage | Boolean |
| WorkOrderInitialQuantity | String |
| WorkOrderOperationId | String |
| _Id | String |
| ParentWorkOrderOperationFolderId | String |
| ParentWorkOrderOperationFolderName | String |
| TargetQuantity | String |
| ToBeCollectedDocument | String |
| DM_MaterialTrackingUnitId | String |
| Plant | String |
| IsDynamic | Boolean |
| ActualTargetQuantity | String |
| PlannedTargetQuantity | String |
| DM_MaterialId | String |
| IsFlexibleWOClosed | Boolean |
| WorkOrderEstimatedStartTime | DateTime |
| WorkOrderProducedQuantity | String |
| EstimatedStartTime | DateTime |
| ElectronicSignatureStart | String |
| ElectronicSignaturePause | String |
| ElectronicSignatureComplete | String |
| IsWOFAIRelevant | Boolean |
| IsWOFAICandidate | Boolean |
| IsSkippable | Boolean |
| NId | String |

#### Entity: RM_FullQuantity_SingleSerialized_WorkOrderStep_AllowedRuntimeActionByEquipment

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| WorkOrderStepNId | String |
| WoOpActiveUserOnMachine | String |
| IsUserActiveOnStep | Boolean |
| EquipmentNId | String |
| EquipmentName | String |
| EquipmentLevelNId | String |
| CanStart | String |
| CanComplete | String |
| WorkOrderStepId | String |
| IsPreferredEquipment | Boolean |
| WorkOrderOperationId | String |
| WorkOrderOperationNId | String |
| DM_MaterialTrackingUnitId | String |
| MaterialTrackingUnitNId | String |
| MaterialTrackingUnitCode | String |

#### Entity: RM_ToBeUsedWorkInstruction

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| WorkInstructionDefinitionId | String |
| WorkOrderOperationId | String |
| WorkOrderStepId | String |
| WorkInstructionAssociationTypeNId | String |
| WorkInstructionId | String |
| DM_MaterialTrackingUnitId | String |
| WIToBeShown | String |
| IsWorkInstructionShared | Boolean |
| LastUpdatedOn | DateTime |
| WorkInstructionDefinitionNId | String |
| WorkInstructionDefinitionName | String |
| WorkInstructionDefinitionRevision | String |
| WorkInstructionDefinitionIsCurrent | String |
| WorkInstructionDefinitionIsLocked | Boolean |

#### Entity: RM_FailureDetails

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| _Id | String |
| NId | String |
| Name | String |
| Revision | String |
| Description | String |
| FailureChildId | String |
| CharacteristicSpecificationId | String |
| CharacteristicSpecificationNId | String |
| CharacteristicSpecificationRevision | String |
| DocumentNId | String |
| DocumentRevision | Boolean |
| FailureId | String |
| HasChildren | Boolean |
| IsChildFailure | Boolean |

#### Entity: RM_Note

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| WorkOrderOperation | String |
| OpenDate | DateTime |
| OpenDateTime | DateTime |
| OpenUser | String |
| Status | String |
| CloseDate | DateTime |
| CloseDateTime | DateTime |
| CloseUser | String |
| Message | String |
| WorkOrder | String |
| IsWorkOrderLevel | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| NoteId | String |
| _Id | String |

#### Entity: RM_Prekit

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| HasToBeValidated | Boolean |
| ProducedDM_MTUId | String |
| ToBeConsumedMaterialId | String |
| ToBeConsumedDM_MTUId | String |
| ToBeConsumedPrekittedQuantity | String |
| ToBeConsumedPrekittedCode | String |
| ToBeConsumedPrekittedQuantityUoM | String |
| PrekitId | String |

#### Entity: RM_MaterialTrackingUnit

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| NId | String |
| Code | String |
| Name | String |
| Quantity | String |
| Status | String |
| MaterialNId | String |
| MaterialRevision | String |
| MaterialTrackingUnitId | String |
| QuantityUom | String |
| _Id | String |

#### Entity: RM_OpLandingDocumentList

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| Category | String |
| WorkOrderOperation | String |
| WorkOrderStep | String |
| ExecutionGroupPhase | String |
| SerialNumber | String |
| FileName | String |
| MIMEType | String |
| NId | String |
| _Type | String |
| UId | String |
| FileId | String |
| IconId | String |
| ContainerTypeId | String |
| Revision | String |
| Description | String |
| Name | String |
| Document | Boolean |

#### Entity: RM_Document

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Description | String |
| MIMEType | String |
| FileName | String |
| IconName | String |
| RepositoryType | String |
| Revision | String |
| SourceRevision | String |
| IsCurrent | Boolean |
| CreatedOn | DateTime |
| DocumentId | String |
| _Id | String |

#### Entity: RM_FullQuantity_SingleSerialized_HistoryDisassembledMaterial

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| EquipmentNId | String |
| DisassembledQuantity | String |
| DisassembleDateTime | DateTime |
| ToBeConsumedMaterialId | String |
| ToBeConsumedMaterialNId | String |
| ToBeConsumedMaterialName | String |
| MaterialSpecificationType | String |
| MaterialNId | String |
| MaterialRevision | String |
| MaterialTrackingUnitNId | String |
| MaterialTrackingUnitName | String |
| MaterialTrackingUnitCode | String |
| MaterialTrackingUnitCodeType | String |
| ActiveNonConformanceNumber | String |
| MTUStateMachineStatus | String |
| MaterialTrackingUnitStatus | String |
| DMMaterialTrackingUnitId | String |
| WorkOrderOperationId | String |
| WorkOrderStepId | String |
| ProducedMaterialItemId | String |

#### Entity: RM_WIDefinitionATNInstanceParameterAssociation_Machine

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| EquipmentNId | String |
| EquipmentName | String |
| LevelNId | String |
| WorkInstructionDefinitionNId | String |
| User | String |
| WorkOrderOperationId | String |
| WorkOrderStepId | String |
| WorkInstructionDefinitionId | String |

#### Entity: RM_TobeUsedToolHistory

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| WorkOrderOperationId | String |
| WorkOrderStepId | String |
| TimesToBeUsed | DateTime |
| ToolDefinitionId | String |
| ToolDefinitionNId | String |
| ToolNId | String |
| ToolUsageDuration | String |
| ActionNId | String |
| CreatedOn | DateTime |
| ActualUsageCounter | Integer |
| CanAcquire | String |
| IsHistoryVisible | Boolean |
| ToolDefinitionIsConsumable | String |
| ProducedMaterialItemId | String |

#### Entity: RM_GetFilteredFailureDetails

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Revision | String |
| Description | String |
| HasChildren | Boolean |
| IsChildFailure | Boolean |
| FailureChildId | String |
| CharacteristicSpecificationId | String |
| CharacteristicSpecificationNId | String |
| CharacteristicSpecificationRevision | String |
| DocumentNId | String |
| DocumentRevision | Boolean |
| MaterialNId | String |
| EquipmentNId | String |
| _Context | String |
| _Id | String |
| FailureId | String |

#### Entity: RM_WOOFolderBreadcrumb

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| ChildFolderId | String |
| ChildFolderNId | String |
| ParentFolderId | String |
| Hierarchy | String |
| HierarchyLabel | String |
| _Id | String |

#### Entity: RM_RoutingNode

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| WorkOrderNId | String |
| WorkOrderName | String |
| WorkOrderStatusNId | String |
| EntityNId | String |
| EntityName | String |
| EntitySequence | Integer |
| EntityIsReady | String |
| EntityStatusNId | String |
| EntityHasPendingNC | String |
| ParentEntityId | String |
| ParentEntityType | String |
| EntityHasChildren | String |
| EntityType | String |
| _Id | String |
| WorkOrderId | String |
| EntityId | String |
| CurrentWorkOrderId | String |
| IsSkippable | Boolean |

#### Entity: RM_RoutingEdge

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| FromWorkOrderId | String |
| FromNodeEntityId | String |
| FromNodeEntityType | String |
| ToWorkOrderId | String |
| ToNodeEntityId | String |
| ToNodeEntityType | String |
| DependencyType | String |
| IsExternalDependency | Boolean |
| _Id | String |
| FromNodeParentEntityId | String |
| ToNodeParentEntityId | String |
| CurrentWorkOrderId | String |
| AlternativeGroup | String |
| IsPreferred | Boolean |

#### Entity: RM_ProducedMaterialItem_MTUStatus

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| ParentDM_MTUId | String |
| DM_MaterialTrackingUnitVolume | String |
| DM_MaterialTrackingUnitWeight | String |
| MaterialTrackingUnitNId | String |
| MaterialTrackingUnitName | String |
| MaterialTrackingUnitCode | String |
| MaterialTrackingUnitCodeType | String |
| MaterialTrackingUnitStateMachineNId | String |
| MaterialTrackingUnitStatusNId | String |
| MaterialTrackingUnitQuantity | String |
| MaterialTrackingUnitUoMNId | String |
| MaterialTrackingUnitAggregateId | String |
| DM_MaterialId | String |
| MaterialNId | String |
| MaterialRevision | String |
| _Id | String |
| ProducedMaterialItemId | String |
| DM_MaterialTrackingUnitId | String |
| MaterialTrackingUnitId | String |
| WorkOrderOperationId | String |
| IsOpen | Boolean |
| IsActive | Boolean |
| IsPaused | Boolean |
| IsCompleted | Boolean |
| IsReady | Boolean |
| CanRepeatKO | String |
| IsResultKO | Boolean |
| IsSkipped | Boolean |
| IsUnderRework | Boolean |

#### Entity: RM_MaterialTrackingUnit_AvailableOnWorkOrderOperation_MTUStatus

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| WorkOrderOperationId | String |
| EquipmentNId | String |
| UserId | String |
| ParentDM_MTUId | String |
| DM_MaterialTrackingUnitVolume | String |
| DM_MaterialTrackingUnitWeight | String |
| DM_MaterialId | String |
| MaterialTrackingUnitNId | String |
| MaterialTrackingUnitName | String |
| MaterialTrackingUnitCode | String |
| MaterialTrackingUnitCodeType | String |
| MaterialTrackingUnitStateMachineNId | String |
| MaterialTrackingUnitStatusNId | String |
| MaterialTrackingUnitQuantity | String |
| MaterialTrackingUnitUoMNId | String |
| MaterialTrackingUnitAggregateId | String |
| MaterialNId | String |
| MaterialRevision | String |
| _Id | String |
| DM_MaterialTrackingUnitId | String |
| MaterialTrackingUnitId | String |
| IsStarted | Boolean |
| IsPaused | Boolean |
| IsCompleted | Boolean |

#### Entity: RM_BuyOff

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| WorkOrderOperationId | String |
| WorkOrderStepId | String |
| MaterialTrackingUnitCode | String |
| StatusNId | String |
| _Id | String |
| BuyOffId | String |
| DM_MaterialTrackingUnitId | String |
| MaterialTrackingUnitId | String |
| NonConformanceNId | String |

#### Entity: RM_MaterialTrackingUnit_AvailableOnWorkOrderStep_MTUStatus

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| WorkOrderStepId | String |
| EquipmentNId | String |
| UserId | String |
| ParentDM_MTUId | String |
| DM_MaterialTrackingUnitVolume | String |
| DM_MaterialTrackingUnitWeight | String |
| DM_MaterialId | String |
| MaterialTrackingUnitNId | String |
| MaterialTrackingUnitName | String |
| MaterialTrackingUnitCode | String |
| MaterialTrackingUnitCodeType | String |
| MaterialTrackingUnitStateMachineNId | String |
| MaterialTrackingUnitStatusNId | String |
| MaterialTrackingUnitQuantity | String |
| MaterialTrackingUnitUoMNId | String |
| MaterialTrackingUnitAggregateId | String |
| MaterialNId | String |
| MaterialRevision | String |
| _Id | String |
| DM_MaterialTrackingUnitId | String |
| MaterialTrackingUnitId | String |
| IsStarted | Boolean |
| IsPaused | Boolean |
| IsCompleted | Boolean |

#### Entity: RM_Tool_AvailableForNonConformance

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| ToolDefinitionNId | String |
| ToolDefinitionName | String |
| ToolDefinitionVersion | String |
| ToolDefinitionConsumable | String |
| LogisticClass | String |
| NId | String |
| Name | String |
| Description | String |
| Status | String |
| ExpirationDate | DateTime |
| UsageCounter | Integer |
| UsageCounterMax | Integer |
| UsageDuration | String |
| UsageDurationMax | Integer/Decimal |
| ActiveNonConformanceNr | String |
| Lockable | String |
| IsLock | Boolean |
| IsRuntimeFrozen | Boolean |
| IsScrap | Boolean |
| _Id | String |

#### Entity: RM_Equipment_AvailableForNonConformance

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Description | String |
| LevelNId | String |
| StateMachine | String |
| Status | String |
| _Id | String |

#### Entity: RM_NonProductiveActivity

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| NId | String |
| Description | String |
| _Id | String |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |

#### Entity: RM_NonProductiveActivityContext

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| NId | String |
| _Id | String |

#### Entity: RM_AssignedNonProductiveActivity

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| NpaNId | String |
| NpaDescription | String |
| NpaContext | String |
| NpaContextValue | Integer/Decimal |
| User | String |
| StartDate | DateTime |
| EndDate | DateTime |
| NpaContextValueName | Integer/Decimal |
| _Id | String |
| NpaId | String |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |

#### Entity: RM_CMX_DM_MaterialTrackingUnit

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| CodeType | String |
| Code | String |
| Description | String |
| MaterialNId | String |
| MaterialRevision | String |
| Quantity | String |
| QuantityUom | String |
| _Id | String |
| MaterialTrackingUnitId | String |

#### Entity: RM_Equipment

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Description | String |
| LevelNId | String |
| Status | String |
| PropertyValue | Integer/Decimal |
| CompleteByDifferentUserEnabled | Boolean |
| _Id | String |
| EquipmentId | String |
| EquipmentConfigurationId | String |

#### Entity: RM_CMX_WorkOrder

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| MaterialNId | String |
| MaterialName | String |
| MaterialRevision | String |
| StatusNId | String |
| InitialQuantity | String |
| _Id | String |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |

#### Entity: RM_CMX_ToBeCoProducedMaterial

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Revision | String |
| UoMNId | String |
| SerialNumberProfile | String |
| MaterialSpecificationType | String |
| IsCustomSpecificationType | Boolean |
| Quantity | String |
| RemainingToBeCoProducedQuantity | String |
| ActualCoProducedMaterialQuantity | String |
| LogicalPosition | String |
| HasToBeCoProducedBarcodeRule | Boolean |
| IsTotallyCoProduced | Boolean |
| IsPartiallyCoProduced | Boolean |
| IsNotCoProduced | Boolean |
| ToBeCoProducedCustomMaterialCount | Integer |
| OccurrenceId | String |
| WorkOrderOperationId | String |
| WorkOrderStepId | String |
| _Id | String |
| ToBeCoProducedMaterialId | String |
| Sequence | Integer |
| DM_MaterialId | String |
| MaterialId | String |
| IsMaterialFAIRequired | Boolean |

#### Entity: RM_CMX_ActualCoProducedMaterial

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| NId | String |
| Code | String |
| MaterialItemCoProducedQty | String |
| ActualCoProducedMaterialDM_MTUId | String |
| WorkOrderOperationId | String |
| WorkOrderStepId | String |
| TotalCoProducedQty | String |
| DisassembleMTUCode | String |
| _Id | String |
| ActualCoProducedMaterialId | String |
| ToBeCoProducedMaterialId | String |
| LastUpdatedOn | DateTime |

#### Entity: RM_Container

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Code | String |
| CodeType | String |
| EquipmentNId | String |
| Status | String |
| _Id | String |

#### Entity: RM_CMX_NumberingPatternMaterialAssociation

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| _Id | String |
| EntityTypeNId | String |
| DMMaterialId | String |
| NumberingPatternNId | String |
| NumberingPatternId | String |

#### Entity: RM_CMX_WoOp_FQ_SingleSer_ToBeConsumedMaterial

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Revision | String |
| UoMNId | String |
| SerialNumberProfile | String |
| MaterialSpecificationType | String |
| IsCustomSpecificationType | Boolean |
| Sequence | Integer |
| Quantity | String |
| RemainingToBeConsumedQuantity | String |
| LogicalPosition | String |
| DM_MaterialId | String |
| HasToBeConsumedBarcodeRule | Boolean |
| IsTotallyConsumed | Boolean |
| IsPartiallyConsumed | Boolean |
| IsNotConsumed | Boolean |
| ProducedMTUNId | String |
| ProducedMTUNName | String |
| ProducedMTUNCode | String |
| IsProducedMaterialTotallyConsumed | Boolean |
| IsProducedMaterialPartiallyConsumed | Boolean |
| IsProducedMaterialNotConsumed | Boolean |
| ToBeConsumedCustomMaterialCount | Integer |
| ToBeConsumedMaterialCount | Integer |
| OccurrenceId | String |
| WorkOrderOperationId | String |
| GroupId | String |
| AlternativeSelected | String |
| IsPrekitted | Boolean |
| PrekittedQuantity | String |
| PrekittedDM_MTUCount | Integer |
| AllPrekitsToBeValidated | DateTime |
| CanAcquire | String |
| SelectedFitTargetQty | String |
| SelectedFitActualQty | String |
| _Id | String |
| ToBeConsumedMaterialId | String |
| MaterialId | String |
| ProducedMaterialItemId | String |
| ProducedDM_MTUtId | String |
| ProducedMTUId | String |
| IsMaterialFAIRequired | Boolean |

#### Entity: RM_CMX_WoStep_FQ_SingleSer_ToBeConsumedMaterial

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Revision | String |
| UoMNId | String |
| SerialNumberProfile | String |
| MaterialSpecificationType | String |
| IsCustomSpecificationType | Boolean |
| Sequence | Integer |
| Quantity | String |
| RemainingToBeConsumedQuantity | String |
| LogicalPosition | String |
| DM_MaterialId | String |
| HasToBeConsumedBarcodeRule | Boolean |
| IsTotallyConsumed | Boolean |
| IsPartiallyConsumed | Boolean |
| IsNotConsumed | Boolean |
| ProducedMTUNId | String |
| ProducedMTUNName | String |
| ProducedMTUNCode | String |
| IsProducedMaterialTotallyConsumed | Boolean |
| IsProducedMaterialPartiallyConsumed | Boolean |
| IsProducedMaterialNotConsumed | Boolean |
| ToBeConsumedCustomMaterialCount | Integer |
| ToBeConsumedMaterialCount | Integer |
| OccurrenceId | String |
| WorkOrderStepId | String |
| GroupId | String |
| AlternativeSelected | String |
| IsPrekitted | Boolean |
| PrekittedQuantity | String |
| PrekittedDM_MTUCount | Integer |
| AllPrekitsToBeValidated | DateTime |
| CanAcquire | String |
| SelectedFitTargetQty | String |
| SelectedFitActualQty | String |
| _Id | String |
| ToBeConsumedMaterialId | String |
| MaterialId | String |
| ProducedMaterialItemId | String |
| ProducedDM_MTUtId | String |
| ProducedMTUId | String |
| IsMaterialFAIRequired | Boolean |

#### Entity: RM_WoOp_WITaskStatus

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| WorkInstructionsInProgress | String |
| WorkInstructionsCompleted | String |
| _Id | String |
| WorkOrderOperationId | String |

#### Entity: RM_WoStep_WITaskStatus

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| WorkInstructionsInProgress | String |
| WorkInstructionsCompleted | String |
| _Id | String |
| WorkOrderStepId | String |

#### Entity: RM_ToBeUsedMachineWithEquipmentHierarchy

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| EquipmentNId | String |
| _Id | String |
| RM_ComplexManufacturing_WorkOrderOperation_Id | String |

#### Entity: RM_CMX_Filter_ProducedMaterialItem

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| MaterialTrackingUnitCode | String |
| _Id | String |
| RM_ComplexManufacturing_WorkOrderOperation_Id | String |

#### Entity: RM_CMX_FAIWorkOrder

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| _Id | String |
| WorkOrderId | String |
| NId | String |
| IsRelevant | Boolean |
| IsCandidate | Boolean |
| CandidateDM_MTUId | String |
| IsFAICandidateDeclared | Boolean |

#### Entity: RM_ProducedMaterialItem

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| WorkOrderId | String |
| ParentDM_MTUId | String |
| DM_MaterialTrackingUnitVolume | String |
| DM_MaterialTrackingUnitWeight | String |
| MaterialTrackingUnitNId | String |
| MaterialTrackingUnitName | String |
| MaterialTrackingUnitCode | String |
| MaterialTrackingUnitCodeType | String |
| MaterialTrackingUnitStateMachineNId | String |
| MaterialTrackingUnitStatusNId | String |
| MaterialTrackingUnitQuantity | String |
| MaterialTrackingUnitUoMNId | String |
| MaterialTrackingUnitAggregateId | String |
| DM_MaterialId | String |
| MaterialNId | String |
| MaterialRevision | String |
| _Id | String |
| DM_MaterialTrackingUnitId | String |
| MaterialTrackingUnitId | String |
| WorkOrderOperationId | String |
| ProducedMaterialItemId | String |
| RM_WorkOrderOperationId | String |

#### Entity: RM_CMX_WorkOrderOperation_ActualProducedMaterial

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| PartialWorkedQuantity | String |
| WorkOrderOperationId | String |
| MaterialTrackingUnitNId | String |
| MaterialTrackingUnitCode | String |
| MaterialNId | String |
| _Id | String |
| ActualProducedMaterialId | String |
| MaterialTrackingUnitId | String |

#### Entity: RM_CMX_SetPointItem_Equipment_DefaultValue

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| SetPointId | String |
| NId | String |
| Name | String |
| Description | String |
| DataType | String |
| IsEditable | Boolean |
| EquipmentNId | String |
| DefaultValue | Integer/Decimal |
| _Id | String |
| SetPointItemId | String |

#### Entity: RM_CMX_MaterialTrackingUnitAndContainer

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| NId | String |
| Code | String |
| MaterialTrackingUnitName | String |
| Quantity | String |
| QuantityUom | String |
| Status | String |
| MaterialUId | String |
| MaterialNId | String |
| MaterialRevision | String |
| MaterialTrackingUnitDescription | String |
| MaterialTrackingUnitCodeType | String |
| ContainerId | String |
| ContainerNId | String |
| ContainerName | String |
| _Id | String |
| MaterialTrackingUnitId | String |

#### Entity: RM_CMX_FAIRelevantWorkOrderOperation

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| WorkOrderNId | String |
| WorkOrderOperationNId | String |
| WorkOrderOperationName | String |
| _Id | String |
| WorkOrderId | String |
| WorkOrderOperationId | String |

#### Entity: RM_CMX_ScrewingToolHistory

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| ToBeUsedToolId | String |
| ToolNId | String |
| DM_MaterialTrackingUnitId | String |
| MaterialTrackingUnitNId | String |
| BoltNumber | String |
| Angle | String |
| Torque | String |
| IsOk | Boolean |
| _Id | String |
| CreatedOn | DateTime |

#### Entity: RM_CMX_ToBeUsedScrewingTool

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| ToolDefinitionNId | String |
| AutomationNodeInstanceNId | String |
| ToolId | String |
| ToolNId | String |
| EquipmentNId | String |
| NumberOfBolts | String |
| AngleValueMin | Integer/Decimal |
| AngleValueMax | Integer/Decimal |
| TorqueValueMin | Integer/Decimal |
| TorqueValueMax | Integer/Decimal |
| DM_MaterialTrackingUnitId | String |
| MaterialTrackingUnitNId | String |
| IsActiveMaterialTrackingUnit | Boolean |
| OkCount | Integer |
| NotOkCount | Integer |
| WorkOrderOperationId | String |
| ToolDefinitionName | String |
| ToolDefinitionConsumable | String |
| ToolName | String |
| RemainingBolts | String |
| _Id | String |
| ToBeUsedToolId | String |
| ToolDefinitionId | String |

#### Entity: SkipReason

**Published From:** AppU4DM

| Attribute | Type |
|-----------|------|
| _Id | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| EntityType | String |
| IsLocked | Boolean |
| ToBeCleaned | String |
| NId | String |

#### Entity: ToBeAndActualProducedMaterial_MTUAndContainer

**Published From:** AppU4DM

| Attribute | Type |
|-----------|------|
| WOOperationId | String |
| TOBeDMMtuID | String |
| ToBeProducedQuantity | String |
| DMMtuId | String |
| IsReserved | Boolean |
| ActiveNCNumber | String |
| CompletedQuantity | String |
| Equipment | String |
| NotWorkedQuantity | String |
| PartialWorkedQuantity | String |
| PausedQuantity | String |
| ReopenedQuantity | String |
| ScrappedQuantity | String |
| SentencedQuantity | String |
| SplitQuantity | String |
| UserId | String |
| ExecutionGroupPhaseId | String |
| ActulDMMtuID | String |
| DMMaterialId | String |
| MaterialTrackingUnit | String |
| MaterialNId | String |
| MaterialName | String |
| MaterialDescription | String |
| Code | String |
| CodeType | String |
| StatusNId | String |
| StateMachineNId | String |
| MaterialDefNId | String |
| materialRevision | String |
| materialQuantityValue | Integer/Decimal |
| UoMNid | String |
| ContainerType | String |
| IsReusable | Boolean |
| IsContainerActiveOnCurrentOperation | Boolean |
| IsContainerPauseOnCurrentOperation | Boolean |
| IsContainerCompleteOnCurrentOperation | Boolean |
| ResultStrategy | String |
| ResultValue | Integer/Decimal |
| SuggestedResultIsKo | String |
| SuggestedResultValue | Integer/Decimal |
| _Id | String |
| IsContainer | Boolean |
| MaterialId | String |
| IsResultKo | Boolean |
| IsRepeatOfKo | Boolean |
| CreatedOn | DateTime |

#### Entity: RM_CMX_SetPointHistoryValues_WorkOrderOperation

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| EquipmentNId | String |
| SetPointValue | Integer/Decimal |
| SetPointNId | String |
| SetPointName | String |
| _Id | String |
| WorkOrderOperationId | String |
| CreatedOn | DateTime |

#### Entity: NonConformance

**Published From:** AppU4DM

| Attribute | Type |
|-----------|------|
| _Id | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| EntityType | String |
| IsLocked | Boolean |
| ToBeCleaned | String |
| EndDate | DateTime |
| Equipment | String |
| MaterialNId | String |
| NId | String |
| NonConformanceLifecycle | String |
| Notes | String |
| Severity | String |
| StartDate | DateTime |
| Status | String |
| _Type | String |
| _Context | String |
| User | String |
| InspectedQuantity | String |
| MaterialDefinition | String |
| NonConformanceQuantity | String |
| SerialNumber | String |
| ToolDefinition | String |
| ParentNonConformanceNId | String |
| Quantity | String |
| WorkOrderOperation_Id | String |
| WorkOrder_Id | String |
| ChangeType_Id | String |

#### Entity: ChangePackage

**Published From:** AppU4DM

| Attribute | Type |
|-----------|------|
| _Id | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| EntityType | String |
| IsLocked | Boolean |
| ToBeCleaned | String |
| NId | String |
| WorkOrderOperationNId | String |
| OpenUserId | String |
| OpenTime | DateTime |
| CloseUserId | String |
| CloseTime | DateTime |
| Notes | String |
| OperationOccurrenceUId | String |
| Status_StateMachineNId | String |
| Status_StatusNId | String |
| WoOpPreviousStatus_StateMachineNId | String |
| WoOpPreviousStatus_StatusNId | String |

### Module: OpcenterEXDS_PartProgram

#### Entity: RM_PartProgramBottomUp

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| ParentId | String |
| DNCId | String |
| Name | String |
| DNCExternalId | String |
| Trialcut | String |
| Status | String |
| CanBePreviewed | String |
| CanBeTransferred | String |
| Version | String |
| _Type | String |
| Released | String |
| DM_MaterialId | String |
| MachineToDNCId | String |
| DNCMachine | String |
| EquipmentNId | String |
| StartingDNCId | String |
| _Id | String |
| DNCItemId | String |

#### Entity: RM_PartProgramTopDown

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| ParentId | String |
| DNCId | String |
| Name | String |
| DNCExternalId | String |
| Trialcut | String |
| Status | String |
| CanBePreviewed | String |
| CanBeTransferred | String |
| Version | String |
| _Type | String |
| Released | String |
| DM_MaterialId | String |
| MachineToDNCId | String |
| DNCMachine | String |
| EquipmentNId | String |
| StartingDNCId | String |
| _Id | String |
| DNCItemId | String |

#### Entity: RM_DNCHistory

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| Equipment | String |
| UserId | String |
| Status | String |
| Notes | String |
| DNCItem | String |
| StartingDate | DateTime |
| EndDate | DateTime |
| ReferencedWorkOrderHistoryId | String |
| WorkOrderName | String |
| WorkOrderOperationName | String |
| _Id | String |
| WorkOrderHistoryId | String |
| WorkOrderOperationId | String |
| DNCItemVersion | String |

#### Entity: RM_DNCTransferStatus

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| WorkOrderOperationId | String |
| NId | String |
| _Id | String |

#### Entity: RM_DNCItemTransferAction

**Published From:** RMComplexManuf

| Attribute | Type |
|-----------|------|
| Action | String |
| _Id | String |
| ActionId | String |

---

## 2. Microflow/Action Calls

Microflow and action calls found in the project (MicroflowCall, JavaAction, ExternalAction).

Found 82 call(s):

| Microflow | Module | Call Type | AppName | CommandName |
|-----------|--------|-----------|---------|-------------|
| CommandTransfer | OpcenterEXDS_PartProgram | MicroflowCall | 'AppU4DM'
 | 'DNCTransferDNCItems'
 |
| CommandPreview | OpcenterEXDS_PartProgram | MicroflowCall | 'AppU4DM'
 | 'DNCPreviewDNCItem' |
| ACT_CreateMTUProperty | OpcenterEXDS_OperatorLanding | MicroflowCall | 'Material'
 | 'CreateMaterialTrackingUnitProperties'
 |
| ACT_CreateMTUProperty | OpcenterEXDS_OperatorLanding | MicroflowCall | 'Material'
 | 'UpdateMaterialTrackingUnitProperties'
 |
| CommandWorkOrderHistory | OpcenterEXDS_PartProgram | MicroflowCall | 'AppU4DM'
 | 'CreateWorkOrderAndMaterialItemHistory'
 |
| SUB_ScrapConsumedMaterial_MF | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'AppU4DM'
 | 'ScrapConsumedMaterial' |
| SUB_UADMCreateSnagAndNoteList_MF | OpcenterEXDS_OperatorLanding | MicroflowCall | 'AppU4DM' | 'UADMCreateSnagAndNoteList'
 |
| SUB_UADMCheckCertificationMultiMachineOnOpenWOOperation_MF | OpcenterEXDS_OperatorLanding | MicroflowCall | 'AppU4DM' | 'UADMCheckCertificationMultiMachineOnOpenWOOperation'
 |
| SUB_PropagateSegregationTagsToDocument_MF | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'AppU4DM' | 'PropagateSegregationTagsToDocument' |
| SUB_CompleteWOStepSerialized_MF | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'AppU4DM' | 'CompleteWOStepSerialized' |
| SUB_AcquireToolFromAutomationNodeInstanceParameter_MF | OpcenterEXDS_OperatorLanding | MicroflowCall | 'AppU4DM' | 'AcquireToolFromAutomationNodeInstanceParameter' |
| SUB_LinkWorkInstructionsToWOOperation_MF | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'AppU4DM'
 | 'LinkWorkInstructionsToWOOperation'
 |
| SUB_StartWOStepFullQty_MF | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'AppU4DM' | 'StartWOStepFullQty' |
| SUB_AcquireMTUFromAutomationNodeInstanceParameter_MF | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'AppU4DM' | 'AcquireMTUFromAutomationNodeInstanceParameter' |
| SUB_LinkWorkInstructionsToWOStep_MF | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'AppU4DM'
 | 'LinkWorkInstructionsToWOStep'
 |
| SUB_TransmitEquipmentSetPointToAutomationNodeParameters | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'AppU4DM' | 'TransmitEquipmentSetPointToAutomationNodeParameters' |
| SUB_UADMCompleteWOOperationSerializedList_MF | OpcenterEXDS_OperatorLanding | MicroflowCall | 'AppU4DM' | 'UADMCompleteWOOperationSerializedList' |
| SUB_CompleteWOStepFullQty_MF | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'AppU4DM' | 'CompleteWOStepFullQty' |
| SUB_UADMCompleteWOOperationFullQtyMultiMachineList_MF | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'AppU4DM' | 'UADMCompleteWOOperationFullQtyMultiMachineList' |
| SUB_UADMPauseWorkOrderOperationMultiMachineList_MF | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'AppU4DM' | 'UADMPauseWorkOrderOperationMultiMachineList' |
| SUB_StartWOStepSerialized_MF | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'AppU4DM' | 'StartWOStepSerialized' |
| SUB_UADMStartWOOperationSerializedList_MF | OpcenterEXDS_OperatorLanding | MicroflowCall | 'AppU4DM' | 'UADMStartWOOperationSerializedList' |
| SUB_UADMConfirmSnagAndNoteList_MF | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'AppU4DM' | 'UADMConfirmSnagAndNoteList' |
| SUB_AcquireDCItemValueFromAutomationNodeInstanceParameter_MF | EXFN_WorkInstruction | MicroflowCall | 'AppU4DM' | 'AcquireDCItemValueFromAutomationNodeInstanceParameter' |
| SUB_UADMCreateDocument_MF | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'AppU4DM'
 | 'UADMCreateDocument'
 |
| UpdateSerialNumberList | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'AppU4DM' | 'UpdateSerialNumberList' |
| SUB_UADMCreateChangePackage | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'AppU4DM'
 | 'CreateChangePackage'
 |
| UADMUseToolList_MF | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'AppU4DM'
 | 'UseToolList'
 |
| SUB_UADMStartNonProductiveActivityList_MF | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'AppU4DM'
 | 'UADMStartNonProductiveActivityList' |
| UADMCreateAndAssignProducedMaterialItems_MF | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'AppU4DM'
 | 'UADMCreateAndAssignProducedMaterialItems' |
| AutoGenerateMTUCode | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'AppU4DM' | 'AutoGenerateMTUCode' |
| SUB_UADMCreateToBeUsedDocuments_MF | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'AppU4DM'
 | 'CreateToBeUsedDocuments'
 |
| SetTargetQuantityOnFlexibleWorkOrder_MF | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'AppU4DM'
 | 'SetTargetQuantityOnFlexibleWorkOrder'
 |
| SUB_DisassembleMaterialTypeDisassemble_MF | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'AppU4DM'
 | 'UADMManageToBeConsumedMaterialDisassemble'
 |
| CreateAndAssignProducedMaterialItems_MF | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'AppU4DM'
 | 'CreateAndAssignProducedMaterialItems'
 |
| SUB_UADMCompleteNonProductiveActivityList_MF | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'AppU4DM'
 | 'UADMCompleteNonProductiveActivityList'
 |
| SUB_UADMCreateNonConformanceV3_1_MF | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'AppU4DM'
 | 'UADMCreateNonConformanceV3_1'
 |
| SUB_UADMHoldOperation_MF | OpcenterEXDS_OperatorLanding | MicroflowCall | 'AppU4DM'
 | 'UADMSetWorkOrderHoldList'
 |
| SUB_DisassembleMaterialItem_MF | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'AppU4DM'
 | 'DisassembleMaterialItem'
 |
| SUB_ScrapWorkOrderSerialNumbers_MF | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'AppU4DM'
 | 'ScrapWorkOrderSerialNumbers' |
| SUB_ValidateBarcodeForMaterialConsumption_MF | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'MaterialTraceability'
 | 'ValidateBarcodeForMaterialConsumption'
 |
| AssignProducedMaterialItems_MF | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'AppU4DM'
 | 'AssignProducedMaterialItems'
 |
| SUB_CoProduceMaterialTrackingUnitList_MF | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'AppU4DM'
 | 'CoProduceMaterialTrackingUnit'
 |
| SUB_UADMConsumeMaterialItemList_MF | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'AppU4DM'
 | 'UADMConsumeMaterialItemList'
 |
| DeactivateMaterialTrackingUnitOnScrewingTool | OpcenterEXDS_OperatorLanding_Connector | ExternalAction | AppU4DM | DeactivateMaterialTrackingUnitOnScrewingTool |
| UADMSendBuyOffNotification_MF | OpcenterEXDS_OperatorLanding | MicroflowCall | 'AppU4DM' | 'SendBuyOffNotification'
 |
| ActivateMaterialTrackingUnitOnScrewingTool | OpcenterEXDS_OperatorLanding_Connector | ExternalAction | AppU4DM | ActivateMaterialTrackingUnitOnScrewingTool |
| UADMAbruptlyCloseFlexibleWorkOrder | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'AppU4DM' | 'UADMAbruptlyCloseFlexibleWorkOrder' |
| UADMStartOperation_MF | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'AppU4DM' | 'UADMStartOperation' |
| AutoGenerateWorkOrderNId | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'AppU4DM' | 'AutoGenerateWorkOrderNId' |
| TriggerPrintingOnWorkOrderOperation | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'AppU4DM'
 | 'TriggerPrintingOnWorkOrderOperation'
 |
| ACT_CallSkipFullQty | OpcenterEXDS_OperatorLanding | ExternalAction | AppU4DM | UADMSkipWOOperationFullQty |
| ACT_CallSkipWOOperationSerialized | OpcenterEXDS_OperatorLanding | ExternalAction | AppU4DM | UADMSkipWOOperationSerialized |
| ACT_EditPrekit_MF | OpcenterEXDS_OperatorLanding_Connector | MicroflowCall | 'AppU4DM'
 | 'ReserveMaterialItems'
 |
| LinkDocumentSet | EXFN_DocumentViewer | MicroflowCall | 'Document' | 'LinkDocumentSet'
 |
| UnlinkDocumentSet | EXFN_DocumentViewer | MicroflowCall | 'Document' | 'UnlinkDocumentSet'
 |
| CreateLinkDocument | EXFN_DocumentViewer | MicroflowCall | 'Document' | 'CreateLinkDocument' |
| CreateDocument | EXFN_DocumentViewer | MicroflowCall | 'Document' | 'CreateDocument'
 |
| UnlinkDocument | EXFN_DocumentViewer | MicroflowCall | 'Document' | 'UnlinkDocument' |
| LinkDocument | EXFN_DocumentViewer | MicroflowCall | 'Document' | 'LinkDocument' |
| SUB_Signature_Prepare | EXFN_ElectronicSignature | MicroflowCall | 'AuditTrail' | 'PrepareSignature' |
| SUB_OnSign | EXFN_ElectronicSignature | MicroflowCall | 'AuditTrail' | 'AddSignature' |
| SUB_Signature_Abort | EXFN_ElectronicSignature | MicroflowCall | 'AuditTrail' | 'AbortSignature' |
| SUB_Context_Create | EXFN_Quality | MicroflowCall | 'WorkInstruction' | 'CreateInspectionAcquisitionContext' |
| ACT_ConfirmInspectionSample | EXFN_Quality | MicroflowCall | 'WorkInstruction' | 'ConfirmInspectionSample' |
| SUB_Definition_AssociateFailure | EXFN_Quality | MicroflowCall | 'WorkInstruction' | 'AssociateFailureToInspectionValue' |
| SUB_Definition_DisassociateFailure | EXFN_Quality | MicroflowCall | 'WorkInstruction' | 'DisassociateFailureFromInspectionValue' |
| SUB_ImageGrid_NewDot_Create | EXFN_Quality | MicroflowCall | 'WorkInstruction' | $HelperExport/CommandName |
| ACT_CreateNewVisualSample | EXFN_Quality | MicroflowCall | 'WorkInstruction' | 'CreateVisualDetectedFailurewithES' |
| ACT_VisualInspection_Confirm | EXFN_Quality | MicroflowCall | 'WorkInstruction' | 'CreateVisualDetectedFailurewithES' |
| SUB_Variable_Update_ES | EXFN_Quality | MicroflowCall | 'WorkInstruction' | 'UpdateInspectionValuewithES' |
| SUB_Variable_Create_ES | EXFN_Quality | MicroflowCall | 'WorkInstruction' | 'CreateInspectionValuewithES' |
| SUB_Attributive_Create_ES | EXFN_Quality | MicroflowCall | 'WorkInstruction' | 'CreateInspectionValuewithES' |
| SUB_Attributive_Update_ES | EXFN_Quality | MicroflowCall | 'WorkInstruction' | 'UpdateInspectionValuewithES' |
| SUB_WorkInstructionStepItem_AutoSave | EXFN_WorkInstruction | MicroflowCall | @EXFN_WorkInstruction.AppName | 'AutoSaveWorkInstructionStepItemValue' |
| SUB_CalculateWorkInstructionFormulaValues | EXFN_WorkInstruction | MicroflowCall | 'WorkInstruction' | 'CalculateWorkInstructionFormulaValues' |
| SUB_WorkInstruction_Delete | EXFN_WorkInstruction | MicroflowCall | @EXFN_WorkInstruction.AppName | 'DeleteWorkInstruction' |
| SUB_WorkInstruction_Create | EXFN_WorkInstruction | MicroflowCall | @EXFN_WorkInstruction.AppName | 'CreateWorkInstruction' |
| SUB_WIStep_Confirm | EXFN_WorkInstruction | MicroflowCall | @EXFN_WorkInstruction.AppName | 'ConfirmWIStep' |
| SUB_WIStep_Acknowledge | EXFN_WorkInstruction | MicroflowCall | @EXFN_WorkInstruction.AppName | 'AcknowledgeWIStep' |
| SUB_WIStep_ReEdit | EXFN_WorkInstruction | MicroflowCall | @EXFN_WorkInstruction.AppName | 'ReEditWorkInstructionStep' |
| SUB_WorkInstructionStatus_InEditing | EXFN_WorkInstruction | MicroflowCall | @EXFN_WorkInstruction.AppName | 'InEditingWorkInstruction' |

---

## 3. Signal Manager Subscriptions

Signal subscriptions from Signal Manager widgets (siemens.mxtosignal.MxToSignal).

Found 64 subscription(s):

| Module | Document Type | Document | Signal Name | App Name | Subscription Filter |
|--------|---------------|----------|-------------|----------|---------------------|
| OpcenterEXDS_PartProgram | Page | PartProgramDetailsPopUp | DNCStartTransferEvent | AppU4DM | Yes |
| OpcenterEXDS_PartProgram | Page | PartProgramDetailsPopUp | DNCCompleteTransferEvent | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnFAICompleted | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | StatusChanged | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnChangeFlexibleWOOpStatusToComplete | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnCreateNonConformanceV3_1 | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnCompleteReworkOrder | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnStartWorkOrderOperationSerialized | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | SnagAndNoteNotificationSgn | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | WorkOrderSerialNumbersScrapped | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnFAICandidateDeclared | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnSkipWOOperation | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnCompleteWorkOrderOperationFullQty | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnCompleteWorkOrderOperationSerialized | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnSetHoldWorkOrder | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnStartWorkOrderOperationFullQty | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnPauseWorkOrderOperation | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnSentenceNonConformanceV3_1 | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnReopenWorkOrderOperation | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnPauseWorkOrderOperation | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnCompleteWorkOrderStepSerialized | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnSetHoldWorkOrder | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnAssemblyWorkOrderOperation | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | WorkInstructionStatusChangedSignal | WorkInstruction | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnReopenWorkOrderOperation | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnStartWorkOrderOperationSerialized | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnStartWorkOrderStepFullQty | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnDisassemblyWorkOrderOperation | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnUsedToolSignal | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | WorkOrderSerialNumbersScrapped | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnCreateNonConformanceV3_1 | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnFAIRemoved | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnFAICandidateDeclared | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnStartWorkOrderStepSerialized | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnCompleteWorkOrderStepFullQty | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnFAICompleted | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnSkipWOOperation | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnCompleteWorkOrderOperationSerialized | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnSentenceNonConformanceV3_1 | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | SnagAndNoteNotificationSgn | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | StatusChanged | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnCompleteReworkOrder | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnStartWorkOrderOperationFullQty | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnCompleteWorkOrderOperationFullQty | AppU4DM | Yes |
| OpcenterEXDS_EXFN_WorkInstruction | Snippet | SNP_WorkInstruction_VerticalView | WorkInstructionSectionCompletedSignal | WorkInstruction | Yes |
| OpcenterEXDS_EXFN_WorkInstruction | Snippet | SNP_WorkInstruction_VerticalView | WorkInstructionStepCompletedSignal | WorkInstruction | Yes |
| OpcenterEXDS_OperatorLanding | Snippet | WorkInstructions | OnLinkWIOnDemandToSerialNumber | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Snippet | ToBeUsedTools_Backup | OnUsedToolSignal | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Snippet | ToBeCoByProducedMaterials | CoProductProduced | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Snippet | ToBeCoByProducedMaterials | ByProductProduced | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Snippet | ToBeCoByProducedMaterials | OutputMaterialProduced | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Snippet | PartProgram | DNCStartTransferEvent | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Snippet | PartProgram | DNCCompleteTransferEvent | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Snippet | ToBeProducedMaterial | OnAssemblyWorkOrderOperation | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Snippet | ToBeProducedMaterial | OnDisassemblyWorkOrderOperation | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Snippet | ToBeProducedMaterial | ScrapMaterials | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Snippet | ToBeUsedTools | OnUsedToolSignal | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Snippet | ToBeUsedTools | MaterialTrackingUnitDeactivationEvent | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Snippet | ToBeUsedTools | MaterialTrackingUnitActivationEvent | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Snippet | ToBeUsedTools | ScrewingExecuted | AppU4DM | Yes |
| EXFN_Quality | Snippet | SNP_QualityInspectionContainer | DelayedExecution | WorkInstruction | Yes |
| EXFN_Quality | Snippet | SNP_QualityInspectionContainer | InspectionExecutionChrReprRuntimeNumberChanged | WorkInstruction | Yes |
| EXFN_Quality | Snippet | SNP_QualityInspectionContainer | OnCompleteInspectionSampleScenarioInstance | WorkInstruction | Yes |
| EXFN_Quality | Snippet | SNP_QualityInspectionContainer | InspectionSampleConfirmed | WorkInstruction | Yes |

---

## 4. Navigation Items

| Parent Node | Node | Target Page | User Roles |
|-------------|------|-------------|------------|
| - | Home | OpcenterEXDS_OperatorLanding.OperatorLanding | - |
| - | Configuration | - | - |
| - | Accounts | Administration.Account_Overview | Administrator |

---

## 5. Navigation Page Commands

Command bar actions extracted from navigation pages. Shows buttons in the vertical command bar of the Right placeholder.

Found commands in 1 page(s):

### OpcenterEXDS_OperatorLanding.OperatorLanding

| Caption | Target Page | Target AppName | Target CommandName |
|---------|-------------|----------------|--------------------|
| Open
 | OpcenterEXDS_OperatorLanding.OperatorTerminal, OpcenterEXDS_OperatorLanding.PANEL_OperatorLandingGoToDetailsSN | - | - |
| Start
 | OpcenterEXDS_OperatorLanding.PANEL_OperatorLandingStart, OpcenterEXDS_OperatorLanding.OperatorTerminal, OpcenterEXDS_OperatorLanding.PANEL_OperatorLandingGoToDetailsSN, OpcenterEXDS_OperatorLanding.PANEL_Start_ValidateUser, OpcenterEXDS_OperatorLanding.PANEL_Start_ValidateUser_FromDetails | - | - |
| Pause
 | OpcenterEXDS_OperatorLanding.PANEL_OperatorLandingPause, OpcenterEXDS_OperatorLanding.PANEL_OperatorLandingPause, OpcenterEXDS_OperatorLanding.PANEL_Pause_ValidateUser | - | - |
| Hold
 | OpcenterEXDS_OperatorLanding.PANEL_OperatorLandingHold | - | - |
| Skip | OpcenterEXDS_OperatorLanding.PANEL_SkipWOOperation | - | - |
| Complete | OpcenterEXDS_OperatorLanding.PANEL_OperatorLandingComplete, OpcenterEXDS_OperatorLanding.PANEL_Complete_ValidateUser | - | - |
| Documents
 | - | - | - |
| NC | OpcenterEXDS_OperatorLanding.PANEL_VerticalCommandBarNC, OpcenterEXDS_OperatorLanding.PANEL_VerticalCommandBarNC_Details | - | - |
| Change Package | OpcenterEXDS_OperatorLanding.PANEL_ChangePackage | - | - |
| Add Documents | OpcenterEXDS_OperatorLanding.AddDocumentsPopup | - | - |
| Notes
 | OpcenterEXDS_OperatorLanding.PANEL_Notes | - | - |
| Scrap | OpcenterEXDS_OperatorLanding.ScrapProducedMaterialPopup | - | - |
| More
 | OpcenterEXDS_OperatorLanding.PANEL_VerticalCommandBarMore, OpcenterEXDS_OperatorLanding.PANEL_VerticalCommandBarMore_Details | - | - |

---

## 6. Pages/Panels Commands Hierarchy

Microflows and nanoflows called by each page/panel, showing recursive call hierarchy up to 5 levels (in YAML structure). Microflows called transitively are loaded on-the-fly from the database when needed.

```yaml
pages:
  - name: PartProgramHistoryPopUp
    module: OpcenterEXDS_PartProgram
    flows:
      []
    target_commands:
      []

  - name: PartProgramDetailsPopUp
    module: OpcenterEXDS_PartProgram
    flows:
      - name: OpcenterEXDS_PartProgram.EVT_DNCCompleteTransfer
        calls:
          - name: OpcenterEXDS_PartProgram.ACT_CreateMTUProperty
    target_commands:
      - Material.CreateMaterialTrackingUnitProperties
      - Material.UpdateMaterialTrackingUnitProperties

  - name: PANEL_SelectEquipmentAcquireWI
    module: OpcenterEXDS_EXFN_WorkInstruction
    flows:
      - name: OpcenterEXDS_EXFN_WorkInstruction.ACT_EquipmentSelectionPanel_AcquireDC
        calls:
          - name: OpcenterEXDS_OperatorLanding_Connector.SUB_AcquireDCItemValueFromAutomationNodeInstanceParameter_MF
          - name: OpcenterEXDS_EXFN_WorkInstruction.SUB_UpdateStepItemValue
            - name: OpcenterEXDS_EXFN_WorkInstruction.SUB_WorkInstruction_GetById
              - name: OpcenterEXDS_EXFN_WorkInstruction.SUB_WorkInstruction_Details_Load
                - name: EXFN_WorkInstruction.ACT_WorkInstructionStatus_InEditing
                  - name: EXFN_WorkInstruction.SUB_WorkInstructionStatus_InEditing
    target_commands:
      - AppU4DM.AcquireDCItemValueFromAutomationNodeInstanceParameter
      - @EXFN_WorkInstruction.AppName.InEditingWorkInstruction

  - name: ScrapMaterialConsumptionPopup
    module: OpcenterEXDS_OperatorLanding
    flows:
      - name: OpcenterEXDS_OperatorLanding.ACT_ScrapConsumedMaterial
        calls:
          - name: OpcenterEXDS_OperatorLanding_Connector.ScrapConsumedMaterial
            - name: OpcenterEXDS_OperatorLanding_Connector.SUB_ScrapConsumedMaterial_MF
    target_commands:
      - AppU4DM.ScrapConsumedMaterial

  - name: AddDocumentsPopup
    module: OpcenterEXDS_OperatorLanding
    flows:
      - name: OpcenterEXDS_OperatorLanding.ACT_Submit_AddDocuments
        calls:
          - name: OpcenterEXDS_OperatorLanding_Connector.UADMCreateToBeUsedDocuments
            - name: OpcenterEXDS_OperatorLanding_Connector.SUB_UADMCreateToBeUsedDocuments_MF
    target_commands:
      - AppU4DM.CreateToBeUsedDocuments

  - name: OperatorLanding
    module: OpcenterEXDS_OperatorLanding
    flows:
      - name: OpcenterEXDS_OperatorLanding.ACT_WorkOrderOperation_Tile_GoToDetails_Click
        calls:
          - name: OpcenterEXDS_OperatorLanding.ACT_GoToDetails_Click
            - name: OpcenterEXDS_OperatorLanding_Connector.UADMCheckCertificationMultiMachineOnOpenWOOperation
              - name: OpcenterEXDS_OperatorLanding_Connector.SUB_UADMCheckCertificationMultiMachineOnOpenWOOperation_MF
    target_commands:
      - AppU4DM.UADMCheckCertificationMultiMachineOnOpenWOOperation

  - name: NonConformancePopup
    module: OpcenterEXDS_OperatorLanding
    flows:
      - name: OpcenterEXDS_OperatorLanding.ACT_Submit_NonConformance
        calls:
          - name: OpcenterEXDS_OperatorLanding_Connector.UADMCreateNonConformanceV3_1
            - name: OpcenterEXDS_OperatorLanding_Connector.SUB_UADMCreateNonConformanceV3_1_MF
    target_commands:
      - AppU4DM.UADMCreateNonConformanceV3_1

  - name: OperatorTerminal
    module: OpcenterEXDS_OperatorLanding
    flows:
      []
    target_commands:
      []

  - name: ScrapProducedMaterialPopup
    module: OpcenterEXDS_OperatorLanding
    flows:
      - name: OpcenterEXDS_OperatorLanding.ACT_ScrapWorkOrderSerialNumbers
        calls:
          - name: OpcenterEXDS_OperatorLanding_Connector.ScrapWorkOrderSerialNumbers
            - name: OpcenterEXDS_OperatorLanding_Connector.SUB_ScrapWorkOrderSerialNumbers_MF
    target_commands:
      - AppU4DM.ScrapWorkOrderSerialNumbers

  - name: PANEL_DisassembleMaterialTrackingUnit
    module: OpcenterEXDS_OperatorLanding
    flows:
      - name: OpcenterEXDS_OperatorLanding.ACT_DisassembleMaterialWithNotes
        calls:
          - name: OpcenterEXDS_OperatorLanding_Connector.DisassembleMaterialItem
            - name: OpcenterEXDS_OperatorLanding_Connector.SUB_DisassembleMaterialItem_MF
    target_commands:
      - AppU4DM.DisassembleMaterialItem

  - name: PANEL_ImportDocument
    module: OpcenterEXDS_OperatorLanding
    flows:
      - name: OpcenterEXDS_OperatorLanding.ACT_CreateAndLinkDocument
        calls:
          - name: OpcenterEXDS_OperatorLanding_Connector.UADMCreateDocument
            - name: OpcenterEXDS_OperatorLanding_Connector.SUB_UADMCreateDocument_MF
          - name: OpcenterEXDS_OperatorLanding_Connector.SUB_PropagateSegregationTagsToDocument_MF
    target_commands:
      - AppU4DM.UADMCreateDocument
      - AppU4DM.PropagateSegregationTagsToDocument

  - name: PANEL_LinkDocuments
    module: OpcenterEXDS_OperatorLanding
    flows:
      []
    target_commands:
      []

  - name: PANEL_SelectToBeConsumedMTU
    module: OpcenterEXDS_OperatorLanding
    flows:
      []
    target_commands:
      []

  - name: PANEL_SelectDestinationContainer
    module: OpcenterEXDS_OperatorLanding
    flows:
      []
    target_commands:
      []

  - name: PANEL_AddDocumentsImportDocument
    module: OpcenterEXDS_OperatorLanding
    flows:
      - name: OpcenterEXDS_OperatorLanding.ACT_AddDocumentsCreateAndLinkDocument
        calls:
          - name: OpcenterEXDS_OperatorLanding_Connector.UADMCreateDocument
            - name: OpcenterEXDS_OperatorLanding_Connector.SUB_UADMCreateDocument_MF
          - name: OpcenterEXDS_OperatorLanding_Connector.SUB_PropagateSegregationTagsToDocument_MF
    target_commands:
      - AppU4DM.UADMCreateDocument
      - AppU4DM.PropagateSegregationTagsToDocument

  - name: PANEL_ChangePackage
    module: OpcenterEXDS_OperatorLanding
    flows:
      - name: OpcenterEXDS_OperatorLanding.ACT_Submit_ChangePackage
        calls:
          - name: OpcenterEXDS_OperatorLanding_Connector.UADMCreateChangePackage
            - name: OpcenterEXDS_OperatorLanding_Connector.SUB_UADMCreateChangePackage
    target_commands:
      - AppU4DM.CreateChangePackage

  - name: PANEL_SkipWOOperation
    module: OpcenterEXDS_OperatorLanding
    flows:
      - name: OpcenterEXDS_OperatorLanding.ACT_CallSkipOperation
        calls:
          - name: OpcenterEXDS_OperatorLanding.ACT_CallSkipFullQty
          - name: OpcenterEXDS_OperatorLanding.ACT_CallSkipWOOperationSerialized
    target_commands:
      - AppU4DM.UADMSkipWOOperationFullQty
      - AppU4DM.UADMSkipWOOperationSerialized

  - name: PANEL_Notes
    module: OpcenterEXDS_OperatorLanding
    flows:
      - name: OpcenterEXDS_OperatorLanding.ACT_Create_Panel_Note
        calls:
          - name: OpcenterEXDS_OperatorLanding_Connector.UADMCreateSnagAndNoteList
            - name: OpcenterEXDS_OperatorLanding_Connector.SUB_UADMCreateSnagAndNoteList_MF
      - name: OpcenterEXDS_OperatorLanding.ACT_Acknowledge_Note
        calls:
          - name: OpcenterEXDS_OperatorLanding_Connector.UADMConfirmSnagAndNoteList
            - name: OpcenterEXDS_OperatorLanding_Connector.SUB_UADMConfirmSnagAndNoteList_MF
    target_commands:
      - AppU4DM.UADMCreateSnagAndNoteList
      - AppU4DM.UADMConfirmSnagAndNoteList

  - name: PANEL_VerticalCommandBarMore
    module: OpcenterEXDS_OperatorLanding
    flows:
      []
    target_commands:
      []

  - name: PANEL_OperatorDetailsCompleteStep
    module: OpcenterEXDS_OperatorLanding
    flows:
      - name: OpcenterEXDS_OperatorLanding.ACT_Complete_Panel_Step
        calls:
          - name: OpcenterEXDS_OperatorLanding.SUB_Call_Complete_Step
            - name: OpcenterEXDS_OperatorLanding_Connector.CompleteWOStepSerialized
              - name: OpcenterEXDS_OperatorLanding_Connector.SUB_CompleteWOStepSerialized_MF
            - name: OpcenterEXDS_OperatorLanding_Connector.CompleteWOStepFullQty
              - name: OpcenterEXDS_OperatorLanding_Connector.SUB_CompleteWOStepFullQty_MF
    target_commands:
      - AppU4DM.CompleteWOStepSerialized
      - AppU4DM.CompleteWOStepFullQty

  - name: PANEL_SetPoint
    module: OpcenterEXDS_OperatorLanding
    flows:
      - name: OpcenterEXDS_OperatorLanding.ACT_TransmitEquipmentSetPointToAutomationNodeParameters
        calls:
          - name: OpcenterEXDS_OperatorLanding_Connector.TransmitEquipmentSetPointToAutomationNodeParameters
            - name: OpcenterEXDS_OperatorLanding_Connector.SUB_TransmitEquipmentSetPointToAutomationNodeParameters
      - name: OpcenterEXDS_OperatorLanding.ACT_TransmitAllEquipmentSetPointToAutomationNodeParameters
        calls:
          - name: OpcenterEXDS_OperatorLanding_Connector.TransmitEquipmentSetPointToAutomationNodeParameters
            - name: OpcenterEXDS_OperatorLanding_Connector.SUB_TransmitEquipmentSetPointToAutomationNodeParameters
    target_commands:
      - AppU4DM.TransmitEquipmentSetPointToAutomationNodeParameters

  - name: PANEL_VerticalCommandBarMore_Details
    module: OpcenterEXDS_OperatorLanding
    flows:
      []
    target_commands:
      []

  - name: PANEL_ActiveUserList
    module: OpcenterEXDS_OperatorLanding
    flows:
      []
    target_commands:
      []

  - name: PANEL_OperatorLandingStart
    module: OpcenterEXDS_OperatorLanding
    flows:
      - name: OpcenterEXDS_OperatorLanding.ACT_StartPanel_Operation
        calls:
          - name: OpcenterEXDS_OperatorLanding.SUB_Start_Operation
            - name: OpcenterEXDS_OperatorLanding_Connector.UADMStartWOOperationSerializedList
              - name: OpcenterEXDS_OperatorLanding_Connector.SUB_UADMStartWOOperationSerializedList_MF
            - name: OpcenterEXDS_OperatorLanding_Connector.UADMStartOperation
              - name: OpcenterEXDS_OperatorLanding_Connector.UADMStartOperation_MF
          - name: OpcenterEXDS_OperatorLanding.SUB_GoToDetails_If_AutomaticRedirect_Is_True
            - name: OpcenterEXDS_OperatorLanding.ACT_GoToDetails_Click
              - name: OpcenterEXDS_OperatorLanding_Connector.UADMCheckCertificationMultiMachineOnOpenWOOperation
                - name: OpcenterEXDS_OperatorLanding_Connector.SUB_UADMCheckCertificationMultiMachineOnOpenWOOperation_MF
      - name: OpcenterEXDS_OperatorLanding.ACT_StartAllSNButton
        calls:
          - name: OpcenterEXDS_OperatorLanding.ACT_StartPanel_Operation
            - name: OpcenterEXDS_OperatorLanding.SUB_Start_Operation
              - name: OpcenterEXDS_OperatorLanding_Connector.UADMStartWOOperationSerializedList
                - name: OpcenterEXDS_OperatorLanding_Connector.SUB_UADMStartWOOperationSerializedList_MF
              - name: OpcenterEXDS_OperatorLanding_Connector.UADMStartOperation
                - name: OpcenterEXDS_OperatorLanding_Connector.UADMStartOperation_MF
            - name: OpcenterEXDS_OperatorLanding.SUB_GoToDetails_If_AutomaticRedirect_Is_True
              - name: OpcenterEXDS_OperatorLanding.ACT_GoToDetails_Click
                - name: OpcenterEXDS_OperatorLanding_Connector.UADMCheckCertificationMultiMachineOnOpenWOOperation
                  - name: OpcenterEXDS_OperatorLanding_Connector.SUB_UADMCheckCertificationMultiMachineOnOpenWOOperation_MF
      - name: OpcenterEXDS_OperatorLanding.ACT_AssociateNewSerialNumber
        calls:
          - name: OpcenterEXDS_OperatorLanding_Connector.AssignProducedMaterialItems_MF
      - name: OpcenterEXDS_OperatorLanding.ACT_AddNewSerialNumbersFromNId
        calls:
          - name: OpcenterEXDS_OperatorLanding_Connector.CreateAndAssignProducedMaterialItems_MF
      - name: OpcenterEXDS_OperatorLanding.ACT_AddNewSerialNumbersFromNumber
        calls:
          - name: OpcenterEXDS_OperatorLanding_Connector.UADMCreateAndAssignProducedMaterialItems_MF
    target_commands:
      - AppU4DM.UADMStartWOOperationSerializedList
      - AppU4DM.UADMStartOperation
      - AppU4DM.UADMCheckCertificationMultiMachineOnOpenWOOperation
      - AppU4DM.AssignProducedMaterialItems
      - AppU4DM.CreateAndAssignProducedMaterialItems
      - AppU4DM.UADMCreateAndAssignProducedMaterialItems

  - name: PANEL_OperatorLandingGoToDetailsSN
    module: OpcenterEXDS_OperatorLanding
    flows:
      []
    target_commands:
      []

  - name: PANEL_OperatorLandingHold
    module: OpcenterEXDS_OperatorLanding
    flows:
      - name: OpcenterEXDS_OperatorLanding.ACT_Hold_Panel_Operation
        calls:
          - name: OpcenterEXDS_OperatorLanding_Connector.UADMHoldOperation
            - name: OpcenterEXDS_OperatorLanding_Connector.SUB_UADMHoldOperation_MF
    target_commands:
      - AppU4DM.UADMSetWorkOrderHoldList

  - name: PANEL_CloseFlexibleWorkOrder
    module: OpcenterEXDS_OperatorLanding
    flows:
      - name: OpcenterEXDS_OperatorLanding.ACT_UADMAbruptlyCloseFlexibleWorkOrder
        calls:
          - name: OpcenterEXDS_OperatorLanding_Connector.UADMAbruptlyCloseFlexibleWorkOrder
      - name: OpcenterEXDS_OperatorLanding.ACT_GenerateWONId
        calls:
          - name: OpcenterEXDS_OperatorLanding.ACT_AutoGenerateWorkOrderNId
            - name: OpcenterEXDS_OperatorLanding_Connector.AutoGenerateWorkOrderNId
    target_commands:
      - AppU4DM.UADMAbruptlyCloseFlexibleWorkOrder
      - AppU4DM.AutoGenerateWorkOrderNId

  - name: PANEL_OperatorLandingComplete
    module: OpcenterEXDS_OperatorLanding
    flows:
      - name: OpcenterEXDS_OperatorLanding.ACT_Complete_Panel_Operation
        calls:
          - name: OpcenterEXDS_OperatorLanding.SUB_Call_Complete_Operation
            - name: OpcenterEXDS_OperatorLanding_Connector.UADMCompleteWOOperationSerializedList
              - name: OpcenterEXDS_OperatorLanding_Connector.SUB_UADMCompleteWOOperationSerializedList_MF
            - name: OpcenterEXDS_OperatorLanding_Connector.UADMCompleteWOOperationFullQtyMultiMachineList
              - name: OpcenterEXDS_OperatorLanding_Connector.SUB_UADMCompleteWOOperationFullQtyMultiMachineList_MF
      - name: OpcenterEXDS_OperatorLanding.ACT_CompleteAllSnButton
        calls:
          - name: OpcenterEXDS_OperatorLanding.ACT_Complete_Panel_Operation
            - name: OpcenterEXDS_OperatorLanding.SUB_Call_Complete_Operation
              - name: OpcenterEXDS_OperatorLanding_Connector.UADMCompleteWOOperationSerializedList
                - name: OpcenterEXDS_OperatorLanding_Connector.SUB_UADMCompleteWOOperationSerializedList_MF
              - name: OpcenterEXDS_OperatorLanding_Connector.UADMCompleteWOOperationFullQtyMultiMachineList
                - name: OpcenterEXDS_OperatorLanding_Connector.SUB_UADMCompleteWOOperationFullQtyMultiMachineList_MF
    target_commands:
      - AppU4DM.UADMCompleteWOOperationSerializedList
      - AppU4DM.UADMCompleteWOOperationFullQtyMultiMachineList

  - name: PANEL_SetPointHistory
    module: OpcenterEXDS_OperatorLanding
    flows:
      []
    target_commands:
      []

  - name: PANEL_OperatorDetailsStartStep
    module: OpcenterEXDS_OperatorLanding
    flows:
      - name: OpcenterEXDS_OperatorLanding.ACT_StartPanel_Step
        calls:
          - name: OpcenterEXDS_OperatorLanding.SUB_Start_Step
            - name: OpcenterEXDS_OperatorLanding_Connector.StartWOStepSerialized
              - name: OpcenterEXDS_OperatorLanding_Connector.SUB_StartWOStepSerialized_MF
            - name: OpcenterEXDS_OperatorLanding_Connector.StartWOStepFullQty
              - name: OpcenterEXDS_OperatorLanding_Connector.SUB_StartWOStepFullQty_MF
    target_commands:
      - AppU4DM.StartWOStepSerialized
      - AppU4DM.StartWOStepFullQty

  - name: PANEL_OperatorLandingPause
    module: OpcenterEXDS_OperatorLanding
    flows:
      - name: OpcenterEXDS_OperatorLanding.ACT_Pause_Panel_Operation
        calls:
          - name: OpcenterEXDS_OperatorLanding_Connector.UADMPauseWorkOrderOperationMultiMachineList
            - name: OpcenterEXDS_OperatorLanding_Connector.SUB_UADMPauseWorkOrderOperationMultiMachineList_MF
    target_commands:
      - AppU4DM.UADMPauseWorkOrderOperationMultiMachineList

  - name: PANEL_ScrewingDetails
    module: OpcenterEXDS_OperatorLanding
    flows:
      []
    target_commands:
      []

  - name: PANEL_ChangeSN
    module: OpcenterEXDS_OperatorLanding
    flows:
      - name: OpcenterEXDS_OperatorLanding.ACT_UpdateSerialNumberList
        calls:
          - name: OpcenterEXDS_OperatorLanding_Connector.UpdateSerialNumberList
      - name: OpcenterEXDS_OperatorLanding.ACT_GenerateNewSerialNumber
        calls:
          - name: OpcenterEXDS_OperatorLanding.ACT_AutoGenerateMTUCode
            - name: OpcenterEXDS_OperatorLanding_Connector.AutoGenerateMTUCode
    target_commands:
      - AppU4DM.UpdateSerialNumberList
      - AppU4DM.AutoGenerateMTUCode

  - name: PANEL_Start_ValidateUser_FromDetails
    module: OpcenterEXDS_OperatorLanding
    flows:
      - name: OpcenterEXDS_OperatorLanding.ACT_Start_ValidateUser_FromDetails
        calls:
          - name: OpcenterEXDS_OperatorLanding.SUB_Start_Operation_AnyCase
            - name: OpcenterEXDS_OperatorLanding.SUB_Start_Operation
              - name: OpcenterEXDS_OperatorLanding_Connector.UADMStartWOOperationSerializedList
                - name: OpcenterEXDS_OperatorLanding_Connector.SUB_UADMStartWOOperationSerializedList_MF
              - name: OpcenterEXDS_OperatorLanding_Connector.UADMStartOperation
                - name: OpcenterEXDS_OperatorLanding_Connector.UADMStartOperation_MF
            - name: OpcenterEXDS_OperatorLanding.SUB_GoToDetails_If_AutomaticRedirect_Is_True
              - name: OpcenterEXDS_OperatorLanding.ACT_GoToDetails_Click
                - name: OpcenterEXDS_OperatorLanding_Connector.UADMCheckCertificationMultiMachineOnOpenWOOperation
                  - name: OpcenterEXDS_OperatorLanding_Connector.SUB_UADMCheckCertificationMultiMachineOnOpenWOOperation_MF
    target_commands:
      - AppU4DM.UADMStartWOOperationSerializedList
      - AppU4DM.UADMStartOperation
      - AppU4DM.UADMCheckCertificationMultiMachineOnOpenWOOperation

  - name: PANEL_Start_ValidateUser
    module: OpcenterEXDS_OperatorLanding
    flows:
      - name: OpcenterEXDS_OperatorLanding.ACT_Start_ValidateUser
        calls:
          - name: OpcenterEXDS_OperatorLanding.SUB_Start_Operation_AnyCase
            - name: OpcenterEXDS_OperatorLanding.SUB_Start_Operation
              - name: OpcenterEXDS_OperatorLanding_Connector.UADMStartWOOperationSerializedList
                - name: OpcenterEXDS_OperatorLanding_Connector.SUB_UADMStartWOOperationSerializedList_MF
              - name: OpcenterEXDS_OperatorLanding_Connector.UADMStartOperation
                - name: OpcenterEXDS_OperatorLanding_Connector.UADMStartOperation_MF
            - name: OpcenterEXDS_OperatorLanding.SUB_GoToDetails_If_AutomaticRedirect_Is_True
              - name: OpcenterEXDS_OperatorLanding.ACT_GoToDetails_Click
                - name: OpcenterEXDS_OperatorLanding_Connector.UADMCheckCertificationMultiMachineOnOpenWOOperation
                  - name: OpcenterEXDS_OperatorLanding_Connector.SUB_UADMCheckCertificationMultiMachineOnOpenWOOperation_MF
    target_commands:
      - AppU4DM.UADMStartWOOperationSerializedList
      - AppU4DM.UADMStartOperation
      - AppU4DM.UADMCheckCertificationMultiMachineOnOpenWOOperation

  - name: PANEL_Pause_ValidateUser
    module: OpcenterEXDS_OperatorLanding
    flows:
      []
    target_commands:
      []

  - name: PANEL_Complete_ValidateUser
    module: OpcenterEXDS_OperatorLanding
    flows:
      - name: OpcenterEXDS_OperatorLanding.ACT_Complete_ValidateUser
        calls:
          - name: OpcenterEXDS_OperatorLanding.SUB_OperatorLanding_Complete
            - name: OpcenterEXDS_OperatorLanding.SUB_Complete_Operation
              - name: OpcenterEXDS_OperatorLanding.SUB_Call_Complete_Operation
                - name: OpcenterEXDS_OperatorLanding_Connector.UADMCompleteWOOperationSerializedList
                  - name: OpcenterEXDS_OperatorLanding_Connector.SUB_UADMCompleteWOOperationSerializedList_MF
                - name: OpcenterEXDS_OperatorLanding_Connector.UADMCompleteWOOperationFullQtyMultiMachineList
                  - name: OpcenterEXDS_OperatorLanding_Connector.SUB_UADMCompleteWOOperationFullQtyMultiMachineList_MF
      - name: OpcenterEXDS_OperatorLanding.ACT_Complete_ValidateUser_FromDetails
        calls:
          - name: OpcenterEXDS_OperatorLanding.SUB_Details_Complete
            - name: OpcenterEXDS_OperatorLanding.SUB_Complete_Operation
              - name: OpcenterEXDS_OperatorLanding.SUB_Call_Complete_Operation
                - name: OpcenterEXDS_OperatorLanding_Connector.UADMCompleteWOOperationSerializedList
                  - name: OpcenterEXDS_OperatorLanding_Connector.SUB_UADMCompleteWOOperationSerializedList_MF
                - name: OpcenterEXDS_OperatorLanding_Connector.UADMCompleteWOOperationFullQtyMultiMachineList
                  - name: OpcenterEXDS_OperatorLanding_Connector.SUB_UADMCompleteWOOperationFullQtyMultiMachineList_MF
            - name: OpcenterEXDS_OperatorLanding.ACT_Complete_Step
              - name: OpcenterEXDS_OperatorLanding.SUB_Call_Complete_Step
                - name: OpcenterEXDS_OperatorLanding_Connector.CompleteWOStepSerialized
                  - name: OpcenterEXDS_OperatorLanding_Connector.SUB_CompleteWOStepSerialized_MF
                - name: OpcenterEXDS_OperatorLanding_Connector.CompleteWOStepFullQty
                  - name: OpcenterEXDS_OperatorLanding_Connector.SUB_CompleteWOStepFullQty_MF
    target_commands:
      - AppU4DM.UADMCompleteWOOperationSerializedList
      - AppU4DM.UADMCompleteWOOperationFullQtyMultiMachineList
      - AppU4DM.CompleteWOStepSerialized
      - AppU4DM.CompleteWOStepFullQty

  - name: Routing_Popup
    module: OpcenterEXDS_OperatorLanding
    flows:
      []
    target_commands:
      []

  - name: PANEL_WorkInstructionsToWOStep_Add
    module: OpcenterEXDS_OperatorLanding
    flows:
      - name: OpcenterEXDS_OperatorLanding.ACT_LinkWorkInstructionsToWOStep
        calls:
          - name: OpcenterEXDS_OperatorLanding_Connector.SUB_LinkWorkInstructionsToWOStep_MF
    target_commands:
      - AppU4DM.LinkWorkInstructionsToWOStep

  - name: PANEL_WorkInstructionsToWOOperation_Add
    module: OpcenterEXDS_OperatorLanding
    flows:
      - name: OpcenterEXDS_OperatorLanding.ACT_LinkWorkInstructionsToWOOperation
        calls:
          - name: OpcenterEXDS_OperatorLanding_Connector.SUB_LinkWorkInstructionsToWOOperation_MF
    target_commands:
      - AppU4DM.LinkWorkInstructionsToWOOperation

  - name: Routing
    module: OpcenterEXDS_OperatorLanding
    flows:
      []
    target_commands:
      []

  - name: PANEL_SelectTool
    module: OpcenterEXDS_OperatorLanding
    flows:
      []
    target_commands:
      []

  - name: PANEL_CompleteAssignedNonProductiveActivities
    module: OpcenterEXDS_OperatorLanding
    flows:
      - name: OpcenterEXDS_OperatorLanding.ACT_CompleteNonProductiveActivites
        calls:
          - name: OpcenterEXDS_OperatorLanding_Connector.UADMCompleteNonProductiveActivityList
            - name: OpcenterEXDS_OperatorLanding_Connector.SUB_UADMCompleteNonProductiveActivityList_MF
    target_commands:
      - AppU4DM.UADMCompleteNonProductiveActivityList

  - name: PANEL_AssignAndStartNonProductiveActivities
    module: OpcenterEXDS_OperatorLanding
    flows:
      - name: OpcenterEXDS_OperatorLanding.ACT_StartNonProductiveActivites
        calls:
          - name: OpcenterEXDS_OperatorLanding_Connector.UADMStartNonProductiveActivityList
            - name: OpcenterEXDS_OperatorLanding_Connector.SUB_UADMStartNonProductiveActivityList_MF
    target_commands:
      - AppU4DM.UADMStartNonProductiveActivityList

  - name: Home_Web
    module: OpcenterEXDS_OperatorLanding
    flows:
      []
    target_commands:
      []

  - name: PANEL_WOOP_NonConformance
    module: OpcenterEXDS_OperatorLanding
    flows:
      []
    target_commands:
      []

  - name: PANEL_VerticalCommandBarNC
    module: OpcenterEXDS_OperatorLanding
    flows:
      []
    target_commands:
      []

  - name: PANEL_VerticalCommandBarNC_Details
    module: OpcenterEXDS_OperatorLanding
    flows:
      []
    target_commands:
      []

```

---

## 7. PageCommands

Simplified view showing only the target commands for each page/panel.

| Page/Panel | Module | Target AppName | Target CommandName |
|------------|--------|----------------|--------------------|
| PartProgramHistoryPopUp | OpcenterEXDS_PartProgram | - | - |
| PartProgramDetailsPopUp | OpcenterEXDS_PartProgram | Material<br>Material | CreateMaterialTrackingUnitProperties<br>UpdateMaterialTrackingUnitProperties |
| PANEL_SelectEquipmentAcquireWI | OpcenterEXDS_EXFN_WorkInstruction | AppU4DM<br>@EXFN_WorkInstruction.AppName | AcquireDCItemValueFromAutomationNodeInstanceParameter<br>InEditingWorkInstruction |
| ScrapMaterialConsumptionPopup | OpcenterEXDS_OperatorLanding | AppU4DM | ScrapConsumedMaterial |
| AddDocumentsPopup | OpcenterEXDS_OperatorLanding | AppU4DM | CreateToBeUsedDocuments |
| OperatorLanding | OpcenterEXDS_OperatorLanding | AppU4DM | UADMCheckCertificationMultiMachineOnOpenWOOperation |
| NonConformancePopup | OpcenterEXDS_OperatorLanding | AppU4DM | UADMCreateNonConformanceV3_1 |
| OperatorTerminal | OpcenterEXDS_OperatorLanding | - | - |
| ScrapProducedMaterialPopup | OpcenterEXDS_OperatorLanding | AppU4DM | ScrapWorkOrderSerialNumbers |
| PANEL_DisassembleMaterialTrackingUnit | OpcenterEXDS_OperatorLanding | AppU4DM | DisassembleMaterialItem |
| PANEL_ImportDocument | OpcenterEXDS_OperatorLanding | AppU4DM<br>AppU4DM | UADMCreateDocument<br>PropagateSegregationTagsToDocument |
| PANEL_LinkDocuments | OpcenterEXDS_OperatorLanding | - | - |
| PANEL_SelectToBeConsumedMTU | OpcenterEXDS_OperatorLanding | - | - |
| PANEL_SelectDestinationContainer | OpcenterEXDS_OperatorLanding | - | - |
| PANEL_AddDocumentsImportDocument | OpcenterEXDS_OperatorLanding | AppU4DM<br>AppU4DM | UADMCreateDocument<br>PropagateSegregationTagsToDocument |
| PANEL_ChangePackage | OpcenterEXDS_OperatorLanding | AppU4DM | CreateChangePackage |
| PANEL_SkipWOOperation | OpcenterEXDS_OperatorLanding | AppU4DM<br>AppU4DM | UADMSkipWOOperationFullQty<br>UADMSkipWOOperationSerialized |
| PANEL_Notes | OpcenterEXDS_OperatorLanding | AppU4DM<br>AppU4DM | UADMCreateSnagAndNoteList<br>UADMConfirmSnagAndNoteList |
| PANEL_VerticalCommandBarMore | OpcenterEXDS_OperatorLanding | - | - |
| PANEL_OperatorDetailsCompleteStep | OpcenterEXDS_OperatorLanding | AppU4DM<br>AppU4DM | CompleteWOStepSerialized<br>CompleteWOStepFullQty |
| PANEL_SetPoint | OpcenterEXDS_OperatorLanding | AppU4DM | TransmitEquipmentSetPointToAutomationNodeParameters |
| PANEL_VerticalCommandBarMore_Details | OpcenterEXDS_OperatorLanding | - | - |
| PANEL_ActiveUserList | OpcenterEXDS_OperatorLanding | - | - |
| PANEL_OperatorLandingStart | OpcenterEXDS_OperatorLanding | AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM | UADMStartWOOperationSerializedList<br>UADMStartOperation<br>UADMCheckCertificationMultiMachineOnOpenWOOperation<br>AssignProducedMaterialItems<br>CreateAndAssignProducedMaterialItems<br>UADMCreateAndAssignProducedMaterialItems |
| PANEL_OperatorLandingGoToDetailsSN | OpcenterEXDS_OperatorLanding | - | - |
| PANEL_OperatorLandingHold | OpcenterEXDS_OperatorLanding | AppU4DM | UADMSetWorkOrderHoldList |
| PANEL_CloseFlexibleWorkOrder | OpcenterEXDS_OperatorLanding | AppU4DM<br>AppU4DM | UADMAbruptlyCloseFlexibleWorkOrder<br>AutoGenerateWorkOrderNId |
| PANEL_OperatorLandingComplete | OpcenterEXDS_OperatorLanding | AppU4DM<br>AppU4DM | UADMCompleteWOOperationSerializedList<br>UADMCompleteWOOperationFullQtyMultiMachineList |
| PANEL_SetPointHistory | OpcenterEXDS_OperatorLanding | - | - |
| PANEL_OperatorDetailsStartStep | OpcenterEXDS_OperatorLanding | AppU4DM<br>AppU4DM | StartWOStepSerialized<br>StartWOStepFullQty |
| PANEL_OperatorLandingPause | OpcenterEXDS_OperatorLanding | AppU4DM | UADMPauseWorkOrderOperationMultiMachineList |
| PANEL_ScrewingDetails | OpcenterEXDS_OperatorLanding | - | - |
| PANEL_ChangeSN | OpcenterEXDS_OperatorLanding | AppU4DM<br>AppU4DM | UpdateSerialNumberList<br>AutoGenerateMTUCode |
| PANEL_Start_ValidateUser_FromDetails | OpcenterEXDS_OperatorLanding | AppU4DM<br>AppU4DM<br>AppU4DM | UADMStartWOOperationSerializedList<br>UADMStartOperation<br>UADMCheckCertificationMultiMachineOnOpenWOOperation |
| PANEL_Start_ValidateUser | OpcenterEXDS_OperatorLanding | AppU4DM<br>AppU4DM<br>AppU4DM | UADMStartWOOperationSerializedList<br>UADMStartOperation<br>UADMCheckCertificationMultiMachineOnOpenWOOperation |
| PANEL_Pause_ValidateUser | OpcenterEXDS_OperatorLanding | - | - |
| PANEL_Complete_ValidateUser | OpcenterEXDS_OperatorLanding | AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM | UADMCompleteWOOperationSerializedList<br>UADMCompleteWOOperationFullQtyMultiMachineList<br>CompleteWOStepSerialized<br>CompleteWOStepFullQty |
| Routing_Popup | OpcenterEXDS_OperatorLanding | - | - |
| PANEL_WorkInstructionsToWOStep_Add | OpcenterEXDS_OperatorLanding | AppU4DM | LinkWorkInstructionsToWOStep |
| PANEL_WorkInstructionsToWOOperation_Add | OpcenterEXDS_OperatorLanding | AppU4DM | LinkWorkInstructionsToWOOperation |
| Routing | OpcenterEXDS_OperatorLanding | - | - |
| PANEL_SelectTool | OpcenterEXDS_OperatorLanding | - | - |
| PANEL_CompleteAssignedNonProductiveActivities | OpcenterEXDS_OperatorLanding | AppU4DM | UADMCompleteNonProductiveActivityList |
| PANEL_AssignAndStartNonProductiveActivities | OpcenterEXDS_OperatorLanding | AppU4DM | UADMStartNonProductiveActivityList |
| Home_Web | OpcenterEXDS_OperatorLanding | - | - |
| PANEL_WOOP_NonConformance | OpcenterEXDS_OperatorLanding | - | - |
| PANEL_VerticalCommandBarNC | OpcenterEXDS_OperatorLanding | - | - |
| PANEL_VerticalCommandBarNC_Details | OpcenterEXDS_OperatorLanding | - | - |

---

_Report generated by export_manifest tool_
