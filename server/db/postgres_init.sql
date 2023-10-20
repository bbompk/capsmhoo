-- Create a new database
-- CREATE DATABASE capsmhoo;

-- Switch to the new database
-- \c capsmhoo;

-- Grant privileges
-- ALTER ROLE 'admin' SET client_encoding TO 'utf8';
-- ALTER ROLE 'admin' SET default_transaction_isolation TO 'read committed';
-- ALTER ROLE 'admin' SET timezone TO 'UTC';
-- GRANT ALL PRIVILEGES ON DATABASE capsmhoo TO 'admin';

-- SET default_tablespace = '';
-- SET default_table_access_method = heap;


-- Create a "user" table
CREATE TABLE users (
    id VARCHAR(255) PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);

-- Create a "professor" table with a foreign key to the "user" table
CREATE TABLE professors (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    profile TEXT,
    user_id VARCHAR(255) REFERENCES users(id)
);

-- Create a "team" table
CREATE TABLE teams (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    profile TEXT
);

-- Create a "professor" table with a foreign key to the "user" and "team" table
CREATE TABLE students (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    user_id VARCHAR(255) REFERENCES users(id),
    team_id VARCHAR(255) REFERENCES teams(id)
);

CREATE TABLE projects (
    project_id VARCHAR(255) PRIMARY KEY,
    team_id VARCHAR(255) REFERENCES teams(id),
    professor_id VARCHAR(255) REFERENCES users(id),
    name VARCHAR(255),
    description TEXT
);

CREATE TABLE team_join_requests (
    id VARCHAR(255) PRIMARY KEY,
    team_id VARCHAR(255) REFERENCES teams(id),
    student_id VARCHAR(255) REFERENCES students(id)
);