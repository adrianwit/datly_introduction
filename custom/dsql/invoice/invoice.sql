/* {"ResponseBody":{"From":"Invoice"}} */

import (
	"custom.Invoice"
)


#set($_ = $Invoice<[]*Invoice>(body/))
$sequencer.Allocate("invoice", $Invoice, "Id")
$sequencer.Allocate("invoice_list_item", $Invoice, "Item/Id")
#foreach($recInvoice in $Unsafe.Invoice)
#if($recInvoice)
    $sql.Insert($recInvoice, "invoice");
      #foreach($recItem in $recInvoice.Item)
      #if($recItem)
          #set($recItem.InvoiceId = $recInvoice.Id)
          $sql.Insert($recItem, "invoice_list_item");
      #end
      #end
#end
#end