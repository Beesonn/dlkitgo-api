# dlkitgo API

<div align="center">
  
![GitHub stars](https://img.shields.io/github/stars/Beesonn/dlkitgo-api?style=for-the-badge)
![GitHub forks](https://img.shields.io/github/forks/Beesonn/dlkitgo-api?style=for-the-badge)
![License](https://img.shields.io/badge/license-MIT-blue?style=for-the-badge)

**An ultra-fast API for downloading, searching, and fetching information from popular platforms and third-party services.**

</div>

## Endpoints

### Spotify 
| Action | Endpoint |
| :--- | :--- |
| **Stream** | `/spotify/stream?url={url}` |
| **Info** | `/spotify/info?url={url}` |
| **Search** | `/spotify/search?q={song_name}` |

### Instagram 
| Action | Endpoint |
| :--- | :--- |
| **Stream** | `/instagram/stream?url={url}` |
| **Info** | `/instagram/info?url={url}` |

### YouTube 
| Action | Endpoint |
| :--- | :--- |
| **Stream** | `/youtube/stream?url={url}` |

---

## Installation

### Run Locally

To run this API on your local machine, ensure you have **Go** installed, then run the following commands:

```
git clone https://github.com/Beesonn/dlkitgo-api
cd dlkitgo-api
go build .
./dlkitgo-api
```
