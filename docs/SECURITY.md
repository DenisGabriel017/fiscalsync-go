# Segurança e privacidade

## Princípios

- negar acesso por padrão;
- validar autorização dentro de cada caso de uso;
- aplicar `organization_id` em todas as consultas;
- coletar e conservar somente o necessário;
- não registrar conteúdo fiscal, senhas ou tokens.

## Controles mínimos

- senha com Argon2id ou bcrypt;
- access token curto e refresh token rotacionado;
- refresh tokens armazenados como hash;
- RBAC para administrador, analista e leitor;
- limite de tamanho e allowlist de extensões;
- validação pelo conteúdo do arquivo;
- nomes aleatórios no armazenamento;
- rate limiting em login, upload e exportação;
- SQL parametrizado;
- CORS restrito por ambiente;
- segredos somente em variáveis de ambiente;
- dependências verificadas no CI.

## Multi-tenancy

Cada método de repositório recebe `tenantID`. O ID da URL nunca basta para autorizar. Testes devem tentar acessar recursos de outra organização e esperar bloqueio.

## Arquivos

- limitar linhas, memória e tempo de parsing;
- processar em streaming;
- usar dados sintéticos nos testes;
- definir retenção, por exemplo 30 dias;
- permitir remoção administrativa.

## Checklist

- [ ] nenhum segredo no histórico Git;
- [ ] `.env` ignorado;
- [ ] exemplos sem dados reais;
- [ ] logs revisados;
- [ ] testes multi-tenant passando;
- [ ] mensagens externas sem detalhes internos;
- [ ] dependências sem vulnerabilidades críticas conhecidas.
