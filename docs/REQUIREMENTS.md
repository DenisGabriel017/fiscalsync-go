# Requisitos

## Perfis

- **Administrador:** gerencia membros e todas as conciliações.
- **Analista:** cria, processa, consulta e exporta.
- **Leitor:** apenas consulta.

## Requisitos funcionais

| ID | Requisito | Prioridade |
|---|---|---|
| RF-01 | Cadastrar usuário e organização | Alta |
| RF-02 | Autenticar e renovar sessão | Alta |
| RF-03 | Criar conciliação por competência | Alta |
| RF-04 | Enviar arquivos CSV ou XLSX | Alta |
| RF-05 | Validar tamanho, extensão e cabeçalhos | Alta |
| RF-06 | Normalizar identificadores, datas e valores | Alta |
| RF-07 | Processar em segundo plano | Alta |
| RF-08 | Classificar conciliados, ausentes e divergentes | Alta |
| RF-09 | Filtrar e paginar divergências | Alta |
| RF-10 | Exibir totais e status | Alta |
| RF-11 | Exportar o resultado em CSV | Média |
| RF-12 | Registrar auditoria | Média |

## Regras de negócio

- todo registro pertence a uma organização;
- uma conciliação possui competência e duas fontes;
- identificadores perdem pontuação e espaços na normalização;
- dinheiro é armazenado em centavos;
- igualdade usa identificador, competência e valor;
- diferença de valor gera `VALUE_MISMATCH`;
- presença em só uma fonte gera `MISSING_IN_A` ou `MISSING_IN_B`;
- arquivo inválido não inicia processamento;
- reprocessamento preserva o histórico anterior.

## Requisitos não funcionais

- arquivos de até 20 MB no MVP;
- idempotência ao iniciar processamento;
- isolamento multi-tenant no serviço e no SQL;
- logs estruturados sem conteúdo fiscal sensível;
- senhas com Argon2id ou bcrypt;
- testes de unidade do motor e testes de integração do banco;
- endpoints comuns com p95 abaixo de 500 ms, sem contar upload.

## Fora do escopo

Captura automática em sites, integração direta com sistemas contábeis, decisão tributária, assinatura digital e aplicativo móvel.
