using MasterDetailsBase.Widget;
using Microsoft.Playwright;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace OperatorTerminals_EXDS.SidePanels
{
    public class StartPanel    {


        public StartPanelLocators Locators { get; protected set; } = new StartPanelLocators();
        public class StartPanelLocators
        {
            public virtual string panelTitle => "Start Work Order Operation";
            public virtual string quantity => "Quantity";
            public virtual string mxQuantity => "mx-name-StartQuantityOperation";
            public virtual string mxPanelClass => "mx-name-dataView1";
            public virtual string mxStartButton => "mx-name-startActionButton";
        }


    }
}
