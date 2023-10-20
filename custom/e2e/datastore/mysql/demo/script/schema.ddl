DROP TABLE IF EXISTS INVOICE_LIST_ITEM;
DROP TABLE IF EXISTS INVOICE;
DROP TABLE IF EXISTS VENDOR;
DROP TABLE IF EXISTS PRODUCT;
DROP TABLE IF EXISTS DISCOUNT_CODE;
DROP TABLE IF EXISTS USER;
DROP TABLE IF EXISTS USER_ROLE;
DROP TABLE IF EXISTS PLATFORM_ROLE;
DROP TABLE IF EXISTS USER_FEATURE;
DROP TABLE IF EXISTS PLATFORM_FEATURE;



CREATE TABLE USER (
      id INT AUTO_INCREMENT PRIMARY KEY,
      first_name VARCHAR(50) NOT NULL,
      last_name VARCHAR(50) NOT NULL,
      email VARCHAR(100) UNIQUE,
      phone_number VARCHAR(20),
      join_date DATE NOT NULL
);


CREATE TABLE VENDOR(
       id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
       name  VARCHAR(255),
       account_id INT,
       created DATETIME,
       user_created INT,
       updated DATETIME,
       user_updated INT
);

CREATE TABLE PRODUCT (
     id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
     name  VARCHAR(255),
     vendor_id INT,
     status INT,
     created DATETIME,
     user_created INT,
     updated DATETIME,
     user_updated INT,
     FOREIGN KEY (vendor_id) REFERENCES VENDOR(id)
);

CREATE TABLE DISCOUNT (
      code VARCHAR(255) PRIMARY KEY,
      pct DECIMAL(10,2),
      start_date DATETIME,
      end_date DATETIME
);


CREATE TABLE INVOICE (
     id INT AUTO_INCREMENT PRIMARY KEY,
     customer_name VARCHAR(255),
     customer_id INT,
     invoice_date DATE,
     due_date DATE,
     status int,
     total_amount DECIMAL(10,2),
     discount_code VARCHAR(255),
     created DATETIME,
     user_created INT,
     updated DATETIME,
     user_updated INT,
     FOREIGN KEY (discount_code) REFERENCES DISCOUNT(code),
     FOREIGN KEY (user_created) REFERENCES USER(id),
     FOREIGN KEY (user_updated) REFERENCES USER(id)
);



CREATE TABLE INVOICE_LIST_ITEM (
   id INT AUTO_INCREMENT PRIMARY KEY,
   invoice_id INT,
   product_id INT,
   product_name VARCHAR(255),
   quantity INT,
   price DECIMAL(10,2),
   total DECIMAL(10,2),
   created DATETIME,
   user_created INT,
   updated DATETIME,
   user_updated INT,
   FOREIGN KEY (invoice_id) REFERENCES INVOICE(id),
   FOREIGN KEY (user_created) REFERENCES USER(id),
   FOREIGN KEY (user_updated) REFERENCES USER(id),
   FOREIGN KEY (product_id) REFERENCES PRODUCT(id)
);

CREATE TABLE PLATFORM_ROLE (
                               id INT AUTO_INCREMENT PRIMARY KEY,
                               name VARCHAR(50) NOT NULL
);


CREATE TABLE USER_ROLE(
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    role_id INT,
    FOREIGN KEY (user_id) REFERENCES USER(id),
    FOREIGN KEY (role_id) REFERENCES PLATFORM_ROLE(id)
);



CREATE TABLE PLATFORM_FEATURE (
                                  id INT AUTO_INCREMENT PRIMARY KEY,
                                  name VARCHAR(50) NOT NULL
);


CREATE TABLE USER_FEATURE(
  id INT AUTO_INCREMENT PRIMARY KEY,
  user_id INT,
  feature_id INT,
  FOREIGN KEY (user_id) REFERENCES USER(id),
  FOREIGN KEY (feature_id) REFERENCES PLATFORM_FEATURE(id)
);




DROP  FUNCTION IF EXISTS  HasUserRole;

DELIMITER $$
CREATE FUNCTION HasUserRole(userId INT, role text)
    RETURNS BOOLEAN
BEGIN
  DECLARE hasRole INTEGER;

SELECT 1 INTO hasRole FROM USER_ROLE r
                      JOIN PLATFORM_ROLE p ON p.ID = r.role_id
                      WHERE  r.user_id = userID
                         AND p.name = role LIMIT 1;
IF hasRole IS NULL THEN
   return false;
END IF;
return true;
END
$$

DELIMITER ;



DROP  FUNCTION IF EXISTS  HasFeatureEnabled;


DELIMITER $$
CREATE FUNCTION HasFeatureEnabled(userId INT, feature text)
    RETURNS BOOLEAN
BEGIN
  DECLARE hasEnabled INTEGER;

SELECT 1 INTO hasEnabled FROM USER_FEATURE uf
    JOIN PLATFORM_FEATURE f ON f.id = uf.feature_id WHERE  f.name = feature
                                                     AND uf.user_id = userId LIMIT 1;
IF hasEnabled IS NULL THEN
   return false;
END IF;
return true;
END
$$

DELIMITER ;