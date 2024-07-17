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
