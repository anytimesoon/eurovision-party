import { env } from '$env/dynamic/public';

export const domainName = env.PUBLIC_DOMAIN_NAME + "/"
const apiBase = env.PUBLIC_GO_HOST + "/"
const feBase = "/"
const restricted = "restricted/"
const svelteAPI = "endpoints/"
const goAPI = "api/"
const countrySvelteURL = feBase + svelteAPI + "country/"
const userSvelteURL = feBase + svelteAPI + "user/"
const voteSvelteURL = feBase + svelteAPI + "vote/"
const staticSvelteURL = feBase + svelteAPI + "content/"
const countryGoURL = apiBase + restricted + goAPI + "country/"
const userGoURL = apiBase + restricted + goAPI + "user/"
const voteGoURL = apiBase + restricted  + goAPI + "vote/"


export const chatEP = env.PUBLIC_CHAT + "/chat/connect/";

export const countrySvelteEP = {
    ALL: countrySvelteURL,
    UPDATE: countrySvelteURL,
    FIND_ONE: countrySvelteURL,
    PARTICIPATING: countrySvelteURL + "participating"
}

export const countryGoEP = {
    ALL: countryGoURL,
    UPDATE: countryGoURL,
    FIND_ONE: countryGoURL,
    PARTICIPATING: countryGoURL + "participating"
}

export const staticSvelteEP = {
    AVATAR_IMG: staticSvelteURL + "user/avatar/",
    CHAT_IMG: staticSvelteURL + "user/chat/",
    CREATE_CHAT_IMG: staticSvelteURL + "chat/"
}

export const staticGoEP = {
    IMG: apiBase + "content/static/",
    AVATAR_IMG: apiBase + "content/user/avatar/",
    CHAT_IMG: apiBase + "content/user/chat/",
    CREATE_CHAT_IMG: apiBase + "content/user/chat"
}

export const userSvelteEP = {
    ALL: userSvelteURL,
    UPDATE: userSvelteURL,
    FIND_ONE: userSvelteURL,
    REMOVE: userSvelteURL,
    REGISTERED: userSvelteURL + "registered"
}

export const userGoEP = {
    ALL: userGoURL,
    UPDATE: userGoURL,
    UPDATE_IMAGE: userGoURL + "image",
    FIND_ONE: userGoURL, //append with user slug
    REMOVE: userGoURL,
    REGISTERED: userGoURL + "registered"
}

export const voteSvelteEP = {
    RESULTS: voteSvelteURL,
    UPDATE: voteSvelteURL,
    BY_COUNTRY_AND_USER: voteSvelteURL + "countryanduser/" //append with country slug
}

export const voteGoEP = {
    RESULTS: voteGoURL + "results/",
    CREATE: voteGoURL,
    UPDATE: voteGoURL,
    BY_COUNTRY_AND_USER: voteGoURL + "countryanduser/" //append with country slug
}

export const authEP = {
    LOGIN: apiBase + "api/login",
    SVELTE_LOGIN: domainName + "login/",
    REGISTER: userGoURL + "register"
}