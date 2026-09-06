# Extraer Master Boot Record

Este software extrae los primeros **512** bytes del disco duro
El Master Boot Record se encuentra en los primeros **512** bytes del disco duro, Así que decidí extraerlo en un archivo binario como practica.

# Requisitos:
- Go
- Make

# Compilar:

`make`

# Como usarlo?
Windows: 
`mbrextractor-go.exe \\.\PhysicalDrive0 mbr.bin`

Linux:
`sudo ./mbrextractor-go /dev/sda mbr.bin`


> [!IMPORTANT]
> 
> Asegurate de ser root/admin!
