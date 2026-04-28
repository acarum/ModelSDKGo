# PANEL_OperatorLandingStart

---

## 📄 PANEL_OperatorLandingStart

**Title:** Start Work Order Operation

### Main

**Layout:** `OpcenterEXFN_DISW_DesignSystem.EXFN_ModalPanel.Main`

#### Contents

| # | Type | Name | Caption |
|---|---|---|---|
| 1 | Label | label1 | Actual Target Quantity: |
| 2 | Label | label2 | Planned Target Quantity: |
| 3 | Label | label4 | Show available Serial Numbers |
| 4 | Label | label3 | Select a Serial Number to associate |
| 5 | Label | label6 | Associate or auto-generate Serial Number |
| 6 | Label | label8 | Associate the Serial Number found in system |
| 7 | TextBox | textBox2 |  |
| 8 | Label | label5 | OR |
| 9 | TextBox | textBox3 |  |
| 10 | Gallery | WOStartPanelSNGallery |  |
| 11 | Gallery | WOStartPanelEquipmentGallery |  |

##### Gallery: WOStartPanelSNGallery

- **Tile container:** container8
- **Sort by:** drop_downSort2
- **Search:** textFilter2

##### Gallery: WOStartPanelEquipmentGallery

- **Tile container:** container4
- **Sort by:** drop_downSort1
- **Search:** textFilter1

**Buttons:**

| # | Name | Caption |
|---|---|---|
| 1 | AssociateSNFromDropDown | Associate |
| 2 | GenerateAndAssociateSNFromNId | Generate and Associate |
| 3 | GenerateAndAssociateSNFromQuantity | Generate and Associate |
| 4 | AssociateSNFromInputNId | Associate |
| 5 | GenerateAndAssociateSNFromNIdPlaceholder | Generate and Associate |
| 6 | startActionButton | Start |
| 7 | startActionButton1 | Start |
| 8 | startActionButton2 |   Start All   |
| 9 | cancelActionButton | Cancel |

---

## Summary

- **Total Buttons:** 9
