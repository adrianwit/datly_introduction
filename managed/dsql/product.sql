SELECT product.* EXCEPT VENDOR_ID, vendor.*, performance.* EXCEPT product_id
FROM (SELECT * FROM PRODUCT t) product
JOIN (SELECT * FROM VENDOR t ) vendor  ON product.VENDOR_ID = vendor.ID AND 1 = 1
JOIN (
  SELECT location_id, product_id, SUM(quantity) AS quantity, AVG(payment) * 1.25 AS price
  FROM `$DB["bqdev"].bqdev.product_performance` t
  WHERE 1=1
#if($Unsafe.period == "today")
    AND TIMESTAMP_TRUNC(t.timestamp, DAY) = TIMESTAMP_TRUNC(CURRENT_TIMESTAMP(), DAY)
#elseif ($Unsafe.period == "yesterday")
    AND TIMESTAMP_TRUNC(t.timestamp, DAY) = TIMESTAMP_TRUNC(TIMESTAMP_ADD(CURRENT_TIMESTAMP(),
        INTERVAL -1 DAY),  DAY)
#end
GROUP BY 1, 2
) performance ON performance.product_id = product.ID





