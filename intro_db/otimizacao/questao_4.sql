-- Adicionar um filme à tabela de movies.
INSERT INTO movies (created_at, updated_at, title, rating, awards, release_date, length, genre_id)
VALUES ("2014-11-06", "2014-11-06", "Interestelar", 9.2, 6, "2014-11-06", 169, 5);

-- Adicione um gênero à tabela de genres.
INSERT INTO genres (created_at, name, ranking, active)
VALUES ("2014-11-06", "Romance", 13, 1);

-- Associe o gênero criado no ponto 2 ao filme no ponto 1. gênero.
UPDATE movies
SET genre_id = 13
where id = 22

-- Modifique a tabela de atores para que pelo menos um ator tenha como favorito o filme adicionado no ponto 1.
UPDATE actors
SET favorite_movie_id = 22
WHERE id = 47

-- Crie uma cópia temporária da tabela de movies.
CREATE TEMPORARY TABLE copia AS
SELECT * FROM movies;
SELECT * FROM copia;

-- Remova dessa tabela temporária todos os filmes que ganharam menos de 5 awards.
SET SQL_SAFE_UPDATES = 0;
DELETE FROM copia
WHERE awards < 5;
SELECT * FROM copia;

-- Obtenha a lista de todos os gêneros que têm pelo menos um movies.
SELECT *
FROM genres
INNER JOIN movies ON genres.id = movies.genre_id

-- Obtenha a lista de atores cujo filme favorito ganhou mais de 3 awards.
SELECT *
FROM actors
INNER JOIN movies ON actors.favorite_movie_id = movies.id
WHERE movies.awards > 3

-- Crie um índice sobre o nome na tabela de movies.
CREATE index title_index 
ON movies (title); 

-- Verifique se o índice foi criado corretamente.]
SHOW INDEX FROM movies;

-- No banco de dados de movies, você notaria uma melhora significativa com a criação de índices? Analise e justifique sua resposta.
-- Não muito, porque não tem muitos registros nessa tabela.
 
-- Em qual outra tabela você criaria um índice e por quê? Justifique sua resposta.
-- Nenhuma tabela tem muitos registros nesse banco de dados.
