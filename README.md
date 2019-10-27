# GoTutorial
Aufgabe 2:

1.Was bedeutet $USER im Kontext der Bash?

2.Weshalb sind die Anführungsstriche wichtig?

ANTWORT:

   1. Alles im Shell, was mit $ beginnt, ist eine Environment-Variable. "$USER" ist eine, die unmittelbar dem Login-Namen im aktuellen System zugeordnet ist.


   2. Der Shell ordnet einen gegebenen Text dem String-Typ zu, wenn der Text in die Anführungszeichen gesetzt ist.
   Außerdem ist der Shell fähig selbst den Datentyp zu definieren. Besteht z.B. ein Text rein aus Ziffern - wird er zum numerischen Datentyp zugeordnet; ist der Text in die Anführungsstriche gesetzt - wird er unabhängig vom Inhalt zu einem String. 

   Die Anführungsstriche lösen auch ein Problem mit den Textlücken. Man kann problemlos die Location home/Sandbox aufrufen, das funktioniert aber nicht mit einem Ordner namens zB home/Parser and tables. In dem Fall sind die Anführungszeichen ein Muss. 








Aufgabe 5:

Unterschied im Error-Handling zwischen Go und Java

ANTWORT:

   Der erste Unterschied liegt daran, dass die Errors in Go, wie Rob Pike im Titel seines Artikels sagt, bestimmte Werte(Values) sind. 
   Dem Go-Team ist es gelungen, auf Exceptions wie z.B in Java zu verzichten und das explizite Ablesen der Fehler zu ermöglichen. Diese Genauigkeit macht den Code zwar schwiereger lesbar, im Nachhinein aber kann man verstehen wo der Fehler ist. 
   Soll ein Go-Call ein Fehler ergeben, wie schon erwähnt, ist er eine Value wie jede andere und bewirkt keine automatische Veränderungen im Ablauf eines Programms. 


   Bei Java sind die Errors "Exceptions". Passiert ein "throw" von einem Exception, wird das ganze Programmablauf abgebrochen bzw. läuft bis zum nächsten Catch. Im Endeffekt kann man nicht darüber bestimmen, ob es weitergehen soll oder nicht, obwohl jeder Call im seinen eigenen "try-catch" "gewrappt" ist.
   Der Vorteil ist es, dass man den Code unmittelbar von Error-Handling trennen kann. 







Aufgabe 7:

1.Was passiert genau, wenn der Receiver-Typ ein Value ist?

2.Was passiert genau, wenn der Receiver-Typ ein Pointer ist?

3.Erkläre in eigenen Worten und vollständigen Sätzen den Unterschied zwischen einem Value- und einem Pointer-Typ.

4.Recherchiere die Bedeutung von call-by-reference und call-by-value. Warum heißt es, dass Java und Go eine call-by-value Semantik haben?

5.Recherchiere über Referenzparameter in C++. Prüfe die Definition von Referenzparametern in C++ genau und vergleiche sie mit der Übergabe von Pointern in C++. Wo sind die Unterschiede?

6.Was hälst du von der Definition von Referenzparametern in C++? Frage deine Kollegen, wer sich mit C++ im Unternehmen gut auskennen könnte und suche ihn auf. Bitte ihn darum, dir den Unterschied zwischen der Übergabe von Referenz- und Pointer-Parametern zu erklären (pass-/call-by-reference vs pass-/call-by-pointer). Diskutiere mit ihm: Ist das mehr als Compiler-Sugar? Dokumentiere die Argumentation und das Ergebnis.

7.Nenne alle Value-Datentypen von Java.

ANTWORT:


   1. Wenn der Receiver in einer Funktion ein Value ist, wird eine Kopie vom Typ T gemacht und an die Funktion angepasst. Im Nachhinein enthält die Funktion den gleichen Typ T (wie angegeben), aber "besitzt" einen anderen, eigenen Platz im Speicher.

   2. Ein Pointer-Receiver dagegen, passt die Adresse vom Typ T zur Funktion an. Die funktion enthält dann "den Weg" zum Typ T, oder seine "Adresse" (auf english - reference).

   3. Der erste Unterschied ist der Anwendungszweck: Wenn es gewollt ist, dass eine Funktion die Werte vom Typ T verändert, soll man im Receiver einen Pointer benutzen. Wenn aber eine Variable unverändbar und berechenbar sein soll - benutzt man den Value-Receiver. 
   Ein noch wichtiger Unterschied liegt daran, dass ein Value sowohl mit einem Pointer als auch mit einem Value aufgerufen werden kann, ein Pointer wird aber nur mit einem anderen Pointer aufgerufen(&).

   4. Eine Call-by-value Funktion arbeitet mit einer Kopie vom Typ T. Demnach wird der Typ T nicht von der Funktion beinträchtigt und hat außer dieser Funktion den gleichen Inhalt, da es ein aktueller und ein von der Funktion abhängiger Argument im Speicher erzeugt werden
      Bei einer Call-by-reference Funktion werden ein aktueller und ein von der Funktion abhängiger Typ T "unter gleicher Adresse" gespeichert, demnach wird Typ T von der Funktion beeinträchtigt.

   5. Sowohl bei Java als auch bei Go gilt: Pass-by-value stellt sicher, dass die Funktion keine unerwartete Effekten auf den Code verursachen wird. Value-Semantik heißt, dass man direkt mit dem Value arbeitet, während es die Kopien davon bei z.B. Neuinitialisierung erstellt werden. Zum Beispiel:
   int x = 1;
   int y = x;
   x = 2; // y bleibt immer noch 1
   
   6. In C++ gibt es eine Alternative, wie man aus einer Funktion auf Variablen des Aufrufers zugreifen kann. Dazu wird statt einer Pointer-Variable eine Referenz als Parameter verwendet. Eine Referenz unterscheidet sich syntaktisch dadurch, dass statt des Sterns ein "&" vorangestellt wird.
  Der Unterschied zwischen Referenz- und Pointer-Parametern liegt daran, dass ein Referenz quasi ein zweiter Name für dieselbe Variable ist. Daher muss man sie nicht immer über die Dereferenzierungs-Operatoren ansprechen. Ein Pointer dazu ist eine Variable, die eine Speicheradresse enthält.
   
   
   7. In Java unterteilt man Datentypen in primitive Datentypen und Objekten.
   Sechs von acht primitiven Datentypen sind numerisch: byte, short, integer, long, float, double - jeder von diesen speichert die Zahlenwerte. Zwei andere primitiven Datentypen sind character und boolean.
   Alle andere Datentypen sind Objekte, man nennt sie auch "komplexe Datentypen", weil sie sich im Gegensatz zu den primitiven Datentypen aus verschiedenen Elementen zusammensetzen. Grob gesagt ist es ein großer Speicherblock, in dem alle zugehörigen Daten festgehalten werden.
   
   
   
   
   
   
   
   
Aufgabe 8:

   Die Beobachtung:

   Laut der Dokumentation ist GeoM eine Struktur, die eine affine Matriz präsentiert.

   Die Struktur ist anzuwenden, wenn man das Objekt umdrehen oder transformieren möchte.
   Bei der Translate-Funktion gilt: je größer der eingesetzte x-Wert ist, desto mehr verschiebt sich das Bild nach rechts;      der größere Y-Wert schiebt das Bild nach unten, da die Y-Achse in Ebiten nach unten orientiert ist. Mit GeoM.Translate(x      float64, y float64) wird die Position also verändert und gespeichert, mit GeoM.Reset() kehrt das Objekt auf die              Anfangsposition zurück. 
