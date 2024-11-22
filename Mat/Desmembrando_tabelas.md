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

# ANOTEÇÕES BRUTAS
  ```txt 
  #titulo Distribuição de Frequências


( imagem da tabela )

ROL = informaçoes organizadas
(A) Amplitude total = maior numero - menor numero (ex: A = 173 - 150 = 23)


(K) N° classes = 

		K = √n		n = Representa a qnt de espaços que estão preenchidos na tabela toda.
	
	exemplo: K = √40 ~= 6,32 ~~= 6 (o numero de classes é igual a 6)

		as classes vão ser o tamanho da qnt de colunas que a tabela vai ter.
_____
| i |
| 1 |
| 2 |
| 3 |
| 4 |
| 5 |
| 6 |

(h) Amplitude do meu intervalo

	     h = A/K 
 h = 23/6 
 h ~= 3,8 ~~= 4 

( imagem da tabela )
_____________________
| i |  classes |  n  | esse "n" é chamado de frequencia absoluta , (novo integrante da tabela)
| 1 | 150|-154 |  4  |
| 2 | 154|-158 |  9  |
| 3 | 158|-162 | 11  |
| 4 | 162|-166 |  8  |
| 5 | 166|-170 |  5  |
| 6 | 170|-174 |  3  |
               |  40 | = SUM  somatorio da frequencia absoluta ( professor pode perguntar!! )

 " Ni "  = frequencia absoluta acumulada  ( vai ser o numero de "n" { frequencia absoluta } + a soma dos anteriores da coluna de "n")
____________________________
| i |  classes |  n  |  Ni |
| 1 | 150|-154 |  4  |  4  | ( 4 + nada = 4 ) 
| 2 | 154|-158 |  9  |  13 | ( 9 + 4 = 13 )
| 3 | 158|-162 | 11  |  24 | ( 11 + 9 + 4 = 24 )
| 4 | 162|-166 |  8  |  32 | ( 8 + 11 + 9 + 4 = 32 )
| 5 | 166|-170 |  5  |  37 | ( 5 + 8 + 11 + 9 + 4 = 37 )
| 6 | 170|-174 |  3  |  40 | ( 3 + 5 + 8 + 11 + 9 + 4 = 40 )
               |  40 | 

 " fi "  = frequencia relativa ( porcetagem ou parte decimal que isso representa da total 40 pesquisas )
Oque fazer = pegar o indece que esta na coluna " n " e dividir pelo total do somatorio de " n " (frequencia absoluta ) { n / SUM(n) }  
____________________________________
| i |  classes |  n  |  Ni |  fi   |
| 1 | 150|-154 |  4  |  4  |  0,1  |  ( 4/40= 10/100 = 0,1 )
| 2 | 154|-158 |  9  |  13 | 0,225 |  ( 9/40 )
| 3 | 158|-162 | 11  |  24 | 0,275 |  ( 11/40 )
| 4 | 162|-166 |  8  |  32 |  0,2  |  ( 8/40 )
| 5 | 166|-170 |  5  |  37 | 0,125 |  ( 5/40 )
| 6 | 170|-174 |  3  |  40 | 0,075 |  ( 3/40 )
               |  40 |     |SUM = 1|  A soma tem q ser igual a 1 ( na decimal tende ser igual a 1 e na porcentagem tem q dar 100 )

 " Fi " Freqncia relativa acumulada ( a mesma coisa da frequencia acumulada )
____________________________________________
| i |  classes |  n  |  Ni |  fi   |  Fi   |
| 1 | 150|-154 |  4  |  4  |  0,1  |  0,1  |  ( 0,1 + nada = 0,1 )
| 2 | 154|-158 |  9  |  13 | 0,225 | 0,325 |  ( 0,225 + 0,1 = 0,325 )
| 3 | 158|-162 | 11  |  24 | 0,275 |  0,6  |  ( 0,275 + 0,225 + 0,1 = 0,6 )
| 4 | 162|-166 |  8  |  32 |  0,2  |  0,8  |  ( 0,2 + 0,275 + 0,225 + 0,1 = 0,8 )
| 5 | 166|-170 |  5  |  37 | 0,125 | 0,925 |  ( 0,125 + 0,2 + 0,275 + 0,225 + 0,1 = 0,925 )
| 6 | 170|-174 |  3  |  40 | 0,075 |   1   |  ( 0,075 + 0,125 + 0,2 + 0,275 + 0,225 + 0,1 = 1 )
               |  40 |     |SUM = 1|


 " x¹ "  Ponto medio ( envolve tirar a media das classes, somar e dividir por 2 { ou a quantidade de numeros que aparecerem na classe relativo a media } )
____________________________________________________
| i |  classes |  n  |  Ni |  fi   |  Fi   |   x¹  |
| 1 | 150|-154 |  4  |  4  |  0,1  |  0,1  |  152  |  ( 150 + 154 = 304 ; 304 / 2 = 152 )
| 2 | 154|-158 |  9  |  13 | 0,225 | 0,325 |  156  |  ( 154 + 158 = 312 ; 312 / 2 = 156 )
| 3 | 158|-162 | 11  |  24 | 0,275 |  0,6  |  160  |  ( 158 + 162 = 320 ; 320 / 2 = 160 )
| 4 | 162|-166 |  8  |  32 |  0,2  |  0,8  |  164  |  ( 160 +  h = 164 )
| 5 | 166|-170 |  5  |  37 | 0,125 | 0,925 |  168  |  ( 164 +  h = 168 )
| 6 | 170|-174 |  3  |  40 | 0,075 |   1   |  172  |  ( 168 +  h = 172 )
               |  40 |     |SUM = 1|
  ```
