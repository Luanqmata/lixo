## SEQUENCIA FIBONACCI :

def fibonacci(numero_de_elementos):
    if numero_de_elementos == 1:
        return 0
    elif numero_de_elementos ==2:
        return[0,1]
    else:
        fibonacci = [0,1]
        for i in range (2,numero_de_elementos):
            novo_elemento_lista = fibonacci[-1]+fibonacci[-2]
            fibonacci.append(novo_elemento_lista)
        return fibonacci
print(fibonacci(int(input("digite o numero de vezes que a sequencia fibonacci deve aparecer: "))))

##AREA TRIANGULO:
def area_triangulo(base,altura):
    conta = (base*altura)/2
    return print(f'A área do triangulo é igual a {conta}')

base_user = float(input('digite o valor da base do triangulo. '))
altura_user = float(input('digite o valor da Altura do triangulo. '))

print(area_triangulo(base_user,altura_user))

def fatorial_cinco():
    numero = int(input('digite o numero. '))
    if numero % 5 == 0:
        return print(f'O numero escolhido foi {numero}.\n E ele é divisivel por 5.')
    else:
        return print(f'o numero não é divisivel por 5.')
    
print(fatorial_cinco())    