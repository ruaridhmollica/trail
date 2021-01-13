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

  const CACHE_NAME = 'Trail-cache-v1';
  const cacheurls = [
      '/',
      '/static',
  ]