import type {PageLoad} from ".$/types";

// export const prerender = 'auto'

export const load:PageLoad =  ( ({ params, cookies }) => {
    const hasLoggedIn:boolean = cookies.get("visited") || false

    if (!hasLoggedIn) {
        cookies.set('visited', 'true', { path: '/' })
    }

    return {
        userId: params.userId,
        token: params.token,
        hasLoggedIn
    }
})