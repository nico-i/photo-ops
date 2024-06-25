# hashtag-service

This service provides endpoints for generating a set of max. 30 hashtags based on a given image caption.

Please refer to the [proto file](../../shared/proto/services/hashtag_service/v1/hashtag_service.proto) for more information on the service's API.

## Requirements

- [Python 3.11.7](https://www.python.org/downloads/release/python-3117/)
  - I recommend to using [pyenv](https://github.com/pyenv/pyenv) to manage python versions
- [gRPCurl](https://github.com/fullstorydev/grpcurl)
- [Make](https://www.gnu.org/software/make/) (optional)
  - should be pre-installed on most Unix-like systems

## Setup

1. Clone [this repository](https://github.com/nico-i/photo-ops/tree/main) and navigate to the root directory.
2. Navigate to the `services/hashtag_service` directory.
3. Set up a python virtual environment with the pre-configured make command.
4. Install the required dependencies to the virtual environment.
5. Start the service. (optionally edit [the make file](./makefile) to change the port or disable debug mode)

The above steps can be executed with the following commands:

```bash
git clone https://github.com/nico-i/photo-ops.git
cd photo-ops
cd services/hashtag_service
make venv
make install
make dev
```

## Usage

If you have followed the setup instructions above, you can utilize the configured make commands to execute the example requests.

Please refer to [the make file](./makefile) or [the table below](#examples) for the available make commands.

## Examples

<table>
  <tr>
    <th>Make job</th>
    <th>RPC</th>
    <th>Input</th>
    <th>Output</th>
    <th>Duration</th>
  </tr>
  <tr>
  <td>
  
  ```bash
  make req_animal
  ```

  </td>
  <td>

  `get_hashtags`

  </td>

  <td>

```json
{ "caption": "a cat is lying on a chair" }
```
  </td>
  <td>
  
```json
{
  "hashtagsCsv": "\"Feline,pet,catsofinstagram,catlovers,catlife,chair,home,interior,animal,petsofinstagram,furryfriend,whiskers,claws,catnap,relaxation,comfort,cozy,domestic,petlife,catmom,catdad,catlady,catman,animallovers,petlovers,whimsical,funny,adorable,cute\""
}
```

  </td>
  <td>
    ~20s
  </td>
  </tr>
  <tr>
  <td>
  
  ```bash
  make req_obj
  ```
  
  </td>
    <td>

  `get_hashtags`

  </td>
  <td>

```json
{ "caption": "a sign on a fence" }
```
</td>
<td>

```json
{
  "hashtagsCsv": "\"sign,fence,handwritten,wooden,metal,public,notice,warning,information,display,marker,markerboard,signage,advertising,communication,structure,property,boundary,markerpost,landmark,markerboard,markerpost,markerboard,markerpost,markerboard,markerpost,markerboard,markerpost,markerboard,markerpost,markerboard\""
}
```

</td>
    <td>~12s</td>
  <tr>
  <td>
  
  ```bash
  make req_human
  ```
  
  </td>
    <td>

  `get_hashtags`

  </td>
<td>

```json
{ "caption": "a girl with a flower in her hair" }
```

</td>
<td>

```json
{
  "hashtagsCsv": "flowers, beauty, hair, girl, style, fashion, floral, bloom, blossom, petals, lovely, lovelygirl, flowerpower, hairaccessories, beautyinspo, girlwithflowers, flowercrowns, hairgoals, fashionista, beautylover, lovelythings, flowerlove, hairdo, girlygirl, flowerpower, beautyful, fashionstyle, hairinspo, lovelyhair, flowercrown, beautyinspiration, girlstyle"
}
```

</td>
<td>~13s</td>
</table>