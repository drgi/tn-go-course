--  выборка фильмов с названием студии;
SELECT *, s.name
  FROM films
           LEFT JOIN studios s
           ON films.studio_id = s.id;

-- выборка фильмов для некоторого актёра;
SELECT *
  FROM films
 WHERE id IN
       (SELECT film_id FROM films_actors WHERE actor_id = (SELECT id FROM actors WHERE actors.name = 'Том Хэнкс'));

-- подсчёт фильмов для некоторого режиссёра;
SELECT COUNT(*)
  FROM films
           INNER JOIN films_directors fd
           ON films.id = fd.film_id
           INNER JOIN directors d
           ON d.id = fd.director_id
 WHERE d.name = 'Роберт Земекис';

-- выборка фильмов для нескольких режиссёров из списка (подзапрос);
SELECT *
  FROM films
           INNER JOIN films_directors fd
           ON films.id = fd.film_id
 WHERE director_id IN (SELECT id FROM directors WHERE birth_date > '1953-01-01');

-- подсчёт количества фильмов для актёра;
SELECT a.name, COUNT(*)
  FROM films
           INNER JOIN films_actors fa
           ON films.id = fa.film_id
           INNER JOIN actors a
           ON fa.actor_id = a.id
 GROUP BY fa.actor_id, a.name;


-- выборка актёров и режиссёров, участвовавших более чем в 2 фильмах;
SELECT *
  FROM (SELECT *
          FROM actors
         WHERE id IN (SELECT actor_id FROM films_actors GROUP BY actor_id HAVING COUNT(film_id) > 1)
         UNION
        SELECT *
          FROM directors
         WHERE id IN (SELECT director_id FROM films_directors GROUP BY director_id HAVING COUNT(film_id) > 1)) AS u;


-- подсчёт количества фильмов со сборами больше 1000;
SELECT COUNT(*)
  FROM films
 WHERE profit > 1000;

-- подсчитать количество режиссёров, фильмы которых собрали больше 1000;
SELECT COUNT(DISTINCT directors.id)
  FROM directors
           LEFT JOIN films_directors fd
           ON directors.id = fd.director_id
           LEFT JOIN films f
           ON fd.film_id = f.id
 WHERE f.profit > 1000
;

-- выборка различных фамилий актёров;
SELECT DISTINCT SPLIT_PART(name, ' ', 2)
  FROM actors;

-- подсчёт количества фильмов, имеющих дубли по названию.
SELECT COUNT(*)
  FROM (SELECT COUNT(*) FROM films GROUP BY name) AS count
 WHERE count.count > 1;
