# Usa la imagen base de Node.js con la versión deseada
FROM node:14

# Directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el archivo package.json y package-lock.json al directorio de trabajo
COPY package*.json ./

# Instala las dependencias
RUN npm install

# Copia el resto de los archivos de la aplicación al directorio de trabajo
COPY . .

# Expone el puerto en el que la aplicación de React estará escuchando
EXPOSE 3000

# Comando para iniciar la aplicación (ajústalo según tus necesidades)
CMD ["npm", "start"]
