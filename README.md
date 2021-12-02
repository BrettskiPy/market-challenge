<p align="center">
  <img width="200" src="https://media.istockphoto.com/photos/colorful-vegetables-and-fruits-vegan-food-in-rainbow-colors-picture-id1284690585?s=612x612" alt="produce logo">
</p>
<h1 align="center">The Supermarket API</h1>

<p> The Supermarket API was designed for grocery store employees to manage and retrieve 
produce item data. This functionality is achieved through the use of a <a href="https://www.redhat.com/en/topics/api/what-is-a-rest-api">REST API.</a>
The specifications for this project are found <a href="https://gist.github.com/tomtoday/ef8e8c01582036ce3edc42fee44a3691">here.</a></p>

<h2>Produce item data fields</h2>
<table>
  <tr>
    <td><b>Data Field</b></td>
    <td><b>Description</b></td>
    <td><b>Data Type</b></td>
    <td><b>Data Sample</b></td>
  </tr>
  <tr>
    <td>name</td>
    <td>A required case insensitive field that represents the official name of a produce item.</td>
    <td>string</td>
    <td>Lettuce</td>
  </tr>
  <tr>
    <td>code</td>
    <td> A required case insensitive, sixteen character, alphanumeric field, with three hyphens between each set of 4 characters. 
This sequence represents a produce item's unique identification.</td>
    <td>string</td>
    <td>A12T-4GH7-QPL9-3N4M</td>
  </tr>
  <tr>
    <td>unit_price</td>
    <td>A required field that represents the total price divided by the total units.</td>
    <td>float64</td>
    <td>3.46</td>
  </tr>
</table>

<h4>Example produce item in JSON:</h4>
<pre><code>{
    "name": "Lettuce",
    "code": "A12T-4GH7-QPL9-3N4M",
    "unit_price": 3.46
}</code></pre>

<h2>How do I use the market APIs?</h2>
<p>API functionality is achieved by using the endpoints below. It is assumed you have a working connection 
to the market website or you are targeting the endpoints via an API tool such as <a href="https://www.postman.com/">postman.</a></p>

<h4>GET</h4>
<table>
  <tr>
    <td><b>Endpoint</b></td>
    <td><b>Description</b></td>
    <td><b>Sample JSON Output</b></td>
  </tr>
  <tr></tr>
  <tr>
    <td>/produce/{name}</td>
    <td>Retrieval of a single produce item and its associated data</td>
    <td><pre><code>{
    "name": "Lettuce",
    "code": "A12T-4GH7-QPL9-3N4M",
    "unit_price": 3.46
}</code></pre></td>
  </tr>
  <tr></tr>
  <tr>
    <td>/produce</td>
    <td>Retrieval of all the produce items and its associated data</td>
    <td><pre><code>[
    {
        "name": "Lettuce",
        "code": "A12T-4GH7-QPL9-3N4M",
        "unit_price": 3.46
    },
    {
        "name": "Peach",
        "code": "E5T6-9UI3-TH15-QR88",
        "unit_price": 2.99
    },
    {
        "name": "Green Pepper",
        "code": "YRT6-72AS-K736-L4AR",
        "unit_price": 0.79
    }
]</code></pre></td>
  </tr>
</table>

<h4>POST</h4>
<table>
  <tr>
    <td><b>Endpoint</b></td>
    <td><b>Description</b></td>
    <td><b>Sample JSON Input</b></td>
  </tr>
  <tr></tr>
  <tr>
    <td>/produce</td>
    <td>Adds a single produce item and its associated data.</td>
    <td><pre><code>{
    "name": "Gala Apple",
    "code": "TQ4C-VV6T-75ZX-1RMR",
    "unit_price": 3.59
}</code></pre></td>
  </tr>
  <tr></tr>
<tr>
    <td>/produce</td>
    <td>Adds multiple produce items and its associated data.</td>
    <td><pre><code>[
    {
        "name": "Lettuce",
        "code": "A12T-4GH7-QPL9-3N4M",
        "unit_price": 3.46
    },
    {
        "name": "Peach",
        "code": "E5T6-9UI3-TH15-QR88",
        "unit_price": 2.99
    }
]</code></pre></td>
  </tr>
</table>

<h4>DELETE</h4>
<table>
  <tr></tr>
  <tr>
    <td><b>Endpoint</b></td>
    <td><b>Description</b></td>
  </tr>
  <tr>
    <td>/produce/{name}</td>
    <td>Deletes a single produce item and its associated data.</td>
  </tr>
</table>





