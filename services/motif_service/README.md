# motif-service

This service provides endpoints for retrieving the bounding box of the motif i.e. the main subject of a given image.

Please refer to the [proto file](../../shared/proto/services/motif_service/v1/motif_service.proto) for more information on the service's API.

## Examples

<table>
  <tr>
    <th>Input</th>
    <th>Debug Image</th>
    <th>Output</th>
  </tr>
  <tr>
    <td><img src="../../test/images/object.jpg" alt="Input image object"></td>
    <td><img src="./docs/img/obj_debug.jpg" alt="Output debug image object"></td>
    <td>
  
  ```json
  {
    "bBox": {
      "x": 428,
      "y": 195,
      "width": 286,
      "height": 463
    }
  }
  ```

  </td>
  </tr>
  <tr>
    <td><img src="../../test/images/human.jpg" alt="Input image human"></td>
    <td><img src="./docs/img/human_debug.jpg" alt="Output image"></td>
       <td>
  
  ```json
  {
    "bBox": {
      "x": 208,
      "y": 176,
      "width": 537,
      "height": 371
    }
  }
  ```

  </td>
  </tr>
  <tr>
    <td><img src="../../test/images/animal.jpg" alt="Input image animal"></td>
    <td><img src="./docs/img/animal_debug.jpg" alt="Output image"></td>
        <td>
  
  ```json
  {
    "bBox": {
      "x": 219,
      "y": 65,
      "width": 680,
      "height": 655
    }
  }
  ```

  </td>
  </tr>
</table>

## Credits

Example images used in the tests are from [Unsplash](https://unsplash.com/):

- [Object example](./docs/examples/obj.jpg) by [Atul Vinayak](https://unsplash.com/@atulvi?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash)
- [Person example](./docs/examples/person.jpg) by [Isabela Drasovean](https://unsplash.com/@isabeladrasovean?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash)
- [Animal example](./docs/examples/animal.jpg) by [Kari Shea](https://unsplash.com/@karishea?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash)
