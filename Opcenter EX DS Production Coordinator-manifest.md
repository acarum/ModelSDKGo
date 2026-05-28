# Manifest Report: Opcenter EX DS Production Coordinator

**Mendix Version:** 11.9.0  
**MPR File:** C:\Workspaces\Mendix\MDUI\PC_AI-Demo-250526-11.9\Opcenter EX DS Production Coordinator.mpr  
**Generated:** 2026-05-27 16:54:28  

---

## Summary

- **External Entities:** 330 (across 9 modules)
- **Microflow/Action Calls:** 732
- **Signal Manager Subscriptions:** 15 subscription(s)
- **Navigation Items:** 39
- **Pages with Commands:** 31 page(s) analyzed
- **Command Buttons:** 142 total, 22 with extracted commands
- **Pages with Entity Datasources:** 55 page(s)

---

## 1. External Entities

External OData entities used in the project, grouped by module.

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

### Module: OpcenterEXDS_Configuration

#### Entity: Hold

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
| CloseComment | String |
| CloseTime | DateTime |
| CloseUserId | String |
| Comment | String |
| FutureHoldId | String |
| IsPresent | Boolean |
| OpenTime | DateTime |
| OpenUserId | String |
| PreviousState | String |
| ReasonId | String |
| _Type | String |
| TypeNId | String |
| HoldReason_Id | String |

#### Entity: HoldReason

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
| _Type | String |

### Module: OpcenterEXDS_Core

#### Entity: EquipmentConfiguration

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
| Name | String |
| Description | String |
| EquipmentTypeNId | String |
| LevelNId | String |

#### Entity: SetPoint

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
| Name | String |
| Description | String |
| Status_StateMachineNId | String |
| Status_StatusNId | String |

#### Entity: EquipmentConfigurationProperty

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
| PropertyValue | Integer/Decimal |
| Operational | String |
| EquipmentConfiguration_Id | String |

#### Entity: NonConformanceSeverity

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
| Val | String |
| NId | String |
| Name | String |
| Description | String |

#### Entity: WorkOrderOperation

**Published From:** AppU4DM

| Attribute | Type |
|-----------|------|
| NId | String |
| _Id | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| EntityType | String |
| IsLocked | Boolean |
| ToBeCleaned | String |
| ActiveNonConformanceNr | String |
| ActualEndTime | DateTime |
| ActualStartTime | DateTime |
| AvailableQuantity | String |
| Description | String |
| ElectronicSignatureComplete | String |
| ElectronicSignaturePause | String |
| ElectronicSignatureStart | String |
| EstimatedEndTime | DateTime |
| EstimatedStartTime | DateTime |
| LastPauseTime | DateTime |
| Name | String |
| OperationNId | String |
| OperationRevision | String |
| OperationUId | String |
| PartialWorkedQuantity | String |
| ProducedQuantity | String |
| RequiredCertificateNId | String |
| ReworkedQuantity | String |
| RequiredInspectionRole | String |
| ScrappedQuantity | String |
| TargetQuantity | String |
| ToBeCollectedDocument | String |
| OperationOccurrenceUId | String |
| ERPConfirmationId | String |
| EstimatedDuration_Ticks | String |
| ExecutionDuration_Ticks | String |
| PauseDuration_Ticks | String |
| IsReady | Boolean |
| Optional | String |
| Priority | String |
| Sequence | Integer |
| Skippable | String |
| WaitingForInspection | String |
| Operation_Id | String |
| WorkOperationType_Id | String |
| OperationStepCategoryId_Id | String |
| WorkOrder_Id | String |
| WOOFolder_Id | String |
| Status_StateMachineNId | String |
| Status_StatusNId | String |
| PreviousStatus_StateMachineNId | String |
| PreviousStatus_StatusNId | String |
| DetachedQuantity | String |

#### Entity: WorkOrder

**Published From:** AppU4DM

| Attribute | Type |
|-----------|------|
| NId | String |
| _Id | String |
| Name | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| EntityType | String |
| IsLocked | Boolean |
| ToBeCleaned | String |
| ActualEndTime | DateTime |
| ActualStartTime | DateTime |
| Notes | String |
| ParentBatch | String |
| PBOPIdentID | String |
| Plant | String |
| AsPlanned | String |
| CreationDate | DateTime |
| DueDate | DateTime |
| Enterprise | String |
| ERPOrder | String |
| EstimatedEndTime | DateTime |
| EstimatedStartTime | DateTime |
| InitialQuantity | String |
| ProcessNId | String |
| ProcessRevision | String |
| ProcessUId | String |
| SchedulingDate | DateTime |
| Sequence | Integer |
| ConfigurationParameter | String |
| RoutingID | String |
| TargetProductID | String |
| PlannedTargetQuantity | String |
| ActualTargetQuantity | String |
| TargetQuantity | String |
| ConfirmationType | String |
| RealInitialQty | String |
| PoC | String |
| Priority | String |
| IsUnderScheduling | Boolean |
| ProducedQuantity | String |
| ReworkedQuantity | String |
| ScrappedQuantity | String |
| Process_Id | String |
| FinalMaterial_Id | String |
| ToBeUsedBillOfFeature_Id | String |
| ReworkOfOrder_Id | String |
| ParentOrder_Id | String |
| ProductionType_Id | String |
| Status_StateMachineNId | String |
| Status_StatusNId | String |
| PreviousStatus_StateMachineNId | String |
| PreviousStatus_StatusNId | String |
| ToBeUsedBoM_NId | String |
| ToBeUsedBoM_Version | String |
| FinalMaterialUoMNId | String |
| DetachedQuantity | String |
| IsFAIRelevant | Boolean |
| IsFAICandidate | Boolean |

#### Entity: WOOFolder

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
| Name | String |
| Sequence | Integer |
| Description | String |
| OpFolderNId | String |
| OpFolderRevision | String |
| OpFolderUId | String |
| ProcessToOpFolderLinkUId | String |
| ProcessToOpFolderLink | String |
| ParentWOOFolder_Id | String |
| WorkOrder_Id | String |

#### Entity: WorkOOperationDependency

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
| DependencyType | String |
| PartialCompleted | String |
| IsFakeDependencyByWOOFolder | Boolean |
| FromWOO_Id | String |
| ToWOO_Id | String |

#### Entity: WorkOperationType

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
| Milestone | String |
| Optional | String |
| Skippable | String |
| ActiveOnlyOne | String |
| AutoStart | String |
| AutoComplete | String |
| MachineAutoStart | String |
| MachineAutoComplete | String |
| MachineAutoPause | String |
| CNC | String |
| AM | String |
| FastOperationConfirmation | String |

#### Entity: ProductionType

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

#### Entity: OperationStepCategory

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
| Description | String |
| Name | String |

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

#### Entity: Equipment

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
| Name | String |
| Description | String |
| WorkCalendarNId | String |
| EquipmentConfigurationId | String |
| LevelNId | String |
| Status_StateMachineNId | String |
| Status_StatusNId | String |

#### Entity: DM_MaterialTrackingUnit

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
| IsReserved | Boolean |
| ActiveNonConformanceNumber | String |
| BufferNId | String |
| CorrelationId | String |
| DM_MaterialId_Id | String |
| ParentDM_MTU_Id | String |
| MaterialTrackingUnit_Id | String |
| Weight_UoMNId | String |
| Weight_QuantityValue | Integer/Decimal |
| Volume_UoMNId | String |
| Volume_QuantityValue | Integer/Decimal |

#### Entity: MaterialTrackingUnit

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
| Name | String |
| Description | String |
| MaterialNId | String |
| MaterialRevision | String |
| MaterialUId | String |
| EquipmentNId | String |
| TemplateNId | String |
| Code | String |
| CodeType | String |
| CurrentLocation | String |
| MaterialLot_Id | String |
| MaterialTrackingUnitAggregate_Id | String |
| Status_StateMachineNId | String |
| Status_StatusNId | String |
| Quantity_UoMNId | String |
| Quantity_QuantityValue | Integer/Decimal |

#### Entity: DM_Material

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
| SerialNumberProfile | String |
| FirstArticleInspection | String |
| IsPhantom | Boolean |
| LogisticClassNId | String |
| CorrelationId | String |
| EffectivityExpression | String |
| Traceable | String |
| Lot | String |
| EnableAutomaticConsumptionQty | String |
| MaxQuantityConsumption | Integer/Decimal |
| FunctionalCodeNId | String |
| Material_Id | String |
| MaterialClass_Id | String |
| Supplier_Id | String |
| FunctionalCode_Id | String |
| Weight_UoMNId | String |
| Weight_QuantityValue | Integer/Decimal |
| Volume_UoMNId | String |
| Volume_QuantityValue | Integer/Decimal |

#### Entity: NonConformanceHistory

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
| Status | String |
| Notes | String |
| Transition | String |
| UserId | String |
| Action | String |
| Severity | String |
| PreviousSeverity | String |
| Equipment | String |
| FailureNId | String |
| ReasonNId | String |
| Date | DateTime |
| Defect_Id | String |
| Tool_Id | String |
| MaterialItem_Id | String |
| NonConformance_Id | String |

#### Entity: NonConformanceItem

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
| NonConformanceSupervisor | String |
| Equipment | String |
| Tool_Id | String |
| NonConformance_Id | String |
| DM_MaterialTrackingUnit_Id | String |

#### Entity: Tool

**Published From:** AppU4DM

| Attribute | Type |
|-----------|------|
| _Id | String |
| NId | String |
| Name | String |
| Description | String |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| ToolDefinition_Id | String |
| Status | String |
| EntityType | String |
| ToBeCleaned | String |
| ExpirationDate | DateTime |
| UsageCounter | Integer |
| UsageCounterMax | Integer |
| ActiveNonConformanceNr | String |
| UsageDuration_Ticks | String |
| UsageDurationMax_Ticks | Integer/Decimal |
| Lockable | String |
| IsLocked | Boolean |
| IsLock | Boolean |
| IsFrozen | Boolean |
| IsRuntimeFrozen | Boolean |
| IsScrap | Boolean |

#### Entity: LogisticWhitelist

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
| ContainedEntityType | String |
| ContainerEntityType | String |
| ToolDefinition_Id | String |
| DMMaterial_Id | String |
| CapacityType_Id | String |
| LogisticClass_Id | String |
| BufferDefinition_Id | String |
| MaxQty_UoMNId | String |
| MaxQty_QuantityValue | Integer/Decimal |
| MaxWeight_UoMNId | String |
| MaxWeight_QuantityValue | Integer/Decimal |
| MaxVolume_UoMNId | String |
| MaxVolume_QuantityValue | Integer/Decimal |
| Min_UoMNId | String |
| Min_QuantityValue | Integer/Decimal |
| Target_UoMNId | String |
| Target_QuantityValue | Integer/Decimal |
| Treshold_UoMNId | String |
| Treshold_QuantityValue | Integer/Decimal |

#### Entity: NonConformanceStatus

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
| Name | String |
| Description | String |

#### Entity: SegregationTag

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
| Name | String |
| Description | String |
| Color | String |
| TagStatus | String |
| IsLogRelevant | Boolean |

#### Entity: NonConformanceTransition

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
| Description | String |
| NId | String |
| Lifecycle | String |
| MailAddressList | String |
| Name | String |
| AuthorizedRoleList | String |
| SendMail | String |
| FromStatus_Id | String |
| ToStatus_Id | String |

#### Entity: WOOpDependencyNavigation

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
| AlternativeGroup | String |
| RelativePath | String |
| FullPath | String |
| ParentNavigation | String |
| IsPreferred | Boolean |
| WorkOOperationDependency_Id | String |

#### Entity: ToBeUsedTool

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
| TimesToBeUsed | DateTime |
| ToolDefinition_Id | String |
| WorkOrderOperation_Id | String |
| WorkOrderStep_Id | String |

#### Entity: ToolDefinition

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
| CorrelationId | String |
| Description | String |
| ExpirationDate | DateTime |
| Name | String |
| NId | String |
| ToolClass | String |
| ToolClassNId | String |
| UId | String |
| UsageCounterMax | Integer |
| UsageDurationMax | Integer/Decimal |
| Version | String |
| Consumable | String |
| RevisionExpression | String |
| Volume | String |
| VolumeUoM | String |
| WeightUoM | String |
| Weight | String |
| Lockable | String |
| LogisticClass_Id | String |

#### Entity: ToBeProducedMaterial

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
| Quantity | String |
| WorkOrderOperation_Id | String |
| DM_MaterialTrackingUnit_Id | String |
| WorkOrderStep_Id | String |

#### Entity: RM_PC_NonConformance

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| NId | String |
| _Type | String |
| _Context | String |
| Severity | String |
| Status | String |
| User | String |
| StartDate | DateTime |
| EndDate | DateTime |
| WorkOrderId | String |
| WorkOrderNId | String |
| WorkOrderOperationId | String |
| WorkOrderOperationNId | String |
| WorkOrderOperationStatus | String |
| WorkOrderOperationSequence | Integer |
| ChangeTypeNId | String |
| Equipment | String |
| ToolNId | String |
| MaterialNId | String |
| MaterialTrackingUnitNId | String |
| MaterialTrackingUnitRevision | String |
| MaterialTrackingUnitCode | String |
| _Id | String |
| NonConformanceId | String |

#### Entity: RM_PC_NonConformanceFailureOrDefect

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| GroupPath | String |
| Revision | String |
| NonConformanceNId | String |
| _Id | String |
| FailureDefectId | String |
| IsFailure | Boolean |
| DefectId | String |
| NonConformanceId | String |

#### Entity: RM_PC_NonConformanceDocuments

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Description | String |
| MIMEType | String |
| IconId | String |
| RepositoryType | String |
| LocalFileId | String |
| FileName | String |
| IconName | String |
| Revision | String |
| Category | String |
| NonconformanceId | String |
| NonconformanceHistoryId | String |
| LinkedEntityId | String |
| _Id | String |
| EntityLinkId | String |
| DocId | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |

#### Entity: RM_PC_NonConformanceMtuOrContainer

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| DM_MaterialTrackingUnitId | String |
| MaterialTrackingUnitId | String |
| ContainerId | String |
| NId | String |
| Code | String |
| CodeType | String |
| MaterialNId | String |
| MaterialRevision | String |
| Quantity | String |
| Status | String |
| svgIconPath | String |
| svgIconSize | String |
| _Id | String |
| NonConformanceNId | String |
| NonConformanceId | String |
| IsContainer | Boolean |

#### Entity: RM_PC_MaterialTrackingUnitAndContainer

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| NId | String |
| Code | String |
| CodeType | String |
| DM_MTUId | String |
| MaterialNId | String |
| MaterialRevision | String |
| Status | String |
| Quantity | String |
| StateMachine | String |
| ContainerType | String |
| IsReusable | Boolean |
| IsContainer | Boolean |
| IsResultKo | Boolean |
| _Id | String |
| MaterialTrackingUnitId | String |

#### Entity: RM_PC_ContainerOrMTUbyWorkOrderOperation

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| NId | String |
| Code | String |
| CodeType | String |
| DMMTUId | String |
| MaterialNId | String |
| MaterialRevision | String |
| StateMachine | String |
| Status | String |
| WorkOrderOperationId | String |
| WorkOrderOperationNId | String |
| IsContainer | Boolean |
| Quantity | String |
| ContainerType | String |
| IsReusable | Boolean |
| _Id | String |
| MaterialTrackingUnitOrContainerId | String |
| IsResultKo | Boolean |

#### Entity: NonConformanceAttachment

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
| Name | String |
| _Type | String |
| Document_Id | String |
| NonConformanceHistory_Id | String |
| Nonconformance_Id | String |

#### Entity: ChangeType

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
| Name | String |
| NId | String |
| Groupp | String |

#### Entity: NonConformanceFacet

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
| NonConformance_Id | String |

#### Entity: ChangeTypeFacet

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
| ChangeType_Id | String |

#### Entity: ToBeConsumedMaterial

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
| Name | String |
| DCDRuntimeTask | DateTime |
| DCDTask | String |
| FAIRuntimeTask | DateTime |
| GroupId | String |
| LogicalPosition | String |
| MaterialSpecificationType | String |
| PrekitSerialNumber | String |
| NId | String |
| OccurenceId | String |
| Unplanned | String |
| QuantityUoMNId | String |
| Quantity | String |
| AlternativeSelected | String |
| Sequence | Integer |
| MaterialDefinition_Id | String |
| WorkOrderOperation_Id | String |
| WorkOrderStep_Id | String |

#### Entity: MaterialClass

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
| Name | String |
| Description | String |

#### Entity: NonConformanceLifecycle

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
| Name | String |
| NId | String |
| _Type | String |
| Description | String |
| InitialStatus_Id | String |

#### Entity: RM_PC_NonConformanceHistory

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| NonConformanceNId | String |
| Status | String |
| Notes | String |
| Transition | String |
| UserId | String |
| Action | String |
| Severity | String |
| PreviousSeverity | String |
| Defect | String |
| MaterialItem | String |
| MaterialTrackingUnitId | String |
| MaterialTrackingUnitNId | String |
| MaterialTrackingUnitName | String |
| MaterialTrackingUnitCode | String |
| Equipment | String |
| Tool | String |
| FailureNId | String |
| ReasonNId | String |
| ContainmentRequestNId | String |
| ContainmentRequestHistory | String |
| DestinationNonConformanceNId | String |
| Date | DateTime |
| NonConformanceId | String |
| ContainerNId | String |
| ToolNId | String |
| _Id | String |
| NonConformanceHistoryId | String |
| IsFrozen | Boolean |
| IsLocked | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |

#### Entity: ScrewingOperationCharacteristics

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
| ToBeUsedTool_Id | String |
| AngleValueMin | Integer/Decimal |
| TorqueValueMin | Integer/Decimal |
| AngleValueMax | Integer/Decimal |
| TorqueValueMax | Integer/Decimal |
| NumberOfBolts | String |

### Module: OpcenterEXDS_MasterData

#### Entity: RM_PC_MaterialTrackingUnitResult

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| WorkOrderId | String |
| WorkOrderNId | String |
| WorkOrderName | String |
| ProductionTypeNId | String |
| WorkOrderOperationId | String |
| WorkOrderOperationNId | String |
| WorkOrderOperationSequence | Integer |
| WorkOrderOperationName | String |
| CategoryNId | String |
| CategoryName | String |
| WorkOperationTypeNId | String |
| MaterialTrackingUnitNId | String |
| EquipmentNId | String |
| ActionNId | String |
| ActionName | String |
| UserId | String |
| ResultIsKO | String |
| ResultValue | Integer/Decimal |
| ResultStrategy | String |
| NonConformanceNId | String |
| ActivityNId | String |
| ActivityType | String |
| Date | DateTime |
| ResultId | String |
| IsPartial | Boolean |
| CreatedWithActivity | String |
| IsParent | Boolean |
| _Id | String |
| ResultHistoryId | String |

#### Entity: RM_PC_DocumentsLinkedToEntityMaterial

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| Category | String |
| NId | String |
| Name | String |
| Description | String |
| MIMEType | String |
| IconId | String |
| RepositoryType | String |
| LocalFileId | String |
| FileName | String |
| IconName | String |
| Revision | String |
| DM_MaterialTrackingUnitDocumentAssociationId | String |
| MaterialTrackingUnitId | String |
| _Id | String |
| EntityLinkId | String |
| DocId | String |

#### Entity: RM_PC_DM_MaterialTrackingUnit

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Code | String |
| CurrentLocation | String |
| Description | String |
| EquipmentNId | String |
| MaterialLotId | String |
| MaterialLotNId | String |
| MaterialNId | String |
| MaterialRevision | String |
| MaterialUId | String |
| Quantity | String |
| QuantityUoMNId | String |
| StateMachineNId | String |
| StatusNId | String |
| TemplateNId | String |
| ActiveNonConformanceNumber | String |
| BufferNId | String |
| DM_MaterialId | String |
| VolumeQuantity | String |
| VolumeQuantityUoMNId | String |
| WeightQuantity | String |
| WeightUoMNId | String |
| MaxRecycleCount | Integer |
| MinQuantity | Integer/Decimal |
| MinQuantityUoMNId | String |
| RecycleCount | Integer |
| MtuAggregateId | String |
| MtuAggregateNId | String |
| _Id | String |
| MtuId | String |
| IsFrozen | Boolean |
| ToBeCleaned | String |
| DM_MaterialTrackingUnitId | String |
| CodeType | String |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |

#### Entity: RM_PC_DM_MaterialTrackingUnitHistory

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| ActionNId | String |
| DM_MaterialTrackingUnitId | String |
| BufferNId | String |
| Code | String |
| CodeType | String |
| Date | DateTime |
| DisassembledFromHistory | String |
| DocumentNId | String |
| EquipmentNId | String |
| ExecutionGroupPhaseId | String |
| ExecutionGroupPhaseNId | String |
| MixedQuantity | String |
| NewCode | String |
| OldCode | String |
| NonConformanceId | String |
| NonConformanceNId | String |
| OldQuantity | String |
| Quantity | String |
| QuantityUoMNId | String |
| RecycleCount | Integer |
| SourceMtuNId | String |
| SourceCode | String |
| TargetMtuNId | String |
| TargetCode | String |
| WorkOrderOperationNId | String |
| UserId | String |
| ContainerNId | String |
| _Id | String |
| CreatedOn | DateTime |
| Status | String |
| ContainmentRequestNId | String |
| CodeTypeName | String |
| Note | String |

#### Entity: RM_PC_BufferFromMTU

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| BufferDefinitionNId | String |
| CapacityTypeNId | String |
| Description | String |
| EquipmentName | String |
| EquipmentNId | String |
| MaxQtyValue | Integer/Decimal |
| MaxQtyUoM | Integer/Decimal |
| MaxVolumeValue | Integer/Decimal |
| MaxVolumeUoM | Integer/Decimal |
| MaxWeightValue | Integer/Decimal |
| MaxWeightUoM | Integer/Decimal |
| Name | String |
| StateMachineNId | String |
| StatusNId | String |
| MaterialTrackingUnitId | String |
| ValidityFrom | String |
| ValidityTo | String |
| NId | String |
| _Id | String |

#### Entity: RM_PC_MTULocationInfo

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| EquipmentName | String |
| EquipmentNId | String |
| EventType | String |
| ProductionItemName | String |
| ProductionItemNId | String |
| Timestamp | DateTime |
| User | String |
| IsCurrent | Boolean |
| _Id | String |

#### Entity: RM_PC_TraceabilityHistory

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| Barcode | String |
| EquipmentNId | String |
| ResultIsKO | String |
| ResultValue | Integer/Decimal |
| Timestamp | DateTime |
| User | String |
| WOOpNId | String |
| WOOpName | String |
| WOOpDescription | String |
| WOOpSequence | Integer |
| Operation | String |
| OperationNId | String |
| OperationRevision | String |
| OperationUId | String |
| Code | String |
| WorkOrderNId | String |
| IsValid | String |
| MaterialNId | String |
| _Id | String |
| WOOp | String |
| MaterialTrackingUnitId | String |

#### Entity: DSResult_AllowedSpecification

**Published From:** Material

| Attribute | Type |
|-----------|------|
| _Id | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| EntityType | String |
| IsLocked | Boolean |
| ToBeCleaned | String |
| SpecificationType | String |
| ResultStrategy_Id | String |

#### Entity: DSResult_ResultType

**Published From:** Material

| Attribute | Type |
|-----------|------|
| _Id | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| EntityType | String |
| IsLocked | Boolean |
| ToBeCleaned | String |
| ResultValue | Integer/Decimal |
| IsKO | Boolean |
| IsDefault | Boolean |

#### Entity: DSResult_ResultStrategy

**Published From:** Material

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

### Module: OpcenterEXDS_ProductionCoordination

#### Entity: Material

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
| Revision | String |
| SourceRevision | String |
| IsCurrent | Boolean |
| UId | String |
| NId | String |
| Name | String |
| Description | String |
| UoMNId | String |
| TemplateNId | String |

#### Entity: BillOfFeature

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
| Revision | String |
| SourceRevision | String |
| IsCurrent | Boolean |
| NId | String |
| Name | String |
| Description | String |

#### Entity: BoM

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
| Name | String |
| Description | String |
| Priority | String |
| ValidityFrom | String |
| ValidityTo | String |
| Version | String |
| NId | String |
| MaterialDefinition_Id | String |
| Status_StateMachineNId | String |
| Status_StatusNId | String |
| Quantity_UoMNId | String |
| Quantity_QuantityValue | Integer/Decimal |
| MaterialProperties_NId | String |
| MaterialProperties_UId | String |
| MaterialProperties_Revision | String |

#### Entity: ProducedMaterialItem

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
| DM_MaterialTrackingUnit_Id | String |
| WorkOrder_Id | String |

#### Entity: AsPlannedBOP

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
| BaselineName | String |
| CorrelationId | String |
| BaselineUId | String |
| PBOPIdentID | String |
| MasterPlanUID | String |
| OrderId | String |
| IsOutOfDate | Boolean |
| EffectivityExpression | String |
| Completed | String |

#### Entity: Process

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
| Name | String |
| Description | String |
| Sequence | Integer |
| NId | String |
| Revision | String |
| UId | String |
| Plant | String |
| CorrelationId | String |
| EffectivityExpression | String |
| Version | String |
| FinalMaterialId_Id | String |
| MaxQuantity_UoMNId | String |
| MaxQuantity_QuantityValue | Integer/Decimal |
| Quantity_UoMNId | String |
| Quantity_QuantityValue | Integer/Decimal |
| MinQuantity_UoMNId | String |
| MinQuantity_QuantityValue | Integer/Decimal |
| DiscreteStatus_StateMachineNId | String |
| DiscreteStatus_StatusNId | String |

#### Entity: StateMachine

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
| Name | String |
| Description | String |
| IsHidden | Boolean |
| IsSystemDefined | Boolean |

#### Entity: Certification

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
| Name | String |
| Description | String |
| CorrelationId | String |

#### Entity: MaterialSpecificationType

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
| HasOutputBehavior | Boolean |

#### Entity: Team

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
| Name | String |
| Description | String |
| MaximumUsers | Integer/Decimal |

#### Entity: TeamUserAssociation

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
| User | String |
| Team_Id | String |

#### Entity: TeamSkillAssociation

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
| Level | String |
| Skill_Id | String |
| Teams_Id | String |

#### Entity: TeamHistory

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
| User | String |
| AssociatedUser | String |
| Date | DateTime |
| Action_Id | String |
| Teams_Id | String |

#### Entity: Skill_BOP

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
| Name | String |
| Description | String |
| Color | String |

#### Entity: TeamHistoryAction

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

#### Entity: NumberingPatternMaterialAssociation

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
| NumberingPatternNId | String |
| EntityTypeNId | String |
| DMMaterialId | String |

#### Entity: UoM

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
| Name | String |
| Description | String |
| IsHidden | Boolean |
| IsSystemDefined | Boolean |
| UoMDimension_Id | String |
| UoMBase_Id | String |

#### Entity: EquipmentType

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
| Name | String |
| Description | String |
| IsDefault | Boolean |
| LevelNId | String |

#### Entity: WorkInstructionDefinition

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
| Revision | String |
| SourceRevision | String |
| IsCurrent | Boolean |
| NId | String |
| Name | String |
| Description | String |
| IsExecutable | Boolean |
| WorkInstructionTemplateNId | String |
| Design_Model | String |

#### Entity: WorkInstructionAssociationType

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

#### Entity: WorkOrderStep

**Published From:** AppU4DM

| Attribute | Type |
|-----------|------|
| NId | String |
| _Id | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| EntityType | String |
| IsLocked | Boolean |
| ToBeCleaned | String |
| ActualEndTime | DateTime |
| ActualStartTime | DateTime |
| Description | String |
| ElectronicSignatureComplete | String |
| ElectronicSignaturePause | String |
| ElectronicSignatureStart | String |
| EstimatedEndTime | DateTime |
| EstimatedStartTime | DateTime |
| LastPauseTime | DateTime |
| Name | String |
| Optional | String |
| PauseDuration | String |
| Priority | String |
| RequiredCertificateNId | String |
| RequiredInspectionRole | String |
| Sequence | Integer |
| Step | String |
| StepNId | String |
| StepRevision | String |
| StepUId | String |
| WaitingForInspection | String |
| TargetQuantity | String |
| AvailableQuantity | String |
| PartialWorkedQuantity | String |
| ProducedQuantity | String |
| ToBeCollectedDocument | String |
| NotWorkedQuantity | String |
| ScrappedQuantity | String |
| ReworkedQuantity | String |
| EstimatedDuration_Ticks | String |
| ExecutionDuration_Ticks | String |
| IsReady | Boolean |
| OperationStepCategoryId_Id | String |
| WorkOrderOperation_Id | String |
| Status_StateMachineNId | String |
| Status_StatusNId | String |

#### Entity: ToBeUsedBillOfFeature

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
| BillOfFeatureNId | String |
| BillOfFeatureRevision | String |

#### Entity: SnagAndNote

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

#### Entity: Status

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
| Name | String |
| Description | String |
| Color | String |
| NId | String |
| IsHidden | Boolean |
| IsSystemDefined | Boolean |
| IsInitial | Boolean |
| Outcome | String |
| StateMachine_Id | String |

#### Entity: WorkOStepDependency

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
| DependencyType | String |
| To_Id | String |
| WosFrom_Id | String |

#### Entity: Step

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
| Name | String |
| Description | String |
| Sequence | Integer |
| Optional | String |
| CorrelationId | String |
| NId | String |
| Revision | String |
| UId | String |
| RequiredInspectionRole | String |
| ElectronicSignatureComplete | String |
| ElectronicSignaturePause | String |
| ElectronicSignatureStart | String |
| RequiredCertificateNId | String |
| EffectivityExpression | String |
| ToBeCollectedDocument | String |
| EstimatedDuration_Ticks | String |
| OperationStepCategoryId_Id | String |

#### Entity: Operation

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
| Name | String |
| Description | String |
| Optional | String |
| Sequence | Integer |
| CorrelationId | String |
| NId | String |
| Revision | String |
| UId | String |
| ElectronicSignatureStart | String |
| ElectronicSignaturePause | String |
| ElectronicSignatureComplete | String |
| RequiredInspectionRole | String |
| RequiredCertificateNId | String |
| ToBeCollectedDocument | String |
| EffectivityExpression | String |
| EstimatedDuration_Ticks | String |
| WorkOperationId_Id | String |
| OperationStepCategoryId_Id | String |

#### Entity: OperationToStepLink

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
| CorrelationId | String |
| Sequence | Integer |
| UId | String |
| EffectivityExpression | String |
| ToBeCollectedDocument | String |
| ParentOperation_Id | String |
| ChildStep_Id | String |
| AsPlannedBOP_Id | String |
| MasterPlan_Id | String |

#### Entity: ProcessToOperationLink

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
| CorrelationId | String |
| Sequence | Integer |
| UId | String |
| EffectivityExpression | String |
| ParentProcess_Id | String |
| ChildOperation_Id | String |
| AsPlannedBOP_Id | String |
| MasterPlan_Id | String |
| ParentProcessToOpFolderLink_Id | String |

#### Entity: MasterPlan

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
| CorrelationId | String |
| UId | String |
| Completed | String |
| ComplementaryEngineeringNeeded | String |
| SpecializedSpecificationSet | String |

#### Entity: CharacteristicRepresentation

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
| Revision | String |
| SourceRevision | String |
| IsCurrent | Boolean |
| Name | String |
| NId | String |
| CanBeSkipped | String |
| CharacteristicRepresentationFrequencyValue | Integer/Decimal |
| TimeBasedScheduleMode | DateTime |
| TimeBasedMandatoryInspectionWarningValue | DateTime |
| TimeBasedMandatoryInspectionWarningUoM | DateTime |
| SampleSize | String |
| CharacteristicSpecification_Id | String |
| CharacteristicRepresentationFrequency_Id | String |

#### Entity: CharacteristicRepresentationFrequency

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
| Name | String |
| Description | String |

#### Entity: ItlkCheck

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
| Name | String |
| Description | String |
| StepRelevant | String |
| OperationRelevant | String |
| IsStandard | Boolean |
| IsInbound | Boolean |
| IsOutbound | Boolean |

#### Entity: WorkOrderOperationSkill

**Published From:** AppU4DM

| Attribute | Type |
|-----------|------|
| _Id | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| ToBeCleaned | String |
| Level | String |
| WorkOrderOperation_Id | String |

#### Entity: ItlkCheckParameter

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
| Name | String |
| DataType | String |
| IsMandatory | Boolean |
| IsList | Boolean |
| ItlkCheck_Id | String |

#### Entity: RM_PC_RoutingNode

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| WorkOrderId | String |
| WorkOrderNId | String |
| WorkOrderName | String |
| WorkOrderStatusNId | String |
| EntityId | String |
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
| IsExternal | Boolean |
| _Id | String |
| CurrentWorkOrderId | String |
| ParentBatch | String |

#### Entity: RM_PC_RoutingEdge

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| FromWorkOrderId | String |
| FromNodeEntityId | String |
| FromNodeEntityType | String |
| FromNodeParentEntityId | String |
| ToWorkOrderId | String |
| ToNodeEntityId | String |
| ToNodeEntityType | String |
| ToNodeParentEntityId | String |
| DependencyType | String |
| IsExternalDependency | Boolean |
| _Id | String |
| CurrentWorkOrderId | String |
| ParentBatch | String |
| AlternativeGroup | String |
| IsPreferred | Boolean |

#### Entity: RM_PC_WOOFolderBreadcrumb

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| ChildFolderNId | String |
| ParentFolderId | String |
| Hierarchy | String |
| HierarchyLabel | String |
| _Id | String |
| ChildFolderId | String |

#### Entity: RM_PC_GetWOSteps

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| Name | String |
| Description | String |
| UID | String |
| Revision | String |
| EstimatedDurationHours | String |
| EstimatedDurationMinutes | Integer/Decimal |
| EstimatedDurationSeconds | String |
| ElectronicSignatureComplete | String |
| Sequence | Integer |
| OperationStepCategoryId | String |
| OperationStepCategoryNId | String |
| ToBeCollectedDocument | String |
| NextSequence | Integer |
| Status | String |
| ActualStartTime | DateTime |
| EstimatedDuration | String |
| _Id | String |
| WorkOrderOperationId | String |
| WorkOrderStepId | String |
| IsReady | Boolean |
| ActualEndTime | DateTime |

#### Entity: WorkOrderHumanResource

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
| CertificationNId | String |
| NumberOfUsers | String |
| WorkOrderOperation_Id | String |
| WorkOrderStep_Id | String |

#### Entity: RM_PC_GetFullEquipment

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| LevelNId | String |
| _Type | String |
| Preferred | String |
| PartProgram | String |
| AssociatedPrintJobFileId | String |
| AssociatedPrintJobFileNId | String |
| AssociatedPrintJobFileNIdRuntimeOnly | DateTime |
| AssociatedPrintJobFileProductSerializable | String |
| AssociatedPrintJobFileProductSerialManagement | String |
| Plant | String |
| RequirementTag | String |
| _Id | String |
| EquipId | String |
| ToBeUsedMachineId | String |
| Locked | Boolean |
| WorkOrderOperationId | String |
| IsWorkOperationAM | Boolean |
| IsWorkOperationCNC | Boolean |

#### Entity: RM_PC_GetAllMaterialDefinitionsForAllMaterialSpecificationTypes

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Behaviour | String |
| LogicalPosition | String |
| MaterialSpecificationType | String |
| WorkOperationId | String |
| WorkOrderStepId | String |
| Quantity | String |
| AlternativeSelected | String |
| GroupName | String |
| PrekitSerialNumber | String |
| HasPrekit | Boolean |
| Sequence | Integer |
| SerialNumberProfile | String |
| Revision | String |
| UoMNId | String |
| OccurenceId | String |
| Fixed | String |
| RequirementTag | String |
| _Id | String |
| MaterialDefinitionId | String |
| ToBeConsumedCoProducedId | String |

#### Entity: MaterialSpecification

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
| UId | String |
| LogicalPosition | String |
| GroupId | String |
| AlternativeSelected | String |
| EffectivityExpression | String |
| CorrelationId | String |
| FunctionalCodeNId | String |
| OccurrenceId | String |
| DM_MaterialId_Id | String |
| Operation_Id | String |
| Step_Id | String |
| MaterialTypeNId_Id | String |
| AsPlannedBOP_Id | String |
| MasterPlan_Id | String |
| Quantity_UoMNId | String |
| Quantity_QuantityValue | Integer/Decimal |

#### Entity: KanbanCall

**Published From:** Kanban

| Attribute | Type |
|-----------|------|
| _Id | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| EntityType | String |
| IsLocked | Boolean |
| ToBeCleaned | String |
| CallSessionId | String |
| LineSidePositionNId | String |
| CurrentPriority | String |
| CurrentLogic | String |
| CurrentStatus_StateMachineNId | String |
| CurrentStatus_StatusNId | String |
| ActualRequestQty_UoMNId | String |
| ActualRequestQty_QuantityValue | Integer/Decimal |
| MessageDefinition_MsgDefinitionNId | String |
| MessageDefinition_MessageDefinitionRevision | String |

#### Entity: LineSidePosition

**Published From:** Kanban

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
| EquipmentNId | String |
| LSMaxBins | Integer/Decimal |
| StartQty | String |
| Released | String |
| BoMConsumptionType | String |
| BoMConsumptionValue | Integer/Decimal |
| BinsToBeRequested | String |
| OnlyManual | String |
| Material_MaterialNId | String |
| Material_MaterialUId | String |
| Material_MaterialRevision | String |
| BinStdQty_UoMNId | String |
| BinStdQty_QuantityValue | Integer/Decimal |
| MaterialSource_MsgDefinitionNId | String |
| MaterialSource_MessageDefinitionRevision | String |

#### Entity: LSPThreshold

**Published From:** Kanban

| Attribute | Type |
|-----------|------|
| _Id | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| EntityType | String |
| IsLocked | Boolean |
| ToBeCleaned | String |
| ThresholdValue | Integer/Decimal |
| Priority | String |
| LineSidePosition_Id | String |

#### Entity: ProducedMTUWithResult

**Published From:** AppU4DM

| Attribute | Type |
|-----------|------|
| WOId | String |
| MTUNId | String |
| Name | String |
| Description | String |
| IsReserved | Boolean |
| MaterialDefinitionId | String |
| Quantity | String |
| SerialNumberCode | String |
| UoMNId | String |
| Status | String |
| MaterialDefinitionNId | String |
| WorkOrderStatus | String |
| RecycleCount | Integer |
| MaxRecycleCount | Integer |
| MinQuantity | Integer/Decimal |
| CodeType | String |
| _Id | String |
| IsResultKo | Boolean |
| MTUId | String |
| DMMTUId | String |

#### Entity: OutMsg

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
| MsgId | String |
| OutMsgDefinitionNId | String |
| MsgHeaderData | String |
| SendTime | DateTime |
| MsgSchemaNId | String |
| IsRegenerated | Boolean |
| PlantNId | String |
| PlantERPNId | String |
| WorkPlaceNId | String |
| DestinationNId | String |
| PFCMilestoneNId | String |
| PFCDiagramNId | String |
| MaterialInstanceNId | String |
| PFCMilestoneSequence | Integer |
| FileId | String |
| OutMsgDefinitionRevision | String |

#### Entity: UserWorkOrderOperationAssociation

**Published From:** AppU4DM

| Attribute | Type |
|-----------|------|
| _Id | String |
| UserId | String |
| WorkOrderOperation_Id | String |

#### Entity: WorkOrderStepSkill

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
| Level | String |
| SkillId_Id | String |
| WorkOrderStep_Id | String |

#### Entity: OutMsgDefinition

**Published From:** Kanban

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
| MsgTypeNId | String |
| IdCreationTemplateNId | String |
| OutMsgDestination_Id | String |

#### Entity: RM_PC_WorkOrderOperation

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Description | String |
| StateMachineNId | String |
| StatusNId | String |
| PreviousStateMachineNId | String |
| PreviousStatusNId | String |
| EstimatedStartTime | DateTime |
| EstimatedEndTime | DateTime |
| ActualStartTime | DateTime |
| ActualEndTime | DateTime |
| EstimatedDuration | String |
| Priority | String |
| Sequence | Integer |
| IsReady | Boolean |
| ExecutionDuration | String |
| LastPauseTime | DateTime |
| IsOptional | Boolean |
| IsSkippable | Boolean |
| PauseDuration | String |
| AvailableQuantity | String |
| TargetQuantity | String |
| ProducedQuantity | String |
| PartialWorkedQuantity | String |
| ReworkedQuantity | String |
| ScrappedQuantity | String |
| ToBeCollectedDocument | String |
| ActiveNonConformanceNr | String |
| OperationId | String |
| OperationNId | String |
| OperationRevision | String |
| OperationUId | String |
| ParentWorkOrderOperationFolderId | String |
| WorkOperationTypeNId | String |
| WorkOperationTypeAM | String |
| WorkOperationTypeAutoStart | String |
| WorkOperationTypeAutoComplete | String |
| KOQuantity | String |
| _Id | String |
| WorkOrderOperationId | String |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| WorkOrderId_Id | String |
| OperationStepCategoryId | String |
| HasInterlockingCheck | Boolean |
| CanReopenQty | String |
| HasTransmittedSetPoints | Boolean |
| ElectronicSignatureComplete | String |
| ElectronicSignaturePause | String |
| ElectronicSignatureStart | String |
| RequiredCertificateNId | String |
| RequiredInspectionRole | String |
| WaitingForInspection | String |

#### Entity: EquipmentConfiguration

**Published From:** Kanban

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
| EquipmentTypeNId | String |
| LevelNId | String |

#### Entity: RM_PC_EquipmentWorkOrderOperationCount

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| EquipmentNId | String |
| EquipmentName | String |
| EquipmentLevel | String |
| EquipmentType | String |
| CompleteByDifferentUser | String |
| _Id | String |
| OpenOccurrence | String |
| QueueOccurrence | String |
| ActiveOccurrence | String |

#### Entity: RM_PC_WorkOrderOperationByEquipmentAndStatus

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| EquipmentNId | String |
| WorkOrderNId | String |
| WorkOrderStatus | String |
| _Id | String |
| WorkOrderId | String |
| IsWorkOrderOperationActive | Boolean |
| IsWorkOrderOperationComplete | Boolean |
| IsWorkOrderOperationOpen | Boolean |
| IsWorkOrderOperationReady | Boolean |
| IsWorkOrderOperationPartial | Boolean |
| WorkOrderOperationNId | String |
| WorkOrderOperationName | String |
| WorkOrderOperationStatus | String |
| WorkOrderOperationReady | String |
| WorkOrderOperationEstimatedStartTime | DateTime |
| WorkOrderOperationEstimatedEndTime | DateTime |
| WorkOrderOperationActualStartTime | DateTime |
| WorkOrderOperationActualEndTime | DateTime |
| WorkOrderOperationEstimatedDuration | String |
| WorkOrderOperationSequence | Integer |
| IsPreferred | Boolean |
| WorkOrderOperationId | String |

#### Entity: RM_PC_AsBuiltMaterials

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| WONId | String |
| WOOpName | String |
| WOOpSequence | Integer |
| WOOpStatusNId | String |
| WOStepId | String |
| WOStepName | String |
| WOStepSequence | Integer |
| WOStepStatusNId | String |
| MTUNId | String |
| Code | String |
| UoMNId | String |
| EquipmentNId | String |
| MaterialSpecificationType | String |
| LogicalPosition | String |
| OccurenceId | String |
| ToBeConsumedQuantityUoMNId | String |
| MaterialNId | String |
| MaterialUoMNId | String |
| MaterialClassNId | String |
| RecycleCount | Integer |
| MaxRecycleCount | Integer |
| MinQuantity | Integer/Decimal |
| HistoryAction | String |
| TargetCode | String |
| QuantityAssembled | String |
| QuantityAssembledUoMNId | String |
| QuantityDisassembled | String |
| QuantityDisassembledUoMNId | String |
| UserId | String |
| ConsumptionDateTime | DateTime |
| EgPhaseNId | String |
| EgNId | String |
| _Id | String |
| WOId | String |
| WOOpId | String |
| DM_MTUId | String |
| CreatedOn | DateTime |
| WOStepNId | String |
| ToBeConsumedDMMaterialId | String |
| TargetNId | String |
| WOOpNId | String |
| TargetMaterialNId | String |
| IsAcquiredWithAGW | Boolean |
| Note | String |
| MaterialRevision | String |
| ToBeConsumedQuantity | String |
| ToBeCoProducedQuantity | String |
| ToBeCoProducedQuantityUoMNId | String |
| QuantityProduced | String |
| QuantityProducedUoMNId | String |
| ProductionDateTime | DateTime |
| Date | DateTime |
| Direction | String |
| MTUName | String |
| MaterialName | String |

#### Entity: RM_PC_ChangePackage

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| NId | String |
| WorkOrderOperationNId | String |
| OpenUserId | String |
| OpenTime | DateTime |
| CloseUserId | String |
| CloseTime | DateTime |
| Notes | String |
| StatusStateMachineNId | String |
| StatusStatusNId | String |
| WoOpPreviousStatusStateMachineNId | String |
| WoOpPreviousStatusStatusNId | String |
| WorkOrderOperationId | String |
| WorkOrderOperationName | String |
| WorkOrderOperationSequence | Integer |
| WorkOrderNId | String |
| WorkOrderName | String |
| WorkOrderId | String |
| _Id | String |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| CloseNotes | String |

#### Entity: RM_PC_LineSideMonitorInfo

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| LineSidePositionNId | String |
| IsManual | Boolean |
| BinsToBeRequested | String |
| ActualQtyValue | Integer/Decimal |
| ActualQtyUoMNId | String |
| EquipmentNId | String |
| MaterialNId | String |
| MaterialRevision | String |
| MaterialUId | String |
| Logic | String |
| CurrentPriority | String |
| CurrentStatus | String |
| KanbanCallId | String |
| KanbanCallUpdatedOn | DateTime |
| KanbanCallCreatedOn | DateTime |
| IsCreateEnabled | Boolean |
| IsCancelEnabled | Boolean |
| IsUpdPrioEnabled | Boolean |
| _Id | String |
| LineSideInstanceId | String |

#### Entity: LineSideInstanceHistory

**Published From:** Kanban

| Attribute | Type |
|-----------|------|
| _Id | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| EntityType | String |
| IsLocked | Boolean |
| ToBeCleaned | String |
| Action | String |
| LineSideInstance_Id | String |
| OldQty_UoMNId | String |
| OldQty_QuantityValue | Integer/Decimal |
| NewQty_UoMNId | String |
| NewQty_QuantityValue | Integer/Decimal |

#### Entity: RM_PC_ChangePackageDocument

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Description | String |
| MIMEType | String |
| IconId | String |
| RepositoryType | String |
| LocalFileId | String |
| FileName | String |
| IconName | String |
| Revision | String |
| EntityLinkId | String |
| Category | String |
| WorkOrderOperationNId | String |
| LinkedEntityId | String |
| ChangePackageNId | String |
| ActionNId | String |
| ChangePackageItemId | String |
| _Id | String |
| DocumentId | String |

#### Entity: RM_PC_ChangePackageSummary

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| ChangePackageNId | String |
| ActionNId | String |
| TypeNId | String |
| EntityNId | String |
| EntityUId | String |
| EntityRevision | String |
| EntityId | String |
| EntityName | String |
| StepName | String |
| StepSequence | Integer |
| StepUId | String |
| _Id | String |
| ChangePackagItemId | String |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| WorkOrderOperationId | String |
| ChPkgWorkOrderOperationId | String |
| WorkOrderOperationNId | String |
| WorkOrderOperationName | String |
| WorkOrderOperationSequence | Integer |
| WorkOrderStepId | String |
| ChPkgWorkOrderStepId | String |
| StepNId | String |

#### Entity: RM_PC_ChangePackageToBeUsedTool

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| tobeId | String |
| WorkOrderOperationId | String |
| WorkOrderStepId | String |
| WorkOrderOperationNId | String |
| NId | String |
| Name | String |
| Version | String |
| ActionNId | String |
| TimesToBeUsed | DateTime |
| ChangePackageNId | String |
| ChangePackageItemId | String |
| _Id | String |

#### Entity: RM_PC_ChangePackageMaterial

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| Name | String |
| MaterialId | String |
| MaterialNId | String |
| MaterialName | String |
| MaterialRevision | String |
| NId | String |
| GroupId | String |
| LogicalPosition | String |
| MaterialSpecificationType | String |
| OccurrenceId | String |
| Quantity | String |
| UoMNId | String |
| Sequence | Integer |
| WorkOrderOperationId | String |
| WorkOrderOperationNId | String |
| ChangePackageNId | String |
| ActionNId | String |
| ChangePackageItemId | String |
| _Id | String |
| ToBeConsumedMaterialId | String |

#### Entity: RM_PC_ChangePackageWorkInstruction

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| WorkOrderOperationId | String |
| WIAssociationTypeNId | String |
| WorkOrderOperationNId | String |
| OriginalWIDefinitionId | String |
| OriginalWIDefinitionNId | String |
| OriginalWIDefinitionRevision | String |
| OriginalWIDefinitionModel | String |
| NewWIDefinitionId | String |
| NewWIDefinitionNId | String |
| NewWIDefinitionRevision | String |
| NewWIDefinitionModel | String |
| ChangePackageNId | String |
| ActionNId | String |
| ChangePackageItemId | String |
| CreatedOn | DateTime |
| _Id | String |
| WIDefinitionToWOOperationAssociationId | String |
| ChangePackageId | String |

#### Entity: KanbanCallHistory

**Published From:** Kanban

| Attribute | Type |
|-----------|------|
| _Id | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| EntityType | String |
| IsLocked | Boolean |
| ToBeCleaned | String |
| InputMsgFileId | String |
| OutMsgId | String |
| IsManual | Boolean |
| Priority | String |
| KanbanCall_Id | String |
| Status_StateMachineNId | String |
| Status_StatusNId | String |
| Quantity_UoMNId | String |
| Quantity_QuantityValue | Integer/Decimal |

#### Entity: RM_PC_WorkOrder

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| ERPOrder | String |
| Sequence | Integer |
| ProductionTypeNId | String |
| Enterprise | String |
| Plant | String |
| Priority | String |
| DueDate | DateTime |
| StateMachineNId | String |
| StatusNId | String |
| EstimatedStartTime | DateTime |
| EstimatedEndTime | DateTime |
| ActualStartTime | DateTime |
| ActualEndTime | DateTime |
| InitialQuantity | String |
| ActualTargetQuantity | String |
| PlannedTargetQuantity | String |
| TargetQuantity | String |
| RealInitialQty | String |
| ProducedQuantity | String |
| ReworkedQuantity | String |
| ScrappedQuantity | String |
| ParentBatch | String |
| ParentOrderId | String |
| IsUnderScheduling | Boolean |
| AsPlannedBoPId | String |
| ConfirmationType | String |
| ToBeUsedBoMNId | String |
| ToBeUsedBoMVersion | String |
| ToBeUsedBillOfFeature | String |
| PBoPIdentId | String |
| PercentageOfCompletion | String |
| ProcessId | String |
| ProcessNId | String |
| ProcessRevision | String |
| ProcessUId | String |
| ReworkOfOrderId | String |
| RoutingId | String |
| SchedulingDate | DateTime |
| TargetProductId | String |
| ConfigurationParameter | String |
| Notes | String |
| DM_MaterialLogisticClassNId | String |
| MaterialNId | String |
| MaterialName | String |
| MaterialRevision | String |
| MaterialUoMNId | String |
| AsPlannedBOPIsOutOfDate | DateTime |
| NewAsPlannedBOP | String |
| _Id | String |
| WorkOrderId | String |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| DM_MaterialId | String |
| MaterialId | String |
| NewAsPlannedBOPExists | String |
| EndItem | String |
| UnitEffectivity | String |
| PreviousStatusNId | String |

#### Entity: RM_PC_ChangePackageWorkOrderStep

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Sequence | Integer |
| Description | String |
| Priority | String |
| ToBeCollectedDocument | String |
| EstimatedDuration | String |
| EstimatedStartTime | DateTime |
| EstimatedEndTime | DateTime |
| ActualStartTime | DateTime |
| ElectronicSignatureStart | String |
| ElectronicSignaturePause | String |
| ElectronicSignatureComplete | String |
| StepNId | String |
| StepRevision | String |
| StepUId | String |
| WorkOrderOperationNId | String |
| ChangePackageNId | String |
| ChangePackageItemId | String |
| ActionNId | String |
| StatusNId | String |
| OperationStepCategoryNId | String |
| IsChangePackageItemWOStep | Boolean |
| _Id | String |
| WOStepId | String |
| WorkOrderOperationId | String |
| ChangePackageId | String |

#### Entity: RM_PC_ChangePackageMachine

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| ToBeUsedMachineName | String |
| ToBeUsedMachineNId | String |
| EquipmentNId | String |
| EquipmentLevel | String |
| EquipmentTypeNId | String |
| WorkOrderOperationId | String |
| WorkOrderOperationNId | String |
| ChangePackageNId | String |
| ActionNId | String |
| ChangePackageItemId | String |
| _Id | String |
| ToBeUsedMachineId | String |

#### Entity: RM_PC_ChangePackageQualityInspection

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| CharacteristicRepresentationNId | String |
| CharacteristicRepresentationRevision | String |
| CharacteristicRepresentationName | String |
| FrequencyNId | String |
| RuntimeCharacteristicRepresentationId | String |
| WorkOrderOperationNId | String |
| ChangePackageNId | String |
| ActionNId | String |
| ChangePackageItemId | String |
| _Id | String |
| WorkOrderOperationId | String |

#### Entity: RM_PC_ChangePackageStepDocument

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Description | String |
| MIMEType | String |
| IconId | String |
| RepositoryType | String |
| LocalFileId | String |
| FileName | String |
| IconName | String |
| Revision | String |
| EntityLinkId | String |
| Category | String |
| WorkOrderStepNId | String |
| LinkedEntityId | String |
| ChangePackageNId | String |
| ActionNId | String |
| ChangePackageItemId | String |
| IsChangePackageItemWOStep | Boolean |
| _Id | String |
| DocumentId | String |

#### Entity: RM_PC_ChangePackageStepToBeUsedTool

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| tobeId | String |
| WorkOrderStepId | String |
| NId | String |
| WorkOrderStepNId | String |
| Name | String |
| Version | String |
| ActionNId | String |
| TimesToBeUsed | DateTime |
| ChangePackageNId | String |
| ChangePackageItemId | String |
| IsChangePackageItemWOStep | Boolean |
| _Id | String |

#### Entity: RM_PC_ChangePackageStepMaterial

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| Name | String |
| MaterialId | String |
| MaterialNId | String |
| MaterialName | String |
| MaterialRevision | String |
| NId | String |
| GroupId | String |
| LogicalPosition | String |
| MaterialSpecificationType | String |
| OccurrenceId | String |
| Quantity | String |
| UoMNId | String |
| Sequence | Integer |
| WorkOrderOperationStepId | String |
| WorkOrderOperationStepNId | String |
| ChangePackageNId | String |
| ActionNId | String |
| ChangePackageItemId | String |
| IsChangePackageItemWOStep | Boolean |
| _Id | String |
| ToBeConsumedMaterialId | String |

#### Entity: RM_PC_ChangePackageStepWorkInstruction

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| WorkOrderStepId | String |
| WIAssociationTypeNId | String |
| WorkOrderStepNId | String |
| OriginalWIDefinitionId | String |
| OriginalWIDefinitionNId | String |
| OriginalWIDefinitionRevision | String |
| OriginalWIDefinitionModel | String |
| NewWIDefinitionId | String |
| NewWIDefinitionNId | String |
| NewWIDefinitionRevision | String |
| NewWIDefinitionModel | String |
| ChangePackageNId | String |
| ActionNId | String |
| ChangePackageItemId | String |
| CreatedOn | DateTime |
| IsChangePackageItemWOStep | Boolean |
| _Id | String |
| WIDefinitionToWOStepAssociationId | String |
| ChangePackageId | String |

#### Entity: RM_PC_ChangePackageStepQualityInspection

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| CharacteristicRepresentationNId | String |
| CharacteristicRepresentationRevision | String |
| CharacteristicRepresentationName | String |
| FrequencyNId | String |
| RuntimeCharacteristicRepresentationId | String |
| WorkOrderStepNId | String |
| ChangePackageNId | String |
| ActionNId | String |
| ChangePackageItemId | String |
| IsChangePackageItemWOStep | Boolean |
| _Id | String |
| WorkOrderStepId | String |

#### Entity: RM_PC_DM_MaterialAndMaterial

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| MaterialClass | String |
| NId | String |
| Name | String |
| Revision | String |
| SerialNumberProfile | String |
| UoMNId | String |
| _Id | String |
| DM_MatId | String |

#### Entity: EXDSAsBuiltSN

**Published From:** AppU4DM

| Attribute | Type |
|-----------|------|
| IsKO | Boolean |
| MtuNId | String |
| Code | String |
| MtuDescription | String |
| MtuName | String |
| CodeType | String |
| MtuStatus | String |
| Quantity | String |
| UomNId | String |
| WeightUoMNId | String |
| WeightQty | String |
| VolumeUoMNId | String |
| VolumeQty | String |
| DMMatNId | String |
| MaterialRevision | String |
| _Id | String |
| WOId | String |
| IsFrozen | Boolean |

#### Entity: RM_PC_GetEquipmentForToBeUsedMachine

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| EquipmentToBeLinkedId | String |
| EquipmentToBeLinked | String |
| EquipmentToBeLinkedName | String |
| EquipmentToBeLinkedLevel | String |
| OperationId | String |
| _Id | String |

#### Entity: RM_PC_DM_MaterialTrackingUnitHistoryWithATN

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| DM_MaterialTrackingUnitId | String |
| DM_MaterialTrackingUnitCode | String |
| TargetDM_MaterialTrackingUnitId | String |
| TargetDM_MaterialTrackingUnitCode | String |
| ATNNodeId | String |
| ATNParameterName | String |
| ATNParameterQualityStatus | String |
| ATNParameterTimestamp | DateTime |
| ATNParameterValue | Integer/Decimal |
| DSParameterNId | String |
| AcquiredFromShopfloorDate | DateTime |
| _Id | String |

#### Entity: ItlkCheckWOOpAssociation

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
| IsInbound | Boolean |
| IsOutbound | Boolean |
| ItlkCheck_Id | String |
| WorkOrderOperation_Id | String |

#### Entity: WorkOrderHistory

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
| _Context | String |
| DNCItem | String |
| Document | Boolean |
| Equipment | String |
| Message | String |
| PrintJobFile | String |
| ProcessUId | String |
| UserId | String |
| SerialNumber | String |
| ExecutionGroupPhaseId | String |
| BaseLine | String |
| OldBaseLine | String |
| SourceWorkOrderNId | String |
| TargetWorkOrderNId | String |
| Notes | String |
| Status | String |
| PartProgramName | String |
| Date | DateTime |
| Team_Id | String |
| Action_Id | String |
| WorkOrderOperation_Id | String |
| WorkOrderStep_Id | String |
| WorkOrder_Id | String |
| ReferencedWorkOrderHistory_Id | String |

#### Entity: ItlkCheckWOOpAssociationParameterValue

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
| ParameterValue | Integer/Decimal |
| ItlkCheckParameter_Id | String |
| ItlkCheckWOOpAssociation_Id | String |

#### Entity: WorkOrderHistoryMaterialItem

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
| Quantity | String |
| WorkOrderHistory_Id | String |
| DM_MaterialTrackingUnit_Id | String |

#### Entity: WorkInstruction

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
| Name | String |
| Description | String |
| WorkInstructionDefinitionNId | String |
| WorkInstructionDefinitionRevision | String |
| IsReEditEnabled | Boolean |
| Design_Model | String |
| Status_StateMachineNId | String |
| Status_StatusNId | String |

#### Entity: WorkOrderHistoryAction

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

#### Entity: WorkInstructionToExecutionGroupPhaseAssociation

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
| ExecutionGroupPhase_Id | String |
| WorkInstruction_Id | String |
| DM_MaterialTrackingUnit_Id | String |

#### Entity: RM_PC_DM_MaterialTrackingUnitAndMTU

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| Name | String |
| MaterialDefinition | String |
| NId | String |
| Code | String |
| CodeType | String |
| StatusNId | String |
| WOId | String |
| ProductionType | String |
| _Id | String |
| IsResultKo | Boolean |
| ProducedId | String |
| MTUId | String |

#### Entity: WorkInstructionToWorkOrderStepAssociation

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
| WorkOrderStep_Id | String |
| WorkInstruction_Id | String |
| DM_MaterialTrackingUnit_Id | String |

#### Entity: RM_PC_GetAllDM_MaterialTrackingUnitToAssignWorkOrder

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| Name | String |
| Code | String |
| NId | String |
| Description | String |
| Status | String |
| Quantity | String |
| MaterialNId | String |
| IsReserved | Boolean |
| DM_MaterialId | String |
| WorkOrderId | String |
| _Id | String |
| IsResultKo | Boolean |

#### Entity: WorkInstructionToWorkOrderOperationAssociation

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
| HideOnExecGroup | String |
| WIDefToWOOpAssociationId | String |
| WorkOrderOperation_Id | String |
| WorkInstruction_Id | String |
| DM_MaterialTrackingUnit_Id | String |

#### Entity: RM_PC_GetWIDetailsForWorkOrderOperationAndStep

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| WorkInstructionDefinitionNId | String |
| WorkInstructionDefinitionName | String |
| WorkInstructionDefinitionRevision | String |
| SectionName | String |
| StepName | String |
| StepType | String |
| DCName | String |
| DCValue | Integer/Decimal |
| DCDefaultValue | Integer/Decimal |
| DCLowerLimit | String |
| DCUpperLimit | String |
| DCUOM | String |
| AssociationType | String |
| WorkOrderOperationId | String |
| WorkOrderStepId | String |
| _Id | String |
| AssociationId | String |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| HideOnExecGroup | String |
| IsShared | Boolean |
| WorkInstructionDefinitionId | String |

#### Entity: RM_PC_GetAlternativeGroupsFromWorkOrderId

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| WorkOrderId | String |
| FromId | String |
| AlternativeGroup | String |
| PreferredWOOpNId | String |
| ParentWOOpNId | String |
| _Id | String |
| PreferredWOOpId | String |
| ParentWOOpId | String |

#### Entity: RM_PC_GetExecutionGroupsForWorkOrder

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Description | String |
| EstimatedDuration | String |
| EstimatedEndTime | DateTime |
| EstimatedStartTime | DateTime |
| TotalQuantity | String |
| Status | String |
| WorkOrderId | String |
| _Id | String |
| ExecutionGroupId | String |

#### Entity: RM_PC_GetDataFromFeatures

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| FeatureNId | String |
| FeatureName | String |
| FeatureDescription | String |
| FeatureValueNId | String |
| FeatureValueName | Integer/Decimal |
| FeatureValueDescription | Integer/Decimal |
| IsOption | Boolean |
| BillOfFeaturesNId | String |
| BillOfFeaturesRevision | String |
| _Id | String |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |

#### Entity: WorkInstructionSection

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
| Title | String |
| CompletedBy | String |
| Sequence | Integer |
| IsCompleted | Boolean |
| WorkInstruction_Id | String |

#### Entity: RM_PC_ItlkCheckWOOpAssociation

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| Name | String |
| Description | String |
| IsInbound | Boolean |
| IsOutbound | Boolean |
| WorkOrderOperationId | String |
| IsStandard | Boolean |
| _Id | String |
| AssociationId | String |
| NId | String |

#### Entity: WorkInstructionParameter

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
| ParameterValue | Integer/Decimal |
| ParameterUoMNId | String |
| IsReadOnly | Boolean |
| Direction | String |
| ParameterType | String |
| WorkInstruction_Id | String |

#### Entity: RM_PC_WorkOrderDependencies

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| DependencyType | String |
| FromWOFolderId | String |
| FromWOFolderNId | String |
| FromWOFolderName | String |
| ToWOFolderId | String |
| ToWOFolderNId | String |
| ToWOFolderName | String |
| FromWOOpId | String |
| FromWOOpNId | String |
| FromWOOpName | String |
| ToWOOpId | String |
| ToWOOpNId | String |
| ToWOOpName | String |
| DestinationWorkOrder | String |
| PartialCompleted | String |
| SourceWorkOrder | String |
| DestinationWorkOrderName | String |
| SourceWorkOrderName | String |
| WorkOrderId | String |
| _Id | String |
| DependencyId | String |

#### Entity: WorkInstructionStep

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
| Title | String |
| InstanceNumber | String |
| Instructions | String |
| AcknowledgeOn | String |
| CompletedBy | String |
| Sequence | Integer |
| IsSignaturePending | Boolean |
| ScenarioInstanceId | String |
| AvoidAcknowledge | String |
| _Type | String |
| IsCompleted | Boolean |
| WorkInstructionSection_Id | String |
| ScenarioConfiguration_NId | String |
| ScenarioConfiguration_ActionLabel | String |

#### Entity: RM_PC_CharacteristicRepresentation

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| NId | String |
| SampleSize | String |
| CanBeSkipped | String |
| Name | String |
| ReferencedOperationNId | String |
| ReferencedOperationName | String |
| ReferencedOperationSequence | Integer |
| Frequency | String |
| FrequencyValue | Integer/Decimal |
| IsWOSpecific | Boolean |
| TimeBasedScheduleMode | DateTime |
| TimeBasedMandatoryInspectionWarningValue | DateTime |
| TimeBasedMandatoryInspectionWarningUoM | DateTime |
| TimeBasedMandatoryInspectionWarningUoMValue | DateTime |
| OperationID | String |
| StepID | String |
| AsPlannedBOPId | String |
| _Id | String |
| RtChrReprId | String |
| ToBeUsedInspectionId | String |
| GetFromContainer | String |
| Revision | String |
| FrequencyPercentage | String |

#### Entity: WorkInstructionStepItem

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
| ItemNId | String |
| StepNId | String |
| ItemValue | Integer/Decimal |
| DataType | String |
| Binding | String |
| Label | String |
| Caption | String |
| UIControl | String |
| IsRequired | Boolean |
| IsReadOnly | Boolean |
| UoM | String |
| Format | String |
| DefaultValue | Integer/Decimal |
| ExecutionDate | DateTime |
| CompletedBy | String |
| IsModifiable | Boolean |
| IsForceBlank | Boolean |
| IsOutOfSpec | Boolean |
| Length | String |
| LastValueUpdatedOn | DateTime |
| Placeholder | String |
| IsCompleted | Boolean |
| InstanceNumber | String |
| AcquisitionBehavior | String |
| WorkInstructionStep_Id | String |

#### Entity: RM_PC_ProcessMeasurement

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| Equipment | String |
| MaterialTrackingUnitNId | String |
| WorkOrderOperationNId | String |
| TimeStamp | DateTime |
| Property | String |
| DataType | String |
| Value | Integer/Decimal |
| ValueUoM | Integer/Decimal |
| MinValue | Integer/Decimal |
| MinValueUoM | Integer/Decimal |
| MaxValue | Integer/Decimal |
| MaxValueUoM | Integer/Decimal |
| TargetValue | Integer/Decimal |
| TargetValueUoM | Integer/Decimal |
| Description | String |
| Notes | String |
| CorrelationId | String |
| AcquisitionDate | DateTime |
| WorkOrderId | String |
| WorkOrderNId | String |
| WorkOrderOperationId | String |
| GroupId | String |
| _Id | String |
| ProductionEvent | String |
| Tool | String |
| Category | String |
| MaterialTrackingUnitCode | String |

#### Entity: WorkInstructionStepItemLimit

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
| ItemNId | String |
| Target | String |
| TargetBinding | String |
| LowLimit | String |
| LowLimitBinding | String |
| HighLimit | String |
| HighLimitBinding | String |
| AlertMin | Integer/Decimal |
| AlertMinBinding | Integer/Decimal |
| CautionMin | Integer/Decimal |
| CautionMinBinding | Integer/Decimal |
| CautionMax | Integer/Decimal |
| CautionMaxBinding | Integer/Decimal |
| AlertMax | Integer/Decimal |
| AlertMaxBinding | Integer/Decimal |
| OutOfLimitBhv | String |
| WorkInstructionStepItem_Id | String |

#### Entity: ItlkCheckWOStepAssociation

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
| IsInbound | Boolean |
| IsOutbound | Boolean |
| ItlkCheck_Id | String |
| WorkOrderStep_Id | String |

#### Entity: ItlkCheckWOStepAssociationParameterValue

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
| ParameterValue | Integer/Decimal |
| ItlkCheckParameter_Id | String |
| ItlkCheckWOStepAssociation_Id | String |

#### Entity: ItlkCheckHistory

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
| WorkOrderOperationId | String |
| WorkOrderOperationNId | String |
| WorkOrderStepId | String |
| WorkOrderStepNId | String |
| MTUCode | String |
| Quantity | String |
| EquipmentNId | String |
| User | String |
| OperationStepCategoryId | String |
| OperationStepCategoryNId | String |
| ItlkCheckSessionId | String |
| DM_MTUId | String |
| MTUCodeType | String |
| ItlkCheckResult | String |
| Date | DateTime |
| ItlkCheckOperationStepCategoryAssociation_Id | String |
| ItlkCheckWOOpAssociation_Id | String |
| ItlkCheckWOStepAssociation_Id | String |

#### Entity: ExecutionGroupPhase

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
| Name | String |
| Sequence | Integer |
| EstimatedStartTime | DateTime |
| EstimatedEndTime | DateTime |
| PreferredMachineNId | String |
| ToBeCollectedDocument | String |
| EstimatedDuration_Ticks | String |
| ExecutionGroup_Id | String |
| Status_StateMachineNId | String |
| Status_StatusNId | String |

#### Entity: ExecutionGroupPhaseWOOAssociation

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
| Quantity | String |
| ExecutionGroupPhase_Id | String |
| WorkOrderOperation_Id | String |

#### Entity: ExecutionGroup

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
| Name | String |
| Description | String |
| EstimatedStartTime | DateTime |
| EstimatedEndTime | DateTime |
| IsUnderScheduling | Boolean |
| SchedulingDate | DateTime |
| DueDate | DateTime |
| EstimatedDuration_Ticks | String |
| Status_StateMachineNId | String |
| Status_StatusNId | String |

#### Entity: RM_PC_ItlkCheckWOStepAssociation

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Description | String |
| IsStandard | Boolean |
| IsInbound | Boolean |
| IsOutbound | Boolean |
| WorkOrderStepId | String |
| _Id | String |
| AssociationId | String |

#### Entity: ActualProducedMaterial

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
| WaitingForQualityQuantity | String |
| UnderQualityInspectionQuantity | String |
| WorkOrderOperations_Id | String |
| ToBeProducedMaterial_Id | String |
| DM_MaterialTrackingUnit_Id | String |
| WorkOrderStep_Id | String |

#### Entity: RM_PC_ActualUsedTool

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| ToolNId | String |
| ToolDescription | String |
| ToolName | String |
| ToolDefinitionName | String |
| UsedTimes | DateTime |
| CurentThickness | String |
| ActualUsageCounter | Integer |
| AmountSetForThicknessDecrease | Integer/Decimal |
| TreatmentCount | Integer |
| MaterialTrackingUnitCode | String |
| ToolHistoryUsageDurationTicks | String |
| QualityGateNId | String |
| QGEquipmentNId | String |
| ActivityNId | String |
| WorkOrderOperationSequence | Integer |
| WorkOrderStepSequence | Integer |
| WorkOrderId | String |
| WorkOrderOperationId | String |
| WorkOrderOperationNId | String |
| WorkOrderStepId | String |
| WorkOrderStepNId | String |
| ToolId | String |
| ToolDefinitionId | String |
| ToBeUsedToolId | String |
| MaterialTrackingUnitId | String |
| DMMaterialTrackingUnitId | String |
| ActionNId | String |
| _Id | String |
| WorkOrderOperationName | String |
| WorkOrderStepName | String |
| LogisticClassNId | String |

#### Entity: RM_PC_ToolHistory

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| ToolDefinitionName | String |
| ToolNId | String |
| ToolName | String |
| UsedToolCount | Integer |
| ExpirationDate | DateTime |
| User | String |
| DMMaterialTrackingUnitId | String |
| ToolId | String |
| WorkOrderOperationId | String |
| WorkOrderStepId | String |
| _Id | String |
| Date | DateTime |

#### Entity: RM_PC_ToolHistoryDetail

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| ToolDefinitionName | String |
| ToolNId | String |
| ToolName | String |
| UsedToolCount | Integer |
| ExpirationDate | DateTime |
| User | String |
| Discriminator | Integer/Decimal |
| DMMaterialTrackingUnitId | String |
| ToolId | String |
| WorkOrderOperationId | String |
| WorkOrderStepId | String |
| ATNACNodeId | String |
| ATNACParameterName | String |
| ATNACParameterQuality | String |
| ATNACParameterValue | Integer/Decimal |
| ATNACParameterTS | String |
| ATNACFromShopFloor | String |
| ATNTRNodeId | String |
| ATNTRParameterName | String |
| ATNTRSetPointItemNId | String |
| ATNTRSetPointItemValue | Integer/Decimal |
| ATNTRParameterQuality | String |
| ATNTRParameterTS | String |
| ATNTRSetPointNId | String |
| ATNTRTransmissionSucceded | String |
| CONTContainerNId | String |
| QGQualityGateNId | String |
| _Id | String |
| ToolHistoryId | String |

#### Entity: RM_PC_DM_MaterialTrackingUnitForReOpen

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| MaterialTrackingUnitNId | String |
| MaterialTrackingUnitName | String |
| MaterialTrackingUnitCode | String |
| MaterialNId | String |
| WorkOrderOperationId | String |
| WorkOrderStepId | String |
| CompletedQuantity | String |
| NotWorkedQuantity | String |
| ReopenedQuantity | String |
| Quantity | String |
| _Id | String |

#### Entity: RM_PC_WorkOrderHistoryATNTRSetPoint

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| ATNTRNodeId | String |
| ATNTRParameterName | String |
| ATNTRParameterQuality | String |
| ATNTRParameterTimestamp | DateTime |
| ATNTRSetPointItemNId | String |
| ATNTRSetPointItemValue | Integer/Decimal |
| ATNTRSetPointNId | String |
| ATNTRTransmissionSucceeded | String |
| Date | DateTime |
| Equipment | String |
| SerialNumber | String |
| User | String |
| WorkOrderOperationId | String |
| WorkOrderStepId | String |
| WorkOrderId | String |
| Discriminator | Integer/Decimal |
| _Id | String |
| WorkOrderHistoryId | String |

#### Entity: RM_PC_QualityInspectionSamples

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| WorkOrderOperationNId | String |
| WorkOrderOperationName | String |
| WorkOrderOperationSequence | Integer |
| WorkOrderStepNId | String |
| WorkOrderStepName | String |
| WorkOrderStepSequence | Integer |
| MaterialTrackingUnitNId | String |
| MeasuredAttributiveValue | Integer/Decimal |
| MeasuredVariableValue | Integer/Decimal |
| VisualFailureNId | String |
| VisualFailureCount | Integer |
| User | String |
| Timestamp | DateTime |
| EquipmentNId | String |
| AnyViolation | String |
| CharacteristicRepresentationNId | String |
| InspectionValueId | String |
| _Id | String |
| InspectionSampleId | String |
| WorkOrderNId | String |
| WorkOrderId | String |
| MaterialTrackingUnitCode | String |
| MaterialNId | String |
| MaterialRevision | String |
| CharacteristicType | String |

#### Entity: RM_PC_SnagAndNote

**Published From:** RMEXDSProdCoord

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
| WorkOrderOperationNId | String |
| _Id | String |
| SnagAndNoteId | String |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| WorkOrderOperationSequence | Integer |
| WorkOrderOperationName | String |

#### Entity: ChangeOperation

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
| NonConformance_Id | String |
| CloseNotes | String |
| NewOperationName | String |
| Description | String |
| EquipmentName | String |
| OperationNameList | String |
| RefNumber | String |
| RoutingOperationAfter | String |
| SequenceAfter | Integer |
| OldOperationName | String |
| ChangeType | String |

#### Entity: RM_PC_WorkOrderStep

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| ActualEndTime | DateTime |
| ActualStartTime | DateTime |
| Description | String |
| Optional | String |
| EstimatedDuration | String |
| EstimatedEndTime | DateTime |
| EstimatedStartTime | DateTime |
| ExecutionDuration | String |
| NId | String |
| Name | String |
| PauseDuration | String |
| Priority | String |
| Sequence | Integer |
| StepNId | String |
| WorkOrderOperationId | String |
| Status | String |
| WorkOrderId | String |
| HasInterlockingCheck | Boolean |
| _Id | String |
| WorkOrderOperationNId | String |
| WorkOrderOperationName | String |
| OperationStepCategoryId | String |

#### Entity: RM_PC_WorkOrderRuntimeDocument

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| WorkOrderId | String |
| WorkOrderOperationNId | String |
| WorkOrderOperationName | String |
| WorkOrderStepId | String |
| WorkOrderStepNId | String |
| WorkOrderStepName | String |
| DM_MaterialTrackingUnitId | String |
| MaterialTrackingUnitNId | String |
| MaterialTrackingUnitCode | String |
| ExecutionGroupPhaseId | String |
| egPhaseNId | String |
| egPhaseName | String |
| egNId | String |
| Document | Boolean |
| NId | String |
| Name | String |
| Description | String |
| MIMEType | String |
| FileName | String |
| IconName | String |
| RepositoryType | String |
| IconId | String |
| LocalFileId | String |
| Revision | String |
| SourceRevision | String |
| IsCurrent | Boolean |
| _Id | String |
| ActualCollectedDocumentId | String |
| CreatedOn | DateTime |
| WorkOrderOperationId | String |

#### Entity: ChangeTool

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
| NonConformance_Id | String |
| NewToolDefinitionName | String |
| RefNumber | String |
| CloseNotes | String |
| ToolUsageNumber | String |
| ChangeType | String |

#### Entity: ChangePart

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
| NonConformance_Id | String |
| CloseNotes | String |
| LogicalPosition | String |
| NewPart | String |
| NewPartQuantity | String |
| OldPart | String |
| OldPartQuantity | String |
| RefNumber | String |
| NewPartNId | String |
| NewMaterialSpecificationType | String |
| MaterialSpecificationType | String |
| NewLogicalPosition | String |
| OldPartNId | String |
| ChangeType | String |

#### Entity: ScrewingToolHistory

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
| ToolNId | String |
| User | String |
| UsageDate | DateTime |
| DM_MTUId | String |
| MTUNId | String |
| MTUCode | String |
| Reason | String |
| Torque | String |
| Angle | String |
| IsOk | Boolean |
| ItemId | String |
| ToBeUsedTool_Id | String |

#### Entity: RM_PC_WorkOrderHistoryAsBuilt

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| WorkOrderId | String |
| WorkOrderOperationId | String |
| WorkOrderStepId | String |
| Code | String |
| ExecutionGroupNId | String |
| ExecutionGroupPhaseNId | String |
| WorkOrderOperationSequence | Integer |
| WorkOrderOperationName | String |
| WorkOrderStepSequence | Integer |
| WorkOrderStepName | String |
| User | String |
| TeamNId | String |
| TeamName | String |
| ProcessUId | String |
| Message | String |
| Equipment | String |
| ContainerNId | String |
| BaseLine | String |
| OldBaseLine | String |
| PrintJobFile | String |
| PartProgramName | String |
| Notes | String |
| Status | String |
| SerialNumber | String |
| ActionNId | String |
| Quantity | String |
| QuantityUoMNId | String |
| IsKO | Boolean |
| ResultValue | Integer/Decimal |
| ElectronicSignatureComplete | String |
| ElectronicSignaturePause | String |
| ElectronicSignatureStart | String |
| WorkOrderOperationNId | String |
| WorkOrderStepNId | String |
| _Id | String |
| Date | DateTime |
| DNCItem | String |
| Document | Boolean |
| _Context | String |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| WorkOrderHistoryId | String |

#### Entity: ChangeDataCollection

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
| NonConformance_Id | String |
| CloseNotes | String |
| ReferenceWorkorderOperations | String |
| RefNumber | String |
| WIDefinition | String |
| ChangeType | String |

#### Entity: RM_PC_WorkInstructionAsBuilt

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| WorkInstructionNId | String |
| WorkInstructionName | String |
| WorkInstructionDescription | String |
| WorkInstructionStatus | String |
| WorkInstructionIsReEditEnabled | Boolean |
| HideOnExecGroup | String |
| WorkOrderOperationId | String |
| WorkOrderStepId | String |
| DM_MaterialTrackingUnitId | String |
| MaterialTrackingUnitNId | String |
| MaterialTrackingUnitCode | String |
| SectionId | String |
| SectionNId | String |
| SectionTitle | String |
| StepNId | String |
| StepType | String |
| StepCompletedBy | String |
| StepIsCompleted | String |
| StepAcknowledgeOn | String |
| StepLastUpdatedOn | DateTime |
| ItemDefaultValue | Integer/Decimal |
| ItemValue | Integer/Decimal |
| ItemIsCompleted | String |
| ItemLastValueUpdatedOn | DateTime |
| ItemLastUpdatedOn | DateTime |
| ItemNId | String |
| ItemLabel | String |
| ItemUoM | String |
| LimitItemNId | String |
| LimitTarget | String |
| LimitTargetBinding | String |
| LowLimit | String |
| LowLimitBinding | String |
| HighLimit | String |
| HighLimitBinding | String |
| LimitAlertMin | Integer/Decimal |
| LimitAlertMinBinding | Integer/Decimal |
| LimitCautionMin | Integer/Decimal |
| LimitCautionMinBinding | Integer/Decimal |
| LimitCautionMax | Integer/Decimal |
| LimitCautionMaxBinding | Integer/Decimal |
| LimitAlertMax | Integer/Decimal |
| LimitAlertMaxBinding | Integer/Decimal |
| OutOfLimitBhv | String |
| WorkOrderOperationNId | String |
| WorkOrderStepNId | String |
| WorkOrderId | String |
| _Id | String |
| AssociationId | String |
| WorkInstructionId | String |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| WorkOrderOperationName | String |
| WorkOrderStepName | String |
| WorkOrderOperationSequence | Integer |
| WorkOrderStepSequence | Integer |

#### Entity: ChangeRouting

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
| NonConformance_Id | String |
| CloseNotes | String |
| RefNumber | String |
| NewOperationFrom | String |
| NewOperationTo | String |
| NewDependency | String |
| OldOperationFrom | String |
| OldOperationTo | String |
| OldDependency | String |
| ChangeType | String |

#### Entity: RM_PC_WorkOrderOperationUser

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| EquipmentNId | String |
| WorkOrderOperationId | String |
| User | String |
| _Id | String |
| UserWoopAssociationId | String |
| EquipmentType | String |
| WOOpStatus | String |

#### Entity: RM_PC_WorkOrderOperationSkill

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| EquipmentNId | String |
| WorkOrderOperationId | String |
| SkillNId | String |
| SkillName | String |
| SkillDescription | String |
| SkillColor | String |
| SkillWoopAssociationLevel | String |
| EquipmentType | String |
| WOOpStatus | String |
| _Id | String |

#### Entity: RM_PC_UserNonProductiveActivity

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| User | String |
| NpaNId | String |
| StartDate | DateTime |
| EndDate | DateTime |
| NpaContextNId | String |
| _Id | String |
| UserNpaId | String |
| NPAContextValue | Integer/Decimal |
| NPAContextValueIdentifier | Integer/Decimal |
| NpaDescription | String |

#### Entity: RM_PC_UserProductiveActivity

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| ActionNId | String |
| UserId | String |
| Message | String |
| BaseLine | String |
| OldBaseLine | String |
| Date | DateTime |
| TeamNId | String |
| TeamName | String |
| WorkOrderNId | String |
| WorkOrderName | String |
| WorkOrderOperationNId | String |
| WorkOrderOperationName | String |
| WorkOrderStepNId | String |
| WorkOrderStepName | String |
| EgNId | String |
| EgName | String |
| EgPhaseNId | String |
| EgPhaseName | String |
| EgPhaseSequence | Integer |
| _Id | String |
| ContainerNId | String |
| ContainerQuantity | String |
| ResultValue | Integer/Decimal |
| IsResultKO | Boolean |
| Quantity | String |
| QuantityUoMNId | String |
| WorkOrderOperationSequence | Integer |
| WorkOrderStepSequence | Integer |
| SerialNumber | String |
| Equipment | String |
| Team | String |
| ProcessNId | String |

#### Entity: NonProductiveActivity

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
| Description | String |

#### Entity: NonProductiveActivityContext

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

#### Entity: RM_PC_ProducedMaterialItem

**Published From:** RMEXDSProdCoord

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
| WorkOrderId | String |

#### Entity: RM_PC_Buffer

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Description | String |
| EquipmentNId | String |
| EquipmentName | String |
| MaxQty | Integer/Decimal |
| MaxQtyUoMNId | String |
| MaxVolume | Integer/Decimal |
| MaxVolumeUoMNId | String |
| MaxWeight | Integer/Decimal |
| MaxWeightUoMNId | String |
| Status | String |
| ValidityFrom | String |
| ValidityTo | String |
| BufferDefinitionNId | String |
| BufferDefinitionVersion | String |
| BufferDefinitionName | String |
| BufferDefinitionIsValid | String |
| CapacityTypeNId | String |
| _Id | String |
| BufferId | String |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| StateMachine | String |
| BufferDefinitionDescription | String |

#### Entity: RM_PC_BufferMaterialItemAssociation

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| DM_MaterialTrackingUnitId | String |
| BufferId | String |
| MaterialTrackingUnitNId | String |
| MaterialTrackingUnitCode | String |
| MaterialTrackingUnitCodeType | String |
| MaterialTrackingUnitName | String |
| MaterialTrackingUnitDescription | String |
| MaterialTrackingUnitQuantity | String |
| MaterialTrackingUnitQuantityUoM | String |
| MaterialTrackingUnitStatus | String |
| DMMaterialTrackingUnitVolume | String |
| DMMaterialTrackingUnitVolumeUoM | String |
| DMMaterialTrackingUnitWeight | String |
| DMMaterialTrackingUnitWeightUoM | String |
| DMMaterialTrackingUnitActiveNCNumber | String |
| MaterialNId | String |
| MaterialName | String |
| MaterialRevision | String |
| _Id | String |
| BufferDMMTUAssociationId | String |
| MaterialTrackingUnitId | String |
| DMMaterialId | String |
| MaterialId | String |

#### Entity: RM_PC_BufferHistory

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| BufferId | String |
| UserId | String |
| ActionNId | String |
| Date | DateTime |
| MTUCode | String |
| MTUNId | String |
| Quantity | String |
| UoMNId | String |
| _Id | String |
| BufferHistoryId | String |

#### Entity: BufferDefinition

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
| Name | String |
| Description | String |
| Version | String |
| IsValid | String |
| CapacityType_Id | String |
| MaxQty_UoMNId | String |
| MaxQty_QuantityValue | Integer/Decimal |
| MaxWeight_UoMNId | String |
| MaxWeight_QuantityValue | Integer/Decimal |
| MaxVolume_UoMNId | String |
| MaxVolume_QuantityValue | Integer/Decimal |

#### Entity: RuntimeLogisticWhitelist

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
| CapacityType | String |
| ContainedEntityType | String |
| ContainerEntityType | String |
| ToolDefinition | String |
| LogisticClass_Id | String |
| DM_Material_Id | String |
| Buffer_Id | String |
| MaxQty_UoMNId | String |
| MaxQty_QuantityValue | Integer/Decimal |
| MaxVolume_UoMNId | String |
| MaxVolume_QuantityValue | Integer/Decimal |
| MaxWeight_UoMNId | String |
| MaxWeight_QuantityValue | Integer/Decimal |
| Min_UoMNId | String |
| Min_QuantityValue | Integer/Decimal |
| Target_UoMNId | String |
| Target_QuantityValue | Integer/Decimal |
| Threshold_UoMNId | String |
| Threshold_QuantityValue | Integer/Decimal |

#### Entity: FutureHold

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
| CloseTime | DateTime |
| CloseUserId | String |
| Comment | String |
| IsPresent | Boolean |
| Lifecycle | String |
| OpenTime | DateTime |
| OpenUserId | String |
| PreviousStatus | String |
| ReasonId | String |
| WorkOrderNId | String |
| WorkOrderOperationNId | String |

#### Entity: HandlingUnit

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
| LocationNId | String |
| NId | String |
| TransportOperation_Id | String |
| Status_StateMachineNId | String |
| Status_StatusNId | String |
| Weight_UoMNId | String |
| Weight_QuantityValue | Integer/Decimal |

#### Entity: HandlingUnitHistory

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
| BatchId | String |
| Date | DateTime |
| DestinationBufferNId | String |
| LocationNId | String |
| HandlingUnitNId | String |
| SerialNumber | String |
| SourceBufferNId | String |
| StatusNId | String |
| TransportOperationNId | String |
| UserId | String |
| HandlingUnit_Id | String |
| Action_Id | String |
| MaterialItemQuantity_UoMNId | String |
| MaterialItemQuantity_QuantityValue | Integer/Decimal |

#### Entity: HandlingUnitMaterialItemAssociation

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
| DM_MaterialTrackingUnit_Id | String |
| HandlingUnit_Id | String |

#### Entity: DMMtuWithResult

**Published From:** AppU4DM

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Status | String |
| Description | String |
| Quantity | String |
| UoM | String |
| Volume | String |
| VolumeUoM | String |
| Weight | String |
| WeightUoM | String |
| CodeType | String |
| Code | String |
| IsReserved | Boolean |
| DM_MaterialId | String |
| MaterialNId | String |
| MaterialRevision | String |
| _Id | String |
| IsResultKo | Boolean |

#### Entity: StatusTransition

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
| Verb | String |
| DoRaiseEvent | Boolean |
| IsHidden | Boolean |
| IsSystemDefined | Boolean |
| TargetStatus_Id | String |
| SourceStatus_Id | String |

#### Entity: TransportOperation

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
| RoleId | String |
| EstimatedDuration_Ticks | String |
| Status_StateMachineNId | String |
| Status_StatusNId | String |

#### Entity: TransportOperationDestination

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
| IsSelected | Boolean |
| Priority | String |
| TransportOperation_Id | String |
| Buffer_Id | String |
| LogisticRequestItem_Id | String |

#### Entity: TransportOperationSource

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
| IsSelected | Boolean |
| Priority | String |
| TransportOperation_Id | String |
| Buffer_Id | String |
| LogisticRequestItem_Id | String |

#### Entity: TransportOperationHistory

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
| UserId | String |
| Date | DateTime |
| DestinationBufferNId | String |
| SourceBufferNId | String |
| AssignedUserId | String |
| TransportOperation_Id | String |
| Action_Id | String |

#### Entity: Buffer_2

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
| Description | String |
| EquipmentNId | String |
| EquipmentName | String |
| Name | String |
| NId | String |
| ValidityFrom | String |
| ValidityTo | String |
| BufferDefinition_Id | String |
| CapacityType_Id | String |
| MaxQty_UoMNId | String |
| MaxQty_QuantityValue | Integer/Decimal |
| MaxVolume_UoMNId | String |
| MaxVolume_QuantityValue | Integer/Decimal |
| MaxWeight_UoMNId | String |
| MaxWeight_QuantityValue | Integer/Decimal |
| Status_StateMachineNId | String |
| Status_StatusNId | String |

#### Entity: LogisticRequest

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
| DueDate | DateTime |
| UserId | String |
| Name | String |
| NId | String |
| Priority | String |
| Destination_Id | String |
| Buffer_Id | String |
| Status_StateMachineNId | String |
| Status_StatusNId | String |

#### Entity: LogisticRequestItem

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
| ToolDefinition | String |
| _Type | String |
| DM_Material_Id | String |
| LogisticRequest_Id | String |
| Source_Id | String |
| TransportOperation_Id | String |
| Quantity_UoMNId | String |
| Quantity_QuantityValue | Integer/Decimal |

#### Entity: ProductConfiguration

**Published From:** ProductConfiguration

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
| FinalProductTypeNId | String |
| FinalProductFamilyNId | String |
| Name | String |
| Description | String |

#### Entity: ProductConfigurationStatus

**Published From:** ProductConfiguration

| Attribute | Type |
|-----------|------|
| _Id | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| EntityType | String |
| IsLocked | Boolean |
| ToBeCleaned | String |
| FinalProductTypeNId | String |
| FinalProductFamilyNId | String |
| ProductConfigurationRevision | String |
| Producibility | String |
| LastAnalysis | String |
| ProductConfiguration_Id | String |

#### Entity: UoM_2

**Published From:** ERPOrder

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
| IsHidden | Boolean |
| IsSystemDefined | Boolean |
| UoMDimension_Id | String |
| UoMBase_Id | String |

#### Entity: ERPOrderBoP_MasterPlan

**Published From:** ERPOrder

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
| CorrelationId | String |
| UId | String |
| Completed | String |
| ComplementaryEngineeringNeeded | String |
| SpecializedSpecificationSet | String |

#### Entity: OfflineAction

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
| Name | String |
| CommandName | String |
| CommandApp | String |
| CommandPayload | String |
| WorkOrderOperationNId | String |
| UserId | String |
| CheckInDate | DateTime |
| CheckInResponse | String |
| Date | DateTime |
| InternalId | String |
| OfflineSession_Id | String |
| CheckInStatus_StateMachineNId | String |
| CheckInStatus_StatusNId | String |

#### Entity: OfflineSessionWOOperationAssociation

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
| WorkOrderOperationNId | String |
| OfflineSession_Id | String |

#### Entity: FinalProductType_2

**Published From:** CEM_ERPOrder

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
| FinalProductTypeFacetNav_Id | String |

#### Entity: OfflineActionAmend

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
| UserId | String |
| CommandPayload | String |
| Date | DateTime |
| OfflineAction_Id | String |

#### Entity: FinalProductTypeFacet

**Published From:** CEM_ERPOrder

| Attribute | Type |
|-----------|------|
| _Id | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| EntityType | String |
| IsLocked | Boolean |
| ToBeCleaned | String |
| ExecutionType_MasterPlanNId | String |
| ExecutionType_Type | String |
| ExecutionType_IsWorkOrderCreationDirect | String |

#### Entity: CapacityType

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

#### Entity: BufferEquipmentAssoication

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
| EquipmentNId | String |
| Priority | String |
| Buffer_Id | String |

#### Entity: ProcessMeasurement

**Published From:** CEM_AppU4DM

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
| Val | String |
| DataType | String |
| MaterialTrackingUnitNId | String |
| EquipmentNId | String |
| WorkOrderOperationNId | String |
| ValUoMNId | String |
| AcquisitionDate | DateTime |
| MinVal | Integer/Decimal |
| MinValUoMNId | String |
| MaxValUoMNId | String |
| MaxVal | Integer/Decimal |
| TargetValUoMNId | String |
| TargetVal | String |
| Description | String |
| Notes | String |
| CorrelationId | String |
| GroupId | String |
| ProductionEvent | String |
| TimeStamp | DateTime |
| Category | String |
| Tool | String |
| ProcessMeasurementFacetNav_Id | String |

#### Entity: ToBeUsedMachinePJF

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
| PrintJobFileExternalId | String |
| EG_PhaseId | String |
| PrintJobFileName | String |
| PrintJobFileSource | String |
| IsSourcePJF | Boolean |
| PrintJobFile_Id | String |
| ToBeUsedMachine_Id | String |

#### Entity: ToBeUsedMachine

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
| Machine | String |
| EquipmentType | String |
| PartProgram | String |
| PrintJobFile | String |
| PlannedSetupStartTime | DateTime |
| PlannedSetupEndTime | DateTime |
| Preferred | String |
| WorkOrderOperation_Id | String |

#### Entity: PrintJobFile

**Published From:** PrintJobFile

| Attribute | Type |
|-----------|------|
| _Id | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| EntityType | String |
| IsLocked | Boolean |
| ToBeCleaned | String |
| BuildtrayLayout | String |
| EstimatedMaterialConsumptionQuantity | String |
| EstimatedMaterialConsumptionUom | String |
| ExternalId | String |
| FileName | String |
| MaxProducedHeightQuantity | Integer/Decimal |
| ProductSerializable | String |
| RunTimeOnly | DateTime |
| ProductSerialNumberManagement | String |
| Source | String |
| TestBars | String |
| TestBarsNumber | String |
| TestBarsSerializable | String |
| TestBarsSerialNumberManagement | String |
| ProductionPrinter | String |
| FilePath | String |
| CanBePreTransferred | String |
| EstimatedBuildProcessDuration_Ticks | String |
| EstimatedPrintDuration_Ticks | String |
| ParentPJF_Id | String |

#### Entity: Printer3D

**Published From:** PrintJobFile

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
| StatusNId | String |
| AMMachine | String |
| TransferPlugin | String |
| Description | String |
| IsPreTransferEnabled | Boolean |
| IsInvalid | String |
| Technology | String |
| NoRecycle | String |
| NoQtyDecreaseOnLoadedBatch | String |
| PrinterType_Id | String |

#### Entity: PrintJobFile3DPrinterAssociation

**Published From:** PrintJobFile

| Attribute | Type |
|-----------|------|
| _Id | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| EntityType | String |
| IsLocked | Boolean |
| ToBeCleaned | String |
| Printer_Id | String |
| PrintJobFile_Id | String |

#### Entity: PrintJobFileMaterialAssociation

**Published From:** PrintJobFile

| Attribute | Type |
|-----------|------|
| _Id | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| EntityType | String |
| IsLocked | Boolean |
| ToBeCleaned | String |
| MaterialNId | String |
| MaterialRevision | String |
| MaterialUId | String |
| Quantity | String |
| QuantityUoMNId | String |
| PrintJobFile_Id | String |

#### Entity: HandlingUnitHistoryAction

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

#### Entity: WIDefinitionToEGPhaseAssociation

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
| WorkInstructionAssociationType_Id | String |
| WorkInstructionDefinition_Id | String |
| EGPhase_Id | String |

#### Entity: AMPowderRequiredOn3DPrinter

**Published From:** PowderMgt

| Attribute | Type |
|-----------|------|
| _Id | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| EntityType | String |
| IsLocked | Boolean |
| ToBeCleaned | String |
| AMPowderNId | String |
| AMPowderRevision | String |
| AMPowderUId | String |
| AMPowderMinQuantity | Integer/Decimal |
| AMPowderUoMNId | String |
| IsCheckMandatory | Boolean |
| CheckAMPowderQuantity | String |
| CheckAMPowderBatch | String |
| Printer3DNId | String |
| AMPowderName | String |
| Printer3DBayNId | String |
| Printer3DBaySequence | Integer |
| IsSupportMaterial | Boolean |
| BoPCtx_Id | String |
| WorkOrderCtx_Id | String |
| ExecGroupCtx_Id | String |

#### Entity: AMPowder

**Published From:** PowderMgt

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
| UId | String |
| Name | String |
| Description | String |
| UoMNId | String |
| MinQuantity | Integer/Decimal |
| MaxRecycleCount | Integer |
| IsInvalid | String |

#### Entity: ExecGroupCtx

**Published From:** PowderMgt

| Attribute | Type |
|-----------|------|
| _Id | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| EntityType | String |
| IsLocked | Boolean |
| ToBeCleaned | String |
| ExecGroupNId | String |
| ExecGroupPhaseNId | String |

#### Entity: ToBeConsumedMaterialPrekitHistory

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
| ActionNId | String |
| DateHistory | DateTime |
| UserId | String |
| WorkOrder_Id | String |
| WorkOrderStep_Id | String |
| PrekitDM_MaterialTrackingUnit_Id | String |
| WorkOrderOperation_Id | String |
| TargetMaterialItem_Id | String |

#### Entity: ToBeConsumedMaterialPrekit

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
| PrekitSerialNumber | String |
| PreKitBatchCode | String |
| ToBeValidated | DateTime |
| UoMNId | String |
| ToBeConsumedMaterial_Id | String |
| DM_MaterialTrackingUnit_Id | String |
| PrekitMaterialItem_Id | String |

#### Entity: RM_PC_FAIRecord

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| FAIRecordNId | String |
| FinalMaterialNId | String |
| FinalMaterialRevision | String |
| FinalMaterialId | String |
| ToBeConsumedMaterialId | String |
| ToBeConsumedMaterialNId | String |
| ToBeConsumedMaterialRevision | String |
| IsFullFAI | Boolean |
| StatusNId | String |
| DM_MTUId | String |
| WoCandidateId | String |
| WoCandidateNId | String |
| WoCandidateName | DateTime |
| FAIRelatedId | String |
| FAIRelatedNId | String |
| FullFAIId | String |
| FullFAINId | String |
| WorOrdersRelevantNId | String |
| _Id | String |
| FAIRecordId | String |
| MTUCandidateNId | String |
| MTUCandidateCode | DateTime |
| FAIRecordEntityType | String |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |

#### Entity: RM_PC_FAIWorkOrderRelevant

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| FAIRecordNId | String |
| WorkOrderId | String |
| WorkOrderName | String |
| StatusNId | String |
| EstimatedStartTime | DateTime |
| EstimatedEndTime | DateTime |
| IsCandidate | Boolean |
| _Id | String |
| FAIRecordId | String |
| WorkOrderNId | String |
| CreatedOn | DateTime |
| ProductionType | String |

#### Entity: RM_PC_FAIReport_Form1

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| PartNumber | String |
| PartName | String |
| SerialNumber | String |
| FAIRNumber | String |
| PartRevisionLevel | String |
| DrawingNumber | String |
| DrawingRevisionLevel | String |
| AdditionalChanges | String |
| ManufacturingProcessReference | String |
| OrganizationName | String |
| SupplierCode | String |
| PONumber | String |
| IsDetailPart | Boolean |
| IsFullFAI | Boolean |
| IsPartialFAI | Boolean |
| BaselinePartNumber | String |
| ReasonForPartialFAI | String |
| Signature | String |
| IsFAICompleted | Boolean |
| SignatureDate | DateTime |
| ReviewedBy | String |
| ReviewedDate | DateTime |
| CustomerApproval | String |
| CustomerApprovalDate | DateTime |
| _Id | String |
| IsAssemblyFAI | Boolean |
| NonConformancesExist | String |
| FAIRecordId | String |
| IsFAINotCompleted | Boolean |

#### Entity: RM_PC_FAIReportPart_Form1

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| PartNumber | String |
| PartName | String |
| FAIRNumber | String |
| _Id | String |
| FAIRecordId | String |
| SerialNumber | String |

#### Entity: FAIInspection

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
| CharacteristicRepresentationNId | String |
| CharacteristicRepresentationRevision | String |
| InspectionSample | String |
| InspectionSampleAnyViolation | String |
| FAIRecord_Id | String |

#### Entity: RM_PC_FAIReport_Form2

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| PartNumber | String |
| PartName | String |
| SerialNumber | String |
| FAIRNumber | String |
| Signature | String |
| SignatureDate | DateTime |
| _Id | String |
| FAIRecordId | String |

#### Entity: RM_PC_FAIReport_Form3

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| PartNumber | String |
| PartName | String |
| SerialNumber | String |
| FAIRNumber | String |
| Signature | String |
| SignatureDate | DateTime |
| _Id | String |
| FAIRecordId | String |

#### Entity: RM_PC_FAIReportCharacteristic_Form3

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| CharNo | String |
| ReferenceLocation | String |
| CharacteristicDesignator | String |
| Requirement | String |
| Results | String |
| DesignedQualifiedTooling | String |
| NonConformanceNumber | String |
| _Id | String |
| FAIRecordId | String |

#### Entity: RM_PC_FAIRecordHistory

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| FAIRecordId | String |
| User | String |
| Date | DateTime |
| Action | String |
| _Id | String |

#### Entity: RM_PC_ContainmentRequest

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| Description | String |
| ExternalId | String |
| Name | String |
| NId | String |
| Notes | String |
| Sequence | Integer |
| StateMachineNId | String |
| StatusNId | String |
| System | String |
| CanBeReleased | String |
| CanBeClosed | String |
| CanBeSetAsReady | String |
| IsImportedFromExtSystem | Boolean |
| IsClosed | Boolean |
| _Id | String |
| ContainmentRequestId | String |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |

#### Entity: FAIInspectionHistory

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
| UserId | String |
| InspectionSample | String |
| NonConformanceId | String |
| Date | DateTime |
| Action_Id | String |
| FAIInspection_Id | String |

#### Entity: RM_PC_ContainmentRequestHistory

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| Action | String |
| ContainmentRequestNId | String |
| ContainmentRequestExternalId | String |
| Notes | String |
| System | String |
| UserId | String |
| OldStatus | String |
| Status | String |
| IsAcknowledged | Boolean |
| AcknowledgedBy | String |
| ItemType | String |
| MaterialTrackingUnitNId | String |
| MaterialTrackingUnitCodeType | String |
| MaterialNId | String |
| MaterialRevision | String |
| MaterialUId | String |
| Date | DateTime |
| MaterialTrackingUnitCode | String |
| ContainmentRequestId | String |
| ContainmentRequestHistoryItemId | String |
| ItemNId | String |
| ContainmentRequestReleaseExternalId | String |
| _Id | String |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| ContainmentRequestHistoryId | String |

#### Entity: RM_PC_FAIInspectionHistory

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| User | String |
| Date | DateTime |
| FAIInspectionId | String |
| ActionNId | String |
| NonConformanceId | String |
| NonConformanceNId | String |
| NonConformanceSerialNumber | String |
| NonConformanceSeverity | String |
| NonConformanceStatus | String |
| InspectionSampleId | String |
| SampleAnyViolation | String |
| IsSampleConfirmed | Boolean |
| IsSampleSignaturePending | Boolean |
| _Id | String |
| FAIInspectionHistoryId | String |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |

#### Entity: RM_PC_ContainmentRequestItem_MTU

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| ContainmentRequestId | String |
| MaterialTrackingUnitNId | String |
| MaterialTrackingUnitCodeType | String |
| MaterialNId | String |
| MaterialRevision | String |
| MaterialUId | String |
| MaterialTrackingUnitStatus | String |
| IsMaterialTrackingUnitInContainment | Boolean |
| MaterialTrackingUnitQuantity | String |
| MaterialTrackingUnitUoM | String |
| MaterialTrackingUnitCode | String |
| IsContainmentRequestItemReleased | Boolean |
| IsContainmentRequestItemScrapped | Boolean |
| MaterialTrackingUnitName | String |
| _Id | String |
| ContainmentRequestItemId | String |

#### Entity: RM_PC_ContainmentRequest_CloseStatus

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| StateMachineNId | String |
| SourceStatus | String |
| TargetStatus | String |
| _Id | String |

#### Entity: RM_PC_ContainmentRequest_ReadyStatus

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| StateMachineNId | String |
| SourceStatus | String |
| TargetStatus | String |
| _Id | String |

#### Entity: RM_PC_ContainmentRequest_ReleaseStatus

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| StateMachineNId | String |
| SourceStatus | String |
| TargetStatus | String |
| _Id | String |

#### Entity: RM_PC_NonConformanceScrapStatus

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| Lifecycle | String |
| Status | String |
| _Id | String |

#### Entity: RM_PC_ContainmentRequestByMaterialTrackingUnit

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| MaterialTrackingUnitNId | String |
| MaterialTrackingUnitName | String |
| MaterialTrackingUnitCode | String |
| MaterialNId | String |
| MaterialRevision | String |
| MaterialUId | String |
| _Id | String |
| ContainmentRequest_Id | String |

#### Entity: ContainmentRequestExt

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
| ExternalSystem | String |
| ContainmentRequestExternalId | String |
| EventDateTime | DateTime |
| ImportDateTime | DateTime |
| Sequence | Integer |
| ImportedSuccessfully | String |

#### Entity: WorkingSession

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
| User | String |
| ClientId | String |
| Team | String |
| SessionNId | String |
| MaterialTrackingUnitId | String |

#### Entity: RM_PC_ContainmentRequestMonitoring

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| ContainmentRequestExternalIdExt | String |
| ContainmentRequestReleaseExternalId | String |
| EventDateTime | DateTime |
| ImportDateTime | DateTime |
| ExternalSystem | String |
| Sequence | Integer |
| ContainmentRequestId | String |
| ContainmentRequestNId | String |
| ContainmentRequestName | String |
| ContainmentRequestSequence | Integer |
| ContainmentRequestExternalId | String |
| _Id | String |
| ImportedRequestId | String |
| ImportedSuccessfully | String |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |

#### Entity: RM_PC_ScanMaterialTrackingUnit

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| NId | String |
| Code | String |
| CodeType | String |
| Name | String |
| Description | String |
| Quantity | String |
| QuantityUom | String |
| Status | String |
| StateMachineNId | String |
| DMMaterialId | String |
| IsReserved | Boolean |
| VolumeQty | String |
| VolumeUoM | String |
| WeightQty | String |
| WeightUoM | String |
| MaterialNId | String |
| MaterialRevision | String |
| MaterialUId | String |
| CurrentLocation | String |
| EquipmentNId | String |
| SessionNId | String |
| User | String |
| Team | String |
| ClientId | String |
| _Id | String |
| MaterialTrackingUnitId | String |
| DMMaterialTrackingUnitId | String |
| ActiveNonConformanceNumber | String |
| InContainment | String |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |

#### Entity: ConfigurationKey_2

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
| _Type | String |
| Val | String |
| IsProtected | Boolean |
| NId | String |
| Category | String |
| IsMultiValue | Boolean |

#### Entity: ActualCoProducedMaterial

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
| MTUNId | String |
| MTUCode | String |
| DM_MTUId | String |
| DisassembleWorkOrderOperationNId | String |
| DisassembleWorkOrderStepNId | String |
| DisassembleMTUCode | String |
| DisassembleWorkOrderNId | String |
| DisassembleWorkOrderOperationName | String |
| DisassembleMTUName | String |
| DisassembleWorkOrderStepName | String |
| MaterialItemCoProducedQtyUoMNId | String |
| ProducedMaterialItemActualQtyUoMNId | String |
| QuantityUoMNId | String |
| MTUId | String |
| Quantity | String |
| MaterialItemCoProducedQty | String |
| ProducedMaterialItemActualQty | String |
| WorkOrderOperation_Id | String |
| WorkOrderStep_Id | String |
| ToBeCoProducedMaterial_Id | String |

#### Entity: RM_PC_ChangePackageWorkOrderOperation

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Sequence | Integer |
| Description | String |
| Priority | String |
| ToBeCollectedDocument | String |
| EstimatedDuration | String |
| EstimatedStartTime | DateTime |
| EstimatedEndTime | DateTime |
| ActualStartTime | DateTime |
| ElectronicSignatureStart | String |
| ElectronicSignaturePause | String |
| ElectronicSignatureComplete | String |
| OperationNId | String |
| OperationRevision | String |
| OperationUId | String |
| WorkOrderNId | String |
| ChangePackageNId | String |
| ChangePackageItemId | String |
| ActionNId | String |
| StatusNId | String |
| OperationStepCategoryNId | String |
| IsChangePackageItemWOOperation | Boolean |
| IsChangePackageOpenedOn | Boolean |
| _Id | String |
| WOOperationId | String |
| WorkOrderId | String |
| ChangePackageId | String |

#### Entity: RM_PC_ChangePackageWorkOOperationDependency

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| ChangePackageNId | String |
| ChangePackageItemId | String |
| ActionNId | String |
| FromOperationNId | String |
| FromOperationName | String |
| FromOperationSequence | Integer |
| ToOperationNId | String |
| ToOperationName | String |
| ToOperationSequence | Integer |
| DependencyType | String |
| WorkOOperationDependencyId | String |
| ChangePackageWorkOOperationDependencyId | String |
| _Id | String |
| ChangePackageId | String |

#### Entity: RM_PC_ChangePackageRoutingNode

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| ChangePackageNId | String |
| WorkOrderNId | String |
| WorkOrderName | String |
| WorkOrderStatusNId | String |
| EntityNId | String |
| EntityName | String |
| EntitySequence | Integer |
| ParentEntityId | String |
| ParentEntityType | String |
| EntityHasChildren | String |
| EntityIsReady | String |
| EntityStatusNId | String |
| EntityHasPendingNC | String |
| EntityType | String |
| IsExternal | Boolean |
| _Id | String |
| ChangePackageId | String |
| WorkOrderId | String |
| EntityId | String |

#### Entity: RM_PC_ChangePackageRoutingEdge

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| ChangePackageNId | String |
| CurrentWorkOrderId | String |
| FromWorkOrderId | String |
| FromNodeEntityId | String |
| FromNodeEntityType | String |
| FromNodeParentEntityId | String |
| ToWorkOrderId | String |
| ToNodeEntityId | String |
| ToNodeEntityType | String |
| ToNodeParentEntityId | String |
| DependencyType | String |
| AlternativeGroup | String |
| IsPreferred | Boolean |
| _Id | String |
| ChangePackageId | String |

#### Entity: RM_PC_PrekitSerialBatchMaterials

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| WorkOrderId | String |
| Code | String |
| UoMNId | String |
| StatusNId | String |
| StateMachineNId | String |
| DMMaterialId | String |
| MaterialName | String |
| MaterialSpecificationType | String |
| MaterialQuantity | String |
| MaterialSequence | Integer |
| SerialNumberProfile | String |
| MaterialRevision | String |
| MaterialNId | String |
| ToBeConsumedMaterialPrekitId | String |
| PrekitSerialNumber | String |
| PrekitBatchCode | String |
| ToBeValidated | DateTime |
| PrekittedQuantity | String |
| WorkOrderOperationId | String |
| WorkOrderOperationName | String |
| WorkOrderOperationSequence | Integer |
| WorkOrderOperationNId | String |
| WorkOrderOperationTargetQuantity | String |
| WorkOrderStepId | String |
| WorkOrderStepName | String |
| WorkOrderStepSequence | Integer |
| WorkOrderStepNId | String |
| WorkOrderStepTargetQuantity | String |
| DMMaterialTrackingUnitId | String |
| _Id | String |
| ProducedMaterialItemId | String |
| ToBeConsumedMaterialId | String |
| MTUToBeProducedMaximumQuantity | Integer/Decimal |

#### Entity: RM_PC_ContainmentRequestByEquipment

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| ContainmentRequestNId | String |
| EquipNId | String |
| EquipName | String |
| EquipLevel | String |
| EquipStatus | String |
| IsContainmentRequestItemReleased | Boolean |
| IsContainmentRequestItemScrapped | Boolean |
| _Id | String |

### Module: OpcenterEXDS_ProductionCoordination_Connector

#### Entity: UADMGetAllOfflineSessionsLists_RF

**Published From:** Consumed_ReadingFunctions

| Attribute | Type |
|-----------|------|
| UserId | String |
| DownloadDate | Boolean |
| UploadDate | DateTime |
| CheckInDate | DateTime |
| DeviceId | String |
| Status_StateMachineNId | String |
| Status_StatusNId | String |
| ActionCheckedIn | String |
| LastUpdatedOn | DateTime |
| _Id | String |
| CreatedOn | DateTime |

#### Entity: WorkOrderNotification

**Published From:** AppU4DM

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: GetExecutionGroupListWithEstimatedDuration_RF

**Published From:** Consumed_ReadingFunctions

| Attribute | Type |
|-----------|------|
| TotalQuantity | String |
| NId | String |
| Name | String |
| Description | String |
| EstimatedStartTime | DateTime |
| EstimatedEndTime | DateTime |
| EstimatedDuration | String |
| Status | String |
| _Id | String |
| DueDate | DateTime |

#### Entity: GetFinalMaterial_RF

**Published From:** Consumed_ReadingFunctions

| Attribute | Type |
|-----------|------|
| NId | String |
| Revision | String |
| UId | String |
| SerialNumberProfile | String |
| UoMNId | String |
| Name | String |
| Description | String |

#### Entity: GetERPOrderWithValidity_RF

**Published From:** Consumed_ReadingFunctions

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Description | String |
| FinalProductFamilyNId | String |
| FinalProductTypeNId | String |
| ProductionDate | DateTime |
| DeliveryDate | DateTime |
| CustomerTypeNId | String |
| _Id | String |
| IsProducible | Boolean |
| Sequence | Integer |
| SerialNumber | String |
| ProductConfigurationRevision | String |
| FrozenStructureContainer_Id | String |
| MaterialInstanceNId | String |
| BaselineUId | String |
| MasterPlanNId | String |
| StatusNId | String |
| StateMachineNId | String |
| FinalMaterialNId | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| ToBeCleaned | String |
| FinalMaterialUId | String |
| FinalMaterialRevision | String |
| QuantityValue | Integer/Decimal |
| QuantityUoM | String |
| IsWorkOrderCreationDirect | Boolean |
| ProductionType | String |

#### Entity: UADMGetWOByProcessRevisions

**Published From:** Consumed_ReadingFunctions

| Attribute | Type |
|-----------|------|
| _Id | String |
| Name | String |
| NId | String |
| Description | String |
| CreationDate | DateTime |
| Process | String |
| ProcessName | String |
| ProcessRevision | String |
| PBOPIdentID | String |
| Status | String |

#### Entity: GetBOMLinkedERP_RF

**Published From:** Consumed_ReadingFunctions

| Attribute | Type |
|-----------|------|
| MaterialNId | String |
| MaterialName | String |
| MaterialDescription | String |
| MaterialFunctionalCodeNId | String |
| MaterialFunctionalCodeName | String |
| UoMNId | String |
| QuantityValue | Integer/Decimal |
| MaterialUId | String |
| MaterialRevision | String |
| MaterialFunctionalCodeDescription | String |
| DM_MaterialId | String |
| MaterialItemId | String |
| _Id | String |
| FuncFilter_ERPOrderId | String |

#### Entity: GetBOFLinkedERP_RF

**Published From:** Consumed_ReadingFunctions

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Description | String |
| FeatureValue | Integer/Decimal |
| OptionValue | Integer/Decimal |
| FeatureValueName | Integer/Decimal |
| FeatureValueDescription | Integer/Decimal |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdateOn | DateTime |
| ToBeCleaned | String |
| FuncFilter_ERPOrderId | String |

#### Entity: GetSuppliersFromMaterial_RF

**Published From:** Consumed_ReadingFunctions

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Description | String |
| FuncFilter_MaterialItemId | String |
| FuncFilter_DM_MaterialId | String |

#### Entity: UADMGetAllMaterialForPrekit_RF

**Published From:** Consumed_ReadingFunctions

| Attribute | Type |
|-----------|------|
| _Id | String |
| CreatedOn | DateTime |
| AlternativeSelected | String |
| LogicalPosition | String |
| SerialNumberProfile | String |
| MaterialSpecificationType | String |
| MaterialDefinitionNId | String |
| NId | String |
| Name | String |
| PrekitSerialNumber | String |
| WorkOrderOperationNId | String |
| WorkOrderOperationName | String |
| WorkOrderOperationSequence | Integer |
| WorkOrderStepNId | String |
| WorkOrderStepName | String |
| Sequence | Integer |
| Quantity | String |
| WorkOrderOperationId | String |
| FuncFilter_woId | String |
| WorkOrderStepSequence | Integer |
| WorkOrderStepId | String |

#### Entity: UADMGetEventLogs_RF

**Published From:** Consumed_ReadingFunctions

| Attribute | Type |
|-----------|------|
| WorkOrder | String |
| WorkOrderOperation | String |
| CompletedQuantity | String |
| ERPOrder | String |
| Status | String |
| ErrorMessage | String |
| LastUpdate | DateTime |
| EventEntityType | String |
| EventID | String |
| PlantID | String |

#### Entity: UADMGetDNCItemsFromAssociations_RF

**Published From:** Consumed_ReadingFunctions

| Attribute | Type |
|-----------|------|
| DNCItemId | String |
| DNCId | String |
| DNCExternalId | String |
| FuncFilter_EquipmentNId | String |
| FuncFilter_DM_MaterialNId | String |

### Module: OpcenterEXDS_ShopfloorExecution

#### Entity: SetPointToolAssociation

**Published From:** AppU4DM

| Attribute | Type |
|-----------|------|
| _Id | String |
| SetPoint_Id | String |
| ToolNId | String |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| EntityType | String |
| ToBeCleaned | String |
| IsFrozen | Boolean |
| IsLocked | Boolean |
| IsDefault | Boolean |

#### Entity: ToolHistory

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
| Date | DateTime |
| Equipment | String |
| OldThickness | String |
| OldTreatmentCount | Integer |
| Thickness | String |
| TreatmentCount | Integer |
| UserId | String |
| UsageDuration_Ticks | String |
| UsedToolCount | Integer |
| Action_Id | String |
| WorkOrderOperation_Id | String |
| NonConformance_Id | String |
| Tool_Id | String |
| DM_MaterialTrackingUnit_Id | String |
| WorkOrderStep_Id | String |

#### Entity: ToolHistoryAction

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

#### Entity: AutomationNodeInstance

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
| Revision | String |
| SourceRevision | String |
| IsCurrent | Boolean |
| NId | String |
| Name | String |
| Description | String |
| ChannelNId | String |
| AutomationNodeType_Id | String |
| ActivationStatus_Dirty | String |
| ActivationStatus_Activated | String |
| ActivationStatus_Ready | String |
| ActivationStatus_Dismissed | String |
| ActivationStatus_ConfirmedOn | DateTime |

#### Entity: ToolAutomationNodeInstanceAssociation

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
| ToolNId | String |
| AutomationNodeInstanceNId | String |

#### Entity: LogisticClass

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
| Name | String |

#### Entity: SubstrateDefinition

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
| ToolDefinition_Id | String |
| AutomaticThicknessDecrease | String |
| MaxTreatmentCountHold | Integer |
| MaxTreatmentCountWarning | Integer |
| MinThicknessHold | Integer/Decimal |
| MinThicknessWarning | Integer/Decimal |
| ThicknessDecreaseAmount | Integer/Decimal |

#### Entity: EXDSContainerOrMTU

**Published From:** AppU4DM

| Attribute | Type |
|-----------|------|
| NId | String |
| Code | String |
| CodeType | String |
| MaterialNId | String |
| MaterialRevision | String |
| Status | String |
| Quantity | String |
| StateMachine | String |
| ContainerType | String |
| IsReusable | Boolean |
| IsContainer | Boolean |
| DM_MTUId | String |
| IsResultKo | Boolean |
| _Id | String |

#### Entity: EXDSGetReworkCode

**Published From:** AppU4DM

| Attribute | Type |
|-----------|------|
| DefectTypeId | String |
| DefectGroupId | String |
| DM_MaterialNId | String |
| FailureNId | String |
| FailureName | String |
| ProcessId | String |
| ProcessNId | String |
| ProcessRevision | String |
| UId | String |
| TypeDescription | String |
| GroupDescription | String |
| DefectGroupNId | String |
| _Id | String |
| CreatedOn | DateTime |

#### Entity: EXDSContainerOrMTUbyWorkOrderOperation

**Published From:** AppU4DM

| Attribute | Type |
|-----------|------|
| NId | String |
| Code | String |
| CodeType | String |
| DM_MTUId | String |
| MaterialNId | String |
| MaterialRevision | String |
| StateMachine | String |
| Status | String |
| WorkOrderOperationId | String |
| IsContainer | Boolean |
| Quantity | String |
| ContainerType | String |
| IsReusable | Boolean |
| _Id | String |
| IsResultKo | Boolean |

#### Entity: EXDSGetNCDocumentAssociation

**Published From:** AppU4DM

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Description | String |
| MIMEType | String |
| IconId | String |
| RepositoryType | String |
| LocalFileId | String |
| FileName | String |
| IconName | String |
| Revision | String |
| NonconformanceId | String |
| LinkedEntityId | String |
| Category | String |
| NonconformanceHistoryId | String |
| _Id | String |
| EntityLinkId | String |
| DocId | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |

#### Entity: Defect

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
| GroupPath | String |
| FailureNId | String |
| DefectType_Id | String |
| NonConformance_Id | String |

#### Entity: DefectType

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
| Description | String |
| Name | String |

#### Entity: RM_PC_ToolAutomationNodeInstance

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| ToolNId | String |
| ToolName | String |
| ToolDescription | String |
| AutomationNodeInstanceNId | String |
| AutomationNodeInstanceName | String |
| AutomationNodeInstanceDescription | String |
| AutomationNodeInstanceRevision | String |
| AutomationNodeTypeNId | String |
| Status | String |
| _Id | String |
| ToolId | String |
| ToolAutomationNodeInstanceAssId | String |

#### Entity: RMI_PotentialFailureDetails

**Published From:** WorkInstruction

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Revision | String |
| Description | String |
| CharacteristicSpecificationId | String |
| CharacteristicSpecificationNId | String |
| CharacteristicSpecificationRevision | String |
| FailureWitheDetailsId | String |
| DocumentNId | String |
| DocumentRevision | Boolean |
| IsCurrent | Boolean |
| NumberOfChildren | String |
| NumberOfParents | String |
| FailureParentNId | String |
| FailureParentRevision | String |
| _Id | String |
| FailureId | String |

#### Entity: RM_PC_GenealogyWorkOrderOperation

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Sequence | Integer |
| WorkOrderId | String |
| Status | String |
| DmMtuId | String |
| SNStatus | String |
| EstimatedDuration | String |
| ExecutionDuration | String |
| _Id | String |
| WorkOrderOperationId | String |
| ProducedMaterialItemId | String |

#### Entity: RM_PC_GenealogyWorkOrderStep

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Sequence | Integer |
| WorkOrderOperationId | String |
| Status | String |
| EstimatedDuration | String |
| ExecutionDuration | String |
| _Id | String |
| WorkOrderStepId | String |

#### Entity: RM_PC_GenealogyConsumedMaterialOperation

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| WorkOrderOperationId | String |
| ProducedMaterialItemId | String |
| ConsumedDmMtuId | String |
| _Type | String |
| Code | String |
| MaterialNId | String |
| UserId | String |
| MaterialItemAssembledQty | String |
| MaterialItemAssembledQtyUoMNId | String |
| Date | DateTime |
| MaxRecycleCount | Integer |
| Sequence | Integer |
| _Id | String |
| NId | String |
| WorkOrderId | String |
| SerialNumberProfile | String |
| OccurenceId | String |
| ConsumedMTUStatus | String |
| IsAMPowder | Boolean |
| ProducedDM_MaterialTrackingUnitId | String |
| ProducedMTUNId | String |
| ProducedMaterialItemActualQty | String |
| WorkOrderNId | String |
| DM_MaterialId | String |
| ProducedMTUCode | String |
| ProducedMaterialNId | String |
| MaterialId | String |
| HasATNDetails | Boolean |

#### Entity: RM_PC_GenealogyToolOperation

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| WorkOrderOperationId | String |
| ProducedDM_MaterialTrackingUnitId | String |
| ToolDefinitionNId | String |
| ToolDefinitionName | String |
| ToolClassNId | String |
| ToolNId | String |
| ToolName | String |
| UsageCounter | Integer |
| UsedTimes | DateTime |
| UsageDurationTicks | String |
| QualityGateNId | String |
| EquipmentNId | String |
| ActivityNId | String |
| ToolDefinitionRevision | String |
| _Id | String |
| HasATNDetails | Boolean |
| HasScrewingDetails | Boolean |
| ToBeUsedToolId | String |
| BoltsNumber | String |
| AngleValueMax | Integer/Decimal |
| AngleValueMin | Integer/Decimal |
| TorqueValueMax | Integer/Decimal |
| TorqueValueMin | Integer/Decimal |

#### Entity: RM_PC_GenealogyEquipmentOperation

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| MachineNId | String |
| MachineName | String |
| WorkOrderOperationId | String |
| _Id | String |
| DM_MaterialTrackingUnitId | String |
| HasATNDetails | Boolean |

#### Entity: RM_PC_GenealogyNonConformance

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| NId | String |
| WorkOrderOperationId | String |
| Status | String |
| Severity | String |
| _Context | String |
| _Type | String |
| NonConformanceLifecycle | String |
| Quantity | String |
| User | String |
| _Id | String |
| NonConformanceId | String |
| CreatedOn | DateTime |
| DM_MaterialTrackingUnitId | String |

#### Entity: RM_PC_GenealogyNonConformanceDefect

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| NonConformanceId | String |
| _Id | String |
| NId | String |
| Name | String |
| Description | String |
| IsFailure | Boolean |

#### Entity: RM_PC_GenealogyActualCoProducedMaterialOperation

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| WorkOrderOperationId | String |
| MTUNId | String |
| MTUCode | String |
| MaterialItemCoProducedQty | String |
| MaterialItemCoProducedQtyUoMNId | String |
| MaterialSpecificationTypeNId | String |
| UserId | String |
| MaxRecycleCount | Integer |
| _Id | String |
| Date | DateTime |

#### Entity: RM_PC_GenealogyQuality

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| WorkOrderOperationId | String |
| ChrReprNId | String |
| WorkOrderNId | String |
| WorkOrderOperationNId | String |
| MeasuredAttributeValue | Integer/Decimal |
| MeasuredVariableValue | Integer/Decimal |
| MaterialTrackingUnitNId | String |
| MaterialTrackingUnitId | String |
| DM_MaterialTrackingUnitId | String |
| Timestamp | DateTime |
| User | String |
| FailureNId | String |
| VisualFailureCount | Integer |
| VisualFailureNId | String |
| _Id | String |
| CharacteristicSpecificationType | String |
| IsBuyOff | Boolean |

#### Entity: RM_PC_GenealogyWorkOrderHistory

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| WorkOrderId | String |
| SourceWorkOrderNId | String |
| TargetWorkOrderNId | String |
| _Id | String |
| WorkOrderHistoryId | String |

#### Entity: RM_PC_GenealogyWorkOrder

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| WorkOrderNId | String |
| Status | String |
| InitialQuantity | String |
| ActualTargetQuantity | String |
| TargetQuantity | String |
| PlannedTargetQuantity | String |
| RealInitialQty | String |
| Sequence | Integer |
| FinalMaterialNId | String |
| ProductionTypeNId | String |
| ParentBatch | String |
| _Id | String |
| WorkOrderId | String |

#### Entity: RM_PC_GenealogyProducedMaterialTrackingUnit

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| WorkOrderId | String |
| MaterialNId | String |
| NId | String |
| Code | String |
| CodeType | String |
| EquipmentNId | String |
| MaxRecycleCount | Integer |
| Status | String |
| Quantity | String |
| SerialNumberProfile | String |
| DM_MaterialTrackingUnitId | String |
| _Id | String |
| ProducedMaterialItemId | String |

#### Entity: RM_PC_GenealogyToolStep

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| WorkOrderStepId | String |
| ProducedDM_MaterialTrackingUnitId | String |
| ToolDefinitionNId | String |
| ToolDefinitionName | String |
| ToolClassNId | String |
| ToolNId | String |
| ToolName | String |
| UsageCounter | Integer |
| UsedTimes | DateTime |
| UsageDurationTicks | String |
| QualityGateNId | String |
| EquipmentNId | String |
| ActivityNId | String |
| ToolDefinitionRevision | String |
| _Id | String |
| HasATNDetails | Boolean |
| HasScrewingDetails | Boolean |
| ToBeUsedToolId | String |
| BoltsNumber | String |
| AngleValueMax | Integer/Decimal |
| AngleValueMin | Integer/Decimal |
| TorqueValueMax | Integer/Decimal |
| TorqueValueMin | Integer/Decimal |

#### Entity: RM_PC_GenealogyWorkInstructionOperation

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| WorkOrderOperationId | String |
| WorkInstructionId | String |
| DM_MaterialTrackingUnitId | String |
| NId | String |
| Name | String |
| Description | String |
| Status | String |
| IsOk | Boolean |
| IsRepeated | Boolean |
| SectionNId | String |
| SectionSequence | Integer |
| StepNId | String |
| StepTitle | String |
| StepSequence | Integer |
| ItemNId | String |
| ItemValue | Integer/Decimal |
| LowLimit | String |
| HighLimit | String |
| Target | String |
| ItemLabel | String |
| WorkInstructionDefinitionNId | String |
| StepType | String |
| StepCompletedBy | String |
| LastValueUpdatedOn | DateTime |
| IsItemCompleted | Boolean |
| IsStepCompleted | Boolean |
| ItemCompletedBy | String |
| ItemUoM | String |
| _Id | String |
| WorkInstructionToWOOpId | String |
| StepId | String |
| AcquisitionStatus | String |

#### Entity: RM_PC_GenealogyWorkInstructionStep

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| WorkOrderStepId | String |
| WorkInstructionId | String |
| DM_MaterialTrackingUnitId | String |
| NId | String |
| Name | String |
| Description | String |
| Status | String |
| IsOk | Boolean |
| IsRepeated | Boolean |
| SectionNId | String |
| SectionSequence | Integer |
| StepNId | String |
| StepTitle | String |
| StepSequence | Integer |
| ItemNId | String |
| ItemValue | Integer/Decimal |
| LowLimit | String |
| HighLimit | String |
| Target | String |
| ItemLabel | String |
| StepType | String |
| StepCompletedBy | String |
| IsStepCompleted | Boolean |
| ItemCompletedBy | String |
| LastValueUpdatedOn | DateTime |
| IsItemCompleted | Boolean |
| ItemUoM | String |
| _Id | String |
| WorkInstructionToWOStepId | String |
| StepId | String |
| WorkInstructionDefinitionNId | String |
| AcquisitionStatus | String |

#### Entity: RM_PC_GenealogyDocumentWorkOrderOperation

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| NId | String |
| Revision | String |
| FileName | String |
| CategoryNId | String |
| WorkOrderOperationId | String |
| MIMEType | String |
| DocumentName | Boolean |
| DocumentDescription | Boolean |
| _Id | String |
| FileId | String |
| ActualCollectedDocumentId | String |

#### Entity: RM_PC_GenealogyDocumentWorkOrderStep

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| NId | String |
| Revision | String |
| FileName | String |
| CategoryNId | String |
| WorkOrderStepId | String |
| MIMEType | String |
| DocumentName | Boolean |
| DocumentDescription | Boolean |
| _Id | String |
| FileId | String |
| ActualCollectedDocumentId | String |

#### Entity: RM_PC_GenealogyDocumentDM_MaterialTrackingUnit

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| NId | String |
| Revision | String |
| FileName | String |
| CategoryNId | String |
| DM_MaterialTrackingUnitId | String |
| MIMEType | String |
| DocumentName | Boolean |
| DocumentDescription | Boolean |
| _Id | String |
| FileId | String |
| DocumentDM_MTUAssociationId | String |

#### Entity: RM_PC_GenealogyQualityStep

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| WorkOrderStepId | String |
| ChrReprNId | String |
| WorkOrderNId | String |
| WorkOrderOperationNId | String |
| WorkOrderStepNId | String |
| MeasuredAttributeValue | Integer/Decimal |
| MeasuredVariableValue | Integer/Decimal |
| MaterialTrackingUnitNId | String |
| MaterialTrackingUnitId | String |
| DM_MaterialTrackingUnitId | String |
| Timestamp | DateTime |
| User | String |
| FailureNId | String |
| VisualFailureNId | String |
| VisualFailureCount | Integer |
| _Id | String |
| CharacteristicSpecificationType | String |
| IsBuyOff | Boolean |

#### Entity: RM_PC_GenealogyConsumedMaterialStep

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| WorkOrderStepId | String |
| ProducedMaterialItemId | String |
| ConsumedDmMtuId | String |
| _Type | String |
| Code | String |
| MaterialNId | String |
| UserId | String |
| MaterialItemAssembledQty | String |
| MaterialItemAssembledQtyUoMNId | String |
| Date | DateTime |
| MaxRecycleCount | Integer |
| Sequence | Integer |
| NId | String |
| OccurenceId | String |
| SerialNumberProfile | String |
| ConsumedMTUStatus | String |
| IsAMPowder | Boolean |
| _Id | String |
| ProducedDM_MaterialTrackingUnitId | String |
| ProducedMTUNId | String |
| ProducedMaterialItemActualQty | String |
| WorkOrderNId | String |
| DM_MaterialId | String |
| ProducedMTUCode | String |
| ProducedMaterialNId | String |
| MaterialId | String |
| HasATNDetails | Boolean |

#### Entity: RM_PC_GenealogySignersFromBuyOffHistory

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| MaterialTrackingUnitCode | String |
| WorkOrderOperationId | String |
| WorkOrderStepId | String |
| StatusTransitionNId | String |
| CharacteristicRepresentationNId | String |
| BuyOffId | String |
| Comment | String |
| Name | String |
| Signer | String |
| _Id | String |
| StatusTransitionId | String |
| TimeStamp | DateTime |
| DM_MaterialTrackingUnitId | String |

#### Entity: RM_PC_GenealogyWorkOrderHistoryMaterialItem

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| DM_MaterialTrackingUnitId | String |
| WorkOrderOperationId | String |
| WorkOrderStepId | String |
| Quantity | String |
| UoM | String |
| Date | DateTime |
| Equipment | String |
| ActionNId | String |
| ActionName | String |
| UserId | String |
| TeamNId | String |
| _Id | String |
| Message | String |

#### Entity: RM_PC_GenealogyActualCoProducedMaterialStep

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| WorkOrderStepId | String |
| MTUNId | String |
| MTUCode | String |
| MaterialItemCoProducedQty | String |
| MaterialItemCoProducedQtyUoMNId | String |
| MaterialSpecificationTypeNId | String |
| UserId | String |
| MaxRecycleCount | Integer |
| _Id | String |
| Date | DateTime |

#### Entity: RM_PC_GenealogyEquipmentDetails

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| WorkOrderOperationId | String |
| Equipment | String |
| ATNNodeId | String |
| ATNParameterName | String |
| SetPointItemNId | String |
| SetPointItemValue | Integer/Decimal |
| TransmissionSuccees | String |
| ATNParameterTimestamp | DateTime |
| ATNParameterQuality | String |
| SetPointNId | String |
| _Id | String |
| WorkOrderHistoryId | String |

#### Entity: RM_PC_GenealogyToolDetailsStep

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| WorkOrderStepId | String |
| ATNNodeId | String |
| ATNParameterName | String |
| SetPointItemNId | String |
| SetPointItemValue | Integer/Decimal |
| TransmissionSuccees | String |
| ATNParameterTimestamp | DateTime |
| ATNParameterQuality | String |
| SetPointNId | String |
| DM_MaterialTrackingUnitId | String |
| ToolId | String |
| DSParameterNId | String |
| AcquireFromShopfloorDate | DateTime |
| ATNDataType | String |
| ATNParameterValue | Integer/Decimal |
| _Id | String |

#### Entity: RM_PC_GenealogyToolDetailsOperation

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| WorkOrderOperationId | String |
| ATNNodeId | String |
| ATNParameterName | String |
| SetPointItemNId | String |
| SetPointItemValue | Integer/Decimal |
| TransmissionSuccees | String |
| ATNParameterTimestamp | DateTime |
| ATNParameterQuality | String |
| SetPointNId | String |
| DM_MaterialTrackingUnitId | String |
| ToolId | String |
| DSParameterNId | String |
| AcquireFromShopfloorDate | DateTime |
| ATNDataType | String |
| ATNParameterValue | Integer/Decimal |
| _Id | String |

#### Entity: RM_PC_GenealogyScrewingToolHistory

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| BoltNumber | String |
| Angle | String |
| Torque | String |
| IsOk | Boolean |
| ToBeUsedToolId | String |
| DM_MaterialTrackingUnitId | String |
| _Id | String |
| HistoryId | String |

#### Entity: RM_PC_GenealogyImport

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| CategoryNId | String |
| DocumentNId | String |
| DocumentName | Boolean |
| DocumentDescription | Boolean |
| DocumentMIMEType | Boolean |
| DocumentFileName | Boolean |
| _Id | String |
| DocumentLocalFileId | String |

#### Entity: RM_PC_GenealogyConsumedMaterialDetailsOperation

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| WorkOrderOperationId | String |
| DM_MaterialTrackingUnitId | String |
| ATNNodeId | String |
| ATNParameterName | String |
| ATNParameterTimestamp | DateTime |
| ATNParameterQuality | String |
| DSParameterNId | String |
| AcquireFromShopfloorDate | DateTime |
| ATNParameterValue | Integer/Decimal |
| _Id | String |

#### Entity: RM_PC_GenealogyConsumedMaterialDetailsStep

**Published From:** RMEXDSProdCoord

| Attribute | Type |
|-----------|------|
| WorkOrderStepId | String |
| DM_MaterialTrackingUnitId | String |
| ATNNodeId | String |
| ATNParameterName | String |
| ATNParameterTimestamp | DateTime |
| ATNParameterQuality | String |
| DSParameterNId | String |
| AcquireFromShopfloorDate | DateTime |
| ATNParameterValue | Integer/Decimal |
| _Id | String |

### Module: OpcenterEXDS_ShopfloorExecution_Connector

#### Entity: GetDocumentsLinkedToEntities_RF

**Published From:** Consumed_ReadingFunctions

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Description | String |
| MIMEType | String |
| IconId | String |
| RepositoryType | String |
| LocalFileId | String |
| FileName | String |
| IconName | String |
| EntityLinkId | String |
| _Id | String |
| Category | String |
| FuncFilter_EntityId | String |

#### Entity: UADMGetAllDocumentsFromDM_MTUIds_RF

**Published From:** Consumed_ReadingFunctions

| Attribute | Type |
|-----------|------|
| MIMEType | String |
| File_Id_Id | String |
| _Id | String |
| AId | String |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| FileName | String |
| NId | String |
| CorrelationId | String |
| DatasetNId | String |
| _Type | String |
| UId | String |
| Category | String |
| Description | String |
| MaterialTrackingUnitDocumentAssociationId | String |
| FuncFilter_DM_MTUIds | String |

---

## 2. Microflow/Action Calls

Microflow and action calls found in the project (MicroflowCall, JavaAction, ExternalAction).

Found 732 call(s):

| Microflow | Module | Call Type | AppName | CommandName |
|-----------|--------|-----------|---------|-------------|
| DSMaterial_UADMCreateDM_MTU | OpcenterEXDS_MasterData_Connector | MicroflowCall | 'Material'
 | 'DSMaterial_UADMCreateDM_MTU' |
| DSMaterial_UADMGenerateAndAssociateMTUCode | OpcenterEXDS_MasterData_Connector | MicroflowCall | 'AppU4DM' | 'DSMaterial_UADMGenerateAndAssociateMTUCode'
 |
| DSResult_AmendResultList | OpcenterEXDS_MasterData_Connector | MicroflowCall | 'Material'
 | 'DSResult_AmendResultList'
 |
| DSMaterial_UnlinkDocumentListToDM_MaterialTrackingUnit | OpcenterEXDS_MasterData_Connector | MicroflowCall | 'Material'
 | 'DSMaterial_UnlinkDocumentListToDM_MaterialTrackingUnit'
 |
| DSMaterial_LinkDocumentListToDM_MaterialTrackingUnit | OpcenterEXDS_MasterData_Connector | MicroflowCall | 'Material'
 | 'DSMaterial_LinkDocumentListToDM_MaterialTrackingUnit'
 |
| DSMaterial_MaterialItemSplitBatchByPartialPreview | OpcenterEXDS_MasterData_Connector | MicroflowCall | 'Material'
 | 'DSMaterial_MaterialItemSplitBatchByPartialPreview'
 |
| DSMaterial_UpdateMaterialTrackingUnitAndDM_MaterialTrackingUnit | OpcenterEXDS_MasterData_Connector | MicroflowCall | 'Material'
 | 'DSMaterial_UpdateMaterialTrackingUnitAndDM_MaterialTrackingUnit'
 |
| DSMaterial_CreateDM_MaterialTrackingUnitandMaterialTrackingUnit | OpcenterEXDS_MasterData_Connector | MicroflowCall | 'Material'
 | 'DSMaterial_CreateDM_MaterialTrackingUnitandMaterialTrackingUnit'
 |
| DSMaterial_MaterialItemSplitBatch | OpcenterEXDS_MasterData_Connector | MicroflowCall | 'Material'
 | 'DSMaterial_MaterialItemSplitBatch'
 |
| UADMRemoveHoldList | OpcenterEXDS_Configuration_Connector | JavaAction | 'AppU4DM' | 'UADMRemoveHoldList' |
| UADMSetLocationHold | OpcenterEXDS_Configuration_Connector | JavaAction | 'AppU4DM' | 'UADMSetLocationHold' |
| UADMSetWorkOrderHoldList | OpcenterEXDS_Configuration_Connector | JavaAction | 'AppU4DM' | 'UADMSetWorkOrderHoldList' |
| SUB_DeleteCommand_StartArray | OpcenterEXDS_Configuration_Connector | MicroflowCall | <parameter> | <parameter> |
| SUB_DeleteCommand | OpcenterEXDS_Configuration_Connector | MicroflowCall | <parameter> | <parameter> |
| UpdateLabelTemplate | OpcenterEXFN_MasterData_Connector | JavaAction | 'Label' | 'UpdateLabelTemplate' |
| CreateLabelTemplate | OpcenterEXFN_MasterData_Connector | JavaAction | 'Label' | 'CreateLabelTemplate' |
| CreateLabelPrinter | OpcenterEXFN_MasterData_Connector | JavaAction | 'Label' | 'CreateLabelPrinter' |
| CopyLabelPrinter | OpcenterEXFN_MasterData_Connector | JavaAction | 'Label' | 'CopyLabelPrinter' |
| UpdateLabelPrinter | OpcenterEXFN_MasterData_Connector | JavaAction | 'Label' | 'UpdateLabelPrinter' |
| ReprintLabel | OpcenterEXFN_MasterData_Connector | JavaAction | 'Label' | 'ReprintLabelsAndAddPrintHistory' |
| CreateLabelTag | OpcenterEXFN_MasterData_Connector | JavaAction | 'Label' | 'CreateLabelTag' |
| UpdateLabelTag | OpcenterEXFN_MasterData_Connector | JavaAction | 'Label' | 'UpdateLabelTag' |
| UpdateLabelType | OpcenterEXFN_MasterData_Connector | JavaAction | 'Label' | 'UpdateLabelType' |
| CopyLabelType | OpcenterEXFN_MasterData_Connector | JavaAction | 'Label' | 'CopyLabelType' |
| CreateLabelType | OpcenterEXFN_MasterData_Connector | JavaAction | 'Label' | 'CreateLabelType' |
| CreateSignatureConfigurations | OpcenterEXFN_MasterData_Connector | JavaAction | 'AuditTrail' | 'CreateSignatureConfigurations' |
| CreateScenarioConfiguration | OpcenterEXFN_MasterData_Connector | JavaAction | 'AuditTrail' | 'CreateScenarioConfiguration' |
| UpdateSignatureConfigurations | OpcenterEXFN_MasterData_Connector | JavaAction | 'AuditTrail' | 'UpdateSignatureConfigurations' |
| UpdateScenarioConfiguration | OpcenterEXFN_MasterData_Connector | JavaAction | 'AuditTrail' | 'UpdateScenarioConfiguration' |
| CreateBarcodeRule | OpcenterEXFN_MasterData_Connector | JavaAction | 'Barcode' | 'CreateBarcodeRule' |
| UpdateBarcodeRule | OpcenterEXFN_MasterData_Connector | JavaAction | 'Barcode' | 'UpdateBarcodeRule' |
| MoveBarcodeRulePart | OpcenterEXFN_MasterData_Connector | JavaAction | 'Barcode' | 'MoveBarcodeRulePart' |
| CreateBarcodeRulePart | OpcenterEXFN_MasterData_Connector | JavaAction | 'Barcode' | 'CreateBarcodeRulePart' |
| UpdateBarcodeRulePart | OpcenterEXFN_MasterData_Connector | JavaAction | 'Barcode' | 'UpdateBarcodeRulePart' |
| ActivateAutomationChannel | OpcenterEXFN_MasterData_Connector | ExternalAction | Automation | ActivateAutomationChannel |
| ApproveAutomationChannel | OpcenterEXFN_MasterData_Connector | ExternalAction | Automation | ApproveAutomationChannel |
| CreateOpcUaAutomationChannel | OpcenterEXFN_MasterData_Connector | JavaAction | 'Automation' | 'CreateOpcUaAutomationChannel' |
| SecureCheckAutomationChannelConnection | OpcenterEXFN_MasterData_Connector | JavaAction | 'Automation' | 'SecureCheckAutomationChannelConnection' |
| UpdateOpcUaAutomationChannel | OpcenterEXFN_MasterData_Connector | JavaAction | 'Automation' | 'UpdateOpcUaAutomationChannel' |
| CreateAutomationNodeInstance | OpcenterEXFN_MasterData_Connector | ExternalAction | Automation | CreateAutomationNodeInstance |
| UpdateAutomationNodeInstanceParameter | OpcenterEXFN_MasterData_Connector | JavaAction | 'Automation' | 'UpdateAutomationNodeInstanceParameter' |
| UpdateAutomationNodeInstanceParameter | OpcenterEXFN_MasterData_Connector | JavaAction | 'Automation' | 'UpdateAutomationNodeInstanceParameter' |
| DeleteAutomationNodeInstance | OpcenterEXFN_MasterData_Connector | JavaAction | 'Automation' | 'DeleteAutomationNodeInstance' |
| UnfreezeAutomationNodeInstance | OpcenterEXFN_MasterData_Connector | JavaAction | 'Automation' | 'UnfreezeAutomationNodeInstance' |
| RevertAutomationNodeInstance | OpcenterEXFN_MasterData_Connector | ExternalAction | Automation | RevertAutomationNodeInstance |
| CopyAutomationNodeInstance | OpcenterEXFN_MasterData_Connector | ExternalAction | Automation | CopyAutomationNodeInstance |
| ApproveAllAutomationNodeInstances | OpcenterEXFN_MasterData_Connector | ExternalAction | Automation | ApproveAllAutomationNodeInstances |
| ApproveAutomationNodeInstance | OpcenterEXFN_MasterData_Connector | ExternalAction | Automation | ApproveAutomationNodeInstance |
| UpdateAutomationNodeInstance | OpcenterEXFN_MasterData_Connector | JavaAction | 'Automation' | 'UpdateAutomationNodeInstance' |
| UpdateAutomationNodeInstance | OpcenterEXFN_MasterData_Connector | JavaAction | 'Automation' | 'UpdateAutomationNodeInstance' |
| FreezeAutomationNodeInstance | OpcenterEXFN_MasterData_Connector | JavaAction | 'Automation' | 'FreezeAutomationNodeInstance' |
| WriteAutomationNodeParameters | OpcenterEXFN_MasterData_Connector | JavaAction | 'Automation' | 'WriteAutomationNodeParameters' |
| ExportAutomationConfigBeforeMigrate | OpcenterEXFN_MasterData_Connector | JavaAction | 'Automation' | 'ExportAutomationConfigBeforeMigrate' |
| ImportAutomationConfig | OpcenterEXFN_MasterData_Connector | JavaAction | 'Automation' | 'ImportAutomationConfig' |
| IsMigrationNecessary | OpcenterEXFN_MasterData_Connector | ExternalAction | Automation | IsMigrationNecessary |
| ExportAutomationConfig | OpcenterEXFN_MasterData_Connector | JavaAction | 'Automation' | 'ExportAutomationConfig' |
| ResetEnvironment | OpcenterEXFN_MasterData_Connector | ExternalAction | Automation | ResetEnvironment |
| MigrateAutomation | OpcenterEXFN_MasterData_Connector | ExternalAction | Automation | MigrateAutomation |
| RevertAutomationNodeType | OpcenterEXFN_MasterData_Connector | ExternalAction | Automation | RevertAutomationNodeType |
| UpdateAutomationNodeType | OpcenterEXFN_MasterData_Connector | JavaAction | 'Automation' | 'UpdateAutomationNodeType' |
| DeleteAutomationNodeType | OpcenterEXFN_MasterData_Connector | JavaAction | 'Automation' | 'DeleteAutomationNodeType' |
| CreateAutomationNodeType | OpcenterEXFN_MasterData_Connector | JavaAction | 'Automation' | 'CreateAutomationNodeType' |
| Activate | OpcenterEXFN_MasterData_Connector | ExternalAction | Automation | Activate |
| CreateAutomationNodeTypeParameter | OpcenterEXFN_MasterData_Connector | JavaAction | 'Automation' | 'UpdateAutomationNodeType' |
| DeleteAutomationNodeTypeParameter | OpcenterEXFN_MasterData_Connector | JavaAction | 'Automation' | 'DeleteAutomationNodeTypeParameter' |
| UpdateAutomationNodeTypeParameter | OpcenterEXFN_MasterData_Connector | JavaAction | 'Automation' | 'UpdateAutomationNodeTypeParameter' |
| UnfreezeAutomationNodeType | OpcenterEXFN_MasterData_Connector | JavaAction | 'Automation' | 'UnfreezeAutomationNodeType' |
| DeleteAutomationNodeTypeAndSingleInstances | OpcenterEXFN_MasterData_Connector | JavaAction | 'Automation' | 'DeleteAutomationNodeTypeAndSingleInstances' |
| UpdateAutomationNodeTypeAndSingleInstances | OpcenterEXFN_MasterData_Connector | JavaAction | 'Automation' | 'UpdateAutomationNodeTypeAndSingleInstances' |
| ApproveAutomationNodeType | OpcenterEXFN_MasterData_Connector | JavaAction | 'Automation' | 'ApproveAutomationNodeType' |
| ApproveAllAutomationNodeTypes | OpcenterEXFN_MasterData_Connector | ExternalAction | Automation | ApproveAllAutomationNodeTypes |
| FreezeAutomationNodeType | OpcenterEXFN_MasterData_Connector | JavaAction | 'Automation' | 'FreezeAutomationNodeType' |
| DeletePerson | OpcenterEXFN_MasterData_Connector | JavaAction | 'Personnel' | 'DeletePerson' |
| UpdatePerson | OpcenterEXFN_MasterData_Connector | JavaAction | 'Personnel' | 'UpdatePerson' |
| ImportUsersByGroup | OpcenterEXFN_MasterData_Connector | JavaAction | 'Personnel' | 'ImportUsersByGroup' |
| CreatePerson | OpcenterEXFN_MasterData_Connector | JavaAction | 'Personnel' | 'CreatePerson' |
| ManageSegregationTagsToPerson | OpcenterEXFN_MasterData_Connector | JavaAction | 'Personnel' | 'ManageSegregationTagsToPersonsAssociation' |
| ImportUsersByUserList | OpcenterEXFN_MasterData_Connector | JavaAction | 'Personnel' | 'ImportUsersByUserList' |
| ManageSegregationTagsToPersonGroupsAssociation | OpcenterEXFN_MasterData_Connector | JavaAction | 'Personnel' | 'ManageSegregationTagsToPersonGroupsAssociation' |
| DisassociatePersonsFromPersonGroup | OpcenterEXFN_MasterData_Connector | JavaAction | 'Personnel' | 'DisassociatePersonsFromPersonGroup' |
| CreatePersonGroup | OpcenterEXFN_MasterData_Connector | JavaAction | 'Personnel' | 'CreatePersonGroup' |
| AssociatePersonsWithPersonGroup | OpcenterEXFN_MasterData_Connector | JavaAction | 'Personnel' | 'AssociatePersonsWithPersonGroup' |
| UpdatePersonGroup | OpcenterEXFN_MasterData_Connector | JavaAction | 'Personnel' | 'UpdatePersonGroup' |
| CompleteInspectionOrder | OpcenterEXFN_MasterData_Connector | MicroflowCall | 'WorkInstruction' | 'CompleteInspectionOrder' |
| CreateInspectionOrderFromInspectionOperation | OpcenterEXFN_MasterData_Connector | MicroflowCall | 'WorkInstruction' | 'CreateInspectionOrderFromInspectionOperation'
 |
| ReleaseInspectionOrder | OpcenterEXFN_MasterData_Connector | MicroflowCall | 'WorkInstruction' | 'ReleaseInspectionOrder' |
| DeleteInspectionOrder | OpcenterEXFN_MasterData_Connector | MicroflowCall | 'WorkInstruction' | 'DeleteInspectionOrder' |
| CreateInspectionOrderFromBoQ | OpcenterEXFN_MasterData_Connector | MicroflowCall | 'WorkInstruction' | 'CreateInspectionOrderFromBoQ'
 |
| UpdateMaterialGroup | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'UpdateMaterialGroup' |
| CreateMaterialGroup | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'CreateMaterialGroup' |
| AssociateChildrenMaterialGroupFromMaterialGroup | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'AssociateChildrenMaterialGroupWithMaterialGroup' |
| AssociateMaterialWithMaterialGroup | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'AssociateMaterialWithMaterialGroup' |
| AssociateParentsMaterialGroupWithMaterialGroup | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'AssociateParentsMaterialGroupWithMaterialGroup' |
| UpdateMaterialProperty | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'UpdateMaterialProperties' |
| AssociateMaterialGroupsWithMaterial | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'AssociateMaterialGroupsWithMaterial' |
| CreateBaseMaterial | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'CreateMaterial' |
| CreateMaterialProperty | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'CreateMaterialProperties' |
| CreateNewMaterialRevision | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'CreateNewMaterialRevision' |
| CopyMaterialRevision | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'CopyMaterialRevision' |
| UpdateMaterial | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'UpdateMaterial' |
| CreateMaterialTemplate | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'CreateMaterialTemplate' |
| CreateMaterialTemplateProperties | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'CreateMaterialTemplateProperties' |
| UpdateMaterialTemplate | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'UpdateMaterialTemplate' |
| UpdateMaterialTemplateProperties | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'UpdateMaterialTemplateProperties' |
| SetQuantityMaterialLot | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'SetMaterialLotQuantity' |
| CreateMaterialLotProperty | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'CreateMaterialLotProperties' |
| SetMaterialMaterialLot | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'SetMaterialLotMaterial' |
| SetStateMachineMaterialLot | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'SetMaterialLotStateMachine' |
| UpdateMaterialLot | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'UpdateMaterialLot' |
| UpdateMaterialLotProperty | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'UpdateMaterialLotProperties' |
| SetMaterialLotStatus | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'SetMaterialLotStatus' |
| CreateMaterialLot | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'CreateMaterialLot' |
| CreateMaterialTrackingUnitAggregatesProperties | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'CreateMaterialTrackingUnitAggregateProperties' |
| SetMaterialTrackingUnitAggregateStatus | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'SetMaterialTrackingUnitAggregateStatus' |
| LoadMaterialTrackingUnitAggregateToEquipment | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'LoadMaterialTrackingUnitAggregateToEquipment' |
| SetMaterialTrackingUnitAggregateStateMachine | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'SetMaterialTrackingUnitAggregateStateMachine' |
| SetMaterialTrackingUnitAggregateQuantity | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'SetMaterialTrackingUnitAggregateQuantity' |
| UpdateMaterialTrackingUnitAggregate | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'UpdateMaterialTrackingUnitAggregate' |
| CreateMaterialTrackingUnitAggregate | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'CreateMaterialTrackingUnitAggregate' |
| UpdateMaterialTrackingUnitAggregateProperties | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'UpdateMaterialTrackingUnitAggregateProperties' |
| SetMaterialTrackingUnitAggregateMaterial | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'SetMaterialTrackingUnitAggregateMaterial' |
| LoadMaterialTrackingUnitAggregateToMaterialTrackingUnitAggregate | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'LoadMaterialTrackingUnitAggregateToMaterialTrackingUnitAggregate' |
| CreateMaterialTrackingUnitAggregateTemplateProperties | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'CreateMaterialTrackingUnitAggregateTemplateProperties' |
| UpdateMaterialTrackingUnitAggregateTemplate | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'UpdateMaterialTrackingUnitAggregateTemplate' |
| UpdateMaterialTrackingUnitAggregateTemplateProperties | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'UpdateMaterialTrackingUnitAggregateTemplateProperties' |
| CreateMaterialTrackingUnitAggregateTemplate | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'CreateMaterialTrackingUnitAggregateTemplate' |
| CreateMaterialTrackingUnit | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'CreateMaterialTrackingUnit' |
| UpdateMaterialTrackingUnitProperties | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'UpdateMaterialTrackingUnitProperties' |
| AssociateMaterialTrackingUnitsWithMaterialLot | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'AssociateMaterialTrackingUnitsWithMaterialLot' |
| SetMaterialTrackingUnitStatus | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'SetMaterialTrackingUnitStatus' |
| SetStateMachineMaterialTrackingUnit | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'SetMaterialTrackingUnitStateMachine' |
| UpdateMaterialTrackingUnit | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'UpdateMaterialTrackingUnit' |
| MoveMaterialTrackingUnitToMaterialTrackingUnitAggregate | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'MoveMaterialTrackingUnitToMaterialTrackingUnitAggregate' |
| CreateMaterialTrackingUnitProperties | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'CreateMaterialTrackingUnitProperties' |
| SetMaterialTrackingUnitMaterial | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'SetMaterialTrackingUnitMaterial' |
| SetMaterialTrackingUnitQuantity | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'SetMaterialTrackingUnitQuantity' |
| MoveMaterialTrackingUnitToEquipment | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'MoveMaterialTrackingUnitToEquipment' |
| CreateMaterialTrackingUnitTemplate | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'CreateMaterialTrackingUnitTemplate' |
| CreateMaterialTrackingUnitTemplatesProperties | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'CreateMaterialTrackingUnitTemplateProperties' |
| UpdateMaterialTrackingUnitTemplate | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'UpdateMaterialTrackingUnitTemplate' |
| UpdateMaterialTrackingUnitTemplateProperties | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'UpdateMaterialTrackingUnitTemplateProperties' |
| UpdateMaterialLotTemplate | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'UpdateMaterialLotTemplate' |
| CreateMaterialLotTemplate | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'CreateMaterialLotTemplate' |
| CreateMaterialLotTemplateProperties | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'CreateMaterialLotTemplateProperties' |
| UpdateMaterialLotTemplateProperties | OpcenterEXFN_MasterData_Connector | JavaAction | 'Material' | 'UpdateMaterialLotTemplateProperties' |
| UpdateDocument | OpcenterEXFN_MasterData_Connector | JavaAction | 'Document' | 'UpdateDocument' |
| CreateDocument | OpcenterEXFN_MasterData_Connector | JavaAction | 'Document' | 'CreateDocument' |
| LinkCategories | OpcenterEXFN_MasterData_Connector | JavaAction | 'Document' | 'LinkCategories' |
| SetDocumentFile | OpcenterEXFN_MasterData_Connector | JavaAction | 'Document' | 'SetDocumentFile' |
| UnlinkCategories | OpcenterEXFN_MasterData_Connector | JavaAction | 'Document' | 'UnlinkCategories' |
| SetDocumentIcon | OpcenterEXFN_MasterData_Connector | JavaAction | 'Document' | 'SetDocumentIcon' |
| UploadDocuments | OpcenterEXFN_MasterData_Connector | JavaAction | 'Document' | 'UploadDocuments' |
| CreateNewDocumentRevision | OpcenterEXFN_MasterData_Connector | JavaAction | 'Document' | 'CreateNewDocumentRevision' |
| CreateExternalDocument | OpcenterEXFN_MasterData_Connector | JavaAction | 'Document' | 'CreateExternalDocument' |
| CopyDocumentRevision | OpcenterEXFN_MasterData_Connector | JavaAction | 'Document' | 'CopyDocumentRevision' |
| UpdateDocumentCategory | OpcenterEXFN_MasterData_Connector | JavaAction | 'Document' | 'UpdateDocumentCategory' |
| CreateDocumentCategory | OpcenterEXFN_MasterData_Connector | JavaAction | 'Document' | 'CreateDocumentCategory' |
| DeleteDocuments | OpcenterEXFN_MasterData_Connector | JavaAction | 'Document' | 'DeleteDocuments' |
| DefectEquipment_AssociateEquipmentConfigurationToFailure | OpcenterEXFN_MasterData_Connector | JavaAction | 'Defect' | 'DefectEquipment_AssociateEquipmentConfigurationToFailure' |
| CreateFailure | OpcenterEXFN_MasterData_Connector | JavaAction | 'Defect' | 'CreateFailure' |
| DefectMaterial_AssociateMaterialsToFailure | OpcenterEXFN_MasterData_Connector | JavaAction | 'Defect' | 'DefectMaterial_AssociateMaterialsToFailure' |
| CreateNewFailureRevision | OpcenterEXFN_MasterData_Connector | JavaAction | 'Defect' | 'CreateNewFailureRevision' |
| UpdateFailure | OpcenterEXFN_MasterData_Connector | JavaAction | 'Defect' | 'UpdateFailure' |
| AssociateQualityActionsWithFailure | OpcenterEXFN_MasterData_Connector | JavaAction | 'Defect' | 'AssociateQualityActionsWithFailure' |
| UpdateSubQualityAction | OpcenterEXFN_MasterData_Connector | JavaAction | 'Defect' | 'UpdateSubQualityAction' |
| CreateBaseQualityAction | OpcenterEXFN_MasterData_Connector | JavaAction | 'Defect' | 'CreateBaseQualityAction' |
| CreateSubQualityAction | OpcenterEXFN_MasterData_Connector | JavaAction | 'Defect' | 'CreateSubQualityAction' |
| UpdateBaseQualityAction | OpcenterEXFN_MasterData_Connector | JavaAction | 'Defect' | 'UpdateBaseQualityAction' |
| WIForOPCUA_Associate_DTO | OpcenterEXFN_MasterData_Connector | JavaAction | 'WorkInstruction' | 'WIForOPCUA_AssociateAutomationNodeInstanceToCharacteristicRepresentation' |
| CopyCharacteristicRepresentationRevision | OpcenterEXFN_MasterData_Connector | JavaAction | 'WorkInstruction' | 'CopyCharacteristicRepresentationRevision' |
| WIForOPCUA_AssociateAutomationNodeInstanceToCharacteristicRepresentation | OpcenterEXFN_MasterData_Connector | JavaAction | 'WorkInstruction' | 'WIForOPCUA_AssociateAutomationNodeInstanceToCharacteristicRepresentation' |
| AssociateReferencedMaterialsToCharacteristicRepresentation | OpcenterEXFN_MasterData_Connector | JavaAction | 'WorkInstruction' | 'AssociateReferencedMaterialsToCharacteristicRepresentation' |
| AssociateReferencedEquipmentConfigurationsToCharacteristicRepresentation | OpcenterEXFN_MasterData_Connector | JavaAction | 'WorkInstruction' | 'AssociateReferencedEquipmentConfigurationsToCharacteristicRepresentation' |
| CreateCharacteristicRepresentationRevision | OpcenterEXFN_MasterData_Connector | JavaAction | 'WorkInstruction' | 'CreateCharacteristicRepresentationRevision' |
| CreateCharacteristicRepresentation | OpcenterEXFN_MasterData_Connector | JavaAction | 'WorkInstruction' | 'CreateCharacteristicRepresentation' |
| Export | OpcenterEXFN_MasterData_Connector | JavaAction | 'System' | 'Export' |
| UpdateCharacteristicRepresentation | OpcenterEXFN_MasterData_Connector | JavaAction | 'WorkInstruction' | 'UpdateCharacteristicRepresentation' |
| LockControlMethod | OpcenterEXFN_MasterData_Connector | JavaAction | 'WorkInstruction' | 'LockControlMethod' |
| CreateControlMethodRevision | OpcenterEXFN_MasterData_Connector | JavaAction | 'WorkInstruction' | 'CreateControlMethodRevision' |
| UpdateControlMethod | OpcenterEXFN_MasterData_Connector | JavaAction | 'WorkInstruction' | 'UpdateControlMethod' |
| DeleteControlMethod | OpcenterEXFN_MasterData_Connector | JavaAction | 'WorkInstruction' | 'DeleteControlMethod' |
| CreateControlMethod | OpcenterEXFN_MasterData_Connector | JavaAction | 'WorkInstruction' | 'CreateControlMethod' |
| UnsetControlMethodRevisionCurrent | OpcenterEXFN_MasterData_Connector | JavaAction | 'WorkInstruction' | 'UnsetControlMethodRevisionCurrent' |
| UnlockControlMethod | OpcenterEXFN_MasterData_Connector | JavaAction | 'WorkInstruction' | 'UnlockControlMethod' |
| CopyControlMethodRevision | OpcenterEXFN_MasterData_Connector | JavaAction | 'WorkInstruction' | 'CopyControlMethodRevision' |
| SetControlMethodRevisionCurrent | OpcenterEXFN_MasterData_Connector | JavaAction | 'WorkInstruction' | 'SetControlMethodRevisionCurrent' |
| AssociateInspectionOperationsWithBillOfQuality | OpcenterEXFN_MasterData_Connector | JavaAction | 'WorkInstruction' | 'AssociateInspectionOperationsWithBillOfQuality' |
| CreateBillOfQualityRevision | OpcenterEXFN_MasterData_Connector | JavaAction | 'WorkInstruction' | 'CreateBillOfQualityRevision' |
| UpdateBillOfQuality | OpcenterEXFN_MasterData_Connector | JavaAction | 'WorkInstruction' | 'UpdateBillOfQuality' |
| DiassociateInspectionOperationsFromBillOfQuality_Custom | OpcenterEXFN_MasterData_Connector_2 | JavaAction | 'WorkInstruction' | 'DiassociateInspectionOperationsFromBillOfQuality' |
| CopyBillOfQualityRevision | OpcenterEXFN_MasterData_Connector | JavaAction | 'WorkInstruction' | 'CopyBillOfQualityRevision' |
| CreateBillOfQuality | OpcenterEXFN_MasterData_Connector | JavaAction | 'WorkInstruction' | 'CreateBillOfQuality' |
| UpdateVariableCharacteristicSpecification | OpcenterEXFN_MasterData_Connector | JavaAction | 'WorkInstruction' | 'UpdateVariableCharacteristicSpecification' |
| CreateVariableCharacteristicSpecification | OpcenterEXFN_MasterData_Connector | JavaAction | 'WorkInstruction' | 'CreateVariableCharacteristicSpecification' |
| UpdateAttributiveCharacteristicSpecification | OpcenterEXFN_MasterData_Connector | JavaAction | 'WorkInstruction' | 'UpdateAttributiveCharacteristicSpecification' |
| AssociateFailuresWithCharacteristicSpecification | OpcenterEXFN_MasterData_Connector | JavaAction | 'WorkInstruction' | 'AssociateFailuresWithCharacteristicSpecification' |
| CreateVisualCharacteristicSpecification | OpcenterEXFN_MasterData_Connector | JavaAction | 'WorkInstruction' | 'CreateVisualCharacteristicSpecification' |
| CreateAttributiveCharacteristicSpecification | OpcenterEXFN_MasterData_Connector | JavaAction | 'WorkInstruction' | 'CreateAttributiveCharacteristicSpecification' |
| UpdateVisualCharacteristicSpecification | OpcenterEXFN_MasterData_Connector | JavaAction | 'WorkInstruction' | 'UpdateVisualCharacteristicSpecification' |
| CreateLinkDocument | OpcenterEXFN_MasterData_Connector | JavaAction | 'Document' | 'CreateLinkDocument' |
| LinkDocument | OpcenterEXFN_MasterData_Connector | JavaAction | 'Document' | 'LinkDocument' |
| UpdateSegregationTag | OpcenterEXFN_MasterData_Connector | JavaAction | 'DataSegregation' | 'UpdateSegregationTag' |
| CreateSegregationTag | OpcenterEXFN_MasterData_Connector | JavaAction | 'DataSegregation' | 'CreateSegregationTag' |
| CreateEquipmentHierarchyConfiguration | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'CreateEquipmentGraphConfiguration' |
| BulkImportEquipmentHierarchy | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'BulkImportEquipmentHierarchy' |
| UpdateHierarchyConfiguration | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'UpdateEquipmentGraphConfiguration' |
| UpdateEquipmentLevel | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'UpdateEquipmentLevel' |
| CreateEquipmentLevel | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'CreateEquipmentLevel' |
| OPCUAConnect_AssociateAutomationNodeInstancesWithEquipmentConfiguration | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'OPCUAConnect_AssociateAutomationNodeInstancesWithEquipmentConfiguration' |
| AddEquipmentConfigurationStateMachines | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'AddEquipmentConfigurationStateMachines' |
| CreateEquipmentConfigurationPropertyAttribute | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'UpdateEquipmentConfigurationProperty' |
| UpdateEquipmentConfigurationPropertyAttribute | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'UpdateEquipmentConfigurationPropertyAttribute' |
| UpdateEquipmentConfiguration | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'UpdateEquipmentConfiguration' |
| SynchronizeEquipment | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'SynchronizeEquipment' |
| CreateEquipmentConfiguration | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'CreateEquipmentConfiguration' |
| AssociateEquipmentGroupConfigurationsWithEquipmentConfiguration | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'AssociateEquipmentGroupConfigurationsWithEquipmentConfiguration' |
| UpdateEquipmentConfigurationProperties | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'UpdateEquipmentConfigurationProperty' |
| CreateEquipmentConfigurationProperties | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'UpdateEquipmentConfiguration' |
| UpdateEquipmentConfigurationProperties | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'UpdateEquipmentConfigurationProperty' |
| BulkImportEquipmentGraphConfiguration | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'BulkImportEquipmentGraphConfiguration' |
| CreateEquipmentGraphConfiguration | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'CreateEquipmentGraphConfiguration' |
| UpdateEquipmentGraphConfiguration | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'UpdateEquipmentGraphConfiguration' |
| CreateEquipmentGraphNodeConfiguration | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'CreateEquipmentGraphNodeConfiguration' |
| UpdateEquipmentGraphLinkConfiguration | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'UpdateEquipmentGraphLinkConfiguration' |
| AssociateChildrenEquipmentGroupConfWithEquipmentGroupConf | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'AssociateChildrenEquipmentGroupConfWithEquipmentGroupConf' |
| AssociateParentsEquipmentGroupConfWithEquipmentGroupConf | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'AssociateParentsEquipmentGroupConfWithEquipmentGroupConf' |
| UpdateEquipmentGroupConfigurationScreen | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'UpdateEquipmentGroupConfiguration' |
| AssociateEquipmentConfigurationsWithEquipmentGroupConfiguration | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'AssociateEquipmentConfigurationsWithEquipmentGroupConfiguration' |
| CreateEquipmentConfigurationScreen | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'CreateEquipmentGroupConfiguration' |
| SetEquipmentPropertyAttributeValue | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'SetEquipmentPropertyAttributeValue' |
| SetEquipmentPropertyValue | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'SetEquipmentPropertyValue' |
| SetEquipmentStatus | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'SetEquipmentStatus' |
| AddEquipmentTypeStateMachines | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'AddEquipmentTypeStateMachines' |
| DeleteEquipmentTypeStateMachine | OpcenterEXFN_MasterData_Connector | MicroflowCall | <parameter> | <parameter> |
| CreateEquipmentTypePropertyAttribute | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'UpdateEquipmentTypeProperty' |
| UpdateEquipmentTypeProperty | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'UpdateEquipmentTypeProperty' |
| CreateEquipmentType | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'CreateEquipmentType' |
| OPCUAConnect_AssociateAutomationNodeTypesWithEquipmentType | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'OPCUAConnect_AssociateAutomationNodeTypesWithEquipmentType' |
| CreateEquipmentTypeProperty | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'UpdateEquipmentType' |
| UpdateEquipmentTypePropertyAttribute | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'UpdateEquipmentTypePropertyAttribute' |
| UpdateEquipmentType | OpcenterEXFN_MasterData_Connector | JavaAction | 'Equipment' | 'UpdateEquipmentType' |
| LinkHazardAndSafetyConfiguration | OpcenterEXFN_MasterData_Connector | JavaAction | 'HazardAndSafety' | 'LinkHazardAndSafetyConfiguration' |
| SUB_CreateLinkDocument | Documents | MicroflowCall | 'Document' | 'CreateLinkDocument' |
| SUB_LinkDocument | Documents | MicroflowCall | 'Document' | 'LinkDocument' |
| SUB_UpdateScenarioConfiguration | Documents | MicroflowCall | 'AuditTrail' | 'UpdateScenarioConfiguration' |
| SUB_UpdateSignatureConfigurations | Documents | MicroflowCall | 'AuditTrail' | 'UpdateSignatureConfigurations' |
| SUB_CreateSignatureConfigurations | Documents | MicroflowCall | 'AuditTrail' | 'CreateSignatureConfigurations' |
| SUB_CreateScenarioConfiguration | Documents | MicroflowCall | 'AuditTrail' | 'CreateScenarioConfiguration' |
| SUB_DeleteDocuments | OpcenterEXFN_MasterData_Connector | MicroflowCall | 'Document' | 'DeleteDocuments' |
| SUB_UpdateDocumentCategory | Documents | MicroflowCall | 'Document' | 'UpdateDocumentCategory' |
| SUB_CreateDocumentCategory | Documents | MicroflowCall | 'Document' | 'CreateDocumentCategory' |
| SUB_LinkCategories | Documents | MicroflowCall | 'Document' | 'LinkCategories'
 |
| SUB_CreateNewDocumentRevision | Documents | MicroflowCall | 'Document' | 'CreateNewDocumentRevision'
 |
| SUB_CreateDocument | Documents | MicroflowCall | 'Document' | 'CreateDocument'
 |
| SUB_CopyDocumentRevision | Documents | MicroflowCall | 'Document' | 'CopyDocumentRevision'
 |
| SUB_SetDocumentFile | Documents | MicroflowCall | 'Document' | 'SetDocumentFile'
 |
| SUB_UpdateDocument | Documents | MicroflowCall | 'Document' | 'UpdateDocument'
 |
| SUB_UnlinkCategories | Documents | MicroflowCall | 'Document' | 'UnlinkCategories'
 |
| SUB_SetDocumentIcon | Documents | MicroflowCall | 'Document' | 'SetDocumentIcon'
 |
| SUB_CreateExternalDocument | OpcenterEXFN_MasterData_Connector | MicroflowCall | 'Document' | 'CreateExternalDocument' |
| SUB_CreateMaterialProperty | Documents | MicroflowCall | 'Material' | 'CreateMaterialProperties' |
| SUB_UpdateMaterial | Documents | MicroflowCall | 'Material' | 'UpdateMaterial' |
| SUB_CreateBaseMaterial | Documents | MicroflowCall | 'Material' | 'CreateMaterial' |
| SUB_UpdateMaterialProperty | Documents | MicroflowCall | 'Material' | 'UpdateMaterialProperties' |
| SUB_CreateMaterialTrackingUnitAggregateTemplate | Documents | MicroflowCall | 'Personnel' | 'CreatePerson' |
| SUB_UpdateMaterialTrackingUnitAggregateTemplateProperty | Documents | MicroflowCall | 'Material' | 'UpdateMaterialTrackingUnitTemplateProperties' |
| SUB_UpdateMaterialTrackingUnitAggregateTemplate | Documents | MicroflowCall | 'Material' | 'UpdateMaterialTrackingUnitTemplate' |
| SUB_CreateMaterialTrackingUnitAggregateTemplateProperties | Documents | MicroflowCall | 'Material' | 'CreateMaterialTrackingUnitTemplateProperties' |
| SUB_AssociateMaterialWithMaterialGroup | Documents | MicroflowCall | 'Material' | 'AssociateMaterialWithMaterialGroup' |
| SUB_AssociateParentsMaterialGroupWithMaterialGroup | Documents | MicroflowCall | 'Material' | 'AssociateParentsMaterialGroupWithMaterialGroup' |
| SUB_CreateMaterialGroup | Documents | MicroflowCall | 'Material' | 'CreateMaterialGroup' |
| SUB_UpdateMaterialGroup | Documents | MicroflowCall | 'Material' | 'UpdateMaterialGroup' |
| SUB_AssociateChildrenMaterialGroupWithMaterialGroup | Documents | MicroflowCall | 'Material' | 'AssociateChildrenMaterialGroupWithMaterialGroup' |
| SUB_CreateMaterialTrackingUnitTemplateProperties | Documents | MicroflowCall | 'Material' | 'CreateMaterialTrackingUnitTemplateProperties' |
| SUB_CreateMaterialTrackingUnitTemplate | Documents | MicroflowCall | 'Personnel' | 'CreatePerson' |
| SUB_UpdateMaterialTrackingUnitTemplate | Documents | MicroflowCall | 'Material' | 'UpdateMaterialTrackingUnitTemplate' |
| SUB_UpdateMaterialTrackingUnitTemplateProperty | Documents | MicroflowCall | 'Material' | 'UpdateMaterialTrackingUnitTemplateProperties' |
| SUB_SetStatusMaterialTrackingUnitAggregate | Documents | MicroflowCall | 'Material' | 'SetMaterialTrackingUnitAggregateStatus' |
| SUB_UpdateMaterialTrackingUnitInMaterialTrackingUnitAggregate | Documents | MicroflowCall | 'Material' | 'UpdateMaterialTrackingUnit' |
| SUB_LoadMaterialTrackingUnitAggregateToMaterialTrackingUnitAggregate | Documents | MicroflowCall | 'Material' | 'LoadMaterialTrackingUnitAggregateToMaterialTrackingUnitAggregate' |
| SUB_CreateMaterialTrackingUnitAggregateProperties | Documents | MicroflowCall | 'Material' | 'CreateMaterialTrackingUnitAggregateProperties' |
| SUB_UpdateMaterialTrackingUnitAggregateProperty | Documents | MicroflowCall | 'Material' | 'UpdateMaterialTrackingUnitAggregateProperties' |
| SUB_CreateMaterialTrackingUnitInMaterialTrackingUnitAggregate | Documents | MicroflowCall | 'Material' | 'CreateMaterialTrackingUnit' |
| SUB_CreateMaterialTrackingUnitAggregate | Documents | MicroflowCall | 'Material' | 'CreateMaterialTrackingUnitAggregate' |
| SUB_SetQuantityMaterialTrackingUnitAggregate | Documents | MicroflowCall | 'Material' | 'SetMaterialTrackingUnitAggregateQuantity' |
| SUB_MoveMaterialTrackingUnitToMaterialTrackingUnitAggregate | OpcenterEXFN_MasterData_Connector | MicroflowCall | 'Material' | 'MoveMaterialTrackingUnitToMaterialTrackingUnitAggregate' |
| SUB_UpdateMaterialTrackingUnitAggregate | Documents | MicroflowCall | 'Material' | 'UpdateMaterialTrackingUnitAggregate' |
| SUB_SetMaterialOfMaterialTrackingUnitAggregate | Documents | MicroflowCall | 'Material' | 'SetMaterialTrackingUnitAggregateMaterial' |
| SUB_SetStateMachineMaterialTrackingUnitAggregate | Documents | MicroflowCall | 'Material' | 'SetMaterialTrackingUnitAggregateStateMachine' |
| SUB_LoadEquipmentToMaterialTrackingUnitAggregate | Documents | MicroflowCall | 'Material' | 'LoadMaterialTrackingUnitAggregateToEquipment' |
| SUB_CreateMaterialLot | Documents | MicroflowCall | 'Material' | 'CreateMaterialLot' |
| SUB_UpdateMaterialLot | Documents | MicroflowCall | 'Material' | 'UpdateMaterialLot' |
| SUB_SetStateMachineMaterialLot | OpcenterEXFN_MasterData_Connector | MicroflowCall | 'Material' | 'SetMaterialLotStateMachine' |
| SUB_CreateMaterialLotProperty | OpcenterEXFN_MasterData_Connector | MicroflowCall | 'Material' | 'CreateMaterialLotProperties' |
| SUB_UpdateMaterialLotProperty | Documents | MicroflowCall | 'Material' | 'UpdateMaterialLotProperties' |
| SUB_SetQuantityMaterialLot | Documents | MicroflowCall | 'Material' | 'SetMaterialLotQuantity' |
| SUB_SetMaterialMaterialLot | OpcenterEXFN_MasterData_Connector | MicroflowCall | 'Material' | 'SetMaterialLotMaterial' |
| SUB_SetMaterialLotStatus | OpcenterEXFN_MasterData_Connector | MicroflowCall | 'Material' | 'SetMaterialLotStatus' |
| SUB_CreateMaterialLotTemplate | Documents | MicroflowCall | 'Material' | 'CreateMaterialLotTemplate' |
| SUB_CreateMaterialLotTemplateProperties | OpcenterEXFN_MasterData_Connector | MicroflowCall | 'Material' | 'CreateMaterialLotTemplateProperties' |
| SUB_UpdateMaterialLotTemplate | OpcenterEXFN_MasterData_Connector | MicroflowCall | 'Material' | 'UpdateMaterialLotTemplate' |
| SUB_UpdateMaterialLotTemplateProperties | OpcenterEXFN_MasterData_Connector | MicroflowCall | 'Material'
 | 'UpdateMaterialLotTemplateProperties' |
| SUB_DisassociateMaterialTrackingUnitsFromMaterialLot | Documents | MicroflowCall | 'Material' | 'DisassociateMaterialTrackingUnitsFromMaterialLot' |
| SUB_SetMaterialOfMaterialTrackingUnit | Documents | MicroflowCall | 'Material' | 'SetMaterialTrackingUnitMaterial' |
| SUB_SetStateMachineMaterialTrackingUnit | Documents | MicroflowCall | 'Material' | 'SetMaterialTrackingUnitStateMachine' |
| SUB_CreateMaterialTrackingUnit | Documents | MicroflowCall | 'Material' | 'CreateMaterialTrackingUnit' |
| SUB_AssociateMaterialTrackingUnitsWithMaterialLot | Documents | MicroflowCall | 'Material' | 'AssociateMaterialTrackingUnitsWithMaterialLot' |
| SUB_SetQuantityMaterialTrackingUnit | Documents | MicroflowCall | 'Material' | 'SetMaterialTrackingUnitQuantity' |
| SUB_UpdateMaterialTrackingUnit | Documents | MicroflowCall | 'Material' | 'UpdateMaterialTrackingUnit' |
| SUB_MoveMaterialTrackingUnitToEquipment | Documents | MicroflowCall | 'Material' | 'MoveMaterialTrackingUnitToEquipment' |
| SUB_CreateMaterialTrackingUnitProperties | Documents | MicroflowCall | 'Material' | 'CreateMaterialTrackingUnitProperties' |
| SUB_UpdateMaterialTrackingUnitProperties | Documents | MicroflowCall | 'Material' | 'UpdateMaterialTrackingUnitProperties' |
| SUB_SetStatusMaterialTrackingUnit | Documents | MicroflowCall | 'Material' | 'SetMaterialTrackingUnitStatus' |
| SUB_MoveMaterialTrackingUnitToMaterialTrackingUnitAggregate_2 | Documents | MicroflowCall | 'Material' | 'MoveMaterialTrackingUnitToMaterialTrackingUnitAggregate' |
| SUB_CreateMaterialTemplate | OpcenterEXFN_MasterData_Connector | MicroflowCall | 'Material' | 'CreateMaterialTemplate' |
| SUB_UpdateMaterialTemplateProperties | Documents | MicroflowCall | 'Material' | 'UpdateMaterialTemplateProperties' |
| SUB_CreateMaterialTemplateProperties | Documents | MicroflowCall | 'Material' | 'CreateMaterialTemplateProperties' |
| SUB_UpdateMaterialTemplate | Documents | MicroflowCall | 'Material' | 'UpdateMaterialTemplate' |
| SUB_AssociateReferencedEquipmentConfigurationsToCharacteristicRepresentation | Documents | MicroflowCall | 'WorkInstruction' | 'AssociateReferencedEquipmentConfigurationsToCharacteristicRepresentation' |
| SUB_DisassociateCommand_ReferencedMaterialsToCharacteristicRepresentation | OpcenterEXFN_MasterData_Connector | MicroflowCall | <parameter> | <parameter> |
| SUB_AssociateReferencedMaterialsToCharacteristicRepresentation | Documents | MicroflowCall | 'WorkInstruction' | 'AssociateReferencedMaterialsToCharacteristicRepresentation' |
| SUB_UpdateCharacteristicRepresentation | Documents | MicroflowCall | 'WorkInstruction' | 'UpdateCharacteristicRepresentation' |
| SUB_CreateCharacteristicRepresentationRevision | Documents | MicroflowCall | 'WorkInstruction' | 'CreateCharacteristicRepresentationRevision' |
| SUB_CreateCharacteristicRepresentation | Documents | MicroflowCall | 'WorkInstruction' | 'CreateCharacteristicRepresentation' |
| SUB_CopyCharacteristicRepresentationRevision | Documents | MicroflowCall | 'WorkInstruction' | 'CopyCharacteristicRepresentationRevision' |
| SUB_UpdateVariableCharacteristicSpecification | Documents | MicroflowCall | 'WorkInstruction' | 'UpdateVariableCharacteristicSpecification' |
| SUB_CreateAttributiveCharacteristicSpecification | Documents | MicroflowCall | 'WorkInstruction' | 'CreateAttributiveCharacteristicSpecification' |
| SUB_CreateVariableCharacteristicSpecification | Documents | MicroflowCall | 'WorkInstruction' | 'CreateVariableCharacteristicSpecification' |
| SUB_AssociateFailuresWithCharacteristicSpecification | Documents | MicroflowCall | 'WorkInstruction' | 'AssociateFailuresWithCharacteristicSpecification' |
| SUB_UpdateVisualCharacteristicSpecification | Documents | MicroflowCall | 'WorkInstruction' | 'UpdateVisualCharacteristicSpecification' |
| SUB_CreateVisualCharacteristicSpecification | Documents | MicroflowCall | 'WorkInstruction' | 'CreateVisualCharacteristicSpecification' |
| SUB_UpdateAttributiveCharacteristicSpecification | Documents | MicroflowCall | 'WorkInstruction' | 'UpdateAttributiveCharacteristicSpecification' |
| SUB_CreateEquipmentLevel | Documents | MicroflowCall | 'Equipment' | 'CreateEquipmentLevel' |
| SUB_UpdateEquipmentLevel | Documents | MicroflowCall | 'Equipment' | 'UpdateEquipmentLevel' |
| SUB_AddEquipmentConfigurationStateMachines | Documents | MicroflowCall | 'Equipment' | 'AddEquipmentConfigurationStateMachines' |
| SUB_UpdateEquipmentConfiguration | Documents | MicroflowCall | 'Equipment' | 'UpdateEquipmentConfiguration' |
| SUB_CreateEquipmentConfigurationProperties | Documents | MicroflowCall | 'Equipment' | 'UpdateEquipmentConfiguration' |
| SUB_CreateEquipmentConfiguration | Documents | MicroflowCall | 'Equipment' | 'CreateEquipmentConfiguration' |
| SUB_SynchronizeEquipment | Documents | MicroflowCall | 'Equipment' | 'SynchronizeEquipment' |
| SUB_AssociateEquipmentGroupConfigurationsWithEquipmentConfiguration | Documents | MicroflowCall | 'Equipment' | 'AssociateEquipmentGroupConfigurationsWithEquipmentConfiguration' |
| SUB_UpdateEquipmentConfigurationProperties | Documents | MicroflowCall | 'Equipment' | 'UpdateEquipmentConfigurationProperty' |
| SUB_CreateEquipmentTypeProperties | Documents | MicroflowCall | 'Equipment' | 'UpdateEquipmentType' |
| SUB_CreateEquipmentType | OpcenterEXFN_MasterData_Connector | MicroflowCall | 'Equipment' | 'CreateEquipmentType' |
| SUB_UpdateEquipmentTypeProperties | Documents | MicroflowCall | 'Equipment' | 'UpdateEquipmentTypeProperty' |
| SUB_UpdateEquipmentType | Documents | MicroflowCall | 'Equipment' | 'UpdateEquipmentType' |
| SUB_AddEquipmentTypeStateMachines | Documents | MicroflowCall | 'Equipment' | 'AddEquipmentTypeStateMachines' |
| SUB_AssociateParentsEquipmentGroupConfWithEquipmentGroupConf | Documents | MicroflowCall | 'Equipment' | 'AssociateParentsEquipmentGroupConfWithEquipmentGroupConf' |
| SUB_AssociateEquipmentConfigurationsWithEquipmentGroupConfiguration | Documents | MicroflowCall | 'Equipment' | 'AssociateEquipmentConfigurationsWithEquipmentGroupConfiguration' |
| SUB_UpdateEquipmentGroupConfigurationScreen | Documents | MicroflowCall | 'Equipment' | 'UpdateEquipmentGroupConfiguration' |
| SUB_CreateEquipmentConfigurationScreen | Documents | MicroflowCall | 'Equipment' | 'CreateEquipmentGroupConfiguration' |
| SUB_AssociateChildrenEquipmentGroupConfWithEquipmentGroupConf | Documents | MicroflowCall | 'Equipment' | 'AssociateChildrenEquipmentGroupConfWithEquipmentGroupConf' |
| SUB_CreateEquipmentHierarchyConfiguration | Documents | MicroflowCall | 'Equipment' | 'CreateEquipmentGraphConfiguration' |
| SUB_UpdateHierarchyConfiguration | Documents | MicroflowCall | 'Equipment' | 'UpdateEquipmentGraphConfiguration' |
| SUB_UpdateEquipmentGraphConfiguration | Documents | MicroflowCall | 'Equipment' | 'UpdateEquipmentGraphConfiguration' |
| SUB_CreateEquipmentGraphConfiguration | Documents | MicroflowCall | 'Equipment' | 'CreateEquipmentGraphConfiguration' |
| SUB_UpdateBarcodeRulePart | Documents | MicroflowCall | 'Barcode' | 'UpdateBarcodeRulePart' |
| SUB_CreateBarcodeRulePart | Documents | MicroflowCall | 'Barcode' | 'CreateBarcodeRulePart' |
| SUB_CreateBarcodeRule | Documents | MicroflowCall | 'Barcode' | 'CreateBarcodeRule' |
| SUB_UpdateBarcodeRule | Documents | MicroflowCall | 'Barcode' | 'UpdateBarcodeRule' |
| SUB_MoveBarcodeRulePart | OpcenterEXFN_MasterData_Connector | MicroflowCall | 'Barcode' | 'MoveBarcodeRulePart' |
| SUB_CopyLabelPrinter | OpcenterEXFN_MasterData_Connector | MicroflowCall | 'Label' | 'CopyLabelPrinter' |
| SUB_CreateLabelPrinter | OpcenterEXFN_MasterData_Connector | MicroflowCall | 'Label' | 'CreateLabelPrinter'
 |
| SUB_UpdateLabelPrinter | OpcenterEXFN_MasterData_Connector | MicroflowCall | 'Label' | 'UpdateLabelPrinter' |
| SUB_CreateLabelType | OpcenterEXFN_MasterData_Connector | MicroflowCall | 'Label' | 'CreateLabelType' |
| SUB_UpdateLabelType | OpcenterEXFN_MasterData_Connector | MicroflowCall | 'Label' | 'UpdateLabelType' |
| SUB_CopyLabelType | OpcenterEXFN_MasterData_Connector | MicroflowCall | 'Label' | 'CopyLabelType' |
| SUB_CreateLabelTemplate | OpcenterEXFN_MasterData_Connector | MicroflowCall | 'Label' | 'CreateLabelTemplate' |
| SUB_RemovePrinter | OpcenterEXFN_MasterData_Connector | MicroflowCall | 'Label' | 'RemovePrinter' |
| SUB_UpdateLabelTemplate | OpcenterEXFN_MasterData_Connector | MicroflowCall | 'Label' | 'UpdateLabelTemplate' |
| SUB_SetDefaultLabelTemplate | OpcenterEXFN_MasterData_Connector | MicroflowCall | 'Label' | 'SetDefaultLabelTemplate' |
| SUB_AddPrinterToTemplate | OpcenterEXFN_MasterData_Connector | MicroflowCall | 'Label' | 'AddPrinterToTemplate' |
| SUB_SetDefaultLabelPrinter | OpcenterEXFN_MasterData_Connector | MicroflowCall | 'Label' | 'SetDefaultLabelPrinter' |
| SUB_CreateLabelTag | OpcenterEXFN_MasterData_Connector | MicroflowCall | 'Label' | 'CreateLabelTag' |
| SUB_UpdateLabelTag | OpcenterEXFN_MasterData_Connector | MicroflowCall | 'Label' | 'UpdateLabelTag' |
| SUB_CreateNewFailureRevision | Documents | MicroflowCall | 'Defect' | 'CreateNewFailureRevision' |
| SUB_CreateFailure | Documents | MicroflowCall | 'Defect' | 'CreateFailure' |
| SUB_UpdateFailure | Documents | MicroflowCall | 'Defect' | 'UpdateFailure' |
| SUB_DefectMaterial_AssociateMaterialsToFailure | Documents | MicroflowCall | 'Defect' | 'DefectMaterial_AssociateMaterialsToFailure' |
| SUB_AssociateQualityActionsWithFailure | Documents | MicroflowCall | 'Defect' | 'AssociateQualityActionsWithFailure' |
| SUB_DefectEquipment_AssociateEquipmentConfigurationToFailure | Documents | MicroflowCall | 'Defect' | 'DefectEquipment_AssociateEquipmentConfigurationToFailure' |
| SUB_CreateSubQualityAction | Documents | MicroflowCall | 'Defect' | 'CreateSubQualityAction' |
| SUB_UpdateBaseQualityAction | Documents | MicroflowCall | 'Defect' | 'UpdateBaseQualityAction' |
| SUB_CreateBaseQualityAction | Documents | MicroflowCall | 'Defect' | 'CreateBaseQualityAction' |
| SUB_UpdateSubQualityAction | Documents | MicroflowCall | 'Defect' | 'UpdateSubQualityAction' |
| SUB_DeleteCommand | OpcenterEXFN_MasterData_Connector | MicroflowCall | <parameter> | <parameter> |
| SUB_ResetCounterCommand | OpcenterEXFN_MasterData_Connector | MicroflowCall | <parameter> | <parameter> |
| SUB_DeleteCommand_StartArray | OpcenterEXFN_MasterData_Connector | MicroflowCall | <parameter> | <parameter> |
| SUB_SynchronizeCommand | OpcenterEXFN_MasterData_Connector | MicroflowCall | <parameter> | <parameter> |
| SUB_FreezeAndUnfreezeCommand | OpcenterEXFN_MasterData_Connector | MicroflowCall | <parameter> | <parameter> |
| SUB_DeleteCommand_StartArrayWithTwoParameter | OpcenterEXFN_MasterData_Connector | MicroflowCall | <parameter> | <parameter> |
| SUB_DeleteCommand_TwoParameter | OpcenterEXFN_MasterData_Connector | MicroflowCall | <parameter> | <parameter> |
| SUB_JoinAndDisjoinCommand | OpcenterEXFN_MasterData_Connector | MicroflowCall | <parameter> | <parameter> |
| CancelKanbanCall | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'Kanban' | 'CancelKanbanCall' |
| CreateOnlyManualKanbanCall | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'Kanban' | 'CreateOnlyManualKanbanCall' |
| CreateManualKanbanCall | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'Kanban' | 'CreateManualKanbanCall' |
| UpdateLSPAvailableQuantity | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'Kanban' | 'UpdateLSPAvailableQuantity' |
| IncreaseKanbanCallPriority | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'Kanban' | 'IncreaseKanbanCallPriority' |
| DeclareKanbanCallDelivered | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'Kanban' | 'DeclareKanbanCallDelivered' |
| CreateWorkOrderHistory | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'CreateWorkOrderHistory'
 |
| UpdateWorkOrderHistory | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UpdateWorkOrderHistory'
 |
| UpdateWorkOrderOperationActualExecutionTimes | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UpdateWorkOrderOperationActualExecutionTimes'
 |
| DeleteWorkOrderHistory | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'DeleteWorkOrderHistory'
 |
| AcceptMaterialRequest | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'AcceptMaterialRequest' |
| RejectMaterialRequest | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'RejectMaterialRequest' |
| RejectChangePackage | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'RejectChangePackage' |
| AcceptChangePackage | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'UADMAcceptChangePackage' |
| AddChangePackageItemWIDefinition | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'AddChangePackageItemWIDefinition' |
| AddChangePackageItemQualityInspection | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'AddChangePackageItemQualityInspection' |
| AddChangePackageItemMachine | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'AddChangePackageItemMachine' |
| UADMEditChangePackageItemStepWIDefinition | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'UADMEditChangePackageItemStepWIDefinition' |
| UADMEditChangePackageItemWIDefinition | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'UADMEditChangePackageItemWIDefinition' |
| AddChangePackageItemWorkOrderStep | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'AddChangePackageItemWorkOrderStep' |
| AddChangePackageItemDocument | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'AddChangePackageItemDocument' |
| AddChangePackageItemNewDocument | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'AddChangePackageItemNewDocument' |
| AddChangePackageItemStepWIDefinition | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'AddChangePackageItemStepWIDefinition' |
| AddChangePackageItemMaterial | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'AddChangePackageItemMaterial' |
| AddChangePackageItemTool | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'AddChangePackageItemTool' |
| BulkMaterialTrackingUnitsReleaseFromNonConformanceItems | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'BulkMaterialTrackingUnitsReleaseFromNonConformanceItems'
 |
| BulkMaterialTrackingUnitsScrapFromNonConformanceItems | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'BulkMaterialTrackingUnitsScrapFromNonConformanceItems'
 |
| CreateWorkingSession | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'CreateWorkingSession'
 |
| BulkMaterialTrackingUnitsReleaseFromFile | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'BulkMaterialTrackingUnitsReleaseFromFile'
 |
| AddMaterialTrackingUnitsToWorkingSession | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'AddMaterialTrackingUnitsToWorkingSession'
 |
| RemoveMaterialTrackingUnitsFromWorkingSession | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'RemoveMaterialTrackingUnitsFromWorkingSession'
 |
| SetContainmentRequestStatus | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'SetContainmentRequestStatus'
 |
| ManageContainmentRequestReleaseFromFile | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'ManageContainmentRequestReleaseFromFile'
 |
| GenerateOPNumberingPattern | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'GenerateOPNumberingPattern'
 |
| BulkMaterialTrackingUnitsScrapFromFile | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'BulkMaterialTrackingUnitsScrapFromFile'
 |
| ManageContainmentRequestRelease | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'ManageContainmentRequestRelease'
 |
| ManageContainmentRequestFullRelease | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'ManageContainmentRequestFullRelease'
 |
| BulkContainmentRequestRelease | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'BulkContainmentRequestRelease'
 |
| BulkMaterialTrackingUnitsScrap | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'BulkMaterialTrackingUnitsScrap'
 |
| CreateContainmentRequest | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'CreateContainmentRequest'
 |
| BulkMaterialTrackingUnitsScrapFromContainmentRequest | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'BulkMaterialTrackingUnitsScrapFromContainmentRequest'
 |
| UnLinkSkillsFromTeam_Custom | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'UnLinkSkillsFromTeam' |
| UADMUnLinkUsersFromTeam_Custom | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'UADMUnLinkUsersFromTeam' |
| UADMUnLinkUsersFromTeam | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UADMUnLinkUsersFromTeam' |
| CreateTeam | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'CreateTeam' |
| UpdateTeam | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UpdateTeam' |
| UpdateTeam_Custom | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'UpdateTeam' |
| UADMLinkUsersToTeam | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UADMLinkUsersToTeam' |
| UnLinkSkillsFromTeam | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UnLinkSkillsFromTeam' |
| UADMLinkSkillsToTeam_Custom | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'UADMLinkSkillsToTeam' |
| UADMLinkUsersToTeam_Custom | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'UADMLinkUsersToTeam' |
| CreateTeam_Custom | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'CreateTeam' |
| UADMLinkSkillsToTeam | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UADMLinkSkillsToTeam' |
| DeleteTeam_Custom | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'DeleteTeam' |
| DeleteTeam | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'DeleteTeam' |
| DeleteLSPThreshold | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'Kanban' | 'DeleteLSPThreshold' |
| DeleteLineSidePosition | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'Kanban' | 'DeleteLineSidePosition' |
| UnreleaseLineSidePosition | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'Kanban' | 'UnreleaseLineSidePosition' |
| ReleaseLineSidePosition | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'Kanban' | 'ReleaseLineSidePosition' |
| CreateLineSidePosition_NotOnlyManual | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'Kanban' | 'CreateLineSidePosition' |
| CreateLSPThreshold | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'Kanban' | 'CreateLSPThreshold' |
| UpdateLSPThreshold | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'Kanban' | 'UpdateLSPThreshold' |
| CreateLineSidePosition_OnlyManual | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'Kanban'
 | 'CreateLineSidePosition'
 |
| UpdateLineSidePosition_OnlyManual | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'Kanban'
 | 'UpdateLineSidePosition'
 |
| UpdateLineSidePosition_NotOnlyManual | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'Kanban' | 'UpdateLineSidePosition' |
| CreateNonConformanceAttachmentList | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'CreateNonConformanceAttachmentList'
 |
| ERPOrderBoP_CreateBoPERPOrder | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'ERPOrder' | 'ERPOrderBoP_CreateBoPERPOrder' |
| ERPOrderBoP_SetBoPInfoToERPOrder | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'ERPOrder' | 'ERPOrderBoP_SetBoPInfoToERPOrder' |
| SetERPOrderListStatus | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'ERPOrder' | 'SetERPOrderListStatus' |
| UpdateERPOrder | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'ERPOrder' | 'UpdateERPOrder' |
| CreateERPOrder | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'ERPOrder' | 'CreateERPOrder' |
| DeleteImportedContainmentRequestResult | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'DeleteImportedContainmentResults'
 |
| GetContainmentReleaseBySequence | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'GetContainmentReleaseBySequence'
 |
| GetContainmentRequestBySequence | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'GetContainmentRequestBySequence'
 |
| CreateHandlingUnit | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'CreateHandlingUnit'
 |
| UADMLinkMaterialItemsToHandlingUnit | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UADMLinkMaterialItemsToHandlingUnit'
 |
| ChangeHandlingUnitStatus | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'ChangeHandlingUnitStatus'
 |
| UnLoadHandlingUnit | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UnLoadHandlingUnit'
 |
| UnLinkMaterialItemsFromHandlingUnit | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UnLinkMaterialItemsFromHandlingUnit'
 |
| CreateWorkOrderDependency | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'CreateWorkOrderDependency'
 |
| UADMAcceptChangeReplaceToBeConsumedMaterial | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'UADMAcceptChangeReplaceToBeConsumedMaterial' |
| UADMAcceptChangeChangeDependency | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'UADMAcceptChangeChangeDependency' |
| UADMAcceptChangeAddOperation | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'UADMAcceptChangeAddOperation' |
| UADMCreateChangeNonConformance | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UADMCreateChangeNonConformance' |
| UADMAcceptChangeAddToBeConsumedMaterial | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'UADMAcceptChangeAddToBeConsumedMaterial' |
| UADMAcceptChangeChangeToBeConsumedMaterialQuantity | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'UADMAcceptChangeChangeToBeConsumedMaterialQuantity' |
| AcceptChangeAddProcessOperation | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'AcceptChangeAddProcessOperation' |
| GetProductionProcessFromAsPlanned | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM'
 | 'GetProductionProcessFromAsPlanned'
 |
| UADMAcceptChangeRemoveDependency | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'UADMAcceptChangeRemoveDependency' |
| UADMAcceptChangeRemoveOperation | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'UADMAcceptChangeRemoveOperation' |
| UADMAcceptChangeRemoveToBeConsumedMaterial | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'UADMAcceptChangeRemoveToBeConsumedMaterial' |
| UADMAcceptChangeRepeatOperation | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UADMAcceptChangeRepeatOperation'
 |
| UADMAcceptChangeAddWorkInstruction | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'UADMAcceptChangeAddWorkInstruction' |
| UADMAcceptChangeAddWorkInstruction | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'UADMAcceptChangeAddWorkInstruction' |
| UADMRejectChangeNonConformance | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'UADMRejectChangeNonConformance' |
| UADMAcceptChangeAddToBeUsedTool | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'UADMAcceptChangeAddToBeUsedTool' |
| UADMCreateWorkOrdersFromAsPlannedBOP | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UADMCreateWorkOrdersFromAsPlannedBOP'
 |
| UADMAbruptlyCloseFlexibleWorkOrder | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UADMAbruptlyCloseFlexibleWorkOrder' |
| UADMDisAssignProducedMaterialitems | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UADMDisAssignProducedMaterialitems' |
| UADMDisAssignProducedMaterialitems_Custom | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'UADMDisAssignProducedMaterialitems' |
| SplitWorkOrder | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UADMSplitWorkOrder'
 |
| CreateWOOpDependencyNavigationList | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'CreateWOOpDependencyNavigationList'
 |
| AutoGenerateWorkOrderNId | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'AutoGenerateWorkOrderNId'
 |
| CreateWorkOrderOutMsg_Custom | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'CreateWorkOrderOutMsg' |
| CreateWorkOrderOutMsg | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'CreateWorkOrderOutMsg' |
| LinkWOOperationsToWOOFolder | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'LinkWOOperationsToWOOFolder'
 |
| LinkWIDefinitionsToWOStep | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM'
 | 'LinkWIDefinitionsToWOStep'
 |
| UnlinkWorkInstructionsFromWOStep | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UnlinkWorkInstructionsFromWOStep' |
| CreateToBeUsedMachine_Custom | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'CreateToBeUsedMachine' |
| CreateToBeUsedMachine | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'CreateToBeUsedMachine' |
| UnlinkWorkOrderHumanResource | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UnlinkWorkOrderHumanResource' |
| UnlinkUserToWorkOrderOperationList | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UnlinkUserToWorkOrderOperationList' |
| CreateWorkOrderOperationFromMPProcessOperation | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'CreateWorkOrderOperationFromMPProcessOperation'
 |
| CreateWorkOrderSteps | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'CreateWorkOrderSteps' |
| UnLinkToBeUsedMachineToPartProgram | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UnLinkToBeUsedMachineToPartProgram' |
| DeleteToBeUsedMachine_Custom | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'DeleteToBeUsedMachine' |
| DeleteToBeUsedMachine | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'DeleteToBeUsedMachine' |
| LinkWIDefinitionsToWOOperation | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM'
 | 'LinkWIDefinitionsToWOOperation'
 |
| EditWorkOrderOperation | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'EditWorkOrderOperation'
 |
| CreateWorkOrderStepFromStep | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'CreateWorkOrderStepFromStep' |
| DeleteWorkOrderStep | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'DeleteWorkOrderStep' |
| LinkSkillsToWorkOrderStep | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'LinkSkillsToWorkOrderStep' |
| SetPreferredMachine_Custom | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'SetPreferredMachine' |
| SetPreferredMachine | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'SetPreferredMachine' |
| TriggerPrintingOnWorkOrderOperation | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'TriggerPrintingOnWorkOrderOperation' |
| LinkToBeUsedMachineToPartProgram | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'LinkToBeUsedMachineToPartProgram' |
| GetUserDetailsList | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'GetUserDetailsList'
 |
| CreateWOStepDependency | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'CreateWOStepDependency' |
| DeleteToBeUsedInspection | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'DeleteToBeUsedInspection' |
| TriggerPrintingOnWorkOrderStep | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'TriggerPrintingOnWorkOrderStep' |
| UpdateWorkOrderStep | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UpdateWorkOrderStep' |
| UADMCreateWorkOrderOperationFromProcessOperation | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UADMCreateWorkOrderOperationFromProcessOperation'
 |
| DeleteWorkOrderOperation | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'DeleteWorkOrderOperation' |
| DeleteWorkOrderOperation_Custom | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'DeleteWorkOrderOperation' |
| UnlinkWorkInstructionsFromWOOperation | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM'
 | 'UnlinkWorkInstructionsFromWOOperation'
 |
| LinkWorkOrderHumanResource | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'LinkWorkOrderHumanResource'
 |
| LinkSkillsToWorkOrderOperation | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'LinkSkillsToWorkOrderOperation'
 |
| LinkUserToWorkOrderOperationList | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'LinkUserToWorkOrderOperationList'
 |
| DeleteWorkOStepDependency | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'DeleteWorkOStepDependency' |
| EditWorkOrderHumanResource | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'EditWorkOrderHumanResource' |
| DeleteToBeConsumedMaterial_Custom | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'DeleteToBeConsumedMaterial' |
| DeleteToBeConsumedMaterial | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'DeleteToBeConsumedMaterial' |
| CreateRuntimeCharacteristicRepresentationContainer | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'CreateRuntimeCharacteristicRepresentationContainer'
 |
| CreateToBeConsumedMaterials_Custom | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'CreateToBeConsumedMaterials' |
| CreateToBeConsumedMaterials | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'CreateToBeConsumedMaterials' |
| DeleteToBeUsedTool | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'DeleteToBeUsedTool' |
| CreateWorkOrderOperationManually | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'CreateWorkOrderOperation' |
| CreateToBeUsedTools | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'CreateToBeUsedTools' |
| UnlinkWOOperationsFromWOOFolder | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UnlinkWOOperationsFromWOOFolder'
 |
| UnlinkWOOFoldersFromWOOFolder | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UnlinkWOOFoldersFromWOOFolder'
 |
| SetTargetQuantityOnFlexibleWorkOrder | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'SetTargetQuantityOnFlexibleWorkOrder' |
| SetTargetQuantityOnFlexibleWorkOrder_Custom | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'SetTargetQuantityOnFlexibleWorkOrder' |
| LinkItlkCheckToWorkOrderStepList | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'LinkItlkCheckToWorkOrderStepList'
 |
| DeleteWorkOrder | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'DeleteWorkOrder' |
| DeleteWOOFolderAndWOOpDependencies | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'DeleteWOOFolderAndWOOpDependencies'
 |
| UnlinkCollectedDocumentList | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UnlinkCollectedDocumentList'
 |
| CreateWorkOrderFromProcess | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UADMCreateWorkOrderFromProcess' |
| MarkForCleaningWorkOrderList_Custom | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'MarkForCleaningWorkOrderList' |
| MarkForCleaningWorkOrderList | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'MarkForCleaningWorkOrderList' |
| CreateWOOFolderAndWOOpDependencies | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'CreateWOOFolderAndWOOpDependencies' |
| AutoGenerateMTUCode | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'AutoGenerateMTUCode'
 |
| ChangeWorkOrderStatusToEdit | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'ChangeWorkOrderStatusToEdit' |
| ChangeWorkOrderStatusToEdit_Custom | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'ChangeWorkOrderStatusToEdit' |
| UnlinkItlkCheckToWorkOrderOperationList | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UnlinkItlkCheckToWorkOrderOperationList' |
| CreateWorkOrderFromMasterPlanWithQC | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'CreateWorkOrderFromMasterPlanWithQC'
 |
| UADMAbortWorkOrder | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UADMAbortWorkOrder' |
| UADMAbortWorkOrder_Custom | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'UADMAbortWorkOrder' |
| CreateWOOFolder | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'CreateWOOFolder'
 |
| EditWOOFolder | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'EditWOOFolder'
 |
| LinkWOOFoldersToWOOFolder | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'LinkWOOFoldersToWOOFolder'
 |
| TriggerPrintingOnWorkOrder | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'TriggerPrintingOnWorkOrder' |
| TriggerPrintingOnWorkOrder_Custom | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'TriggerPrintingOnWorkOrder' |
| EditWorkOrder | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'EditWorkOrder' |
| UpdateItlkCheckWOStepAssociationList | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UpdateItlkCheckWOStepAssociationList'
 |
| CreateWorkOrderFromMasterPlanWithEffectivity | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'CreateWorkOrderFromMasterPlanWithEffectivity'
 |
| LinkItlkCheckToWorkOrderOperationList | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'LinkItlkCheckToWorkOrderOperationList'
 |
| UADMCreateAndAssignProducedMaterialItems | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UADMCreateAndAssignProducedMaterialItems'
 |
| CreateWorkOrderFromMasterPlanBOMResolution | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'PLMCreateWorkOrderFromMasterPlanBOMResolution'
 |
| CreateWorkOrderHeader | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'CreateWorkOrderHeader'
 |
| AssignProducedMaterialItems | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'AssignProducedMaterialItems'
 |
| DeleteWOOFolders | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'DeleteWOOFolders'
 |
| UpdateItlkCheckWOOpAssociationList | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UpdateItlkCheckWOOpAssociationList'
 |
| UADMSetWorkOrderForScheduling | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UADMSetWorkOrderForScheduling' |
| UADMSetWorkOrderForScheduling_Custom | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'UADMSetWorkOrderForScheduling' |
| UADMReleaseWorkOrderList | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UADMReleaseWorkOrder' |
| UADMReleaseWorkOrder_Custom | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'UADMReleaseWorkOrder' |
| UnlinkItlkCheckToWorkOrderStepList | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UnlinkItlkCheckToWorkOrderStepList' |
| UpdatePreferredWOOpDependencyNavigation | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UpdatePreferredWOOpDependencyNavigation'
 |
| CreateWorkOrderManually | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'CreateWorkOrder' |
| MergeWOHeaderWithProcess | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'MergeWOHeaderWithProcess'
 |
| UADMCompleteNonProductiveActivityList | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UADMCompleteNonProductiveActivityList'
 |
| UADMRegisterNonProductiveActivityForUser | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UADMRegisterNonProductiveActivityForUser'
 |
| UADMUpdateNonProductiveActivityForUser | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UADMUpdateNonProductiveActivityForUser'
 |
| UADMDeleteNonProductiveActivityListForUser | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UADMDeleteNonProductiveActivityListForUser'
 |
| ReserveMaterialItems | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'ReserveMaterialItems' |
| UADMUpdateWOByBoPList | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'UADMUpdateWOByBoPList' |
| CreateBuffer | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'CreateBuffer'
 |
| UADMLinkMaterialItemsToBuffer | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UADMLinkMaterialItemsToBuffer'
 |
| DeleteBuffer | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'DeleteBuffer'
 |
| UpdateBuffer | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UpdateBuffer'
 |
| UnLinkMaterialItemFromBuffer | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UnLinkMaterialItemFromBuffer'
 |
| ChangeBufferStatus | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'ChangeBufferStatus'
 |
| QualityCharacteristicSPCEvaluation | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'WorkInstruction' | 'QualityCharacteristicSPCEvaluation'
 |
| CompleteFAIRecord | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'CompleteFAIRecord' |
| DeclareFAICandidate | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'DeclareFAICandidate' |
| CopyFAIRecord | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'CopyFAIRecord' |
| ChangeFAICandidate | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'ChangeFAICandidate' |
| AbortFAIRecord | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'AbortFAIRecord' |
| RevokeFAICandidate | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'RevokeFAICandidate' |
| SetReadyForSchedulingExecutionGroup | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'ScheduleExecutionGroup' |
| UnlinkWIDefinitionsFromEGPhase | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UnlinkWIDefinitionsFromEGPhase'
 |
| UpdateExecutionGroup | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'UpdateExecutionGroup' |
| RequestTransferPrintJobFile | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'PrintJobFile' | 'RequestTransferPrintJobFile' |
| SetAMPowderLoadedCheckOnEgPhase | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'PowderMgt' | 'SetAMPowderLoadedCheckOnEgPhase' |
| CreateExecution | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'UADMCreateExecutionGroup' |
| UnlinkWOOperationsFromExecutionGroup | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UnlinkWOOperationsFromExecutionGroup'
 |
| RemoveAMPowderLoadedCheckFromEgPhaseList | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'PowderMgt' | 'RemoveAMPowderLoadedCheckFromEgPhaseList' |
| LinkWOOperationListToExecutionGroup | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'LinkWOOperationListToExecutionGroup'
 |
| LinkWIDefinitionsToEGPhase | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'LinkWIDefinitionsToEGPhase' |
| GetReadyForPrintJob | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'AppU4DMAMN_GetReadyPrintJobs' |
| SetExecGroupListInEditStatus | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'SetExecGroupListInEditStatus' |
| ReleaseExecutionGroup | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'ReleaseExecutionGroup' |
| UpdateWOOpQuantityAssociatedToEG | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'UpdateWOOpQuantityAssociatedToEG' |
| AbortExecutionGroup | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'AbortExecutionGroupList' |
| LinkToBeUsedMachineToPJFForEGPhase | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'LinkToBeUsedMachineToPJFForEGPhase'
 |
| EditAMPowderLoadedCheckOnEgPhase | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'PowderMgt' | 'EditAMPowderLoadedCheckOnEgPhase' |
| UpdateExecutionGroupPhase | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'UpdateExecutionGroupPhase' |
| SUB_UADMCreateSnagAndNoteList | OpcenterEXDS_ProductionCoordination | MicroflowCall | 'AppU4DM' | 'UADMCreateSnagAndNoteList'
 |
| UADMReleaseWorkOrder | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UADMReleaseWorkOrder' |
| UADMReleaseWorkOrder_Custom_2 | OpcenterEXFN_MasterData_Connector_2 | JavaAction | 'AppU4DM' | 'UADMReleaseWorkOrder' |
| CreateToBeUsedMachine_2 | OpcenterEXFN_MasterData_Connector_2 | MicroflowCall | 'AppU4DM' | 'CreateToBeUsedMachine' |
| CreateToBeUsedMachine_Custom_2 | OpcenterEXFN_MasterData_Connector_2 | JavaAction | 'AppU4DM' | 'CreateToBeUsedMachine' |
| SetWorkOrderForScheduling | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'SetWorkOrderForScheduling' |
| SetWorkOrderForScheduling_Custom | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'SetWorkOrderForScheduling' |
| SetPreferredMachine_Custom_2 | OpcenterEXFN_MasterData_Connector_2 | JavaAction | 'AppU4DM' | 'SetPreferredMachine' |
| SetPreferredMachine_2 | OpcenterEXFN_MasterData_Connector_2 | MicroflowCall | 'AppU4DM' | 'SetPreferredMachine' |
| SUB_UADMConfirmSnagAndNoteList | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UADMConfirmSnagAndNoteList' |
| SetWorkOrderOperationEstimatedTimes | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'SetWorkOrderOperationEstimatedTimes' |
| SUB_LinkUserToWorkOrderOperationList | EXFN_ServiceLayer | MicroflowCall | 'AppU4DM' | 'LinkUserToWorkOrderOperationList'
 |
| SUB_GetUserDetailsList | EXFN_ServiceLayer | MicroflowCall | 'AppU4DM' | 'GetUserDetailsList'
 |
| SUB_UnlinkUserToWorkOrderOperationList | EXFN_ServiceLayer | MicroflowCall | 'AppU4DM' | 'UnlinkUserToWorkOrderOperationList'
 |
| UADMRemoveWorkOrderOperationFutureHoldList | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UADMRemoveWorkOrderOperationFutureHoldList'
 |
| UADMSetWorkOrderOperationFutureHoldList | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UADMSetWorkOrderOperationFutureHoldList'
 |
| AmendOfflineAction | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'AmendOfflineAction' |
| UpdateOfflineAction | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'UpdateOfflineAction' |
| DiscardOfflineSessionList | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'DiscardOfflineSessionList' |
| CompleteCheckIn | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'CompleteCheckIn' |
| ChangeOfflineSessionStatus | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'ChangeOfflineSessionStatus' |
| TraceProcessMeasurements | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'TraceProcessMeasurements' |
| UADMUpdateOutOfDateWOByBoPList | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UADMUpdateOutOfDateWOByBoPList'
 |
| PLMUpdateWorkOrderByCCList | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'PLMUpdateWorkOrderByCCList'
 |
| ReopenWorkOrderOperationSerialized | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'ReopenWorkOrderOperationSerialized'
 |
| UADMCreateDocument | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'UADMCreateDocument'
 |
| CreateToBeUsedDocuments | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'CreateToBeUsedDocuments'
 |
| AdministrativePause | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'AdministrativePause'
 |
| ReopenWorkOrderOperationTransferBatch | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'ReopenWorkOrderOperationTransferBatch'
 |
| DisassembleMaterialItem | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'DisassembleMaterialItem'
 |
| ReopenWorkOrderOperationList | OpcenterEXDS_ProductionCoordination_Connector | MicroflowCall | 'AppU4DM' | 'ReopenWorkOrderOperationList'
 |
| ResendConfirmWorkOrderOperation | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'APPU4DM' | 'ResendConfirmWorkOrderOperation' |
| ACT_DeleteWOOpDependencyNavigationList_WithConfirmation | OpcenterEXDS_ProductionCoordination | MicroflowCall | 'AppU4DM' | 'DeleteWOOpDependencyNavigationList' |
| ACT_UnlinkSkillsFromWOStep | OpcenterEXDS_ProductionCoordination | MicroflowCall | 'AppU4DM' | 'UnlinkSkillsToWorkOrderStep'
 |
| ACT_UnlinkSkillsFromWOOperation | OpcenterEXDS_ProductionCoordination | MicroflowCall | 'AppU4DM' | 'UnlinkSkillsFromWOOperation'
 |
| ACT_DeleteToBeUsedTool_WithConfirmation | OpcenterEXDS_Core | MicroflowCall | 'AppU4DM' | 'DeleteToBeUsedTool' |
| ACT_GenerateFAIRNumber | OpcenterEXDS_ProductionCoordination | MicroflowCall | 'AppU4DM' | 'GenerateFAIRNumber' |
| ACT_GenerateFAIRNumber_OLD | OpcenterEXDS_ProductionCoordination | MicroflowCall | 'AppU4DM' | 'GenerateFAIRNumber' |
| ACT_ReleaseLineSidePosition_WithConfirmation | OpcenterEXDS_ProductionCoordination | MicroflowCall | 'Kanban' | 'ReleaseLineSidePosition' |
| ACT_DeleteLineSidePosition_WithConfirmation | OpcenterEXDS_ProductionCoordination | MicroflowCall | 'Kanban' | 'DeleteLineSidePosition' |
| ACT_DeleteLSPThreshold_WithConfirmation | OpcenterEXDS_ProductionCoordination | MicroflowCall | 'Kanban' | 'DeleteLSPThreshold' |
| ACT_UnreleaseLineSidePosition_WithConfirmation | OpcenterEXDS_ProductionCoordination | MicroflowCall | 'Kanban' | 'UnreleaseLineSidePosition' |
| ACT_ChangeOfflineSessionStatus | OpcenterEXDS_ProductionCoordination | MicroflowCall | $IteratorOfflineAction/CommandApp | $IteratorOfflineAction/CommandName |
| ACT_ChangeOfflineSessionStatus | OpcenterEXDS_ProductionCoordination | MicroflowCall | $IteratorOfflineAction/CommandApp | $IteratorOfflineAction/CommandName |
| ACT_CheckInActionPayloadToCmd | OpcenterEXDS_ProductionCoordination | MicroflowCall | $CmdApp
 | <parameter> |
| ACT_CheckInActionPayloadToCmd | OpcenterEXDS_ProductionCoordination | MicroflowCall | $CmdApp
 | <parameter> |
| ACT_IncreaseKanbanCallPriority_WithConfirmation | OpcenterEXDS_ProductionCoordination | MicroflowCall | 'Kanban' | 'IncreaseKanbanCallPriority' |
| ACT_CreateManualKanbanCall_WithConfirmation | OpcenterEXDS_ProductionCoordination | MicroflowCall | 'Kanban' | 'CreateManualKanbanCall' |
| ACT_CreateOnlyManualKanbanCall_WithConfirmation | OpcenterEXDS_ProductionCoordination | MicroflowCall | 'Kanban' | 'CreateOnlyManualKanbanCall' |
| ACT_CancelKanbanCall_WithConfirmation | OpcenterEXDS_ProductionCoordination | MicroflowCall | 'Kanban' | 'CancelKanbanCall' |
| ACT_AddChangePackageWorkOrderOperation_Command | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'AddChangePackageItemWorkOrderOperation' |
| ACT_AddChangePackageWorkOrderOperationDependency_Command | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'AddChangePackageItemWorkOOperationDependency' |
| ACT_RemoveChangePackageWorkOrderOperation_Command | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'RemoveChangePackageItemWorkOrderOperation' |
| ACT_RemoveChangePackageWorkOrderOperationDependency_Command | OpcenterEXDS_ProductionCoordination_Connector | JavaAction | 'AppU4DM' | 'RemoveChangePackageItemWorkOOperationDependency' |
| DS_GetNcStatusToWithRoles | OpcenterEXDS_Core | MicroflowCall | 'AppU4DM' | 'GetCurrentUserRoles' |
| ACT_RemoveDefectsFromNonConformance_WithConfirmation | OpcenterEXDS_Core | MicroflowCall | 'AppU4DM' | 'RemoveDefectsFromNonConformance' |
| ACT_DeleteNonConformanceAttachmentList_WithConfirmation | OpcenterEXDS_Core | MicroflowCall | 'AppU4DM' | 'DeleteNonConformanceAttachmentList' |
| ACT_DeleteTool_WithConfirmation | OpcenterEXDS_Core | MicroflowCall | 'AppU4DM' | 'DeleteTool' |
| ACT_UnlinkSubstrateToMachine_WithConfirmation | OpcenterEXDS_ShopfloorExecution_Connector | MicroflowCall | 'AppU4DM' | 'UnlinkSubstrateToMachine' |
| ToolMaintenance | OpcenterEXDS_ShopfloorExecution_Connector | JavaAction | 'AppU4DM' | 'ToolMaintenance' |
| UpdateTool | OpcenterEXDS_ShopfloorExecution_Connector | JavaAction | 'AppU4DM' | 'UpdateTool' |
| LinkToolToAutomationNodeInstanceList | OpcenterEXDS_ShopfloorExecution_Connector | JavaAction | 'AppU4DM' | 'LinkToolToAutomationNodeInstanceList' |
| UADMCreateToolDefinition | OpcenterEXDS_ShopfloorExecution_Connector | MicroflowCall | 'AppU4DM' | 'UADMCreateToolDefinition' |
| UnlinkSubstrateToMachine | OpcenterEXDS_ShopfloorExecution_Connector | JavaAction | 'AppU4DM' | 'UnlinkSubstrateToMachine' |
| CreateLogisticClass | OpcenterEXDS_ShopfloorExecution_Connector | MicroflowCall | 'AppU4DM' | 'CreateLogisticClass' |
| LinkSubstrateToMachine | OpcenterEXDS_ShopfloorExecution_Connector | JavaAction | 'AppU4DM' | 'LinkSubstrateToMachine' |
| UnlinkToolToAutomationNodeInstanceList | OpcenterEXDS_ShopfloorExecution_Connector | JavaAction | 'AppU4DM' | 'UnlinkToolToAutomationNodeInstanceList' |
| CreateTool | OpcenterEXDS_ShopfloorExecution_Connector | JavaAction | 'AppU4DM' | 'CreateTool' |
| DeleteTool | OpcenterEXDS_ShopfloorExecution_Connector | JavaAction | 'AppU4DM' | 'DeleteTool' |
| RemoveDefectsFromNonConformance | OpcenterEXDS_ShopfloorExecution_Connector | JavaAction | 'AppU4DM' | 'RemoveDefectsFromNonConformance' |
| UADMSentenceNonConformanceV3_1 | OpcenterEXDS_ShopfloorExecution_Connector | JavaAction | 'AppU4DM' | 'UADMSentenceNonConformanceV3_1' |
| UADMCreateNonConformanceV3_1 | OpcenterEXDS_ShopfloorExecution_Connector | JavaAction | 'AppU4DM' | 'UADMCreateNonConformanceV3_1' |
| UADMLinkGenericItemsToNonConformance | OpcenterEXDS_ShopfloorExecution_Connector | JavaAction | 'AppU4DM' | 'UADMLinkGenericItemsToNonConformance' |
| UADMUpdateNonConformance | OpcenterEXDS_ShopfloorExecution_Connector | JavaAction | 'AppU4DM' | 'UADMUpdateNonConformance' |
| DeleteNonConformanceAttachmentList | OpcenterEXDS_ShopfloorExecution_Connector | JavaAction | 'AppU4DM' | 'DeleteNonConformanceAttachmentList' |
| UploadDocuments | OpcenterEXFN_MasterData_Connector | MicroflowCall | 'AppU4DM' | 'UploadDocuments' |
| UADMCreateNonConformanceDefectList | OpcenterEXDS_ShopfloorExecution_Connector | JavaAction | 'AppU4DM' | 'UADMCreateNonConformanceDefectList' |
| CreateNonConformanceAttachmentList | OpcenterEXDS_ShopfloorExecution_Connector | JavaAction | 'AppU4DM' | 'CreateNonConformanceAttachmentList' |
| UADMUnLinkGenericItemsFromNonConformance | OpcenterEXDS_ShopfloorExecution_Connector | JavaAction | 'AppU4DM' | 'UADMUnLinkGenericItemsFromNonConformance' |
| UADMExportGenealogy | OpcenterEXDS_ShopfloorExecution_Connector | MicroflowCall | 'AppU4DM' | 'UADMExportGenealogy' |
| UnlinkDocumentListToDM_MaterialTrackingUnit | OpcenterEXDS_ShopfloorExecution_Connector | MicroflowCall | 'AppU4DM' | 'UnlinkDocumentListToDM_MaterialTrackingUnit' |
| UADMCreateDocument | OpcenterEXDS_ShopfloorExecution | MicroflowCall | 'AppU4DM' | 'UADMCreateDocument' |
| UnlinkCollectedDocumentList | OpcenterEXDS_ShopfloorExecution_Connector | MicroflowCall | 'AppU4DM' | 'UnlinkCollectedDocumentList' |
| SUB_Signature_Abort | EXFN_ElectronicSignature | MicroflowCall | 'AuditTrail' | 'AbortSignature' |
| SUB_Signature_Prepare | EXFN_ElectronicSignature | MicroflowCall | 'AuditTrail' | 'PrepareSignature' |
| SUB_OnSign | EXFN_ElectronicSignature | MicroflowCall | 'AuditTrail' | 'AddSignature' |
| SUB_DeleteCommand_TwoParameter | OpcenterEXDS_Core | MicroflowCall | <parameter> | <parameter> |
| SUB_DeleteCommand | OpcenterEXDS_Core | MicroflowCall | <parameter> | <parameter> |
| SUB_ResetCounterCommand | OpcenterEXDS_Core | MicroflowCall | <parameter> | <parameter> |
| SUB_SynchronizeCommand | OpcenterEXDS_Core | MicroflowCall | <parameter> | <parameter> |
| SUB_DeleteCommand_StartArrayWithTwoParameter | OpcenterEXDS_Core | MicroflowCall | <parameter> | <parameter> |
| SUB_JoinAndDisjoinCommand | OpcenterEXDS_Core | MicroflowCall | <parameter> | <parameter> |
| SUB_FreezeAndUnfreezeCommand | OpcenterEXDS_Core | MicroflowCall | <parameter> | <parameter> |
| SUB_DeleteCommand_StartArray | OpcenterEXDS_Core | MicroflowCall | <parameter> | <parameter> |
| SUB_ImageGrid_NewDot_Create | EXFN_Quality | MicroflowCall | 'WorkInstruction' | $HelperExport/CommandName |
| ACT_CreateNewVisualSample | EXFN_Quality | MicroflowCall | 'WorkInstruction' | 'CreateVisualDetectedFailurewithES' |
| ACT_VisualInspection_Confirm | EXFN_Quality | MicroflowCall | 'WorkInstruction' | 'CreateVisualDetectedFailurewithES' |
| SUB_Attributive_Create_ES | EXFN_Quality | MicroflowCall | 'WorkInstruction' | 'CreateInspectionValuewithES' |
| SUB_Attributive_Update_ES | EXFN_Quality | MicroflowCall | 'WorkInstruction' | 'UpdateInspectionValuewithES' |
| SUB_Variable_Update_ES | EXFN_Quality | MicroflowCall | 'WorkInstruction' | 'UpdateInspectionValuewithES' |
| SUB_Variable_Create_ES | EXFN_Quality | MicroflowCall | 'WorkInstruction' | 'CreateInspectionValuewithES' |
| SUB_Context_Create | EXFN_Quality | MicroflowCall | 'WorkInstruction' | 'CreateInspectionAcquisitionContext' |
| ACT_ConfirmInspectionSample | EXFN_Quality | MicroflowCall | 'WorkInstruction' | 'ConfirmInspectionSample' |
| SUB_Definition_AssociateFailure | EXFN_Quality | MicroflowCall | 'WorkInstruction' | 'AssociateFailureToInspectionValue' |
| SUB_Definition_DisassociateFailure | EXFN_Quality | MicroflowCall | 'WorkInstruction' | 'DisassociateFailureFromInspectionValue' |
| SUB_WIStep_ReEdit | EXFN_WorkInstruction | MicroflowCall | @EXFN_WorkInstruction.AppName | 'ReEditWorkInstructionStep' |
| SUB_WIStep_Confirm | EXFN_WorkInstruction | MicroflowCall | @EXFN_WorkInstruction.AppName | 'ConfirmWIStep' |
| SUB_WIStep_Acknowledge | EXFN_WorkInstruction | MicroflowCall | @EXFN_WorkInstruction.AppName | 'AcknowledgeWIStep' |
| SUB_WorkInstruction_Create | EXFN_WorkInstruction | MicroflowCall | @EXFN_WorkInstruction.AppName | 'CreateWorkInstruction' |
| SUB_WorkInstruction_Delete | EXFN_WorkInstruction | MicroflowCall | @EXFN_WorkInstruction.AppName | 'DeleteWorkInstruction' |
| SUB_WorkInstructionStatus_InEditing | EXFN_WorkInstruction | MicroflowCall | @EXFN_WorkInstruction.AppName | 'InEditingWorkInstruction' |
| SUB_WorkInstructionStepItem_AutoSave | EXFN_WorkInstruction | MicroflowCall | @EXFN_WorkInstruction.AppName | 'AutoSaveWorkInstructionStepItemValue' |
| SUB_CalculateWorkInstructionFormulaValues | EXFN_WorkInstruction | MicroflowCall | 'WorkInstruction' | 'CalculateWorkInstructionFormulaValues' |

---

## 3. Signal Manager Subscriptions

Signal subscriptions from Signal Manager widgets (siemens.mxtosignal.MxToSignal).

Found 15 subscription(s):

| Module | Document Type | Document | Signal Name | App Name | Subscription Filter |
|--------|---------------|----------|-------------|----------|---------------------|
| OpcenterEXDS_ProductionCoordination | Page | WorkOrder_Master | WorkOrderStatusChanged | AppU4DM | No |
| OpcenterEXDS_ProductionCoordination | Page | WorkOrderOperation_Details | WorkOrderOperationStatusChanged | AppU4DM | No |
| OpcenterEXDS_ProductionCoordination | Page | WorkOrder_Details | WorkOrderStatusChanged | AppU4DM | No |
| OpcenterEXDS_ProductionCoordination | Page | WorkOrder_Details | WorkOrderOperationStatusChanged | AppU4DM | No |
| OpcenterEXDS_ProductionCoordination | Page | WorkOrder_AsBuilt | WorkOrderStatusChanged | AppU4DM | No |
| OpcenterEXDS_ProductionCoordination | Page | WorkOrder_AsBuilt | WorkOrderOperationStatusChanged | AppU4DM | No |
| OpcenterEXDS_ProductionCoordination | Page | WorkOrderUpdateCheck_Master | WorkOrderStatusChanged | AppU4DM | No |
| OpcenterEXDS_ProductionCoordination | Page | ContainmentRequests_Master | ContainmentRequestImportedSignal | AppU4DM | No |
| OpcenterEXDS_ProductionCoordination | Page | ContainmentRequests_Master | ContainmentRequestReleaseImportedSignal | AppU4DM | No |
| OpcenterEXDS_ProductionCoordination | Page | ContainmentRequestMonitoring_Master | ContainmentRequestImportedSignal | AppU4DM | No |
| OpcenterEXDS_ProductionCoordination | Page | ContainmentRequestMonitoring_Master | ContainmentRequestReleaseImportedSignal | AppU4DM | No |
| EXFN_Quality | Snippet | SNP_QualityInspectionContainer | OnCompleteInspectionSampleScenarioInstance | WorkInstruction | Yes |
| EXFN_Quality | Snippet | SNP_QualityInspectionContainer | InspectionSampleConfirmed | WorkInstruction | Yes |
| EXFN_Quality | Snippet | SNP_QualityInspectionContainer | DelayedExecution | WorkInstruction | Yes |
| EXFN_Quality | Snippet | SNP_QualityInspectionContainer | InspectionExecutionChrReprRuntimeNumberChanged | WorkInstruction | Yes |

---

## 4. Navigation Items

| Parent Node | Node | Target Page | User Roles |
|-------------|------|-------------|------------|
| - | Production Coordination | - | - |
| Production Coordination | Automation Node Viewer | OpcenterEXFN_MasterData.AutomationNodeViewer_Master | - |
| Production Coordination | Equipment | OpcenterEXFN_MasterData.Equipment_Master | - |
| Production Coordination | Inspection Orders | OpcenterEXFN_MasterData.InspectionOrders_Master | - |
| Production Coordination | Unlinked Documents | OpcenterEXFN_MasterData.UnlinkedDocument_Master | - |
| Production Coordination | SPC Evaluation | OpcenterEXDS_ProductionCoordination.SPCEvaluation_Master | Administrator, User |
| Production Coordination | Hold | OpcenterEXDS_Configuration.Hold_Master | Administrator, User |
| Production Coordination | Teams | OpcenterEXDS_ProductionCoordination.Teams_Master | Administrator, User |
| Production Coordination | Work Orders | OpcenterEXDS_ProductionCoordination.WorkOrder_Master | Administrator, User |
| Production Coordination | Work Order Update | OpcenterEXDS_ProductionCoordination.WorkOrderUpdate_Master | Administrator, User |
| Production Coordination | Work Order Update Check | OpcenterEXDS_ProductionCoordination.WorkOrderUpdateCheck_Master | Administrator, User |
| Production Coordination | Work Order Network | OpcenterEXDS_ProductionCoordination.WorkOrderNetwork_Master | Administrator, User |
| Production Coordination | Production Coordinator Dashboard | OpcenterEXDS_ProductionCoordination.ProductionCoordinatorDashboard_Master | Administrator, User |
| Production Coordination | Change Packages | OpcenterEXDS_ProductionCoordination.ChangePackages_Master | Administrator, User |
| Production Coordination | Change Requests | OpcenterEXDS_ProductionCoordination.ChangeRequest_Master | Administrator, User |
| Production Coordination | Line Side Positions | OpcenterEXDS_ProductionCoordination.LineSidePositions_Master | Administrator, User |
| Production Coordination | Line Side Positions Monitoring | OpcenterEXDS_ProductionCoordination.LineSidePositionsMonitoring_Master | Administrator, User |
| Production Coordination | Material Lots | OpcenterEXFN_MasterData.MaterialLot_Master | - |
| Production Coordination | Material Tracking Units | OpcenterEXDS_MasterData.MaterialTrackingUnit_Master | Administrator, User |
| Production Coordination | Material Tracking Unit Aggregates | OpcenterEXFN_MasterData.MaterialTrackingUnitAggregates_Master | - |
| Production Coordination | Users | OpcenterEXDS_ProductionCoordination.Users_Master | Administrator, User |
| Production Coordination | Future Hold | OpcenterEXDS_ProductionCoordination.FutureHoldManagement_Master | Administrator, User |
| Production Coordination | Production Coordinator Time Update | OpcenterEXDS_ProductionCoordination.PCTimeUpdate_Master | Administrator, User |
| Production Coordination | Logistic Requests | OpcenterEXDS_ProductionCoordination.LogisticRequest_Master | Administrator, User |
| Production Coordination | Handling Units | OpcenterEXDS_ProductionCoordination.HandlingUnits_Master | Administrator, User |
| Production Coordination | ERP Order | OpcenterEXDS_ProductionCoordination.ERPOrder_Master | Administrator, User |
| Production Coordination | Execution Group | OpcenterEXDS_ProductionCoordination.ExecutionGroup_Master | Administrator, User |
| Production Coordination | Offline Sessions | OpcenterEXDS_ProductionCoordination.OfflineSession_Master | Administrator, User |
| Production Coordination | Measurements | OpcenterEXDS_ProductionCoordination.Measurements_Master | Administrator, User |
| Production Coordination | Work Order Pre-Kitting | OpcenterEXDS_ProductionCoordination.WorkOrderPreKitting_Master | Administrator, User |
| Production Coordination | Integration Events | OpcenterEXDS_ProductionCoordination.IntegrationEvents_Master | Administrator, User |
| Production Coordination | Tools | OpcenterEXDS_ShopfloorExecution.Tool_Master | Administrator, User |
| Production Coordination | Non-Conformances | OpcenterEXDS_ShopfloorExecution.NonConformance_Master | Administrator, User |
| Production Coordination | Buffers | OpcenterEXDS_ProductionCoordination.Buffers_Master | Administrator, User |
| Production Coordination | First Article Inspections | OpcenterEXDS_ProductionCoordination.FirstArticleInspection_Master | Administrator, User |
| Production Coordination | Containment Requests | OpcenterEXDS_ProductionCoordination.ContainmentRequests_Master | Administrator, User |
| Production Coordination | Containment Requests Monitoring | OpcenterEXDS_ProductionCoordination.ContainmentRequestMonitoring_Master | Administrator, User |
| Production Coordination | Scan Material Tracking Units | OpcenterEXDS_ProductionCoordination.ScanMaterialTrackingUnits_Master | Administrator, User |
| Production Coordination | Genealogy | OpcenterEXDS_ShopfloorExecution.Genealogy_MasterV2 | Administrator, User |

---

## 5. Navigation Page Commands

Command bar actions extracted from navigation pages. Shows buttons in the vertical command bar of the Right placeholder.

Found commands in 31 page(s):

### OpcenterEXDS_ProductionCoordination.SPCEvaluation_Master

| Caption | Target Page | Target AppName | Target CommandName |
|---------|-------------|----------------|--------------------|
| Open | OpcenterEXDS_ProductionCoordination.SPCEvaluation_Details | - | - |

### OpcenterEXDS_ProductionCoordination.WorkOrderNetwork_Master

| Caption | Target Page | Target AppName | Target CommandName |
|---------|-------------|----------------|--------------------|
| Details | OpcenterEXDS_ProductionCoordination.WorkOrderRouting_Details | - | - |
| Add Dependency | OpcenterEXDS_ProductionCoordination.PANEL_CreateWorkOrderDependency | - | - |

### OpcenterEXDS_ProductionCoordination.LineSidePositionsMonitoring_Master

| Caption | Target Page | Target AppName | Target CommandName |
|---------|-------------|----------------|--------------------|
| Reload | - | - | - |
| Open | OpcenterEXDS_ProductionCoordination.LineSidePositionMonitoring_Details | - | CallCommand_MF |
| Request | - | - | CallCommand_MF |
| Manual Request | - | - | CallCommand_MF |
| Cancel Request | - | - | CallCommand_MF |
| Incrase Priority | - | - | CallCommand_MF |
| Update Quantity | OpcenterEXDS_ProductionCoordination.PANEL_UpdateLSPAvailableQuantity | - | - |

### OpcenterEXDS_MasterData.MaterialTrackingUnit_Master

| Caption | Target Page | Target AppName | Target CommandName |
|---------|-------------|----------------|--------------------|
| Create | OpcenterEXDS_MasterData.PANEL_CreateMaterialTrackingUnit | - | - |
| Open | OpcenterEXDS_MasterData.MaterialTrackingUnit_Details | - | SUB_FreezeAndUnfreezeCommand |
| Edit | OpcenterEXDS_MasterData.PANEL_EditMaterialTrackingUnit | - | - |
| Freeze | - | - | SUB_FreezeAndUnfreezeCommand |
| Set Status | OpcenterEXDS_MasterData.PANEL_SetStatusMaterialTrackingUnit | - | - |
| Set State Machine | OpcenterEXDS_MasterData.PANEL_SetMaterialTrackingUnitStateMachine | - | - |
| Set Material  | OpcenterEXDS_MasterData.PANEL_SetMaterialTrackingUnitStatus | - | - |
| Set Quantity | OpcenterEXDS_MasterData.PANEL_SetMaterialTrackingUnitQuantity | - | - |
| Associate Material Lot | OpcenterEXDS_MasterData.PANEL_AssociateMaterialTrackingUnitsWithMaterialLot | - | - |
| Disassociate Material Lot | - | - | SUB_DeleteCommand_StartArrayWithTwoParameter |
| Move To a MTU | OpcenterEXDS_MasterData.PANEL_MoveMaterialTrackingUnitToMaterialTrackingUnitAggregate | - | - |
| Move to Equipment | OpcenterEXDS_MasterData.PANEL_MoveMaterialTrackingUnitToEquipment | - | - |
| UnFreeze | - | - | SUB_FreezeAndUnfreezeCommand |
| Delete | - | - | SUB_DeleteCommand |
| Generate and Associte | - | - | - |
| Split | OpcenterEXDS_MasterData.PANEL_MTUSplit | - | - |

### OpcenterEXDS_ProductionCoordination.FutureHoldManagement_Master

| Caption | Target Page | Target AppName | Target CommandName |
|---------|-------------|----------------|--------------------|
| Create | OpcenterEXDS_ProductionCoordination.PANEL_UADMSetWorkOrderOperationFutureHoldList | - | - |
| Details | OpcenterEXDS_ProductionCoordination.FutureHoldManagement_Details | - | - |
| Close | - | - | - |

### OpcenterEXDS_ProductionCoordination.LogisticRequest_Master

| Caption | Target Page | Target AppName | Target CommandName |
|---------|-------------|----------------|--------------------|
| Accept | OpcenterEXDS_ProductionCoordination.PANEL_AcceptMaterialRequest | - | - |
| Reject | OpcenterEXDS_ProductionCoordination.PANEL_RejectMaterialRequest | - | - |

### OpcenterEXDS_ProductionCoordination.FirstArticleInspection_Master

| Caption | Target Page | Target AppName | Target CommandName |
|---------|-------------|----------------|--------------------|
| Refresh | - | - | - |
| Details | OpcenterEXDS_ProductionCoordination.FirstArticleInspection_Details | - | - |
| Declare Candidate | OpcenterEXDS_ProductionCoordination.PANEL_DeclareFAICandidate | - | - |
| Change Candidate | OpcenterEXDS_ProductionCoordination.PANEL_ChangeFAICandidate | - | - |
| Revoke Candidate | - | - | - |
| Abort | - | - | - |
| Complete | - | - | - |
| Copy | OpcenterEXDS_ProductionCoordination.PANEL_CopyFAIRecordWizard, OpcenterEXDS_ProductionCoordination.PANEL_CopyFAIRecordOnlyCompleteRelated | - | - |
| Download Report | OpcenterEXDS_ProductionCoordination.PANEL_ToastNotification | - | - |

### OpcenterEXDS_ProductionCoordination.ContainmentRequestMonitoring_Master

| Caption | Target Page | Target AppName | Target CommandName |
|---------|-------------|----------------|--------------------|
| Open | OpcenterEXDS_ProductionCoordination.ContainmentRequests_Details | - | - |
| Delete | - | - | - |
| Import | - | - | - |
| Refresh | - | - | - |

### OpcenterEXDS_ProductionCoordination.WorkOrderUpdate_Master

| Caption | Target Page | Target AppName | Target CommandName |
|---------|-------------|----------------|--------------------|
| Details | OpcenterEXDS_ProductionCoordination.WorkOrderUpdate_Details | - | - |

### OpcenterEXDS_ProductionCoordination.ChangeRequest_Master

| Caption | Target Page | Target AppName | Target CommandName |
|---------|-------------|----------------|--------------------|
| Create | OpcenterEXDS_ProductionCoordination.PANEL_CreateChangeNonConformance | - | - |
| Open | OpcenterEXDS_ProductionCoordination.ChangeRequest_Details | - | - |
| Accept | - | - | - |
| Reject  | OpcenterEXDS_ProductionCoordination.PANEL_UADMRejectChangeNonConformance | - | - |

### OpcenterEXDS_ProductionCoordination.LineSidePositions_Master

| Caption | Target Page | Target AppName | Target CommandName |
|---------|-------------|----------------|--------------------|
| Create | OpcenterEXDS_ProductionCoordination.PANEL_CreateLineSidePosition | - | - |
| Open | OpcenterEXDS_ProductionCoordination.LineSidePosition_Details | - | CallCommand_MF |
| Edit | OpcenterEXDS_ProductionCoordination.PANEL_UpdateLineSidePosition | - | - |
| Release | - | - | CallCommand_MF |
| Unset | - | - | CallCommand_MF |
| Delete | - | - | CallCommand_MF |

### OpcenterEXDS_ProductionCoordination.OfflineSession_Master

| Caption | Target Page | Target AppName | Target CommandName |
|---------|-------------|----------------|--------------------|
| Details | OpcenterEXDS_ProductionCoordination.OfflineSession_Details | - | - |
| Refresh | - | - | - |
| Check-In | - | - | - |
| Discard | - | - | - |

### OpcenterEXDS_ProductionCoordination.WorkOrderPreKitting_Master

| Caption | Target Page | Target AppName | Target CommandName |
|---------|-------------|----------------|--------------------|
| Details | OpcenterEXDS_ProductionCoordination.WorkOrderPreKitting_Details | - | - |

### OpcenterEXDS_ShopfloorExecution.Tool_Master

| Caption | Target Page | Target AppName | Target CommandName |
|---------|-------------|----------------|--------------------|
| Create | OpcenterEXDS_ShopfloorExecution.PANEL_CreateTool | - | - |
| Open | OpcenterEXDS_ShopfloorExecution.ToolDetails_Details | - | - |
| Edit | OpcenterEXDS_ShopfloorExecution.PANEL_UpdateTool | - | - |
| Tool Maintenance | OpcenterEXDS_ShopfloorExecution.PANEL_ToolMaintenance | - | - |
| Add Defect | OpcenterEXDS_ShopfloorExecution.PANEL_BrowseFailures | - | - |
| Delete Tool | - | - | CallCommand_MF |

### OpcenterEXDS_ProductionCoordination.Buffers_Master

| Caption | Target Page | Target AppName | Target CommandName |
|---------|-------------|----------------|--------------------|
| Create | OpcenterEXDS_ProductionCoordination.PANEL_CreateBuffer | - | - |
| Details | OpcenterEXDS_ProductionCoordination.Buffers_Details | - | - |
| Edit | OpcenterEXDS_ProductionCoordination.PANEL_UpdateBuffer | - | - |
| Change Buffer Status | OpcenterEXDS_ProductionCoordination.PANEL_ChangeBufferStatus | - | - |
| Delete | - | - | - |

### OpcenterEXDS_ProductionCoordination.ContainmentRequests_Master

| Caption | Target Page | Target AppName | Target CommandName |
|---------|-------------|----------------|--------------------|
| Create | OpcenterEXDS_ProductionCoordination.Panel_CreateContainmentRequest | - | - |
| Open | OpcenterEXDS_ProductionCoordination.ContainmentRequests_Details | - | - |
| Close | OpcenterEXDS_ProductionCoordination.PANEL_CloseContainmentRequest | - | - |
| Release | OpcenterEXDS_ProductionCoordination.PANEL_ManageContainmentRequestFullRelease | - | - |
| Scrap | OpcenterEXDS_ProductionCoordination.PANEL_BulkMaterialTrackingUnitsScrapFromContainmentRequest | - | - |
| Set As Ready | OpcenterEXDS_ProductionCoordination.PANEL_SetContainmentRequestAsReady | - | - |
| Refresh | - | - | - |

### OpcenterEXDS_ProductionCoordination.Measurements_Master

| Caption | Target Page | Target AppName | Target CommandName |
|---------|-------------|----------------|--------------------|
| Create | OpcenterEXDS_ProductionCoordination.PANEL_CreateMeasurements | - | - |
| Refresh | - | - | - |
| Download | - | - | - |

### OpcenterEXDS_ShopfloorExecution.NonConformance_Master

| Caption | Target Page | Target AppName | Target CommandName |
|---------|-------------|----------------|--------------------|
| Create | OpcenterEXDS_ShopfloorExecution.PANEL_DeclareNonConformance | - | - |
| Open | OpcenterEXDS_ShopfloorExecution.NonConformance_Details | - | CallCommand_MF |
| Edit | OpcenterEXDS_ShopfloorExecution.PANEL_UADMUpdateNonConformance | - | - |
| Change Status | OpcenterEXDS_ShopfloorExecution.PANEL_UADMSentenceNonConformanceV3_1 | - | - |

### OpcenterEXDS_ProductionCoordination.ProductionCoordinatorDashboard_Master

| Caption | Target Page | Target AppName | Target CommandName |
|---------|-------------|----------------|--------------------|
| Refresh Equipment | - | - | - |
| Details | OpcenterEXDS_ProductionCoordination.ProductionCoordinatorDashboard_Details | - | - |

### OpcenterEXDS_ProductionCoordination.ExecutionGroup_Master

| Caption | Target Page | Target AppName | Target CommandName |
|---------|-------------|----------------|--------------------|
| Create | OpcenterEXDS_ProductionCoordination.PANEL_CreateExecutionGroup | - | - |
| Details | OpcenterEXDS_ProductionCoordination.ExecutionGroup_Details | - | SUB_DeleteCommand_StartArray |
| Refresh | - | - | - |
| Edit | OpcenterEXDS_ProductionCoordination.PANEL_UpdateExecutionGroup, OpcenterEXDS_ProductionCoordination.PANEL_UpdateExecutionGroup_Details | - | - |
| Release | - | - | - |
| Ready For Scheduling | - | - | - |
| Change Status To Edit | - | - | - |
| Abort | - | - | - |
| Delete | - | - | SUB_DeleteCommand_StartArray |
| Get Ready Print Jobs | - | - | - |

### OpcenterEXDS_ProductionCoordination.Teams_Master

| Caption | Target Page | Target AppName | Target CommandName |
|---------|-------------|----------------|--------------------|
| Create | OpcenterEXDS_ProductionCoordination.PANEL_CreateTeam | - | - |
| Details | OpcenterEXDS_ProductionCoordination.Teams_Details | - | - |
| Edit | OpcenterEXDS_ProductionCoordination.PANEL_UpdateTeam | - | - |
| Delete | - | - | - |

### OpcenterEXDS_ProductionCoordination.WorkOrder_Master

| Caption | Target Page | Target AppName | Target CommandName |
|---------|-------------|----------------|--------------------|
| Open | OpcenterEXDS_ProductionCoordination.WorkOrder_Details | - | CallCommand_MF |
| Edit | OpcenterEXDS_ProductionCoordination.PANEL_EditWorkOrder | - | - |
| Merge | OpcenterEXDS_ProductionCoordination.PANEL_MergeWOHeaderWithProcess | - | - |
| Status To Edit | - | - | - |
| Ready for Scheduling | - | - | - |
| As Built | OpcenterEXDS_ProductionCoordination.WorkOrder_AsBuilt | - | - |
| Print Label | - | - | - |
| Release | - | - | - |
| Split | OpcenterEXDS_ProductionCoordination.PANEL_SplitWorkOrder | - | - |
| Complete | - | - | - |
| Abort | - | - | - |
| Close Flexible | OpcenterEXDS_ProductionCoordination.PANEL_CloseFlexibleWorkOrder | - | - |
| Delete | - | - | - |
| To Be Cleaned | - | - | - |

### OpcenterEXDS_ProductionCoordination.ChangePackages_Master

| Caption | Target Page | Target AppName | Target CommandName |
|---------|-------------|----------------|--------------------|
| Refresh | - | - | - |
| Details | OpcenterEXDS_ProductionCoordination.ChangePackages_WorkOrderOverview_Details | - | SUB_DeleteCommand_StartArray |
| Accept | OpcenterEXDS_ProductionCoordination.PANEL_AcceptChangePackage | - | - |
| Reject  | OpcenterEXDS_ProductionCoordination.PANEL_RejectChangePackage | - | - |

### OpcenterEXDS_ProductionCoordination.PCTimeUpdate_Master

| Caption | Target Page | Target AppName | Target CommandName |
|---------|-------------|----------------|--------------------|
| Details | OpcenterEXDS_ProductionCoordination.PCTimeUpdate_Details | - | - |

### OpcenterEXDS_ProductionCoordination.HandlingUnits_Master

| Caption | Target Page | Target AppName | Target CommandName |
|---------|-------------|----------------|--------------------|
| - | OpcenterEXDS_ProductionCoordination.PANEL_CreateHandlingUnit | - | - |
| - | OpcenterEXDS_ProductionCoordination.HandlingUnit_Details | - | - |

### OpcenterEXDS_ProductionCoordination.WorkOrderUpdateCheck_Master

| Caption | Target Page | Target AppName | Target CommandName |
|---------|-------------|----------------|--------------------|
| Refresh | - | - | - |
| Update | - | - | - |
| Request New BOP | - | - | - |

### OpcenterEXDS_ProductionCoordination.Users_Master

| Caption | Target Page | Target AppName | Target CommandName |
|---------|-------------|----------------|--------------------|
| Details | OpcenterEXDS_ProductionCoordination.Users_Details | - | - |

### OpcenterEXDS_ProductionCoordination.ERPOrder_Master

| Caption | Target Page | Target AppName | Target CommandName |
|---------|-------------|----------------|--------------------|
| Add | OpcenterEXDS_ProductionCoordination.PANEL_CreateERPOrder | - | - |
| Reload Data | - | - | - |
| Details | OpcenterEXDS_ProductionCoordination.ERPOrder_Details | - | SUB_DeleteCommand |
| Edit | OpcenterEXDS_ProductionCoordination.PANEL_UpdateERPOrder | - | - |
| Schedule | - | - | - |
| Schedule | - | - | - |
| Unschedule | - | - | - |
| Unschedule | - | - | - |
| Delete | - | - | SUB_DeleteCommand |

### OpcenterEXDS_Configuration.Hold_Master

| Caption | Target Page | Target AppName | Target CommandName |
|---------|-------------|----------------|--------------------|
| Close | OpcenterEXDS_Configuration.PANEL_UADMRemoveHoldList | - | - |
| Details | OpcenterEXDS_Configuration.Hold_Details | - | - |

### OpcenterEXDS_ProductionCoordination.IntegrationEvents_Master

| Caption | Target Page | Target AppName | Target CommandName |
|---------|-------------|----------------|--------------------|
| Resend Failed Event | - | - | - |

### OpcenterEXDS_ProductionCoordination.ScanMaterialTrackingUnits_Master

| Caption | Target Page | Target AppName | Target CommandName |
|---------|-------------|----------------|--------------------|
| Scrap | OpcenterEXDS_ProductionCoordination.Panel_MaterialTrackingUnit_Scrap | - | - |
| Release | OpcenterEXDS_ProductionCoordination.Panel_MaterialTrackingUnit_Release | - | - |
| Scrap By File | OpcenterEXDS_ProductionCoordination.Panel_MaterialTrackingUnit_ScrapByFile | - | - |
| Release By File | OpcenterEXDS_ProductionCoordination.Panel_MaterialTrackingUnit_ReleaseByFile | - | - |

---

## 6. Pages/Panels Commands Hierarchy

Microflows and nanoflows called by each page/panel, showing recursive call hierarchy up to 5 levels (in YAML structure). Microflows called transitively are loaded on-the-fly from the database when needed.

```yaml
pages:
  - name: PANEL_CreateHandlingUnit
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateHandlingUnit
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateHandlingUnit
    target_commands:
      - AppU4DM.CreateHandlingUnit

  - name: HandlingUnits_Master
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_UnLinkMaterialItemsFromHandlingUnit
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UnlinkMaterialItemsToHandlingUnit
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UnLinkMaterialItemsFromHandlingUnit
    target_commands:
      - AppU4DM.UnLinkMaterialItemsFromHandlingUnit

  - name: PANEL_UnLoadHandlingUnit
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UnLoadHandlingUnit
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UnLoadHandlingUnit
    target_commands:
      - AppU4DM.UnLoadHandlingUnit

  - name: PANEL_UADMLinkMaterialItemsToHandlingUnit
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UADMLinkMaterialItemsToHandlingUnit
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMLinkMaterialItemsToHandlingUnit
    target_commands:
      - AppU4DM.UADMLinkMaterialItemsToHandlingUnit

  - name: PANEL_ChangeHandlingUnitStatus
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_ChangeHandlingUnitStatus
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.ChangeHandlingUnitStatus
    target_commands:
      - AppU4DM.ChangeHandlingUnitStatus

  - name: HandlingUnit_Details
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: LogisticRequest_Master
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_RejectMaterialRequest
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_RejectMaterialRequest
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.RejectMaterialRequest
    target_commands:
      - AppU4DM.RejectMaterialRequest

  - name: PANEL_AddBuffer
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_AcceptMaterialRequest
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_AcceptMaterialRequest
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.AcceptMaterialRequest
    target_commands:
      - AppU4DM.AcceptMaterialRequest

  - name: PANEL_UpdateERPOrder
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UpdateERPOrder
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_UpdateERPOrder_ERPOrderBoP_CreateBoPERPOrder
            - name: OpcenterEXDS_ProductionCoordination_Connector.UpdateERPOrder
            - name: OpcenterEXDS_ProductionCoordination_Connector.ERPOrderBoP_SetBoPInfoToERPOrder
    target_commands:
      - ERPOrder.UpdateERPOrder
      - ERPOrder.ERPOrderBoP_SetBoPInfoToERPOrder

  - name: PANEL_CreateERPOrder
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateERPOrder
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_CreateERPOrder_ERPOrderBoP_CreateBoPERPOrder
            - name: OpcenterEXDS_ProductionCoordination_Connector.CreateERPOrder
            - name: OpcenterEXDS_ProductionCoordination_Connector.ERPOrderBoP_CreateBoPERPOrder
    target_commands:
      - ERPOrder.CreateERPOrder
      - ERPOrder.ERPOrderBoP_CreateBoPERPOrder

  - name: ERPOrder_Master
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_SetERPOrderListStatus
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.SetERPOrderListStatus
      - name: OpcenterEXDS_ProductionCoordination.ACT_SetERPOrderListStatuses
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.SetERPOrderListStatus
      - name: OpcenterEXDS_ProductionCoordination.ACT_DeleteERPOrder
        calls:
          - name: OpcenterEXFN_MasterData_Connector.SUB_DeleteCommand
    target_commands:
      - ERPOrder.SetERPOrderListStatus
      - AppName.CommandName

  - name: MaterialSuppliers_Details
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: ERPOrder_Details
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_DeleteERPOrder
        calls:
          - name: OpcenterEXFN_MasterData_Connector.SUB_DeleteCommand
    target_commands:
      - AppName.CommandName

  - name: Measurements_Master
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_CreateMeasurements
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_TraceProcessMeasurements
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.TraceProcessMeasurements
    target_commands:
      - AppU4DM.TraceProcessMeasurements

  - name: PANEL_UADMSetWorkOrderOperationFutureHoldList
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UADMSetWorkOrderOperationFutureHoldList
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMSetWorkOrderOperationFutureHoldList
    target_commands:
      - AppU4DM.UADMSetWorkOrderOperationFutureHoldList

  - name: PANEL_UADMRemoveWorkOrderOperationFutureHoldList
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UADMRemoveWorkOrderOperationFutureHoldList
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMRemoveWorkOrderOperationFutureHoldList
    target_commands:
      - AppU4DM.UADMRemoveWorkOrderOperationFutureHoldList

  - name: FutureHoldManagement_Master
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: FutureHoldManagement_Details
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_UADMLinkMaterialItemsToBuffer
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UADMLinkMaterialItemsToBuffer
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMLinkMaterialItemsToBuffer
    target_commands:
      - AppU4DM.UADMLinkMaterialItemsToBuffer

  - name: Buffers_Details
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UnLinkMaterialItemFromBuffer
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UnLinkMaterialItemFromBuffer
      - name: OpcenterEXDS_ProductionCoordination.ACT_DeleteBuffer
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.DeleteBuffer
    target_commands:
      - AppU4DM.UnLinkMaterialItemFromBuffer
      - AppU4DM.DeleteBuffer

  - name: PANEL_ChangeBufferStatus
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_ChangeBufferStatus
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.ChangeBufferStatus
    target_commands:
      - AppU4DM.ChangeBufferStatus

  - name: PANEL_CreateBuffer
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateBuffer
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateBuffer
    target_commands:
      - AppU4DM.CreateBuffer

  - name: PANEL_UpdateBuffer
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UpdateBuffer
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UpdateBuffer
    target_commands:
      - AppU4DM.UpdateBuffer

  - name: Buffers_Master
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_DeleteBuffer
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.DeleteBuffer
    target_commands:
      - AppU4DM.DeleteBuffer

  - name: WorkOrderUpdate_Master
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: Process_Details
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UADMUpdateWOByBoPList_WithConfirmation_ShowPage
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMUpdateWOByBoPList
    target_commands:
      - AppU4DM.UADMUpdateWOByBoPList

  - name: WorkOrderUpdate_Details
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UADMUpdateWOByBoPList_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMUpdateWOByBoPList
    target_commands:
      - AppU4DM.UADMUpdateWOByBoPList

  - name: PANEL_SplitWorkOrder
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_GenerateSplitWorkOrderId
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_GenerateSplitWorkOrderId_MF
            - name: OpcenterEXDS_ProductionCoordination_Connector.AutoGenerateWorkOrderNId
      - name: OpcenterEXDS_ProductionCoordination.ACT_SplitWorkOrder
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.SplitWorkOrder
    target_commands:
      - AppU4DM.AutoGenerateWorkOrderNId
      - AppU4DM.UADMSplitWorkOrder

  - name: PANEL_AssociateSerialNumbersToSplitWo
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_CreateWOOFolderAndWOOpDependencies
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateWOOFolderAndWOOpDependencies
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateWOOFolderAndWOOpDependencies
    target_commands:
      - AppU4DM.CreateWOOFolderAndWOOpDependencies

  - name: PANEL_CloseFlexibleWorkOrder
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CloseFlexibleWorkOrder
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMAbruptlyCloseFlexibleWorkOrder
    target_commands:
      - AppU4DM.UADMAbruptlyCloseFlexibleWorkOrder

  - name: PANEL_CreateWorkOrder_FromMasterPlan_Step1
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_GenerateWorkOrderId
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_GenerateWorkOrderId_MF
            - name: OpcenterEXDS_ProductionCoordination_Connector.AutoGenerateWorkOrderNId
      - name: OpcenterEXDS_ProductionCoordination.ACT_GenerateBatchId
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_GenerateBatchId_MF
            - name: OpcenterEXDS_ProductionCoordination_Connector.AutoGenerateMTUCode
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateWorkOrderFromMasterPlan_Step1
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateWorkOrderFromMasterPlanBOMResolution
    target_commands:
      - AppU4DM.AutoGenerateWorkOrderNId
      - AppU4DM.AutoGenerateMTUCode
      - AppU4DM.PLMCreateWorkOrderFromMasterPlanBOMResolution

  - name: PANEL_CreateWorkOrder_FromMasterPlan_Step2
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateWorkOrderFromMasterPlan_Step2
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateWorkOrderFromMasterPlanBOMResolution
    target_commands:
      - AppU4DM.PLMCreateWorkOrderFromMasterPlanBOMResolution

  - name: PANEL_CreateWorkOrder_FromMasterPlanWithQC_Step1
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateWorkOrderFromMasterPlanWithQC_Step1
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateWorkOrderFromMasterPlanWithQC
      - name: OpcenterEXDS_ProductionCoordination.ACT_GenerateWorkOrderId
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_GenerateWorkOrderId_MF
            - name: OpcenterEXDS_ProductionCoordination_Connector.AutoGenerateWorkOrderNId
      - name: OpcenterEXDS_ProductionCoordination.ACT_GenerateBatchId
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_GenerateBatchId_MF
            - name: OpcenterEXDS_ProductionCoordination_Connector.AutoGenerateMTUCode
    target_commands:
      - AppU4DM.CreateWorkOrderFromMasterPlanWithQC
      - AppU4DM.AutoGenerateWorkOrderNId
      - AppU4DM.AutoGenerateMTUCode

  - name: PANEL_CreateWorkOrder_FromMasterPlanWithQC_Step2
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_CreateWorkOrder_FromProcess
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateWorkOrderFromProcess
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateWorkOrderFromProcess
      - name: OpcenterEXDS_ProductionCoordination.ACT_GenerateWorkOrderId
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_GenerateWorkOrderId_MF
            - name: OpcenterEXDS_ProductionCoordination_Connector.AutoGenerateWorkOrderNId
      - name: OpcenterEXDS_ProductionCoordination.ACT_GenerateBatchId
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_GenerateBatchId_MF
            - name: OpcenterEXDS_ProductionCoordination_Connector.AutoGenerateMTUCode
    target_commands:
      - AppU4DM.UADMCreateWorkOrderFromProcess
      - AppU4DM.AutoGenerateWorkOrderNId
      - AppU4DM.AutoGenerateMTUCode

  - name: PANEL_CreateWorkOrder_Manually
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateWorkOrderManually
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateWorkOrderManually
    target_commands:
      - AppU4DM.CreateWorkOrder

  - name: PANEL_CreateWorkOrder_AsPlanned
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateWorkOrderAsPlanned
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMCreateWorkOrdersFromAsPlannedBOP
    target_commands:
      - AppU4DM.UADMCreateWorkOrdersFromAsPlannedBOP

  - name: PANEL_CreateWorkOrder_Header
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateWorkOrderHeader
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateWorkOrderHeader
      - name: OpcenterEXDS_ProductionCoordination.ACT_GenerateWorkOrderId
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_GenerateWorkOrderId_MF
            - name: OpcenterEXDS_ProductionCoordination_Connector.AutoGenerateWorkOrderNId
      - name: OpcenterEXDS_ProductionCoordination.ACT_GenerateBatchId
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_GenerateBatchId_MF
            - name: OpcenterEXDS_ProductionCoordination_Connector.AutoGenerateMTUCode
    target_commands:
      - AppU4DM.CreateWorkOrderHeader
      - AppU4DM.AutoGenerateWorkOrderNId
      - AppU4DM.AutoGenerateMTUCode

  - name: PANEL_CreateWorkOrder_FromMasterPlanWithEffectivity
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_GenerateWorkOrderId
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_GenerateWorkOrderId_MF
            - name: OpcenterEXDS_ProductionCoordination_Connector.AutoGenerateWorkOrderNId
      - name: OpcenterEXDS_ProductionCoordination.ACT_GenerateBatchId
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_GenerateBatchId_MF
            - name: OpcenterEXDS_ProductionCoordination_Connector.AutoGenerateMTUCode
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateWorkOrder_FromMasterPlanWithEffectivity
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateWorkOrderFromMasterPlanWithEffectivity
    target_commands:
      - AppU4DM.AutoGenerateWorkOrderNId
      - AppU4DM.AutoGenerateMTUCode
      - AppU4DM.CreateWorkOrderFromMasterPlanWithEffectivity

  - name: PANEL_MergeWOHeaderWithProcess
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_MergeWOHeaderWithProcess
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.MergeWOHeaderWithProcess
    target_commands:
      - AppU4DM.MergeWOHeaderWithProcess

  - name: PANEL_EditWorkOrder
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_EditWorkOrder
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.EditWorkOrder
    target_commands:
      - AppU4DM.EditWorkOrder

  - name: WorkOrder_Master
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_ChangeWorkOrderStatusToEdit_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.ChangeWorkOrderStatusToEdit
      - name: OpcenterEXDS_ProductionCoordination.ACT_SetWorkOrderForScheduling_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMSetWorkOrderForScheduling
      - name: OpcenterEXDS_ProductionCoordination.ACT_TriggerPrintingOnWorkOrder_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.TriggerPrintingOnWorkOrder
      - name: OpcenterEXDS_ProductionCoordination.ACT_ReleaseWorkOrder_List_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_ReleaseWorkOrder_List_WithConfirmation_MF
            - name: OpcenterEXDS_ProductionCoordination_Connector.UADMReleaseWorkOrderList
      - name: OpcenterEXDS_ProductionCoordination.ACT_SetTargetQuantityOnFlexibleWorkOrder_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.SetTargetQuantityOnFlexibleWorkOrder
      - name: OpcenterEXDS_ProductionCoordination.ACT_AbortWorkOrder_List_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_AbortWorkOrder_List_WithConfirmation_MF
            - name: OpcenterEXDS_ProductionCoordination_Connector.UADMAbortWorkOrder
      - name: OpcenterEXDS_ProductionCoordination.ACT_DeleteWorkOrder_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.DeleteWorkOrder
      - name: OpcenterEXDS_ProductionCoordination.ACT_MarkForCleaningWorkOrder_List_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_MarkForCleaningWorkOrder_List_WithConfirmation_MF
            - name: OpcenterEXDS_ProductionCoordination_Connector.MarkForCleaningWorkOrderList
      - name: OpcenterEXDS_ProductionCoordination.ACT_SendMessageWorkOrder_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateWorkOrderOutMsg
      - name: OpcenterEXDS_ProductionCoordination.ACT_SendMessageWorkOrderAsBuilt
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateWorkOrderOutMsg
    target_commands:
      - AppU4DM.ChangeWorkOrderStatusToEdit
      - AppU4DM.UADMSetWorkOrderForScheduling
      - AppU4DM.TriggerPrintingOnWorkOrder
      - AppU4DM.UADMReleaseWorkOrder
      - AppU4DM.SetTargetQuantityOnFlexibleWorkOrder
      - AppU4DM.UADMAbortWorkOrder
      - AppU4DM.DeleteWorkOrder
      - AppU4DM.MarkForCleaningWorkOrderList
      - AppU4DM.CreateWorkOrderOutMsg

  - name: PANEL_AssociateSerialNumbers
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_GenerateProducedMaterialItems
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMCreateAndAssignProducedMaterialItems
      - name: OpcenterEXDS_ProductionCoordination.ACT_AssignProducedMaterialItems
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.AssignProducedMaterialItems
    target_commands:
      - AppU4DM.UADMCreateAndAssignProducedMaterialItems
      - AppU4DM.AssignProducedMaterialItems

  - name: PANEL_EditAlternativeOperationGroup
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UpdatePreferredWOOpDependencyNavigation
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_UpdatePreferredWOOpDependencyNavigation_MF
            - name: OpcenterEXDS_ProductionCoordination_Connector.UpdatePreferredWOOpDependencyNavigation
    target_commands:
      - AppU4DM.UpdatePreferredWOOpDependencyNavigation

  - name: PANEL_AddAlternativeOperationGroup
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateWOOpDependencyNavigationList
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_CreateWOOpDependencyNavigationListCommand_MF
            - name: OpcenterEXDS_ProductionCoordination_Connector.CreateWOOpDependencyNavigationList
    target_commands:
      - AppU4DM.CreateWOOpDependencyNavigationList

  - name: PANEL_LinkWIDefinitionsToWOOperation
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_LinkWIDefinitionsToWOOperation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.LinkWIDefinitionsToWOOperation
    target_commands:
      - AppU4DM.LinkWIDefinitionsToWOOperation

  - name: PANEL_LinkWIDefinitionsToWOStep
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_LinkWIDefinitionsToWOStep
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.LinkWIDefinitionsToWOStep
    target_commands:
      - AppU4DM.LinkWIDefinitionsToWOStep

  - name: PANEL_EditWorkOrderHumanResource_Step
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_EditWorkOrderHumanResource_Command_Step
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.EditWorkOrderHumanResource
    target_commands:
      - AppU4DM.EditWorkOrderHumanResource

  - name: PANEL_LinkWorkOrderHumanResource_Step
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_LinkWorkOrderHumanResource_Step
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_LinkWorkOrderHumanResource_Command
            - name: OpcenterEXDS_ProductionCoordination_Connector.LinkWorkOrderHumanResource
    target_commands:
      - AppU4DM.LinkWorkOrderHumanResource

  - name: PANEL_CreateWOStepDependency_FromOpDetails
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateWOStepDependency_FromOpDetails
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateWOStepDependency
    target_commands:
      - AppU4DM.CreateWOStepDependency

  - name: PANEL_CreateWOStepDependency
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateWOStepDependency
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateWOStepDependency
    target_commands:
      - AppU4DM.CreateWOStepDependency

  - name: PANEL_LinkSkillsToWorkOrderStep
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_LinkSkillsToWorkOrderStep
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_LinkSkillsToWorkOrderStep_Command
            - name: OpcenterEXDS_ProductionCoordination_Connector.LinkSkillsToWorkOrderStep
    target_commands:
      - AppU4DM.LinkSkillsToWorkOrderStep

  - name: PANEL_CreateToBeConsumedMaterials_Step
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateToBeConsumedMaterials_Step
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateToBeConsumedMaterials
    target_commands:
      - AppU4DM.CreateToBeConsumedMaterials

  - name: PANEL_CreateRuntimeCharacteristicRepresentationContainer_Step
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateRuntimeCharacteristicRepresentationContainer
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateRuntimeCharacteristicRepresentationContainer
    target_commands:
      - AppU4DM.CreateRuntimeCharacteristicRepresentationContainer

  - name: PANEL_EditWorkOrderStep
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_EditWorkOrderStep
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UpdateWorkOrderStep
    target_commands:
      - AppU4DM.UpdateWorkOrderStep

  - name: PANEL_LinkStepToOperation
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_LinkStepToOperation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateWorkOrderStepFromStep
    target_commands:
      - AppU4DM.CreateWorkOrderStepFromStep

  - name: PANEL_CreateWorkOrderStepsManually
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateWorkOrderSteps
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateWorkOrderSteps
    target_commands:
      - AppU4DM.CreateWorkOrderSteps

  - name: PANEL_AddTool_Step
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateToBeUsedTools_Step
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateToBeUsedTools
    target_commands:
      - AppU4DM.CreateToBeUsedTools

  - name: PANEL_EditInterlockingChecks_Step
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UpdateItlkCheckWOStepAssociationList
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UpdateItlkCheckWOStepAssociationList
    target_commands:
      - AppU4DM.UpdateItlkCheckWOStepAssociationList

  - name: PANEL_LinkItlkCheckToWorkOrderStep
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_LinkItlkCheckToWorkOrderStepList
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.LinkItlkCheckToWorkOrderStepList
    target_commands:
      - AppU4DM.LinkItlkCheckToWorkOrderStepList

  - name: WorkOrderOperationSteps_Details
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_DeleteWorkOStepDependency_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.DeleteWorkOStepDependency
      - name: OpcenterEXDS_ProductionCoordination.ACT_DeleteToBeConsumedMaterial_Step_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.DeleteToBeConsumedMaterial
      - name: OpcenterEXDS_ProductionCoordination.ACT_DeleteToBeUsedTool_WithConfirmation
      - name: OpcenterEXDS_ProductionCoordination.ACT_UnlinkWorkInstructionsFromWOStep
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UnlinkWorkInstructionsFromWOStep
      - name: OpcenterEXDS_ProductionCoordination.ACT_DeleteToBeUsedInspection
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.DeleteToBeUsedInspection
      - name: OpcenterEXDS_ProductionCoordination.ACT_UnlinkWorkOrderHumanResource_Step
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UnlinkWorkOrderHumanResource
      - name: OpcenterEXDS_ProductionCoordination.ACT_UnlinkSkillsFromWOStep
      - name: OpcenterEXDS_ProductionCoordination.ACT_UnlinkItlkCheckToWorkOrderStepList
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UnlinkItlkCheckToWorkOrderStepList
      - name: OpcenterEXDS_ProductionCoordination.ACT_TriggerPrintingOnWorkOrderStep
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.TriggerPrintingOnWorkOrderStep
    target_commands:
      - AppU4DM.DeleteWorkOStepDependency
      - AppU4DM.DeleteToBeConsumedMaterial
      - AppU4DM.DeleteToBeUsedTool
      - AppU4DM.UnlinkWorkInstructionsFromWOStep
      - AppU4DM.DeleteToBeUsedInspection
      - AppU4DM.UnlinkWorkOrderHumanResource
      - AppU4DM.UnlinkSkillsToWorkOrderStep
      - AppU4DM.UnlinkItlkCheckToWorkOrderStepList
      - AppU4DM.TriggerPrintingOnWorkOrderStep

  - name: PANEL_LinkSkillsToWorkOrderOperation
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_LinkSkillsToWorkOrderOperation
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_LinkSkillsToWorkOrderOperation_Command
            - name: OpcenterEXDS_ProductionCoordination_Connector.LinkSkillsToWorkOrderOperation
    target_commands:
      - AppU4DM.LinkSkillsToWorkOrderOperation

  - name: PANEL_AddTool
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateToBeUsedTools
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateToBeUsedTools
    target_commands:
      - AppU4DM.CreateToBeUsedTools

  - name: PANEL_ShowScrewing
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_LinkUserToWorkOrderOperationList
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_LinkUserToWorkOrderOperationList
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_LinkUserToWorkOrderOperationList_Command
            - name: OpcenterEXDS_ProductionCoordination_Connector.LinkUserToWorkOrderOperationList
      - name: OpcenterEXDS_ProductionCoordination.ACT_GetUserDetails
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_GetUserDetails_MF
            - name: OpcenterEXDS_ProductionCoordination_Connector.GetUserDetailsList
    target_commands:
      - AppU4DM.LinkUserToWorkOrderOperationList
      - AppU4DM.GetUserDetailsList

  - name: PANEL_EditInterlockingChecks
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UpdateItlkCheckWOOpAssociationList
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UpdateItlkCheckWOOpAssociationList
    target_commands:
      - AppU4DM.UpdateItlkCheckWOOpAssociationList

  - name: PANEL_DetailsInterlockingChecks
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_LinkItlkCheckToWorkOrderOperationList
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_LinkItlkCheckToWorkOrderOperationList
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.LinkItlkCheckToWorkOrderOperationList
    target_commands:
      - AppU4DM.LinkItlkCheckToWorkOrderOperationList

  - name: PANEL_CreateToBeConsumedMaterials
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateToBeConsumedMaterials
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateToBeConsumedMaterials
    target_commands:
      - AppU4DM.CreateToBeConsumedMaterials

  - name: PANEL_CreateToBeUsedMachine_PartProgram
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateToBeUsedMachine
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateToBeUsedMachine
    target_commands:
      - AppU4DM.CreateToBeUsedMachine

  - name: PANEL_CreateToBeUsedMachine
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateToBeUsedMachine
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateToBeUsedMachine
    target_commands:
      - AppU4DM.CreateToBeUsedMachine

  - name: PANEL_LinkToBeUsedMachineToPartProgram
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_LinkToBeUsedMachineToPartProgram
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.LinkToBeUsedMachineToPartProgram
    target_commands:
      - AppU4DM.LinkToBeUsedMachineToPartProgram

  - name: PANEL_EditWorkOrderHumanResource
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_EditWorkOrderHumanResource_Command
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.EditWorkOrderHumanResource
    target_commands:
      - AppU4DM.EditWorkOrderHumanResource

  - name: PANEL_LinkWorkOrderHumanResource
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_LinkWorkOrderHumanResource
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_LinkWorkOrderHumanResource_Command
            - name: OpcenterEXDS_ProductionCoordination_Connector.LinkWorkOrderHumanResource
    target_commands:
      - AppU4DM.LinkWorkOrderHumanResource

  - name: PANEL_CreateWorkOrderOperation_FromProcess
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateWorkOrderOperation_FromProcess
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMCreateWorkOrderOperationFromProcessOperation
    target_commands:
      - AppU4DM.UADMCreateWorkOrderOperationFromProcessOperation

  - name: PANEL_CreateWorkOrderOperation_FromMasterPlan
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateWorkOrderOperation_FromMasterPlan
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateWorkOrderOperationFromMPProcessOperation
    target_commands:
      - AppU4DM.CreateWorkOrderOperationFromMPProcessOperation

  - name: PANEL_CreateWorkOrderOperation_Manually
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateWorkOrderOperationManually
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateWorkOrderOperationManually
    target_commands:
      - AppU4DM.CreateWorkOrderOperation

  - name: PANEL_EditWorkOrderOperation
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_EditWorkOrderOperation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.EditWorkOrderOperation
    target_commands:
      - AppU4DM.EditWorkOrderOperation

  - name: PANEL_CharacteristicRepresentation
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_CreateRuntimeCharacteristicRepresentationContainer
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateRuntimeCharacteristicRepresentationContainer
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateRuntimeCharacteristicRepresentationContainer
    target_commands:
      - AppU4DM.CreateRuntimeCharacteristicRepresentationContainer

  - name: WorkOrderOperation_Details
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_DeleteWOOFolderAndWOOpDependencies_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.DeleteWOOFolderAndWOOpDependencies
      - name: OpcenterEXDS_ProductionCoordination.ACT_UnLinkToBeUsedMachineToPartProgram
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UnLinkToBeUsedMachineToPartProgram
      - name: OpcenterEXDS_ProductionCoordination.ACT_DeleteToBeUsedMachine_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.DeleteToBeUsedMachine
      - name: OpcenterEXDS_ProductionCoordination.ACT_SetPreferredMachine_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.SetPreferredMachine
      - name: OpcenterEXDS_ProductionCoordination.ACT_DeleteToBeConsumedMaterial_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.DeleteToBeConsumedMaterial
      - name: OpcenterEXDS_ProductionCoordination.ACT_DeleteToBeUsedTool_WithConfirmation
      - name: OpcenterEXDS_ProductionCoordination.ACT_UnlinkWorkInstructionsFromWOOperation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UnlinkWorkInstructionsFromWOOperation
      - name: OpcenterEXDS_ProductionCoordination.ACT_UnlinkSkillsFromWOOperation
      - name: OpcenterEXDS_ProductionCoordination.ACT_DeleteWorkOrderStep_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.DeleteWorkOrderStep
      - name: OpcenterEXDS_ProductionCoordination.ACT_TriggerPrintingOnWorkOrderStep
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.TriggerPrintingOnWorkOrderStep
      - name: OpcenterEXDS_ProductionCoordination.ACT_DeleteToBeUsedInspection
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.DeleteToBeUsedInspection
      - name: OpcenterEXDS_ProductionCoordination.ACT_UnlinkWorkOrderHumanResource
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UnlinkWorkOrderHumanResource
      - name: OpcenterEXDS_ProductionCoordination.ACT_UnlinkUserToWorkOrderOperationList
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_UnlinkUserToWorkOrderOperationList_Command
            - name: OpcenterEXDS_ProductionCoordination_Connector.UnlinkUserToWorkOrderOperationList
      - name: OpcenterEXDS_ProductionCoordination.ACT_UnlinkItlkCheckToWorkOrderOperationList
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UnlinkItlkCheckToWorkOrderOperationList
      - name: OpcenterEXDS_ProductionCoordination.ACT_TriggerPrintingOnWorkOrderOperation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.TriggerPrintingOnWorkOrderOperation
      - name: OpcenterEXDS_ProductionCoordination.ACT_SendMessageWorkOrderOperation_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateWorkOrderOutMsg
      - name: OpcenterEXDS_ProductionCoordination.ACT_SendMessageWorkOrderOperationAsBuilt_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateWorkOrderOutMsg
    target_commands:
      - AppU4DM.DeleteWOOFolderAndWOOpDependencies
      - AppU4DM.UnLinkToBeUsedMachineToPartProgram
      - AppU4DM.DeleteToBeUsedMachine
      - AppU4DM.SetPreferredMachine
      - AppU4DM.DeleteToBeConsumedMaterial
      - AppU4DM.DeleteToBeUsedTool
      - AppU4DM.UnlinkWorkInstructionsFromWOOperation
      - AppU4DM.UnlinkSkillsFromWOOperation
      - AppU4DM.DeleteWorkOrderStep
      - AppU4DM.TriggerPrintingOnWorkOrderStep
      - AppU4DM.DeleteToBeUsedInspection
      - AppU4DM.UnlinkWorkOrderHumanResource
      - AppU4DM.UnlinkUserToWorkOrderOperationList
      - AppU4DM.UnlinkItlkCheckToWorkOrderOperationList
      - AppU4DM.TriggerPrintingOnWorkOrderOperation
      - AppU4DM.CreateWorkOrderOutMsg

  - name: ExecutionGroup_Details_2
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UnlinkWOOperationsFromExecutionGroup
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UnlinkWOOperationsFromExecutionGroup
      - name: OpcenterEXDS_ProductionCoordination.ACT_ReleaseExecGroup_Details
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.ReleaseExecutionGroup
      - name: OpcenterEXDS_ProductionCoordination.ACT_SetReadyExecGroup_Details
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.SetReadyForSchedulingExecutionGroup
      - name: OpcenterEXDS_ProductionCoordination.ACT_SetEditExecGroup_Details
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.SetExecGroupListInEditStatus
      - name: OpcenterEXDS_ProductionCoordination.ACT_AbortExecGroup_Details
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.AbortExecutionGroup
      - name: OpcenterEXDS_ProductionCoordination.DeleteExecutionGroup
        calls:
          - name: OpcenterEXFN_MasterData_Connector.SUB_DeleteCommand_StartArray
    target_commands:
      - AppU4DM.UnlinkWOOperationsFromExecutionGroup
      - AppU4DM.ReleaseExecutionGroup
      - AppU4DM.ScheduleExecutionGroup
      - AppU4DM.SetExecGroupListInEditStatus
      - AppU4DM.AbortExecutionGroupList
      - AppName.CommandName

  - name: ExecutionGroupPhase_Details_2
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UnlinkWIToEGPhase
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UnlinkWIDefinitionsFromEGPhase
      - name: OpcenterEXDS_ProductionCoordination.ACT_DeletePowder
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.RemoveAMPowderLoadedCheckFromEgPhaseList
    target_commands:
      - AppU4DM.UnlinkWIDefinitionsFromEGPhase
      - PowderMgt.RemoveAMPowderLoadedCheckFromEgPhaseList

  - name: PANEL_EditWorkOrderFolder
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_EditWorkOrderFolder
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.EditWOOFolder
    target_commands:
      - AppU4DM.EditWOOFolder

  - name: PANEL_CreateSubFolder
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateWorkOrderFolder
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateWOOFolder
    target_commands:
      - AppU4DM.CreateWOOFolder

  - name: PANEL_LinkWOOFoldersToWOOFolder
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_LinkWOOFoldersToWOOFolder
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.LinkWOOFoldersToWOOFolder
    target_commands:
      - AppU4DM.LinkWOOFoldersToWOOFolder

  - name: PANEL_CreateWorkOrderFolder
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateWorkOrderFolder
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateWOOFolder
    target_commands:
      - AppU4DM.CreateWOOFolder

  - name: PANEL_LinkWOOperationsToWOOFolder
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_LinkWOOperationsToWOOFolder
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.LinkWOOperationsToWOOFolder
    target_commands:
      - AppU4DM.LinkWOOperationsToWOOFolder

  - name: OperationFolders_Details
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UnlinkWOOperationsFromWOOFolder
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UnlinkWOOperationsFromWOOFolder
      - name: OpcenterEXDS_ProductionCoordination.ACT_DeleteWorkOrderOperation_WithConfirmation_FromList
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_DeleteWorkOrderOperation_WithConfirmation
            - name: OpcenterEXDS_ProductionCoordination_Connector.DeleteWorkOrderOperation
      - name: OpcenterEXDS_ProductionCoordination.ACT_DeleteWOOFolderAndWOOpDependencies_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.DeleteWOOFolderAndWOOpDependencies
      - name: OpcenterEXDS_ProductionCoordination.ACT_UnlinkSubFolders
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UnlinkWOOFoldersFromWOOFolder
      - name: OpcenterEXDS_ProductionCoordination.ACT_DeleteWOOFolders
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.DeleteWOOFolders
    target_commands:
      - AppU4DM.UnlinkWOOperationsFromWOOFolder
      - AppU4DM.DeleteWorkOrderOperation
      - AppU4DM.DeleteWOOFolderAndWOOpDependencies
      - AppU4DM.UnlinkWOOFoldersFromWOOFolder
      - AppU4DM.DeleteWOOFolders

  - name: WorkOrder_Details
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_DeleteWorkOrderOperation_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.DeleteWorkOrderOperation
      - name: OpcenterEXDS_ProductionCoordination.ACT_TriggerPrintingOnWorkOrderOperation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.TriggerPrintingOnWorkOrderOperation
      - name: OpcenterEXDS_ProductionCoordination.ACT_SendMessageWorkOrderOperation_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateWorkOrderOutMsg
      - name: OpcenterEXDS_ProductionCoordination.ACT_SendMessageWorkOrderOperationAsBuilt_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateWorkOrderOutMsg
      - name: OpcenterEXDS_ProductionCoordination.ACT_DisAssignProducedMaterialitems
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMDisAssignProducedMaterialitems
      - name: OpcenterEXDS_ProductionCoordination.ACT_DeleteWOOFolders
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.DeleteWOOFolders
      - name: OpcenterEXDS_ProductionCoordination.ACT_DeleteWOOpDependencyNavigationList_WithConfirmation
      - name: OpcenterEXDS_ProductionCoordination.ACT_ChangeWorkOrderStatusToEdit_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.ChangeWorkOrderStatusToEdit
      - name: OpcenterEXDS_ProductionCoordination.ACT_SetWorkOrderForScheduling_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMSetWorkOrderForScheduling
      - name: OpcenterEXDS_ProductionCoordination.ACT_TriggerPrintingOnWorkOrder_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.TriggerPrintingOnWorkOrder
      - name: OpcenterEXDS_ProductionCoordination.ACT_ReleaseWorkOrder_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMReleaseWorkOrderList
      - name: OpcenterEXDS_ProductionCoordination.ACT_SetTargetQuantityOnFlexibleWorkOrder_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.SetTargetQuantityOnFlexibleWorkOrder
      - name: OpcenterEXDS_ProductionCoordination.ACT_AbortWorkOrder_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMAbortWorkOrder
      - name: OpcenterEXDS_ProductionCoordination.ACT_DeleteWorkOrder_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.DeleteWorkOrder
      - name: OpcenterEXDS_ProductionCoordination.ACT_MarkForCleaningWorkOrder_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.MarkForCleaningWorkOrderList
      - name: OpcenterEXDS_ProductionCoordination.ACT_SendMessageWorkOrder_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateWorkOrderOutMsg
      - name: OpcenterEXDS_ProductionCoordination.ACT_SendMessageWorkOrderAsBuilt
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateWorkOrderOutMsg
    target_commands:
      - AppU4DM.DeleteWorkOrderOperation
      - AppU4DM.TriggerPrintingOnWorkOrderOperation
      - AppU4DM.CreateWorkOrderOutMsg
      - AppU4DM.UADMDisAssignProducedMaterialitems
      - AppU4DM.DeleteWOOFolders
      - AppU4DM.DeleteWOOpDependencyNavigationList
      - AppU4DM.ChangeWorkOrderStatusToEdit
      - AppU4DM.UADMSetWorkOrderForScheduling
      - AppU4DM.TriggerPrintingOnWorkOrder
      - AppU4DM.UADMReleaseWorkOrder
      - AppU4DM.SetTargetQuantityOnFlexibleWorkOrder
      - AppU4DM.UADMAbortWorkOrder
      - AppU4DM.DeleteWorkOrder
      - AppU4DM.MarkForCleaningWorkOrderList

  - name: WorkOrderPreKitting_Master
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_SelectSerialNumber
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_PreKitSerialNumber
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_ReserveMaterialItems
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.ReserveMaterialItems
      - name: OpcenterEXDS_ProductionCoordination.ACT_GetPrekitSerialNumber
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_GetPrekitSerialNumber_Table
    target_commands:
      - AppU4DM.ReserveMaterialItems
      - OpcenterEXDS_ProductionCoordination.Java_MapPrekitMaterials

  - name: PANEL_PreKitSerialNumber_NoData
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: WorkOrderPreKitting_Details
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_Note_Management
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_Note_Create
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.SUB_UADMCreateSnagAndNoteList
      - name: OpcenterEXDS_ProductionCoordination.ACT_Note_Acknowledge
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.SUB_UADMConfirmSnagAndNoteList
    target_commands:
      - AppU4DM.UADMCreateSnagAndNoteList
      - AppU4DM.UADMConfirmSnagAndNoteList

  - name: PANEL_CreateToBeUsedMachine_PCD
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateToBeUsedMachine_PCD
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateToBeUsedMachine
    target_commands:
      - AppU4DM.CreateToBeUsedMachine

  - name: PANEL_SetWorkOrderOperationEstimatedTimes
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_SetWorkOrderOperationEstimatedTimes
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.SetWorkOrderOperationEstimatedTimes
    target_commands:
      - AppU4DM.SetWorkOrderOperationEstimatedTimes

  - name: PANEL_LinkUserToWorkOrderOperation
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.DS_GetUserDetailsList
        calls:
          - name: OpcenterEXDS_ProductionCoordination.DS_GetUserDetailsList_MF
            - name: OpcenterEXDS_ProductionCoordination_Connector.SUB_GetUserDetailsList
      - name: OpcenterEXDS_ProductionCoordination.ACT_LinkUserToWorkOrderOperation
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_LinkUserToWorkOrderOperation_MF
            - name: OpcenterEXDS_ProductionCoordination_Connector.SUB_LinkUserToWorkOrderOperationList
    target_commands:
      - AppU4DM.GetUserDetailsList
      - AppU4DM.LinkUserToWorkOrderOperationList

  - name: ProductionCoordinatorDashboard_Details
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_Operation_OnChange_By_NPE
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_SetWorkOrderOperationEstimatedTimes
            - name: OpcenterEXDS_ProductionCoordination_Connector.SetWorkOrderOperationEstimatedTimes
      - name: OpcenterEXDS_ProductionCoordination.ACT_UnlinkUserToWorkOrderOperation
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_UnlinkUserToWorkOrderOperation_MF
            - name: OpcenterEXDS_ProductionCoordination_Connector.SUB_UnlinkUserToWorkOrderOperationList
      - name: OpcenterEXDS_ProductionCoordination.ACT_SetPreferredMachine
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.SetPreferredMachine
      - name: OpcenterEXDS_ProductionCoordination.ACT_UADMReleaseWorkOrder
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_ReleaseWorkOrder_By_Id_WithConfirmation
            - name: OpcenterEXDS_ProductionCoordination_Connector.UADMReleaseWorkOrderList
      - name: OpcenterEXDS_ProductionCoordination.ACT_SetWorkOrderForScheduling
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_SetWorkOrderForScheduling_By_Id_WithConfirmation
            - name: OpcenterEXDS_ProductionCoordination_Connector.UADMSetWorkOrderForScheduling
    target_commands:
      - AppU4DM.SetWorkOrderOperationEstimatedTimes
      - AppU4DM.UnlinkUserToWorkOrderOperationList
      - AppU4DM.SetPreferredMachine
      - AppU4DM.UADMReleaseWorkOrder
      - AppU4DM.UADMSetWorkOrderForScheduling

  - name: ProductionCoordinatorDashboard_Master
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: QualityInspectionVisualView
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_TeamHistory
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_StepITLKCheckHistory
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_ToolHistory
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_ScrewingHistory
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_OperationSetpoint
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_ITLKCheckHistory
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_ReopenSerializedOperation
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_ReopenWorkOrderOperationSerialized_Command
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.ReopenWorkOrderOperationSerialized
    target_commands:
      - AppU4DM.ReopenWorkOrderOperationSerialized

  - name: PANEL_ReopenBatchOperation
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_ReopenWorkOrderOperationTransferBatch_Command
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.ReopenWorkOrderOperationTransferBatch
    target_commands:
      - AppU4DM.ReopenWorkOrderOperationTransferBatch

  - name: PANEL_DisassembleMaterialItem
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_DisassembleMaterialItem
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_DisassembleMaterialItem_Command
            - name: OpcenterEXDS_ProductionCoordination_Connector.DisassembleMaterialItem
    target_commands:
      - AppU4DM.DisassembleMaterialItem

  - name: PANEL_AcquisitionHistory
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_LinkExistingDocument
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_LinkExistingDocument
        calls:
          - name: OpcenterEXDS_MasterData_Connector.DSMaterial_LinkDocumentListToDM_MaterialTrackingUnit
    target_commands:
      - Material.DSMaterial_LinkDocumentListToDM_MaterialTrackingUnit

  - name: PANEL_LinkNewDocument
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_LinkNewDocument
        calls:
          - name: OpcenterEXDS_MasterData_Connector.DSMaterial_LinkDocumentListToDM_MaterialTrackingUnit
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMCreateDocument
    target_commands:
      - Material.DSMaterial_LinkDocumentListToDM_MaterialTrackingUnit
      - AppU4DM.UADMCreateDocument

  - name: Change_Details
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: WorkOrder_AsBuilt
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_SendMessageWorkOrderOperationAsBuilt
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateWorkOrderOutMsg
      - name: OpcenterEXDS_ProductionCoordination.ACT_AdministrativePause_Command
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.AdministrativePause
      - name: OpcenterEXDS_ProductionCoordination.ACT_ReopenWorkOrderOperation
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_ReopenWorkOrderOperationList_Command
            - name: OpcenterEXDS_ProductionCoordination_Connector.ReopenWorkOrderOperationList
      - name: OpcenterEXDS_ProductionCoordination.NAV_QualityInspectionVisualView
        calls:
          - name: OpcenterEXDS_ProductionCoordination.SUB_RuntimeInspectionDefinition_Context_Create
            - name: EXFN_Quality.SUB_ManageContexts
              - name: EXFN_Quality.ACT_Context_Create
                - name: EXFN_Quality.SUB_Context_Create
            - name: EXFN_Quality.SUB_RuntimeInspection_Visual_Create
              - name: EXFN_Quality.ACT_VisualInspection_Confirm_NF
                - name: EXFN_Quality.ACT_VisualInspection_Confirm
      - name: OpcenterEXDS_ProductionCoordination.ACT_UploadDocumentsToTeamcenterShare
        calls:
          - name: OpcenterEXDS_ShopfloorExecution_Connector.UploadDocuments
      - name: OpcenterEXDS_ProductionCoordination.ACT_UnlinkCollectedDocumentList_WithConfirmation
        calls:
          - name: OpcenterEXDS_MasterData_Connector.DSMaterial_UnlinkDocumentListToDM_MaterialTrackingUnit
      - name: OpcenterEXDS_ProductionCoordination.ACT_SendMessageWorkOrderAsBuilt
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateWorkOrderOutMsg
    target_commands:
      - AppU4DM.CreateWorkOrderOutMsg
      - AppU4DM.AdministrativePause
      - AppU4DM.ReopenWorkOrderOperationList
      - WorkInstruction.CreateInspectionAcquisitionContext
      - WorkInstruction.CreateVisualDetectedFailurewithES
      - AppU4DM.UploadDocuments
      - Material.DSMaterial_UnlinkDocumentListToDM_MaterialTrackingUnit

  - name: FirstArticleInspection_Master
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_RevokeFAICandidate
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.RevokeFAICandidate
      - name: OpcenterEXDS_ProductionCoordination.ACT_AbortFAIRecord
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.AbortFAIRecord
      - name: OpcenterEXDS_ProductionCoordination.ACT_CompleteFAIRecord
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CompleteFAIRecord
      - name: OpcenterEXDS_ProductionCoordination.NAV_CopyFAIRecord
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CopyFAIRecord
      - name: OpcenterEXDS_ProductionCoordination.DS_DownloadFAIReport
        calls:
          - name: OpcenterEXDS_ProductionCoordination.DS_CreateAndDownloadFAIReport
            - name: EXFN_DocumentViewer.CreateLinkDocument
            - name: OpcenterEXDS_ProductionCoordination.ACT_GenerateFAIRNumber
            - name: OpcenterEXDS_ProductionCoordination.DS_CreateFAIReport
              - name: OpcenterEXDS_ProductionCoordination.DS_CreateFAIReportForm1
                - name: OpcenterEXDS_ProductionCoordination.DS_CreateForm1Page
              - name: OpcenterEXDS_ProductionCoordination.DS_CreateFAIReportForm2
                - name: OpcenterEXDS_ProductionCoordination.DS_CreateForm2Page
              - name: OpcenterEXDS_ProductionCoordination.DS_CreateFAIReportForm3
                - name: OpcenterEXDS_ProductionCoordination.DS_CreateForm3Page
    target_commands:
      - AppU4DM.RevokeFAICandidate
      - AppU4DM.AbortFAIRecord
      - AppU4DM.CompleteFAIRecord
      - AppU4DM.CopyFAIRecord
      - OpcenterEXDS_ProductionCoordination.JA_HtmlToPdf
      - CommunityCommons.Base64EncodeFile
      - Document.CreateLinkDocument
      - AppU4DM.GenerateFAIRNumber
      - OpcenterEXDS_ProductionCoordination.JA_ReadTemplate
      - CommunityCommons.SubstringBefore
      - CommunityCommons.SubstringAfter
      - OpcenterEXDS_ProductionCoordination.JA_ReplacePlaceholdersForm1
      - CommunityCommons.RegexReplaceAll
      - OpcenterEXDS_ProductionCoordination.JA_ReplacePlaceholdersForm2
      - OpcenterEXDS_ProductionCoordination.JA_ReplacePlaceholdersForm3

  - name: FirstArticleInspection_Details
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_RevokeFAICandidate
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.RevokeFAICandidate
      - name: OpcenterEXDS_ProductionCoordination.ACT_AbortFAIRecord
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.AbortFAIRecord
      - name: OpcenterEXDS_ProductionCoordination.ACT_CompleteFAIRecord
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CompleteFAIRecord
      - name: OpcenterEXDS_ProductionCoordination.NAV_CopyFAIRecord
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CopyFAIRecord
      - name: OpcenterEXDS_ProductionCoordination.DS_DownloadFAIReport
        calls:
          - name: OpcenterEXDS_ProductionCoordination.DS_CreateAndDownloadFAIReport
            - name: EXFN_DocumentViewer.CreateLinkDocument
            - name: OpcenterEXDS_ProductionCoordination.ACT_GenerateFAIRNumber
            - name: OpcenterEXDS_ProductionCoordination.DS_CreateFAIReport
              - name: OpcenterEXDS_ProductionCoordination.DS_CreateFAIReportForm1
                - name: OpcenterEXDS_ProductionCoordination.DS_CreateForm1Page
              - name: OpcenterEXDS_ProductionCoordination.DS_CreateFAIReportForm2
                - name: OpcenterEXDS_ProductionCoordination.DS_CreateForm2Page
              - name: OpcenterEXDS_ProductionCoordination.DS_CreateFAIReportForm3
                - name: OpcenterEXDS_ProductionCoordination.DS_CreateForm3Page
    target_commands:
      - AppU4DM.RevokeFAICandidate
      - AppU4DM.AbortFAIRecord
      - AppU4DM.CompleteFAIRecord
      - AppU4DM.CopyFAIRecord
      - OpcenterEXDS_ProductionCoordination.JA_HtmlToPdf
      - CommunityCommons.Base64EncodeFile
      - Document.CreateLinkDocument
      - AppU4DM.GenerateFAIRNumber
      - OpcenterEXDS_ProductionCoordination.JA_ReadTemplate
      - CommunityCommons.SubstringBefore
      - CommunityCommons.SubstringAfter
      - OpcenterEXDS_ProductionCoordination.JA_ReplacePlaceholdersForm1
      - CommunityCommons.RegexReplaceAll
      - OpcenterEXDS_ProductionCoordination.JA_ReplacePlaceholdersForm2
      - OpcenterEXDS_ProductionCoordination.JA_ReplacePlaceholdersForm3

  - name: FirstArticleInspection_InspectionSampleValues
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.NAV_QualityInspectionVisualView
        calls:
          - name: OpcenterEXDS_ProductionCoordination.SUB_RuntimeInspectionDefinition_Context_Create
            - name: EXFN_Quality.SUB_ManageContexts
              - name: EXFN_Quality.ACT_Context_Create
                - name: EXFN_Quality.SUB_Context_Create
            - name: EXFN_Quality.SUB_RuntimeInspection_Visual_Create
              - name: EXFN_Quality.ACT_VisualInspection_Confirm_NF
                - name: EXFN_Quality.ACT_VisualInspection_Confirm
    target_commands:
      - WorkInstruction.CreateInspectionAcquisitionContext
      - WorkInstruction.CreateVisualDetectedFailurewithES

  - name: FirstArticleInspection_InspectionHistory
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_ToastNotification
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_CopyFAIRecordWizard
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CopyFAIRecord
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CopyFAIRecord
    target_commands:
      - AppU4DM.CopyFAIRecord

  - name: PANEL_CopyFAIRecordOnlyCompleteRelated
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CopyFAIRecord_OnlyCompleteRelated
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CopyFAIRecord
    target_commands:
      - AppU4DM.CopyFAIRecord

  - name: PANEL_ChangeFAICandidate
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_ChangeFAICandidate
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.ChangeFAICandidate
    target_commands:
      - AppU4DM.ChangeFAICandidate

  - name: PANEL_DeclareFAICandidate
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_DeclareFAICandidate
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.DeclareFAICandidate
    target_commands:
      - AppU4DM.DeclareFAICandidate

  - name: PANEL_UpdateLSPThreshold
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UpdateLSPThreshold
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UpdateLSPThreshold
    target_commands:
      - Kanban.UpdateLSPThreshold

  - name: PANEL_UpdateLineSidePosition
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UpdateLineSidePosition
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UpdateLineSidePosition_OnlyManual
          - name: OpcenterEXDS_ProductionCoordination_Connector.UpdateLineSidePosition_NotOnlyManual
    target_commands:
      - Kanban.UpdateLineSidePosition

  - name: PANEL_CreateLSPThreshold
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateLSPThreshold
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateLSPThreshold
    target_commands:
      - Kanban.CreateLSPThreshold

  - name: PANEL_CreateLineSidePosition
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateLineSidePosition
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateLineSidePosition_OnlyManual
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateLineSidePosition_NotOnlyManual
    target_commands:
      - Kanban.CreateLineSidePosition

  - name: LineSidePosition_Details
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_DeleteLSPThreshold_WithConfirmation
      - name: OpcenterEXDS_ProductionCoordination.ACT_ReleaseLineSidePosition_WithConfirmation
      - name: OpcenterEXDS_ProductionCoordination.ACT_UnreleaseLineSidePosition_WithConfirmation
    target_commands:
      - Kanban.DeleteLSPThreshold
      - Kanban.ReleaseLineSidePosition
      - Kanban.UnreleaseLineSidePosition

  - name: LineSidePositions_Master
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_ReleaseLineSidePosition_WithConfirmation
      - name: OpcenterEXDS_ProductionCoordination.ACT_UnreleaseLineSidePosition_WithConfirmation
      - name: OpcenterEXDS_ProductionCoordination.ACT_DeleteLineSidePosition_WithConfirmation
    target_commands:
      - Kanban.ReleaseLineSidePosition
      - Kanban.UnreleaseLineSidePosition
      - Kanban.DeleteLineSidePosition

  - name: Panel_Scanner
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.OCH_MaterialTrackingUnit_Scan
        calls:
          - name: OpcenterEXDS_ProductionCoordination.SUB_MaterialTrackingUnit_Get
            - name: OpcenterEXDS_ProductionCoordination_Connector.AddMaterialTrackingUnitsToWorkingSession
            - name: OpcenterEXDS_ProductionCoordination.SUB_VerifySession
              - name: OpcenterEXDS_ProductionCoordination_Connector.CreateWorkingSession
    target_commands:
      - OpcenterEXDS_ProductionCoordination.JAVA_GetClientId
      - AppU4DM.AddMaterialTrackingUnitsToWorkingSession
      - AppU4DM.CreateWorkingSession

  - name: ScanMaterialTrackingUnits_Master
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UpdateSession
        calls:
          - name: OpcenterEXDS_ProductionCoordination.SUB_GetClientId
      - name: OpcenterEXDS_ProductionCoordination.OCH_MaterialTrackingUnit_Scan
        calls:
          - name: OpcenterEXDS_ProductionCoordination.SUB_MaterialTrackingUnit_Get
            - name: OpcenterEXDS_ProductionCoordination_Connector.AddMaterialTrackingUnitsToWorkingSession
            - name: OpcenterEXDS_ProductionCoordination.SUB_VerifySession
              - name: OpcenterEXDS_ProductionCoordination_Connector.CreateWorkingSession
      - name: OpcenterEXDS_ProductionCoordination.ACT_InitializeMaterialTrackingUnit_MTUsDTO
        calls:
          - name: OpcenterEXDS_ProductionCoordination.SUB_MaterialTrackingUnit_MTUsDTO_Get
            - name: OpcenterEXDS_ProductionCoordination_Connector.CreateWorkingSession
          - name: OpcenterEXDS_ProductionCoordination.SUB_GetClientId
      - name: OpcenterEXDS_ProductionCoordination.OCH_MaterialTrackingUnit_Discard
        calls:
          - name: OpcenterEXDS_ProductionCoordination.SUB_MaterialTrackingUnit_Discard
            - name: OpcenterEXDS_ProductionCoordination_Connector.RemoveMaterialTrackingUnitsFromWorkingSession
    target_commands:
      - OpcenterEXDS_ProductionCoordination.JAVA_GetClientId
      - AppU4DM.AddMaterialTrackingUnitsToWorkingSession
      - AppU4DM.CreateWorkingSession
      - AppU4DM.RemoveMaterialTrackingUnitsFromWorkingSession

  - name: Panel_MaterialTrackingUnit_ReleaseByFile
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_MTUReleaseByFile
        calls:
          - name: OpcenterEXDS_ProductionCoordination.SUB_MTUReleaseByFile_BulkMaterialTrackingUnitsReleaseFromFile
            - name: OpcenterEXDS_ProductionCoordination_Connector.BulkMaterialTrackingUnitsReleaseFromFile
      - name: OpcenterEXDS_ProductionCoordination.ACT_GetMTUReleaseByFileItemsTotalCount
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_ProcessUploadedCSV
    target_commands:
      - AppU4DM.BulkMaterialTrackingUnitsReleaseFromFile
      - OpcenterEXDS_ProductionCoordination.JAVA_GetItemsTotalCount
      - OpcenterEXDS_ProductionCoordination.JAVA_ValidateFileHeader

  - name: Panel_MaterialTrackingUnit_Scrap
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_GenerateMTUScrapOPNumberingPattern
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.GenerateOPNumberingPattern
      - name: OpcenterEXDS_ProductionCoordination.ACT_MTUScrap
        calls:
          - name: OpcenterEXDS_ProductionCoordination.SUB_ACT_MTUScrap
            - name: OpcenterEXDS_ProductionCoordination.SUB_BulkMaterialTrackingUnitsScrapFromNonConformanceItems
              - name: OpcenterEXDS_ProductionCoordination_Connector.BulkMaterialTrackingUnitsScrapFromNonConformanceItems
            - name: OpcenterEXDS_ProductionCoordination.SUB_MTUScrap_BulkMaterialTrackingUnitsScrap_SetNIdAutomatically
              - name: OpcenterEXDS_ProductionCoordination_Connector.BulkMaterialTrackingUnitsScrap
            - name: OpcenterEXDS_ProductionCoordination.SUB_MTUScrap_BulkMaterialTrackingUnitsScrap_SetNId
              - name: OpcenterEXDS_ProductionCoordination_Connector.BulkMaterialTrackingUnitsScrap
            - name: OpcenterEXDS_ProductionCoordination_Connector.BulkMaterialTrackingUnitsScrap
    target_commands:
      - AppU4DM.GenerateOPNumberingPattern
      - AppU4DM.BulkMaterialTrackingUnitsScrapFromNonConformanceItems
      - AppU4DM.BulkMaterialTrackingUnitsScrap

  - name: Panel_MaterialTrackingUnit_Release
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_BulkMaterialTrackingUnitsReleaseFromNonConformanceItems
        calls:
          - name: OpcenterEXDS_ProductionCoordination.SUB_BulkMaterialTrackingUnitsReleaseFromNonConformanceItems
            - name: OpcenterEXDS_ProductionCoordination_Connector.BulkMaterialTrackingUnitsReleaseFromNonConformanceItems
    target_commands:
      - AppU4DM.BulkMaterialTrackingUnitsReleaseFromNonConformanceItems

  - name: Panel_MaterialTrackingUnit_ScrapByFile
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_MTUScrapByFile
        calls:
          - name: OpcenterEXDS_ProductionCoordination.SUB_MTUScrapByFile_BulkMaterialTrackingUnitsScrap
            - name: OpcenterEXDS_ProductionCoordination_Connector.BulkMaterialTrackingUnitsScrapFromFile
      - name: OpcenterEXDS_ProductionCoordination.ACT_GetMTUScrapByFileItemsTotalCount
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_ProcessUploadedCSV
      - name: OpcenterEXDS_ProductionCoordination.ACT_GenerateMTUScrapByFileOPNumberingPattern
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.GenerateOPNumberingPattern
    target_commands:
      - AppU4DM.BulkMaterialTrackingUnitsScrapFromFile
      - OpcenterEXDS_ProductionCoordination.JAVA_GetItemsTotalCount
      - OpcenterEXDS_ProductionCoordination.JAVA_ValidateFileHeader
      - AppU4DM.GenerateOPNumberingPattern

  - name: PANEL_EditNPAForUser
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UADMUpdateNonProductiveActivityForUser
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMUpdateNonProductiveActivityForUser
    target_commands:
      - AppU4DM.UADMUpdateNonProductiveActivityForUser

  - name: PANEL_CreateNPAForUser
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UADMRegisterNonProductiveActivityForUser
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMRegisterNonProductiveActivityForUser
    target_commands:
      - AppU4DM.UADMRegisterNonProductiveActivityForUser

  - name: Users_Details
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UADMDeleteNonProductiveActivityListForUser
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMDeleteNonProductiveActivityListForUser
      - name: OpcenterEXDS_ProductionCoordination.ACT_UADMCompleteNonProductiveActivityList
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMCompleteNonProductiveActivityList
    target_commands:
      - AppU4DM.UADMDeleteNonProductiveActivityListForUser
      - AppU4DM.UADMCompleteNonProductiveActivityList

  - name: Users_Master
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_UpdateExecutionGroup
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UpdateExecutionGroup
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UpdateExecutionGroup
    target_commands:
      - AppU4DM.UpdateExecutionGroup

  - name: PANEL_CreateExecutionGroup
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateExecutionGroup
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateExecution
    target_commands:
      - AppU4DM.UADMCreateExecutionGroup

  - name: ExecutionGroup_Master
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_ReleaseExecGroup_Master
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.ReleaseExecutionGroup
      - name: OpcenterEXDS_ProductionCoordination.ACT_ChangeStatusToSetReady
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.SetReadyForSchedulingExecutionGroup
      - name: OpcenterEXDS_ProductionCoordination.ACT_ChangeStatusToEdit
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.SetExecGroupListInEditStatus
      - name: OpcenterEXDS_ProductionCoordination.ACT_AbortExecGroup
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.AbortExecutionGroup
      - name: OpcenterEXDS_ProductionCoordination.DeleteExecutionGroup
        calls:
          - name: OpcenterEXFN_MasterData_Connector.SUB_DeleteCommand_StartArray
      - name: OpcenterEXDS_ProductionCoordination_Connector.GetReadyForPrintJob
    target_commands:
      - AppU4DM.ReleaseExecutionGroup
      - AppU4DM.ScheduleExecutionGroup
      - AppU4DM.SetExecGroupListInEditStatus
      - AppU4DM.AbortExecutionGroupList
      - AppName.CommandName
      - AppU4DM.AppU4DMAMN_GetReadyPrintJobs

  - name: PANEL_EditWOOP
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UpdateWOOP
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UpdateWOOpQuantityAssociatedToEG
    target_commands:
      - AppU4DM.UpdateWOOpQuantityAssociatedToEG

  - name: PANEL_LinkWOOperationToExecutionGroup
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_LinkWOOperationToExecutionGroup
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.LinkWOOperationListToExecutionGroup
    target_commands:
      - AppU4DM.LinkWOOperationListToExecutionGroup

  - name: PANEL_UpdateExecutionGroup_Details
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UpdateExecutionGroup_Details
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UpdateExecutionGroup
    target_commands:
      - AppU4DM.UpdateExecutionGroup

  - name: ExecutionGroup_Details
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UnlinkWOOperationsFromExecutionGroup
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UnlinkWOOperationsFromExecutionGroup
      - name: OpcenterEXDS_ProductionCoordination.ACT_ReleaseExecGroup_Details
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.ReleaseExecutionGroup
      - name: OpcenterEXDS_ProductionCoordination.ACT_SetReadyExecGroup_Details
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.SetReadyForSchedulingExecutionGroup
      - name: OpcenterEXDS_ProductionCoordination.ACT_SetEditExecGroup_Details
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.SetExecGroupListInEditStatus
      - name: OpcenterEXDS_ProductionCoordination.ACT_AbortExecGroup_Details
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.AbortExecutionGroup
      - name: OpcenterEXDS_ProductionCoordination.DeleteExecutionGroup
        calls:
          - name: OpcenterEXFN_MasterData_Connector.SUB_DeleteCommand_StartArray
    target_commands:
      - AppU4DM.UnlinkWOOperationsFromExecutionGroup
      - AppU4DM.ReleaseExecutionGroup
      - AppU4DM.ScheduleExecutionGroup
      - AppU4DM.SetExecGroupListInEditStatus
      - AppU4DM.AbortExecutionGroupList
      - AppName.CommandName

  - name: ExecutionGroupPhase_Details
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UnlinkWIToEGPhase
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UnlinkWIDefinitionsFromEGPhase
      - name: OpcenterEXDS_ProductionCoordination.ACT_DeletePowder
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.RemoveAMPowderLoadedCheckFromEgPhaseList
    target_commands:
      - AppU4DM.UnlinkWIDefinitionsFromEGPhase
      - PowderMgt.RemoveAMPowderLoadedCheckFromEgPhaseList

  - name: PANEL_LinkPrintJobFile
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_LinkPrintJobFile
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.LinkToBeUsedMachineToPJFForEGPhase
    target_commands:
      - AppU4DM.LinkToBeUsedMachineToPJFForEGPhase

  - name: PANEL_EditAMPowder
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_EditAMPowder
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.EditAMPowderLoadedCheckOnEgPhase
    target_commands:
      - PowderMgt.EditAMPowderLoadedCheckOnEgPhase

  - name: PANEL_UpdateExecutionGroupPhase
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UpdateExecutionGroupPhase
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UpdateExecutionGroupPhase
    target_commands:
      - AppU4DM.UpdateExecutionGroupPhase

  - name: PANEL_PreTransfer
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CallPreTransfer
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_PreTransfer
            - name: OpcenterEXDS_ProductionCoordination_Connector.RequestTransferPrintJobFile
    target_commands:
      - PrintJobFile.RequestTransferPrintJobFile

  - name: PANEL_LinkWInstructionToExecutionGroupPhase
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_LinkWIToEGPhase
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.LinkWIDefinitionsToEGPhase
    target_commands:
      - AppU4DM.LinkWIDefinitionsToEGPhase

  - name: PANEL_LoadAMPowder
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_SetPowder
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.SetAMPowderLoadedCheckOnEgPhase
    target_commands:
      - PowderMgt.SetAMPowderLoadedCheckOnEgPhase

  - name: OfflineSession_Details
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CheckInOfflineAction
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_CheckInActionPayloadToCmd
            - name: OpcenterEXDS_ProductionCoordination_Connector.ChangeOfflineSessionStatus
      - name: OpcenterEXDS_ProductionCoordination.ACT_DiscardOfflineAction
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UpdateOfflineAction
          - name: OpcenterEXDS_ProductionCoordination_Connector.CompleteCheckIn
    target_commands:
      - OpcenterEXDS_ProductionCoordination.Java_ConvertToCmdActionAmend
      - CmdApp.CmdName
      - OpcenterEXDS_ProductionCoordination.Java_ConvertToCmdAction
      - AppU4DM.ChangeOfflineSessionStatus
      - AppU4DM.UpdateOfflineAction
      - AppU4DM.CompleteCheckIn

  - name: PANEL_EditOfflineAction
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_OpenEditAction
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_GetOfflineActionAmendPayload
      - name: OpcenterEXDS_ProductionCoordination.ACT_AmmendOfflineAction
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.AmendOfflineAction
    target_commands:
      - OpcenterEXDS_ProductionCoordination.Java_DecodeBinaryActionAmendToJson
      - OpcenterEXDS_ProductionCoordination.Java_DecodedActionToJson
      - AppU4DM.AmendOfflineAction

  - name: OfflineSession_Master
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CheckInOfflineSession
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_ChangeOfflineSessionStatus
            - name: OpcenterEXDS_ProductionCoordination_Connector.ChangeOfflineSessionStatus
            - name: OpcenterEXDS_ProductionCoordination_Connector.CompleteCheckIn
      - name: OpcenterEXDS_ProductionCoordination.ACT_DiscardOfflineSessionList
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.DiscardOfflineSessionList
    target_commands:
      - OpcenterEXDS_ProductionCoordination.Java_ConvertToCmdActionAmend
      - OpcenterEXDS_ProductionCoordination.Java_ConvertToCmdAction
      - IteratorOfflineAction/CommandApp.IteratorOfflineAction/CommandName
      - AppU4DM.ChangeOfflineSessionStatus
      - AppU4DM.CompleteCheckIn
      - AppU4DM.DiscardOfflineSessionList

  - name: WorkOrderUpdateCheck_Master
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UADMUpdateOutOfDateWOByBoPList_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMUpdateOutOfDateWOByBoPList
      - name: OpcenterEXDS_ProductionCoordination.ACT_PLMUpdateWorkOrderByCCList_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.PLMUpdateWorkOrderByCCList
    target_commands:
      - AppU4DM.UADMUpdateOutOfDateWOByBoPList
      - AppU4DM.PLMUpdateWorkOrderByCCList

  - name: PANEL_UpdateLSPAvailableQuantity
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UpdateLSPAvailableQuantity
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UpdateLSPAvailableQuantity
    target_commands:
      - Kanban.UpdateLSPAvailableQuantity

  - name: PANEL_DeclareKanbanCallDelivered
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_DeclareKanbanCallDelivered
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.DeclareKanbanCallDelivered
    target_commands:
      - Kanban.DeclareKanbanCallDelivered

  - name: LineSidePositionMonitoring_KanbanRequest_Details
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: LineSidePositionMonitoring_Details
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateManualKanbanCall_WithConfirmation
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateOnlyManualKanbanCall_WithConfirmation
      - name: OpcenterEXDS_ProductionCoordination.ACT_CancelKanbanCall_WithConfirmation
      - name: OpcenterEXDS_ProductionCoordination.ACT_IncreaseKanbanCallPriority_WithConfirmation
    target_commands:
      - Kanban.CreateManualKanbanCall
      - Kanban.CreateOnlyManualKanbanCall
      - Kanban.CancelKanbanCall
      - Kanban.IncreaseKanbanCallPriority

  - name: LineSidePositionsMonitoring_Master
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateManualKanbanCall_WithConfirmation
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateOnlyManualKanbanCall_WithConfirmation
      - name: OpcenterEXDS_ProductionCoordination.ACT_CancelKanbanCall_WithConfirmation
      - name: OpcenterEXDS_ProductionCoordination.ACT_IncreaseKanbanCallPriority_WithConfirmation
    target_commands:
      - Kanban.CreateManualKanbanCall
      - Kanban.CreateOnlyManualKanbanCall
      - Kanban.CancelKanbanCall
      - Kanban.IncreaseKanbanCallPriority

  - name: ContainmentRequests_Master
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: ContainmentRequests_Details
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_SetContainmentRequestAsReady
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_SetContainmentRequestAsReady
        calls:
          - name: OpcenterEXDS_ProductionCoordination.SUB_SetContainmentRequestStatus
            - name: OpcenterEXDS_ProductionCoordination_Connector.SetContainmentRequestStatus
    target_commands:
      - AppU4DM.SetContainmentRequestStatus

  - name: Panel_ManageContainmentRequestRelease
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_GetReleaseFileItemsTotalCount
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_ProcessUploadedCSV
      - name: OpcenterEXDS_ProductionCoordination.ACT_ManageContainmentRequestRelease
        calls:
          - name: OpcenterEXDS_ProductionCoordination.SUB_ProcessContainmentRequestRelease
            - name: OpcenterEXDS_ProductionCoordination_Connector.BulkContainmentRequestRelease
    target_commands:
      - OpcenterEXDS_ProductionCoordination.JAVA_GetItemsTotalCount
      - OpcenterEXDS_ProductionCoordination.JAVA_ValidateFileHeader
      - AppU4DM.BulkContainmentRequestRelease

  - name: PANEL_BulkMaterialTrackingUnitsScrapFromContainmentRequest
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_BulkMaterialTrackingUnitsScrapFromContainmentRequest
        calls:
          - name: OpcenterEXDS_ProductionCoordination.SUB_ACT_BulkMaterialTrackingUnitsScrapFromContainmentRequest
            - name: OpcenterEXDS_ProductionCoordination.SUB_BulkMaterialTrackingUnitsScrapFromContainmentRequest_NewNonConformance_SetNIdAutomatically
              - name: OpcenterEXDS_ProductionCoordination_Connector.BulkMaterialTrackingUnitsScrapFromContainmentRequest
            - name: OpcenterEXDS_ProductionCoordination.SUB_BulkMaterialTrackingUnitsScrapFromContainmentRequest_NewNonConformance_SetNId
              - name: OpcenterEXDS_ProductionCoordination_Connector.BulkMaterialTrackingUnitsScrapFromContainmentRequest
            - name: OpcenterEXDS_ProductionCoordination.SUB_BulkMaterialTrackingUnitsScrapFromContainmentRequest_NoNC
              - name: OpcenterEXDS_ProductionCoordination_Connector.BulkMaterialTrackingUnitsScrapFromContainmentRequest
      - name: OpcenterEXDS_ProductionCoordination.ACT_GenerateBulkScrapOPNumberingPattern
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.GenerateOPNumberingPattern
    target_commands:
      - AppU4DM.BulkMaterialTrackingUnitsScrapFromContainmentRequest
      - AppU4DM.GenerateOPNumberingPattern

  - name: PANEL_ManageContainmentRequestFullRelease
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_ManageContainmentRequestFullRelease
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.BulkContainmentRequestRelease
    target_commands:
      - AppU4DM.BulkContainmentRequestRelease

  - name: Panel_CreateContainmentRequest
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateContainmentRequest
        calls:
          - name: OpcenterEXDS_ProductionCoordination.SUB_CreateContainmentRequest
            - name: OpcenterEXDS_ProductionCoordination_Connector.CreateContainmentRequest
    target_commands:
      - AppU4DM.CreateContainmentRequest

  - name: PANEL_CloseContainmentRequest
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CloseContainmentRequest
        calls:
          - name: OpcenterEXDS_ProductionCoordination.SUB_CloseContainmentRequest
            - name: OpcenterEXDS_ProductionCoordination_Connector.SetContainmentRequestStatus
    target_commands:
      - AppU4DM.SetContainmentRequestStatus

  - name: PANEL_BulkMaterialTrackingUnitsScrap
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_GetScrapFileItemsTotalCount
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_ProcessUploadedCSV
      - name: OpcenterEXDS_ProductionCoordination.ACT_GenerateScrapOPNumberingPattern
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.GenerateOPNumberingPattern
      - name: OpcenterEXDS_ProductionCoordination.ACT_BulkMaterialTrackingUnitsScrap
        calls:
          - name: OpcenterEXDS_ProductionCoordination.SUB_ACT_BulkMaterialTrackingUnitsScrap
            - name: OpcenterEXDS_ProductionCoordination.SUB_BulkMaterialTrackingUnitsScrap_SetNIdAutomatically
              - name: OpcenterEXDS_ProductionCoordination_Connector.BulkMaterialTrackingUnitsScrap
              - name: OpcenterEXDS_ProductionCoordination_Connector.BulkMaterialTrackingUnitsScrapFromFile
            - name: OpcenterEXDS_ProductionCoordination.SUB_BulkMaterialTrackingUnitsScrap_NoNC
              - name: OpcenterEXDS_ProductionCoordination_Connector.BulkMaterialTrackingUnitsScrap
    target_commands:
      - OpcenterEXDS_ProductionCoordination.JAVA_GetItemsTotalCount
      - OpcenterEXDS_ProductionCoordination.JAVA_ValidateFileHeader
      - AppU4DM.GenerateOPNumberingPattern
      - AppU4DM.BulkMaterialTrackingUnitsScrap
      - AppU4DM.BulkMaterialTrackingUnitsScrapFromFile

  - name: PANEL_CompareChangePackageItem_ForMachine
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_AddChangePackageItemMachine_Workcenter
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_AddChangePackageItemMachine
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.AddChangePackageItemMachine
    target_commands:
      - AppU4DM.AddChangePackageItemMachine

  - name: PANEL_AddChangePackageItemMachine_EquipmentType
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_AddChangePackageItemMachine
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.AddChangePackageItemMachine
    target_commands:
      - AppU4DM.AddChangePackageItemMachine

  - name: PANEL_AddChangePackageItemMachine_Unit
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_AddChangePackageItemMachine
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.AddChangePackageItemMachine
    target_commands:
      - AppU4DM.AddChangePackageItemMachine

  - name: PANEL_AddChangePackageItemMaterial
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_AddChangePackageItemMaterial
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.AddChangePackageItemMaterial
    target_commands:
      - AppU4DM.AddChangePackageItemMaterial

  - name: PANEL_CompareChangePackageItem_ForMaterial
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_CompareChangePackageItem
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: WorkInstruction_PreviewForNewInstructions
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_AddChangePackageItemWIDefinition
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_AddChangePackageItemWIDefinition
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.AddChangePackageItemWIDefinition
    target_commands:
      - AppU4DM.AddChangePackageItemWIDefinition

  - name: PANEL_UADMEditChangePackageItemWIDefinition
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UADMEditChangePackageItemWIDefinition
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMEditChangePackageItemWIDefinition
    target_commands:
      - AppU4DM.UADMEditChangePackageItemWIDefinition

  - name: PANEL_AddChangePackageItemWorkOrderStep
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_AddChangePackageItemWorkOrderStep
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.AddChangePackageItemWorkOrderStep
    target_commands:
      - AppU4DM.AddChangePackageItemWorkOrderStep

  - name: PANEL_CompareChangePackageItem_ForStep
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_AddChangePackageItemStepWIDefinition
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_AddChangePackageItemStepWIDefinition
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.AddChangePackageItemStepWIDefinition
    target_commands:
      - AppU4DM.AddChangePackageItemStepWIDefinition

  - name: WorkInstruction_PreviewForNewInstructions_ForWIStep
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_UADMEditChangePackageItemStepWIDefinition
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UADMEditChangePackageItemStepWIDefinition
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMEditChangePackageItemStepWIDefinition
    target_commands:
      - AppU4DM.UADMEditChangePackageItemStepWIDefinition

  - name: PANEL_CompareChangePackageItem_ForWIStep
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_AddChangePackageItemTool_ForStepTool
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_AddChangePackageItemTool_ForStepTool
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.AddChangePackageItemTool
    target_commands:
      - AppU4DM.AddChangePackageItemTool

  - name: PANEL_AddChangePackageItemMaterial_ForStep
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_AddChangePackageItemMaterial_ForStep
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.AddChangePackageItemMaterial
    target_commands:
      - AppU4DM.AddChangePackageItemMaterial

  - name: PANEL_AddChangePackageItemQualityInspection_Step
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_AddChangePackageItemQualityInspection_Step
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.AddChangePackageItemQualityInspection
    target_commands:
      - AppU4DM.AddChangePackageItemQualityInspection

  - name: PANEL_AddChangePackageItemDocument_ForStep
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_AddChangePackageItemDocument_ForStep
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.AddChangePackageItemDocument
    target_commands:
      - AppU4DM.AddChangePackageItemDocument

  - name: PANEL_AddChangePackageItemNewDocument_ForStep
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_AddChangePackageItemNewDocument_ForStep
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.AddChangePackageItemNewDocument
    target_commands:
      - AppU4DM.AddChangePackageItemNewDocument

  - name: Steps_Details
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_RevertChangePackageItemList_ForStepDocument
        calls:
          - name: OpcenterEXDS_Core.SUB_DeleteCommand_StartArray
      - name: OpcenterEXDS_ProductionCoordination.ACT_RemoveChangePackageItemDocument_ForStep
        calls:
          - name: OpcenterEXDS_Core.SUB_DeleteCommand_TwoParameter
      - name: OpcenterEXDS_ProductionCoordination.ACT_RevertChangePackageItemList_ForStepTool
        calls:
          - name: OpcenterEXDS_Core.SUB_DeleteCommand_StartArray
      - name: OpcenterEXDS_ProductionCoordination.ACT_RemoveChangePackageItemTool_2
        calls:
          - name: OpcenterEXDS_Core.SUB_DeleteCommand_TwoParameter
      - name: OpcenterEXDS_ProductionCoordination.ACT_RevertChangePackageItemList_ForStepMaterial
        calls:
          - name: OpcenterEXDS_Core.SUB_DeleteCommand_StartArray
      - name: OpcenterEXDS_ProductionCoordination.ACT_RemoveChangePackageItemMaterial_2
        calls:
          - name: OpcenterEXDS_Core.SUB_DeleteCommand_TwoParameter
      - name: OpcenterEXDS_ProductionCoordination.ACT_RevertChangePackageItemList_ForStepWI
        calls:
          - name: OpcenterEXDS_Core.SUB_DeleteCommand_StartArray
      - name: OpcenterEXDS_ProductionCoordination.ACT_RemoveChangePackageItemStepWIDefinition
        calls:
          - name: OpcenterEXDS_Core.SUB_DeleteCommand_TwoParameter
      - name: OpcenterEXDS_ProductionCoordination.ACT_RevertChangePackageItemList_ForStepQI
        calls:
          - name: OpcenterEXDS_Core.SUB_DeleteCommand_StartArray
      - name: OpcenterEXDS_ProductionCoordination.ACT_RemoveChangePackageItemQualityInspection_2
        calls:
          - name: OpcenterEXDS_Core.SUB_DeleteCommand_TwoParameter
    target_commands:
      - AppName.CommandName

  - name: PANEL_AddChangePackageWorkOrderOperation
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_AddChangePackageWorkOrderOperation
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_AddChangePackageWorkOrderOperation_Command
    target_commands:
      - AppU4DM.AddChangePackageItemWorkOrderOperation

  - name: PANEL_ChangePackageAddWOOpDependencies
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_AddChangePackageWorkOrderOperationDependency
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_AddChangePackageWorkOrderOperationDependency_Command
    target_commands:
      - AppU4DM.AddChangePackageItemWorkOOperationDependency

  - name: PANEL_CompareChangePackageItem_ForSummary
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_AddChangePackageItemQualityInspection
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_AddChangePackageItemQualityInspection
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.AddChangePackageItemQualityInspection
    target_commands:
      - AppU4DM.AddChangePackageItemQualityInspection

  - name: PANEL_CompareChangePackageItem_ForQI
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_AddChangePackageItemNewDocument
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_AddChangePackageItemNewDocument
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.AddChangePackageItemNewDocument
    target_commands:
      - AppU4DM.AddChangePackageItemNewDocument

  - name: PANEL_AddChangePackageItemDocument
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_AddChangePackageItemDocument
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.AddChangePackageItemDocument
    target_commands:
      - AppU4DM.AddChangePackageItemDocument

  - name: PANEL_CompareChangePackageItem_ForDocument
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PreviewRouting
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_AddChangePackageItemTool
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_AddChangePackageItemTool
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.AddChangePackageItemTool
    target_commands:
      - AppU4DM.AddChangePackageItemTool

  - name: PANEL_CompareChangePackageItem_ForTools
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: ChangePackages_WorkOrderOverview_Details
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_RevertChangePackageItemList_ForWorkOrderOperation
        calls:
          - name: OpcenterEXDS_Core.SUB_DeleteCommand_StartArray
      - name: OpcenterEXDS_ProductionCoordination.ACT_RemoveChangePackageWorkOrderOperation
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_RemoveChangePackageWorkOrderOperation_Command
      - name: OpcenterEXDS_ProductionCoordination.ACT_RevertChangePackageItemList_ForWorkOrderOperation_Dependency
        calls:
          - name: OpcenterEXDS_Core.SUB_DeleteCommand_StartArray
      - name: OpcenterEXDS_ProductionCoordination.ACT_RemoveChangePackageWorkOrderOperationDependency
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_RemoveChangePackageWorkOrderOperationDependency_Command
    target_commands:
      - AppName.CommandName
      - AppU4DM.RemoveChangePackageItemWorkOrderOperation
      - AppU4DM.RemoveChangePackageItemWorkOOperationDependency

  - name: ChangePackages_Details
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_RevertChangePackageItemList_ForDocument
        calls:
          - name: OpcenterEXDS_Core.SUB_DeleteCommand_StartArray
      - name: OpcenterEXDS_ProductionCoordination.ACT_RemoveChangePackageItemDocument
        calls:
          - name: OpcenterEXDS_Core.SUB_DeleteCommand_TwoParameter
      - name: OpcenterEXDS_ProductionCoordination.ACT_RevertChangePackageItemList_ForTool
        calls:
          - name: OpcenterEXDS_Core.SUB_DeleteCommand_StartArray
      - name: OpcenterEXDS_ProductionCoordination.ACT_RemoveChangePackageItemTool
        calls:
          - name: OpcenterEXDS_Core.SUB_DeleteCommand_TwoParameter
      - name: OpcenterEXDS_ProductionCoordination.ACT_RevertChangePackageItemList_ForMaterial
        calls:
          - name: OpcenterEXDS_Core.SUB_DeleteCommand_StartArray
      - name: OpcenterEXDS_ProductionCoordination.ACT_RemoveChangePackageItemMaterial
        calls:
          - name: OpcenterEXDS_Core.SUB_DeleteCommand_TwoParameter
      - name: OpcenterEXDS_ProductionCoordination.ACT_RevertChangePackageItemList_ForWI
        calls:
          - name: OpcenterEXDS_Core.SUB_DeleteCommand_StartArray
      - name: OpcenterEXDS_ProductionCoordination.ACT_RemoveChangePackageItemWIDefinition
        calls:
          - name: OpcenterEXDS_Core.SUB_DeleteCommand_TwoParameter
      - name: OpcenterEXDS_ProductionCoordination.ACT_RevertChangePackageItemList_ForSteps
        calls:
          - name: OpcenterEXDS_Core.SUB_DeleteCommand_StartArray
      - name: OpcenterEXDS_ProductionCoordination.ACT_RemoveChangePackageItemWorkOrderStep
        calls:
          - name: OpcenterEXDS_Core.SUB_DeleteCommand_TwoParameter
      - name: OpcenterEXDS_ProductionCoordination.ACT_RevertChangePackageItemList_ForMachine
        calls:
          - name: OpcenterEXDS_Core.SUB_DeleteCommand_StartArray
      - name: OpcenterEXDS_ProductionCoordination.ACT_RemoveChangePackageItemMachine
        calls:
          - name: OpcenterEXDS_Core.SUB_DeleteCommand_TwoParameter
      - name: OpcenterEXDS_ProductionCoordination.ACT_RevertChangePackageItemList_ForQI
        calls:
          - name: OpcenterEXDS_Core.SUB_DeleteCommand_StartArray
      - name: OpcenterEXDS_ProductionCoordination.ACT_RemoveChangePackageItemQualityInspection
        calls:
          - name: OpcenterEXDS_Core.SUB_DeleteCommand_TwoParameter
    target_commands:
      - AppName.CommandName

  - name: PANEL_AcceptChangePackage
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UADMAcceptChangePackage
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.AcceptChangePackage
    target_commands:
      - AppU4DM.UADMAcceptChangePackage

  - name: PANEL_RejectChangePackage
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_RejectChangePackage
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.RejectChangePackage
    target_commands:
      - AppU4DM.RejectChangePackage

  - name: ChangePackages_Master
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: IntegrationEvents_Master
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_ResendFailedEvent
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.ResendConfirmWorkOrderOperation
    target_commands:
      - APPU4DM.ResendConfirmWorkOrderOperation

  - name: PANEL_UADMLinkSkillsToTeam
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UADMLinkSkillsToTeam
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMLinkSkillsToTeam
    target_commands:
      - AppU4DM.UADMLinkSkillsToTeam

  - name: PANEL_UADMLinkUsersToTeam
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UADMLinkUsersToTeam
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMLinkUsersToTeam
    target_commands:
      - AppU4DM.UADMLinkUsersToTeam

  - name: Teams_Details
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UnLinkSkillsFromTeam_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UnLinkSkillsFromTeam
      - name: OpcenterEXDS_ProductionCoordination.ACT_UADMUnLinkUsersFromTeam_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMUnLinkUsersFromTeam
    target_commands:
      - AppU4DM.UnLinkSkillsFromTeam
      - AppU4DM.UADMUnLinkUsersFromTeam

  - name: PANEL_UpdateTeam
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UpdateTeam
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UpdateTeam
    target_commands:
      - AppU4DM.UpdateTeam

  - name: PANEL_ExportTeams
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_CreateTeam
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateTeam
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateTeam
    target_commands:
      - AppU4DM.CreateTeam

  - name: Teams_Master
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_DeleteTeam_WithConfirmation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.DeleteTeam
    target_commands:
      - AppU4DM.DeleteTeam

  - name: SPCEvaluation_Details
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CalculatedSPCEvaluationResult
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_CalculatedSPCEvaluationResult_MF
            - name: OpcenterEXDS_ProductionCoordination_Connector.QualityCharacteristicSPCEvaluation
            - name: OpcenterEXDS_ProductionCoordination_Connector.CalculatedSPCEvaluationResult
    target_commands:
      - WorkInstruction.QualityCharacteristicSPCEvaluation
      - OpcenterEXDS_ProductionCoordination_Connector.FromBlobToJson

  - name: SPCEvaluation_Master
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PCTimeUpdate_Master
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_CreateWorkOrderHistory
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateWorkOrderHistory
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_CreateWorkOrderHistory_MF
            - name: OpcenterEXDS_ProductionCoordination_Connector.CreateWorkOrderHistory
    target_commands:
      - AppU4DM.CreateWorkOrderHistory

  - name: PANEL_UpdateWorkOrderHistory
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UpdateWorkOrderHistory
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UpdateWorkOrderHistory
    target_commands:
      - AppU4DM.UpdateWorkOrderHistory

  - name: PANEL_UpdateWorkOrderOperationActualExecutionTimes
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UpdateWorkOrderOperationActualExecutionTimes
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UpdateWorkOrderOperationActualExecutionTimes
    target_commands:
      - AppU4DM.UpdateWorkOrderOperationActualExecutionTimes

  - name: PCTimeUpdate_Details
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_DeleteWorkOrderHistory
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.DeleteWorkOrderHistory
    target_commands:
      - AppU4DM.DeleteWorkOrderHistory

  - name: WorkOrderRouting_Details
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: WorkOrderNetwork_Master
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: PANEL_CreateWorkOrderDependency
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_CreateWorkOrderDependency
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.CreateWorkOrderDependency
    target_commands:
      - AppU4DM.CreateWorkOrderDependency

  - name: PANEL_CreateChangeNonConformance
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UADMCreateChangeNonConformance
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMCreateChangeNonConformance
    target_commands:
      - AppU4DM.UADMCreateChangeNonConformance

  - name: PANEL_UADMAcceptChangeAddOperation
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UADMAcceptChangeAddOperation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMAcceptChangeAddOperation
    target_commands:
      - AppU4DM.UADMAcceptChangeAddOperation

  - name: PANEL_UADMAcceptChangeAddWorkInstruction
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UADMAcceptChangeAddWorkInstruction
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMAcceptChangeAddWorkInstruction
    target_commands:
      - AppU4DM.UADMAcceptChangeAddWorkInstruction

  - name: PANEL_UADMAcceptChangeRemoveDependency
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UADMAcceptChangeRemoveDependency
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMAcceptChangeRemoveDependency
    target_commands:
      - AppU4DM.UADMAcceptChangeRemoveDependency

  - name: PANEL_UADMAcceptChangeReplaceToBeConsumedMaterial
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UADMAcceptChangeReplaceToBeConsumedMaterial
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMAcceptChangeReplaceToBeConsumedMaterial
    target_commands:
      - AppU4DM.UADMAcceptChangeReplaceToBeConsumedMaterial

  - name: PANEL_UADMAcceptChangeAddToBeConsumedMaterial
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UADMAcceptChangeAddToBeConsumedMaterial
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMAcceptChangeAddToBeConsumedMaterial
    target_commands:
      - AppU4DM.UADMAcceptChangeAddToBeConsumedMaterial

  - name: PANEL_UADMAcceptChangeRemoveOperation
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UADMAcceptChangeRemoveOperation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMAcceptChangeRemoveOperation
    target_commands:
      - AppU4DM.UADMAcceptChangeRemoveOperation

  - name: PANEL_UADMAcceptChangeChangeDependency
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UADMAcceptChangeChangeDependency
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMAcceptChangeChangeDependency
    target_commands:
      - AppU4DM.UADMAcceptChangeChangeDependency

  - name: PANEL_UADMAcceptChangeChangeToBeConsumedMaterialQuantity
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UADMAcceptChangeChangeToBeConsumedMaterialQuantity
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMAcceptChangeChangeToBeConsumedMaterialQuantity
    target_commands:
      - AppU4DM.UADMAcceptChangeChangeToBeConsumedMaterialQuantity

  - name: PANEL_UADMAcceptChangeRepeatOperation
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UADMAcceptChangeRepeatOperation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMAcceptChangeRepeatOperation
    target_commands:
      - AppU4DM.UADMAcceptChangeRepeatOperation

  - name: PANEL_UADMAcceptChangeAddWorkInstruction_SN
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UADMAcceptChangeAddWorkInstruction
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMAcceptChangeAddWorkInstruction
    target_commands:
      - AppU4DM.UADMAcceptChangeAddWorkInstruction

  - name: PANEL_UADMAcceptChangeAddToBeUsedTool
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UADMAcceptChangeAddToBeUsedTool
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMAcceptChangeAddToBeUsedTool
    target_commands:
      - AppU4DM.UADMAcceptChangeAddToBeUsedTool

  - name: PANEL_UADMAcceptChangeRemoveToBeConsumedMaterial
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UADMAcceptChangeRemoveToBeConsumedMaterial
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMAcceptChangeRemoveToBeConsumedMaterial
    target_commands:
      - AppU4DM.UADMAcceptChangeRemoveToBeConsumedMaterial

  - name: PANEL_AcceptChangeAddProcessOperation
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_GetProcessesFromAsPlanned
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.GetProductionProcessFromAsPlanned
      - name: OpcenterEXDS_ProductionCoordination.ACT_AcceptChangeAddProcessOperation
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.AcceptChangeAddProcessOperation
    target_commands:
      - AppU4DM.GetProductionProcessFromAsPlanned
      - AppU4DM.AcceptChangeAddProcessOperation

  - name: PANEL_UADMRejectChangeNonConformance
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_UADMRejectChangeNonConformance
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMRejectChangeNonConformance
    target_commands:
      - AppU4DM.UADMRejectChangeNonConformance

  - name: ChangeRequest_Details
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: ChangeRequest_Master
    module: OpcenterEXDS_ProductionCoordination
    flows:
      []
    target_commands:
      []

  - name: ContainmentRequestMonitoring_Master
    module: OpcenterEXDS_ProductionCoordination
    flows:
      - name: OpcenterEXDS_ProductionCoordination.ACT_DeleteImportedContainmentRequestResult
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.DeleteImportedContainmentRequestResult
      - name: OpcenterEXDS_ProductionCoordination.ACT_ImportContainmentRequest
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.GetContainmentReleaseBySequence
          - name: OpcenterEXDS_ProductionCoordination_Connector.GetContainmentRequestBySequence
    target_commands:
      - AppU4DM.DeleteImportedContainmentResults
      - AppU4DM.GetContainmentReleaseBySequence
      - AppU4DM.GetContainmentRequestBySequence

  - name: Genealogy_MasterV2
    module: OpcenterEXDS_ShopfloorExecution
    flows:
      - name: OpcenterEXDS_ShopfloorExecution.DS_GetNodes
        calls:
          - name: OpcenterEXDS_ShopfloorExecution.ACT_GetChildNodes
            - name: OpcenterEXDS_ShopfloorExecution.ACT_LoadStepChilds
            - name: OpcenterEXDS_ShopfloorExecution.ACT_LoadWOOChilds
              - name: OpcenterEXDS_ShopfloorExecution.ACT_LoadWOOSteps
                - name: OpcenterEXDS_ShopfloorExecution.ACT_SplitNodeId
            - name: OpcenterEXDS_ShopfloorExecution.ACT_SplitNodeId
          - name: OpcenterEXDS_ShopfloorExecution.ACT_GetImportedGenealogyChildNodes
            - name: OpcenterEXDS_ShopfloorExecution.ACT_LoadStepChild_Import
              - name: OpcenterEXDS_ShopfloorExecution.ACT_LoadStepDataCollectionWOOp
                - name: OpcenterEXDS_ShopfloorExecution.ACT_SplitNodeId
              - name: OpcenterEXDS_ShopfloorExecution.ACT_LoadStepActionsWOOp
                - name: OpcenterEXDS_ShopfloorExecution.ACT_SplitNodeId
            - name: OpcenterEXDS_ShopfloorExecution.ACT_LoadWOOpChild_Import
              - name: OpcenterEXDS_ShopfloorExecution.ACT_LoadWOOpSteps
                - name: OpcenterEXDS_ShopfloorExecution.ACT_SplitNodeId
              - name: OpcenterEXDS_ShopfloorExecution.ACT_LoadWOOpDataCollection
                - name: OpcenterEXDS_ShopfloorExecution.ACT_SplitNodeId
              - name: OpcenterEXDS_ShopfloorExecution.ACT_LoadWOOpActions
                - name: OpcenterEXDS_ShopfloorExecution.ACT_SplitNodeId
              - name: OpcenterEXDS_ShopfloorExecution.ACT_LoadWOOpNonConformances
                - name: OpcenterEXDS_ShopfloorExecution.ACT_SplitNodeId
            - name: OpcenterEXDS_ShopfloorExecution.ACT_SplitNodeId
      - name: OpcenterEXDS_ShopfloorExecution.ACT_ExportGenealogy
        calls:
          - name: OpcenterEXDS_ShopfloorExecution_Connector.UADMExportGenealogy
      - name: OpcenterEXDS_ShopfloorExecution.ACT_UnlinkGenealogy
        calls:
          - name: OpcenterEXDS_ShopfloorExecution_Connector.UnlinkCollectedDocumentList
          - name: OpcenterEXDS_ShopfloorExecution_Connector.UnlinkDocumentListToDM_MaterialTrackingUnit
    target_commands:
      - CommunityCommons.RandomHash
      - CommunityCommons.StringSplit
      - AppU4DM.UADMExportGenealogy
      - AppU4DM.UnlinkCollectedDocumentList
      - AppU4DM.UnlinkDocumentListToDM_MaterialTrackingUnit

  - name: PANEL_AddDocumentGenealogy
    module: OpcenterEXDS_ShopfloorExecution
    flows:
      - name: OpcenterEXDS_ShopfloorExecution.ACT_AddDocumentGenealogy
        calls:
          - name: OpcenterEXDS_ShopfloorExecution_Connector.UADMCreateDocument
    target_commands:
      - AppU4DM.UADMCreateDocument

  - name: PANEL_ImportGenealogy
    module: OpcenterEXDS_ShopfloorExecution
    flows:
      []
    target_commands:
      []

  - name: PANEL_DisassembleMaterialItem
    module: OpcenterEXDS_ShopfloorExecution
    flows:
      - name: OpcenterEXDS_ShopfloorExecution.ACT_DisassembleMaterialItem
        calls:
          - name: OpcenterEXDS_ProductionCoordination.ACT_DisassembleMaterialItem_Command
            - name: OpcenterEXDS_ProductionCoordination_Connector.DisassembleMaterialItem
    target_commands:
      - AppU4DM.DisassembleMaterialItem

  - name: NonConformance_Master
    module: OpcenterEXDS_ShopfloorExecution
    flows:
      []
    target_commands:
      []

  - name: PANEL_UADMSentenceNonConformanceV3_1
    module: OpcenterEXDS_ShopfloorExecution
    flows:
      - name: OpcenterEXDS_ShopfloorExecution.DS_GetNcStatusToWithRoles
      - name: OpcenterEXDS_ShopfloorExecution.ACT_UADMSentenceNonConformanceV3_1
        calls:
          - name: OpcenterEXDS_ShopfloorExecution_Connector.UADMSentenceNonConformanceV3_1
    target_commands:
      - AppU4DM.GetCurrentUserRoles
      - AppU4DM.UADMSentenceNonConformanceV3_1

  - name: PANEL_UADMLinkGenericItemsToNonConformance
    module: OpcenterEXDS_ShopfloorExecution
    flows:
      - name: OpcenterEXDS_ShopfloorExecution.ACT_SelectionItemsAttachmentOnExistingNC
        calls:
          - name: OpcenterEXDS_ShopfloorExecution.ACT_UADMLinkGenericItemsToNonConformance
            - name: OpcenterEXDS_ShopfloorExecution_Connector.UADMLinkGenericItemsToNonConformance
    target_commands:
      - AppU4DM.UADMLinkGenericItemsToNonConformance

  - name: Panel_UADMCreateDocument
    module: OpcenterEXDS_ShopfloorExecution
    flows:
      - name: OpcenterEXDS_ShopfloorExecution.ACT_UADMCreateDocument
        calls:
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMCreateDocument
          - name: OpcenterEXDS_ShopfloorExecution.ACT_CreateNonConformanceAttachment
            - name: OpcenterEXDS_ProductionCoordination_Connector.CreateNonConformanceAttachmentList
    target_commands:
      - AppU4DM.UADMCreateDocument
      - AppU4DM.CreateNonConformanceAttachmentList

  - name: PANEL_DeclareNonConformance
    module: OpcenterEXDS_ShopfloorExecution
    flows:
      - name: OpcenterEXDS_ShopfloorExecution.ACT_SelectionItemsAttachmentOnCreateNC
        calls:
          - name: OpcenterEXDS_ShopfloorExecution.ACT_UADMCreateNonConformanceV3_1
            - name: OpcenterEXDS_ShopfloorExecution_Connector.UADMCreateNonConformanceV3_1
    target_commands:
      - AppU4DM.UADMCreateNonConformanceV3_1

  - name: PANEL_UADMCreateNonConformanceDefectList
    module: OpcenterEXDS_ShopfloorExecution
    flows:
      - name: OpcenterEXDS_ShopfloorExecution.ACT_CreateDefectListWIthSelection
        calls:
          - name: OpcenterEXDS_ShopfloorExecution.ACT_UADMCreateNonConformanceDefectList
            - name: OpcenterEXDS_ShopfloorExecution_Connector.UADMCreateNonConformanceDefectList
    target_commands:
      - AppU4DM.UADMCreateNonConformanceDefectList

  - name: PANEL_CreateNonConformanceAttachmentList
    module: OpcenterEXDS_ShopfloorExecution
    flows:
      - name: OpcenterEXDS_ShopfloorExecution.ACT_SelectionDocumentsAttachmentOnCreateNC
        calls:
          - name: OpcenterEXDS_ShopfloorExecution.ACT_CreateNonConformanceAttachmentList
            - name: OpcenterEXDS_ShopfloorExecution_Connector.CreateNonConformanceAttachmentList
    target_commands:
      - AppU4DM.CreateNonConformanceAttachmentList

  - name: PANEL_AddNCNewDocument
    module: OpcenterEXDS_ShopfloorExecution
    flows:
      - name: OpcenterEXDS_ShopfloorExecution.ACT_AddNCNewDocument
        calls:
          - name: OpcenterEXFN_MasterData_Connector.CreateDocument
          - name: OpcenterEXDS_ShopfloorExecution.ACT_CreateNonConformanceAttachmentList
            - name: OpcenterEXDS_ShopfloorExecution_Connector.CreateNonConformanceAttachmentList
    target_commands:
      - Document.CreateDocument
      - AppU4DM.CreateNonConformanceAttachmentList

  - name: PANEL_UADMUpdateNonConformance
    module: OpcenterEXDS_ShopfloorExecution
    flows:
      - name: OpcenterEXDS_ShopfloorExecution.ACT_UADMUpdateNonConformance
        calls:
          - name: OpcenterEXDS_ShopfloorExecution_Connector.UADMUpdateNonConformance
    target_commands:
      - AppU4DM.UADMUpdateNonConformance

  - name: NonConformance_Details
    module: OpcenterEXDS_ShopfloorExecution
    flows:
      - name: OpcenterEXDS_ShopfloorExecution.ACT_RemoveDefectsFromNonConformance_WithConfirmation
      - name: OpcenterEXDS_ShopfloorExecution.ACT_DeleteNonConformanceAttachmentList_WithConfirmation
      - name: OpcenterEXDS_ShopfloorExecution.ACT_UploadDocumentsToTeamcenterShare
        calls:
          - name: OpcenterEXDS_ShopfloorExecution_Connector.UploadDocuments
      - name: OpcenterEXDS_ShopfloorExecution.ACT_SelectionNCItemsDetachmentOnExistingNC
        calls:
          - name: OpcenterEXDS_ShopfloorExecution.ACT_UADMUnLinkGenericItemsFromNonConformance_WithConfirmation
            - name: OpcenterEXDS_ShopfloorExecution_Connector.UADMUnLinkGenericItemsFromNonConformance
      - name: OpcenterEXDS_ShopfloorExecution.ACT_SelectionMTUsDetachmentOnExistingNC
        calls:
          - name: OpcenterEXDS_ShopfloorExecution.ACT_UADMUnLinkGenericItemsFromNonConformance_WithConfirmation
            - name: OpcenterEXDS_ShopfloorExecution_Connector.UADMUnLinkGenericItemsFromNonConformance
    target_commands:
      - AppU4DM.RemoveDefectsFromNonConformance
      - AppU4DM.DeleteNonConformanceAttachmentList
      - AppU4DM.UploadDocuments
      - AppU4DM.UADMUnLinkGenericItemsFromNonConformance

  - name: Tool_Master
    module: OpcenterEXDS_ShopfloorExecution
    flows:
      - name: OpcenterEXDS_ShopfloorExecution.ACT_DeleteTool_WithConfirmation
    target_commands:
      - AppU4DM.DeleteTool

  - name: PANEL_DeclareToolNonConformance
    module: OpcenterEXDS_ShopfloorExecution
    flows:
      - name: OpcenterEXDS_ShopfloorExecution.ACT_UADMCreateNonConformanceV3_1
        calls:
          - name: OpcenterEXDS_ShopfloorExecution_Connector.UADMCreateNonConformanceV3_1
    target_commands:
      - AppU4DM.UADMCreateNonConformanceV3_1

  - name: PANEL_BrowseFailures
    module: OpcenterEXDS_ShopfloorExecution
    flows:
      []
    target_commands:
      []

  - name: PANEL_UpdateTool
    module: OpcenterEXDS_ShopfloorExecution
    flows:
      - name: OpcenterEXDS_ShopfloorExecution.ACT_UpdateTool
        calls:
          - name: OpcenterEXDS_ShopfloorExecution_Connector.UpdateTool
    target_commands:
      - AppU4DM.UpdateTool

  - name: PANEL_ToolMaintenance
    module: OpcenterEXDS_ShopfloorExecution
    flows:
      - name: OpcenterEXDS_ShopfloorExecution.ACT_ToolMaintenance
        calls:
          - name: OpcenterEXDS_ShopfloorExecution_Connector.ToolMaintenance
    target_commands:
      - AppU4DM.ToolMaintenance

  - name: PANEL_CreateNewLogisticClass
    module: OpcenterEXDS_ShopfloorExecution
    flows:
      - name: OpcenterEXDS_ShopfloorExecution.ACT_CreateLogisticClass
        calls:
          - name: OpcenterEXDS_ShopfloorExecution_Connector.CreateLogisticClass
    target_commands:
      - AppU4DM.CreateLogisticClass

  - name: PANEL_UADMCreateToolDefinition
    module: OpcenterEXDS_ShopfloorExecution
    flows:
      - name: OpcenterEXDS_ShopfloorExecution.ACT_UADMCreateToolDefinition
        calls:
          - name: OpcenterEXDS_ShopfloorExecution_Connector.UADMCreateToolDefinition
    target_commands:
      - AppU4DM.UADMCreateToolDefinition

  - name: PANEL_CreateTool
    module: OpcenterEXDS_ShopfloorExecution
    flows:
      - name: OpcenterEXDS_ShopfloorExecution.ACT_CreateTool
        calls:
          - name: OpcenterEXDS_ShopfloorExecution_Connector.CreateTool
    target_commands:
      - AppU4DM.CreateTool

  - name: PANEL_LinkToolToAutomationNodeInstanceList
    module: OpcenterEXDS_ShopfloorExecution
    flows:
      - name: OpcenterEXDS_ShopfloorExecution.ACT_LinkToolToAutomationNodeInstanceList
        calls:
          - name: OpcenterEXDS_ShopfloorExecution_Connector.LinkToolToAutomationNodeInstanceList
    target_commands:
      - AppU4DM.LinkToolToAutomationNodeInstanceList

  - name: PANEL_LinkSubstrateToMachine
    module: OpcenterEXDS_ShopfloorExecution
    flows:
      - name: OpcenterEXDS_ShopfloorExecution.ACT_LinkSubstrateToMachine
        calls:
          - name: OpcenterEXDS_ShopfloorExecution_Connector.LinkSubstrateToMachine
    target_commands:
      - AppU4DM.LinkSubstrateToMachine

  - name: ToolDetails_Details
    module: OpcenterEXDS_ShopfloorExecution
    flows:
      - name: OpcenterEXDS_ShopfloorExecution.ACT_UnlinkSubstrateToMachine_WithConfirmation
      - name: OpcenterEXDS_ShopfloorExecution.ACT_UnlinkToolToAutomationNodeInstanceList_WithConfirmation
        calls:
          - name: OpcenterEXDS_ShopfloorExecution_Connector.UnlinkToolToAutomationNodeInstanceList
    target_commands:
      - AppU4DM.UnlinkSubstrateToMachine
      - AppU4DM.UnlinkToolToAutomationNodeInstanceList

  - name: PANEL_UADMSetLocationHold
    module: OpcenterEXDS_Configuration
    flows:
      - name: OpcenterEXDS_Configuration.ACT_UADMSetLocationHold
        calls:
          - name: OpcenterEXDS_Configuration_Connector.UADMSetLocationHold
    target_commands:
      - AppU4DM.UADMSetLocationHold

  - name: PANEL_UADMSetWorkOrderHoldList
    module: OpcenterEXDS_Configuration
    flows:
      - name: OpcenterEXDS_Configuration.ACT_UADMSetWorkOrderHoldList
        calls:
          - name: OpcenterEXDS_Configuration_Connector.UADMSetWorkOrderHoldList
    target_commands:
      - AppU4DM.UADMSetWorkOrderHoldList

  - name: PANEL_UADMRemoveHoldList
    module: OpcenterEXDS_Configuration
    flows:
      - name: OpcenterEXDS_Configuration.ACT_UADMRemoveHoldList
        calls:
          - name: OpcenterEXDS_Configuration_Connector.UADMRemoveHoldList
    target_commands:
      - AppU4DM.UADMRemoveHoldList

  - name: Hold_Master
    module: OpcenterEXDS_Configuration
    flows:
      []
    target_commands:
      []

  - name: Hold_Details
    module: OpcenterEXDS_Configuration
    flows:
      []
    target_commands:
      []

  - name: PANEL_MTUSplit
    module: OpcenterEXDS_MasterData
    flows:
      - name: OpcenterEXDS_MasterData.ACT_MTUSplit
        calls:
          - name: OpcenterEXDS_MasterData_Connector.DSMaterial_MaterialItemSplitBatch
      - name: OpcenterEXDS_MasterData.ACT_MTUSplitPreview
        calls:
          - name: OpcenterEXDS_MasterData_Connector.DSMaterial_MaterialItemSplitBatchByPartialPreview
    target_commands:
      - Material.DSMaterial_MaterialItemSplitBatch
      - Material.DSMaterial_MaterialItemSplitBatchByPartialPreview

  - name: PANEL_SetMaterialTrackingUnitQuantity
    module: OpcenterEXDS_MasterData
    flows:
      - name: OpcenterEXDS_MasterData.ACT_SetMaterialTrackingUnitQuantity
        calls:
          - name: OpcenterEXFN_MasterData_Connector.SetMaterialTrackingUnitQuantity
    target_commands:
      - Material.SetMaterialTrackingUnitQuantity

  - name: PANEL_AssociateMaterialTrackingUnitsWithMaterialLot
    module: OpcenterEXDS_MasterData
    flows:
      - name: OpcenterEXDS_MasterData.ACT_AssociateMaterialTrackingUnitsWithMaterialLot
        calls:
          - name: OpcenterEXFN_MasterData_Connector.AssociateMaterialTrackingUnitsWithMaterialLot
    target_commands:
      - Material.AssociateMaterialTrackingUnitsWithMaterialLot

  - name: PANEL_MoveMaterialTrackingUnitToEquipment
    module: OpcenterEXDS_MasterData
    flows:
      - name: OpcenterEXDS_MasterData.ACT_MoveMaterialTrackingUnitToEquipment
        calls:
          - name: OpcenterEXFN_MasterData_Connector.MoveMaterialTrackingUnitToEquipment
    target_commands:
      - Material.MoveMaterialTrackingUnitToEquipment

  - name: PANEL_MoveMaterialTrackingUnitToMaterialTrackingUnitAggregate
    module: OpcenterEXDS_MasterData
    flows:
      - name: OpcenterEXDS_MasterData.ACT_MoveMaterialTrackingUnitToMaterialTrackingUnitAggregate
        calls:
          - name: OpcenterEXFN_MasterData_Connector.MoveMaterialTrackingUnitToMaterialTrackingUnitAggregate
    target_commands:
      - Material.MoveMaterialTrackingUnitToMaterialTrackingUnitAggregate

  - name: PANEL_SetMaterialTrackingUnitStatus
    module: OpcenterEXDS_MasterData
    flows:
      - name: OpcenterEXDS_MasterData.ACT_SetMaterialTrackingUnitStatus
        calls:
          - name: OpcenterEXFN_MasterData_Connector.SetMaterialTrackingUnitMaterial
    target_commands:
      - Material.SetMaterialTrackingUnitMaterial

  - name: PANEL_CreateMaterialTrackingUnit
    module: OpcenterEXDS_MasterData
    flows:
      - name: OpcenterEXDS_MasterData.ACT_CreateMaterialTrackingUnit
        calls:
          - name: OpcenterEXDS_MasterData_Connector.DSMaterial_CreateDM_MaterialTrackingUnitandMaterialTrackingUnit
    target_commands:
      - Material.DSMaterial_CreateDM_MaterialTrackingUnitandMaterialTrackingUnit

  - name: PANEL_EditMaterialTrackingUnit
    module: OpcenterEXDS_MasterData
    flows:
      - name: OpcenterEXDS_MasterData.ACT_EditMaterialTrackingUnit
        calls:
          - name: OpcenterEXDS_MasterData_Connector.DSMaterial_UpdateMaterialTrackingUnitAndDM_MaterialTrackingUnit
    target_commands:
      - Material.DSMaterial_UpdateMaterialTrackingUnitAndDM_MaterialTrackingUnit

  - name: PANEL_SetStatusMaterialTrackingUnit
    module: OpcenterEXDS_MasterData
    flows:
      - name: OpcenterEXDS_MasterData.ACT_SetStatusMaterialTrackingUnit
        calls:
          - name: OpcenterEXFN_MasterData_Connector.SetMaterialTrackingUnitStatus
    target_commands:
      - Material.SetMaterialTrackingUnitStatus

  - name: PANEL_SetMaterialTrackingUnitStateMachine
    module: OpcenterEXDS_MasterData
    flows:
      - name: OpcenterEXDS_MasterData.ACT_SetMaterialTrackingUnitStateMachine
        calls:
          - name: OpcenterEXFN_MasterData_Connector.SetStateMachineMaterialTrackingUnit
    target_commands:
      - Material.SetMaterialTrackingUnitStateMachine

  - name: PANEL_LinkNewDocument
    module: OpcenterEXDS_MasterData
    flows:
      - name: OpcenterEXDS_MasterData.ACT_LinkNewDocument
        calls:
          - name: OpcenterEXDS_MasterData_Connector.DSMaterial_LinkDocumentListToDM_MaterialTrackingUnit
          - name: OpcenterEXDS_ProductionCoordination_Connector.UADMCreateDocument
    target_commands:
      - Material.DSMaterial_LinkDocumentListToDM_MaterialTrackingUnit
      - AppU4DM.UADMCreateDocument

  - name: PANEL_LinkExistingDocument
    module: OpcenterEXDS_MasterData
    flows:
      - name: OpcenterEXDS_MasterData.ACT_LinkExistingDocument
        calls:
          - name: OpcenterEXDS_MasterData_Connector.DSMaterial_LinkDocumentListToDM_MaterialTrackingUnit
    target_commands:
      - Material.DSMaterial_LinkDocumentListToDM_MaterialTrackingUnit

  - name: PANEL_ResultAmend
    module: OpcenterEXDS_MasterData
    flows:
      - name: OpcenterEXDS_MasterData.ACT_ResultAmend
        calls:
          - name: OpcenterEXDS_MasterData_Connector.DSResult_AmendResultList
    target_commands:
      - Material.DSResult_AmendResultList

  - name: PANEL_ResultHistory
    module: OpcenterEXDS_MasterData
    flows:
      []
    target_commands:
      []

  - name: MaterialTrackingUnit_Details
    module: OpcenterEXDS_MasterData
    flows:
      - name: OpcenterEXDS_MasterData.ACT_UnlinkCollectedDocumentList_WithConfirmation
        calls:
          - name: OpcenterEXDS_MasterData_Connector.DSMaterial_UnlinkDocumentListToDM_MaterialTrackingUnit
      - name: OpcenterEXDS_MasterData.ACT_FreezeMaterialTrackingUnit
        calls:
          - name: OpcenterEXFN_MasterData_Connector.SUB_FreezeAndUnfreezeCommand
      - name: OpcenterEXDS_MasterData.ACT_DisassociateMaterialTrackingUnitsFromMaterialLot
        calls:
          - name: OpcenterEXFN_MasterData_Connector.SUB_DeleteCommand_StartArrayWithTwoParameter
      - name: OpcenterEXDS_MasterData.ACT_UnfreezeMaterialTrackingUnit
        calls:
          - name: OpcenterEXFN_MasterData_Connector.SUB_FreezeAndUnfreezeCommand
    target_commands:
      - Material.DSMaterial_UnlinkDocumentListToDM_MaterialTrackingUnit
      - AppName.CommandName

  - name: MaterialTrackingUnit_Master
    module: OpcenterEXDS_MasterData
    flows:
      - name: OpcenterEXDS_MasterData.ACT_FreezeMaterialTrackingUnit
        calls:
          - name: OpcenterEXFN_MasterData_Connector.SUB_FreezeAndUnfreezeCommand
      - name: OpcenterEXDS_MasterData.ACT_DisassociateMaterialTrackingUnitsFromMaterialLot
        calls:
          - name: OpcenterEXFN_MasterData_Connector.SUB_DeleteCommand_StartArrayWithTwoParameter
      - name: OpcenterEXDS_MasterData.ACT_UnfreezeMaterialTrackingUnit
        calls:
          - name: OpcenterEXFN_MasterData_Connector.SUB_FreezeAndUnfreezeCommand
      - name: OpcenterEXDS_MasterData.ACT_DeleteMaterialTrackingUnit
        calls:
          - name: OpcenterEXFN_MasterData_Connector.SUB_DeleteCommand
      - name: OpcenterEXDS_MasterData.ACT_GenerateAndAssociateMTUCode
        calls:
          - name: OpcenterEXDS_MasterData_Connector.DSMaterial_UADMGenerateAndAssociateMTUCode
    target_commands:
      - AppName.CommandName
      - AppU4DM.DSMaterial_UADMGenerateAndAssociateMTUCode

  - name: PANEL_CreateMaterialTrackingUnitFromMaterialLot
    module: OpcenterEXDS_MasterData
    flows:
      - name: OpcenterEXDS_MasterData.ACT_CreateMaterialTrackingUnitForMaterialLot
        calls:
          - name: OpcenterEXDS_MasterData_Connector.CreateMTUAndDMMTUFromLotWrapper
            - name: OpcenterEXFN_MasterData_Connector.CreateMaterialTrackingUnit
            - name: OpcenterEXDS_MasterData_Connector.DSMaterial_UADMCreateDM_MTU
    target_commands:
      - Material.CreateMaterialTrackingUnit
      - Material.DSMaterial_UADMCreateDM_MTU

```

---

## 7. PageCommands

Simplified view showing only the target commands for each page/panel.

| Page/Panel | Module | Target AppName | Target CommandName |
|------------|--------|----------------|--------------------|
| PANEL_CreateHandlingUnit | OpcenterEXDS_ProductionCoordination | AppU4DM | CreateHandlingUnit |
| HandlingUnits_Master | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_UnLinkMaterialItemsFromHandlingUnit | OpcenterEXDS_ProductionCoordination | AppU4DM | UnLinkMaterialItemsFromHandlingUnit |
| PANEL_UnLoadHandlingUnit | OpcenterEXDS_ProductionCoordination | AppU4DM | UnLoadHandlingUnit |
| PANEL_UADMLinkMaterialItemsToHandlingUnit | OpcenterEXDS_ProductionCoordination | AppU4DM | UADMLinkMaterialItemsToHandlingUnit |
| PANEL_ChangeHandlingUnitStatus | OpcenterEXDS_ProductionCoordination | AppU4DM | ChangeHandlingUnitStatus |
| HandlingUnit_Details | OpcenterEXDS_ProductionCoordination | - | - |
| LogisticRequest_Master | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_RejectMaterialRequest | OpcenterEXDS_ProductionCoordination | AppU4DM | RejectMaterialRequest |
| PANEL_AddBuffer | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_AcceptMaterialRequest | OpcenterEXDS_ProductionCoordination | AppU4DM | AcceptMaterialRequest |
| PANEL_UpdateERPOrder | OpcenterEXDS_ProductionCoordination | ERPOrder<br>ERPOrder | UpdateERPOrder<br>ERPOrderBoP_SetBoPInfoToERPOrder |
| PANEL_CreateERPOrder | OpcenterEXDS_ProductionCoordination | ERPOrder<br>ERPOrder | CreateERPOrder<br>ERPOrderBoP_CreateBoPERPOrder |
| ERPOrder_Master | OpcenterEXDS_ProductionCoordination | ERPOrder<br>AppName | SetERPOrderListStatus<br>CommandName |
| MaterialSuppliers_Details | OpcenterEXDS_ProductionCoordination | - | - |
| ERPOrder_Details | OpcenterEXDS_ProductionCoordination | AppName | CommandName |
| Measurements_Master | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_CreateMeasurements | OpcenterEXDS_ProductionCoordination | AppU4DM | TraceProcessMeasurements |
| PANEL_UADMSetWorkOrderOperationFutureHoldList | OpcenterEXDS_ProductionCoordination | AppU4DM | UADMSetWorkOrderOperationFutureHoldList |
| PANEL_UADMRemoveWorkOrderOperationFutureHoldList | OpcenterEXDS_ProductionCoordination | AppU4DM | UADMRemoveWorkOrderOperationFutureHoldList |
| FutureHoldManagement_Master | OpcenterEXDS_ProductionCoordination | - | - |
| FutureHoldManagement_Details | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_UADMLinkMaterialItemsToBuffer | OpcenterEXDS_ProductionCoordination | AppU4DM | UADMLinkMaterialItemsToBuffer |
| Buffers_Details | OpcenterEXDS_ProductionCoordination | AppU4DM<br>AppU4DM | UnLinkMaterialItemFromBuffer<br>DeleteBuffer |
| PANEL_ChangeBufferStatus | OpcenterEXDS_ProductionCoordination | AppU4DM | ChangeBufferStatus |
| PANEL_CreateBuffer | OpcenterEXDS_ProductionCoordination | AppU4DM | CreateBuffer |
| PANEL_UpdateBuffer | OpcenterEXDS_ProductionCoordination | AppU4DM | UpdateBuffer |
| Buffers_Master | OpcenterEXDS_ProductionCoordination | AppU4DM | DeleteBuffer |
| WorkOrderUpdate_Master | OpcenterEXDS_ProductionCoordination | - | - |
| Process_Details | OpcenterEXDS_ProductionCoordination | AppU4DM | UADMUpdateWOByBoPList |
| WorkOrderUpdate_Details | OpcenterEXDS_ProductionCoordination | AppU4DM | UADMUpdateWOByBoPList |
| PANEL_SplitWorkOrder | OpcenterEXDS_ProductionCoordination | AppU4DM<br>AppU4DM | AutoGenerateWorkOrderNId<br>UADMSplitWorkOrder |
| PANEL_AssociateSerialNumbersToSplitWo | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_CreateWOOFolderAndWOOpDependencies | OpcenterEXDS_ProductionCoordination | AppU4DM | CreateWOOFolderAndWOOpDependencies |
| PANEL_CloseFlexibleWorkOrder | OpcenterEXDS_ProductionCoordination | AppU4DM | UADMAbruptlyCloseFlexibleWorkOrder |
| PANEL_CreateWorkOrder_FromMasterPlan_Step1 | OpcenterEXDS_ProductionCoordination | AppU4DM<br>AppU4DM<br>AppU4DM | AutoGenerateWorkOrderNId<br>AutoGenerateMTUCode<br>PLMCreateWorkOrderFromMasterPlanBOMResolution |
| PANEL_CreateWorkOrder_FromMasterPlan_Step2 | OpcenterEXDS_ProductionCoordination | AppU4DM | PLMCreateWorkOrderFromMasterPlanBOMResolution |
| PANEL_CreateWorkOrder_FromMasterPlanWithQC_Step1 | OpcenterEXDS_ProductionCoordination | AppU4DM<br>AppU4DM<br>AppU4DM | CreateWorkOrderFromMasterPlanWithQC<br>AutoGenerateWorkOrderNId<br>AutoGenerateMTUCode |
| PANEL_CreateWorkOrder_FromMasterPlanWithQC_Step2 | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_CreateWorkOrder_FromProcess | OpcenterEXDS_ProductionCoordination | AppU4DM<br>AppU4DM<br>AppU4DM | UADMCreateWorkOrderFromProcess<br>AutoGenerateWorkOrderNId<br>AutoGenerateMTUCode |
| PANEL_CreateWorkOrder_Manually | OpcenterEXDS_ProductionCoordination | AppU4DM | CreateWorkOrder |
| PANEL_CreateWorkOrder_AsPlanned | OpcenterEXDS_ProductionCoordination | AppU4DM | UADMCreateWorkOrdersFromAsPlannedBOP |
| PANEL_CreateWorkOrder_Header | OpcenterEXDS_ProductionCoordination | AppU4DM<br>AppU4DM<br>AppU4DM | CreateWorkOrderHeader<br>AutoGenerateWorkOrderNId<br>AutoGenerateMTUCode |
| PANEL_CreateWorkOrder_FromMasterPlanWithEffectivity | OpcenterEXDS_ProductionCoordination | AppU4DM<br>AppU4DM<br>AppU4DM | AutoGenerateWorkOrderNId<br>AutoGenerateMTUCode<br>CreateWorkOrderFromMasterPlanWithEffectivity |
| PANEL_MergeWOHeaderWithProcess | OpcenterEXDS_ProductionCoordination | AppU4DM | MergeWOHeaderWithProcess |
| PANEL_EditWorkOrder | OpcenterEXDS_ProductionCoordination | AppU4DM | EditWorkOrder |
| WorkOrder_Master | OpcenterEXDS_ProductionCoordination | AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM | ChangeWorkOrderStatusToEdit<br>UADMSetWorkOrderForScheduling<br>TriggerPrintingOnWorkOrder<br>UADMReleaseWorkOrder<br>SetTargetQuantityOnFlexibleWorkOrder<br>UADMAbortWorkOrder<br>DeleteWorkOrder<br>MarkForCleaningWorkOrderList<br>CreateWorkOrderOutMsg |
| PANEL_AssociateSerialNumbers | OpcenterEXDS_ProductionCoordination | AppU4DM<br>AppU4DM | UADMCreateAndAssignProducedMaterialItems<br>AssignProducedMaterialItems |
| PANEL_EditAlternativeOperationGroup | OpcenterEXDS_ProductionCoordination | AppU4DM | UpdatePreferredWOOpDependencyNavigation |
| PANEL_AddAlternativeOperationGroup | OpcenterEXDS_ProductionCoordination | AppU4DM | CreateWOOpDependencyNavigationList |
| PANEL_LinkWIDefinitionsToWOOperation | OpcenterEXDS_ProductionCoordination | AppU4DM | LinkWIDefinitionsToWOOperation |
| PANEL_LinkWIDefinitionsToWOStep | OpcenterEXDS_ProductionCoordination | AppU4DM | LinkWIDefinitionsToWOStep |
| PANEL_EditWorkOrderHumanResource_Step | OpcenterEXDS_ProductionCoordination | AppU4DM | EditWorkOrderHumanResource |
| PANEL_LinkWorkOrderHumanResource_Step | OpcenterEXDS_ProductionCoordination | AppU4DM | LinkWorkOrderHumanResource |
| PANEL_CreateWOStepDependency_FromOpDetails | OpcenterEXDS_ProductionCoordination | AppU4DM | CreateWOStepDependency |
| PANEL_CreateWOStepDependency | OpcenterEXDS_ProductionCoordination | AppU4DM | CreateWOStepDependency |
| PANEL_LinkSkillsToWorkOrderStep | OpcenterEXDS_ProductionCoordination | AppU4DM | LinkSkillsToWorkOrderStep |
| PANEL_CreateToBeConsumedMaterials_Step | OpcenterEXDS_ProductionCoordination | AppU4DM | CreateToBeConsumedMaterials |
| PANEL_CreateRuntimeCharacteristicRepresentationContainer_Step | OpcenterEXDS_ProductionCoordination | AppU4DM | CreateRuntimeCharacteristicRepresentationContainer |
| PANEL_EditWorkOrderStep | OpcenterEXDS_ProductionCoordination | AppU4DM | UpdateWorkOrderStep |
| PANEL_LinkStepToOperation | OpcenterEXDS_ProductionCoordination | AppU4DM | CreateWorkOrderStepFromStep |
| PANEL_CreateWorkOrderStepsManually | OpcenterEXDS_ProductionCoordination | AppU4DM | CreateWorkOrderSteps |
| PANEL_AddTool_Step | OpcenterEXDS_ProductionCoordination | AppU4DM | CreateToBeUsedTools |
| PANEL_EditInterlockingChecks_Step | OpcenterEXDS_ProductionCoordination | AppU4DM | UpdateItlkCheckWOStepAssociationList |
| PANEL_LinkItlkCheckToWorkOrderStep | OpcenterEXDS_ProductionCoordination | AppU4DM | LinkItlkCheckToWorkOrderStepList |
| WorkOrderOperationSteps_Details | OpcenterEXDS_ProductionCoordination | AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM | DeleteWorkOStepDependency<br>DeleteToBeConsumedMaterial<br>DeleteToBeUsedTool<br>UnlinkWorkInstructionsFromWOStep<br>DeleteToBeUsedInspection<br>UnlinkWorkOrderHumanResource<br>UnlinkSkillsToWorkOrderStep<br>UnlinkItlkCheckToWorkOrderStepList<br>TriggerPrintingOnWorkOrderStep |
| PANEL_LinkSkillsToWorkOrderOperation | OpcenterEXDS_ProductionCoordination | AppU4DM | LinkSkillsToWorkOrderOperation |
| PANEL_AddTool | OpcenterEXDS_ProductionCoordination | AppU4DM | CreateToBeUsedTools |
| PANEL_ShowScrewing | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_LinkUserToWorkOrderOperationList | OpcenterEXDS_ProductionCoordination | AppU4DM<br>AppU4DM | LinkUserToWorkOrderOperationList<br>GetUserDetailsList |
| PANEL_EditInterlockingChecks | OpcenterEXDS_ProductionCoordination | AppU4DM | UpdateItlkCheckWOOpAssociationList |
| PANEL_DetailsInterlockingChecks | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_LinkItlkCheckToWorkOrderOperationList | OpcenterEXDS_ProductionCoordination | AppU4DM | LinkItlkCheckToWorkOrderOperationList |
| PANEL_CreateToBeConsumedMaterials | OpcenterEXDS_ProductionCoordination | AppU4DM | CreateToBeConsumedMaterials |
| PANEL_CreateToBeUsedMachine_PartProgram | OpcenterEXDS_ProductionCoordination | AppU4DM | CreateToBeUsedMachine |
| PANEL_CreateToBeUsedMachine | OpcenterEXDS_ProductionCoordination | AppU4DM | CreateToBeUsedMachine |
| PANEL_LinkToBeUsedMachineToPartProgram | OpcenterEXDS_ProductionCoordination | AppU4DM | LinkToBeUsedMachineToPartProgram |
| PANEL_EditWorkOrderHumanResource | OpcenterEXDS_ProductionCoordination | AppU4DM | EditWorkOrderHumanResource |
| PANEL_LinkWorkOrderHumanResource | OpcenterEXDS_ProductionCoordination | AppU4DM | LinkWorkOrderHumanResource |
| PANEL_CreateWorkOrderOperation_FromProcess | OpcenterEXDS_ProductionCoordination | AppU4DM | UADMCreateWorkOrderOperationFromProcessOperation |
| PANEL_CreateWorkOrderOperation_FromMasterPlan | OpcenterEXDS_ProductionCoordination | AppU4DM | CreateWorkOrderOperationFromMPProcessOperation |
| PANEL_CreateWorkOrderOperation_Manually | OpcenterEXDS_ProductionCoordination | AppU4DM | CreateWorkOrderOperation |
| PANEL_EditWorkOrderOperation | OpcenterEXDS_ProductionCoordination | AppU4DM | EditWorkOrderOperation |
| PANEL_CharacteristicRepresentation | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_CreateRuntimeCharacteristicRepresentationContainer | OpcenterEXDS_ProductionCoordination | AppU4DM | CreateRuntimeCharacteristicRepresentationContainer |
| WorkOrderOperation_Details | OpcenterEXDS_ProductionCoordination | AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM | DeleteWOOFolderAndWOOpDependencies<br>UnLinkToBeUsedMachineToPartProgram<br>DeleteToBeUsedMachine<br>SetPreferredMachine<br>DeleteToBeConsumedMaterial<br>DeleteToBeUsedTool<br>UnlinkWorkInstructionsFromWOOperation<br>UnlinkSkillsFromWOOperation<br>DeleteWorkOrderStep<br>TriggerPrintingOnWorkOrderStep<br>DeleteToBeUsedInspection<br>UnlinkWorkOrderHumanResource<br>UnlinkUserToWorkOrderOperationList<br>UnlinkItlkCheckToWorkOrderOperationList<br>TriggerPrintingOnWorkOrderOperation<br>CreateWorkOrderOutMsg |
| ExecutionGroup_Details_2 | OpcenterEXDS_ProductionCoordination | AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppName | UnlinkWOOperationsFromExecutionGroup<br>ReleaseExecutionGroup<br>ScheduleExecutionGroup<br>SetExecGroupListInEditStatus<br>AbortExecutionGroupList<br>CommandName |
| ExecutionGroupPhase_Details_2 | OpcenterEXDS_ProductionCoordination | AppU4DM<br>PowderMgt | UnlinkWIDefinitionsFromEGPhase<br>RemoveAMPowderLoadedCheckFromEgPhaseList |
| PANEL_EditWorkOrderFolder | OpcenterEXDS_ProductionCoordination | AppU4DM | EditWOOFolder |
| PANEL_CreateSubFolder | OpcenterEXDS_ProductionCoordination | AppU4DM | CreateWOOFolder |
| PANEL_LinkWOOFoldersToWOOFolder | OpcenterEXDS_ProductionCoordination | AppU4DM | LinkWOOFoldersToWOOFolder |
| PANEL_CreateWorkOrderFolder | OpcenterEXDS_ProductionCoordination | AppU4DM | CreateWOOFolder |
| PANEL_LinkWOOperationsToWOOFolder | OpcenterEXDS_ProductionCoordination | AppU4DM | LinkWOOperationsToWOOFolder |
| OperationFolders_Details | OpcenterEXDS_ProductionCoordination | AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM | UnlinkWOOperationsFromWOOFolder<br>DeleteWorkOrderOperation<br>DeleteWOOFolderAndWOOpDependencies<br>UnlinkWOOFoldersFromWOOFolder<br>DeleteWOOFolders |
| WorkOrder_Details | OpcenterEXDS_ProductionCoordination | AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM | DeleteWorkOrderOperation<br>TriggerPrintingOnWorkOrderOperation<br>CreateWorkOrderOutMsg<br>UADMDisAssignProducedMaterialitems<br>DeleteWOOFolders<br>DeleteWOOpDependencyNavigationList<br>ChangeWorkOrderStatusToEdit<br>UADMSetWorkOrderForScheduling<br>TriggerPrintingOnWorkOrder<br>UADMReleaseWorkOrder<br>SetTargetQuantityOnFlexibleWorkOrder<br>UADMAbortWorkOrder<br>DeleteWorkOrder<br>MarkForCleaningWorkOrderList |
| WorkOrderPreKitting_Master | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_SelectSerialNumber | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_PreKitSerialNumber | OpcenterEXDS_ProductionCoordination | AppU4DM<br>OpcenterEXDS_ProductionCoordination | ReserveMaterialItems<br>Java_MapPrekitMaterials |
| PANEL_PreKitSerialNumber_NoData | OpcenterEXDS_ProductionCoordination | - | - |
| WorkOrderPreKitting_Details | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_Note_Management | OpcenterEXDS_ProductionCoordination | AppU4DM<br>AppU4DM | UADMCreateSnagAndNoteList<br>UADMConfirmSnagAndNoteList |
| PANEL_CreateToBeUsedMachine_PCD | OpcenterEXDS_ProductionCoordination | AppU4DM | CreateToBeUsedMachine |
| PANEL_SetWorkOrderOperationEstimatedTimes | OpcenterEXDS_ProductionCoordination | AppU4DM | SetWorkOrderOperationEstimatedTimes |
| PANEL_LinkUserToWorkOrderOperation | OpcenterEXDS_ProductionCoordination | AppU4DM<br>AppU4DM | GetUserDetailsList<br>LinkUserToWorkOrderOperationList |
| ProductionCoordinatorDashboard_Details | OpcenterEXDS_ProductionCoordination | AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM | SetWorkOrderOperationEstimatedTimes<br>UnlinkUserToWorkOrderOperationList<br>SetPreferredMachine<br>UADMReleaseWorkOrder<br>UADMSetWorkOrderForScheduling |
| ProductionCoordinatorDashboard_Master | OpcenterEXDS_ProductionCoordination | - | - |
| QualityInspectionVisualView | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_TeamHistory | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_StepITLKCheckHistory | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_ToolHistory | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_ScrewingHistory | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_OperationSetpoint | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_ITLKCheckHistory | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_ReopenSerializedOperation | OpcenterEXDS_ProductionCoordination | AppU4DM | ReopenWorkOrderOperationSerialized |
| PANEL_ReopenBatchOperation | OpcenterEXDS_ProductionCoordination | AppU4DM | ReopenWorkOrderOperationTransferBatch |
| PANEL_DisassembleMaterialItem | OpcenterEXDS_ProductionCoordination | AppU4DM | DisassembleMaterialItem |
| PANEL_AcquisitionHistory | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_LinkExistingDocument | OpcenterEXDS_ProductionCoordination | Material | DSMaterial_LinkDocumentListToDM_MaterialTrackingUnit |
| PANEL_LinkNewDocument | OpcenterEXDS_ProductionCoordination | Material<br>AppU4DM | DSMaterial_LinkDocumentListToDM_MaterialTrackingUnit<br>UADMCreateDocument |
| Change_Details | OpcenterEXDS_ProductionCoordination | - | - |
| WorkOrder_AsBuilt | OpcenterEXDS_ProductionCoordination | AppU4DM<br>AppU4DM<br>AppU4DM<br>WorkInstruction<br>WorkInstruction<br>AppU4DM<br>Material | CreateWorkOrderOutMsg<br>AdministrativePause<br>ReopenWorkOrderOperationList<br>CreateInspectionAcquisitionContext<br>CreateVisualDetectedFailurewithES<br>UploadDocuments<br>DSMaterial_UnlinkDocumentListToDM_MaterialTrackingUnit |
| FirstArticleInspection_Master | OpcenterEXDS_ProductionCoordination | AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>OpcenterEXDS_ProductionCoordination<br>CommunityCommons<br>Document<br>AppU4DM<br>OpcenterEXDS_ProductionCoordination<br>CommunityCommons<br>CommunityCommons<br>OpcenterEXDS_ProductionCoordination<br>CommunityCommons<br>OpcenterEXDS_ProductionCoordination<br>OpcenterEXDS_ProductionCoordination | RevokeFAICandidate<br>AbortFAIRecord<br>CompleteFAIRecord<br>CopyFAIRecord<br>JA_HtmlToPdf<br>Base64EncodeFile<br>CreateLinkDocument<br>GenerateFAIRNumber<br>JA_ReadTemplate<br>SubstringBefore<br>SubstringAfter<br>JA_ReplacePlaceholdersForm1<br>RegexReplaceAll<br>JA_ReplacePlaceholdersForm2<br>JA_ReplacePlaceholdersForm3 |
| FirstArticleInspection_Details | OpcenterEXDS_ProductionCoordination | AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>OpcenterEXDS_ProductionCoordination<br>CommunityCommons<br>Document<br>AppU4DM<br>OpcenterEXDS_ProductionCoordination<br>CommunityCommons<br>CommunityCommons<br>OpcenterEXDS_ProductionCoordination<br>CommunityCommons<br>OpcenterEXDS_ProductionCoordination<br>OpcenterEXDS_ProductionCoordination | RevokeFAICandidate<br>AbortFAIRecord<br>CompleteFAIRecord<br>CopyFAIRecord<br>JA_HtmlToPdf<br>Base64EncodeFile<br>CreateLinkDocument<br>GenerateFAIRNumber<br>JA_ReadTemplate<br>SubstringBefore<br>SubstringAfter<br>JA_ReplacePlaceholdersForm1<br>RegexReplaceAll<br>JA_ReplacePlaceholdersForm2<br>JA_ReplacePlaceholdersForm3 |
| FirstArticleInspection_InspectionSampleValues | OpcenterEXDS_ProductionCoordination | WorkInstruction<br>WorkInstruction | CreateInspectionAcquisitionContext<br>CreateVisualDetectedFailurewithES |
| FirstArticleInspection_InspectionHistory | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_ToastNotification | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_CopyFAIRecordWizard | OpcenterEXDS_ProductionCoordination | AppU4DM | CopyFAIRecord |
| PANEL_CopyFAIRecordOnlyCompleteRelated | OpcenterEXDS_ProductionCoordination | AppU4DM | CopyFAIRecord |
| PANEL_ChangeFAICandidate | OpcenterEXDS_ProductionCoordination | AppU4DM | ChangeFAICandidate |
| PANEL_DeclareFAICandidate | OpcenterEXDS_ProductionCoordination | AppU4DM | DeclareFAICandidate |
| PANEL_UpdateLSPThreshold | OpcenterEXDS_ProductionCoordination | Kanban | UpdateLSPThreshold |
| PANEL_UpdateLineSidePosition | OpcenterEXDS_ProductionCoordination | Kanban | UpdateLineSidePosition |
| PANEL_CreateLSPThreshold | OpcenterEXDS_ProductionCoordination | Kanban | CreateLSPThreshold |
| PANEL_CreateLineSidePosition | OpcenterEXDS_ProductionCoordination | Kanban | CreateLineSidePosition |
| LineSidePosition_Details | OpcenterEXDS_ProductionCoordination | Kanban<br>Kanban<br>Kanban | DeleteLSPThreshold<br>ReleaseLineSidePosition<br>UnreleaseLineSidePosition |
| LineSidePositions_Master | OpcenterEXDS_ProductionCoordination | Kanban<br>Kanban<br>Kanban | ReleaseLineSidePosition<br>UnreleaseLineSidePosition<br>DeleteLineSidePosition |
| Panel_Scanner | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination<br>AppU4DM<br>AppU4DM | JAVA_GetClientId<br>AddMaterialTrackingUnitsToWorkingSession<br>CreateWorkingSession |
| ScanMaterialTrackingUnits_Master | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination<br>AppU4DM<br>AppU4DM<br>AppU4DM | JAVA_GetClientId<br>AddMaterialTrackingUnitsToWorkingSession<br>CreateWorkingSession<br>RemoveMaterialTrackingUnitsFromWorkingSession |
| Panel_MaterialTrackingUnit_ReleaseByFile | OpcenterEXDS_ProductionCoordination | AppU4DM<br>OpcenterEXDS_ProductionCoordination<br>OpcenterEXDS_ProductionCoordination | BulkMaterialTrackingUnitsReleaseFromFile<br>JAVA_GetItemsTotalCount<br>JAVA_ValidateFileHeader |
| Panel_MaterialTrackingUnit_Scrap | OpcenterEXDS_ProductionCoordination | AppU4DM<br>AppU4DM<br>AppU4DM | GenerateOPNumberingPattern<br>BulkMaterialTrackingUnitsScrapFromNonConformanceItems<br>BulkMaterialTrackingUnitsScrap |
| Panel_MaterialTrackingUnit_Release | OpcenterEXDS_ProductionCoordination | AppU4DM | BulkMaterialTrackingUnitsReleaseFromNonConformanceItems |
| Panel_MaterialTrackingUnit_ScrapByFile | OpcenterEXDS_ProductionCoordination | AppU4DM<br>OpcenterEXDS_ProductionCoordination<br>OpcenterEXDS_ProductionCoordination<br>AppU4DM | BulkMaterialTrackingUnitsScrapFromFile<br>JAVA_GetItemsTotalCount<br>JAVA_ValidateFileHeader<br>GenerateOPNumberingPattern |
| PANEL_EditNPAForUser | OpcenterEXDS_ProductionCoordination | AppU4DM | UADMUpdateNonProductiveActivityForUser |
| PANEL_CreateNPAForUser | OpcenterEXDS_ProductionCoordination | AppU4DM | UADMRegisterNonProductiveActivityForUser |
| Users_Details | OpcenterEXDS_ProductionCoordination | AppU4DM<br>AppU4DM | UADMDeleteNonProductiveActivityListForUser<br>UADMCompleteNonProductiveActivityList |
| Users_Master | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_UpdateExecutionGroup | OpcenterEXDS_ProductionCoordination | AppU4DM | UpdateExecutionGroup |
| PANEL_CreateExecutionGroup | OpcenterEXDS_ProductionCoordination | AppU4DM | UADMCreateExecutionGroup |
| ExecutionGroup_Master | OpcenterEXDS_ProductionCoordination | AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppName<br>AppU4DM | ReleaseExecutionGroup<br>ScheduleExecutionGroup<br>SetExecGroupListInEditStatus<br>AbortExecutionGroupList<br>CommandName<br>AppU4DMAMN_GetReadyPrintJobs |
| PANEL_EditWOOP | OpcenterEXDS_ProductionCoordination | AppU4DM | UpdateWOOpQuantityAssociatedToEG |
| PANEL_LinkWOOperationToExecutionGroup | OpcenterEXDS_ProductionCoordination | AppU4DM | LinkWOOperationListToExecutionGroup |
| PANEL_UpdateExecutionGroup_Details | OpcenterEXDS_ProductionCoordination | AppU4DM | UpdateExecutionGroup |
| ExecutionGroup_Details | OpcenterEXDS_ProductionCoordination | AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM<br>AppName | UnlinkWOOperationsFromExecutionGroup<br>ReleaseExecutionGroup<br>ScheduleExecutionGroup<br>SetExecGroupListInEditStatus<br>AbortExecutionGroupList<br>CommandName |
| ExecutionGroupPhase_Details | OpcenterEXDS_ProductionCoordination | AppU4DM<br>PowderMgt | UnlinkWIDefinitionsFromEGPhase<br>RemoveAMPowderLoadedCheckFromEgPhaseList |
| PANEL_LinkPrintJobFile | OpcenterEXDS_ProductionCoordination | AppU4DM | LinkToBeUsedMachineToPJFForEGPhase |
| PANEL_EditAMPowder | OpcenterEXDS_ProductionCoordination | PowderMgt | EditAMPowderLoadedCheckOnEgPhase |
| PANEL_UpdateExecutionGroupPhase | OpcenterEXDS_ProductionCoordination | AppU4DM | UpdateExecutionGroupPhase |
| PANEL_PreTransfer | OpcenterEXDS_ProductionCoordination | PrintJobFile | RequestTransferPrintJobFile |
| PANEL_LinkWInstructionToExecutionGroupPhase | OpcenterEXDS_ProductionCoordination | AppU4DM | LinkWIDefinitionsToEGPhase |
| PANEL_LoadAMPowder | OpcenterEXDS_ProductionCoordination | PowderMgt | SetAMPowderLoadedCheckOnEgPhase |
| OfflineSession_Details | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination<br>CmdApp<br>OpcenterEXDS_ProductionCoordination<br>AppU4DM<br>AppU4DM<br>AppU4DM | Java_ConvertToCmdActionAmend<br>CmdName<br>Java_ConvertToCmdAction<br>ChangeOfflineSessionStatus<br>UpdateOfflineAction<br>CompleteCheckIn |
| PANEL_EditOfflineAction | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination<br>OpcenterEXDS_ProductionCoordination<br>AppU4DM | Java_DecodeBinaryActionAmendToJson<br>Java_DecodedActionToJson<br>AmendOfflineAction |
| OfflineSession_Master | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination<br>OpcenterEXDS_ProductionCoordination<br>IteratorOfflineAction/CommandApp<br>AppU4DM<br>AppU4DM<br>AppU4DM | Java_ConvertToCmdActionAmend<br>Java_ConvertToCmdAction<br>IteratorOfflineAction/CommandName<br>ChangeOfflineSessionStatus<br>CompleteCheckIn<br>DiscardOfflineSessionList |
| WorkOrderUpdateCheck_Master | OpcenterEXDS_ProductionCoordination | AppU4DM<br>AppU4DM | UADMUpdateOutOfDateWOByBoPList<br>PLMUpdateWorkOrderByCCList |
| PANEL_UpdateLSPAvailableQuantity | OpcenterEXDS_ProductionCoordination | Kanban | UpdateLSPAvailableQuantity |
| PANEL_DeclareKanbanCallDelivered | OpcenterEXDS_ProductionCoordination | Kanban | DeclareKanbanCallDelivered |
| LineSidePositionMonitoring_KanbanRequest_Details | OpcenterEXDS_ProductionCoordination | - | - |
| LineSidePositionMonitoring_Details | OpcenterEXDS_ProductionCoordination | Kanban<br>Kanban<br>Kanban<br>Kanban | CreateManualKanbanCall<br>CreateOnlyManualKanbanCall<br>CancelKanbanCall<br>IncreaseKanbanCallPriority |
| LineSidePositionsMonitoring_Master | OpcenterEXDS_ProductionCoordination | Kanban<br>Kanban<br>Kanban<br>Kanban | CreateManualKanbanCall<br>CreateOnlyManualKanbanCall<br>CancelKanbanCall<br>IncreaseKanbanCallPriority |
| ContainmentRequests_Master | OpcenterEXDS_ProductionCoordination | - | - |
| ContainmentRequests_Details | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_SetContainmentRequestAsReady | OpcenterEXDS_ProductionCoordination | AppU4DM | SetContainmentRequestStatus |
| Panel_ManageContainmentRequestRelease | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination<br>OpcenterEXDS_ProductionCoordination<br>AppU4DM | JAVA_GetItemsTotalCount<br>JAVA_ValidateFileHeader<br>BulkContainmentRequestRelease |
| PANEL_BulkMaterialTrackingUnitsScrapFromContainmentRequest | OpcenterEXDS_ProductionCoordination | AppU4DM<br>AppU4DM | BulkMaterialTrackingUnitsScrapFromContainmentRequest<br>GenerateOPNumberingPattern |
| PANEL_ManageContainmentRequestFullRelease | OpcenterEXDS_ProductionCoordination | AppU4DM | BulkContainmentRequestRelease |
| Panel_CreateContainmentRequest | OpcenterEXDS_ProductionCoordination | AppU4DM | CreateContainmentRequest |
| PANEL_CloseContainmentRequest | OpcenterEXDS_ProductionCoordination | AppU4DM | SetContainmentRequestStatus |
| PANEL_BulkMaterialTrackingUnitsScrap | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination<br>OpcenterEXDS_ProductionCoordination<br>AppU4DM<br>AppU4DM<br>AppU4DM | JAVA_GetItemsTotalCount<br>JAVA_ValidateFileHeader<br>GenerateOPNumberingPattern<br>BulkMaterialTrackingUnitsScrap<br>BulkMaterialTrackingUnitsScrapFromFile |
| PANEL_CompareChangePackageItem_ForMachine | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_AddChangePackageItemMachine_Workcenter | OpcenterEXDS_ProductionCoordination | AppU4DM | AddChangePackageItemMachine |
| PANEL_AddChangePackageItemMachine_EquipmentType | OpcenterEXDS_ProductionCoordination | AppU4DM | AddChangePackageItemMachine |
| PANEL_AddChangePackageItemMachine_Unit | OpcenterEXDS_ProductionCoordination | AppU4DM | AddChangePackageItemMachine |
| PANEL_AddChangePackageItemMaterial | OpcenterEXDS_ProductionCoordination | AppU4DM | AddChangePackageItemMaterial |
| PANEL_CompareChangePackageItem_ForMaterial | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_CompareChangePackageItem | OpcenterEXDS_ProductionCoordination | - | - |
| WorkInstruction_PreviewForNewInstructions | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_AddChangePackageItemWIDefinition | OpcenterEXDS_ProductionCoordination | AppU4DM | AddChangePackageItemWIDefinition |
| PANEL_UADMEditChangePackageItemWIDefinition | OpcenterEXDS_ProductionCoordination | AppU4DM | UADMEditChangePackageItemWIDefinition |
| PANEL_AddChangePackageItemWorkOrderStep | OpcenterEXDS_ProductionCoordination | AppU4DM | AddChangePackageItemWorkOrderStep |
| PANEL_CompareChangePackageItem_ForStep | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_AddChangePackageItemStepWIDefinition | OpcenterEXDS_ProductionCoordination | AppU4DM | AddChangePackageItemStepWIDefinition |
| WorkInstruction_PreviewForNewInstructions_ForWIStep | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_UADMEditChangePackageItemStepWIDefinition | OpcenterEXDS_ProductionCoordination | AppU4DM | UADMEditChangePackageItemStepWIDefinition |
| PANEL_CompareChangePackageItem_ForWIStep | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_AddChangePackageItemTool_ForStepTool | OpcenterEXDS_ProductionCoordination | AppU4DM | AddChangePackageItemTool |
| PANEL_AddChangePackageItemMaterial_ForStep | OpcenterEXDS_ProductionCoordination | AppU4DM | AddChangePackageItemMaterial |
| PANEL_AddChangePackageItemQualityInspection_Step | OpcenterEXDS_ProductionCoordination | AppU4DM | AddChangePackageItemQualityInspection |
| PANEL_AddChangePackageItemDocument_ForStep | OpcenterEXDS_ProductionCoordination | AppU4DM | AddChangePackageItemDocument |
| PANEL_AddChangePackageItemNewDocument_ForStep | OpcenterEXDS_ProductionCoordination | AppU4DM | AddChangePackageItemNewDocument |
| Steps_Details | OpcenterEXDS_ProductionCoordination | AppName | CommandName |
| PANEL_AddChangePackageWorkOrderOperation | OpcenterEXDS_ProductionCoordination | AppU4DM | AddChangePackageItemWorkOrderOperation |
| PANEL_ChangePackageAddWOOpDependencies | OpcenterEXDS_ProductionCoordination | AppU4DM | AddChangePackageItemWorkOOperationDependency |
| PANEL_CompareChangePackageItem_ForSummary | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_AddChangePackageItemQualityInspection | OpcenterEXDS_ProductionCoordination | AppU4DM | AddChangePackageItemQualityInspection |
| PANEL_CompareChangePackageItem_ForQI | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_AddChangePackageItemNewDocument | OpcenterEXDS_ProductionCoordination | AppU4DM | AddChangePackageItemNewDocument |
| PANEL_AddChangePackageItemDocument | OpcenterEXDS_ProductionCoordination | AppU4DM | AddChangePackageItemDocument |
| PANEL_CompareChangePackageItem_ForDocument | OpcenterEXDS_ProductionCoordination | - | - |
| PreviewRouting | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_AddChangePackageItemTool | OpcenterEXDS_ProductionCoordination | AppU4DM | AddChangePackageItemTool |
| PANEL_CompareChangePackageItem_ForTools | OpcenterEXDS_ProductionCoordination | - | - |
| ChangePackages_WorkOrderOverview_Details | OpcenterEXDS_ProductionCoordination | AppName<br>AppU4DM<br>AppU4DM | CommandName<br>RemoveChangePackageItemWorkOrderOperation<br>RemoveChangePackageItemWorkOOperationDependency |
| ChangePackages_Details | OpcenterEXDS_ProductionCoordination | AppName | CommandName |
| PANEL_AcceptChangePackage | OpcenterEXDS_ProductionCoordination | AppU4DM | UADMAcceptChangePackage |
| PANEL_RejectChangePackage | OpcenterEXDS_ProductionCoordination | AppU4DM | RejectChangePackage |
| ChangePackages_Master | OpcenterEXDS_ProductionCoordination | - | - |
| IntegrationEvents_Master | OpcenterEXDS_ProductionCoordination | APPU4DM | ResendConfirmWorkOrderOperation |
| PANEL_UADMLinkSkillsToTeam | OpcenterEXDS_ProductionCoordination | AppU4DM | UADMLinkSkillsToTeam |
| PANEL_UADMLinkUsersToTeam | OpcenterEXDS_ProductionCoordination | AppU4DM | UADMLinkUsersToTeam |
| Teams_Details | OpcenterEXDS_ProductionCoordination | AppU4DM<br>AppU4DM | UnLinkSkillsFromTeam<br>UADMUnLinkUsersFromTeam |
| PANEL_UpdateTeam | OpcenterEXDS_ProductionCoordination | AppU4DM | UpdateTeam |
| PANEL_ExportTeams | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_CreateTeam | OpcenterEXDS_ProductionCoordination | AppU4DM | CreateTeam |
| Teams_Master | OpcenterEXDS_ProductionCoordination | AppU4DM | DeleteTeam |
| SPCEvaluation_Details | OpcenterEXDS_ProductionCoordination | WorkInstruction<br>OpcenterEXDS_ProductionCoordination_Connector | QualityCharacteristicSPCEvaluation<br>FromBlobToJson |
| SPCEvaluation_Master | OpcenterEXDS_ProductionCoordination | - | - |
| PCTimeUpdate_Master | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_CreateWorkOrderHistory | OpcenterEXDS_ProductionCoordination | AppU4DM | CreateWorkOrderHistory |
| PANEL_UpdateWorkOrderHistory | OpcenterEXDS_ProductionCoordination | AppU4DM | UpdateWorkOrderHistory |
| PANEL_UpdateWorkOrderOperationActualExecutionTimes | OpcenterEXDS_ProductionCoordination | AppU4DM | UpdateWorkOrderOperationActualExecutionTimes |
| PCTimeUpdate_Details | OpcenterEXDS_ProductionCoordination | AppU4DM | DeleteWorkOrderHistory |
| WorkOrderRouting_Details | OpcenterEXDS_ProductionCoordination | - | - |
| WorkOrderNetwork_Master | OpcenterEXDS_ProductionCoordination | - | - |
| PANEL_CreateWorkOrderDependency | OpcenterEXDS_ProductionCoordination | AppU4DM | CreateWorkOrderDependency |
| PANEL_CreateChangeNonConformance | OpcenterEXDS_ProductionCoordination | AppU4DM | UADMCreateChangeNonConformance |
| PANEL_UADMAcceptChangeAddOperation | OpcenterEXDS_ProductionCoordination | AppU4DM | UADMAcceptChangeAddOperation |
| PANEL_UADMAcceptChangeAddWorkInstruction | OpcenterEXDS_ProductionCoordination | AppU4DM | UADMAcceptChangeAddWorkInstruction |
| PANEL_UADMAcceptChangeRemoveDependency | OpcenterEXDS_ProductionCoordination | AppU4DM | UADMAcceptChangeRemoveDependency |
| PANEL_UADMAcceptChangeReplaceToBeConsumedMaterial | OpcenterEXDS_ProductionCoordination | AppU4DM | UADMAcceptChangeReplaceToBeConsumedMaterial |
| PANEL_UADMAcceptChangeAddToBeConsumedMaterial | OpcenterEXDS_ProductionCoordination | AppU4DM | UADMAcceptChangeAddToBeConsumedMaterial |
| PANEL_UADMAcceptChangeRemoveOperation | OpcenterEXDS_ProductionCoordination | AppU4DM | UADMAcceptChangeRemoveOperation |
| PANEL_UADMAcceptChangeChangeDependency | OpcenterEXDS_ProductionCoordination | AppU4DM | UADMAcceptChangeChangeDependency |
| PANEL_UADMAcceptChangeChangeToBeConsumedMaterialQuantity | OpcenterEXDS_ProductionCoordination | AppU4DM | UADMAcceptChangeChangeToBeConsumedMaterialQuantity |
| PANEL_UADMAcceptChangeRepeatOperation | OpcenterEXDS_ProductionCoordination | AppU4DM | UADMAcceptChangeRepeatOperation |
| PANEL_UADMAcceptChangeAddWorkInstruction_SN | OpcenterEXDS_ProductionCoordination | AppU4DM | UADMAcceptChangeAddWorkInstruction |
| PANEL_UADMAcceptChangeAddToBeUsedTool | OpcenterEXDS_ProductionCoordination | AppU4DM | UADMAcceptChangeAddToBeUsedTool |
| PANEL_UADMAcceptChangeRemoveToBeConsumedMaterial | OpcenterEXDS_ProductionCoordination | AppU4DM | UADMAcceptChangeRemoveToBeConsumedMaterial |
| PANEL_AcceptChangeAddProcessOperation | OpcenterEXDS_ProductionCoordination | AppU4DM<br>AppU4DM | GetProductionProcessFromAsPlanned<br>AcceptChangeAddProcessOperation |
| PANEL_UADMRejectChangeNonConformance | OpcenterEXDS_ProductionCoordination | AppU4DM | UADMRejectChangeNonConformance |
| ChangeRequest_Details | OpcenterEXDS_ProductionCoordination | - | - |
| ChangeRequest_Master | OpcenterEXDS_ProductionCoordination | - | - |
| ContainmentRequestMonitoring_Master | OpcenterEXDS_ProductionCoordination | AppU4DM<br>AppU4DM<br>AppU4DM | DeleteImportedContainmentResults<br>GetContainmentReleaseBySequence<br>GetContainmentRequestBySequence |
| Genealogy_MasterV2 | OpcenterEXDS_ShopfloorExecution | CommunityCommons<br>CommunityCommons<br>AppU4DM<br>AppU4DM<br>AppU4DM | RandomHash<br>StringSplit<br>UADMExportGenealogy<br>UnlinkCollectedDocumentList<br>UnlinkDocumentListToDM_MaterialTrackingUnit |
| PANEL_AddDocumentGenealogy | OpcenterEXDS_ShopfloorExecution | AppU4DM | UADMCreateDocument |
| PANEL_ImportGenealogy | OpcenterEXDS_ShopfloorExecution | - | - |
| PANEL_DisassembleMaterialItem | OpcenterEXDS_ShopfloorExecution | AppU4DM | DisassembleMaterialItem |
| NonConformance_Master | OpcenterEXDS_ShopfloorExecution | - | - |
| PANEL_UADMSentenceNonConformanceV3_1 | OpcenterEXDS_ShopfloorExecution | AppU4DM<br>AppU4DM | GetCurrentUserRoles<br>UADMSentenceNonConformanceV3_1 |
| PANEL_UADMLinkGenericItemsToNonConformance | OpcenterEXDS_ShopfloorExecution | AppU4DM | UADMLinkGenericItemsToNonConformance |
| Panel_UADMCreateDocument | OpcenterEXDS_ShopfloorExecution | AppU4DM<br>AppU4DM | UADMCreateDocument<br>CreateNonConformanceAttachmentList |
| PANEL_DeclareNonConformance | OpcenterEXDS_ShopfloorExecution | AppU4DM | UADMCreateNonConformanceV3_1 |
| PANEL_UADMCreateNonConformanceDefectList | OpcenterEXDS_ShopfloorExecution | AppU4DM | UADMCreateNonConformanceDefectList |
| PANEL_CreateNonConformanceAttachmentList | OpcenterEXDS_ShopfloorExecution | AppU4DM | CreateNonConformanceAttachmentList |
| PANEL_AddNCNewDocument | OpcenterEXDS_ShopfloorExecution | Document<br>AppU4DM | CreateDocument<br>CreateNonConformanceAttachmentList |
| PANEL_UADMUpdateNonConformance | OpcenterEXDS_ShopfloorExecution | AppU4DM | UADMUpdateNonConformance |
| NonConformance_Details | OpcenterEXDS_ShopfloorExecution | AppU4DM<br>AppU4DM<br>AppU4DM<br>AppU4DM | RemoveDefectsFromNonConformance<br>DeleteNonConformanceAttachmentList<br>UploadDocuments<br>UADMUnLinkGenericItemsFromNonConformance |
| Tool_Master | OpcenterEXDS_ShopfloorExecution | AppU4DM | DeleteTool |
| PANEL_DeclareToolNonConformance | OpcenterEXDS_ShopfloorExecution | AppU4DM | UADMCreateNonConformanceV3_1 |
| PANEL_BrowseFailures | OpcenterEXDS_ShopfloorExecution | - | - |
| PANEL_UpdateTool | OpcenterEXDS_ShopfloorExecution | AppU4DM | UpdateTool |
| PANEL_ToolMaintenance | OpcenterEXDS_ShopfloorExecution | AppU4DM | ToolMaintenance |
| PANEL_CreateNewLogisticClass | OpcenterEXDS_ShopfloorExecution | AppU4DM | CreateLogisticClass |
| PANEL_UADMCreateToolDefinition | OpcenterEXDS_ShopfloorExecution | AppU4DM | UADMCreateToolDefinition |
| PANEL_CreateTool | OpcenterEXDS_ShopfloorExecution | AppU4DM | CreateTool |
| PANEL_LinkToolToAutomationNodeInstanceList | OpcenterEXDS_ShopfloorExecution | AppU4DM | LinkToolToAutomationNodeInstanceList |
| PANEL_LinkSubstrateToMachine | OpcenterEXDS_ShopfloorExecution | AppU4DM | LinkSubstrateToMachine |
| ToolDetails_Details | OpcenterEXDS_ShopfloorExecution | AppU4DM<br>AppU4DM | UnlinkSubstrateToMachine<br>UnlinkToolToAutomationNodeInstanceList |
| PANEL_UADMSetLocationHold | OpcenterEXDS_Configuration | AppU4DM | UADMSetLocationHold |
| PANEL_UADMSetWorkOrderHoldList | OpcenterEXDS_Configuration | AppU4DM | UADMSetWorkOrderHoldList |
| PANEL_UADMRemoveHoldList | OpcenterEXDS_Configuration | AppU4DM | UADMRemoveHoldList |
| Hold_Master | OpcenterEXDS_Configuration | - | - |
| Hold_Details | OpcenterEXDS_Configuration | - | - |
| PANEL_MTUSplit | OpcenterEXDS_MasterData | Material<br>Material | DSMaterial_MaterialItemSplitBatch<br>DSMaterial_MaterialItemSplitBatchByPartialPreview |
| PANEL_SetMaterialTrackingUnitQuantity | OpcenterEXDS_MasterData | Material | SetMaterialTrackingUnitQuantity |
| PANEL_AssociateMaterialTrackingUnitsWithMaterialLot | OpcenterEXDS_MasterData | Material | AssociateMaterialTrackingUnitsWithMaterialLot |
| PANEL_MoveMaterialTrackingUnitToEquipment | OpcenterEXDS_MasterData | Material | MoveMaterialTrackingUnitToEquipment |
| PANEL_MoveMaterialTrackingUnitToMaterialTrackingUnitAggregate | OpcenterEXDS_MasterData | Material | MoveMaterialTrackingUnitToMaterialTrackingUnitAggregate |
| PANEL_SetMaterialTrackingUnitStatus | OpcenterEXDS_MasterData | Material | SetMaterialTrackingUnitMaterial |
| PANEL_CreateMaterialTrackingUnit | OpcenterEXDS_MasterData | Material | DSMaterial_CreateDM_MaterialTrackingUnitandMaterialTrackingUnit |
| PANEL_EditMaterialTrackingUnit | OpcenterEXDS_MasterData | Material | DSMaterial_UpdateMaterialTrackingUnitAndDM_MaterialTrackingUnit |
| PANEL_SetStatusMaterialTrackingUnit | OpcenterEXDS_MasterData | Material | SetMaterialTrackingUnitStatus |
| PANEL_SetMaterialTrackingUnitStateMachine | OpcenterEXDS_MasterData | Material | SetMaterialTrackingUnitStateMachine |
| PANEL_LinkNewDocument | OpcenterEXDS_MasterData | Material<br>AppU4DM | DSMaterial_LinkDocumentListToDM_MaterialTrackingUnit<br>UADMCreateDocument |
| PANEL_LinkExistingDocument | OpcenterEXDS_MasterData | Material | DSMaterial_LinkDocumentListToDM_MaterialTrackingUnit |
| PANEL_ResultAmend | OpcenterEXDS_MasterData | Material | DSResult_AmendResultList |
| PANEL_ResultHistory | OpcenterEXDS_MasterData | - | - |
| MaterialTrackingUnit_Details | OpcenterEXDS_MasterData | Material<br>AppName | DSMaterial_UnlinkDocumentListToDM_MaterialTrackingUnit<br>CommandName |
| MaterialTrackingUnit_Master | OpcenterEXDS_MasterData | AppName<br>AppU4DM | CommandName<br>DSMaterial_UADMGenerateAndAssociateMTUCode |
| PANEL_CreateMaterialTrackingUnitFromMaterialLot | OpcenterEXDS_MasterData | Material<br>Material | CreateMaterialTrackingUnit<br>DSMaterial_UADMCreateDM_MTU |

---

## 9. Page Entities

Pages using the EXFN_Master layout that load entities via DataGrid2 or Gallery datasources.

| Page | Module | Entities |
|------|--------|----------|
| Hold_Master | OpcenterEXDS_Configuration | OpcenterEXDS_Configuration.Hold |
| MaterialTrackingUnit_Details | OpcenterEXDS_MasterData | OpcenterEXDS_MasterData.RM_PC_BufferFromMTU<br>OpcenterEXDS_MasterData.RM_PC_DM_MaterialTrackingUnitHistory<br>OpcenterEXDS_MasterData.RM_PC_DocumentsLinkedToEntityMaterial<br>OpcenterEXDS_MasterData.RM_PC_MTULocationInfo<br>OpcenterEXDS_MasterData.RM_PC_MaterialTrackingUnitResult<br>OpcenterEXDS_MasterData.RM_PC_TraceabilityHistory |
| MaterialTrackingUnit_Master | OpcenterEXDS_MasterData | OpcenterEXDS_MasterData.RM_PC_DM_MaterialTrackingUnit |
| Buffers_Details | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination.RM_PC_BufferHistory<br>OpcenterEXDS_ProductionCoordination.RM_PC_BufferMaterialItemAssociation |
| Buffers_Master | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination.RM_PC_Buffer |
| ChangePackages_Details | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination.RM_PC_ChangePackageDocument<br>OpcenterEXDS_ProductionCoordination.RM_PC_ChangePackageMachine<br>OpcenterEXDS_ProductionCoordination.RM_PC_ChangePackageMaterial<br>OpcenterEXDS_ProductionCoordination.RM_PC_ChangePackageQualityInspection<br>OpcenterEXDS_ProductionCoordination.RM_PC_ChangePackageSummary<br>OpcenterEXDS_ProductionCoordination.RM_PC_ChangePackageToBeUsedTool<br>OpcenterEXDS_ProductionCoordination.RM_PC_ChangePackageWorkInstruction<br>OpcenterEXDS_ProductionCoordination.RM_PC_ChangePackageWorkOrderStep |
| ChangePackages_Master | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination.RM_PC_ChangePackage |
| ChangePackages_WorkOrderOverview_Details | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination.RM_PC_ChangePackageSummary<br>OpcenterEXDS_ProductionCoordination.RM_PC_ChangePackageWorkOOperationDependency<br>OpcenterEXDS_ProductionCoordination.RM_PC_ChangePackageWorkOrderOperation |
| ChangeRequest_Master | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_Core.NonConformance |
| ContainmentRequestMonitoring_Master | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination.RM_PC_ContainmentRequestMonitoring |
| ContainmentRequests_Master | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination.RM_PC_ContainmentRequest |
| ERPOrder_Details | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination_Connector.GetBOFLinkedERP_RF<br>OpcenterEXDS_ProductionCoordination_Connector.GetBOMLinkedERP_RF |
| ERPOrder_Master | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination_Connector.GetERPOrderWithValidity_RF |
| ExecutionGroupPhase_Details | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination.AMPowderRequiredOn3DPrinter |
| ExecutionGroupPhase_Details_2 | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination.AMPowderRequiredOn3DPrinter |
| ExecutionGroup_Master | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination_Connector.GetExecutionGroupListWithEstimatedDuration_RF |
| FirstArticleInspection_Details | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination.FAIInspection<br>OpcenterEXDS_ProductionCoordination.RM_PC_FAIRecordHistory<br>OpcenterEXDS_ProductionCoordination.RM_PC_FAIWorkOrderRelevant |
| FirstArticleInspection_InspectionHistory | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination.RM_PC_FAIInspectionHistory |
| FirstArticleInspection_InspectionSampleValues | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination.RM_PC_QualityInspectionSamples |
| FirstArticleInspection_Master | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination.RM_PC_FAIRecord |
| FutureHoldManagement_Master | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination.FutureHold |
| HandlingUnit_Details | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination.HandlingUnitHistory |
| HandlingUnits_Master | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination.HandlingUnit |
| IntegrationEvents_Master | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination_Connector.UADMGetEventLogs_RF |
| LineSidePositionMonitoring_Details | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination.KanbanCall<br>OpcenterEXDS_ProductionCoordination.LineSideInstanceHistory |
| LineSidePositionsMonitoring_Master | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination.RM_PC_LineSideMonitorInfo |
| LineSidePositions_Master | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination.LineSidePosition |
| LogisticRequest_Master | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination.LogisticRequestItem |
| MaterialSuppliers_Details | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination_Connector.GetSuppliersFromMaterial_RF |
| Measurements_Master | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination.RM_PC_ProcessMeasurement |
| OfflineSession_Details | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination.OfflineAction |
| OfflineSession_Master | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination_Connector.UADMGetAllOfflineSessionsLists_RF |
| OperationFolders_Details | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_Core.WOOFolder<br>OpcenterEXDS_ProductionCoordination.RM_PC_WorkOrderDependencies |
| PCTimeUpdate_Details | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination.RM_PC_WorkOrderHistoryAsBuilt<br>OpcenterEXDS_ProductionCoordination.RM_PC_WorkOrderOperation |
| PCTimeUpdate_Master | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination.RM_PC_WorkOrder |
| ProductionCoordinatorDashboard_Master | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination.RM_PC_EquipmentWorkOrderOperationCount |
| SPCEvaluation_Master | OpcenterEXDS_ProductionCoordination | OpcenterEXFN_MasterData.CharacteristicSpecification |
| Steps_Details | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination.RM_PC_ChangePackageStepDocument<br>OpcenterEXDS_ProductionCoordination.RM_PC_ChangePackageStepMaterial<br>OpcenterEXDS_ProductionCoordination.RM_PC_ChangePackageStepQualityInspection<br>OpcenterEXDS_ProductionCoordination.RM_PC_ChangePackageStepToBeUsedTool<br>OpcenterEXDS_ProductionCoordination.RM_PC_ChangePackageStepWorkInstruction |
| Teams_Details | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination.TeamHistory<br>OpcenterEXDS_ProductionCoordination.TeamSkillAssociation<br>OpcenterEXDS_ProductionCoordination.TeamUserAssociation |
| Teams_Master | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination.Team |
| Users_Details | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination.RM_PC_UserNonProductiveActivity<br>OpcenterEXDS_ProductionCoordination.RM_PC_UserProductiveActivity |
| WorkOrderNetwork_Master | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination.RM_PC_WorkOrder |
| WorkOrderOperationSteps_Details | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_Core.ToBeUsedTool<br>OpcenterEXDS_ProductionCoordination.RM_PC_CharacteristicRepresentation<br>OpcenterEXDS_ProductionCoordination.RM_PC_GetAllMaterialDefinitionsForAllMaterialSpecificationTypes<br>OpcenterEXDS_ProductionCoordination.RM_PC_GetWIDetailsForWorkOrderOperationAndStep<br>OpcenterEXDS_ProductionCoordination.RM_PC_ItlkCheckWOStepAssociation<br>OpcenterEXDS_ProductionCoordination.WorkOStepDependency<br>OpcenterEXDS_ProductionCoordination.WorkOrderStepSkill |
| WorkOrderOperation_Details | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_Core.ToBeUsedTool<br>OpcenterEXDS_ProductionCoordination.OutMsg<br>OpcenterEXDS_ProductionCoordination.RM_PC_CharacteristicRepresentation<br>OpcenterEXDS_ProductionCoordination.RM_PC_GetAllMaterialDefinitionsForAllMaterialSpecificationTypes<br>OpcenterEXDS_ProductionCoordination.RM_PC_GetFullEquipment<br>OpcenterEXDS_ProductionCoordination.RM_PC_GetWIDetailsForWorkOrderOperationAndStep<br>OpcenterEXDS_ProductionCoordination.RM_PC_GetWOSteps<br>OpcenterEXDS_ProductionCoordination.RM_PC_ItlkCheckWOOpAssociation<br>OpcenterEXDS_ProductionCoordination.RM_PC_WorkOrderDependencies<br>OpcenterEXDS_ProductionCoordination.UserWorkOrderOperationAssociation<br>OpcenterEXDS_ProductionCoordination.WorkOrderOperationSkill |
| WorkOrderPreKitting_Details | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination.ToBeConsumedMaterialPrekitHistory<br>OpcenterEXDS_ProductionCoordination_Connector.UADMGetAllMaterialForPrekit_RF |
| WorkOrderPreKitting_Master | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_Core.WorkOrder |
| WorkOrderUpdateCheck_Master | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination.RM_PC_WorkOrder |
| WorkOrderUpdate_Master | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_ProductionCoordination_Connector.UADMGetWOByProcessRevisions |
| WorkOrder_AsBuilt | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_Configuration.Hold<br>OpcenterEXDS_Core.NonConformance<br>OpcenterEXDS_ProductionCoordination.EXDSAsBuiltSN<br>OpcenterEXDS_ProductionCoordination.RM_PC_ActualUsedTool<br>OpcenterEXDS_ProductionCoordination.RM_PC_AsBuiltMaterials<br>OpcenterEXDS_ProductionCoordination.RM_PC_ProcessMeasurement<br>OpcenterEXDS_ProductionCoordination.RM_PC_QualityInspectionSamples<br>OpcenterEXDS_ProductionCoordination.RM_PC_SnagAndNote<br>OpcenterEXDS_ProductionCoordination.RM_PC_WorkInstructionAsBuilt<br>OpcenterEXDS_ProductionCoordination.RM_PC_WorkOrderHistoryAsBuilt<br>OpcenterEXDS_ProductionCoordination.RM_PC_WorkOrderOperation<br>OpcenterEXDS_ProductionCoordination.RM_PC_WorkOrderRuntimeDocument<br>OpcenterEXDS_ProductionCoordination.RM_PC_WorkOrderStep<br>OpcenterEXDS_ProductionCoordination.WorkOrderHistory |
| WorkOrder_Details | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_Core.WOOFolder<br>OpcenterEXDS_ProductionCoordination.OutMsg<br>OpcenterEXDS_ProductionCoordination.RM_PC_DM_MaterialTrackingUnitAndMTU<br>OpcenterEXDS_ProductionCoordination.RM_PC_GetAlternativeGroupsFromWorkOrderId<br>OpcenterEXDS_ProductionCoordination.RM_PC_GetDataFromFeatures<br>OpcenterEXDS_ProductionCoordination.RM_PC_GetExecutionGroupsForWorkOrder |
| WorkOrder_Master | OpcenterEXDS_ProductionCoordination | OpcenterEXDS_Core.WorkOrder |
| NonConformance_Details | OpcenterEXDS_ShopfloorExecution | OpcenterEXDS_Core.RM_PC_NonConformanceDocuments<br>OpcenterEXDS_Core.RM_PC_NonConformanceFailureOrDefect<br>OpcenterEXDS_Core.RM_PC_NonConformanceHistory<br>OpcenterEXDS_Core.RM_PC_NonConformanceMtuOrContainer |
| NonConformance_Master | OpcenterEXDS_ShopfloorExecution | OpcenterEXDS_Core.RM_PC_NonConformance |
| ToolDetails_Details | OpcenterEXDS_ShopfloorExecution | OpcenterEXDS_ShopfloorExecution.RM_PC_ToolAutomationNodeInstance<br>OpcenterEXDS_ShopfloorExecution.SetPointToolAssociation<br>OpcenterEXDS_ShopfloorExecution.ToolHistory |
| Tool_Master | OpcenterEXDS_ShopfloorExecution | OpcenterEXDS_Core.Tool |

---

_Report generated by export_manifest tool_
