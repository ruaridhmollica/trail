![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ruaridhmollica/Trail?label=Powered%20by%20Go&logo=Go)

# Trail - An Innovative Way To Learn About Trees on the HWU Campus
A Progressive Web App (PWA) that allows a user to get information about the trees on the Heriot-Watt using QR, Spatial and Speech Synth technologies.
* Work In Progress Deployed at: https://www.thetrailapp.com/
- This application is created by Ruaridh Mollica, the creative direction of this application idea is owned by Ruaridh Mollica. As per section 2 of Heriot-Watts Intelectual Property Policy which states: 
> Any Intellectual Property created by a Student during the course of his/her study shall
remain the property of the Student except where the provisions of Clause Three
herein apply. In particular first ownership of any copyright created rests with the
Student in accordance with the Copyright, Designs and Patents Act 1988 except
where the provisions of Clause Three herein apply.


## Stack

### Hosting
Trail is hosted on [Heroku](https://www.heroku.com/) and has a full valid SSH certification as well as a custom domain.
- Heroku also hosts the applications database.
- Heroku builds and deploys the web application automatically.

### Database
Trail makes use of a [PostgreSQL](https://www.postgresql.org/) database with the [postGIS](https://postgis.net/) extension installed for the ability to query and store geospatial data.

### Backend
This web application is powered by Golang v1.13.1 and handles the following:
* File serving
* Page routing, made easier by the use of the [Gin-Gonic/gin framework](https://github.com/gin-gonic/gin)
* Database connection and querying (using [Go-SQL-driver](https://github.com/go-sql-driver/mysql)
* Outlining package details and dependencies to Heroku for hosting

### Front-end
To make this web app usable and pretty the following technologies are used:
* HTML
* CSS
* JavaScript (to provide more complex functionality and handle API calls)
* [JQuery](https://jquery.com/)
* [Material Design Bootstrap](https://mdbootstrap.com/) (to make everything that little bit easier)

### APIs
* [Google Maps Javascript API](https://developers.google.com/maps/documentation/javascript/overview) is used to generate custom, cross-compatible, interactive maps.

### JavaScript Libraries
* [html5-qrcode](https://blog.minhazav.dev/HTML5-QR-Code-scanning-launched-v1.0.1/#how-to-use) by Minhaz - this library was crucial in development and allowed for a QR code scanner to be embedded into the web application.
* [WriteIT.js](https://khushit-shah.github.io/WriteIt.js/) by khushit-shah - used for the homepage greeting.
* Google Maps' [Marker Clusterer Library](https://googlemaps.github.io/v3-utility-library/classes/_google_markerclustererplus.markerclusterer.html)
* [SpeechSynthesis Web Speech API](https://developer.mozilla.org/en-US/docs/Web/API/SpeechSynthesis) - used for text to speech synthesis
* [User Agent Parser](https://github.com/faisalman/ua-parser-js) by faisalman on GitHub - used to get the users operating system to prevent vibration API functions calling when a user is running web app on an IOS or Mac OS device. (vibration API. is unsupported by IOS)
