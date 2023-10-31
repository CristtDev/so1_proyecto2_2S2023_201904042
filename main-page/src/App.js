import React, { useState, useEffect } from 'react';
import './App.css';
import socket from "./socket/Socket.ts";

function App() {
  const [data, setData] = useState([]); // Estado para almacenar los datos de estudiantes

  useEffect(() => {
    socket.emit("estudiantes", "Hola");
    socket.on("estudiantes", (respuesta) => {
      console.log(respuesta);
      setData(respuesta); // Almacena la respuesta en el estado 'data'
    });
  }, []);

  return (
    <div className="App">
      <h1>Lista de Estudiantes</h1>
      <header className="App-header">
        <div className='tableContainer'>
          <table>
            <thead>
              <tr>
                <th>Carnet</th>
                <th>Nombre</th>
                <th>Curso</th>
                <th>Nota</th>
                <th>Semestre</th>
                <th>Año</th>
              </tr>
            </thead>
            <tbody>
              {data.map((item, index) => (
                <tr key={index}>
                  <td>{item.carnet}</td>
                  <td>{item.nombre}</td>
                  <td>{item.curso}</td>
                  <td>{item.nota}</td>
                  <td>{item.semestre}</td>
                  <td>{item.year}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </header>
    </div>
  );
}

export default App;
