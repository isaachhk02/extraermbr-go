# MBR Extractor

This software extracts the first **512** bytes from hard disk.
The Master Boot Record is stored in the first 512 bytes; therefore, this utility extracts those first 512 bytes into a binary file.

# Requirements:
- Go
- Make

# Build:

`make`

# How to use?
Windows: 
`mbrextractor-go.exe \\.\PhysicalDrive0 mbr.bin`

Linux:
`sudo ./mbrextractor-go /dev/sda mbr.bin`


> [!IMPORTANT]
> 
> MAKE SURE RUN THIS SOFTWARE AS ROOT/ADMIN!
