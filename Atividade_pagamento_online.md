# 💳 Sistema de Pagamento Online – README

## Parte A – Respostas Simples

### 🔹 Tipos de falhas de comunicação

- **Falha 1:** A mensagem não chega no servidor de pagamento. O cliente pensa que pagou, mas o pagamento nunca aconteceu.  
- **Falha 2:** O pagamento acontece, mas a resposta não volta para o frontend. O cliente pode tentar pagar de novo e acabar sendo cobrado duas vezes.  
- **Falha 3:** A mesma mensagem é enviada mais de uma vez. O sistema pode tentar cobrar o mesmo valor repetidamente.  

---

### 🔹 Solução para evitar duplicação

- Usar o **id_transacao** como uma “etiqueta única” para cada pagamento.  
- Se o servidor de pagamento receber o mesmo id outra vez, ele não faz a cobrança de novo, só devolve a resposta anterior.  
- Assim, mesmo que a mensagem chegue repetida, o cliente só será cobrado uma vez.  

---

### 🔹 Qual semântica escolher?

- “No máximo uma vez” pode falhar, porque se a mensagem se perder o pagamento nem acontece.  
- A melhor opção é **idempotência**, porque permite reenviar a operação sem perigo de cobrar duas vezes.  

📚 **Discussão estilo estudantes de faculdade (6º semestre):**  
> A semântica mais adequada para uma operação crítica é a **idempotência**.  
> Isso porque, em sistemas distribuídos, pode acontecer de a mensagem ser enviada mais de uma vez (por falha de rede, reenvio automático, etc.).  
> Se usarmos "no máximo uma vez", corremos o risco de perder a operação (caso a mensagem falhe, ela nunca mais será executada).  
> Já com idempotência, mesmo que a operação seja repetida, o resultado final será o mesmo, garantindo segurança e consistência.  

🔄 **Tipos de comunicação relacionados:**  
- **Assíncrona e persistente:** funciona bem, porque a mensagem fica guardada até ser entregue, mesmo se houver falhas.  
- Portanto, a melhor escolha é **idempotência + comunicação assíncrona e persistente**, para garantir que a operação sempre aconteça e não cause erro mesmo se for repetida.  

---

## Parte B – Pontos da Discussão

- O **id_transacao** é essencial para diferenciar um pagamento novo de um repetido.  
- O servidor precisa guardar os **ids já processados** para não cobrar outra vez.  
- O desafio é equilibrar:  
  - Garantir que o pagamento chegue mesmo com falhas.  
  - Mas não deixar que isso vire cobrança duplicada.  
