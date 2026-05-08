# Manifest Report: Opcenter EX DS Complex Manufacturing

**Mendix Version:** 11.6.4  
**MPR File:** C:\Workspaces\Mendix\Complex\main\Opcenter EX DS Complex Manufacturing.mpr  
**Generated:** 2026-05-08 09:05:56  

---

## Summary

- **External Entities:** 87 (across 6 modules)
- **Microflow/Action Calls:** 82
- **Signal Manager Subscriptions:** 64 subscription(s)
- **Navigation Items:** 3
- **Pages/Panels:** 154 analyzed, 3 with commands

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
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnSkipWOOperation | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnSetHoldWorkOrder | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnSentenceNonConformanceV3_1 | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | StatusChanged | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | WorkOrderSerialNumbersScrapped | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnStartWorkOrderOperationFullQty | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnCompleteWorkOrderOperationFullQty | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnCompleteWorkOrderOperationSerialized | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnPauseWorkOrderOperation | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnReopenWorkOrderOperation | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnStartWorkOrderOperationSerialized | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | SnagAndNoteNotificationSgn | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnChangeFlexibleWOOpStatusToComplete | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnCreateNonConformanceV3_1 | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnCompleteReworkOrder | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnFAICompleted | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnFAICandidateDeclared | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | StatusChanged | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | WorkOrderSerialNumbersScrapped | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnCompleteWorkOrderOperationSerialized | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnPauseWorkOrderOperation | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnStartWorkOrderStepFullQty | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | SnagAndNoteNotificationSgn | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnUsedToolSignal | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | WorkInstructionStatusChangedSignal | WorkInstruction | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnReopenWorkOrderOperation | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnFAICompleted | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnCompleteWorkOrderOperationFullQty | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnSetHoldWorkOrder | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnCreateNonConformanceV3_1 | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnFAICandidateDeclared | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnStartWorkOrderOperationSerialized | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnCompleteWorkOrderStepFullQty | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnCompleteReworkOrder | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnFAIRemoved | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnSkipWOOperation | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnStartWorkOrderOperationFullQty | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnStartWorkOrderStepSerialized | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnCompleteWorkOrderStepSerialized | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnSentenceNonConformanceV3_1 | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnAssemblyWorkOrderOperation | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnDisassemblyWorkOrderOperation | AppU4DM | Yes |
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

## 5. Pages/Panels

All pages and panels with their referenced commands (excluding marketplace and UI modules).

**Note:** Commands are found by recursively following microflow call chains (Page→Microflow→...→Command). The analysis traces MicroflowCall, JavaAction, and ExternalAction types.

**Limitation:** Inline nanoflows (nanoflows embedded directly in pages, not stored as separate Units) are not currently traced. Only standalone microflows stored in the Unit table can be analyzed recursively.

Found 154 page(s)/panel(s):

| Name | Module | Target Commands |
|------|--------|----------------|
| ScrapMaterialConsumptionQuantity | OpcenterEXDS_OperatorLanding | - |
| HeaderBar | OpcenterEXDS_OperatorLanding | - |
| WorkInstructions | OpcenterEXDS_OperatorLanding | - |
| ToBeConsumedMaterialsGridCustomPart | OpcenterEXDS_OperatorLanding | - |
| DocumentViewer_ForDocSetLinkedEntity | EXFN_DocumentViewer | - |
| PANEL_VerticalCommandBarMore | OpcenterEXDS_OperatorLanding | - |
| ToBeConsumedMaterialsGrid | OpcenterEXDS_OperatorLanding | - |
| SNP_WorkInstruction_VerticalView | OpcenterEXDS_EXFN_WorkInstruction | - |
| NonConformanceMaterialTrackingUnits_MaterialTrackingUnitContext | OpcenterEXDS_OperatorLanding | - |
| SelectedToolBadge_ToolContext | OpcenterEXDS_OperatorLanding | - |
| DocumentViewerSnippet | OpcenterEXDS_EXFN_DocumentViewer | - |
| Home_Web | OpcenterEXDS_OperatorLanding | - |
| ScrapMaterialConsumptionPopup | OpcenterEXDS_OperatorLanding | - |
| NonConformance_ToolContext | OpcenterEXDS_OperatorLanding | - |
| LazyLoader | OpcenterEXDS_OperatorLanding | - |
| PANEL_WorkInstructionsToWOStep_Add | OpcenterEXDS_OperatorLanding | - |
| PANEL_OperatorDetailsCompleteStep | OpcenterEXDS_OperatorLanding | - |
| Datatime | EXFN_WorkInstruction | - |
| AddDocumentsMaterialTrackingUnits | OpcenterEXDS_OperatorLanding | - |
| PANEL_VerticalCommandBarNC_Details | OpcenterEXDS_OperatorLanding | - |
| PANEL_SetPoint | OpcenterEXDS_OperatorLanding | - |
| SelectedMaterialTrackingUnitBadge_WorkOrderOperationContext | OpcenterEXDS_OperatorLanding | - |
| SNP_WorkInstruction_Classic | EXFN_WorkInstruction | - |
| PANEL_VerticalCommandBarMore_Details | OpcenterEXDS_OperatorLanding | - |
| ToBeConsumedMaterialsDisassemble | OpcenterEXDS_OperatorLanding | - |
| WorkInstructionDefinitions_Overview_Example | EXFN_WorkInstruction | - |
| VerticalCommandBar_More | OpcenterEXDS_OperatorLanding | AppU4DM
.TriggerPrintingOnWorkOrderOperation
 |
| DocumentViewerSnippet_DNC | OpcenterEXDS_PartProgram | - |
| VerticalCommandBar | OpcenterEXDS_OperatorLanding | - |
| PANEL_ActiveUserList | OpcenterEXDS_OperatorLanding | - |
| ToBeConsumedMaterialsGridDisassemble | OpcenterEXDS_OperatorLanding | - |
| AddDocumentsSummary | OpcenterEXDS_OperatorLanding | - |
| PANEL_Start_ValidateUser_FromDetails | OpcenterEXDS_OperatorLanding | - |
| FailureContextualMenu | EXFN_Quality | - |
| FailureBrowser | OpcenterEXDS_OperatorLanding | - |
| ToBeUsedTools_Backup | OpcenterEXDS_OperatorLanding | - |
| ToBeConsumedMaterialsSelectedFit | OpcenterEXDS_OperatorLanding | - |
| PANEL_SelectToBeConsumedMTU | OpcenterEXDS_OperatorLanding | - |
| PANEL_LinkExistingDocument | EXFN_DocumentViewer | - |
| Integer | EXFN_WorkInstruction | - |
| WorkInstructionDefinitions_Overview | EXFN_WorkInstruction | - |
| Routing | OpcenterEXDS_OperatorLanding | - |
| PANEL_SkipWOOperation | OpcenterEXDS_OperatorLanding | - |
| SPCChart_VisualAndAttributive | EXFN_Quality | - |
| SNP_WorkInstruction_VerticalView | EXFN_WorkInstruction | - |
| ToBeConsumedMaterialCustomProduced | OpcenterEXDS_OperatorLanding | - |
| SelectedMaterialTrackingUnitBadge_MaterialTrackingUnitContext | OpcenterEXDS_OperatorLanding | - |
| SNP_QualityInspection_Attributive | EXFN_Quality | - |
| AddDocumentsNavigationWizard | OpcenterEXDS_OperatorLanding | - |
| SNP_QualityInspection_Visual | EXFN_Quality | - |
| ToBeConsumedMaterialsGridRangePartsAsRequired | OpcenterEXDS_OperatorLanding | - |
| PANEL_ImportDocument | OpcenterEXDS_OperatorLanding | - |
| PANEL_OperatorLandingStart | OpcenterEXDS_OperatorLanding | - |
| WorkInstruction_View | EXFN_WorkInstruction | - |
| ToBeCoByProducedMaterials | OpcenterEXDS_OperatorLanding | - |
| PANEL_WOOP_NonConformance | OpcenterEXDS_OperatorLanding | - |
| PANEL_OperatorLandingGoToDetailsSN | OpcenterEXDS_OperatorLanding | - |
| PANEL_OperatorLandingHold | OpcenterEXDS_OperatorLanding | - |
| QualityInspectionVisual_FullscreenPopup | EXFN_Quality | - |
| PANEL_Notes | OpcenterEXDS_OperatorLanding | - |
| AddDocumentsPopup | OpcenterEXDS_OperatorLanding | - |
| ToBeConsumedMaterialsAlternative | OpcenterEXDS_OperatorLanding | - |
| WorkInstruction_View_Vertical | EXFN_WorkInstruction | - |
| Dropdown | EXFN_WorkInstruction | - |
| PANEL_CloseFlexibleWorkOrder | OpcenterEXDS_OperatorLanding | AppU4DM.UADMAbruptlyCloseFlexibleWorkOrder |
| PANEL_CompleteAssignedNonProductiveActivities | OpcenterEXDS_OperatorLanding | - |
| ScrapFailureBrowser | OpcenterEXDS_OperatorLanding | - |
| Decimal | EXFN_WorkInstruction | - |
| NonConformanceMaterialTrackingUnits_WorkOrderOperationContext | OpcenterEXDS_OperatorLanding | - |
| PANEL_OperatorLandingComplete | OpcenterEXDS_OperatorLanding | - |
| Document_Overview | EXFN_DocumentViewer | - |
| ToBeConsumedMaterialsGridNormalPart | OpcenterEXDS_OperatorLanding | - |
| ToBeConsumedMaterialsGridSelectedFit | OpcenterEXDS_OperatorLanding | - |
| SNP_QualityInspection | EXFN_Quality | - |
| OperationContainer | OpcenterEXDS_OperatorLanding | - |
| AddDocumentsSelectedMaterialTrackingUnitBadge | OpcenterEXDS_OperatorLanding | - |
| ToBeConsumedMaterialCoByProduced | OpcenterEXDS_OperatorLanding | - |
| PartProgramHistoryPopUp | OpcenterEXDS_PartProgram | - |
| NonConformancesSummary | OpcenterEXDS_OperatorLanding | - |
| OperatorLanding | OpcenterEXDS_OperatorLanding | - |
| PartProgram | OpcenterEXDS_OperatorLanding | - |
| PANEL_Complete_ValidateUser | OpcenterEXDS_OperatorLanding | - |
| NonConformancesInfo | OpcenterEXDS_OperatorLanding | - |
| PANEL_SelectEquipmentAcquireWI | OpcenterEXDS_EXFN_WorkInstruction | - |
| ToBeConsumedMaterialGridCoByProduced | OpcenterEXDS_OperatorLanding | - |
| ScenarioInstanceView | EXFN_ElectronicSignature | - |
| AddDocumentsSelectedDocumentBadge | OpcenterEXDS_OperatorLanding | - |
| SelectedFailureBadge | OpcenterEXDS_OperatorLanding | - |
| RoutingNode | OpcenterEXDS_OperatorLanding | - |
| NonConformanceEquipment_EquipmentContext | OpcenterEXDS_OperatorLanding | - |
| ToBeConsumedMaterialsCustomPart | OpcenterEXDS_OperatorLanding | - |
| Failure_Overview | EXFN_Quality | - |
| PANEL_SetPointHistory | OpcenterEXDS_OperatorLanding | - |
| PANEL_LinkExistingDocumentSet | EXFN_DocumentViewer | - |
| DataCollector | EXFN_WorkInstruction | - |
| PANEL_AssignAndStartNonProductiveActivities | OpcenterEXDS_OperatorLanding | - |
| PANEL_ChangePackage | OpcenterEXDS_OperatorLanding | - |
| AddDocuments | OpcenterEXDS_OperatorLanding | - |
| NonConformancePopup | OpcenterEXDS_OperatorLanding | - |
| SPCChart_Variable | EXFN_Quality | - |
| PANEL_CreateLinkDocument | EXFN_DocumentViewer | - |
| ToBeConsumedMaterialsNormalPart | OpcenterEXDS_OperatorLanding | - |
| ToBeProducedMaterial | OpcenterEXDS_OperatorLanding | - |
| ToBeUsedTools | OpcenterEXDS_OperatorLanding | - |
| VerticalCommandBarPartProgram | OpcenterEXDS_PartProgram | - |
| QualityInspections | OpcenterEXDS_OperatorLanding | - |
| WorkInstruction_Overview | EXFN_WorkInstruction | - |
| DocumentViewerSnippet | EXFN_DocumentViewer | - |
| ToBeConsumedMaterialsReference | OpcenterEXDS_OperatorLanding | - |
| PANEL_SelectTool | OpcenterEXDS_OperatorLanding | - |
| ScrapMaterialTrackingUnits | OpcenterEXDS_OperatorLanding | - |
| ES_Button | EXFN_ElectronicSignature | - |
| SPCResponseMessages | EXFN_Quality | - |
| PANEL_DisassembleMaterialTrackingUnit | OpcenterEXDS_OperatorLanding | - |
| SelectedDocumentBadge | OpcenterEXDS_OperatorLanding | - |
| PANEL_OperatorDetailsStartStep | OpcenterEXDS_OperatorLanding | - |
| ToBeConsumedMaterialsGridReference | OpcenterEXDS_OperatorLanding | - |
| PANEL_Start_ValidateUser | OpcenterEXDS_OperatorLanding | - |
| MultipleChoice | EXFN_WorkInstruction | - |
| WorkInstruction_New_Step2 | EXFN_WorkInstruction | - |
| PANEL_VerticalCommandBarNC | OpcenterEXDS_OperatorLanding | - |
| PartProgramDetailsPopUp | OpcenterEXDS_PartProgram | - |
| DocumentSelection_View | EXFN_DocumentViewer | - |
| ToBeConsumedMaterialGridAlternative | OpcenterEXDS_OperatorLanding | - |
| VerticalCommandBar_NC | OpcenterEXDS_OperatorLanding | - |
| PANEL_OperatorLandingPause | OpcenterEXDS_OperatorLanding | - |
| ToBeConsumedMaterialsPrekit | OpcenterEXDS_OperatorLanding | - |
| OperationContainerHeaderBar | OpcenterEXDS_OperatorLanding | - |
| WorkInstruction_VerticalView_Example | EXFN_WorkInstruction | - |
| PANEL_LinkDocuments | OpcenterEXDS_OperatorLanding | - |
| OperatorTerminal | OpcenterEXDS_OperatorLanding | - |
| ReadMe | EXFN_Quality | - |
| PANEL_ScrewingDetails | OpcenterEXDS_OperatorLanding | - |
| PANEL_ChangeSN | OpcenterEXDS_OperatorLanding | AppU4DM.UpdateSerialNumberList |
| DocumentViewer_ForLinkedEntity | EXFN_DocumentViewer | - |
| ToBeConsumedMaterialsRangePartsAndAsRequired | OpcenterEXDS_OperatorLanding | - |
| PANEL_AddDocumentsImportDocument | OpcenterEXDS_OperatorLanding | - |
| PANEL_WorkInstructionsToWOOperation_Add | OpcenterEXDS_OperatorLanding | - |
| PANEL_SelectDestinationContainer | OpcenterEXDS_OperatorLanding | - |
| PANEL_Pause_ValidateUser | OpcenterEXDS_OperatorLanding | - |
| ScrapProducedMaterialPopup | OpcenterEXDS_OperatorLanding | - |
| NonConformanceDocuments | OpcenterEXDS_OperatorLanding | - |
| Routing_Popup | OpcenterEXDS_OperatorLanding | - |
| SNP_QualityInspectionContainer | EXFN_Quality | - |
| SNP_QualityInspection_Variable | EXFN_Quality | - |
| NonConformancesNavigationWizard | OpcenterEXDS_OperatorLanding | - |
| Checkbox | EXFN_WorkInstruction | - |
| WorkInstruction_Overview_Example | EXFN_WorkInstruction | - |
| OperationList | OpcenterEXDS_OperatorLanding | - |
| SelectedEquipmentBadge_EquipmentContext | OpcenterEXDS_OperatorLanding | - |
| WorkInstruction_New_Step1 | EXFN_WorkInstruction | - |
| Text | EXFN_WorkInstruction | - |
| WorkInstruction_Preview | EXFN_WorkInstruction | - |
| Multiline | EXFN_WorkInstruction | - |

---

## 6. Pages/Panels Commands Hierarchy

Microflows and nanoflows called by each page/panel (in YAML structure).

```yaml
pages:
  - name: ScrapMaterialConsumptionQuantity
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_SearchToBeConsumedMTU
      - OpcenterEXDS_OperatorLanding.ACT_ShowSelectToBeConsumedMTUPanel
      - OpcenterEXDS_OperatorLanding.ACT_ChangeScrappedQuantity
      - OpcenterEXDS_OperatorLanding.ACT_IncreaseScrapAssembledQuantity
      - OpcenterEXDS_OperatorLanding.ACT_DescreaseScrapAssembledQuantity
      - OpcenterEXDS_OperatorLanding.ACT_SearchDestinationContainer
      - OpcenterEXDS_OperatorLanding.ACT_ShowSelectDestinationContainerPanel
    target_commands:
      []

  - name: HeaderBar
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_TimerRefreshRemainingTime
      - OpcenterEXDS_OperatorLanding.DS_CreateTimerContext
      - OpcenterEXDS_OperatorLanding.ACT_Show_SelectSerialNumber_Panel
      - OpcenterEXDS_OperatorLanding.ACT_HeaderBar_EquipmentList_Click
    target_commands:
      []

  - name: WorkInstructions
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - EXFN_Authentication.Signal_Access_Token
      - OpcenterEXDS_OperatorLanding.EVT_LinkWIOnDemandToSerialNumber
      - EXFN_Authentication.HandleUnauthorizedBehavior
      - OpcenterEXDS_OperatorLanding.DS_GetWorkInstructionView
      - OpcenterEXDS_EXFN_WorkInstruction.DS_GetWorkInstructionContextHelper
      - OpcenterEXDS_OperatorLanding.DS_GetOrCreateWorkInstructionContext
    target_commands:
      []

  - name: ToBeConsumedMaterialsGridCustomPart
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_GetToBeConsumeMaterial_GridMode_CustomPart
      - OpcenterEXDS_OperatorLanding.ACT_ValidateMaterialConsumption
      - OpcenterEXDS_OperatorLanding.ACT_AcquireMTUList
      - OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialToBeConsumedPage
      - OpcenterEXDS_OperatorLanding.ACT_ConsumeMaterial
      - OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialConsumedPage_Grid
      - OpcenterEXDS_OperatorLanding.ACT_DisassembleMaterial_Grid
    target_commands:
      []

  - name: DocumentViewer_ForDocSetLinkedEntity
    module: EXFN_DocumentViewer
    microflows:
      - EXFN_DocumentViewer.ACT_UnlinkDocumentSetList
      - EXFN_DocumentViewer.DS_GetDocumentsByDocumentSetsLinkedToEntity
    target_commands:
      []

  - name: PANEL_VerticalCommandBarMore
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_ClosePanel_VerticalCommandBarMore
    target_commands:
      []

  - name: ToBeConsumedMaterialsGrid
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_ToBeConsumedMaterialHistory
      - OpcenterEXDS_OperatorLanding.DS_GetToBeConsumedMaterial_GridMode
      - OpcenterEXDS_OperatorLanding.SetVisibilitySnippet_GridMode
    target_commands:
      []

  - name: SNP_WorkInstruction_VerticalView
    module: OpcenterEXDS_EXFN_WorkInstruction
    microflows:
      - OpcenterEXDS_EXFN_WorkInstruction.ACT_WISectionRefresh
      - EXFN_Authentication.Signal_Access_Token
      - OpcenterEXDS_EXFN_WorkInstruction.EVT_WorkInstructionSectionCompleted_Details
      - EXFN_Authentication.HandleUnauthorizedBehavior
      - OpcenterEXDS_EXFN_WorkInstruction.DS_CreateWISignalConfiguration
      - EXFN_WorkInstruction.SUB_GetCalculateWorkInstructionFormulaValues_BySection
      - OpcenterEXDS_EXFN_WorkInstruction.EVT_WorkInstructionStepCompleted_Details
      - EXFN_WorkInstruction.SUB_GetCalculateWorkInstructionFormulaValues_ByStep
      - OpcenterEXDS_EXFN_WorkInstruction.DS_Get_WI_Step_Instructions
      - OpcenterEXDS_OperatorLanding.DS_SetAutoGenerateId
      - EXFN_WorkInstruction.DS_WI_SortedItemViewList
      - EXFN_WorkInstruction.DS_ScenarioConfiguration
      - EXFN_WorkInstruction.DS_ScenarioInstanceView
      - OpcenterEXDS_EXFN_WorkInstruction.ACT_WIStep_Confirm
      - OpcenterEXDS_EXFN_WorkInstruction.ACT_WIStep_Acknowledge
      - OpcenterEXDS_EXFN_WorkInstruction.ACT_WIStep_Acquire
      - OpcenterEXDS_EXFN_WorkInstruction.EVT_WIStep_ES_Acquired
      - OpcenterEXDS_EXFN_WorkInstruction.ACT_WIStep_ReEdit
    target_commands:
      []

  - name: NonConformanceMaterialTrackingUnits_MaterialTrackingUnitContext
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_NonConformanceMaterialTrackingUnits_ShowOnlyRelatedButtonClick
      - OpcenterEXDS_OperatorLanding.DS_GetMaterialTrackingUnitsRelatedToTheSelectedWorkOrderOperation
      - OpcenterEXDS_OperatorLanding.ACT_SelectNonConformanceMaterialTrackingUnit_MaterialTrackingUnitContext
      - OpcenterEXDS_OperatorLanding.ACT_SelectNonConformanceMaterialTrackingUnit_ByView_MaterialTrackingUnitContext
      - OpcenterEXDS_OperatorLanding.DS_GetLayoutHorizontalMode
    target_commands:
      []

  - name: SelectedToolBadge_ToolContext
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_DeselectTool
      - OpcenterEXDS_OperatorLanding.DS_GetSelectedTools_Vertical
    target_commands:
      []

  - name: DocumentViewerSnippet
    module: OpcenterEXDS_EXFN_DocumentViewer
    microflows:
      - EXFN_DocumentViewer.ACT_SetPMIVisibility
      - EXFN_DocumentViewer.ACT_FitToView
      - OpcenterEXDS_EXFN_DocumentViewer.ACT_EXDSOpenDocument
      - EXFN_DocumentViewer.ACT_ChangeDetailsView
      - EXFN_DocumentViewer.DS_GetFile
      - EXFN_ServiceLayer.DS_GetApplicationURL
      - EXFN_DocumentViewer.ACT_CloseJTFullScreen
      - OpcenterEXDS_EXFN_DocumentViewer.ACT_AddGenerateMarkupToOperation
      - EXFN_DocumentViewer.ToggleMarkup
      - EXFN_DocumentViewer.ACT_ToggleToolbox
      - EXFN_DocumentViewer.ACT_SetMarkupColor
      - OpcenterEXDS_EXFN_DocumentViewer.EVT_OnSelectedObject
      - EXFN_DocumentViewer.ACT_ChangeModel3D
      - EXFN_DocumentViewer.DS_GetCategories
      - EXFN_DocumentViewer.DS_GetDocumentByCategory
      - EXFN_DocumentViewer.ACT_Carousel
      - EXFN_DocumentViewer.ACT_ChangeDocumentSelection
      - EXFN_DocumentViewer.ACT_ButtonDownload
      - EXFN_DocumentViewer.DS_GetDocumentSetList
    target_commands:
      []

  - name: Home_Web
    module: OpcenterEXDS_OperatorLanding
    microflows:
      []
    target_commands:
      []

  - name: ScrapMaterialConsumptionPopup
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_ScrapMaterialConsumedPage_RegisterClosePageAction
      - OpcenterEXDS_OperatorLanding.DS_GetSelectedFailuresForScrap
      - OpcenterEXDS_OperatorLanding.ACT_DeselectScrapFailure
      - OpcenterEXDS_OperatorLanding.ACT_ScrapConsumedMaterial
      - OpcenterEXDS_OperatorLanding.NAV_Back
    target_commands:
      []

  - name: NonConformance_ToolContext
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_SelectNonConformanceTool
      - OpcenterEXDS_OperatorLanding.DS_GetLayoutHorizontalMode
    target_commands:
      []

  - name: LazyLoader
    module: OpcenterEXDS_OperatorLanding
    microflows:
      []
    target_commands:
      []

  - name: PANEL_WorkInstructionsToWOStep_Add
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_LinkWorkInstructionsToWOStep
    target_commands:
      []

  - name: PANEL_OperatorDetailsCompleteStep
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_Complete_Panel_Step
      - OpcenterEXDS_OperatorLanding.ACT_Select_CompleteStepPanel_Equipment
    target_commands:
      []

  - name: Datatime
    module: EXFN_WorkInstruction
    microflows:
      - EXFN_WorkInstruction.DS_ItemView_Datatime
      - EXFN_WorkInstruction.ACT_Date_change
      - EXFN_WorkInstruction.ACT_dateTime_null_Set
      - EXFN_WorkInstruction.ACT_Time_Change
      - EXFN_WorkInstruction.ACT_DateTime_New
      - EXFN_WorkInstruction.ACT_CalculateWorkInstructionFormulaValues_Datetime
    target_commands:
      []

  - name: AddDocumentsMaterialTrackingUnits
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_GetActiveMaterialTrackingUnitsRelatedToTheSelectedWorkOrderOperationOrStep
      - OpcenterEXDS_OperatorLanding.ACT_AddDocumentsSelectMaterialTrackingUnit
      - OpcenterEXDS_OperatorLanding.DS_GetLayoutHorizontalMode
    target_commands:
      []

  - name: PANEL_VerticalCommandBarNC_Details
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_ClosePanel_VerticalCommandBarMore
    target_commands:
      []

  - name: PANEL_SetPoint
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_OpenSetPointHistoryPanel
      - OpcenterEXDS_OperatorLanding.DS_SetPointVariables
      - OpcenterEXDS_OperatorLanding.ACT_TransmitEquipmentSetPointToAutomationNodeParameters
      - OpcenterEXDS_OperatorLanding.ACT_TransmitAllEquipmentSetPointToAutomationNodeParameters
    target_commands:
      []

  - name: SelectedMaterialTrackingUnitBadge_WorkOrderOperationContext
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_DeselectMaterialTrackingUnit_MaterialTrackingUnitContext
      - OpcenterEXDS_OperatorLanding.DS_GetSelectedMaterialTrackingUnits_Vertical_WorkOrderOperationContext
      - OpcenterEXDS_OperatorLanding.ACT_DeselectMaterialTrackingUnit_WorkOrderOperationContext
    target_commands:
      []

  - name: SNP_WorkInstruction_Classic
    module: EXFN_WorkInstruction
    microflows:
      - EXFN_WorkInstruction.SUB_GetCalculateWorkInstructionFormulaValues_BySection
      - EXFN_WorkInstruction.SUB_GetCalculateWorkInstructionFormulaValues_ByStep
      - EXFN_WorkInstruction.DS_WI_SortedItemViewList
      - EXFN_WorkInstruction.DS_ScenarioInstanceView
      - EXFN_WorkInstruction.DS_ScenarioConfiguration
      - EXFN_WorkInstruction.ACT_WIStep_Confirm
      - EXFN_WorkInstruction.ACT_WIStep_Acknowledge
      - EXFN_WorkInstruction.EVT_WIStep_ES_Acquired
      - EXFN_WorkInstruction.ACT_WIStep_ReEdit
      - EXFN_WorkInstruction.DS_StepView
      - EXFN_WorkInstruction.ACT_StepNavigation_First
      - EXFN_WorkInstruction.ACT_StepNavigation_Previous
      - EXFN_WorkInstruction.DS_Step_Sequence
      - EXFN_WorkInstruction.ACT_StepNavigation_Set
      - EXFN_WorkInstruction.ACT_StepNavigation_Next
      - EXFN_WorkInstruction.ACT_StepNavigation_Last
    target_commands:
      []

  - name: PANEL_VerticalCommandBarMore_Details
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_ClosePanel_VerticalCommandBarMore
    target_commands:
      []

  - name: ToBeConsumedMaterialsDisassemble
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_GetToBeConsumeMaterial_Disassemble
      - OpcenterEXDS_OperatorLanding.DS_GetHistoryDisassembledMaterial
      - OpcenterEXDS_OperatorLanding.ACT_ValidateMaterialConsumption
      - OpcenterEXDS_OperatorLanding.ACT_DisassembleMaterialSpecificationType
      - OpcenterEXDS_EXFN_DocumentViewer.ACT_ShowAllPartson3dViewer
    target_commands:
      []

  - name: WorkInstructionDefinitions_Overview_Example
    module: EXFN_WorkInstruction
    microflows:
      - EXFN_WorkInstruction.NAV_WorkInstructionPreview_Example
    target_commands:
      []

  - name: VerticalCommandBar_More
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_History
      - OpcenterEXDS_OperatorLanding.ACT_Open_Genealogy
      - OpcenterEXDS_OperatorLanding.ACT_Open_AsBuilt
      - OpcenterEXDS_OperatorLanding.ACT_GoTo_ChangePackage
      - OpcenterEXDS_OperatorLanding.ACT_Open_NonProductiveActivities
      - OpcenterEXDS_OperatorLanding.ACT_Open_ExternalIntegration
      - OpcenterEXDS_OperatorLanding.ACT_SetTargetQuantityOnFlexibleWorkOrder
      - OpcenterEXDS_OperatorLanding.ACT_OpenCloseFlexiblePanel
      - OpcenterEXDS_OperatorLanding.ACT_TriggerPrintingOnWorkOrderOperation
      - OpcenterEXDS_OperatorLanding.ACT_OpenChangeSNPanel
      - OpcenterEXDS_OperatorLanding.ACT_OpenSetPointPanel
      - OpcenterEXDS_OperatorLanding.DS_GetOrCreateVerticalCommandBarContext
    target_commands:
      - AppU4DM
.TriggerPrintingOnWorkOrderOperation


  - name: DocumentViewerSnippet_DNC
    module: OpcenterEXDS_PartProgram
    microflows:
      - OpcenterEXDS_PartProgram.DS_Tree_Context
      - OpcenterEXDS_PartProgram.ACT_SetHeight
      - OpcenterEXDS_PartProgram.ACT_Search
      - OpcenterEXDS_PartProgram.ACT_GetTree
      - OpcenterEXDS_PartProgram.ACT_SetSelected
      - OpcenterEXDS_PartProgram.ACT_ShowWorkOrderTree_Toogle
    target_commands:
      []

  - name: VerticalCommandBar
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_GetOrCreateVerticalCommandBarContext
      - OpcenterEXDS_OperatorLanding.ACT_GoToDetails_Click
      - OpcenterEXDS_OperatorLanding.ACT_Start
      - OpcenterEXDS_OperatorLanding.ACT_Pause_Operation
      - OpcenterEXDS_OperatorLanding.ACT_Hold_Operation
      - OpcenterEXDS_OperatorLanding.ACT_OpenOperationSkip
      - OpcenterEXDS_OperatorLanding.ACT_Complete
      - OpcenterEXDS_OperatorLanding.ACT_ShowDocuments
      - OpcenterEXDS_OperatorLanding.ACT_Open_NC_Popup
      - OpcenterEXDS_OperatorLanding.ACT_OpenChangePackage
      - OpcenterEXDS_OperatorLanding.ACT_AddDocuments_Operation
      - OpcenterEXDS_OperatorLanding.ACT_Open_Notes
      - OpcenterEXDS_OperatorLanding.ACT_Open_ScrapProducedMaterialPage
      - OpcenterEXDS_OperatorLanding.ACT_More
      - OpcenterEXDS_OperatorLanding.ACT_ClosePanel_VerticalCommandBarMore
    target_commands:
      []

  - name: PANEL_ActiveUserList
    module: OpcenterEXDS_OperatorLanding
    microflows:
      []
    target_commands:
      []

  - name: ToBeConsumedMaterialsGridDisassemble
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_GetToBeConsumeMaterial_GridMode_Disassemble
      - OpcenterEXDS_OperatorLanding.ACT_ValidateMaterialConsumption
      - OpcenterEXDS_OperatorLanding.ACT_DisassembleMaterialSpecificationType
    target_commands:
      []

  - name: AddDocumentsSummary
    module: OpcenterEXDS_OperatorLanding
    microflows:
      []
    target_commands:
      []

  - name: PANEL_Start_ValidateUser_FromDetails
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_Start_ValidateUser_FromDetails
      - OpcenterEXDS_OperatorLanding.DS_Create_ESContext
    target_commands:
      []

  - name: FailureContextualMenu
    module: EXFN_Quality
    microflows:
      - EXFN_Quality.DS_QualityInspectionVisual_ContextualMenu_GetLimitedFailureList_WithOthersOption
      - EXFN_Quality.ACT_QualityInspectionVisual_ContextualMenu_SelectionChanged
      - EXFN_Quality.ACT_ContextualMenu_OnMouseLeave
    target_commands:
      []

  - name: FailureBrowser
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_GetSortedBreadcrumb
      - OpcenterEXDS_OperatorLanding.ACT_Navigate_breadcrumb
      - OpcenterEXDS_OperatorLanding.ACT_FailureList
      - OpcenterEXDS_OperatorLanding.ACT_SelectFailure
      - OpcenterEXDS_OperatorLanding.NAV_FailureChildren_Overview
    target_commands:
      []

  - name: ToBeUsedTools_Backup
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_CreateToBeUsedToolContext
      - OpcenterEXDS_OperatorLanding.DS_GetOrCreateSignalEnvelope_UseTool
      - EXFN_Authentication.Signal_Access_Token
      - OpcenterEXDS_OperatorLanding.EVT_OnUsedToolSignal
      - EXFN_Authentication.HandleUnauthorizedBehavior
      - OpcenterEXDS_OperatorLanding.ACT_ToolHistory_ToggleVisibility
      - OpcenterEXDS_OperatorLanding.ACT_ToolInput_ToggleVisibility
      - OpcenterEXDS_OperatorLanding.OCH_ToolNId
      - OpcenterEXDS_OperatorLanding.ACT_ShowSelectToolPanel
      - OpcenterEXDS_OperatorLanding.ACT_UseTool
      - OpcenterEXDS_OperatorLanding.ACT_AcquireTool
      - OpcenterEXDS_OperatorLanding.DS_GetToolEntity
      - OpcenterEXDS_OperatorLanding.ACT_ToBeUsedToolHistory
      - OpcenterEXDS_OperatorLanding.DS_GetTobeUsedToolHistory
      - OpcenterEXDS_OperatorLanding.ACT_ShowToolPanelForGrid
      - OpcenterEXDS_OperatorLanding.ACT_Grid_AcquireTool
      - OpcenterEXDS_OperatorLanding.ACT_Grid_UseTool
      - OpcenterEXDS_OperatorLanding.ACT_UseAllTools
      - OpcenterEXDS_OperatorLanding.ACT_Grid_UseAllTools
    target_commands:
      []

  - name: ToBeConsumedMaterialsSelectedFit
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_GetToBeConsumeMaterial_SelectedFit
      - OpcenterEXDS_OperatorLanding.ACT_ActualMaterials_ToggleVisibility
      - OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialConsumedPage
      - OpcenterEXDS_OperatorLanding.ACT_DisassembleMaterial
      - OpcenterEXDS_EXFN_DocumentViewer.ACT_ShowAllPartson3dViewer
      - OpcenterEXDS_OperatorLanding.ACT_ValidateMaterialConsumption
      - OpcenterEXDS_OperatorLanding.ACT_AcquireMTUList
      - OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialToBeConsumedPage
      - OpcenterEXDS_OperatorLanding.ACT_ConsumeMaterial
    target_commands:
      []

  - name: PANEL_SelectToBeConsumedMTU
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_Cancel_SelectToBeConsumedPanel
      - OpcenterEXDS_OperatorLanding.ACT_Save_SelectToBeConsumedPanel
      - OpcenterEXDS_OperatorLanding.ACT_SelectToBeConsumedMTU
    target_commands:
      []

  - name: PANEL_LinkExistingDocument
    module: EXFN_DocumentViewer
    microflows:
      - EXFN_DocumentViewer.ACT_LinkDocument
      - EXFN_DocumentViewer.ACT_LinkDocument_ChangeUseCurrentDocumentRevision
    target_commands:
      []

  - name: Integer
    module: EXFN_WorkInstruction
    microflows:
      - EXFN_WorkInstruction.DS_ItemView_Integer
      - EXFN_WorkInstruction.ACT_ItemView_Integer_OnChange
      - EXFN_WorkInstruction.ACT_CalculateWorkInstructionFormulaValues_Integer
    target_commands:
      []

  - name: WorkInstructionDefinitions_Overview
    module: EXFN_WorkInstruction
    microflows:
      - EXFN_WorkInstruction.NAV_WorkInstruction_Preview
    target_commands:
      []

  - name: Routing
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_NavigateToOperation
      - OpcenterEXDS_OperatorLanding.DependencyGraphContext_SetHeight
      - OpcenterEXDS_OperatorLanding.ACT_OnClick_RoutingNode
      - OpcenterEXDS_OperatorLanding.DS_GetBreadcrumb
    target_commands:
      []

  - name: PANEL_SkipWOOperation
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_SelectSkipSerialNumber
      - OpcenterEXDS_OperatorLanding.ACT_SelectSkipReason
      - OpcenterEXDS_OperatorLanding.ACT_CallSkipOperation
    target_commands:
      []

  - name: SPCChart_VisualAndAttributive
    module: EXFN_Quality
    microflows:
      - EXFN_Quality.SUB_RetrieveFirstMessage
    target_commands:
      []

  - name: SNP_WorkInstruction_VerticalView
    module: EXFN_WorkInstruction
    microflows:
      - EXFN_WorkInstruction.SUB_GetCalculateWorkInstructionFormulaValues_BySection
      - EXFN_WorkInstruction.SUB_GetCalculateWorkInstructionFormulaValues_ByStep
      - EXFN_WorkInstruction.DS_WI_SortedItemViewList
      - EXFN_WorkInstruction.DS_ScenarioInstanceView
      - EXFN_WorkInstruction.ACT_WIStep_Acquire
      - EXFN_WorkInstruction.ACT_WIStep_Confirm
      - EXFN_WorkInstruction.ACT_WIStep_Acknowledge
      - EXFN_WorkInstruction.EVT_WIStep_ES_Acquired
      - EXFN_WorkInstruction.ACT_WIStep_ReEdit
    target_commands:
      []

  - name: ToBeConsumedMaterialCustomProduced
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_GetToBeCustomProductMaterial
      - OpcenterEXDS_OperatorLanding.ACT_ActualProducedMaterials_ToggleVisibility
      - OpcenterEXDS_OperatorLanding.ACT_ValidateMaterialConsumption
      - OpcenterEXDS_OperatorLanding.ACT_CoProduceMaterial
      - OpcenterEXDS_OperatorLanding.ACT_AcquireMTUList
    target_commands:
      []

  - name: SelectedMaterialTrackingUnitBadge_MaterialTrackingUnitContext
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_DeselectMaterialTrackingUnit_MaterialTrackingUnitContext
      - OpcenterEXDS_OperatorLanding.DS_GetSelectedMaterialTrackingUnits_Vertical_MaterialTrackingUnitContext
    target_commands:
      []

  - name: SNP_QualityInspection_Attributive
    module: EXFN_Quality
    microflows:
      - EXFN_Quality.ACT_SetOkValue
      - EXFN_Quality.ACT_SetNOkValue
      - EXFN_Quality.ACT_Attributive_OnChange
      - EXFN_Quality.ACT_ShowPmiByCharName
      - EXFN_Quality.ACT_DeselectAttributeValueToUpdate
      - EXFN_Quality.ACT_GetAttributiveViewForHistory
      - EXFN_Quality.ACT_Attributive_DisassociateFailure
      - EXFN_Quality.ACT_SelectAttributiveValueToUpdate
      - EXFN_Quality.ACT_SetInspectionValueAndOpenFailure
      - EXFN_Quality.DS_ScenarioInstanceView
      - EXFN_Quality.ACT_ConfirmSample_Attributive
      - EXFN_Quality.ACT_NewSample_Attributive
      - EXFN_Quality.SUB_RetrieveCalcualatedJson
    target_commands:
      []

  - name: AddDocumentsNavigationWizard
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_AddDocumentsNavigationWizard_Documents_Click
      - OpcenterEXDS_OperatorLanding.DS_GetLayoutVerticalMode
      - OpcenterEXDS_OperatorLanding.ACT_AddDocumentsNavigationWizard_MaterialTrackingUnits_Click
      - OpcenterEXDS_OperatorLanding.ACT_AddDocumentsNavigationWizard_Submit_Click
    target_commands:
      []

  - name: SNP_QualityInspection_Visual
    module: EXFN_Quality
    microflows:
      - EXFN_Quality.ACT_ImageGrid_OnClick
      - EXFN_Quality.DS_Failure_Retrieve_ListGallery
      - EXFN_Quality.ACT_CurrentFailure_Change
      - EXFN_Quality.ACT_QualityInspectionVisual_Fullscreen
      - EXFN_Quality.ACT_NewSample_Visual
      - EXFN_Quality.ACT_ConfirmSample_Visual
      - EXFN_Quality.ACT_OpenFailureSelectionPanel
      - EXFN_Quality.DS_ScenarioInstanceView
      - EXFN_Quality.SUB_RetrieveCalcualatedJson
    target_commands:
      []

  - name: ToBeConsumedMaterialsGridRangePartsAsRequired
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_GetToBeConsumeMaterial_GridMode_RangeParts
      - OpcenterEXDS_OperatorLanding.ACT_ValidateMaterialConsumption
      - OpcenterEXDS_OperatorLanding.ACT_AcquireMTUList
      - OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialToBeConsumedPage
      - OpcenterEXDS_OperatorLanding.ACT_ConsumeMaterial
      - OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialConsumedPage_Grid
      - OpcenterEXDS_OperatorLanding.ACT_DisassembleMaterial_Grid
      - OpcenterEXDS_OperatorLanding.DS_GetToBeConsumeMaterial_GridMode_AsRequired
    target_commands:
      []

  - name: PANEL_ImportDocument
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_NonConformancesInfo_OnOperationDetailsChange
      - OpcenterEXDS_OperatorLanding.ACT_CreateAndLinkDocument
    target_commands:
      []

  - name: PANEL_OperatorLandingStart
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_StartPanel_Operation
      - OpcenterEXDS_OperatorLanding.ACT_StartAllSNButton
      - OpcenterEXDS_OperatorLanding.DS_SerialNumber_GetNotAlreadyAssociated
      - OpcenterEXDS_OperatorLanding.ACT_AssociateNewSerialNumber
      - OpcenterEXDS_OperatorLanding.OCH_SerialNumber_CheckIfExists
      - OpcenterEXDS_OperatorLanding.OCH_FlexSerialized_Quantity
      - OpcenterEXDS_OperatorLanding.ACT_AddNewSerialNumbersFromNId
      - OpcenterEXDS_OperatorLanding.ACT_AddNewSerialNumbersFromNumber
      - OpcenterEXDS_OperatorLanding.DS_StartCompletePanel_GetSN
      - OpcenterEXDS_OperatorLanding.ACT_SelectedSerialNumberInStartCompletePanel
      - OpcenterEXDS_OperatorLanding.DS_StartCompletePanel_GetEquipments
      - OpcenterEXDS_OperatorLanding.ACT_Select_Equipment_Operation
    target_commands:
      []

  - name: WorkInstruction_View
    module: EXFN_WorkInstruction
    microflows:
      []
    target_commands:
      []

  - name: ToBeCoByProducedMaterials
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_GetOrCreateSignalEnvelope_ToBeCoByProduct
      - EXFN_Authentication.Signal_Access_Token
      - OpcenterEXDS_OperatorLanding.EVT_OnCoProductProduced
      - OpcenterEXDS_OperatorLanding.EVT_OnByProductProduced
      - OpcenterEXDS_OperatorLanding.EVT_OnOutputMaterialProduced
      - EXFN_Authentication.HandleUnauthorizedBehavior
      - OpcenterEXDS_OperatorLanding.ACT_ToBeConsumedMaterialHistory
      - OpcenterEXDS_OperatorLanding.DS_GetToBeCoProducedMaterial_Grid
      - OpcenterEXDS_OperatorLanding.ACT_ValidateMaterialConsumption
      - OpcenterEXDS_OperatorLanding.ACT_CoProduceMaterial
      - OpcenterEXDS_OperatorLanding.DS_ChangeAssembleAllVisibility
      - OpcenterEXDS_OperatorLanding.ACT_CoProduceMaterialsAll
      - OpcenterEXDS_OperatorLanding.DS_CreateCoByProductContext
    target_commands:
      []

  - name: PANEL_WOOP_NonConformance
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_Open_NonConformance_PC
      - OpcenterEXDS_OperatorLanding.ACT_CreateNC_From_Panel
    target_commands:
      []

  - name: PANEL_OperatorLandingGoToDetailsSN
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_SelectedSerialNumberInPanel
      - OpcenterEXDS_OperatorLanding.DS_GetSelectedSerialNumberContext
      - OpcenterEXDS_OperatorLanding.ACT_Select_GoToDetails_SerialNumber
      - OpcenterEXDS_OperatorLanding.ACT_Select_SerialNumber
    target_commands:
      []

  - name: PANEL_OperatorLandingHold
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_Select_HoldReason
      - OpcenterEXDS_OperatorLanding.ACT_Hold_Panel_Operation
    target_commands:
      []

  - name: QualityInspectionVisual_FullscreenPopup
    module: EXFN_Quality
    microflows:
      - EXFN_Quality.DS_GetOrCreateFullscreenFailureContextualMenuContext
      - EXFN_Quality.ACT_QualityInspectionVisual_FullscreenPopup_Close
    target_commands:
      []

  - name: PANEL_Notes
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_Create_Panel_Note
      - OpcenterEXDS_OperatorLanding.ACT_Create_NoteToBeCreated_Operation
      - OpcenterEXDS_OperatorLanding.ACT_Acknowledge_Note
      - OpcenterEXDS_OperatorLanding.ACT_Create_NoteToBeCreated_WorkOrder
    target_commands:
      []

  - name: AddDocumentsPopup
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_AddDocumentsPage_RegisterClosePageAction
      - OpcenterEXDS_OperatorLanding.DS_GetLayoutHorizontalMode
      - OpcenterEXDS_OperatorLanding.ACT_Submit_AddDocuments
      - OpcenterEXDS_OperatorLanding.NAV_Back
    target_commands:
      []

  - name: ToBeConsumedMaterialsAlternative
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_GetToBeConsumeMaterial_Alternative
      - OpcenterEXDS_OperatorLanding.ACT_ActualMaterials_ToggleVisibility
      - OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialConsumedPage
      - OpcenterEXDS_OperatorLanding.ACT_DisassembleMaterial
      - OpcenterEXDS_EXFN_DocumentViewer.ACT_ShowAllPartson3dViewer
      - OpcenterEXDS_OperatorLanding.ACT_ValidateMaterialConsumption
      - OpcenterEXDS_OperatorLanding.ACT_AcquireMTUList
      - OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialToBeConsumedPage
      - OpcenterEXDS_OperatorLanding.ACT_ConsumeMaterial
    target_commands:
      []

  - name: WorkInstruction_View_Vertical
    module: EXFN_WorkInstruction
    microflows:
      []
    target_commands:
      []

  - name: Dropdown
    module: EXFN_WorkInstruction
    microflows:
      - EXFN_WorkInstruction.DS_ItemView_Dropdown
      - EXFN_WorkInstruction.ACT_ItemView_Dropdown_OnChange
      - EXFN_WorkInstruction.DS_DropdownItemList
    target_commands:
      []

  - name: PANEL_CloseFlexibleWorkOrder
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_UADMAbruptlyCloseFlexibleWorkOrder
      - OpcenterEXDS_OperatorLanding.ACT_GenerateWONId
      - OpcenterEXDS_OperatorLanding.DS_GetOrCreateCloseFlexibleContext
    target_commands:
      - AppU4DM.UADMAbruptlyCloseFlexibleWorkOrder

  - name: PANEL_CompleteAssignedNonProductiveActivities
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_SelectAssignedNonProductiveActivities
      - OpcenterEXDS_OperatorLanding.ACT_CompleteNonProductiveActivites
    target_commands:
      []

  - name: ScrapFailureBrowser
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_GetScrapFailureSortedBreadcrumb
      - OpcenterEXDS_OperatorLanding.ACT_Navigate_ScrapFailureBreadcrumb
      - OpcenterEXDS_OperatorLanding.ACT_ScrapFailureList
      - OpcenterEXDS_OperatorLanding.NAV_ScrapFailureChildren_Overview
      - OpcenterEXDS_OperatorLanding.ACT_SelectScrapFailure
    target_commands:
      []

  - name: Decimal
    module: EXFN_WorkInstruction
    microflows:
      - EXFN_WorkInstruction.DS_ItemView_Decimal
      - EXFN_WorkInstruction.ACT_ItemView_Decimal_OnChange
      - EXFN_WorkInstruction.ACT_CalculateWorkInstructionFormulaValues_Decimal
    target_commands:
      []

  - name: NonConformanceMaterialTrackingUnits_WorkOrderOperationContext
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_GetMaterialTrackingUnitsRelatedToTheSelectedWorkOrderOperation
      - OpcenterEXDS_OperatorLanding.ACT_SelectNonConformanceMaterialTrackingUnit_WorkOrderOperationContext
      - OpcenterEXDS_OperatorLanding.DS_GetLayoutHorizontalMode
    target_commands:
      []

  - name: PANEL_OperatorLandingComplete
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_SelectedSerialNumberInStartCompletePanel
      - OpcenterEXDS_OperatorLanding.DS_StartCompletePanel_GetEquipments
      - OpcenterEXDS_OperatorLanding.ACT_Select_CompleteOperationPanel_Equipment
      - OpcenterEXDS_OperatorLanding.ACT_Complete_Panel_Operation
      - OpcenterEXDS_OperatorLanding.ACT_CompleteAllSnButton
    target_commands:
      []

  - name: Document_Overview
    module: EXFN_DocumentViewer
    microflows:
      []
    target_commands:
      []

  - name: ToBeConsumedMaterialsGridNormalPart
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_GetToBeConsumeMaterial_GridMode_NormalPart
      - OpcenterEXDS_OperatorLanding.ACT_ValidateMaterialConsumption
      - OpcenterEXDS_OperatorLanding.ACT_AcquireMTUList
      - OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialToBeConsumedPage
      - OpcenterEXDS_OperatorLanding.ACT_ConsumeMaterial
      - OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialConsumedPage_Grid
      - OpcenterEXDS_OperatorLanding.ACT_DisassembleMaterial_Grid
    target_commands:
      []

  - name: ToBeConsumedMaterialsGridSelectedFit
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_GetToBeConsumeMaterial_GridMode_SelectedFit
      - OpcenterEXDS_OperatorLanding.ACT_ValidateMaterialConsumption
      - OpcenterEXDS_OperatorLanding.ACT_AcquireMTUList
      - OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialToBeConsumedPage
      - OpcenterEXDS_OperatorLanding.ACT_ConsumeMaterial
      - OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialConsumedPage_Grid
      - OpcenterEXDS_OperatorLanding.ACT_DisassembleMaterial_Grid
    target_commands:
      []

  - name: SNP_QualityInspection
    module: EXFN_Quality
    microflows:
      - EXFN_Quality.DS_GetOrCreateFailureContextualMenuContext
    target_commands:
      []

  - name: OperationContainer
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_GetBuyOffStatus
      - OpcenterEXDS_OperatorLanding.ACT_OperationContainer_SendBuyOffNotification
      - OpcenterEXDS_OperatorLanding.ACT_OperationContainer_Complete
      - OpcenterEXDS_OperatorLanding.ACT_OperationContainer_Item_Click
    target_commands:
      []

  - name: AddDocumentsSelectedMaterialTrackingUnitBadge
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_AddDocumentsDeselectMaterialTrackingUnit
      - OpcenterEXDS_OperatorLanding.DS_GetAddDocumentsSelectedMaterialTrackingUnits_Vertical
    target_commands:
      []

  - name: ToBeConsumedMaterialCoByProduced
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_GetToBeCoProductMaterial
      - OpcenterEXDS_OperatorLanding.ACT_ActualProducedMaterials_ToggleVisibility
      - OpcenterEXDS_OperatorLanding.ACT_ValidateMaterialConsumption
      - OpcenterEXDS_OperatorLanding.ACT_CoProduceMaterial
      - OpcenterEXDS_OperatorLanding.ACT_AcquireMTUList
    target_commands:
      []

  - name: PartProgramHistoryPopUp
    module: OpcenterEXDS_PartProgram
    microflows:
      - OpcenterEXDS_PartProgram.DS_RegisterClosePageAction
      - OpcenterEXDS_OperatorLanding.NAV_Back
    target_commands:
      []

  - name: NonConformancesSummary
    module: OpcenterEXDS_OperatorLanding
    microflows:
      []
    target_commands:
      []

  - name: OperatorLanding
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_CreateSignalConfiguration
      - EXFN_Authentication.Signal_Access_Token
      - OpcenterEXDS_OperatorLanding.EVT_WorkOrderOperationStartedSerialized
      - OpcenterEXDS_OperatorLanding.EVT_WorkOrderOperationStartedFullQty
      - OpcenterEXDS_OperatorLanding.EVT_WorkOrderOperationCompletedFullQty
      - OpcenterEXDS_OperatorLanding.EVT_WorkOrderOperationCompletedSerialized
      - OpcenterEXDS_OperatorLanding.EVT_WorkOrderOperationPaused
      - OpcenterEXDS_OperatorLanding.EVT_WorkOrderOperationHold
      - OpcenterEXDS_OperatorLanding.EVT_OnSentenceNonConformance
      - OpcenterEXDS_OperatorLanding.EVT_CreateOrAcknowledgeNote
      - OpcenterEXDS_OperatorLanding.EVT_WorkOrderOperationActivedCP
      - OpcenterEXDS_OperatorLanding.EVT_WorkOrderSerialNumbersScrapped
      - OpcenterEXDS_OperatorLanding.EVT_OnReopenWorkOrderOperation
      - OpcenterEXDS_OperatorLanding.SUB_RefreshOperatorLanding
      - OpcenterEXDS_OperatorLanding.EVT_CreateNonConformance
      - OpcenterEXDS_OperatorLanding.EVT_OnCompleteReworkOrder
      - OpcenterEXDS_OperatorLanding.EVT_OnFAICompleted
      - OpcenterEXDS_OperatorLanding.EVT_OnFAICandidateDeclared
      - OpcenterEXDS_OperatorLanding.EVT_OnSkipWOOperation
      - EXFN_Authentication.HandleUnauthorizedBehavior
      - OpcenterEXDS_OperatorLanding.ACT_Set_FilterCriteria_OnChanged
      - OpcenterEXDS_OperatorLanding.ACT_SearchBy_Product
      - OpcenterEXDS_OperatorLanding.ACT_SearchBy_Machine
      - OpcenterEXDS_OperatorLanding.ACT_SearchCommandBar_ClearAll_Click
      - OpcenterEXDS_OperatorLanding.DS_GetOrCreate_WorkOrderOperationHelper
      - OpcenterEXDS_OperatorLanding.NAV_Back
      - OpcenterEXDS_OperatorLanding.ACT_WorkOrderOperation_Tile_GoToDetails_Click
      - OpcenterEXDS_OperatorLanding.ACT_WorkOrderOperation_Selected
      - OpcenterEXDS_OperatorLanding.DS_GetBreadcrumb
      - OpcenterEXDS_OperatorLanding.DS_CreateDependencyGraphContext
      - OpcenterEXDS_OperatorLanding.ACT_Routing_ShowPopup
    target_commands:
      []

  - name: PartProgram
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_PartProgram.SUB_SetPartProgramContext
      - EXFN_Authentication.Signal_Access_Token
      - OpcenterEXDS_PartProgram.EVT_DNCStartTransfer
      - OpcenterEXDS_PartProgram.EVT_DNCCompleteTransfer
      - EXFN_Authentication.HandleUnauthorizedBehavior
      - OpcenterEXDS_OperatorLanding.DS_GetOrCreateSignalEnvelope_PartProgram
      - OpcenterEXDS_PartProgram.ACT_CommandTransferDefault
      - OpcenterEXDS_PartProgram.NAV_ProgramPartDetails
      - OpcenterEXDS_PartProgram.NAV_ProgramPartHistory
      - OpcenterEXDS_PartProgram.SUB_WorkOrderOperationOrStepContext_SetDncItemId
    target_commands:
      []

  - name: PANEL_Complete_ValidateUser
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_Create_ESContext
      - OpcenterEXDS_OperatorLanding.ACT_Complete_ValidateUser
      - OpcenterEXDS_OperatorLanding.ACT_Complete_ValidateUser_FromDetails
    target_commands:
      []

  - name: NonConformancesInfo
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_Select_WorkOrderOperation_Context
      - OpcenterEXDS_OperatorLanding.ACT_Select_MaterialTrackingUnit_Context
      - OpcenterEXDS_OperatorLanding.ACT_Select_Equipment
      - OpcenterEXDS_OperatorLanding.ACT_Select_Tool_Context
      - OpcenterEXDS_OperatorLanding.ACT_NonConformancesInfo_OnOperationDetailsChange
      - OpcenterEXDS_OperatorLanding.DS_Get_Severity
    target_commands:
      []

  - name: PANEL_SelectEquipmentAcquireWI
    module: OpcenterEXDS_EXFN_WorkInstruction
    microflows:
      - OpcenterEXDS_EXFN_WorkInstruction.ACT_Select_AcquireWIPanel_Equipment
      - OpcenterEXDS_EXFN_WorkInstruction.ACT_EquipmentSelectionPanel_AcquireDC
    target_commands:
      []

  - name: ToBeConsumedMaterialGridCoByProduced
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_GetToBeCoProductMaterial_GridMode
      - OpcenterEXDS_OperatorLanding.ACT_ValidateMaterialConsumption
      - OpcenterEXDS_OperatorLanding.ACT_CoProduceMaterial
      - OpcenterEXDS_OperatorLanding.ACT_AcquireMTUList
    target_commands:
      []

  - name: ScenarioInstanceView
    module: EXFN_ElectronicSignature
    microflows:
      - EXFN_ElectronicSignature.ACT_ScnInstanceView_AccordionOpen
      - EXFN_ElectronicSignature.ACT_ScnInstanceView_AccordionCollapsed
      - EXFN_ElectronicSignature.ACT_SignatureAction_Set
      - EXFN_ElectronicSignature.SWAC_OnReady
      - EXFN_ElectronicSignature.SWAC_OnFailure
      - EXFN_ElectronicSignature.EVT_AccordionStatusChanged
    target_commands:
      []

  - name: AddDocumentsSelectedDocumentBadge
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_AddDocumentsDeselectDocument
    target_commands:
      []

  - name: SelectedFailureBadge
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_DeselectFailure
    target_commands:
      []

  - name: RoutingNode
    module: OpcenterEXDS_OperatorLanding
    microflows:
      []
    target_commands:
      []

  - name: NonConformanceEquipment_EquipmentContext
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_SelectNonConformanceEquipment_EquipmentContext
      - OpcenterEXDS_OperatorLanding.DS_GetLayoutHorizontalMode
    target_commands:
      []

  - name: ToBeConsumedMaterialsCustomPart
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_GetToBeConsumeMaterial_CustomPart
      - OpcenterEXDS_OperatorLanding.ACT_ActualMaterials_ToggleVisibility
      - OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialConsumedPage
      - OpcenterEXDS_OperatorLanding.ACT_DisassembleMaterial
      - OpcenterEXDS_EXFN_DocumentViewer.ACT_ShowAllPartson3dViewer
      - OpcenterEXDS_OperatorLanding.ACT_ValidateMaterialConsumption
      - OpcenterEXDS_OperatorLanding.ACT_AcquireMTUList
      - OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialToBeConsumedPage
      - OpcenterEXDS_OperatorLanding.ACT_ConsumeMaterial
    target_commands:
      []

  - name: Failure_Overview
    module: EXFN_Quality
    microflows:
      - EXFN_Quality.ACT_Failure_Cancel
      - EXFN_Quality.ACT_Failure_Select_And_Close
      - EXFN_Quality.NAV_FirstLevelFailure_Overview
      - EXFN_Quality.NAV_FailureChildren_Overview
      - EXFN_Quality.NAV_FailureParent_Overview
      - EXFN_Quality.DS_FailureList_Retrive
      - EXFN_Quality.ACT_FailureGalleryItemSelection
      - EXFN_Quality.ACT_PotentialFailure_Retrieve
    target_commands:
      []

  - name: PANEL_SetPointHistory
    module: OpcenterEXDS_OperatorLanding
    microflows:
      []
    target_commands:
      []

  - name: PANEL_LinkExistingDocumentSet
    module: EXFN_DocumentViewer
    microflows:
      - EXFN_DocumentViewer.OCH_LinkDocumentSet_ChangeUseCurrentDocumentSetRevision
      - EXFN_DocumentViewer.ACT_LinkDocumentSet
    target_commands:
      []

  - name: DataCollector
    module: EXFN_WorkInstruction
    microflows:
      []
    target_commands:
      []

  - name: PANEL_AssignAndStartNonProductiveActivities
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_StartNonProductiveActivites
      - OpcenterEXDS_OperatorLanding.ACT_Close_CreateAndStartNonProductiveActivities
      - OpcenterEXDS_OperatorLanding.ACT_SelectNonProductiveActivities
      - OpcenterEXDS_OperatorLanding.ACT_SetContextInfo
      - OpcenterEXDS_OperatorLanding.ACT_SelectWorkOrder
      - OpcenterEXDS_OperatorLanding.ACT_SelectWorkOrderOperation
    target_commands:
      []

  - name: PANEL_ChangePackage
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_Submit_ChangePackage
    target_commands:
      []

  - name: AddDocuments
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_SearchBy_Product
      - OpcenterEXDS_OperatorLanding.ACT_AddDocumentDropdownDocType
      - OpcenterEXDS_OperatorLanding.ACT_Show_AddDocumentsImportPanel
      - OpcenterEXDS_OperatorLanding.DS_CreateFile
      - OpcenterEXDS_OperatorLanding.DS_GetOrCreate_DocSearchCommandBar_Context
      - OpcenterEXDS_OperatorLanding.DS_GetDocument
      - OpcenterEXDS_OperatorLanding.ACT_AddDocumentsSelectAndLinkDocument
      - OpcenterEXDS_EXFN_DocumentViewer.ACT_OpenDocument
      - OpcenterEXDS_OperatorLanding.DS_CreateDocumentViewer
      - OpcenterEXDS_OperatorLanding.DS_PreviewDocumentSelection
    target_commands:
      []

  - name: NonConformancePopup
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_NonConformancePage_RegisterClosePageAction
      - OpcenterEXDS_OperatorLanding.DS_GetLayoutHorizontalMode
      - OpcenterEXDS_OperatorLanding.ACT_Submit_NonConformance
      - OpcenterEXDS_OperatorLanding.NAV_Back
    target_commands:
      []

  - name: SPCChart_Variable
    module: EXFN_Quality
    microflows:
      - EXFN_Quality.SUB_RetrieveFirstMessage
    target_commands:
      []

  - name: PANEL_CreateLinkDocument
    module: EXFN_DocumentViewer
    microflows:
      - EXFN_DocumentViewer.ACT_CreateLinkDocument
    target_commands:
      []

  - name: ToBeConsumedMaterialsNormalPart
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding_Connector.DS_GetMTUWithContainer
      - OpcenterEXDS_OperatorLanding.DS_GetToBeConsumeMaterial_NormalPart
      - OpcenterEXDS_OperatorLanding.ACT_ActualMaterials_ToggleVisibility
      - OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialConsumedPage
      - OpcenterEXDS_OperatorLanding.ACT_DisassembleMaterial
      - OpcenterEXDS_EXFN_DocumentViewer.ACT_ShowAllPartson3dViewer
      - OpcenterEXDS_OperatorLanding.ACT_ValidateMaterialConsumption
      - OpcenterEXDS_OperatorLanding.ACT_AcquireMTUList
      - OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialToBeConsumedPage
      - OpcenterEXDS_OperatorLanding.ACT_ConsumeMaterial
    target_commands:
      []

  - name: ToBeProducedMaterial
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_CreateMatConsContext
      - OpcenterEXDS_OperatorLanding.ACT_Set_AutoConsume_BatchMaterials
      - OpcenterEXDS_OperatorLanding.DS_GetOrCreateSignalEnvelope_ConsumeMaterial
      - EXFN_Authentication.Signal_Access_Token
      - OpcenterEXDS_OperatorLanding.EVT_OnAssemblyWorkOrderOperation
      - OpcenterEXDS_OperatorLanding.EVT_OnDisassemblyWorkOrderOperation
      - OpcenterEXDS_OperatorLanding.EVT_ScrapMaterials_Details
      - EXFN_Authentication.HandleUnauthorizedBehavior
      - OpcenterEXDS_OperatorLanding.ACT_ConsumeMaterialsAll
      - OpcenterEXDS_EXFN_DocumentViewer.ACT_ShowAllPartson3dViewer
      - OpcenterEXDS_OperatorLanding.DS_ChangeAssembleAllVisibility
      - OpcenterEXDS_OperatorLanding_Connector.DS_CreateContainerDTO
    target_commands:
      []

  - name: ToBeUsedTools
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_GetOrCreateSignalEnvelope_UseTool
      - EXFN_Authentication.Signal_Access_Token
      - OpcenterEXDS_OperatorLanding.EVT_OnUsedToolSignal
      - OpcenterEXDS_OperatorLanding.EVT_OnScrewingDeactivation
      - OpcenterEXDS_OperatorLanding.EVT_OnScrewingActivation
      - OpcenterEXDS_OperatorLanding.EVT_OnScrewingExecuted
      - EXFN_Authentication.HandleUnauthorizedBehavior
      - OpcenterEXDS_OperatorLanding.ACT_ToolHistory_ToggleVisibility
      - OpcenterEXDS_OperatorLanding.ACT_ToolInput_ToggleVisibility
      - OpcenterEXDS_OperatorLanding.OCH_ToolNId
      - OpcenterEXDS_OperatorLanding.ACT_ShowSelectToolPanel
      - OpcenterEXDS_OperatorLanding.ACT_UseTool
      - OpcenterEXDS_OperatorLanding.ACT_AcquireTool
      - OpcenterEXDS_OperatorLanding.DS_GetToolEntity
      - OpcenterEXDS_OperatorLanding.DS_CreateScrewingToolUsage
      - OpcenterEXDS_OperatorLanding.ACT_ToggleIsActiveMaterialTrackingUnit
      - OpcenterEXDS_OperatorLanding.DS_GetScrewingToolEntity
      - OpcenterEXDS_OperatorLanding.ACT_ToBeUsedToolHistory
      - OpcenterEXDS_OperatorLanding.DS_GetTobeUsedToolHistory
      - OpcenterEXDS_OperatorLanding.ACT_ShowToolPanelForGrid
      - OpcenterEXDS_OperatorLanding.ACT_ValidateDurationToolInGridMode
      - OpcenterEXDS_OperatorLanding.ACT_Grid_AcquireTool
      - OpcenterEXDS_OperatorLanding.ACT_Grid_UseTool
      - OpcenterEXDS_OperatorLanding.ACT_UseAllTools
      - OpcenterEXDS_OperatorLanding.ACT_Grid_UseAllTools
      - OpcenterEXDS_OperatorLanding.DS_CreateToBeUsedToolContext
    target_commands:
      []

  - name: VerticalCommandBarPartProgram
    module: OpcenterEXDS_PartProgram
    microflows:
      - OpcenterEXDS_PartProgram.ACT_CommandTransfer
      - OpcenterEXDS_PartProgram.ACT_CommandDownloadAndManualTransfer
      - OpcenterEXDS_PartProgram.ACT_CommandPreview_NEW
      - OpcenterEXDS_PartProgram.ACT_Close
    target_commands:
      []

  - name: QualityInspections
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_ParameterView_Create
    target_commands:
      []

  - name: WorkInstruction_Overview
    module: EXFN_WorkInstruction
    microflows:
      - EXFN_WorkInstruction.DS_WorkInstructionView_List
      - EXFN_WorkInstruction.NAV_WorkInstruction_View
      - EXFN_WorkInstruction.NAV_WorkInstruction_View_Vertical
      - EXFN_WorkInstruction.ACT_WorkInstruction_Delete
    target_commands:
      []

  - name: DocumentViewerSnippet
    module: EXFN_DocumentViewer
    microflows:
      - EXFN_DocumentViewer.ACT_GetDocumentSelection
      - EXFN_DocumentViewer.ACT_SetPMIVisibility
      - EXFN_DocumentViewer.ACT_FitToView
      - EXFN_DocumentViewer.ACT_OpenDocument
      - EXFN_DocumentViewer.ACT_ChangeDetailsView
      - EXFN_DocumentViewer.DS_GetFile
      - EXFN_ServiceLayer.DS_GetApplicationURL
      - EXFN_DocumentViewer.ACT_ChangeModel3D
      - EXFN_DocumentViewer.ACT_CloseJTFullScreen
      - EXFN_DocumentViewer.ACT_GenerateMarkupSnapshot
      - EXFN_DocumentViewer.ToggleMarkup
      - EXFN_DocumentViewer.ACT_ToggleToolbox
      - EXFN_DocumentViewer.ACT_SetMarkupColor
      - EXFN_DocumentViewer.EVT_OnSelectedObject
      - EXFN_DocumentViewer.EVT_LoadJT
      - EXFN_DocumentViewer.DS_GetCategories
      - EXFN_DocumentViewer.DS_GetDocumentByCategory
      - EXFN_DocumentViewer.ACT_Carousel
      - EXFN_DocumentViewer.ACT_ChangeDocumentSelection
      - EXFN_DocumentViewer.ACT_ButtonDownload
      - EXFN_DocumentViewer.DS_GetDocumentSetList
    target_commands:
      []

  - name: ToBeConsumedMaterialsReference
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_GetToBeConsumeMaterial_Reference
    target_commands:
      []

  - name: PANEL_SelectTool
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_ToggleToolSelection
      - OpcenterEXDS_OperatorLanding.ACT_SelectTool
    target_commands:
      []

  - name: ScrapMaterialTrackingUnits
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_GetMaterialTrackingUnitsActive
      - OpcenterEXDS_OperatorLanding.ACT_SelectScrapMTU
    target_commands:
      []

  - name: ES_Button
    module: EXFN_ElectronicSignature
    microflows:
      - EXFN_ElectronicSignature.ACT_Show_ScenarioInstanceView
    target_commands:
      []

  - name: SPCResponseMessages
    module: EXFN_Quality
    microflows:
      []
    target_commands:
      []

  - name: PANEL_DisassembleMaterialTrackingUnit
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_DisassembleMaterialWithNotes
    target_commands:
      []

  - name: SelectedDocumentBadge
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_DeselectDocument
    target_commands:
      []

  - name: PANEL_OperatorDetailsStartStep
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_Select_Equipment_Step
      - OpcenterEXDS_OperatorLanding.ACT_StartPanel_Step
    target_commands:
      []

  - name: ToBeConsumedMaterialsGridReference
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_GetToBeConsumeMaterial_Reference
    target_commands:
      []

  - name: PANEL_Start_ValidateUser
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_Start_ValidateUser
      - OpcenterEXDS_OperatorLanding.DS_Create_ESContext
    target_commands:
      []

  - name: MultipleChoice
    module: EXFN_WorkInstruction
    microflows:
      - EXFN_WorkInstruction.ACT_ItemView_MultipleChoice_OnChange
      - EXFN_WorkInstruction.DS_MultipleChoiceItem
      - EXFN_WorkInstruction.DS_ItemView_MultipleChoice
    target_commands:
      []

  - name: WorkInstruction_New_Step2
    module: EXFN_WorkInstruction
    microflows:
      - EXFN_WorkInstruction.ACT_WorkInstruction_Create
      - EXFN_WorkInstruction.ACT_WorkInstructionView_New_MF
    target_commands:
      []

  - name: PANEL_VerticalCommandBarNC
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_ClosePanel_VerticalCommandBarMore
    target_commands:
      []

  - name: PartProgramDetailsPopUp
    module: OpcenterEXDS_PartProgram
    microflows:
      - OpcenterEXDS_PartProgram.DS_RegisterClosePageAction
      - OpcenterEXDS_OperatorLanding.NAV_Back
      - OpcenterEXDS_PartProgram.DS_CreateSignalConfiguration
      - EXFN_Authentication.Signal_Access_Token
      - OpcenterEXDS_PartProgram.EVT_DNCStartTransfer
      - OpcenterEXDS_PartProgram.EVT_DNCCompleteTransfer
      - EXFN_Authentication.HandleUnauthorizedBehavior
    target_commands:
      []

  - name: DocumentSelection_View
    module: EXFN_DocumentViewer
    microflows:
      - EXFN_DocumentViewer.DS_GetFile
    target_commands:
      []

  - name: ToBeConsumedMaterialGridAlternative
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_GetToBeConsumeMaterial_GridMode_Alternative
      - OpcenterEXDS_OperatorLanding.ACT_ValidateMaterialConsumption
      - OpcenterEXDS_OperatorLanding.ACT_AcquireMTUList
      - OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialToBeConsumedPage
      - OpcenterEXDS_OperatorLanding.ACT_ConsumeMaterial
      - OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialConsumedPage_Grid
      - OpcenterEXDS_OperatorLanding.ACT_DisassembleMaterial_Grid
    target_commands:
      []

  - name: VerticalCommandBar_NC
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_NonConformance_Operation
      - OpcenterEXDS_OperatorLanding.ACT_OpenNonConformanceList
      - OpcenterEXDS_OperatorLanding.DS_GetOrCreateVerticalCommandBarContext
    target_commands:
      []

  - name: PANEL_OperatorLandingPause
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_Pause_Panel_Operation
      - OpcenterEXDS_OperatorLanding.ACT_Select_PauseReason
      - OpcenterEXDS_OperatorLanding.ACT_Select_PausePanel_Equipment
    target_commands:
      []

  - name: ToBeConsumedMaterialsPrekit
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_GetToBeConsumeMaterial_Prekit
      - OpcenterEXDS_OperatorLanding.DS_TogglePrekitEdit
      - OpcenterEXDS_OperatorLanding.ACT_ValidatePrekitCode
      - OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialToBeConsumedPage
      - OpcenterEXDS_OperatorLanding.ACT_ConsumeMaterial_Prekitted
      - OpcenterEXDS_OperatorLanding.ACT_EditPrekit
      - OpcenterEXDS_EXFN_DocumentViewer.ACT_ShowAllPartson3dViewer
    target_commands:
      []

  - name: OperationContainerHeaderBar
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.SUB_SetSkillsByWorkOrderOperationOrStep
      - OpcenterEXDS_OperatorLanding.DS_ChangeWorkOrderOpOrStepContext_WIOnDemand
      - OpcenterEXDS_OperatorLanding.ACT_LinkWorkInstructionsToWOOperation_ShowPanel
      - OpcenterEXDS_OperatorLanding.ACT_LinkWorkInstructionsToWOStep_ShowPanel
      - OpcenterEXDS_OperatorLanding.DS_GetBuyOffStatus
      - OpcenterEXDS_OperatorLanding.ACT_OperationContainer_SendBuyOffNotification
      - OpcenterEXDS_OperatorLanding.ACT_OperationContainer_Fullscreen
      - OpcenterEXDS_OperatorLanding.ACT_OperationContainer_ExitFullscreen
      - OpcenterEXDS_OperatorLanding.DS_SetAllowedRuntimeActionProperties_ToWorkOrderOperationOrStepContext
      - OpcenterEXDS_OperatorLanding.ACT_OperationContainer_Start
      - OpcenterEXDS_OperatorLanding.ACT_QuickStart
    target_commands:
      []

  - name: WorkInstruction_VerticalView_Example
    module: EXFN_WorkInstruction
    microflows:
      []
    target_commands:
      []

  - name: PANEL_LinkDocuments
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_SearchBy_Product
      - OpcenterEXDS_OperatorLanding.ACT_DropdownDocType
      - OpcenterEXDS_OperatorLanding.DS_GetOrCreate_DocSearchCommandBar_Context
      - OpcenterEXDS_OperatorLanding.ACT_LinkDocument
      - OpcenterEXDS_OperatorLanding.DS_GetDocument
      - OpcenterEXDS_EXFN_DocumentViewer.ACT_OpenDocument
      - OpcenterEXDS_OperatorLanding.DS_CreateDocumentViewer
      - OpcenterEXDS_OperatorLanding.DS_PreviewDocumentSelection
      - OpcenterEXDS_OperatorLanding.ACT_MultiSelection_onLinkDocPanel
    target_commands:
      []

  - name: OperatorTerminal
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_Get_WOOId_First
      - EXFN_Authentication.Signal_Access_Token
      - OpcenterEXDS_OperatorLanding.EVT_WorkOrderOperationStartedSerialized_Details
      - OpcenterEXDS_OperatorLanding.EVT_WorkOrderOperationStartedFullQty_Details
      - OpcenterEXDS_OperatorLanding.EVT_WorkOrderOperationCompletedFullQty_Details
      - OpcenterEXDS_OperatorLanding.EVT_WorkOrderOperationCompletedSerialized_Details
      - OpcenterEXDS_OperatorLanding.EVT_WorkOrderOperationPaused_Details
      - OpcenterEXDS_OperatorLanding.EVT_WorkOrderStepStartedSerialized_Details
      - OpcenterEXDS_OperatorLanding.EVT_WorkOrderStepStartedFullQty_Details
      - OpcenterEXDS_OperatorLanding.EVT_WorkOrderStepCompletedSerialized_Details
      - OpcenterEXDS_OperatorLanding.EVT_WorkOrderStepCompletedFullQty_Details
      - OpcenterEXDS_OperatorLanding.EVT_WorkOrderOperationHold_Details
      - OpcenterEXDS_OperatorLanding.EVT_OnSentenceNonConformance_Details
      - OpcenterEXDS_OperatorLanding.EVT_CreateOrAcknowledgeNote_Details
      - OpcenterEXDS_OperatorLanding.EVT_RefreshActivityIndicatorOnAssemblyWorkOrderOperation
      - OpcenterEXDS_OperatorLanding.EVT_RefreshActivityIndicatorOnDisassemblyWorkOrderOperation
      - OpcenterEXDS_OperatorLanding.EVT_RefreshActivityIndicatorOnUsedTool
      - OpcenterEXDS_OperatorLanding.EVT_RefreshActivityIndicatorOnWIStatusChangedSignal
      - OpcenterEXDS_OperatorLanding.EVT_WorkOrderOperationActivedCP_Details
      - OpcenterEXDS_OperatorLanding.EVT_WorkOrderSerialNumbersScrapped_Details
      - OpcenterEXDS_OperatorLanding.EVT_OnReopenWorkOrderOperation_Details
      - OpcenterEXDS_OperatorLanding.EVT_OnCreateNonConformance_Details
      - OpcenterEXDS_OperatorLanding.EVT_OnCompleteReworkOrder_Details
      - OpcenterEXDS_OperatorLanding.EVT_OnFAIRemoved_Details
      - OpcenterEXDS_OperatorLanding.EVT_OnFAICompleted_Details
      - OpcenterEXDS_OperatorLanding.EVT_OnFAICandidateDeclared_Details
      - OpcenterEXDS_OperatorLanding.EVT_OnSkipWOOperation_Details
      - EXFN_Authentication.HandleUnauthorizedBehavior
      - OpcenterEXDS_OperatorLanding.DS_CreateSignalConfiguration
      - OpcenterEXDS_OperatorLanding.DS_FAIWoBadgeVisibility
      - OpcenterEXDS_OperatorLanding.DS_Gallery_OperationContainer_GetWorkOrderOperationContextList
      - OpcenterEXDS_OperatorLanding.DS_GetAssociatedDocuments_MF
      - OpcenterEXDS_OperatorLanding.NAV_Back
    target_commands:
      []

  - name: ReadMe
    module: EXFN_Quality
    microflows:
      []
    target_commands:
      []

  - name: PANEL_ScrewingDetails
    module: OpcenterEXDS_OperatorLanding
    microflows:
      []
    target_commands:
      []

  - name: PANEL_ChangeSN
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_UpdateSerialNumberList
      - OpcenterEXDS_OperatorLanding.ACT_GenerateNewSerialNumber
    target_commands:
      - AppU4DM.UpdateSerialNumberList

  - name: DocumentViewer_ForLinkedEntity
    module: EXFN_DocumentViewer
    microflows:
      - EXFN_DocumentViewer.DS_GetDocumentsLinkedToEntity
      - EXFN_DocumentViewer.ACT_UnlinkDocument
    target_commands:
      []

  - name: ToBeConsumedMaterialsRangePartsAndAsRequired
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_GetToBeConsumeMaterial_RangePartsAndAsRequired
      - OpcenterEXDS_OperatorLanding.ACT_ActualMaterials_ToggleVisibility
      - OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialConsumedPage
      - OpcenterEXDS_OperatorLanding.ACT_DisassembleMaterial
      - OpcenterEXDS_EXFN_DocumentViewer.ACT_ShowAllPartson3dViewer
      - OpcenterEXDS_OperatorLanding.ACT_ValidateMaterialConsumption
      - OpcenterEXDS_OperatorLanding.ACT_AcquireMTUList
      - OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialToBeConsumedPage
      - OpcenterEXDS_OperatorLanding.ACT_ConsumeMaterial
    target_commands:
      []

  - name: PANEL_AddDocumentsImportDocument
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_CancelImportDocument
      - OpcenterEXDS_OperatorLanding.ACT_AddDocumentsCreateAndLinkDocument
    target_commands:
      []

  - name: PANEL_WorkInstructionsToWOOperation_Add
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_LinkWorkInstructionsToWOOperation
    target_commands:
      []

  - name: PANEL_SelectDestinationContainer
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_SelectDestinationContainer
      - OpcenterEXDS_OperatorLanding.ACT_Cancel_SelectDestinationContainerPanel
      - OpcenterEXDS_OperatorLanding.ACT_Save_SelectDestinationContainerPanel
    target_commands:
      []

  - name: PANEL_Pause_ValidateUser
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_Pause_ValidateUser
      - OpcenterEXDS_OperatorLanding.DS_Create_ESContext
    target_commands:
      []

  - name: ScrapProducedMaterialPopup
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_ScrapPage_RegisterClosePageAction
      - OpcenterEXDS_OperatorLanding.ACT_ScrapWorkOrderSerialNumbers
      - OpcenterEXDS_OperatorLanding.NAV_Back
      - OpcenterEXDS_OperatorLanding.DS_GetSelectedMaterialTrackingUnitsForScrap
      - OpcenterEXDS_OperatorLanding.ACT_DeselectScrapMTU
      - OpcenterEXDS_OperatorLanding.DS_GetSelectedFailuresForScrap
      - OpcenterEXDS_OperatorLanding.ACT_DeselectScrapFailure
    target_commands:
      []

  - name: NonConformanceDocuments
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.DS_GetDocumentList
      - OpcenterEXDS_OperatorLanding.DS_CreateDocumentSelection
      - OpcenterEXDS_OperatorLanding.DS_CreateDocumentViewer
      - OpcenterEXDS_OperatorLanding.ACT_SelectDocument
      - OpcenterEXDS_EXFN_DocumentViewer.ACT_OpenDocument
      - OpcenterEXDS_OperatorLanding.DS_GetDocument
      - OpcenterEXDS_OperatorLanding.ACT_ShowDocumentsPanel
      - OpcenterEXDS_OperatorLanding.ACT_Show_ImportPanel
    target_commands:
      []

  - name: Routing_Popup
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.SUB_DependencyGraphContext_RegisterClosePageAction
      - OpcenterEXDS_OperatorLanding.DS_GetBreadcrumb
      - OpcenterEXDS_OperatorLanding.NAV_Back
    target_commands:
      []

  - name: SNP_QualityInspectionContainer
    module: EXFN_Quality
    microflows:
      - EXFN_Quality.ACT_CreateSignalConfiguration
      - EXFN_Quality.ACT_RuntimeInspectionDefinition_Create
      - EXFN_Authentication.Signal_Access_Token
      - EXFN_Quality.EVT_TerminalDelayedExecution
      - EXFN_Quality.EVT_TerminalInspectionExecutionChrReprRuntimeNumberChanged
      - EXFN_Quality.EVT_CompletedInspectionSampleSignIn
      - EXFN_Quality.EVT_InspectionSampleConfirmed
      - EXFN_Quality.SUB_InspectionEngine_Retrieve
    target_commands:
      []

  - name: SNP_QualityInspection_Variable
    module: EXFN_Quality
    microflows:
      - EXFN_Quality.ACT_DrawBar
      - EXFN_Quality.ACT_Variable_OnChange
      - EXFN_Quality.ACT_RefreshVariableView
      - EXFN_Quality.ACT_ShowPmiByCharName
      - EXFN_Quality.ACT_DeselectValueToUpdate
      - EXFN_Quality.ACT_Variable_DisassociateFailure
      - EXFN_Quality.ACT_SelectValueToUpdate
      - EXFN_Quality.ACT_SetInspectionValueAndOpenFailure
      - EXFN_Quality.DS_ScenarioInstanceView
      - EXFN_Quality.ACT_ConfirmSample_Variable
      - EXFN_Quality.ACT_NewSample_Variable
      - EXFN_Quality.SUB_RetrieveCalcualatedJson
    target_commands:
      []

  - name: NonConformancesNavigationWizard
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_NonConformancesNavigationWizard_Details_Click
      - OpcenterEXDS_OperatorLanding.DS_GetLayoutVerticalMode
      - OpcenterEXDS_OperatorLanding.ACT_NonConformancesNavigationWizard_Failure_Click
      - OpcenterEXDS_OperatorLanding.ACT_NonConformancesNavigationWizard_Documents_Click
      - OpcenterEXDS_OperatorLanding.ACT_NonConformancesNavigationWizard_Submit_Click
    target_commands:
      []

  - name: Checkbox
    module: EXFN_WorkInstruction
    microflows:
      - EXFN_WorkInstruction.ACT_ItemView_Checkbox_OnChange_Visibility
      - EXFN_WorkInstruction.ACT_ItemView_Checkbox_OnChange_Value
      - EXFN_WorkInstruction.ACT_CalculateWorkInstructionFormulaValues_Checkbox
      - EXFN_WorkInstruction.DS_ItemView_Checkbox
    target_commands:
      []

  - name: WorkInstruction_Overview_Example
    module: EXFN_WorkInstruction
    microflows:
      - EXFN_WorkInstruction.DS_WorkInstructionView_List
      - EXFN_WorkInstruction.NAV_WorkInstruction_VerticalView_Example
    target_commands:
      []

  - name: OperationList
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_Collapse_OperationList
      - OpcenterEXDS_OperatorLanding.ACT_Expand_OperationList
      - OpcenterEXDS_OperatorLanding.DS_GetSelectedWorkOrderOperationsAndSteps
      - OpcenterEXDS_OperatorLanding.ACT_ShowSteps_OperationList
      - OpcenterEXDS_OperatorLanding.ACT_HideSteps_OperationList
      - OpcenterEXDS_OperatorLanding.ACT_OperationList_Icon_Click
      - OpcenterEXDS_OperatorLanding.ACT_ShowRouting
      - OpcenterEXDS_OperatorLanding.ACT_Switch_ShowCompletedOperation
      - OpcenterEXDS_OperatorLanding.ACT_Switch_HideCompletedOperation
      - OpcenterEXDS_OperatorLanding.ACT_OperationList_Item_Click
    target_commands:
      []

  - name: SelectedEquipmentBadge_EquipmentContext
    module: OpcenterEXDS_OperatorLanding
    microflows:
      - OpcenterEXDS_OperatorLanding.ACT_DeselectEquipment_EquipmentContext
      - OpcenterEXDS_OperatorLanding.DS_GetSelectedEquipments_Vertical_EquipmentContext
    target_commands:
      []

  - name: WorkInstruction_New_Step1
    module: EXFN_WorkInstruction
    microflows:
      []
    target_commands:
      []

  - name: Text
    module: EXFN_WorkInstruction
    microflows:
      - EXFN_WorkInstruction.DS_ItemView_Text
      - EXFN_WorkInstruction.ACT_ItemView_Text_OnChange
      - EXFN_WorkInstruction.ACT_CalculateWorkInstructionFormulaValues_Text
    target_commands:
      []

  - name: WorkInstruction_Preview
    module: EXFN_WorkInstruction
    microflows:
      - EXFN_WorkInstruction.DS_WI_SortedItemViewList
      - EXFN_WorkInstruction.DS_StepView
      - EXFN_WorkInstruction.ACT_StepNavigation_First
      - EXFN_WorkInstruction.ACT_StepNavigation_Previous
      - EXFN_WorkInstruction.DS_Step_Sequence
      - EXFN_WorkInstruction.ACT_StepNavigation_Set
      - EXFN_WorkInstruction.ACT_StepNavigation_Next
      - EXFN_WorkInstruction.ACT_StepNavigation_Last
    target_commands:
      []

  - name: Multiline
    module: EXFN_WorkInstruction
    microflows:
      - EXFN_WorkInstruction.DS_ItemView_Multiline
      - EXFN_WorkInstruction.ACT_ItemView_Multiline_OnChange
      - EXFN_WorkInstruction.ACT_CalculateWorkInstructionFormulaValues_Multiline
    target_commands:
      []

```

---

_Report generated by export_manifest tool_
