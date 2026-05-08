# Manifest Report: OC EX System

**Mendix Version:** 11.10.0  
**MPR File:** C:\Workspaces\Mendix\MDUI\System_Mendix_CLI\OC EX System.mpr  
**Generated:** 2026-05-08 17:08:20  

---

## Summary

- **External Entities:** 22 (across 3 modules)
- **Microflow/Action Calls:** 80
- **Signal Manager Subscriptions:** 3 subscription(s)
- **Navigation Items:** 13
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

Found 80 call(s):

| Microflow | Module | Call Type | AppName | CommandName |
|-----------|--------|-----------|---------|-------------|
| Microflow | Documents | ExternalAction | TripPinServiceRW | ResetDataSource |
| UpdateStatus | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | UpdateStatus |
| SUB_SetStatusAsInitial | OpcenterEXFN_ReferenceData | ExternalAction | Reference | SetStatusAsInitial |
| SUB_DeleteCommand_StartArrayWithTwoParameter | OpcenterEXFN_ReferenceData_Connector | MicroflowCall | <parameter> | <parameter> |
| SUB_HideAndUnhideCommand | OpcenterEXFN_ReferenceData_Connector | MicroflowCall | <parameter> | <parameter> |
| SUB_HideAndUnhideCommand | OpcenterEXFN_ReferenceData_Connector | JavaAction | <JavaAppName> | <JavaCommandName> |
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
| UpdateUoMFactor_Action | OpcenterEXFN_ReferenceData_Connector | ExternalAction | Reference | UpdateUoMFactor |
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
| UpdateUoMFactor_MF | OpcenterEXFN_ReferenceData_Connector | MicroflowCall | 'APPNAME'
 | 'COMMANDNAME'
 |
| UpdateUoMFactor_MF_JAVA | OpcenterEXFN_ReferenceData_Connector | JavaAction | <JavaAppName> | <JavaCommandName> |

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

## 5. Page Commands

Command bar actions extracted from navigation pages. Shows buttons in the vertical command bar of the Right placeholder.

Found commands in 6 page(s):

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

### OpcenterEXFN_ReferenceData.StatusTransitionDefinition_Master

| Caption | Target Page | Target Command |
|---------|-------------|----------------|
| Create | OpcenterEXFN_ReferenceData.PANEL_CreateStatusTransitionDefinition | CreateStatusTransitionDefinition |
| Hide | - | HideStatusTransitionDefinition |
| Unhide | - | UnhideStatusTransitionDefinition |
| Freeze | - | FreezeStatusTransitionDefinition |
| Unfreeze | - | UnfreezeStatusTransitionDefinition |
| Delete | - | DeleteStatusTransitionDefinition |

---

## 6. Pages/Panels Commands Hierarchy

Microflows and nanoflows called by each page/panel, showing recursive call hierarchy up to 5 levels (in YAML structure). Microflows called transitively are loaded on-the-fly from the database when needed.

**Limitation:** Inline nanoflows (nanoflows embedded directly in pages, not stored as separate Units) are not currently traced.

```yaml
pages:
  - name: PANEL_UpdateUoMDimension
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateUoMDimension
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.UpdateUoMDimension
    target_commands:
      - Reference.UpdateUoMDimension

  - name: PANEL_CreateBaseUoMDimension
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_CreateBaseUoMDimension
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.CreateBaseUoMDimension
    target_commands:
      - Reference.CreateUoMDimension

  - name: PANEL_CreateBaseUoM_UoMDimension
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_CreateBaseUoM_UoMDimension
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.CreateBaseUoM_UoMDimension
    target_commands:
      - Reference.CreateBaseUoM

  - name: PANEL_UoMDimensionUpdateUoM
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateUoM_UoMDimension
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.UpdateUoM_UoMDimension
    target_commands:
      - Reference.UpdateUoM

  - name: UoMDimension_Details
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_UnhideUoM
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.UnhideUoM
      - name: OpcenterEXFN_ReferenceData.ACT_HideUoM
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.HideUoM
      - name: OpcenterEXFN_ReferenceData.ACT_DeleteUoM_SingleSelection
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.DeleteUoM
    target_commands:
      - Reference.UnhideUoM
      - Reference.HideUoM
      - Reference.DeleteUoM

  - name: UoMDimension_Master_SingleSelection
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_UnhideUoMDimesnion
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.UnhideUoMDImension
      - name: OpcenterEXFN_ReferenceData.ACT_HideUoMDimension
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.HideUoMDImension
      - name: OpcenterEXFN_ReferenceData.ACT_FreezeUoMDimension
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.FreezeUoMDimension
      - name: OpcenterEXFN_ReferenceData.ACT_UnfreezeUoMDimension
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.UnfreezeUoMDimension
      - name: OpcenterEXFN_ReferenceData.ACT_DeleteUoMDimension_SingleSelection
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.DeleteUoMDimension
    target_commands:
      - Reference.UnhideUoMDimension
      - Reference.HideUoMDimension
      - Reference.FreezeUoMDimension
      - Reference.UnfreezeUoMDimension
      - Reference.DeleteUoMDimension

  - name: UoMDimension_Master_MultiSelection
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_DeleteUoM_MultiSelection
        calls:
          - name: OpcenterEXFN_ReferenceData.ACT_DeleteUoM_SingleSelection
            - name: OpcenterEXFN_ReferenceData_Connector.DeleteUoM
    target_commands:
      - Reference.DeleteUoM

  - name: UoM_Master
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_UnhideUoM
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.UnhideUoM
      - name: OpcenterEXFN_ReferenceData.ACT_HideUoM
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.HideUoM
      - name: OpcenterEXFN_ReferenceData.ACT_DeleteUoM_SingleSelection
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.DeleteUoM
    target_commands:
      - Reference.UnhideUoM
      - Reference.HideUoM
      - Reference.DeleteUoM

  - name: PANEL_CreateBaseUoM
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_CreateBaseUoM
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.CreateBaseUoM
    target_commands:
      - Reference.CreateBaseUoM

  - name: PANEL_UpdateUoM
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateUoM
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.UpdateUoM
    target_commands:
      - Reference.UpdateUoM

  - name: PANEL_UpdateUoMSubmultiple
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateUoMSubmultiple
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.UpdateUoMSubmultiple
    target_commands:
      - Reference.UpdateMultipleOrSubmultipleUoM

  - name: PANEL_CreateUoMSubmultiple
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_CreateUoMSubmultiple
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.CreateUoMSubmultiple
    target_commands:
      - Reference.CreateMultipleOrSubmultipleUoM

  - name: PANEL_CreateUoMFactor
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_CreateUoMFactor
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.CreateUoMFactor
    target_commands:
      - Reference.CreateUoMFactor

  - name: PANEL_UpdateUoMFactor
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateUoMFactor_1
        calls:
          - name: OpcenterEXFN_ReferenceData.ACT_UpdateUoMFactor_2
            - name: OpcenterEXFN_ReferenceData_Connector.UpdateUoMFactor_Action
            - name: OpcenterEXFN_ReferenceData_Connector.UpdateUoMFactor_MF
            - name: OpcenterEXFN_ReferenceData_Connector.UpdateUoMFactor_MF_JAVA
    target_commands:
      - Reference.UpdateUoMFactor
      - APPNAME.COMMANDNAME
      - EXFN_ServiceLayer.CallCommandAction

  - name: UoM_Details
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_UnhideUoMSubmultiple
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.UnhideUoM
      - name: OpcenterEXFN_ReferenceData.ACT_HideUoMSubmultiple
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.HideUoM
      - name: OpcenterEXFN_ReferenceData.ACT_DeleteUoMSubmultiple
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.DeleteUoM
      - name: OpcenterEXFN_ReferenceData.ACT_DeleteUoMFactor
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.DeleteUoMFactor
    target_commands:
      - Reference.UnhideUoM
      - Reference.HideUoM
      - Reference.DeleteUoM
      - Reference.DeleteUoMFactor

  - name: NumberingPattern_Master
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_FreezeNumberingPattern
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.FreezeNumberingPattern
      - name: OpcenterEXFN_ReferenceData.ACT_UnfreezeNumberingPattern
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.UnfreezeNumberingPattern
      - name: OpcenterEXFN_ReferenceData.ACT_DeleteNumberingPattern
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.DeleteNumberingPattern
    target_commands:
      - Reference.FreezeNumberingPattern
      - Reference.UnfreezeNumberingPattern
      - Reference.DeleteNumberingPattern

  - name: PANEL_PreviewNumberingPattern
    module: OpcenterEXFN_ReferenceData
    flows:
      []
    target_commands:
      []

  - name: PANEL_UpdateNumberingPattern
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateNumberingPattern
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.UpdateNumberingPattern
    target_commands:
      - Reference.UpdateNumberingPattern

  - name: PANEL_CreateNumberingPattern
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_CreateNumberingPattern
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.CreateNumberingPattern
    target_commands:
      - Reference.CreateNumberingPattern

  - name: PANEL_UpdateNumberingPatternPart_ValidatePart
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateNumberingPatternPart
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.UpdateNumberingPatternPart
    target_commands:
      - Reference.UpdateNumberingPatternPart

  - name: PANEL_UpdateNumberingPatternPart
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateNumberingPatternPart
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.UpdateNumberingPatternPart
    target_commands:
      - Reference.UpdateNumberingPatternPart

  - name: PANEL_UpdateNumberingPatternForNumberingPatternPart
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateNumberingPatternForNumberingPatternPart
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.UpdateNumberingPatternForNumberingPatternPart
    target_commands:
      - Reference.UpdateNumberingPattern

  - name: NumberingPattern_Details
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_MoveNumberingPatternPartUp
        calls:
          - name: OpcenterEXFN_ReferenceData.ACT_MoveNumberingPatternPart
            - name: OpcenterEXFN_ReferenceData_Connector.MoveNumberingPatternPart
      - name: OpcenterEXFN_ReferenceData.ACT_MoveNumberingPatternPartDown
        calls:
          - name: OpcenterEXFN_ReferenceData.ACT_MoveNumberingPatternPart
            - name: OpcenterEXFN_ReferenceData_Connector.MoveNumberingPatternPart
      - name: OpcenterEXFN_ReferenceData.ACT_DeleteNumberingPatternPart
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.DeleteNumberingPatternPart
      - name: OpcenterEXFN_ReferenceData.ACT_FreezeNumberingPattern
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.FreezeNumberingPattern
      - name: OpcenterEXFN_ReferenceData.ACT_UnfreezeNumberingPattern
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.UnfreezeNumberingPattern
    target_commands:
      - Reference.MoveNumberingPatternPart
      - Reference.DeleteNumberingPatternPart
      - Reference.FreezeNumberingPattern
      - Reference.UnfreezeNumberingPattern

  - name: Counter_Master
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_HideCounter
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.HideCounter
      - name: OpcenterEXFN_ReferenceData.ACT_UnhideCounter
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.UnhideCounter
      - name: OpcenterEXFN_ReferenceData.ACT_FreezeCounter
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.FreezeCounter
      - name: OpcenterEXFN_ReferenceData.ACT_UnfreezeCounter
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.UnfreezeCounter
      - name: OpcenterEXFN_ReferenceData.ACT_ResetCounter
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.ResetCounter
      - name: OpcenterEXFN_ReferenceData.ACT_DeleteCounter
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.DeleteCounter
    target_commands:
      - Reference.HideCounter
      - Reference.UnhideCounter
      - Reference.FreezeCounter
      - Reference.UnfreezeCounter
      - Reference.ResetCounter
      - Reference.DeleteCounter
      - AppName.CommandName

  - name: PANEL_UpdateCounter
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateCounter
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.UpdateCounter
    target_commands:
      - Reference.UpdateCounter

  - name: PANEL_CreateCounter
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_CreateCounter
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.CreateCounter
    target_commands:
      - Reference.CreateCounter

  - name: Counter_Details
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_HideCounter
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.HideCounter
      - name: OpcenterEXFN_ReferenceData.ACT_UnhideCounter
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.UnhideCounter
      - name: OpcenterEXFN_ReferenceData.ACT_FreezeCounter
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.FreezeCounter
      - name: OpcenterEXFN_ReferenceData.ACT_UnfreezeCounter
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.UnfreezeCounter
      - name: OpcenterEXFN_ReferenceData.ACT_ResetCounter
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.ResetCounter
    target_commands:
      - Reference.HideCounter
      - Reference.UnhideCounter
      - Reference.FreezeCounter
      - Reference.UnfreezeCounter
      - Reference.ResetCounter

  - name: StatusDefinition_Master
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_HideStatusDefinition
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.HideStatusDefinition
      - name: OpcenterEXFN_ReferenceData.ACT_UnhideStatusDefinition
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.UnhideStatusDefinition
      - name: OpcenterEXFN_ReferenceData.ACT_FreezeStatusDefinition
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.FreezeStatusDefinition
      - name: OpcenterEXFN_ReferenceData.ACT_UnfreezeStatusDefinition
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.UnfreezeStatusDefinition
      - name: OpcenterEXFN_ReferenceData.ACT_DeleteStatusDefinition
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.DeleteStatusDefinition
    target_commands:
      - Reference.HideStatusDefinition
      - Reference.UnhideStatusDefinition
      - Reference.FreezeStatusDefinition
      - Reference.UnfreezeStatusDefinition
      - Reference.DeleteStatusDefinition

  - name: PANEL_CreateStatusDefinition
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_CreateStatusDefinition
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.CreateStatusDefinition
    target_commands:
      - Reference.CreateStatusDefinition

  - name: PANEL_UpdateStatusDefinition
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateStatusDefinition
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.UpdateStatusDefinition
    target_commands:
      - Reference.UpdateStatusDefinition

  - name: PANEL_CreateStatusBehaviorDefinition
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_CreateStatusBehaviorDefinition
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.CreateStatusBehaviorDefinition
    target_commands:
      - Reference.CreateStatusBehaviorDefinition

  - name: PANEL_UpdateStatusBehaviorDefinition
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateStatusBehaviorDefinition
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.UpdateStatusBehaviorDefinition
    target_commands:
      - Reference.UpdateStatusBehaviorDefinition

  - name: StatusBehaviorDefinition_Details
    module: OpcenterEXFN_ReferenceData
    flows:
      []
    target_commands:
      []

  - name: StatusBehaviorDefinition_Master
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_DeleteStatusBehaviorDefinition
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.DeleteStatusBehaviorDefinition
    target_commands:
      - Reference.DeleteStatusBehaviorDefinition

  - name: PANEL_UpdateStatus
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateStatus
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.UpdateStatus
    target_commands:
      - Reference.UpdateStatus

  - name: PANEL_CreateStatus
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_CreateStatus
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.CreateStatus
    target_commands:
      - Reference.CreateStatus

  - name: PANEL_AssociateStatusBehaviorDefinitionsWithStatus
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_AssociateStatusBehaviorDefinitionsWithStatus
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.AssociateStatusBehaviorDefinitionsWithStatus
    target_commands:
      - Reference.AssociateStatusBehaviorDefinitionsWithStatus

  - name: Status_Details
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_DisassociateStatusBehaviorDefinitionsFromStatus
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.DisassociateStatusBehaviorDefinitionsFromStatus
      - name: OpcenterEXFN_ReferenceData.ACT_SetStatusInitial
        calls:
          - name: OpcenterEXFN_ReferenceData.SUB_SetStatusAsInitial
    target_commands:
      - Reference.DisassociateStatusBehaviorDefinitionsFromStatus
      - Reference.SetStatusAsInitial

  - name: Status_Master
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_SetStatusInitial
        calls:
          - name: OpcenterEXFN_ReferenceData.SUB_SetStatusAsInitial
      - name: OpcenterEXFN_ReferenceData.ACT_DeleteStatus
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.DeleteStatus
    target_commands:
      - Reference.SetStatusAsInitial
      - Reference.DeleteStatus

  - name: StateMachine_Master
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_HideStateMachine
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.HideStateMachine
      - name: OpcenterEXFN_ReferenceData.ACT_UnhideStateMachine
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.UnhideStateMachine
      - name: OpcenterEXFN_ReferenceData.ACT_FreezeStateMachine
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.FreezeStateMachine
      - name: OpcenterEXFN_ReferenceData.ACT_UnfreezeStateMachine
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.UnfreezeStateMachine
      - name: OpcenterEXFN_ReferenceData.ACT_DeleteStateMachine
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.DeleteStateMachine
    target_commands:
      - Reference.HideStateMachine
      - Reference.UnhideStateMachine
      - Reference.FreezeStateMachine
      - Reference.UnfreezeStateMachine
      - Reference.DeleteStateMachine

  - name: StateMachine_Details
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_SetStatusAsInitial
        calls:
          - name: OpcenterEXFN_ReferenceData.SUB_SetStatusAsInitial
      - name: OpcenterEXFN_ReferenceData.ACT_DeleteStatusStateMachine
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.DeleteStatus
      - name: OpcenterEXFN_ReferenceData.ACT_DeleteStatusTransition
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.DeleteStatusTransition
      - name: OpcenterEXFN_ReferenceData.ACT_HideStateMachine
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.HideStateMachine
      - name: OpcenterEXFN_ReferenceData.ACT_UnhideStateMachine
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.UnhideStateMachine
      - name: OpcenterEXFN_ReferenceData.ACT_FreezeStateMachine
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.FreezeStateMachine
      - name: OpcenterEXFN_ReferenceData.ACT_UnfreezeStateMachine
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.UnfreezeStateMachine
    target_commands:
      - Reference.SetStatusAsInitial
      - Reference.DeleteStatus
      - Reference.DeleteStatusTransition
      - Reference.HideStateMachine
      - Reference.UnhideStateMachine
      - Reference.FreezeStateMachine
      - Reference.UnfreezeStateMachine

  - name: PANEL_UpdateStatusStateMachine
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateStatusStateMachine
        calls:
          - name: OpcenterEXFN_ReferenceData.UpdateStatus
    target_commands:
      - Reference.UpdateStatus

  - name: PANEL_UpdateStatusTransition
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateStatusTransition
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.UpdateStatusTransition
    target_commands:
      - Reference.UpdateStatusTransition

  - name: PANEL_UpdateStateMachine
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_UpdateStateMachine
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.UpdateStateMachine
    target_commands:
      - Reference.UpdateStateMachine

  - name: PANEL_CreateStatusStateMachine
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_CreateStatusStateMachine
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.CreateStatusStateMachine
    target_commands:
      - Reference.CreateStatus

  - name: PANEL_CreateStateMachine
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_CreateStateMachine
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.CreateStateMachine
    target_commands:
      - Reference.CreateStateMachine

  - name: PANEL_CreateStatusTransition
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_CreateStatusTransition
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.CreateStatusTransition
    target_commands:
      - Reference.CreateStatusTransition

  - name: PANEL_CreateStatusTransitionDefinition
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_CreateStatusTransitionDefinition
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.CreateStatusTransitionDefinition
    target_commands:
      - Reference.CreateStatusTransitionDefinition

  - name: StatusTransitionDefinition_Master
    module: OpcenterEXFN_ReferenceData
    flows:
      - name: OpcenterEXFN_ReferenceData.ACT_HideStatusTransitionDefinition
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.HideStatusTransitionDefinition
      - name: OpcenterEXFN_ReferenceData.ACT_UnhideStatusTransitionDefinition
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.UnhideStatusTransitionDefinition
      - name: OpcenterEXFN_ReferenceData.ACT_FreezeStatusTransitionDefinition
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.FreezeStatusTransitionDefinition
      - name: OpcenterEXFN_ReferenceData.ACT_UnfreezeStatusTransitionDefinition
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.UnfreezeStatusTransitionDefinition
      - name: OpcenterEXFN_ReferenceData.ACT_DeleteStatusTransitionDefinition
        calls:
          - name: OpcenterEXFN_ReferenceData_Connector.DeleteStatusTransitionDefinition
    target_commands:
      - Reference.HideStatusTransitionDefinition
      - Reference.UnhideStatusTransitionDefinition
      - Reference.FreezeStatusTransitionDefinition
      - Reference.UnfreezeStatusTransitionDefinition
      - Reference.DeleteStatusTransitionDefinition

```

---

## 7. Pages/Panels Commands

Simplified view showing only the target commands for each page/panel.

| Page/Panel | Module | Target Commands |
|------------|--------|------------------|
| PANEL_UpdateUoMDimension | OpcenterEXFN_ReferenceData | Reference.UpdateUoMDimension |
| PANEL_CreateBaseUoMDimension | OpcenterEXFN_ReferenceData | Reference.CreateUoMDimension |
| PANEL_CreateBaseUoM_UoMDimension | OpcenterEXFN_ReferenceData | Reference.CreateBaseUoM |
| PANEL_UoMDimensionUpdateUoM | OpcenterEXFN_ReferenceData | Reference.UpdateUoM |
| UoMDimension_Details | OpcenterEXFN_ReferenceData | Reference.UnhideUoM<br>Reference.HideUoM<br>Reference.DeleteUoM |
| UoMDimension_Master_SingleSelection | OpcenterEXFN_ReferenceData | Reference.UnhideUoMDimension<br>Reference.HideUoMDimension<br>Reference.FreezeUoMDimension<br>Reference.UnfreezeUoMDimension<br>Reference.DeleteUoMDimension |
| UoMDimension_Master_MultiSelection | OpcenterEXFN_ReferenceData | Reference.DeleteUoM |
| UoM_Master | OpcenterEXFN_ReferenceData | Reference.UnhideUoM<br>Reference.HideUoM<br>Reference.DeleteUoM |
| PANEL_CreateBaseUoM | OpcenterEXFN_ReferenceData | Reference.CreateBaseUoM |
| PANEL_UpdateUoM | OpcenterEXFN_ReferenceData | Reference.UpdateUoM |
| PANEL_UpdateUoMSubmultiple | OpcenterEXFN_ReferenceData | Reference.UpdateMultipleOrSubmultipleUoM |
| PANEL_CreateUoMSubmultiple | OpcenterEXFN_ReferenceData | Reference.CreateMultipleOrSubmultipleUoM |
| PANEL_CreateUoMFactor | OpcenterEXFN_ReferenceData | Reference.CreateUoMFactor |
| PANEL_UpdateUoMFactor | OpcenterEXFN_ReferenceData | Reference.UpdateUoMFactor<br>APPNAME.COMMANDNAME<br>EXFN_ServiceLayer.CallCommandAction |
| UoM_Details | OpcenterEXFN_ReferenceData | Reference.UnhideUoM<br>Reference.HideUoM<br>Reference.DeleteUoM<br>Reference.DeleteUoMFactor |
| NumberingPattern_Master | OpcenterEXFN_ReferenceData | Reference.FreezeNumberingPattern<br>Reference.UnfreezeNumberingPattern<br>Reference.DeleteNumberingPattern |
| PANEL_PreviewNumberingPattern | OpcenterEXFN_ReferenceData | - |
| PANEL_UpdateNumberingPattern | OpcenterEXFN_ReferenceData | Reference.UpdateNumberingPattern |
| PANEL_CreateNumberingPattern | OpcenterEXFN_ReferenceData | Reference.CreateNumberingPattern |
| PANEL_UpdateNumberingPatternPart_ValidatePart | OpcenterEXFN_ReferenceData | Reference.UpdateNumberingPatternPart |
| PANEL_UpdateNumberingPatternPart | OpcenterEXFN_ReferenceData | Reference.UpdateNumberingPatternPart |
| PANEL_UpdateNumberingPatternForNumberingPatternPart | OpcenterEXFN_ReferenceData | Reference.UpdateNumberingPattern |
| NumberingPattern_Details | OpcenterEXFN_ReferenceData | Reference.MoveNumberingPatternPart<br>Reference.DeleteNumberingPatternPart<br>Reference.FreezeNumberingPattern<br>Reference.UnfreezeNumberingPattern |
| Counter_Master | OpcenterEXFN_ReferenceData | Reference.HideCounter<br>Reference.UnhideCounter<br>Reference.FreezeCounter<br>Reference.UnfreezeCounter<br>Reference.ResetCounter<br>Reference.DeleteCounter<br>AppName.CommandName |
| PANEL_UpdateCounter | OpcenterEXFN_ReferenceData | Reference.UpdateCounter |
| PANEL_CreateCounter | OpcenterEXFN_ReferenceData | Reference.CreateCounter |
| Counter_Details | OpcenterEXFN_ReferenceData | Reference.HideCounter<br>Reference.UnhideCounter<br>Reference.FreezeCounter<br>Reference.UnfreezeCounter<br>Reference.ResetCounter |
| StatusDefinition_Master | OpcenterEXFN_ReferenceData | Reference.HideStatusDefinition<br>Reference.UnhideStatusDefinition<br>Reference.FreezeStatusDefinition<br>Reference.UnfreezeStatusDefinition<br>Reference.DeleteStatusDefinition |
| PANEL_CreateStatusDefinition | OpcenterEXFN_ReferenceData | Reference.CreateStatusDefinition |
| PANEL_UpdateStatusDefinition | OpcenterEXFN_ReferenceData | Reference.UpdateStatusDefinition |
| PANEL_CreateStatusBehaviorDefinition | OpcenterEXFN_ReferenceData | Reference.CreateStatusBehaviorDefinition |
| PANEL_UpdateStatusBehaviorDefinition | OpcenterEXFN_ReferenceData | Reference.UpdateStatusBehaviorDefinition |
| StatusBehaviorDefinition_Details | OpcenterEXFN_ReferenceData | - |
| StatusBehaviorDefinition_Master | OpcenterEXFN_ReferenceData | Reference.DeleteStatusBehaviorDefinition |
| PANEL_UpdateStatus | OpcenterEXFN_ReferenceData | Reference.UpdateStatus |
| PANEL_CreateStatus | OpcenterEXFN_ReferenceData | Reference.CreateStatus |
| PANEL_AssociateStatusBehaviorDefinitionsWithStatus | OpcenterEXFN_ReferenceData | Reference.AssociateStatusBehaviorDefinitionsWithStatus |
| Status_Details | OpcenterEXFN_ReferenceData | Reference.DisassociateStatusBehaviorDefinitionsFromStatus<br>Reference.SetStatusAsInitial |
| Status_Master | OpcenterEXFN_ReferenceData | Reference.SetStatusAsInitial<br>Reference.DeleteStatus |
| StateMachine_Master | OpcenterEXFN_ReferenceData | Reference.HideStateMachine<br>Reference.UnhideStateMachine<br>Reference.FreezeStateMachine<br>Reference.UnfreezeStateMachine<br>Reference.DeleteStateMachine |
| StateMachine_Details | OpcenterEXFN_ReferenceData | Reference.SetStatusAsInitial<br>Reference.DeleteStatus<br>Reference.DeleteStatusTransition<br>Reference.HideStateMachine<br>Reference.UnhideStateMachine<br>Reference.FreezeStateMachine<br>Reference.UnfreezeStateMachine |
| PANEL_UpdateStatusStateMachine | OpcenterEXFN_ReferenceData | Reference.UpdateStatus |
| PANEL_UpdateStatusTransition | OpcenterEXFN_ReferenceData | Reference.UpdateStatusTransition |
| PANEL_UpdateStateMachine | OpcenterEXFN_ReferenceData | Reference.UpdateStateMachine |
| PANEL_CreateStatusStateMachine | OpcenterEXFN_ReferenceData | Reference.CreateStatus |
| PANEL_CreateStateMachine | OpcenterEXFN_ReferenceData | Reference.CreateStateMachine |
| PANEL_CreateStatusTransition | OpcenterEXFN_ReferenceData | Reference.CreateStatusTransition |
| PANEL_CreateStatusTransitionDefinition | OpcenterEXFN_ReferenceData | Reference.CreateStatusTransitionDefinition |
| StatusTransitionDefinition_Master | OpcenterEXFN_ReferenceData | Reference.HideStatusTransitionDefinition<br>Reference.UnhideStatusTransitionDefinition<br>Reference.FreezeStatusTransitionDefinition<br>Reference.UnfreezeStatusTransitionDefinition<br>Reference.DeleteStatusTransitionDefinition |

---

_Report generated by export_manifest tool_
