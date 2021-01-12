const cacheName = 'trail-v3';

self.addEventListener( "install", function( event ){
    event.waitUntil(
        caches.open( "static-cache" )
              .then(function( cache ){
            return cache.addAll([
                'Trees.geojson',
                'manifest.json',
                '/css/',
                '/font/',
                '/img/',
                '/js/',
                '/templates/index.html',
                '/templates/footer.html',
                '/templates/head.html',
                '/templates/map.html',
                '/templates/mobnav.html',
                '/templates/nav.html',
                '/templates/scan.html',
                '/templates/scripts.html',
                '/templates/settings.html',
                '/templates/tour.html',
                "/img/favicon.png"
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
