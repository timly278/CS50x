-- write a SQL query to list the titles of all movies in which both Bradley Cooper and Jennifer Lawrence starred.
-- Your query should output a table with a single column for the title of each movie.
-- You may assume that there is only one person in the database with the name Bradley Cooper.
-- You may assume that there is only one person in the database with the name Jennifer Lawrence.

SELECT title, rating, year FROM 
(
    SELECT title, movie_id, year FROM movies INNER JOIN
    (
        SELECT m1 as movie_id 
        FROM 
        (SELECT movie_id AS m1, person_id AS p1 FROM stars INNER JOIN people ON stars.person_id = people.id WHERE people.name = 'Bradley Cooper'),
        (SELECT movie_id AS m2, person_id AS p2 FROM stars INNER JOIN people ON stars.person_id = people.id WHERE people.name = 'Jennifer Lawrence')
        WHERE m1 = m2
    )
    ON movie_id = movies.id
) AS newTable
INNER JOIN ratings ON ratings.movie_id = newTable.movie_id;

/* test other way */
SELECT title FROM movies INNER JOIN
(
    SELECT m1 as movie_id FROM 
    (SELECT movie_id AS m1, person_id AS p1 FROM stars INNER JOIN people ON stars.person_id = people.id WHERE people.name = 'Bradley Cooper')
    INNER JOIN
    (SELECT movie_id AS m2, person_id AS p2 FROM stars INNER JOIN people ON stars.person_id = people.id WHERE people.name = 'Jennifer Lawrence')
    ON m1 = m2
)
ON movie_id = movies.id;

-- Bradley Cooper if Jennifer Lawrence movies.
+-------------------------+--------+------+
|          title          | rating | year |
+-------------------------+--------+------+
| Silver Linings Playbook | 7.7    | 2012 |
| Serena                  | 5.4    | 2014 |
| American Hustle         | 7.2    | 2013 |
| Joy                     | 6.6    | 2015 |
+-------------------------+--------+------+