================================================================================
Wprowadzenie
================================================================================

Architektura mikroserwisów
--------------------------------------------------------------------------------

Mikroserwisy (lub mikrousługi) to pomniejsze fragmenty programu, które
są ze sobą w relacji i tworzą razem większą aplikację. Każdy z nich ma
własną unikatową budowę i logikę. Owa niezależność pozwala na niezależne
rozwijanie każdej z usług w aplikacji, pojedyncze wdrażanie dodatkowych
usług do programu, a także upraszcza jego budowę. Nie jest ona jednak
pozbawiona wad - modyfikacji programu trudniej dokonać globalnie, gdyż
konieczne jest wprowadzenie zmian w każdej z komórek. Może pojawić się
również problem zgodności z modelem danych organizacji wykorzystującej
program oraz redundancja danych.

Testowanie oprogramowania
--------------------------------------------------------------------------------

Kiedy projektujemy urządzenie, chcemy być pewni, że zadziała zgodnie z
jego przeznaczeniem oraz nie ulegnie awarii. Podobnie jest z programem
komputerowym - testowanie jest kluczowym elementem tworzenia kodu - na
każdym etapie procesu.

Model piramidowy
````````````````````````````````````````````````````````````````````````````````

W modelu tym wyróżniamy 3 rodzaje testów, które ułożone są wg hierarchii
wykonywania. Na samym dole znajdują się testy jednostkowe -
przeprowadzane na pojedynczych fragmentach kodu, nie sprawdzają
interakcji między różnymi elementami programu - są najszybsze w
wykonaniu i najczęściej stosowane (ok. 70% wszystkich testów).

Pośrodku mamy testy integracyjne. Sprawdzają one, jak poszczególne
komponenty współpracują między sobą i z usługami zewnętrznymi. Stanowią
w przybliżeniu 20% testów.

Na samym szczycie piramidy są testy akceptacyjne (end-to-end). Badana w
nich jest pełna funkcjonalność programu od początku do końca działania,
a także działanie interfejsu użytkownika. Circa 10% testów to testy
akceptacyjne.

.. figure:: /_static/model_piramidowy.jpg
   :alt: Model piramidowy - schemat

   Model piramidowy - schemat

Testy jednostkowe
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Testy weryfikujące poprawność działania pojedynczych fragmentów
(jednostek) programu nazywamy testami jednostkowymi. Polegają one na
wykonywaniu małych fragmentów programu i porównywania wyników ich
działania z oczekiwanymi wynikami. Testy jednostkowe mogą być
zautomatyzowane tak, by wykonywały się na bieżąco przy wprowadzaniu
jakichkolwiek modyfikacji do programu. Istnieją różne podejścia do tego
typu testów:

-  **analiza ścieżek** - określany jest punkt początkowy i końcowy do
   przeprowadzenia testów, a następnie badany jest przebieg możliwych
   ścieżek pomiędzy nimi,
-  **użycie klas równoważności** - jeśli testowi podlega zbiór
   wieloelementowy, to można z niego wybrać próbę reprezentatywną, po
   sprawdzeniu której można założyć, że te same rezultaty będą zachodzić
   dla całego zbioru,
-  **testowanie wartości brzegowych** - rozwinięcie wyżej wspomnianego
   punktu, testy koncentrują się na punktach w pobliżu wartości
   brzegowych,
-  **testowanie składniowe** - jego celem jest sprawdzenie poprawności
   danych wprowadzanych do systemu.

Testy integracyjne
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Po przetestowaniu poszczególnych komponentów należy upewnić się, że
relacje między komponentami są prawidłowo zbudowane, a także nie
zachodzą żadne błędy w ich wzajemnych interakcjach. Wyróżniamy
następujące podejścia do testów integracyjnych:

-  **top - down**- jako pierwsze sterowane są moduły znajdujące się na
   najwyższym poziomie, w tym czasie niższe moduły zastępowane są przez
   zaślepki. Przetestowane moduły używane są następnie do niżej
   położonych komponentów, aż przetestowane są fragmenty programu na
   najniższym poziomie.
-  **bottom - up -**\ przeciwieństwo powyższego podejścia, procesy
   powyżej aktualnie testowanych zastępowane są tzw. driverami. Najpierw
   testowane są komponmnty
-  **big bang -**\ większość opracowanych modułów jest łączona ze sobą,
   by stworzyć kompletny system wykorzystywany do testowania integracji.
   Testowanie systemu jako całość powoduje jednak, że błędy występujące
   w interfejsach komponentów mogą być wykryte zbyt późno lub niewykryte
   w ogóle, trudniejsze jest również zlokalizowanie przyczyn błędów
   wykrytych.

Testy akceptacyjne
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Są to końcowe testy przeprowadzane w celu potwierdzenia jakości
tworzonego oprogramowania przed wprowadzeniem go na rynek bądź przed
udostępnieniem dla zleceniodawcy. W ich ramach mogą być przeprowadzone
testy alfa (testy wykonywane w firmie, która wyprodukowała program, ale
przez inny zespół niezależny od deweloperów programu) oraz testy beta
(przeprowadzane przez podmioty zewnętrzne). Według prawa w Polsce przez
testy akceptacyjne należy rozumieć zbiór danych pozwalający na
stwierdzenie poprawności współpracy systemów informatycznych,

Smoke & Sanity testing
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Smoke test to sprawdzenie, czy program uruchamia się i reaguje na
działania użytkownika. Testuje on najbardziej krytyczne funkcje programu
bez zagłębiania się w detale oprogramowania. Jeśli smoke test
przebiegnie pomyślnie, można przejść do testu sanity. W tymże sprawdzane
są pojedyncze funkcjonalności programu oraz zgodność logiki aplikacji z
wymaganiami. Oba te rodzaje są testami ramowymi, obejmującymi szeroki
zakres i nie zagłębiającymi się w szczegóły działania testowanej
funkcjonalności.

Testowanie biało- i czarnoskrzynkowe
````````````````````````````````````````````````````````````````````````````````

Testy białoskrzynkowe (strukturalne) polegają na podawaniu takich danych
wejściowych, by program przeszedł przez każdą zaimplementowaną ścieżkę.
Są one używane do dokładnego sprawdzenia operacji wykonywanych w
zaimplementowanych metodach. Program jest testowany na poziomie kodu
źródłowego.

Z kolei testy czarnoskrzynkowe (funkcjonalne) wykonywane są przez osoby
nie znające budowy programru. Mają one na celu sprawdzenie, jak
potencjalny użytkownik będzie sobie radził z obsługą programu i czy
program prawidłowo wykonuje swoje funkcje. Łatwo za pomocą takiego testu
wykryć istnienie błędu, ale jego przyczyna może pozostać nieznana.

Testowanie FIT - Fault Injection Testing
````````````````````````````````````````````````````````````````````````````````

Jest to celowe wywoływanie (wstrzykiwanie) błędów w oprogramowaniu.
Pozwala to na sprawdzenie mechanizmów error handlingu, a także
zapobieganie ewentualnym awariom.

Metoda ta ma na celu sprawdzenie, jak program wykrywa błędy i
minimalizuje ich szkodliwość. Przykładowo programiści Netflixa
wpuszczają ‘małpki’ do programu, które mają przeszkadzać w jego
prawidłowej egzekucji. Jedną z takich małpek jest sztuczne wydłużenie
czasu odpowiedzi serwera - pozwala to przygotować program na problemy z
łącznością. Ważne jest, by do programu wprowadzać błędy, które nie będą
miały trwałego szkodliwego wpływu na projekt - celem deweloperów nie
jest wprowadzenie chaosu w programie, a błędy “wstrzykiwane” są bardzo
precyzyjnie. FIT pomaga też w budowaniu narzędzi do szybszej i
trafniejszej diagnozy błędów oraz automatyzacji testów.

TDD - Test driven development
--------------------------------------------------------------------------------

Jest to technika, w której programowanie zaczyna się od skryptów
testujących jeszcze nieistniejący element programu. Na początku program
nie powinien przechodzić testów - jest on jednak stopniowo rozwijany
tak, by dostosować się do wymagań testowych. Podejście to może być
używane do tworzenia nowych programów, a także rozwijania istniejących
kodów. Wadą takiego podejścia jest konieczność napisania testów przed
rozpoczęciem właściwego developmentu - jednak zasoby przeznaczone na
tworzenie tych testów zwrócą się, gdy błędy w programie zostaną wykryte
natychmiast po kompilacji. Z racji na cykliczną i iteratywną naturę TDD,
można je stosować przy metodyce Agile, gdzie projekt jest podzielony na
części realizowane w kolejnych sprintach, a po każdym sprincie zgodnie z
założeniami następuje faza testów oraz planowania dalszej pracy.

