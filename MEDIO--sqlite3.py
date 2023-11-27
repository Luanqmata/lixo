#criar aplicação capaz de armazenar dados de usuarios (nome ,email,idade) via terminal . crie um menu via terminal ,permitindo add 
#remover via chave primaria e vizualizar a lista de user 
import sqlite3

def criar_tabela():
    # Conectar ao banco de dados ou criar se não existir
    connection = sqlite3.connect('dados_usuarios.db')
    
    # Abrir o cursor para executar comandos SQL
    cursor = connection.cursor()
    
    # Criar a tabela se não existir
    cursor.execute("""
        CREATE TABLE IF NOT EXISTS usuarios (
        id INTEGER PRIMARY KEY AUTOINCREMENT, 
        nome TEXT NOT NULL,
        email TEXT NOT NULL,
        idade INTEGER NOT NULL
        )            
    """)
    
    # Commit para salvar as alterações no banco
    connection.commit()
    
    # Fechar a conexão com o banco de dados
    connection.close()

def vizualizar_Lista():
    # Corrected the variable name from conncetion to connection
    connection = sqlite3.connect('dados_usuarios.db')
    cursor = connection.cursor()
    
    # Selecionar o registro da tabela
    cursor.execute("SELECT * FROM usuarios")
    rows = cursor.fetchall()
    
    print('\nLista de Usuários:')
    for row in rows:
        print(row)
        
    connection.close()

def adicionar_membro():
    nome = input('digite o nome do usuario: ')
    email = input('digite o email do usuario')
    idade = int(input('digite sua idade: '))

    connection = sqlite3.connect('dados_usuarios.db')
    cursor = connection.cursor()
    
    #inserir novo registro 
    cursor.execute('INSERT INTO usuarios(nome,email,idade)VALUES (?,?,?)',(nome,email,idade))
    
    connection.commit()
    connection.close()
    print(f'\nO usuario {nome} foi registrado com sucesso')

def remover_usuario():
    id_remover = int(input('Digite o id do usuario a ser removido: '))

    connection = sqlite3.connect('dados_usuarios.db')
    cursor = connection.cursor()
    #remover o registro do banco com base no ID 
    cursor.execute('DELETE FROM usuarios WHERE id = ?' , (id_remover,))
    connection.commit()
    connection.close()
    print(f'\n o usuario ID:{id_remover}  foi removido.')
    
if __name__=="__main__":
    criar_tabela()
    
    while True:
        print('\n escolha uma opcão: ')
        print('Opção 1: vizualizar lista')
        print('Opção 2: ADD membro')
        print('Opção 3: Remover Membro')
        print('Opção 4: sair.')
    
        user_escolha = input('Opcão: ')
        
        if user_escolha == '1':
            vizualizar_Lista()
        elif user_escolha == '2':
            adicionar_membro()
        elif user_escolha == '3':
            remover_usuario()
        elif user_escolha == '4':
            print('Saindo...')
            break
        else:
            print('Escolha dentre as opções do MENU.')

            
