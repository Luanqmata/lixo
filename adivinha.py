import random
import time
def jogo_adivinha():
    numero_aleatorio = random.randint(0,10)
    return numero_aleatorio

numero_do_jogo = jogo_adivinha()
tentativas = 0


while True:
    numero_escolhido = int(input("Digite o numero do jogo adivinha: '0 a 10'. "))
    tentativas += 1
    if numero_do_jogo == numero_escolhido :
        print(f"parabens você acertou. e teve {tentativas} tentativas!")
        break
    else:
        print(f"você errou tente novamente.")
        time.sleep(0.7)
