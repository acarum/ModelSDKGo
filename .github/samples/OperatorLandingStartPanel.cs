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
            public virtual string actualTargetQuantity => "Actual Target Quantity:";
            public virtual string mxActualTargetQuantity => "mx-name-label1";
            public virtual string plannedTargetQuantity => "Planned Target Quantity:";
            public virtual string mxPlannedTargetQuantity => "mx-name-label2";
            public virtual string showAvailableSerialNumbers => "Show available Serial Numbers";
            public virtual string mxShowAvailableSerialNumbers => "mx-name-label4";
            public virtual string selectASerialNumberToAssociate => "Select a Serial Number to associate";
            public virtual string mxSelectASerialNumberToAssociate => "mx-name-label3";
            public virtual string associateOrAutoGenerateSerialNumber => "Associate or auto-generate Serial Number";
            public virtual string mxAssociateOrAutoGenerateSerialNumber => "mx-name-label6";
            public virtual string associateTheSerialNumberFoundInSystem => "Associate the Serial Number found in system";
            public virtual string mxAssociateTheSerialNumberFoundInSystem => "mx-name-label8";
            public virtual string textBox2 => "textBox2";
            public virtual string mxTextBox2 => "mx-name-textBox2";
            public virtual string or => "OR";
            public virtual string mxOr => "mx-name-label5";
            public virtual string textBox3 => "textBox3";
            public virtual string mxTextBox3 => "mx-name-textBox3";
            public virtual string woStartPanelSNGallery => "WOStartPanelSNGallery";
            public virtual string mxWoStartPanelSNGallery => "mx-name-WOStartPanelSNGallery";
            public virtual string woStartPanelEquipmentGallery => "WOStartPanelEquipmentGallery";
            public virtual string mxWoStartPanelEquipmentGallery => "mx-name-WOStartPanelEquipmentGallery";
        }
        
        public class ButtonsLocators
        {
            public virtual string associate => "Associate";
            public virtual string mxAssociate => "mx-name-AssociateSNFromDropDown";
            public virtual string generateAndAssociate => "Generate and Associate";
            public virtual string mxGenerateAndAssociate => "mx-name-GenerateAndAssociateSNFromNId";
            public virtual string generateAndAssociate_1 => "Generate and Associate";
            public virtual string mxGenerateAndAssociate_1 => "mx-name-GenerateAndAssociateSNFromQuantity";
            public virtual string associate_1 => "Associate";
            public virtual string mxAssociate_1 => "mx-name-AssociateSNFromInputNId";
            public virtual string generateAndAssociate_2 => "Generate and Associate";
            public virtual string mxGenerateAndAssociate_2 => "mx-name-GenerateAndAssociateSNFromNIdPlaceholder";
            public virtual string start => "Start";
            public virtual string mxStart => "mx-name-startActionButton";
            public virtual string start_1 => "Start";
            public virtual string mxStart_1 => "mx-name-startActionButton1";
            public virtual string startAll => "Start All";
            public virtual string mxStartAll => "mx-name-startActionButton2";
            public virtual string cancel => "Cancel";
            public virtual string mxCancel => "mx-name-cancelActionButton";
        }
        
        public class GalleryLocators
        {
            public virtual string woStartPanelSNGallery => "container8";
            public virtual string woStartPanelEquipmentGallery => "container4";
        }
    }
}
