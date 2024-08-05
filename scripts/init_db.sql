CREATE DATABASE IF NOT EXISTS go_import_mng;

-- Connect to the newly created database
USE go_import_mng;

CREATE TABLE IF NOT EXISTS records (
    id INT AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    company_name VARCHAR(100),
    address VARCHAR(255),
    city VARCHAR(100),
    county VARCHAR(100),
    postal VARCHAR(20),
    phone VARCHAR(15),
    email VARCHAR(100) NOT NULL,
    web VARCHAR(100)
);

-- mysql -u root -p go_import_mng < scripts/init_db.sql
