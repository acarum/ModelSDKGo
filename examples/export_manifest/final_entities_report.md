# Manifest Report: OC EX System

**Mendix Version:** 11.9.0  
**Generated:** 2026-04-26 13:26:44  

---

## Summary

- **External Entities:** 21 (across 2 modules)

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

---

_Report generated by export_manifest tool_
