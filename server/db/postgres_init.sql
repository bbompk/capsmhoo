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
    role VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);

-- Create a "professor" table with a foreign key to the "user" table
CREATE TABLE professors (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    profile TEXT,
    user_id VARCHAR(255) REFERENCES users(id) ON DELETE CASCADE
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
    user_id VARCHAR(255) REFERENCES users(id) ON DELETE CASCADE,
    team_id VARCHAR(255) REFERENCES teams(id) ON DELETE CASCADE
);



CREATE TABLE projects (
    project_id VARCHAR(255) PRIMARY KEY,
    team_id VARCHAR(255) REFERENCES teams(id) ON DELETE CASCADE,
    professor_id VARCHAR(255) REFERENCES professors(id) ON DELETE CASCADE,
    name VARCHAR(255),
    description TEXT,
    status VARCHAR(255),
    label VARCHAR(255)
);

-- Create a "project_request" table with a foreign key to the "project" and "team" table
CREATE TABLE project_requests (
    project_request_id VARCHAR(255) PRIMARY KEY,
    team_id VARCHAR(255) REFERENCES teams(id) ON DELETE CASCADE,
    project_id VARCHAR(255) REFERENCES projects(project_id) ON DELETE CASCADE,
    message VARCHAR(255),
    status VARCHAR(255)
);

CREATE TABLE team_join_requests (
    id VARCHAR(255) PRIMARY KEY,
    team_id VARCHAR(255) REFERENCES teams(id) ON DELETE CASCADE,
    student_id VARCHAR(255) REFERENCES students(id) ON DELETE CASCADE
);