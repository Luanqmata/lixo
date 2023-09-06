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
________________________________________________________________________________________________________________________________________
import random

def jokenpo():
    opcoes_do_jogo = ['pedra', 'papel', 'tesoura']
    selecao_do_computador = random.choice(opcoes_do_jogo)
    selecao_do_usuario = opcoes_do_jogo[int(input('Escolha entre pedra, papel ou tesoura: \n 1. pedra \n 2. papel \n 3. tesoura \n'))-1]
    print(f'O computador escolheu {selecao_do_computador}.')
    print(f'O usuário escolheu {selecao_do_usuario}.')

    if selecao_do_usuario == selecao_do_computador:
        print('Empate!')
    elif selecao_do_usuario == 'pedra' and selecao_do_computador == 'tesoura':
        print('Você ganhou!')
    elif selecao_do_usuario == 'papel' and selecao_do_computador == 'pedra':
        print('Você ganhou!')
    elif selecao_do_usuario == 'tesoura' and selecao_do_computador == 'papel':
        print('Você ganhou!')
    else:
        print('Você perdeu!')
    input('Pressione ENTER para jogar novamente...')
    jokenpo()

jokenpo()


