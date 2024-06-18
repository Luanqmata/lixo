function fatorial(numero) {
    if (numero < 0) { //se o numero for menor que 0
        return "Fatorial não está definido para números negativos.";
    }
    if (numero === 0) {
        return 1;
    }
    return numero * fatorial(numero - 1);
}
//console.log(fatorial(5)); // Saída: 120

function maior_valor(lista,indice = 0,maior=Infinity){
    if(indice === lista.length){
        return `o Maior indice da listsa é ${indice}`
    }
    if(lista[indice] > maior){
        maior = lista[indice]
    }
    return maior_valor(lista,indice +1 , maior)
}
//console.log(maior_valor([1,2,3,4,5,6]))
function e_palindromo(palavra) {
    if (palavra.length <= 1) {
        return 'É um palíndromo ' + palavra; 
    }
    if (palavra[0] !== palavra[palavra.length - 1]) {
        return 'Não é um palíndromo '+ palavra; 
    }
    return e_palindromo(palavra.slice(1, palavra.length - 1));
}
//console.log(e_palindromo('arara'));
function soma_valores(lista, indice = 0, soma = 0) {
    if (indice === lista.length) { 
        return `A soma dos valores é ${soma}.`;
    }
    soma += lista[indice]; //*2;
    return soma_valores(lista, indice + 1, soma);
}
//console.log(soma_valores([1,5,6,7,8,2]));
function menorValor(lista, indice = 0, menor = Infinity) {
    console.log('incio')
    if (indice === lista.length) {
        //console.log(menor)
        return menor;
    }
    if (lista[indice] < menor) {
        //console.log(menor)
        menor = lista[indice];
    }
    //console.log(menor)
    return menorValor(lista, indice + 1, menor);
}
//let menor = menorValor([122, 33, 42, 12, 432]);
//console.log("O menor valor é: " + menor);
function reverter_String(palavra, indice = 0) {
    if (indice === palavra.length) {
        return '';
    } else {
        return reverter_String(palavra, indice + 1) + palavra[indice];
    }
}
//console.log(reverter_String('teste'));
function somaDigitos(numero, indice = 0) {
    let num_Str = numero.toString();

    if (indice === num_Str.length) {
        return 0;
    } else {
        let num_int = parseInt(num_Str[indice]);
        return num_int + somaDigitos(numero, indice + 1);
    }
}
//console.log(somaDigitos(1));
function ehPrimo(numero, divisor = 2) {
    if (numero <= 1) {
        return `o numero ${numero} não é primo.`;
    }
    if (divisor === numero) {
        return `O numero ${numero} é primo.`;
    }
    if (numero % divisor === 0) {
        return  `o numero ${numero} não é primo.`;
    }
    return ehPrimo(numero, divisor + 1);
}
//console.log(ehPrimo(7)); // Saída: true (7 é primo)
//luan
                                                                            function soma_elementos(array,soma){
                                                                                soma = 0;
                                                                                array.map(num=>soma+=num)
                                                                                return soma
                                                                            }
                                                                            //mensagem(soma_elementos([1,2,3]))

//halbert
                                                                            function somaElementos(lista,soma){
                                                                                if(lista.length==1){
                                                                                    soma += lista.pop();
                                                                                    return soma;
                                                                                } else {
                                                                                    soma += lista.pop();
                                                                                    return somaElementos(lista,soma)
                                                                                }
                                                                            }
                                                                                    //var total = somaElementos([12,34,54],0)
                                                                                    //mensagem(total)

                                                                                    //sandir
                                                                                    function SomaElementos(lista,indice = 0){
                                                                                        if(indice===lista.length){
                                                                                            return 0
                                                                                        }
                                                                                        return lista[indice]+SomaElementos(lista,indice+1)
                                                                                    }
                                                                                    //mensagem(SomaElementos([1,2,3]))
                                                                                    
                                                                                    //----------------------------------
        function dobrarValores(array){
            return array.map(num=>num*2);
        }
        //mensagem(dobrarValores([2,4,6]))
