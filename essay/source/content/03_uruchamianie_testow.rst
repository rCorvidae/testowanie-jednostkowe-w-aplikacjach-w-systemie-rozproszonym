================================================================================
Testowanie aplikacji
================================================================================

Projekt aplikacji
--------------------------------------------------------------------------------

.. note::

    Aplikacja pisana w Go może być interesująca, gdyż nadal nie jest to wyjątkowo
    powszechny język programowania: Tiobe, Maj 2019: 19 miejsce dla Go, 4 - python,
    3 - C++, 1 - Java.



.. figure:: /_static/architektura_aplikacji.png
   :alt: Architektura prostej aplikacji opartej o mikrousługi

   Architektura prostej aplikacji opartej o mikrousługi

Projekt aplikacji rozproszonej bazującej na mikrousługach:
Liczenie PI na podstawie całkowania numerycznego: `Obliczanie PI metodą Simpsona <https://www.mcs.anl.gov/research/projects/mpi/tutorial/mpiexmpl/src/pi/C/solution.txt>`_.
Przepisać to tylko na microusługi. Załóżmy, że jest usługa master, która wysyła odbiera polecenie od użytkownika:
GET pi()

Usługa ta robi operację Bcast dla wszystkich dostępnych usług liczących (jedna aplikacja, 3 repliki). Taka usługa przyjmie jako wartość zakres całkowania i krok. Zakres całkowania to n/3, krok jakiś dowolny się wybierze.

Po wykonaniu tych kroków, aplikacja master scala wynik i zwraca klientowi.

Wspomnieć o tym, że aplikacja przypomina MPI.

Węzeł główny (Gateway)
````````````````````````````````````````````````````````````````````````````````

.. note::
   Krótki opis zadań

Węzeł roboczy (ikona silnika)
````````````````````````````````````````````````````````````````````````````````

.. note::
   Krótki opis zadań

Testowanie jednostkowe aplikacji
--------------------------------------------------------------------------------

Węzeł główny 
````````````````````````````````````````````````````````````````````````````````

Węzeł roboczy 
````````````````````````````````````````````````````````````````````````````````

Dockeryzacja aplikacji i testów
--------------------------------------------------------------------------------

Dockeryzacja
````````````````````````````````````````````````````````````````````````````````

Docker-compose: aplikacja
````````````````````````````````````````````````````````````````````````````````

Docker-compose: testy
````````````````````````````````````````````````````````````````````````````````

Kubernetes: aplikacja
````````````````````````````````````````````````````````````````````````````````


Uruchomienie testów w Jenkinsie
--------------------------------------------------------------------------------

Opis podstawowego deploymentu Jenkinsa z użyciem docker-compose lokalnie.
Uruchomić następnie obraz do testów z każdym przychodzącym commitem

Można Jenkinsa rzucić też K8s jeśli gdzieś są gotowe YAML'e pod to - poszukać
w Helm'e albo innym Rancher'e.

Awaria aplikacji w chmurze
--------------------------------------------------------------------------------

Po etapie implementacji i przetestowania modułów aplikacji z pomocą unit testów,
doprowadzić do jej awarii z powodu przeciążenia sieci lub nagłego usunięcia
node'a w Kubernetesie.

Na podstawie tego etapu decyduje się o wnioskach. Ten punkt ma udowodnić, że
same unit testy nie są w stanie zapewnić działania aplikacji, pomiimo dobrych
chęci.
