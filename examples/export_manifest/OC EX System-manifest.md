# Manifest Report: OC EX System

**Mendix Version:** 11.10.0  
**MPR File:** C:\Workspaces\Mendix\MDUI\System_Mendix_CLI\OC EX System.mpr  
**Generated:** 2026-05-04 11:40:00  

---

## Summary

- **External Entities:** 21 (across 2 modules)
- **Microflow/Action Calls:** 78
- **Signal Manager Subscriptions:** 3 subscription(s)
- **Navigation Items:** 10
- **Pages with Commands:** 6 page(s) analyzed
- **Command Buttons:** 38 total, 34 with extracted commands

---

## 1. External Entities

External OData entities used in the project, grouped by module.

### Module: EXFN_AuditTrailViewer

#### Entity: AuditTrailRecord

**Published From:** AuditTrailViewer

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

#### Entity: UoM

**Published From:** Reference

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

**Published From:** Reference

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

**Published From:** Reference

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

**Published From:** Reference

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

**Published From:** Reference

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

**Published From:** Reference

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

**Published From:** Reference

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

**Published From:** Reference

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

**Published From:** Reference

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

**Published From:** Reference

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

**Published From:** Reference

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

**Published From:** Reference

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

**Published From:** Reference

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

**Published From:** Reference

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

**Published From:** Reference

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

**Published From:** Reference

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

**Published From:** Reference

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

**Published From:** Reference

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

**Published From:** Reference

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

**Published From:** Reference

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

## 2. Microflow/Action Calls

Microflow and action calls found in the project (MicroflowCall, JavaAction, ExternalAction).

Found 78 call(s):

| Microflow | Module | Call Type | AppName | CommandName |
|-----------|--------|-----------|---------|-------------|
| Microflow | Documents | ExternalAction | TripPinServiceRW | ResetDataSource |
| UpdateStatus | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | UpdateStatus |
| SUB_SetStatusAsInitial | OpcenterEXFN_ReferenceData | ExternalAction | Reference | SetStatusAsInitial |
| SUB_DeleteCommand_StartArrayWithTwoParameter | OpcenterEXFN_ReferenceData_Connector | MicroflowCall | <parameter> | <parameter> |
| SUB_HideAndUnhideCommand | OpcenterEXFN_ReferenceData_Connector | MicroflowCall | <parameter> | <parameter> |
| SUB_HideAndUnhideCommand | OpcenterEXFN_ReferenceData_Connector | JavaAction | 'AppName' | 'CommandName' |
| SUB_DeleteCommand | OpcenterEXFN_ReferenceData_Connector | MicroflowCall | <parameter> | <parameter> |
| SUB_FreezeAndUnfreezeCommand | OpcenterEXFN_ReferenceData_Connector | MicroflowCall | <parameter> | <parameter> |
| SUB_CallCommand | OpcenterEXFN_ReferenceData_Connector | MicroflowCall | <parameter> | <parameter> |
| HideStatusDefinition | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | HideStatusDefinition |
| UnfreezeStatusDefinition | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | UnfreezeStatusDefinition |
| FreezeStatusDefinition | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | FreezeStatusDefinition |
| UnhideStatusDefinition | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | UnhideStatusDefinition |
| DeleteStatusDefinition | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | DeleteStatusDefinition |
| CreateStatusDefinition | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | CreateStatusDefinition |
| UpdateStatusDefinition | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | UpdateStatusDefinition |
| FreezeStatusTransitionDefinition | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | FreezeStatusTransitionDefinition |
| UnfreezeStatusTransitionDefinition | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | UnfreezeStatusTransitionDefinition |
| CreateStatusTransitionDefinition | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | CreateStatusTransitionDefinition |
| HideStatusTransitionDefinition | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | HideStatusTransitionDefinition |
| UnhideStatusTransitionDefinition | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | UnhideStatusTransitionDefinition |
| DeleteStatusTransitionDefinition | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | DeleteStatusTransitionDefinition |
| CreateStatusBehaviorDefinition | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | CreateStatusBehaviorDefinition |
| UpdateStatusBehaviorDefinition | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | UpdateStatusBehaviorDefinition |
| DeleteStatusBehaviorDefinition | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | DeleteStatusBehaviorDefinition |
| FreezeStateMachine | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | FreezeStateMachine |
| UpdateStatusTransition | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | UpdateStatusTransition |
| CreateStatusStateMachine | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | CreateStatus |
| DeleteStateMachine | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | DeleteStateMachine |
| UpdateStateMachine | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | UpdateStateMachine |
| HideStateMachine | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | HideStateMachine |
| UnfreezeStateMachine | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | UnfreezeStateMachine |
| CreateStatusTransition | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | CreateStatusTransition |
| CreateStateMachine | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | CreateStateMachine |
| UnhideStateMachine | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | UnhideStateMachine |
| DeleteStatusTransition | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | DeleteStatusTransition |
| AssociateStatusBehaviorDefinitionsWithStatus | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | AssociateStatusBehaviorDefinitionsWithStatus |
| UpdateStatus | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | UpdateStatus |
| CreateStatus | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | CreateStatus |
| DeleteStatus | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | DeleteStatus |
| DisassociateStatusBehaviorDefinitionsFromStatus | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | DisassociateStatusBehaviorDefinitionsFromStatus |
| UnfreezeCounter | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | UnfreezeCounter |
| ResetCounter | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | ResetCounter |
| DeleteCounter | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | DeleteCounter |
| DeleteCounter | OpcenterEXFN_ReferenceData_Connector | MicroflowCall | 'AppName'
 | 'CommandName'
 |
| UpdateCounter | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | UpdateCounter |
| HideCounter | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | HideCounter |
| CreateCounter | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | CreateCounter |
| FreezeCounter | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | FreezeCounter |
| UnhideCounter | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | UnhideCounter |
| UpdateNumberingPattern | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | UpdateNumberingPattern |
| DeleteNumberingPatternPart | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | DeleteNumberingPatternPart |
| DeleteNumberingPattern | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | DeleteNumberingPattern |
| MoveNumberingPatternPart | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | MoveNumberingPatternPart |
| UpdateNumberingPatternForNumberingPatternPart | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | UpdateNumberingPattern |
| UnfreezeNumberingPattern | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | UnfreezeNumberingPattern |
| UpdateNumberingPatternPart | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | UpdateNumberingPatternPart |
| FreezeNumberingPattern | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | FreezeNumberingPattern |
| CreateNumberingPattern | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | CreateNumberingPattern |
| HideUoM | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | HideUoM |
| UpdateUoMSubmultiple | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | UpdateMultipleOrSubmultipleUoM |
| DeleteUoMFactor | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | DeleteUoMFactor |
| UpdateUoM | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | UpdateUoM |
| UpdateUoMFactor | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | UpdateUoMFactor |
| CreateBaseUoM | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | CreateBaseUoM |
| UnhideUoM | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | UnhideUoM |
| DeleteUoM | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | DeleteUoM |
| CreateUoMFactor | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | CreateUoMFactor |
| CreateUoMSubmultiple | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | CreateMultipleOrSubmultipleUoM |
| DeleteUoMDimension | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | DeleteUoMDimension |
| UpdateUoM_UoMDimension | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | UpdateUoM |
| CreateBaseUoMDimension | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | CreateUoMDimension |
| UnfreezeUoMDimension | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | UnfreezeUoMDimension |
| UnhideUoMDImension | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | UnhideUoMDimension |
| HideUoMDImension | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | HideUoMDimension |
| UpdateUoMDimension | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | UpdateUoMDimension |
| CreateBaseUoM_UoMDimension | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | CreateBaseUoM |
| FreezeUoMDimension | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | FreezeUoMDimension |

---

## 3. Signal Manager Subscriptions

Signal subscriptions from Signal Manager widgets (siemens.mxtosignal.MxToSignal).

Found 3 subscription(s):

| Module | Document Type | Document | Signal Name | App Name | Subscription Filter |
|--------|---------------|----------|-------------|----------|---------------------|
| OpcenterEXFN_ReferenceData | Page | StateMachine_Details | SN | APPName | Yes |
| OpcenterEXFN_ReferenceData | Page | StateMachine_Details | SN2 | AN2 | Yes |
| OpcenterEXFN_ReferenceData | Snippet | MySnippet | SNIPPET_SN | SNIPPET_APPName | Yes |

---

## 4. Navigation Items

| Parent Node | Node | Target Page | User Roles |
|-------------|------|-------------|------------|
| - | System | - | - |
| System | Counters | OpcenterEXFN_ReferenceData.Counter_Master | Administrator, User, Manager |
| System | Numbering Patterns | OpcenterEXFN_ReferenceData.NumberingPattern_Master | Administrator, User, Manager |
| System | State Machines | OpcenterEXFN_ReferenceData.StateMachine_Master | Administrator, User, Manager |
| System | Statuses | OpcenterEXFN_ReferenceData.Status_Master | Administrator, User, Manager |
| System | Status Behavior Definitions | OpcenterEXFN_ReferenceData.StatusBehaviorDefinition_Master | Administrator, User, Manager |
| System | Status Definitions | OpcenterEXFN_ReferenceData.StatusDefinition_Master | Administrator, User, Manager |
| System | Status Transition Definitions | OpcenterEXFN_ReferenceData.StatusTransitionDefinition_Master | Administrator, User, Manager |
| System | Unit of Measures | OpcenterEXFN_ReferenceData.UoM_Master | Administrator, User, Manager |
| System | UoM Dimensions | OpcenterEXFN_ReferenceData.UoMDimension_Master_SingleSelection | Administrator, User, Manager |

---

## 5. Page Commands

Command bar actions extracted from navigation pages. Shows buttons in the vertical command bar of the Right placeholder.

Found commands in 6 page(s):

### OpcenterEXFN_ReferenceData.StatusTransitionDefinition_Master

| Caption | Target Page | Target Command |
|---------|-------------|----------------|
| Create | OpcenterEXFN_ReferenceData.PANEL_CreateStatusTransitionDefinition | CreateStatusTransitionDefinition |
| Hide | - | HideStatusTransitionDefinition |
| Unhide | - | UnhideStatusTransitionDefinition |
| Freeze | - | FreezeStatusTransitionDefinition |
| Unfreeze | - | UnfreezeStatusTransitionDefinition |
| Delete | - | DeleteStatusTransitionDefinition |

### OpcenterEXFN_ReferenceData.UoMDimension_Master_SingleSelection

| Caption | Target Page | Target Command |
|---------|-------------|----------------|
| Create | OpcenterEXFN_ReferenceData.PANEL_CreateBaseUoMDimension | CreateUoMDimension |
| Details | OpcenterEXFN_ReferenceData.UoMDimension_Details | - |
| Edit | OpcenterEXFN_ReferenceData.PANEL_UpdateUoMDimension | UpdateUoMDimension |
| Unhide | - | UnhideUoMDimension |
| Hide | - | HideUoMDimension |
| Freeze | - | FreezeUoMDimension |
| Unfreeze | - | UnfreezeUoMDimension |
| Delete | - | DeleteUoMDimension |

### OpcenterEXFN_ReferenceData.StateMachine_Master

| Caption | Target Page | Target Command |
|---------|-------------|----------------|
| Create | OpcenterEXFN_ReferenceData.PANEL_CreateStateMachine | CreateStateMachine |
| Details | OpcenterEXFN_ReferenceData.StateMachine_Details | - |
| Edit | OpcenterEXFN_ReferenceData.PANEL_UpdateStateMachine | UpdateStateMachine |
| Hide | - | HideStateMachine |
| Unhide | - | UnhideStateMachine |
| Freeze | - | FreezeStateMachine |
| Unfreeze | - | UnfreezeStateMachine |
| Delete | - | DeleteStateMachine |

### OpcenterEXFN_ReferenceData.Status_Master

| Caption | Target Page | Target Command |
|---------|-------------|----------------|
| Create | OpcenterEXFN_ReferenceData.PANEL_CreateStatus | CreateStatus |
| Details | - | - |
| Edit | OpcenterEXFN_ReferenceData.PANEL_UpdateStatus | UpdateStatus |
| Initial | - | SetStatusAsInitial |
| Delete | - | DeleteStatus |

### OpcenterEXFN_ReferenceData.StatusBehaviorDefinition_Master

| Caption | Target Page | Target Command |
|---------|-------------|----------------|
| Create | OpcenterEXFN_ReferenceData.PANEL_CreateStatusBehaviorDefinition | CreateStatusBehaviorDefinition |
| Details | OpcenterEXFN_ReferenceData.StatusBehaviorDefinition_Details | - |
| Edit | OpcenterEXFN_ReferenceData.PANEL_UpdateStatusBehaviorDefinition | UpdateStatusBehaviorDefinition |
| Delete | - | DeleteStatusBehaviorDefinition |

### OpcenterEXFN_ReferenceData.StatusDefinition_Master

| Caption | Target Page | Target Command |
|---------|-------------|----------------|
| Create | OpcenterEXFN_ReferenceData.PANEL_CreateStatusDefinition | CreateStatusDefinition |
| Edit | OpcenterEXFN_ReferenceData.PANEL_UpdateStatusDefinition | UpdateStatusDefinition |
| Hide | - | HideStatusDefinition |
| Unhide | - | UnhideStatusDefinition |
| Freeze | - | FreezeStatusDefinition |
| Unfreeze | - | UnfreezeStatusDefinition |
| Delete | - | DeleteStatusDefinition |

---

_Report generated by export_manifest tool_
