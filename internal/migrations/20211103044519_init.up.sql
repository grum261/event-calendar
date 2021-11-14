CREATE TABLE tags (
    id int UNIQUE generated always AS identity,
    name varchar UNIQUE NOT NULL,
    PRIMARY KEY (id, name)
);

CREATE TABLE cities (
    id int UNIQUE generated always AS identity,
    name varchar UNIQUE NOT NULL,
    timezone int NOT NULL,
    PRIMARY KEY (id, name)
);

CREATE TABLE EVENTS (
    id int PRIMARY KEY generated always AS identity,
    name varchar NOT NULL,
    start_date timestamp without time zone NOT NULL,
    end_date timestamp without time zone NOT NULL,
    url varchar
);

CREATE TABLE events_cities (
    event_id int NOT NULL REFERENCES EVENTS ON DELETE CASCADE,
    city_id int NOT NULL REFERENCES cities (id) ON DELETE CASCADE
);

CREATE INDEX ON events_cities (event_id, city_id);

CREATE TABLE events_parts (
    id int PRIMARY KEY generated always AS identity,
    name varchar NOT NULL,
    start_time timestamp without time zone NOT NULL,
    end_time timestamp without time zone NOT NULL,
    description varchar,
    address varchar,
    place varchar NOT NULL,
    event_id int NOT NULL REFERENCES EVENTS ON DELETE CASCADE,
    age int NOT NULL DEFAULT 16
);

CREATE TABLE events_parts_tags (
    event_part_id int NOT NULL REFERENCES events_parts ON DELETE CASCADE,
    tag_id int NOT NULL REFERENCES tags (id) ON DELETE CASCADE
);

CREATE INDEX ON events_parts_tags (event_part_id, tag_id);