# Manifest Report: OC EX System

**Mendix Version:** 11.9.0  
**MPR File:** C:\Workspaces\Mendix\MDUI\System_Mendix_CLI\OC EX System.mpr  
**Generated:** 2026-04-26 15:14:22  

---

## Summary

- **Microflow/Action Calls:** 78
- **Signal Manager Widgets:** 3 subscription(s)
- **Navigation Items:** 10

---

## 2. Microflow/Action Calls

Microflow and action calls found in the project (MicroflowCall, JavaAction, ExternalAction).

Found 78 call(s):

| Microflow | Module | Call Type | AppName | CommandName |
|-----------|--------|-----------|---------|-------------|
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
| Microflow | Documents | ExternalAction | TripPinServiceRW | ResetDataSource |

---

## 3. Signal Manager Widgets

Signal subscriptions from Signal Manager widgets (siemens.mxtosignal.MxToSignal).

Found 3 subscription(s):

| Module | Document Type | Document | Signal Name | App Name | Subscription Filter |
|--------|---------------|----------|-------------|----------|---------------------|
| OpcenterEXFN_ReferenceData | Page | StateMachine_Details | SN | APPName | filter0 |
| OpcenterEXFN_ReferenceData | Page | StateMachine_Details | SN2 | AN2 | filter |
| OpcenterEXFN_ReferenceData | Snippet | MySnippet | SNIPPET_SN | SNIPPET_APPName | SNIPPET_filter0 |

---

## 4. Navigation Items

| Parent Node | Node | User Roles |
|-------------|------|------------|
| - | System | - |
| System | Counters | Administrator, User |
| System | Numbering Patterns | Administrator, User |
| System | State Machines | Administrator, User |
| System | Statuses | Administrator, User |
| System | Status Behavior Definitions | Administrator, User |
| System | Status Definitions | Administrator, User |
| System | Status Transition Definitions | Administrator, User |
| System | Unit of Measures | Administrator, User |
| System | UoM Dimensions | Administrator, User |

---

_Report generated by export_manifest tool_
