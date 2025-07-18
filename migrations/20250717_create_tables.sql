BEGIN;

CREATE TABLE actors (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    gender      VARCHAR(10)  NOT NULL,
    birth_date  DATE         NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE TABLE movies (
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(150) NOT NULL CHECK (char_length(title) BETWEEN 1 AND 150),
    description VARCHAR(1000),
    release_date DATE,
    rating      NUMERIC(3,1) CHECK (rating BETWEEN 0 AND 10),
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE TABLE movie_actors (
    movie_id INTEGER NOT NULL REFERENCES movies(id) ON DELETE CASCADE,
    actor_id INTEGER NOT NULL REFERENCES actors(id) ON DELETE CASCADE,
    PRIMARY KEY (movie_id, actor_id)
);

CREATE INDEX idx_movies_title_lower   ON movies (lower(title));
CREATE INDEX idx_movies_rating        ON movies (rating);
CREATE INDEX idx_movies_release_date  ON movies (release_date);
CREATE INDEX idx_actors_name_lower    ON actors (lower(name));

COMMIT;
