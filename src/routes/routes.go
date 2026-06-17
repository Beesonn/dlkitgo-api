package routes

import (
	"github.com/Beesonn/dlkitgo-api/src/routes/instagram"
	"github.com/Beesonn/dlkitgo-api/src/routes/pinterest"
	"github.com/Beesonn/dlkitgo-api/src/routes/spotify"
	"github.com/Beesonn/dlkitgo-api/src/routes/youtube"
	"github.com/gin-gonic/gin"
)

func SetupAllRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		c.String(200, `<!DOCTYPE html>
<html>
  <head>
    <title>dlkitgo API Reference</title>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=yes" />
    <meta name="description" content="Ultra-fast API for downloading, searching, and fetching information from popular platforms." />
    <meta name="theme-color" content="#000000" />
    <style>
      * {
        margin: 0;
        padding: 0;
        box-sizing: border-box;
      }
      html, body {
        width: 100%;
        height: 100%;
        overflow-x: hidden;
      }
      body {
        font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, sans-serif;
        background: #0a0a0a;
        color: #ffffff;
        margin: 0;
        padding: 0;
      }
      #app {
        width: 100%;
        max-width: 100vw;
        min-height: 100vh;
        overflow-x: hidden;
      }
      .scalar-api-reference {
        width: 100% !important;
        max-width: 100% !important;
        overflow-x: hidden !important;
      }
      .scalar-api-reference * {
        max-width: 100% !important;
        word-wrap: break-word !important;
      }
      @media (max-width: 768px) {
        .scalar-api-reference {
          padding: 0 !important;
        }
        .scalar-api-reference .sidebar {
          width: 100% !important;
          max-width: 100% !important;
        }
        .scalar-api-reference .content {
          width: 100% !important;
          max-width: 100% !important;
          padding: 10px !important;
        }
      }
    </style>
  </head>
  <body>
    <div id="app"></div>
    <script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
    <script>
      Scalar.createApiReference('#app', {
        url: window.location.origin + '/openapi.json',
        proxyUrl: 'https://proxy.scalar.com',
        theme: 'dark',
        layout: 'modern',
        hideDownloadButton: false,
        hideTestRequestButton: false,
        defaultHttpClient: {
          target: 'fetch',
          proxyUrl: 'https://proxy.scalar.com',
        },
        searchHotKey: 'k',
        darkMode: true,
        showSidebar: true,
        hideModels: false,
        hideSchemaPattern: false,
        hideSecurity: true,
        hideInfo: false,
        hideLogo: false,
        hideSearch: false,
        hideExamples: false,
        hideRequestPayload: false,
        defaultOpenAllTags: false,
      })
    </script>
  </body>
</html>`)
	})

	r.GET("/openapi.json", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.String(200, `{
  "openapi": "3.0.0",
  "info": {
    "title": "dlkitgo API",
    "version": "1.0.0",
    "description": "Ultra-fast API for downloading, searching, and fetching information from popular platforms."
  },
  "paths": {
    "/spotify/search": {
      "get": {
        "summary": "Search Spotify",
        "description": "Search for tracks, artists, or albums on Spotify",
        "parameters": [
          {
            "name": "q",
            "in": "query",
            "required": true,
            "schema": {
              "type": "string"
            },
            "description": "Search query"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response"
          },
          "400": {
            "description": "Missing query parameter"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    },
    "/spotify/info": {
      "get": {
        "summary": "Get Spotify track info",
        "description": "Get detailed information about a Spotify track",
        "parameters": [
          {
            "name": "url",
            "in": "query",
            "required": true,
            "schema": {
              "type": "string"
            },
            "description": "Spotify track URL"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response"
          },
          "400": {
            "description": "Missing URL parameter"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    },
    "/spotify/stream": {
      "get": {
        "summary": "Stream Spotify track",
        "description": "Get stream sources for a Spotify track",
        "parameters": [
          {
            "name": "url",
            "in": "query",
            "required": true,
            "schema": {
              "type": "string"
            },
            "description": "Spotify track URL"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response"
          },
          "400": {
            "description": "Missing URL parameter"
          },
          "404": {
            "description": "No stream sources found"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    },
    "/youtube/search": {
      "get": {
        "summary": "Search YouTube",
        "description": "Search for videos on YouTube",
        "parameters": [
          {
            "name": "q",
            "in": "query",
            "required": true,
            "schema": {
              "type": "string"
            },
            "description": "Search query"
          },
          {
            "name": "lim",
            "in": "query",
            "required": false,
            "schema": {
              "type": "integer",
              "default": 0
            },
            "description": "Limit results (0 for unlimited)"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response"
          },
          "400": {
            "description": "Missing query parameter"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    },
    "/youtube/stream": {
      "get": {
        "summary": "Stream YouTube video",
        "description": "Get stream sources for a YouTube video",
        "parameters": [
          {
            "name": "url",
            "in": "query",
            "required": true,
            "schema": {
              "type": "string"
            },
            "description": "YouTube video URL"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response"
          },
          "400": {
            "description": "Missing URL parameter"
          },
          "404": {
            "description": "No stream sources found"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    },
    "/youtube/info": {
      "get": {
        "summary": "Get YouTube video info",
        "description": "Get detailed information about a YouTube video",
        "parameters": [
          {
            "name": "url",
            "in": "query",
            "required": true,
            "schema": {
              "type": "string"
            },
            "description": "YouTube video URL"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response"
          },
          "400": {
            "description": "Missing URL parameter"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    },
    "/instagram/stream": {
      "get": {
        "summary": "Stream Instagram media",
        "description": "Get stream sources for Instagram media (photo, video, reel)",
        "parameters": [
          {
            "name": "url",
            "in": "query",
            "required": true,
            "schema": {
              "type": "string"
            },
            "description": "Instagram post/reel URL"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response"
          },
          "400": {
            "description": "Missing URL parameter"
          },
          "404": {
            "description": "No stream sources found"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    },
    "/instagram/info": {
      "get": {
        "summary": "Get Instagram media info",
        "description": "Get detailed information about Instagram media",
        "parameters": [
          {
            "name": "url",
            "in": "query",
            "required": true,
            "schema": {
              "type": "string"
            },
            "description": "Instagram post/reel URL"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response"
          },
          "400": {
            "description": "Missing URL parameter"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    },
    "/pinterest/stream": {
      "get": {
        "summary": "Stream Pinterest media",
        "description": "Get stream sources for Pinterest pins",
        "parameters": [
          {
            "name": "url",
            "in": "query",
            "required": true,
            "schema": {
              "type": "string"
            },
            "description": "Pinterest pin URL"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful response"
          },
          "400": {
            "description": "Missing URL parameter"
          },
          "404": {
            "description": "No stream sources found"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    }
  }
}`)
	})

	spotify.SetupRoutes(r)
	youtube.SetupRoutes(r)
	instagram.SetupRoutes(r)
	pinterest.SetupRoutes(r)
}