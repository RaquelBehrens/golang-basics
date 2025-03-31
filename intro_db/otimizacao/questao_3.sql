SELECT *
FROM seasons
WHERE number = 1;

-- Sem índice: 0.0025 sec
-- Com íncice: 0.00097 sec
-- Adicionando o index em number de seasons, o tempo de pesquisa de seasons pelo número reduziu!
-- Escolhi criar nessa tabela, porque todas as tabelas já tinham index na primary key, e a tabela seasons era uma das com mais registros
-- Mas o critério é criar index para os atributos que são usados muito em pesquisas, e se esses registros não são atualizados com muita frequência

CREATE index number_index 
ON seasons (number); 
