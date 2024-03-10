# Website-Monitor

## Visão Geral

O `Monitorador de Sites` é um programa escrito em Go (Golang) que realiza o monitoramento de sites, verificando se estão online e registrando logs. O programa oferece uma interface de linha de comando (CLI) para o usuário interagir e executar as seguintes funcionalidades:

1. **Iniciar Monitoramento:**
   - O usuário pode fornecer uma lista de URLs a serem monitoradas, especificar o intervalo de tempo entre as requisições e a duração total do monitoramento.
   - O programa realiza solicitações HTTP às URLs fornecidas, exibindo o status de cada site e registrando logs.

2. **Exibir Logs:**
   - O usuário pode visualizar os logs gerados durante o monitoramento, incluindo timestamps, URLs e o status de disponibilidade dos sites.

3. **Sair do Programa:**
   - Encerra a execução do programa.

## Estrutura do Código

O código está organizado em funções para melhor legibilidade e manutenibilidade. Abaixo estão as principais funções e suas responsabilidades:

- **`main`:**
  - Função principal que controla o fluxo do programa.
  - Apresenta um menu interativo para o usuário escolher as opções.

- **`exibeIntroducao`:**
  - Apresenta uma mensagem de boas-vindas e informa a versão do programa.

- **`exibeMenu`:**
  - Exibe o menu de opções disponíveis para o usuário.

- **`leComando`:**
  - Lê e retorna o comando escolhido pelo usuário.

- **`iniciarMonitoramento`:**
  - Solicita ao usuário as URLs a serem monitoradas, o intervalo de tempo entre as requisições e a duração do monitoramento.
  - Realiza o monitoramento das URLs, exibindo o status de cada site e registrando logs.

- **`obterUrls`:**
  - Solicita ao usuário inserir as URLs que deseja monitorar e as armazena em um slice.

- **`testaUrl`:**
  - Realiza uma solicitação HTTP à URL fornecida, exibindo o status do site.
  - Registra logs com timestamps e status de disponibilidade.

- **`registrarLog`:**
  - Registra logs em um arquivo chamado `log.txt` com timestamps, URLs e status de disponibilidade.

- **`imprimirLogs`:**
  - Lê e exibe o conteúdo do arquivo de logs (`log.txt`).

## Utilização

1. **Compilação e Execução:**
   - Compile o programa com o comando `go build` e execute o binário resultante.
   - Ou, execute diretamente com o comando `go run monitorador.go`.

2. **Escolha de Comandos:**
   - Use o menu interativo para escolher entre iniciar o monitoramento, exibir logs ou sair do programa.

3. **Entrada de URLs:**
   - Durante o monitoramento, insira as URLs que deseja monitorar. Digite 'sair' para encerrar a inserção.

4. **Configuração de Intervalos:**
   - Forneça o intervalo de tempo entre as requisições e a duração total do monitoramento em segundos.

5. **Logs Gerados:**
   - Os logs são registrados no arquivo `log.txt` e podem ser visualizados selecionando a opção de exibir logs no menu.

## Dependências

- O programa utiliza a biblioteca padrão `net/http` para realizar solicitações HTTP às URLs.

## Estrutura de Arquivos

- O programa cria e atualiza o arquivo `log.txt` para armazenar registros de monitoramento.

## Conclusão

O `Monitorador de Sites` é uma ferramenta simples, escrita em Go, que permite aos usuários monitorar a disponibilidade de sites e manter registros para referência futura. Sua interface de linha de comando torna fácil a execução e monitoramento em ambientes variados.
