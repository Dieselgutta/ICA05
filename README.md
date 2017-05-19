# ICA05
Server på Ubuntu ved http://158.39.77.237:8001/

* Deltakere: Simen Fredriksen, Stian Blankenberg, Jone Manneråk, Kristian Hagberg, Tarjei Taxerås og Rasmus Sørby*


Tankegang for kode: 
Golang:
Vi satte opp en enkel webserver ved å bruke HandleFunc og ListenAndServe på port 8001. 
vi brukte en funksjon som laget klienten i func basicHandler. Denne sørger for 
at vi får en respons på siden og den knytter template/index.html til serveren for å få
det opp som en nettside. 
Vi brukte en API fra OpenWeatherMap som vi dekodet i Func decode. Denne decoder
dataen som blir hentet ut fra funksjonen func getData. Denne dataen er da tilgjengelig, 
og vi bruker den til å sette opp "structs". De forskjellige parameterne som er gruppert
i api-en blir samlet i en struct som videre blir samlet i en struct kalt "Weather". 
Her blir også variablene som ikke har en gruppa lagt inn (eks. Name string). 
I API-en er temperaturen definert ved Kelvin-måleenheten. Dette blir gjort om til Celsius
ved å ta temperaturen minus det absolutte nullpunktet(-273.15). 

HTML: 
Index.html blir brukt som en mal (template) for nettsiden. Her bruker man diverse
HTML-koder for å utforme sidens innhold. Vi bruker et stylesheet fra w3schools.com for å definere 
de forskjellige HTML-kodene som vi bruker. Dette innebærer endringer som farge og andre designmuligheter. 
HTML-siden kaller på Golang sine structs for å hente ut/vise fram informasjon hentet fra API-en. 
Vi har også embedded et kart fra google maps som bruker Name og Country fra Struct for å definere
lokasjon. Denne ble enkelt generert ved hjelp av http://www.embedgooglemap.net/. 



ekstra kommentar: Deler av koden er basert på gruppe 12 sin kode. 
