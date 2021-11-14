CREATE TABLE events_positions (
    event_id int NOT NULL REFERENCES EVENTS,
    event_date timestamp without time zone NOT NULL,
    position int NOT NULL
);

CREATE INDEX ON events_positions (event_id, event_date);