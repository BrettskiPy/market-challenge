<p align="center">
  <img width="200" src="https://media.istockphoto.com/photos/colorful-vegetables-and-fruits-vegan-food-in-rainbow-colors-picture-id1284690585?s=612x612" alt="produce logo">
</p>
<h1 align="center">The Supermarket API</h1>


<h2>What is this?</h2>

<p> The Supermarket API was designed for grocery store employees to manage and retrieve 
produce item data. This is achieved through the use of a <a href="https://www.redhat.com/en/topics/api/what-is-a-rest-api">RESTful API.</a>
The specifications for this project are found <a href="https://gist.github.com/tomtoday/ef8e8c01582036ce3edc42fee44a3691">here.</a></p>


<h2>What does the API do?</h2>

<h3>Item retrievals via GET</h3>
<p>Allows a user to retrieve a single or multiple produce items.</p>
<h3>Item additions via POST</h3>
<p>Allows a user to add a single or multiple produce items via JSON.</p>
<h3>Item deletions via DELETE</h3>
<p>Allows a user to delete a single produce item.</p>

<h2>What data fields define a produce item?</h2>

<h3>Produce item field descriptions:</h3>
<ul>
  <li><b>name</b> - A required field that represents the official name of a produce item.</li>
  <li><b>code</b> - A required sixteen character alphanumeric field with three hyphens between each set of 4 characters. This sequence represents a produce item's unique identification. </li>
  <li><b>unit_price</b> - A required field that represents the total price divided by the total units.
<a href="https://www.onlinemathlearning.com/unit-price.html">click here</a> for more information about what a unit price is.</li>
</ul>

<h3>Example of a produce item's fields and values in JSON:</h3>
{"name": "Apple", "code": "Z931-44D4-A12T-1224", "unit_price": 5.33}

<h2>How do I use the market APIs?</h2>

Below are the descriptions of each API endpoint and functionality. It is assumed you have a working connection 
to the market website or targeting the endpoints via an api tool such as <a href="https://www.postman.com/">postman.</a>
<h3>GET</h3>
<table>
  <tr>
    <td><b>Endpoint</b></td>
    <td><b>Description</b></td>
    <td><b>Sample Output</b></td>
  </tr>
  <tr>
    <td>/produce/{name}</td>
    <td>Retrieval of a single produce item and its associated data</td>
    <td>{
    "name": "Lettuce",
    "code": "A12T-4GH7-QPL9-3N4M",
    "unit_price": 3.46
}</td>
  </tr>
  <tr>
    <td>/produce</td>
    <td>Retrieval of all the produce items and their associated data</td>
    <td>[
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
    },
    {
        "name": "Gala Apple",
        "code": "TQ4C-VV6T-75ZX-1RMR",
        "unit_price": 3.59
    }
]</td>
  </tr>
</table>


<h3>POST</h3>
<table>
  <tr>
    <td><b>Endpoint</b></td>
    <td><b>Description</b></td>
  </tr>
  <tr>
    <td>/produce/{name or names}</td>
    <td>Adds a single or multiple produce item(s) and their associated data via JSON.</td>
  </tr>
</table>

<h3>DELETE</h3>
<table>
  <tr>
    <td><b>Endpoint</b></td>
    <td><b>Description</b></td>
  </tr>
  <tr>
    <td>/produce/{name}</td>
    <td>Deletes a single produce item and its associated data.</td>
  </tr>
</table>








