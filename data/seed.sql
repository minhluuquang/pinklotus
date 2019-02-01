
CREATE TABLE IF NOT EXISTS vessels (
        id varchar(36) PRIMARY KEY NOT NULL,
        name VARCHAR(512) NOT NULL,
        capacity INT NOT NULL,
        max_weight INT NOT NULL,
        created_at timestamp DEFAULT now() NOT NULL,
        updated_at timestamp DEFAULT now() NOT NULL,
        deleted_at timestamp
);

CREATE TABLE IF NOT EXISTS consignments (
        id varchar(36) PRIMARY KEY NOT NULL,
        vessel_id VARCHAR(36)  REFERENCES vessels (id) ON DELETE CASCADE,
        containers_number INT NOT NULL,
        weight INT NOT NULL,
        description varchar(512),
        created_at timestamp DEFAULT now() NOT NULL,
        updated_at timestamp DEFAULT now() NOT NULL,
        deleted_at timestamp
);

INSERT INTO consignments (id, containers_number, weight, description) VALUES 
('e01d17a3-8f26-4af9-ba69-015a0728d374', 5, 50, 'consignment 1'),
('5336db7f-0abe-42e9-b9ff-cb95cc0fe58b', 7, 15, 'consignment 2'),
('13082d81-95e0-447b-a9ba-93277f1fe8d8', 10, 70, 'consignment 3');

INSERT INTO vessels (id, name, capacity, max_weight) VALUES 
('aeb5b577-fd12-4af0-b791-69ecd8a53226', 'vessel 1', 10, 30),
('2635e49c-85a2-431e-b01c-d16eafd797fb', 'vessel 2', 6, 17),
('6bf93edf-0690-4a97-b90a-29e092189622', 'vessel 3', 12, 100);