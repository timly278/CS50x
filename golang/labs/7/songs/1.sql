SELECT name FROM songs;
/* want to see the composer of Shape of You song */
SELECT artists.id, artists.name as artist, songs.name as song_name, songs.id as song_id
FROM artists 
INNER JOIN songs ON artists.id = songs.artist_id 
WHERE songs.name = 'Shape of You';