# Manifest Report: OC EX System

**Mendix Version:** 11.9.0  
**Generated:** 2026-04-25 16:03:54  

---

## Summary

- **External Entities:** 156 (across 3 modules)
- **Microflow/Action Calls:** _Coming soon_
- **Signal Manager Widgets:** _Coming soon_

---

## 1. External Entities

External OData entities used in the project, grouped by module.

### Module: EXFN_AuditTrailViewer

Found 1 external entity/entities:

#### Entity: AuditTrailRecord

**Published From:** AuditTrailRecord

| Attribute | Type |
|-----------|------|
| _Id | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| EntityType | String |
| IsLocked | Boolean |
| ToBeCleaned | String |
| RootEntityId | String |
| RootEntityName | String |
| RootEntityType | String |
| ChangedEntityId | String |
| ChangedEntityName | String |
| ChangedEntityType | String |
| RootCommand | String |
| Action | String |
| UserName | String |
| UpdatedOn | DateTime |
| Environment | String |
| ComputerName | String |
| ChangedEntity | String |
| ChangedLargeProperties | String |
| CorrelationId | String |
| ElectronicSignatureId | String |
| Domain | Boolean |
| Ordering | String |
| TransactionId | String |
| AssociatedRows | String |

### Module: OpcenterEXFN_ReferenceData

Found 20 external entity/entities:

#### Entity: UoM

**Published From:** UoM

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

#### Entity: UoMFactor

**Published From:** UoMFactor

| Attribute | Type |
|-----------|------|
| _Id | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| EntityType | String |
| IsLocked | Boolean |
| ToBeCleaned | String |
| K0Multiplier | Decimal |
| Addend | Decimal |
| K1Multiplier | Decimal |
| Exponent | Decimal |
| IsHidden | Boolean |
| IsSystemDefined | Boolean |
| TargetUoM_Id | String |
| SourceUoM_Id | String |

#### Entity: UoMDimension

**Published From:** UoMDimension

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

#### Entity: NumberingPatternFacet

**Published From:** NumberingPatternFacet

| Attribute | Type |
|-----------|------|
| _Id | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| EntityType | String |
| IsLocked | Boolean |
| ToBeCleaned | String |
| NumberingPattern_Id | String |

#### Entity: CustomPart

**Published From:** CustomPart

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
| EntityName | String |
| AppName | String |
| Field | String |

#### Entity: NumberingPatternPartFacet

**Published From:** NumberingPatternPartFacet

| Attribute | Type |
|-----------|------|
| _Id | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| EntityType | String |
| IsLocked | Boolean |
| ToBeCleaned | String |
| NumberingPatternPart_Id | String |

#### Entity: NumberingPatternPart

**Published From:** NumberingPatternPart

| Attribute | Type |
|-----------|------|
| _Id | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| EntityType | String |
| IsLocked | Boolean |
| ToBeCleaned | String |
| Constant | String |
| Timestamp | DateTime |
| ParameterNId | String |
| RegularExpression | String |
| CustomPart | String |
| Sequence | Integer |
| NumberingPattern_Id | String |
| Counter_Id | String |

#### Entity: Counter

**Published From:** Counter

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
| Seed | Integer |
| Increment | Integer |
| MaxValue | Integer/Decimal |
| LeadingZeros | String |
| TimeBasedReset | DateTime |
| IsTransactional | Boolean |
| IsHidden | Boolean |

#### Entity: CounterFacet

**Published From:** CounterFacet

| Attribute | Type |
|-----------|------|
| _Id | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| EntityType | String |
| IsLocked | Boolean |
| ToBeCleaned | String |
| Counter_Id | String |

#### Entity: NumberingPatternEntityInfo

**Published From:** NumberingPatternEntityInfo

| Attribute | Type |
|-----------|------|
| _Id | String |
| IsFrozen | Boolean |
| CreatedOn | DateTime |
| LastUpdatedOn | DateTime |
| EntityType | String |
| IsLocked | Boolean |
| ToBeCleaned | String |
| AppName | String |
| EntityName | String |
| EntityProperty | String |
| NumberingPatternEntityNId_Id | String |

#### Entity: NumberingPatternEntity

**Published From:** NumberingPatternEntity

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

#### Entity: NumberingPatternParameter

**Published From:** NumberingPatternParameter

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
| NumberingPatternEntity_Id | String |

#### Entity: PrefilledRegularExpression

**Published From:** PrefilledRegularExpression

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
| TextualExpression | String |

#### Entity: NumberingPattern

**Published From:** NumberingPattern

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
| EntityTypeNId | String |
| DestinationProperty | String |

#### Entity: Status

**Published From:** Status

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

#### Entity: StatusTransition

**Published From:** StatusTransition

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

#### Entity: StatusDefinition

**Published From:** StatusDefinition

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
| IsHidden | Boolean |
| IsSystemDefined | Boolean |
| Outcome | String |

#### Entity: StatusBehaviorDefinition

**Published From:** StatusBehaviorDefinition

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
| NId | String |

#### Entity: StatusTransitionDefinition

**Published From:** StatusTransitionDefinition

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
| IsHidden | Boolean |
| IsSystemDefined | Boolean |

#### Entity: StateMachine

**Published From:** StateMachine

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

### Module: OpcenterEXFN_ReferenceData_Connector

Found 135 external entity/entities:

#### Entity: CreateBaseUoM

**Published From:** CreateBaseUoM

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Description | String |
| UoMDimensionId | String |

#### Entity: AdditionalArgumentsDefinition

**Published From:** AdditionalArgumentsDefinition

_No attributes_

#### Entity: CreateBaseUoMResponse

**Published From:** CreateBaseUoMResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |
| UoMId | String |

#### Entity: CreateUoMDimension_Command

**Published From:** CreateUoMDimension

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Description | String |

#### Entity: CreateUoMDimensionResponse

**Published From:** CreateUoMDimensionResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |
| UoMDimensionId | String |

#### Entity: CreateUoMSubmultiple_Command

**Published From:** CreateMultipleOrSubmultipleUoM

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Description | String |
| BaseUoMId | String |

#### Entity: CreateMultipleOrSubmultipleUoMResponse

**Published From:** CreateMultipleOrSubmultipleUoMResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |
| UoMId | String |

#### Entity: CreateUoMSubmultiple_UoMFactor

**Published From:** UoMFactorParameterType

| Attribute | Type |
|-----------|------|
| K0Multiplier | Decimal |
| K1Multiplier | Decimal |
| Exponent | Decimal |
| Addend | Decimal |

#### Entity: CreateStatus_Command

**Published From:** CreateStatus

| Attribute | Type |
|-----------|------|
| StatusDefinitionId | String |
| StateMachineId | String |
| Name | String |
| Description | String |
| IsInitial | Boolean |

#### Entity: CreateStatusResponse

**Published From:** CreateStatusResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |
| StatusId | String |

#### Entity: UpdateUoMSubmultiple_Command

**Published From:** UpdateMultipleOrSubmultipleUoM

| Attribute | Type |
|-----------|------|
| _Id | String |
| Name | String |
| Description | String |
| K1Multiplier | Decimal |

#### Entity: UpdateMultipleOrSubmultipleUoMResponse

**Published From:** UpdateMultipleOrSubmultipleUoMResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: CreateStateMachine_Command

**Published From:** CreateStateMachine

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Description | String |

#### Entity: CreateStateMachineResponse

**Published From:** CreateStateMachineResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |
| StateMachineId | String |

#### Entity: CreateStatusTransition_Command

**Published From:** CreateStatusTransition

| Attribute | Type |
|-----------|------|
| SourceStatusId | String |
| TargetStatusId | String |
| DoRaiseEvent | Boolean |
| StatusTransitionDefinitionVerb | String |

#### Entity: CreateStatusTransitionResponse

**Published From:** CreateStatusTransitionResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |
| StatusTransitionId | String |

#### Entity: CreateStatusTransitionDefinition_Command

**Published From:** CreateStatusTransitionDefinition

| Attribute | Type |
|-----------|------|
| Verb | String |

#### Entity: CreateStatusTransitionDefinitionResponse

**Published From:** CreateStatusTransitionDefinitionResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |
| StatusTransitionDefinitionId | String |

#### Entity: UpdateStateMachine_Command

**Published From:** UpdateStateMachine

| Attribute | Type |
|-----------|------|
| _Id | String |
| Name | String |
| Description | String |

#### Entity: UpdateStateMachineResponse_2

**Published From:** UpdateStateMachineResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: UpdateUoMDimension_Command

**Published From:** UpdateUoMDimension

| Attribute | Type |
|-----------|------|
| _Id | String |
| Name | String |
| Description | String |

#### Entity: UpdateUoMDimensionResponse

**Published From:** UpdateUoMDimensionResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: CreateStatusBehaviorDefinition_Command

**Published From:** CreateStatusBehaviorDefinition

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Description | String |

#### Entity: CreateStatusBehaviorDefinitionResponse

**Published From:** CreateStatusBehaviorDefinitionResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |
| StatusBehaviorDefinitionId | String |

#### Entity: UpdateStatusBehaviorDefinition_Command

**Published From:** UpdateStatusBehaviorDefinition

| Attribute | Type |
|-----------|------|
| _Id | String |
| Name | String |
| Description | String |

#### Entity: UpdateStatusBehaviorDefinitionResponse

**Published From:** UpdateStatusBehaviorDefinitionResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: UpdateUoM_UoMDimension_Command

**Published From:** UpdateUoM

| Attribute | Type |
|-----------|------|
| _Id | String |
| Name | String |
| Description | String |

#### Entity: UpdateUoMResponse

**Published From:** UpdateUoMResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: UpdateStatusTransition_Command

**Published From:** UpdateStatusTransition

| Attribute | Type |
|-----------|------|
| _Id | String |
| DoRaiseEvent | Boolean |
| TargetStatusId | String |

#### Entity: UpdateStatusTransitionResponse

**Published From:** UpdateStatusTransitionResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: SetStatusAsInitial_Command

**Published From:** SetStatusAsInitial

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: SetStatusAsInitialResponse

**Published From:** SetStatusAsInitialResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: MoveNumberingPatternPart_Command

**Published From:** MoveNumberingPatternPart

| Attribute | Type |
|-----------|------|
| _Id | String |
| Shift | String |

#### Entity: MoveNumberingPatternPartResponse

**Published From:** MoveNumberingPatternPartResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |
| NumberingPatternPartId | String |

#### Entity: CreateNumberingPattern_Command

**Published From:** CreateNumberingPattern

| Attribute | Type |
|-----------|------|
| EntityTypeNId | String |
| NId | String |
| Name | String |
| Description | String |
| DestinationProperty | String |

#### Entity: CreateNumberingPatternResponse

**Published From:** CreateNumberingPatternResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |
| NumberingPatternId | String |

#### Entity: NumberingPatternPartsParameterType

**Published From:** NumberingPatternPartsParameterType

| Attribute | Type |
|-----------|------|
| Constant | String |
| CounterId | String |
| ParameterNId | String |
| Timestamp | DateTime |
| RegularExpression | String |
| CustomPart | String |

#### Entity: CreateStatusDefinition_Command

**Published From:** CreateStatusDefinition

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Color | String |
| Description | String |
| Outcome | String |

#### Entity: CreateStatusDefinitionResponse

**Published From:** CreateStatusDefinitionResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |
| StatusDefinitionId | String |

#### Entity: UpdateStatusDefinition_Command

**Published From:** UpdateStatusDefinition

| Attribute | Type |
|-----------|------|
| _Id | String |
| Name | String |
| Color | String |
| Description | String |
| Outcome | String |

#### Entity: UpdateStatusDefinitionResponse

**Published From:** UpdateStatusDefinitionResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: UpdateStatus_Command

**Published From:** UpdateStatus

| Attribute | Type |
|-----------|------|
| _Id | String |
| Name | String |
| Color | String |
| Description | String |
| IsInitial | Boolean |
| Outcome | String |

#### Entity: UpdateStatusResponse

**Published From:** UpdateStatusResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: CreateCounter_Command

**Published From:** CreateCounter

| Attribute | Type |
|-----------|------|
| NId | String |
| Name | String |
| Description | String |
| Seed | Integer |
| Increment | Integer |
| MaxValue | Integer/Decimal |
| LeadingZeros | String |
| TimeBasedReset | DateTime |
| IsTransactional | Boolean |

#### Entity: CreateCounterResponse

**Published From:** CreateCounterResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |
| CounterId | String |

#### Entity: UpdateCounter_Command

**Published From:** UpdateCounter

| Attribute | Type |
|-----------|------|
| _Id | String |
| Name | String |
| Description | String |
| Seed | Integer |
| Increment | Integer |
| MaxValue | Integer/Decimal |
| LeadingZeros | String |
| TimeBasedReset | DateTime |
| IsTransactional | Boolean |

#### Entity: UpdateCounterResponse

**Published From:** UpdateCounterResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: UpdateNumberingPattern_Command

**Published From:** UpdateNumberingPattern

| Attribute | Type |
|-----------|------|
| _Id | String |
| Name | String |
| Description | String |

#### Entity: UpdateNumberingPatternResponse

**Published From:** UpdateNumberingPatternResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |
| NumberingPatternPartId | String |

#### Entity: UpdateNumberingPatternPart_Command

**Published From:** UpdateNumberingPatternPart

| Attribute | Type |
|-----------|------|
| _Id | String |
| Constant | String |
| CounterId | String |
| ParameterNId | String |
| Timestamp | DateTime |
| RegularExpression | String |
| CustomPart | String |

#### Entity: UpdateNumberingPatternPartResponse

**Published From:** UpdateNumberingPatternPartResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: AssociateStatusBehaviorDefinitionsWithStatus_Command

**Published From:** AssociateStatusBehaviorDefinitionsWithStatus

| Attribute | Type |
|-----------|------|
| StatusId | String |

#### Entity: AssociateStatusBehaviorDefinitionsWithStatusStatusBehaviorId

**Published From:** AssociateStatusBehaviorDefinitionsWithStatusStatusBehaviorId

| Attribute | Type |
|-----------|------|
| StatusBehaviorId | String |

#### Entity: AssociateStatusBehaviorDefinitionsWithStatusResponse

**Published From:** AssociateStatusBehaviorDefinitionsWithStatusResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: CreateUoMFactor_Command

**Published From:** CreateUoMFactor

| Attribute | Type |
|-----------|------|
| K0Multiplier | Decimal |
| Addend | Decimal |
| K1Multiplier | Decimal |
| Exponent | Decimal |
| TargetUoMId | String |
| SourceUoMId | String |

#### Entity: CreateUoMFactorResponse

**Published From:** CreateUoMFactorResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |
| UoMFactorId | String |

#### Entity: UpdateUoMFactor_Command

**Published From:** UpdateUoMFactor

| Attribute | Type |
|-----------|------|
| _Id | String |
| K0Multiplier | Decimal |
| Addend | Decimal |
| K1Multiplier | Decimal |
| Exponent | Decimal |

#### Entity: UpdateUoMFactorResponse

**Published From:** UpdateUoMFactorResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: DeleteStateMachine_Command

**Published From:** DeleteStateMachine

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: DeleteStateMachineResponse

**Published From:** DeleteStateMachineResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: DeleteStatusTransitionDefinition_Command

**Published From:** DeleteStatusTransitionDefinition

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: DeleteStatusTransitionDefinitionResponse

**Published From:** DeleteStatusTransitionDefinitionResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: DeleteStatus_Command

**Published From:** DeleteStatus

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: DeleteStatusResponse

**Published From:** DeleteStatusResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: DeleteStatusTransition_Command

**Published From:** DeleteStatusTransition

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: DeleteStatusTransitionResponse

**Published From:** DeleteStatusTransitionResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: DeleteStatusBehaviorDefinition_Command

**Published From:** DeleteStatusBehaviorDefinition

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: DeleteStatusBehaviorDefinitionResponse

**Published From:** DeleteStatusBehaviorDefinitionResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: DeleteStatusDefinition_Command

**Published From:** DeleteStatusDefinition

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: DeleteStatusDefinitionResponse

**Published From:** DeleteStatusDefinitionResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: DeleteCounter_Command

**Published From:** DeleteCounter

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: DeleteCounterResponse

**Published From:** DeleteCounterResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: DeleteNumberingPatternPart_Command

**Published From:** DeleteNumberingPatternPart

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: DeleteNumberingPatternPartResponse

**Published From:** DeleteNumberingPatternPartResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: DeleteNumberingPattern_Command

**Published From:** DeleteNumberingPattern

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: DeleteNumberingPatternResponse

**Published From:** DeleteNumberingPatternResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: DeleteUoMFactor_Command

**Published From:** DeleteUoMFactor

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: DeleteUoMFactorResponse

**Published From:** DeleteUoMFactorResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: DeleteUoM_Command

**Published From:** DeleteUoM

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: DeleteUoMResponse

**Published From:** DeleteUoMResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: DeleteUoMDimension_Command

**Published From:** DeleteUoMDimension

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: DeleteUoMDimensionResponse

**Published From:** DeleteUoMDimensionResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: ResetCounter

**Published From:** ResetCounter

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: ResetCounterResponse

**Published From:** ResetCounterResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: HideCounter

**Published From:** HideCounter

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: HideCounterResponse

**Published From:** HideCounterResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: UnhideCounter

**Published From:** UnhideCounter

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: UnhideCounterResponse

**Published From:** UnhideCounterResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: HideUoM

**Published From:** HideUoM

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: HideUoMResponse

**Published From:** HideUoMResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: UnhideUoM

**Published From:** UnhideUoM

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: UnhideUoMResponse

**Published From:** UnhideUoMResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: HideUoMDimension

**Published From:** HideUoMDimension

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: HideUoMDimensionResponse

**Published From:** HideUoMDimensionResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: UnhideUoMDimension

**Published From:** UnhideUoMDimension

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: UnhideUoMDimensionResponse

**Published From:** UnhideUoMDimensionResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: HideStateMachine

**Published From:** HideStateMachine

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: HideStateMachineResponse

**Published From:** HideStateMachineResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: UnhideStateMachine

**Published From:** UnhideStateMachine

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: UnhideStateMachineResponse

**Published From:** UnhideStateMachineResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: HideStatusDefinition

**Published From:** HideStatusDefinition

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: HideStatusDefinitionResponse

**Published From:** HideStatusDefinitionResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: UnhideStatusDefinition

**Published From:** UnhideStatusDefinition

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: UnhideStatusDefinitionResponse

**Published From:** UnhideStatusDefinitionResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: HideStatusTransitionDefinition

**Published From:** HideStatusTransitionDefinition

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: HideStatusTransitionDefinitionResponse

**Published From:** HideStatusTransitionDefinitionResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: UnhideStatusTransitionDefinition

**Published From:** UnhideStatusTransitionDefinition

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: UnhideStatusTransitionDefinitionResponse

**Published From:** UnhideStatusTransitionDefinitionResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: FreezeCounter

**Published From:** FreezeCounter

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: FreezeCounterResponse

**Published From:** FreezeCounterResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: UnfreezeCounter

**Published From:** UnfreezeCounter

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: UnfreezeCounterResponse

**Published From:** UnfreezeCounterResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: FreezeUoMDimension

**Published From:** FreezeUoMDimension

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: FreezeUoMDimensionResponse

**Published From:** FreezeUoMDimensionResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: UnfreezeUoMDimension

**Published From:** UnfreezeUoMDimension

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: UnfreezeUoMDimensionResponse

**Published From:** UnfreezeUoMDimensionResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: FreezeNumberingPattern

**Published From:** FreezeNumberingPattern

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: FreezeNumberingPatternResponse

**Published From:** FreezeNumberingPatternResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: UnfreezeNumberingPattern

**Published From:** UnfreezeNumberingPattern

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: UnfreezeNumberingPatternResponse

**Published From:** UnfreezeNumberingPatternResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: FreezeStatusDefinition

**Published From:** FreezeStatusDefinition

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: FreezeStatusDefinitionResponse

**Published From:** FreezeStatusDefinitionResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: UnfreezeStatusDefinition

**Published From:** UnfreezeStatusDefinition

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: UnfreezeStatusDefinitionResponse

**Published From:** UnfreezeStatusDefinitionResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: FreezeStateMachine

**Published From:** FreezeStateMachine

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: FreezeStateMachineResponse

**Published From:** FreezeStateMachineResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: UnfreezeStateMachine

**Published From:** UnfreezeStateMachine

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: UnfreezeStateMachineResponse

**Published From:** UnfreezeStateMachineResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: FreezeStatusTransitionDefinition

**Published From:** FreezeStatusTransitionDefinition

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: FreezeStatusTransitionDefinitionResponse

**Published From:** FreezeStatusTransitionDefinitionResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: UnfreezeStatusTransitionDefinition

**Published From:** UnfreezeStatusTransitionDefinition

| Attribute | Type |
|-----------|------|
| _Id | String |

#### Entity: UnfreezeStatusTransitionDefinitionResponse

**Published From:** UnfreezeStatusTransitionDefinitionResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

#### Entity: DisassociateStatusBehaviorDefinitionsFromStatus

**Published From:** DisassociateStatusBehaviorDefinitionsFromStatus

| Attribute | Type |
|-----------|------|
| StatusId | String |

#### Entity: DisassociateStatusBehaviorDefinitionsFromStatusStatusBehaviorId

**Published From:** DisassociateStatusBehaviorDefinitionsFromStatusStatusBehaviorId

| Attribute | Type |
|-----------|------|
| StatusBehaviorId | String |

#### Entity: DisassociateStatusBehaviorDefinitionsFromStatusResponse

**Published From:** DisassociateStatusBehaviorDefinitionsFromStatusResponse

| Attribute | Type |
|-----------|------|
| Succeeded | Boolean |

---

## 2. Microflow/Action Calls

_Section not yet implemented._

---

## 3. Signal Manager Widgets

_Section not yet implemented._

---

_Report generated by export_manifest tool_
