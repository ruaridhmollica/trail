const cacheName = 'trail-v2';
const staticAssets = [
   './',
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
   '/templates/tour.html'
];

self.addEventListener('install', async e => {
    const cache = await caches.open(cacheName);
    await cache.addAll(staticAssets);
    return self.skipWaiting();
});

self.addEventListener('active', e => {
    self.ClientRectList.claim();
});


