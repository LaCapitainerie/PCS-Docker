DROP TABLE IF EXISTS sidebar;
DROP TABLE IF EXISTS message;
DROP TABLE IF EXISTS ticket;
DROP TABLE IF EXISTS chat_user;
DROP TABLE IF EXISTS chat;
DROP TABLE IF EXISTS review_lessor_to_service;
DROP TABLE IF EXISTS review_traveler_to_service;
DROP TABLE IF EXISTS review_traveler_to_property;
DROP TABLE IF EXISTS reservation_service;
DROP TABLE IF EXISTS reservation_bill;
DROP TABLE IF EXISTS reservation;
DROP TABLE IF EXISTS property_service;
DROP TABLE IF EXISTS bill;
DROP TABLE IF EXISTS provider_licence;
DROP TABLE IF EXISTS service_type;
DROP TABLE IF EXISTS type_of_service;
DROP TABLE IF EXISTS service_unavailability;
DROP TABLE IF EXISTS service;
DROP TABLE IF EXISTS property_image;
DROP TABLE IF EXISTS property_unavailability;
DROP TABLE IF EXISTS property;
DROP TABLE IF EXISTS subscribe_traveler;
DROP TABLE IF EXISTS subscribe;
DROP TABLE IF EXISTS lessor;
DROP TABLE IF EXISTS provider;
DROP TABLE IF EXISTS traveler;
DROP TABLE IF EXISTS administrator;
DROP TABLE IF EXISTS users;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID PRIMARY KEY,
    mail VARCHAR(320) NOT NULL,
    password VARCHAR(64) NOT NULL,
    avatar VARCHAR(255),
    phone_number VARCHAR(15),
    type         VARCHAR(15),
    description TEXT,
    register_date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_connection_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE administrator (
    id  UUID PRIMARY KEY,
    site VARCHAR(64),
    nickname VARCHAR(64) NOT NULL,
    user_id  UUID NOT NULL,
    FOREIGN KEY (user_id ) REFERENCES users(id)
);

CREATE TABLE traveler (
    id  UUID PRIMARY KEY,
    first_name VARCHAR(64) NOT NULL,
    last_name VARCHAR(64) NOT NULL,
    user_id  UUID NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE provider (
    id  UUID PRIMARY KEY,
    first_name VARCHAR(64) NOT NULL,
    last_name VARCHAR(64) NOT NULL,
    nickname VARCHAR(64) NOT NULL,
    user_id  UUID NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE lessor (
    id  UUID PRIMARY KEY,
    first_name VARCHAR(64) NOT NULL,
    last_name VARCHAR(64) NOT NULL,
    user_id  UUID NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE subscribe (
    id  UUID PRIMARY KEY,
    type VARCHAR(64) NOT NULL,
    price NUMERIC(10, 2) NOT NULL
);

CREATE TABLE subscribe_traveler (
    id  UUID PRIMARY KEY,
    begin_date TIMESTAMP WITH TIME ZONE NOT NULL,
    end_date TIMESTAMP WITH TIME ZONE NOT NULL,
    traveler_id  UUID NOT NULL,
    subscribe_id  UUID NOT NULL,
    FOREIGN KEY (traveler_id) REFERENCES traveler(id),
    FOREIGN KEY (subscribe_id) REFERENCES subscribe(id)
);

CREATE TABLE property (
    id  UUID PRIMARY KEY,
    name VARCHAR(64) NOT NULL,
    type VARCHAR(64) NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    surface INTEGER NOT NULL,
    room INTEGER NOT NULL,
    bathroom INTEGER NOT NULL,
    garage INTEGER,
    description TEXT,
    address VARCHAR(64) NOT NULL,
    city VARCHAR(64) NOT NULL,
    zip_code VARCHAR(6) NOT NULL,
    country VARCHAR(64) NOT NULL,
    lon DOUBLE PRECISION,
    lat DOUBLE PRECISION,
    administrator_validation BOOLEAN DEFAULT FALSE,
    lessor_id  UUID NOT NULL,
    FOREIGN KEY (lessor_id) REFERENCES lessor(id)
);

CREATE TABLE property_unavailability (
    id  UUID PRIMARY KEY,
    begin_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL,
    property_id  UUID NOT NULL,
    FOREIGN KEY (property_id) REFERENCES property(id)
);

CREATE TABLE property_image (
    id  UUID PRIMARY KEY,
    path VARCHAR(255) NOT NULL,
    property_id  UUID NOT NULL,
    FOREIGN KEY (property_id) REFERENCES property(id)
);

CREATE TABLE service(
    id  UUID PRIMARY KEY,
    price NUMERIC(10,2) NOT NULL,
    target_customer VARCHAR(8) NOT NULL, -- Peut prendre que les valeurs "all" "lessor" ou "traveler"
    address VARCHAR(64) NOT NULL,
    city VARCHAR(64) NOT NULL,
    zip_code VARCHAR(6) NOT NULL,
    country VARCHAR(64) NOT NULL,
    range_action INTEGER, -- null = infinie
    description TEXT,
    provider_id  UUID NOT NULL,
    FOREIGN KEY (provider_id) REFERENCES provider(id)
);

CREATE TABLE service_unavailability(
    id  UUID PRIMARY KEY,
    begin_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL,
    service_id  UUID NOT NULL,
    FOREIGN KEY (service_id) REFERENCES service(id)
);

CREATE TABLE type_of_service(
    id  UUID PRIMARY KEY,
    type VARCHAR(64) NOT NULL,
    licence BOOLEAN DEFAULT FALSE
);

CREATE TABLE service_type
(
    service_id  UUID NOT NULL,
    type_of_service_id  UUID NOT NULL,
    FOREIGN KEY (service_id ) REFERENCES service (id),
    FOREIGN KEY (type_of_service_id ) REFERENCES type_of_service (id),
    PRIMARY KEY (service_id, type_of_service_id)
);

CREATE TABLE provider_licence(
    id  UUID PRIMARY KEY,
    begin_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP,
    validation BOOLEAN DEFAULT FALSE,
    path_proof VARCHAR(255) NOT NULL,
    provider_id  UUID NOT NULL,
    FOREIGN KEY (provider_id) REFERENCES provider(id)
);

CREATE TABLE bill(
    id  UUID PRIMARY KEY,
    price NUMERIC(10, 2) NOT NULL,
    date TIMESTAMP NOT NULL,
    type VARCHAR(64),
    content TEXT
);

CREATE TABLE property_service(
    property_id  UUID NOT NULL,
    service_id  UUID NOT NULL,
    bill_id  UUID NOT NULL,
    date TIMESTAMP NOT NULL,
    FOREIGN KEY (bill_id) REFERENCES bill(id),
    FOREIGN KEY (property_id) REFERENCES property(id),
    FOREIGN KEY (service_id) REFERENCES service(id),
    PRIMARY KEY (property_id, service_id)
);

CREATE TABLE reservation(
    id  UUID PRIMARY KEY,
    traveler_id  UUID NOT NULL,
    property_id  UUID NOT NULL,
    begin_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL,
    FOREIGN KEY (traveler_id) REFERENCES traveler(id),
    FOREIGN KEY (property_id) REFERENCES property(id)
);

CREATE TABLE reservation_bill(
    reservation_id  UUID NOT NULL,
    bill_id  UUID NOT NULL,
    FOREIGN KEY (reservation_id) REFERENCES reservation(id),
    FOREIGN KEY (bill_id) REFERENCES bill(id),
    PRIMARY KEY (reservation_id, bill_id)
);

CREATE TABLE reservation_service(
    reservation_id  UUID NOT NULL,
    service_id  UUID NOT NULL,
    date TIMESTAMP NOT NULL,
    FOREIGN KEY (reservation_id ) REFERENCES reservation(id),
    FOREIGN KEY (service_id ) REFERENCES service(id),
    PRIMARY KEY (reservation_id, service_id)
);

CREATE TABLE review_traveler_to_property(
    traveler_id  UUID NOT NULL,
    property_id  UUID NOT NULL,
    note numeric(10, 1) NOT NULL,
    comment TEXT,
    FOREIGN KEY (traveler_id ) REFERENCES traveler(id),
    FOREIGN KEY (property_id ) REFERENCES property(id),
    PRIMARY KEY (traveler_id, property_id)
);

CREATE TABLE review_traveler_to_service(
    traveler_id  UUID NOT NULL,
    service_id  UUID NOT NULL,
    note numeric(10, 1) NOT NULL,
    comment TEXT,
    FOREIGN KEY (traveler_id ) REFERENCES traveler(id),
    FOREIGN KEY (service_id ) REFERENCES service(id),
    PRIMARY KEY (traveler_id, service_id)
);

CREATE TABLE review_lessor_to_service (
    lessor_id  UUID NOT NULL,
    service_id  UUID NOT NULL,
    note numeric(10, 1) NOT NULL,
    comment TEXT,
    FOREIGN KEY (lessor_id) REFERENCES lessor(id),
    FOREIGN KEY (service_id) REFERENCES service(id),
    PRIMARY KEY (lessor_id, service_id)
);

CREATE TABLE chat (
    id  UUID PRIMARY KEY,
    view BOOLEAN DEFAULT FALSE
);

CREATE TABLE chat_user (
    user_id UUID,
    chat_id UUID,
    PRIMARY KEY (user_id, chat_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (chat_id) REFERENCES chat(id)
);

CREATE TABLE ticket
(
    id         UUID PRIMARY KEY,
    type        VARCHAR(64) NOT NULL,
    state       VARCHAR(16) NOT NULL,
    description TEXT        NOT NULL,
    chat_id   UUID        NOT NULL,
    FOREIGN KEY (chat_id) REFERENCES chat (id)
);

CREATE TABLE message (
    id  UUID PRIMARY KEY,
    content TEXT NOT NULL,
    date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    type VARCHAR(5), -- "text" ou "image"
    user_id  UUID NOT NULL,
    chat_id  UUID NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (chat_id) REFERENCES chat(id)
);

CREATE TABLE sidebar (
    id UUID PRIMARY KEY,
    permission INT,
    icon VARCHAR(255),
    hover VARCHAR(128),
    href VARCHAR(255)
);

CREATE TABLE log (
    id UUID PRIMARY KEY,
    from_user_id UUID,
    date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    type VARCHAR(64) NOT NULL,
    description TEXT NOT NULL
);



-- Insertion des données de base --

INSERT INTO sidebar (id, permission, icon, hover, href)
VALUES
    (uuid_generate_v4(), 1, 'Home', 'property', '/traveler/property'),
    (uuid_generate_v4(), 1, 'Msg', 'messages', '/traveler/messages'),
    (uuid_generate_v4(), 2, 'Home', 'property', '/provider/property'),
    (uuid_generate_v4(), 2, 'Msg', 'messages', '/provider/messages'),
    (uuid_generate_v4(), 2, 'Gauge', 'dashboard', '/provider/dashboard'),
    (uuid_generate_v4(), 3, 'Home', 'property', '/lessor/property'),
    (uuid_generate_v4(), 3, 'Msg', 'messages', '/lessor/messages'),
    (uuid_generate_v4(), 3, 'Gauge', 'dashboard', '/lessor/dashboard'),
    (uuid_generate_v4(), 4, 'Gauge', 'dashboard', '/admin/dashboard');

INSERT INTO users (id, mail, password, avatar, description, register_date, last_connection_date, phone_number) VALUES
    (
        'a0e12f8a-4776-4ed3-91d5-673fcef79d5c',
        'traveler@gmail.com',
        '$2a$10$07qrkk/uZXKX47WWwO.OVO2Ims8sJz2V8pUKleziMhcWW49xXvoYy',
        '',
        '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, '06 23 00 01 47'
    ),
    (
        'c3c99ccc-4844-4f78-9b27-8daabbc7f8f8',
        'provider@gmail.com',
        '$2a$10$16FESXE6QgyRPTOif2Mwfu/5ocarRihkfk3vuL0YfgffTmtiqSbFu',
        '',
        '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, '06 21 29 77 06'
    ),
    (
        'b2a88bbb-3822-4d56-8b36-7c9a44dc6a7e',
        'lessor@gmail.com',
        '$2y$10$7nrgYo5DHC6Pr1eeLWX5GuFQKn082oAETDRxIc1PRtBD/o1UMT10e',
        '',
        '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, '06 29 24 07 06'
    ),
    (
        'a0e12f8a-4776-4ed3-91d5-673cef79c3ec',
        'admin@gmail.com',
        '$2a$10$39lQGnhapB17TmtXQduhhuHOqYrr2yJ1FlE.SOCW4nvI4qutiOvnS',
        '',
        '', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, '06 21 69 77 06'
    );

INSERT INTO administrator (id, site, nickname, user_id)
VALUES
    (uuid_generate_v4(), 'Paris', 'Antoine', 'a0e12f8a-4776-4ed3-91d5-673cef79c3ec');

INSERT INTO lessor (id, first_name, last_name, user_id)
VALUES
    (uuid_generate_v4(), 'Marie', 'Durand', 'b2a88bbb-3822-4d56-8b36-7c9a44dc6a7e');

INSERT INTO provider (id, first_name, last_name, nickname, user_id)
VALUES
    (uuid_generate_v4(), 'Sophie', 'Bernard', 'Sophie', 'c3c99ccc-4844-4f78-9b27-8daabbc7f8f8');

INSERT INTO traveler (id, first_name, last_name, user_id)
VALUES
    (uuid_generate_v4(), 'Paul', 'Lefevre', 'a0e12f8a-4776-4ed3-91d5-673fcef79d5c');






