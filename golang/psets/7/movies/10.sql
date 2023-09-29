-- write a SQL query to list the names of all people who have directed a movie that received a rating of at least 9.0.
-- Your query should output a table with a single column for the name of each person.
-- If a person directed more than one movie that received a rating of at least 9.0, they should only appear in your results once.

SELECT COUNT(name) FROM people INNER JOIN
(SELECT DISTINCT person_id FROM directors INNER JOIN ratings ON directors.movie_id = ratings.movie_id WHERE rating >= 9.0)
ON people.id = person_id;

SELECT COUNT(name) FROM people INNER JOIN
(SELECT person_id FROM directors INNER JOIN ratings ON directors.movie_id = ratings.movie_id WHERE rating >= 9.0 GROUP BY person_id)
ON people.id = person_id;