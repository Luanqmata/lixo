import sqlite3

# Criar a pasta 
connection = sqlite3.connect('novo_db.db')

cursor = connection.cursor()

# Criar uma tabela 
cursor.execute(
    'CREATE TABLE IF NOT EXISTS usuarios(id INT PRIMARY KEY, nome TEXT, idade INT)'
)

"""
# Inserir dados na tabela 
def inserir_dados(nome, idade):
    cursor.execute(
        f"INSERT INTO usuarios (nome, idade) VALUES ('{nome}', {idade})"
    )
"""
# Inserir dados na tabela 
cursor.execute(
    "INSERT INTO usuarios (nome,idade) VALUES ('seucu', 32)"
)

#fazer consulta no banco de DADOS 

cursor.execute(
    'SELECT * FROM usuarios ' #varrer todos os usuarios 
)
rows =  cursor.fetchall() # puxar todo mundo 
# forr lop para armazenar 
for row in rows: # a linha nas linhas é igual a printt linha
    print(row)

connection.commit()
connection.close()
