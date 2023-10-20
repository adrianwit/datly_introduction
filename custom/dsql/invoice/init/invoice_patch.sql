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