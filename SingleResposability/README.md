## Single Responsibility Principle

Ele define que apenas uma função, deve ter apenas uma unica função. Que uma classe, deva ter apenas uma função de oficio. 

Ex: A classe/função, buscarPao deve ter apenas a função de buscar o pão. Não podendo adicionar um preço ao pão por exemplo.


No exemplo do código na página 'main.go', o código onde começa a apresentar as noticias, buscar a noticias de certos locais, acaba que sai do escopo inicial. Pois, mesmo estando no mesmo contexto, acaba que o mesmo sai da sua função/logica inicial da aplicação.