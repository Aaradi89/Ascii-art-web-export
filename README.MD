Description : 
* An HTML server that convert a normal text to an !!! ART !!! using a pre-defined banners.



Authors :
* Ali Alaradi (alalaradi)
* Ali Alaradi (alialaradi0)
* Ali Meshaimea (ameshaim)



Usage: how to run
* go run .



Implementation details: algorithm

* Server is build using GO, HTML & CSS.
* Home page is bulit in html & CSS.
* In the homepage, the user will add one value (InputText) & choose the required banner name (banner).
* These values will be sent to Go server and convert the text into ASCII art.
* Using POST method, go server will take the final results and write it in a new page in the html server (/ascii-art).
* User will be able test on the same page.

* Invalid input (Not supported language) will give an error -  500 Input Error. 
* Navigating between unavailable webpages will give an error - 404 Error Page is not found.
* Navigating to (/ascii-art) without input will give an error - 400 Error Bad Request.