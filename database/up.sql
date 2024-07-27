DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id VARCHAR(255) PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
DROP TABLE IF EXISTS characters;
CREATE TABLE characters (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    jname VARCHAR(255) NOT NULL,
    rname VARCHAR(255) NOT NULL
);
DROP TABLE IF EXISTS characters_images;
CREATE TABLE characters_images (
    id VARCHAR(255) PRIMARY KEY,
    character_id VARCHAR(255) NOT NULL,
    image_url VARCHAR(255) NOT NULL,
    FOREIGN KEY (character_id) REFERENCES characters(id)
);
DROP TABLE IF EXISTS devil_fruits;
CREATE TABLE devil_fruits (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL
);
DROP TABLE IF EXISTS characters_full;
CREATE TABLE characters_full (
    id VARCHAR(255) NOT NULL,
    age VARCHAR(2) NOT NULL,
    birth VARCHAR(255) NOT NULL,
    affiliation VARCHAR(255),
    occupation VARCHAR(255),
    origin VARCHAR(255),
    height VARCHAR(255),
    character_id VARCHAR(255) NOT NULL,
    devil_fruits_id VARCHAR(255),
    character_images_id VARCHAR(255),
    PRIMARY KEY (id, character_id),
    FOREIGN KEY (character_id) REFERENCES characters(id),
    FOREIGN KEY (devil_fruits_id) REFERENCES devil_fruits(id),
    FOREIGN KEY (character_images_id) REFERENCES characters_images(id)
);