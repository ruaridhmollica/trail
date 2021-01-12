self.addEventListener( "install", function( event ){
    event.waitUntil(
        caches.open( "static-cache" )
              .then(function( cache ){
            return cache.addAll([
                'Trees.geojson',
                'manifest.json',
                '/static/css/main.css',
                '/static/css/bootstrap.css',
                '/static/css/mdb.css',
                '/static/img/marker.svg',
                '/static/img/tree1.svg',
                '/static/js/WriteIt.js',
                '/static/js/jquery.js',
                '/static/templates/index.html',
                '/static/templates/footer.html',
                '/static/templates/head.html',
                '/static/templates/map.html',
                '/static/templates/mobnav.html',
                '/static/templates/nav.html',
                '/static/templates/scan.html',
                '/static/templates/scripts.html',
                '/static/templates/settings.html',
                '/static/templates/tour.html',
                "/static/img/favicon.png"
            ]);
        })
    );
});

self.addEventListener( "fetch", event => {
    const request = event.request,
                    url = request.url;
    // If we are requesting an HTML page.
    if ( request.headers.get("Accept").includes("text/html") ) {
        event.respondWith(
            // Check the cache first to see if the asset exists, and if it does, return the cached asset.
            caches.match( request ).then( cached_result => {
                if ( cached_result ) {
                    return cached_result;
                }
                // If the asset is not in the cache, fallback to a network request for the asset, and proceed to cache the result.
                return fetch( request ).then( response => {
                    const copy = response.clone();
                    // Wait until the response we received is added to the cache.
                    event.waitUntil(
                        caches.open( "pages" )
                              .then( cache => {
                            return cache.put( request, response );
                        })
                    );
                    return response;
                })
                // If the network is unavailable to make a request, pull the offline page out of the cache.
                .catch(() => caches.match( "/templates/settings.html" ));
            })
        ); // end respondWith
    } // end if HTML
});
