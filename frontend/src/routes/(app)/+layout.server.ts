import {redirect} from "@sveltejs/kit";

export const prerender = true

export const load = (async ({cookies}) => {
    if(cookies.get("session") == null) {
        throw redirect(303, "/login")
    }
})