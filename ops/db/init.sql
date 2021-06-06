-- CREATE DATABASE IF NOT EXISTS abac_auth_system;
-- USE abac_auth_system;

-- Create Partner Table
CREATE TABLE IF NOT EXISTS partner (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	name VARCHAR(255) NOT NULL,
	created_at datetime DEFAULT CURRENT_TIMESTAMP,
	updated_at datetime
);

-- Create Product Table
CREATE TABLE IF NOT EXISTS product (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	name VARCHAR(255) NOT NULL,
	base_url VARCHAR(255) NOT NULL,
	created_at datetime DEFAULT CURRENT_TIMESTAMP,
	updated_at datetime	
);

-- Create Client Table
CREATE TABLE IF NOT EXISTS client (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	name VARCHAR(255),
	products int[] DEFAULT '{}',
	partner_id INTEGER,
	created_at datetime DEFAULT CURRENT_TIMESTAMP,
	updated_at datetime
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