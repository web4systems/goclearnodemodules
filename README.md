<h3 align="center">
  Go-Clear-Node-Modules
</h3>

<p align="center">
  <a href="https://farsoft.com.br">
    <img alt="Feito para Farsoft Systems" src="https://img.shields.io/badge/made%20by-Farsoft%20Systems-purple%2306b656?style=flat-square">
  </a>

  <a href="https://www.github.com/farnetani/">
    <img alt="Feito por Arlei F. Farnetani Junior" src="https://img.shields.io/badge/solved%20by-Arlei%20F.%20Farnetani%20Junior-%2306b656?style=flat-square">
  </a>
</p>

<br>

## :rocket: Sobre o aplicativo

Criado para deletar recursivamente as pastas `node_modules` para facilitar backups de copiar e colar do windows para aplicações frontend.

<br>

## :wrench: Instalação e Uso

```bash
# Abra um terminal e copie este repositório com o comando
git clone https://github.com/web4systems/goclearnodemodules.git

# ou baixe o repositório clicando em Code > Download ZIP.

# Entre na pasta com
cd <nomedapasta>

# Na pasta exemplo-varredura
Trocar no arquivo scan_recursivo.go o diretorio.

# Listar node_modules arquivo por arquivo
go run varredura/scan_recursivo.go

# Listar os node_modules
go run listar_nodemodules.go

## Exemplo do result:
/home/farnetani/cursos/nextjs/farsoft.com.br/node_modules
/home/farnetani/cursos/nextjs/app-do-curso/node_modules
/home/farnetani/cursos/nextjs/boilerplate/node_modules
/home/farnetani/cursos/nextjs/farsoft.com.br/node_modules | 196MB | 205706320.00 | Arquivos: 28537
/home/farnetani/cursos/nextjs/app-do-curso/node_modules | 237MB | 248804606.00 | Arquivos: 36859
/home/farnetani/cursos/nextjs/boilerplate/node_modules | 241MB | 252826173.00 | Arquivos: 34861
3 diretórios: 675MB arquivos: 100257

# Para deletar usar a instrução abaixo:
go run deletar_nodemodules.go

```

<br>

## :memo: Licença

Esse projeto está sob a licença MIT. Veja o arquivo [LICENSE](/LICENSE) para mais detalhes.


Feito com :heart: por [Arlei F. Farnetani Junior](https://github.com/farnetani)

[![Github Badge](https://img.shields.io/github/followers/farnetani?style=social)](https://img.shields.io/github/followers/farnetani?style=social)
[![Instagram Badge](https://img.shields.io/badge/-farnetanijr-purple?style=flat-square&logo=Instagram&logoColor=white&link=https://www.instagram.com/farnetanijr/)](https://www.instagram.com/farnetanijr)
[![Facebook Badge](https://img.shields.io/badge/-farnetanijr-navy?style=flat-square&logo=Facebook&logoColor=white&link=https://www.facebook.com/farnetanijr/)](https://www.facebook.com/farnetanijr)
[![Gmail Badge](https://img.shields.io/badge/-farnetani@gmail.com-c14438?style=flat-square&logo=Gmail&logoColor=white&link=mailto:farnetani@gmail.com)](mailto:farnetani@gmail.com)
[![Linkedin Badge](https://img.shields.io/badge/-Arlei%20F.%20Farnetani%20Junior-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https://www.linkedin.com/in/farnetani/)](https://www.linkedin.com/in/farnetani/)

[![Github Badge](https://img.shields.io/github/followers/farnetani?label=Clique%20aqui%20para%20me%20seguir%20no%20Github&style=plastic)](https://img.shields.io/github/followers/farnetani?label=Clique%20aqui%20para%20me%20seguir%20no%20Github&style=plastic)
