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

self.addEventListener('fetch', event => {
    event.respondWith(
        caches.match(event.request)
        .then(response => {
            return response || fetch(event.request);
        })
    );
});

let deferredPrompt;

window.addEventListener('beforeinstallprompt', (e) => {
    e.preventDefault();
    deferredPrompt = e;
    btnAdd.style.display = 'block';
});

btnAdd.addEventListener('click', (e)=> {
    deferredPrompt.prompt();
    deferredPrompt.userChoice.then((choiceResult) => {
        if (choiceResult.outcome === 'accepted'){
            console.log('User accepted A2HS prompt');
        }
        deferredPrompt = null;
    });
});

window.addEventListener('appinstalled', (evt) =>{
    app.logEvent('a2hs', 'installed');
});

  const CACHE_NAME = 'Trail-cache-v1';
  const cacheurls = [
      '/static/Trees.geojson',
      '/static/css/bootstrap.css',
      '/static/css/main.css',
      '/static/css/mdb.css',
      '/static/img/favicon.png',
      '/static/img/marker.svg',
      '/static/img/tree1.svg',
      '/static/js/WriteIt.js',	
      '/static/js/jquery.js',
      'manifest.json',
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
      '/static/font/coolvetica_rg-webfont.woff'	
  ]