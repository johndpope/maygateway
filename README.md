![diagram](/../master/maygateway-diagram.jpg?raw=true)

# maygateway
Repositorio para trabalho de sistemas distribuidos

## Formalização
O projeto consiste em um gateway de microservicos, onde as API's serao cadastradas atraves da UI e serao salvas 
em um banco postgres passando por microservicos do Manager. As API's cadastradas serao sincrozadas com o redis que sera
utilizado para leitura do gateway ao receber uma requisição. Alem disso todas as requisicões serão publicadas em uma fila
e salvas no elasticsearch para geraçao de logs, estes serão exibidos na UI.

## Componentes
- Gateway
- Manager
- UI
- Postgres
- Redis
- Elasticsearch
- Fila (RabbitMQ, afins...)
