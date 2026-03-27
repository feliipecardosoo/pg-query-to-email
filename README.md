# pg-query-to-email
Microserviço em Go que consulta dados em um BD e envia os resultados por e-mail automaticamente. (Depretada)

Alterei um pouco a infraestrutura do meu projeto em si, e inves de se conectar direto em um banco de dados, o sistema fará uma requisicao em um caminho da minha API, aonde ela ira retornar os usuarios.
Será uma rota privilegiada, aonde terei que passar um token de seguranca para conseguir dar o get.
