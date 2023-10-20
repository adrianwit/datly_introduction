## Managed mode

This project demonstrate datly usge with mysql/BigQuery and DynamoDB


## Prerequisites
All you need to build, to run API and run e2e is the latest version of 
- Viant [Datly](https://github.com/viant/datly) 
- [Endly](https://github.com/viant/endly/releases/tag/v0.64.1) 



## Getting Started
Now after you have installed [Endly](https://github.com/viant/endly/releases/tag/v0.64.1) and [Datly](https://github.com/viant/datly) tools in you `/usr/local/bin`, you need to:<br>
1. To generate credentials file to enable endly exec service to run on localhost:
   Provide a username and password to login to your box.
   ```shell
   mkdir $HOME/.secret
   ssh-keygen -b 1024 -t rsa -f id_rsa -P "" -f $HOME/.secret/id_rsa
   touch ~/.ssh/authorized_keys
   cat $HOME/.secret/id_rsa.pub >>  ~/.ssh/authorized_keys
   chmod u+w authorized_keys
   endly -c=localhost -k=~/.secret/id_rsa
   ```
   

2. Generate secrets for mysql, execure next command with Username:`root` and Password, `dev`:
    ```shell
    endly -c=mysql-e2e -endpoint=127.0.0.1 
    ```
3. Generate GCP secret credentials as documented blow
   -  https://github.com/viant/endly/tree/master/doc/secrets#google-cloud-credentials
      Copy created service account secret  to ~/.secret/gcp-e2e.json   

3. To run full workflow of building, running and testing Datly endpoints go to `e2e/` folder and run `endly`:
    ```shell
    cd e2e
    endly authWith=gcp-e2e
    ```
Now you should see all execution flow and tests results in your console logs.<br>
Datly API remains running on 8080 port.



## Reader Service

### Use Cases

#### Invoice Service

```bash
   
    datly dsql -c='demo|mysql|demo:demo@tcp(127.0.0.1:3306)/demo?parseTime=true' \
               -s=invoice.sql -P=8080
    open http://127.0.0.1:8080/v1/api/meta/struct/dev/   
    
```


#### Trader Service
```bash
    datly dsql -c='demo|mysql|demo:demo@tcp(127.0.0.1:3306)/demo?parseTime=true' \
               -c='dyndb|dynamodb|dynamodb://localhost:8000/us-west-1?key=dummy&secret=dummy' \
               -s=trader.sql -P=8080
    open http://127.0.0.1:8080/v1/api/meta/struct/dev/    
```

#### Product Service

```bash
    datly dsql -c='demo|mysql|demo:demo@tcp(127.0.0.1:3306)/demo?parseTime=true' \
                -c='bqdev|bigquery|bigquery://viant-e2e/bqdev'  \
      -s=product.sql -P=8080
    open http://127.0.0.1:8080/v1/api/meta/struct/dev/    
```

#### Audience Service

```bash
    datly dsql -c='demo|mysql|demo:demo@tcp(127.0.0.1:3306)/demo?parseTime=true' \
      -s=audience.sql -P=8080
    open http://127.0.0.1:8080/v1/api/meta/struct/dev/    
 
```



## Executor Service

Generate insert operation

```bash
    cd dsql/init

    datly gen -o=post -s=invoice_ins.sql -g=domain \
     -c='demo|mysql|demo:demo@tcp(127.0.0.1:3306)/demo?parseTime=true'
```
As a resul the following files get created:

- **init/pkg/domain.entity.go**
```go
type Entity struct {
	Entity []*Invoice `typeName:"Invoice"`
}

type Invoice struct {
	Id           int         `sqlx:"name=id,autoincrement,primaryKey,required"`
	CustomerName *string     `sqlx:"name=customer_name" json:",omitempty"`
	InvoiceDate  *time.Time  `sqlx:"name=invoice_date" json:",omitempty"`
	DueDate      *time.Time  `sqlx:"name=due_date" json:",omitempty"`
	TotalAmount  *float64    `sqlx:"name=total_amount" json:",omitempty"`
	Item         []*Item     `typeName:"Item" sqlx:"-" datly:"relName=item,relColumn=id,refColumn=invoice_id,refTable=invoice_list_item" sql:"SELECT * FROM invoice_list_item"`
	Has          *InvoiceHas `presenceIndex:"true" typeName:"InvoiceHas" json:"-" diff:"presence=true" sqlx:"presence=true" validate:"presence=true"`
}

type Item struct {
	Id          int      `sqlx:"name=id,autoincrement,primaryKey,required"`
	InvoiceId   *int     `sqlx:"name=invoice_id,refTable=invoice,refColumn=id" json:",omitempty"`
	ProductName *string  `sqlx:"name=product_name" json:",omitempty"`
	Quantity    *int     `sqlx:"name=quantity" json:",omitempty"`
	Price       *float64 `sqlx:"name=price" json:",omitempty"`
	Total       *float64 `sqlx:"name=total" json:",omitempty"`
	Has         *ItemHas `presenceIndex:"true" typeName:"ItemHas" json:"-" diff:"presence=true" sqlx:"presence=true" validate:"presence=true"`
}
```

- **init/dsql/invoice.sql**
```sql
/* {"Method":"POST","ResponseBody":{"From":"Invoice"}} */

import (

	"domain.Invoice"
	"domain.Entity"
	"domain.Item"
)


#set($_ = $Invoice<*Invoice>(body/Entity))


$sequencer.Allocate("invoice", $Invoice, "Id")
$sequencer.Allocate("invoice_list_item", $Invoice, "Item/Id")
#if($Unsafe.Invoice)
  $sql.Insert($Invoice, "invoice");
    #foreach($recItem in $Unsafe.Invoice.Item)
    #if($recItem)
        #set($recItem.InvoiceId = $Unsafe.Invoice.Id)
        $sql.Insert($recItem, "invoice_list_item");
    #end
    #end
#end
```

