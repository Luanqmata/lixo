## 1
def multiplo_cinco():
    numero_dado = int(input('Digite o numero.'))
    if numero_dado % 5 == 0:
        return True
    else:
        return False
print(multiplo_cinco())
## 2 
import time

def multiplo_cinco_tres():
    numero_user = int(input('Digite um numero.'))
    time.sleep(1)
    
    if numero_user % 5 ==0:
        print('\nMultiplo de 5.\n')
        return True 
    elif numero_user % 3 ==0:
        print('\nMultiplo de 3.\n')
        return True
    else:
        print('\neste numero nao é divisivel por 3 e nem por 5.\n')
        return False

print(multiplo_cinco_tres())
## 3 
## 4 
import time
def lista_frutas():
    lista = []
    while True:
        
        frutas = input("Digite a Fruta. ou 'sair' para sair. ")
        frutasminusculo = frutas.lower()
        if frutas == 'sair':
            print('fazendo anotações.')
            time.sleep(2)
            print('Voce saiu.\nOs objetos que foram adicionados a lista são:')
            print(lista)
            exit()
        lista.append(frutasminusculo)
        print('A fruta Foi adicionada a Lista.')
                   
print(lista_frutas())
## 5
def area_circulo():
    raio = float(input('Digite o valor do raio da circuferencia.'))
    pi = 3.141516
    area = pi * raio**2
    area_arredondada = "{:.2f}".format(area)
    perimetro = 2*pi*raio
    perimetro_arredondado = "{:.2f}".format(perimetro)
    return print(f'A area do circulo é igual a: {area_arredondada}m².\n O perimetro do circulo é igual a: {perimetro_arredondado}m².')

print(area_circulo())
## 9
import random
import time

def jogo_adivinha():
    numero_aleatorio = random.randint(0,10)
    print('\nA maquina esta pensando um numero.')
    time.sleep(1)
    print('pensando.\n')
    time.sleep(1)
    print('pensado...\n')
    
    numero_user = int(input('Digite o numero que a Maquina pensou.'))
    print('analisando...')
    time.sleep(1)

    if numero_aleatorio == numero_user:
        print('Você Acertou parabens.')
    elif numero_user < 0:
        print('o numero deve ser maior que 0')
    elif numero_user > 10:
        print('o numero deve ser menor que 10')
    else:
        print('Você errou infelizmente.')

print(jogo_adivinha())
## 10 
def contador_de_a():
    frase = input('digite a frase: ')
    print("a volgal A aparece:", frase.count('a')+frase.count('A') ,"Vezes.")

print(contador_de_a())
