const express = require('express');
const app = express();
const http = require('http');
const server = http.createServer(app);
const cors = require('cors');
const Redis = require('ioredis');

const redis = new Redis({
    host: 'redis',
    port: 6379,
    db: 15,
});

app.use(cors());
app.use(express.urlencoded({
    extended: true
}));
app.use(express.json());

const { Server } = require('socket.io');
const io = new Server(server, {
    cors: {
        origin: "*"
    }
});

io.on('connection', (socket) => {
    console.log("Se conectó un cliente");

    const emitEstudiantes = async () => {
        try {
            const estudiantes = await redis.hgetall('estudiantes');
            const estudiantesList = Object.values(estudiantes).map(JSON.parse);

            console.log("Emitiendo");
            io.emit("estudiantes", estudiantesList);
        } catch (error) {
            console.error('Error al obtener los estudiantes desde Redis', error);
            io.emit("estudiantes", 'Error al obtener los estudiantes desde Redis');
        }
    }

    // Realizar la primera consulta al conectarse el cliente
    emitEstudiantes();

    // Establecer un intervalo para consultar Redis cada cierto tiempo (ejemplo: cada 10 segundos)
    const interval = setInterval(emitEstudiantes, 1500);

    // Manejar la desconexión del cliente
    socket.on("disconnect", () => {
        console.log("Cliente desconectado");
        clearInterval(interval); // Detener el intervalo cuando el cliente se desconecta
    });
});

server.listen(4000, () => {
    console.log("Server on port 4000");
});
