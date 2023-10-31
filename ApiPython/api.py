from flask import Flask, request, jsonify
import redis
import json
import mysql.connector

app = Flask(__name__)
redisClient = redis.StrictRedis(host="redis", port=6379, db=15)

# Configuración para la base de datos de Google Cloud SQL
db_config = {
    "host": "34.122.19.115",
    "user": "root",
    "password": "test123$",
    "database": "proyecto2"
}

@app.route('/agregarEstudiante', methods=['POST'])
def agregar_estudiante():
    try:
        data = request.get_json()
        carnet = data.get("carnet")
        estudiante_json = json.dumps(data)

        # Insertar en Redis
        redisClient.hset("estudiantes", carnet, estudiante_json)

        # Insertar en Google Cloud SQL
        connection = mysql.connector.connect(**db_config)
        cursor = connection.cursor()

        # Insertar en la tabla ALUMNO
        insert_alumno_query = "INSERT INTO ALUMNO (carnet, nombre) VALUES (%s, %s)"
        cursor.execute(insert_alumno_query, (carnet, data["nombre"]))

        # Commit para confirmar la inserción en la tabla ALUMNO
        connection.commit()

        # Insertar en la tabla CALIFICACION
        insert_calificacion_query = "INSERT INTO CALIFICACION (nota, anio, carnet, codigo, idSemestre) VALUES (%s, %s, %s, %s, %s)"
        cursor.execute(insert_calificacion_query, (data["nota"], data["year"], carnet, data["curso"], data["semestre"]))

        # Commit para confirmar la inserción en la tabla CALIFICACION
        connection.commit()

        cursor.close()
        connection.close()

        return "Estudiante agregado con éxito", 200
    except Exception as e:
        return str(e), 500

@app.route('/obtenerEstudiantes', methods=['GET'])
def obtener_estudiantes():
    try:
        estudiantes = redisClient.hgetall("estudiantes")
        estudiantes_list = []
        for carnet, estudiante_json in estudiantes.items():
            estudiante_data = json.loads(estudiante_json)
            estudiantes_list.append(estudiante_data)
        return jsonify(estudiantes_list), 200
    except Exception as e:
        return str(e), 500

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=3002)
