CREATE TABLE urls
    (
        id SERIAL
            PRIMARY KEY,
        url text,
        string_id varchar(255)
    );

CREATE UNIQUE INDEX idx_urls_string_id ON urls (string_id);
