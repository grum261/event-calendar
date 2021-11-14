CREATE TABLE cities (
    id int UNIQUE generated always AS identity,
    name varchar UNIQUE NOT NULL,
    timezone int NOT NULL,
    PRIMARY KEY (id, name)
);

CREATE TABLE events_cities (
    event_id int NOT NULL REFERENCES EVENTS ON DELETE CASCADE,
    city_id int NOT NULL REFERENCES cities (id) ON DELETE CASCADE
);

CREATE INDEX ON events_cities (event_id, city_id);