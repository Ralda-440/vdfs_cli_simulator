#Version de node
FROM node:latest

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

# Comando para iniciar la aplicación
CMD ["npm", "run", "start"]

# Exponer el puerto 3000
EXPOSE 3000