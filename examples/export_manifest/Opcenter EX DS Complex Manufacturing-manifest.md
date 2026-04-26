# Manifest Report: Opcenter EX DS Complex Manufacturing

**Mendix Version:** 11.6.4  
**MPR File:** C:\Workspaces\Mendix\2601S186_SPX\CMX\Opcenter EX DS Complex Manufacturing.mpr  
**Generated:** 2026-04-26 15:18:25  

---

## Summary

- **Microflow/Action Calls:** 82
- **Signal Manager Widgets:** 64 subscription(s)
- **Navigation Items:** 1

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

## 3. Signal Manager Widgets

Signal subscriptions from Signal Manager widgets (siemens.mxtosignal.MxToSignal).

Found 64 subscription(s):

| Module | Document Type | Document | Signal Name | App Name | Subscription Filter |
|--------|---------------|----------|-------------|----------|---------------------|
| OpcenterEXDS_PartProgram | Page | PartProgramDetailsPopUp | DNCStartTransferEvent | AppU4DM | WorkOrderOperationNId eq ''' + $ProgramPartDetailsContext/WorkOrderOperationNId + ''' and DNCMachine eq ''' + $ProgramPartDetailsContext/DNCMachine +  |
| OpcenterEXDS_PartProgram | Page | PartProgramDetailsPopUp | DNCCompleteTransferEvent | AppU4DM | WorkOrderOperationNId eq ''' + $ProgramPartDetailsContext/WorkOrderOperationNId + ''' and DNCMachine eq ''' + $ProgramPartDetailsContext/DNCMachine +  |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnChangeFlexibleWOOpStatusToComplete | AppU4DM | - |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnFAICandidateDeclared | AppU4DM | - |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnStartWorkOrderOperationFullQty | AppU4DM | - |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnCreateNonConformanceV3_1 | AppU4DM | - |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnCompleteWorkOrderOperationSerialized | AppU4DM | - |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | StatusChanged | AppU4DM | EnvelopeCategory eq ''ChangePackageStatusChanged |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | WorkOrderSerialNumbersScrapped | AppU4DM | - |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnCompleteReworkOrder | AppU4DM | - |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnFAICompleted | AppU4DM | - |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnPauseWorkOrderOperation | AppU4DM | - |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnSetHoldWorkOrder | AppU4DM | - |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | SnagAndNoteNotificationSgn | AppU4DM | - |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnSkipWOOperation | AppU4DM | - |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnStartWorkOrderOperationSerialized | AppU4DM | - |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnCompleteWorkOrderOperationFullQty | AppU4DM | - |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnSentenceNonConformanceV3_1 | AppU4DM | - |
| OpcenterEXDS_OperatorLanding | Page | OperatorLanding | OnReopenWorkOrderOperation | AppU4DM | - |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnPauseWorkOrderOperation | AppU4DM | if(trim($currentObject/WorkOrderNId) != '') then
'EnvelopeUserField1 eq ''' + $currentObject/WorkOrderNId + ''''
else  |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnCompleteWorkOrderStepSerialized | AppU4DM | if(trim($currentObject/WorkOrderNId) != '') then
'EnvelopeUserField1 eq ''' + $currentObject/WorkOrderNId + ''''
else  |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnCompleteWorkOrderStepFullQty | AppU4DM | if(trim($currentObject/WorkOrderNId) != '') then
'EnvelopeUserField1 eq ''' + $currentObject/WorkOrderNId + ''''
else  |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnSentenceNonConformanceV3_1 | AppU4DM | if(trim($currentObject/WorkOrderNId) != '') then
'EnvelopeUserField4 eq ''' + $currentObject/WorkOrderNId + ''''
else  |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | WorkOrderSerialNumbersScrapped | AppU4DM | if(trim($currentObject/WorkOrderNId) != '') then
'WorkOrderNId eq ''' + $currentObject/WorkOrderNId + ''''
else  |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnFAICompleted | AppU4DM | if(trim($currentObject/WorkOrderNId) != '') then
'EnvelopeUserField1 eq ''' + $currentObject/WorkOrderNId + ''''
else  |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnFAICandidateDeclared | AppU4DM | if(trim($currentObject/WorkOrderNId) != '') then
'EnvelopeUserField1 eq ''' + $currentObject/WorkOrderNId + ''''
else  |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | WorkInstructionStatusChangedSignal | WorkInstruction | if(trim($currentObject/WorkOrderNId) != '') then
'Context/CtxEntityValue eq ''' + $currentObject/WorkOrderNId + ''''
else  |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | StatusChanged | AppU4DM | EnvelopeCategory eq ''ChangePackageStatusChanged |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnCompleteReworkOrder | AppU4DM | if(trim($currentObject/WorkOrderNId) != '') then
'WorkOrderNId eq ''' + $currentObject/WorkOrderNId + ''''
else  |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnSkipWOOperation | AppU4DM | if(trim($currentObject/WorkOrderNId) != '') then
'WorkOrderNId eq ''' + $currentObject/WorkOrderNId + ''''
else  |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnStartWorkOrderOperationFullQty | AppU4DM | if(trim($currentObject/WorkOrderNId) != '') then
'EnvelopeUserField1 eq ''' + $currentObject/WorkOrderNId + ''''
else  |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnSetHoldWorkOrder | AppU4DM | if(trim($currentObject/WorkOrderNId) != '') then
'EnvelopeUserField1 eq ''' + $currentObject/WorkOrderNId + ''''
else  |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnDisassemblyWorkOrderOperation | AppU4DM | if(trim($currentObject/WorkOrderNId) != '') then
'EnvelopeUserField1 eq ''' + $currentObject/WorkOrderNId + ''''
else  |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnReopenWorkOrderOperation | AppU4DM | if(trim($currentObject/WorkOrderNId) != '') then
'EnvelopeUserField4 eq ''' + $currentObject/WorkOrderNId + ''''
else  |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnFAIRemoved | AppU4DM | if(trim($currentObject/WorkOrderNId) != '') then
'EnvelopeUserField1 eq ''' + $currentObject/WorkOrderNId + ''''
else  |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnStartWorkOrderOperationSerialized | AppU4DM | if(trim($currentObject/WorkOrderNId) != '') then
'EnvelopeUserField1 eq ''' + $currentObject/WorkOrderNId + ''''
else  |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnCompleteWorkOrderOperationSerialized | AppU4DM | if(trim($currentObject/WorkOrderNId) != '') then
'EnvelopeUserField1 eq ''' + $currentObject/WorkOrderNId + ''''
else  |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnStartWorkOrderStepSerialized | AppU4DM | if(trim($currentObject/WorkOrderNId) != '') then
'EnvelopeUserField1 eq ''' + $currentObject/WorkOrderNId + ''''
else  |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnStartWorkOrderStepFullQty | AppU4DM | if(trim($currentObject/WorkOrderNId) != '') then
'EnvelopeUserField1 eq ''' + $currentObject/WorkOrderNId + ''''
else  |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | SnagAndNoteNotificationSgn | AppU4DM | if(trim($currentObject/WorkOrderNId) != '') then
'WorkOrderNId eq ''' + $currentObject/WorkOrderNId + ''''
else  |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnAssemblyWorkOrderOperation | AppU4DM | if(trim($currentObject/WorkOrderNId) != '') then
'EnvelopeUserField1 eq ''' + $currentObject/WorkOrderNId + ''''
else  |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnUsedToolSignal | AppU4DM | if(trim($currentObject/WorkOrderNId) != '') then
'EnvelopeUserField1 eq ''' + $currentObject/WorkOrderNId + ''''
else  |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnCreateNonConformanceV3_1 | AppU4DM | if(trim($currentObject/WorkOrderNId) != '') then
'WorkOrderNId eq ''' + $currentObject/WorkOrderNId + ''''
else  |
| OpcenterEXDS_OperatorLanding | Page | OperatorTerminal | OnCompleteWorkOrderOperationFullQty | AppU4DM | if(trim($currentObject/WorkOrderNId) != '') then
'EnvelopeUserField1 eq ''' + $currentObject/WorkOrderNId + ''''
else  |
| OpcenterEXDS_EXFN_WorkInstruction | Snippet | SNP_WorkInstruction_VerticalView | WorkInstructionSectionCompletedSignal | WorkInstruction | if($WorkInstructionContextHelper != empty and trim($WorkInstructionContextHelper/WorkOrderNId) != '') then
'Context/CtxEntityValue eq ''' + $WorkInstructionContextHelper/WorkOrderNId + ''''
else 'EnvelopeCategory eq ''WorkInstructionSectionCompleted |
| OpcenterEXDS_EXFN_WorkInstruction | Snippet | SNP_WorkInstruction_VerticalView | WorkInstructionStepCompletedSignal | WorkInstruction | if($WorkInstructionContextHelper != empty and trim($WorkInstructionContextHelper/WorkOrderNId) != '') then
'Context/CtxEntityValue eq ''' + $WorkInstructionContextHelper/WorkOrderNId + ''''
else 'EnvelopeCategory eq ''WorkInstructionStepCompleted |
| OpcenterEXDS_OperatorLanding | Snippet | WorkInstructions | OnLinkWIOnDemandToSerialNumber | AppU4DM | if($WorkOrderOperationOrStepContext/WorkOrderStepId = empty) then 
	'EnvelopeUserField1 eq ''' + $WorkOrderOperationOrStepContext/WorkOrderOperationId + ''''
else
	'EnvelopeUserField1 eq ''' + $WorkOrderOperationOrStepContext/WorkOrderStepId +  |
| OpcenterEXDS_OperatorLanding | Snippet | ToBeUsedTools_Backup | OnUsedToolSignal | AppU4DM | if(trim($WorkOrderOperationOrStepContext/WorkOrderNId) != '') then
'EnvelopeUserField1 eq ''' + $WorkOrderOperationOrStepContext/WorkOrderNId + ''''
else  |
| OpcenterEXDS_OperatorLanding | Snippet | ToBeCoByProducedMaterials | CoProductProduced | AppU4DM | if(trim($WorkOrderOperationOrStepContext/WorkOrderNId) != '') then
'CoProductProducedParameter/WorkOrderNId eq ''' + $WorkOrderOperationOrStepContext/WorkOrderNId + ''''
else  |
| OpcenterEXDS_OperatorLanding | Snippet | ToBeCoByProducedMaterials | ByProductProduced | AppU4DM | if(trim($WorkOrderOperationOrStepContext/WorkOrderNId) != '') then
'ByProductProducedParameter/WorkOrderNId eq ''' + $WorkOrderOperationOrStepContext/WorkOrderNId + ''''
else  |
| OpcenterEXDS_OperatorLanding | Snippet | ToBeCoByProducedMaterials | OutputMaterialProduced | AppU4DM | if(trim($WorkOrderOperationOrStepContext/WorkOrderNId) != '') then
'ProducedMaterial/WorkOrderNId eq ''' + $WorkOrderOperationOrStepContext/WorkOrderNId + ''''
else  |
| OpcenterEXDS_OperatorLanding | Snippet | PartProgram | DNCCompleteTransferEvent | AppU4DM | WorkOrderOperationNId eq ''' + $dataView3/WorkOrderOperationNId + ''' and DNCMachine eq ''' + $dataView3/DNCMachine +  |
| OpcenterEXDS_OperatorLanding | Snippet | PartProgram | DNCStartTransferEvent | AppU4DM | WorkOrderOperationNId eq ''' + $dataView3/WorkOrderOperationNId + ''' and DNCMachine eq ''' + $dataView3/DNCMachine +  |
| OpcenterEXDS_OperatorLanding | Snippet | ToBeProducedMaterial | OnAssemblyWorkOrderOperation | AppU4DM | if(trim($WorkOrderOperationOrStepContext/WorkOrderNId) != '') then
'EnvelopeUserField1 eq ''' + $WorkOrderOperationOrStepContext/WorkOrderNId + ''''
else  |
| OpcenterEXDS_OperatorLanding | Snippet | ToBeProducedMaterial | OnDisassemblyWorkOrderOperation | AppU4DM | if(trim($WorkOrderOperationOrStepContext/WorkOrderNId) != '') then
'EnvelopeUserField1 eq ''' + $WorkOrderOperationOrStepContext/WorkOrderNId + ''''
else  |
| OpcenterEXDS_OperatorLanding | Snippet | ToBeProducedMaterial | ScrapMaterials | AppU4DM | if(trim($WorkOrderOperationOrStepContext/WorkOrderNId) != '') then
'WorkOrderNId eq ''' + $WorkOrderOperationOrStepContext/WorkOrderNId + ''''
else  |
| OpcenterEXDS_OperatorLanding | Snippet | ToBeUsedTools | OnUsedToolSignal | AppU4DM | if(trim($WorkOrderOperationOrStepContext/WorkOrderNId) != '') then
'EnvelopeUserField1 eq ''' + $WorkOrderOperationOrStepContext/WorkOrderNId + ''''
else  |
| OpcenterEXDS_OperatorLanding | Snippet | ToBeUsedTools | MaterialTrackingUnitDeactivationEvent | AppU4DM | if(trim($WorkOrderOperationOrStepContext/WorkOrderNId) != '') then
'EnvelopeUserField1 eq ''' + $WorkOrderOperationOrStepContext/WorkOrderNId + ''''
else  |
| OpcenterEXDS_OperatorLanding | Snippet | ToBeUsedTools | MaterialTrackingUnitActivationEvent | AppU4DM | if(trim($WorkOrderOperationOrStepContext/WorkOrderNId) != '') then
'EnvelopeUserField1 eq ''' + $WorkOrderOperationOrStepContext/WorkOrderNId + ''''
else  |
| OpcenterEXDS_OperatorLanding | Snippet | ToBeUsedTools | ScrewingExecuted | AppU4DM | if(trim($WorkOrderOperationOrStepContext/WorkOrderNId) != '') then
'EnvelopeUserField1 eq ''' + $WorkOrderOperationOrStepContext/WorkOrderNId + ''''
else  |
| EXFN_Quality | Snippet | SNP_QualityInspectionContainer | DelayedExecution | WorkInstruction | RuntimeChrReprContainerNId eq ''' + $ParameterView/ContainerNId +  |
| EXFN_Quality | Snippet | SNP_QualityInspectionContainer | InspectionExecutionChrReprRuntimeNumberChanged | WorkInstruction | RuntimeChrRepresentationContainerNId eq ''' + $ParameterView/ContainerNId +  |
| EXFN_Quality | Snippet | SNP_QualityInspectionContainer | OnCompleteInspectionSampleScenarioInstance | WorkInstruction | RuntimeChrRepresentationContainerNId eq ''' + $ParameterView/ContainerNId +  |
| EXFN_Quality | Snippet | SNP_QualityInspectionContainer | InspectionSampleConfirmed | WorkInstruction | RuntimeChrRepresentationContainerNId eq ''' + $ParameterView/ContainerNId +  |

---

## 4. Navigation Items

| Parent Node | Node | User Roles |
|-------------|------|------------|
| - | Home | - |

---

_Report generated by export_manifest tool_
