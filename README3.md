## COMPILATION IN GO

FILE =>>>>>

### 1.LEXER(Scanner)
    Transforme le code source en Tokens (chaque token est un mot cle, identifiant, opérateur...)
    RQ=> Tokens ne contient aucune structure il ne sait pas les instruction
        EX : x = a + b * c
                            ==>>  [x] [=] [a] [+] [b] [*] [c]

* Objectif: Nettoyer le texte et le rendre compréhensible pour le parser

### 2.Parser 
    Prend le Token du lexer et construit AST (abstract Syntax Tree)
        AST est une représentation arborescente de notre Programme

* Objectif: Structurer le code selon la syntaxe du langage cad  le parser transforme cette liste de token en strucutr hiérarchique cad il comprend :
            La priorite des opérateurs, la forme des instructions, structures des blocs, les appels des fonction, déclaration...

            EX:
```GO
                    x = a + b * c

                    Assignment
                        ├── Var: x
                        └── Add
                            ├── a
                            └── Multiply
                                ├── b
                                └── c


```

POURQUOI?? ; Le compilateur doit comprendre la structure logique, SANS PARSER, LE COMPILER ne sait pas ou commence le bloc quelle instruction appartient au if pr exemple ect ...

RQ: le parser produit un AST mais ne connait pas encore comment ca va etre exécuter pu stocke en mémoire 

### Semantic Analysis (Type Checker):
    Verifier que le code fait sens ; types, regles d langages ..., s'il ya une erreur au niveau d ca le programme stop la et retourne une erreur .

### Intermediate Representation (IR):
    Transfert l' AST en une representation intermédiaire plus proche d la machine; pour etre lisible par CPU et peut interagir avec la mémoire (allocation ect )

        AST -> IR -> CODEMACHINE(BYTECODE)

        EX: 

        CODE :
            a := 5                
            b := 10            
            c := a + b

        TOKENS 
        AST:
            Assign(c)
                └── Add
                    ├── Var(a)
                    └── Var(b)

        IR (SSA ; single Static Assignment ):
            on transfert AST => IR afin de rendre des instructions lineaires simples a manipuler ect,
            Le compilateur transforme les varibales temporaires SSA pour eviter les réassignations ect 
        
            EX :
                x := 5 ; x = x + 2 ; y := x + 3

                <<=>>
                t1 = 5       // x initial
                t2 = t1 + 2  // nouvelle valeur de x
                t3 = t2 + 3  // y


### MACHINE CODE GENERATION :

    A ce stade le compilateur a  déja le pseudo-code IR/SSA :
        chaque instruction ce transforme an code machine natif que le CPU exécute directement 

* Mapping IR -> registres/mémoire  => ALLOCATION EN MÉMOIRE (décide ou stocker chaque variable)
    Le programme chargé en RAM ; le compilateur doit mapper chaque variable temporaire sur un emplacement physique

        RAM : Organisée en cases mémoires identifiées par des adresses 
            chaque adresse peut stocker 1-8 octets selon l'architecture 

```GO
    //ALLOCATION EN MÉMOIRE : 
        /*
            INT(ex:int64 ): 
                stocke ds 8 cases, chaque case contient 1 octet(8bits) : 
                    x:=5 => (0x1000 - 0x1007)

            BOOL:
                1 OCTET -> adresse pr exemple 0x3000 => 0(False) ou 1(True)

            STRING: 
                type string struct {
                    ptr *byte  //adresse mémoire du premier caractere (8 octets)
                    len int => 8octets 
                }

            ARRAY:
                l'adresse de l'array = adresse du premier element 
                
        
        
        
        */

```
        