using MasterDetailsBase.Widget;
using Microsoft.Playwright;

namespace OperatorTerminals_EXDS.SidePanels
{
    public class OperatorLandingStartPanel
    {
        public OperatorLandingStartPanelLocators Locators { get; protected set; } = new OperatorLandingStartPanelLocators();
        
        public class OperatorLandingStartPanelLocators
        {
            public virtual string panelTitle => "Start Work Order Operation";
            
            public ContentsLocators Contents { get; } = new ContentsLocators();
            public ButtonsLocators Buttons { get; } = new ButtonsLocators();
            public GalleryLocators Gallery { get; } = new GalleryLocators();
        }
        
        public class ContentsLocators
        {
            public virtual string actualTargetQuantityCaption => "Actual Target Quantity:";
            public virtual string mxActualTargetQuantity => "mx-name-label1";
            public virtual string plannedTargetQuantityCaption => "Planned Target Quantity:";
            public virtual string mxPlannedTargetQuantity => "mx-name-label2";
            public virtual string showAvailableSerialNumbersCaption => "Show available Serial Numbers";
            public virtual string mxShowAvailableSerialNumbers => "mx-name-label4";
            public virtual string selectASerialNumberToAssociateCaption => "Select a Serial Number to associate";
            public virtual string mxSelectASerialNumberToAssociate => "mx-name-label3";
            public virtual string associateOrAutoGenerateSerialNumberCaption => "Associate or auto-generate Serial Number";
            public virtual string mxAssociateOrAutoGenerateSerialNumber => "mx-name-label6";
            public virtual string associateTheSerialNumberFoundInSystemCaption => "Associate the Serial Number found in system";
            public virtual string mxAssociateTheSerialNumberFoundInSystem => "mx-name-label8";
            public virtual string textBox2Caption => "textBox2";
            public virtual string mxTextBox2 => "mx-name-textBox2";
            public virtual string orCaption => "OR";
            public virtual string mxOr => "mx-name-label5";
            public virtual string textBox3Caption => "textBox3";
            public virtual string mxTextBox3 => "mx-name-textBox3";
            public virtual string woStartPanelSNGalleryCaption => "WOStartPanelSNGallery";
            public virtual string mxWoStartPanelSNGallery => "mx-name-WOStartPanelSNGallery";
            public virtual string woStartPanelEquipmentGalleryCaption => "WOStartPanelEquipmentGallery";
            public virtual string mxWoStartPanelEquipmentGallery => "mx-name-WOStartPanelEquipmentGallery";
        }
        
        public class ButtonsLocators
        {
            public virtual string associateCaption => "Associate";
            public virtual string mxAssociate => "mx-name-AssociateSNFromDropDown";
            public virtual string generateAndAssociateCaption => "Generate and Associate";
            public virtual string mxGenerateAndAssociate => "mx-name-GenerateAndAssociateSNFromNId";
            public virtual string generateAndAssociateCaption_1 => "Generate and Associate";
            public virtual string mxGenerateAndAssociate_1 => "mx-name-GenerateAndAssociateSNFromQuantity";
            public virtual string associateCaption_1 => "Associate";
            public virtual string mxAssociate_1 => "mx-name-AssociateSNFromInputNId";
            public virtual string generateAndAssociateCaption_2 => "Generate and Associate";
            public virtual string mxGenerateAndAssociate_2 => "mx-name-GenerateAndAssociateSNFromNIdPlaceholder";
            public virtual string startCaption => "Start";
            public virtual string mxStart => "mx-name-startActionButton";
            public virtual string startCaption_1 => "Start";
            public virtual string mxStart_1 => "mx-name-startActionButton1";
            public virtual string startAllCaption => "Start All";
            public virtual string mxStartAll => "mx-name-startActionButton2";
            public virtual string cancelCaption => "Cancel";
            public virtual string mxCancel => "mx-name-cancelActionButton";
        }
        
        public class GalleryLocators
        {
            public virtual string woStartPanelSNGallery => "container8";
            public virtual string woStartPanelEquipmentGallery => "container4";
        }
    }
}
