# Contrato inicial da API

Base: `/api/v1`. Durante a implementação, mantenha o contrato executável em `openapi.yaml`.

## Rotas

| Método | Rota | Descrição |
|---|---|---|
| POST | `/auth/register` | Cria usuário e organização |
| POST | `/auth/login` | Retorna access e refresh token |
| POST | `/auth/refresh` | Renova sessão |
| POST | `/auth/logout` | Revoga refresh token |
| POST | `/reconciliations` | Cria conciliação |
| GET | `/reconciliations` | Lista com filtros e paginação |
| GET | `/reconciliations/{id}` | Detalhes e último resumo |
| POST | `/reconciliations/{id}/files` | Envia fonte A ou B |
| POST | `/reconciliations/{id}/runs` | Inicia processamento |
| GET | `/reconciliations/{id}/runs/{runId}` | Status e totais |
| GET | `/reconciliations/{id}/runs/{runId}/discrepancies` | Lista divergências |
| GET | `/reconciliations/{id}/runs/{runId}/export` | Exporta CSV |

## Criação

```json
{
  "name": "Conciliação julho de 2026",
  "reference_month": "2026-07-01"
}
```

Resposta `201` contém `id`, `status: DRAFT` e `created_at`.

## Upload

`multipart/form-data` com `source` (`A` ou `B`) e `file` (`.csv` ou `.xlsx`). Retorne `202 Accepted`. O processamento só inicia após duas fontes válidas.

## Resumo

```json
{
  "status": "COMPLETED",
  "totals": {
    "source_a": 1200,
    "source_b": 1192,
    "matched": 1178,
    "missing_in_a": 4,
    "missing_in_b": 12,
    "value_mismatch": 6
  }
}
```

## Erros

Use `application/problem+json` com `type`, `title`, `status`, `detail`, `instance` e `request_id`. Nunca retorne stack trace.

## Paginação e idempotência

Use `page` e `page_size`, máximo 100. O início de uma execução aceita `Idempotency-Key`; repetições retornam a mesma execução.
