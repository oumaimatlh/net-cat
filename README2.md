## TCP (Transfer Control Protocol):


## TCP (Transmission Control Protocol):
   
* HEADER TCP :
    Source Port (Port d'expediteur) 
    Destination Port
    Séquence Number (Num du premier pctet envoyé)
    Acknowledgment Number (Num du prochain octet attendu)
    Flags (SYN=> debut d conn, ACK=>reception, RST=>reset, FIN=>close...)
    Window SIZE (taille de buffer disponible)
    Checksum (verification d'erreur)

* DATA (Payload)
    Octets
/-------------------------------------------------------------/
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

        -Prepare un packet TCP => SYN SYNchronize(demande de connexion)
    
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

   Un segment TCP [ HEADER TCP ] + [ DATA ]


### Comment TCP (ordre chronologique)
    SOCKET TCP => state CLOSED
    CONNEXION (3-way handshake)

* Client -> construit un Segment TCP ->  SYN
    FLAG => SYN = 1 
    SEQ = x 
    state = SYN-SENT 
    ACK = 0 
    Pas de donnees

* Serveur -> SYN + ACK
    SYN => 1, ACK => 1, SEQ => y, ACK=x+1
    STATE = SYN-RECEIVED

* Client -> ACK
    ACK = 1, ACK = y+1, STATE= ESTABLISHED

### recoit des paquets

le 1er objet TCP qui recoit le segment  TCP c'est un struct sock  (verification SEQ/ACK, enleve HEADER, garde donnees, ces donnees sont mises ds sk_recieve_queue ):
    RQ: 
        c'est quoi sk_recieve_queue => struct sk_buff_head sk_receive_queue;


read(fd, buffer, size)
    Le kernel prend le fd
    Il retrouve struct file
    Puis struct socket
    Puis le struct sock
    Il regarde dans sk_receive_queue


TCP fait deux choses :

    1/Réassembler les segments dans l’ordre en fonction de seq number
    2/Copier seulement les octets consécutifs attendus dans sk_receive_queue
        rcv_nxt = 101
        Segment reçu → seq=103-110
        ↓ stocké temporairement (out-of-order)
        Segment reçu → seq=101-102
        ↓ mis dans receive queue
        ↓ rcv_nxt = 111
        ↓ maintenant segment seq=103-110 est ajouté en ordre


```go
    struct sock
    └── sk_receive_queue (struct sk_buff_head)
            └── liste de struct sk_buff

```


### STRUCT SOCK
    c'est une structure (objet d controle) (kernel space) alloue dynamiquement en mémoire kernel
    elle represente un endpoint d'une connexion réseau 
    contient ; l'etat TCP , les filles d'attente, les pointeurs veers protocole, ...
```GO
    struct sock {
        socket_state sk_state;
        struct sk_buff_head sk_receive_queue; //    Donnees recus 
        struct sk_buff_head sk_write_queue; //      Donnees envoyer 
        struct sk_buff_head sk_backlog;
        struct proto *sk_prot;
        void *sk_protinfo; // pointe vers tcp_sock
        ...
};

```

### STRUCT SK_BUFF_HEAD
    c'est une structure de gestion de liste chainee => pointeur vers struct sk_buff
```GO
    struct sk_buff_head {
        struct sk_buff *next; // premier element
        struct sk_buff *prev; //dernier element 
        __u32 qlen; //nbr d elements 
        spinlock_t lock;
};

```

### STRUCT SK_BUFF
    Represente un  paquet reseau (donnees + metadata )

```