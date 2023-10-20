# Custom datly

The custom project datly integration project

### Initialise new custom project 
In this step, new project with custom go module is created.
The goal of custom go module is to register custom data type to datly type registry 

```bash
cd github.com/adrianwit/datly_introduction/custom 
datly initExt  #initialises custom datly integration project see datly  initExt -h 
```

Generate patch for invoice/line items and all other input needed to business logic
[dsql/invoice/init/invoice_patch.sql](dsql/invoice/init/invoice_patch.sql)
```sql
SELECT
    invoice.* /* { "Cardinality": "One", "Field":"Entity" } */,
    item.*,
    product.*,
    acl.*,
    features.*
FROM (SELECT * FROM INVOICE) invoice
         JOIN (SELECT * FROM INVOICE_LIST_ITEM) item ON item.invoice_id = invoice.id
         LEFT JOIN  (SELECT ID, STATUS FROM (PRODUCT)  ) product ON item.product_id = product.id
         LEFT JOIN (SELECT * FROM DISCOUNT) discount ON discount.code = invoice.discount_code
         LEFT JOIN (SELECT ID AS USER_ID,
                           HasUserRole(ID, 'ADMIN') IS_ADMIN,
                           HasUserRole(ID, 'VIEWER') IS_VIEWER
                    FROM (USER) WHERE ID = 1
) acl  ON acl.USER_ID = invoice.user_created
         LEFT JOIN (SELECT ID AS USER_ID,
                           HasFeatureEnabled(ID, 'SET_DISCOUNT') CAN_SET_DISCOUNT
                    FROM (USER) WHERE ID = 1
) features  ON features.USER_ID = invoice.user_created
```

Generate Patch DSQL

```bash
datly gen -o=patch  -s=invoice_patch.sql -c='demo|mysql|root:dev@tcp(127.0.0.1:3306)/demo?parseTime=true' -g=invoice -p=datly_introduction/custom 
```


