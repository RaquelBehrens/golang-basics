-- CREATE TEMPORARY TABLE twd (
--     id int,
--     id_episode int
-- );

DROP TABLE IF EXISTS twd;

CREATE TEMPORARY TABLE twd AS
SELECT e.id, e.title, s.number
FROM episodes AS e
INNER JOIN seasons AS s ON e.season_id = s.id
INNER JOIN series ON series.id = s.serie_id
WHERE series.title = "The Walking Dead";

SELECT id, title, number
FROM twd
WHERE number = 1;