# Distribuição de Frequências

## Introdução

Este projeto apresenta uma tabela de **Distribuição de Frequências** com cálculos detalhados para compreender dados organizados em classes. Aqui estão incluídos o cálculo da amplitude, número de classes, frequências absolutas e relativas, frequências acumuladas e o ponto médio de cada classe.

---

## Estrutura da Tabela

### Rol e Amplitude Total (A)

O **ROL** é composto pelos dados organizados em ordem crescente.  
A **Amplitude Total** (A) é calculada como:

\[
A = \text{Maior número} - \text{Menor número}
\]

Exemplo:  
\[
A = 173 - 150 = 23
\]

---

### Número de Classes (K)

O número de classes (**K**) é determinado pela fórmula:

\[
K = \sqrt{n}
\]

Onde **n** é o total de elementos no rol.  
Exemplo:  
\[
K = \sqrt{40} \approx 6,32 \approx 6
\]

O número de classes define a quantidade de intervalos na tabela.

---

### Amplitude do Intervalo (h)

A **Amplitude do Intervalo** (**h**) é calculada como:

\[
h = \frac{A}{K}
\]

Exemplo:  
\[
h = \frac{23}{6} \approx 3,8 \approx 4
\]

---

### Construção da Tabela

#### Tabela Inicial com Frequência Absoluta (n)

| **i** | **Classes** | **n (Frequência Absoluta)** |
|-------|-------------|-----------------------------|
| 1     | 150 |- 154  | 4                           |
| 2     | 154 |- 158  | 9                           |
| 3     | 158 |- 162  | 11                          |
| 4     | 162 |- 166  | 8                           |
| 5     | 166 |- 170  | 5                           |
| 6     | 170 |- 174  | 3                           |
|       |             | **40**                      |

O total de **n** (soma das frequências absolutas) é **40**.

---

#### Frequência Absoluta Acumulada (Ni)

| **i** | **Classes** | **n** | **Ni (Frequência Absoluta Acumulada)** |
|-------|-------------|-------|---------------------------------------|
| 1     | 150 |- 154  | 4     |
