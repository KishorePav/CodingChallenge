# CodingChallenge


API Coding Challenge Summary
A GraphQL API or REST API for storing metadata of music tracks to a Relational Database,
fetching the metadata from the external Spotify REST API.
Framework Requirements:
 Implement in GoLang using a lightweight node framework such (Gorilla Mux, Echo, or
Gin etc.)
 We would like to see an ORM in place too (GORM)
 If you opt for REST API, then to document the API you are creating, feel free to use
swagger.io.
Code Challenge API Details:
 Write access: There should be a single endpoint to trigger for the creation of a track,
which takes a single value “ISRC”
o Input is just the ISRC
 An ISRC is more or less an identifier of a music track, see “Spotify API”
section below

o As soon as the call comes in, use the Spotify API to fetch the following metadata:
 Spotify Image URI
 Title
 Artist Name List
o In case the SpotifyAPI returns multiple tracks take the track with highest
popularity (an attribute in the Json)
o Store the ISRC and the additional metadata into the DB
 No need to care about updating an already existing ISRC, skipping or give
back an error is enough.

o Read access: There should be 2 endpoints to retrieve metadata
o By “ISRC”: single result
o By “artist”: multiple results
 &quot;like&quot; search in the DB
 Multiple results (list/array)

Spotify API Details:
 You should use this API to retrieve the metadata of a music track which is identified by
an ISRC
 Create a developer account at Spotify: https://developer.spotify.com/
 As the challenge is about creating an API which uses the Spotify API, OAuth “Client
Credentials Flow” should be used as authorization:
https://developer.spotify.com/documentation/general/guides/authorization-guide/#client-c
redentials-flow
o No need to cache the token for the code challenge, it&#39;s ok to grab it on each API
Call
o Spotify API Endpoints you might use
o https://api.spotify.com/v1/search?q=thetitle&amp;type=track

 Search for track using the title
 You can use that call to find out ISRCs of music tracks
 Some sample ISRC you can use: USVT10300001,
USEE10001992,
GBAYE0601498, USWB11403680, GBAYE0601477
o Feel free to search for some by yourself, the get by artists
endpoint might need more tracks for the same artist
o https://api.spotify.com/v1/search?q=isrc:USEE10001992&amp;type=track
 Get track data by ISRC

Details, Suggestions etc.:
 Choose a Database you are most productive with, MSSQL, MySQL, PostgreSQL
whatever.
 You can choose the way you store the metadata. Normalized table(s), semistructured
tables by storing JSON/Arrays into the DB, etc.
o For the track-&gt;artist relation, it is not necessary to create a many-to-many
relationship. A simple parent (track) -&gt; child (artists) relation is enough
 The design of the REST API is up to you (GET/POST/PUT, query parameters, naming,
hierarchy etc.)
 Of course, we want to see the API in action ;)
o Use the client of your choice (Postman, Curl, Httpie etc.) to demo it
 How would you secure your API endpoints? (in detail)
