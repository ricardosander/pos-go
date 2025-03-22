# pos-go
Projeto com exercícios e desafios da Pós Gradução Go Expert da Full Cycle.

Cada diretório na raiz desse projeto é um projeto go independente com um exercício ou desafio.

`busca-cep-cli` -> busca CEP via CLI. Cada argumento é um CEP a ser buscado. O resultao é escrito em um arquivo chamado `ceps.txt`

`busca-cep-api` -> servidor web REST wue recebe chamada no endpoint `:8080/busca-cep` com um query param `cep` e responde com os dados do endereço do cep informado.