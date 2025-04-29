# Manual de Instalación y Despliegue de Harbor con HTTPS y Aplicación en Kubernetes

## Parte 1: Instalación de Harbor con HTTPS (Certificado Autofirmado)

### Requisitos Previos
- Docker y Docker Compose instalados
- Servidor con Ubuntu 20.04 o superior
- IP pública (ej. 35.222.163.140)

### Paso 1: Descarga e instalación de Harbor
```bash
wget https://github.com/goharbor/harbor/releases/download/v2.10.0/harbor-online-installer-v2.10.0.tgz
tar -xvzf harbor-online-installer-v2.10.0.tgz
cd harbor
```
