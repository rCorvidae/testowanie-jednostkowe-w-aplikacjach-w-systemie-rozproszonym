================================================================================
Techniki i narzędzia dla testów jednostkowych
================================================================================

Techniki testowania jednostkowego
--------------------------------------------------------------------------------

Pojedyncze testy
````````````````````````````````````````````````````````````````````````````````

Test na zbiorze danych
````````````````````````````````````````````````````````````````````````````````

Stubbing
````````````````````````````````````````````````````````````````````````````````

Mock objects
````````````````````````````````````````````````````````````````````````````````

Frameworki do testów
--------------------------------------------------------------------------------

.. note::
    Omówić narzędzie pod kątem wyżej wymienionych podpunktów. Jeśli
    framework testówy nie obsługuje np. Mock Objectów (np. w c++), uzasadnić
    dlaczego (brak refleksji statycznej/dynamicznej).

Java: JUnit 5
````````````````````````````````````````````````````````````````````````````````

Python: unittest
````````````````````````````````````````````````````````````````````````````````

Go: testing
````````````````````````````````````````````````````````````````````````````````

Narzędzia CI/CD
--------------------------------------------------------------------------------

.. note::
    Stosunkowo krótko o wybranych narzędziach. Jeśli możlwie, to bazować
    na dodatkowych linkach (w języku angielskim). Przedstawić rację
    istnienia danego narzęcia. Ten pod-rozdział odnosi się do rozdziału trzeciego

Jenkins
````````````````````````````````````````````````````````````````````````````````

Docker
````````````````````````````````````````````````````````````````````````````````

Podstawowe informacje
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

.. note::
    Ze względu na złożoność narzędzia Docker, przedstawia się jedynie
    najważniejsze cechy i informacje. W przypadku braku wyjaśnienia
    pewnych terminów, posiadać będą one link do dokumentacji i definicji
    tego terminu.

**Docker** jest narzędziem umożliwiającym uproszczenie procesu wytwarzania 
i instalowania oprogramowania z wykorzystaniem kontenerów. Kontener pozwala
na hermetyzację zależności pomiędzy bibliotekami systemowymi a wymaganymi
przez aplikację. Dzięki tej funkcjonalności umożliwia się łatwą instalację
oprogramowania unikając problemów m.in. kompatybilności ABI
wykorzystywanych/instalowanych bibliotek.

`Architektura <https://docs.docker.com/engine/docker-overview/#docker-architecture>`_
Dockera jest następująca:

    * *dockerd* - daemon Dockera, serwer nasłuchujący zapytań Docker API.
      Zarządza cyklem życia kontenerów, montowaniem zasobów (w tym sieciowych).
      Podstawowymi technologiami wykorzystywanymi przez dockerd są
      linuksowe: namespaces i control groups (cgroups):

      * Namespaces - zapewniają odizolowaną przestrzeń użytkownika,
        zwaną kontenerem. Odizolowanie polega na stworzeniu własnego
        drzewa id procesów (PID), samodzielnych ustawień sieci, punktów
        montowania dysków itp.
      * Cgroups - umożliwiają ograniczenie zasobów (m.in. CPU, pamięci), dla
        uruchomionych aplikacji w danym kontenerze.

    * *docker* - jest narzędziem CLI umożliwiającym łatwe wysyłanie zapytań
      do daemona Dockera. Jest głównym interfejsem za pomocą którego
      operator uruchamia kontenery.
    * *rejestr* (Docker Registry) - repozytorium obrazów Dockera.


Wykorzystując Docker'a należy zrozumieć dwa następujące terminy:

    * *image* (obraz) - szablon tylko-do-odczytu posiadający niezbędny zestaw
      plików konfiguracyjnych, aplikacji i bibliotek umożliwiający uruchomienie
      wybranej aplikacji). Właściwa aplikacja najczęściej również znajduje się
      w obrazie.
    * *container* (kontener) - uruchomiony i aktywny obraz. Wraz z uruchmionym
      kontenerem, aplikacja świadczy swoje usługi.

Po zainstalowaniu Docker'a, warto sprawdzić, czy narzędzie działa. Sprawdzić to
można z użyciem następującej komendy::

    $ docker run hello-world

Z pomocą tego polecenia, Docker pobiera obraz o nazwie hello-world oraz go 
uruchamia. W efekcie, wyświetlony zostaje następujący komunikat::

    Hello from Docker!
    This message shows that your installation appears to be working correctly.

    To generate this message, Docker took the following steps:
      1. The Docker client contacted the Docker daemon.
      2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
         (amd64)
      3. The Docker daemon created a new container from that image which runs the
         executable that produces the output you are currently reading.
      4. The Docker daemon streamed that output to the Docker client, which sent it 
         to your terminal.
    [...]

Dockerfile
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Docker-compose
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Kubernetes
````````````````````````````````````````````````````````````````````````````````
