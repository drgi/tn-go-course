INSERT INTO studios (id, name)
VALUES (1, 'Castle Rock Entertainment'),
       (2, 'Darkwoods Productions'),
       (3, 'Warner Bros.');

ALTER SEQUENCE studios_id_seq RESTART WITH 4;

INSERT INTO actors (id, name, birth_date)
VALUES (1, 'Том Хэнкс', '1956-07-09'),
       (2, 'Тим Роббинс', '1958-10-16'),
       (3, 'Морган Фриман', '1937-06-01'),
       (4, 'Боб Гантон', '1945-11-15');

ALTER SEQUENCE actors_id_seq RESTART WITH 5;

INSERT INTO directors (id, name, birth_date)
VALUES (1, 'Фрэнк Дарабонт', '1959-01-28'),
       (2, 'Роберт Земекис', '1952-05-14');

ALTER SEQUENCE directors_id_seq RESTART WITH 3;

INSERT INTO films (id, name, realise_year, profit, rating, studio_id)
VALUES (1, 'Побег из Шоушенка', 1994, 28418687, 'PG-18', 1),
       (2, 'Зеленая миля', 1999, 286801374, 'PG-10', 2),
       (3, 'Форрест Гамп', 1994, 286801374, 'PG-13', 3);

ALTER SEQUENCE films_id_seq RESTART WITH 4;

INSERT INTO films_actors (film_id, actor_id)
VALUES (1, 2),
       (1, 3),
       (2, 1);


INSERT INTO films_directors (film_id, director_id)
VALUES (1, 1),
       (2, 1),
       (3, 2);