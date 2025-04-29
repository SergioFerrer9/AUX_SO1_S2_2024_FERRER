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

## Parte 2: Subir Imágenes a Harbor desde Docker

### Paso 1: Iniciar Sesión Docker en Harbor
```bash
mkdir -p ~/.docker/certs.d/35.222.163.140
cp harbor.crt ~/.docker/certs.d/35.222.163.140/ca.crt
docker login https://35.222.163.140
```
### Paso 2: Etiquetar y Subir Imagen
```bash
docker tag go-redis-app:1.0 35.222.163.140/library/go-redis-app:1.0
docker push 35.222.163.140/library/go-redis-app:1.0
```

## Parte 3: Kubernetes (GKE) con Aplicación Go y Redis

### Paso 1: Crear Secreto para Redis
```bash
apiVersion: v1
kind: Secret
metadata:
  name: redis-secret
type: Opaque
data:
  password: cGFzc3dvcmRfc29wZXNfMjAyNQ== # "password_sopes_2025" en base64
```

### Paso 2: Crear Secreto para Acceso a Harbor
```bash
gcloud auth configure-docker
kubectl create secret docker-registry harbor-registry-secret \
  --docker-server=35.222.163.140 \
  --docker-username=admin \
  --docker-password="<tu_password>" \
  --docker-email=example@example.com
```

### Paso 3: Deployment para la Aplicación Go-Redis
```bash
apiVersion: apps/v1
kind: Deployment
metadata:
  name: redis-go-app
spec:
  replicas: 1
  selector:
    matchLabels:
      app: redis-go-app
  template:
    metadata:
      labels:
        app: redis-go-app
    spec:
      imagePullSecrets:
      - name: harbor-registry-secret
      containers:
      - name: app
        image: 35.222.163.140/library/go-redis-app:1.0
        ports:
        - containerPort: 8080
        env:
        - name: REDIS_HOST
          value: "redis-service:6379"
        - name: REDIS_PASSWORD
          valueFrom:
            secretKeyRef:
              name: redis-secret
              key: password
```

### Paso 4: Verificar Estado del Pod
```bash
kubectl get pods
kubectl logs <nombre_del_pod>
```

## Parte 4: Alternativa si Certificados Autofirmados Fallan en GKE

### Subir Imagen a Google Artifact Registry
```bash
docker tag 35.222.163.140/library/go-redis-app:1.0 \
  us-central1-docker.pkg.dev/<PROJECT_ID>/<REPO_NAME>/go-redis-app:1.0
docker push us-central1-docker.pkg.dev/<PROJECT_ID>/<REPO_NAME>/go-redis-app:1.0
```

### Modificar Deployment para Usar Artifact Registry
```bash
containers:
- name: app
  image: us-central1-docker.pkg.dev/<PROJECT_ID>/<REPO_NAME>/go-redis-app:1.0
```

