# Page Manifest — Layout Placeholders & Tabs

| Metric | Count |
|---|---|
| Pages with Main placeholder | 1 |
| Pages with Right placeholder | 1 |
| Total tabs found | 5 |

---

## 📄 StateMachine_Details

### Main

_Placeholder:_ `OpcenterEXFN_DISW_DesignSystem.EXFN_Master.Main`

#### Tab 1 — `tabPage1` _Overview_

| Widget Type | Name | Caption |
|---|---|---|
| TextBox | textBox1 | Id |
| TextBox | textBox2 | Name |
| TextArea | textArea4 | Description |
| RadioButtonGroup | radioButtons5 | Is System Defined |
| RadioButtonGroup | radioButtons4 | Is Hidden |
| RadioButtonGroup | radioButtons1 | Is Frozen |
| DatePicker | datePicker1 | Created On |
| DatePicker | datePicker2 | Last Updated On |

#### Tab 2 — `tabPage2` _Statuses_

| Widget Type | Name | Caption |
|---|---|---|
| DataGrid2 | dataGrid21 | (Data Grid 2) |

  **Columns of `dataGrid21`:**

  | # | Column Name | Caption |
  |---|---|---|
  | 1 | col1 | Id |
  | 2 | col2 | Name |
  | 3 | col3 | Color (Hex Color) |
  | 4 | col4 | Outcome |
  | 5 | col5 | Is System Defined |
  | 6 | col6 | Is Initial |
  | 7 | col7 | Created On |
  | 8 | col8 | Last Updated On |

  **Contextual CommandBar Buttons of `dataGrid21`:**

  | # | Button Name | Caption |
  |---|---|---|
  | 1 | actionButton4 | Create |
  | 2 | actionButton7 | Edit  |
  | 3 | actionButton10 | Set Status As Initial |
  | 4 | actionButton11 | Delete |


#### Tab 3 — `tabPage3` _Transitions_

| Widget Type | Name | Caption |
|---|---|---|
| DataGrid2 | dataGrid22 | (Data Grid 2) |

  **Columns of `dataGrid22`:**

  | # | Column Name | Caption |
  |---|---|---|
  | 1 | col1 | Source Status |
  | 2 | col2 | Target Status |
  | 3 | col3 | Verb |
  | 4 | col4 | Raise Event |
  | 5 | col5 | Is System Defined |
  | 6 | col6 | Created On |
  | 7 | col7 | Last Updated On |

  **Contextual CommandBar Buttons of `dataGrid22`:**

  | # | Button Name | Caption |
  |---|---|---|
  | 1 | actionButton1 | Create |
  | 2 | actionButton2 | Edit |
  | 3 | actionButton3 | Delete |


#### Tab 4 — `tabPage4` _State Machine Graph_

_No widgets with captions found._

#### Tab 5 — `tabPage5` _Audit Trail_

_No widgets with captions found._

### Right

_Placeholder:_ `OpcenterEXFN_DISW_DesignSystem.EXFN_Master.Right`

**Vertical CommandBar Buttons:**

| # | Button Name | Caption | Nanoflow | Show Page |
|---|---|---|---|---|
| 1 | actionButton12 | Edit State Machine | ACT_UpdateStateMachine_ShowPanel | PANEL_UpdateStateMachine |
| 2 | actionButton13 | Hide State Machine | ACT_HideStateMachine | - |
| 3 | actionButton14 | Unhide State Machine | ACT_UnhideStateMachine | - |
| 4 | actionButton8 | Freeze State Machine | ACT_FreezeStateMachine | - |
| 5 | actionButton9 | Unfreeze State Machine | ACT_UnfreezeStateMachine | - |

**Button → Page Navigation:**

| Button Caption | Nanoflow | Page Opened |
|---|---|---|
| Edit State Machine | ACT_UpdateStateMachine_ShowPanel | PANEL_UpdateStateMachine |

**Panel: `PANEL_UpdateStateMachine`**

| Widget Type | Name | Caption |
|---|---|---|
| ActionButton | actionButton1 | Save |
| ActionButton | actionButton2 | Cancel |
| TextBox | textBox7 | Id |
| TextBox | textBox5 | Name |
| TextBox | textBox6 | Description |

---

## Summary

- **Pages with placeholders:** 1
- **Total tabs:** 5
