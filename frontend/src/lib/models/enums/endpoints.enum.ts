import { env } from '$env/dynamic/public';

const apiBase = env.PUBLIC_GO_HOST + "/"
const feBase = "/"
const restricted = "restricted/"
const svelteAPI = "endpoints/"
const goAPI = "api/"
const countrySvelteURL = feBase + svelteAPI + "country/"
const userSvelteURL = feBase + svelteAPI + "user/"
const voteSvelteURL = feBase + svelteAPI + "vote/"
const countryGoURL = apiBase + restricted + goAPI + "country/"
const userGoURL = apiBase + restricted + goAPI + "user/"
const voteGoURL = apiBase + restricted  + goAPI + "vote/"



export const chatEP = (env.PUBLIC_CHAT ? env.PUBLIC_CHAT : "wss://eurovision-dev.delartigue.xyz") + "/restricted/chat/connect";

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

export const staticEP = {
    IMG: apiBase
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
    SVELTE_LOGIN: feBase + "login",
    REGISTER: userGoURL + "register"
}