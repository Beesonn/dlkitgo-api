# dlkitgo API

<div align="center">
  
![GitHub stars](https://img.shields.io/github/stars/Beesonn/dlkitgo-api?style=for-the-badge)
![GitHub forks](https://img.shields.io/github/forks/Beesonn/dlkitgo-api?style=for-the-badge)
![License](https://img.shields.io/badge/license-MIT-blue?style=for-the-badge)

**Ultra-fast API** for downloading, searching, and fetching information from popular platforms and third party

</div>

## Endpoints

### Spotify 
* /spotify/stream?url=url
* /spotify/info?url=url
* /spotify/search?q=song

### Instagram 
* /instagram/stream?url=url
* /instagram/info?url=url

### YouTube 
* /youtube/stream?url=url

## Locally

```bash
git clone https://github.com/Beesonn/dlkitgo-api
cd dlkitgo-api
go build .
./dlkitgo-api
```
