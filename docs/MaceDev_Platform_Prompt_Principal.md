# MaceDev Platform — Visão Arquitetural Principal

## 1. Propósito do documento

Este documento é a fonte principal da visão arquitetural e de produto da MaceDev Platform. Ele orienta decisões futuras, preserva os princípios do projeto e delimita a evolução incremental da plataforma.

Ele não descreve funcionalidades já entregues, nem impõe tecnologias, serviços ou decisões de implementação que ainda não tenham sido definidos e aprovados.

## 2. Visão de produto

A MaceDev Platform é uma plataforma de Engenharia de Software, Dados, Arquitetura e IA orientada por agentes.

A visão é permitir análise arquitetural e governada de dados, bancos de dados, SQL e APIs, gerando diagnósticos e recomendações baseados em evidências. A plataforma deverá preparar informações para consumo seguro e contextualizado por agentes de IA e permitir ações controladas quando houver autorização explícita.

Capacidades previstas incluem:

- análise de arquitetura de dados;
- descoberta e análise de bancos de dados;
- conexão inicialmente somente leitura com Oracle e PostgreSQL;
- análise de schemas, tabelas, views, procedures, functions, triggers, índices, constraints, relacionamentos e dependências;
- análise de SQL;
- análise, arquitetura e governança de APIs;
- governança de dados e arquitetura;
- ontologia como camada semântica transversal;
- preparação de dados para consumo por agentes de IA, sob o conceito de AI-Ready Data;
- diagnóstico arquitetural;
- recomendações baseadas em evidências;
- execução controlada de ações por agentes;
- interfaces CLI, API e Frontend;
- segurança, autorização, auditoria, observabilidade e rastreabilidade completa das ações dos agentes.

Essas capacidades são visão de evolução, não confirmação de implementação atual.

## 3. Arquitetura conceitual

```text
MaceDev Platform
    |
    +-- CLI
    +-- API
    +-- Frontend
    |
    +-- Agent Platform
            |
            +-- Dados
            +-- Governança
            +-- Ontologia
            +-- APIs
            +-- Arquitetura
            |
            +-- AI-Ready Data
            |
            +-- Diagnóstico
            +-- Recomendação
            +-- Ação controlada
            +-- Auditoria
```

As interfaces são pontos de interação. A Agent Platform representa a capacidade conceitual central de reunir contexto, análise, governança, evidências, recomendações e ações controladas. Esta estrutura não determina, por si só, fronteiras físicas de processos, serviços ou módulos.

## 4. Princípios arquiteturais

| Princípio | Diretriz |
| --- | --- |
| Governance by Design | Governança deve fazer parte do desenho das capacidades, e não ser adicionada apenas depois da análise ou da ação. |
| Security by Design | Segurança deve orientar acessos, ferramentas, integrações, segredos e fluxos desde o início. |
| API-First | Capacidades compartilhadas devem poder ser expostas por contratos claros, sem depender de uma interface específica. |
| AI-Ready by Design | Dados e contexto devem ser preparados para uso confiável, governado e compreensível por IA. |
| Ontology-Aware | A semântica e as relações entre conceitos devem acompanhar a informação técnica analisada. |
| Human-in-the-Loop | Pessoas mantêm autoridade sobre decisões e ações sensíveis. |
| Read-Only by Default | Consultas e descobertas devem ser somente leitura por padrão. |
| Evidence-Based Recommendations | Diagnósticos e recomendações devem apontar para evidências verificáveis. |
| Full Action Traceability | Decisões e ações relevantes devem produzir registros rastreáveis de ponta a ponta. |
| Separation of Concerns | Interfaces, capacidades de domínio, infraestrutura e automações devem ter responsabilidades claras. |
| Least Privilege | Cada identidade, integração e ferramenta deve receber apenas o acesso necessário. |
| Observability by Design | Eventos, erros, métricas e fluxos relevantes devem poder ser observados. |
| Automate Before Manual Execution | Automação controlada deve ser priorizada para atividades repetitivas e verificáveis. |
| Version Everything | Código, contratos, documentação, regras e mudanças relevantes devem ser versionados quando aplicável. |
| Architecture Before Implementation | Mudanças relevantes devem partir de contexto, dependências e decisão arquitetural explícita. |

## 5. Segurança e autorização

A plataforma deverá adotar leitura como padrão, inclusive para descoberta e análise de bancos de dados. Operações destrutivas não devem ser executadas por padrão.

Qualquer ação sensível deverá exigir autorização explícita, respeitar o princípio do menor privilégio e permanecer sujeita a human-in-the-loop. Ferramentas usadas por agentes devem ser controladas, com escopo conhecido e apropriado à ação solicitada. Segredos devem ser protegidos e não devem ser expostos em código, logs ou evidências.

Auditoria e rastreabilidade devem registrar contexto, autorizações, ferramentas e resultados das atividades relevantes.

## 6. Rastreabilidade de agentes

O conceito de **Agent Action Lineage** descreve a cadeia mínima de significado e evidência para decisões ou ações relevantes:

```text
Usuário
→ solicitação
→ agente
→ conceito ontológico
→ regra de governança
→ autorização
→ ferramenta
→ fonte
→ evidência
→ análise
→ recomendação
→ aprovação
→ ação
→ resultado
```

Toda decisão ou ação relevante de um agente deverá possuir contexto, significado semântico, governança, evidência e rastreabilidade. A cadeia deve permitir compreender o que foi solicitado, por qual agente foi tratado, quais fontes e ferramentas foram usadas, qual análise foi produzida, que autorização existiu e qual resultado ocorreu.

## 7. Arquitetura inicial de agentes

A visão inicial prevê agentes especializados com responsabilidades conceituais distintas:

- **Orchestrator Agent**: coordena solicitações, contexto, etapas e encaminhamentos entre capacidades.
- **Database Agent**: realiza descoberta e análise de estruturas e metadados de bancos de dados autorizados.
- **SQL Agent**: analisa consultas SQL e seus aspectos técnicos, semânticos e de governança.
- **API Agent**: analisa contratos, arquitetura, governança e integrações de APIs.
- **Governance Agent**: aplica ou avalia regras e controles de governança.
- **Ontology Agent**: relaciona ativos técnicos a conceitos e significados semânticos.
- **Architecture Agent**: produz diagnósticos sobre arquitetura de dados, APIs e componentes.
- **Recommendation Agent**: gera recomendações baseadas nas evidências e análises disponíveis.

Essa divisão é uma visão arquitetural inicial. Ela não significa que todos esses agentes existam ou precisem ser implementados imediatamente.

## 8. Dados, metadados, semântica e IA

A evolução conceitual esperada para dados é:

```text
Database
→ Discovery
→ Metadata
→ Semantics
→ Governance
→ AI-Ready Data
→ Agents
```

A plataforma deverá manter distinções claras entre:

- **dados**: conteúdo e ativos de informação;
- **metadados**: descrições técnicas e estruturais dos ativos;
- **semântica**: significado, conceitos, relações e contexto ontológico;
- **governança**: regras, classificações, controles, responsabilidades e autorizações;
- **consumo por IA**: uso de contexto preparado e controlado por agentes.

AI-Ready Data representa a preparação de dados, metadados, semântica e governança para que agentes de IA consumam informação com contexto, evidência e limites apropriados.

## 9. APIs

API-First é um princípio de evolução da plataforma. Futuramente, capacidades expostas por API deverão considerar:

- contratos claros;
- OpenAPI;
- autenticação;
- autorização;
- versionamento;
- auditoria;
- observabilidade;
- rate limiting;
- integração com agentes.

Esses itens são requisitos arquiteturais futuros; não representam APIs já disponíveis.

## 10. Interfaces

A plataforma prevê três interfaces complementares:

- **CLI**: interface para fluxos de desenvolvimento, automação e operação assistida;
- **API**: interface contratual para integração programática;
- **Frontend**: interface visual para exploração, acompanhamento, decisão e operação autorizada.

As interfaces devem consumir as capacidades centrais da plataforma sem duplicar regras de domínio, segurança, governança ou lógica de análise.

## 11. Estrutura arquitetural inicial do repositório

| Diretório | Responsabilidade arquitetural |
| --- | --- |
| `apps/` | Aplicações executáveis e interfaces de entrada, como CLI, API ou Frontend. |
| `services/` | Capacidades de domínio e serviços que venham a ser separados por uma responsabilidade concreta. |
| `infra/` | Definições e documentação de infraestrutura, ambientes e implantação quando forem necessárias. |
| `scripts/` | Automações de desenvolvimento, validação, operação e manutenção, quando justificadas. |
| `docs/` | Visão arquitetural, decisões, guias, referências e documentação operacional. |

Essa organização define responsabilidade inicial, mas não exige a criação imediata de componentes em todos os diretórios.

## 12. Roadmap incremental

O roadmap a seguir é planejamento, não implementação confirmada:

| Sprint | Foco planejado |
| --- | --- |
| Sprint 0 | Ambiente |
| Sprint 1 | Git + AI Engineering |
| Sprint 2 | Prompt Engineering |
| Sprint 3 | Arquitetura antes do código |
| Sprint 4 | Go |
| Sprint 5 | Python |
| Sprint 6 | Agents, Tools, Context e Security |
| Sprint 7 | Data Discovery |
| Sprint 8 | Governance |
| Sprint 9 | Ontology |
| Sprint 10 | AI-Ready Data |
| Sprint 11 | APIs |
| Sprint 12 | Frontend |
| Sprint 13 | Docker/Linux |
| Sprint 14 | Cloud |
| Sprint 15 | CI/CD/Observability |
| Sprint 16 | MaceDev Architecture Advisor |

Cada sprint deve ser decomposto em objetivos pequenos, com dependências, validação e autorização adequadas antes de qualquer implementação relevante.

## 13. Estado atual da implementação

Atualmente, o projeto possui apenas a primeira aplicação CLI em Go, localizada em `apps/cli-go`. Os demais componentes descritos neste documento — incluindo serviços, agentes especializados, conexões com bancos de dados, APIs, Frontend, infraestrutura e automações — são evolução planejada e não estão implementados.

Este documento deve ser consultado antes de mudanças arquiteturais ou funcionalidades relevantes, em conjunto com as instruções do repositório.
