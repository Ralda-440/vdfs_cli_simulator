#Version de node
FROM node:lts-alpine

# Crear directorio de trabajo
WORKDIR /app

# Copiar el archivo package.json
COPY package.json .

# Instalar dependencias
RUN npm install

# Copiar la aplicación
COPY . .

# Construir la aplicación
RUN npm run build

# Exponer el puerto 3000
EXPOSE 3000

# Variables de entorno PORT
ENV PORT=3000

# Comando para iniciar la aplicación
CMD ["npm", "run", "start"]