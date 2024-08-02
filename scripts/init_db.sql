CREATE DATABASE IF NOT EXISTS go_import_mng;

USE go_import_mng;

CREATE TABLE IF NOT EXISTS records (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    phone VARCHAR(15) NOT NULL
);
-- mysql -u root -p go_import_mng < scripts/init_db.sql
