# wallet
Projet de Wallet


## Arquitetura Hexagonal

A arquitetura que será adota, será a hexagonal, pois a arquitetura hexagonal é 
um padrão de design que visa criar aplicações com um núcleo de negócios bem definido 
e isolado de tecnologias externas, como bancos de dados, frameworks e interfaces de usuário. 
Essa abordagem promove alta testabilidade, reutilização de código e facilita a evolução do sistema ao longo do tempo.

### Por que usar arquitetura hexagonal em Golang?

<strong>Isolamento da lógica de negócio:</strong> O núcleo da aplicação, onde reside a regra de negócio, fica completamente 
isolado de detalhes de implementação. Isso facilita a escrita de testes unitários e torna o código mais fácil de entender e manter.  
<strong>Testabilidade:</strong> A separação de responsabilidades permite testar cada componente de forma isolada, 
garantindo a qualidade do código.  
<strong>Flexibilidade:</strong> A arquitetura hexagonal permite trocar facilmente tecnologias e frameworks sem afetar 
o núcleo da aplicação.  
<strong>Reutilização de código:</strong> A lógica de negócio pode ser reutilizada em diferentes contextos, pois 
está desacoplada de detalhes de implementação.  

### Conceitos chave:

<strong>Núcleo:</strong> Contém a lógica de negócio da aplicação, as entidades e as regras de negócio. É completamente 
independente de qualquer tecnologia externa.  
<strong>Portas:</strong> São interfaces que definem como o núcleo se comunica com o mundo externo.  
<strong>Adaptadores:</strong> Implementam as portas, conectando o núcleo a tecnologias específicas, como bancos de dados, frameworks web, etc.  
<strong>Atores:</strong> São os elementos externos que interagem com a aplicação, como usuários, sistemas legados, etc.  

### Benefícios de usar arquitetura hexagonal em Golang:

<strong>Código mais limpo e organizado:</strong> A separação de responsabilidades torna o código mais fácil de entender e manter.  
<strong>Melhora a qualidade do código:</strong> A alta testabilidade permite identificar e corrigir bugs mais rapidamente.  
<strong>Facilita a evolução do sistema:</strong> A arquitetura hexagonal permite adicionar novas funcionalidades e trocar tecnologias
 sem afetar o núcleo da aplicação.  
<strong>Promove a colaboração entre equipes:</strong> A arquitetura hexagonal facilita a colaboração entre equipes, pois cada 
equipe pode se concentrar em uma área específica da aplicação.  

## Estrutura do Projeto

![Estrutura](imagens/estrutura_projeto.png)

Os projetos terão como estrutura o modelo apresentando acima.

<strong>internal:</strong> Contém todo o código interno da aplicação.  
<strong>domain:</strong> O núcleo da aplicação, com entidades, portas e casos de uso.  
<strong>entities:</strong> Representam os conceitos do domínio, como User, Product, etc.  
<strong>ports:</strong> Definem as interfaces que o domínio expõe.  
<strong>usecases:</strong> Encapsulam a lógica de negócio, utilizando as portas para interagir com o mundo externo.  
<strong>adapters:</strong> Implementam as portas, conectando o domínio a tecnologias específicas.  
<strong>input:</strong> Adaptadores de entrada, como HTTP handlers.  
<strong>output:</strong> Adaptadores de saída, como repositórios de banco de dados.  
<strong>config:</strong> Contém configurações da aplicação.  
<strong>cmd:</strong> Contém os comandos da aplicação, como o servidor HTTP.  
<strong>main.go:</strong> Ponto de entrada da aplicação.  