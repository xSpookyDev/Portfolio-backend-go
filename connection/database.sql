CREATE TABLE usuarios (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    apellido VARCHAR(100) NOT NULL,
    contrasena VARCHAR(200) NOT NULL,
    comentario VARCHAR(500) 
);
