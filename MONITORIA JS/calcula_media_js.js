function mensagem(text){
    document.write(text);
    document.write('<br>')
}
    _notas_do_aluno_ = []
    _soma_notas_ = 0
    _qnt_materias_ = 0
    materias = [' PORTUGUES ',' MATEMATICA ',' HISTORIA ',' GEOGRAFIA ',' EDUCAÇÃO FISICA ']

function solicitar_notas(){
        for(let i =0;i<materias.length;i++){
            _nota_aluno_ = Number(prompt('Digite a sua nota na materia de' + materias[i] + ':'))
            if (_nota_aluno_<0){
                mensagem('QUANDO VC DIGITA UM NUMERO NEGATIVO VOCE DECIDE PARAR . <br> ')
                break
            }   else if (_nota_aluno_>0 && _nota_aluno_ <11){
                _notas_do_aluno_.push(materias[i] + _nota_aluno_)
                _soma_notas_ += _nota_aluno_
                _qnt_materias_ += 1 
            } else {
                alert('Nota valida somente de 1 a 10.')
            }
        }
    return ' E SUAS NOTAS registradas foram: ' + _notas_do_aluno_ + ' <br> a soma das notas foi:' + _soma_notas_ + '<br> E foram registrados ' + _qnt_materias_ + ' materias.'
}
//mensagem(solicitar_notas())

function calcularMedia(SomaNotas,QntNotas){
    media = SomaNotas / QntNotas
    mensagem('A media das suas notas é de : ' + Math.round(media))
    return Math.round(media)
}
//calcularMedia(_soma_notas_,_qnt_materias_)

function classificarDesempenho(media){
    if (media >= 9.0){
        mensagem('EXELENTE!');
    }  else if (media >= 8.0 && media < 9.0) {
        mensagem('MUITO BOM!!')
    } else if (media >= 7.0 && media <8.0){
        mensagem('BOM')
    } else if (media >= 7.0 && media <6.0){
        mensagem('SATISFATORIO')
    } else {
        mensagem('MUITO RUIM!!')
    }
}
//classificarDesempenho(media);

function remover_materia(){
    let materia = Number(prompt('Qual materia i(INDICE) você deseja remover?\n' + materias)) -1
    if (materia >= 0 && materia < materias.length){
        materias.splice(materia,1)
        alert('Materia removida com sucesso!')
        mensagem(materias)
        } else {
            mensagem('INDICE INVALIDO.')
    }
}
//remover_materia();

function adicionar_materia(){
    materia = prompt('Digite o nome da materia que você deseja adicionar: ')
    materias.push(materia);
    //mensagem(materias)
}
//adicionar_materia();
function menu(){
    let _user_escolha_ = Number(prompt('Digite o Número de acordo com o MENU: \n 1. CALCULAR MÉDIA \n 2. ADICIONAR MATÉRIA \n 3. REMOVER MATÉRIA (por índice) \n 4. PRINTAR LISTA DE MATÉRIAS'));
    if (_user_escolha_ == 1) {
        mensagem(solicitar_notas()); 
        let media = calcularMedia(_soma_notas_, _qnt_materias_);
        classificarDesempenho(media);
        //menu();
    } else if (_user_escolha_ == 2) {
        adicionar_materia();
        menu();
    } else if (_user_escolha_ == 3) {
        remover_materia();
        menu();
    } else if (_user_escolha_ == 4) {
        alert(materias);
        menu();
    } else {
        alert('DIGITE UM NUMERO ENTRE 1 E 4.');
        menu();
    }
}
menu()