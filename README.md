# Bot Oficial Coders Community

<p align="center">
    <img src="./assets/cc-logo.png" alt="Coders Community Logo" style="max-width:70%;">
</p>

Bem-vindo ao repositório do nosso bot! Ele foi feito totalmente utilizando a linguagem Go, e desenvolvido com contribuições principalmente da nossa staff, bem como da comunidade como um todo.

### Por quê Go?

Golang, ou Go para os mais íntimos é uma linguagem que tem crescido muito em popularidade nos últimos anos, e já está sendo utilizada amplamente no mercado. Ela é compilada, rápida, e fácil de se ler e aprender. Então, além de ser do interesse da maior parte dos desenvolvedores atualmente aprender essa fantástica linguagem, acreditamos que mesmo quem não está tão abituado a ela conseguirá contribuir tranquilamente ao nosso bot.

<p align="center">
    <img src="./assets/gopher1.png" alt="Gopher" style="max-width:100%;">
</p>

### Regras do repositório

Primeiramente, se você está considerando contribuir para o nosso bot, nós da Coders Community agradecemos imensamente! Acreditamos fortemente no desenvolvimento de aplicações FOSS, então a sua contribuição faz toda a diferença!

Agora, antes de contribuir, vamos dar uma olhada em algumas regras do nosso repositório:

#### Formatação

Para manter uma certa consistência na formatação do projeto, estamos utilizando a ferramenta [Editorconfig](https://editorconfig.org). Para mais informações, dê uma olhada no site do projeto deles.

#### Convenções

Declarações de variáveis ou funções e mensagens de commits **devem ser feitas em inglês**. Comentários no código, mensagens em PRs (Pull Requests) e Issues **devem** ser feitas em português.

Fora isso estamos tentando seguir o máximo possível o [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md).

#### Branches

Utilizamos duas branchs principais em nosso projeto: **main** e **dev**. Todas novas implementações são primeiramente feitas na branch dev, e após testadas, juntadas a branch principal.

Para efetuar um PR você deverá seguir os seguintes passos:

1. Realizar um fork do repositório.

2. Criar uma nova branch a partir da branch dev:
   
   ```
   git checkout -b my-new-feature dev
   ```

3. Adicionar seus arquivos e realizar o commit:
   
   ```
   git add .
   git commit -m "Add new feature"
   ```

4. E finalmente, realizar o push:
   
   ```
   git push origin my-new-feature
   ```

5. Depois disso, abra um PR para o merge da sua branch, com a nossa branch dev.
