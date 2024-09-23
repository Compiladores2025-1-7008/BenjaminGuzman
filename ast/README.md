# AST

Abstract Syntax Tree

## Programa en Go

```shell
go build
./ast main.go
```

Copiar y pegar el output en https://dreampuf.github.io/GraphvizOnline/

[Ejemplo](https://dreampuf.github.io/GraphvizOnline/#digraph%20G%20%7B%0Afuncion_4806%20%5Blabel%3D%22funcion%22%2C%20style%3Dfilled%2C%20fontcolor%3Dwhite%2C%20fillcolor%3Dpalevioletred4%2C%20xlabel%3D%22doc%0A%22%5D%3B%0An_4a3b%20%5Blabel%3D%22n%22%2C%20fontsize%3D8%2C%20shape%3Dinvtriangle%2C%20width%3D0.5%2C%20height%3D0.5%5D%3B%0Aint_63de%20%5Blabel%3D%22int%22%2C%20fontsize%3D4%2C%20width%3D0.3%2C%20height%3D0.3%5D%3B%0Aint_63de%20-%3E%20n_4a3b%20%5Bdir%3Dnone%5D%3B%0An_4a3b%20-%3E%20funcion_4806%20%5Bdir%3Dnone%5D%3B%0Aint64_546a%20%5Blabel%3D%22int64%22%2C%20shape%3Dtriangle%2C%20fontsize%3D4%2C%20width%3D0.5%2C%20height%3D0.5%5D%3B%0Aint64_546a%20-%3E%20funcion_4806%20%5Bdir%3Dnone%5D%3B%0Aassignment_8d96%20%5Blabel%3D%22%3D%22%5D%3B%0Aid_bb49%20%5Blabel%3D%22one%22%5D%3B%0Aassignment_8d96%20-%3E%20id_bb49%3B%0Alit_1_93f1%20%5Blabel%3D%221%22%5D%3B%0Aassignment_8d96%20-%3E%20lit_1_93f1%3B%0Afuncion_4806%20-%3E%20assignment_8d96%3B%0Aassignment_a047%20%5Blabel%3D%22%3D%22%5D%3B%0Aid_b6eb%20%5Blabel%3D%22n_plus_one%22%5D%3B%0Aassignment_a047%20-%3E%20id_b6eb%3B%0Abinary_operator_b8a7%20%5Blabel%3D%22%2B%22%5D%3B%0Aid_4232%20%5Blabel%3D%22n%22%5D%3B%0Abinary_operator_b8a7%20-%3E%20id_4232%3B%0Aid_a4d5%20%5Blabel%3D%22one%22%5D%3B%0Abinary_operator_b8a7%20-%3E%20id_a4d5%3B%0Aassignment_a047%20-%3E%20binary_operator_b8a7%3B%0Afuncion_4806%20-%3E%20assignment_a047%3B%0Aassignment_61ce%20%5Blabel%3D%22%3D%22%5D%3B%0Aid_7a25%20%5Blabel%3D%22gauss_sum%22%5D%3B%0Aassignment_61ce%20-%3E%20id_7a25%3B%0Abinary_operator_e245%20%5Blabel%3D%22%2F%22%5D%3B%0Abinary_operator_1a9b%20%5Blabel%3D%22*%22%5D%3B%0Aid_248c%20%5Blabel%3D%22n%22%5D%3B%0Abinary_operator_1a9b%20-%3E%20id_248c%3B%0Aid_8c0e%20%5Blabel%3D%22n_plus_one%22%5D%3B%0Abinary_operator_1a9b%20-%3E%20id_8c0e%3B%0Abinary_operator_e245%20-%3E%20binary_operator_1a9b%3B%0Alit_2_4579%20%5Blabel%3D%222%22%5D%3B%0Abinary_operator_e245%20-%3E%20lit_2_4579%3B%0Aassignment_61ce%20-%3E%20binary_operator_e245%3B%0Afuncion_4806%20-%3E%20assignment_61ce%3B%0Areturn_0f7d%20%5Blabel%3D%22return%22%5D%3B%0Aint64_abaa%20%5Blabel%3D%22call%20int64%22%5D%3B%0Aid_95cf%20%5Blabel%3D%22gauss_sum%22%5D%3B%0Aint64_abaa%20-%3E%20id_95cf%3B%0Areturn_0f7d%20-%3E%20int64_abaa%3B%0Afuncion_4806%20-%3E%20return_0f7d%3B%0A%7D%0A)

### Para qué me sirve??

- Entender porqué un compilador arroja ciertos errores
- Self mutating code (hacking)
- Escribir `dot` bonito... Podríamos construir un AST de dot, en lugar de usar printf's...
- Ejercicio mental

### Notas y preguntas

- La mayoría de lenguajes de programación tienen un módulo de AST
- **Todos los lenguajes de programación necesitan un AST?** No, e.g., assembly no necesita, ¿porqué? ¿cuándo es conveniente?
- Gran forma de aprender: Leer código de los demás, leer el código del módulo de AST de Go, Java, lo que sea. e.g., 
  qué pasa cuando en un Hashmap hay colisión de valores? Se crea un arreglo y se itera, pero...
- Gran forma de aprender: Usar un debugger.
- Pro tip: Con el correo institucional pueden sacar licencia de JetBrains

## Misceláneo

### Programas de investigación

- Programa delfin: https://www.programadelfin.org.mx/
- ENLACE: https://resilientmaterials.ucsd.edu/ENLACE
- El que tú quieras
