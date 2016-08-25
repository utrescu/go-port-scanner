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

Per defecte escanneja el port 22 però amb el paràmetre *-p* o *-port* es pot especificar un altre port. Per exemple per veure quines màquines de la xarxa 192.168.9.0 tenen el port 80 (http) obert:

    $ go run list-ip.go -p 80 192.168.1.0/24 
    Màquines
    ---------------
    192.168.1.1
    192.168.1.2

    durada: 1.023340366s

Es pot fer posant les xarxes que calgui com a paràmetres: 

    $ go run list-ip.go -p 80 192.168.0.0/24 192.168.1.0/24
    Màquines
    ---------------
    192.168.0.2
    192.168.1.2

    durada: 1.039590773s

El temps d'espera per establir connexions s'ha definit per defecte a 1 segon però és pot canviar per un altre valor amb *-timeout*. Per exemple 200 milisegons: 

    $ go run list-ip.go -timeout 250ms  -p 80  192.168.8.0/24                             
    Màquines 
    ---------------    
    192.168.8.2    
    192.168.8.1    
    
    durada: 274.959328ms

El programa l'he fet per generar inventaris per Ansible i per això si en la generació s'hi posa el paràmetre 'tag' es genera un fitxer d'inventari per Ansible amb el mateix nom de l'etiqueta: 

    $ go run list-ip.go -timeout 250ms -tag alumnes 192.168.50.0/24
    [alumnes]
    192.168.50.12
    192.168.50.16
    192.168.50.23

En aquest cas a més de generar l'inventari per pantalla també crea un fitxer amb el mateix nom:

    $ ls                                   
    alumnes  list-ip.go  README.md
    $ cat alumnes
    [alumnes]
    192.168.50.12
    192.168.50.16
    192.168.50.23

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

