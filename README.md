<p align="center">
  <img width="200" src="https://media.istockphoto.com/photos/colorful-vegetables-and-fruits-vegan-food-in-rainbow-colors-picture-id1284690585?s=612x612" alt="produce logo">
</p>
<h1 align="center">The Supermarket API</h1>


<h2>What is it?</h2>
___
<p> The Supermarket API was designed for grocery store employees to manage and retrieve 
produce item data through the use of a RESTful API.</p>


<h2>What does the API do?</h2>
___
<h3>Item Retrieval Via GET</h3>
<p>Allows the user retrieve single or multiple produce items.</p>
<h3>Item Addition Via POST</h3>
<p>Allows the user to add single or multiple produce items via JSON.</p>
<h3>Item Deletion Via DELETE</h3>
<p>Allows the user to delete single produce items.</p>

<h2>What data fields are contained within a produce item?</h2>
___
<h3>Produce item field descriptions:</h3>
<ul>
  <li><b>name</b> - A required field that represents the official name of a produce item.</li>
  <li><b>code</b> - A required sixteen character alphanumeric field that is separated by three hyphens. This code represents a produce item's unique identification. </li>
  <li><b>unit_price</b> - A required field that represents the total price divided by the total units <a href="https://www.onlinemathlearning.com/unit-price.html">click here</a> for more information </li>
</ul>

<h3>Example of a produce item in JSON:</h3>
{"name": "Apple", "code": "Z931-44D4-A12T-1224", "unit_price": 5.33}

<h2>How?</h2>
___
Below are the commands (in bold) followed by descriptions of each API command functionality. It is assumed you have a working connection 
to the market website or targeting the endpoints via an api tool such as <a href="https://www.postman.com/">postman.</a>
<h3>GET</h3>
<ul>

  <li><b>/produce/{name}</b> Retrieval of a single produce item and its associated data </li>
  <li><b>/produce</b> Retrieval of all the produce items and their associated data</li>
</ul>

<h3>POST</h3>
<ul>
  <li>Capable of adding a single produce item and its associated data via JSON</li>
  <li>Capable of adding multiple produce items and their associated data via JSON</li>
</ul>







