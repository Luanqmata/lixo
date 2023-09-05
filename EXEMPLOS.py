## ARRAY                     ARRAY
def lista_de_usuarios():
    lista_de_usuarios = ["João", "Maria", "José", "Pedro", "Ana"]
    usuario_selecionado = int(input("digite a colocação do usuario: "))
    usuario_selecionado = lista_de_usuarios[usuario_selecionado -1].strip().split()
    return usuario_selecionado[0]

print(lista_de_usuarios())
______________________________________________________________________________________
