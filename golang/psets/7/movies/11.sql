-- write a SQL query to list the titles of the five highest rated movies (in order) that Chadwick Boseman starred in, 
--      starting with the highest rated.
-- Your query should output a table with a single column for the title of each movie.
-- You may assume that there is only one person in the database with the name Chadwick Boseman.

SELECT title, ratings.rating, year
FROM ratings INNER JOIN
(
    SELECT movies.title, movie_id AS film_id, actor, year FROM movies
    INNER JOIN
        (
            SELECT stars.movie_id, people.name AS actor FROM people 
            INNER JOIN stars 
            ON people.id = stars.person_id
            WHERE people.name = 'Chadwick Boseman'
        )
    ON movies.id = movie_id
)
ON film_id = ratings.movie_id ORDER BY rating DESC;