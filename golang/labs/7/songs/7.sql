-- write a SQL query that returns the average energy of songs that are by Drake.
-- output a table with a single column and a single row containing the average energy.
-- should not make any assumptions about what Drake’s artist_id is

SELECT AVG(energy) AS Drake_AverageEnergy FROM
(SELECT songs.energy FROM songs INNER JOIN artists ON artists.id = songs.artist_id WHERE artists.name = 'Drake');