## TCP (Transfer Control Protocol): 
   Un segment TCP [ HEADER TCP ] + [ DATA ]

* HEADER TCP :
    Source Port (Port d'expediteur) 
    Destination Port
    Séquence Number (Num du premier pctet envoyé)
    Acknowledgment Number (Num du prochain octet attendu)
    Flags (SYN=> debut d conn, ACK=>reception, RST=>reset, FIN=>close...)
    Window SIZE (taille de biffer disponible)
    Checksum (verification d'erreur)

* DATA (Payload)
    Octets

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

le 1er objet TCP qui recoit un segment est un struct sock  (verification SEQ/ACK, enleve HEADER, garde donnees, ces donnees sont mises ds sk_recieve_queue ):

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
