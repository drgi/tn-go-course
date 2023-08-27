CREATE TABLE studios
    (
        id SERIAL
            PRIMARY KEY,
        name varchar(255)
    );

CREATE TABLE actors
    (
        id SERIAL
            PRIMARY KEY,
        name varchar(255),
        birth_date DATE NOT NULL
    );

CREATE TABLE directors
    (
        id SERIAL
            PRIMARY KEY,
        name varchar(255),
        birth_date DATE NOT NULL
    );

CREATE TABLE films
    (
        id SERIAL
            PRIMARY KEY,
        name varchar(255) NOT NULL,
        realise_year INTEGER NOT NULL
            CHECK ( realise_year >= 1800  ),
        profit numeric NOT NULL,
        rating varchar(255)
            CHECK ( rating IN ('PG-10', 'PG-13', 'PG-18') ),
        studio_id bigint,
        CONSTRAINT fk_films_studio
            FOREIGN KEY (studio_id) REFERENCES studios (id)
    );

CREATE TABLE films_actors
    (
        film_id bigint NOT NULL,
        actor_id bigint NOT NULL,
        CONSTRAINT fk_films
            FOREIGN KEY (film_id) REFERENCES films (id),
        CONSTRAINT fk_actors
            FOREIGN KEY (actor_id) REFERENCES actors (id)
    );

CREATE TABLE films_directors
    (
        film_id bigint NOT NULL,
        director_id bigint NOT NULL,
        CONSTRAINT fk_films
            FOREIGN KEY (film_id) REFERENCES films (id),
        CONSTRAINT fk_actors
            FOREIGN KEY (director_id) REFERENCES directors (id)
    );

CREATE UNIQUE INDEX idx_films_name_year_unique ON films (realise_year, name);
CREATE UNIQUE INDEX idx_films_actors ON films_actors (actor_id, film_id);
CREATE UNIQUE INDEX idx_films_directors ON films_directors (director_id, film_id);
