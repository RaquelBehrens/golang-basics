-- Listar todos os clientes
SELECT * FROM cliente;

-- Listar todos os planos
SELECT * FROM plano;

-- Listar clientes e seus respectivos planos (join)
SELECT c.nome, c.sobrenome, p.velocidade, p.preco
FROM cliente c
JOIN plano p ON c.planoId = p.identificacao;

-- Clientes que moram em um determinado estado
SELECT nome, sobrenome FROM cliente WHERE estado = 'Estado X';

-- Planos com velocidade acima de um certo valor
SELECT identificacao, velocidade, preco FROM plano WHERE velocidade > 400;

-- Média do preço dos planos
SELECT AVG(preco) AS media_preco_planos FROM plano;

-- Clientes que nasceram após uma determinada data
SELECT nome, sobrenome FROM cliente WHERE data_nasc > '1990-01-01';

-- Contagem de clientes por plano
SELECT planoId, COUNT(*) AS total_clientes FROM cliente GROUP BY planoId;

-- Detalhes do plano mais caro
SELECT * FROM plano ORDER BY preco DESC LIMIT 1;

-- Clientes que têm um plano com desconto
SELECT c.nome, c.sobrenome
FROM cliente c
JOIN plano p ON c.planoId = p.identificacao
WHERE p.desconto > 0;