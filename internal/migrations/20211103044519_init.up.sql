CREATE TABLE tags (
    id int unique generated always AS identity,
    name varchar unique NOT NULL,
    PRIMARY KEY (id, name)
);

CREATE TABLE cities (
    id int unique generated always AS identity,
    name varchar unique NOT NULL,
    timezone int NOT NULL,
    PRIMARY KEY (id, name)
);

CREATE TABLE EVENTS (
    id int PRIMARY KEY generated always AS identity,
    name varchar not null,
    start_date timestamp without time zone NOT NULL,
    end_date timestamp without time zone NOT NULL,
    url varchar
);

CREATE TABLE events_cities (
    event_id int NOT NULL REFERENCES EVENTS,
    city_id int NOT NULL REFERENCES cities (id)
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
    event_id int NOT NULL REFERENCES EVENTS,
    age int not null default 16
);

CREATE TABLE events_parts_tags (
    event_part_id int NOT NULL REFERENCES events_parts,
    tag_id int NOT NULL REFERENCES tags (id)
);

CREATE INDEX ON events_parts_tags (event_part_id, tag_id);