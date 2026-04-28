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
            public virtual string plannedTargetQuantity => "Planned Target Quantity:";
            public virtual string showAvailableSerialNumbers => "Show available Serial Numbers";
            public virtual string selectASerialNumberToAssociate => "Select a Serial Number to associate";
            public virtual string associateOrAutoGenerateSerialNumber => "Associate or auto-generate Serial Number";
            public virtual string associateTheSerialNumberFoundInSystem => "Associate the Serial Number found in system";
            public virtual string textBox2 => "textBox2";
            public virtual string or => "OR";
            public virtual string textBox3 => "textBox3";
            public virtual string woStartPanelSNGallery => "WOStartPanelSNGallery";
            public virtual string woStartPanelEquipmentGallery => "WOStartPanelEquipmentGallery";
        }
        
        public class ButtonsLocators
        {
            public virtual string associate => "Associate";
            public virtual string generateAndAssociate => "Generate and Associate";
            public virtual string generateAndAssociate_1 => "Generate and Associate";
            public virtual string associate_1 => "Associate";
            public virtual string generateAndAssociate_2 => "Generate and Associate";
            public virtual string start => "Start";
            public virtual string start_1 => "Start";
            public virtual string startAll => "Start All";
            public virtual string cancel => "Cancel";
        }
        
        public class GalleryLocators
        {
            public virtual string woStartPanelSNGallery => "container8";
            public virtual string woStartPanelEquipmentGallery => "container4";
        }
    }
}
