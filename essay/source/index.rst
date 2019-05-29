.. Testy jednostkowe w aplikacjach rozproszonych. documentation master file, created by
   sphinx-quickstart on Tue May 28 19:16:40 2019.
   You can adapt this file completely to your liking, but it should at least
   contain the root `toctree` directive.

Testy jednostkowe w aplikacjach rozproszonych.
==========================================================================

Wstęp
--------------------------------------------------------------------------------

Celem dokumentu jest zaprezentowanie możliwych technik testowania
jednostkowego aplikacji rozproszonych z naciskiem na testowanie
mikrousług.

Teza
--------------------------------------------------------------------------------

Tworzenie aplikacji jest skomplikowanym procesem podatnym na błędy ludzkie.
Jest to w szczególności widoczne w postaci systemów wielowątkowych i rozproszonych.
Bezpieczne zarządzanie danymi w takich systemach jest wyzwaniem. Wytwarzanie
aplikacji w architektrzue mikrousług kładzie szczególny nacisk na wymianę
danych pomiędzy komponentami danych. Komponenty te, wytwarzane przez oddzielne
zespoły, a nawet firmy, muszą ze sobą współgrać, by ostatecznie dostarczyć
gotowy produkt w postaci aplikacji.

Aspekt komunikacji pomiędzy komponentami (usługami) jest niezwykle istotny
a jednocześnie trudny to testowania. Stawia się tezę, iż aby dostarczyć
produkt dostatecznej jakości, testy jednostkowe poszczególnych usług są
niewystarczające. 

.. toctree::
   :maxdepth: 2
   :caption: Spis treści:

   content/01_wprowadzenie
   content/02_techniki_i_narzedzia
   content/03_uruchamianie_testow
   content/04_wnioski
