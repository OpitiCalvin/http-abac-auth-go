-- CREATE DATABASE IF NOT EXISTS abac_auth_system;
-- USE abac_auth_system;

-- Create Partner Table
CREATE TABLE IF NOT EXISTS partner (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	partner_name VARCHAR(255) NOT NULL,
	created_at datetime DEFAULT CURRENT_TIMESTAMP,
	updated_at datetime
);

-- Create Product Table
CREATE TABLE IF NOT EXISTS product (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	product_name VARCHAR(255) NOT NULL,
	base_url VARCHAR(255) NOT NULL,
	created_at datetime DEFAULT CURRENT_TIMESTAMP,
	updated_at datetime	
);

-- Create Client Table
CREATE TABLE IF NOT EXISTS client (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	client_name VARCHAR(255),
	partner_id INTEGER,
	created_at datetime DEFAULT CURRENT_TIMESTAMP,
	updated_at datetime
);

-- Create ClientProducts Table
CREATE TABLE IF NOT EXISTS client_product (
	client_id INTEGER NOT NULL,
	product_id INTEGER NOT NULL,
	created_at datetime DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (client_id, product_id)
);

-- Create user table
CREATE TABLE IF NOT EXISTS user (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	email varchar(100) NOT NULL,
	username varchar(100) NOT NULL,
	password varchar(250) NOT NULL,
	client_id INTEGER NOT NULL,
	created_at datetime DEFAULT CURRENT_TIMESTAMP,
	updated_at datatime,
	FOREIGN KEY (client_id) REFERENCES client(id) ON DELETE CASCADE
);