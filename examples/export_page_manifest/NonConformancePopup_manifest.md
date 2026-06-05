# NonConformancePopup

---

## 📄 NonConformancePopup

**Title:** Declare Non-Conformance

### Main

**Layout:** `OpcenterEXFN_DISW_DesignSystem.CanvasContainer.Main`

**DataView:** `NCPageContextDataView`

#### Wizard Area

##### Area 1: `container1`

| # | Source | Type | Name | Caption |
|---|---|---|---|---|
| 1 | OpcenterEXDS_OperatorLanding.NonConformancesNavigationWizard | Label | label7 | 1 |
| 2 | OpcenterEXDS_OperatorLanding.NonConformancesNavigationWizard | DynamicText | text3 | Add Details |
| 3 | OpcenterEXDS_OperatorLanding.NonConformancesNavigationWizard | Label | label5 | 2 |
| 4 | OpcenterEXDS_OperatorLanding.NonConformancesNavigationWizard | DynamicText | text2 | Add Defects |
| 5 | OpcenterEXDS_OperatorLanding.NonConformancesNavigationWizard | Label | label1 | 3 |
| 6 | OpcenterEXDS_OperatorLanding.NonConformancesNavigationWizard | DynamicText | text4 | Add Documents |
| 7 | OpcenterEXDS_OperatorLanding.NonConformancesNavigationWizard | Label | label3 | 4 |
| 8 | OpcenterEXDS_OperatorLanding.NonConformancesNavigationWizard | DynamicText | text5 | Submit |
| 9 | OpcenterEXDS_OperatorLanding.NonConformancesNavigationWizard | Container | wizardStep2 | OnClick: OpcenterEXDS_OperatorLanding.ACT_NonConformancesNavigationWizard_Details_Click |
| 10 | OpcenterEXDS_OperatorLanding.NonConformancesNavigationWizard | Container | wizardStep1 | OnClick: OpcenterEXDS_OperatorLanding.ACT_NonConformancesNavigationWizard_Failure_Click |
| 11 | OpcenterEXDS_OperatorLanding.NonConformancesNavigationWizard | Container | wizardStep3 | OnClick: OpcenterEXDS_OperatorLanding.ACT_NonConformancesNavigationWizard_Documents_Click |
| 12 | OpcenterEXDS_OperatorLanding.NonConformancesNavigationWizard | Container | wizardStep4 | OnClick: OpcenterEXDS_OperatorLanding.ACT_NonConformancesNavigationWizard_Submit_Click |

**Steps**

##### Section 1: `containerFailures`

| # | Source | Type | Name | Caption |
|---|---|---|---|---|
| 1 | containerFailures | DynamicText | NCFailureStep1_text | 2. Select one or more Failures from the list below |
| 2 | OpcenterEXDS_OperatorLanding.FailureBrowser | Gallery | NCFailuresBrowser_BreadCrumbGallery | (no caption) |
| 3 | OpcenterEXDS_OperatorLanding.FailureBrowser | DynamicText | text1 | {1} |
| 4 | OpcenterEXDS_OperatorLanding.FailureBrowser | Gallery | TileGallery | (no caption) |

##### Section 2: `containerDetails`

| # | Source | Type | Name | Caption |
|---|---|---|---|---|
| 1 | containerDetails | DynamicText | NCFailureStep2_text | 1. Fill data in the fields below |
| 2 | OpcenterEXDS_OperatorLanding.NonConformancesInfo | ActionButton | buttonWorkOrderOperationContext1 | Work Order Operation |
| 3 | OpcenterEXDS_OperatorLanding.NonConformancesInfo | ActionButton | buttonMaterialTrackingUnitContext1 | Material Tracking Unit |
| 4 | OpcenterEXDS_OperatorLanding.NonConformancesInfo | ActionButton | buttonToolContext | Equipment |
| 5 | OpcenterEXDS_OperatorLanding.NonConformancesInfo | ActionButton | buttonToolContext1 | Tool |
| 6 | OpcenterEXDS_OperatorLanding.NonConformancesInfo | DynamicText | text11 | Identifier |
| 7 | OpcenterEXDS_OperatorLanding.NonConformancesInfo | DynamicText | text13 | Severity |
| 8 | OpcenterEXDS_OperatorLanding.NonConformancesInfo | TextBox | NCInfo_NCID | (no caption) |
| 9 | OpcenterEXDS_OperatorLanding.NonConformancesInfo | DynamicText | ToBeConsMatNP_ErrorMessage | Identifier is  mandatory |
| 10 | OpcenterEXDS_OperatorLanding.NonConformancesInfo | ReferenceSelector | referenceSelector1 | (no caption) |
| 11 | OpcenterEXDS_OperatorLanding.NonConformancesInfo | DynamicText | text4 | Work Order |
| 12 | OpcenterEXDS_OperatorLanding.NonConformancesInfo | DynamicText | text5 | Work Order Operation |
| 13 | OpcenterEXDS_OperatorLanding.NonConformancesInfo | TextBox | NCInfo_WO | (no caption) |
| 14 | OpcenterEXDS_OperatorLanding.NonConformancesInfo | TextBox | NCInfo_WOO | (no caption) |
| 15 | OpcenterEXDS_OperatorLanding.NonConformancesInfo | DynamicText | text10 | Quantity |
| 16 | OpcenterEXDS_OperatorLanding.NonConformancesInfo | TextBox | NCInfo_Qty | (no caption) |
| 17 | OpcenterEXDS_OperatorLanding.NonConformancesInfo | DynamicText | ToBeConsMatNP_QuantityErrorMsg | Quantity must be greater than 0 and lower than {1} |
| 18 | OpcenterEXDS_OperatorLanding.NonConformancesInfo | DynamicText | text6 | Comments |
| 19 | OpcenterEXDS_OperatorLanding.NonConformancesInfo | TextArea | NCInfo_Comment | (no caption) |

##### Section 3: `containerDocuments`

| # | Source | Type | Name | Caption |
|---|---|---|---|---|
| 1 | containerDocuments | DynamicText | NCFailureStep3_text1 | 3. Select one or more Documents from the list below |
| 2 | OpcenterEXDS_OperatorLanding.NonConformanceDocuments | Gallery | NCDocuments_TileGallery | (no caption) |

##### Section 4: `containerSubmit`

| # | Source | Type | Name | Caption |
|---|---|---|---|---|
| 1 | containerSubmit | DynamicText | NCFailureStep3_text | 4. Review the information and Submit |
| 2 | OpcenterEXDS_OperatorLanding.NonConformancesSummary | DynamicText | text1 | Failures |
| 3 | OpcenterEXDS_OperatorLanding.NonConformancesSummary | Label | NCSummary_labelfailuretype | Failure Type |
| 4 | OpcenterEXDS_OperatorLanding.NonConformancesSummary | Gallery | NCSummary_failuresgallery | (no caption) |
| 5 | OpcenterEXDS_OperatorLanding.NonConformancesSummary | DynamicText | text2 | Details |
| 6 | OpcenterEXDS_OperatorLanding.NonConformancesSummary | Label | NCSummary_identifier | Identifier |
| 7 | OpcenterEXDS_OperatorLanding.NonConformancesSummary | DynamicText | NCSummary_ncid | {1} |
| 8 | OpcenterEXDS_OperatorLanding.NonConformancesSummary | Label | NCSummary_labelcontext | Context |
| 9 | OpcenterEXDS_OperatorLanding.NonConformancesSummary | DynamicText | NCSummary_context | {1} |
| 10 | OpcenterEXDS_OperatorLanding.NonConformancesSummary | Label | NCSummary_labelseverity | Severity |
| 11 | OpcenterEXDS_OperatorLanding.NonConformancesSummary | DynamicText | NCSummary_severity | {1} |
| 12 | OpcenterEXDS_OperatorLanding.NonConformancesSummary | Label | NCSummary_labelwo | Work Order |
| 13 | OpcenterEXDS_OperatorLanding.NonConformancesSummary | DynamicText | NCSummary_wo | {1} |
| 14 | OpcenterEXDS_OperatorLanding.NonConformancesSummary | Label | NCSummary_labelwoo | Work Order Operation |
| 15 | OpcenterEXDS_OperatorLanding.NonConformancesSummary | DynamicText | NCSummary_woo | {1} |
| 16 | OpcenterEXDS_OperatorLanding.NonConformancesSummary | Label | NCSummary_labelwoo1 | Quantity |
| 17 | OpcenterEXDS_OperatorLanding.NonConformancesSummary | DynamicText | NCSummary_woo1 | {1} |
| 18 | OpcenterEXDS_OperatorLanding.NonConformancesSummary | Label | NCSummary_labelcomments | Comments |
| 19 | OpcenterEXDS_OperatorLanding.NonConformancesSummary | DynamicText | NCSummary_comments | {1} |
| 20 | OpcenterEXDS_OperatorLanding.NonConformancesSummary | DynamicText | text4 | Documents |
| 21 | OpcenterEXDS_OperatorLanding.NonConformancesSummary | Label | NCSummary_labelDocument | Identifier |
| 22 | OpcenterEXDS_OperatorLanding.NonConformancesSummary | Gallery | NCSummary_documentsgallery | (no caption) |
| 23 | OpcenterEXDS_OperatorLanding.NonConformancesSummary | DynamicText | text3 | Material Tracking Units |
| 24 | OpcenterEXDS_OperatorLanding.NonConformancesSummary | Label | NCSummary_labelMTU | Identifier |
| 25 | OpcenterEXDS_OperatorLanding.NonConformancesSummary | Gallery | NCSummary_mtusgallery1 | (no caption) |
| 26 | OpcenterEXDS_OperatorLanding.NonConformancesSummary | Gallery | NCSummary_mtusgallery | (no caption) |
| 27 | OpcenterEXDS_OperatorLanding.NonConformancesSummary | DynamicText | text5 | Tools |
| 28 | OpcenterEXDS_OperatorLanding.NonConformancesSummary | Label | NCSummary_labelMTU1 | Identifier |
| 29 | OpcenterEXDS_OperatorLanding.NonConformancesSummary | Gallery | NCSummary_toolsgallery2 | (no caption) |
| 30 | OpcenterEXDS_OperatorLanding.NonConformancesSummary | DynamicText | text6 | Equipment |
| 31 | OpcenterEXDS_OperatorLanding.NonConformancesSummary | Label | NCSummary_labelMTU2 | Identifier |
| 32 | OpcenterEXDS_OperatorLanding.NonConformancesSummary | Gallery | NCSummary_toolsgallery3 | (no caption) |
| 33 | containerSubmit | ActionButton | NC_SubmitButton | Submit |
| 34 | containerSubmit | ActionButton | cancelToolSelectionActionButton | Cancel |

---

## Summary

- **Pages with placeholders:** 1
- **Total tabs:** 0
