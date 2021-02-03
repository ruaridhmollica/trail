console.log("Service Worker Waking UP :)");

self.addEventListener('install', function(event) {
    event.waitUntil(
        caches.open(CACHE_NAME)
        .then(cache => {
            return cache.addAll(cacheurls);
        })
    );
});

self.addEventListener('activate', function(event) {
    console.log("NEW Service Worker Activated :)");
  });

/*self.addEventListener('fetch', event => {
    event.respondWith(
        caches.match(event.request)
        .then(response => {
            return response || fetch(event.request);
        })
    );
});*/

self.addEventListener('fetch', event => {
    console.log('[Trail - ServiceWorker] Fetch event fired.', event.request.url);
    event.respondWith(
        caches.match(event.request).then(function(response) {
            if (response) {
                console.log('[Trail - ServiceWorker] Retrieving from cache...');
                return response;
            }
            console.log('[Trail - ServiceWorker] Retrieving from URL...');
            return fetch(event.request).catch(function (e) {
               //you might want to do more error checking here too,
               //eg, check what e is returning..
               alert('You appear to be offline, please try again when back online');
            });
        })
    );
});

  const CACHE_NAME = 'Trail-cache-v3';
  const cacheurls = [
      './',
      '/static/TreesEdi.geojson',
      '/static/css/bootstrap.css',
      '/static/css/main.css',
      '/static/css/mdb.css',
      '/static/img/favicon.png',
      '/static/img/marker.svg',
      '/static/js/WriteIt.js',	
      '/static/js/jquery.js',
      'manifest.webmanifest',
      'sw.js',
      '/static/facts.json',
      '/static/templates/footer.html',
      '/static/templates/head.html',
      '/static/templates/index.html',
      '/static/templates/map.html',
      '/static/templates/mobnav.html',
      '/static/templates/nav.html',
      '/static/templates/scan.html',
      '/static/templates/scripts.html',
      '/static/templates/settings.html',
      '/static/templates/tour.html',
      '/static/font/coolvetica_rg-webfont.woff2',
      '/static/font/coolvetica_rg-webfont.woff',
      '/static/fontawesome/css/all.min.css',
      '/static/css/bootstrap.min.css',
      '/static/js/jquery.min.js',
      '/static/js/popper.min.js',
      '/static/js/bootstrap.min.js',
      '/static/js/mdb.min.js',
      '/static/js/WriteIt.min.js',
      '/static/js/html5-qrcode.min.js',
      '/static/img/icon-144.png',
      '/static/fontawesome/webfonts/fa-solid-900.woff2',
      '/static/templates/infomodal.html',
      '/static/TreesHWU.geojson',
      '/static/img/tree1.svg',
      '/static/img/tree2.svg',
      '/static/img/tree3.svg',
      '/static/img/tree4.svg',
      '/static/img/tree5.svg',
      '/static/img/tree6.svg',
      '/static/img/tree7.svg',
      '/static/img/tree8.svg',
      '/static/img/tree9.svg',
      '/static/img/tree10.svg',
      '/static/img/tree11.svg',
      '/static/img/tree12.svg',
      '/static/img/tree13.svg',
      '/static/img/tree14.svg',
      '/static/img/tree15.svg',
      '/static/img/tree16.svg',
      '/static/img/tree17.svg',
      '/static/img/tree18.svg',
      '/static/img/tree19.svg',
      '/static/img/tree20.svg',
      '/static/img/tree21.svg',
      '/static/img/tree22.svg',
      '/static/img/tree23.svg',
      '/static/img/tree24.svg',
      '/static/img/tree25.svg',
      '/static/img/tree26.svg',
      '/static/img/tree27.svg',
      '/static/img/tree28.svg',
      '/static/img/tree29.svg',
      '/static/img/tree30.svg',
      '/static/img/tree31.svg',
      '/static/img/tree32.svg',
      '/static/img/tree33.svg',
      '/static/img/tree34.svg',
      '/static/img/tree35.svg',
      '/static/img/tree36.svg',
      '/static/img/tree37.svg',
      '/static/img/tree38.svg',
      '/static/img/tree39.svg',
      '/static/img/tree40.svg',
      '/static/img/tree41.svg',
      '/static/img/tree42.svg',
      '/static/img/tree43.svg',
      '/static/img/tree44.svg',
      '/static/img/tree45.svg',
      '/static/img/tree46.svg',
      '/static/img/tree47.svg'
  ]