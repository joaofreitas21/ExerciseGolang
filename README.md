#Exercicio golang

Objetivos:
- Criar a base de dados inicial 
- Login - POST /api/login
- Operações - GET /api/persons/{id}
            - POST /api/create


TODO:
- Verificar se é um admin a realizar as operações com utilizadores
    --> Extrair token, e através do ID que contém verificar a "Role"
- Se falhar login 3x em menos de 10s, impedir o login durante 30s (Go routines)
- Alterar a biblioteca JWT utilizada


Alguma Documentação utilizada:
- https://www.youtube.com/watch?v=YA6cVebkwJE&ab_channel=EricLau (arquitetura e organização/jwt)
- https://www.youtube.com/watch?v=hWmR8YtlFlE&ab_channel=DailyCodeBuffer (jwt)


How to Use:
- cd src/main ->> go install 
- Execute .\main.exe in bin directory