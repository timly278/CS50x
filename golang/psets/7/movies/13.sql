-- write a SQL query to list the names of all people who starred in a movie in which Kevin Bacon also starred.
-- Your query should output a table with a single column for the name of each person.
-- There may be multiple people named Kevin Bacon in the database. Be sure to only select the Kevin Bacon born in 1958.
-- Kevin Bacon himself should not be included in the resulting list


SELECT COUNT(name) FROM people INNER JOIN
(    
    SELECT stars.person_id AS starsID FROM stars,
    (
        SELECT movie_id AS kevin_movie, person_id AS kevin_id FROM stars INNER JOIN people ON stars.person_id = people.id 
        WHERE people.name = 'Kevin Bacon' AND people.birth = 1958
    )
    WHERE (stars.movie_id = kevin_movie) AND (stars.person_id <> kevin_id) GROUP BY starsID
)
ON people.id = starsID;

/* not using INNER JOIN */
SELECT COUNT(name) FROM people,
(    
    SELECT stars.person_id AS starsID FROM stars,
    (
        SELECT movie_id AS kevin_movie, person_id AS kevin_id FROM stars INNER JOIN people ON stars.person_id = people.id 
        WHERE people.name = 'Kevin Bacon' AND people.birth = 1958
    )
    WHERE (stars.movie_id = kevin_movie) AND (stars.person_id <> kevin_id) GROUP BY starsID
)
WHERE people.id = starsID;
