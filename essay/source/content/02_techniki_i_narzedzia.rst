================================================================================
Techniki i narzędzia dla testów jednostkowych
================================================================================

Techniki testowania jednostkowego
--------------------------------------------------------------------------------

Poniższy rozdział bazuje głównie na doświadczeniach własnych z zakresu
testowania aplikacji. Ma on jednak odbicie w doświadczeniach innych
programistów (zamieszczone stosowne referencje). W przypadku konfliktów
terminologii, prezentowane są źródła na których treść jest bazowana.

Pojedyncze testy
````````````````````````````````````````````````````````````````````````````````

Testowanie jednostkowe jest, jak nazwa wskazuje, techniką testowania
pojedynczych jednostek oprogramowania np. funkcji, klas, modułów - przeważnie
funkcji/metod.

W przypadku takich testów, najczęściej testuje się pojedynczą funkcję
i analizuje względem oczekiwanego rezultatu. Jeśli wynik danego testu
jednotkowego jest rozbieżny z oczekiwanym rezultatem, kończy się on
niepowodzeniem. Ilość powodzeń/niepowodzeń zbioru testów jest następnie
prezentowana w postaci podsumowania generowanego przez wybrany framework.

Przykładowo, od jednostki oprogramowania (funkcji) oczekuje się zwrócenie
wartości liczby znajdującym się na wybranej pozycji w ciągu Fibbonacciego,
który przybiera następującą postać::

    index: 0, 1, 2, 3, 4, 5, 6
    value: 1, 1, 2, 3, 5, 8, 13, ...

Oczekiwania wobec funkcji Fibbonaciego prezentowane są w postaci następujących
testów. Testy prezentują nie tylko oczekiwany efekt, ale również definiują
API, co jest jednym z założeń TDD. Oto testy::

    assert 1 == fib(0)
    assert 5 == fib(4)
    assert 8 == fib(5)

Kod spełniający wyżej wymienione wymagania (inaczej testy), jest następujący::

    def fib(index):
        if index < 2:
            return 1

        a, b = 2, 3
        for _ in range(2, index):
            a, b = b, a + b

        return a

Przykład ilustruje potęgę testowania jednostkowego, ale również i jego wady.
Do zalet z pewnością jest zaliczenie testów do narzędzi umożliwiających
stworzenie oprogramowania lepszej jakości. Do wad należy często czasochłonność
oraz niemożność całkowitego pokrycia testami każdego przypadku. Dla powyższego
przykładu należy się zastanowić, czy wartość index < 0 powinna być dozwolona,
czy funkcja nie powinna zwrócić np. wyjątku.

Data Driven Unit Test
````````````````````````````````````````````````````````````````````````````````

Data Driven Unit Test jest rozszerzeniem koncepcji pojedynczego testu. Jest
to sposób testowania wybranego komponentu *obszernym* pakietem danych. Taki
sposób testowania jest wspierany przez niektóre systemy testowania np.
`Qt Tests <https://doc.qt.io/qt-5/qttestlib-tutorial2-example.html>`_,
`MS Unit Test Framework <https://docs.microsoft.com/en-us/visualstudio/test/how-to-create-a-data-driven-unit-test?view=vs-2019>`_.

Technika ta polega na wykorzystaniu danego źródła danych np. z pliku CSV
i przekazaniu tych danych do testowanej funkcji. W przypadku niepowodzenia,
framework testowy stosowanie poinformuje o problemie. Pozwala to na 
automatyzację testowania i zwiększenia prawdopodobieństwa wykrycia potencjalnych
awarii.

Stubbing i Mock obiekty
````````````````````````````````````````````````````````````````````````````````

Technika wykorzystująca **mock obiekty** (atrapa obiektu [#martin_agile]_) polega
na zastąpieniu kodu dziedzinowego kodem symulującym jego **zachowaniem** 
[#endotesting]_. Przykładowo, testowany pojedynczy komponent (np. funkcja)
wymaga połączenia z bazą danych. Wymieniona technika w takim przypadku
zaleca zasymulować działanie tej bazy poprzez zwrócenie oczekiwanej wartości,
eliminując konieczność instalacji, połączenia i odpytywania 
produktu bazodanowego. Wykorzystanie atrapy obiektu redukuje złożoność
testu i czyni go stabilniejszym, gdyż nie wywołuje kaskady innych zapytań,
funkcji, połączeń sieciowych mogących ulec awarii.

Stosowanie tej techniki ma następujące korzyści:

    #. Zastosowanie reguły DIP

        Stosowanie atrap nakazuje programiście zastanowienie się, czy dany
        komponent powinien bazować na zaprogramowanej na sztywno
        infrastrukturze. Dobrą praktyką jest zastosowanie reguły DIP, czyli
        odwrócenie zależności. W takim przypadku, testowany komponent
        uzależniany jest od interfejsu, a nie od konkretnej klasy, co pozwala
        na rozszerzalność kodu.

    #. Możliwość odtworzenia rzadko występujących błędów/awarii

        Wprowadzenie atrap pozwala na zastosowanie "awarii na żądanie".
        Zakładając, iż testowany komponent wymaga połączenia z serwerem,
        technika Mock Object pozwala na implementację odpowiedzi serwera
        z wybranym kodem błędu. Możliwe staje się tym samym przetestowanie
        funkcjonalności w sytuacjach wyjątkowych, które innym sposobem jest
        ciężko odworzyć.

    #. Łatwość użycia z nowoczesnymi językami i bibliotekami

        Technika obiektu atrapy jest łatwo stosowalna z nowoczesnymi językami
        udostępniającymi refleksję np. Java i framework `Mockito <https://site.mockito.org/>`_.
        Stworzenie mock-objectu ogranicza się do zastosowania kilku linii kodu.

Czasem może wystąpić potrzeba przetestowania funkcjonalności testowanej
funkcji pod względem *stanu* obiektu do którego ta funkcjonalność posiada
referencję. Innymi słowy, oczekuje się, iż testowana funkcjonalność zmieni
stan innego obiektu. Technikę tę nazywa się *Stubbing*. Do badania stanu
końcowego, wykorzystuje się zaimlementowany obiekt o nazwie stub, co
oznacza, cytując:

    "**Stub** zapewnia odpowiedź do zapytania utworzonego podczas testów,
    nie odpowiadając innym zapytaniom, poza tymi zaprogramowanymi w teście"
    -- źródło: [#mocksArentStubs]_

Cykl testowania wykorzystujący Stub'y może przyjąć następującą postać
(za `źródłem <https://stackoverflow.com/a/17810004/11084875>`_):

    #. Konfiguracja - przygotowanie obiektu testu i stworzenie instancji klasy implementującej stan Stub'a

    #. Wykonanie - wykonanie testowanego komponentu

    #. Weryfikacja stanu - sprawdzenie, czy stan obiektu Stub'a jest zgodny z oczekiwaniem

    #. Zniszczenie obiektu testowego - uruchomienie destruktorów lub Garbage Collectora

Narzędzia CI/CD
--------------------------------------------------------------------------------

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

`Dockerfile <https://docs.docker.com/engine/reference/builder/>`_ 
jest sposobem na stworzenie obrazu Dockera. Jest to plik tekstowy deklarujący
postać obrazu.

Do stworzenia obrazu należy wykorzystać dostępne słowa dla języka plików
Dockerfile. Najczęściej przybiera to postać trzech kroków: 

    #) wybór obrazu bazowego

    Obraz bazowy pozwala na stworzenie finalnego obrazu dostosowanego do potrzeb
    użytkownika. W tym celu wykorzystuje się słowo kluczowe "FROM" oraz
    nazwę obrazu (z opcjonalną jego wersją - tagiem). Należy pamiętać, by 
    budowany obraz był jak najmniejszy oraz spełniał 
    `dobre praktyki <https://docs.docker.com/develop/develop-images/dockerfile_best-practices/>`_

    #) przygotowanie środowiska

    W ramach przygotowania środowiska należy doinstalować wybrane pakiety
    oprogramowania z użyciem menadżera pakietów dostępnego dla wybranego
    obrazu Dockera. Kolejnym krokiem jest przygotowanie aplikacji, skopiowanie
    kodu źródłowego i jego ewentualna kompilacja.

    W przypadku budowania zaawansowanych aplikacji zaleca się stosowania
    techniki `multi-stage build <https://docs.docker.com/develop/develop-images/multistage-build/>`_
    dla redukcji rozmiaru ostatecznego obrazu. Do budowy obrazu wykorzystuje
    się m.in. polecenia: RUN - uruchamia skrypt powłoki, COPY - kopiuje pliki.

    Podstawowym założeniem jest jeden stworzenie jednego obrazu dla
    jednej aplikacji. Jest to dobrą praktyką, która później ułatwia
    uruchomienie oprogramowania w systemie orkiestratorowanym np.
    w Kubernetesie.

    #) uruchomienie aplikacji

    Wybranie domyślnej aplikacji, która powinna być wykonywana przy starcie
    kontenera. W tym celu wykorzystuje się jedno z poleceń: CMD, ENTRYPOINT.

Ostatecznie, obraz przyjmuje podobną postać do poniższej::

    FROM alpine:3.7

    RUN apk update --purge && \
        apk add --purge --no-cache bash

    CMD ["/usr/bin/env", "bash", "-c", "echo 'Hello World'"]

Po zbudowaniu obrazu z pomocą polecenia (Dockerfile musi się znajdować
w katalogu dostępnego w $(pwd) )::

    $ docker build -t test-img .

Docker buduje obraz, który może następnie zostać uruchomiony::

    $ docker run -it --rm test-img
    Hello World


Docker-compose
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

`Docker Compose <https://docs.docker.com/compose/>`_
pozwala połączyć i uruchomić aplikację składającą się z wielu kontenerów.
Docker-compose nie jest dostępny domyślnie po instalacji Docker Engine na
komputerze. Instalacja jest wyjątkowo prosta na każdej z dystrybucji linuksa
i ogranicza się do::

    # pip install docker-compose

Docker-compose pozwala na szybkie przetestowanie gotowej aplikacji. Owa szybkość
bierze się z automatycznej konfiguracji sieci prywatnej dla i przypisywania
do niej każdego z kontenerów. Pozwala również na proste określenie adresu
contenera poprzez wykorzystanie nazwy usługi.

Poniższy docker-compose.yml pozwala na uruchomienie serwera NGINX i wysłanie
10 zapytań::

    version: '3'
    services:
        website:
            image: nginx:stable-alpine
            ports:
                - "80:80"
        client:
            image: "centos:7"
            command:
                - "/usr/bin/env"
                - "bash"
                - "-c"
                - "for _ in {1..10}; do curl website:80; sleep 1; done"

Tak skonstruowaną aplikacje uruchamia się za pomocą polecenia::

    $ docker-compose up

Efektem będą wyświetlane logi z domyślnej strony serwera NGINX. Docker-compose
automatycznie uruchomił dwa kontenery/usługi o nazwach "website" i "client".
Sieć pomiędzy kontenerami została skonfigurowana w taki sposób, że "client"
może odwoływać się do towarzyszących mu kontenerów poprzez nazwę usług,
co udowadnia polecenie "curl website:80".

Powyżej skonstruowana aplikacja nadal działa z poziomu systemu operacyjnego
hosta. Warto zaznaczyć, iż format pliku docker-compose.yml umożliwia
stosunkowo bezproblemową integrację z Docker Swarm, rozproszonego orkiestratora
kontenerów. W dalszej części powyższa aplikacja uruchomiona zostanie
w clustrze Kubernetesa.

Kubernetes
````````````````````````````````````````````````````````````````````````````````

`Kubernetes <https://kubernetes.io/>`_ (w skrócie K8s) jest otwartym systemem 
zarządzania aplikacji skontenteneryzowanych. Umożliwia on szereg czynności
jak zarządzanie połączeniem sieciowym, montowaniem zasobów dyskowych
w systemie rozproszonym, monitorowaniem obciążenia, skalowania i czuwania
nad stanem kontenerów i inne.

Architektura Kubernetesa jest skomplikowana i zaleca się zapoznanie poprzez
`dokumentację <https://kubernetes.io/docs/concepts/overview/components/>`_.
K8s wyróżnia następujące elementy:

    * Dla węzła głównego (master), m.in.:

      * *kube-apiserver*:  udostępnia Rest API
      * *etcd*: baza danych klucz-wartość
      * *kube-scheduler*: scheduler

    * Dla pozostałych węzłów (nodes), m.in.:

      * *kubelet*: agent uruchamiający *Pod*\y na danym węźle
      * *kube-proxy*: zarządzanie połączeniem sieciowym
      * *container-runtime*: jeden z - Docker, containerd, rkt, cri-o

Ze względu na złożoność architektury Kubernetesa, również jego instalacja
nie jest zadaniem trywialnym. Należy pamiętać o wielu zależnościach,
zabezpieczeniach etc. Poleca się stosowanie gotowych systemów chmurowych.
Niemniej jednak istnieją instalatory upraszczające znacząco instalację
klastra, m.in. `kubeadm <https://kubernetes.io/docs/setup/independent/create-cluster-kubeadm/>`_,
`kubespray <https://github.com/kubernetes-sigs/kubespray>`_,
`kops <https://github.com/kubernetes/kops>`_ i inne. Istnieją również
uproszczone metody uruchomienia środowiska pod postacią
`minikube <https://kubernetes.io/docs/tasks/tools/install-minikube/>`_,
`k3s <https://k3s.io/>`_,
`microk8s <https://microk8s.io/>`_ i inne systemy. Wybór właściwego zależy
od potrzeb użytkownika i złożoności posiadanej infrastruktury.

Obiekty
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

`Obiektem <https://kubernetes.io/docs/concepts/overview/working-with-objects/kubernetes-objects/>`_
w Kubernetesie jest każda intencja przekazana do klastra. Taką intencją może
być chęć uruchomienia konkretnego obrazu, konfiguracja usługi sieciowej,
przechowywanie zaszyfrowanych danych wrażliwych, konfiguracja przestrzeni
użytkownika, akcje dla cron'a, itp.

Obiekty najczęściej definiuje się z użyciem plików w formacie `YAML <https://yaml.org/>`_.
Podstawowy schemat jest następujący::

    apiVersion: <<wersja api w zależności od porządanej funkcjonalności>>
    kind: <<rodzaj obiektu>>
    metadata:
      name: <<unikalna nazwa obiektu>>
      labels: <<etykiety umożliwiające identyfikację zadań przez K8s>
        key: value
    spec: <<specyfikacja obiektu, zależy od jego rodzaju>>

Niezwykle istotnym elementem w systemie Kubernetesa są `label'e (etykiety) <https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/>`_.
Umożliwiają one identyfikację i grupowanie zadań. Grupowanie zapewnia
element K8s'a zwany selektorem (label selector).

.. note::
    Poniższy opis dostępnych obiektów w K8s zapewnia jedynie minimalną
    wiedzę z zakresu działania narzędzia. Zaleca się zapoznanie w pełni
    z dokumentacją.

Pod
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

`Pod <https://kubernetes.io/docs/concepts/workloads/pods/pod-overview/>`_ jest
podstawowym obiektem w Kubernetesie zapewniającym działania i uruchomienie
kontenerów. Pod jest abstrakcją kontenera w Kubernetesie. Pozwala jednakże
na uruchomienie wielu kontenerów jednym Podzie.

Uruchomienie prostego standardowego hello-world odbywa się w następujących
krokach:

    #. Stworzenie pliku YAML z następującą treścią (wcięcia są istotne!)::

        apiVersion: v1
        kind: Pod
        metadata:
          name: hello-world-pod
          labels:
            app: hello-world
        spec:
          containers:
          - name: hello-world-container
            image: hello-world

    #. Wykonanie polecenia::
    
        $ kubectl create -f nazwa_pliku.yml

    #. Sprawdzenie logów::

        $ kubectl logs hello-world-pod

Wraz z zakończeniem działania Poda, nie zostaje on uruchomiony ponownie.
Ma to istotne znaczenie, o ile próbujemy zapewnić pewne działanie aplikacji.
Aby zapewnić ciągłość działania, należy wykorzystać obiekty wyższego rzędu
zapewnione przez K8s.


Deployment
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

`Deployment <https://kubernetes.io/docs/concepts/workloads/controllers/deployment/>`_
jest obiektem umożliwiającym zdefiniowanie porządanego stanu aplikacji. Po
utworzeniu tego obiektu, automatycznie on zarządza ilością replik wybranego
obrazu pilnując by zawsze określona ich ilość była funkcjonalna.

Doskonałym zastosowaniem tego obiektu jest uruchomienie aplikacji
bezstanowej np. prostego front-endu który komunikuje się z backendem. Dodatkowo,
umożliwia on wykonywanie czynności niezwykle istotnych z punktu widzenia
HA (High Availability) m.in. rolling deployments czy skalowanie.

Przykład uruchamiający serwer NGINX jest następujący::
    
    apiVersion: apps/v1
    kind: Deployment
    metadata:
      name: nginx-deployment
      labels:
        app: nginx-server
        purpose: test-hello-world
    spec:
      replicas: 1
      selector:
        matchLabels:
          app: nginx
          purpose: test-hello-world
      template:
        metadata:
          labels:
            app: nginx
            purpose: test-hello-world
        spec:
          containers:
          - name: nginx-container
            image: nginx
            ports:
            - containerPort: 80

Tym sposobem uruchomiony zostaje serwer nginx. Nie został on jednak
udostępniony światu. Nie można w rozsądny sposób skomunikować się 
z serwerem. W tym celu należy stworzyć nowy obiekt zwany: Service.

Service
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

`Service (Usługi) <https://kubernetes.io/docs/concepts/services-networking/service/>`_
zapewnia dostęp do kontenerów poprzez sieć. Jest to o tyle istotne, iż kontenery
w swej naturze są śmiertelne, tj. mogą zostać zniszczone. Nie byłoby zalecane
ręczne infrastruktury sieciowej manualnie. Usługi automatyzują dynamiczną
konfigurację sieci pomiędzy kontenerami a światem zewnętrznym.

Ich działanie, w głównej mierze opiera się na wykorzystaniu i dopasowaniu
etykiet. Z pomocą właśnie etykiet Usługa wie, dla których Pod'ów należy
skonfigurować DNSy i udostępnić właściwe porty.

Przykład współpracujący z Deployment'em określonym powyżej::

    apiVersion: v1
    kind: Service
    metadata:
      name: nginx-service
      labels:
        app: nginx
        purpose: test-hello-world
    spec:
      selector:
        app: nginx
        purpose: test-hello-world
      ports:
      - protocol: TCP
        port: 80
        targetPort: 80

Tworzy on obiekt Service typu ClusterIP (domyślny). Umożliwia on
dostęp do strony kontenera pod stałym adresem IP, niezależnie
od obecnego adresu IP kontenera. Aby uzyskać udostępniony
adres IP, należy użyć poniższego polecenia::

    $ kubectl get service nginx-service -o jsonpath='{ .spec.clusterIP }'

Po uzyskaniu adresu można połączyć się z wybraną aplikacją pod warunkiem,
że znajduje się na tam, gdzie został zainstalowany klaster.

Jenkins
````````````````````````````````````````````````````````````````````````````````

`Jenkins <https://jenkins.io/>`_ jest otwarty system automatyzacji umożliwiający
w łatwy sposób wdrożenie CI/CD.

Jenkins pozwala na stworzenie rodzaju linii produkcyjnej dla oprogramowania,
w której oprogramowanie jest budowane, testowane i na samym końcu dostarczane
do klienta. W celu zapewnienia szeregu użytecznych cech i rozszerzalności
narzędzie wprowadza możlwiość stosowania dodatkowych pluginów, dzięki którym
wprowadza się nowe funkcjonalności do systemu.

.. warning::
    Zobacz tutaj: https://www.stratoscale.com/blog/devops/practical-devops-use-case-github-jenkins-docker/

Instalacja
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Jenkinsa zainstalować można na platformach obsługujących maszynę wirtualną
Javy. Oznacza to, iż możliwa jest instalacja na gołym metalu (bare-metal).
Niemniej jednak, instalacja z użyciem poniższych środowisk pozwala na znacznie
szybszą konfigurację i jest wstępem do rozwiązania chmurowego.

**Docker**

Instalacja z pomocą Docker'a ogranicza się praktycznie do jednolinijkowego
polecenia. W wersji minimalistycznej, poleceniem tym jest::

    $ docker run -p 8080:8080 -p 50000:50000 jenkins/jenkins:lts

Powyższa instalacja jest niemniej jednak dość naiwna, gdyż nie zachowuje
konfiguracji w systemie plików. Dokładniejsze informacje znajdują się
w `dokumentacji <https://github.com/jenkinsci/docker/blob/master/README.md>`_.

**Kubernetes i Helm**

Jednym ze sposobów instalacji Jenkinsa w Kubernetesie jest wykorzystanie
menadżera pakietów, `Helm <https://helm.sh/>`_.

Zakładając, iż Helm został poprawnie zainstalowany i skonfigurowany,
instalacja Jenkinsa ogranicza się do edycji pliku `konfiguracyjnego <https://github.com/helm/charts/blob/master/stable/jenkins/values.yaml>`
i wykonania poniższej instrukcji::

    $ helm install --name jenkins stable/jenkins -f confg_file.yml

Opis dostępnych parametrów znajduje się w `dokumentacji <https://github.com/helm/charts/tree/master/stable/jenkins>`_.
Niezbędne jest również udostępnienie przestrzeni dyskowej działającej 
w systemie rozproszonym np. NFS. Taki zasób należy następnie zamontować
w Kubernetesie za pomocą `PersistentVolume/PersistentVolumeClaim <https://kubernetes.io/docs/concepts/storage/persistent-volumes/>`_ albo z pomocą `StorageClass <https://kubernetes.io/docs/concepts/storage/storage-classes/>`_.

Pipeline CI/CD
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

#. Commit

   Programista tworzy kod. Wraz z wysłaniem kodu do repozytorium, Jenkins
   zostaje poinformowanych o zmianach za pomocą `webhook-a <https://en.wikipedia.org/wiki/Webhook>`_.
   Jenkins przechodzi do kolejnego kroku.

#. Zbuduj

   W tym etapie budowana jest aplikacja, co najmniej w wersji do testów.
   Etap budowania jest również swego rodzaju testem, sprawdzający integralność
   projektu sugerującym, że kod po ściągnięci z repozytorium również skompiluje
   się na maszynie lokalnej. W przypadku nieudanego "builda", właściwie
   skonfigurowany Jenkins poinformuje winowajcę, team i menadżera..., dzięki
   czemu na przyszłość kod stanie się lepszy.

   Budowanie na tym etapie zalecane jest do zbudowania kontenera
   i ewentulane udostępnienie go we własciwym rejesetrze 
   (np. `Docker Hub <https://hub.docker.com/>`_).

#. Uruchom testy

   Posiadając zbudowany kontener ze zbudowaną aplikacją do testów, pozostaje
   uruchomić testy przez Jenkinsa. Obdywa się to w sposób automatyczny.
   W zależności od rezultatów, Jenkins pozwoli na dalsze budowanie projektu.

   Zupełnym minimum produkcyjnym jest uruchomienie testów jednostkowych.
   Nie jest to wystarczający zestaw testów. W tak skonfigurowanym łańcuchu
   testów należy zbadać również np. testy systemowe zostały uruchomione na
   pewnym etapie. Zależy to oczywiście od rodzaju projektu a przede
   wszystkim od norm obowiązującyh w danym przedsiębiorstwie.

#. Release
   
   W tej fazie, Jenkins buduje kontener z wersją aplikacji gotową do
   użytkowania w systemie. 

#. Deploy/Deliver

   Zakładając pomyślnie przejście wcześniejszych etapów, aplikację
   można zainstalować w zasobach docelowych.

#. "Dalsze testy na produkcji" :)  

   Powyższe humorystyczne stwierdzenie, dość często spotykane - niestety,
   prezentuje jednak dość ważny fakt. Nie ważne jak dobrze przetestowane
   jest oprogramowanie, finalne testy przeprowadzane są dopiero
   w środowisku docelowym klienta. Dopiero w środowisko docelowe
   pozwala na pełne przetestowanie software'u zgodnie ze wszystkimi
   scenariuszami. Kompania testowa pozwala w znaczącym stopniu 
   zredukować możliwości wystąpienia problemów.

.. rubric:: Referencje

.. [#martin_agile] Martin Robert C., *Zwinne wytwarzanie oprogramowania. Najlepsze zasady, wzorce i praktyki*, Helion, 2015 - str. 47
.. [#endotesting] McKinnon T., Freeman S., Craig P., *Endo-Testing: Unit Testing with Mock Objects*, 2000
.. [#mocksArentStubs] Cokelaer T, *Mocks Aren't Stubs'*, https://martinfowler.com/articles/mocksArentStubs.html, z dn. 2007-01-02
