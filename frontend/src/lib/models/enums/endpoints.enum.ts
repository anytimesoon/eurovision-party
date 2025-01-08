import { env } from '$env/dynamic/public';

export const domainName = env.PUBLIC_DOMAIN_NAME + "/"
const apiBase = env.PUBLIC_GO_HOST + "/"
const restricted = "restricted/"
const goAPI = "api/"
const countryGoURL = apiBase + restricted + goAPI + "country/"
const userGoURL = apiBase + restricted + goAPI + "user/"
const voteGoURL = apiBase + restricted  + goAPI + "vote/"


export const chatEP = env.PUBLIC_CHAT + "/chat/connect/";

export const countryEP = {
    ALL: countryGoURL,
    UPDATE: countryGoURL,
    FIND_ONE: countryGoURL,
    PARTICIPATING: countryGoURL + "participating"
}

export const staticEP = {
    IMG: apiBase + "content/static/",
    AVATAR_IMG: apiBase + "content/user/avatar/",
    CHAT_IMG: apiBase + "content/user/chat/",
    CREATE_CHAT_IMG: apiBase + "content/user/chat"
}

export const userEP = {
    ALL: userGoURL,
    UPDATE: userGoURL,
    UPDATE_IMAGE: userGoURL + "image",
    FIND_ONE: userGoURL, //append with user slug
    REMOVE: userGoURL,
    REGISTERED: userGoURL + "registered"
}

export const voteEP = {
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