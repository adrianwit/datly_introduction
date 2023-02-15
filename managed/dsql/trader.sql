SELECT
    trader.*,
    acl.*
FROM (SELECT * FROM trader) trader
JOIN (SELECT
          USER_ID,
          ARRAY_EXISTS(ROLE, 'READ_ONLY') AS IS_READONLY,
          ARRAY_EXISTS(PERMISSION, 'FEATURE1') AS CAN_USE_FEATURE1
    FROM $DB["dyndb"].USER_ACL ) acl ON acl.USER_ID = trader.id AND 1=1
