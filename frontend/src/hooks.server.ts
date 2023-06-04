import type { Handle, HandleFetch } from '@sveltejs/kit';

export const handle: Handle = async ( { event, resolve}) => {
    console.log("handle: " + event.request.url)
    event.request.headers.set("accept", "application/json, image/*")
    return resolve(event);
};

export const handleFetch:HandleFetch = async ({ event, request, fetch }) => {
    console.log("handle fetch: " + request.url)
    event.request.headers.set("accept", "application/json, image/*")
    request.headers.set('cookie', event.request.headers.get('cookie'));
    return fetch(request);
}