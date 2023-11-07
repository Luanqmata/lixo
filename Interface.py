"""
dpkg -1 | grep python3-tk

import tkinter as tk

def mostrar_mensagem():    
    label.config(text='Ola iesgo!')

#criar uma janela 

janela = tk.Tk()
janela.title('Exemplo de gui em python')

#criar um rotulo 

label = tk.Label(janela,text='Bemm vindo a interfaçe grafica de usuario')
label.pack(padx=10,pady=10)

# criar um botao do tipo CTA (call to action)

botao = tk.Button(janela,text="Clique aqui!",command=mostrar_mensagem)
botao.pack(padx=10,pady=10)
botao.config(width=15,height=20)

"""
#__________________________________________________________________________________________________


import tkinter as tk
from tkinter import messagebox 
# import pesquisar_produto import messagebox # nao vai rodar porque estao no mesmo codigo tem q dividir no meio 
import pesquisa_produto


def abrir_janela(mensagem):
    nova_janela = tk.Toplevel()
    nova_janela.title('Ação qualquer....')

    label = tk.Label(nova_janela,text= mensagem,padx=20,pady=20)
    label.pack()

    botao_fechar= tk.Button(nova_janela,text='Sair',command=nova_janela.destroy)
    botao_fechar.pack(padx=20,pady=10)

def pesquisar_produtos():
    abrir_janela('Pesquisar Produto ')

def checar_estoque():
    abrir_janela('Checar estoque')

def remover_estoque():
    abrir_janela('Removquer Produto.')

def acrescentar_produto():
    abrir_janela('Cadastrar Produto.')

def cadastrar_produto():
    abrir_janela('Cadastrar produto.')
 
#criar a janela principal

janela_principal = tk.Tk()
janela_principal.title('Sistema Erp Iesgo.')
janela_principal.attributes('-fullscreen',True)


# site de incones Iconfinder.com
#como chamar colcar uma imagem

icon_pesquisar = tk.PhotoImage(file="pesquisar.png")

icon_estoque = tk.PhotoImage(file="/Users/luan/Documents/Github/Iesgo-LTP1/UnidadeB/estoque.png")

icon_remover_estoque = tk.PhotoImage(file='remover.png')

icon_acrescentar_produto = tk.PhotoImage(file='acrescentar.png')

icon_cadastrar_produto = tk.PhotoImage(fila='cadastrar.png')

#criar botões com icones 
botao_pesquisar = tk.Button(janela_principal,image=icon_pesquisar,command=pesquisar_produtos)
botao_estoque =tk.Button(janela_principal,image= icon_estoque,command=checar_estoque)
botao_remover =tk.Button(janela_principal,image= icon_remover_estoque, command=remover_estoque)
botao_acrescentar= tk.Button(janela_principal,image= icon_acrescentar_produto,command= acrescentar_produto)
botao_cadastrar = tk.Button(janela_principal,image=icon_cadastrar_produto,command=cadastrar_produto )

#Função do loop 

janela_principal.mainloop()


#----------------------------------------- PARTE PARA DIVIDIR 
# criar outra pagina  pesquisar_produtos.py / arquivo.py com o nome de pesquisar produto

"""
import tkinter as tk

def abrir_janela_de_pesquisar_produto():
    janela_pesquisa = tk.Toplevel()
    janela_pesquisa.title('pesquisa e produtos ')
    # Funcionalidade da pesquisa

    #widgets 

    #... conteudo da função 


"""
