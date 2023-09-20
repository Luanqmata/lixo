# FOR FOR FOR FOR FOR FOR FOR FOR FOR FOR FOR FOR FOR FOR FOR FOR        

                                                            ## CONTAGEM DO '1 AO 10'
def contagem():
    começo = int(input('Digite o começo da contagem.'))
    final = int(input('Digite o final da contagem.'))
    for i in range(começo, final):                     ##  {  E PARA FAZER A CONTAGEM DESCRESCENTE= for i in range(final,começo,-1):}
        print(i)
contagem()

## _______________________________________________________________________________________________________
                                                        
                                                         ##TABUADA DE QUALQR NUMERO.
import time                                       
def tabuada(numero_escolhido):
    print('\n@@@@@@@@@@@@ Gerador de tabuada. @@@@@@@@@@@@@@\n')
    for multiplicador in range(1, 11):
        calculo = numero_escolhido * multiplicador
        time.sleep(1)
        print(f'{numero_escolhido} x {multiplicador} = {calculo}')

numero_da_tabuada_user = int(input('Digite o numero para o computador fazer a tabuada dele: '))
tabuada(numero_da_tabuada_user)

## __________________________________________________________________________________________________________

                                                            ## TABUADA DO 1 AO 10
 
import time 
def tabuada_1_10():
    print('\n@@@@@@@@@@@@@ TABUADA DO 1 ao 10. @@@@@@@@@@@@@\n')
    for multiplo in range(1,11):
        time.sleep(0.5)
        print(f'\nTABUADA DO {multiplo} : \n')
        for expoente in range(1,11):
            conta = multiplo * expoente
            print(f"    {multiplo} x {expoente} = {conta}")
        
        
tabuada_1_10()

##____________________________________________________________________________________________________________

## WHILE WHILE WHILE WHILE WHILE WHILE WHILE WHILE WHILE WHILE WHILE WHILE WHILE WHILE WHILE WHILE WHILE WHILE 

##                                                                         CONTADOR DO 1 AO 10 

## DO 1 AO 10 
incio = 1
final = 10
while incio <= final:
    print(incio)
    incio +=1
## DO 10 AO 1 
inicio = 10
while inicio >=1:
    print(inicio)
    inicio -=1

##__________________________________________________________________________________________________________________

import time
def tabuada_while(numero_da_tabuada):
    print('\n')
    i = 1
    while i <= 10:
        conta = i * numero_da_tabuada
        print(f'  {i} x {user_tabuada} = {conta}')
        time.sleep(0.5)
        i +=1
        
user_tabuada = int(input(' digite o numero para realizar a tabuada: '))
tabuada_while(user_tabuada)

##_____________________________________________________________________________________________________________________

def tabuada_1_10_while():
    print('\n@@@@@@@@@@@@ Tabuada do 1 ao 10 WHILE @@@@@@@@@@@@\n')
    a = 1
    while a <= 10:
        b = 1 # Reinicializa o contador b a cada iteração de a
        print(f'\nTABUADA DO {a}\n')
        while b <=10:
            conta =a*b
            print(f'{a} x {b} = {conta}')
            b+=1
        a +=1
            
tabuada_1_10_while()

