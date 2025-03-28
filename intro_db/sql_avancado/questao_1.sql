-- Exibir o título e o nome do gênero de todas as séries.
SELECT s.title, g.name
FROM series s INNER JOIN genres g 
WHERE s.genre_id = g.id

-- Mostre o título dos episódios, o nome e o sobrenome dos atores que trabalham em cada episódio.
SELECT e.title, ae.episode_id, ae.actor_id
FROM episodes AS e 
INNER JOIN actor_episode AS ae ON e.id = ae.episode_id 
INNER JOIN actors AS a ON ae.actor_id = a.id

-- Mostre o título de todas as séries e o número total de temporadas de cada série.
SELECT COUNT(*), series.title
FROM series
LEFT JOIN seasons
ON series.id = seasons.serie_id
GROUP BY series.id

-- Mostre o nome de todos os gêneros e o número total de filmes de cada gênero, desde que seja maior ou igual a 3.
SELECT COUNT(movies.genre_id) AS amount_movies, genres.id as genre_id, genres.name
FROM genres
LEFT JOIN movies
ON genres.id = movies.genre_id
GROUP BY genres.id
HAVING amount_movies >= 3

-- Mostre apenas o nome e o sobrenome dos atores que trabalharam em todos os filmes de Guerra nas Estrelas e não os repita.
SELECT a.first_name, a.last_name
FROM movies AS m
INNER JOIN actor_movie AS am ON m.id = am.movie_id
INNER JOIN actors AS a ON a.id = am.actor_id
WHERE m.title LIKE 'La Guerra de las Galaxias%'
GROUP BY a.first_name, a.last_name
HAVING COUNT(DISTINCT m.id) = (SELECT COUNT(*) FROM movies WHERE title LIKE 'La Guerra de las Galaxias%')
