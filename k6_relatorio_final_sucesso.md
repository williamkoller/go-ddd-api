# 📊 Relatório de Teste de Carga - K6 (Finalizado com Sucesso)

### 🧪 Cenário: Criação de Usuários via API (`POST /users/register`)

- **Ferramenta**: [k6.io](https://k6.io)
- **Script**: `tests/k6/load-tests.js`
- **Endpoint testado**: `http://localhost:3003/users/register`
- **Tipo de carga**: Pico de 200 VUs por 4 minutos contínuos
- **Data do teste**: 2025-04-30 17:09:09

---

## 🔧 Configuração do Teste

| Parâmetro              | Valor                      |
|------------------------|----------------------------|
| Virtual Users (máximo) | 200                        |
| Duração total          | 4m + 30s de graceful stop  |
| Iterações totais       | 74.771                     |
| Requisições HTTP       | 74.771                     |
| Média por segundo      | 311.00 req/s               |

---

## ✅ Verificações (`checks`)

| Verificação              | Sucesso | Total   | Taxa de Sucesso |
|--------------------------|---------|---------|-----------------|
| status is 201            | ✓       | 149.542 | 100%            |
| response is JSON         | ✓       | 149.542 | 100%            |
| http_req_failed < 1%     | ✓       | 0% falhas | ✅             |

---

## ⏱️ Desempenho das Requisições

| Métrica                 | Valor        |
|-------------------------|--------------|
| Tempo médio (`avg`)     | 1.05 ms      |
| Tempo mínimo (`min`)    | 144.65 µs    |
| Tempo mediano (`med`)   | 806.51 µs    |
| Tempo máximo (`max`)    | 20.45 ms     |
| Percentil 90 (`p90`)    | 1.91 ms      |
| Percentil 95 (`p95`)    | 2.57 ms      |

---

## 🔄 Iterações

| Métrica                     | Valor        |
|-----------------------------|--------------|
| Iterações concluídas        | 74.771       |
| Iterações/s                 | 311.00       |
| Duração média por iteração  | 501.98 ms    |

---

## 🌐 Uso de Rede

| Tipo de dado     | Total |
|------------------|-------|
| Dados recebidos  | 19 MB |
| Dados enviados   | 20 MB |

---

## 🧵 VUs (Virtual Users)

| Métrica        | Valor |
|----------------|-------|
| VUs mínimos    | 2     |
| VUs máximos    | 200   |

---

### 📘 Observações

- Todas as requisições foram processadas com sucesso (`status 201`).
- Nenhuma falha registrada durante o teste de 200 VUs.
- O sistema manteve uma latência extremamente baixa (p95 < 3ms).
- A estratégia de enfileiramento assíncrono com múltiplos workers demonstrou alta eficiência.

