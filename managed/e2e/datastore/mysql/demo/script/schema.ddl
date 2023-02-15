DROP TABLE IF EXISTS invoice_list_item;
DROP TABLE IF EXISTS invoice;
DROP TABLE IF EXISTS trader;
DROP TABLE IF EXISTS VENDOR;
DROP TABLE IF EXISTS PRODUCT;
DROP TABLE IF EXISTS AUDIENCE;
DROP TABLE IF EXISTS DEALS;


CREATE TABLE invoice (
                         id INT PRIMARY KEY,
                         customer_name VARCHAR(255),
                         invoice_date DATE,
                         due_date DATE,
                         total_amount DECIMAL(10,2)
);



CREATE TABLE invoice_list_item (
                                   id INT PRIMARY KEY,
                                   invoice_id INT,
                                   product_name VARCHAR(255),
                                   quantity INT,
                                   price DECIMAL(10,2),
                                   total DECIMAL(10,2),
                                   FOREIGN KEY (invoice_id) REFERENCES invoice(id)
);

CREATE TABLE trader (
                        id INT PRIMARY KEY,
                        first_name VARCHAR(50) NOT NULL,
                        last_name VARCHAR(50) NOT NULL,
                        email VARCHAR(100) UNIQUE,
                        phone_number VARCHAR(20),
                        join_date DATE NOT NULL
);


CREATE TABLE VENDOR(
                       ID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
                       NAME  VARCHAR(255),
                       ACCOUNT_ID INT,
                       CREATED DATETIME,
                       USER_CREATED INT,
                       UPDATED DATETIME,
                       USER_UPDATED INT
);

CREATE TABLE PRODUCT (
                         ID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
                         NAME  VARCHAR(255),
                         VENDOR_ID INT,
                         STATUS INT,
                         CREATED DATETIME,
                         USER_CREATED INT,
                         UPDATED DATETIME,
                         USER_UPDATED INT
);



CREATE TABLE AUDIENCE(
    ID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    NAME  VARCHAR(255),
    MATCH_EXPRESSION VARCHAR(1024)
);


CREATE TABLE DEAL(
    ID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    NAME  VARCHAR(255),
    FEE DECIMAL(10,2)
);

