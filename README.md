## This Go application running in a Docker container automatically runs OCR on the PDF files in a mounted folder.

1. Download this repository
1. Modify docker-compose.yml 
1. (Optional) Add languages in the languages folder
1. Start the container with `docker compose up`

<br/>

### Filename format
`DoOcr-deu+eng-XXXX.pdf`  
`1⮥------2⮥-----3⮥`

1. `DoOcr` Part that triggers the script  
1. `deu+eng` Languages for the ocr (Must be added during build)  
1. `XXXX` User specified part  

Example:  
`DoOcr-deu+eng-blablup.pdf`

<br/>
<br/>

Credits to [OCRmyPDF](https://github.com/ocrmypdf/OCRmyPDF) and [tesseract](https://github.com/tesseract-ocr/tesseract).
