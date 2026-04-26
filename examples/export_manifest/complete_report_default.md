# Manifest Report: OC EX System

**Mendix Version:** 11.9.0  
**MPR File:** C:\Workspaces\Mendix\MDUI\System_Mendix_CLI\OC EX System.mpr  
**Generated:** 2026-04-26 14:03:37  

---

## Summary

- **Microflow/Action Calls:** 78
- **Signal Manager Widgets:** 3 subscription(s)
- **Navigation Items:** 10
- **System Roles:** 2
- **Pages/Snippets:** 79

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

## 5. System Roles & Page Accessibility

### 5.1 System Roles

Found 2 system role(s):

| Role Name | Module |
|-----------|--------|
| Administrator | System |
| User | System |

### 5.2 Page Accessibility by Role

| System Role | Accessible Pages Count | Pages |
|-------------|------------------------|-------|
| User | 51 | OpcenterEXFN_DISW_DesignSystem.RuntimeUIApplicationsByPlant (OpcenterEXFN_DISW_DesignSystem), Opcent... |
| Public (No Restrictions) | 20 | Administration.MyAccount (Administration), Administration.ChangeMyPasswordForm (Administration), Adm... |
| Administrator | 59 | Administration.Account_Edit (Administration), Administration.Account_New (Administration), Administr... |

### 5.3 Page Accessibility by Page

Found 79 page(s)/snippet(s):

| Page Name | Module | Type | Allowed Roles | Access |
|-----------|--------|------|---------------|--------|
| Administration.MyAccount | Administration | Page | - | **Public** |
| Administration.ChangeMyPasswordForm | Administration | Page | - | **Public** |
| Administration.Account_Edit | Administration | Page | Administrator | Restricted |
| Administration.Account_New | Administration | Page | Administrator | Restricted |
| Administration.ChangePasswordForm | Administration | Page | - | **Public** |
| Administration.Account_Overview | Administration | Page | Administrator | Restricted |
| Administration.RuntimeInstances | Administration | Page | Administrator | Restricted |
| Administration.ScheduledEvents | Administration | Page | Administrator | Restricted |
| Administration.ActiveSessions | Administration | Page | Administrator | Restricted |
| Administration.ReadMe | Administration | Snippet | - | **Public** |
| OpcenterEXFN_DISW_DesignSystem.ClientConfiguration_NewEdit | OpcenterEXFN_DISW_DesignSystem | Page | Administrator | Restricted |
| OpcenterEXFN_DISW_DesignSystem.Snippet_Configuration | OpcenterEXFN_DISW_DesignSystem | Snippet | - | **Public** |
| OpcenterEXFN_DISW_DesignSystem.ClientConfiguration_Overview | OpcenterEXFN_DISW_DesignSystem | Page | Administrator | Restricted |
| OpcenterEXFN_DISW_DesignSystem.RuntimeUIApplicationsByPlant | OpcenterEXFN_DISW_DesignSystem | Page | Administrator, User | Restricted |
| OpcenterEXFN_DISW_DesignSystem.UserAvatar | OpcenterEXFN_DISW_DesignSystem | Snippet | - | **Public** |
| OpcenterEXFN_DISW_DesignSystem.OC_EX_HomePage | OpcenterEXFN_DISW_DesignSystem | Page | Administrator, User | Restricted |
| OpcenterEXFN_DISW_DesignSystem.DataArchivingManager | OpcenterEXFN_DISW_DesignSystem | Snippet | - | **Public** |
| OpcenterEXFN_DISW_DesignSystem.SessionManager | OpcenterEXFN_DISW_DesignSystem | Snippet | - | **Public** |
| EXFN_AuditTrailViewer.AuditTrailRecord_View | EXFN_AuditTrailViewer | Page | Administrator, User | Restricted |
| EXFN_AuditTrailViewer.AuditTrail | EXFN_AuditTrailViewer | Page | Administrator, User | Restricted |
| EXFN_AuditTrailViewer.AuditTrailViewer | EXFN_AuditTrailViewer | Snippet | - | **Public** |
| OpcenterEXFN_ReferenceData.PANEL_UpdateUoMDimension | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.PANEL_CreateBaseUoMDimension | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.PANEL_CreateBaseUoM_UoMDimension | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.PANEL_UoMDimensionUpdateUoM | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.UoMDimension_Details | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.UoMDimension_Master_SingleSelection | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.UoMDimension_Master_MultiSelection | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.UoM_Master | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.PANEL_CreateBaseUoM | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.PANEL_UpdateUoM | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.PANEL_UpdateUoMSubmultiple | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.PANEL_CreateUoMSubmultiple | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.PANEL_CreateUoMFactor | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.PANEL_UpdateUoMFactor | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.UoM_Details | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.NumberingPattern_Master | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.PANEL_PreviewNumberingPattern | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.PANEL_UpdateNumberingPattern | OpcenterEXFN_ReferenceData | Page | - | **Public** |
| OpcenterEXFN_ReferenceData.PANEL_CreateNumberingPattern | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.PANEL_UpdateNumberingPatternPart_ValidatePart | OpcenterEXFN_ReferenceData | Page | - | **Public** |
| OpcenterEXFN_ReferenceData.PANEL_UpdateNumberingPatternPart | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.PANEL_UpdateNumberingPatternForNumberingPatternPart | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.NumberingPattern_Details | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.Counter_Master | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.PANEL_UpdateCounter | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.PANEL_CreateCounter | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.Counter_Details | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.StatusDefinition_Master | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.PANEL_CreateStatusDefinition | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.PANEL_UpdateStatusDefinition | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.PANEL_CreateStatusBehaviorDefinition | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.PANEL_UpdateStatusBehaviorDefinition | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.StatusBehaviorDefinition_Details | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.StatusBehaviorDefinition_Master | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.PANEL_UpdateStatus | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.PANEL_CreateStatus | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.PANEL_AssociateStatusBehaviorDefinitionsWithStatus | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.Status_Details | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.Status_Master | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.StateMachine_Master | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.StateMachine_Details | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.PANEL_UpdateStatusStateMachine | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.PANEL_UpdateStatusTransition | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.PANEL_UpdateStateMachine | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.PANEL_CreateStatusStateMachine | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.PANEL_CreateStateMachine | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.PANEL_CreateStatusTransition | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.PANEL_CreateStatusTransitionDefinition | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| OpcenterEXFN_ReferenceData.StatusTransitionDefinition_Master | OpcenterEXFN_ReferenceData | Page | Administrator, User | Restricted |
| DISW_DesignSystem.SN_NavbarProfile | DISW_DesignSystem | Snippet | - | **Public** |
| DISW_DesignSystem.BottomBarInverse_Template | DISW_DesignSystem | Snippet | - | **Public** |
| DISW_DesignSystem.SN_Native_NavbarProfile | DISW_DesignSystem | Snippet | - | **Public** |
| DISW_DesignSystem.FeedbackWidget_Template | DISW_DesignSystem | Snippet | - | **Public** |
| DISW_DesignSystem.MobileFeatures | DISW_DesignSystem | Snippet | - | **Public** |
| Atlas_Core.FeedbackWidget | Atlas_Core | Snippet | - | **Public** |
| Atlas_Core.LanguageSelectorWidget | Atlas_Core | Snippet | - | **Public** |
| MyModule.SNIPPET_VerticalCommandBar_Button | MyModule | Snippet | - | **Public** |
| OpcenterEXFN_ReferenceData.MySnippet | OpcenterEXFN_ReferenceData | Snippet | - | **Public** |

---

_Report generated by export_manifest tool_
