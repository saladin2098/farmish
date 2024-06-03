-- Create the enum type for feeding_index
CREATE TYPE feeding_index AS ENUM ('1', '2', '3');

-- Create the animals table
CREATE TABLE IF NOT EXISTS animals (
    id INT PRIMARY KEY,
    type VARCHAR(255),
    birth TIMESTAMP,
    weight FLOAT,
    avg_consumption FLOAT,
    avg_water FLOAT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at INT DEFAULT 0
);

-- Create the feeding_schedule table
CREATE TABLE IF NOT EXISTS feeding_schedules (
    id INT PRIMARY KEY,
    animal_type VARCHAR(255),
    last_fed_index feeding_index,
    sch1 INT DEFAULT 1,
    sch2 INT DEFAULT 2,
    sch3 INT DEFAULT 3
);

-- Create the schedules table
CREATE TABLE IF NOT EXISTS schedules (
    id INT PRIMARY KEY,
    animal_type VARCHAR(255),
    time1 TIME,
    time2 TIME,
    time3 TIME
);

-- Create the medications table
CREATE TABLE IF NOT EXISTS medications (
    id INT PRIMARY KEY,
    type VARCHAR(255),
    name VARCHAR(255),
    quantity FLOAT DEFAULT 0
);

-- Create the health_conditions table with a foreign key constraint
CREATE TABLE IF NOT EXISTS health_conditions (
    id INT PRIMARY KEY,
    animal_id INT REFERENCES animals(id),
    is_healthy BOOLEAN,
    condition VARCHAR(255),
    medication VARCHAR(255),
    is_treated BOOLEAN
);