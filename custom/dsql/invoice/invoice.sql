/* {"Method":"PATCH","ResponseBody":{"From":"Invoice"},
   "DateFormat":"2006-01-02 15:04:05",
   "CustomValidation":true
   } */

import (
	"invoice.Invoice"
	"invoice.Item"
	"invoice.Product"
	"invoice.Discount"
	"invoice.Acl"
	"invoice.Features"
)

#set($_ = $Jwt<string>(Header/Authorization).WithCodec(JwtClaim).WithStatusCode(401))


#set($_ = $Invoice<*Invoice>(body/Entity))

#set($_ = $CurInvoice<*Invoice>(data_view/CurInvoice) /* ?
  SELECT * FROM INVOICE
  WHERE id = $Invoice.Id
  */
)

#set($_ = $CurItem<[]*Item>(data_view/CurItem) /* ? 
  SELECT * FROM INVOICE_LIST_ITEM
  WHERE invoice_id = $Invoice.Id
  */
)

#set($_ = $ProductIds<?>(param/Invoice) /*
   ? SELECT ARRAY_AGG(ProductId) AS Values FROM `/Item/` LIMIT 1
   */
)

#set($_ = $CurProduct<[]*Product>(data_view/CurProduct) /* ? 
  SELECT ID, STATUS FROM PRODUCT WHERE ID $criteria.In($ProductIds)
 */
)

#set($_ = $CurDiscount<*Discount>(data_view/CurDiscount) /* ?
  SELECT * FROM DISCOUNT  WHERE CODE = $Invoice.DiscountCode
  */
)

#set($_ = $Acl<[]*Acl>(data_view/Acl) /* ?
  SELECT ID AS USER_ID,
         HasUserRole(ID, 'ADMIN') IS_ADMIN,
         HasUserRole(ID, 'VIEWER') IS_VIEWER
    FROM (USER) WHERE ID = $Jwt.UserID
  */
)

#set($_ = $Features<[]*Features>(data_view/Features) /* ?
  SELECT ID AS USER_ID,
    HasFeatureEnabled(ID, 'SET_DISCOUNT') CAN_SET_DISCOUNT
  FROM (USER) WHERE ID = $Jwt.UserID
  */
)


$sequencer.Allocate("INVOICE", $Invoice, "Id")
$sequencer.Allocate("INVOICE_LIST_ITEM", $Invoice, "Item/Id")
#set($invoiceById = $CurInvoice.IndexBy("Id"))
#set($itemById = $CurItem.IndexBy("Id"))

#if($Unsafe.Invoice)
  #set($init = $Invoice.Init($CurInvoice, $Acl, $Features, $CurDiscount))
  #set($info = $Invoice.Validate($CurInvoice, $Acl, $CurProduct))
  #if($info.Failed ==  true)
        $response.StatusCode(401)
        $response.Failf("%v",$info.Error)
  #end

  #if(($invoiceById.HasKey($Unsafe.Invoice.Id) == true))
    $sql.Update($Invoice, "INVOICE");
  #else
    $sql.Insert($Invoice, "INVOICE");
  #end
    #foreach($recItem in $Unsafe.Invoice.Item)
        #if(($itemById.HasKey($recItem.Id) == true))
          $sql.Update($recItem, "INVOICE_LIST_ITEM");
        #else
          $sql.Insert($recItem, "INVOICE_LIST_ITEM");
        #end
    #end
#end