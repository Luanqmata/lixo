"""
import array as ar 

obj = ar.array('i',[1,2,3])
print(type(obj),obj)
# print(obj[1]) selecionado 
# print(obj[1:2]) fatiado 



obj2 = ar.array('u','YESGO') # ARCAICO
print(type(obj2),obj2)
obj2[0] = 'I'
print(obj2)   ## TROCA LETRA ERRADO POR LETRA CERTA  


obj3 = ar.array('d',[1.233 , 1.2334 , 3.3])
print(type (obj3),obj3)
"""


""" exemplos 
typecode    Python data type     Byte size
  
'b' signed interger 1 
'B' unsingned interger 
'u' inicode character 
'h' signed interger 
'H' 
"""
#PARTE 3
"""
#____________________________________________________________________________________________

import array as arr

#criar um array de inteiro vazios 

meu_array = arr.array('i')

#adicionar 5 numeros ao array 

for i in range(5):
    num = int(input('insira os numeros inteiros. '))
    meu_array.append(num)

print('Array resultante',meu_array)

# somar todos os numeros do array 

print(len(meu_array))#contar qyantidade de numeros 

print(sum(meu_array))#somar Numeros.

# encontrar o min e o max 
min_valor = min(meu_array)
max_valor = max(meu_array)

print(f'O menor valor é: {min_valor}, o maior valor é: {max_valor}')

# Remover o último elemento 

print(meu_array)
meu_array.pop()
print('O ultimo elemento foi excluido:',meu_array)

# inverter a Ordem

meu_array.reverse()
print('lista invertida', meu_array )

"""
"""

try:
    num = int(input('Digite um numero inteiro'))
    resultado = 10/num
except(ZeroDivisionError,ValueError)
    print('Certifique se você digitou um numero valido , ou não digitar o numero 0 ')
else:
    print(f'O resutlado da operaçao de 10/{num} = {resultado}')

"""
