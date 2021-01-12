const cacheName = 'trail-v1';
const staticAssets = [
    '/main.go',
   '/templates/',
   '/static/',
   '/bin/'
];

self.addEventListener('install', async e => {
    const cache = await caches.open(cacheName);
    await cache.addAll(staticAssets);
    return self.skipWaiting();
});



