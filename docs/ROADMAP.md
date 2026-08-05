# Roadmap

Construa em etapas pequenas. Não implemente tudo de uma vez.

## Fase 0 — Fundação

- [ ] módulo Go e configuração;
- [ ] logger e servidor com shutdown gracioso;
- [ ] Docker Compose com PostgreSQL;
- [ ] migrations e CI.

**Pronto quando:** build e testes passam localmente e no GitHub Actions.

## Fase 1 — Identidade

- [ ] usuários e organizações;
- [ ] cadastro, login, refresh e logout;
- [ ] middleware, RBAC e isolamento multi-tenant;
- [ ] testes de integração.

**Pronto quando:** dois tenants não acessam recursos um do outro.

## Fase 2 — Upload e parsing

- [ ] criação de conciliação;
- [ ] upload e parser CSV;
- [ ] validação de cabeçalhos;
- [ ] normalização de ID, data e dinheiro;
- [ ] suporte XLSX;
- [ ] arquivos sintéticos de exemplo.

**Pronto quando:** arquivos válidos geram registros canônicos e inválidos retornam erros claros.

## Fase 3 — Motor

- [ ] tabela e consumidor de jobs;
- [ ] comparação por ID, competência e valor;
- [ ] divergências e contadores;
- [ ] reprocessamento e idempotência;
- [ ] testes do conjunto de referência.

## Fase 4 — Consulta

- [ ] paginação e filtros;
- [ ] resumo;
- [ ] exportação CSV em streaming;
- [ ] auditoria;
- [ ] OpenAPI.

## Fase 5 — Portfólio

- [ ] dados demonstrativos reproduzíveis;
- [ ] diagramas atualizados;
- [ ] benchmark;
- [ ] deploy de demonstração;
- [ ] vídeo curto apresentando fluxo e arquitetura.

## Só depois do MVP

S3, Redis, regras configuráveis, layouts específicos, webhooks e frontend.
