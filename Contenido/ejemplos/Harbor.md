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

### Paso 2: Generar certificados autofirmados
```bash
mkdir ~/harbor-certs && cd ~/harbor-certs
openssl genrsa -out harbor.key 4096
openssl req -new -x509 -key harbor.key -out harbor.crt -days 365 \
  -subj "/C=GT/ST=Guatemala/L=Guatemala City/O=Harbor GCP/OU=IT/CN=35.222.163.140"
```

### Paso 3: Configurar Harbor con HTTPS
- Modificar harbor.yml:
```bash
hostname: 35.222.163.140
https:
  port: 443
  certificate: /data/secret/cert/harbor.crt
  private_key: /data/secret/cert/harbor.key
```
- Copiar los certificados al directorio configurado:
```bash
sudo mkdir -p /data/secret/cert
sudo cp harbor.crt /data/secret/cert/
sudo cp harbor.key /data/secret/cert/
```

### Paso 4: Instalar Harbor
```bash
sudo ./install.sh
```
- Verifica acceso en navegador: https://35.222.163.140


