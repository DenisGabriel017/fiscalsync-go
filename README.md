# FiscalSync Go

API em Go para importar relatórios de duas fontes, normalizar registros e identificar divergências automaticamente.

> Projeto de portfólio. Use somente dados fictícios ou anonimizados.

## Problema

Equipes precisam comparar manualmente relatórios exportados por portais e sistemas internos. Diferenças de formato, nomes de colunas e valores tornam o processo lento e sujeito a erros. O FiscalSync recebe dois arquivos, transforma os registros em um modelo comum e classifica correspondências, ausências e divergências.

## MVP

- autenticação JWT e perfis de acesso;
- organizações isoladas por `tenant_id`;
- upload de CSV e XLSX;
- validação e normalização dos registros;
- processamento assíncrono com worker;
- conciliação por identificador, competência e valor;
- consulta paginada e exportação CSV;
- logs de auditoria.

## Stack sugerida

Go 1.24+, `chi`, PostgreSQL, `pgx`, `sqlc`, Goose, JWT, Docker Compose, OpenAPI 3.1 e Testcontainers.

## Fluxo

1. Usuário cria uma conciliação por competência.
2. Envia o arquivo da fonte A e da fonte B.
3. A API valida tipo, tamanho e cabeçalhos.
4. O worker normaliza os registros.
5. O motor compara as fontes.
6. A API apresenta o resumo e as divergências.
7. O usuário exporta o resultado.

## Estrutura planejada

```text
cmd/api
cmd/worker
internal/auth
internal/organization
internal/importfile
internal/reconciliation
internal/audit
internal/platform
migrations
docs
testdata
```

## Documentação

- [Requisitos](docs/REQUIREMENTS.md)
- [Arquitetura](docs/ARCHITECTURE.md)
- [Modelo de dados](docs/DATA_MODEL.md)
- [Contrato da API](docs/API.md)
- [Segurança](docs/SECURITY.md)
- [Roadmap](docs/ROADMAP.md)

## MVP concluído quando

- `docker compose up` inicia API e banco;
- migrations funcionam do zero;
- autenticação, upload, processamento e consulta estão integrados;
- dados de organizações diferentes nunca se misturam;
- casos críticos possuem testes;
- OpenAPI está acessível;
- CI executa lint, testes e build.

## Limites

O sistema não acessa portais governamentais, não substitui análise profissional e não toma decisões fiscais. Ele processa somente os arquivos fornecidos pelo usuário.

## Licença

MIT.
