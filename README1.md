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
- IPS => Ne change jamais , configureeclient manuelle ou par serveur pr exemple localhost est tjr statique 127.0.0.1

- IPD => Change a chaque connexion (securite , facile a gerer...)


#### IP Version 6 (IPV6)
- 128bits 
- en base 16 HEXADECIMALclient
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

ON a le lancement d'un service (ex ; serveur web ) , ce service va demander au OS d'ouvrir un port 
    OS reserve un port + creation d'une table interne(une structure qui relie un port 80 au programme(service)) + ecouter les paquets (chaque paquet arrivee sur ce port il sera envoye au service lie )

    C/C  Un canal  est en realite une entree ds la table de l'os qui relie Port -> Service -> Socket

##### Socket 
    UN SOCKET est un point d communication entre programme et réseau 
    (IP, PORT) C'est une structure interne qui relie IP => PORT  => Programme (service )
    L’OS crée une structure interne (socket) (PORT, IP, SERVICE) mettre ce socket en état "Écoute" ,point de connexion réel
    - IL sert juste a accepter la connexion
    - Chaque socket utilise une  Mémoire (structure Socket) + descripteur systeme (chaque socket = fichier interne au  kernel  )

###### Comment ?? la creation de ce socket ? 
    * Programme est un fichier sur disque 
    * Execution
    * Kernel => Creer un processus + allour de la mémoire + initialise une tablde de file descriptors
    * Programme demande une socket TCP
    * kernel :RQ:
            -Socket()
            -Creation de la structure plus haut niveau Couche BSD=>struct socket *sock;
                ROLE: c'est un adaptateur Programme  <->  struct socket  <->  TCP (struct sock / tcp_sock)
                composants :
                    famille d'adresse (IPV4 ou IPV6 ou ) + Type d socket (SOCK_STREM(TCP) ou SOCK_DGRAM(UDP) + SOCET_RAW(paquet brut)) + STATE de socket(CLOSED...)+ sk = null
            
            -Creation de struct sock => c'es t ka structure réseau générique 
                Composants: files d'attente paquets + buffers(envoie et réception) + pointeur vers protocoles 


            -Creation de struct tcp_sock: l'objet qui fait fonctionner le protocole TCP 
            
            -liason entre struct socket <=> struct sock  => double le lien 
                le programme part du FD => Socket => TCP 
                le reseau recoit un paquet => TCP=> socket=>programme

            -Kernel fait le lien entre processus=> objet socket creation d'un file descriptor => creation d'un struct file => socket file 

            -Bind() => Associe ce socket au port 8080 + IP (recoit une identite locale ) => il enregistre cette identite ds struct sock

            -Listen() => ds struct tcp_sock change l etat au TCP listen, ds ce point la pile réseau commence a traiter les SYN entrantes 




    *  La creation d'un nv socket pour chaque client car TCP est connexion orientée 
        Lorsque le client envoie un SYN le kernel verifie le socket listen => creer automaitquement un nv struct sock + tcp_sock + socket + fd pour cette connexion specifique => le kernel fait accept() le kernel sort ce socket de client de la accpet queue  

## Modele client | serveur : 
- Un serveur est une machine ou programme qui fourni des services ou ressources au clients 
    Ecoute les connexions entrantes , recevoir des demandes des clients, traiter des demandes, Envoyer des reponses

- CLient est une machine ou un programme qui demande un service a un serveur =>
    Initier une connexion avec un serveur, envoyer une requete, recevoir et utiliser la reponse du serveur 

### qui initie la connexion ?? 
- ds l e modele client | serveur tjr le client qui initie la connexion , le serveur reste écouter sur un port 
    lorsque le service demande a OS creer un socket  et fait bind() + listen() le socket est cree en memoire + relie a IP + PORT + Service => Etat Listen cad pret a accepter des connexions mais aucun connexion n'existe encore juste le serveur ecoute 

    * Pourquoi .??? :
        Le serveur ne sait pas a l'avance qui veut se connecter ; donc il reste a l'ecoute sur un port cepandant le client sait qu'il veut acceder a un service sur une machine 

### Relation 1 <-> 1 (SERVEUR <=> CLIENT)
- Le Serveur communiquer avec un seul client a la fois une fois connecte , le serveur est occupe avec ce client il ne peut pas accepter d'autre connexions

### Relation 1 <-> 1> (SERVEUR <=> CLIENTs)
- Un serveur peut accepter plusieurs clients en meme temps 
- Chaque client a sa connexion independante avec le serveur 
- Le serveur utilise des threads, socket dedies pour chaque client 
    Le serveur n’a pas besoin de créer un nouveau socket
    Le même socket LISTEN peut servir directement à échanger les données
    ####  Comment ? 
        - Socket listen => serveur est a l ecoute sur un port 
        - Chaque client se connecte   le socket listen ne devient pas la connexion reelle ,=> la creation d'un nv socket interne  contient 
                            socket_client = {
                                IP serveur
                                port serveur
                                IP client       // unique pour ce client
                                port client     // numéro temporaire choisi par le client
                                service/programme
                                état = ESTABLISHED
                            }
        - On va parle sur les THREADS 
### Pourquoi le serveur ne sait rien d'avance des clients
    les clients peuvent etre dynamique et nombreux  et ne peut pas savoir leur IP (IP dynamique )et leurs ports.
    Si le serveur devait connaitre tous les client a l'avance 
        -IL doit pre-alloue une connexion en memoire chaque client a son socket interne  +> des que le serveur demarrer il aurait tous ces sockets deja ouverts et prete de recevoir => c'est impossible 

## TCP (Transmission Control Protocol):
    -TCP est un protocole de transfert entre 2 processus, est un protocole oriente connexion.
        RQ: 
            Un processus est une instance active d’un programme, créée et gérée par le système d’exploitation, avec sa propre mémoire et ses ressources.
    -segmenter des données en paquets
    - numerotation des segments pour garder l'ordre
    - retransmission en cas de perte 

    write(clientfd(file descriptor), "hello world", 11)
        -copie ces octets ds buffer de sortie
        -segmentation 
        -L'entete TCP 
            chaque segment devient => [En-tête TCP | Données]
            contient :
                Source port + destination port + numero d séquence + NUmero d'ACK + Flags ...
        -Passage a IP layer
            Ip encapsule le segement ds un paquet IP
            .
            .
            .
            ...
   ### Comment La creation d'une Connexion  TCP?...
#### 3-WAY HANDSHAKE  SYN(CLIENT) => SYN + ACK(SERVEUR) => ACK(CLIENT) 
    1- Service démarre

    2- Le client veut envoyer au serveur
         - kernel creer un struct socket client + struct sock + struct tcp_sock + FD associe au socket + TCP closed  + Client appelle connect(fd, &addr_serveur, sizeof(addr_serveur)); + si le client ne fait pas bind afin de lier un port a ce socket  il choisi un port libre  => Initialisation d'un TCP (closed => SYN_SENT) => envoie d'un paquet SYN => KERNEL client attend SYN+ACK
        ex: navigateur

        le navigateur  dit a L'OS  je veux envoyer des donnees  a Ip sur le port 80 en TCP
        -Os choisit un port local libre
        -Creation d'un socket client 
        -Prepare un packet TCP => SYN SYNchronize(demande de connexion)
        -Envoie ce paquets via IP
    
    3- le serveur Recoit cette demande 
        -OS regarde le port et IP de destination => trouve le socket en etat listen 
        -OS creer un nv socket (IP+Port +IP et PORT du client)=> Mode RECEIVED  SYN+ACK(ACKnowledge) Recu de la demande 

    4-Client confirme 
        -client recoit le SYN+ACK
        -OS client socket => MOde ESTABLISHED
        -envoie ACK
ESTABLISHED
    5-Recoit du ACK
        -OS serveur socket => MOde ESTABLISHED
        -Connexion TCP existe 

C/C:
    une connexion TCP est ouverte, chaque côté (client et serveur) possède un socket TCP qui contient 
        Socket TCP =
            {
            IP locale,
            Port local,
            IP distante,
            Port distant,
            État = ESTABLISHED
            }

    RQ :
        le socket LISTEN ne peut jamais read ou write  => IL SERT D 'ACCEPTER JUSTE LES NV CONNEXIONS => ETAT LISTEN

        NV SOCKET CLIENT => ETAT TCP : ESTABLISHED => READ AND WRITE 
