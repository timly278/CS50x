-- write a SQL query that lists the names of songs that are by Post Malone
-- output a table with a single column for the name of each song.

SELECT songs.name AS NameSong, artists.id FROM songs INNER JOIN artists ON artists.id = songs.artist_id WHERE artists.name = 'Post Malone';