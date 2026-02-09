# Projet Net-Cat :

## NET-CAT
Ce projet consiste a réecrire Net-Cat
#### c'est quoi Net-Cat :
   - est un outil réseau en ligne de commande trés simple permet d'envoyer et de recevoir des données brutes entre 2 machines via le réseau, en utilisant TCP ou UDP.
   - est considérer comme un tuyau qui transporte les données en octets
   - joue 2 roles => Mode Serveur, Mode Client 
   - Net Cat connait juste une connexion et un flux de données (ouvre une socket + lire/ecrire ds un flux de données )

## Réseaux Informatique
### Réseau :
- Un réseau est une infrastructure qui permet l’échange de données entre machines.(machines(PC,...)+ Moyen de communication(cable, wifi,...) + des regles de communication(Protocoles) ).

### LAN vs WAN
#### LAN (Local Area Network):
- réseau local
#### WAN (Wide Area Network):
- réseau distant (composée de plusieurs LAN interconnectés )

### Adresse IP (Internet Protocol):
- Une IP sert a identifier un noeud(machine, ..) ds un réseau afin d'envoyer et recevoir des données . (Changement d réseau => changement d IP) => chaqu'un son adresse IP unique (réseau+machine)

#### IP Version 4 (IPV4)
- Cette IP a 32 bits au total , est composée de 4 nombres  chaque nombre est compris entre 0 et 225
pourquoi entre 0 et 255
l'ordinateur compris que du binaire (0, 1), chaque partie du IP utilise 8 bits
8 bits + 8 bits + 8 bits + 8 bits = 32 bits
EX: 
    192.168.1.25
        192....1 => Identifier le réseau 
        .25 => identifier la machine 
    RQ:
        a chaque reseau 2 adresse sont reservees 
                192.168.1.0 => identifie le reseau lui meme
                192.168.1.255 => Adresse de broadcast pour envpyer des un message a toutes les machines du réseau

     IP pub vs IP prv
- Identification de réseau local =>  a l'interieur on a des adresse IP privées 

##### c'est quoi un routeur:
- celui qui relie ton LAN a a INternet (WAN)
    Routeur a 2 faces 
        LAN:
           Il attribue des IP privee a chaque noeud ds reseau local

        WAN  
            Le routeur a une IP publique fournie par FAI 
            le routeur transforme IP prv en publique grace au NAT (network addresss TRanslation) , Il garde en mémoire quel appareil ds le LAN fait la demande 
#### IP dynamique vs IP statique 
- IPS => Ne change jamais , configuree manuelle ou par serveur pr exemple localhost est tjr statique 127.0.0.1

- IPD => Change a chaque connexion (securite , facile a gerer...)


#### IP Version 6 (IPV6)
- 128bits 
- en base 16 HEXADECIMAL
- 8 groupes de 4 caracteres =>  2001:0db8:85a3:0000:0000:8a2e:0370:7334

RQ: 
- la  la difference entre IPV4 est necessite plus de NAT qui va generer une IP publique qui identifie le reseau , IPV6 donne a chaque machine son IP unique  
-  Pour IPV4 ne securise rien par defaut , on doit ajouter des protocoles externes ex: IPsec , cepandant IPV6  la sécurisation est deja integrée IPSEC permet de crypter les données + Authentification

### PORT:
- IP sert a identifier une machine sur un réseau, mais une machine peut avoir plusieurs  services  donc un port est un numero qui identifie un service ou une application sur la machine .
    EX => 192.168.1.25:80 => 80 est un service web 

    * Port est un nombre entre 0 et 65535
        DNS = 53
        SSH = 22
        HTTPS = 43 , HTTP = 80
    
    * OS  fournit ou lancer des services chaque service a un port  afin de communiquer donc c'est le point d'entree  pour que le service 
    ex
        192.168.1.10:80 => cad , une machine écoute les connexions entrantes sur le port 80
    #### Ecoute sur un PORT ??? 
    -  UN SERVICE ou UN PROGRAMME demande a OS je veux recevoir tous les données qui arrivent sur ce port 
    - OS creer un canal virtuel pour ce port
    - les paquets reseau arrive au IP:PORT est redirigé vers ce service.

ON A le lancement d'un service (ex ; serveur web )
