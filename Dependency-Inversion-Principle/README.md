## Dependency Inversion Principle

Dependa de abstrações e não de implementações. 

1. Módulo de alto nivel não devem depender de módulos de baixo nivel. 
Mabos devem depender da abstração. 

2. Abstrações não devem depender de detalhes.
Detalhes devem depender de abstrações.

A melhor opção em um cenario de CRUD, onde o nosso serviço precisa trabalhar com a persistencia em um banco de dados por exemplo, seria interessante que funciona-se na seguinte ordem de chama:

Services -> IntefaceRepository -> implementação do repositorio

Services chamando a interface de repository, independentemente do que aconteça da interface para trás, o services não muda e para é como se nada tivesse acontecido. 