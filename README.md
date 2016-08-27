Escanner de ports oberts
=============================
Es tracta d'un programa que comprova quantes màquines en una xarxa o xarxes tenen un determinat port obert

Execució
----------------
La comanda bàsica consisteix en especificar la xarxa a escannejar en format CIDR: 

    $ go run list-ip.go 192.168.1.0/24
    Màquines
    ---------------
    192.168.1.2

    durada: 1.028944403s

També funciona posant IP:

    $ go run list-ip.go 192.168.1.2
    Màquines
    ---------------
    192.168.1.2

    durada: 830.384µs

Es pot fer posant les xarxes o IP que calgui com a paràmetres: 

    $ go run list-ip.go 192.168.0.0/24 192.168.1.0/24 192.168.8.1
    Màquines
    ---------------
    192.168.0.2
    192.168.1.2

    durada: 1.039590773s

### Definir el port 

Per defecte escanneja el port 22 però amb el paràmetre *-p* o *-port* es pot especificar un altre port. Per exemple per veure quines màquines de la xarxa 192.168.9.0 tenen el port 80 (http) obert:

    $ go run list-ip.go -p 80 192.168.1.0/24 
    Màquines
    ---------------
    192.168.1.1
    192.168.1.2

    durada: 1.023340366s

### Definir el temps d'espera

El temps d'espera per establir connexions s'ha definit per defecte a 1 segon però és pot canviar per un altre valor amb *-timeout*. Per exemple 200 milisegons: 

    $ go run list-ip.go -timeout 250ms  -p 80  192.168.8.0/24                             
    Màquines 
    ---------------    
    192.168.8.2    
    192.168.8.1    
    
    durada: 274.959328ms


Compilar
----------
En Go els programes es poden compilar per generar executables nadius del sistema:

    $ go build list-ip.go 
    $ ls
    list-ip  list-ip.go  README.md
    $ ./list-ip 
    Hosts  254
    Obert 192.168.88.1
    [192.168.88.1]

    Fet en 1.003207812s

