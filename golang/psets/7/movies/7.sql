-- write a SQL query to list all movies released in 2010 and their ratings, in descending order by rating. 
--      For movies with the same rating, order them alphabetically by title.
-- Your query should output a table with two columns, one for the title of each movie and one for the rating of each movie.
-- Movies that do not have ratings should not be included in the result.

SELECT movies.title, ratings.rating FROM movies INNER JOIN ratings ON movies.id = ratings.movie_id
WHERE movies.year = 2010
ORDER BY ratings.rating DESC, movies.title ASC;


/* Find the best movies of favorite actors */
SELECT title, ratings.rating, year
FROM ratings INNER JOIN
(
    SELECT movies.title, movie_id AS film_id, actor, year FROM movies
    INNER JOIN
        (
            SELECT stars.movie_id, people.name AS actor FROM people 
            INNER JOIN stars 
            ON people.id = stars.person_id
            WHERE people.name = 'Anne Hathaway'
        )
    ON movies.id = movie_id
)
ON film_id = ratings.movie_id ORDER BY rating DESC;