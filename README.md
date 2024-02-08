
# Desafio Tunts.Rocks 2024

Parte integrante do processo seletivo da Tunts.Rocks, o desafio de programação tem como objetivo  principal a avaliação das habilidades de programação do candidato. Levando em conta não  apenas o êxito de implementação da funcionalidade desejada, mas também uma análise da  solução de forma estrutural, semântica e performática. 

### REGRAS: 

#### Link da tabela do Google Sheets
https://docs.google.com/spreadsheets/d/1jJCYe__gprZznPm-iecd4x5p0W9oXxXyhlHOqhF7IAY/edit?usp=sharing



Calcular a situação de cada aluno baseado na média das 3 provas (P1, P2 e P3), conforme a  tabela: 


Média (m) Situação:

m<5  - Reprovado por Nota

5<=m<7  - Exame Final

m>=7  - Aprovado

Caso o número de faltas ultrapasse 25% do número total de aulas o aluno terá a situação  "Reprovado por Falta", independente da média.  Caso a situação seja "Exame Final" é necessário calcular a "Nota para Aprovação Final"(naf) de  cada aluno de acordo com seguinte fórmula: 

5 <= (m + naf)/2

Caso a situação do aluno seja diferente de "Exame Final", preencha o campo "Nota para  Aprovação Final" com 0. 

Arredondar o resultado para o próximo número inteiro (aumentar) caso necessário. Utilizar linhas de logs para acompanhamento das atividades da aplicação. 


## Instalação

A aplicação foi feita em Go e "dockerizada" para facilitar a execução e evitar erros

```bash
  git clone git@github.com:markallenarchviz/desafioTuntsGo.git
  cd desafioTuntsGo
```
É preciso ter o Docker previamente instalado para rodar a aplicação

```bash
  docker build -t desafio-tunts .
```
Após o docker terminar de buildar a imagem da aplicação use o comando
```bash
  docker run desafio-tunts
```
Isso irá iniciar o contêiner e executar a aplicação, modificando a planilha do Google Sheets, sempre que o comando de iniciar contêiner for executado, a aplicação irá calcular as notas na planilha

Caso queira rodar a aplicação localmente, tenha instalado o Go na versão 1.21.6 e use o seguinte comando(certifique-se de estar dentro da pasta principal do projeto):
```bash
  go run cmd/main.go
```
