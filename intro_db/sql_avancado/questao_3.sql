-- Enumera os dados dos autores.
SELECT *
FROM autor

-- Indica o nome e a idade dos alunos
SELECT nombre, edad
FROM estudiante

-- Que alunos pertencem ao curso de informática?
SELECT nombre
FROM estudiante
WHERE carrera = "Informática"

-- Quais são os autores de nacionalidade francesa ou italiana?
SELECT nombre
FROM autor
WHERE nacionalidade IN ("francesa", "italiana");

-- Quais os livros que não são da área da Internet?
SELECT titulo
FROM libro
WHERE area != "Internet"

-- Enumera os livros publicados pela Salamandra.
SELECT titulo
FROM libro
WHERE editorial = "Salamandra"

-- Enumera os nomes dos alunos cuja idade é superior à média.
SELECT nombre
FROM estudiante
WHERE edad > (SELECT AVG(edad) FROM estudiante)

-- Enumera os nomes dos alunos cujo apelido começa com a letra G.
SELECT nombre
FROM estudiante
WHERE apellido lIKE "G%"

-- Faz uma lista dos autores do livro "O Universo: Guia de Viagem". (Apenas os nomes devem ser indicados).
SELECT a.nombre
FROM autor a
INNER JOIN libroautor as la ON a.idAutor = la.idAutor
INNER JOIN livro as l ON l.idLibro = la.idLibro
WHERE titulo = "O Universo: Guia de Viagem"

-- Que livros emprestaste ao leitor "Filippo Galli"?
SELECT l.titulo
FROM estudiante e
INNER JOIN prestamo p ON p.idLector = e.idLector
INNER JOIN livro l ON l.idLibro = p.idLibro
WHERE e.nombre = "Filippo Galli"

-- Indica o nome do aluno mais novo.
SELECT nombre
FROM estudiante
WHERE edad = (SELECT MIN(edad) FROM estudiante)

-- Enumera os nomes dos alunos a quem foram emprestados livros da Base de Dados.
SELECT estudiante.nombre
FROM estudiante
INNER JOIN prestamo ON estudiante.idLector = prestamo.idLector

-- Enumera os livros que pertencem à autora J.K. Rowling.
SELECT l.titulo
FROM libro l
INNER JOIN libroautor la ON la.idLibro = l.idLibro
INNER JOIN autor a ON la.idAutor = a.idAutor
WHERE a.nombre = "J. K. Rowling"

-- Enumera os títulos dos livros que deviam ser devolvidos em 16/07/2021.
SELECT l.titulo
FROM libro l
INNER JOIN prestamo p ON p.idLibro = l.idLibro
WHERE p.fechaDevolucion <= "2021-07-16"
AND p.devuelto = 0
