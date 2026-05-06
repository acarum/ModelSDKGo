# Manifest Report: Opcenter EX DS Complex Manufacturing

**Mendix Version:** 11.6.4  
**MPR File:** C:\Workspaces\Mendix\Complex\main\Opcenter EX DS Complex Manufacturing.mpr  
**Generated:** 2026-05-06 09:37:46  

---

## Summary

- **External Entities:** 87 (across 6 modules)
- **Microflow/Action Calls:** 82
- **Pages/Panels with Microflow Calls:** 138
- **Signal Manager Subscriptions:** 64 subscription(s)
- **Navigation Items:** 3

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
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnSentenceNonConformanceV3_1 | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | StatusChanged | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | WorkOrderSerialNumbersScrapped | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnReopenWorkOrderOperation | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnStartWorkOrderOperationSerialized | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnChangeFlexibleWOOpStatusToComplete | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnCreateNonConformanceV3_1 | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnFAICompleted | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnSkipWOOperation | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnStartWorkOrderOperationFullQty | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | SnagAndNoteNotificationSgn | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnPauseWorkOrderOperation | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnCompleteReworkOrder | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnFAICandidateDeclared | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnCompleteWorkOrderOperationFullQty | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnCompleteWorkOrderOperationSerialized | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnSetHoldWorkOrder | AppU4DM | No |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnSetHoldWorkOrder | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnDisassemblyWorkOrderOperation | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnReopenWorkOrderOperation | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnFAICompleted | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnFAICandidateDeclared | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnSkipWOOperation | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnCompleteWorkOrderOperationFullQty | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnCompleteWorkOrderOperationSerialized | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnStartWorkOrderStepSerialized | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnSentenceNonConformanceV3_1 | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | StatusChanged | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | WorkOrderSerialNumbersScrapped | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnCreateNonConformanceV3_1 | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnCompleteReworkOrder | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnAssemblyWorkOrderOperation | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnStartWorkOrderOperationFullQty | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnPauseWorkOrderOperation | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnStartWorkOrderStepFullQty | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnCompleteWorkOrderStepFullQty | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | SnagAndNoteNotificationSgn | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnUsedToolSignal | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | WorkInstructionStatusChangedSignal | WorkInstruction | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnFAIRemoved | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnStartWorkOrderOperationSerialized | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnCompleteWorkOrderStepSerialized | AppU4DM | Yes |
| OpcenterEXDS_EXFN_WorkInstruction | Snippet | SNP_WorkInstruction_VerticalView | WorkInstructionSectionCompletedSignal | WorkInstruction | Yes |
| OpcenterEXDS_EXFN_WorkInstruction | Snippet | SNP_WorkInstruction_VerticalView | WorkInstructionStepCompletedSignal | WorkInstruction | Yes |
| OpcenterEXDS_OperatorLanding | Snippet | WorkInstructions | OnLinkWIOnDemandToSerialNumber | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Snippet | ToBeUsedTools_Backup | OnUsedToolSignal | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Snippet | ToBeCoByProducedMaterials | ByProductProduced | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Snippet | ToBeCoByProducedMaterials | OutputMaterialProduced | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Snippet | ToBeCoByProducedMaterials | CoProductProduced | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Snippet | PartProgram | DNCStartTransferEvent | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Snippet | PartProgram | DNCCompleteTransferEvent | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Snippet | ToBeProducedMaterial | OnAssemblyWorkOrderOperation | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Snippet | ToBeProducedMaterial | OnDisassemblyWorkOrderOperation | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Snippet | ToBeProducedMaterial | ScrapMaterials | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Snippet | ToBeUsedTools | MaterialTrackingUnitActivationEvent | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Snippet | ToBeUsedTools | ScrewingExecuted | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Snippet | ToBeUsedTools | OnUsedToolSignal | AppU4DM | Yes |
| OpcenterEXDS_OperatorLanding | Snippet | ToBeUsedTools | MaterialTrackingUnitDeactivationEvent | AppU4DM | Yes |
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

Found 138 page(s)/panel(s) with microflow calls:

| Name | Module | Target Commands |
|------|--------|----------------|
| DocumentSelection_View | EXFN_DocumentViewer | EXFN_DocumentViewer.DS_GetFile |
| DocumentViewerSnippet | EXFN_DocumentViewer | EXFN_DocumentViewer.ACT_ButtonDownload<br>EXFN_DocumentViewer.ACT_Carousel<br>EXFN_DocumentViewer.ACT_ChangeDetailsView<br>EXFN_DocumentViewer.ACT_ChangeDocumentSelection<br>EXFN_DocumentViewer.ACT_ChangeModel3D<br>EXFN_DocumentViewer.ACT_CloseJTFullScreen<br>EXFN_DocumentViewer.ACT_FitToView<br>EXFN_DocumentViewer.ACT_GenerateMarkupSnapshot<br>EXFN_DocumentViewer.ACT_GetDocumentSelection<br>EXFN_DocumentViewer.ACT_OpenDocument<br>EXFN_DocumentViewer.ACT_SetMarkupColor<br>EXFN_DocumentViewer.ACT_SetPMIVisibility<br>EXFN_DocumentViewer.ACT_ToggleToolbox<br>EXFN_DocumentViewer.DS_GetCategories<br>EXFN_DocumentViewer.DS_GetDocumentByCategory<br>EXFN_DocumentViewer.DS_GetDocumentSetList<br>EXFN_DocumentViewer.DS_GetFile<br>EXFN_DocumentViewer.EVT_LoadJT<br>EXFN_DocumentViewer.EVT_OnSelectedObject<br>EXFN_DocumentViewer.ToggleMarkup<br>EXFN_ServiceLayer.DS_GetApplicationURL |
| DocumentViewer_ForDocSetLinkedEntity | EXFN_DocumentViewer | EXFN_DocumentViewer.ACT_UnlinkDocumentSetList<br>EXFN_DocumentViewer.DS_GetDocumentsByDocumentSetsLinkedToEntity |
| DocumentViewer_ForLinkedEntity | EXFN_DocumentViewer | EXFN_DocumentViewer.ACT_UnlinkDocument<br>EXFN_DocumentViewer.DS_GetDocumentsLinkedToEntity |
| PANEL_CreateLinkDocument | EXFN_DocumentViewer | EXFN_DocumentViewer.ACT_CreateLinkDocument |
| PANEL_LinkExistingDocument | EXFN_DocumentViewer | EXFN_DocumentViewer.ACT_LinkDocument<br>EXFN_DocumentViewer.ACT_LinkDocument_ChangeUseCurrentDocumentRevision |
| PANEL_LinkExistingDocumentSet | EXFN_DocumentViewer | EXFN_DocumentViewer.ACT_LinkDocumentSet<br>EXFN_DocumentViewer.OCH_LinkDocumentSet_ChangeUseCurrentDocumentSetRevision |
| ES_Button | EXFN_ElectronicSignature | EXFN_ElectronicSignature.ACT_Show_ScenarioInstanceView |
| ScenarioInstanceView | EXFN_ElectronicSignature | EXFN_ElectronicSignature.ACT_ScnInstanceView_AccordionCollapsed<br>EXFN_ElectronicSignature.ACT_ScnInstanceView_AccordionOpen<br>EXFN_ElectronicSignature.ACT_SignatureAction_Set<br>EXFN_ElectronicSignature.EVT_AccordionStatusChanged<br>EXFN_ElectronicSignature.SWAC_OnFailure<br>EXFN_ElectronicSignature.SWAC_OnReady |
| FailureContextualMenu | EXFN_Quality | EXFN_Quality.ACT_ContextualMenu_OnMouseLeave<br>EXFN_Quality.ACT_QualityInspectionVisual_ContextualMenu_SelectionChanged<br>EXFN_Quality.DS_QualityInspectionVisual_ContextualMenu_GetLimitedFailureList_WithOthersOption |
| Failure_Overview | EXFN_Quality | EXFN_Quality.ACT_FailureGalleryItemSelection<br>EXFN_Quality.ACT_Failure_Cancel<br>EXFN_Quality.ACT_Failure_Select_And_Close<br>EXFN_Quality.ACT_PotentialFailure_Retrieve<br>EXFN_Quality.DS_FailureList_Retrive<br>EXFN_Quality.NAV_FailureChildren_Overview<br>EXFN_Quality.NAV_FailureParent_Overview<br>EXFN_Quality.NAV_FirstLevelFailure_Overview |
| QualityInspectionVisual_FullscreenPopup | EXFN_Quality | EXFN_Quality.ACT_QualityInspectionVisual_FullscreenPopup_Close<br>EXFN_Quality.DS_GetOrCreateFullscreenFailureContextualMenuContext |
| SNP_QualityInspection | EXFN_Quality | EXFN_Quality.DS_GetOrCreateFailureContextualMenuContext |
| SNP_QualityInspectionContainer | EXFN_Quality | EXFN_Authentication.Signal_Access_Token<br>EXFN_Quality.ACT_CreateSignalConfiguration<br>EXFN_Quality.ACT_RuntimeInspectionDefinition_Create<br>EXFN_Quality.EVT_CompletedInspectionSampleSignIn<br>EXFN_Quality.EVT_InspectionSampleConfirmed<br>EXFN_Quality.EVT_TerminalDelayedExecution<br>EXFN_Quality.EVT_TerminalInspectionExecutionChrReprRuntimeNumberChanged<br>EXFN_Quality.SUB_InspectionEngine_Retrieve |
| SNP_QualityInspection_Attributive | EXFN_Quality | EXFN_Quality.ACT_Attributive_DisassociateFailure<br>EXFN_Quality.ACT_Attributive_OnChange<br>EXFN_Quality.ACT_ConfirmSample_Attributive<br>EXFN_Quality.ACT_DeselectAttributeValueToUpdate<br>EXFN_Quality.ACT_GetAttributiveViewForHistory<br>EXFN_Quality.ACT_NewSample_Attributive<br>EXFN_Quality.ACT_SelectAttributiveValueToUpdate<br>EXFN_Quality.ACT_SetInspectionValueAndOpenFailure<br>EXFN_Quality.ACT_SetNOkValue<br>EXFN_Quality.ACT_SetOkValue<br>EXFN_Quality.ACT_ShowPmiByCharName<br>EXFN_Quality.DS_ScenarioInstanceView<br>EXFN_Quality.SUB_RetrieveCalcualatedJson |
| SNP_QualityInspection_Variable | EXFN_Quality | EXFN_Quality.ACT_ConfirmSample_Variable<br>EXFN_Quality.ACT_DeselectValueToUpdate<br>EXFN_Quality.ACT_DrawBar<br>EXFN_Quality.ACT_NewSample_Variable<br>EXFN_Quality.ACT_RefreshVariableView<br>EXFN_Quality.ACT_SelectValueToUpdate<br>EXFN_Quality.ACT_SetInspectionValueAndOpenFailure<br>EXFN_Quality.ACT_ShowPmiByCharName<br>EXFN_Quality.ACT_Variable_DisassociateFailure<br>EXFN_Quality.ACT_Variable_OnChange<br>EXFN_Quality.DS_ScenarioInstanceView<br>EXFN_Quality.SUB_RetrieveCalcualatedJson |
| SNP_QualityInspection_Visual | EXFN_Quality | EXFN_Quality.ACT_ConfirmSample_Visual<br>EXFN_Quality.ACT_CurrentFailure_Change<br>EXFN_Quality.ACT_ImageGrid_OnClick<br>EXFN_Quality.ACT_NewSample_Visual<br>EXFN_Quality.ACT_OpenFailureSelectionPanel<br>EXFN_Quality.ACT_QualityInspectionVisual_Fullscreen<br>EXFN_Quality.DS_Failure_Retrieve_ListGallery<br>EXFN_Quality.DS_ScenarioInstanceView<br>EXFN_Quality.SUB_RetrieveCalcualatedJson |
| SPCChart_Variable | EXFN_Quality | EXFN_Quality.SUB_RetrieveFirstMessage |
| SPCChart_VisualAndAttributive | EXFN_Quality | EXFN_Quality.SUB_RetrieveFirstMessage |
| Checkbox | EXFN_WorkInstruction | EXFN_WorkInstruction.ACT_CalculateWorkInstructionFormulaValues_Checkbox<br>EXFN_WorkInstruction.ACT_ItemView_Checkbox_OnChange_Value<br>EXFN_WorkInstruction.ACT_ItemView_Checkbox_OnChange_Visibility<br>EXFN_WorkInstruction.DS_ItemView_Checkbox |
| Datatime | EXFN_WorkInstruction | EXFN_WorkInstruction.ACT_CalculateWorkInstructionFormulaValues_Datetime<br>EXFN_WorkInstruction.ACT_DateTime_New<br>EXFN_WorkInstruction.ACT_Date_change<br>EXFN_WorkInstruction.ACT_Time_Change<br>EXFN_WorkInstruction.ACT_dateTime_null_Set<br>EXFN_WorkInstruction.DS_ItemView_Datatime |
| Decimal | EXFN_WorkInstruction | EXFN_WorkInstruction.ACT_CalculateWorkInstructionFormulaValues_Decimal<br>EXFN_WorkInstruction.ACT_ItemView_Decimal_OnChange<br>EXFN_WorkInstruction.DS_ItemView_Decimal |
| Dropdown | EXFN_WorkInstruction | EXFN_WorkInstruction.ACT_ItemView_Dropdown_OnChange<br>EXFN_WorkInstruction.DS_DropdownItemList<br>EXFN_WorkInstruction.DS_ItemView_Dropdown |
| Integer | EXFN_WorkInstruction | EXFN_WorkInstruction.ACT_CalculateWorkInstructionFormulaValues_Integer<br>EXFN_WorkInstruction.ACT_ItemView_Integer_OnChange<br>EXFN_WorkInstruction.DS_ItemView_Integer |
| Multiline | EXFN_WorkInstruction | EXFN_WorkInstruction.ACT_CalculateWorkInstructionFormulaValues_Multiline<br>EXFN_WorkInstruction.ACT_ItemView_Multiline_OnChange<br>EXFN_WorkInstruction.DS_ItemView_Multiline |
| MultipleChoice | EXFN_WorkInstruction | EXFN_WorkInstruction.ACT_ItemView_MultipleChoice_OnChange<br>EXFN_WorkInstruction.DS_ItemView_MultipleChoice<br>EXFN_WorkInstruction.DS_MultipleChoiceItem |
| SNP_WorkInstruction_Classic | EXFN_WorkInstruction | EXFN_WorkInstruction.ACT_StepNavigation_First<br>EXFN_WorkInstruction.ACT_StepNavigation_Last<br>EXFN_WorkInstruction.ACT_StepNavigation_Next<br>EXFN_WorkInstruction.ACT_StepNavigation_Previous<br>EXFN_WorkInstruction.ACT_StepNavigation_Set<br>EXFN_WorkInstruction.ACT_WIStep_Acknowledge<br>EXFN_WorkInstruction.ACT_WIStep_Confirm<br>EXFN_WorkInstruction.ACT_WIStep_ReEdit<br>EXFN_WorkInstruction.DS_ScenarioConfiguration<br>EXFN_WorkInstruction.DS_ScenarioInstanceView<br>EXFN_WorkInstruction.DS_StepView<br>EXFN_WorkInstruction.DS_Step_Sequence<br>EXFN_WorkInstruction.DS_WI_SortedItemViewList<br>EXFN_WorkInstruction.EVT_WIStep_ES_Acquired<br>EXFN_WorkInstruction.SUB_GetCalculateWorkInstructionFormulaValues_BySection<br>EXFN_WorkInstruction.SUB_GetCalculateWorkInstructionFormulaValues_ByStep |
| SNP_WorkInstruction_VerticalView | EXFN_WorkInstruction | EXFN_WorkInstruction.ACT_WIStep_Acknowledge<br>EXFN_WorkInstruction.ACT_WIStep_Acquire<br>EXFN_WorkInstruction.ACT_WIStep_Confirm<br>EXFN_WorkInstruction.ACT_WIStep_ReEdit<br>EXFN_WorkInstruction.DS_ScenarioInstanceView<br>EXFN_WorkInstruction.DS_WI_SortedItemViewList<br>EXFN_WorkInstruction.EVT_WIStep_ES_Acquired<br>EXFN_WorkInstruction.SUB_GetCalculateWorkInstructionFormulaValues_BySection<br>EXFN_WorkInstruction.SUB_GetCalculateWorkInstructionFormulaValues_ByStep |
| Text | EXFN_WorkInstruction | EXFN_WorkInstruction.ACT_CalculateWorkInstructionFormulaValues_Text<br>EXFN_WorkInstruction.ACT_ItemView_Text_OnChange<br>EXFN_WorkInstruction.DS_ItemView_Text |
| WorkInstructionDefinitions_Overview | EXFN_WorkInstruction | EXFN_WorkInstruction.NAV_WorkInstruction_Preview |
| WorkInstructionDefinitions_Overview_Example | EXFN_WorkInstruction | EXFN_WorkInstruction.NAV_WorkInstructionPreview_Example |
| WorkInstruction_New_Step2 | EXFN_WorkInstruction | EXFN_WorkInstruction.ACT_WorkInstructionView_New_MF<br>EXFN_WorkInstruction.ACT_WorkInstruction_Create |
| WorkInstruction_Overview | EXFN_WorkInstruction | EXFN_WorkInstruction.ACT_WorkInstruction_Delete<br>EXFN_WorkInstruction.DS_WorkInstructionView_List<br>EXFN_WorkInstruction.NAV_WorkInstruction_View<br>EXFN_WorkInstruction.NAV_WorkInstruction_View_Vertical |
| WorkInstruction_Overview_Example | EXFN_WorkInstruction | EXFN_WorkInstruction.DS_WorkInstructionView_List<br>EXFN_WorkInstruction.NAV_WorkInstruction_VerticalView_Example |
| WorkInstruction_Preview | EXFN_WorkInstruction | EXFN_WorkInstruction.ACT_StepNavigation_First<br>EXFN_WorkInstruction.ACT_StepNavigation_Last<br>EXFN_WorkInstruction.ACT_StepNavigation_Next<br>EXFN_WorkInstruction.ACT_StepNavigation_Previous<br>EXFN_WorkInstruction.ACT_StepNavigation_Set<br>EXFN_WorkInstruction.DS_StepView<br>EXFN_WorkInstruction.DS_Step_Sequence<br>EXFN_WorkInstruction.DS_WI_SortedItemViewList |
| DocumentViewerSnippet | OpcenterEXDS_EXFN_DocumentViewer | EXFN_DocumentViewer.ACT_ButtonDownload<br>EXFN_DocumentViewer.ACT_Carousel<br>EXFN_DocumentViewer.ACT_ChangeDetailsView<br>EXFN_DocumentViewer.ACT_ChangeDocumentSelection<br>EXFN_DocumentViewer.ACT_ChangeModel3D<br>EXFN_DocumentViewer.ACT_CloseJTFullScreen<br>EXFN_DocumentViewer.ACT_FitToView<br>EXFN_DocumentViewer.ACT_SetMarkupColor<br>EXFN_DocumentViewer.ACT_SetPMIVisibility<br>EXFN_DocumentViewer.ACT_ToggleToolbox<br>EXFN_DocumentViewer.DS_GetCategories<br>EXFN_DocumentViewer.DS_GetDocumentByCategory<br>EXFN_DocumentViewer.DS_GetDocumentSetList<br>EXFN_DocumentViewer.DS_GetFile<br>EXFN_DocumentViewer.ToggleMarkup<br>EXFN_ServiceLayer.DS_GetApplicationURL<br>OpcenterEXDS_EXFN_DocumentViewer.ACT_AddGenerateMarkupToOperation<br>OpcenterEXDS_EXFN_DocumentViewer.ACT_EXDSOpenDocument<br>OpcenterEXDS_EXFN_DocumentViewer.EVT_OnSelectedObject |
| PANEL_SelectEquipmentAcquireWI | OpcenterEXDS_EXFN_WorkInstruction | OpcenterEXDS_EXFN_WorkInstruction.ACT_EquipmentSelectionPanel_AcquireDC<br>OpcenterEXDS_EXFN_WorkInstruction.ACT_Select_AcquireWIPanel_Equipment |
| SNP_WorkInstruction_VerticalView | OpcenterEXDS_EXFN_WorkInstruction | EXFN_Authentication.HandleUnauthorizedBehavior<br>EXFN_Authentication.Signal_Access_Token<br>EXFN_WorkInstruction.DS_ScenarioConfiguration<br>EXFN_WorkInstruction.DS_ScenarioInstanceView<br>EXFN_WorkInstruction.DS_WI_SortedItemViewList<br>EXFN_WorkInstruction.SUB_GetCalculateWorkInstructionFormulaValues_BySection<br>EXFN_WorkInstruction.SUB_GetCalculateWorkInstructionFormulaValues_ByStep<br>OpcenterEXDS_EXFN_WorkInstruction.ACT_WISectionRefresh<br>OpcenterEXDS_EXFN_WorkInstruction.ACT_WIStep_Acknowledge<br>OpcenterEXDS_EXFN_WorkInstruction.ACT_WIStep_Acquire<br>OpcenterEXDS_EXFN_WorkInstruction.ACT_WIStep_Confirm<br>OpcenterEXDS_EXFN_WorkInstruction.ACT_WIStep_ReEdit<br>OpcenterEXDS_EXFN_WorkInstruction.DS_CreateWISignalConfiguration<br>OpcenterEXDS_EXFN_WorkInstruction.DS_Get_WI_Step_Instructions<br>OpcenterEXDS_EXFN_WorkInstruction.EVT_WIStep_ES_Acquired<br>OpcenterEXDS_EXFN_WorkInstruction.EVT_WorkInstructionSectionCompleted_Details<br>OpcenterEXDS_EXFN_WorkInstruction.EVT_WorkInstructionStepCompleted_Details<br>OpcenterEXDS_OperatorLanding.DS_SetAutoGenerateId |
| AddDocuments | OpcenterEXDS_OperatorLanding | OpcenterEXDS_EXFN_DocumentViewer.ACT_OpenDocument<br>OpcenterEXDS_OperatorLanding.ACT_AddDocumentDropdownDocType<br>OpcenterEXDS_OperatorLanding.ACT_AddDocumentsSelectAndLinkDocument<br>OpcenterEXDS_OperatorLanding.ACT_SearchBy_Product<br>OpcenterEXDS_OperatorLanding.ACT_Show_AddDocumentsImportPanel<br>OpcenterEXDS_OperatorLanding.DS_CreateDocumentViewer<br>OpcenterEXDS_OperatorLanding.DS_CreateFile<br>OpcenterEXDS_OperatorLanding.DS_GetDocument<br>OpcenterEXDS_OperatorLanding.DS_GetOrCreate_DocSearchCommandBar_Context<br>OpcenterEXDS_OperatorLanding.DS_PreviewDocumentSelection |
| AddDocumentsMaterialTrackingUnits | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_AddDocumentsSelectMaterialTrackingUnit<br>OpcenterEXDS_OperatorLanding.DS_GetActiveMaterialTrackingUnitsRelatedToTheSelectedWorkOrderOperationOrStep<br>OpcenterEXDS_OperatorLanding.DS_GetLayoutHorizontalMode |
| AddDocumentsNavigationWizard | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_AddDocumentsNavigationWizard_Documents_Click<br>OpcenterEXDS_OperatorLanding.ACT_AddDocumentsNavigationWizard_MaterialTrackingUnits_Click<br>OpcenterEXDS_OperatorLanding.ACT_AddDocumentsNavigationWizard_Submit_Click<br>OpcenterEXDS_OperatorLanding.DS_GetLayoutVerticalMode |
| AddDocumentsPopup | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_Submit_AddDocuments<br>OpcenterEXDS_OperatorLanding.DS_AddDocumentsPage_RegisterClosePageAction<br>OpcenterEXDS_OperatorLanding.DS_GetLayoutHorizontalMode<br>OpcenterEXDS_OperatorLanding.NAV_Back |
| AddDocumentsSelectedDocumentBadge | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_AddDocumentsDeselectDocument |
| AddDocumentsSelectedMaterialTrackingUnitBadge | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_AddDocumentsDeselectMaterialTrackingUnit<br>OpcenterEXDS_OperatorLanding.DS_GetAddDocumentsSelectedMaterialTrackingUnits_Vertical |
| FailureBrowser | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_FailureList<br>OpcenterEXDS_OperatorLanding.ACT_Navigate_breadcrumb<br>OpcenterEXDS_OperatorLanding.ACT_SelectFailure<br>OpcenterEXDS_OperatorLanding.DS_GetSortedBreadcrumb<br>OpcenterEXDS_OperatorLanding.NAV_FailureChildren_Overview |
| HeaderBar | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_HeaderBar_EquipmentList_Click<br>OpcenterEXDS_OperatorLanding.ACT_Show_SelectSerialNumber_Panel<br>OpcenterEXDS_OperatorLanding.DS_CreateTimerContext<br>OpcenterEXDS_OperatorLanding.DS_TimerRefreshRemainingTime |
| NonConformanceDocuments | OpcenterEXDS_OperatorLanding | OpcenterEXDS_EXFN_DocumentViewer.ACT_OpenDocument<br>OpcenterEXDS_OperatorLanding.ACT_SelectDocument<br>OpcenterEXDS_OperatorLanding.ACT_ShowDocumentsPanel<br>OpcenterEXDS_OperatorLanding.ACT_Show_ImportPanel<br>OpcenterEXDS_OperatorLanding.DS_CreateDocumentSelection<br>OpcenterEXDS_OperatorLanding.DS_CreateDocumentViewer<br>OpcenterEXDS_OperatorLanding.DS_GetDocument<br>OpcenterEXDS_OperatorLanding.DS_GetDocumentList |
| NonConformanceEquipment_EquipmentContext | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_SelectNonConformanceEquipment_EquipmentContext<br>OpcenterEXDS_OperatorLanding.DS_GetLayoutHorizontalMode |
| NonConformanceMaterialTrackingUnits_MaterialTrackingUnitContext | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_NonConformanceMaterialTrackingUnits_ShowOnlyRelatedButtonClick<br>OpcenterEXDS_OperatorLanding.ACT_SelectNonConformanceMaterialTrackingUnit_ByView_MaterialTrackingUnitContext<br>OpcenterEXDS_OperatorLanding.ACT_SelectNonConformanceMaterialTrackingUnit_MaterialTrackingUnitContext<br>OpcenterEXDS_OperatorLanding.DS_GetLayoutHorizontalMode<br>OpcenterEXDS_OperatorLanding.DS_GetMaterialTrackingUnitsRelatedToTheSelectedWorkOrderOperation |
| NonConformanceMaterialTrackingUnits_WorkOrderOperationContext | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_SelectNonConformanceMaterialTrackingUnit_WorkOrderOperationContext<br>OpcenterEXDS_OperatorLanding.DS_GetLayoutHorizontalMode<br>OpcenterEXDS_OperatorLanding.DS_GetMaterialTrackingUnitsRelatedToTheSelectedWorkOrderOperation |
| NonConformancePopup | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_Submit_NonConformance<br>OpcenterEXDS_OperatorLanding.DS_GetLayoutHorizontalMode<br>OpcenterEXDS_OperatorLanding.DS_NonConformancePage_RegisterClosePageAction<br>OpcenterEXDS_OperatorLanding.NAV_Back |
| NonConformance_ToolContext | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_SelectNonConformanceTool<br>OpcenterEXDS_OperatorLanding.DS_GetLayoutHorizontalMode |
| NonConformancesInfo | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_NonConformancesInfo_OnOperationDetailsChange<br>OpcenterEXDS_OperatorLanding.ACT_Select_Equipment<br>OpcenterEXDS_OperatorLanding.ACT_Select_MaterialTrackingUnit_Context<br>OpcenterEXDS_OperatorLanding.ACT_Select_Tool_Context<br>OpcenterEXDS_OperatorLanding.ACT_Select_WorkOrderOperation_Context<br>OpcenterEXDS_OperatorLanding.DS_Get_Severity |
| NonConformancesNavigationWizard | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_NonConformancesNavigationWizard_Details_Click<br>OpcenterEXDS_OperatorLanding.ACT_NonConformancesNavigationWizard_Documents_Click<br>OpcenterEXDS_OperatorLanding.ACT_NonConformancesNavigationWizard_Failure_Click<br>OpcenterEXDS_OperatorLanding.ACT_NonConformancesNavigationWizard_Submit_Click<br>OpcenterEXDS_OperatorLanding.DS_GetLayoutVerticalMode |
| OperationContainer | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_OperationContainer_Complete<br>OpcenterEXDS_OperatorLanding.ACT_OperationContainer_Item_Click<br>OpcenterEXDS_OperatorLanding.ACT_OperationContainer_SendBuyOffNotification<br>OpcenterEXDS_OperatorLanding.DS_GetBuyOffStatus |
| OperationContainerHeaderBar | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_LinkWorkInstructionsToWOOperation_ShowPanel<br>OpcenterEXDS_OperatorLanding.ACT_LinkWorkInstructionsToWOStep_ShowPanel<br>OpcenterEXDS_OperatorLanding.ACT_OperationContainer_ExitFullscreen<br>OpcenterEXDS_OperatorLanding.ACT_OperationContainer_Fullscreen<br>OpcenterEXDS_OperatorLanding.ACT_OperationContainer_SendBuyOffNotification<br>OpcenterEXDS_OperatorLanding.ACT_OperationContainer_Start<br>OpcenterEXDS_OperatorLanding.ACT_QuickStart<br>OpcenterEXDS_OperatorLanding.DS_ChangeWorkOrderOpOrStepContext_WIOnDemand<br>OpcenterEXDS_OperatorLanding.DS_GetBuyOffStatus<br>OpcenterEXDS_OperatorLanding.DS_SetAllowedRuntimeActionProperties_ToWorkOrderOperationOrStepContext<br>OpcenterEXDS_OperatorLanding.SUB_SetSkillsByWorkOrderOperationOrStep |
| OperationList | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_Collapse_OperationList<br>OpcenterEXDS_OperatorLanding.ACT_Expand_OperationList<br>OpcenterEXDS_OperatorLanding.ACT_HideSteps_OperationList<br>OpcenterEXDS_OperatorLanding.ACT_OperationList_Icon_Click<br>OpcenterEXDS_OperatorLanding.ACT_OperationList_Item_Click<br>OpcenterEXDS_OperatorLanding.ACT_ShowRouting<br>OpcenterEXDS_OperatorLanding.ACT_ShowSteps_OperationList<br>OpcenterEXDS_OperatorLanding.ACT_Switch_HideCompletedOperation<br>OpcenterEXDS_OperatorLanding.ACT_Switch_ShowCompletedOperation<br>OpcenterEXDS_OperatorLanding.DS_GetSelectedWorkOrderOperationsAndSteps |
| OperatorLanding | OpcenterEXDS_OperatorLanding | EXFN_Authentication.HandleUnauthorizedBehavior<br>EXFN_Authentication.Signal_Access_Token<br>OpcenterEXDS_OperatorLanding.ACT_Routing_ShowPopup<br>OpcenterEXDS_OperatorLanding.ACT_SearchBy_Machine<br>OpcenterEXDS_OperatorLanding.ACT_SearchBy_Product<br>OpcenterEXDS_OperatorLanding.ACT_SearchCommandBar_ClearAll_Click<br>OpcenterEXDS_OperatorLanding.ACT_Set_FilterCriteria_OnChanged<br>OpcenterEXDS_OperatorLanding.ACT_WorkOrderOperation_Selected<br>OpcenterEXDS_OperatorLanding.ACT_WorkOrderOperation_Tile_GoToDetails_Click<br>OpcenterEXDS_OperatorLanding.DS_CreateDependencyGraphContext<br>OpcenterEXDS_OperatorLanding.DS_CreateSignalConfiguration<br>OpcenterEXDS_OperatorLanding.DS_GetBreadcrumb<br>OpcenterEXDS_OperatorLanding.DS_GetOrCreate_WorkOrderOperationHelper<br>OpcenterEXDS_OperatorLanding.EVT_CreateNonConformance<br>OpcenterEXDS_OperatorLanding.EVT_CreateOrAcknowledgeNote<br>OpcenterEXDS_OperatorLanding.EVT_OnCompleteReworkOrder<br>OpcenterEXDS_OperatorLanding.EVT_OnFAICandidateDeclared<br>OpcenterEXDS_OperatorLanding.EVT_OnFAICompleted<br>OpcenterEXDS_OperatorLanding.EVT_OnReopenWorkOrderOperation<br>OpcenterEXDS_OperatorLanding.EVT_OnSentenceNonConformance<br>OpcenterEXDS_OperatorLanding.EVT_OnSkipWOOperation<br>OpcenterEXDS_OperatorLanding.EVT_WorkOrderOperationActivedCP<br>OpcenterEXDS_OperatorLanding.EVT_WorkOrderOperationCompletedFullQty<br>OpcenterEXDS_OperatorLanding.EVT_WorkOrderOperationCompletedSerialized<br>OpcenterEXDS_OperatorLanding.EVT_WorkOrderOperationHold<br>OpcenterEXDS_OperatorLanding.EVT_WorkOrderOperationPaused<br>OpcenterEXDS_OperatorLanding.EVT_WorkOrderOperationStartedFullQty<br>OpcenterEXDS_OperatorLanding.EVT_WorkOrderOperationStartedSerialized<br>OpcenterEXDS_OperatorLanding.EVT_WorkOrderSerialNumbersScrapped<br>OpcenterEXDS_OperatorLanding.NAV_Back<br>OpcenterEXDS_OperatorLanding.SUB_RefreshOperatorLanding |
| OperatorTerminal | OpcenterEXDS_OperatorLanding | EXFN_Authentication.HandleUnauthorizedBehavior<br>EXFN_Authentication.Signal_Access_Token<br>OpcenterEXDS_OperatorLanding.DS_CreateSignalConfiguration<br>OpcenterEXDS_OperatorLanding.DS_FAIWoBadgeVisibility<br>OpcenterEXDS_OperatorLanding.DS_Gallery_OperationContainer_GetWorkOrderOperationContextList<br>OpcenterEXDS_OperatorLanding.DS_GetAssociatedDocuments_MF<br>OpcenterEXDS_OperatorLanding.DS_Get_WOOId_First<br>OpcenterEXDS_OperatorLanding.EVT_CreateOrAcknowledgeNote_Details<br>OpcenterEXDS_OperatorLanding.EVT_OnCompleteReworkOrder_Details<br>OpcenterEXDS_OperatorLanding.EVT_OnCreateNonConformance_Details<br>OpcenterEXDS_OperatorLanding.EVT_OnFAICandidateDeclared_Details<br>OpcenterEXDS_OperatorLanding.EVT_OnFAICompleted_Details<br>OpcenterEXDS_OperatorLanding.EVT_OnFAIRemoved_Details<br>OpcenterEXDS_OperatorLanding.EVT_OnReopenWorkOrderOperation_Details<br>OpcenterEXDS_OperatorLanding.EVT_OnSentenceNonConformance_Details<br>OpcenterEXDS_OperatorLanding.EVT_OnSkipWOOperation_Details<br>OpcenterEXDS_OperatorLanding.EVT_RefreshActivityIndicatorOnAssemblyWorkOrderOperation<br>OpcenterEXDS_OperatorLanding.EVT_RefreshActivityIndicatorOnDisassemblyWorkOrderOperation<br>OpcenterEXDS_OperatorLanding.EVT_RefreshActivityIndicatorOnUsedTool<br>OpcenterEXDS_OperatorLanding.EVT_RefreshActivityIndicatorOnWIStatusChangedSignal<br>OpcenterEXDS_OperatorLanding.EVT_WorkOrderOperationActivedCP_Details<br>OpcenterEXDS_OperatorLanding.EVT_WorkOrderOperationCompletedFullQty_Details<br>OpcenterEXDS_OperatorLanding.EVT_WorkOrderOperationCompletedSerialized_Details<br>OpcenterEXDS_OperatorLanding.EVT_WorkOrderOperationHold_Details<br>OpcenterEXDS_OperatorLanding.EVT_WorkOrderOperationPaused_Details<br>OpcenterEXDS_OperatorLanding.EVT_WorkOrderOperationStartedFullQty_Details<br>OpcenterEXDS_OperatorLanding.EVT_WorkOrderOperationStartedSerialized_Details<br>OpcenterEXDS_OperatorLanding.EVT_WorkOrderSerialNumbersScrapped_Details<br>OpcenterEXDS_OperatorLanding.EVT_WorkOrderStepCompletedFullQty_Details<br>OpcenterEXDS_OperatorLanding.EVT_WorkOrderStepCompletedSerialized_Details<br>OpcenterEXDS_OperatorLanding.EVT_WorkOrderStepStartedFullQty_Details<br>OpcenterEXDS_OperatorLanding.EVT_WorkOrderStepStartedSerialized_Details<br>OpcenterEXDS_OperatorLanding.NAV_Back |
| PANEL_AddDocumentsImportDocument | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_AddDocumentsCreateAndLinkDocument<br>OpcenterEXDS_OperatorLanding.ACT_CancelImportDocument |
| PANEL_AssignAndStartNonProductiveActivities | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_Close_CreateAndStartNonProductiveActivities<br>OpcenterEXDS_OperatorLanding.ACT_SelectNonProductiveActivities<br>OpcenterEXDS_OperatorLanding.ACT_SelectWorkOrder<br>OpcenterEXDS_OperatorLanding.ACT_SelectWorkOrderOperation<br>OpcenterEXDS_OperatorLanding.ACT_SetContextInfo<br>OpcenterEXDS_OperatorLanding.ACT_StartNonProductiveActivites |
| PANEL_ChangePackage | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_Submit_ChangePackage |
| PANEL_ChangeSN | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_GenerateNewSerialNumber<br>OpcenterEXDS_OperatorLanding.ACT_UpdateSerialNumberList |
| PANEL_CloseFlexibleWorkOrder | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_GenerateWONId<br>OpcenterEXDS_OperatorLanding.ACT_UADMAbruptlyCloseFlexibleWorkOrder<br>OpcenterEXDS_OperatorLanding.DS_GetOrCreateCloseFlexibleContext |
| PANEL_CompleteAssignedNonProductiveActivities | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_CompleteNonProductiveActivites<br>OpcenterEXDS_OperatorLanding.ACT_SelectAssignedNonProductiveActivities |
| PANEL_Complete_ValidateUser | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_Complete_ValidateUser<br>OpcenterEXDS_OperatorLanding.ACT_Complete_ValidateUser_FromDetails<br>OpcenterEXDS_OperatorLanding.DS_Create_ESContext |
| PANEL_DisassembleMaterialTrackingUnit | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_DisassembleMaterialWithNotes |
| PANEL_ImportDocument | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_CreateAndLinkDocument<br>OpcenterEXDS_OperatorLanding.ACT_NonConformancesInfo_OnOperationDetailsChange |
| PANEL_LinkDocuments | OpcenterEXDS_OperatorLanding | OpcenterEXDS_EXFN_DocumentViewer.ACT_OpenDocument<br>OpcenterEXDS_OperatorLanding.ACT_DropdownDocType<br>OpcenterEXDS_OperatorLanding.ACT_LinkDocument<br>OpcenterEXDS_OperatorLanding.ACT_MultiSelection_onLinkDocPanel<br>OpcenterEXDS_OperatorLanding.ACT_SearchBy_Product<br>OpcenterEXDS_OperatorLanding.DS_CreateDocumentViewer<br>OpcenterEXDS_OperatorLanding.DS_GetDocument<br>OpcenterEXDS_OperatorLanding.DS_GetOrCreate_DocSearchCommandBar_Context<br>OpcenterEXDS_OperatorLanding.DS_PreviewDocumentSelection |
| PANEL_Notes | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_Acknowledge_Note<br>OpcenterEXDS_OperatorLanding.ACT_Create_NoteToBeCreated_Operation<br>OpcenterEXDS_OperatorLanding.ACT_Create_NoteToBeCreated_WorkOrder<br>OpcenterEXDS_OperatorLanding.ACT_Create_Panel_Note |
| PANEL_OperatorDetailsCompleteStep | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_Complete_Panel_Step<br>OpcenterEXDS_OperatorLanding.ACT_Select_CompleteStepPanel_Equipment |
| PANEL_OperatorDetailsStartStep | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_Select_Equipment_Step<br>OpcenterEXDS_OperatorLanding.ACT_StartPanel_Step |
| PANEL_OperatorLandingComplete | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_CompleteAllSnButton<br>OpcenterEXDS_OperatorLanding.ACT_Complete_Panel_Operation<br>OpcenterEXDS_OperatorLanding.ACT_Select_CompleteOperationPanel_Equipment<br>OpcenterEXDS_OperatorLanding.ACT_SelectedSerialNumberInStartCompletePanel<br>OpcenterEXDS_OperatorLanding.DS_StartCompletePanel_GetEquipments |
| PANEL_OperatorLandingGoToDetailsSN | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_Select_GoToDetails_SerialNumber<br>OpcenterEXDS_OperatorLanding.ACT_Select_SerialNumber<br>OpcenterEXDS_OperatorLanding.ACT_SelectedSerialNumberInPanel<br>OpcenterEXDS_OperatorLanding.DS_GetSelectedSerialNumberContext |
| PANEL_OperatorLandingHold | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_Hold_Panel_Operation<br>OpcenterEXDS_OperatorLanding.ACT_Select_HoldReason |
| PANEL_OperatorLandingPause | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_Pause_Panel_Operation<br>OpcenterEXDS_OperatorLanding.ACT_Select_PausePanel_Equipment<br>OpcenterEXDS_OperatorLanding.ACT_Select_PauseReason |
| PANEL_OperatorLandingStart | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_AddNewSerialNumbersFromNId<br>OpcenterEXDS_OperatorLanding.ACT_AddNewSerialNumbersFromNumber<br>OpcenterEXDS_OperatorLanding.ACT_AssociateNewSerialNumber<br>OpcenterEXDS_OperatorLanding.ACT_Select_Equipment_Operation<br>OpcenterEXDS_OperatorLanding.ACT_SelectedSerialNumberInStartCompletePanel<br>OpcenterEXDS_OperatorLanding.ACT_StartAllSNButton<br>OpcenterEXDS_OperatorLanding.ACT_StartPanel_Operation<br>OpcenterEXDS_OperatorLanding.DS_SerialNumber_GetNotAlreadyAssociated<br>OpcenterEXDS_OperatorLanding.DS_StartCompletePanel_GetEquipments<br>OpcenterEXDS_OperatorLanding.DS_StartCompletePanel_GetSN<br>OpcenterEXDS_OperatorLanding.OCH_FlexSerialized_Quantity<br>OpcenterEXDS_OperatorLanding.OCH_SerialNumber_CheckIfExists |
| PANEL_Pause_ValidateUser | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_Pause_ValidateUser<br>OpcenterEXDS_OperatorLanding.DS_Create_ESContext |
| PANEL_SelectDestinationContainer | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_Cancel_SelectDestinationContainerPanel<br>OpcenterEXDS_OperatorLanding.ACT_Save_SelectDestinationContainerPanel<br>OpcenterEXDS_OperatorLanding.ACT_SelectDestinationContainer |
| PANEL_SelectToBeConsumedMTU | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_Cancel_SelectToBeConsumedPanel<br>OpcenterEXDS_OperatorLanding.ACT_Save_SelectToBeConsumedPanel<br>OpcenterEXDS_OperatorLanding.ACT_SelectToBeConsumedMTU |
| PANEL_SelectTool | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_SelectTool<br>OpcenterEXDS_OperatorLanding.ACT_ToggleToolSelection |
| PANEL_SetPoint | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_OpenSetPointHistoryPanel<br>OpcenterEXDS_OperatorLanding.ACT_TransmitAllEquipmentSetPointToAutomationNodeParameters<br>OpcenterEXDS_OperatorLanding.ACT_TransmitEquipmentSetPointToAutomationNodeParameters<br>OpcenterEXDS_OperatorLanding.DS_SetPointVariables |
| PANEL_SkipWOOperation | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_CallSkipOperation<br>OpcenterEXDS_OperatorLanding.ACT_SelectSkipReason<br>OpcenterEXDS_OperatorLanding.ACT_SelectSkipSerialNumber |
| PANEL_Start_ValidateUser | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_Start_ValidateUser<br>OpcenterEXDS_OperatorLanding.DS_Create_ESContext |
| PANEL_Start_ValidateUser_FromDetails | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_Start_ValidateUser_FromDetails<br>OpcenterEXDS_OperatorLanding.DS_Create_ESContext |
| PANEL_VerticalCommandBarMore | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_ClosePanel_VerticalCommandBarMore |
| PANEL_VerticalCommandBarMore_Details | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_ClosePanel_VerticalCommandBarMore |
| PANEL_VerticalCommandBarNC | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_ClosePanel_VerticalCommandBarMore |
| PANEL_VerticalCommandBarNC_Details | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_ClosePanel_VerticalCommandBarMore |
| PANEL_WOOP_NonConformance | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_CreateNC_From_Panel<br>OpcenterEXDS_OperatorLanding.ACT_Open_NonConformance_PC |
| PANEL_WorkInstructionsToWOOperation_Add | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_LinkWorkInstructionsToWOOperation |
| PANEL_WorkInstructionsToWOStep_Add | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_LinkWorkInstructionsToWOStep |
| PartProgram | OpcenterEXDS_OperatorLanding | EXFN_Authentication.HandleUnauthorizedBehavior<br>EXFN_Authentication.Signal_Access_Token<br>OpcenterEXDS_OperatorLanding.DS_GetOrCreateSignalEnvelope_PartProgram<br>OpcenterEXDS_PartProgram.ACT_CommandTransferDefault<br>OpcenterEXDS_PartProgram.EVT_DNCCompleteTransfer<br>OpcenterEXDS_PartProgram.EVT_DNCStartTransfer<br>OpcenterEXDS_PartProgram.NAV_ProgramPartDetails<br>OpcenterEXDS_PartProgram.NAV_ProgramPartHistory<br>OpcenterEXDS_PartProgram.SUB_SetPartProgramContext<br>OpcenterEXDS_PartProgram.SUB_WorkOrderOperationOrStepContext_SetDncItemId |
| QualityInspections | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_ParameterView_Create |
| Routing | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_NavigateToOperation<br>OpcenterEXDS_OperatorLanding.ACT_OnClick_RoutingNode<br>OpcenterEXDS_OperatorLanding.DS_GetBreadcrumb<br>OpcenterEXDS_OperatorLanding.DependencyGraphContext_SetHeight |
| Routing_Popup | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.DS_GetBreadcrumb<br>OpcenterEXDS_OperatorLanding.NAV_Back<br>OpcenterEXDS_OperatorLanding.SUB_DependencyGraphContext_RegisterClosePageAction |
| ScrapFailureBrowser | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_Navigate_ScrapFailureBreadcrumb<br>OpcenterEXDS_OperatorLanding.ACT_ScrapFailureList<br>OpcenterEXDS_OperatorLanding.ACT_SelectScrapFailure<br>OpcenterEXDS_OperatorLanding.DS_GetScrapFailureSortedBreadcrumb<br>OpcenterEXDS_OperatorLanding.NAV_ScrapFailureChildren_Overview |
| ScrapMaterialConsumptionPopup | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_DeselectScrapFailure<br>OpcenterEXDS_OperatorLanding.ACT_ScrapConsumedMaterial<br>OpcenterEXDS_OperatorLanding.DS_GetSelectedFailuresForScrap<br>OpcenterEXDS_OperatorLanding.DS_ScrapMaterialConsumedPage_RegisterClosePageAction<br>OpcenterEXDS_OperatorLanding.NAV_Back |
| ScrapMaterialConsumptionQuantity | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_ChangeScrappedQuantity<br>OpcenterEXDS_OperatorLanding.ACT_DescreaseScrapAssembledQuantity<br>OpcenterEXDS_OperatorLanding.ACT_IncreaseScrapAssembledQuantity<br>OpcenterEXDS_OperatorLanding.ACT_SearchDestinationContainer<br>OpcenterEXDS_OperatorLanding.ACT_SearchToBeConsumedMTU<br>OpcenterEXDS_OperatorLanding.ACT_ShowSelectDestinationContainerPanel<br>OpcenterEXDS_OperatorLanding.ACT_ShowSelectToBeConsumedMTUPanel |
| ScrapMaterialTrackingUnits | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_SelectScrapMTU<br>OpcenterEXDS_OperatorLanding.DS_GetMaterialTrackingUnitsActive |
| ScrapProducedMaterialPopup | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_DeselectScrapFailure<br>OpcenterEXDS_OperatorLanding.ACT_DeselectScrapMTU<br>OpcenterEXDS_OperatorLanding.ACT_ScrapWorkOrderSerialNumbers<br>OpcenterEXDS_OperatorLanding.DS_GetSelectedFailuresForScrap<br>OpcenterEXDS_OperatorLanding.DS_GetSelectedMaterialTrackingUnitsForScrap<br>OpcenterEXDS_OperatorLanding.DS_ScrapPage_RegisterClosePageAction<br>OpcenterEXDS_OperatorLanding.NAV_Back |
| SelectedDocumentBadge | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_DeselectDocument |
| SelectedEquipmentBadge_EquipmentContext | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_DeselectEquipment_EquipmentContext<br>OpcenterEXDS_OperatorLanding.DS_GetSelectedEquipments_Vertical_EquipmentContext |
| SelectedFailureBadge | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_DeselectFailure |
| SelectedMaterialTrackingUnitBadge_MaterialTrackingUnitContext | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_DeselectMaterialTrackingUnit_MaterialTrackingUnitContext<br>OpcenterEXDS_OperatorLanding.DS_GetSelectedMaterialTrackingUnits_Vertical_MaterialTrackingUnitContext |
| SelectedMaterialTrackingUnitBadge_WorkOrderOperationContext | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_DeselectMaterialTrackingUnit_MaterialTrackingUnitContext<br>OpcenterEXDS_OperatorLanding.ACT_DeselectMaterialTrackingUnit_WorkOrderOperationContext<br>OpcenterEXDS_OperatorLanding.DS_GetSelectedMaterialTrackingUnits_Vertical_WorkOrderOperationContext |
| SelectedToolBadge_ToolContext | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_DeselectTool<br>OpcenterEXDS_OperatorLanding.DS_GetSelectedTools_Vertical |
| ToBeCoByProducedMaterials | OpcenterEXDS_OperatorLanding | EXFN_Authentication.HandleUnauthorizedBehavior<br>EXFN_Authentication.Signal_Access_Token<br>OpcenterEXDS_OperatorLanding.ACT_CoProduceMaterial<br>OpcenterEXDS_OperatorLanding.ACT_CoProduceMaterialsAll<br>OpcenterEXDS_OperatorLanding.ACT_ToBeConsumedMaterialHistory<br>OpcenterEXDS_OperatorLanding.ACT_ValidateMaterialConsumption<br>OpcenterEXDS_OperatorLanding.DS_ChangeAssembleAllVisibility<br>OpcenterEXDS_OperatorLanding.DS_CreateCoByProductContext<br>OpcenterEXDS_OperatorLanding.DS_GetOrCreateSignalEnvelope_ToBeCoByProduct<br>OpcenterEXDS_OperatorLanding.DS_GetToBeCoProducedMaterial_Grid<br>OpcenterEXDS_OperatorLanding.EVT_OnByProductProduced<br>OpcenterEXDS_OperatorLanding.EVT_OnCoProductProduced<br>OpcenterEXDS_OperatorLanding.EVT_OnOutputMaterialProduced |
| ToBeConsumedMaterialCoByProduced | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_AcquireMTUList<br>OpcenterEXDS_OperatorLanding.ACT_ActualProducedMaterials_ToggleVisibility<br>OpcenterEXDS_OperatorLanding.ACT_CoProduceMaterial<br>OpcenterEXDS_OperatorLanding.ACT_ValidateMaterialConsumption<br>OpcenterEXDS_OperatorLanding.DS_GetToBeCoProductMaterial |
| ToBeConsumedMaterialCustomProduced | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_AcquireMTUList<br>OpcenterEXDS_OperatorLanding.ACT_ActualProducedMaterials_ToggleVisibility<br>OpcenterEXDS_OperatorLanding.ACT_CoProduceMaterial<br>OpcenterEXDS_OperatorLanding.ACT_ValidateMaterialConsumption<br>OpcenterEXDS_OperatorLanding.DS_GetToBeCustomProductMaterial |
| ToBeConsumedMaterialGridAlternative | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_AcquireMTUList<br>OpcenterEXDS_OperatorLanding.ACT_ConsumeMaterial<br>OpcenterEXDS_OperatorLanding.ACT_DisassembleMaterial_Grid<br>OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialConsumedPage_Grid<br>OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialToBeConsumedPage<br>OpcenterEXDS_OperatorLanding.ACT_ValidateMaterialConsumption<br>OpcenterEXDS_OperatorLanding.DS_GetToBeConsumeMaterial_GridMode_Alternative |
| ToBeConsumedMaterialGridCoByProduced | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_AcquireMTUList<br>OpcenterEXDS_OperatorLanding.ACT_CoProduceMaterial<br>OpcenterEXDS_OperatorLanding.ACT_ValidateMaterialConsumption<br>OpcenterEXDS_OperatorLanding.DS_GetToBeCoProductMaterial_GridMode |
| ToBeConsumedMaterialsAlternative | OpcenterEXDS_OperatorLanding | OpcenterEXDS_EXFN_DocumentViewer.ACT_ShowAllPartson3dViewer<br>OpcenterEXDS_OperatorLanding.ACT_AcquireMTUList<br>OpcenterEXDS_OperatorLanding.ACT_ActualMaterials_ToggleVisibility<br>OpcenterEXDS_OperatorLanding.ACT_ConsumeMaterial<br>OpcenterEXDS_OperatorLanding.ACT_DisassembleMaterial<br>OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialConsumedPage<br>OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialToBeConsumedPage<br>OpcenterEXDS_OperatorLanding.ACT_ValidateMaterialConsumption<br>OpcenterEXDS_OperatorLanding.DS_GetToBeConsumeMaterial_Alternative |
| ToBeConsumedMaterialsCustomPart | OpcenterEXDS_OperatorLanding | OpcenterEXDS_EXFN_DocumentViewer.ACT_ShowAllPartson3dViewer<br>OpcenterEXDS_OperatorLanding.ACT_AcquireMTUList<br>OpcenterEXDS_OperatorLanding.ACT_ActualMaterials_ToggleVisibility<br>OpcenterEXDS_OperatorLanding.ACT_ConsumeMaterial<br>OpcenterEXDS_OperatorLanding.ACT_DisassembleMaterial<br>OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialConsumedPage<br>OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialToBeConsumedPage<br>OpcenterEXDS_OperatorLanding.ACT_ValidateMaterialConsumption<br>OpcenterEXDS_OperatorLanding.DS_GetToBeConsumeMaterial_CustomPart |
| ToBeConsumedMaterialsDisassemble | OpcenterEXDS_OperatorLanding | OpcenterEXDS_EXFN_DocumentViewer.ACT_ShowAllPartson3dViewer<br>OpcenterEXDS_OperatorLanding.ACT_DisassembleMaterialSpecificationType<br>OpcenterEXDS_OperatorLanding.ACT_ValidateMaterialConsumption<br>OpcenterEXDS_OperatorLanding.DS_GetHistoryDisassembledMaterial<br>OpcenterEXDS_OperatorLanding.DS_GetToBeConsumeMaterial_Disassemble |
| ToBeConsumedMaterialsGrid | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_ToBeConsumedMaterialHistory<br>OpcenterEXDS_OperatorLanding.DS_GetToBeConsumedMaterial_GridMode<br>OpcenterEXDS_OperatorLanding.SetVisibilitySnippet_GridMode |
| ToBeConsumedMaterialsGridCustomPart | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_AcquireMTUList<br>OpcenterEXDS_OperatorLanding.ACT_ConsumeMaterial<br>OpcenterEXDS_OperatorLanding.ACT_DisassembleMaterial_Grid<br>OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialConsumedPage_Grid<br>OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialToBeConsumedPage<br>OpcenterEXDS_OperatorLanding.ACT_ValidateMaterialConsumption<br>OpcenterEXDS_OperatorLanding.DS_GetToBeConsumeMaterial_GridMode_CustomPart |
| ToBeConsumedMaterialsGridDisassemble | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_DisassembleMaterialSpecificationType<br>OpcenterEXDS_OperatorLanding.ACT_ValidateMaterialConsumption<br>OpcenterEXDS_OperatorLanding.DS_GetToBeConsumeMaterial_GridMode_Disassemble |
| ToBeConsumedMaterialsGridNormalPart | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_AcquireMTUList<br>OpcenterEXDS_OperatorLanding.ACT_ConsumeMaterial<br>OpcenterEXDS_OperatorLanding.ACT_DisassembleMaterial_Grid<br>OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialConsumedPage_Grid<br>OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialToBeConsumedPage<br>OpcenterEXDS_OperatorLanding.ACT_ValidateMaterialConsumption<br>OpcenterEXDS_OperatorLanding.DS_GetToBeConsumeMaterial_GridMode_NormalPart |
| ToBeConsumedMaterialsGridRangePartsAsRequired | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_AcquireMTUList<br>OpcenterEXDS_OperatorLanding.ACT_ConsumeMaterial<br>OpcenterEXDS_OperatorLanding.ACT_DisassembleMaterial_Grid<br>OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialConsumedPage_Grid<br>OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialToBeConsumedPage<br>OpcenterEXDS_OperatorLanding.ACT_ValidateMaterialConsumption<br>OpcenterEXDS_OperatorLanding.DS_GetToBeConsumeMaterial_GridMode_AsRequired<br>OpcenterEXDS_OperatorLanding.DS_GetToBeConsumeMaterial_GridMode_RangeParts |
| ToBeConsumedMaterialsGridReference | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.DS_GetToBeConsumeMaterial_Reference |
| ToBeConsumedMaterialsGridSelectedFit | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_AcquireMTUList<br>OpcenterEXDS_OperatorLanding.ACT_ConsumeMaterial<br>OpcenterEXDS_OperatorLanding.ACT_DisassembleMaterial_Grid<br>OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialConsumedPage_Grid<br>OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialToBeConsumedPage<br>OpcenterEXDS_OperatorLanding.ACT_ValidateMaterialConsumption<br>OpcenterEXDS_OperatorLanding.DS_GetToBeConsumeMaterial_GridMode_SelectedFit |
| ToBeConsumedMaterialsNormalPart | OpcenterEXDS_OperatorLanding | OpcenterEXDS_EXFN_DocumentViewer.ACT_ShowAllPartson3dViewer<br>OpcenterEXDS_OperatorLanding.ACT_AcquireMTUList<br>OpcenterEXDS_OperatorLanding.ACT_ActualMaterials_ToggleVisibility<br>OpcenterEXDS_OperatorLanding.ACT_ConsumeMaterial<br>OpcenterEXDS_OperatorLanding.ACT_DisassembleMaterial<br>OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialConsumedPage<br>OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialToBeConsumedPage<br>OpcenterEXDS_OperatorLanding.ACT_ValidateMaterialConsumption<br>OpcenterEXDS_OperatorLanding.DS_GetToBeConsumeMaterial_NormalPart<br>OpcenterEXDS_OperatorLanding_Connector.DS_GetMTUWithContainer |
| ToBeConsumedMaterialsPrekit | OpcenterEXDS_OperatorLanding | OpcenterEXDS_EXFN_DocumentViewer.ACT_ShowAllPartson3dViewer<br>OpcenterEXDS_OperatorLanding.ACT_ConsumeMaterial_Prekitted<br>OpcenterEXDS_OperatorLanding.ACT_EditPrekit<br>OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialToBeConsumedPage<br>OpcenterEXDS_OperatorLanding.ACT_ValidatePrekitCode<br>OpcenterEXDS_OperatorLanding.DS_GetToBeConsumeMaterial_Prekit<br>OpcenterEXDS_OperatorLanding.DS_TogglePrekitEdit |
| ToBeConsumedMaterialsRangePartsAndAsRequired | OpcenterEXDS_OperatorLanding | OpcenterEXDS_EXFN_DocumentViewer.ACT_ShowAllPartson3dViewer<br>OpcenterEXDS_OperatorLanding.ACT_AcquireMTUList<br>OpcenterEXDS_OperatorLanding.ACT_ActualMaterials_ToggleVisibility<br>OpcenterEXDS_OperatorLanding.ACT_ConsumeMaterial<br>OpcenterEXDS_OperatorLanding.ACT_DisassembleMaterial<br>OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialConsumedPage<br>OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialToBeConsumedPage<br>OpcenterEXDS_OperatorLanding.ACT_ValidateMaterialConsumption<br>OpcenterEXDS_OperatorLanding.DS_GetToBeConsumeMaterial_RangePartsAndAsRequired |
| ToBeConsumedMaterialsReference | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.DS_GetToBeConsumeMaterial_Reference |
| ToBeConsumedMaterialsSelectedFit | OpcenterEXDS_OperatorLanding | OpcenterEXDS_EXFN_DocumentViewer.ACT_ShowAllPartson3dViewer<br>OpcenterEXDS_OperatorLanding.ACT_AcquireMTUList<br>OpcenterEXDS_OperatorLanding.ACT_ActualMaterials_ToggleVisibility<br>OpcenterEXDS_OperatorLanding.ACT_ConsumeMaterial<br>OpcenterEXDS_OperatorLanding.ACT_DisassembleMaterial<br>OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialConsumedPage<br>OpcenterEXDS_OperatorLanding.ACT_Open_ScrapMaterialToBeConsumedPage<br>OpcenterEXDS_OperatorLanding.ACT_ValidateMaterialConsumption<br>OpcenterEXDS_OperatorLanding.DS_GetToBeConsumeMaterial_SelectedFit |
| ToBeProducedMaterial | OpcenterEXDS_OperatorLanding | EXFN_Authentication.HandleUnauthorizedBehavior<br>EXFN_Authentication.Signal_Access_Token<br>OpcenterEXDS_EXFN_DocumentViewer.ACT_ShowAllPartson3dViewer<br>OpcenterEXDS_OperatorLanding.ACT_ConsumeMaterialsAll<br>OpcenterEXDS_OperatorLanding.ACT_Set_AutoConsume_BatchMaterials<br>OpcenterEXDS_OperatorLanding.DS_ChangeAssembleAllVisibility<br>OpcenterEXDS_OperatorLanding.DS_CreateMatConsContext<br>OpcenterEXDS_OperatorLanding.DS_GetOrCreateSignalEnvelope_ConsumeMaterial<br>OpcenterEXDS_OperatorLanding.EVT_OnAssemblyWorkOrderOperation<br>OpcenterEXDS_OperatorLanding.EVT_OnDisassemblyWorkOrderOperation<br>OpcenterEXDS_OperatorLanding.EVT_ScrapMaterials_Details<br>OpcenterEXDS_OperatorLanding_Connector.DS_CreateContainerDTO |
| ToBeUsedTools | OpcenterEXDS_OperatorLanding | EXFN_Authentication.HandleUnauthorizedBehavior<br>EXFN_Authentication.Signal_Access_Token<br>OpcenterEXDS_OperatorLanding.ACT_AcquireTool<br>OpcenterEXDS_OperatorLanding.ACT_Grid_AcquireTool<br>OpcenterEXDS_OperatorLanding.ACT_Grid_UseAllTools<br>OpcenterEXDS_OperatorLanding.ACT_Grid_UseTool<br>OpcenterEXDS_OperatorLanding.ACT_ShowSelectToolPanel<br>OpcenterEXDS_OperatorLanding.ACT_ShowToolPanelForGrid<br>OpcenterEXDS_OperatorLanding.ACT_ToBeUsedToolHistory<br>OpcenterEXDS_OperatorLanding.ACT_ToggleIsActiveMaterialTrackingUnit<br>OpcenterEXDS_OperatorLanding.ACT_ToolHistory_ToggleVisibility<br>OpcenterEXDS_OperatorLanding.ACT_ToolInput_ToggleVisibility<br>OpcenterEXDS_OperatorLanding.ACT_UseAllTools<br>OpcenterEXDS_OperatorLanding.ACT_UseTool<br>OpcenterEXDS_OperatorLanding.ACT_ValidateDurationToolInGridMode<br>OpcenterEXDS_OperatorLanding.DS_CreateScrewingToolUsage<br>OpcenterEXDS_OperatorLanding.DS_CreateToBeUsedToolContext<br>OpcenterEXDS_OperatorLanding.DS_GetOrCreateSignalEnvelope_UseTool<br>OpcenterEXDS_OperatorLanding.DS_GetScrewingToolEntity<br>OpcenterEXDS_OperatorLanding.DS_GetTobeUsedToolHistory<br>OpcenterEXDS_OperatorLanding.DS_GetToolEntity<br>OpcenterEXDS_OperatorLanding.EVT_OnScrewingActivation<br>OpcenterEXDS_OperatorLanding.EVT_OnScrewingDeactivation<br>OpcenterEXDS_OperatorLanding.EVT_OnScrewingExecuted<br>OpcenterEXDS_OperatorLanding.EVT_OnUsedToolSignal<br>OpcenterEXDS_OperatorLanding.OCH_ToolNId |
| ToBeUsedTools_Backup | OpcenterEXDS_OperatorLanding | EXFN_Authentication.HandleUnauthorizedBehavior<br>EXFN_Authentication.Signal_Access_Token<br>OpcenterEXDS_OperatorLanding.ACT_AcquireTool<br>OpcenterEXDS_OperatorLanding.ACT_Grid_AcquireTool<br>OpcenterEXDS_OperatorLanding.ACT_Grid_UseAllTools<br>OpcenterEXDS_OperatorLanding.ACT_Grid_UseTool<br>OpcenterEXDS_OperatorLanding.ACT_ShowSelectToolPanel<br>OpcenterEXDS_OperatorLanding.ACT_ShowToolPanelForGrid<br>OpcenterEXDS_OperatorLanding.ACT_ToBeUsedToolHistory<br>OpcenterEXDS_OperatorLanding.ACT_ToolHistory_ToggleVisibility<br>OpcenterEXDS_OperatorLanding.ACT_ToolInput_ToggleVisibility<br>OpcenterEXDS_OperatorLanding.ACT_UseAllTools<br>OpcenterEXDS_OperatorLanding.ACT_UseTool<br>OpcenterEXDS_OperatorLanding.DS_CreateToBeUsedToolContext<br>OpcenterEXDS_OperatorLanding.DS_GetOrCreateSignalEnvelope_UseTool<br>OpcenterEXDS_OperatorLanding.DS_GetTobeUsedToolHistory<br>OpcenterEXDS_OperatorLanding.DS_GetToolEntity<br>OpcenterEXDS_OperatorLanding.EVT_OnUsedToolSignal<br>OpcenterEXDS_OperatorLanding.OCH_ToolNId |
| VerticalCommandBar | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_AddDocuments_Operation<br>OpcenterEXDS_OperatorLanding.ACT_ClosePanel_VerticalCommandBarMore<br>OpcenterEXDS_OperatorLanding.ACT_Complete<br>OpcenterEXDS_OperatorLanding.ACT_GoToDetails_Click<br>OpcenterEXDS_OperatorLanding.ACT_Hold_Operation<br>OpcenterEXDS_OperatorLanding.ACT_More<br>OpcenterEXDS_OperatorLanding.ACT_OpenChangePackage<br>OpcenterEXDS_OperatorLanding.ACT_OpenOperationSkip<br>OpcenterEXDS_OperatorLanding.ACT_Open_NC_Popup<br>OpcenterEXDS_OperatorLanding.ACT_Open_Notes<br>OpcenterEXDS_OperatorLanding.ACT_Open_ScrapProducedMaterialPage<br>OpcenterEXDS_OperatorLanding.ACT_Pause_Operation<br>OpcenterEXDS_OperatorLanding.ACT_ShowDocuments<br>OpcenterEXDS_OperatorLanding.ACT_Start<br>OpcenterEXDS_OperatorLanding.DS_GetOrCreateVerticalCommandBarContext |
| VerticalCommandBar_More | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_GoTo_ChangePackage<br>OpcenterEXDS_OperatorLanding.ACT_History<br>OpcenterEXDS_OperatorLanding.ACT_OpenChangeSNPanel<br>OpcenterEXDS_OperatorLanding.ACT_OpenCloseFlexiblePanel<br>OpcenterEXDS_OperatorLanding.ACT_OpenSetPointPanel<br>OpcenterEXDS_OperatorLanding.ACT_Open_AsBuilt<br>OpcenterEXDS_OperatorLanding.ACT_Open_ExternalIntegration<br>OpcenterEXDS_OperatorLanding.ACT_Open_Genealogy<br>OpcenterEXDS_OperatorLanding.ACT_Open_NonProductiveActivities<br>OpcenterEXDS_OperatorLanding.ACT_SetTargetQuantityOnFlexibleWorkOrder<br>OpcenterEXDS_OperatorLanding.ACT_TriggerPrintingOnWorkOrderOperation<br>OpcenterEXDS_OperatorLanding.DS_GetOrCreateVerticalCommandBarContext |
| VerticalCommandBar_NC | OpcenterEXDS_OperatorLanding | OpcenterEXDS_OperatorLanding.ACT_NonConformance_Operation<br>OpcenterEXDS_OperatorLanding.ACT_OpenNonConformanceList<br>OpcenterEXDS_OperatorLanding.DS_GetOrCreateVerticalCommandBarContext |
| WorkInstructions | OpcenterEXDS_OperatorLanding | EXFN_Authentication.HandleUnauthorizedBehavior<br>EXFN_Authentication.Signal_Access_Token<br>OpcenterEXDS_EXFN_WorkInstruction.DS_GetWorkInstructionContextHelper<br>OpcenterEXDS_OperatorLanding.DS_GetOrCreateWorkInstructionContext<br>OpcenterEXDS_OperatorLanding.DS_GetWorkInstructionView<br>OpcenterEXDS_OperatorLanding.EVT_LinkWIOnDemandToSerialNumber |
| DocumentViewerSnippet_DNC | OpcenterEXDS_PartProgram | OpcenterEXDS_PartProgram.ACT_GetTree<br>OpcenterEXDS_PartProgram.ACT_Search<br>OpcenterEXDS_PartProgram.ACT_SetHeight<br>OpcenterEXDS_PartProgram.ACT_SetSelected<br>OpcenterEXDS_PartProgram.ACT_ShowWorkOrderTree_Toogle<br>OpcenterEXDS_PartProgram.DS_Tree_Context |
| PartProgramDetailsPopUp | OpcenterEXDS_PartProgram | EXFN_Authentication.HandleUnauthorizedBehavior<br>EXFN_Authentication.Signal_Access_Token<br>OpcenterEXDS_OperatorLanding.NAV_Back<br>OpcenterEXDS_PartProgram.DS_CreateSignalConfiguration<br>OpcenterEXDS_PartProgram.DS_RegisterClosePageAction<br>OpcenterEXDS_PartProgram.EVT_DNCCompleteTransfer<br>OpcenterEXDS_PartProgram.EVT_DNCStartTransfer |
| PartProgramHistoryPopUp | OpcenterEXDS_PartProgram | OpcenterEXDS_OperatorLanding.NAV_Back<br>OpcenterEXDS_PartProgram.DS_RegisterClosePageAction |
| VerticalCommandBarPartProgram | OpcenterEXDS_PartProgram | OpcenterEXDS_PartProgram.ACT_Close<br>OpcenterEXDS_PartProgram.ACT_CommandDownloadAndManualTransfer<br>OpcenterEXDS_PartProgram.ACT_CommandPreview_NEW<br>OpcenterEXDS_PartProgram.ACT_CommandTransfer |

---

_Report generated by export_manifest tool_
