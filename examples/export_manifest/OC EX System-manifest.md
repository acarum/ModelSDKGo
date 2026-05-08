# Manifest Report: OC EX System

**Mendix Version:** 11.10.0  
**MPR File:** C:\Workspaces\Mendix\MDUI\System_Mendix_CLI\OC EX System.mpr  
**Generated:** 2026-05-08 12:35:47  

---

## Summary

- **External Entities:** 22 (across 3 modules)
- **Microflow/Action Calls:** 78
- **Signal Manager Subscriptions:** 3 subscription(s)
- **Navigation Items:** 13

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

### Module: MyModule

#### Entity: Feature

**Published From:** BoFService

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
| IsOption | Boolean |

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
| System | Counters | OpcenterEXFN_ReferenceData.Counter_Master | Administrator, User |
| System | Numbering Patterns | OpcenterEXFN_ReferenceData.NumberingPattern_Master | Administrator, User |
| System | State Machines | OpcenterEXFN_ReferenceData.StateMachine_Master | Administrator, User |
| System | Statuses | OpcenterEXFN_ReferenceData.Status_Master | Administrator, User |
| System | Status Behavior Definitions | OpcenterEXFN_ReferenceData.StatusBehaviorDefinition_Master | Administrator, User |
| System | Status Definitions | OpcenterEXFN_ReferenceData.StatusDefinition_Master | Administrator, User |
| System | Status Transition Definitions | OpcenterEXFN_ReferenceData.StatusTransitionDefinition_Master | Administrator, User |
| System | Unit of Measures | OpcenterEXFN_ReferenceData.UoM_Master | Administrator, User |
| System | UoM Dimensions | OpcenterEXFN_ReferenceData.UoMDimension_Master_SingleSelection | Administrator, User |
| - | Accounts | Administration.Account_Overview | Administrator |
| - | Cross | - | - |
| Cross | Work Orders | - | - |

---

## 5. Pages/Panels Commands Hierarchy

Microflows and nanoflows called by each page/panel, showing recursive call hierarchy up to 5 levels (in YAML structure). Microflows called transitively are loaded on-the-fly from the database when needed.

**Limitation:** Inline nanoflows (nanoflows embedded directly in pages, not stored as separate Units) are not currently traced.

```yaml
pages:
  - name: PANEL_UpdateUoMDimension
    module: OpcenterEXFN_ReferenceData
    flows:
      - OpcenterEXFN_ReferenceData.ACT_UpdateUoMDimension
    target_commands:
      []

  - name: PANEL_CreateBaseUoMDimension
    module: OpcenterEXFN_ReferenceData
    flows:
      - OpcenterEXFN_ReferenceData.ACT_CreateBaseUoMDimension
    target_commands:
      []

  - name: PANEL_CreateBaseUoM_UoMDimension
    module: OpcenterEXFN_ReferenceData
    flows:
      - OpcenterEXFN_ReferenceData.ACT_CreateBaseUoM_UoMDimension
    target_commands:
      []

  - name: PANEL_UoMDimensionUpdateUoM
    module: OpcenterEXFN_ReferenceData
    flows:
      - OpcenterEXFN_ReferenceData.ACT_UpdateUoM_UoMDimension
    target_commands:
      []

  - name: UoMDimension_Details
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateUoM_UoMDimension_ShowPanel
      - name: OpcenterEXFN_ReferenceData.ACT_UnhideUoM
      - name: OpcenterEXFN_ReferenceData.ACT_HideUoM
      - name: OpcenterEXFN_ReferenceData.ACT_DeleteUoM_SingleSelection
      - name: EXFN_AuditTrailViewer.DS_AuditTrailContext
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateUoMDimension_ShowPanel_SingleSelection
    target_commands:
      []

  - name: UoMDimension_Master_SingleSelection
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_DISW_DesignSystem.ExportGridToExcel_MasterScreen
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateUoMDimension_ShowPanel_SingleSelection
      - name: OpcenterEXFN_ReferenceData.ACT_UnhideUoMDimesnion
      - name: OpcenterEXFN_ReferenceData.ACT_HideUoMDimension
      - name: OpcenterEXFN_ReferenceData.ACT_FreezeUoMDimension
      - name: OpcenterEXFN_ReferenceData.ACT_UnfreezeUoMDimension
      - name: OpcenterEXFN_ReferenceData.ACT_DeleteUoMDimension_SingleSelection
    target_commands:
      []

  - name: UoMDimension_Master_MultiSelection
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.DS_UoM_GetCommandBarContext
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateUoM_ShowPanel_MultiSelection
      - name: OpcenterEXFN_ReferenceData.ACT_DeleteUoM_MultiSelection
    target_commands:
      []

  - name: UoM_Master
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_DISW_DesignSystem.ExportGridToExcel_MasterScreen
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateUoM_ShowPanel_SingleSelection
      - name: OpcenterEXFN_ReferenceData.ACT_UnhideUoM
      - name: OpcenterEXFN_ReferenceData.ACT_HideUoM
      - name: OpcenterEXFN_ReferenceData.ACT_DeleteUoM_SingleSelection
    target_commands:
      []

  - name: PANEL_CreateBaseUoM
    module: OpcenterEXFN_ReferenceData
    flows:
      - OpcenterEXFN_ReferenceData.ACT_CreateBaseUoM
    target_commands:
      []

  - name: PANEL_UpdateUoM
    module: OpcenterEXFN_ReferenceData
    flows:
      - OpcenterEXFN_ReferenceData.ACT_UpdateUoM
    target_commands:
      []

  - name: PANEL_UpdateUoMSubmultiple
    module: OpcenterEXFN_ReferenceData
    flows:
      - OpcenterEXFN_ReferenceData.ACT_UpdateUoMSubmultiple
    target_commands:
      []

  - name: PANEL_CreateUoMSubmultiple
    module: OpcenterEXFN_ReferenceData
    flows:
      - OpcenterEXFN_ReferenceData.ACT_CreateUoMSubmultiple
    target_commands:
      []

  - name: PANEL_CreateUoMFactor
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.DS_UoMFactor_GetTargetUoMList
      - name: OpcenterEXFN_ReferenceData.ACT_CreateUoMFactor
    target_commands:
      []

  - name: PANEL_UpdateUoMFactor
    module: OpcenterEXFN_ReferenceData
    flows:
      - OpcenterEXFN_ReferenceData.ACT_UpdateUoMFactor_1
    target_commands:
      []

  - name: UoM_Details
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.DS_UoMSubmultiple_GetUoMFactorForUoM
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateUoMSubmultiple_ShowPanel
      - name: OpcenterEXFN_ReferenceData.ACT_UnhideUoMSubmultiple
      - name: OpcenterEXFN_ReferenceData.ACT_HideUoMSubmultiple
      - name: OpcenterEXFN_ReferenceData.ACT_DeleteUoMSubmultiple
      - name: OpcenterEXFN_ReferenceData.DS_UoMFactor_ListForUoM
      - name: OpcenterEXFN_ReferenceData.DS_UoMFactor_Refresh
      - name: OpcenterEXFN_ReferenceData.ACT_CreateUoMFactor_ShowPanel
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateUoMFactor_ShowPanel
      - name: OpcenterEXFN_ReferenceData.ACT_DeleteUoMFactor
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateUoM_ShowPanel_SingleSelection
    target_commands:
      []

  - name: NumberingPattern_Master
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_DISW_DesignSystem.ExportGridToExcel_MasterScreen
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateNumberingPattern_ShowPanel
      - name: OpcenterEXFN_ReferenceData.ACT_FreezeNumberingPattern
      - name: OpcenterEXFN_ReferenceData.ACT_UnfreezeNumberingPattern
      - name: OpcenterEXFN_ReferenceData.ACT_PreviewNumberingPattern_ShowPanel
      - name: OpcenterEXFN_ReferenceData.ACT_DeleteNumberingPattern
    target_commands:
      []

  - name: PANEL_PreviewNumberingPattern
    module: OpcenterEXFN_ReferenceData
    flows:
      []
    target_commands:
      []

  - name: PANEL_UpdateNumberingPattern
    module: OpcenterEXFN_ReferenceData
    flows:
      - OpcenterEXFN_ReferenceData.ACT_UpdateNumberingPattern
    target_commands:
      []

  - name: PANEL_CreateNumberingPattern
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.DS_NumberingPatternInfo_EntityName
      - name: OpcenterEXFN_ReferenceData.DS_NumberingPatternInfo_EntityProperty
      - name: OpcenterEXFN_ReferenceData.ACT_CreateNumberingPattern
    target_commands:
      []

  - name: PANEL_UpdateNumberingPatternPart_ValidatePart
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateNumberingPatternPart
      - name: OpcenterEXFN_ReferenceData.EVT_OnChangeValidatePart
    target_commands:
      []

  - name: PANEL_UpdateNumberingPatternPart
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateNumberingPatternPart
      - name: OpcenterEXFN_ReferenceData.EVT_OnSelectionChangeForPrefilledRegularExpression_NumberingPatternPart
    target_commands:
      []

  - name: PANEL_UpdateNumberingPatternForNumberingPatternPart
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateNumberingPatternForNumberingPatternPart
      - name: OpcenterEXFN_ReferenceData.EVT_OnSelectionChangeForPrefilledRegularExpression
    target_commands:
      []

  - name: NumberingPattern_Details
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.DS_NumberingPatternDetails
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateNumberingPatternForNumberingPatternPart_ShowPanel
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateNumberingPatternPart_ShowPanel
      - name: OpcenterEXFN_ReferenceData.ACT_MoveNumberingPatternPartUp
      - name: OpcenterEXFN_ReferenceData.ACT_MoveNumberingPatternPartDown
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateNumberingPatternPart_ValidatePart
      - name: OpcenterEXFN_ReferenceData.ACT_DeleteNumberingPatternPart
      - name: EXFN_AuditTrailViewer.DS_AuditTrailContext
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateNumberingPattern_ShowPanel
      - name: OpcenterEXFN_ReferenceData.ACT_FreezeNumberingPattern
      - name: OpcenterEXFN_ReferenceData.ACT_UnfreezeNumberingPattern
      - name: OpcenterEXFN_ReferenceData.ACT_PreviewNumberingPattern_ShowPanel
    target_commands:
      []

  - name: Counter_Master
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_DISW_DesignSystem.ExportGridToExcel_MasterScreen
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateCounter_ShowPanel
      - name: OpcenterEXFN_ReferenceData.ACT_HideCounter
      - name: OpcenterEXFN_ReferenceData.ACT_UnhideCounter
      - name: OpcenterEXFN_ReferenceData.ACT_FreezeCounter
      - name: OpcenterEXFN_ReferenceData.ACT_UnfreezeCounter
      - name: OpcenterEXFN_ReferenceData.ACT_ResetCounter
      - name: OpcenterEXFN_ReferenceData.ACT_DeleteCounter
    target_commands:
      []

  - name: PANEL_UpdateCounter
    module: OpcenterEXFN_ReferenceData
    flows:
      - OpcenterEXFN_ReferenceData.ACT_UpdateCounter
    target_commands:
      []

  - name: PANEL_CreateCounter
    module: OpcenterEXFN_ReferenceData
    flows:
      - OpcenterEXFN_ReferenceData.ACT_CreateCounter
    target_commands:
      []

  - name: Counter_Details
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: EXFN_AuditTrailViewer.DS_AuditTrailContext
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateCounter_ShowPanel
      - name: OpcenterEXFN_ReferenceData.ACT_HideCounter
      - name: OpcenterEXFN_ReferenceData.ACT_UnhideCounter
      - name: OpcenterEXFN_ReferenceData.ACT_FreezeCounter
      - name: OpcenterEXFN_ReferenceData.ACT_UnfreezeCounter
      - name: OpcenterEXFN_ReferenceData.ACT_ResetCounter
    target_commands:
      []

  - name: StatusDefinition_Master
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_DISW_DesignSystem.ExportGridToExcel_MasterScreen
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateStatusDefinition_ShowPanel
      - name: OpcenterEXFN_ReferenceData.ACT_HideStatusDefinition
      - name: OpcenterEXFN_ReferenceData.ACT_UnhideStatusDefinition
      - name: OpcenterEXFN_ReferenceData.ACT_FreezeStatusDefinition
      - name: OpcenterEXFN_ReferenceData.ACT_UnfreezeStatusDefinition
      - name: OpcenterEXFN_ReferenceData.ACT_DeleteStatusDefinition
    target_commands:
      []

  - name: PANEL_CreateStatusDefinition
    module: OpcenterEXFN_ReferenceData
    flows:
      - OpcenterEXFN_ReferenceData.ACT_CreateStatusDefinition
    target_commands:
      []

  - name: PANEL_UpdateStatusDefinition
    module: OpcenterEXFN_ReferenceData
    flows:
      - OpcenterEXFN_ReferenceData.ACT_UpdateStatusDefinition
    target_commands:
      []

  - name: PANEL_CreateStatusBehaviorDefinition
    module: OpcenterEXFN_ReferenceData
    flows:
      - OpcenterEXFN_ReferenceData.ACT_CreateStatusBehaviorDefinition
    target_commands:
      []

  - name: PANEL_UpdateStatusBehaviorDefinition
    module: OpcenterEXFN_ReferenceData
    flows:
      - OpcenterEXFN_ReferenceData.ACT_UpdateStatusBehaviorDefinition
    target_commands:
      []

  - name: StatusBehaviorDefinition_Details
    module: OpcenterEXFN_ReferenceData
    flows:
      []
    target_commands:
      []

  - name: StatusBehaviorDefinition_Master
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_DISW_DesignSystem.ExportGridToExcel_MasterScreen
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateStatusBehaviorDefinition_ShowPanel
      - name: OpcenterEXFN_ReferenceData.ACT_DeleteStatusBehaviorDefinition
    target_commands:
      []

  - name: PANEL_UpdateStatus
    module: OpcenterEXFN_ReferenceData
    flows:
      - OpcenterEXFN_ReferenceData.ACT_UpdateStatus
    target_commands:
      []

  - name: PANEL_CreateStatus
    module: OpcenterEXFN_ReferenceData
    flows:
      - OpcenterEXFN_ReferenceData.ACT_CreateStatus
    target_commands:
      []

  - name: PANEL_AssociateStatusBehaviorDefinitionsWithStatus
    module: OpcenterEXFN_ReferenceData
    flows:
      - OpcenterEXFN_ReferenceData.ACT_AssociateStatusBehaviorDefinitionsWithStatus
    target_commands:
      []

  - name: Status_Details
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.DS_StatusBehavior
      - name: OpcenterEXFN_ReferenceData.ACT_DisassociateStatusBehaviorDefinitionsFromStatus
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateStatus_ShowPanel
      - name: OpcenterEXFN_ReferenceData.ACT_SetStatusInitial
    target_commands:
      []

  - name: Status_Master
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_DISW_DesignSystem.ExportGridToExcel_MasterScreen
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateStatus_ShowPanel
      - name: OpcenterEXFN_ReferenceData.ACT_SetStatusInitial
      - name: OpcenterEXFN_ReferenceData.ACT_DeleteStatus
    target_commands:
      []

  - name: StateMachine_Master
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_DISW_DesignSystem.ExportGridToExcel_MasterScreen
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateStateMachine_ShowPanel
      - name: OpcenterEXFN_ReferenceData.ACT_HideStateMachine
      - name: OpcenterEXFN_ReferenceData.ACT_UnhideStateMachine
      - name: OpcenterEXFN_ReferenceData.ACT_FreezeStateMachine
      - name: OpcenterEXFN_ReferenceData.ACT_UnfreezeStateMachine
      - name: OpcenterEXFN_ReferenceData.ACT_DeleteStateMachine
    target_commands:
      []

  - name: StateMachine_Details
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: EXFN_Authentication.Signal_Access_Token
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateStatusStateMachine_ShowPanel
      - name: OpcenterEXFN_ReferenceData.ACT_SetStatusAsInitial
      - name: OpcenterEXFN_ReferenceData.ACT_DeleteStatusStateMachine
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateStatusTransitionShowPanel
      - name: OpcenterEXFN_ReferenceData.ACT_DeleteStatusTransition
      - name: EXFN_AuditTrailViewer.DS_AuditTrailContext
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateStateMachine_ShowPanel
      - name: OpcenterEXFN_ReferenceData.ACT_HideStateMachine
      - name: OpcenterEXFN_ReferenceData.ACT_UnhideStateMachine
      - name: OpcenterEXFN_ReferenceData.ACT_FreezeStateMachine
      - name: OpcenterEXFN_ReferenceData.ACT_UnfreezeStateMachine
    target_commands:
      []

  - name: PANEL_UpdateStatusStateMachine
    module: OpcenterEXFN_ReferenceData
    flows:
      - OpcenterEXFN_ReferenceData.ACT_UpdateStatusStateMachine
    target_commands:
      []

  - name: PANEL_UpdateStatusTransition
    module: OpcenterEXFN_ReferenceData
    flows:
      - OpcenterEXFN_ReferenceData.ACT_UpdateStatusTransition
    target_commands:
      []

  - name: PANEL_UpdateStateMachine
    module: OpcenterEXFN_ReferenceData
    flows:
      - OpcenterEXFN_ReferenceData.ACT_UpdateStateMachine
    target_commands:
      []

  - name: PANEL_CreateStatusStateMachine
    module: OpcenterEXFN_ReferenceData
    flows:
      - OpcenterEXFN_ReferenceData.ACT_CreateStatusStateMachine
    target_commands:
      []

  - name: PANEL_CreateStateMachine
    module: OpcenterEXFN_ReferenceData
    flows:
      - OpcenterEXFN_ReferenceData.ACT_CreateStateMachine
    target_commands:
      []

  - name: PANEL_CreateStatusTransition
    module: OpcenterEXFN_ReferenceData
    flows:
      - OpcenterEXFN_ReferenceData.ACT_CreateStatusTransition
    target_commands:
      []

  - name: PANEL_CreateStatusTransitionDefinition
    module: OpcenterEXFN_ReferenceData
    flows:
      - OpcenterEXFN_ReferenceData.ACT_CreateStatusTransitionDefinition
    target_commands:
      []

  - name: StatusTransitionDefinition_Master
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_DISW_DesignSystem.ExportGridToExcel_MasterScreen
      - name: OpcenterEXFN_ReferenceData.ACT_HideStatusTransitionDefinition
      - name: OpcenterEXFN_ReferenceData.ACT_UnhideStatusTransitionDefinition
      - name: OpcenterEXFN_ReferenceData.ACT_FreezeStatusTransitionDefinition
      - name: OpcenterEXFN_ReferenceData.ACT_UnfreezeStatusTransitionDefinition
      - name: OpcenterEXFN_ReferenceData.ACT_DeleteStatusTransitionDefinition
    target_commands:
      []

```

---

_Report generated by export_manifest tool_
