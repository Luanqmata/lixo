                                                  
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
