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
DROP TABLE IF EXISTS log;
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
    ('a0e12f8a-4776-4ed3-91d5-673fcef79d5c', 'lessor1@example.com', '$2y$10$7nrgYo5DHC6Pr1eeLWX5GuFQKn082oAETDRxIc1PRtBD/o1UMT10e', 'https://example.com/avatar1.jpg', 'Description de user1', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, '0123456789'),
    ('b2a88bbb-3822-4d56-8b36-7c9a44dc6a7e', 'lessor2@example.com', '$2y$10$7nrgYo5DHC6Pr1eeLWX5GuFQKn082oAETDRxIc1PRtBD/o1UMT10e', 'https://example.com/avatar2.jpg', 'Description de user2', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, '9876543210'),
    ('c3c99ccc-4844-4f78-9b27-8daabbc7f8f8', 'lessor3@example.com', '$2y$10$7nrgYo5DHC6Pr1eeLWX5GuFQKn082oAETDRxIc1PRtBD/o1UMT10e', 'https://example.com/avatar3.jpg', 'Description de user3', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, '1234567890'),
    ('a0e12f8a-4776-4ed3-91d5-673cef79c3ec', 'provider@example.com', '$2y$10$7nrgYo5DHC6Pr1eeLWX5GuFQKn082oAETDRxIc1PRtBD/o1UMT10e', 'https://example.com/avatar1.jpg', 'Description de user1', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, '0123696789'),
    ('b2a88bbb-3822-4d56-8b36-7c9a4489b7ea', 'traveler@example.com', '$2y$10$7nrgYo5DHC6Pr1eeLWX5GuFQKn082oAETDRxIc1PRtBD/o1UMT10e', 'https://example.com/avatar2.jpg', 'Description de user2', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, '9821543210');

INSERT INTO lessor (id, first_name, last_name, user_id) VALUES
    ('98765432-12d3-e456-b426-426614174000', 'John', 'Doe', 'a0e12f8a-4776-4ed3-91d5-673fcef79d5c'),
    ('123e4567-e89b-12d3-a456-426614174000', 'Jane', 'Smith', 'b2a88bbb-3822-4d56-8b36-7c9a44dc6a7e'),
    ('647d216d-d534-4c7e-b1f1-0c2d815bd3f4', 'Emily', 'Brown', 'c3c99ccc-4844-4f78-9b27-8daabbc7f8f8');

INSERT INTO provider (id, first_name, last_name, nickname, user_id) VALUES
    ('123e4567-e89b-13f3-a456-426618174000', 'Jerry', 'Escobar', 'Escopuelo', 'a0e12f8a-4776-4ed3-91d5-673cef79c3ec');

INSERT INTO traveler (id, first_name, last_name, user_id) VALUES
    ('123e4567-e89b-13f3-a456-426618174901', 'Noah', 'Picard', 'b2a88bbb-3822-4d56-8b36-7c9a4489b7ea');

INSERT INTO property (id, name, type, price, surface, room, bathroom, garage, description, address, city, zip_code, country, administrator_validation, lessor_id)
VALUES
    ('f47ac10b-58cc-4372-a567-0e02b2c3d479', 'Belle Maison en Centre-Ville', 'Maison', 250000.00, 180, 5, 3, 1, 'Belle maison située en plein centre-ville, proche des commerces et écoles.', '12 Rue de la Liberté', 'Paris', '75001', 'France', TRUE, '123e4567-e89b-12d3-a456-426614174000'),
    ('1ed3b7b1-f37b-4b5f-8e6b-382fae57640a', 'Appartement Moderne avec Vue sur Mer', 'Appartement', 150000.00, 90, 3, 2, 0, 'Appartement récemment rénové offrant une vue imprenable sur la mer.', '8 Rue des Palmiers', 'Nice', '06000', 'France', FALSE, '123e4567-e89b-12d3-a456-426614174000'),
    ('ab9d50e8-3b15-4a43-95aa-41745c87ff5e', 'Villa de Luxe avec Piscine', 'Villa', 750000.00, 300, 7, 5, 1, 'Villa de luxe avec piscine privée, jardin et vue panoramique sur la ville.', '25 Avenue des Roses', 'Cannes', '06400', 'France', TRUE, '98765432-12d3-e456-b426-426614174000'),
    ('7beed29c-2623-4b67-baf2-73c284f0f39a', 'Maison Traditionnelle avec Jardin', 'Maison', 180000.00, 150, 4, 2, 1, 'Charmante maison traditionnelle avec un grand jardin arboré.', '4 Rue des Chênes', 'Bordeaux', '33000', 'France', FALSE, '98765432-12d3-e456-b426-426614174000'),
    ('6d3474bb-218e-497f-bdc9-774af449e215', 'Appartement Cosy en Centre Historique', 'Appartement', 120000.00, 75, 2, 1, 0, 'Appartement cosy situé au cœur du centre historique, à proximité des monuments.', '6 Place du Marché', 'Strasbourg', '67000', 'France', TRUE, '123e4567-e89b-12d3-a456-426614174000'),
    ('d40e5e8d-1a26-41a0-b65a-0a30ed21e77f', 'Villa Familiale avec Vue sur les Montagnes', 'Villa', 400000.00, 250, 6, 4, 1, 'Superbe villa familiale offrant une vue panoramique sur les montagnes environnantes.', '10 Chemin des Cimes', 'Grenoble', '38000', 'France', TRUE, '98765432-12d3-e456-b426-426614174000'),
    ('ad65e803-81d5-4d02-97ab-503c6eab6f9f', 'Maison de Campagne avec Grand Terrain', 'Maison', 220000.00, 200, 5, 3, 1, 'Maison de campagne avec un grand terrain, idéale pour les amoureux de la nature.', '2 Route des Champs', 'Lyon', '69000', 'France', FALSE, '123e4567-e89b-12d3-a456-426614174000'),
    ('c18dfc9f-4d96-4d14-af5a-2e0332876e5d', 'Appartement Lumineux avec Balcon', 'Appartement', 95000.00, 60, 2, 1, 0, 'Appartement lumineux avec un balcon offrant une vue dégagée.', '15 Avenue du Soleil', 'Marseille', '13000', 'France', TRUE, '98765432-12d3-e456-b426-426614174000'),
    ('d43e501e-77c3-42c4-a9a2-42f013e1a5b1', 'Villa Moderne avec Piscine et Spa', 'Villa', 680000.00, 320, 8, 5, 1, 'Villa moderne équipée d une piscine, d un spa et d un grand jardin.', '18 Boulevard des Palmiers', 'Nice', '06000', 'France', FALSE, '123e4567-e89b-12d3-a456-426614174000'),
    ('7fc56270-a7a7-4ec5-9ec1-57c5860b0026', 'Maison de Ville avec Cour Intérieure', 'Maison', 195000.00, 120, 4, 2, 1, 'Maison de ville avec une charmante cour intérieure, proche des commodités.', '3 Rue des Moulins', 'Lille', '59000', 'France', TRUE, '98765432-12d3-e456-b426-426614174000');

-- M. Delon me tuerait mais pas grave je lui ferai un pot de vin à coup de pizza à 2€ de Lidl
ALTER TABLE message DROP CONSTRAINT message_user_id_fkey;
ALTER TABLE message ALTER COLUMN user_id DROP NOT NULL;
ALTER TABLE message ADD CONSTRAINT message_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL;
ALTER TABLE provider DROP CONSTRAINT provider_user_id_fkey;
ALTER TABLE provider ALTER COLUMN user_id DROP NOT NULL;
ALTER TABLE provider ADD CONSTRAINT provider_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL;
ALTER TABLE lessor DROP CONSTRAINT lessor_user_id_fkey;
ALTER TABLE lessor ALTER COLUMN user_id DROP NOT NULL;
ALTER TABLE lessor ADD CONSTRAINT lessor_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL;
ALTER TABLE traveler DROP CONSTRAINT traveler_user_id_fkey;
ALTER TABLE traveler ALTER COLUMN user_id DROP NOT NULL;
ALTER TABLE traveler ADD CONSTRAINT traveler_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL;
ALTER TABLE administrator DROP CONSTRAINT administrator_user_id_fkey;
ALTER TABLE administrator ALTER COLUMN user_id DROP NOT NULL;
ALTER TABLE administrator ADD CONSTRAINT administrator_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL;

ALTER TABLE service ADD COLUMN lat DOUBLE PRECISION;
ALTER TABLE service ADD COLUMN lon DOUBLE PRECISION;
ALTER TABLE service ALTER COLUMN description SET NOT NULL;

ALTER TABLE reservation ADD COLUMN annulation BOOLEAN DEFAULT FALSE;
DROP TABLE property_unavailability;
DROP TABLE service_unavailability;

ALTER TABLE bill RENAME COLUMN type TO statut;
ALTER TABLE property ADD COLUMN id_stripe VARCHAR(32);
ALTER TABLE service ADD COLUMN id_stripe VARCHAR(32);
ALTER TABLE service ADD COLUMN name VARCHAR(64);

INSERT INTO Users (id, mail, password, avatar, type, description, phone_number, register_date, last_connection_date)
VALUES (
           '123e9567-e89b-12d3-a456-426214174000',
           'admin@gmail.com',
           '$2a$12$iNWU.QqnpEgZ5uUTIdqkO.Z8guqrGRHb5Nv4hD3c.Dp2XGgrESiue',
           'oui.png',
           'admin',
           'Premier administrateur !',
           '0669662851',
           CURRENT_TIMESTAMP,
           CURRENT_TIMESTAMP
);

INSERT INTO administrator (id, site, nickname, user_id)
VALUES (
           '123e4567-e89c-15d3-a456-420814174000',
           'Paris',
           'administrateur',
           '123e9567-e89b-12d3-a456-426214174000'
);

INSERT INTO chat (id, view) VALUES ('e02934d9-cb1b-475f-9972-90816d402518', FALSE);
INSERT INTO chat_user (user_id, chat_id) VALUES ('123e4567-e89b-12d3-a456-426214174000', 'e02934d9-cb1b-475f-9972-90816d402518');
INSERT INTO ticket (id, type, state, description, chat_id)
VALUES
    ('123e4567-e89b-12d3-a486-426614174001', 'paiement', 'progress', 'Problème avec le serveur', 'e02934d9-cb1b-475f-9972-90816d402518');