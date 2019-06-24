================================================================================
Testowanie aplikacji
================================================================================

Projekt aplikacji
--------------------------------------------------------------------------------

.. note::

    Aplikacja pisana w Go może być interesująca, gdyż nadal nie jest to wyjątkowo
    powszechny język programowania: Tiobe, Maj 2019: 19 miejsce dla Go, 4 - python,
    3 - C++, 1 - Java.



.. figure:: /_static/archi.jpg
   :alt: Architektura prostej aplikacji opartej o mikrousługi

   Architektura prostej aplikacji opartej o mikrousługi

Podstawowe założenia
--------------------------------------------------------------------------------

#. Pierwsza aplikacja zbiera dane sprawdza ich poprawność po czym przesyła je do kolejenj mikrousługi.
#. Druga aplikacja sprawdza czy użytkownik istnieje(tworzy go, jeżli nie), generuje UID oraz wysyła dane do apliakcji numer 3.
#. Aplikacja 3 tworzy losowy PIN i numer karty(spełniający określone wymagania) i zwraca wynik do apliakcji numer 3.
#. Aplikacja 2 zbiera UID, PIN oraz numer karty i wysłya do aplikacji 1.
#. Aplikacja 1 wyświetla wszystkie dane łącznie z UIDem, PINem, numerem karty 

Testowanie jednostkowe aplikacji
--------------------------------------------------------------------------------

- Testy jednostkowe znajdujął się w plikach o nazwie `main_test.go`. Umieszczone zostały w każdym folderze z mikroserwisem. 
- Same testy są napisane również w Go z wykorzystaniem bibliotek standartowych oraz `github.com/stretchr/testify/assert`.
- Przykładowa funkcja testująca status aplikcaji: 

.. code-block:: go

    func TestGetStatus(t *testing.T) {

        e := echo.New()
        req := httptest.NewRequest(http.MethodGet, "/status", nil)
        rec := httptest.NewRecorder()
        c := e.NewContext(req, rec)

        // Assertions
        if assert.NoError(t, getStatus(c)) {
            assert.Equal(t, http.StatusOK, rec.Code)
            assert.Equal(t, msg1, rec.Body.String())
        }
    }

- Wynik uruchomienia testów jednostkowych dla poszczególnych mikroserwisach:

.. code-block:: bash

    sky@A1$ go test -v
    === RUN   TestGetStatus
    --- PASS: TestGetStatus (0.00s)
    === RUN   TestBasicValidate
    --- PASS: TestBasicValidate (0.00s)
    === RUN   TestSendDataToApp2
    --- PASS: TestSendDataToApp2 (0.00s)
    === RUN   TestCreateNewCard
    --- PASS: TestCreateNewCard (0.00s)
    PASS
    ok  	_/home/sky/Documents/code/Go/testowanie-jednostkowe-w-aplikacjach-w-systemie-rozproszonym/3sky/good_apps/A1	0.020s

.. code-block:: bash

    sky@A2$ go test -v
    === RUN   TestCheckUserExist
    --- PASS: TestCheckUserExist (0.00s)
    === RUN   TestGetStatus
    --- PASS: TestGetStatus (0.00s)
    === RUN   TestCreateNewCard
    --- PASS: TestCreateNewCard (0.00s)
    === RUN   TestCreateClient
    --- PASS: TestCreateClient (0.00s)
    PASS
    ok  	_/home/sky/Documents/code/Go/testowanie-jednostkowe-w-aplikacjach-w-systemie-rozproszonym/3sky/good_apps/A2	0.017s

.. code-block:: bash

    sky@A3$ go test -v
    === RUN   TestCreateCard
    --- PASS: TestCreateCard (0.00s)
    PASS
    ok  	_/home/sky/Documents/code/Go/testowanie-jednostkowe-w-aplikacjach-w-systemie-rozproszonym/3sky/good_apps/A3	0.014s

Dockeryzacja
````````````````````````````````````````````````````````````````````````````````

- Aplikacje napisane w języku GO bardzo łatwo można umieścić w kontenerze, same kontenery potrafiał zajmować ok 6MB.
- Przykładowy `Dockerfile` dla wykorzystywanej apliakcji:

.. code-block:: python

    FROM golang:alpine as builder
    RUN apk add --no-cache git gcc libc-dev
    RUN go get github.com/labstack/echo && go get github.com/stretchr/testify/assert
    ADD . .
    RUN go test -v
    RUN go build -o main

    FROM alpine
    COPY --from=builder /go/main main
    EXPOSE 5000
    CMD ["./main"]

- Do samego budowania wykorzystano multi-stage building tak aby finalne binaria nie posiadały niepotrzebnych bibliotek
- W trakcie budowania obrazów były uruchomiane testy jednostkowe ich pozytywne zakończenie decydowaly o poprawnym zbudowaniu obrazu dockerowego

Docker-compose: aplikacja
````````````````````````````````````````````````````````````````````````````````
- W celu stworzenia serwisów wykorzystano docker-composa, tak aby za w łatwiejszy sposób kontrolować wszystkie mikrousługi
- Sam plik docker-compose.yml również nie jest zbyt rozbudowany, tworzy 3 niezależne mikroserwisy:

.. code-block:: yml

    version: "2"

    services:
    app1:
        container_name: 'good-app1'
        build: A1/
        ports:
        - "5000:5000"
        network_mode: host

    app2:
        container_name: 'good-app2'
        build: A2/
        ports:
        - "5001:5001"
        network_mode: host

    app3:
        container_name: 'good-app3'
        build: A3/
        ports:
        - "5002:5002"
        network_mode: host

Docker-compose: testy czyli `Przykładowa symulacja sytuacji, która mogłaby mieć miejsce w prawdziwym projekcie`  
````````````````````````````````````````````````````````````````````````````````````````````````````````````````````````````````````````````````````````````````
#. Stworzono dwa odseparowane pliki docker-compose, do budowy dwóch serwisów
#. Zrobiono nieznacznał zmiane w kodzie
    - Osoba odpowiedzialna za apliakcje 2 stwierdziła, że tag `nuid` jest lepsza niż `uid`, po czym ja zmieniła
    .. code-block:: go

        type Account struct {
            UID      string `json:"nuid"`
            CardData Card
        }
    - Jednak osoba odpowiedzialna za applikacje 1, nie wiedziała o zmianie i wziąż używała tagu `uid`.
    .. code-block:: go
        type Account struct {
            UID      string `json:"uid"`
            CardData Card
        }

#. Testy jednostkowe wciaż, działały poprawnie i nie wykryły błędu - w końcu odbywały się w ramach poszczególnych mikroserwisów.
#. Jednak wynik wygladał następująco:

.. figure:: /_static/Testing.png
   :alt: Zrzut ekrany serwisów uruchomionych za pomocą popularnej apliakcji *tmux*

#. Wystarczy się przyjrzeć aby zauważyć, że wartość pola UID w systemie o zmienionych tagach jest pusta.


Wnioski
````````````````````````````````````````````````````````````````````````````````
Jak zaprezentowano w powyższym przykładzie testy jednostkowe są potrzebne i pozwalają wyeliminować znaczął cześć błedów. Jednak nie są one wystarczające w świecie mikroserwisów oraz systemów rozproszonych.
W przypadku takiej architektury najważniejszym elementem jest komunikacja zarówno między usługami jak i między developwerami. 
Problemy te można rozwiązać stasując testy integracyjne, które wymagają więcej pracy oraz zdecydowanie lepszej organizacji. Jednak są w stanie oszczędzić zespołom nieprzyjemnych sytuacjach. 