# API endpoints
## Channel controls

| Verb | Path | Description |
| --- | --- | --- |
| GET | /channels | Get list of channels |
| GET | /channels/{id} | Get channel info (what's playing, others?) |
| POST | /channels | Manage new channel |
| DELETE | /channels/{id} | Forget channel |
| PUT | /channels/{id}/connect | Connect bot to channel |
| PUT | /channels/{id}/disconnect | Disconnect bot from channel |
| PUT | /channels/{id}/play/{id} | Play song in channel |
| PUT | /channels/{id}/stop | Stop song in channel |
| PUT | /channels/{id}/pause | Pause song in channel |

## MP3s

| Verb | Path | Description |
| --- | --- | --- |
| GET | /songs | List all the MP3s available |
| GET | /songs?sort=["theme","ASC"] | Order all songs by theme |
| GET | /songs?range=[0, 24] | Paginate list of all MP3s |
| GET | /songs?filter={"status":"new"} | List newly added songs |
| GET | /songs/{song_id} | Show song info |
| PUT | /songs/scan | Scan for changes (add/delete songs) |
| PUT | /songs/{song_id} | Edit song information |
| PUT | /songs?filter={"song_id":[1,2,3]} | Edit multiple songs |
| DELETE | /songs/{song_id} | Delete song from filesystem and db |
| DELETE | /songs?filter={"song_id":[1,2,3]} | Delete multiple songs from the filesystem and the db |

## Playlists

| Verb | Path | Description |
| --- | --- | --- |
| GET | /playlists | List all the playlists |
| GET | /playlists?sort | |
| GET | /playlists?range | |
| GET | /playlists?filter | |
| GET | /playlists/{id} | Get playlist information |
| POST | /playlists | Create a new playlist |
| PUT | /playlists/{id} | Edit playlist |
| DELETE | /playlists/{id} | Delete a playlist |
| GET | /playlists/{id}/songs | List all the songs in a playlist |
| GET | /playlists/{id}/songs?sort=["theme","ASC"] | Order all songs in the playlist by theme |
| GET | /playlists/{id}/songs?range=[0, 24] | Paginate list in the playlist |
| GET | /playlists/{id}/songs?filter={"theme":"a"} | List all songs in the playlist matching a theme |
| POST | /playlists/{id}/songs/{id} | Add song to a playlist |
| POST | /playlists/{id}/songs?filter={"song_id":[1,2,3]} | |
| DELETE | /playlists/{id}/songs/{id} | Remove song from a playlist |
| DELETE | /playlists/{id}/songs?filter={"song_id":[1,2,3]} | |

## Tags

| Verb | Path | Description |
| --- | --- | --- |
| GET | /tags | List all tags |
| GET | /tags?sort | |
| GET | /tags?range | |
| GET | /tags?filter | |
| PUT | /tags/{name} | Edit a tag name (Change name for all songs) |
| DELETE | /tags/{name} | Delete a tag (Untag all the songs) |
| GET | /tags/{name}/songs | List all songs with that tag |
| GET | /tags/{name}/songs?sort | |
| GET | /tags/{name}/songs?range | |
| GET | /tags/{name}/songs?filter | |
| POST | /tags/{name}/songs/{song_id} | Tag a song (create tag if it doesn't exist) |
| DELETE | /tags/{name}/songs/{song_id} | Untag a song |
