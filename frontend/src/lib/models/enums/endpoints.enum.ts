import { env } from '$env/dynamic/public';

export const domainName = env.PUBLIC_DOMAIN_NAME + "/"
const apiBase = env.PUBLIC_GO_HOST + "/"
const restricted = "restricted/"
const goAPI = "api/"
const countryURL = apiBase + restricted + goAPI + "country/"
const userURL = apiBase + restricted + goAPI + "user/"
const voteURL = apiBase + restricted  + goAPI + "vote/"
const commentURL = apiBase + restricted  + goAPI + "comment/"


export const chatEP = env.PUBLIC_CHAT + "/chat/connect/";

export const countryEP = {
    ALL: countryURL,
    UPDATE: countryURL,
    FIND_ONE: countryURL,
    PARTICIPATING: countryURL + "participating"
}

export const staticEP = {
    IMG: apiBase + "content/static/",
    AVATAR_IMG: apiBase + "content/user/avatar/",
    CHAT_IMG: apiBase + "content/user/chat/",
    CREATE_CHAT_IMG: apiBase + "content/user/chat"
}

export const userEP = {
    ALL: userURL,
    UPDATE: userURL,
    UPDATE_IMAGE: userURL + "image",
    FIND_ONE: userURL, //append with user slug
    BAN: userURL + "ban/",
    REGISTERED: userURL + "registered/"
}

export const voteEP = {
    RESULTS: voteURL + "results/",
    CREATE: voteURL,
    UPDATE: voteURL,
    BY_COUNTRY_AND_USER: voteURL + "countryanduser/" //append with country slug
}

export const authEP = {
    LOGIN: apiBase + "api/login",
    SVELTE_LOGIN: domainName + "login/",
    REGISTER: userURL + "register"
}

export const commentEP = {
    REACTIONS: commentURL + "reactions",
}