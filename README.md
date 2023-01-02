## Estudando GoLang - Elevando Stack
<p align="center">
  <a href="https://go.dev/" target="blank"><img src="https://th.bing.com/th/id/OIP.iwrDhp5zhgKs69Ad4WJ7xAHaEK?w=312&h=180&c=7&r=0&o=5&pid=1.7" width="250" alt="Nest Logo" /></a>
</p>

---
<br>

### Conceitos fundamentais estudados
- Tipos
- Manipulação de pacotes
- Fundamentos da Programação
- Controladores
- API Golang
- Goroutine
- Channel Buffer
- Slices - Maps - Arrays

## Como instalar o projeto - terminal
```bash
  -> open the terminal 

  $git clone https://github.com/fymorGod/curso-golang
  
  $cd curso-golang

  -> Open this folder on terminal

  $code .

  -> open terminal in VS Code
  $go run *name_do_arquivo

  OR

  to start code 
  -> command with keys Ctrl + Alt + N 
```
## Concorrência vs Paralelismo
<p width="500">
    
    Paralelismo - é quando temos processamentos simultâneos, ou seja, processos ativos no mesmo instante. Trabalhando de forma paralela, ao mesmo tempo.

    Concorrência - é a capacidade de administrar multiplas tarefas, podendo ocorrer em um único processador (Core/Núcleo).
  
    É possível que a concorrência seja melhor que o parelelismo, pois este exige muito mais do SO e do hardware.

---
</p>

## Threads - Conceitos e aplicações
<br>
<p align="center">
  <a href="https://go.dev/" target="blank"><img src="https://miro.medium.com/max/1400/0*CFiDiTCqF_hmEvdX.png" width="300" alt="Nest Logo" /></a>
</p>

<p>
    Uma Thread seria uma função capaz de destribuir multiplas funções para serem executadas simultaneamente.
    <h3>Um bom exemplo seria:</h3>
    
    Há um casal em um shopping que foram para assistir um filme. 
    Nesse contexto, eles seguem um fluxo de atividades(TASKS) antes de iniciar a sessão de cinema.

  ### Tasks:
  - Comprar o Lanche
  - Comprar os ingressos

  ### Resolução:
    Com isso, cria-se uma Thread -> para separar as tarefas entre eles.
    exemplo:
    - João vai comprar o lanche
    - Maria vai comprar os ingressos
</p>