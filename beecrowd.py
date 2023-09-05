## 1011
def volume_esfera():
    raio = float(input())
    volume = (4/3) * 3.14159 * raio**3
    volume_format = "{:.3f}".format(volume)
    return volume_format

    resultado = volume_esfera()
    print('VOLUME =',resultado)

## 1010
def leitor1():
    codigo1 = int(input())
    pecas1 = int(input())
    valor1 = float(input())
    codigo2 = int(input())
    pecas2 = int(input())
    valor2 = float(input())
    
    total_a_pagar = (pecas1 * valor1) + (pecas2 * valor2)
    return total_a_pagar

total = leitor1()
print('VALOR A PAGAR:', total)
 
 
##1013
def maior():
    num1 = int(input())
    num2 = int(input())
    num3 = int(input())
    if num1>num2 and num1>num3:
        print(num1,'eh o maior')
    elif num2>num1 and num2>num3:
        print(num2,'eh o maior')
    else: num3>num1 and num3>num2
    print(num3,'eh o maior')
print(maior())
