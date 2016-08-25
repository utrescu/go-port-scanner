Escanner de ports oberts
=============================
Es tracta d'un programa que comprova quantes màquines en una xarxa o xarxes tenen un determinat port obert

Execució
----------------
La comanda bàsica consisteix en especificar la xarxa a escannejar en format CIDR: 

    $ go run list-ip.go 192.168.1.0/24
    Hosts  256
    Obert 192.168.1.1
    [192.168.1.2]

Per defecte escanneja el port 22 però amb el paràmetre *-p* o *-port* es pot especificar un altre port. Per exemple per veure quines màquines de la xarxa 192.168.9.0 tenen el port 80 (http) obert:

    $ go run list-ip.go -p 80 192.168.1.0/24 
    Hosts  256
    Obert 192.168.1.2
    Obert 192.168.1.1
    [192.168.1.2 192.168.1.1]

    Fet en 1.023404099s

Es pot fer posant les xarxes que calgui com a paràmetres: 

    $ go run list-ip.go -p 80 192.168.0.0/24 192.168.1.0/24
    Hosts  508
    Obert 192.168.0.2
    [192.168.0.2]

    Fet en 1.040980273s

El temps d'espera per establir connexions s'ha definit per defecte a 1 segon però és pot canviar per un altre valor amb *-timeout*. Per exemple 200 milisegons: 

    $ go run list-ip.go -p 22 -timeout 200ms 192.168.1.0/24
    Hosts  254
    Obert 192.168.1.2
    [192.168.1.2]

    Fet en 223.323989ms

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

