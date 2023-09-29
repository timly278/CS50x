-- write a SQL query to list the names of all people who starred in a movie released in 2004, ordered by birth year.
-- Your query should output a table with a single column for the name of each person.
-- People with the same birth year may be listed in any order.
-- No need to worry about people who have no birth year listed, so long as those who do have a birth year are listed in order.
-- If a person appeared in more than one movie in 2004, they should only appear in your results once.

/* test if any person has no birth year */
SELECT * FROM people WHERE birth IS NULL;

SELECT name, birth FROM people INNER JOIN 
(SELECT DISTINCT person_id FROM movies INNER JOIN stars ON movies.id = stars.movie_id WHERE movies.year = 2004)
ON people.id = person_id ORDER BY birth DESC;

/* Try another way using order by */
SELECT count(name) AS num, name, person_id FROM people INNER JOIN
(SELECT person_id FROM movies INNER JOIN stars ON movies.id = stars.movie_id WHERE movies.year = 2004)
ON people.id = person_id GROUP BY person_id ORDER BY num DESC LIMIT 40;