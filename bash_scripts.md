# 2. Creación de Scripts en Bash

## 1.2.1. Introducción a Bash

Bash es el acrónimo de "Bourne Again SHell". Es el shell por defecto en muchas distribuciones de Unix y Linux, y es un lenguaje de comandos que permite a los usuarios interactuar con el sistema operativo para ejecutar comandos, realizar tareas de administración y automatizar procesos.

## 1.2.2. Estructura básica de un script

Un script de Bash es simplemente un archivo de texto que contiene una serie de comandos que se ejecutan secuencialmente. La estructura básica incluye:

- **Shebang**: Indica al sistema qué intérprete usar. Ejemplo: `#!/bin/bash`
- **Comentarios**: Empiezan con `#` y no son ejecutados. Ejemplo: `# Este es un comentario`
- **Comandos**: Cualquier comando de Bash. Ejemplo: `echo "Hola, mundo"`

```bash
#!/bin/bash
# Este es un script básico de Bash
echo "Hola, mundo"
```

## 1.2.3. Variables y control de flujo
En Bash, puedes definir variables y usar estructuras de control de flujo como bucles y condicionales.

- **Variables**: Las variables en Bash se definen simplemente asignando un valor a un nombre. Para acceder al valor de una variable, se utiliza el símbolo `$` seguido del nombre de la variable.

```bash
#!/bin/bash
# Definir una variable
nombre="Mundo"
echo "Hola, $nombre"
```

- ** Condicionales**: Las estructuras condicionales en Bash permiten ejecutar comandos basados en ciertas condiciones. La estructura básica de un condicional `if` es la siguiente:

```bash
#!/bin/bash
# Condicional if-else
if [ "$nombre" = "Mundo" ]; then
  echo "Hola, Mundo"
else
  echo "Hola, desconocido"
fi
```
