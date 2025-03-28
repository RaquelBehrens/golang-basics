-- Selecciona o nome, a posição e a localização dos departamentos onde os vendedores trabalham.
SELECT f.nome, f.sobrenome, f.posto, d.localidad
FROM funcionario AS f
INNER JOIN departamento as d
ON f.dpto_nro = d.depto_nro

-- Mostra os departamentos com mais de cinco empregados.
SELECT COUNT(*), d.nombre_depto
FROM funcionario AS f
INNER JOIN departamento as d
ON f.dpto_nro = d.dpto_nro
GROUP BY d.dpto_nro

-- Mostra o nome, o salário e o nome do departamento dos empregados que têm a mesma posição que "Mito Barchuk".
SELECT f2.nome, f2.sobrenome, f2.salario, d.nombre_depto
FROM funcionario f2
INNER JOIN departamento as d
ON f2.dpto_nro = d.dpto_nro
WHERE f2.posto = (SELECT posto FROM funcionario AS f WHERE f.nome = "Mito" AND f.sobrenome = "Barchuk")

-- Mostra os detalhes dos empregados que trabalham no departamento de contabilidade, ordenados por nome.
SELECT f.*
FROM departamento as d
INNER JOIN funcionario as f
ON f.depto_nro = d.depto_nro
WHERE d.nombre_depto = "Contabilidade"
ORDER BY f.nome 

-- Mostra o nome do empregado com o salário mais baixo.
SELECT f.nome, f.sobrenome
FROM funcionario as f
WHERE f.salario = (SELECT MIN(salario) FROM funcionario);

-- Mostra os detalhes do empregado com o salário mais alto no departamento de "Vendas".
SELECT f.*
FROM funcionario as f
INNER JOIN departamento as d
WHERE d.nombre_depto = "Vendas" AND f.salario = (SELECT MAX(salario) from funcionario)
