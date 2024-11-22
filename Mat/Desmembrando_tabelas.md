# Distribuição de Frequências

## Introdução

Distribuição de frequências é uma ferramenta estatística utilizada para organizar e interpretar dados em grupos ou classes. Este guia mostra como criar uma tabela de distribuição de frequência a partir de dados brutos, destacando cálculos e conceitos envolvidos.

---
![image](https://github.com/user-attachments/assets/716937ea-d9ae-475e-a850-15077d7a2ca7)

## Passo 1: ROL (Organização dos Dados)
Os dados devem ser organizados em ordem crescente.

---

## Passo 2: Cálculo da Amplitude Total (A)

**Fórmula:**  
\[ A = *Maior Número* - *Menor Número* \]  
**Exemplo:**  
\[ A = 173 - 150 = 23 \]

---

## Passo 3: Determinação do Número de Classes (K)

**Fórmula:**  
\[ K = √n \]  
Onde \( n \) é o total de dados na tabela.  

**Exemplo:**  
\[ K =  √40 ~= 6,32 ~~= 6 \]  
As classes formam a quantidade de linhas da tabela.
| \( i \) |
|---------|
| 1       |
| 2       |
| 3       |
| 4       |
| 5       |
| 6       |

---

## Passo 4: Cálculo da Amplitude do Intervalo (h)

**Fórmula:**  
\[ h = A/K \]  
**Exemplo:**  
\[ h = 23/6 ~= 3,8 ~~= 4 \]

---

## Passo 5: Tabela Inicial com Classes e Frequência Absoluta \( n \)

| \( i \) | Classes     | \( n \) |
|---------|-------------|---------|
| 1       | 150 - 154   | 4       |
| 2       | 154 - 158   | 9       |
| 3       | 158 - 162   | 11      |
| 4       | 162 - 166   | 8       |
| 5       | 166 - 170   | 5       |
| 6       | 170 - 174   | 3       |
| **Σ**   |             | **40**  |

---

## Passo 6: Frequência Absoluta Acumulada (\ N_i \)

**Definição:** Soma acumulada dos valores de \( n \).  

| \( i \) | Classes     | \( n \) | \( N_i \) |
|---------|-------------|---------|-----------|
| 1       | 150 - 154   | 4       | 4         |  ( 4 + nada = 4 ) 
| 2       | 154 - 158   | 9       | 13        |  ( 9 + 4 = 13 )
| 3       | 158 - 162   | 11      | 24        |  ( 11 + 9 + 4 = 24 )
| 4       | 162 - 166   | 8       | 32        |  ( 8 + 11 + 9 + 4 = 32 )
| 5       | 166 - 170   | 5       | 37        |  ( 5 + 8 + 11 + 9 + 4 = 37 )
| 6       | 170 - 174   | 3       | 40        |  ( 3 + 5 + 8 + 11 + 9 + 4 = 40 )

---

## Passo 7: Frequência Relativa (\( f_i \))

**Fórmula:**  
\[ f_i = \frac{n}{\Sigma n} \]  
**Exemplo:** Para \( n = 4 \):  
\[ f_i = \frac{4}{40} = 0,1 \]

| \( i \) | Classes     | \( n \) | \( N_i \) | \( f_i \) |
|---------|-------------|---------|-----------|-----------|
| 1       | 150 - 154   | 4       | 4         | 0,1       |
| 2       | 154 - 158   | 9       | 13        | 0,225     |
| 3       | 158 - 162   | 11      | 24        | 0,275     |
| 4       | 162 - 166   | 8       | 32        | 0,2       |
| 5       | 166 - 170   | 5       | 37        | 0,125     |
| 6       | 170 - 174   | 3       | 40        | 0,075     |
| **Σ**   |             | **40**  |           | **1,0**   |

---

## Passo 8: Frequência Relativa Acumulada (\( F_i \))

**Definição:** Soma acumulada dos valores de \( f_i \).

| \( i \) | Classes     | \( n \) | \( N_i \) | \( f_i \) | \( F_i \) |
|---------|-------------|---------|-----------|-----------|-----------|
| 1       | 150 - 154   | 4       | 4         | 0,1       | 0,1       |
| 2       | 154 - 158   | 9       | 13        | 0,225     | 0,325     |
| 3       | 158 - 162   | 11      | 24        | 0,275     | 0,6       |
| 4       | 162 - 166   | 8       | 32        | 0,2       | 0,8       |
| 5       | 166 - 170   | 5       | 37        | 0,125     | 0,925     |
| 6       | 170 - 174   | 3       | 40        | 0,075     | 1,0       |

---

## Passo 9: Ponto Médio (\( x^1 \))

**Fórmula:**  
\[ x^1 = \frac{\text{Limite Inferior} + \text{Limite Superior}}{2} \]  
**Exemplo:** Para \( 150 - 154 \):  
\[ x^1 = \frac{150 + 154}{2} = 152 \]

| \( i \) | Classes     | \( n \) | \( N_i \) | \( f_i \) | \( F_i \) | \( x^1 \) |
|---------|-------------|---------|-----------|-----------|-----------|-----------|
| 1       | 150 - 154   | 4       | 4         | 0,1       | 0,1       | 152       |
| 2       | 154 - 158   | 9       | 13        | 0,225     | 0,325     | 156       |
| 3       | 158 - 162   | 11      | 24        | 0,275     | 0,6       | 160       |
| 4       | 162 - 166   | 8       | 32        | 0,2       | 0,8       | 164       |
| 5       | 166 - 170   | 5       | 37        | 0,125     | 0,925     | 168       |
| 6       | 170 - 174   | 3       | 40        | 0,075     | 1,0       | 172       |


---

## Considerações Finais

Este documento resume os passos para construir uma tabela de distribuição de frequências com \( n \), \( N_i \), \( f_i \), \( F_i \) e \( x^1 \), fornecendo uma visão clara e ordenada dos dados.
